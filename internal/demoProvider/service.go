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
)

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

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

func sub(topic string) mqtt.Token {
	token := mClient.Subscribe(topic, 1, nil)
	level.Debug(logger).Log("Subscribed to topic", topic)

	return token
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
	var foundFile string

	if fileExists(filepath) {
		foundFile = filepath
	} else {
		if strings.HasPrefix(filepath, "../") {
			foundFile = strings.ReplaceAll(filepath, "../.", "")
			if !fileExists(foundFile) {
				level.Error(tlog).Log("error", "FILE NOT FOUND", "file", foundFile)
				return
			}
		}
	}

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
		tparts := strings.Split(topic, "/")
		payload := strings.Join(parts[1:], " ")

		idx++
		for k, v := range mClient.subs {
			msg := newMockMessage(topic, idx, 1, false, []byte(payload))
			if strings.Contains(k,topic) {
				v.callback(mClient, msg)
			} else if len(tparts) >= 3 {
				if strings.Contains(k, tparts[2]) {
					v.callback(mClient, msg)
				}
			}
			if !limit {  // runmode if false, discovery mode if true
					mClient.cbDefault(mClient, msg)
			}
		}

		if limit {
			if idx >= 106 {
				break  // todo break then return
			}
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

	var foundFile string

	if !fileExists(cfg.Dbc.DemoSource) {
		if strings.HasPrefix(cfg.Dbc.DemoSource, "../") {
			foundFile = strings.ReplaceAll(cfg.Dbc.DemoSource, "../.", "")
			if !fileExists(foundFile) {
				level.Error(logger).Log("error", "FILE NOT FOUND", "file", foundFile)
				return nil, nil, nil, fmt.Errorf("FILE NOT FOUND: %s", foundFile) // vs panic()
			}
		}
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