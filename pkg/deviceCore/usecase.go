package deviceCore

import (
	"fmt"
	"github.com/go-kit/kit/log/level"
)

/*
  	deviceCore/usecase.go:

	- Implements CoreService
	- Provides interface to outside UIs
*/

// The active Service
var (
	em *coreService
)

/* Service implementation
 *
 * CoreService for UI or externals
 */
func (em *coreService) AllNetworks() SiteNetworks {
	level.Debug(em.logger).Log("method", "AllNetworks() called")
	return siteNetworks
}
func (em *coreService) PrivateSiteNetworks() *SiteNetworks {
	level.Debug(em.logger).Log("method", "PrivateSiteNetworks() called")
	return &siteNetworks
}

func (em *coreService) NetworkByName(networkName string) Network {
	level.Debug(em.logger).Log("method", "NetworkByName() called")
	network := siteNetworks.DeviceNetworks[networkName]
	return network
}

// DeviceByName devices are mapped by deviceName
func (em *coreService) DeviceByName(deviceName, networkName string) (Device, error) {
	level.Debug(em.logger).Log("method", "DeviceByNameFromNetwork() called")
	var err error
	var device = Device{}
	var found  = false
	device, found = siteNetworks.DeviceNetworks[networkName].Devices[deviceName]
	if !found {
		err = fmt.Errorf("device with name {%s} was  not found on {%s} network", deviceName, networkName)
		level.Error(em.logger).Log("error", err.Error())
	}
	return device, err
}
func (em *coreService) DeviceByID(deviceID string, networkName string) (Device, error) {
	var err error
	level.Debug(em.logger).Log("method", "DeviceByIDFromNetwork() called")
	var device Device = Device{}
	
	for _, dev := range siteNetworks.DeviceNetworks[networkName].Devices {
		if string(dev.ID) == deviceID {
			device = dev
			break
		}
	}
	if device.ElementType != CoreTypeDevice {
		err = fmt.Errorf("device with id {%s} was  not found on {%s} network", deviceID, networkName)
		level.Error(em.logger).Log("error", err.Error())
	}
	return device, err
}
func (em *coreService) RemoveDeviceByID(deviceID string, networkName string) error {
	var err error
	level.Debug(em.logger).Log("method", "RemoveDeviceByIDFromNetwork() called")

	var device Device = Device{}
	for _, dev := range siteNetworks.DeviceNetworks[networkName].Devices {
		if string(dev.ID) == deviceID {
			device = dev
			break
		}
	}
	if device.ElementType != CoreTypeDevice {
		err = fmt.Errorf("device with id {%s} was  not found on {%s} network", deviceID, networkName)
		level.Error(em.logger).Log("error", err.Error())
	} else {
		delete(siteNetworks.DeviceNetworks[networkName].Devices, device.Name)

		// Tell the world to delete this device
		dv := DeviceMessage{}
		for _, topic := range TopicsFromDevice(device) {
			dm := DeviceMessage{
				NetworkID: []byte(device.Parent),
				DeviceID: []byte(device.Name),
				HomieType: CoreTypeDeviceDelete,
				TopicS: topic,
				Value: nil,
				Qosb: 1,
				RetainedB: false,
			}
			dv = dm
			em.dsp.PublishToStreamProvider(dm)
			level.Debug(em.logger).Log("publishing to", device.Name, "Topic", topic)
		}
		
		// remove from db
		_ = em.repo.Remove(dv)

		// Update Scheduler
		if em.scp != nil {
			schedule := em.scp.FindScheduleByDeviceID(string(device.ID))
			if schedule.ElementType == CoreTypeSchedule {
				err = em.scp.DeleteSchedule(schedule.ID)
			}
		}
	}
	return err
}

func (em *coreService) PublishNetworkMessage(dm DeviceMessage) {
	em.dsp.PublishToStreamProvider(dm)
}

