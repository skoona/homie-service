package deviceSource

/*
  deviceSource/usecase-stream.go:

  DeviceSource Service Implementation
*/

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	cc "github.com/skoona/homie-service/internal/utils"
)

/*
 * PublishToStreamProvider
 */
func (s *deviceSource) PublishToStreamProvider(dm dc.DeviceMessage) {
	logger := log.With(s.logger, "method", "PublishDeviceStream()")

	s.dStream.GetPublishChannel() <- dm

	level.Debug(logger).Log("DeviceID ", dm.DeviceID)
}

// handle incoming device stream events
func (s *deviceSource) ConsumeDeviceStream(dm dc.DeviceMessage) error {
	var err error
	tlog := log.With(s.logger, "method", "ConsumeDeviceStream")

	err = s.repository.Store(dm)
	if err != nil {
		level.Error(tlog).Log("error", err)
		return err
	}

	s.PublishCoreEvent(dm)

	level.Debug(tlog).Log("DeviceID ", dm.DeviceID)

	return err
}

/**
 * ConsumeFromDeviceStream
 * - Stream Listener
 */
func ConsumeFromStreamProvider(consumer chan dc.DeviceMessage, plog log.Logger) {
	/*
	 * Create a Go Routine for the Providers Channel to
	 */
	go func(dmChan chan dc.DeviceMessage, tlog log.Logger) {
		level.Debug(tlog).Log("event", "ConsumeFromDeviceStream(gofunc) called")
		for msg := range dmChan { // read until closed
			err := dvService.ConsumeDeviceStream(msg)
			if err != nil {
				level.Error(tlog).Log("method", "ConsumeFromDeviceStream(gofunc)", "error", err.Error())
			}

			level.Debug(tlog).Log("method", "ConsumeFromDeviceStream(gofunc)", "queue depth", len(dmChan), "device", msg.DeviceID)
		}
		level.Debug(tlog).Log("method", "ConsumeFromDeviceStream()", "event", "Completed")
	}(consumer, plog)
}

// Receive/send DM from Channel
// Receive/send OTA from Channel
// Receive/Send to Scheduler
// Receive/Send to Core

// Route OTA to/from Scheduler
// Route DM to Repository and Core Service
