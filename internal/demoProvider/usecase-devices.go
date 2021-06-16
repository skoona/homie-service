package demoProvider

/**
  demoProvider/usecase-devices.go:

  Handle Interaction with deviceSource
  - Send network state messages
  - Listen for network delete messages
  - store/update network traffic to repository
  - Implements the dss.StreamProvider to deviceSource
  - Implements the
*/
import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	dss "github.com/skoona/homie-service/internal/deviceSource"
	"strings"
)

type (
	deviceStream struct {
		notifyChannel  chan dc.DeviceMessage // in
		publishChannel chan dc.DeviceMessage // in
		logger         log.Logger
	}
)

var (
	dStream *deviceStream
)

/*
 * DeviceStream Implementation */

func NewStreamProvider(plog log.Logger) dss.StreamProvider {
	dStream = &deviceStream{
		logger: log.With(plog, "service", "StreamProvider"),
	}
	return dStream
}
func (s *deviceStream) ActivateNotifications() chan dc.DeviceMessage {
	enableNetworkTraffic(s.logger)
	return s.GetNotifyChannel()
}
func (s *deviceStream) CreateDemoDeviceMessage(topic string, payload []byte, idCounter uint16, retained bool, qos byte) dc.DeviceMessage {
	level.Debug(s.logger).Log("method", "CreateDemoDeviceMessage() called")
	dm, _ := dc.NewDeviceMessage(topic, payload, idCounter, retained, qos, s.logger)
	return dm
}

func (s *deviceStream) CreateQueueDeviceMessage(qmsg dc.QueueMessage) dc.DeviceMessage {
	level.Debug(s.logger).Log("method", "CreateQueueDeviceMessage() called")
	dm, _ := dc.NewQueueMessage(qmsg, s.logger)
	return dm
}

func (s *deviceStream) GetPublishChannel() chan dc.DeviceMessage {
	level.Debug(s.logger).Log("method", "GetPublishChannel()")
	if s.publishChannel == nil {
		s.publishChannel = make(chan dc.DeviceMessage, 32)
		establishPublishing(s.publishChannel, s.logger)
	}
	return s.publishChannel
}
func (s *deviceStream) GetNotifyChannel() chan dc.DeviceMessage {
	level.Debug(s.logger).Log("method", "GetNotifyChannel()")
	if s.notifyChannel == nil {
		s.notifyChannel = make(chan dc.DeviceMessage, 200)
	}
	return s.notifyChannel
}

/**
 * establishPublishing()
 */
func establishPublishing(pubChan chan dc.DeviceMessage, tlog log.Logger) {
	// Listen on incoming channel for device delete messages
	go func(consumer chan dc.DeviceMessage) {
		level.Debug(tlog).Log("method", "establishPublishing()", "event", "establishPublishing(gofunc) called")
		for msg := range consumer { // read until closed
			publish(msg.Topic(), msg.Payload(), msg.Retained(), msg.Qos())
			level.Debug(tlog).Log("queue depth",len(consumer), "publishing to", msg.DeviceID, "Topic", msg.Topic())
		}
		level.Debug(tlog).Log("method", "establishPublishing()", "event", "Completed")
	}(pubChan)
}

/*
 * defaultOnMessage
 * Default Incoming Message Handler
 */
var defaultOnMessage mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	// if this is a trigger route to scheduler
	// triggers:
	// sknSensors/GarageMonitor/$state ready
	// sknSensors/GarageMonitor/$implementation/ota/enabled true
	var trigger bool = false
	if strings.HasSuffix(msg.Topic(), "$fw/") {
		trigger = true
	}

	dm := dStream.CreateQueueDeviceMessage(msg)

	// firmware load messages require payload to be transformed to smaller message
	if strings.Contains(msg.Topic(), "$implementation/ota/firmware/") {
		dm.Value = []byte( fmt.Sprintf("MessageBytes=%d", len(dm.Value)))
	}

	if trigger {
		if otastream != nil {
			if otastream.notifyChannel != nil {
				otastream.notifyChannel <- dm
			}
		}
	}
	// plus always send to source
	if dStream != nil {
		if dStream.notifyChannel != nil {
			dStream.notifyChannel <- dm
		}
	}
}
