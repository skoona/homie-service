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
	Broadcasts     []Broadcast
	Firmwares      []Firmware
	Schedules      map[string]Schedule // by EID
	DeviceNetworks map[string]Network  // by Device Name
}

// NewNetworks Creates Component
func NewSiteNetworks(siteName, siteTitle string, networks []string, firmwares []Firmware, schedules map[string]Schedule) *SiteNetworks {
	var err error
	level.Debug(cdss.logger).Log("event", "NewSiteNetworks() called")

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
		siteNetworks.DeviceNetworks[nName] = newNetwork(nName, nName)
	}

	level.Debug(cdss.logger).Log("event", "NewSiteNetworks() completed")
	return &siteNetworks
}

// newNetwork Creates Component
func newNetwork(title, name string) Network {
	level.Debug(cdss.logger).Log("event", "newNetwork() called", "title", title, "name", name)
	return Network{
		ID:          NewEID(),
		ElementType: CoreTypeNetwork,
		Title:       title,
		Name:        name,
		Devices:     make(map[string]Device, 16),
	}
}

/* EntityBuilder
 * based on Network Entity
 * manages creation and deletion of components
 */
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

// Apply device updates to Network
func (hn *Network) apply(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "apply() called")

	// ensure this device is in our network
	_, ok := hn.Devices[string(dm.DeviceID)]
	if !ok {
		err = fmt.Errorf("device{%s} not found in network={%s}, should be created", dm.DeviceID, hn.Name)
		level.Warn(cdss.logger).Log("action", err.Error())
	}

	/*
	 * Delete the device if value is nil
	 */
	if ok && string(dm.Value) == "" {
		delete(hn.Devices, string(dm.DeviceID))
		err = fmt.Errorf("device{%s} on network{%s} was deleted since value was nil", dm.DeviceID, hn.Name)
		level.Warn(cdss.logger).Log("action", err.Error())
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

	level.Debug(cdss.logger).Log("event", "apply() completed")
	return err
}

