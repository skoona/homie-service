package deviceSource

/*
  deviceSource/usecase-core.go:

  DeviceSource Service Implementation
*/

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	cc "github.com/skoona/homie-service/internal/utils"
)

func (s *deviceSource) PublishCoreEvent(dm dc.DeviceMessage) {
	logger := log.With(s.logger, "method", "PublishCoreEvent")

	// also sent it to core
	// s.coreSvc.FromDeviceSource(dm)

	level.Debug(logger).Log("DeviceID ", dm.DeviceID)
}

// handle incoming core events
func (s *deviceSource) ConsumeCoreEvent(dm dc.DeviceMessage) error {
	var err error
	plog := log.With(s.logger, "method", "ConsumeCoreEvent")

	// can only be a delete request
	err = s.repository.Store(dm)
	if err != nil {
		level.Error(plog).Log("error", err)
		return err
	}

	s.PublishToStreamProvider(dm)

	level.Debug(plog).Log("DeviceID ", dm.DeviceID)
}

/**
 * ConsumeFromCore
 * Handles incoming channel DM message
 */
func ConsumeFromCore(consumer chan dc.DeviceMessage, plog log.Logger) error {
	/*
	 * Create a Go Routine for the Providers Channel to
	 */
	go func(dmChan chan dc.DeviceMessage, tlog log.Logger) {
		level.Debug(tlog).Log("event", "ConsumeFromCore(gofunc) called")
		for msg := range dmChan { // read until closed
			err := dvService.ConsumeCoreEvent(msg)
			if err != nil {
				level.Error(tlog).Log("method", "ConsumeFromCore(gofunc)", "error", err.Error())
			}

			level.Debug(tlog).Log("method", "ConsumeFromCore(gofunc)", "queue depth", len(dmChan), "device", msg.DeviceID)
		}
		level.Debug(tlog).Log("method", "ConsumeFromCore()", "event", "Completed")
	}(consumer, plog)

	return nil
}
