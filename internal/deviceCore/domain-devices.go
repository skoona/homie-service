package deviceCore

import (
	"crypto/md5"
	"fmt"
)

/**
	deviceCore/domain-devices.go

The design goal for this file is:
	* Establish the core Device data model

*/

// Device represent a physical device
type Device struct {
	ID          string
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
	ID          string
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
	ID          string
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
	ID          string
	ElementType CoreType
	Parent      string
	Name        string
	Value       string
}

// Node representing the capabilities or features of device
type DeviceNode struct {
	ID          string
	ElementType CoreType
	Parent      string
	Name        string
	Attrs       map[string]DeviceNodeAttribute
	Props       map[string]DeviceNodeProperty
}

// Attribute used by Devices, Nodes, and Properties.  Used to describe format
type DeviceNodeAttribute struct {
	ID          string
	ElementType CoreType
	Parent      string
	Name        string
	Value       string
}

// NodeProperty of a Node, representing a single value or measurement
type DeviceNodeProperty struct {
	ID          string
	ElementType CoreType
	Parent      string
	Name        string
	Value       string
	Attrs       map[string]DeviceNodePropertyAttribute
}

// DeviceNodePropertyAttribute used by Devices, Nodes, and Properties.  Used to describe format
type DeviceNodePropertyAttribute struct {
	ID          string
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
func NewDevice(parent, name string) Device {
	return Device{
		ID:          fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s.%s", parent, name)))),
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
		ID:         fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s.%s", parent, name)))),
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
		ID:          fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s.%s", parent, name)))),
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
		ID:          fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s.%s", parent, name)))),
		ElementType: CoreTypeDeviceAttributePropertyProperty,
		Parent:      parent,
		Name:        name,
		Value:       value,
	}
}

// NewDeviceNode Creates Component
func NewDeviceNode(parent, name string) DeviceNode {
	return DeviceNode{
		ID:          fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s.%s", parent, name)))),
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
		ID:          fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s.%s", parent, name)))),
		ElementType: CoreTypeDeviceNodeAttribute,
		Parent:      parent,
		Name:        name,
		Value:       value,
	}
}

// NewDeviceNodeProperty Creates Component
func NewDeviceNodeProperty(parent, name, value string) DeviceNodeProperty {
	return DeviceNodeProperty{
		ID:          fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s.%s", parent, name)))),
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
		ID:          fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s.%s", parent, name)))),
		ElementType: CoreTypeDeviceNodePropertyAttribute,
		Parent:      parent,
		Name:        name,
		Value:       value,
	}
}
