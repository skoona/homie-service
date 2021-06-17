package deviceScheduler

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

/**
  deviceScheduler/scheduler.go:

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
			$fw/checksum → "any-value"
		This is the state the device is in when it is connected to the MQTT broker,
		has sent all Homie messages and is ready to operate.

		$implementation/ota/firmware: If the update is needed, you
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

	ref: https://github.com/skoona/HomieMonitor/blob/master/mains/homie/components/subscription.rb

		# pending
        #   -- send $implementation/ota/firmware/checksum -> firmware
        # wait
        #   -- recv $implementation/ota/firmware/checksum -> MessageBytes=ddd
        # status
        #   -- recv $implementation/ota/status -> 200 ddd/ddddd


*/


func buildScheduleCatalog(plog log.Logger) map[string]dc.Schedule {
	level.Debug(plog).Log("event", "buildScheduleCatalog() called")
	return sch.repo.LoadSchedules()
}

// handleOTATrigger()
// - schedule.Status = "waiting", "initializing", "current", "firmware error", "flash error", "failure", "dd/dddd", "success"
// - schedule.State = "pending", "active", "aborted", "rejected", "complete"
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
		schedule.Status = "current"
		schedule.State = "complete"
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

	otaStream.EnableNotificationsFor(string(dvm.NetworkID), string(dvm.DeviceID), false)
	otaStream.OtaPublish(dvm)

	level.Debug(plog).Log("event", "handleOTATrigger() completed")
	return err
}
func handleOTAStatus(dm dc.DeviceMessage, schedule *dc.Schedule, plog log.Logger) error  {
	level.Debug(plog).Log("event", "Calling handleOTAStatus()")
	if dm.OTACurrent() {
		schedule.Completed = time.Now()
		schedule.Status = "current"
		schedule.State = "complete"
		otaStream.EnableNotificationsFor(string(dm.NetworkID), string(dm.DeviceID), false)
	}  else if dm.OTAPublishMessage() { // some other controller is publishing
		schedule.State = "active"
		otaStream.EnableNotificationsFor(string(dm.NetworkID), string(dm.DeviceID), true)
	} else if dm.OTAActive() {
		schedule.Status = string(dm.Value)[3:] // 203 dd/dddd
		schedule.State = "active"
	} else if dm.OTAAborted() {
		schedule.State = "pending"
		if dm.OTAFlashError() {
			schedule.Status = "flash error"
			schedule.Retries += 1
		} else if dm.OTAFirmwareError() {
			schedule.Status = "firmware error"
			schedule.Retries += 1
		} else {
			schedule.Status = "failure"
			schedule.State = "complete"
			schedule.Completed = time.Now()
		}
		if schedule.Retries >= 3 {
			schedule.State = "complete"
			schedule.Completed = time.Now()
		}
		otaStream.EnableNotificationsFor(string(dm.NetworkID), string(dm.DeviceID), false)
	} else if dm.OTARejected() {
		schedule.Status = "rejected"
		schedule.State = "complete"
		schedule.Completed = time.Now()
		otaStream.EnableNotificationsFor(string(dm.NetworkID), string(dm.DeviceID), false)
	}
	level.Debug(plog).Log("event", "handleOTAStatus()", "schedule", schedule.ID, "status", schedule.Status)

	level.Debug(plog).Log("event", "handleOTAStatus() completed")
	return nil
}
func handleOTAComplete(dm dc.DeviceMessage, schedule *dc.Schedule, plog log.Logger) error  {
	level.Debug(plog).Log("event", "Calling handleOTAComplete()")
	err := otaStream.EnableNotificationsFor(string(dm.NetworkID), string(dm.DeviceID), false)
	schedule.Completed = time.Now()
	schedule.Status = "success"
	schedule.State = "complete"
	level.Debug(plog).Log("event", "handleOTAComplete() completed", "schedule",
		schedule.ID, "status", schedule.Status, "state", schedule.State)
	return err
}

// scheduleProcessor()
// - handle notifications, triggers, and new core requests
func scheduleProcessor(dm dc.DeviceMessage, plog log.Logger) error {
	level.Debug(plog).Log("event", "Calling scheduleProcessor()")
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
	} else 	if dm.OTAPublishMessage() {
		err = handleOTAStatus(dm, schedule, plog)
	}

	level.Debug(plog).Log("event", "schedulePprocessor() completed")
	return err
}

