package mqttProvider

/**
  mqttProvider/service.go:

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
	sch "github.com/skoona/homie-service/internal/deviceScheduler"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	stringset "github.com/jjeffery/stringset"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	dss "github.com/skoona/homie-service/internal/deviceSource"
	cc "github.com/skoona/homie-service/internal/utils"
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
 * defaultOnMessage
 * Default Incoming Message Handler
 */
var defaultOnMessage mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	// if this is a trigger route to scheduler
	if otahandler.notifyChannel != nil {
		otahandler.notifyChannel <- msg
	}
	// plus aways send to source
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	level.Debug(logger).Log("status", "MQTT Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	level.Debug(logger).Log("status", "lost MQTT connection", "error", err.Error())
}

func publish(topic string, payload []byte, retain bool, qos byte) {
	token := client.Publish(topic, qos, retain, payload)
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
 * disableNetworkTraffic Unsubscribe and shutdown cleanly
 */
func disableNetworkTraffic() error {
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

/**
 * Activate Subscriptions starting network traffic
 */
func enableNetworkTraffic() error {
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

/**
 * Discover existing Homie Networks
 */
func doNetworkDiscovery() {
	// allow for network discovery
	subWithHandler(config.DiscoveryTopic, networksHandler) // Homie Discovery Topic
	for {
		time.Sleep(3 * time.Second) // delay long enough to collect networks
		if len(DiscoveredNetworks()) >= 1 {
			token := client.Unsubscribe(config.DiscoveryTopic)
			token.Wait()
			break
		}
	}
}

var (
	config         cc.MQTTConfig
	fromDMService  chan dc.DeviceMessage // in
	fromOTAService chan dc.DeviceMessage // in
	client         mqtt.Client
	dmh            dss.DeviceMessageHandler
	nNetworks      = stringset.New()
	logger         log.Logger
)

/*
 * Initialize()
 *
 * Initialize this service
 */
func Start() (sch.OTAInteractor, error) {
	var err error
	dmh = s

	// ensure Initialize() is called first
	if logger == nil {
		panic(fmt.Errorf("you must call Initialize() in this package before calling Start()"))
	}
	level.Debug(logger).Log("event", "Calling Start()")

	NewOTAHandler(logger) // creates otahandler

	// Initialize a Message Channel
	// one use is to delete devices based on ui input
	establishProviderChannels()

	// start active processing on discovered networks
	enableNetworkTraffic()
	level.Debug(logger).Log("event", "Start() completed")

	return otahandler, err
}

//
func Initialize(cfg cc.Config) ([]string, error) {
	config = cfg.Mqc
	logger = log.With(cfg.Logger, "pkg", "mqttProvider")
	var err error
	level.Debug(logger).Log("event", "Calling Initialize()", "broker", config.BrokerIP)

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
	doNetworkDiscovery()

	level.Debug(logger).Log("event", "Initialize() completed", "networks discovered", strings.Join(DiscoveredNetworks(), ","))
	return DiscoveredNetworks(), err
}

func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")
	// Unsubscribe and shutdown cleanly
	disableNetworkTraffic()

	if otahandler.notifyChannel != nil {
		close(otahandler.notifyChannel) // we own chan, cleanup when done
		otahandler.notifyChannel = nil
	}
	if otahandler.publishChannel != nil {
		close(otahandler.publishChannel) // we own chan, cleanup when done
		otahandler.publishChannel = nil
	}

	client.Disconnect(250)

	time.Sleep(time.Second)
	level.Debug(logger).Log("event", "Stop() completed")
}
