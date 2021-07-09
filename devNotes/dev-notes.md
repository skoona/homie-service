# HomieExplorer Tech Notes

### Swagger build spec
    $  make
    $  swagger generate spec -o ./swagger.yaml --scan-models

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
 *   0       1            2                   3            4
 * homie / $broadcast / level : message
 * homie / device ID  / $state : last      (LWT)
 * homie / device ID  / $device-attribute
 * homie / device ID  / $device-attribute / property 
 * homie / device ID  / $device-attribute / property   / property
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
 ** Special MQTT Message formats
 * BROADCAST MESSAGE FORMAT (low)
 * homie / $broadcast / level : message
 *
 * OTA RESPONSE MESSAGE FORMAT (high)
 * homie / device ID  / $implementation / ota    / status   : 206 written/total
 *
 * OTA FIRMWARE DOWNLOADS MESSAGE FORMAT
 * homie / device ID  / $implementation / ota    / firmware / <md5> : <ota-transport-formatted>
 *
 * CONFIGURATION CHANGE MESSAGE FORMAT
 * homie / device ID  / $implementation / config / set      : <json>
 * 
 */

/*  boltDB Operations
 *     0         1           2                   3                 4
 * sknSensors/device/node/version: 3.0.0
 *  * homie / device ID / $device-attribute
 *    NetworkBucket(DeviceBucket(AttributeBucket(topic:value)))
 *
 *  * homie / device ID / $device-attribute / attribute-property 
 *    NetworkBucket(DeviceBucket(AttributeBucket(PropertyBucket(topic:value)))))
 *
 *  * homie / device ID / $device-attribute / attribute-property / attr-prop-attr
 *	  NetworkBucket(DeviceBucket(AttributeBucket(PropertyBucket(PropertyBucket(topic:value)))))
 *
 *  * homie / device ID / node ID / $node-attribute
 *	  NetworkBucket(DeviceBucket(NodeBucket(AttributeBucket(topic:value))))
 *
 *  * homie / device ID / node ID / node-property
 *	  NetworkBucket(DeviceBucket(NodeBucket(PropertyBucket(topic:value))))
 *
 *  * homie / device ID / node ID /      property ID /        $property-attribute
 *	  NetworkBucket(DeviceBucket(NodeBucket(PropertyBucket(AttributeBucket(topic:value)))))
*/
```

Database Console log of a single device
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
```json
{
  "Device": {
    "ID": "36ec4f3d3eedbe07b28e5fa6ab6ddb9b",
    "Title": "",
    "ElementType": 10,
    "OTAEnabled": true,
    "Parent": "sknSensors",
    "Name": "GarageMonitor",
    "Attrs": {
      "$fw": {
        "ID": "5817b977b7409ef2cf561ce29f4bdcdb",
        "ElementType": 11,
        "Parent": "GarageMonitor",
        "Name": "$fw",
        "Value": "",
        "Props": {
          "checksum": {
            "ID": "fe6b1c6a54ad924d80f355d1bf75c439",
            "ElementType": 12,
            "Parent": "$fw",
            "Name": "checksum",
            "Value": "615fed382ab44bd43fe83508aecac682",
            "Props": {}
          },
          "name": {
            "ID": "a657075ef75d1195e53fe8e20b275d14",
            "ElementType": 12,
            "Parent": "$fw",
            "Name": "name",
            "Value": "Monitor-SHT31-RCWL-Metrics",
            "Props": {}
          },
          "version": {
            "ID": "04781f846170e30259f618f54d01ef31",
            "ElementType": 12,
            "Parent": "$fw",
            "Name": "version",
            "Value": "2.0.0",
            "Props": {}
          }
        }
      },
      "$homie": {
        "ID": "79ee226e00a555813dc00fc179823921",
        "ElementType": 11,
        "Parent": "GarageMonitor",
        "Name": "$homie",
        "Value": "3.0.1",
        "Props": {}
      },
      "$implementation": {
        "ID": "28918a63addfd6bc3e75be47fb17f715",
        "ElementType": 11,
        "Parent": "GarageMonitor",
        "Name": "$implementation",
        "Value": "esp8266",
        "Props": {
          "config": {
            "ID": "4219e9c9ac2367a64e932df573592645",
            "ElementType": 12,
            "Parent": "$implementation",
            "Name": "config",
            "Value": "{\"name\":\"Garage Monitor\",\"device_id\":\"GarageMonitor\",\"device_stats_interval\":900,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"sensorInterval\":900,\"motionHoldInterval\":60}}",
            "Props": {}
          },
          "ota": {
            "ID": "3f46256b1699878ec144137e6602dc21",
            "ElementType": 12,
            "Parent": "$implementation",
            "Name": "ota",
            "Value": "",
            "Props": {
              "ota": {
                "ID": "0f3e50d8a8d10cef1cd8fc1def8acb2a",
                "ElementType": 13,
                "Parent": "ota",
                "Name": "enabled",
                "Value": "true"
              }
            }
          },
          "version": {
            "ID": "c3b88496b22cdf143def8442d90f75a2",
            "ElementType": 12,
            "Parent": "$implementation",
            "Name": "version",
            "Value": "3.0.0",
            "Props": {}
          }
        }
      },
      "$localip": {
        "ID": "8f2ed99b783adafad42243fbfd055443",
        "ElementType": 11,
        "Parent": "GarageMonitor",
        "Name": "$localip",
        "Value": "10.100.1.177",
        "Props": {}
      },
      "$mac": {
        "ID": "d836298c7743e12aacf721ee01efaa66",
        "ElementType": 11,
        "Parent": "GarageMonitor",
        "Name": "$mac",
        "Value": "B4:E6:2D:1B:5C:4D",
        "Props": {}
      },
      "$name": {
        "ID": "91c16f8f99059c72025c7ed3f15e0075",
        "ElementType": 11,
        "Parent": "GarageMonitor",
        "Name": "$name",
        "Value": "Garage Monitor",
        "Props": {}
      },
      "$nodes": {
        "ID": "bca6297465a2d9752db71a890563a8e4",
        "ElementType": 11,
        "Parent": "GarageMonitor",
        "Name": "$nodes",
        "Value": "Ambient,Presence,hardware",
        "Props": {}
      },
      "$state": {
        "ID": "3f88c3287bd4243247fd1e501bf2dd51",
        "ElementType": 11,
        "Parent": "GarageMonitor",
        "Name": "$state",
        "Value": "ready",
        "Props": {}
      },
      "$stats": {
        "ID": "fdd2e51ca3d87381f172ee313948b2bf",
        "ElementType": 11,
        "Parent": "GarageMonitor",
        "Name": "$stats",
        "Value": "uptime",
        "Props": {
          "interval": {
            "ID": "c569da83b8806917b24e96c8223d6aeb",
            "ElementType": 12,
            "Parent": "$stats",
            "Name": "interval",
            "Value": "905",
            "Props": {}
          },
          "signal": {
            "ID": "c1b3f57f589675e6efc4a571f8a94db6",
            "ElementType": 12,
            "Parent": "$stats",
            "Name": "signal",
            "Value": "54",
            "Props": {}
          },
          "uptime": {
            "ID": "8d3fac2566c604a9b6dc2dae5731b209",
            "ElementType": 12,
            "Parent": "$stats",
            "Name": "uptime",
            "Value": "592996",
            "Props": {}
          }
        }
      }
    },
    "Nodes": {
      "Ambient": {
        "ID": "c7da7f594ada6c9de0d59d60db38bc56",
        "ElementType": 14,
        "Parent": "GarageMonitor",
        "Name": "Ambient",
        "Attrs": {
          "$name": {
            "ID": "f8844bbbb57867e023539da7ae983bbe",
            "ElementType": 15,
            "Parent": "Ambient",
            "Name": "$name",
            "Value": "Temperature and Humidity Sensor"
          },
          "$properties": {
            "ID": "ab38f5261f392c6650496e63e5bb5008",
            "ElementType": 15,
            "Parent": "Ambient",
            "Name": "$properties",
            "Value": "humidity,temperature"
          },
          "$type": {
            "ID": "093fcbc870c04c6c01dde0151f8362b9",
            "ElementType": 15,
            "Parent": "Ambient",
            "Name": "$type",
            "Value": "sensor"
          }
        },
        "Props": {
          "humidity": {
            "ID": "11f6b15bfc89ef5646c42ba6742383b1",
            "ElementType": 16,
            "Parent": "Ambient",
            "Name": "humidity",
            "Value": "57.18",
            "Attrs": {
              "$datatype": {
                "ID": "053841b63999a23d464f003c5a4e68b1",
                "ElementType": 17,
                "Parent": "humidity",
                "Name": "$datatype",
                "Value": "float"
              },
              "$name": {
                "ID": "7a05d3fa3ce54025579710a3f6761e5e",
                "ElementType": 17,
                "Parent": "humidity",
                "Name": "$name",
                "Value": "Humidity"
              },
              "$unit": {
                "ID": "e760067c9de61076946cad16101e5e82",
                "ElementType": 17,
                "Parent": "humidity",
                "Name": "$unit",
                "Value": "%rH"
              }
            }
          },
          "temperature": {
            "ID": "20cfb60f0333fc4c66343a68385b33cf",
            "ElementType": 16,
            "Parent": "Ambient",
            "Name": "temperature",
            "Value": "26.18",
            "Attrs": {
              "$datatype": {
                "ID": "c4c28f794a906b4d834370d18c22d8c5",
                "ElementType": 17,
                "Parent": "temperature",
                "Name": "$datatype",
                "Value": "float"
              },
              "$name": {
                "ID": "c1656f96196e4dabe8be3ca905c6c2c1",
                "ElementType": 17,
                "Parent": "temperature",
                "Name": "$name",
                "Value": "Temperature"
              },
              "$unit": {
                "ID": "ca4f9f26919cf7e7f338e37f4b68f940",
                "ElementType": 17,
                "Parent": "temperature",
                "Name": "$unit",
                "Value": "°F"
              }
            }
          }
        }
      },
      "Presence": {
        "ID": "2f65a44bc42e84868999ba0b191e46cf",
        "ElementType": 14,
        "Parent": "GarageMonitor",
        "Name": "Presence",
        "Attrs": {
          "$name": {
            "ID": "10437163a083fc69f69e29d94c463f9b",
            "ElementType": 15,
            "Parent": "Presence",
            "Name": "$name",
            "Value": "Motion Sensor"
          },
          "$properties": {
            "ID": "0bd3e099d4bcd7c5b2517e37de79735f",
            "ElementType": 15,
            "Parent": "Presence",
            "Name": "$properties",
            "Value": "motion"
          },
          "$type": {
            "ID": "7a142a71a0783dd1e5d2cd26d21e1478",
            "ElementType": 15,
            "Parent": "Presence",
            "Name": "$type",
            "Value": "sensor"
          }
        },
        "Props": {
          "motion": {
            "ID": "4545ba00e494362df252ad5e7ba04c9a",
            "ElementType": 16,
            "Parent": "Presence",
            "Name": "motion",
            "Value": "OPEN",
            "Attrs": {
              "$datatype": {
                "ID": "01905e1a2d74d6a06bfc2ece5e6887bc",
                "ElementType": 17,
                "Parent": "motion",
                "Name": "$datatype",
                "Value": "enum"
              },
              "$format": {
                "ID": "1c8de77efd3aacaaeefc8c9a2d3432e8",
                "ElementType": 17,
                "Parent": "motion",
                "Name": "$format",
                "Value": "OPEN,CLOSED"
              },
              "$name": {
                "ID": "b8933a0139861da662973531466bb569",
                "ElementType": 17,
                "Parent": "motion",
                "Name": "$name",
                "Value": "Motion"
              }
            }
          }
        }
      },
      "hardware": {
        "ID": "3bbd96f4d01f9cc91b26c7e8c4dbeccd",
        "ElementType": 14,
        "Parent": "GarageMonitor",
        "Name": "hardware",
        "Attrs": {
          "$name": {
            "ID": "0c827538089c683b52e1a095a1583f56",
            "ElementType": 15,
            "Parent": "hardware",
            "Name": "$name",
            "Value": "Device Info"
          },
          "$properties": {
            "ID": "486a8e024f6a900549888c296bc8300e",
            "ElementType": 15,
            "Parent": "hardware",
            "Name": "$properties",
            "Value": "signal,mac,resetReason,voltage"
          },
          "$type": {
            "ID": "e4d96164a6cc491bf95aa7a9aab3f6d8",
            "ElementType": 15,
            "Parent": "hardware",
            "Name": "$type",
            "Value": "sensor"
          }
        },
        "Props": {
          "mac": {
            "ID": "063d81598ffa82e2be5d0cf79e56e57f",
            "ElementType": 16,
            "Parent": "hardware",
            "Name": "mac",
            "Value": "B4:E6:2D:1B:5C:4D",
            "Attrs": {
              "$datatype": {
                "ID": "f167195c8f3693a30e3a9c4c9d73ad06",
                "ElementType": 17,
                "Parent": "mac",
                "Name": "$datatype",
                "Value": "sring"
              },
              "$name": {
                "ID": "6a8f186d267239295a0b5c40e287be97",
                "ElementType": 17,
                "Parent": "mac",
                "Name": "$name",
                "Value": "mac"
              }
            }
          },
          "resetReason": {
            "ID": "a2cf73ed3f489461f525fef416c3c5ec",
            "ElementType": 16,
            "Parent": "hardware",
            "Name": "resetReason",
            "Value": "External System",
            "Attrs": {
              "$datatype": {
                "ID": "296bee8074585df5af42f8aaccec68b5",
                "ElementType": 17,
                "Parent": "resetReason",
                "Name": "$datatype",
                "Value": "string"
              },
              "$name": {
                "ID": "3369da1e1c5706db5583abbc40a096b0",
                "ElementType": 17,
                "Parent": "resetReason",
                "Name": "$name",
                "Value": "Last Reset Reason"
              }
            }
          },
          "signal": {
            "ID": "379fb17a6f73870bb6ea6933c8f90b12",
            "ElementType": 16,
            "Parent": "hardware",
            "Name": "signal",
            "Value": "-73",
            "Attrs": {
              "$datatype": {
                "ID": "9ff83c074a3c5789a6a0cea8169f97b7",
                "ElementType": 17,
                "Parent": "signal",
                "Name": "$datatype",
                "Value": "integer"
              },
              "$name": {
                "ID": "0f1a0b5052048a2b49370af8d42b5238",
                "ElementType": 17,
                "Parent": "signal",
                "Name": "$name",
                "Value": "RSSI"
              },
              "$unit": {
                "ID": "4fd073ab6614bd005b9a25873b05a1de",
                "ElementType": 17,
                "Parent": "signal",
                "Name": "$unit",
                "Value": "dBm"
              }
            }
          },
          "voltage": {
            "ID": "3fca43b5ad8f03c6b0419e2fdfe0c92a",
            "ElementType": 16,
            "Parent": "hardware",
            "Name": "voltage",
            "Value": "3.05",
            "Attrs": {
              "$datatype": {
                "ID": "3ffd557321b8e543be764ff95ad684bd",
                "ElementType": 17,
                "Parent": "voltage",
                "Name": "$datatype",
                "Value": "float"
              },
              "$name": {
                "ID": "b28b79070e57c3f7e8c382bcd8b6867f",
                "ElementType": 17,
                "Parent": "voltage",
                "Name": "$name",
                "Value": "3.3V Supply"
              },
              "$unit": {
                "ID": "a7a9db6a7cca212cdb4d17cd54a96e8f",
                "ElementType": 17,
                "Parent": "voltage",
                "Name": "$unit",
                "Value": "V"
              }
            }
          }
        }
      }
    }
  }
}

