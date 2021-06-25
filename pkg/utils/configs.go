package utils

/*
  utils/configContext.go:

  Application Configs
*/

import (
	goflag "flag"
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"path/filepath"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
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

	ApiConfig struct {
		BindAddress   string
	}

	Config struct {
		Title   string
		RunMode string
		Version string
		Dbc     DBConfig
		Mqc     MQTTConfig
		Api     ApiConfig
		Logger  log.Logger
	}
)

/*
 * init()
 * using it to avoid test failures caused by Viper / "flag redefined" panic
 *	i.e. flag.StringVar(&configPath, "config", envConfigPath, "path to config file")
 *  cannot define 'config' twice, so we do it once in init()
 */
var (
	configPath string
	debugFlag bool
	envConfigPath string
	bindAddress string
)

func init() {
	fmt.Println("Configs.go init() called")
	pflag.StringVar(&configPath, "config", "", "path to config file")
	pflag.BoolVar(&debugFlag, "debug", false, "enable debug logging : default is true")
	pflag.StringVar(&bindAddress, "bind", "localhost:9090", "Bind address for the server, ex: localhost:9090")

	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	fmt.Printf("Configs.go init() completed:  %s \n", configPath)
}

func handleCommandLineParams() (string, bool) {
	envConfigPath = os.Getenv("HOMIE_SERVICE_CONFIG_FILE")
	if "" == envConfigPath {
		envConfigPath = "mqtt-config"
	}

	pflag.Parse()

	if pflag.Lookup("config") != nil {
		configPath = pflag.Lookup("config").Value.String()
		if configPath == "" {
			configPath = envConfigPath
		}
	} else {
		configPath = envConfigPath
	}
	fmt.Printf("event=calling handleCommandLineParams(), Config=%s, Environment=%s, debug=%v \n", configPath, envConfigPath, debugFlag)
	return configPath, debugFlag
}

func BuildConfigForCLI(configFilename string, log log.Logger) (*viper.Viper, error) {

	cfg := viper.New()
	cfg.SetConfigName(configFilename) // name of config file (without extension)
	cfg.SetConfigType("yaml")                     // REQUIRED if the config file does not have the extension in the name
	cfg.AddConfigPath("../../config/")            // path to look for the config file in
	cfg.AddConfigPath("../config/")            // path to look for the config file in
	cfg.AddConfigPath("./config/")                // path to look for the config file in
	cfg.AddConfigPath(".")                        // path to look for the config file in

	cfg.SetDefault("homiemonitor.datasources.demoNetworks", []string{"sknSensors"})
	cfg.SetDefault("homiemonitor.datasources.firmwareStorage", "./dataDB/firmwares")
	cfg.SetDefault("homiemonitor.datasources.dataStorage", "./dataDB/dataDir/devices.db")
	cfg.SetDefault("homiemonitor.datasources.demoSource", "./dataDB/demoData/mosquitto.log")

	cfg.SetDefault("homiemonitor.mqtt.homiesubscriptiontopic", "sknSensors/#")
	cfg.SetDefault("homiemonitor.mqtt.homiediscoverytopic", "+/+/$name")
	cfg.SetDefault("homiemonitor.mqtt.lwttopic", "sknSensors/$broadcast/LWT")
	cfg.SetDefault("homiemonitor.mqtt.lwtmsg", "HomieMonitor Offline!")
	cfg.SetDefault("homiemonitor.mqtt.port", 1883)

	cfg.SetDefault("homiemonitor.api.bindAddress", "localhost:9090")

	err := cfg.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		err = fmt.Errorf("fatal error config file: %s \n", err.Error())
		return nil, err
	}

	// Overwrite config from ENV is present
	cfg.SetEnvPrefix("HS")
	if err = cfg.BindEnv("firmware_storage"); err == nil {
		cfg.Set("homiemonitor.datasources.firmwareStorage", cfg.Get("firmware_storage")) // ./dataDB/firmwares
	}
	if err = cfg.BindEnv("data_storage"); err == nil {
		cfg.Set("homiemonitor.datasources.dataStorage", cfg.Get("data_storage")) // ./dataDB/dataDir/devices.db
	}
	if err = cfg.BindEnv("demo_source"); err == nil {
		cfg.Set("homiemonitor.datasources.demoSource", cfg.Get("demo_source")) // ./dataDB/demoData/mosquitto.log
	}

	cfg.SetEnvPrefix("mqtt")
	if err = cfg.BindEnv("broker"); err == nil {
		cfg.Set("homiemonitor.mqtt.broker", cfg.Get("broker"))
	}
	if err = cfg.BindEnv("username"); err == nil {
		cfg.Set("homiemonitor.mqtt.username", cfg.Get("username"))
	}
	if err = cfg.BindEnv("password"); err == nil {
		cfg.Set("homiemonitor.mqtt.userpassword", cfg.Get("password"))
	}
	if err = cfg.BindEnv("discovery_topic"); err == nil {
		cfg.Set("homiemonitor.mqtt.homiediscoverytopic", cfg.Get("discovery_topic"))
	}
	if err = cfg.BindEnv("subscription_topic"); err == nil {
		cfg.Set("homiemonitor.mqtt.homiesubscriptiontopic", cfg.Get("subscription_topic"))
	}

	cfg.SetEnvPrefix("api")
	if err = cfg.BindEnv("bind_address"); err == nil {
		cfg.Set("homiemonitor.api.bindAddress", cfg.Get("bind_address"))
	}

	//envHostURL := os.Getenv("BIND_ADDRESS", false, ":9090", "Bind address for the server")


	return cfg, err
}

func BuildLogger(moduleName string, debugOn bool) log.Logger {
	var opt level.Option
	if debugOn {
		opt = level.AllowDebug()
	} else {
		opt = level.AllowInfo()
	}
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = level.NewFilter(logger, opt)  // set log level
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"module", moduleName,
		"time:", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)
	level.Debug(logger).Log("event", "called buildLogger()")

	return logger
}

func BuildAppConfig(cfg *viper.Viper, log log.Logger) Config {
	level.Debug(log).Log("event", "calling buildAppConfig()")

	return Config{
		Title:   cfg.GetString("homiemonitor.title"),
		RunMode: cfg.GetString("homiemonitor.runmode"),
		Version: cfg.GetString("homiemonitor.version"),
		Logger:  log,
		Dbc: DBConfig{
			DemoSource:      filepath.FromSlash(cfg.GetString("homiemonitor.datasources.demoSource")),
			DataStorage:     filepath.FromSlash(cfg.GetString("homiemonitor.datasources.dataStorage")),
			FirmwareStorage: filepath.FromSlash(cfg.GetString("homiemonitor.datasources.firmwareStorage")),
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
		Api: ApiConfig{
			BindAddress: cfg.GetString("homiemonitor.api.bindAddress"),
		},
	}
}

func BuildRuntimeConfig(moduleName string) (Config, error) {
	cfile, debug := handleCommandLineParams()
	logger := BuildLogger(moduleName, debug)
	cliConfig, err := BuildConfigForCLI(cfile, logger)
	appConfig := BuildAppConfig(cliConfig, logger)

	return appConfig, err
}
