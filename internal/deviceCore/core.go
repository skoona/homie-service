package deviceCore

/*
	deviceCore/core.go

	Defines the site collection of network models


 * Component layout
 *                                    Topic Place
 * Device                             X/D
 * DeviceAttribute                    X/D/$A
 * DeviceAttributeProperty            X/D/$A/P
 * DeviceAttributePropertyProperty    X/D/$A/P/P
 * DeviceNode                         X/D/N
 * DeviceNodeAttribute                X/D/N/$A
 * DeviceNodeProperty                 X/D/N/P
 * DeviceNodePropertyAttribute        X/D/N/P/$A
 *
 * X=Network, D=Device, N=Node, $A=Attribute, P=Property
*/


import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"

	"github.com/go-kit/kit/log/level"
)

var siteNetworks SiteNetworks

// Globals required for DeviceMessage Types
type CoreType int

// OTATransport Flag used to choose OTA transport Format
type OTATransport int

//EID entity EID
type EID string

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
	CoreTypeDeviceDelete
	CoreTypePublishMessage
	CoreTypeBroadcast
	CoreTypeFirmware
	CoreTypeSchedule
	CoreTypeNetwork
	CoreTypeSiteNetworks
)

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
	ID          string
	ElementType CoreType
	DeviceID    string
	FirmwareID  EID
	State       string
	Status      string
	Retries     int
	Transport   OTATransport
	Scheduled   time.Time
	Completed   time.Time
}

// Network contains all known devices in application
type Network struct {
	ID          string
	Title       string
	ElementType CoreType // network name
	Name        string   // any
	Devices     map[string]Device
}

// SiteNetworks contains all known/valid Collections
type SiteNetworks struct {
	ID             EID
	ElementType    CoreType // site name
	SiteName       string
	Title          string
	Names          []string // any
	Broadcasts     []Broadcast
	Firmwares      []Firmware
	Schedules      map[string]Schedule   // by EID
	DeviceNetworks map[string]Network // by Device Name
}

// Broadcast Alert Messages on  Network
type Broadcast struct {
	ID          string
	ElementType CoreType
	Parent      string
	Topic       string
	Level       string
	Value       string
	Received    time.Time
}

// Device represent a physical device
type Device struct {
	ID          EID
	Title       string
	ElementType CoreType
	OTAEnabled  bool
	Parent      string
	Name        string
	Attrs       map[string]DeviceAttribute
	Nodes       map[string]DeviceNode
}

// DeviceAttribute of a Device allowing sub properties, representing a single value or measurement
type DeviceAttribute struct {
	ID          EID
	ElementType CoreType
	Parent      string
	Name        string
	Value       string
	Props       map[string]DeviceAttributeProperty
}

// DeviceAttributeProperty
// of a Device Attribute, representing a single value or measurement
// $implementation/ota/status -- where status will be a regular attribute
type DeviceAttributeProperty struct {
	ID          EID
	ElementType CoreType
	Parent      string
	Name        string
	Value       string
	Props       map[string]DeviceAttributePropertyProperty
}

// DeviceAttributePropertyProperty
// of a Device Attribute, representing a single value or measurement
// $implementation/ota/status -- where status will be a regular attribute
type DeviceAttributePropertyProperty struct {
	ID          EID
	ElementType CoreType
	Parent      string
	Name        string
	Value       string
}

// Node representing the capabilities or features of device
type DeviceNode struct {
	ID          EID
	ElementType CoreType
	Parent      string
	Name        string
	Attrs       map[string]DeviceNodeAttribute
	Props       map[string]DeviceNodeProperty
}

// Attribute used by Devices, Nodes, and Properties.  Used to describe format
type DeviceNodeAttribute struct {
	ID          EID
	ElementType CoreType
	Parent      string
	Name        string
	Value       string
}

// NodeProperty of a Node, representing a single value or measurement
type DeviceNodeProperty struct {
	ID          EID
	ElementType CoreType
	Parent      string
	Name        string
	Value       string
	Attrs       map[string]DeviceNodePropertyAttribute
}

// DeviceNodePropertyAttribute used by Devices, Nodes, and Properties.  Used to describe format
type DeviceNodePropertyAttribute struct {
	ID          EID
	ElementType CoreType
	Parent      string
	Name        string
	Value       string
}