func (em *coreService) AllSchedules() []Schedule {
	level.Debug(em.logger).Log("method", "AllSchedules() called")
	return em.scp.Schedules()
}
func (em *coreService) CreateSchedule(networkName string, deviceID string, transport OTATransport, firmwareID EID) (string, error) {
	level.Debug(em.logger).Log("method", "CreateSchedule() called", "device", deviceID, "firmware", firmwareID)
	return em.scp.CreateSchedule(networkName, deviceID, transport, firmwareID)
}
func (em *coreService) RemoveSchedule(scheduleID string) error {
	level.Debug(em.logger).Log("method", "RemoveSchedule() called", "schedule", scheduleID)
	return em.scp.DeleteSchedule(scheduleID)
}
func (em *coreService) ScheduleByID(scheduleID string) Schedule {
	level.Debug(em.logger).Log("method", "ScheduleByID() called", "schedule", scheduleID)
	schedule := siteNetworks.Schedules[scheduleID]
	return schedule
}
func (em *coreService) ScheduleByDeviceID(deviceID string) *Schedule {
	level.Debug(em.logger).Log("method", "ScheduleByDeviceID() called")
	return em.scp.FindScheduleByDeviceID(deviceID)
}
func (em *coreService) AllFirmwares() []Firmware {
	level.Debug(em.logger).Log("method", "AllFirmwares() called")
	return em.scp.Firmwares()
}
func (em *coreService) CreateFirmware(srcFile, dstFile string) (EID, error) {
	level.Debug(em.logger).Log("method", "CreateFirmware() called")
	return em.scp.CreateFirmware(srcFile, dstFile)
}
func (em *coreService) RemoveFirmwareByID(firmwareEID EID) {
	level.Debug(em.logger).Log("method", "RemoveFirmwareByEID() called")
	_ = em.scp.DeleteFirmware(firmwareEID)
}
func (em *coreService) FirmwareByName(firmwareName string) (Firmware, error) {
	level.Debug(em.logger).Log("method", "FirmwareByName() called")
	var firmware Firmware = Firmware{}
	var err error
	for _, fw := range siteNetworks.Firmwares {
		if fw.Name == firmwareName {
			firmware = fw
			break
		}
	}
	if (firmware.ElementType != CoreTypeFirmware) {
		err = fmt.Errorf("firmware with name {%s} was  not found", firmwareName)
		level.Error(em.logger).Log("error", err.Error())
	}
	return firmware, err
}
func (em *coreService) FirmwareByID(firmwareID EID) (Firmware, error) {
	level.Debug(em.logger).Log("method", "FirmwareByID() called")
	return em.scp.GetFirmware(firmwareID)
}

func (em *coreService) AllBroadcasts() []Broadcast {
	level.Debug(em.logger).Log("method", "AllBroadcasts() called")
	return siteNetworks.Broadcasts
}
func (em *coreService) RemoveBroadcastByID(broadcastID string) {
	level.Debug(em.logger).Log("method", "RemoveBroadcastByID() called")
	var index int
	var broadcast Broadcast = Broadcast{}
	for idx, bc := range siteNetworks.Broadcasts {
		if bc.ID == broadcastID {
			broadcast = bc
			index = idx
			break
		}
	}

	if broadcast.ElementType == CoreTypeBroadcast {

		// maybe send null messages to remove from network
		//  NewBroadcast(string(dm.NetworkID), string(dm.AttributeID), string(dm.PropertyID), string(dm.Value))
		// NewBroadcast(parent, topic, level, value string)
		var topic string
		if broadcast.Level == "" {
			topic = fmt.Sprintf("%s/$broadcast/%s", broadcast.Parent, broadcast.Topic)
		} else {
			topic = fmt.Sprintf("%s/$broadcast/%s/%s", broadcast.Parent, broadcast.Topic, broadcast.Level)
		}
		dm := DeviceMessage{
			NetworkID: []byte(broadcast.Parent),
			DeviceID: []byte(broadcast.Level),
			HomieType: CoreTypeDeviceDelete,
			TopicS: topic,
			Value: []byte(broadcast.Value),
			Qosb: 1,
			RetainedB: false,
		}
		em.dsp.PublishToStreamProvider(dm)

		siteNetworks.Broadcasts = append(siteNetworks.Broadcasts[:index], siteNetworks.Broadcasts[index+1:]...) // remove from slice
	}
}
func (em *coreService) BroadcastByID(broadcastID string) (Broadcast, error) {
	level.Debug(em.logger).Log("method", "BroadcastByID() called")
	var broadcast Broadcast = Broadcast{}
	var err error
	for _, bc := range siteNetworks.Broadcasts {
		if bc.ID == broadcastID {
			broadcast = bc
			break
		}
	}
	if (broadcast.ElementType != CoreTypeBroadcast) {
		err = fmt.Errorf("broadcast with id {%s} was  not found", broadcastID)
		level.Error(em.logger).Log("error", err.Error())
	}
	return broadcast, err
}
