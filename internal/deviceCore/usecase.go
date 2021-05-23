package deviceCore

import (
	"errors"

	"github.com/go-kit/kit/log/level"
)

/*
  	deviceCore/service.go:

	Build Device Network Inventory
	Provide Interaction to CLI
	Provide Interaction fo HTTP
	Consider Scheduler Feature
*/

// The active Service
var (
	em               coreService
	fromDeviceSource chan DeviceMessage // in
	toDeviceSource   chan DeviceMessage // out
)

// EventHandler manages mqtt input messages
type EventHandler interface {
	CreateEvent(coreType CoreType, otaTrigger, retain bool, qos byte, networkName, deviceName, nodeName, propertyName, attributeName, value []byte, topic string) (*DeviceMessage, error)
	ApplyEvent(dm *DeviceMessage) error
	GetCorePublishingChannel() (chan DeviceMessage, error)
	GetCoreSubscribingChannel() (chan DeviceMessage, error)
}

func (em *coreService) CreateEvent(coreType CoreType, otaTrigger, retain bool, qos byte, networkName, deviceName, nodeName, propertyName, attributeName, value []byte, topic string) (*DeviceMessage, error) {
	level.Debug(em.logger).Log("method", "CreateEvent() called")
	return nil, nil
}
func (em *coreService) ApplyEvent(dm *DeviceMessage) error {
	level.Debug(em.logger).Log("method", "ApplyEvent() called")
	return nil
}

// caller sends on this channel
func (em *coreService) GetCorePublishingChannel() (chan DeviceMessage, error) {
	level.Debug(em.logger).Log("method", "GetCorePublishingChannel() called")
	var err error
	if fromDeviceSource == nil {
		fromDeviceSource = make(chan DeviceMessage, 256) // averages 120 on startup
		// err = ConsumeFromDeviceSource(fromDeviceSource)
	}

	if nil == fromDeviceSource {
		err = errors.New("create core publishing channel failed")
		level.Error(em.logger).Log("error", err.Error())
		return nil, err
	}
	return fromDeviceSource, err
}

// caller listens on this channel
func (em *coreService) GetCoreSubscribingChannel() (chan DeviceMessage, error) {
	level.Debug(em.logger).Log("method", "GetCoreSubscribingChannel() called")
	var err error
	if toDeviceSource == nil {
		toDeviceSource = make(chan DeviceMessage, 256) // averages 120 on startup
	}

	if nil == toDeviceSource {
		err = errors.New("create core subscribing channel failed")
		level.Error(em.logger).Log("error", err.Error())
		return nil, err
	}
	return toDeviceSource, err
}
