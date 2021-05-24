package deviceCore

import (
	"fmt"

	"github.com/go-kit/kit/log/level"
)

/*
	Defines the site collection of network models
*/

// Network contains all known devices in application
type Network struct {
	ID          EID
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
	DeviceNetworks map[string]Network
	Broadcasts     []Broadcast
	Firmwares      []Firmware
	Schedules      map[string]Schedule
}

// NewNetworks Creates Component
func NewSiteNetworks(siteName, siteTitle string, networks []string, firmwares []Firmware, schedules map[string]Schedule) (*SiteNetworks, error) {
	var err error

	siteNetworks = SiteNetworks{
		ID:             NewEID(),
		ElementType:    CoreTypeSiteNetworks,
		SiteName:       siteName,
		Title:          siteTitle,
		Names:          networks,
		DeviceNetworks: make(map[string]Network, len(networks)+1),
		Broadcasts:     []Broadcast{},
		Firmwares:      firmwares,
		Schedules:      schedules,
	}

	for _, nName := range networks {
		siteNetworks.DeviceNetworks[nName] = newNetwork(nName, nName)
	}

	return &siteNetworks, err
}

// newNetwork Creates Component
func newNetwork(title, name string) Network {
	hn := Network{
		ID:          NewEID(),
		ElementType: CoreTypeNetwork,
		Title:       title,
		Name:        name,
		Devices:     make(map[string]Device, 16),
	}

	siteNetworks.DeviceNetworks[name] = hn

	return hn
}

// EntityFinder finder utility for Networks
// based on Network Entity
type EntityFinder interface {
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

// Apply device updates to Network
func (hn *Network) apply(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "apply() called")

	// ensure this device is in our network
	_, ok := hn.Devices[string(dm.DeviceID)]
	if !ok {
		return fmt.Errorf("Device{%s} not found in network={%s}", dm.DeviceID, hn.Name)
	}

	switch dm.HomieType {
	case CoreTypeBroadcast:
		err = hn.handleBroadcast(dm)
		break
	case CoreTypeDeviceAttributePropertyProperty:
		err = hn.handleDeviceAttributePropertyProperty(dm)
		break
	case CoreTypeDeviceAttributeProperty:
		err = hn.handleDeviceAttributeProperty(dm)
		break
	case CoreTypeDeviceAttribute:
		err = hn.handleDeviceAttribute(dm)
		break
	case CoreTypeDeviceNodePropertyAttribute:
		err = hn.handleNodePropertyAttribute(dm)
		break
	case CoreTypeDeviceNodeAttribute:
		err = hn.handleNodeAttribute(dm)
		break
	case CoreTypeDeviceNodeProperty:
		err = hn.handleNodeProperty(dm)
		break
	case CoreTypeDeviceNode:
		err = hn.handleNode(dm)
		break
	case CoreTypeDevice:
		err = hn.handleDevice(dm)
		break
	}

	level.Debug(cdss.logger).Log("event", "apply() completed")
	return err
}

// HandleDeviceAttribute on HomieNetork
func (hn *Network) handleDeviceAttributeProperty(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleDeviceAttributeProperty() called")

	// hd, ok := hn.Devices[string(dm.DeviceID)]
	// if !ok {
	// 	return fmt.Errorf("device [%s] not found.", string(dm.DeviceID))
	// }

	// attr, ok := hd.Attrs[string(dm.AttributeID)]
	// if !ok {
	// 	attr, err := NewDeviceAttributeProperty(string(dm.AttributeID), string(dm.PropertyID), string(dm.Value))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	hd.Attrs[string(dm.AttributeID)] = attr
	// 	return nil
	// }

	// attr.Value = string(dm.Value)

	level.Debug(cdss.logger).Log("event", "handleDeviceAttributeProperty() completed")
	return err
}

// HandleDeviceAttribute on HomieNetork
func (hn *Network) handleDeviceAttributePropertyProperty(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleDeviceAttributePropertyProperty() called")

	// hd, ok := hn.Devices[string(dm.DeviceID)]
	// if !ok {
	// 	return fmt.Errorf("device [%s] not found.", string(dm.DeviceID))
	// }

	// attr, ok := hd.Attrs[string(dm.AttributeID)]
	// if !ok {
	// 	attr, err := NewDeviceAttributePropertyProperty(string(dm.PropertyID), string(dm.PPropertyID), string(dm.Value))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	hd.Attrs[string(dm.AttributeID)] = attr
	// 	return nil
	// }

	// attr.Value = string(dm.Value)

	level.Debug(cdss.logger).Log("event", "handleDeviceAttributePropertyProperty() completed")
	return err
}

