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
	* Implements sch.OTAInteractor
	* implements dss.StreamProvider
*/

import (
	"errors"
	"fmt"
	"strings"
	"time"

	sch "github.com/skoona/homie-service/internal/deviceScheduler"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	stringset "github.com/jjeffery/stringset"
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
 * networkDiscovery
 * Handles '+/+/$name' a discovery technique which list HomieDevices on any network prefix
 */
var networkDiscovery mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	parts := strings.Split(msg.Topic(), "/")
	if !nNetworks.Contains(parts[0]) {
		nNetworks.Add(parts[0])
		level.Debug(logger).Log("Network Discovery", parts[0])
	}
	level.Debug(logger).Log("Network Discovery topic", msg.Topic(), "payload", msg.Payload())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	level.Debug(logger).Log("status", "MQTT Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	level.Debug(logger).Log("status", "lost MQTT connection", "error", err.Error())
	//time.Sleep(10 * time.Second)
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
// consider
// AddRoute(topic string, callback MessageHandler)
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
func enableNetworkTraffic( plog log.Logger) error {
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
	subWithHandler(config.DiscoveryTopic, networkDiscovery) // Homie Discovery Topic
	for {
		time.Sleep(5 * time.Second) // delay long enough to collect networks
		if len(DiscoveredNetworks()) >= 1 {
			token := client.Unsubscribe(config.DiscoveryTopic)
			token.Wait()
			break
		}
	}
}

var (
	config    cc.MQTTConfig
	client    mqtt.Client
	nNetworks = stringset.New()
	logger    log.Logger
	bInitialized  bool

)

/*
 * Start()
 *
 * Start this service
 */
func Start() error {
	var err error

	// ensure Initialize() is called first
	if !bInitialized {
		return fmt.Errorf("you must call Initialize() in this package before calling Start()")
	}
	level.Debug(logger).Log("event", "Calling Start()")

	// start active processing on discovered networks
	// enableNetworkTraffic()  -- wait for dss request

	level.Debug(logger).Log("event", "Start() completed")

	return err
}

func Initialize(cfg cc.Config) (sch.OTAInteractor, dss.StreamProvider, []string, error) {
	config = cfg.Mqc
	logger = log.With(cfg.Logger, "pkg", "mqttProvider")
	var err error
	level.Debug(logger).Log("event", "Calling Initialize()", "broker", config.BrokerIP )

	if config.BrokerIP == "" || len(config.BrokerIP) < 9 {
		err := errors.New("empty or invalid broker ip address")
		return nil, nil, nil, err
	}

	/* Initialize MQTT */
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.BrokerIP, config.BrokerPort))
	opts.SetClientID(config.ClientID)
	opts.SetUsername(config.UserName)
	opts.SetPassword(config.UserPassword)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(10 * time.Second)
	opts.SetMaxReconnectInterval(30 * time.Minute)
	opts.SetOrderMatters(true)
	opts.SetKeepAlive(120)
	opts.SetPingTimeout(10)
	if len(config.LwtTopic) > 0 && len(config.LwtMessage) > 0 {
		opts.SetWill(config.LwtTopic, config.LwtMessage, 0, true)
	}
	opts.SetDefaultPublishHandler(defaultOnMessage)
	opts.SetOnConnectHandler(connectHandler)
	opts.SetConnectionLostHandler(connectLostHandler)

	client = mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		level.Error(logger).Log("cause", "connect() failed", "error", token.Error())
		return nil, nil, nil, token.Error()
	}

	NewOTAStream(logger)      // creates otastream
	NewStreamProvider(logger) // creates stream provider

	// allow for network discovery
	doNetworkDiscovery()

	bInitialized = true

	level.Debug(logger).Log("event", "Initialize() completed", "networks discovered", strings.Join(DiscoveredNetworks(), ","), "clientID", opts.ClientID)
	return otastream, dStream, DiscoveredNetworks(), err
}

func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")

	// Unsubscribe and shutdown cleanly
	disableNetworkTraffic()
	client.Disconnect(250)

	if otastream.notifyChannel != nil {
		close(otastream.notifyChannel) // we own chan, cleanup when done
		otastream.notifyChannel = nil
	}
	if otastream.publishChannel != nil {
		close(otastream.publishChannel) // we own chan, cleanup when done
		otastream.publishChannel = nil
	}
	if dStream.notifyChannel != nil {
		close(dStream.notifyChannel) // we own chan, cleanup when done
		dStream.notifyChannel = nil
	}
	if dStream.publishChannel != nil {
		close(dStream.publishChannel) // we own chan, cleanup when done
		dStream.publishChannel = nil
	}

	time.Sleep(3 * time.Second)
	level.Debug(logger).Log("event", "Stop() completed")
}
