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
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	dc "github.com/skoona/homie-service/pkg/deviceCore"
	cc "github.com/skoona/homie-service/pkg/utils"
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

			err := scheduleProcessor(msg, plog)
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
func NewSchedule(networkName string, deviceID string, transport dc.OTATransport, firmwareEID dc.EID) dc.Schedule {
	return dc.Schedule{
		ID:          fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s.%s", networkName, deviceID)))),
		ElementType: dc.CoreTypeSchedule,
		DeviceID:    deviceID,
		FirmwareID:  firmwareEID,
		State:       "pending",
		Status:      "waiting",
		Transport:   transport,
		Scheduled:   time.Now(),
	}
}


// newFirmware Creates Component
func NewFirmware(sFile, dFile string) (dc.Firmware, error) {
	destinationFile := filepath.FromSlash(dFile)
	sourceFile := filepath.FromSlash(sFile)

	level.Debug(logger).Log("event", "NewFirmware() called", "source", sourceFile, "destination", destinationFile)
	var dstFile string

	if destinationFile != "" { // will be empty when creating library
		fileInfo, err := os.Stat(sourceFile)
		if err != nil {
			return dc.Firmware{}, fmt.Errorf("NewFirmware() file[%s] is not found: %s", sourceFile, err.Error())
		}

		// apply destination path
		// -- may fail with windows path input when running on linux server
		if fileInfo.IsDir() {
			dstFile = fmt.Sprintf("%s%s%s", cfg.Dbc.FirmwareStorage, string(os.PathSeparator), filepath.Base(sourceFile))
		} else {
			dstFile = fmt.Sprintf("%s%s%s", cfg.Dbc.FirmwareStorage, string(os.PathSeparator), filepath.Base(destinationFile))
		}

		// Copy file to storage location
		input, err := ioutil.ReadFile(sourceFile)
		if err != nil {
			return dc.Firmware{}, fmt.Errorf("NewFirmware() error reading source[%s]; %s", sourceFile, err.Error())
		}

		err = ioutil.WriteFile(dstFile, input, 0644)
		if err != nil {
			return dc.Firmware{}, fmt.Errorf("NewFirmware() error writing destination[%s]; %s", destinationFile, err.Error())
		}
	} else {
		dstFile = sourceFile
	}

	// call resolver routine
	md5Sum, fwName, fwVersion, fwBrand, fsize, modtime, err := firmwareDetails(dstFile)
	if err != nil {
		os.Remove(dstFile)  // cleanup if not firmware
		return dc.Firmware{}, err
	}

	// create firmware object
	fw := dc.Firmware{
		ID:          dc.EID(md5Sum),
		ElementType: dc.CoreTypeFirmware,
		Name:        fwName,
		FileName:    dstFile,
		Version:     fwVersion,
		Path:        dstFile,
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
