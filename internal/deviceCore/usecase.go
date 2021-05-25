package deviceCore

import (
	"errors"
	"fmt"

	"github.com/go-kit/kit/log/level"
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
	em   coreService
	cdss coreDeviceSourceService

	fromDeviceSource chan DeviceMessage // in
	toDeviceSource   chan DeviceMessage // out
)

/**
 * ConsumeFromDeviceSource
 * Handles incoming channel DM message
 */
func ConsumeFromDeviceSource(consumer chan DeviceMessage) error {
	/*
	 * Create a Go Routine for the Providers Channel to
	 */
	go func(dsChan chan DeviceMessage) {
		level.Debug(cdss.logger).Log("event", "ConsumeFromDeviceSource(gofunc) called")
		for msg := range dsChan { // read until closed
			network, ok := siteNetworks.DeviceNetworks[string(msg.NetworkID)]
			if !ok {
				err := fmt.Errorf("device{%s} not found in network={%s}", msg.DeviceID, network.Name)
				level.Error(cdss.logger).Log("method", "ConsumeFromDeviceSource(gofunc)", "error", err.Error())
			} else {
				network.apply(msg)
			}

			level.Debug(cdss.logger).Log("method", "ConsumeFromDeviceSource(gofunc)", "queue depth", len(dsChan), "network id", msg.NetworkID, "msg id", msg.ID, "deviceID", msg.DeviceID)
		}
		level.Debug(cdss.logger).Log("method", "ConsumeFromDeviceSource()", "event", "Completed")
	}(consumer)

	return nil
}

// DeviceSourceInteractor
func (cdss *coreDeviceSourceService) CreateDemoDeviceMessage(topic string, payload []byte, idCounter uint16, retained bool, qos byte) (DeviceMessage, error) {
	level.Debug(cdss.logger).Log("method", "CreateDemoDeviceMessage() called")
	return NewDeviceMessage(topic, payload, idCounter, retained, qos)
}
func (cdss *coreDeviceSourceService) CreateQueueDeviceMessage(qmsg QueueMessage) (DeviceMessage, error) {
	level.Debug(cdss.logger).Log("method", "CreateQueueDeviceMessage() called")
	return NewQueueMessage(qmsg)
}
func (cdss *coreDeviceSourceService) GetCoreRequestChannel() (chan DeviceMessage, error) {
	level.Debug(cdss.logger).Log("method", "GetCoreRequestChannel() called")
	var err error
	if fromDeviceSource == nil {
		fromDeviceSource = make(chan DeviceMessage, 256) // averages 120 on startup
		if fromDeviceSource != nil {
			err = ConsumeFromDeviceSource(fromDeviceSource)
		}
	}

	if nil == fromDeviceSource {
		err = errors.New("create core publishing channel failed")
		level.Error(em.logger).Log("error", err.Error())
		return nil, err
	}

	return fromDeviceSource, err
}
func (cdss *coreDeviceSourceService) GetCoreResponseChannel() (chan DeviceMessage, error) {
	level.Debug(cdss.logger).Log("method", "GetCoreResponseChannel() called")
	var err error
	if toDeviceSource == nil {
		toDeviceSource = make(chan DeviceMessage, 256) // averages 120 on startup
	}

	if nil == toDeviceSource {
		err = errors.New("create core subscribing channel failed")
		level.Error(em.logger).Log("error", err.Error())
		return nil, err
	}
	return toDeviceSource, err
}

// EventHandler manages mqtt input messages
// Includes channel to/from a device Source

func (em *coreService) ApplyEvent(dm *DeviceMessage) error {
	level.Debug(em.logger).Log("method", "ApplyEvent() called")
	return nil
}

// AccessHomieNetworks

// HomieNetworkByName Returns a HomieNetwork of Devices
func NetworkByName(network string) Network {
	return siteNetworks.DeviceNetworks[network]
}

// DeviceSchedules returns map[device-name]schedule of all know schedules
func DeviceSchedules() *map[string]Schedule {
	return &siteNetworks.Schedules
}

// DeviceFirmwares returns firmware inventory
func DeviceFirmwares() *[]Firmware {
	return &siteNetworks.Firmwares
}

// NetworkBroadcasts returns a list of network broadcast
func NetworkBroadcasts() *[]Broadcast {
	return &siteNetworks.Broadcasts
}
