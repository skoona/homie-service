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
	OTAInteractor interface {
		EnableNotifications()  // starts gofunc to receive on status updates
		GetPublishingChannel() // channel which scheduler send out OTA
		SendNotifications()    // sends message to request chan, keeping chan control here
	}

	otaChannel struct {
		requestChannel  chan dc.QueueMessage
		responseChannel chan dc.QueueMessage
		cfg             cc.Config
		logger          log.Logger
	}
)

func NewOTAInteractor(dfg cc.Config) OTAInteractor {
	// make channels
	req := make(chan dc.QueueMessage, 256) // averages 120 on startup
	rsp := make(chan dc.QueueMessage, 256) // averages 120 on startup
	ota = &otaChannel{
		requestChannel:  req,
		responseChannel: rsp,
		cfg:             dfg,
		logger:          log.With(dfg.Logger, "pkg", "deviceScheduler", "service", "OTAInteractor"),
	}
	return ota
}

/*
 * OTAInteractor implementation */
func (s *otaChannel) EnableNotifications() {
	level.Debug(s.logger).Log("event", "EnableNotifications() called")
}

func (s *otaChannel) GetPublishingChannel() {
	level.Debug(s.logger).Log("event", "GetPublishingChannel() called")
}

func (s *otaChannel) SendNotifications(msg dc.QueueMessage) {
	level.Debug(s.logger).Log("event", "SendNotifications() called")
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
