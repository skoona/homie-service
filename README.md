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
* Support Mutliple Homie Networks; /homie/#, /custom/#, ...
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