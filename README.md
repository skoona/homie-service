# homie-service
<a href="https://homieiot.github.io/">
  <img src="https://homieiot.github.io/img/works-with-homie.png" alt="works with MQTT Homie">
</a>

Rewrite of HomieMonitor in GOlang.


This application is designed to act as a `Homie Controller`, or `Monitor`, 
in support of IOT/Devices using [Homie-esp8266](https://github.com/homieiot/homie-esp8266); although any `Homie Device` implementation should be supported.

#### References: 
* [Homie: An MQTT Convention for IOT/M2M](https://homieiot.github.io/specification/)
* [Homie-ESP8266 Example of RCWL-0516 Microwave Presence Detector and DHT22 Temperature and Humidity sensors](https://github.com/skoona/EnvironmentMonitor_DHT)

## Status
* Support Mutliple Homie Networks (auto discovery); /homie/#, /custom/#, ...
* [MQTT](https://github.com/eclipse/paho.mqtt.golang) Unsecured connection.
* Produce MQTT messages decoded to Homie Device Model.
* Produce MQTT messages from MQTT Logfile for Demo use.
* [bBoltDB](https://github.com/boltdb/bolt) data storage for decoded device messages
* CoreLogic 
* * decode and transform data into Homie's Broadcast, Device, Node, Property, and Attribute collections.
* * encode firmware information into the Firmware Scheduling Service
* * implement firmware schedule
* * implement data retention on change service
* Create Web UI with WedSocket Driven Components
* Enable SSL/TLS Connection to MQTT
* Enabled OTA Scheduling

### Project Tree
```
├── LICENSE
├── README.md
├── bin
├── cmd
│   └── cli
│       └── main.go
├── config
│   ├── demo-config.yml
│   └── mqtt-config.yml
├── dataDB
│   ├── dataDir
│   │   └── devices.db
│   ├── demoData
│   │   └── mosquitto.log
│   └── firmwares
│       ├── Environment-DS18B20.bin
│       ├── Monitor-DHT-RCWL-Metrics-200.bin
│       └── Monitor-SHT31-RCWL-Metrics-200.bin
├── demo-runtime.log
├── docs
│   └── notes.md
├── go.mod
├── go.sum
├── internal
│   ├── demoProvider
│   │   └── service.go
│   ├── mqttProvider
│   │   └── service.go
│   ├── deviceSource
│   │   ├── service.go
│   │   └── usecase.go
│   ├── deviceStorage
│   │   ├── repository.go
│   │   └── usecase.go
│   ├── deviceScheduler
│   │   ├── firmwares.go
│   │   ├── schedules.go
│   │   ├── service.go
│   │   └── usecase.go
│   ├── deviceCore
│   │   ├── broadcasts.go
│   │   ├── devices.go
│   │   ├── domain.go
│   │   ├── events.go
│   │   ├── networks.go
│   │   ├── service.go
│   │   └── usecase.go
│   └── utils
│       └── configs.go
├── main
└── mqtt-runtime.log

18 directories, 34 files
```

#### Package Description
* (A) demoProvider
  - Converts Mosquitto logfile into device messages for deviceSource
* (A) mqttProvider
  - Captures Mosquitto messages into device messages for deviceSource
* (A) deviceStorage
  - Stores device messages from deviceSource to LevelDB/boltDB 
* (A) deviceScheduler
  - Schedules firmware OTA updates to network devices thru deviceSource
* (B) deviceSource
  - Device message manager for sourcing device info for storage, scheduling, and network device collection
* (C) deviceCore
  - Network device collection of all known/active devices on any Homie based network.  Fully managed with adds, updates, deletes, queries, firmware uploads, and ota schedule creation
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
* MQTT_HOST
* MQTT_USERNAME
* MQTT_PASSWORD

Optional:
* MQTT_SUBSCRIPTION_TOPIC
* MQTT_DISCOVERY_TOPIC


#### Configuration Files: ./config/demo-config.yml or live-config.yml
* ./config/<anyname>.yml     -- any named YAML file can be used via command-line options
* ./config/demo-config.yml   -- demo version, uses an internal mqtt.log file versus MQTT
* ./config/live-config.yml   -- default configuration file

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
    demoNetworks:
    - sknSensors
    - homie
  
```

#### Command-line Help
* program -h
```
Usage of ./program
  -config string
        path to config file (default "live-config")
```

* program 
```
$ go build cmd/cli/main.go 
$ ./program --config mqtt-config
$ ./program --config demo-config
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