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
	dc "github.com/skoona/homie-service/internal/deviceCore"
)


// ConsumeFromDeviceStream
// - Stream Listener
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