// handleDeviceAttributePropertyProperty on HomieNetork
func (hn *Network) handleDeviceAttributePropertyProperty(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleDeviceAttributePropertyProperty() called")

	dev, found := hn.Devices[string(dm.DeviceID)]
	if !found {
		err = fmt.Errorf("device [%s] not found, will create it", string(dm.DeviceID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		dev = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
		hn.Devices[string(dm.DeviceID)] = dev
	}

	devattr, found := dev.Attrs[string(dm.AttributeID)[1:]]
	if !found {
		err = fmt.Errorf("attribute [%s] not found, will create it", string(dm.AttributeID)[1:])
		level.Warn(cdss.logger).Log("warning", err.Error())
		devattr = NewDeviceAttribute(string(dm.DeviceID), string(dm.AttributeID)[1:], "")
		dev.Attrs[string(dm.AttributeID)[1:]] = devattr
	}

	devattrprop, found := devattr.Props[string(dm.PropertyID)]
	if !found {
		err = fmt.Errorf("property [%s] not found, will create it", string(dm.PropertyID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		devattrprop = NewDeviceAttributeProperty(string(dm.AttributeID)[1:], string(dm.PropertyID), "")
		devattr.Props[string(dm.PropertyID)] = devattrprop
	}

	devattrpropprop, found := devattrprop.Props[string(dm.PPropertyID)]
	if !found {
		err = fmt.Errorf("property pproperty [%s] not found, will create it", string(dm.PPropertyID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		devattrpropprop = NewDeviceAttributePropertyProperty(string(dm.PropertyID), string(dm.PPropertyID), string(dm.Value))
		devattrprop.Props[string(dm.PropertyID)] = devattrpropprop
	}

	level.Debug(cdss.logger).Log("event", "handleDeviceAttributePropertyProperty() completed", "id", devattrpropprop.ID, "name", devattrpropprop.Name)
	return err
}

// handleDeviceAttributeProperty on HomieNetork
func (hn *Network) handleDeviceAttributeProperty(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleDeviceAttributeProperty() called")

	dev, found := hn.Devices[string(dm.DeviceID)]
	if !found {
		err = fmt.Errorf("device [%s] not found, will create it", string(dm.DeviceID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		dev = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
		hn.Devices[string(dm.DeviceID)] = dev
	}

	devattr, found := dev.Attrs[string(dm.AttributeID)[1:]]
	if !found {
		err = fmt.Errorf("attribute [%s] not found, will create it", string(dm.AttributeID)[1:])
		level.Warn(cdss.logger).Log("warning", err.Error())
		devattr = NewDeviceAttribute(string(dm.DeviceID), string(dm.AttributeID)[1:], "")
		dev.Attrs[string(dm.AttributeID)] = devattr
	}

	devattrprop, found := devattr.Props[string(dm.PropertyID)]
	if !found {
		err = fmt.Errorf("attribute property [%s] not found, will create it", string(dm.PropertyID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		devattrprop = NewDeviceAttributeProperty(string(dm.AttributeID)[1:], string(dm.PropertyID), string(dm.Value))
		devattr.Props[string(dm.PropertyID)] = devattrprop
	}

	level.Debug(cdss.logger).Log("event", "handleDeviceAttributeProperty() completed", "id", devattrprop.ID, "name", devattrprop.Name)
	return err
}

// handleDeviceAttribute on HomieNetork
func (hn *Network) handleDeviceAttribute(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleDeviceAttribute() called")

	dev, found := hn.Devices[string(dm.DeviceID)]
	if !found {
		err = fmt.Errorf("device [%s] not found, will create it", string(dm.DeviceID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		dev = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
		hn.Devices[string(dm.DeviceID)] = dev
	}

	devattr, found := dev.Attrs[string(dm.AttributeID)]
	if !found {
		err = fmt.Errorf("attribute [%s] not found, will create it", string(dm.AttributeID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		devattr = NewDeviceAttribute(string(dm.DeviceID), string(dm.AttributeID), string(dm.Value))
		dev.Attrs[string(dm.AttributeID)] = devattr
	}

	level.Debug(cdss.logger).Log("event", "handleDeviceAttribute() completed", "id", devattr.ID, "name", devattr.Name)
	return err
}

// handleDevice on HomieNetork
func (hn *Network) handleDevice(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleDevice() called")

	dev, found := hn.Devices[string(dm.DeviceID)]
	if !found {
		err = fmt.Errorf("device [%s] not found, will create it", string(dm.DeviceID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		dev = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
		hn.Devices[string(dm.DeviceID)] = dev
	}

	level.Debug(cdss.logger).Log("event", "handleDevice() completed", "id", dev.ID, "name", dev.Name)
	return err
}

// handleNodePropertyAttribute on HomieNetork
func (hn *Network) handleNodePropertyAttribute(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleNodePropertyAttribute() called")

	dev, found := hn.Devices[string(dm.DeviceID)]
	if !found {
		err = fmt.Errorf("device [%s] not found, will create it", string(dm.DeviceID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		dev = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
		hn.Devices[string(dm.DeviceID)] = dev
	}

	node, found := dev.Nodes[string(dm.NodeID)]
	if !found {
		err = fmt.Errorf("node [%s] not found, will create it", string(dm.NodeID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		node = NewDeviceNode(string(dm.DeviceID), string(dm.NodeID))
		dev.Nodes[string(dm.NodeID)] = node
	}

	nodeprop, found := node.Props[string(dm.PropertyID)]
	if !found {
		err = fmt.Errorf("property [%s] not found, will create it", string(dm.PropertyID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		nodeprop = NewDeviceNodeProperty(string(dm.NodeID), string(dm.PropertyID), "")
		node.Props[string(dm.PropertyID)] = nodeprop
	}

	nodepropattr, found := nodeprop.Attrs[string(dm.AttributeID)]
	if !found {
		err = fmt.Errorf("property attribute [%s] not found, will create it", string(dm.AttributeID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		nodepropattr = NewDeviceNodePropertyAttribute(string(dm.PropertyID), string(dm.AttributeID), string(dm.Value))
		nodeprop.Attrs[string(dm.AttributeID)] = nodepropattr
	}

	level.Debug(cdss.logger).Log("event", "handleNodePropertyAttribute() completed", "id", nodepropattr.ID, "name", nodepropattr.Name)
	return err
}

// HandleNodeProperty on HomieNetork
func (hn *Network) handleNodeProperty(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleNodeProperty() called")

	dev, found := hn.Devices[string(dm.DeviceID)]
	if !found {
		err = fmt.Errorf("device [%s] not found, will create it", string(dm.DeviceID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		dev = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
		hn.Devices[string(dm.DeviceID)] = dev
	}

	node, found := dev.Nodes[string(dm.NodeID)]
	if !found {
		err = fmt.Errorf("node [%s] not found, will create it", string(dm.NodeID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		node = NewDeviceNode(string(dm.DeviceID), string(dm.NodeID))
		dev.Nodes[string(dm.NodeID)] = node
	}

	nodeprop, found := node.Props[string(dm.PropertyID)]
	if !found {
		err = fmt.Errorf("property [%s] not found, will create it", string(dm.PropertyID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		nodeprop = NewDeviceNodeProperty(string(dm.NodeID), string(dm.PropertyID), string(dm.Value))
		node.Props[string(dm.PropertyID)] = nodeprop
	}

	level.Debug(cdss.logger).Log("event", "handleNodeProperty() completed", "id", nodeprop.ID, "name", nodeprop.Name)
	return err
}

// HandleNodeAttribute on HomieNetork
func (hn *Network) handleNodeAttribute(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleNodeAttribute() called")

	dev, found := hn.Devices[string(dm.DeviceID)]
	if !found {
		err = fmt.Errorf("device [%s] not found, will create it", string(dm.DeviceID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		dev = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
		hn.Devices[string(dm.DeviceID)] = dev
	}

	node, found := dev.Nodes[string(dm.NodeID)]
	if !found {
		err = fmt.Errorf("node [%s] not found, will create it", string(dm.NodeID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		node = NewDeviceNode(string(dm.DeviceID), string(dm.NodeID))
		dev.Nodes[string(dm.NodeID)] = node
	}

	nodeattr, found := node.Attrs[string(dm.AttributeID)]
	if !found {
		err = fmt.Errorf("attribute [%s] not found, will create it", string(dm.AttributeID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		nodeattr = NewDeviceNodeAttribute(string(dm.NodeID), string(dm.AttributeID), string(dm.Value))
		node.Attrs[string(dm.AttributeID)] = nodeattr
	}

	level.Debug(cdss.logger).Log("event", "handleNodeAttribute() completed", "id", nodeattr.ID, "name", nodeattr.Name)
	return err
}

// handleNode on HomieNetork
func (hn *Network) handleNode(dm DeviceMessage) error {
	var err error
	level.Debug(cdss.logger).Log("event", "handleNode() called")

	dev, found := hn.Devices[string(dm.DeviceID)]
	if !found {
		err = fmt.Errorf("device [%s] not found, will create it", string(dm.DeviceID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		dev = NewDevice(string(dm.NetworkID), string(dm.DeviceID))
		hn.Devices[string(dm.DeviceID)] = dev
	}

	node, found := dev.Nodes[string(dm.NodeID)]
	if !found {
		err = fmt.Errorf("node [%s] not found, will create it", string(dm.NodeID))
		level.Warn(cdss.logger).Log("warning", err.Error())
		node = NewDeviceNode(string(dm.DeviceID), string(dm.NodeID))
		dev.Nodes[string(dm.NodeID)] = node
	}

	level.Debug(cdss.logger).Log("event", "handleNode() completed", "id", node.ID, "name", node.Name)
	return err
}

// func RemoveIndexFromSlice(slice []Broadcast, index int) []Broadcast {
// 	return append(slice[:index], slice[index+1:]...)
// }

// HandleBroadcast on HomieNetork
func (hn *Network) handleBroadcast(dm DeviceMessage) error {
	var err error
	var found bool
	var bcn Broadcast
	var index int

	level.Debug(cdss.logger).Log("event", "handleBroadcast() called")

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
		level.Warn(cdss.logger).Log("action", err.Error())
	} else { // add it blindly
		err = fmt.Errorf("broadcast [%s] not found, will create it", string(dm.AttributeID))
		level.Warn(cdss.logger).Log("action", err.Error())
		bcn = NewBroadcast(string(dm.NetworkID), string(dm.AttributeID), string(dm.PropertyID), string(dm.Value))
		siteNetworks.Broadcasts = append(siteNetworks.Broadcasts, bcn)
	}

	level.Debug(cdss.logger).Log("event", "handleBroadcast() completed", "id", bcn.ID, "topic", bcn.Topic, "value", bcn.Value)
	return err
}
