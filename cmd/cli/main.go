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
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dp "github.com/skoona/homie-service/pkg/services/demoProvider"
	dss "github.com/skoona/homie-service/pkg/services/deviceSource"
	dds "github.com/skoona/homie-service/pkg/services/deviceStorage"
	mq "github.com/skoona/homie-service/pkg/services/liveProvider"
	cc "github.com/skoona/homie-service/pkg/utils"
)

func shutdownDemo() {
	level.Debug(logger).Log("event", "shutdownDemo() called")
	dss.Stop()
	dds.Stop()
	dp.Stop()
	level.Debug(logger).Log("event", "shutdownDemo() completed")
}

func shutdownLive() {
	level.Debug(logger).Log("event", "shutdownLive() called")
	mq.Stop()
	dss.Stop()
	dds.Stop()
	level.Debug(logger).Log("event", "shutdownLive() completed")
}

func runLive(cfg cc.Config) error {
	level.Debug(logger).Log("event", "runLive() called")
	repo, err := dds.Start(cfg)
	dssService, err = dss.Start(cfg, repo)
	mq.Start(cfg, dssService)
	level.Debug(logger).Log("event", "runLive() completed")
	return err
}

func runDemo(cfg cc.Config) error {
	level.Debug(logger).Log("event", "runDemo() called")
	repo, err := dds.Start(cfg)
	dssService, err = dss.Start(cfg, repo)
	dp.Start(cfg, dssService)
	level.Debug(logger).Log("event", "runDemo() completed")
	return err
}

var logger log.Logger
var dssService dss.Service

func main() {
	// var hns *cl.HomieNetworks
	var err error

	cfg := cc.BuildRuntimeConfig("Homie-Service")
	logger = log.With(cfg.Logger, "pkg", "main")

	level.Debug(logger).Log("event", "service started")

	/* Prepare for clean exit */
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

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

	// wait on ctrl-c
	sig := <-sigs
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

	level.Debug(logger).Log("event", "service ended")

	os.Exit(0)
}