// HandleDeviceAttribute on HomieNetork
func (hn *Network) handleDeviceAttribute(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleDeviceAttribute() called")

	// hd, ok := hn.Devices[string(dm.DeviceID)]
	// if !ok {
	// 	return fmt.Errorf("device [%s] not found.", string(dm.DeviceID))
	// }

	// attr, ok := hd.Attrs[string(dm.AttributeID)]
	// if !ok {
	// 	attr, err = NewDeviceAttribute(string(dm.DeviceID), string(dm.AttributeID), string(dm.Value))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	hd.Attrs[string(dm.AttributeID)] = attr
	// 	return nil
	// }

	// attr.Value = string(dm.Value)

	level.Debug(cdss.logger).Log("event", "handleDeviceAttribute() completed")
	return err
}

// handleDevice on HomieNetork
func (hn *Network) handleDevice(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleDevice() called")

	// hd, ok := hn.Devices[string(dm.DeviceID)]
	// if ok {
	// 	return err
	// }

	// hd, err = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
	// if err != nil {
	// 	log.Printf("[ERROR] CoreLogic::handleDevice() -> %v", err.Error())
	// 	return err
	// }

	// hn.Devices[string(dm.DeviceID)] = hd

	level.Debug(cdss.logger).Log("event", "handleDevice() completed")
	return err
}

// HandleNodePropertyAttribute on HomieNetork
func (hn *Network) handleNodePropertyAttribute(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleNodePropertyAttribute() called")

	// hdnp, err := hn.handleNodeProperty(dm)
	// if err != nil {
	// 	return err
	// }

	// attr, ok := hdnp.Attrs[string(dm.AttributeID)]
	// if !ok {
	// 	attr, err := NewDeviceNodePropertyAttribute(string(dm.PropertyID), string(dm.AttributeID), string(dm.Value))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	hdnp.Attrs[string(dm.AttributeID)] = attr
	// }
	// attr.Value = string(dm.Value)

	level.Debug(cdss.logger).Log("event", "handleNodePropertyAttribute() completed")
	return err
}

// HandleNodeProperty on HomieNetork
func (hn *Network) handleNodeProperty(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleNodeProperty() called")

	// hdn, err := hn.handleNode(dm)
	// if err != nil {
	// 	return err
	// }

	// prop, ok := hdn.Props[string(dm.PropertyID)]
	// if !ok {
	// 	prop, err = NewDeviceNodeProperty(string(dm.NodeID), string(dm.PropertyID), string(dm.Value))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	hdn.Props[string(dm.PropertyID)] = prop
	// }
	// prop.Value = string(dm.Value)

	level.Debug(cdss.logger).Log("event", "handleNodeProperty() completed")
	return err
}

// HandleNodeAttribute on HomieNetork
func (hn *Network) handleNodeAttribute(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleNodeAttribute() called")

	// hdn, err := hn.handleNode(dm)
	// if err != nil {
	// 	return err
	// }

	// attr, ok := hdn.Attrs[string(dm.AttributeID)]
	// if !ok {
	// 	attr, err = NewDeviceNodeAttribute(string(dm.NodeID), string(dm.AttributeID), string(dm.Value))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	hdn.Attrs[string(dm.AttributeID)] = attr
	// 	return nil
	// }

	// attr.Value = string(dm.Value)

	level.Debug(cdss.logger).Log("event", "handleNodeAttribute() completed")
	return err
}

func (hn *Network) handleNode(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleNode() called")

	// hd, err := hn.handleDevice(dm)
	// if err != nil {
	// 	return err
	// }

	// hdn, ok := hd.Nodes[string(dm.NodeID)]
	// if !ok {
	// 	hdn, err = NewDeviceNode(string(dm.NodeID), string(dm.NodeID))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	hd.Nodes[string(dm.NodeID)] = hdn
	// }

	level.Debug(cdss.logger).Log("event", "handleNode() completed")
	return err
}

// HandleBroadcast on HomieNetork
func (hn *Network) handleBroadcast(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleBroadcast() called")

	// bc, err := NewBroadcast(string(dm.NetworkID), string(dm.DeviceID), string(dm.AttributeID), string(dm.Value))
	// if err == nil {
	// 	hn.Broadcasts = append(hn.Broadcasts, bc)
	// }

	level.Debug(cdss.logger).Log("event", "handleBroadcast() completed")
	return err
}