```
### Map Data for a Device
```
map[
 "":[[GarageMonitor] [GarageMonitor]]
 GarageMonitor:[[sknSensors/GarageMonitor/$name sknSensors/GarageMonitor/$homie sknSensors/GarageMonitor/$mac sknSensors/GarageMonitor/$nodes sknSensors/GarageMonitor/$state sknSensors/GarageMonitor/$localip sknSensors/GarageMonitor/hardware sknSensors/GarageMonitor/Ambient sknSensors/GarageMonitor/Presence] []]
 sknSensors/GarageMonitor/$fw/checksum:[[] [checksum 615fed382ab44bd43fe83508aecac682]]
 sknSensors/GarageMonitor/$fw/name:[[] [name Monitor-SHT31-RCWL-Metrics]]
 sknSensors/GarageMonitor/$fw/version:[[] [version 2.0.0]]
 sknSensors/GarageMonitor/$homie:[[] [$homie 3.0.1]]
 sknSensors/GarageMonitor/$implementation/config:[[] [config {"name":"Garage Monitor","device_id":"GarageMonitor","device_stats_interval":900,"wifi":{"ssid":"SFNSS1-24G"},"mqtt":{"host":"openhab.skoona.net","port":1883,"base_topic":"sknSensors/","auth":true},"ota":{"enabled":true},"settings":{"sensorInterval":900,"motionHoldInterval":60}}]]
 sknSensors/GarageMonitor/$implementation/ota/enabled:[[] [enabled true]]
 sknSensors/GarageMonitor/$implementation/version:[[] [version 3.0.0]]
 sknSensors/GarageMonitor/$localip:[[] [$localip 10.100.1.177]]
 sknSensors/GarageMonitor/$mac:[[] [$mac B4:E6:2D:1B:5C:4D]]
 sknSensors/GarageMonitor/$name:[[] [$name Garage Monitor]]
 sknSensors/GarageMonitor/$nodes:[[] [$nodes Ambient,Presence,hardware]]
 sknSensors/GarageMonitor/$state:[[] [$state ready]]
 sknSensors/GarageMonitor/$stats/interval:[[] [interval 905]]
 sknSensors/GarageMonitor/$stats/signal:[[] [signal 46]]
 sknSensors/GarageMonitor/$stats/uptime:[[] [uptime 12727347]]
 sknSensors/GarageMonitor/Ambient:[[sknSensors/GarageMonitor/Ambient/$type sknSensors/GarageMonitor/Ambient/$name sknSensors/GarageMonitor/Ambient/humidity sknSensors/GarageMonitor/Ambient/temperature] [Ambient]]
 sknSensors/GarageMonitor/Ambient/$name:[[] [$name Temperature and Humidity Sensor]]
 sknSensors/GarageMonitor/Ambient/$type:[[] [$type sensor]]
 sknSensors/GarageMonitor/Ambient/humidity:[[] [humidity 66.43]]
 sknSensors/GarageMonitor/Ambient/temperature:[[] [temperature 81.98]]
 sknSensors/GarageMonitor/Presence:[[sknSensors/GarageMonitor/Presence/motion] [Presence]]
 sknSensors/GarageMonitor/Presence/motion:[[] [motion CLOSED]]
 sknSensors/GarageMonitor/hardware:[[sknSensors/GarageMonitor/hardware/resetReason sknSensors/GarageMonitor/hardware/signal sknSensors/GarageMonitor/hardware/voltage sknSensors/GarageMonitor/hardware/mac] [hardware]]
 sknSensors/GarageMonitor/hardware/mac:[[] [mac B4:E6:2D:1B:5C:4D]]
 sknSensors/GarageMonitor/hardware/resetReason:[[] [resetReason External System]]
 sknSensors/GarageMonitor/hardware/signal:[[] [signal -76]]
 sknSensors/GarageMonitor/hardware/voltage:[[] [voltage 3.03]]
 ]
