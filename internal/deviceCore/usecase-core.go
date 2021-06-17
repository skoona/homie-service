package deviceCore

import (
	"fmt"
	"github.com/go-kit/kit/log/level"
	"os"
)

/*
  	deviceCore/usecase-core.go:

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
	copyOfSiteNetworks := siteNetworks
	return copyOfSiteNetworks
}
func (em *coreService) NetworkByName(networkName string) Network {
	level.Debug(em.logger).Log("method", "NetworkByName() called")
	copyOfNetwork := siteNetworks.DeviceNetworks[networkName]
	return copyOfNetwork
}

// TODO Sort out if devices are mapped by deviceName or by Device.ID
func (em *coreService) DeviceByName(deviceName, networkName string) (Device, error) {
	var err error
	level.Debug(em.logger).Log("method", "DeviceByNameFromNetwork() called")
	copyOfDevice, found := siteNetworks.DeviceNetworks[networkName].Devices[deviceName]
	if !found {
		err = fmt.Errorf("device with name {%s} was  not found on {%s} network", deviceName, networkName)
		level.Error(em.logger).Log("error", err.Error())
	}
	return copyOfDevice, err
}
func (em *coreService) DeviceByID(deviceID string, networkName string) (Device, error) {
	var err error
	level.Debug(em.logger).Log("method", "DeviceByIDFromNetwork() called")
	var ptrToDevice *Device
	var copyOfDevice Device
	for _, device := range siteNetworks.DeviceNetworks[networkName].Devices {
		if device.ID == deviceID {
			ptrToDevice = &device
			break
		}
	}
	if ptrToDevice == nil {
		err = fmt.Errorf("device with id {%s} was  not found on {%s} network", deviceID, networkName)
		level.Error(em.logger).Log("error", err.Error())
	} else {
		copyOfDevice = *ptrToDevice // TODO: hoping this copies vs references
	}
	return copyOfDevice, err
}
func (em *coreService) RemoveDeviceByID(deviceID string, networkName string) error {
	var err error
	level.Debug(em.logger).Log("method", "RemoveDeviceByIDFromNetwork() called")

	var ptrToDevice *Device
	for _, device := range siteNetworks.DeviceNetworks[networkName].Devices {
		if device.ID == deviceID {
			ptrToDevice = &device
			break
		}
	}
	if ptrToDevice == nil {
		err = fmt.Errorf("device with id {%s} was  not found on {%s} network", deviceID, networkName)
		level.Error(em.logger).Log("error", err.Error())
	} else {
		delete(siteNetworks.DeviceNetworks[networkName].Devices, ptrToDevice.Name)

		// Tell the world to delete this device
		dv := DeviceMessage{}
		for _, topic := range TopicsFromDevice(*ptrToDevice) {
			dm := DeviceMessage{
				NetworkID: []byte(ptrToDevice.Parent),
				DeviceID: []byte(ptrToDevice.Name),
				TopicS: topic,
				Value: nil,
				Qosb: 0,
				RetainedB: false,
			}
			dv = dm
			em.dsp.PublishToStreamProvider(dm)
			level.Debug(em.logger).Log("publishing to", ptrToDevice.Name, "Topic", topic)
		}
		
		// remove from db
		em.repo.Remove(dv)

		// Update Scheduler
		if em.scp != nil {
			schedule := em.scp.FindScheduleByDeviceID(ptrToDevice.ID)
			err = em.scp.DeleteSchedule(schedule.ID)
		}
	}
	return err
}

func (em *coreService) PublishNetworkMessage(dm DeviceMessage) {
	em.dsp.PublishToStreamProvider(dm)
}

func (em *coreService) AllSchedules() []Schedule {
	level.Debug(em.logger).Log("method", "AllSchedules() called")
	schedules := make([]Schedule, len(siteNetworks.Schedules))
	for _, schedule := range siteNetworks.Schedules {
		schedules = append(schedules, schedule)
	}
	return schedules
}
func (em *coreService) CreateSchedule(networkName string, deviceID string, transport OTATransport, firmwareID EID) (string, error) {
	level.Debug(em.logger).Log("method", "CreateSchedule() called")
	return em.scp.CreateSchedule(networkName, deviceID, transport, firmwareID)
}
func (em *coreService) RemoveSchedule(scheduleID string) {
	level.Debug(em.logger).Log("method", "RemoveSchedule() called")
	delete(siteNetworks.Schedules, scheduleID)
}
func (em *coreService) ScheduleByID(scheduleID string) Schedule {
	level.Debug(em.logger).Log("method", "ScheduleByID() called")
	schedule := siteNetworks.Schedules[scheduleID]
	return schedule
}
func (em *coreService) ScheduleByDeviceID(deviceID string) Schedule {
	level.Debug(em.logger).Log("method", "ScheduleByDeviceID() called")
	var schedule Schedule
	for _, obj := range siteNetworks.Schedules {
		if obj.DeviceID == deviceID {
			schedule = obj
			break
		}
	}
	return schedule
}
func (em *coreService) AllFirmwares() []Firmware {
	level.Debug(em.logger).Log("method", "AllFirmwares() called")
	firmwares := siteNetworks.Firmwares // presumed to actually copy
	return firmwares
}
func (em *coreService) CreateFirmware(path string) (EID, error) {
	level.Debug(em.logger).Log("method", "CreateFirmware() called")
	return em.scp.CreateFirmware(path)
}
func (em *coreService) RemoveFirmwareByID(firmwareEID EID) {
	level.Debug(em.logger).Log("method", "RemoveFirmwareByEID() called")
	var firmware Firmware
	var index int
	for idx, fw := range siteNetworks.Firmwares {
		if fw.ID == firmwareEID {
			firmware = fw
			index = idx
			break
		}
	}
	if (firmware != Firmware{}) {
		siteNetworks.Firmwares = append(siteNetworks.Firmwares[:index], siteNetworks.Firmwares[index+1:]...) // remove from slice
		if _, err := os.Stat(firmware.Path); err == nil {
			err = os.Remove(firmware.Path)
			if err != nil {
				err = fmt.Errorf("firmware with name {%s} was  not found to remove(): error=%s", firmware.Name, err.Error())
				level.Error(em.logger).Log("error", err.Error())
			}

		} else if os.IsNotExist(err) {
			err = fmt.Errorf("firmware with name {%s} was  not found to remove(): error=%s", firmware.Name, err.Error())
			level.Error(em.logger).Log("error", err.Error())
		}
	}
}
func (em *coreService) FirmwareByName(firmwareName string) (Firmware, error) {
	level.Debug(em.logger).Log("method", "FirmwareByName() called")
	var firmware Firmware
	var err error
	for _, fw := range siteNetworks.Firmwares {
		if fw.Name == firmwareName {
			firmware = fw
			break
		}
	}
	if (firmware == Firmware{}) {
		err = fmt.Errorf("firmware with name {%s} was  not found", firmwareName)
		level.Error(em.logger).Log("error", err.Error())
	}
	return firmware, err
}
func (em *coreService) FirmwareByID(firmwareID EID) (Firmware, error) {
	level.Debug(em.logger).Log("method", "FirmwareByID() called")
	var firmware Firmware
	var err error
	for _, fw := range siteNetworks.Firmwares {
		if fw.ID == firmwareID {
			firmware = fw
			break
		}
	}
	if (firmware == Firmware{}) {
		err = fmt.Errorf("firmware with id {%s} was  not found", firmwareID)
		level.Error(em.logger).Log("error", err.Error())
	}
	return firmware, err
}

func (em *coreService) AllBroadcasts() []Broadcast {
	level.Debug(em.logger).Log("method", "AllBroadcasts() called")
	broadcasts := make([]Broadcast, 3)
	copy(broadcasts, siteNetworks.Broadcasts) // presumed to actually copy
	return broadcasts
}
func (em *coreService) RemoveBroadcastByID(broadcastID string) {
	level.Debug(em.logger).Log("method", "RemoveBroadcastByID() called")
	var index int
	var broadcast Broadcast
	for idx, bc := range siteNetworks.Broadcasts {
		if bc.ID == broadcastID {
			broadcast = bc
			index = idx
			break
		}
	}
	if (broadcast != Broadcast{}) {
		siteNetworks.Broadcasts = append(siteNetworks.Broadcasts[:index], siteNetworks.Broadcasts[index+1:]...) // remove from slice
	}
}
func (em *coreService) BroadcastByID(broadcastID string) (Broadcast, error) {
	level.Debug(em.logger).Log("method", "BroadcastByID() called")
	var broadcast Broadcast
	var err error
	for _, bc := range siteNetworks.Broadcasts {
		if bc.ID == broadcastID {
			broadcast = bc
			break
		}
	}
	if (broadcast == Broadcast{}) {
		err = fmt.Errorf("broadcast with id {%s} was  not found", broadcastID)
		level.Error(em.logger).Log("error", err.Error())
	}
	return broadcast, err
}
