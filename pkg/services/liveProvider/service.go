package liveProvider

/**
  liveProvider/service.go:

  Generate MQTT DeviceMessages

  mqttclient

  The design goal for this file is:
	* Establish a un-secured connection to local MQTT broker
	* Discover which Homie network prefixes are being used
	* Subscribe to all found Homie networks
	* Listen for network messages and forward them unaltered on `mqttChannel`
	* Listen to `mqttPublish` channel and publish outgoing as needed
	* Auto reconnect when connection is lost.
	* Allow for the replacement or deletion of both.

*/

import (
	"fmt"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	stringset "github.com/jjeffery/stringset"
	dc "github.com/skoona/homie-service/pkg/services/deviceCore"
	dss "github.com/skoona/homie-service/pkg/services/deviceSource"
	cc "github.com/skoona/homie-service/pkg/utils"
)

/*
 * NeworkDiscovery()
 * Discovers available Homie networks
 */
func DiscoveredNetworks() []string {
	return nNetworks.Values()
}

/*
 * networksHandler
 * Handles '+/+/$name' a discovery technique which list HomieDevices on any network prefix
 */
var networksHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	parts := strings.Split(msg.Topic(), "/")
	if !nNetworks.Contains(parts[0]) {
		nNetworks.Add(parts[0])
		level.Debug(logger).Log("Network Discovery", parts[0])
	}
	level.Debug(logger).Log("Network Discovery topic", msg.Topic(), "payload", msg.Payload())
}

/*
 * otaResponses
 * Source OTA Responses to Scheduler
 */
var otaResponses mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	// msg.Payload()[0] = byte{0}
	dm, err := dc.NewQueueMessage(msg)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return
	}
	toOTAService <- dm // send it to channel
}

/*
 * defaultOnMessage
 * Default Incoming Message Handler
 */
var defaultOnMessage mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	dm, err := dc.NewQueueMessage(msg)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return
	}
	toDMService <- dm // send it to channel
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	level.Debug(logger).Log("status", "MQTT Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	level.Debug(logger).Log("status", "lost MQTT connection", "error", err.Error())
}

func publish(topic string, payload string, retain bool, qos byte) {
	token := client.Publish(topic, qos, retain, []byte(payload))
	token.Wait()
	level.Debug(logger).Log("Published topic", topic, "payload", payload)
}

func sub(topic string) mqtt.Token {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	level.Debug(logger).Log("Subscribed to topic", topic)

	return token
}
func subWithHandler(topic string, callback mqtt.MessageHandler) mqtt.Token {
	token := client.Subscribe(topic, 1, callback)
	token.Wait()
	level.Debug(logger).Log("Subscribed to topic", topic)

	return token
}

/**
 * ConsumeOTAMessages
 * Handles OTA Streaming responses
 */
func ConsumeOTAMessages(network, name string, enabled bool) error {
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

// ProduceMQTTMessages to group topic
func ProduceMQTTMessages() error {
	var token mqtt.Token
	networks := DiscoveredNetworks()

	if len(networks) > 0 {
		for _, network := range networks {
			networkTopic := fmt.Sprintf("%s/#", network)
			token = subWithHandler(networkTopic, defaultOnMessage) // Available Topic
		}
	} else {
		token = subWithHandler(config.SubscriptionTopic, defaultOnMessage) // Default Topic
	}
	return token.Error()
}

// RemoveSubscriptions Unsubscribe and shutdown cleanly
func removeSubscriptions() error {
	networks := DiscoveredNetworks()
	if len(networks) > 0 {
		for _, network := range networks {
			networkTopic := fmt.Sprintf("%s/#", network)
			token := client.Unsubscribe(networkTopic)
			token.Wait()
		}
	} else {
		token := client.Unsubscribe(config.SubscriptionTopic)
		token.Wait()
	}

	return nil
}

var (
	config         cc.MQTTConfig
	fromDMService  chan dc.DeviceMessage // in
	toDMService    chan dc.DeviceMessage // out
	toOTAService   chan dc.DeviceMessage // out
	fromOTAService chan dc.DeviceMessage // in
	client         mqtt.Client
	nNetworks      = stringset.New()
	deviceService  dss.Service
	logger         log.Logger
)

// chan mqtt.Message, chan mqtt.Message, chan stor.DeviceMessage
func Start(cfg cc.Config, svc dss.Service) ([]string, error) {
	config = cfg.Mqc
	logger = log.With(cfg.Logger, "pkg", "liveProvider")
	deviceService = svc
	var err error
	level.Debug(logger).Log("event", "Calling Start()", "broker", config.BrokerIP)

	/* Initialize MQTT */
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.BrokerIP, config.BrokerPort))
	opts.SetClientID(config.ClientID)
	opts.SetUsername(config.UserName)
	opts.SetPassword(config.UserPassword)
	opts.SetAutoReconnect(true)
	opts.SetOrderMatters(true)
	if len(config.LwtTopic) > 0 && len(config.LwtMessage) > 0 {
		opts.SetWill(config.LwtTopic, config.LwtMessage, 0, true)
	}
	opts.SetDefaultPublishHandler(defaultOnMessage)
	opts.SetOnConnectHandler(connectHandler)
	opts.SetConnectionLostHandler(connectLostHandler)

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		level.Error(logger).Log("cause", "connect() failed", "error", token.Error())
		panic(token.Error())
	}

	// allow for network discovery
	subWithHandler(config.DiscoveryTopic, networksHandler) // Homie Discovery Topic
	for {
		time.Sleep(2 * time.Second) // delay long enough to collect networks
		if len(DiscoveredNetworks()) >= 1 {
			token := client.Unsubscribe(config.DiscoveryTopic)
			token.Wait()
			break
		}
	}

	// Initialize a Message Channel
	fromDMService, toDMService, err = dss.ChannelsForDMProviders()
	if err != nil {
		level.Error(logger).Log("event", "DM Channels offline", "error", err.Error())
		client.Disconnect(250)
		panic(err.Error())
	}
	fromOTAService, toOTAService, err = dss.ChannelsForOTAProviders()
	if err != nil {
		level.Error(logger).Log("event", "OTA Channels offline", "error", err.Error())
		client.Disconnect(250)
		panic(err.Error())
	}

	ProduceMQTTMessages()

	level.Debug(logger).Log("event", "Start() Completed", "networks discovered", nNetworks.GoString())
	return DiscoveredNetworks(), err
}

func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")
	// Unsubscribe and shutdown cleanly
	removeSubscriptions()

	if nil != toDMService {
		close(toDMService) // only the send should shutdown channels
	}
	if nil != toOTAService {
		close(toOTAService)
	}

	client.Disconnect(250)

	time.Sleep(time.Second)
	level.Debug(logger).Log("event", "Stop() completed")
}