// NewNetworks Creates Component
func NewSiteNetworks(siteName, siteTitle string, networks []string, firmwares []Firmware, schedules map[string]Schedule) *SiteNetworks {
	//level.Debug(em.logger).Log("event", "NewSiteNetworks() called")

	siteNetworks = SiteNetworks{
		ID:             NewEID(),
		ElementType:    CoreTypeSiteNetworks,
		SiteName:       siteName,
		Title:          siteTitle,
		Names:          networks,
		DeviceNetworks: make(map[string]Network, len(networks)+2),
		Broadcasts:     []Broadcast{},
		Firmwares:      firmwares,
		Schedules:      schedules, // make key id ID
	}

	for _, nName := range networks {
		siteNetworks.DeviceNetworks[nName] = NewNetwork(nName, nName)
	}

	//level.Debug(em.logger).Log("event", "NewSiteNetworks() completed")
	return &siteNetworks
}

// NewNetwork Creates Component
func NewNetwork(title, name string) Network {
	//level.Debug(em.logger).Log("event", "NewNetwork() called", "title", title, "name", name)
	return Network{
		ID:          fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s.%s", title, name)))),
		ElementType: CoreTypeNetwork,
		Title:       title,
		Name:        name,
		Devices:     make(map[string]Device, 16),
	}
}


// newBroadcast Creates Component
func NewBroadcast(parent, topic, level, value string) Broadcast {
	bc := Broadcast{
		ID:          fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s.%s.%s", parent, topic, level)))),
		ElementType: CoreTypeBroadcast,
		Parent:      parent,
		Topic:       topic,
		Level:       level,
		Value:       value,
		Received:    time.Now(),
	}

	return bc
}


//NewID create a new entity ID
func NewEID() EID {
	uuid := uuid.New()
	return EID(uuid.String())
}

// NewDevice Creates Component
// read $implementation/ota/enabled -> bool and set the
// OTAEnabled for Schedules
func NewDevice(parent, name string) Device {
	return Device{
		ID:          NewEID(),
		ElementType: CoreTypeDevice,
		OTAEnabled:  true,
		Parent:      parent,
		Name:        name,
		Attrs:       make(map[string]DeviceAttribute, 4),
		Nodes:       make(map[string]DeviceNode, 4),
	}
}

// NewDeviceAttribute Creates Component
func NewDeviceAttribute(parent, name, value string) DeviceAttribute {

	return DeviceAttribute{
		ID:          NewEID(),
		ElementType: CoreTypeDeviceAttribute,
		Parent:      parent,
		Name:        name,
		Value:       value,
		Props:       make(map[string]DeviceAttributeProperty, 4),
	}
}

// NewDeviceAttributeProperty Creates Component
// of a Device Attribute, representing a single value or measurement
// $implementation/fw -- where status will be a regular attribute
//  $DeviceAttr / Property
func NewDeviceAttributeProperty(parent, name, value string) DeviceAttributeProperty {
	return DeviceAttributeProperty{
		ID:          NewEID(),
		ElementType: CoreTypeDeviceAttributeProperty,
		Parent:      parent,
		Name:        name,
		Value:       value,
		Props:       make(map[string]DeviceAttributePropertyProperty, 4),
	}
}

// NewDeviceAttributePropertyProperty Creates Component
// of a Device Attribute, representing a single value or measurement
// $implementation/fw/version -- where status will be a regular attribute
//  $DeviceAttr / Property / Property
func NewDeviceAttributePropertyProperty(parent, name, value string) DeviceAttributePropertyProperty {
	return DeviceAttributePropertyProperty{
		ID:          NewEID(),
		ElementType: CoreTypeDeviceAttributePropertyProperty,
		Parent:      parent,
		Name:        name,
		Value:       value,
	}
}

// NewDeviceNode Creates Component
func NewDeviceNode(parent, name string) DeviceNode {
	return DeviceNode{
		ID:          NewEID(),
		ElementType: CoreTypeDeviceNode,
		Parent:      parent,
		Name:        name,
		Attrs:       make(map[string]DeviceNodeAttribute, 4),
		Props:       make(map[string]DeviceNodeProperty, 4),
	}
}

