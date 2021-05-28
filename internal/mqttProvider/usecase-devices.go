package mqttProvider

/**
  mqttProvider/usecase-devices.go:

  Handle Interaction with deviceSource
  - Send network state messages
  - Listen for network delete messages
  - store/update network traffic to repository
*/
import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	dss "github.com/skoona/homie-service/internal/deviceSource"
	"strings"
)

type (
	deviceStream struct {
		notifyChannel chan dc.QueueMessage // in
		publishChannel chan dc.QueueMessage // in
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

func (s *deviceStream) GetPublishChannel() chan dc.QueueMessage {
	level.Debug(s.logger).Log("method", "GetPublishChannel()")
	if s.publishChannel == nil {
		s.publishChannel = make(chan dc.QueueMessage, 120)
		establishPublishing(s.publishChannel, s.logger)
	}
	return s.publishChannel
}
func (s *deviceStream) GetNotifyChannel() chan dc.QueueMessage {
	level.Debug(s.logger).Log("method", "GetNotifyChannel()")
	if s.notifyChannel == nil {
		s.notifyChannel = make(chan dc.QueueMessage, 200)
	}
	return s.notifyChannel
}

/**
 * establishPublishing()
*/
func establishPublishing(pubChan chan dc.QueueMessage, tlog log.Logger) {
	// Listen on incoming channel for device delete messages
	go func(consumer chan dc.QueueMessage) {
		level.Debug(tlog).Log("event", "establishPublishing(gofunc) called")
		for msg := range consumer { // read until closed

			// DO PUBLISHING WORK HERE
			// DO Device DELETES HERE

			level.Debug(tlog).Log("method", "establishPublishing(gofunc)", "queue depth", len(consumer), "topic", msg.Topic())
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
	if strings.Contains(msg.Topic(), "/$state") &&
		string(msg.Payload()) == "ready" {
		trigger = true
	} else if strings.Contains(msg.Topic(), "$implementation/ota/enabled") &&
		string(msg.Payload()) == "true" {
		trigger = true
	} else if strings.Contains(msg.Topic(), "$implementation/ota/status") {
		trigger = true
	}

	if trigger {
		if otahandler.notifyChannel != nil {
			otahandler.notifyChannel <- msg
		}
	}
	// plus always send to source
	if dStream.notifyChannel != nil {
		dStream.notifyChannel <- msg
	}
}
