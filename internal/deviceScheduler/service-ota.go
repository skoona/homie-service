package deviceScheduler

/**
  deviceScheduler/service-ota.go:

  Interactions with MQTT Providers
*/

import (
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	cc "github.com/skoona/homie-service/internal/utils"
)

type (
	/*
	 * call EnableTriggers() to see messages with ota enabled flags from devices
	 * when triggered send the EnableNotificationFor(..., true), then OtaPublish(ota-payload
	*/
	OTAInteractor interface {
		EnableTriggers() chan dc.QueueMessage  // open notification and trigger channel
		EnableNotificationsFor(networkName, deviceName string, enabledOrDisable bool) error // watch a device ota progres
		OtaPublish(otaMessage dc.QueueMessage) // sending a ota
	}

)

var (
	otaNotify chan dc.QueueMessage
)

/*
 * Receive Triggers and Notifications from OTA Provider  */
func enableNotifications(s OTAInteractor) error {
	level.Debug(s.logger).Log("method", "enableNotifications()")
	otaNotify = make(chan dc.QueueMessage, 120)
	err := consumeNotificationsFromOTAProviders(otaNotify)
	if err == nil {
		err = s.EnableNotifications(otaNotify)
		if err != nil {
			level.Error(logger).Log("error", err.Error())
		}
	} else {
		level.Error(logger).Log("error", err.Error())
	}
	return err
}

/**
 * ConsumeFromOTAProviders
 * routes from OTA to Scheduler
 */
func consumeNotificationsFromOTAProviders(consumer chan dc.QueueMessage) error {
	/*
	 * Create a Go Routine for the Providers Channel to
	 */
	go func(consumeChan chan dc.QueueMessage) {
		level.Debug(logger).Log("event", "consumeNotificationsFromOTAProviders(gofunc) called")
		for msg := range consumeChan { // read until closed

			err := otaChannel.ApplyOTAEvent(msg)
			if err != nil {
				level.Error(logger).Log("method", "consumeNotificationsFromOTAProviders(gofunc)", "error", err.Error())
			}

			level.Debug(logger).Log("method", "consumeNotificationsFromOTAProviders(gofunc)", "consume queue depth", len(consumeChan))
		}
		level.Debug(logger).Log("method", "consumeNotificationsFromOTAProviders()", "event", "Completed")
	}(consumer)

	return nil
}
