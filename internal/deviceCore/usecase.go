package deviceCore

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	cc "github.com/skoona/homie-service/internal/utils"
)

/*
  	deviceCore/service.go:

	- Build Device Network Inventory
	- deviceSource Conversation
	  IN (from ds)
		* Create DeviceMessage from MQTT Message via QueueMessage
		* Create Device Message from Demo Log
		* Get Request Channel
		* Get Response Channel
	  OUT (to ds)
	  	* Delete Device
		* Activate Schedule
		* Cancel Schedule


*/

// The active Service
var (
	em   *coreService
	dlog = log.With(cc.DefaultLogger, "pkg", "deviceCore", "service", "coreService") // default logger
)

func GetSiteNetworks() *SiteNetworks {
	return &siteNetworks
}

/**
 * ConsumeFromDeviceSource
 * Handles incoming channel DM message
 */
func ConsumeFromDeviceSource(dm DeviceMessage) error {
	var err error
	network, ok := siteNetworks.DeviceNetworks[string(dm.NetworkID)]
	if !ok {
		err := fmt.Errorf("device{%s} not found in network={%s}", dm.DeviceID, network.Name)
		level.Error(em.logger).Log("method", "ConsumeFromDeviceSource(gofunc)", "error", err.Error())
	} else {
		network.apply(dm)
	}

	level.Debug(em.logger).Log("method", "ConsumeFromDeviceSource(gofunc)", "network id", dm.NetworkID, "msg id", dm.ID, "deviceID", dm.DeviceID)

	return err
}

/* Service implementation
 *
 * CoreService for UI or externals
 */
func (em *coreService) ApplyEvent(dm *DeviceMessage) error {
	level.Debug(em.logger).Log("method", "ApplyEvent() called")
	return nil
}
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
func (em *coreService) DeviceByNameFromNetwork(deviceName, networkName string) (Device, error) {
	var err error
	level.Debug(em.logger).Log("method", "DeviceByNameFromNetwork() called")
	copyOfDevice, found := siteNetworks.DeviceNetworks[networkName].Devices[deviceName]
	if !found {
		err = fmt.Errorf("device with name {%s} was  not found on {%s} network", deviceName, networkName)
		level.Error(em.logger).Log("error", err.Error())
	}
	return copyOfDevice, err
}
func (em *coreService) DeviceByEIDFromNetwork(deviceEID EID, networkName string) (Device, error) {
	var err error
	level.Debug(em.logger).Log("method", "DeviceByEIDFromNetwork() called")
	var ptrToDevice *Device
	var copyOfDevice Device
	for _, device := range siteNetworks.DeviceNetworks[networkName].Devices {
		if device.ID == deviceEID {
			ptrToDevice = &device
			break
		}
	}
	if ptrToDevice == nil {
		err = fmt.Errorf("device with id {%s} was  not found on {%s} network", deviceEID, networkName)
		level.Error(em.logger).Log("error", err.Error())
	}
	copyOfDevice = *ptrToDevice // TODO: hoping this copies vs references
	return copyOfDevice, err
}
func (em *coreService) RemoveDeviceByEIDFromNetwork(deviceEID EID, networkName string) error {
	var err error
	level.Debug(em.logger).Log("method", "RemoveDeviceByEIDFromNetwork() called")
	var ptrToDevice *Device
	for _, device := range siteNetworks.DeviceNetworks[networkName].Devices {
		if device.ID == deviceEID {
			ptrToDevice = &device
			break
		}
	}
	if ptrToDevice == nil {
		err = fmt.Errorf("device with id {%s} was  not found on {%s} network", deviceEID, networkName)
		level.Error(em.logger).Log("error", err.Error())
	} else {
		delete(siteNetworks.DeviceNetworks[networkName].Devices, ptrToDevice.Name)
		// TODO: send delete command to deviceSource and Scheduler
	}
	return err
}
func (em *coreService) AllSchedules() []Schedule {
	level.Debug(em.logger).Log("method", "AllSchedules() called")
	schedules := make([]Schedule, len(siteNetworks.Schedules))
	for _, schedule := range siteNetworks.Schedules {
		schedules = append(schedules, schedule)
	}
	return schedules
}
func (em *coreService) AddSchedule(schedule Schedule) {
	level.Debug(em.logger).Log("method", "AddSchedule() called")
	siteNetworks.Schedules[schedule.ID] = schedule
}
func (em *coreService) RemoveSchedule(scheduleID EID) {
	level.Debug(em.logger).Log("method", "RemoveSchedule() called")
	delete(siteNetworks.Schedules, scheduleID)
}
func (em *coreService) ScheduleByEID(scheduleID EID) Schedule {
	level.Debug(em.logger).Log("method", "ScheduleByEID() called")
	schedule := siteNetworks.Schedules[scheduleID]
	return schedule
}
func (em *coreService) ScheduleByDeviceName(deviceName string) Schedule {
	level.Debug(em.logger).Log("method", "ScheduleByDeviceName() called")
	var schedule Schedule
	for _, obj := range siteNetworks.Schedules {
		if obj.DeviceName == deviceName {
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
func (em *coreService) AddFirmware(firmware Firmware) {
	level.Debug(em.logger).Log("method", "AddFirmware() called")
	siteNetworks.Firmwares = append(siteNetworks.Firmwares, firmware)
}
func (em *coreService) RemoveFirmwareByEID(firmwareEID EID) {
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
func (em *coreService) FirmwareByEID(firmwareEID EID) (Firmware, error) {
	level.Debug(em.logger).Log("method", "FirmwareByID() called")
	var firmware Firmware
	var err error
	for _, fw := range siteNetworks.Firmwares {
		if fw.ID == firmwareEID {
			firmware = fw
			break
		}
	}
	if (firmware == Firmware{}) {
		err = fmt.Errorf("firmware with id {%s} was  not found", firmwareEID)
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
func (em *coreService) AddBroadcast(broadcast Broadcast) {
	level.Debug(em.logger).Log("method", "AddBroadcast() called")
	siteNetworks.Broadcasts = append(siteNetworks.Broadcasts, broadcast)
}
func (em *coreService) RemoveBroadcastByEID(broadcastEID EID) {
	level.Debug(em.logger).Log("method", "RemoveBroadcastByEID() called")
	var index int
	var broadcast Broadcast
	for idx, bc := range siteNetworks.Broadcasts {
		if bc.ID == broadcastEID {
			broadcast = bc
			index = idx
			break
		}
	}
	if (broadcast != Broadcast{}) {
		siteNetworks.Broadcasts = append(siteNetworks.Broadcasts[:index], siteNetworks.Broadcasts[index+1:]...) // remove from slice
	}
}
func (em *coreService) BroadcastByEID(broadcastEID EID) (Broadcast, error) {
	level.Debug(em.logger).Log("method", "BroadcastByID() called")
	var broadcast Broadcast
	var err error
	for _, bc := range siteNetworks.Broadcasts {
		if bc.ID == broadcastEID {
			broadcast = bc
			break
		}
	}
	if (broadcast == Broadcast{}) {
		err = fmt.Errorf("broadcast with id {%s} was  not found", broadcastEID)
		level.Error(em.logger).Log("error", err.Error())
	}
	return broadcast, err
}