// NewDeviceNodeAttribute Creates Component
func NewDeviceNodeAttribute(parent, name, value string) DeviceNodeAttribute {
	return DeviceNodeAttribute{
		ID:          NewEID(),
		ElementType: CoreTypeDeviceNodeAttribute,
		Parent:      parent,
		Name:        name,
		Value:       value,
	}
}

// NewDeviceNodeProperty Creates Component
func NewDeviceNodeProperty(parent, name, value string) DeviceNodeProperty {
	return DeviceNodeProperty{
		ID:          NewEID(),
		ElementType: CoreTypeDeviceNodeProperty,
		Parent:      parent,
		Name:        name,
		Value:       value,
		Attrs:       make(map[string]DeviceNodePropertyAttribute, 4),
	}
}

// NewDeviceNodePropertyAttribute Creates Component
func NewDeviceNodePropertyAttribute(parent, name, value string) DeviceNodePropertyAttribute {
	return DeviceNodePropertyAttribute{
		ID:          NewEID(),
		ElementType: CoreTypeDeviceNodePropertyAttribute,
		Parent:      parent,
		Name:        name,
		Value:       value,
	}
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
	return fmt.Sprintf("id=%s, device=%s, state=%s, status=%s, scheduled=%v, firmware.id=%s",
		s.ID, s.DeviceID, s.State, s.Status, s.Scheduled, s.FirmwareID)
}


// EntityBuilder
// - based on Network Entity
// - manages creation and deletion of components
type EntityBuilder interface {
	apply(dm DeviceMessage) error
	handleBroadcast(dm DeviceMessage) error
	handleDeviceAttributePropertyProperty(dm DeviceMessage) error
	handleDeviceAttributeProperty(dm DeviceMessage) error
	handleDeviceAttribute(dm DeviceMessage) error
	handleDevice(dm DeviceMessage) error
	handleNodePropertyAttribute(dm DeviceMessage) error
	handleNodeProperty(dm DeviceMessage) error
	handleNodeAttribute(dm DeviceMessage) error
	handleNode(dm DeviceMessage) error
}

// removeAttributePrefix()
// - Removes prefixes added by dm builders
// - prefix default = "_$"
func removeAttributePrefix(attribute []byte, prefix string) string{
	var nValue string
	if strings.HasPrefix(string(attribute), prefix){
		startAtPos := len(prefix)
		nValue = string(attribute)[startAtPos:]
	} else{
		nValue = string(attribute)
	}
	return nValue
}

// Apply device updates to Network
func (hn *Network) apply(dm DeviceMessage) error {
	var err error
	level.Debug(em.logger).Log("event", "apply() called")

	// ensure this device is in our network
	dv, ok := hn.Devices[string(dm.DeviceID)]
	if !ok {
		err = fmt.Errorf("device{%s} not found in network={%s}, should be created", dm.DeviceID, hn.Name)
		level.Debug(em.logger).Log("action", err.Error())
	}

	/*
	 * Delete the device if value is nil
	 * TODO determine if this nil'ed message is looping
	 * -- i.e we send it, and also receive it which makes us send it again...
	 */
	if ok && (string(dm.Value) == "" || dm.Value == nil) {
		err := em.RemoveDeviceByID(string(dv.ID), dv.Parent)

		err = fmt.Errorf("device{%s} on network{%s} was deleted since value was nil", dm.DeviceID, hn.Name)
		level.Warn(em.logger).Log("action", err.Error())
		return err
	}

	/*
	 * Create Components
	 */
	switch dm.HomieType {
	case CoreTypeBroadcast:
		err = hn.handleBroadcast(dm)
	case CoreTypeDeviceAttributePropertyProperty:
		err = hn.handleDeviceAttributePropertyProperty(dm)
	case CoreTypeDeviceAttributeProperty:
		err = hn.handleDeviceAttributeProperty(dm)
	case CoreTypeDeviceAttribute:
		err = hn.handleDeviceAttribute(dm)
	case CoreTypeDevice:
		err = hn.handleDevice(dm)
	case CoreTypeDeviceNodePropertyAttribute:
		err = hn.handleNodePropertyAttribute(dm)
	case CoreTypeDeviceNodeAttribute:
		err = hn.handleNodeAttribute(dm)
	case CoreTypeDeviceNodeProperty:
		err = hn.handleNodeProperty(dm)
	case CoreTypeDeviceNode:
		err = hn.handleNode(dm)
	}

	level.Debug(em.logger).Log("event", "apply() completed")
	return err
}

