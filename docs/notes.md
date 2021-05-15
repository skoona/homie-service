# HomieExplorer Tech Notes

## Work flows
* From infra.MQTT --> intf.MQTT --> usec.MQTT --> Domain
* From infra.WebAPI --> intf.WebAPI --> usec.WebAPI --> Domain
* From infra.Services --> intf.Service --> intf.MQTT --> infra.MQTT
* --
* From Web to Core
* From Core{Scheduler} to MQTT
* From Core{Set-Commands} to MQTT
* From Core{Database} to bBoltDB

## Code Structure
* domains{core}
* use-cases{services}
* interfaces{DB-Out/Repository}
* interfaces{WebApi-In/Controller}
* interfaces{MQTT-In/Controller}
* infrastructure{DBHandler/Impl}
* infrastructure{WebApi/Router}
* infrastructure{MQTT/Client}
* main.go{configuration/startup/shutdown}

```
Domain <- UseCases <- Interfaces <- Infrastructure: <== Main(Dependency Injection)
```

## Project Tree
```
❯ homieExplorer
.
├── LICENSE
├── README.md
├── go.mod
├── go.sum
├── bin
├── config
│   ├── demo-config.yml
│   └── live-config.yml
├── dataDB
│   ├── demo
│   │   └── homiedemo.log
│   ├── firmwares
│   │   ├── Environment-DS18B20.bin
│   │   ├── Monitor-DHT-RCWL-Metrics-200.bin
│   │   └── Monitor-SHT31-RCWL-Metrics-200.bin
│   └── live
├── docs
│   └── notes.md
└── src
    ├── main.go
    ├── entities
    │   ├── domain.go
    │   ├── devices.go
    │   ├── firmwares.go
    │   ├── networks.go
    │   └── schedules.go
    ├── infrastructure
    │   ├── db
    │   │   ├── bbolt-db.go
    │   │   └── db.go
    │   ├── router
    │   │   ├── mux-router.go
    │   │   └── router.go
    │   └── stream
    │       ├── stream.go
    │       ├── mqtt-stream.go
    │       └── stream.go
    ├── interfaces
    │   ├── controllers
    │   │   ├── controller.go
    │   │   ├── configuration-controller.go
    │   │   ├── firmwares-controller.go
    │   │   ├── networks-controller.go
    │   │   └── schedule-controller.go
    │   ├── repository
    │   │   ├── repository.go
    │   │   ├── firmware-repository.go
    │   │   ├── networks-repository.go
    │   │   └── schedules-repository.go
    │   └── services
    │       ├── services.go
    │       └── stream-service.go
    └── usecases
        ├── usecases.go
        ├── uc-devices.go
        ├── uc-firmwares.go
        ├── uc-networks.go
        └── uc-schedules.go

```

### General Notes:
#### Generate Interface Stubs
    https://github.com/josharian/impl

    $ impl 'dm *DeviceMessage' github.com/skoona/homieExplorer/src/services/datasources.DeviceEventIntf

#### MQTT Notes
    MQTT Spec: http://docs.oasis-open.org/mqtt/mqtt/v3.1.1/os/mqtt-v3.1.1-os.html

*   QoS 0: At most once delivery  CAUTION: PUBLISHING QOS=0 IS NOT ACKNOWLEDGED!!!
*   QoS 1: At least once delivery
*   QoS 2: Exactly once delivery

#### Starting a new project
    go mod init github.com/skoona/<project-dir>

#### Importing
    All files in dir are of the same package

    import "github.com/skoona/<project>/src/<dir-name>   -- all files in dir are of the same package
    or
    import <short-name> "github.com/skoona/<project>/src/<dir-name>

#### Building
    $ go build src/main.go
    $ go clean --cache
    $ go run


