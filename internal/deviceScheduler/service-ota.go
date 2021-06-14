package deviceScheduler

/**
  deviceScheduler/service-ota.go:

  Interactions with MQTT Providers
*/

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
)

type (
	/*
	 * call EnableTriggers() to see messages with ota enabled flags from devices
	 * when triggered send the EnableNotificationFor(..., true), then OtaPublish(ota-payload
	 */
	OTAInteractor interface {
		EnableTriggers() chan dc.DeviceMessage                                              // open notification and trigger channel
		EnableNotificationsFor(networkName, deviceName string, enabledOrDisable bool) error // watch a device ota progres
		OtaPublish(otaMessage dc.DeviceMessage)                                             // sending a ota
	}
)

func publishOTAStream(dm dc.DeviceMessage, plog log.Logger) {
	level.Debug(plog).Log("event", "Calling publishToDeviceSource()")

	otaStream.OtaPublish(dm)

	level.Debug(plog).Log("event", "publishToDeviceSource() completed")
	return
}

func consumeOTAStream(dm dc.DeviceMessage, plog log.Logger) error {
	level.Debug(plog).Log("event", "Calling consumeOTAStream()")

	err := processSchedulerMessages(dm, plog)
	if err != nil {
		level.Error(plog).Log("method", "consumeOTAStream()", "error", err.Error(), "device", dm.DeviceID)
	}

	level.Debug(plog).Log("method", "consumeOTAStream()", "device", dm.DeviceID, "topic", dm.Topic(), "value", dm.Value)

	level.Debug(plog).Log("event", "consumeOTAStream() completed")
	return err
}

/**
 * consumeFromStreamProvider
 * Handles incoming channel DM message
 */
func consumeFromOTAStreamProvider(consumer chan dc.DeviceMessage, plog log.Logger) {
	/*
	 * Create a Go Routine for the Schedulers Channel to
	 */
	go func(consumeChan chan dc.DeviceMessage, tlog log.Logger) {
		level.Debug(tlog).Log("event", "consumeFromOTAStreamProvider(gofunc) called")

		for msg := range consumeChan { // read until closed

			err := consumeOTAStream(msg, tlog)
			if err != nil {
				level.Error(tlog).Log("method", "consumeFromOTAStreamProvider(gofunc)", "error", err.Error())
			}

			// if nil != publishChan {
			// 	publishChan <- msg
			// }
			level.Debug(tlog).Log("method", "consumeFromOTAStreamProvider(gofunc)", "consume queue depth", len(consumeChan), "device", msg.DeviceID)
		}
		level.Debug(tlog).Log("method", "consumeFromOTAStreamProvider()", "event", "Completed")
	}(consumer, plog)
}
