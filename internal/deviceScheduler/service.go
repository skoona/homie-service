package deviceScheduler

/**
  deviceScheduler/service.go:

  Consume OTA Incoming messages and manage the schedule of firmware downloads
*/

import (
	"crypto/md5"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"time"

	dc "github.com/skoona/homie-service/internal/deviceCore"
	cc "github.com/skoona/homie-service/internal/utils"
)

var (
	cfg       cc.Config
	logger    log.Logger
	otaStream OTAInteractor
	sch       *schedulerProvider
)

type (
	/*
	 * call EnableTriggers() to see messages with ota enabled flags from devices
	 * when triggered send the EnableNotificationFor(..., true), then OtaPublish(ota-payload
	 */
	OTAInteractor interface {
		EnableTriggers() chan dc.DeviceMessage                                              // open notification and trigger channel
		EnableNotificationsFor(networkName, deviceName string, enabledOrDisable bool) error // watch a device ota progres
		OtaPublish(otaMessage dc.DeviceMessage)                                             // sending a ota
	}
)

/**
 * consumeFromStreamProvider
 * Handles incoming channel DM message
 */
func consumeFromOTAStreamProvider(consumer chan dc.DeviceMessage, plog log.Logger) {
	/*
	 * Create a Go Routine for the Schedulers Channel to
	 */
	go func(consumeChan chan dc.DeviceMessage, tlog log.Logger) {
		level.Debug(tlog).Log("event", "consumeFromOTAStreamProvider(gofunc) called")

		for msg := range consumeChan { // read until closed

			err := processSchedulerMessages(msg, plog)
			if err != nil {
				level.Error(tlog).Log("method", "consumeFromOTAStreamProvider(gofunc)", "error", err.Error())
			}
			level.Debug(tlog).Log("method", "consumeFromOTAStreamProvider(gofunc)", "consume queue depth", len(consumeChan), "device", msg.DeviceID)
		}
		level.Debug(tlog).Log("method", "consumeFromOTAStreamProvider()", "event", "Completed")
	}(consumer, plog)
}

// NewSchedule Creates Component
// Get related device
// verify OTAEnabled is true [default]
// if false return error
func NewSchedule(networkName string, deviceID string, transport dc.OTATransport, firmware *dc.Firmware) dc.Schedule {
	return dc.Schedule{
		ID:          fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s.%s", networkName, deviceID)))),
		ElementType: dc.CoreTypeSchedule,
		DeviceID:   deviceID,
		Package:     *firmware,
		State:       "pending",
		Status:      "waiting",
		Transport:   transport,
		Scheduled:   time.Now(),
	}
}


// newFirmware Creates Component
func NewFirmware(path string) (dc.Firmware, error) {
	level.Debug(logger).Log("event", "NewFirmware() called", "path", path)
	// call resolver routine
	md5Sum, fwName, fwVersion, fwBrand, fsize, modtime, err := firmwareDetails(path)

	fw := dc.Firmware{
		ID:          dc.EID(md5Sum),
		ElementType: dc.CoreTypeFirmware,
		Name:        fwName,
		FileName:    path,
		Version:     fwVersion,
		Path:        path,
		Size:        fsize,
		MD5Digest:   md5Sum,
		Brand:       fwBrand,
		Created:     modtime,
	}

	return fw, err
}


/**
 * Start()
 *
 * Initialize this service
*/
func Start(dfg cc.Config, s OTAInteractor, r dc.Repository) dc.SchedulerProvider {
	cfg = dfg
	otaStream = s

	NewSchedulerProvider(dfg, s, r)
	logger = sch.logger
	level.Debug(logger).Log("event", "Calling Start()")

	level.Debug(logger).Log("event", "Start() completed")

	return sch
}

/**
 * Stop
 * Cleans up this service
*/
func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")

	level.Debug(logger).Log("event", "Stop() completed")
}
