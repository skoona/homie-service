/*
 * services/service.go:
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

package services

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dp "github.com/skoona/homie-service/pkg/demoProvider"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	sch "github.com/skoona/homie-service/pkg/deviceScheduler"
	dss "github.com/skoona/homie-service/pkg/deviceSource"
	dds "github.com/skoona/homie-service/pkg/deviceStorage"
	mq "github.com/skoona/homie-service/pkg/mqttProvider"
	cc "github.com/skoona/homie-service/pkg/utils"
	"os"
)

func shutdownDemo(logger log.Logger) {
	level.Debug(logger).Log("event", "ShutdownDemo() called")
	dp.Stop()
	dss.Stop()

	// List the Devices Found/Recorded
	//dds.ListHomieDB()

	dds.Stop()
	dc.Stop()
	level.Debug(logger).Log("event", "ShutdownDemo() completed")
}

func shutdownLive(logger log.Logger) {
	level.Debug(logger).Log("event", "ShutdownLive() called")
	mq.Stop()
	sch.Stop()

	// List the Devices Found/Recorded
	//dds.ListHomieDB()

	dss.Stop()
	dds.Stop()
	dc.Stop()
	level.Debug(logger).Log("event", "ShutdownLive() completed")
}

func runLive(cfg cc.Config, logger log.Logger)(dc.CoreService, error) {
	level.Debug(logger).Log("event", "RunLive() called")
	otap, dsp, netw, err := mq.Initialize(cfg)                 // message stream
	if err != nil {
		return nil, err
	}
	networks = netw

	repo, err := dds.Start(cfg)                                    // message db
	if err != nil {
		return nil, err
	}

	dep, _ := dss.Start(cfg, repo, dsp)                          // message aggregation
	sched := sch.Start(cfg, otap, repo)                          // ota scheduler
	coreSvc, _ := dc.Start(cfg, dep, sched, repo, networks) // network logic -- may need scheduler
	err = mq.Start()                                           // activate message stream
	if err != nil {
		return nil, err
	}

	level.Debug(logger).Log("event", "RunLive() completed")
	return coreSvc, err
}

func runDemo(cfg cc.Config, logger log.Logger) (dc.CoreService, error) {
	level.Debug(logger).Log("event", "RunDemo() called")

	otap, dsp, networks, err := dp.Initialize(cfg)                 // message stream
	if err != nil {
		return nil, err
	}

	repo, err := dds.Start(cfg)                                    // message db
	if err != nil {
		return nil, err
	}

	dep, err := dss.Start(cfg, repo, dsp)                          // message aggregation
	if err != nil {
		return nil, err
	}

	sched := sch.Start(cfg, otap, repo)                          // ota scheduler
	coreSvc, _ := dc.Start(cfg, dep, sched, repo, networks) // network logic -- may need scheduler
	err = dp.Start()                                           // activate message stream
	if err != nil {
		return nil, err
	}

	level.Debug(logger).Log("event", "RunDemo() completed")
	return coreSvc, err
}

var (
	cfg			cc.Config
	coreSvc dc.CoreService
	networks []string
)

func startUp() (dc.CoreService, cc.Config, error) {
	var err error

	cfg, err = cc.BuildRuntimeConfig("Homie-Service")
	if err != nil {
		os.Exit(1)
	}
	logger := log.With(cfg.Logger, "pkg", "services")

	level.Info(logger).Log("event", "service started")

	// Run the App
	if cfg.RunMode == "live" {
		coreSvc, err = runLive(cfg, logger)
	} else {
		coreSvc, err = runDemo(cfg, logger)
	}
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		os.Exit(2)
	}

	return coreSvc, cfg,  err
}

// Shutdown shutdown func for Core Service
func Shutdown() {
	defer level.Info(cfg.Logger).Log("event", "service ended")
	if cfg.RunMode == "live" {
		shutdownLive(cfg.Logger)
	} else {
		shutdownDemo(cfg.Logger)
	}
}

// Service main entry point for Core Service
func Service() (*dc.CoreService, *cc.Config, []string) {

	coreSvc, cfg, err := startUp()
	if err != nil {
		e := fmt.Errorf("service startup failure: %s", err.Error())
		panic(e.Error())
	}

	return &coreSvc, &cfg, networks
}

