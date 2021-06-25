
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
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	api "github.com/skoona/homie-service/pkg/api/handlers"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"github.com/skoona/homie-service/pkg/services"
	cc "github.com/skoona/homie-service/pkg/utils"
	stdlog "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartUp() (*dc.CoreService, *cc.Config) {
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

	validator := api.NewValidation()
	ctrl := api.NewApiController(coreSvc, &logger, validator)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// handlers for API
	getR := sm.PathPrefix("/api/v1").
		Methods(http.MethodGet).
		Subrouter()

	getR.HandleFunc("/site", ctrl.AllNetworks)
	getR.HandleFunc("/network/{networkName:[a-z|A-Z]+}", ctrl.NetworkByName)
	getR.HandleFunc("/deviceByName/{networkName:[a-z|A-Z]+}/{deviceName:[a-z|A-Z]+}", ctrl.DeviceByName)
	getR.HandleFunc("/deviceById/{networkName:[a-z|A-Z]+}/{deviceID:[a-z|A-Z|0-9]+}", ctrl.DeviceByID)

	delR := sm.PathPrefix("/api/v1").
		Methods(http.MethodDelete).
		Subrouter()
	delR.HandleFunc("/removeDeviceId/{networkName:[a-z|A-Z]+}/{deviceID:[a-z|A-Z|0-9]+}", ctrl.RemoveDeviceID)

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	/*
	 * Prepare for clean exit
	 */
	errs := make(chan error, 1)
	stdlog.SetOutput(log.NewStdlibAdapter(logger))
	// create a new server
	s := http.Server{
		Addr:         cfg.Api.BindAddress,  // configure the bind address
		Handler:      ch(sm),            // set the default handler
		ErrorLog:     stdlog.Default(),            // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		level.Info(logger).Log("event", "starting server", "serving", cfg.Api.BindAddress)
		errs <- s.ListenAndServe()
	}()

	go func(shutdown chan error) {
		systemSignalChannel := make(chan os.Signal, 1)
		signal.Notify(systemSignalChannel, syscall.SIGINT, syscall.SIGTERM)
		sig := <-systemSignalChannel // wait on ctrl-c
		shutdown <- fmt.Errorf("%s", sig)
	}(errs)
	level.Info(logger).Log("event", "shutdown requested", "cause", <-errs)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
	Shutdown()

	os.Exit(0)
}

