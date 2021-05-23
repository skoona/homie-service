package deviceScheduler

/**
  deviceScheduler/service.go:

  Consume OTA Incoming messages and manage the schedule of firmware downloads
*/

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/pkg/services/deviceCore"
	dss "github.com/skoona/homie-service/pkg/services/deviceSource"
	cc "github.com/skoona/homie-service/pkg/utils"
)

/**
 * ConsumeFromDeviceSource
 * Handles incoming channel DM message
 */
func ConsumeFromDeviceSource(consumer chan dc.DeviceMessage, publisher chan dc.DeviceMessage) error {
	/*
	 * Create a Go Routine for the Schedulers Channel to
	 */
	go func(consumeChan chan dc.DeviceMessage, publishChan chan dc.DeviceMessage) {
		var err error
		level.Debug(logger).Log("event", "ConsumeFromDeviceSource(gofunc) called")

		for msg := range consumeChan { // read until closed
			// err := dvService.ApplyOTAEvent(&msg)
			if err != nil {
				level.Error(logger).Log("method", "ConsumeFromDeviceSource(gofunc)", "error", err.Error())
			}

			// if nil != publishChan {
			// 	publishChan <- msg
			// }

			level.Debug(logger).Log("method", "ConsumeFromDeviceSource(gofunc)", "consume queue depth", len(consumeChan), "Device", msg.DeviceID)
		}
		level.Debug(logger).Log("method", "ConsumeFromDeviceSource()", "event", "Completed")
	}(consumer, publisher)

	return nil
}

func PublishToDeviceSource(dm dc.DeviceMessage, publisher chan dc.DeviceMessage) error {
	var err error

	return err
}
func processSchedulerMessages(dm dc.DeviceMessage) error {
	var err error

	return err
}

var (
	cfg    cc.Config
	logger log.Logger
	fromDS chan dc.DeviceMessage // in
	toDS   chan dc.DeviceMessage // out
)

/*
 * Start()
 *
 * Initialize this service
 */
func Start(dfg cc.Config) error {
	var err error
	cfg = dfg

	logger = log.With(dfg.Logger, "pkg", "deviceScheduler")
	level.Debug(logger).Log("event", "Calling Start()")

	// Initialize a Message Channel
	fromDS, toDS, err = dss.ChannelsForScheduler()
	if err != nil {
		level.Error(logger).Log("event", "Channels offline", "error", err.Error())
		panic(err.Error())
	}

	err = ConsumeFromDeviceSource(fromDS, toDS)

	level.Debug(logger).Log("event", "Start() completed")

	return err
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")
	if nil != toDS {
		close(toDS) // only the send should shutdown channels
	}

	level.Debug(logger).Log("event", "Stop() completed")
}
