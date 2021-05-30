package mqttProvider

/**
  mqttProvider/usecase-ota.go:

  Handle Interaction with Scheduler
  - Send OTA Triggers
  - Execute OTA to device
  - Watch OTA Progress
  - Send Status
*/

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	sch "github.com/skoona/homie-service/internal/deviceScheduler"
)

type (
	otaStream struct {
		notifyChannel  chan dc.QueueMessage
		publishChannel chan dc.QueueMessage
		logger         log.Logger
	}
)

var (
	otastream *otaStream
)

func NewOTAStream(plog log.Logger) sch.OTAInteractor {
	otastream = &otaStream{
		logger: log.With(plog, "service", "OTAInteractor"),
	}
	return otastream
}
func (s *otaStream) EnableTriggers() chan dc.QueueMessage {
	if s.notifyChannel == nil {
		s.notifyChannel = make(chan dc.QueueMessage, 120)
	}
	return s.notifyChannel
}
func (s *otaStream) EnableNotificationsFor(networkName, deviceName string, enabledOrDisable bool) error {
	err := handleOTAMessages(networkName, deviceName, enabledOrDisable)
	if err != nil {
		level.Error(s.logger).Log("error", err.Error())
	}
	return err
}
func (s *otaStream) OtaPublish(otaMessage dc.QueueMessage) {
	if s.publishChannel == nil {
		s.publishChannel = make(chan dc.QueueMessage, 120)
		publishOTAMessages(s.publishChannel, s.logger) // start receiver
	}
	s.publishChannel <- otaMessage
}

/*
 * otaResponses
 * Source OTA Responses to Scheduler
 */
var otaResponses mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	// msg.Payload()[0] = byte{0}
	if otastream.notifyChannel != nil {
		otastream.notifyChannel <- msg
	} else {
		level.Error(logger).Log("error", "ota notification channel offline", "topic", msg.Topic())
	}
}

/**
 * ConsumeOTAMessages
 * Handles OTA Streaming responses
 */
func handleOTAMessages(network, name string, enabled bool) error {
	var err error
	if enabled {
		err = WatchOTAProgress(network, name)
	} else {
		err = UnWatchOTAProgress(network, name)
	}

	if nil != err {
		level.Error(logger).Log("event", "Subscribe Failed", "error", err.Error())
	}
	level.Debug(logger).Log("ConsumeOTAMessages Completed", enabled, "network", network, "name", name)

	return err
}

// WatchOTAProgress for Scheduler
func WatchOTAProgress(network, device string) error {
	topic := fmt.Sprintf("%s/%s/$implementation/ota/status", network, device)
	token := subWithHandler(topic, otaResponses) // OTA Responses
	return token.Error()
}

// UnWatchOTAProgress for Scheduler
func UnWatchOTAProgress(network, device string) error {
	topic := fmt.Sprintf("%s/%s/$implementation/ota/status", network, device)
	token := client.Unsubscribe(topic)
	token.Wait()
	return token.Error()
}

func publishOTAMessages(publisher chan dc.QueueMessage, plog log.Logger) {
	go func(publishChan chan dc.QueueMessage) {
		level.Debug(plog).Log("event", "publishOTAMessages(gofunc) called")
		for otaMessage := range publishChan { // read until closed
			publish(otaMessage.Topic(), otaMessage.Payload(), otaMessage.Retained(), otaMessage.Qos())
			level.Debug(plog).Log("method", "publishOTAMessages(gofunc)", "queue depth", len(publishChan), "device", otaMessage.Topic())
		}
		level.Debug(plog).Log("method", "publishOTAMessages()", "event", "Completed")
	}(publisher)
}
