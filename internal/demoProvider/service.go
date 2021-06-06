package demoProvider

/**
  demoProvider/service.go:


  The design goal for this file is:
	* Generate Demo DeviceMessages from a specified Mosquitto Logfile
	* implements dss.StreamProvider to SEND to deviceSource
		- DOES NOT SUPPORT RECEIVING FROM deviceSource
		- DOES NOT SUPPORT OTA requests from Scheduler
*/

import (
	"bufio"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jjeffery/stringset"
	sch "github.com/skoona/homie-service/internal/deviceScheduler"
	dss "github.com/skoona/homie-service/internal/deviceSource"
	cc "github.com/skoona/homie-service/internal/utils"
	"os"
	"strings"
	"time"
)

var (
	config    cc.MQTTConfig
	cfg			cc.Config
	nNetworks = stringset.New()
	logger    log.Logger
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
}

func publish(topic string, payload []byte, retain bool, qos byte) {
	token := mClient.Publish(topic, qos, retain, payload)
	token.Wait()
	level.Debug(logger).Log("Published topic", topic, "payload", payload)
}

func sub(topic string) mqtt.Token {
	token := mClient.Subscribe(topic, 1, nil)
	token.Wait()
	level.Debug(logger).Log("Subscribed to topic", topic)

	return token
}
// consider
// AddRoute(topic string, callback MessageHandler)
func subWithHandler(topic string, callback mqtt.MessageHandler) mqtt.Token {
	token := mClient.Subscribe(topic, 1, callback)
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
			token := mClient.Unsubscribe(networkTopic)
			token.Wait()
		}
	} else {
		token := mClient.Unsubscribe(config.SubscriptionTopic)
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
	dStream.GetNotifyChannel()
	trafficGenerator(cfg.Dbc.DemoSource, logger)
	return token.Error()
}

func trafficGenerator(demoFile string, plog log.Logger) error {
	level.Debug(plog).Log("event", "calling trafficGenerator()", "source file", demoFile)
	var err error

	/*
	 * Create a Go Routine for the MQTT Channel to
	 * convert mq logs messages into mqtt.Messages
	 */
	go func(fp string, plog log.Logger) {
		demoRender(fp, plog, false)
	}(demoFile, plog)

	level.Debug(plog).Log("event", "trafficGenerator() active")
	return err
}

func demoRender(filepath string, tlog log.Logger, limit bool) {
	var err error
	var file *os.File
	var idx uint16 = 0

	file, err = os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		level.Error(tlog).Log("error", err.Error())
		return // vs panic()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err = scanner.Err(); err != nil {
			continue
		}

		parts := strings.Split(line, " ")
		topic := parts[0]
		payload := strings.Join(parts[1:], " ")

		idx++
		for k, v := range mClient.subscriptions {
			dm := dStream.CreateDemoDeviceMessage(topic, []byte(payload), idx, false, 1)
			if strings.Contains(k,topic) {
				v.callback(mClient, dm)
			} else if len(parts) >= 3 {
				if strings.Contains(k, parts[2]) {
					v.callback(mClient, dm)
				}
			} else if !limit {  // runmode if false, discovery mode if true
					mClient.cbDefault(mClient, dm)
			}
		}

		if limit {
			if idx >= 106 {
				break  // todo break then return
			}
			time.Sleep(2 * time.Millisecond) // slow the pace
		} else {
			time.Sleep(5 * time.Millisecond) // slow the pace
		}
	}

	level.Debug(tlog).Log("event", "demoRender() completed")
}

/**
 * Discover existing Homie Networks
 */
func doNetworkDiscovery() {
	// allow for network discovery
	logger = log.With(cfg.Logger, "method", "doNetworkDiscovery")
	level.Debug(logger).Log("event", "Calling doNetworkDiscovery()")
	subWithHandler(config.DiscoveryTopic, networkDiscovery) // Homie Discovery Topic
	demoRender(cfg.Dbc.DemoSource, logger, true)
	for {
		time.Sleep(5 * time.Second) // delay long enough to collect networks
		if len(DiscoveredNetworks()) >= 1 {
			token := mClient.Unsubscribe(config.DiscoveryTopic)
			token.Wait()
			break
		}
	}
	level.Debug(logger).Log("event", "doNetworkDiscovery() completed", "networks", strings.Join(DiscoveredNetworks(), ","))
}

func Initialize(dfg cc.Config) (sch.OTAInteractor, dss.StreamProvider, []string, error) {
	var err error
	config = cfg.Mqc
	cfg = dfg
	logger = log.With(cfg.Logger, "pkg", "demoProvider")
	level.Debug(logger).Log("event", "Calling Initialize()", "sourceFile", cfg.Dbc.DemoSource)

	file, err := os.OpenFile(cfg.Dbc.DemoSource, os.O_RDONLY, 0666)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return nil, nil, nil,  err // vs panic()
	}
	defer file.Close()

	NewMockClient()
	mClient.SetDefaultHandler(defaultOnMessage)
	mClient.SetConnLostHandler(connectLostHandler)
	mClient.SetConnHandler(connectHandler)
	if token := mClient.Connect(); token.Wait() && token.Error() != nil {
		level.Error(logger).Log("cause", "connect() failed", "error", token.Error())
		return nil, nil, nil, token.Error()
	}

	NewOTAStream(logger)      // creates otastream
	NewStreamProvider(logger) // creates stream provider

	// allow for network discovery
	doNetworkDiscovery()

	level.Debug(logger).Log("event", "Initialize() completed", "networks discovered", strings.Join(DiscoveredNetworks(), ","))
	return otastream, dStream, DiscoveredNetworks(), err
}
func Start() error {
	var err error

	// ensure Initialize() is called first
	if logger == nil {
		return fmt.Errorf("you must call Initialize() in this package before calling Start()")
	}
	level.Debug(logger).Log("event", "Calling Start()")

	// start active processing on discovered networks
	// enableNetworkTraffic()  -- wait for dss request

	time.Sleep(4 * time.Second) // slow the pace

	level.Debug(logger).Log("event", "Start() completed")

	return err
}
func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")

	// Unsubscribe and shutdown cleanly
	disableNetworkTraffic()
	mClient.Disconnect(250)

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