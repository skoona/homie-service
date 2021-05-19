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
	dp "github.com/skoona/homie-service/pkg/services/demoProvider"
	dss "github.com/skoona/homie-service/pkg/services/deviceSource"
	dds "github.com/skoona/homie-service/pkg/services/deviceStorage"
	mq "github.com/skoona/homie-service/pkg/services/liveProvider"
	cc "github.com/skoona/homie-service/pkg/utils"
)

func shutdownDemo() {
	level.Info(logger).Log("msg", "shutdownDemo() called")
	dss.Stop()
	dds.Stop()
	dp.Stop()
}

func shutdownLive() {
	level.Info(logger).Log("msg", "shutdownLive() called")
	mq.Stop()
	dss.Stop()
	dds.Stop()
}

func runLive(ctx context.Context) error {
	level.Info(logger).Log("msg", "runLive() called")
	repo, err := dds.Start(ctx)
	dssService, _ = dss.Start(ctx, repo)
	mq.Start(ctx, dssService)
	return err
}

func runDemo(ctx context.Context) error {
	level.Info(logger).Log("msg", "runDemo() called")
	repo, err := dds.Start(ctx)
	dssService, _ = dss.Start(ctx, repo)
	dp.Start(ctx, dssService)
	return err
}

var logger log.Logger
var dssService dss.Service

func main() {
	// var hns *cl.HomieNetworks
	var err error

	ctx := cc.BuildRuntimeConfigAndContext("Homie-Service")
	logger = log.With(ctx.Value(cc.AppConfig).(cc.Config).Logger, "pkg", "main")

	level.Debug(logger).Log("msg", "service started")
	defer level.Debug(logger).Log("msg", "service ended")

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
	level.Debug(logger).Log("msg", sig)
	level.Debug(logger).Log("msg", "Shutting Down")

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
