
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
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"github.com/skoona/homie-service/pkg/services"
	cc "github.com/skoona/homie-service/pkg/utils"
	"os"
	"os/signal"
	"syscall"
)

func StartUp() (dc.CoreService, cc.Config) {
	return services.Service()
}

func Shutdown() {
	services.Shutdown()
}

func main() {

	coreSvc, cfg := StartUp()
	if coreSvc == nil {
		panic("Service Failed to start")
	}

	logger := log.With(cfg.Logger, "ui", "api")

	/*
	 * Prepare for clean exit
	 */
	errs := make(chan error, 1)
	//go func() {
	//	logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
	//	errs <- http.ListenAndServe(*httpAddr, nil)
	//}()
	go func(shutdown chan error) {
		systemSignalChannel := make(chan os.Signal, 1)
		signal.Notify(systemSignalChannel, syscall.SIGINT, syscall.SIGTERM)
		sig := <-systemSignalChannel // wait on ctrl-c
		shutdown <- fmt.Errorf("%s", sig)
	}(errs)
	level.Info(logger).Log("event", "shutdown requested", "cause", <-errs)

	Shutdown()

	os.Exit(0)
}