// handleDeviceAttributePropertyProperty on HomieNetork
func (hn *Network) handleDeviceAttributePropertyProperty(dm DeviceMessage) error {
	var err error
	level.Debug(em.logger).Log("event", "handleDeviceAttributePropertyProperty() called")
	nValue := removeAttributePrefix(dm.AttributeID, "_$")

	dev, found := hn.Devices[string(dm.DeviceID)]
	if !found {
		err = fmt.Errorf("device [%s] not found, will create it", string(dm.DeviceID))
		level.Debug(em.logger).Log("warning", err.Error())
		dev = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
		hn.Devices[string(dm.DeviceID)] = dev
	}

	devattr, found := dev.Attrs[nValue]
	if !found {
		err = fmt.Errorf("attribute [%s] not found, will create it", nValue)
		level.Debug(em.logger).Log("warning", err.Error())
		devattr = NewDeviceAttribute(string(dm.DeviceID), nValue, "")
		dev.Attrs[nValue] = devattr
	}

	devattrprop, found := devattr.Props[string(dm.PropertyID)]
	if !found {
		err = fmt.Errorf("property [%s] not found, will create it", string(dm.PropertyID))
		level.Debug(em.logger).Log("warning", err.Error())
		devattrprop = NewDeviceAttributeProperty(nValue, string(dm.PropertyID), "")
		devattr.Props[string(dm.PropertyID)] = devattrprop
	}

	devattrpropprop, found := devattrprop.Props[string(dm.PPropertyID)]
	if !found {
		err = fmt.Errorf("property pproperty [%s] not found, will create it", string(dm.PPropertyID))
		level.Debug(em.logger).Log("warning", err.Error())
		devattrpropprop = NewDeviceAttributePropertyProperty(string(dm.PropertyID), string(dm.PPropertyID), string(dm.Value))
		devattrprop.Props[string(dm.PropertyID)] = devattrpropprop
	}

	level.Debug(em.logger).Log("event", "handleDeviceAttributePropertyProperty() completed", "id", devattrpropprop.ID, "name", devattrpropprop.Name)
	return err
}

// handleDeviceAttributeProperty on HomieNetork
func (hn *Network) handleDeviceAttributeProperty(dm DeviceMessage) error {
	var err error
	level.Debug(em.logger).Log("event", "handleDeviceAttributeProperty() called")
	nValue := removeAttributePrefix(dm.AttributeID, "_$")

	dev, found := hn.Devices[string(dm.DeviceID)]
	if !found {
		err = fmt.Errorf("device [%s] not found, will create it", string(dm.DeviceID))
		level.Debug(em.logger).Log("warning", err.Error())
		dev = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
		hn.Devices[string(dm.DeviceID)] = dev
	}

	devattr, found := dev.Attrs[nValue]
	if !found {
		err = fmt.Errorf("attribute [%s] not found, will create it", nValue)
		level.Debug(em.logger).Log("warning", err.Error())
		devattr = NewDeviceAttribute(string(dm.DeviceID), nValue, "")
		dev.Attrs[nValue] = devattr
	}

	devattrprop, found := devattr.Props[string(dm.PropertyID)]
	if !found {
		err = fmt.Errorf("attribute property [%s] not found, will create it", string(dm.PropertyID))
		level.Debug(em.logger).Log("warning", err.Error())
		devattrprop = NewDeviceAttributeProperty(nValue, string(dm.PropertyID), string(dm.Value))
		devattr.Props[string(dm.PropertyID)] = devattrprop
	}

	level.Debug(em.logger).Log("event", "handleDeviceAttributeProperty() completed", "id", devattrprop.ID, "name", devattrprop.Name)
	return err
}

