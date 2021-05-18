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
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dss "github.com/skoona/homie-service/pkg/services/deviceSource"
	cc "github.com/skoona/homie-service/pkg/utils"
)

func shutdownDemo() {
	level.Info(logger).Log("msg", "shutdownDemo called")
}

func shutdownLive() {
	level.Info(logger).Log("msg", "shutdownLive called")
}

func runLive(ctx context.Context) error {
	level.Info(logger).Log("msg", "runLive called")
	return nil
}

func runDemo(ctx context.Context) error {
	level.Info(logger).Log("msg", "runDemo called")
	return nil
}

var logger log.Logger

func main() {
	// var hns *cl.HomieNetworks
	var err error
	var dm dss.DeviceMessage

	ctx := cc.BuildRuntimeConfigAndContext("Homie-Service")
	logger = ctx.Value(cc.AppConfig).(cc.Config).Logger

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	/* Prepare for clean exit */
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	// Run the App
	if ctx.Value(cc.AppConfig).(cc.Config).RunMode == "demo" {
		// demo
		err = runDemo(ctx)
	} else {
		// live
		err = runLive(ctx)
	}

	// wait on ctrl-c
	sig := <-sigs
	level.Info(logger).Log("msg", sig)
	level.Info(logger).Log("msg", "Shutting Down")

	if ctx.Value(cc.AppConfig).(cc.Config).RunMode == "demo" {
		// demo
		shutdownDemo()
	} else {
		// live
		shutdownLive()
	}

	// out, err := json.MarshalIndent(hns, "", "  ")
	if err != nil {
		level.Error(logger).Log("msg", err.Error())
	} else {
		level.Info(logger).Log("DEBUG", "HomieNetworks: --> {JSON}")
	}

	os.Exit(0)

}
