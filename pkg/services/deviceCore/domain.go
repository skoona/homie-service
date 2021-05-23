package deviceCore

import (
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
