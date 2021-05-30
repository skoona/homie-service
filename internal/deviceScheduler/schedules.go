package deviceScheduler

import (
	"time"

	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
)

/**
  deviceScheduler/schedules.go:

The design goal for this file is:
	* Discover existing ESP8266/ESP32 Firmware in dataDB directotry
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
func NewSchedule(networkName, deviceName string, transport dc.OTATransport, firmware *dc.Firmware) dc.Schedule {
	return dc.Schedule{
		ID:          dc.NewEID(),
		ElementType: dc.CoreTypeSchedule,
		DeviceName:  deviceName,
		Package:     *firmware,
		State:       "Pending",
		Status:      "Waiting",
		Transport:   transport,
		Scheduled:   time.Now(),
	}
}

func buildScheduleCatalog() map[dc.EID]dc.Schedule {
	level.Debug(logger).Log("event", "buildScheduleCatalog() called")
	// get from levelDB
	schedMap := sch.repo.ScheduleStore(dc.Schedule{})
	sch.snwk.Schedules = schedMap

	return schedMap
}

/*
 * processSchedulerMessages()
 * - handle notifications, triggers, and new core requests
 */
func processSchedulerMessages(dm dc.DeviceMessage) error {
	level.Debug(logger).Log("event", "Calling processSchedulerMessages()")
	var err error
	level.Debug(logger).Log("topic", dm.TopicS, "device", dm.DeviceID)

	level.Debug(logger).Log("event", "processSchedulerMessages() completed")
	return err
}
