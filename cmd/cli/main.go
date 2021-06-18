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

func ShutdownDemo() {
	level.Debug(logger).Log("event", "ShutdownDemo() called")
	dp.Stop()
	dss.Stop()

	// List the Devices Found/Recorded
	//dds.ListHomieDB()

	dds.Stop()
	dc.Stop()
	level.Debug(logger).Log("event", "ShutdownDemo() completed")
}

func ShutdownLive() {
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

func RunLive(cfg cc.Config)(dc.CoreService, error) {
	level.Debug(cfg.Logger).Log("event", "RunLive() called")
	otap, dsp, networks, err = mq.Initialize(cfg)                 // message stream
	if err != nil {
		return nil, err
	}
	repo, err = dds.Start(cfg)                                    // message db
	if err != nil {
		return nil, err
	}
	dep, _ = dss.Start(cfg, repo, dsp)                          // message aggregation
	sched = sch.Start(cfg, otap, repo)                          // ota scheduler
	coreSvc, siteNetworks = dc.Start(cfg, dep, sched, repo, networks) // network logic -- may need scheduler
	err = mq.Start()                                           // activate message stream
	if err != nil {
		return nil, err
	}
	level.Debug(cfg.Logger).Log("event", "RunLive() completed")
	return coreSvc, err
}

func RunDemo(cfg cc.Config) (dc.CoreService, error) {
	level.Debug(cfg.Logger).Log("event", "RunDemo() called")
	otap, dsp, networks, err = dp.Initialize(cfg)                 // message stream
	if err != nil {
		return nil, err
	}
	repo, err = dds.Start(cfg)                                    // message db
	if err != nil {
		return nil, err
	}
	dep, _ = dss.Start(cfg, repo, dsp)                          // message aggregation
	sched = sch.Start(cfg, otap, repo)                          // ota scheduler
	coreSvc, siteNetworks = dc.Start(cfg, dep, sched, repo, networks) // network logic -- may need scheduler
	err = dp.Start()                                           // activate message stream
	if err != nil {
		return nil, err
	}
	level.Debug(cfg.Logger).Log("event", "RunDemo() completed")
	return coreSvc, err
}

var (
	networks     []string
	logger       log.Logger
	siteNetworks *dc.SiteNetworks
	dep          dc.DeviceEventProvider
	sched        dc.SchedulerProvider
	coreSvc      dc.CoreService
	repo         dc.Repository
	otap         sch.OTAInteractor
	dsp          dss.StreamProvider
	err 		 error
	cfg			cc.Config
)

func StartUp() (dc.CoreService, cc.Config, error) {
	cfg, err = cc.BuildRuntimeConfig("Homie-Service")
	if err != nil {
		os.Exit(1)
	}
	logger = log.With(cfg.Logger, "pkg", "main")

	level.Info(logger).Log("event", "service started")
	defer level.Info(logger).Log("event", "service ended")

	// Run the App
	if cfg.RunMode == "live" {
		coreSvc, err = RunLive(cfg)
	} else {
		coreSvc, err = RunDemo(cfg)
	}
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		os.Exit(2)
	}

	return coreSvc, cfg, err
}

func Shutdown(cfg cc.Config) {
	if cfg.RunMode == "live" {
		ShutdownLive()
	} else {
		ShutdownDemo()
	}
}

func main() {

	_, _, err = StartUp()

	/*
	 * Prepare for clean exit
	*/
	errs := make(chan error, 1)
	//go func() {
	//	logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
	//	errs <- http.ListenAndServe(*httpAddr, nil)
	//}()
	go func(shutdown chan error) {
		systemSignalChannel := make(chan os.Signal)
		signal.Notify(systemSignalChannel, syscall.SIGINT, syscall.SIGTERM)
		sig := <-systemSignalChannel // wait on ctrl-c
		shutdown <- fmt.Errorf("%s", sig)
	}(errs)
	level.Info(logger).Log("event", "shutdown requested", "cause", <-errs)

	Shutdown(cfg)

	if err != nil {
		level.Error(logger).Log("error", err.Error())
	}


	// Dump the SiteNetwork and all nodes as JSON
	//out, err := json.MarshalIndent(siteNetworks, "", "  ")
	//if err != nil {
	//	level.Warn(logger).Log("action", err.Error())
	//} else {
	//	fmt.Println("Site Network")
	//	fmt.Println(string(out))
	//}

	os.Exit(0)
}
