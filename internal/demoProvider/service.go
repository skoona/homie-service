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
	cfg			cc.Config
	nNetworks = stringset.New()
	logger    log.Logger
	bInitialized bool = false
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
	mClient.Publish(topic, qos, retain, payload)
	level.Debug(logger).Log("Published topic", topic)
}

// consider
// AddRoute(topic string, callback MessageHandler)
func subWithHandler(topic string, callback mqtt.MessageHandler) mqtt.Token {
	token := mClient.Subscribe(topic, 1, callback)
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
			mClient.Unsubscribe(networkTopic)
		}
	} else {
		mClient.Unsubscribe(cfg.Mqc.SubscriptionTopic)
	}

	return nil
}
/**
 * Activate Subscriptions starting network traffic
 */
func enableNetworkTraffic( plog log.Logger) error {
	networks := DiscoveredNetworks()

	if len(networks) > 0 {
		for _, network := range networks {
			networkTopic := fmt.Sprintf("%s/#", network)
			subWithHandler(networkTopic, defaultOnMessage) // Available Topic
		}
	} else {
		subWithHandler(cfg.Mqc.SubscriptionTopic, defaultOnMessage) // Default Topic
	}
	dStream.GetNotifyChannel()
	trafficGenerator(cfg.Dbc.DemoSource, plog)
	return nil
}

func trafficGenerator(demoFile string, plog log.Logger) error {
	level.Info(plog).Log("event", "calling trafficGenerator()", "source file", demoFile)

	/*
	 * Create a Go Routine for the MQTT Channel to
	 * convert mq logs messages into mqtt.Messages
	 */
	go func(fp string, plog log.Logger) {
		demoRender(fp, plog, false)
	}(demoFile, plog)

	level.Info(plog).Log("event", "trafficGenerator() active")
	return nil
}

func demoRender(filepath string, tlog log.Logger, limit bool) {
	var err error
	var file *os.File
	var idx uint16 = 0

	foundFile := cc.LocateFile(filepath)

	file, err = os.OpenFile(foundFile, os.O_RDONLY, 0666)
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
		tparts  := strings.Split(topic, "/")
		payload := strings.Join(parts[1:], " ")

		idx++
		msg := newMockMessage(topic, idx, 1, false, []byte(payload))

		// Honor subscriptions
		for k, v := range mClient.subs {
			if strings.Contains(k,topic) { // ota watches
				v.callback(mClient, msg)
				//fmt.Printf("A:[%d] %s <--> %s \n", msg.MessageID(), k, topic)
				break
			} else if len(tparts) >= 3 {
				if strings.Contains(k, tparts[2]) { // +/+/$name  Discovery
					v.callback(mClient, msg)
					//fmt.Printf("B:[%d] %s <--> %s \n", msg.MessageID(), k, tparts[2])
					break
				}
			}
		}

		if !limit {  // runmode if false, discovery mode if true
			mClient.cbDefault(mClient, msg)
			//fmt.Printf("C:[%d] %s\n", msg.MessageID(), msg.Topic())
		}

		if limit {
			if idx >= 106 {
				break
			}
		}
	}

	level.Info(tlog).Log("event", "demoRender() completed", "limit", limit)
}

/**
 * Discover existing Homie Networks
 */
func doNetworkDiscovery() {
	// allow for network discovery
	logger = log.With(cfg.Logger, "method", "doNetworkDiscovery")
	level.Debug(logger).Log("event", "Calling doNetworkDiscovery()")
	subWithHandler(cfg.Mqc.DiscoveryTopic, networkDiscovery) // Homie Discovery Topic
	demoRender(cfg.Dbc.DemoSource, logger, true)
	for {
		time.Sleep(4 * time.Second) // delay long enough to collect networks
		if len(DiscoveredNetworks()) >= 1 {
			mClient.Unsubscribe(cfg.Mqc.DiscoveryTopic)
			break
		}
	}
	level.Debug(logger).Log("event", "doNetworkDiscovery() completed", "networks", strings.Join(DiscoveredNetworks(), ","))
}

func Initialize(dfg cc.Config) (sch.OTAInteractor, dss.StreamProvider, []string, error) {
	var err error
	cfg = dfg
	logger = log.With(dfg.Logger, "pkg", "demoProvider")
	level.Debug(logger).Log("event", "Calling Initialize()", "sourceFile", cfg.Dbc.DemoSource)


	if cc.LocateFile(cfg.Dbc.DemoSource) == "" {
		level.Error(logger).Log("error", "FILE NOT FOUND", "file", cfg.Dbc.DemoSource)
		return nil, nil, nil, fmt.Errorf("FILE NOT FOUND: %s", cfg.Dbc.DemoSource) // vs panic()
	}

	NewMockClient()
	mClient.SetDefaultHandler(defaultOnMessage)
	mClient.SetConnLostHandler(connectLostHandler)
	mClient.SetConnHandler(connectHandler)

	mClient.Connect()

	NewOTAStream(logger)      // creates otastream
	NewStreamProvider(logger) // creates stream provider

	// allow for network discovery
	doNetworkDiscovery()

	bInitialized = true

	level.Debug(logger).Log("event", "Initialize() completed", "networks discovered", strings.Join(DiscoveredNetworks(), ","))
	return otastream, dStream, DiscoveredNetworks(), err
}
func Start() error {
	var err error

	// ensure Initialize() is called first
	if !bInitialized {
		return fmt.Errorf("you must call Initialize() in this package before calling Start()")
	}
	level.Debug(logger).Log("event", "Calling Start()")

	time.Sleep(4 * time.Second) // slow the pace

	level.Debug(logger).Log("event", "Start() completed")

	return err
}
func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")
	bInitialized = false

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