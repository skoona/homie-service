package deviceSource

/*
  deviceSource/sources.go:

  Local StreamProvider Service helpers

  The design goal for this package:
	* Implement Stream Control Helpers in support of deviceCore operations

*/

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"time"
)


// ConsumeFromStreamProvider Stream Listener
func ConsumeFromStreamProvider(consumer chan dc.DeviceMessage, plog log.Logger) {
	/*
	 * Create a Go Routine for the Providers Channel to
	 */
	go func(dmChan chan dc.DeviceMessage, tlog log.Logger) {
		_ = level.Debug(tlog).Log("event", "ConsumeFromDeviceStream(gofunc) called")
		for msg := range dmChan { // read until closed
			tStart := time.Now()

			dvService.sStream <- msg // to repository
			dvService.ApplyDeviceEvent(msg)

			tEnd := time.Now()
			_ = level.Debug(tlog).Log("method", "ConsumeFromDeviceStream(gofunc)", "queue depth", len(dmChan), "device", msg.DeviceID, "duration", tEnd.Sub(tStart))
		}
		_ = level.Debug(tlog).Log("method", "ConsumeFromDeviceStream()", "event", "Completed")
	}(consumer, plog)
}

// StoreDeviceStream Stream Listener
func StoreDeviceStream(consumer chan dc.DeviceMessage, plog log.Logger) {
	/*
	 * Create a Go Routine for the Storage Channel to
	 */
	go func(dmChan chan dc.DeviceMessage, tlog log.Logger) {
		_ = level.Debug(tlog).Log("event", "StoreDeviceStream(gofunc) called")

		for msg := range dmChan { // read until closed
			tStart := time.Now()
			err := dvService.repository.Store(msg)
			if err != nil {
				_ = level.Error(plog).Log("error", err)
			}
			tEnd := time.Now()
			_ = level.Debug(tlog).Log("method", "StoreDeviceStream(gofunc)", "queue depth", len(dmChan), "device", msg.DeviceID, "duration", tEnd.Sub(tStart))
		}
		_ = level.Debug(tlog).Log("method", "StoreDeviceStream()", "event", "Completed")
	}(consumer, plog)
}

