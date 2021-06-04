package deviceCore

import (
	"fmt"
	"github.com/go-kit/kit/log/level"
)

/*
  	deviceCore/usecase-stream.go:

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

