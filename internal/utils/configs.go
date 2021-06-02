package configs

/*
  utils/configContext.go:

  Application Configs
*/

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type (
	DBConfig struct {
		DemoSource      string
		DataStorage     string
		FirmwareStorage string
		DemoNetworks    []string
	}

	MQTTConfig struct {
		ClientID          string
		BrokerIP          string
		BrokerPort        int
		UserName          string
		UserPassword      string
		DiscoveryTopic    string
		SubscriptionTopic string
		LwtTopic          string
		LwtMessage        string
	}

	Config struct {
		Title   string
		RunMode string
		Version string
		Dbc     DBConfig
		Mqc     MQTTConfig
		Logger  log.Logger
	}
)

func buildConfigForCLI(log log.Logger) *viper.Viper {
	level.Debug(log).Log("event", "calling buildConfigForCLI()")

	var configPath string
	envConfigPath := os.Getenv("HOMIE_SERVICE_CONFIG_FILE")  // "mqtt-config"
	if "" == envConfigPath {
		envConfigPath = "mqtt-config"
	}

	config := viper.New()

	flag.StringVar(&configPath, "config", envConfigPath, "path to config file")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	config.BindPFlags(pflag.CommandLine)

	config.SetConfigName(config.GetString("config")) // name of config file (without extension)
	config.SetConfigType("yaml")                     // REQUIRED if the config file does not have the extension in the name
	config.AddConfigPath("./config/")                // path to look for the config file in
	config.AddConfigPath(".")                        // path to look for the config file in

	config.SetDefault("homiemonitor.datasources.demoNetworks", []string{"sknSensors"})
	config.SetDefault("homiemonitor.datasources.firmwareStorage", "./dataDB/firmwares")
	config.SetDefault("homiemonitor.datasources.dataStorage", "./dataDB/dataDir/devices.db")
	config.SetDefault("homiemonitor.datasources.demoSource", "./dataDB/demoData/mosquitto.log")

	config.SetDefault("homiemonitor.mqtt.homiesubscriptiontopic", "sknSensors/#")
	config.SetDefault("homiemonitor.mqtt.homiediscoverytopic", "+/+/$name")
	config.SetDefault("homiemonitor.mqtt.lwttopic", "sknSensors/$broadcast/LWT")
	config.SetDefault("homiemonitor.mqtt.lwtmsg", "HomieMonitor Offline!")
	config.SetDefault("homiemonitor.mqtt.port", 1883)

	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	// Overwrite config from ENV is present
	config.SetEnvPrefix("HS")
	if err = config.BindEnv("firmware_storage"); err == nil {
		config.Set("homiemonitor.datasources.firmwareStorage", config.Get("firmware_storage")) // ./dataDB/firmwares
	}
	if err = config.BindEnv("data_storage"); err == nil {
		config.Set("homiemonitor.datasources.dataStorage", config.Get("data_storage")) // ./dataDB/dataDir/devices.db
	}
	if err = config.BindEnv("demo_source"); err == nil {
		config.Set("homiemonitor.datasources.demoSource", config.Get("demo_source")) // ./dataDB/demoData/mosquitto.log
	}

	config.SetEnvPrefix("mqtt")
	if err = config.BindEnv("broker"); err == nil {
		config.Set("homiemonitor.mqtt.broker", config.Get("broker"))
	}
	if err = config.BindEnv("username"); err == nil {
		config.Set("homiemonitor.mqtt.username", config.Get("username"))
	}
	if err = config.BindEnv("password"); err == nil {
		config.Set("homiemonitor.mqtt.userpassword", config.Get("password"))
	}
	if err = config.BindEnv("discovery_topic"); err == nil {
		config.Set("homiemonitor.mqtt.homiediscoverytopic", config.Get("discovery_topic"))
	}
	if err = config.BindEnv("subscription_topic"); err == nil {
		config.Set("homiemonitor.mqtt.homiesubscriptiontopic", config.Get("subscription_topic"))
	}

	return config
}

var DefaultLogger log.Logger

func buildLogger(moduleName string) log.Logger {

	logger := log.NewLogfmtLogger(os.Stderr)
	logger = level.NewFilter(logger, level.AllowDebug())  // set log level
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"module", moduleName,
		"time:", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)
	DefaultLogger = logger
	level.Debug(logger).Log("event", "called buildLogger()")

	return logger
}

func buildAppConfig(cfg *viper.Viper, log log.Logger) Config {
	level.Debug(log).Log("event", "calling buildAppConfig()")

	return Config{
		Title:   cfg.GetString("homiemonitor.title"),
		RunMode: cfg.GetString("homiemonitor.runmode"),
		Version: cfg.GetString("homiemonitor.version"),
		Logger:  log,
		Dbc: DBConfig{
			DemoSource:      cfg.GetString("homiemonitor.datasources.demoSource"),
			DataStorage:     cfg.GetString("homiemonitor.datasources.dataStorage"),
			FirmwareStorage: cfg.GetString("homiemonitor.datasources.firmwareStorage"),
			DemoNetworks:    cfg.GetStringSlice("homiemonitor.datasources.demoNetworks"),
		},
		Mqc: MQTTConfig{
			ClientID:          cfg.GetString("homiemonitor.mqtt.clientid"),
			BrokerIP:          cfg.GetString("homiemonitor.mqtt.broker"),
			BrokerPort:        cfg.GetInt("homiemonitor.mqtt.port"),
			UserName:          cfg.GetString("homiemonitor.mqtt.username"),
			UserPassword:      cfg.GetString("homiemonitor.mqtt.userpassword"),
			DiscoveryTopic:    cfg.GetString("homiemonitor.mqtt.homiediscoverytopic"),
			SubscriptionTopic: cfg.GetString("homiemonitor.mqtt.homiesubscriptiontopic"),
			LwtTopic:          cfg.GetString("homiemonitor.mqtt.lwttopic"),
			LwtMessage:        cfg.GetString("homiemonitor.mqtt.lwtmsg"),
		},
	}
}

func BuildRuntimeConfig(moduleName string) Config {
	logger := buildLogger(moduleName)
	cliConfig := buildConfigForCLI(logger)
	appConfig := buildAppConfig(cliConfig, logger)

	return appConfig
}
