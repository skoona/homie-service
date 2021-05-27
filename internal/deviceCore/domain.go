package deviceCore

import (
	"fmt"
	"time"
)

var (
	siteNetworks SiteNetworks
)

// Globals required for DeviceMessage Types
type CoreType int

// CoreType Basic object types
const (
	CoreTypeDevice CoreType = (iota + 10)
	CoreTypeDeviceAttribute
	CoreTypeDeviceAttributeProperty
	CoreTypeDeviceAttributePropertyProperty
	CoreTypeDeviceNode
	CoreTypeDeviceNodeAttribute
	CoreTypeDeviceNodeProperty
	CoreTypeDeviceNodePropertyAttribute
	CoreTypeAttribute
	CoreTypeProperty
	CoreTypeBroadcast
	CoreTypeFirmware
	CoreTypeSchedule
	CoreTypeNetwork
	CoreTypeSiteNetworks
)

// OTATransport Flag used to choose OTA transport Format
type OTATransport int

// ENUM Flags used to choose OTA transport Format
const (
	Binary OTATransport = iota + 30
	Base64
	Base64Strict
	RFC4648URLSafeWithPadding
	RFC4648URLSafeWithoutPadding
)

// Firmware basic firmware model
type Firmware struct {
	ID          EID
	ElementType CoreType
	Name        string
	FileName    string
	Version     string
	Path        string
	Size        int64
	MD5Digest   string
	Brand       string
	Created     time.Time
}

// DeviceSchedule the ota schedule details
type Schedule struct {
	ID          EID
	ElementType CoreType
	DeviceName  string
	Package     Firmware
	State       string
	Status      string
	Transport   OTATransport
	Scheduled   time.Time
	Completed   time.Time
}

type FirmwareTools interface {
	String()
}
type ScheduleTools interface {
	String()
}

func (f *Firmware) String() string {
	return fmt.Sprintf("id=%s, name=%s, fileName=%s, path=%s, size=%d, brand=%s, created=%v",
		f.ID, f.Name, f.FileName, f.Path, f.Size, f.Brand, f.Created)
}
func (s *Schedule) String() string {
	return fmt.Sprintf("id=%s, device=%s, state=%s, status=%s, scheduled=%v, firmware.id=%s, firmware.name=%s",
		s.ID, s.DeviceName, s.State, s.Status, s.Scheduled, s.Package.ID, s.Package.Name)
}
