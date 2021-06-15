package deviceScheduler

import (
	"crypto/md5"
	"fmt"
	"github.com/go-kit/kit/log"
	"strings"
	"time"

	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
)

/**
  deviceScheduler/schedules.go:

The design goal for this file is:
	* Discover existing ESP8266/ESP32 Firmware in dataDB directory
	* Attach the discovered collection to the central HomeNetworks struct
	* Create Device OTA schedules and prepare for execution of an OTA download
	* Execute an OTA Download through MQTT using the Homie protocol
	* Allow for the creation of Schedules and the Upload of Firmwares
	* Allow for the replacement or deletion of both.

	It works this way:

	During startup of the Homie for ESP8266 device, it reports the current
	firmware's MD5 to $fw/checksum (in addition to $fw/name and $fw/version).
	The OTA entity may or may not use this information to automatically
	schedule OTA updates

	The OTA entity publishes the latest available firmware payload
	to $implementation/ota/firmware/<md5 checksum>, either as binary or
	as a Base64 encoded string

	If OTA is disabled, Homie for ESP8266 reports 403 to
	$implementation/ota/status and aborts the OTA

	If OTA is enabled and the latest available checksum is the same as
	what is currently running, Homie for ESP8266 reports 304 and aborts the OTA

	If the checksum is not a valid MD5, Homie for ESP8266 reports 400 BAD_CHECKSUM
	to $implementation/ota/status and aborts the OTA

	Homie starts to flash the firmware

	The firmware is updating. Homie for ESP8266 reports progress with
	206 <bytes written>/<bytes total>

	When all bytes are flashed, the firmware is verified (including the
		MD5 if one was set)

	Homie for ESP8266 either reports 200 on success, 400 if the
	firmware in invalid or 500 if there's an internal error

	Homie for ESP8266 reboots on success as soon as the device is idle

	OTA¶
		$implementation/ota/enabled: true if OTA is enabled, false otherwise

		A Good Trigger point is:
			$state → "ready"
		This is the state the device is in when it is connected to the MQTT broker,
		has sent all Homie messages and is ready to operate.

		$implementation/ota/firmware: If the update request is accepted, you
			must send the firmware payload to this topic
		$implementation/ota/status: HTTP-like status code indicating the status
			of the OTA. Might be:

		Code	Description
		200	OTA successfully flashed
		202	OTA request / checksum accepted
		206 465/349680	OTA in progress. The data after the status code corresponds to <bytes written>/<bytes total>
		304	The current firmware is already up-to-date
		400 BAD_FIRMWARE	OTA error from your side. The identifier might be BAD_FIRMWARE, BAD_CHECKSUM, NOT_ENOUGH_SPACE, NOT_REQUESTED
		403	OTA not enabled
		500 FLASH_ERROR	OTA error on the ESP8266. The identifier might be FLASH_ERROR



*/

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

func buildScheduleCatalog(plog log.Logger) map[string]dc.Schedule {
	level.Debug(plog).Log("event", "buildScheduleCatalog() called")
	return sch.repo.LoadSchedules()
}
//
//schedule.Status = "waiting", "initializing", "downloading", "dd/dddd", "complete"
//schedule.State = "pending", "active", ["accepted", "aborted", "rejected"], "complete", "current"