#### Specifications
[Homie Specification Details](https://homieiot.github.io/specification/#toc-0)
[Homie ESP8266 Implementation Details](https://homieiot.github.io/homie-esp8266/docs/develop/others/homie-implementation-specifics/)
```
/*
 * Base / Device / Node / Properties / Attributes
 * Device  -> The physical thing
 * Devices -> A device can expose multiple nodes
 * Nodes   ->  A node can have multiple properties
 * Properties ->  Devices, nodes and properties have (multiple) specific  
 *                attributes characterizing them.
 * Attributes ->  Represented by topic identifier starting with $
 * 
 * homie / $broadcast / level : message
 * homie / device ID  / $state : last      (LWT)
 * homie / device ID  / $device-attribute
 * homie / device ID  / $device-attribute / property 
 * homie / device ID  / $device-attribute / property   / property
 *   0       1            2                   3            4
 * homie / device ID  / node ID           / $node-attribute
 * homie / device ID  / node ID           / property ID
 * homie / device ID  / node ID           / property ID / set-command
 * homie / device ID  / node ID           / property ID / $property-attribute
 *
 * Ignore h/d/n/p/set messages
 * Ignore firmware in progress messages h/d/$implementation/ota/firmware
 *
 * Networks, Devices, and Nodes must be created as needed. The have no Key/Value
 *  relationships.  They are containers for:
 *   - Attributes (device-attrs, node-attrs, and prop-attrs)
 *   - Properties (node-props)
 *   - PropertyAttributes (prop-attr)
 *
 * $device-attributes can be 3 parts verus 1
 * homie / device ID  / $implmentation / ota    / enabled : true
 * homie / device ID  / $stats         / signal           : 92
 * homie / device ID  / $localip                          : 10.100.1.x
 *
 ** Currently $device-attributes longer than 1 are combined as single key
 * homie / device ID  / ($implmentation/ota/enabled)      : true
 * 
 *
 ** Special MQTT Message formats
 * BROADCAST MESSAGE FORMAT (low)
 * homie / $broadcast / level : message
 *
 * OTA RESPONSE MESSAGE FORMAT (high)
 * homie / device ID  / $implementation / ota    / status   : 206 written/total
 *
 * OTA FIRMWARE DOWNLOADS MESSAGE FORMAT
 * homie / device ID  / $implementation / ota    / firmware : <md5>
 *
 * CONFIGURATION CHANGE MESSAGE FORMAT
 * homie / device ID  / $implementation / config / set      : <json>
 * 
 */

/*
 *     0         1           2                   3                 4
 * sknSensors/device/node/version: 3.0.0
 *  * homie / device ID / $device-attribute
 *              NetworkBucket(Bucket(topic:value))
 *
 *  * homie / device ID / $device-attribute / attribute-property 
 *              NetworkBucket(Bucket(Bucket(topic:value)))
 *
 *  * homie / device ID / $device-attribute / attribute-property / attr-prop-attr
 *				NetworkBucket(Bucket(Bucket(Bucket(topic:value))))
 *
 *  * homie / device ID / node ID / $node-attribute
 *				NetworkBucket(Bucket(Bucket(topic:value)))
 *
 *  * homie / device ID / node ID / node-property
 *				NetworkBucket(Bucket(Bucket(topic:value)))
 *
 *  * homie / device ID / node ID /      property ID /        $property-attribute
 *				NetworkBucket(Bucket(Bucket(Bucket(topic:value))))
*/
```

Database Console log od a single device
```
2021/02/22 14:19:07     [11] Device: OutsideMonitor
2021/02/22 14:19:07              $fw/checksum --> 615fed382ab44bd43fe83508aecac682
2021/02/22 14:19:07              $fw/name --> Monitor-SHT31-RCWL-Metrics
2021/02/22 14:19:07              $fw/version --> 2.0.0
2021/02/22 14:19:07              $homie --> 3.0.1
2021/02/22 14:19:07              $implementation --> esp8266
2021/02/22 14:19:07              $implementation/config --> {"name":"Outside Monitor","device_id":"OutsideMonitor","device_stats_interval":900,"wifi":{"ssid":"SFNSS1-24G"},"mqtt":{"host":"openhab.skoona.net","port":1883,"base_topic":"sknSensors/","auth":true},"ota":{"enabled":true},"settings":{"sensorsInterval":900,"motionHoldInterval":60}}
2021/02/22 14:19:07              $implementation/ota/enabled --> true
2021/02/22 14:19:07              $implementation/version --> 3.0.0
2021/02/22 14:19:07              $localip --> 10.100.1.182
2021/02/22 14:19:07              $mac --> 18:FE:34:FD:8C:1B
2021/02/22 14:19:07              $name --> Outside Monitor
2021/02/22 14:19:07              $nodes --> Ambient,Presence,hardware
2021/02/22 14:19:07              $state --> ready
2021/02/22 14:19:07              $stats --> uptime
2021/02/22 14:19:07              $stats/interval --> 905
2021/02/22 14:19:07              $stats/signal --> 92
2021/02/22 14:19:07              $stats/uptime --> 396023
2021/02/22 14:19:07              [Ambient]$name --> Temperature and Humidity Sensor
2021/02/22 14:19:07              [Ambient]$properties --> humidity,temperature
2021/02/22 14:19:07              [Ambient]$type --> sensor
2021/02/22 14:19:07              [Ambient][humidity]$datatype --> float
2021/02/22 14:19:07              [Ambient][humidity]$name --> Humidity
2021/02/22 14:19:07              [Ambient][humidity]$unit --> %rH
2021/02/22 14:19:07              [Ambient][humidity]humidity --> 66.71
2021/02/22 14:19:07              [Ambient][temperature]$datatype --> float
2021/02/22 14:19:07              [Ambient][temperature]$name --> Temperature
2021/02/22 14:19:07              [Ambient][temperature]$unit --> °F
2021/02/22 14:19:07              [Ambient][temperature]temperature --> 20.00
2021/02/22 14:19:07              [Presence]$name --> Motion Sensor
2021/02/22 14:19:07              [Presence]$properties --> motion
2021/02/22 14:19:07              [Presence]$type --> sensor
2021/02/22 14:19:07              [Presence][motion]$datatype --> enum
2021/02/22 14:19:07              [Presence][motion]$format --> OPEN,CLOSED
2021/02/22 14:19:07              [Presence][motion]$name --> Motion
2021/02/22 14:19:07              [Presence][motion]motion --> OPEN
2021/02/22 14:19:07              [hardware]$name --> Device Info
2021/02/22 14:19:07              [hardware]$properties --> signal,mac,resetReason,voltage
2021/02/22 14:19:07              [hardware]$type --> sensor
2021/02/22 14:19:07              [hardware][mac]$datatype --> sring
2021/02/22 14:19:07              [hardware][mac]$name --> mac
2021/02/22 14:19:07              [hardware][mac]mac --> 18:FE:34:FD:8C:1B
2021/02/22 14:19:07              [hardware][resetReason]$datatype --> string
2021/02/22 14:19:07              [hardware][resetReason]$name --> Last Reset Reason
2021/02/22 14:19:07              [hardware][resetReason]resetReason --> External System
2021/02/22 14:19:07              [hardware][signal]$datatype --> integer
2021/02/22 14:19:07              [hardware][signal]$name --> RSSI
2021/02/22 14:19:07              [hardware][signal]$unit --> dBm
2021/02/22 14:19:07              [hardware][signal]signal --> -54
2021/02/22 14:19:07              [hardware][voltage]$datatype --> float
2021/02/22 14:19:07              [hardware][voltage]$name --> 3.3V Supply
2021/02/22 14:19:07              [hardware][voltage]$unit --> V
2021/02/22 14:19:07              [hardware][voltage]voltage --> 3.44
```