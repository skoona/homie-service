package configs

/*
  utils/configContext.go:

  Application Configs
*/

import (
	"context"
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
		demoSource  string
		dataStorage string
	}

	MQTTConfig struct {
		clientID          string
		brokerIP          string
		brokerPort        int
		userName          string
		userPassword      string
		discoveryTopic    string
		subscriptionTopic string
		lwtTopic          string
		lwtMessage        string
	}

	Config struct {
		title   string
		runMode string
		version string
		dbc     DBConfig
		mqc     MQTTConfig
		logger  *log.Logger
	}
)

func buildConfigFromCLI(logger log.logger) *viper.Viper {
	level.Info(logger).Log("msg", "calling configSetup()")

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
	config.SetDefault("homiemonitor.mqtt.port", 1883)

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

var logger log.Logger

func buildLogger(moduleName string) *log.Logger {

	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"module", moduleName,
		"time:", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)

	level.Debug(logger).Log("msg", "logging service started")

	return &logger
}

func buildAppConfig(cfg *viper.Viper, log log.logger) *Config {
	return &Config{
		title:   cfg.GetString("homiemonitor.title"),
		runMode: cfg.GetString("homiemonitor.runmode"),
		version: cfg.GetString("homiemonitor.version"),
		logger:  log,
		dbc: DBConfig{
			demoSource:  cfg.GetString("homiemonitor.datasources.demoSource"),
			dataStorage: cfg.GetString("homiemonitor.datasources.dataStorage"),
		},
		mqc: MQTTConfig{
			clientID:          cfg.GetString("homiemonitor.mqtt.clientid"),
			brokerIP:          cfg.GetString("homiemonitor.mqtt.broker"),
			brokerPort:        cfg.GetInt("homiemonitor.mqtt.port"),
			userName:          cfg.GetString("homiemonitor.mqtt.username"),
			userPassword:      cfg.GetString("homiemonitor.mqtt.userpassword"),
			discoveryTopic:    cfg.GetString("homiemonitor.mqtt.homiediscoverytopic"),
			subscriptionTopic: cfg.GetString("homiemonitor.mqtt.homiesubscriptiontopic"),
			lwtTopic:          cfg.GetString("homiemonitor.mqtt.lwttopic"),
			lwtMessage:        cfg.GetString("homiemonitor.mqtt.lwtmsg"),
		},
	}
}

func BuildRuntimeConfigAndContext(moduleName string) context.Context {
	logger := buildLogger(moduleName)
	cliConfig := buildConfigFromCLI(logger)

	appConfig := buildAppConfig(cliConfig, logger)
	appCtx := context.Background()
	appCtx = appCtx.WithValue(appCtx, "DBConfig", appConfig.dbc)
	appCtx = appCtx.WithValue(appCtx, "MQTTConfig", appConfig.mqc)

	return appCtx.WithValue(appCtx, "appConfig", appConfig)
}
