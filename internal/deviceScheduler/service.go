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

type (
	SchedulerService interface {
		ApplySchedule(scheduleEID, deviceEID dc.EID)
		StartSchedule(scheduleEID dc.EID)
		DeleteSchedule(scheduleEID dc.EID)
		UpdateScheduler(dm dc.DeviceMessage)

		ApplyFirmwareUpload(blob []byte) error
		DeleteFirmware(firmwareEID dc.EID)
		GetFirmware(firmwareEID dc.EID) (dc.Firmware, error)
		GetFirmwareMeta(firmwareEID dc.EID)
	}

	schedulerService struct {
		sites  *dc.SiteNetworks
		cfg    cc.Config
		logger log.Logger
	}
)

func NewSchedulerService(dfg cc.Config, siteNetworks *dc.SiteNetworks) SchedulerService {
	sch = &schedulerService{
		sites:  siteNetworks,
		cfg:    dfg,
		logger: log.With(dfg.Logger, "pkg", "deviceScheduler", "service", "SchedulerService"),
	}
	return sch
}

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
)

/*
 * Start()
 *
 * Initialize this service
 */
func Start(dfg cc.Config, s dss.DeviceMessageHandler) error {
	var err error
	cfg = dfg
	dmh = s

	NewSchedulerService(dfg, dc.GetSiteNetworks())
	logger = sch.logger

	level.Debug(sch.logger).Log("event", "Calling Start()")

	// Initialize a Message Channel
	fromDS, err = s.GetSchedulerResponseChannel()
	if err != nil {
		level.Error(sch.logger).Log("event", "Channels offline", "error", err.Error())
		panic(err.Error())
	}
	err = consumeFromDeviceSource(fromDS)

	level.Debug(sch.logger).Log("event", "Start() completed")

	return err
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")

	level.Debug(logger).Log("event", "Stop() completed")
}
