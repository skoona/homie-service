package mqttProvider

/**
  mqttProvider/usecase-devices.go:

  Handle Interaction with deviceSource
  - Send network state messages
  - Listen for network delete messages
  - store/update network traffic to repository
  - Implements the dss.StreamProvider to deviceSource
  - Implements the
*/
import (
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	dss "github.com/skoona/homie-service/internal/deviceSource"
)

type (
	deviceStream struct {
		notifyChannel  chan dc.DeviceMessage // in
		publishChannel chan dc.Device // in
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

func (s *deviceStream) GetPublishChannel() chan dc.Device {
	level.Debug(s.logger).Log("method", "GetPublishChannel()")
	if s.publishChannel == nil {
		s.publishChannel = make(chan dc.Device, 120)
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
func establishPublishing(pubChan chan dc.Device, tlog log.Logger) {
	// Listen on incoming channel for device delete messages
	go func(consumer chan dc.Device) {
		level.Debug(tlog).Log("event", "establishPublishing(gofunc) called")
		for msg := range consumer { // read until closed

			for _, topic := range dc.TopicsFromDevice(msg) {
				publish(topic, []byte{}, false, 0)
				time.Sleep(5 * time.Millisecond)
				level.Debug(tlog).Log("publishing to", msg.Name, "Topic", topic)
			}

			level.Debug(tlog).Log("method", "establishPublishing(gofunc)", "queue depth", len(consumer), "deviceID", msg.ID)
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
	// -- sknSensors/GarageMonitor/$fw/{name|version|checksum} value
	var trigger bool = false
	if strings.Contains(msg.Topic(), "$fw/") {
		trigger = true
	}

	dm := dStream.CreateQueueDeviceMessage(msg)

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
