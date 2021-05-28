package mqttProvider

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kit/kit/log/level"
	sch "github.com/skoona/homie-service/internal/deviceScheduler"
)

/**
  mqttProvider/usecase-ota.go:

  Handle Interaction with Scheduler
  - Send OTA Triggers
  - Execute OTA to device
  - Watch OTA Progress
  - Send Status
*/

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	cc "github.com/skoona/homie-service/internal/utils"
)

type (
	otaHandler struct {
		notifyChannel chan dc.QueueMessage
		publishChannel chan dc.QueueMessage
		logger log.Logger
	}
)

var (
	otahandler *otaHandler
)

func NewOTAHandler(plog log.Logger) sch.OTAInteractor {
	otahandler = &otaHandler{
		logger: log.With(plog, "service", "OTAInteractor"),
	}
	return otahandler
}
func (s *otaHandler) EnableTriggers() chan dc.QueueMessage {
	if s.notifyChannel == nil {
		s.notifyChannel = make(chan dc.QueueMessage, 120)
	}
	return s.notifyChannel
}
func (s *otaHandler) EnableNotificationsFor(networkName, deviceName string, enabledOrDisable bool) error {
	err := handleOTAMessages(networkName, deviceName, enabledOrDisable)
	if err != nil {
		level.Error(s.logger).Log("error", err.Error())
	}
	return err
}
func (s *otaHandler) OtaPublish(otaMessage dc.QueueMessage) {
		if s.publishChannel == nil {
			s.publishChannel = make(chan dc.QueueMessage, 120)
			publishOTAMessages(s.publishChannel) // start receiver
		}
		s.publishChannel <- otaMessage
}

/*
 * otaResponses
 * Source OTA Responses to Scheduler
 */
var otaResponses mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	// msg.Payload()[0] = byte{0}
	if otahandler.notifyChannel != nil {
		otahandler.notifyChannel <- msg
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

func publishOTAMessages(publisher chan dc.QueueMessage) {
	go func(publishChan chan dc.QueueMessage) {
		level.Debug(logger).Log("event", "publishOTAMessages(gofunc) called")
		for otaMessage := range publishChan { // read until closed
			publish(otaMessage.Topic(), otaMessage.Payload(), otaMessage.Retained(), otaMessage.Qos())
			level.Debug(logger).Log("method", "publishOTAMessages(gofunc)", "queue depth", len(consumeChan), "device", otaMessage.Topic())
		}
		level.Debug(logger).Log("method", "publishOTAMessages()", "event", "Completed")
	}(publisher)
}

func establishOTAListener() {
	var err error

	go func(consumeChan chan dc.QueueMessage) {
		level.Debug(logger).Log("event", "consumeNotificationsFromOTAProviders(gofunc) called")
		for msg := range consumeChan { // read until closed

			err := otaChannel.ApplyOTAEvent(msg)
			if err != nil {
				level.Error(logger).Log("method", "consumeNotificationsFromOTAProviders(gofunc)", "error", err.Error())
			}

			level.Debug(logger).Log("method", "consumeNotificationsFromOTAProviders(gofunc)", "consume queue depth", len(consumeChan))
		}
		level.Debug(logger).Log("method", "consumeNotificationsFromOTAProviders()", "event", "Completed")
	}(consumer)

	if err != nil {
		level.Error(logger).Log("event", "OTA respnse channel offline", "error", err.Error())
		client.Disconnect(250)
		panic(err.Error())
	}
}
