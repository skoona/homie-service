package deviceCore

/*
 *  deviceCore/events.go
 *
 * handle device source Primary input
 */

import (
	"fmt"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

/*
 * DeviceMessage
 * Transposed MQTT Message for Internal Use
 */
type DeviceMessage struct {
	ID          uint16
	HomieType   CoreType
	OTATrigger  bool
	RetainedB   bool
	Qosb        byte
	NetworkID   []byte
	DeviceID    []byte
	NodeID      []byte
	PropertyID  []byte
	PPropertyID []byte
	AttributeID []byte
	TopicS      string
	Value       []byte
}

// QueueMessage maps a mqtt.Message into this domain entity
type QueueMessage interface {
	Topic() string
	Payload() []byte
	MessageID() uint16
	Retained() bool
	Qos() byte
}

/*
 * Device Management Methods
 */
type DeviceEventIntf interface {
	String() string
	Schedulable() bool
	Settable() bool
	OTAactive() bool
	OTAComplete() bool
	Broadcast() bool
	Parts() []string
	PartsLen() int
}

/*
   String ToString
   - utility funciton
*/
func (dm *DeviceMessage) String() string {
	return fmt.Sprintf("id=%06d retained=%-5t device=%-16s node=%-16s property=%-16s pProperty=%-16s attr=%-16s value=%s",
		dm.ID, dm.RetainedB, dm.DeviceID, dm.NodeID, dm.PropertyID, dm.PPropertyID, dm.AttributeID, dm.Value)
}

// Schedulable()
// also -- $stats/uptime
func (dm *DeviceMessage) Schedulable() bool {
	level.Debug(cdss.logger).Log("DeviceMessage", "Schedulable()")
	res := false
	for _, keys := range []string{"$state", "$online", "$fw", "$implementation", "uptime"} {
		if strings.Contains(dm.TopicS, keys) {
			res = true
			break
		}
	}
	return res
}

// Settable() determine is property is settable
func (dm *DeviceMessage) Settable() bool {
	level.Debug(cdss.logger).Log("DeviceMessage", "Settable()")
	return strings.HasSuffix(dm.TopicS, "set")
}

// OTAactive() determines if device is downloading firmware
func (dm *DeviceMessage) OTAactive() bool {
	level.Debug(cdss.logger).Log("DeviceMessage", "OTAactive()")
	return strings.Contains(dm.TopicS, "$implementation/ota/status") &&
		strings.HasPrefix(string(dm.Value), "206")
}

// OTAComplete determines if device has accepted new firmware
// Filtered on create of event message, use Topic to bypass Filtering
func (dm *DeviceMessage) OTAComplete() bool {
	level.Debug(cdss.logger).Log("DeviceMessage", "OTAComplete()")
	return strings.Contains(dm.TopicS, "$implementation/ota/status") &&
		strings.HasPrefix(string(dm.Value), "200")
}

// Broadcast() determines if this is a Homie Broadcast message
func (dm *DeviceMessage) Broadcast() bool {
	level.Debug(cdss.logger).Log("DeviceMessage", "Broadcast()")
	return dm.Parts()[1] == "$broadcast"
}

// Parts() returns the individual parts of the original MQTT message
func (dm *DeviceMessage) Parts() []string {
	level.Debug(cdss.logger).Log("DeviceMessage", "Parts()")
	return strings.Split(dm.TopicS, "/")
}

// PartsLen() returns nuber of parts in Topic
func (dm *DeviceMessage) PartsLen() int {
	level.Debug(cdss.logger).Log("DeviceMessage", "PartsLen()")
	return len(dm.Parts())
}

/*
 * QueueMessage implementation for DeviceMessage
 */
func (dm DeviceMessage) Topic() string {
	return dm.TopicS
}
func (dm DeviceMessage) Payload() []byte {
	return dm.Value
}
func (dm DeviceMessage) MessageID() uint16 {
	return dm.ID
}
func (dm DeviceMessage) Retained() bool {
	return dm.RetainedB
}
func (dm DeviceMessage) Qos() byte {
	return dm.Qosb
}

// X/$B
func homieBroadcast(parts []string) bool {
	if nil != parts && len(parts) >= 3 && strings.HasPrefix(parts[1], "$broad") {
		return true
	}
	return false
}

// X/D/$A
func deviceAttribute(parts []string) bool {
	if nil != parts && len(parts) == 3 && strings.HasPrefix(parts[2], "$") {
		return true
	}
	return false
}

// X/D/$A/P
func deviceAttributeProperty(parts []string) bool {
	if nil != parts && len(parts) == 4 && strings.HasPrefix(parts[2], "$") {
		return true
	}
	return false
}

// X/D/$A/P/P
func deviceAttributePropertyProperty(parts []string) bool {
	if nil != parts && len(parts) == 5 && strings.HasPrefix(parts[2], "$") {
		return true
	}
	return false
}

// X/D/N/$A
func nodeAttribute(parts []string) bool {
	if nil != parts && len(parts) == 4 && strings.HasPrefix(parts[3], "$") {
		return true
	}
	return false
}

// X/D/N/P/$A
func nodePropertyAttribute(parts []string) bool {
	if nil != parts && len(parts) == 5 && strings.HasPrefix(parts[4], "$") {
		return true
	}
	return false
}

// X/D/N/P
// X/D/N/P/{set}
func nodeProperty(parts []string) bool {
	if nil != parts && len(parts) == 4 && !strings.HasPrefix(parts[3], "$") {
		return true
	} else if nil != parts && len(parts) == 5 && strings.HasPrefix(parts[4], "set") {
		return true
	}
	return false
}

/*
 * homieDeviceFilter
 *
 * Base / Device / Node / Properties / Attributes
 * Device  -> The physical thing
 * Devices -> A device can expose multiple nodes
 * Nodes   ->  A node can have multiple properties
 * Properties ->  Devices, nodes and properties have (multiple) specific attributes characterizing them.
 *				  Attributes are represented by topic identifier starting with $
 * homie / device ID / $state : last      (LWT)
 * homie / device ID / $device-attribute
 * homie / device ID / $device-attribute / property ID
 * homie / device ID / $device-attribute / property ID / property ID
 * homie / device ID / node ID / $node-attribute
 * homie / device ID / node ID / property ID / $property-attribute
 * homie / device ID / node ID / property ID / set     {property set command}
 * homie / device ID / node ID / property ID
 *
 * Ignore h/d/n/p/set messages
 * Ignore firmware in progress messages h/d/$implementation/ota/firmware
 */
func homieDeviceFilter(attributeID []byte, parts []string) error {
	if len(attributeID) == 0 { // Ignore nil attributes (since they can be valid)
		return nil
	}

	if len(parts) == 5 && strings.HasPrefix(parts[4], "set") { // ignore commands to set properties
		return fmt.Errorf("filtering Homie Set commands from: %s", strings.Join(parts[1:], "/"))
	}

	filters := []string{"_$implementation/ota/firmware"} // "$implementation/ota/status",
	source := string(attributeID)

	for _, item := range filters {
		if item == source {
			return fmt.Errorf("filtering Homie Attributes: %s, from: %s", source, strings.Join(parts[1:], "/"))
		}
	}

	if len(parts) >= 5 {
		source = strings.Join(parts[2:(len(parts)-1)], "/")
		for _, item := range filters {
			if item == source {
				return fmt.Errorf("filtering Homie Attributes: %s, from: %s", source, strings.Join(parts[1:], "/"))
			}
		}
	}

	return nil
}

/**
 * buildDeviceMessage()
 *
 *  Create a New DeviceMessage and initializes it.
 */
func buildDeviceMessage(topic string, payload []byte, idCounter uint16, retained bool, qos byte) (DeviceMessage, error) {
	logger := log.With(cdss.logger, "method", "buildDeviceMessage")

	var deviceID, nodeID, propertyID, attributeID, networkID, propertyPropertyID []byte
	var dm DeviceMessage
	var typeIntHomie CoreType

	nodeID = []byte{}
	propertyID = []byte{}
	propertyPropertyID = []byte{}
	parts := strings.Split(topic, "/")

	if homieBroadcast(parts) {
		typeIntHomie = CoreTypeBroadcast
		deviceID = []byte(parts[1])
		attributeID = []byte(parts[2])
		propertyID = []byte(strings.Join(parts[3:], "/"))
		level.Debug(logger).Log("ID", idCounter, "HomieBroadcast", attributeID, "level", propertyID)

	} else if deviceAttributePropertyProperty(parts) { // look for OTATrigger(bool) on $state = Init
		typeIntHomie = CoreTypeDeviceAttributePropertyProperty
		deviceID = []byte(parts[1])
		propertyID = []byte(parts[3])
		propertyPropertyID = []byte(parts[4])
		attributeID = []byte(fmt.Sprintf("_%s", parts[2])) // LevelDB bug writing Bucket with $ prefix
		level.Debug(logger).Log("ID", idCounter, "DeviceAttributePropertyProperty", attributeID, "DeviceID", deviceID)

	} else if deviceAttributeProperty(parts) { // look for OTATrigger(bool) on $state = Init
		typeIntHomie = CoreTypeDeviceAttributeProperty
		deviceID = []byte(parts[1])
		propertyID = []byte(parts[3])
		attributeID = []byte(fmt.Sprintf("_%s", parts[2])) // LevelDB bug writing Bucket with $ prefix
		level.Debug(logger).Log("ID", idCounter, "DeviceAttributeProperty", attributeID, "DeviceID", deviceID)

	} else if deviceAttribute(parts) { // look for OTATrigger(bool) on $state = Init
		typeIntHomie = CoreTypeDeviceAttribute
		deviceID = []byte(parts[1])
		attributeID = []byte(parts[2])
		level.Debug(logger).Log("ID", idCounter, "DeviceAttribute", attributeID, "DeviceID", deviceID)

	} else if nodeAttribute(parts) {
		typeIntHomie = CoreTypeDeviceNodeAttribute
		deviceID = []byte(parts[1])
		nodeID = []byte(parts[2])
		attributeID = []byte(parts[3])
		level.Debug(logger).Log("ID", idCounter, "NodeAttribute", attributeID, "NodeID", nodeID)

	} else if nodePropertyAttribute(parts) {
		typeIntHomie = CoreTypeDeviceNodePropertyAttribute
		deviceID = []byte(parts[1])
		nodeID = []byte(parts[2])
		propertyID = []byte(parts[3])
		attributeID = []byte(parts[4])
		level.Debug(logger).Log("ID", idCounter, "NodePropertyAttribute", attributeID, "PropertyID", propertyID, "NodeID", nodeID)

	} else if nodeProperty(parts) {
		typeIntHomie = CoreTypeDeviceNodeProperty
		deviceID = []byte(parts[1])
		nodeID = []byte(parts[2])
		propertyID = []byte(parts[3])
		attributeID = []byte{}
		level.Debug(logger).Log("ID", idCounter, "NodeProperty", propertyID, "NodeID", nodeID)

	} else {
		return DeviceMessage{}, fmt.Errorf("invalid Homie message structure: %s", topic)
	} // Done

	networkID = []byte(parts[0])
	if err := homieDeviceFilter(attributeID, parts); err != nil {
		return DeviceMessage{}, err
	}

	dm = DeviceMessage{ // effectively copy the data
		HomieType:   typeIntHomie,
		ID:          idCounter,
		RetainedB:   retained,
		Qosb:        qos,
		NetworkID:   networkID,
		DeviceID:    deviceID,
		NodeID:      nodeID,
		PropertyID:  propertyID,
		PPropertyID: propertyPropertyID,
		AttributeID: attributeID,
		TopicS:      topic,
		Value:       payload,
	}

	return dm, nil
}
