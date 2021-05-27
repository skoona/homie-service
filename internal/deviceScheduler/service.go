package deviceScheduler

/**
  deviceScheduler/service.go:

  Consume OTA Incoming messages and manage the schedule of firmware downloads
*/

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	dss "github.com/skoona/homie-service/internal/deviceSource"
	cc "github.com/skoona/homie-service/internal/utils"
)

/**
 * ConsumeFromDeviceSource
 * Handles incoming channel DM message
 */
func consumeFromDeviceSource(consumer chan dc.DeviceMessage) error {
	/*
	 * Create a Go Routine for the Schedulers Channel to
	 */
	go func(consumeChan chan dc.DeviceMessage) {
		// var err error
		level.Debug(logger).Log("event", "consumeFromDeviceSource(gofunc) called")

		for msg := range consumeChan { // read until closed
			// err := dvService.ApplyOTAEvent(&msg)
			// if err != nil {
			// 	level.Error(logger).Log("method", "consumeFromDeviceSource(gofunc)", "error", err.Error())
			// }

			// if nil != publishChan {
			// 	publishChan <- msg
			// }

			level.Debug(logger).Log("method", "consumeFromDeviceSource(gofunc)", "consume queue depth", len(consumeChan), "Device", msg.DeviceID)
		}
		level.Debug(logger).Log("method", "consumeFromDeviceSource()", "event", "Completed")
	}(consumer)

	return nil
}

func publishToDeviceSource(dm dc.DeviceMessage) error {
	level.Debug(logger).Log("event", "Calling publishToDeviceSource()")
	var err error

	dmh.FromScheduler(dm)

	level.Debug(logger).Log("event", "publishToDeviceSource() completed")
	return err
}
func processSchedulerMessages(dm dc.DeviceMessage) error {
	level.Debug(logger).Log("event", "Calling processSchedulerMessages()")
	var err error

	level.Debug(logger).Log("event", "processSchedulerMessages() completed")
	return err
}

var (
	cfg    cc.Config
	logger log.Logger
	fromDS chan dc.DeviceMessage // in
	dmh    dss.DeviceMessageHandler
	sch    *schedulerService
	ota    *otaChannel
)

/*
 * Start()
 *
 * Initialize this service
 */
func Start(dfg cc.Config, s dss.DeviceMessageHandler) (SchedulerService, error) {
	var err error
	cfg = dfg
	dmh = s

	NewSchedulerService(dfg, s, dc.GetSiteNetworks())
	logger = sch.logger
	s.BuilFirmwareCatalog(sch) // apply scheduler server and induce building firmware catalogs

	level.Debug(logger).Log("event", "Calling Start()")

	// Initialize a Message Channel
	fromDS, err = s.GetSchedulerResponseChannel()
	if err != nil {
		level.Error(logger).Log("event", "Channels offline", "error", err.Error())
		panic(err.Error())
	}
	err = consumeFromDeviceSource(fromDS)

	level.Debug(logger).Log("event", "Start() completed")

	return sch, err
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")

	level.Debug(logger).Log("event", "Stop() completed")
}