// extractValueFromHexStr()
//   HOMIE_PATTERN   = "\x25\x48\x4f\x4d\x49\x45\x5f\x45\x53\x50\x38\x32\x36\x36\x5f\x46\x57\x25".unpack('H*').first
//   NAME_PATTERN    = ["\xbf\x84\xe4\x13\x54".unpack('H*').first, "\x93\x44\x6b\xa7\x75".unpack('H*').first]
//   VERSION_PATTERN = ["\x6a\x3f\x3e\x0e\xe1".unpack('H*').first, "\xb0\x30\x48\xd4\x1a".unpack('H*').first]
//   BRAND_PATTERN   = ["\xfb\x2a\xf5\x68\xc0".unpack('H*').first, "\x6e\x2f\x0f\xeb\x2d".unpack('H*').first]
func extractValueFromHexStr(str string, startS string, endS string) (result string, found bool) {
	s := strings.Index(str, startS)
	if s == -1 {
		return result, false
	}
	newS := str[s+len(startS):]
	e := strings.Index(newS, endS)
	if e == -1 {
		return result, false
	}
	result = newS[:e]

	data, err := hex.DecodeString(result)
	if err != nil {
		return "", false
	}
	result = string(data)

	return result, true
}

// firmwareDetails()
//  - Firmware and Scheduling
func firmwareDetails(path string) (string, string, string, string, int64, time.Time, error) {
	level.Debug(logger).Log("event", "firmwareDetails() called", "filePath", path)
	fileInfo, err := os.Stat(path)
	if err != nil {
		return "", "", "", "", 0, time.Time{}, fmt.Errorf("firmwareDetails() file[%s] is not found: %s", path, err.Error())
	}
	modtime := fileInfo.ModTime()

	content, err := ioutil.ReadFile(path)
	if err != nil {
		level.Error(logger).Log("file open failed", path, "error", err.Error())
		return "", "", "", "", 0, time.Time{}, fmt.Errorf("fileopen(%s) failed: %s", path, err.Error())
	}

	// md5Sum := hex.EncodeToString(md5.Sum(content))
	md5Sum := fmt.Sprintf("%x", md5.Sum(content))
	hexStr := hex.EncodeToString(content)

	if !strings.Contains(hexStr, "25484f4d49455f455350383236365f465725") {
		return "", "", "", "", 0, time.Time{}, fmt.Errorf("firmwareDetails() file[%s] is not Homie Firmware", path)
	}
	fwName, _ := extractValueFromHexStr(hexStr, "bf84e41354", "93446ba775")
	fwVersion, ok := extractValueFromHexStr(hexStr, "6a3f3e0ee1", "b03048d41a")
	if !ok {
		return "", "", "", "", 0, time.Time{}, fmt.Errorf("firmware %s and version not found", path)
	}
	fwBrand, _ := extractValueFromHexStr(hexStr, "fb2af568c0", "6e2f0feb2d")

	contentLen := int64(len(content))

	return md5Sum, fwName, fwVersion, fwBrand, contentLen, modtime, err
}

// buildFirmwareCatalog()
// -- collect available firmwares
func buildFirmwareCatalog() []dc.Firmware {
	level.Debug(logger).Log("event", "buildFirmwareCatalog() called")

	firmware := []dc.Firmware{}
	dir := cfg.Dbc.FirmwareStorage // use config params for this value
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		level.Error(logger).Log("ioutil.ReadDir()", "failed", "dir", dir, "error", err.Error())
	}
	for _, fi := range fileInfos {
		if len(fi.Name()) > 0 && !fi.IsDir() && strings.HasSuffix(fi.Name(), ".bin") {
			fw, err := NewFirmware(dir + "/" + fi.Name())
			if err == nil {
				firmware = append(firmware, fw)
				level.Debug(logger).Log("created", "Firmware", "brand", fw.Brand, "object", fw.String())
			}
		}
	}

	level.Debug(logger).Log("event", "buildFirmwareCatalog() completed")
	return firmware
}

// buildFirmwarePayload()
// - return checksum, byteBuffer, and error
// - transforms file into desired otaFormat
func buildFirmwarePayload(transport dc.OTATransport, fw dc.Firmware) (string, string, error) {
	var content string

	mlog := log.With(logger,"method", "buildFirmwarePayload()")
	level.Debug(mlog).Log("event", "buildFirmwarePayload() called")

	buffer, err := ioutil.ReadFile(fw.Path)
	if err != nil {
		level.Error(mlog).Log("file open failed", fw.Path, "error", err.Error())
		return "", "", fmt.Errorf("fileopen(%s) failed: %s", fw.Path, err.Error())
	}

	if int64(len(buffer)) != fw.Size {
		return "", "", fmt.Errorf("file(%s) size has changed: %d", fw.FileName, fw.Size)
	}

	switch transport {
	case dc.Binary:
		content = string(buffer)
	case dc.Base64:
		content = base64.StdEncoding.EncodeToString(buffer)
	case dc.Base64Strict:
		content = base64.StdEncoding.Strict().EncodeToString(buffer)
	case dc.RFC4648URLSafeWithPadding:
		content = base64.URLEncoding.EncodeToString(buffer)
	case dc.RFC4648URLSafeWithoutPadding:
		content = base64.RawURLEncoding.EncodeToString(buffer)
	}

	level.Debug(mlog).Log("event", "buildFirmwarePayload() completed")
	return fw.MD5Digest, content, err
}
