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
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func configSetup() *viper.Viper {
	level.Info(logger).Log("msg", "configSetup called")

	var configPath string

	config := viper.New()

	flag.StringVar(&configPath, "config", "config/live-config.yml", "path to config file")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	config.BindPFlags(pflag.CommandLine)

	config.SetConfigName(config.GetString("config")) // name of config file (without extension)
	config.SetConfigType("yaml")                     // REQUIRED if the config file does not have the extension in the name
	config.AddConfigPath("./config/")                // path to look for the config file in
	config.AddConfigPath(".")                        // path to look for the config file in

	config.SetDefault("homiemonitor.mqtt.homiesubscriptiontopic", "sknSensors/#")
	config.SetDefault("homiemonitor.mqtt.homiediscoverytopic", "+/+/$name")
	config.SetDefault("homiemonitor.mqtt.lwttopic", "sknSensors/$broadcast/LWT")
	config.SetDefault("homiemonitor.mqtt.lwtmsg", "HomieMonitor Offline!")

	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	// Overwrite config from ENV is present
	config.SetEnvPrefix("mqtt")
	if err = config.BindEnv("broker"); err == nil {
		config.Set("homiemonitor.mqtt.broker", config.Get("broker"))
	}
	if err = config.BindEnv("username"); err == nil {
		config.Set("homiemonitor.mqtt.username", config.Get("username"))
	}
	if err = config.BindEnv("password"); err == nil {
		config.Set("homiemonitor.mqtt.password", config.Get("password"))
	}
	if err = config.BindEnv("discovery_topic"); err == nil {
		config.Set("homiemonitor.mqtt.homiediscoverytopic", config.Get("discovery_topic"))
	}
	if err = config.BindEnv("subscription_topic"); err == nil {
		config.Set("homiemonitor.mqtt.homiesubscriptiontopic", config.Get("subscription_topic"))
	}

	return config
}

func shutdownDemo() {
	level.Info(logger).Log("msg", "shutdownDemo called")
}

func shutdownLive() {
	level.Info(logger).Log("msg", "shutdownLive called")
}

func runLive(config *viper.Viper) error {
	level.Info(logger).Log("msg", "runLive called")
	return nil
}

func runDemo(config *viper.Viper) error {
	level.Info(logger).Log("msg", "runDemo called")
	return nil
}

var logger log.Logger

func main() {
	// var hns *cl.HomieNetworks
	var err error
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"main", "entry",
		"time:", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller)

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	/* Prepare for clean exit */
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	config := configSetup()

	// Run the App
	if config.GetString("homiemonitor.runmode") == "demo" {
		// demo
		err = runDemo(config)
	} else {
		// live
		err = runLive(config)
	}

	// wait on ctrl-c
	sig := <-sigs
	level.Info(logger).Log("msg", sig)
	level.Info(logger).Log("msg", "Shutting Down")

	if config.GetString("homiemonitor.runmode") == "demo" {
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
