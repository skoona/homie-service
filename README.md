# homie-service
<a href="https://homieiot.github.io/">
  <img src="https://homieiot.github.io/img/works-with-homie.png" alt="works with MQTT Homie">
</a>

Rewrite of [HomieMonitor](https://github.com/skoona/homieMonitor) in Golang.

Application implements the `Homie Controller`, or `Monitor`, supporting development of IOT/Devices 
using [Homie-esp8266](https://github.com/homieiot/homie-esp8266); although any `Homie Device` specification should be supported.

#### References: 
* [Homie: An MQTT Convention for IOT/M2M](https://homieiot.github.io/specification/)
* [Homie-ESP8266 Example of RCWL-0516 Microwave Presence Detector and DHT22 Temperature and Humidity sensors](https://github.com/skoona/EnvironmentMonitor_DHT)

## Status
* [Fyne.io](https://fyne.io) GUI 75% complete
* [go-kit/kit](https://github.com/go-kit/kit) Restful JSON API (with swagger docs) 95% complete  
* Support Mutliple Homie Networks (auto discovery); /homie/#, /custom/#, ...
* [MQTT](https://github.com/eclipse/paho.mqtt.golang) Unsecured connection.
* Produce MQTT messages decoded to internal Device Model.
* Produce MQTT messages from MQTT logfile for Demo use.
* [bBolt DB](https://github.com/etcd-io/bbolt) data storage for decoded device messages, and ota schedules
* CoreLogic 
* * decode and transform data into Homie's Broadcast, Device, Node, Property, and Attribute collections.
* * encode firmware information into the Firmware OTA Scheduling Service
* * implement firmware schedule
* * implement data retention on change service
* Create Web UI with WedSocket Driven Components (TBD)
* Enable SSL/TLS Connection to MQTT (TBD)
* Enabled OTA Scheduling

### Project Tree
```
<homie-service>.
├── LICENSE
├── Makefile
├── README.md
├── bin
├── cmd
│   ├── api
│   │   └── main.go
│   ├── base
│   │   └── main.go
│   └── gui
│       ├── main.go
│       └── svgImages.go
├── config
│   ├── demo-config.yml
│   ├── mqtt-config.yml
│   └── test-config.yml
├── dataDB
│   ├── dataDir
│   │   └── storage.db
│   ├── demoData
│   │   └── mosquitto.log
│   ├── firmwares
│   │   ├── Environment-DS18B20.bin
│   │   ├── Monitor-DHT-RCWL-Metrics-200.bin
│   │   └── Monitor-SHT31-RCWL-Metrics-200.bin
│   └── testData
│       ├── mosquitto_test.log
│       └── storage_test.db
├── demo-runtime.log
├── mqtt-runtime.log
├── devNotes
│   └── dev-notes.md
├── go.mod
├── go.sum
├── pkg
│   ├── UIAdapters
│   │   ├── api
│   │   │   ├── docs
│   │   │   │   └── docs.go
│   │   │   └── handlers
│   │   │       ├── broadcasts.go
│   │   │       ├── controller.go
│   │   │       ├── files.go
│   │   │       ├── firmwares.go
│   │   │       ├── middlewares.go
│   │   │       ├── networks.go
│   │   │       ├── schedules.go
│   │   │       └── validation.go
│   │   └── gui
│   │       ├── components
│   │       │   ├── cards.go
│   │       │   ├── forms.go
│   │       │   ├── homieTheme.go
│   │       │   ├── lists.go
│   │       │   ├── tables.go
│   │       │   ├── trees.go
│   │       │   └── utils.go
│   │       ├── providers
│   │       │   └── provider.go
│   │       └── views
│   │           ├── firmwares.go
│   │           ├── home.go
│   │           ├── mainPage.go
│   │           ├── networks.go
│   │           ├── schedules.go
│   │           ├── sites.go
│   │           └── views.go
│   ├── demoProvider
│   │   ├── demoProvider_suite_test.go
│   │   ├── demo_test.go
│   │   ├── mqttMock.go
│   │   ├── service.go
│   │   ├── usecase-devices.go
│   │   └── usecase-ota.go
│   ├── deviceCore
│   │   ├── core.go
│   │   ├── core_test.go
│   │   ├── deviceCore_suite_test.go
│   │   ├── service.go
│   │   ├── streams.go
│   │   └── usecase.go
│   ├── deviceScheduler
│   │   ├── deviceScheduler_suite_test.go
│   │   ├── scheduler.go
│   │   ├── scheduler_test.go
│   │   ├── service.go
│   │   └── usecase.go
│   ├── deviceSource
│   │   ├── service.go
│   │   ├── sources.go
│   │   └── usercase.go
│   ├── deviceStorage
│   │   ├── deviceStorage_suite_test.go
│   │   ├── repository.go
│   │   ├── storage_test.go
│   │   └── usecase.go
│   ├── mqttProvider
│   │   ├── mqttProvider_suite_test.go
│   │   ├── service.go
│   │   ├── service_test.go
│   │   ├── usecase-devices.go
│   │   └── usecase-ota.go
│   ├── services
│   │   └── service.go
│   └── utils
│       ├── colors.go
│       ├── configs.go
│       ├── configs_test.go
│       ├── imageUtils.go
│       ├── svgImages.go
│       ├── utils.go
│       ├── utils_suite_test.go
│       └── utils_test.go
├── svgResources
│   ├── homie-banner.png
│   ├── insertDriveFile-mdo-24px.svg
│   ├── insertDriveFile-mdr-24px.svg
│   ├── notificationAlert-mbo-24px.svg
│   ├── notificationAlert-mbr-24px.svg
│   ├── sensorsOff-mbo-24px.svg
│   ├── sensorsOff-mbr-24px.svg
│   ├── sensorsOn-mbo-24px.svg
│   ├── sensorsOn-mbr-24px.svg
│   ├── thumbsDown-mdo-24px.svg
│   ├── thumbsDown-mdr-24px.svg
│   ├── thumbsDown-mds-24px.svg
│   ├── thumbsUp-mdo-24px.svg
│   ├── thumbsUp-mdr-24px.svg
│   ├── thumbsUp-mds-24px.svg
│   ├── timeLapse-mbo-24px.svg
│   ├── timeLapse-mbr-24px.svg
│   ├── toggleOff-mbo-24px.svg
│   ├── toggleOff-mbr-24px.svg
│   ├── toggleOn-mbo-24px.svg
│   ├── toggleOn-mbr-24px.svg
│   ├── wall-clock-svgrepo-com.svg
│   └── works-with-homie-svg
└── swagger.yaml

```

#### Package Description
* (A) demoProvider
  - mock implementation of MQTT Broker using mqtt interfaces
  - Converts Mosquitto logfile into device messages for deviceSource
  - Intf: SteamProvider
  - Intf: OTAInteractor
* (A) mqttProvider
  - Captures Mosquitto messages into device messages for deviceSource
  - Intf: SteamProvider
  - Intf: OTAInteractor
* (A) deviceStorage
  - Stores device messages from deviceSource to LevelDB/boltDB 
  - Intf: Repository
* (B) deviceSource
  - Device message manager for sourcing device info for storage, scheduling, and network device collection
  - Intf: SteamProvider
  - Intf: DeviceEventProvider
* (B) deviceScheduler
  - Schedules firmware OTA updates to network devices thru deviceSource
  - Intf: OTAInteractor
  - Intf: SchedulerProvider
* (C) deviceCore
  - Network device collection of all known/active devices on any Homie based network.  Fully managed with adds, updates, deletes, queries, firmware uploads, and ota schedule creation
  - Intf: SchedulerProvider
  - Intf: DeviceEventProvider
  - Intf: CoreService
* (0) utils
  - Configuration and misc utilities
* (TBD) uiApi
  - JSON API for http implementing deviceCore for device network interactions
* (TBD) uiHtml
  - HTML ui implementing deviceCore for device network interactions
  - Potentially with Websocks dynamic updates
* (TBD) uiCli
  - Command line interface allowing data extraction from core network models
  
### Environment Variables and Configuration
MQTT hostname, username, and password can be set in the environment.  If present they will override those in the configuration file.  The variable are:

#### Environment Vars: ~/.profile or ~/.bashrc
* MQTT_BROKER
* MQTT_USERNAME
* MQTT_PASSWORD

Optional:
* MQTT_SUBSCRIPTION_TOPIC
* MQTT_DISCOVERY_TOPIC
* HS_FIRMWARE_STORAGE
* HS_DATA_STORAGE
* HS_DEMO_SOURCE
* HOMIE_SERVICE_CONFIG_FILE   
  - choices: mqtt-config, demo-config, test-config

#### Configuration Files: ./config/demo-config.yml or live-config.yml
* ./config/<anyname>.yml     -- any named YAML file can be used via command-line options
* ./config/demo-config.yml   -- demo version, uses an internal mqtt.log file versus MQTT
* ./config/mqtt-config.yml   -- default configuration file
* ./config/test-config.yml   -- test configuration

#### Config file priority
1. Commandline (--config demo-config)
2. Environment (HOMIE_SERVICE_CONFIG_FILE)
3. Default ('mqtt-config')
#### Config params priority
1. Environment (HS_*, MQTT_*)
2. Config file contents (...)

#### Configuration file Contents
Live Mode: Reads data from MQTT
```
---
# Live Mode
homiemonitor:
  title: Homie Network Monitor
  runmode: live
  version: 0.4.0

  datasources: 
    dataStorage: ./dataDB/dataDir/devices.db
    demoSource: ./dataDB/demoData/mosquitto.log
    firmwareStorage: ./dataDB/firmwares
    demoNetworks:
    - sknSensors
    - homie

  mqtt:
    clientid: homie-monitor
    broker: <mqtt-host-or-ip>
    port: 1883
    username: <mqtt-username>
    userpassword: <mqtt-password>
    homiediscoverytopic: "+/+/$name"
    homiesubscriptiontopic: "sknSensors/#"
    lwttopic: sknSensors/$broadcast/LWT
    lwtmsg: HomieMonitor Offline!

```

Demo Mode: Reads local MQTT logfile as input
```
---
# Demo Config
homiemonitor:
  title: Homie Network Monitor
  runmode: live
  version: 0.4.0

  datasources: 
    dataStorage: ./dataDB/dataDir/devices.db
    demoSource: ./dataDB/demoData/mosquitto.log
    firmwareStorage: ./dataDB/firmwares
    demoNetworks:
    - sknSensors
    - homie
  
```

#### Command-line Help
* program -h
```
Usage of ./program
  --config string
        path to config file (default "live-config")
  --debug true
        turn on debug messaging 
```

* API program 
```
$ go build cmd/api/main.go 
$ ./program --config mqtt-config
$ ./program --config demo-config
 ...
$ curl http://localhost:9090/docs
```
* GUI program
```
$ go build cmd/gui/main.go 
$ ./main --config mqtt-config
$ ./main --config demo-config
```
* basic program
```
$ go build cmd/api/main.go 
$ ./main --config mqtt-config --debug=true
$ ./main --config demo-config --debug=true
 no UI
```

### Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

Find me on Gitter!

### LICENSE
The application is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).