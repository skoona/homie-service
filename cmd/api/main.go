
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

package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/skoona/homie-service/pkg/UIAdapters/api/handlers"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"github.com/skoona/homie-service/pkg/services"
	cc "github.com/skoona/homie-service/pkg/utils"
	"os"
	"os/signal"
	"syscall"
)


func StartUp() (*dc.CoreService, *cc.Config, []string) {
	return services.Service()
}

func Shutdown() {
	services.Shutdown()
}

func main() {

	coreSvc, cfg, _ := StartUp()
	if coreSvc == nil {
		panic("Service Failed to start")
	}

	logger := log.With(cfg.Logger, "ui", "api")

	ctrl := handlers.NewApiProvider(cfg, coreSvc, &logger)

	/*
	 * Prepare for clean exit
	 */
	errs := make(chan error, 1)

	// start the server
	go func(shutdown chan error) {
		level.Info(logger).Log("event", "starting server", "serving", cfg.Api.BindAddress)
		shutdown <- ctrl.Run()
	}(errs)

	go func(shutdown chan error) {
		systemSignalChannel := make(chan os.Signal, 1)
		signal.Notify(systemSignalChannel, syscall.SIGINT, syscall.SIGTERM)
		sig := <-systemSignalChannel // wait on ctrl-c
		shutdown <- fmt.Errorf("%s", sig)
	}(errs)

	// wait for server or interrupt
	level.Info(logger).Log("event", "shutdown requested", "cause", <-errs)

	ctrl.Shutdown()
	Shutdown()

	os.Exit(0)
}

