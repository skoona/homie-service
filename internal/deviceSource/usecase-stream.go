package deviceSource

/*
  deviceSource/usecase-stream.go:

  DeviceSource Service Implementation
*/

import (
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	sch "github.com/skoona/homie-service/internal/deviceScheduler"
	cc "github.com/skoona/homie-service/internal/utils"
)

type (

	StreamProvider interface {
		GetPublishChannel() chan dc.QueueMessage
		GetNotifyChannel() chan dc.QueueMessage
	}

	// Device Source Storage Repository
	Repository interface {
		Store(d dc.DeviceMessage) error
	}

	deviceMessageHandler struct {
		cfg    cc.Config
		repository Repository
		dStream StreamProvider
		logger log.Logger
	}
)

// Retained DeviceSource Service, once created
var (
	dmHandler *deviceMessageHandler
)

/**
 * NewDeviceMessageService()
 *
 *  Handler for incoming device stream messages
 */
func NewDeviceMessageHandler(dfg cc.Config, stream StreamProvider, plog log.Logger) DeviceMessageHandler {
	dmHandler = &deviceMessageHandler{
		cfg:    dfg,
		dStream:  stream,
		logger: log.With(plog, "service", "DeviceMessageHandler"),
	}
	return dmHandler
}

/*
 * DeviceSourceInteractor
 */
/**
 * ConsumeFromProviders
 * Handles incoming channel DM message
 */
func consumeFromProviders(consumer chan dc.DeviceMessage) error {
	/*
	 * Create a Go Routine for the Providers Channel to
	 */
	go func(dmChan chan dc.DeviceMessage) {
		level.Debug(logger).Log("event", "ConsumeFromDMProviders(gofunc) called")
		for msg := range dmChan { // read until closed
			err := dvService.ApplyDMEvent(msg)
			if err != nil {
				level.Error(logger).Log("method", "ConsumeFromDMProviders(gofunc)", "error", err.Error())
			}
			level.Debug(logger).Log("method", "ConsumeFromDMProviders(gofunc)", "queue depth", len(dmChan))
		}
		level.Debug(logger).Log("method", "ConsumeFromDMProviders(gofunc)", "event", "Completed")
	}(consumer)

	return nil
}


// Receive/send DM from Channel
// Receive/send OTA from Channel
// Receive/Send to Scheduler
// Receive/Send to Core

// Route OTA to/from Scheduler
// Route DM to Repository and Core Service
