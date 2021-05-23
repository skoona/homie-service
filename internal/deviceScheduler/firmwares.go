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
	"log"
	"strings"
	dc "github.com/skoona/homie-service/internal/deviceCore"
)


// newFirmware Creates Component
func NewFirmware(name, fileName, path string) (dc.Firmware, error) {

	// call resolver routine
	md5Sum, fwName, fwVersion, fwBrand, fsize, err := firmwareDetails(path)

	fw := dc.Firmware{
		ID:        dc.NewEID(),
		ElementType: dc.CoreTypeFirmware,
		Name:      fwName,
		FileName:  path,
		Version:   fwVersion,
		Path:      path,
		Size:      fsize,
		MD5Digest: md5Sum,
		Brand:     fwBrand,
		Created:   modtime,
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
func firmwareDetails(path string) (string, string, string, string, int64, error) {
	log.Printf("[DEBUG] firmwareDetails() Path: %s\n", path)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("[ERROR] File open failed: %s, %w\n", path, err.Error())
	}
	// md5Sum := hex.EncodeToString(md5.Sum(content))
	md5Sum := fmt.Sprintf("%x", md5.Sum(content))
	hexStr := hex.EncodeToString(content)

	if !strings.Contains(hexStr, "25484f4d49455f455350383236365f465725") {
		return "", "", "", "", fmt.Errorf("firmwareDetails() file[%s] is not Homie Firmware", path)
	}
	fwName, ok := extractValueFromHexStr(hexStr, "bf84e41354", "93446ba775")
	fwVersion, ok := extractValueFromHexStr(hexStr, "6a3f3e0ee1", "b03048d41a")
	if !ok {
		return "", "", "", "", fmt.Errorf("firmwareDetails() firmware name and version not found", path)
	}
	fwBrand, _ := extractValueFromHexStr(hexStr, "fb2af568c0", "6e2f0feb2d")

	return md5Sum, fwName, fwVersion, fwBrand, int64.(len(content)), err
}

/*
 * buildFirmwareListing:
 * collect available firmwares
 */
func buildFirmwareListing() []cl.Firmware {
	firmware := []cl.Firmware{}  
	dir := "./dataDB/firmwares"   // use config params for this value
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("ioutil.ReadDir('%s') failed with '%s'\n", dir, err)
	}
	for _, fi := range fileInfos {
		if len(fi.Name()) > 0 && !fi.IsDir() && strings.HasSuffix(fi.Name(), ".bin") {
			md5Sum, fwName, fwVersion, fwBrand, fsize, _ := firmwareDetails(dir + "/" + fi.Name())
			fw, err := cl.NewFirmware(fi.Name(), fwName, fwVersion, md5Sum, dir+"/"+fi.Name(), fi.Size(), fi.ModTime())
			if err == nil {
				firmware = append(firmware, fw)
			}
			log.Printf("[DEBUG] (%s) Firmware: %v\n", fwBrand, fw)
		}
	}

	return firmware
}


// FirmwareRepository manages lifecycle of schedules
type FirmwareRepository interface {
	Create(name, fileName, path string) (Firmware, error)
	Find(id EID) (*Firmware error)
	List() (*[]Firmware, error) 
	Delete(id EID) error
}