// handleDeviceAttribute on HomieNetork
func (hn *Network) handleDeviceAttribute(dm DeviceMessage) error {
	var err error
	level.Debug(em.logger).Log("event", "handleDeviceAttribute() called")
	nValue := removeAttributePrefix(dm.AttributeID, "_$")

	dev, found := hn.Devices[string(dm.DeviceID)]
	if !found {
		err = fmt.Errorf("device [%s] not found, will create it", string(dm.DeviceID))
		level.Debug(em.logger).Log("warning", err.Error())
		dev = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
		hn.Devices[string(dm.DeviceID)] = dev
	}

	devattr, found := dev.Attrs[nValue]
	if !found {
		err = fmt.Errorf("attribute [%s] not found, will create it", nValue)
		level.Debug(em.logger).Log("warning", err.Error())
		devattr = NewDeviceAttribute(string(dm.DeviceID), nValue, string(dm.Value))
		dev.Attrs[string(dm.AttributeID)] = devattr
	}

	level.Debug(em.logger).Log("event", "handleDeviceAttribute() completed", "id", devattr.ID, "name", devattr.Name)
	return err
}

// handleDevice on HomieNetork
func (hn *Network) handleDevice(dm DeviceMessage) error {
	var err error
	level.Debug(em.logger).Log("event", "handleDevice() called")

	dev, found := hn.Devices[string(dm.DeviceID)]
	if !found {
		err = fmt.Errorf("device [%s] not found, will create it", string(dm.DeviceID))
		level.Debug(em.logger).Log("warning", err.Error())
		dev = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
		hn.Devices[string(dm.DeviceID)] = dev
	}

	level.Debug(em.logger).Log("event", "handleDevice() completed", "id", dev.ID, "name", dev.Name)
	return err
}

// handleNodePropertyAttribute on HomieNetork
func (hn *Network) handleNodePropertyAttribute(dm DeviceMessage) error {
	var err error
	level.Debug(em.logger).Log("event", "handleNodePropertyAttribute() called")
	nValue := removeAttributePrefix(dm.AttributeID, "_$")

	dev, found := hn.Devices[string(dm.DeviceID)]
	if !found {
		err = fmt.Errorf("device [%s] not found, will create it", string(dm.DeviceID))
		level.Debug(em.logger).Log("warning", err.Error())
		dev = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
		hn.Devices[string(dm.DeviceID)] = dev
	}

	node, found := dev.Nodes[string(dm.NodeID)]
	if !found {
		err = fmt.Errorf("node [%s] not found, will create it", string(dm.NodeID))
		level.Debug(em.logger).Log("warning", err.Error())
		node = NewDeviceNode(string(dm.DeviceID), string(dm.NodeID))
		dev.Nodes[string(dm.NodeID)] = node
	}

	nodeprop, found := node.Props[string(dm.PropertyID)]
	if !found {
		err = fmt.Errorf("property [%s] not found, will create it", string(dm.PropertyID))
		level.Debug(em.logger).Log("warning", err.Error())
		nodeprop = NewDeviceNodeProperty(string(dm.NodeID), string(dm.PropertyID), "")
		node.Props[string(dm.PropertyID)] = nodeprop
	}

	nodepropattr, found := nodeprop.Attrs[nValue]
	if !found {
		err = fmt.Errorf("property attribute [%s] not found, will create it", nValue)
		level.Debug(em.logger).Log("warning", err.Error())
		nodepropattr = NewDeviceNodePropertyAttribute(string(dm.PropertyID), nValue, string(dm.Value))
		nodeprop.Attrs[nValue] = nodepropattr
	}

	level.Debug(em.logger).Log("event", "handleNodePropertyAttribute() completed", "id", nodepropattr.ID, "name", nodepropattr.Name)
	return err
}

// HandleNodeProperty on HomieNetork
func (hn *Network) handleNodeProperty(dm DeviceMessage) error {
	var err error
	level.Debug(em.logger).Log("event", "handleNodeProperty() called")

	dev, found := hn.Devices[string(dm.DeviceID)]
	if !found {
		err = fmt.Errorf("device [%s] not found, will create it", string(dm.DeviceID))
		level.Debug(em.logger).Log("warning", err.Error())
		dev = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
		hn.Devices[string(dm.DeviceID)] = dev
	}

	node, found := dev.Nodes[string(dm.NodeID)]
	if !found {
		err = fmt.Errorf("node [%s] not found, will create it", string(dm.NodeID))
		level.Debug(em.logger).Log("warning", err.Error())
		node = NewDeviceNode(string(dm.DeviceID), string(dm.NodeID))
		dev.Nodes[string(dm.NodeID)] = node
	}

	nodeprop, found := node.Props[string(dm.PropertyID)]
	if !found {
		err = fmt.Errorf("property [%s] not found, will create it", string(dm.PropertyID))
		level.Debug(em.logger).Log("warning", err.Error())
		nodeprop = NewDeviceNodeProperty(string(dm.NodeID), string(dm.PropertyID), string(dm.Value))
		node.Props[string(dm.PropertyID)] = nodeprop
	}

	level.Debug(em.logger).Log("event", "handleNodeProperty() completed", "id", nodeprop.ID, "name", nodeprop.Name)
	return err
}

