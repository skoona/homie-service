package main

/*
 * main
 *
 * Subscriptions
 * - Network Discovery   +/+/$name
 * - Device Networks     homie/#, sknSensors/#
 * - OTA Streaming       +/+/$implementation/ota/firmware/<md5>
 *                       +/+/$implementation/ota/firmware <data>
 *                       +/+/$implementation/ota/status
 *
 * Publish
 * - Device Network Commands
 * - OTA Streaming
 *
 * Channels: IN
 * MQTT to Datasource       mqtt.Message  mqttChannel
 * DataSource(MQTT/Demo) to Datasource  DeviceMessage dvcSyncChannel
 * Datasource to CoreLogic  DeviceMessage coreLogicChannel
 * CoreLogic to Scheduler   DeviceMessage schedTrigChannel
 * MQTT to Scheduler        mqtt.Message  otaRspChannel
 *
 * Channels: OUT
 * Scheduler to MQTT Publisher   OTA Streaming   publishChannel
 * CoreLogic to MQTT Publisher   Device Commands publishChannel
 *
 * Datasource
 */

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dp "github.com/skoona/homie-service/internal/demoProvider"
	dc "github.com/skoona/homie-service/internal/deviceCore"

	sch "github.com/skoona/homie-service/internal/deviceScheduler"
	dss "github.com/skoona/homie-service/internal/deviceSource"
	dds "github.com/skoona/homie-service/internal/deviceStorage"
	mq "github.com/skoona/homie-service/internal/mqttProvider"
	cc "github.com/skoona/homie-service/internal/utils"
)

func shutdownDemo() {
	level.Debug(logger).Log("event", "shutdownDemo() called")
	dp.Stop()
	sch.Stop()
	dss.Stop()

	// List the Devices Found/Recorded
	dds.ListHomieDB()

	dds.Stop()
	dc.Stop()
	level.Debug(logger).Log("event", "shutdownDemo() completed")
}

func shutdownLive() {
	level.Debug(logger).Log("event", "shutdownLive() called")
	mq.Stop()
	// sch.Stop()

	// List the Devices Found/Recorded
	dds.ListHomieDB()

	dss.Stop()
	dds.Stop()
	dc.Stop()
	level.Debug(logger).Log("event", "shutdownLive() completed")
}

func runLive(cfg cc.Config) error {
	level.Debug(logger).Log("event", "runLive() called")
	networks, _ := mq.Initialize(cfg)
	coreSvc, _ := dc.Start(cfg, networks)
	repo, err := dds.Start(cfg)
	dmh, err = dss.Start(cfg, repo, coreSvc) // needs coreSvc
	// _ = sch.Start(cfg, dmh)
	err = mq.Start(dmh)
	level.Debug(logger).Log("event", "runLive() completed")
	return err
}

func runDemo(cfg cc.Config) error {
	level.Debug(logger).Log("event", "runDemo() called")
	networks, _ := dp.Initialize(cfg)
	coreSvc, _ := dc.Start(cfg, networks)
	repo, err := dds.Start(cfg)
	dmh, err = dss.Start(cfg, repo, coreSvc) // needs coreSvc
	// _ = sch.Start(cfg, dmh)
	err = dp.Start(dmh)
	level.Debug(logger).Log("event", "runDemo() completed")
	return err
}

var logger log.Logger
var dmh dss.DeviceMessageHandler

func main() {
	// var hns *cl.HomieNetworks
	var err error

	cfg := cc.BuildRuntimeConfig("Homie-Service")
	logger = log.With(cfg.Logger, "pkg", "main")

	level.Debug(logger).Log("event", "service started")
	defer level.Debug(logger).Log("event", "service ended")

	// Run the App
	if cfg.RunMode == "demo" {
		// demo
		err = runDemo(cfg)
	} else {
		// live
		err = runLive(cfg)
	}
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		panic(err.Error())
	}

	/* Prepare for clean exit */
	systemSignalChannel := make(chan os.Signal, 1)
	signal.Notify(systemSignalChannel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	sig := <-systemSignalChannel // wait on ctrl-c
	level.Debug(logger).Log("event", sig)
	level.Debug(logger).Log("event", "Shutting Down")

	if cfg.RunMode == "demo" {
		// demo
		shutdownDemo()
	} else {
		// live
		shutdownLive()
	}

	if err != nil {
		level.Error(logger).Log("error", err.Error())
	}

	// Dump the SiteNetwork and all nodes as JSON
	out, err := json.MarshalIndent(*dc.GetSiteNetworks(), "", "  ")
	if err != nil {
		level.Warn(logger).Log("action", err.Error())
	} else {
		fmt.Println("Site Network")
		fmt.Println(string(out))
	}

	os.Exit(0)
}
