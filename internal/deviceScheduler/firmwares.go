package deviceScheduler

/*
  	deviceScheduler/firmwares.go:

	Manage Homie Firmware catalog

*/

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/go-kit/kit/log"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
)

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

/*
   HOMIE_PATTERN   = "\x25\x48\x4f\x4d\x49\x45\x5f\x45\x53\x50\x38\x32\x36\x36\x5f\x46\x57\x25".unpack('H*').first
   NAME_PATTERN    = ["\xbf\x84\xe4\x13\x54".unpack('H*').first, "\x93\x44\x6b\xa7\x75".unpack('H*').first]
   VERSION_PATTERN = ["\x6a\x3f\x3e\x0e\xe1".unpack('H*').first, "\xb0\x30\x48\xd4\x1a".unpack('H*').first]
   BRAND_PATTERN   = ["\xfb\x2a\xf5\x68\xc0".unpack('H*').first, "\x6e\x2f\x0f\xeb\x2d".unpack('H*').first]
*/
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

/*
 * firmwareDetails
 * Firmware and Scheduling
 */
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

/*
 * buildFirmwareCatalog:
 * collect available firmwares
 */
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
			}
			level.Debug(logger).Log("created", "Firmware", "brand", fw.Brand, "object", fw.String())
		}
	}

	level.Debug(logger).Log("event", "buildFirmwareCatalog() completed")
	return firmware
}

/*
 * buildFirmwarePayload()
 * return checksum, byteBuffer, and error
 * transforms file into desired otaFormat
 */
func buildFirmwarePayload(transport dc.OTATransport, fw *dc.Firmware) (string, string, error) {
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

/*
        HOW TO PUBLISH FIRMWARE
ref: https://github.com/skoona/HomieMonitor/blob/master/mains/homie/components/subscription.rb

		# pending
        #   -- send $implementation/ota/firmware/checksum -> firmware
        # wait
        #   -- recv $implementation/ota/firmware/checksum -> MessageBytes=ddd
        #
        # disabled = 403
        # aborted = 400 | 500
        # accepted = 304
        #
        # inprogress = 206 dd/dd
        #
        # aborted = 400 | 500
        # success = 200



firmware
def ota_format
        @ota_format ||= Homie::Handlers::Stream.ota_type
      end

      def ota_format=(ota)
        @ota_format = ota
      end

      def as_binary
        @path.binread
      end

      def as_base64_no_pad_url
        urlsafe_encode64(as_binary, padding: false)
      end

      def as_base64_with_pad_url
        Base64.urlsafe_encode64(as_binary, padding: true)
      end

      def as_strict_base64
        Base64.strict_encode64(as_binary)
      end

      def as_base64
        Base64.encode64(as_binary)
      end
 */