```

```go
// buildNetworkDevice  Returns string array a device's properties (topic:value)
func buildNetworkDevice(db *bolt.DB, networkName, deviceName string) (dc.Device, error) {
	var found bool
	var device dc.Device
	var deviceAttribute dc.DeviceAttribute
	var deviceAttributeProperty dc.DeviceAttributeProperty
	var deviceAttributePropertyProperty dc.DeviceAttributePropertyProperty
	var node dc.DeviceNode
	var nodeProperty dc.DeviceNodeProperty
	var nodeAttribute dc.DeviceNodeAttribute
	var nodePropertyAttribute dc.DeviceNodePropertyAttribute

	// todo: Certify device retrieval

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(networkName)) // (1) Network Level
		if b == nil {
			return fmt.Errorf("network not found!: %s", networkName)
		}
		//fmt.Printf("(2) X/D...... %s/%s \n", networkName, deviceName)

		b = b.Bucket([]byte(deviceName)) // (2) Device Level
		if b == nil {
			return fmt.Errorf("network not found!: %s", deviceName)
		}
		device = dc.NewDevice(networkName, deviceName)

		/*
		 * examine device bucket for nodes or attributes */
		err := b.ForEach(func(k, v []byte) error {
			//fmt.Printf("(2) X/D...... Bucket[2]=%s, enum Key=%s Value=[%s] \n", deviceName, string(k), string(v))

			c := b.Bucket(k) // (3) Node/Attr Level
			if c == nil {
				return nil
			}
			/*
			 * bucket is either a Node attrs/props or a DeviceAttribute properties
			 * if 2=attr follow device attr path
			 * else follow nodepath
			*/
			if strings.HasPrefix(string(k), "$") {  // device attr path
				// 3 attribute
				// 4 property
				// 5 propertyProperty

				err := c.ForEach(func(kk, vv []byte) error {
					if string(vv) != "" {
						deviceAttribute, found = device.Attrs[string(kk)]
						if !found {
							//fmt.Printf("(3) X/D/A.... %s/%s -->[%s] \n", deviceName, kk, vv)
							deviceAttribute = dc.NewDeviceAttribute(deviceName, string(kk), string(vv))
							device.Attrs[string(kk)] = deviceAttribute
						}
					}
					d := c.Bucket(kk) // (4) Property
					if d == nil {
						return nil
					}
					err := d.ForEach(func(kkk, vvv []byte) error {
						if string(vvv) != "" {
							deviceAttribute, found = device.Attrs[string(k)]
							if !found {
								deviceAttribute = dc.NewDeviceAttribute(deviceName, string(k), string(vv))
								device.Attrs[string(k)] = deviceAttribute
							}
							deviceAttributeProperty, found = deviceAttribute.Props[string(kkk)]
							if !found {
								//fmt.Printf("(4) X/D/A/P.. %s/%s/%s -->[%s]\n", deviceName, k, kkk, vvv)
								deviceAttributeProperty = dc.NewDeviceAttributeProperty(string(k), string(kkk), string(vvv))
								deviceAttribute.Props[string(kkk)] = deviceAttributeProperty
							}
						}
						e := d.Bucket(kkk) // (5) PProperty
						if e == nil {
							return nil
						}
						err := e.ForEach(func(kkkk, vvvv []byte) error {
							deviceAttribute, found = device.Attrs[string(k)]
							if !found {
								deviceAttribute = dc.NewDeviceAttribute(deviceName, string(k), string(vv))
								device.Attrs[string(k)] = deviceAttribute
							}
							deviceAttributeProperty, found = deviceAttribute.Props[string(kk)]
							if !found {
								deviceAttributeProperty = dc.NewDeviceAttributeProperty(string(k), string(kk), string(vvv))
								deviceAttribute.Props[string(kk)] = deviceAttributeProperty
							}
							deviceAttributePropertyProperty, found = deviceAttributeProperty.Props[string(kkkk)]
							if !found {
								//fmt.Printf("(5) X/D/A/P/P %s/%s/%s/%s -->[%s] \n", deviceName, k, kk, kkkk, vvvv)
								deviceAttributePropertyProperty = dc.NewDeviceAttributePropertyProperty(string(kk), string(kkkk), string(vvvv))
								deviceAttributeProperty.Props[string(kkkk)] = deviceAttributePropertyProperty
							}
							return nil
						})
						return err
					})
					return err
				})
				return err

			} else { // node path
				// 3 node
				// 4 property, attribute
				// 5 propertyAttribute

				err := c.ForEach(func(kk, vv []byte) error {
					if string(vv) != "" {
						node, found = device.Nodes[string(k)]
						if !found {
							//fmt.Printf("(3) X/D/N.... %s/%s -->[%s] \n", deviceName, kk, vv)
							node = dc.NewDeviceNode(deviceName, string(kk))
							device.Nodes[string(kk)] = node
						}
					}

					d := c.Bucket(kk) // (4) Attribute | Property
					if d == nil {
						return nil
					}
					err := d.ForEach(func(kkk, vvv []byte) error {
						if string(vvv) != "" {

							if strings.HasPrefix(string(kkk), "$") {
								node, found = device.Nodes[string(k)]
								if !found {
									node = dc.NewDeviceNode(deviceName, string(k))
									device.Nodes[string(k)] = node
								}
								nodeAttribute, found = node.Attrs[string(kkk)]
								if !found {
									//fmt.Printf("(4) X/D/N/A.. %s/%s/%s -->[%s] \n", deviceName, k, kkk, vvv)
									nodeAttribute = dc.NewDeviceNodeAttribute(string(k), string(kkk), string(vvv))
									node.Attrs[string(kkk)] = nodeAttribute
								}

							} else {
								node, found = device.Nodes[string(k)]
								if !found {
									node = dc.NewDeviceNode(deviceName, string(k))
									device.Nodes[string(k)] = node
								}
								nodeProperty, found = node.Props[string(kkk)]
								if !found {
									//fmt.Printf("(4) X/D/N/P.. %s/%s/%s -->[%s] \n", deviceName, k, kkk, vvv)
									nodeProperty = dc.NewDeviceNodeProperty(string(k), string(kkk), string(vvv))
									node.Props[string(kkk)] = nodeProperty
								}

							}
						}

						e := d.Bucket(kkk) // (5) PPropertyAttribute
						if e == nil {
							return nil
						}
						err := e.ForEach(func(kkkk, vvvv []byte) error {
							node, found = device.Nodes[string(k)]
							if !found {
								node = dc.NewDeviceNode(deviceName, string(k))
								device.Nodes[string(k)] = node
							}
							nodeProperty, found = node.Props[string(kk)]
							if !found {
								nodeProperty = dc.NewDeviceNodeProperty(string(k), string(kk), string(vvv))
								node.Props[string(kk)] = nodeProperty
							}
							nodePropertyAttribute, found = nodeProperty.Attrs[string(kkkk)]
							if !found {
								//fmt.Printf("(5) X/D/N/P/A %s/%s/%s/%s -->[%s] \n", deviceName, k, kk, kkkk, vvvv)
								nodePropertyAttribute = dc.NewDeviceNodePropertyAttribute(string(kk), string(kkkk), string(vvvv))
								nodeProperty.Attrs[string(kkkk)] = nodePropertyAttribute
							}
							return nil
						})
						return err
					})
					return err
				})
				return err
			} // end if path
		})
		return err
	})

	return device, err
}

```