// HandleNodeAttribute on HomieNetork
func (hn *Network) handleNodeAttribute(dm DeviceMessage) error {
	var err error
	level.Debug(em.logger).Log("event", "handleNodeAttribute() called")
	nValue := removeAttributePrefix(dm.AttributeID, "_$")

	dev, found := hn.Devices[string(dm.DeviceID)]
	if !found {
		err = fmt.Errorf("device [%s] not found, will create it", string(dm.DeviceID))
		level.Debug(em.logger).Log("warning", err.Error())
		dev = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
		hn.Devices[string(dm.DeviceID)] = dev
	}

	node, found := dev.Nodes[string(dm.NodeID)]
	if !found {
		err = fmt.Errorf("node [%s] not found, will create it", string(dm.NodeID))
		level.Debug(em.logger).Log("warning", err.Error())
		node = NewDeviceNode(string(dm.DeviceID), string(dm.NodeID))
		dev.Nodes[string(dm.NodeID)] = node
	}

	nodeattr, found := node.Attrs[nValue]
	if !found {
		err = fmt.Errorf("attribute [%s] not found, will create it", nValue)
		level.Debug(em.logger).Log("warning", err.Error())
		nodeattr = NewDeviceNodeAttribute(string(dm.NodeID), nValue, string(dm.Value))
		node.Attrs[nValue] = nodeattr
	}

	level.Debug(em.logger).Log("event", "handleNodeAttribute() completed", "id", nodeattr.ID, "name", nodeattr.Name)
	return err
}

// handleNode on HomieNetork
func (hn *Network) handleNode(dm DeviceMessage) error {
	var err error
	level.Debug(em.logger).Log("event", "handleNode() called")

	dev, found := hn.Devices[string(dm.DeviceID)]
	if !found {
		err = fmt.Errorf("device [%s] not found, will create it", string(dm.DeviceID))
		level.Debug(em.logger).Log("warning", err.Error())
		dev = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
		hn.Devices[string(dm.DeviceID)] = dev
	}

	node, found := dev.Nodes[string(dm.NodeID)]
	if !found {
		err = fmt.Errorf("node [%s] not found, will create it", string(dm.NodeID))
		level.Debug(em.logger).Log("warning", err.Error())
		node = NewDeviceNode(string(dm.DeviceID), string(dm.NodeID))
		dev.Nodes[string(dm.NodeID)] = node
	}

	level.Debug(em.logger).Log("event", "handleNode() completed", "id", node.ID, "name", node.Name)
	return err
}

// HandleBroadcast on HomieNetork
func (hn *Network) handleBroadcast(dm DeviceMessage) error {
	var err error
	var found bool
	var bcn Broadcast
	var index int

	level.Debug(em.logger).Log("event", "handleBroadcast() called")

	for idx, bc := range siteNetworks.Broadcasts {
		if bc.Topic == string(dm.AttributeID) {
			found = true
			index = idx
			bcn = bc
			break
		}
	}
	if found && string(dm.Value) == "" { // delete if value is empty
		siteNetworks.Broadcasts = append(siteNetworks.Broadcasts[:index], siteNetworks.Broadcasts[index+1:]...)
		err = fmt.Errorf("broadcast [%s] was deleted", string(dm.AttributeID))
		level.Debug(em.logger).Log("action", err.Error())
	} else { // add it blindly
		err = fmt.Errorf("broadcast [%s] not found, will create it", string(dm.AttributeID))
		level.Debug(em.logger).Log("action", err.Error())
		bcn = NewBroadcast(string(dm.NetworkID), string(dm.AttributeID), string(dm.PropertyID), string(dm.Value))
		siteNetworks.Broadcasts = append(siteNetworks.Broadcasts, bcn)
	}

	level.Debug(em.logger).Log("event", "handleBroadcast() completed", "id", bcn.ID, "topic", bcn.Topic, "value", bcn.Value)
	return err
}
