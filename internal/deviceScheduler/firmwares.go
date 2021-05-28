package deviceScheduler

/*
  	deviceScheduler/firmwares.go:

	Manage Homie Firmware catalog

*/

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
)

// newFirmware Creates Component
func NewFirmware(path string) (dc.Firmware, error) {

	// call resolver routine
	md5Sum, fwName, fwVersion, fwBrand, fsize, modtime, err := firmwareDetails(path)

	fw := dc.Firmware{
		ID:          dc.NewEID(),
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
		return "", "", "", "", 0, time.Time{}, fmt.Errorf("fileopen failed", path, "error", err.Error())
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

	return firmware
}

// FirmwareRepository manages lifecycle of schedules
type FirmwareRepository interface {
	List() []dc.Firmware
	Find(id dc.EID) (dc.Firmware, error)
	Create(path string) error
	Delete(id dc.EID) error
}
type firmwareRepo struct {
	snwk	dc.SiteNetworks
}


func NewFirmwareRepository(nwks dc.SiteNetworks) FirmwareRepository {
	return
}