func handleOTATrigger(dm dc.DeviceMessage, schedule *dc.Schedule, plog log.Logger) error  {
	level.Debug(plog).Log("event", "Calling handleOTATrigger()")
	dvm := dc.DeviceMessage{}
	var err error

	// match existing
	// "$fw/checksum" is the only trigger word
	if !strings.Contains(dm.Topic(), "$fw/checksum") {
		return err
	}

	// already being handled
	if schedule.State != "pending" {
		return err
	}

	// ensure md5's are different
	if strings.Contains(string(dm.Value), schedule.Package.MD5Digest) {
		schedule.Completed = time.Now()
		schedule.Status = "complete"
		schedule.State = "current"
		return err
	}

	// build and send OTA
	mhash, bundle, err := buildFirmwarePayload(schedule.Transport, schedule.Package)
	if err != nil {
		level.Error(plog).Log("event", "handleOTATrigger()", "error", err.Error())
		return err
	}

	dvm.TopicS = fmt.Sprintf("%s/%s/$implementation/ota/firmware/%s", dm.NetworkID, dm.DeviceID, mhash)
	dvm.NetworkID = dm.NetworkID
	dvm.DeviceID = dm.DeviceID
	dvm.Value = []byte(bundle)
	dvm.RetainedB = false
	schedule.Status = "initializing"
	schedule.State = "active"
	schedule.Scheduled = time.Now()

	level.Debug(plog).Log("event", "handleOTATrigger()", "schedule", schedule.ID)

	otaStream.OtaPublish(dvm)

	level.Debug(plog).Log("event", "handleOTATrigger() completed")
	return err
}
func handleOTAStatus(dm dc.DeviceMessage, schedule *dc.Schedule, plog log.Logger) error  {
	level.Debug(plog).Log("event", "Calling handleOTAStatus()")
	if dm.OTACurrent() {
		schedule.Completed = time.Now()
		schedule.Status = "complete"
		schedule.State = "current"
		otaStream.EnableNotificationsFor(string(dm.NetworkID), string(dm.DeviceID), true)
	} else if dm.OTAActive() {
		schedule.Status = string(dm.Value)[3:] // 203 dd/dddd
		schedule.State = "active"
	} else if dm.OTAAborted() {
		schedule.Status = "complete"
		schedule.State = "aborted"
		otaStream.EnableNotificationsFor(string(dm.NetworkID), string(dm.DeviceID), false)
	} else if dm.OTARejected() {
		schedule.Status = "complete"
		schedule.State = "rejected"
		otaStream.EnableNotificationsFor(string(dm.NetworkID), string(dm.DeviceID), false)
	}
	level.Debug(plog).Log("event", "handleOTAStatus()", "schedule", schedule.ID, "status", schedule.Status)

	level.Debug(plog).Log("event", "handleOTAStatus() completed")
	return nil
}
func handleOTAComplete(dm dc.DeviceMessage, schedule *dc.Schedule, plog log.Logger) error  {
	level.Debug(plog).Log("event", "Calling handleOTAComplete()")
	var err error
	if dm.OTAComplete() {
		err = otaStream.EnableNotificationsFor(string(dm.NetworkID), string(dm.DeviceID), false)
		schedule.Completed = time.Now()
		schedule.Status = "complete"
		schedule.State = "complete"
		level.Debug(plog).Log("event", "handleOTAComplete()", "schedule", schedule.ID, "state", "completed")
	}
	level.Debug(plog).Log("event", "handleOTAComplete() completed")
	return err
}

/*
 * processSchedulerMessages()
 * - handle notifications, triggers, and new core requests
 */
func processSchedulerMessages(dm dc.DeviceMessage, plog log.Logger) error {
	level.Debug(plog).Log("event", "Calling processSchedulerMessages()")
	var err error
	deviceID := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s.%s", dm.NetworkID, dm.DeviceID))))
	schedule := sch.FindScheduleByDeviceID(deviceID)
	if schedule.ElementType != dc.CoreTypeSchedule {  // empty schedule or not found
		return fmt.Errorf("no schedule available for device %s", dm.DeviceID)
	}

	level.Debug(plog).Log("topic", dm.TopicS, "device", dm.DeviceID, "value", dm.Value)
	if dm.OTATrigger {
		err = handleOTATrigger(dm, schedule, plog)
	} else 	if dm.OTAComplete() {
		err = handleOTAComplete(dm, schedule, plog)
	} else 	if dm.OTAStatus() {
		err = handleOTAStatus(dm, schedule, plog)
	}

	level.Debug(plog).Log("event", "processSchedulerMessages() completed")
	return err
}
