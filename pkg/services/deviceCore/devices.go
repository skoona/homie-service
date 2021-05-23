package deviceCore

/**
	deviceCore/devices.go

The design goal for this file is:
	* Establish the core Device data model

*/

// Network contains all known devices in application
type Network struct {
	ID          EID
	Title       string
	ElementType CoreType // network name
	Name        string   // any
	Devices     map[string]Device
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

/*
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

// NewDevice Creates Component
// read $implementation/ota/enabled -> bool and set the
// OTAEnabled for Schedules
func NewDevice(parent, name string) (Device, error) {
	return Device{
		ID:          NewEID(),
		ElementType: CoreTypeDevice,
		OTAEnabled:  true,
		Parent:      parent,
		Name:        name,
		Attrs:       make(map[string]DeviceAttribute, 4),
		Nodes:       make(map[string]DeviceNode, 4),
	}, nil
}

// NewDeviceAttribute Creates Component
func NewDeviceAttribute(parent, name, value string) (DeviceAttribute, error) {
	return DeviceAttribute{
		ID:          NewEID(),
		ElementType: CoreTypeDeviceAttribute,
		Parent:      parent,
		Name:        name,
		Value:       value,
		Props:       make(map[string]DeviceAttributeProperty, 4),
	}, nil
}

// NewDeviceAttributeProperty Creates Component
// of a Device Attribute, representing a single value or measurement
// $implementation/fw -- where status will be a regular attribute
//  $DeviceAttr / Property
func NewDeviceAttributeProperty(parent, name, value string) (DeviceAttributeProperty, error) {
	return DeviceAttributeProperty{
		ID:          NewEID(),
		ElementType: CoreTypeDeviceAttributeProperty,
		Parent:      parent,
		Name:        name,
		Value:       value,
		Props:       make(map[string]DeviceAttributePropertyProperty, 4),
	}, nil
}

// NewDeviceAttributePropertyProperty Creates Component
// of a Device Attribute, representing a single value or measurement
// $implementation/fw/version -- where status will be a regular attribute
//  $DeviceAttr / Property / Property
func NewDeviceAttributePropertyProperty(parent, name, value string) (DeviceAttributePropertyProperty, error) {
	return DeviceAttributePropertyProperty{
		ID:          NewEID(),
		ElementType: CoreTypeDeviceAttributePropertyProperty,
		Parent:      parent,
		Name:        name,
		Value:       value,
	}, nil
}

// NewDeviceNode Creates Component
func NewDeviceNode(parent, name string) (DeviceNode, error) {
	return DeviceNode{
		ID:          NewEID(),
		ElementType: CoreTypeDeviceNode,
		Parent:      parent,
		Name:        name,
		Attrs:       make(map[string]DeviceNodeAttribute, 4),
		Props:       make(map[string]DeviceNodeProperty, 4),
	}, nil
}

// NewDeviceNodeAttribute Creates Component
func NewDeviceNodeAttribute(parent, name, value string) (DeviceNodeAttribute, error) {
	return DeviceNodeAttribute{
		ID:          NewEID(),
		ElementType: CoreTypeDeviceNodeAttribute,
		Parent:      parent,
		Name:        name,
		Value:       value,
	}, nil
}

// NewDeviceNodeProperty Creates Component
func NewDeviceNodeProperty(parent, name, value string) (DeviceNodeProperty, error) {
	return DeviceNodeProperty{
		ID:          NewEID(),
		ElementType: CoreTypeDeviceNodeProperty,
		Parent:      parent,
		Name:        name,
		Value:       value,
		Attrs:       make(map[string]DeviceNodePropertyAttribute, 4),
	}, nil
}

// NewDeviceNodePropertyAttribute Creates Component
func NewDeviceNodePropertyAttribute(parent, name, value string) (DeviceNodePropertyAttribute, error) {
	return DeviceNodePropertyAttribute{
		ID:          NewEID(),
		ElementType: CoreTypeDeviceNodePropertyAttribute,
		Parent:      parent,
		Name:        name,
		Value:       value,
	}, nil
}

// EntityFinder finder utility for Networks
// based on Network Entity
type EntityFinder interface {
	FindOrCreateDeviceAttributePropertyProperty(deviceName, AttributeName, propertyName, pPropertyName string) *DeviceAttributePropertyProperty
	FindOrCreateDeviceAttributeProperty(deviceName, AttributeName, propertyName string) *DeviceAttributeProperty
	FindOrCreateDeviceAttribute(deviceName, attributeName string) *DeviceAttribute
	FindOrCreateDevice(deviceName string) *Device
	FindOrCreateDeviceNodePropertyAttribute(deviceName, nodeName, propertyName string) *DeviceNodeProperty
	FindOrCreateDeviceNodeProperty(deviceName, nodeName, propertyName string) *DeviceNodeProperty
	FindOrCreateDeviceNodeAttribute(deviceName, nodeName, propertyName string) *DeviceNodeProperty
	FindOrCreateDeviceNode(deviceName, nodeName string) *DeviceNode
}
