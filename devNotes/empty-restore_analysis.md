
### From Empty
```json
{
  "ID": "e15f81d9-19ff-48b5-95ed-d02d5cc7d00d",
  "ElementType": 24,
  "SiteName": "Skoona Consulting",
  "Title": "Homie Monitor (GOLANG)",
  "Names": [
    "sknSensors"
  ],
  "Broadcasts": [
    {
      "ID": "c254b24c700dbc2ecba237869f9b493b",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "$broadcast",
      "Level": "alert",
      "Value": "openHAB3 Offline",
      "Received": "2021-07-09T09:06:15.97838-04:00"
    },
    {
      "ID": "c254b24c700dbc2ecba237869f9b493b",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "$broadcast",
      "Level": "alert",
      "Value": "Door Open",
      "Received": "2021-07-09T09:06:16.01636-04:00"
    },
    {
      "ID": "80ea5ef8b530e884349d967e9c09c28a",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "$broadcast",
      "Level": "notice",
      "Value": "openHAB3 Online",
      "Received": "2021-07-09T09:06:16.054517-04:00"
    },
    {
      "ID": "c254b24c700dbc2ecba237869f9b493b",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "$broadcast",
      "Level": "alert",
      "Value": "Garage Door Open",
      "Received": "2021-07-09T09:06:16.09152-04:00"
    },
    {
      "ID": "bd7d6615f85ac2322688ab68d96b4c01",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "$broadcast",
      "Level": "alarm",
      "Value": "Media Room Window Broken",
      "Received": "2021-07-09T09:06:16.151447-04:00"
    }
  ],
  "Firmwares": [
    {
      "ID": "5a3a2c57b5084ee4243017beaa0df466",
      "ElementType": 21,
      "Name": "Environment-DS18B20",
      "FileName": "./dataDB/firmwares/Environment-DS18B20.bin",
      "Version": "1.0.2",
      "Path": "./dataDB/firmwares/Environment-DS18B20.bin",
      "Size": 561456,
      "MD5Digest": "5a3a2c57b5084ee4243017beaa0df466",
      "Brand": "SknSensors",
      "Created": "2021-05-13T23:25:08.350185372-04:00"
    },
    {
      "ID": "3ef8fbb48c5b23788a22a998c14a1a6d",
      "ElementType": 21,
      "Name": "Monitor-DHT-RCWL-Metrics",
      "FileName": "./dataDB/firmwares/Monitor-DHT-RCWL-Metrics-200.bin",
      "Version": "2.0.0",
      "Path": "./dataDB/firmwares/Monitor-DHT-RCWL-Metrics-200.bin",
      "Size": 399360,
      "MD5Digest": "3ef8fbb48c5b23788a22a998c14a1a6d",
      "Brand": "SknSensors",
      "Created": "2021-05-13T23:25:08.352490958-04:00"
    },
    {
      "ID": "615fed382ab44bd43fe83508aecac682",
      "ElementType": 21,
      "Name": "Monitor-SHT31-RCWL-Metrics",
      "FileName": "./dataDB/firmwares/Monitor-SHT31-RCWL-Metrics-200.bin",
      "Version": "2.0.0",
      "Path": "./dataDB/firmwares/Monitor-SHT31-RCWL-Metrics-200.bin",
      "Size": 402640,
      "MD5Digest": "615fed382ab44bd43fe83508aecac682",
      "Brand": "SknSensors",
      "Created": "2021-05-13T23:25:08.354699188-04:00"
    }
  ],
  "Schedules": {},
  "DeviceNetworks": {
    "sknSensors": {
      "ID": "6897631feedfa370c777bc6b6f6c9b3f",
      "Title": "Restored",
      "ElementType": 23,
      "Name": "sknSensors",
      "Devices": {
        "D1R1MiniA": {
          "ID": "908a089bd74911ba7558c0092b1c55ed",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "D1R1MiniA",
          "Attrs": {
            "$fw": {
              "ID": "0547c83c-34f2-446c-b302-cc12de1978f2",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "68a31560-0513-4e1c-98ff-7caf06a869d6",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "6867842dfce4674d4c724085467362c9",
                  "Props": {}
                },
                "name": {
                  "ID": "9db4ae9a-a7b5-44cb-95a6-d22848fe9361",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "SknLiquids-128x32-oled",
                  "Props": {}
                },
                "version": {
                  "ID": "2df20268-16b4-478b-864f-3e7443abbd1c",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "1.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "1517bc10-6caf-4a06-bbfb-45f03fb18946",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "de00003b-34e1-461a-920f-233707acf5b3",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$implementation",
              "Value": "esp8266",
              "Props": {
                "config": {
                  "ID": "100fa1c6-28c2-4b58-b60e-4621bc42e12b",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"Furnace\",\"device_id\":\"D1R1MiniA\",\"device_stats_interval\":900,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "d5adb121-83b9-4233-bd54-83d4fb0941ed",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "ota": {
                      "ID": "eaba1f0a-5657-43e0-b881-2440fd4a07b8",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "804efbdd-8708-49e3-8b2e-6b7d8e7103a2",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "433a620b-2b9c-4dad-a516-fedf7b4d6698",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$localip",
              "Value": "10.100.1.181",
              "Props": {}
            },
            "$mac": {
              "ID": "7ed10144-8fd4-4eb5-8554-a6757e56f1b6",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$mac",
              "Value": "84:F3:EB:0C:38:6F",
              "Props": {}
            },
            "$name": {
              "ID": "6e0835cd-0e46-4a3e-9940-9b42fcc316da",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$name",
              "Value": "Furnace",
              "Props": {}
            },
            "$nodes": {
              "ID": "db53baa9-1544-4669-a1f5-cca2a91f1f6e",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$nodes",
              "Value": "Liquids",
              "Props": {}
            },
            "$state": {
              "ID": "daa4676c-b6c1-4469-ab2d-8eebb25c5c51",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "ff69f1d8-7c3c-40e2-8bdf-db400b4960e3",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$stats",
              "Value": "uptime",
              "Props": {
                "interval": {
                  "ID": "8d9c2242-a15f-48d4-95c7-d28b95a5c62e",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "905",
                  "Props": {}
                },
                "signal": {
                  "ID": "f4d16a1f-8c4e-4755-9c8a-6a147b6b1ea9",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "88",
                  "Props": {}
                },
                "uptime": {
                  "ID": "0009408e-5d47-4716-b10d-463c379d60a7",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "114307",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "Liquids": {
              "ID": "70b11eaa-32d2-41fe-b363-1c333e95b87f",
              "ElementType": 14,
              "Parent": "D1R1MiniA",
              "Name": "Liquids",
              "Attrs": {
                "$name": {
                  "ID": "c5b661d6-5f23-4656-9d76-21add26ac45d",
                  "ElementType": 15,
                  "Parent": "Liquids",
                  "Name": "$name",
                  "Value": "Water Sensor"
                },
                "$properties": {
                  "ID": "b2cb8f0d-d8d8-4127-bc00-ddaf3f028a9e",
                  "ElementType": 15,
                  "Parent": "Liquids",
                  "Name": "$properties",
                  "Value": "level,volts"
                },
                "$type": {
                  "ID": "ab57f547-b484-402c-8dc9-6c9e2134a33e",
                  "ElementType": 15,
                  "Parent": "Liquids",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "level": {
                  "ID": "4948ad16-2ae6-4c02-886d-fcb80c24a716",
                  "ElementType": 16,
                  "Parent": "Liquids",
                  "Name": "level",
                  "Value": "7",
                  "Attrs": {
                    "$datatype": {
                      "ID": "d653f311-afec-48bd-b3fb-7ec40c09bf13",
                      "ElementType": 17,
                      "Parent": "level",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$format": {
                      "ID": "484355fd-b9d5-4cfa-9722-1c0ca9f62ebc",
                      "ElementType": 17,
                      "Parent": "level",
                      "Name": "$format",
                      "Value": "0:1000"
                    },
                    "$name": {
                      "ID": "f9fb606a-b2b0-4dec-bce7-c761bcfa1648",
                      "ElementType": 17,
                      "Parent": "level",
                      "Name": "$name",
                      "Value": "Level"
                    },
                    "$unit": {
                      "ID": "aacd2141-0b29-4bb4-8f4e-67655c2f2561",
                      "ElementType": 17,
                      "Parent": "level",
                      "Name": "$unit",
                      "Value": "#"
                    }
                  }
                },
                "volts": {
                  "ID": "5a2756d7-9681-4b3d-85f1-fb866d8745c4",
                  "ElementType": 16,
                  "Parent": "Liquids",
                  "Name": "volts",
                  "Value": "0.0",
                  "Attrs": {
                    "$datatype": {
                      "ID": "2e0ed437-38b6-4839-b9cb-0547e27517c6",
                      "ElementType": 17,
                      "Parent": "volts",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$format": {
                      "ID": "e927cd5b-c469-4ce1-8a47-51a14076231f",
                      "ElementType": 17,
                      "Parent": "volts",
                      "Name": "$format",
                      "Value": "0:4"
                    },
                    "$name": {
                      "ID": "a1e04083-5014-4141-9860-21bbd87a8ee0",
                      "ElementType": 17,
                      "Parent": "volts",
                      "Name": "$name",
                      "Value": "Volts"
                    },
                    "$unit": {
                      "ID": "d0774c34-afef-4973-b5db-a125884fae5d",
                      "ElementType": 17,
                      "Parent": "volts",
                      "Name": "$unit",
                      "Value": "V"
                    }
                  }
                }
              }
            }
          }
        },
        "FamilyRoom": {
          "ID": "2f63939498209cf4584b6e82954d2407",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "FamilyRoom",
          "Attrs": {
            "$fw": {
              "ID": "8210ca75-7a5f-4ef8-8660-4a6339991665",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "2922a996-7bad-45e7-adfc-87cdb7e82dd9",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "3ef8fbb48c5b23788a22a998c14a1a6d",
                  "Props": {}
                },
                "name": {
                  "ID": "9d248e56-c1fb-4d72-8cd7-08af21dca621",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "Monitor-DHT-RCWL-Metrics",
                  "Props": {}
                },
                "version": {
                  "ID": "db575517-2d6f-4b6e-a80f-d99de0ee2dc3",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "2.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "31867492-7680-4119-8ed6-4872f59ea63d",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "b57bd050-bab5-4ab5-8abe-84ad710bc619",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$implementation",
              "Value": "esp8266",
              "Props": {
                "config": {
                  "ID": "57dcbf43-01cc-4da2-89c6-613f50d4d902",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"FamilyRoom\",\"device_id\":\"FamilyRoom\",\"device_stats_interval\":180,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"sensorInterval\":900,\"motionHoldInterval\":60}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "ab2f7fad-bc85-454b-b88d-271aa8e5bd0c",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "ota": {
                      "ID": "46f423c7-7c72-4ae5-bc02-74ac7b450a09",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "aedc1f07-db78-45b8-a6d3-5a2583f78f9a",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "a544e58c-df58-4ef3-9c93-818a646643dc",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$localip",
              "Value": "10.100.1.166",
              "Props": {}
            },
            "$mac": {
              "ID": "efae884a-f65c-4908-b61c-e6f2743b1cb2",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$mac",
              "Value": "BC:DD:C2:E5:C4:24",
              "Props": {}
            },
            "$name": {
              "ID": "249e96d5-b6aa-4194-8cb1-b5f7915935ff",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$name",
              "Value": "FamilyRoom",
              "Props": {}
            },
            "$nodes": {
              "ID": "f25176b2-9e24-4c5b-98ee-bf7e05926c62",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$nodes",
              "Value": "Ambient,Presence,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "53544848-7437-4358-83df-7a5acb64628b",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "833bf6f6-0056-4898-aa1d-26208d0e9862",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$stats",
              "Value": "uptime",
              "Props": {
                "interval": {
                  "ID": "07c4abcd-3463-48dd-8938-d007fb97260b",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "185",
                  "Props": {}
                },
                "signal": {
                  "ID": "5346bbe2-82f4-41f0-8b94-5ae16275c010",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "78",
                  "Props": {}
                },
                "uptime": {
                  "ID": "6945f5a7-a5e8-484c-ae6f-0de50a348685",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "423459",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "Ambient": {
              "ID": "2036d9ed-b55c-4d0f-83e9-ae994f32014b",
              "ElementType": 14,
              "Parent": "FamilyRoom",
              "Name": "Ambient",
              "Attrs": {
                "$name": {
                  "ID": "79dd294c-6bd6-4774-8124-9909db612e06",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$name",
                  "Value": "Temperature and Humidity Sensor"
                },
                "$properties": {
                  "ID": "ea364676-fa4e-45c1-a0c4-ad959035aa0e",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$properties",
                  "Value": "humidity,temperature"
                },
                "$type": {
                  "ID": "f059b653-8a29-405c-ab8d-5693b843fcfa",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "humidity": {
                  "ID": "7ac0f902-0748-4a8c-91d9-9b28701aaf68",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "humidity",
                  "Value": "20.50",
                  "Attrs": {
                    "$datatype": {
                      "ID": "7b069151-07bc-4728-89bf-d88a47813a3b",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "1a47ee9a-532c-43e6-ac27-51a6d66f06f3",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$name",
                      "Value": "Humidity"
                    },
                    "$unit": {
                      "ID": "c8a7d908-a042-48ef-ae50-a63695dca078",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": "%rH"
                    }
                  }
                },
                "temperature": {
                  "ID": "2cd94a92-62dc-4be4-b09a-9d495f574acb",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "temperature",
                  "Value": "67.64",
                  "Attrs": {
                    "$datatype": {
                      "ID": "c05da997-bb53-4da4-8db5-015cc692c93b",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "3a1987f9-be47-4045-925f-7c97c6ba46b3",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$name",
                      "Value": "Temperature"
                    },
                    "$unit": {
                      "ID": "7f6efb56-6bf3-46f7-b4f7-12a3628dac0c",
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
              "ID": "55d3930e-2229-4caa-acbd-4da5742671ba",
              "ElementType": 14,
              "Parent": "FamilyRoom",
              "Name": "Presence",
              "Attrs": {
                "$name": {
                  "ID": "a94f35d7-a259-4b02-ac84-25fd475a2aef",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$name",
                  "Value": "Motion Sensor"
                },
                "$properties": {
                  "ID": "cb74c953-fb33-46e7-8841-baa9f9267923",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$properties",
                  "Value": "motion"
                },
                "$type": {
                  "ID": "022dd8d9-c726-4d14-a85f-10bd16053365",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "motion": {
                  "ID": "65fc2d4b-685b-48b8-8708-d023b537afb6",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "motion",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "d88f5fa5-54e1-4f4e-9a5a-995cf5238cbd",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "9f27ee8e-f66f-41aa-ae9d-d5b4917d58a0",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "252c314c-65b4-4b3a-9dd7-ff5aadae220c",
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
              "ID": "3b926290-1d46-4bbb-8883-36f56a993ecd",
              "ElementType": 14,
              "Parent": "FamilyRoom",
              "Name": "hardware",
              "Attrs": {
                "$name": {
                  "ID": "ec70a4f4-bd2c-43d7-bb0a-36d947a7ec2c",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$name",
                  "Value": "Device Info"
                },
                "$properties": {
                  "ID": "9c7b1a9f-8b9a-41bf-b7b1-1ff683e5bd84",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$properties",
                  "Value": "signal,mac,resetReason,voltage"
                },
                "$type": {
                  "ID": "99c56fb7-8149-44ea-b981-bc096122c39d",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "mac": {
                  "ID": "a92db52c-2ae8-45e3-a346-83f01202dc57",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "mac",
                  "Value": "BC:DD:C2:E5:C4:24",
                  "Attrs": {
                    "$datatype": {
                      "ID": "2c04fb85-4dd6-449b-af64-2f07a2ba53be",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": "sring"
                    },
                    "$name": {
                      "ID": "d3eb0a23-099d-425c-833d-5d70f4077acd",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$name",
                      "Value": "mac"
                    }
                  }
                },
                "resetReason": {
                  "ID": "cea5e1e7-af0f-4b39-bf5c-e4122651d304",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "resetReason",
                  "Value": "External System",
                  "Attrs": {
                    "$datatype": {
                      "ID": "db585aca-0adc-49cd-9d70-86aa5a966360",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$name": {
                      "ID": "424ce6f1-96f5-4d41-9a4a-dff6d41b6f60",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$name",
                      "Value": "Last Reset Reason"
                    }
                  }
                },
                "signal": {
                  "ID": "c394242a-9c2e-4225-adbc-50ca9104b2e8",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "signal",
                  "Value": "-61",
                  "Attrs": {
                    "$datatype": {
                      "ID": "8f855921-9e88-4743-b7fc-8b7fa383db8d",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$name": {
                      "ID": "5aebaca1-cb0f-40c8-865f-20c21f2bc7dc",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$name",
                      "Value": "RSSI"
                    },
                    "$unit": {
                      "ID": "29d5b52f-05fc-4445-8db3-c779048b4282",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": "dBm"
                    }
                  }
                },
                "voltage": {
                  "ID": "cef06b28-d1c7-48a9-b89c-298b0a2d31d6",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "voltage",
                  "Value": "3.03",
                  "Attrs": {
                    "$datatype": {
                      "ID": "9b692055-cf24-4189-a747-2b95b4c990fd",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "f89ac3b6-e2f9-4a2e-a8af-be8011c68f96",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$name",
                      "Value": "3.3V Supply"
                    },
                    "$unit": {
                      "ID": "6c0fd381-4892-46e9-87a6-7153656a8346",
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
        },
        "GarageMonitor": {
          "ID": "36ec4f3d3eedbe07b28e5fa6ab6ddb9b",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "GarageMonitor",
          "Attrs": {
            "$fw": {
              "ID": "95b48939-8af9-41d9-a6bc-b7eee2987a6d",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "c7ce7e09-2fd7-42ac-93e2-7602a2483ca5",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "615fed382ab44bd43fe83508aecac682",
                  "Props": {}
                },
                "name": {
                  "ID": "af00a32e-dab2-4e02-84c9-c41be05434f4",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "Monitor-SHT31-RCWL-Metrics",
                  "Props": {}
                },
                "version": {
                  "ID": "6c59463c-2f0d-48c7-89ba-078fb4e58ba8",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "2.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "6fda83ca-0247-43ff-ade7-cbcc4843cbfc",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "cf607c2e-087b-43fb-a374-7b7e63120222",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$implementation",
              "Value": "esp8266",
              "Props": {
                "config": {
                  "ID": "963bf4ee-4764-4803-937d-45525674c02d",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"Garage Monitor\",\"device_id\":\"GarageMonitor\",\"device_stats_interval\":900,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"sensorInterval\":900,\"motionHoldInterval\":60}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "108fe59c-2824-409a-bf5a-76a41f5ced8e",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "ota": {
                      "ID": "f0ca2969-9435-408f-9a83-299269dc53c0",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "cfbb9716-8d81-4a66-bd6f-d29af6e15f5e",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "89a9dc60-5e9d-48dd-aa4c-d63fd6081d18",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$localip",
              "Value": "10.100.1.177",
              "Props": {}
            },
            "$mac": {
              "ID": "eb819e0f-0772-4db6-b55b-1b4beb366eb6",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$mac",
              "Value": "B4:E6:2D:1B:5C:4D",
              "Props": {}
            },
            "$name": {
              "ID": "559df22f-5489-4391-bcde-df63c8805498",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$name",
              "Value": "Garage Monitor",
              "Props": {}
            },
            "$nodes": {
              "ID": "aceba469-6c5b-4a62-81db-7a7fde4fe331",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$nodes",
              "Value": "Ambient,Presence,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "e2798260-36ae-4d0c-bcf6-62f907b79079",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "4bc37e43-8828-484b-945b-e47db0526952",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$stats",
              "Value": "uptime",
              "Props": {
                "interval": {
                  "ID": "7bf3e53d-8137-4739-b4b8-cb116c42c7ae",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "905",
                  "Props": {}
                },
                "signal": {
                  "ID": "1ed8def7-e741-4256-82ab-a66484170356",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "54",
                  "Props": {}
                },
                "uptime": {
                  "ID": "194ffef9-3667-44b6-ad3e-3218fe16001f",
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
              "ID": "ca14f387-f068-43fd-8c7b-6e63732c1a2b",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "Ambient",
              "Attrs": {
                "$name": {
                  "ID": "08be2a14-aac6-4ff6-a6ee-d15e658d1ee7",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$name",
                  "Value": "Temperature and Humidity Sensor"
                },
                "$properties": {
                  "ID": "580ab46c-2a57-434f-98f1-7eda9873178f",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$properties",
                  "Value": "humidity,temperature"
                },
                "$type": {
                  "ID": "3f6096d4-30fd-4d01-aae4-ff0d83ba3b86",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "humidity": {
                  "ID": "e8af0bbe-b86f-4c43-b542-e26bc83e1363",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "humidity",
                  "Value": "57.18",
                  "Attrs": {
                    "$datatype": {
                      "ID": "aab61f47-1db4-4c0e-a730-c11a2f316623",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "8d3553ec-584d-45f0-a1cd-5b27c027e150",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$name",
                      "Value": "Humidity"
                    },
                    "$unit": {
                      "ID": "56a3e016-5e75-4732-b653-5db5498e6d2e",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": "%rH"
                    }
                  }
                },
                "temperature": {
                  "ID": "d870ed0f-18d1-4c25-9169-09dd615e3919",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "temperature",
                  "Value": "26.18",
                  "Attrs": {
                    "$datatype": {
                      "ID": "21ef06d7-c3fe-48e9-ac7c-191342f3da69",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "2667c37d-c77c-4193-b3e0-a2dd8f052044",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$name",
                      "Value": "Temperature"
                    },
                    "$unit": {
                      "ID": "6fa06759-d3b4-482e-bb11-9c159c4b98fe",
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
              "ID": "cd3c3928-1962-45b2-8b5e-f129164477b8",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "Presence",
              "Attrs": {
                "$name": {
                  "ID": "e5821a67-91fc-4088-8abe-3ce48ab5c80d",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$name",
                  "Value": "Motion Sensor"
                },
                "$properties": {
                  "ID": "b3eab42a-10c5-4ca4-9f18-faecac89912c",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$properties",
                  "Value": "motion"
                },
                "$type": {
                  "ID": "1fc60b84-8554-4430-821f-2a7db01bf2cc",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "motion": {
                  "ID": "895ec992-854d-4606-8e91-bbed85b7b44f",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "motion",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "d44134ac-0763-47ba-8647-aec5f70415ab",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "8d94b0bf-28fb-437c-95bc-6ea174745e41",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "ec9d04c5-d093-4351-b01c-d55d7af480ea",
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
              "ID": "0ec32672-24a9-425b-b02d-65cdb89bc281",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "hardware",
              "Attrs": {
                "$name": {
                  "ID": "9fd052ef-0061-470b-9971-82b754ee87f3",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$name",
                  "Value": "Device Info"
                },
                "$properties": {
                  "ID": "63f4a041-1f48-4c15-8714-797f3b6bdfbf",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$properties",
                  "Value": "signal,mac,resetReason,voltage"
                },
                "$type": {
                  "ID": "31a80368-c947-4749-a0a8-c8d98c81b69a",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "mac": {
                  "ID": "db32a290-84e3-4a7f-bbd1-292edbe139af",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "mac",
                  "Value": "B4:E6:2D:1B:5C:4D",
                  "Attrs": {
                    "$datatype": {
                      "ID": "0ad3761a-b473-4654-ab48-ec243c1e34ea",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": "sring"
                    },
                    "$name": {
                      "ID": "7e8438c7-56f3-4f3a-88ba-7edbf355bdf7",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$name",
                      "Value": "mac"
                    }
                  }
                },
                "resetReason": {
                  "ID": "be23e2b0-f637-4de8-aeb7-30ea899aaaef",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "resetReason",
                  "Value": "External System",
                  "Attrs": {
                    "$datatype": {
                      "ID": "eb88050e-4050-480e-bceb-4e999e058da3",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$name": {
                      "ID": "08e15ebd-e61c-4033-a02e-d8662b40bc61",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$name",
                      "Value": "Last Reset Reason"
                    }
                  }
                },
                "signal": {
                  "ID": "fff52379-29e1-45bf-9f7c-40cf5695ab4a",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "signal",
                  "Value": "-73",
                  "Attrs": {
                    "$datatype": {
                      "ID": "731dbe1f-b516-4c21-8267-23c78bf1b231",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$name": {
                      "ID": "ed279ee2-3b08-40dd-88d7-90b732a90109",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$name",
                      "Value": "RSSI"
                    },
                    "$unit": {
                      "ID": "f89ef9e8-a8df-4686-b818-3ed807e7b168",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": "dBm"
                    }
                  }
                },
                "voltage": {
                  "ID": "83e0bc0f-d897-40d5-a487-30af0416c539",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "voltage",
                  "Value": "3.05",
                  "Attrs": {
                    "$datatype": {
                      "ID": "43b0384e-21b5-4381-b8f8-a3e1c3ef9668",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "ffc887d5-be73-43f3-8e25-5d3fdbe7a663",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$name",
                      "Value": "3.3V Supply"
                    },
                    "$unit": {
                      "ID": "c8dc73d3-52d4-46e3-b66a-9c9d25dc649c",
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
        },
        "GarageProvider": {
          "ID": "71c6d9d39d15fe3565d22ab2ea689028",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "GarageProvider",
          "Attrs": {
            "$fw": {
              "ID": "114af6f0-8eeb-421d-a42c-e8e82a7e837d",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "e0ded65d-e71a-49ed-8148-547e08bd8b00",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "39286ba639a9ef1d5395e1c0e45d13fa",
                  "Props": {}
                },
                "name": {
                  "ID": "67cc6c39-91a0-4d23-98d4-8f6e0ad6ef84",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "GarageProvider",
                  "Props": {}
                },
                "version": {
                  "ID": "acc62064-bc39-4fc4-ad73-6d27fb8ffc80",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "1.1.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "10f10182-7386-4a72-9c03-a6048702f7e6",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "8c245aaa-0c3c-41cd-a609-d989f1d20b0e",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$implementation",
              "Value": "esp32",
              "Props": {
                "config": {
                  "ID": "e955313e-9e5b-45fc-87d2-2c38332193b9",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"Garage Provider\",\"device_id\":\"GarageProvider\",\"device_stats_interval\":900,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"sensorsInterval\":900}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "1716dd2b-0156-4e6a-8680-639a3a04c465",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "ota": {
                      "ID": "d19aa37b-310a-41e0-ae22-a91e1e18e38e",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "2b2bddc6-9a1d-4ccf-bd69-3b30f3171148",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "5a303733-7a90-4022-9fee-69ca318db1a2",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$localip",
              "Value": "10.100.1.218",
              "Props": {}
            },
            "$mac": {
              "ID": "14f8e535-86b3-43b6-a230-a9925204d8d4",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$mac",
              "Value": "24:6F:28:97:63:B8",
              "Props": {}
            },
            "$name": {
              "ID": "6a6f1cec-8b34-43cd-826b-e35e27d23412",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$name",
              "Value": "Garage Provider",
              "Props": {}
            },
            "$nodes": {
              "ID": "02783dc0-0032-4404-82de-1eca96b6b91c",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$nodes",
              "Value": "garageDoor,environmentMonitor",
              "Props": {}
            },
            "$state": {
              "ID": "9d8e3518-a57e-442b-a5a1-14d240902305",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "76ff1bee-d73b-4258-a715-d2fc14b45fab",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$stats",
              "Value": "uptime",
              "Props": {
                "interval": {
                  "ID": "b11463e5-da72-4114-870d-d3124058e48b",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "905",
                  "Props": {}
                },
                "signal": {
                  "ID": "9a943f69-99c8-4ca6-9824-9d05170736e1",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "72",
                  "Props": {}
                },
                "uptime": {
                  "ID": "f33f6924-7ea4-4dee-8008-53670ee7f2fa",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "113713",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "environmentMonitor": {
              "ID": "f0f246af-99b6-462b-9cd6-2544e6b3134c",
              "ElementType": 14,
              "Parent": "GarageProvider",
              "Name": "environmentMonitor",
              "Attrs": {
                "$name": {
                  "ID": "4c20e307-ca10-4d70-845c-851c3127adb9",
                  "ElementType": 15,
                  "Parent": "environmentMonitor",
                  "Name": "$name",
                  "Value": "Garage"
                },
                "$properties": {
                  "ID": "d5236b7d-c64f-4d4f-93a3-f52e55b042bf",
                  "ElementType": 15,
                  "Parent": "environmentMonitor",
                  "Name": "$properties",
                  "Value": "motion,temperature,humdity"
                },
                "$type": {
                  "ID": "403e1057-9bbc-45ea-be99-b033dd688b0d",
                  "ElementType": 15,
                  "Parent": "environmentMonitor",
                  "Name": "$type",
                  "Value": "SknSensors"
                }
              },
              "Props": {
                "humdity": {
                  "ID": "a79a4bbf-9429-451f-afda-bb3df4b076cb",
                  "ElementType": 16,
                  "Parent": "environmentMonitor",
                  "Name": "humdity",
                  "Value": "29.0",
                  "Attrs": {
                    "$datatype": {
                      "ID": "21a59d98-ad48-4996-8d00-9a89af078a32",
                      "ElementType": 17,
                      "Parent": "humdity",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$format": {
                      "ID": "67c84ddc-4275-4322-b60e-46022a9dc9af",
                      "ElementType": 17,
                      "Parent": "humdity",
                      "Name": "$format",
                      "Value": "%.1f"
                    },
                    "$name": {
                      "ID": "91ed206c-4ad1-4020-8bc8-60818ee3d378",
                      "ElementType": 17,
                      "Parent": "humdity",
                      "Name": "$name",
                      "Value": "Humidity"
                    },
                    "$unit": {
                      "ID": "a01f8c17-2ccf-4c9a-9fd2-4a7159b113e6",
                      "ElementType": 17,
                      "Parent": "humdity",
                      "Name": "$unit",
                      "Value": "%"
                    }
                  }
                },
                "motion": {
                  "ID": "c11692d3-3d36-493e-aa51-0cc0cb19dd20",
                  "ElementType": 16,
                  "Parent": "environmentMonitor",
                  "Name": "motion",
                  "Value": "ON",
                  "Attrs": {
                    "$datatype": {
                      "ID": "112c3d22-d824-416b-9568-de07a64ab749",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$format": {
                      "ID": "39c634d5-bc18-437b-999c-72ae7ad6fc75",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": "%s"
                    },
                    "$name": {
                      "ID": "512941c3-f886-4b23-abb4-1d79d842ae18",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$name",
                      "Value": "Motion"
                    }
                  }
                },
                "temperature": {
                  "ID": "ed6c265a-89a0-4143-a1e4-c58fc0a8bdfe",
                  "ElementType": 16,
                  "Parent": "environmentMonitor",
                  "Name": "temperature",
                  "Value": "73.8",
                  "Attrs": {
                    "$datatype": {
                      "ID": "9896d19c-e36e-427a-9e4a-9ac53a9f941f",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$format": {
                      "ID": "33f7dd06-4fc9-4743-bb13-a56d0b26e03b",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$format",
                      "Value": "%.1f"
                    },
                    "$name": {
                      "ID": "479f916b-77a8-447c-ae22-57f82589c0b3",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$name",
                      "Value": "Temperature"
                    },
                    "$unit": {
                      "ID": "183e90ad-e6c7-4ea1-b72f-950820d27c6e",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$unit",
                      "Value": "ºF"
                    }
                  }
                }
              }
            },
            "garageDoor": {
              "ID": "2913de12-d7f2-46b9-88f3-cf15aa0bd987",
              "ElementType": 14,
              "Parent": "GarageProvider",
              "Name": "garageDoor",
              "Attrs": {
                "$name": {
                  "ID": "4857b1eb-dcac-4e40-a269-d0a3aac53b40",
                  "ElementType": 15,
                  "Parent": "garageDoor",
                  "Name": "$name",
                  "Value": "Operations"
                },
                "$properties": {
                  "ID": "bd53d017-d16a-4378-ad77-df19e7046957",
                  "ElementType": 15,
                  "Parent": "garageDoor",
                  "Name": "$properties",
                  "Value": "operator,positon,direction"
                },
                "$type": {
                  "ID": "c7012dbb-77e3-4abb-a1c1-6da2708d372d",
                  "ElementType": 15,
                  "Parent": "garageDoor",
                  "Name": "$type",
                  "Value": "SknSensors"
                }
              },
              "Props": {
                "direction": {
                  "ID": "ab53969a-2b48-41c8-9791-59ecd18f4e24",
                  "ElementType": 16,
                  "Parent": "garageDoor",
                  "Name": "direction",
                  "Value": "OPENING",
                  "Attrs": {
                    "$datatype": {
                      "ID": "4fb9fc0d-7fc7-4f96-9830-38a76bfd9c85",
                      "ElementType": 17,
                      "Parent": "direction",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$format": {
                      "ID": "974f7683-5ce2-417f-a00f-5dcb07225ab9",
                      "ElementType": 17,
                      "Parent": "direction",
                      "Name": "$format",
                      "Value": "%s"
                    },
                    "$name": {
                      "ID": "2be1dea3-65c1-4bea-8bb8-77ef03c583ff",
                      "ElementType": 17,
                      "Parent": "direction",
                      "Name": "$name",
                      "Value": "Travel Direction"
                    }
                  }
                },
                "operator": {
                  "ID": "8f418b51-a9b9-4f36-b4bf-9ee41d51ee72",
                  "ElementType": 16,
                  "Parent": "garageDoor",
                  "Name": "operator",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "8b068bb6-37de-4a0a-a4d8-35eec5a0e506",
                      "ElementType": 17,
                      "Parent": "operator",
                      "Name": "$datatype",
                      "Value": "boolean"
                    },
                    "$format": {
                      "ID": "5be73fe8-7a57-4772-b5cb-47bfa58f7f5c",
                      "ElementType": 17,
                      "Parent": "operator",
                      "Name": "$format",
                      "Value": "%s"
                    },
                    "$name": {
                      "ID": "37347d72-e7d1-4fd1-bc9c-73ef4a7372a3",
                      "ElementType": 17,
                      "Parent": "operator",
                      "Name": "$name",
                      "Value": "Operator"
                    },
                    "$settable": {
                      "ID": "4388fad5-d03f-4e6d-b886-10ff01c6346a",
                      "ElementType": 17,
                      "Parent": "operator",
                      "Name": "$settable",
                      "Value": "true"
                    }
                  }
                },
                "positon": {
                  "ID": "c5896ecf-f23e-4cc4-9c2a-6fd951b7ae87",
                  "ElementType": 16,
                  "Parent": "garageDoor",
                  "Name": "positon",
                  "Value": "19",
                  "Attrs": {
                    "$datatype": {
                      "ID": "878f1130-ce24-4e44-a63b-73e17b00588f",
                      "ElementType": 17,
                      "Parent": "positon",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$format": {
                      "ID": "c31ddfc4-cb1d-4630-884d-8cf119b5b97f",
                      "ElementType": 17,
                      "Parent": "positon",
                      "Name": "$format",
                      "Value": "%d"
                    },
                    "$name": {
                      "ID": "97daeb90-5525-4b2f-b038-ac876c12319a",
                      "ElementType": 17,
                      "Parent": "positon",
                      "Name": "$name",
                      "Value": "Position MM"
                    },
                    "$unit": {
                      "ID": "4e979048-58d2-4702-b76c-b2d12a22e585",
                      "ElementType": 17,
                      "Parent": "positon",
                      "Name": "$unit",
                      "Value": "mm"
                    }
                  }
                }
              }
            }
          }
        },
        "GuestRoom": {
          "ID": "bdab9513512a9d316939b5c0ac1ee099",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "GuestRoom",
          "Attrs": {
            "$fw": {
              "ID": "ca492453-eaea-4635-9bfe-fcc1cbf457fc",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "e8a4c164-1f6c-462a-bcc2-4fd58fbef0a9",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "3ef8fbb48c5b23788a22a998c14a1a6d",
                  "Props": {}
                },
                "name": {
                  "ID": "a14ef0a9-a678-437d-bcc2-66386b27689e",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "Monitor-DHT-RCWL-Metrics",
                  "Props": {}
                },
                "version": {
                  "ID": "1ad38a24-daf6-42e3-9e3b-d5a7d106a93f",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "2.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "cbb9b41f-5cc2-4068-98ce-2251ef3ccad2",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "3478403b-7771-498c-a9da-eafe93f9940d",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$implementation",
              "Value": "esp8266",
              "Props": {
                "config": {
                  "ID": "34614c3a-4a63-4488-8c49-fe758c22ddfe",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"Guest Room\",\"device_id\":\"GuestRoom\",\"device_stats_interval\":180,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"sensorInterval\":900,\"motionHoldInterval\":60}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "972b3b34-d3eb-4894-ae0e-852ae52dc8de",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "ota": {
                      "ID": "ddcb8405-c678-4c3d-b8b7-e763abe44046",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "fcf916c0-3472-46bc-a2ef-739b8e5eedc0",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "8cfb0c48-3ffb-4079-9737-9dccff4aed40",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$localip",
              "Value": "10.100.1.178",
              "Props": {}
            },
            "$mac": {
              "ID": "27a3000f-d946-4cf9-a609-6dbdb5520390",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$mac",
              "Value": "BC:DD:C2:81:04:72",
              "Props": {}
            },
            "$name": {
              "ID": "a84eac42-c4c4-435f-92fd-54f689a15983",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$name",
              "Value": "Guest Room",
              "Props": {}
            },
            "$nodes": {
              "ID": "80874648-f118-438c-8dc7-c6c68b70564e",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$nodes",
              "Value": "Ambient,Presence,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "7aab1f6a-d532-4105-a8d4-5d4c917de57b",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "d3346387-8961-4706-aafc-4e26fe23f70d",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$stats",
              "Value": "uptime",
              "Props": {
                "interval": {
                  "ID": "beaf800c-b322-4814-bd07-78325956e307",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "185",
                  "Props": {}
                },
                "signal": {
                  "ID": "fd101fd5-9662-4075-b82f-81aff9e56722",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "72",
                  "Props": {}
                },
                "uptime": {
                  "ID": "24b32bf1-cf62-4a74-a6a0-b3821a98112c",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "361998",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "Ambient": {
              "ID": "21b4b79b-12a3-4d52-bfe0-2b8601684a4c",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "Ambient",
              "Attrs": {
                "$name": {
                  "ID": "4c1c6278-53d0-43e8-aad0-4051108895e4",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$name",
                  "Value": "Temperature and Humidity Sensor"
                },
                "$properties": {
                  "ID": "f108f6c7-f7f4-4cab-8420-9cb0c197eff8",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$properties",
                  "Value": "humidity,temperature"
                },
                "$type": {
                  "ID": "59fd6561-7884-46e2-a075-a1f694556cb6",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "humidity": {
                  "ID": "2cb7b93b-eaf1-40a2-bcb2-2bc8b5ea672f",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "humidity",
                  "Value": "11.10",
                  "Attrs": {
                    "$datatype": {
                      "ID": "a3ae8ed4-4759-4509-8fe1-04666b94f44b",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "c53c2518-0cc3-4176-907f-e0c69c76c88f",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$name",
                      "Value": "Humidity"
                    },
                    "$unit": {
                      "ID": "a1feba03-cdd0-419e-8367-20b87f231b81",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": "%rH"
                    }
                  }
                },
                "temperature": {
                  "ID": "e6a7c64c-f3e5-4acd-9213-0692bdea271f",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "temperature",
                  "Value": "73.76",
                  "Attrs": {
                    "$datatype": {
                      "ID": "41d51ff3-bb4d-40fa-9306-764275d694ac",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "6fe6aeb9-496c-4c35-87cf-d942eb93d922",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$name",
                      "Value": "Temperature"
                    },
                    "$unit": {
                      "ID": "a3b3634b-70bd-4d87-8504-448402d2a0c5",
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
              "ID": "92c5f4b3-1fbe-4a9c-b6f0-f31367e8aeb0",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "Presence",
              "Attrs": {
                "$name": {
                  "ID": "28995a03-e2c8-45d7-b571-29545ec010e3",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$name",
                  "Value": "Motion Sensor"
                },
                "$properties": {
                  "ID": "424f6b1d-a954-4be7-9deb-049d82bc204b",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$properties",
                  "Value": "motion"
                },
                "$type": {
                  "ID": "8865d23d-b05f-419b-bdb3-65661a76b8ce",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "motion": {
                  "ID": "b88b8ef6-52db-4d38-a6dc-38969f125979",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "motion",
                  "Value": "CLOSED",
                  "Attrs": {
                    "$datatype": {
                      "ID": "a780149e-e87a-4be6-8b6f-af201b7ab9f1",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "3fd674e6-0f41-4964-96a0-6e6ccfb56844",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "f12e7509-f898-4e18-9340-a8c21befdabc",
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
              "ID": "695bc975-5d21-462d-b631-64b704ceccd9",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "hardware",
              "Attrs": {
                "$name": {
                  "ID": "e6cb66c8-246c-4d88-a45d-21da0a858e0a",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$name",
                  "Value": "Device Info"
                },
                "$properties": {
                  "ID": "32102d6e-ac99-4a26-bfc4-33fc782f2ace",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$properties",
                  "Value": "signal,mac,resetReason,voltage"
                },
                "$type": {
                  "ID": "261255aa-3929-4602-99aa-93aa6686d72f",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "mac": {
                  "ID": "24e610f9-8ef2-415d-90de-e3c9d955733a",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "mac",
                  "Value": "BC:DD:C2:81:04:72",
                  "Attrs": {
                    "$datatype": {
                      "ID": "dd594057-149c-432a-aafc-b496eb04f895",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": "sring"
                    },
                    "$name": {
                      "ID": "65772269-33f8-4af6-b7fb-4f12efb4151d",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$name",
                      "Value": "mac"
                    }
                  }
                },
                "resetReason": {
                  "ID": "2840cd72-a9f7-4ae6-aaf5-7516385aa7c2",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "resetReason",
                  "Value": "External System",
                  "Attrs": {
                    "$datatype": {
                      "ID": "09b39b92-a7b8-44c2-b25f-76c3f7fd6eb0",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$name": {
                      "ID": "438fe40f-07a5-4046-883c-3dfd029a4ce9",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$name",
                      "Value": "Last Reset Reason"
                    }
                  }
                },
                "signal": {
                  "ID": "6eac0471-7508-4228-a428-2fc7bae5b324",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "signal",
                  "Value": "-64",
                  "Attrs": {
                    "$datatype": {
                      "ID": "f76caef2-cc65-4783-8c38-0cd4fc53a541",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$name": {
                      "ID": "86af1fff-a154-4938-8141-f2889d02f30a",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$name",
                      "Value": "RSSI"
                    },
                    "$unit": {
                      "ID": "65784a8b-d42d-46d5-8b7d-6bb22e3297f0",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": "dBm"
                    }
                  }
                },
                "voltage": {
                  "ID": "528af9d3-bb63-4803-a5a9-ac982c606f43",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "voltage",
                  "Value": "2.98",
                  "Attrs": {
                    "$datatype": {
                      "ID": "026c9ec0-9f4a-48dd-9f5f-65bfead58cc4",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "e432c91b-d651-401b-87be-0bd8cd5d79f6",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$name",
                      "Value": "3.3V Supply"
                    },
                    "$unit": {
                      "ID": "31477b97-b2ba-476e-ae9a-d50ce5cb26de",
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
        },
        "MCPProvider": {
          "ID": "66e0038fbeb3c18a787483b03732a372",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "MCPProvider",
          "Attrs": {
            "$fw": {
              "ID": "b8a4d5b4-aa4f-403a-90a6-494455994346",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "9c34b8e1-b2cb-4a71-8665-9b2378ba776c",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "b92b40fead22529097629c0597186c05",
                  "Props": {}
                },
                "name": {
                  "ID": "93706154-c57e-48fb-811b-7b6949c8ee75",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "WiredProvider",
                  "Props": {}
                },
                "version": {
                  "ID": "76e940db-04dd-414c-a566-65c1016cecf8",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "1.4.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "ed29fa55-03fe-4c10-967d-1c2484dcce54",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "783fbf56-bf99-423f-85d7-32dcb6a469f5",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$implementation",
              "Value": "esp8266",
              "Props": {
                "config": {
                  "ID": "5d222c59-b646-419d-a1f1-0b94c4629217",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"Hard Wired Alarm Sensors\",\"device_id\":\"MCPProvider\",\"device_stats_interval\":300,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"ipolA\":255,\"ipolB\":255}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "92793726-b1bc-4d2a-92e3-9862d4b433f2",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "ota": {
                      "ID": "7674d4b5-00e9-48e1-8457-3c1a6cc456a3",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "6d7aa36b-e775-49ce-a152-995524310b22",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "6b5d8d8a-5f70-4e11-9197-a52b7251088f",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$localip",
              "Value": "10.100.1.222",
              "Props": {}
            },
            "$mac": {
              "ID": "0ed760af-7725-4eb8-82d8-b9069b7cc6fa",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$mac",
              "Value": "BC:DD:C2:24:B7:3C",
              "Props": {}
            },
            "$name": {
              "ID": "58b73ed4-ef0a-4823-a32c-cded3a2dee60",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$name",
              "Value": "Hard Wired Alarm Sensors",
              "Props": {}
            },
            "$nodes": {
              "ID": "7b07de26-970f-4fbe-bbc7-d363b458674f",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$nodes",
              "Value": "wiredMonitor,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "25735a06-c7f5-4179-9426-7ca241549ce8",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "defacc5c-e2ed-465f-9a23-c1520876473c",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$stats",
              "Value": "uptime",
              "Props": {
                "interval": {
                  "ID": "dfd4dbb9-9601-4c54-9919-9336e6e0e2b3",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "305",
                  "Props": {}
                },
                "signal": {
                  "ID": "0cadb465-e9eb-43a8-9507-4e0642e09e0e",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "56",
                  "Props": {}
                },
                "uptime": {
                  "ID": "a3262176-f0dc-4d65-883a-c5ff5573b563",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "113110",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "hardware": {
              "ID": "1da2837e-562f-413a-aa54-0b266aa60e44",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "hardware",
              "Attrs": {
                "$name": {
                  "ID": "26d63bf8-881f-4b80-8394-4a8b601c6c9e",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$name",
                  "Value": "Device Info"
                },
                "$properties": {
                  "ID": "f22bfa75-be10-4ad7-9475-00d5945ec501",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$properties",
                  "Value": "signal,mac,resetReason,voltage"
                },
                "$type": {
                  "ID": "6a1cc9d4-0877-4f37-9f06-c7526dfb64cc",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "signal": {
                  "ID": "045a7a7b-b382-4c19-9220-7d5e1d55ecbc",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "signal",
                  "Value": "-71",
                  "Attrs": {
                    "$datatype": {
                      "ID": "c9bfdf6e-d7bf-49a2-b3ce-7cb85f712cc3",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$name": {
                      "ID": "71d8fe47-9e47-4635-bae3-798c4d0069cf",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$name",
                      "Value": "RSSI"
                    },
                    "$unit": {
                      "ID": "c4bfe17d-6fee-40e6-a4ce-bcde1951c106",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": "dBm"
                    }
                  }
                }
              }
            },
            "wiredMonitor": {
              "ID": "ddd40200-b352-4b08-afe8-7a08b4c0ca0c",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "wiredMonitor",
              "Attrs": {
                "$name": {
                  "ID": "e3893ae3-a655-417c-80ad-e910647902a9",
                  "ElementType": 15,
                  "Parent": "wiredMonitor",
                  "Name": "$name",
                  "Value": "Wired Sensors"
                },
                "$properties": {
                  "ID": "791ffb5f-afde-4b43-9db7-15672b53655f",
                  "ElementType": 15,
                  "Parent": "wiredMonitor",
                  "Name": "$properties",
                  "Value": "pin0,pin1,pin2,pin3,pin4,pin5,pin6,pin7,pin8,pin9,pin10,pin11,pin12,pin13,pin14,pin15"
                },
                "$type": {
                  "ID": "fbe2bde0-de41-478b-9462-d3f4f764163f",
                  "ElementType": 15,
                  "Parent": "wiredMonitor",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "pin0": {
                  "ID": "81c230df-ec41-4213-9f4d-a9e552ff8d22",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin0",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "e6a79cd2-a163-484d-a797-ee487e1b9d9c",
                      "ElementType": 17,
                      "Parent": "pin0",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "a565b278-df8f-43bb-938d-79d526cc9398",
                      "ElementType": 17,
                      "Parent": "pin0",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "541e142f-7497-48bc-a29a-deab6f6e4b76",
                      "ElementType": 17,
                      "Parent": "pin0",
                      "Name": "$name",
                      "Value": "Pin 0"
                    }
                  }
                },
                "pin1": {
                  "ID": "3677bbf0-1db0-4ee3-abb1-f5795b1e95e3",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin1",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "4c31a201-b625-4e04-956f-0ed5fad7a7aa",
                      "ElementType": 17,
                      "Parent": "pin1",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "ce824ba6-a1d6-42e0-a9ec-647dc2ee6dfc",
                      "ElementType": 17,
                      "Parent": "pin1",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "4423931a-db61-431e-aa0b-9a2a6befcbce",
                      "ElementType": 17,
                      "Parent": "pin1",
                      "Name": "$name",
                      "Value": "Pin 1"
                    }
                  }
                },
                "pin10": {
                  "ID": "e9532dac-4b4f-446d-9aba-31dfeadb8ed9",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin10",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "350b9dd8-ed23-4565-a428-99145856a43b",
                      "ElementType": 17,
                      "Parent": "pin10",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "c065df3c-a4f3-4086-8e15-fe5fde25cb5b",
                      "ElementType": 17,
                      "Parent": "pin10",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "b7c727e8-09a4-48a5-9e9a-8b9bb50a8f5a",
                      "ElementType": 17,
                      "Parent": "pin10",
                      "Name": "$name",
                      "Value": "Pin 10"
                    }
                  }
                },
                "pin11": {
                  "ID": "a97d144b-a6b3-4db7-996f-7280d8498dd4",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin11",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "54110be3-af29-445d-89ed-ddaf73a6233d",
                      "ElementType": 17,
                      "Parent": "pin11",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "f5356ce6-13f4-4fe3-b89f-75ebca219b46",
                      "ElementType": 17,
                      "Parent": "pin11",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "4b1ecdb9-0f4a-40c6-9764-a89fc5db40ac",
                      "ElementType": 17,
                      "Parent": "pin11",
                      "Name": "$name",
                      "Value": "Pin 11"
                    }
                  }
                },
                "pin12": {
                  "ID": "f722e67e-4898-4aac-a36e-6def19de127a",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin12",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "997bbbe6-a52e-427f-95e4-1862175fc75a",
                      "ElementType": 17,
                      "Parent": "pin12",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "3db6e2af-9e40-4809-9d52-f0c53896c8a5",
                      "ElementType": 17,
                      "Parent": "pin12",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "97f4075e-5130-4382-a5cb-aa5f79c36ca0",
                      "ElementType": 17,
                      "Parent": "pin12",
                      "Name": "$name",
                      "Value": "Pin 12"
                    }
                  }
                },
                "pin13": {
                  "ID": "af914914-e1a0-4682-b6f0-722f11b7c524",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin13",
                  "Value": "CLOSED",
                  "Attrs": {
                    "$datatype": {
                      "ID": "b7048f4f-ae5a-4dd3-b672-0e03ba7ed3ed",
                      "ElementType": 17,
                      "Parent": "pin13",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "c4f6b443-eb8e-4840-9bd1-a4e80ec41991",
                      "ElementType": 17,
                      "Parent": "pin13",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "0bf2ba42-c659-4c07-bd19-0ab4285f7e7a",
                      "ElementType": 17,
                      "Parent": "pin13",
                      "Name": "$name",
                      "Value": "Pin 13"
                    }
                  }
                },
                "pin14": {
                  "ID": "d4eb3023-0812-48e3-91d5-a0d950c5263b",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin14",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "fc51bd42-4672-4e1d-a7e9-c380327be677",
                      "ElementType": 17,
                      "Parent": "pin14",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "c61cef95-8b04-4788-8754-d327f9980051",
                      "ElementType": 17,
                      "Parent": "pin14",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "a6f9d240-d3e6-4231-8111-d6f88b5b5799",
                      "ElementType": 17,
                      "Parent": "pin14",
                      "Name": "$name",
                      "Value": "Pin 14"
                    }
                  }
                },
                "pin15": {
                  "ID": "c9bf199f-b9af-443d-8161-1b7960a4c8b7",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin15",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "9c4292ad-bc08-45f3-ace9-4afab6dc0355",
                      "ElementType": 17,
                      "Parent": "pin15",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "33d2d103-4e05-4204-a735-7babe35ec299",
                      "ElementType": 17,
                      "Parent": "pin15",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "22bbc8c6-c462-4443-94b9-93af4894b45a",
                      "ElementType": 17,
                      "Parent": "pin15",
                      "Name": "$name",
                      "Value": "Pin 15"
                    }
                  }
                },
                "pin2": {
                  "ID": "5e97bcd4-e9f8-40ce-ae4e-2c665d73b226",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin2",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "b864ed67-2990-42c5-9b96-5631734b83bc",
                      "ElementType": 17,
                      "Parent": "pin2",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "7f4c6687-f965-44f2-9539-f537da7d0f24",
                      "ElementType": 17,
                      "Parent": "pin2",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "d04b526a-4159-4175-87b5-461f8b629591",
                      "ElementType": 17,
                      "Parent": "pin2",
                      "Name": "$name",
                      "Value": "Pin 2"
                    }
                  }
                },
                "pin3": {
                  "ID": "f7b850e8-f501-4181-89da-52616d640e4f",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin3",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "4e1b9e8b-07c0-44c4-bfac-2f0e1d334a87",
                      "ElementType": 17,
                      "Parent": "pin3",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "31791ea2-13a6-4131-8e2b-dcdf745f5cd2",
                      "ElementType": 17,
                      "Parent": "pin3",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "1ab0d403-f0ed-4028-818c-ce813981b5b4",
                      "ElementType": 17,
                      "Parent": "pin3",
                      "Name": "$name",
                      "Value": "Pin 3"
                    }
                  }
                },
                "pin4": {
                  "ID": "b633ad93-9342-4d47-9718-0323db5e5b5a",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin4",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "1646bd28-0979-4221-8229-47dc86e53162",
                      "ElementType": 17,
                      "Parent": "pin4",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "70041aa4-68ad-40b5-a019-967fbec2155c",
                      "ElementType": 17,
                      "Parent": "pin4",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "5b8b1a64-97cb-478e-9c19-498b640deca7",
                      "ElementType": 17,
                      "Parent": "pin4",
                      "Name": "$name",
                      "Value": "Pin 4"
                    }
                  }
                },
                "pin5": {
                  "ID": "1a2ab47b-df10-47a3-88fc-b591b728b55b",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin5",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "9644bf03-4ec3-42e7-b180-763b65b8aa62",
                      "ElementType": 17,
                      "Parent": "pin5",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "9d8093b9-da9c-41d5-ad35-60cbfa6a2926",
                      "ElementType": 17,
                      "Parent": "pin5",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "6648146c-2a5f-4e30-a9ab-47d6e83fba9b",
                      "ElementType": 17,
                      "Parent": "pin5",
                      "Name": "$name",
                      "Value": "Pin 5"
                    }
                  }
                },
                "pin6": {
                  "ID": "43922b5c-8c7b-4558-aa75-3e2895c06f76",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin6",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "3bb4da20-cd91-48a1-9ae1-94a2f7d0cf3d",
                      "ElementType": 17,
                      "Parent": "pin6",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "116eef44-174e-4cdd-881e-46113a0daeb4",
                      "ElementType": 17,
                      "Parent": "pin6",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "338cc5e6-169c-4202-8d79-71a0c4ddfe98",
                      "ElementType": 17,
                      "Parent": "pin6",
                      "Name": "$name",
                      "Value": "Pin 6"
                    }
                  }
                },
                "pin7": {
                  "ID": "e8a80a35-5a6e-4869-808b-03faa43d1cd2",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin7",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "0aff2742-0336-4f93-988a-4cea936dbeb5",
                      "ElementType": 17,
                      "Parent": "pin7",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "b58ecc94-085e-4813-bc6f-2c1fd083f9e8",
                      "ElementType": 17,
                      "Parent": "pin7",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "f3020f86-a037-472f-b6b8-aaa6ba1e1163",
                      "ElementType": 17,
                      "Parent": "pin7",
                      "Name": "$name",
                      "Value": "Pin 7"
                    }
                  }
                },
                "pin8": {
                  "ID": "27d85c72-be8b-4f76-85c0-9afe496bbc4a",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin8",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "e227df83-5b74-41f1-8e48-9f832372e86e",
                      "ElementType": 17,
                      "Parent": "pin8",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "325f9cb0-0507-4ea3-8691-fa305dd0f7a8",
                      "ElementType": 17,
                      "Parent": "pin8",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "c2b6f5d3-b274-4814-b55d-c21b677f9498",
                      "ElementType": 17,
                      "Parent": "pin8",
                      "Name": "$name",
                      "Value": "Pin 8"
                    }
                  }
                },
                "pin9": {
                  "ID": "af4bf01f-57b8-4661-adbb-9377e097d070",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin9",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "6f08f580-17e8-47e7-a87b-c8153a2280c7",
                      "ElementType": 17,
                      "Parent": "pin9",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "4f8b18f9-efab-4276-bd48-0ae030bbccfc",
                      "ElementType": 17,
                      "Parent": "pin9",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "93df54e9-b639-4f15-a105-434fc4c88d86",
                      "ElementType": 17,
                      "Parent": "pin9",
                      "Name": "$name",
                      "Value": "Pin 9"
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

# From Reload after empty
```json
{
  "ID": "b2c4ab7c-2a48-4d41-b840-13a06175a93c",
  "ElementType": 24,
  "SiteName": "Skoona Consulting",
  "Title": "Homie Monitor (GOLANG)",
  "Names": [
    "sknSensors"
  ],
  "Broadcasts": [
    {
      "ID": "5226dd52a6c288ed3301b9302d325133",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "alarm",
      "Level": "item",
      "Value": "Media Room Window Broken",
      "Received": "2021-07-09T15:38:51.2862-04:00"
    },
    {
      "ID": "8490e6cbf39eeadd991cb2a346580e51",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "alert",
      "Level": "door",
      "Value": "Door Open",
      "Received": "2021-07-09T15:38:51.286203-04:00"
    },
    {
      "ID": "c6c659175e0f9ca804664b3c6e3fda37",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "alert",
      "Level": "item",
      "Value": "Garage Door Open",
      "Received": "2021-07-09T15:38:51.286204-04:00"
    },
    {
      "ID": "df9b91ec197689e2c566aea2add68ad5",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "notice",
      "Level": "item",
      "Value": "openHAB3 Online",
      "Received": "2021-07-09T15:38:51.286206-04:00"
    },
    {
      "ID": "c254b24c700dbc2ecba237869f9b493b",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "$broadcast",
      "Level": "alert",
      "Value": "openHAB3 Offline",
      "Received": "2021-07-09T15:38:55.376377-04:00"
    },
    {
      "ID": "c254b24c700dbc2ecba237869f9b493b",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "$broadcast",
      "Level": "alert",
      "Value": "Door Open",
      "Received": "2021-07-09T15:38:55.414332-04:00"
    }
  ],
  "Firmwares": [
    {
      "ID": "5a3a2c57b5084ee4243017beaa0df466",
      "ElementType": 21,
      "Name": "Environment-DS18B20",
      "FileName": "./dataDB/firmwares/Environment-DS18B20.bin",
      "Version": "1.0.2",
      "Path": "./dataDB/firmwares/Environment-DS18B20.bin",
      "Size": 561456,
      "MD5Digest": "5a3a2c57b5084ee4243017beaa0df466",
      "Brand": "SknSensors",
      "Created": "2021-05-13T23:25:08.350185372-04:00"
    },
    {
      "ID": "3ef8fbb48c5b23788a22a998c14a1a6d",
      "ElementType": 21,
      "Name": "Monitor-DHT-RCWL-Metrics",
      "FileName": "./dataDB/firmwares/Monitor-DHT-RCWL-Metrics-200.bin",
      "Version": "2.0.0",
      "Path": "./dataDB/firmwares/Monitor-DHT-RCWL-Metrics-200.bin",
      "Size": 399360,
      "MD5Digest": "3ef8fbb48c5b23788a22a998c14a1a6d",
      "Brand": "SknSensors",
      "Created": "2021-05-13T23:25:08.352490958-04:00"
    },
    {
      "ID": "615fed382ab44bd43fe83508aecac682",
      "ElementType": 21,
      "Name": "Monitor-SHT31-RCWL-Metrics",
      "FileName": "./dataDB/firmwares/Monitor-SHT31-RCWL-Metrics-200.bin",
      "Version": "2.0.0",
      "Path": "./dataDB/firmwares/Monitor-SHT31-RCWL-Metrics-200.bin",
      "Size": 402640,
      "MD5Digest": "615fed382ab44bd43fe83508aecac682",
      "Brand": "SknSensors",
      "Created": "2021-05-13T23:25:08.354699188-04:00"
    }
  ],
  "Schedules": {},
  "DeviceNetworks": {
    "sknSensors": {
      "ID": "6897631feedfa370c777bc6b6f6c9b3f",
      "Title": "Restored",
      "ElementType": 23,
      "Name": "sknSensors",
      "Devices": {
        "D1R1MiniA": {
          "ID": "908a089bd74911ba7558c0092b1c55ed",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "D1R1MiniA",
          "Attrs": {
            "$fw": {
              "ID": "41b53f5b-4785-44b4-b7ca-e341a1900faf",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "version": {
                  "ID": "cecd1f54-79d0-4bc6-853c-0165378a222e",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "1.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "7b3f1ab4-8c1b-48e4-8bc2-4b85686db637",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "61fec600-b7a3-42f8-8018-8f8d363a04c7",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$implementation",
              "Value": "",
              "Props": {
                "version": {
                  "ID": "cc7505ed-1580-419f-9796-e6e5875f0cac",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "5f4693b0-5a8b-4a49-810f-0f3d3f108457",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$localip",
              "Value": "10.100.1.181",
              "Props": {}
            },
            "$mac": {
              "ID": "e411d48c-7a9f-43f8-9aed-ac76e49e0921",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$mac",
              "Value": "84:F3:EB:0C:38:6F",
              "Props": {}
            },
            "$name": {
              "ID": "ebfa083a-b2a0-49e4-8fcd-d63ea6692b0a",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$name",
              "Value": "Furnace",
              "Props": {}
            },
            "$nodes": {
              "ID": "e1a9165f-93d1-4c57-948e-4f0dcf676d3c",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$nodes",
              "Value": "Liquids",
              "Props": {}
            },
            "$state": {
              "ID": "a00482b3-f9b5-4b96-9ec3-6e6e367aa291",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "1c7dd466-2f17-4586-a6c6-46727a34590b",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$stats",
              "Value": "",
              "Props": {
                "uptime": {
                  "ID": "bb74b4f7-9687-436d-a585-c4bddddec14f",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "114307",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "Liquids": {
              "ID": "ef85a83e-81ea-4d4d-97bd-ac6425b563ad",
              "ElementType": 14,
              "Parent": "D1R1MiniA",
              "Name": "Liquids",
              "Attrs": {},
              "Props": {
                "Liquids": {
                  "ID": "4dfd68ad-751a-4b80-b5dd-e63e7ab8fb8b",
                  "ElementType": 16,
                  "Parent": "Liquids",
                  "Name": "Liquids",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "51f645e8-9fa8-4ca6-9085-9f37d5e80eca",
                      "ElementType": 17,
                      "Parent": "level",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$format": {
                      "ID": "8a7a271c-e8df-4ca9-a9d7-4cd8205feadc",
                      "ElementType": 17,
                      "Parent": "level",
                      "Name": "$format",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "5555de47-917e-496c-ae20-c8a62ccc03f6",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Water Sensor"
                    },
                    "$properties": {
                      "ID": "ab135074-4a59-4332-947b-292417796b15",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "level,volts"
                    },
                    "$type": {
                      "ID": "d15418b9-01b5-48ad-b97a-b5a5213c5d9e",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "$unit": {
                      "ID": "18f570f7-99c6-4023-86d7-e7408d1f4419",
                      "ElementType": 17,
                      "Parent": "level",
                      "Name": "$unit",
                      "Value": ""
                    },
                    "level": {
                      "ID": "44211f9b-98c5-4099-b8be-ce4e15f1aa6c",
                      "ElementType": 17,
                      "Parent": "level",
                      "Name": "level",
                      "Value": "7"
                    },
                    "volts": {
                      "ID": "419d6009-0b2c-42ff-b5a8-4347ebd1b6e6",
                      "ElementType": 17,
                      "Parent": "volts",
                      "Name": "volts",
                      "Value": "0.0"
                    }
                  }
                }
              }
            }
          }
        },
        "DSMonitor": {
          "ID": "280cff5165f95cca0fc6f9bcbe2bbc79",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "DSMonitor",
          "Attrs": {
            "$fw": {
              "ID": "5025af0b-ff5a-40e1-af9f-fe92f6e47c86",
              "ElementType": 11,
              "Parent": "DSMonitor",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "version": {
                  "ID": "fd2c66a9-f65c-4a29-9b27-bb0749db2621",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "1.0.2",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "4d8d874e-a4e3-455a-b772-4977e18779de",
              "ElementType": 11,
              "Parent": "DSMonitor",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "57486d0f-c5e5-4886-9211-7426e9a020a2",
              "ElementType": 11,
              "Parent": "DSMonitor",
              "Name": "$implementation",
              "Value": "",
              "Props": {
                "version": {
                  "ID": "16782b11-ef15-4bba-b637-97b675572db5",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "bcccaa6f-0664-4aba-b991-0aa005b83211",
              "ElementType": 11,
              "Parent": "DSMonitor",
              "Name": "$localip",
              "Value": "10.100.1.164",
              "Props": {}
            },
            "$mac": {
              "ID": "545604c1-4faa-4d0b-947c-56427015ad52",
              "ElementType": 11,
              "Parent": "DSMonitor",
              "Name": "$mac",
              "Value": "DC:4F:22:3B:33:4F",
              "Props": {}
            },
            "$name": {
              "ID": "59b03543-e737-4182-8ccb-0ab2c4578a92",
              "ElementType": 11,
              "Parent": "DSMonitor",
              "Name": "$name",
              "Value": "Environment-DS18B20",
              "Props": {}
            },
            "$nodes": {
              "ID": "4bea5500-1389-489c-b1ae-15222b59b6f1",
              "ElementType": 11,
              "Parent": "DSMonitor",
              "Name": "$nodes",
              "Value": "Ambient[]",
              "Props": {}
            },
            "$state": {
              "ID": "74cd8a2d-e4aa-4e21-9f74-70dd11593dda",
              "ElementType": 11,
              "Parent": "DSMonitor",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "a52d39f6-ea1a-467d-8d20-4ab67e5c5b73",
              "ElementType": 11,
              "Parent": "DSMonitor",
              "Name": "$stats",
              "Value": "",
              "Props": {
                "uptime": {
                  "ID": "304f43e0-3eea-46b7-911c-d5721f29084a",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "5",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "Ambient": {
              "ID": "93c4a239-15e3-4b2e-b9e1-dc041dfe475c",
              "ElementType": 14,
              "Parent": "DSMonitor",
              "Name": "Ambient",
              "Attrs": {},
              "Props": {
                "Ambient": {
                  "ID": "3668f01e-602c-4bb7-8230-aa75df9c6bea",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": "",
                  "Attrs": {
                    "$array": {
                      "ID": "fcdd66a9-72f7-4ddf-b7bc-115ee8e61193",
                      "ElementType": 17,
                      "Parent": "$array",
                      "Name": "$array",
                      "Value": "0-3"
                    },
                    "$datatype": {
                      "ID": "febb8b1a-47e5-4dcf-a644-d1ea80ce6245",
                      "ElementType": 17,
                      "Parent": "state",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$format": {
                      "ID": "29fbb972-033b-4e61-85ea-9efe6989f384",
                      "ElementType": 17,
                      "Parent": "state",
                      "Name": "$format",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "0218c765-e56f-4dda-8ca6-fd8342a51b03",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Environment Monitor"
                    },
                    "$properties": {
                      "ID": "95d64d1d-8dff-402b-8e2d-b92b8df00562",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "state,temperature"
                    },
                    "$type": {
                      "ID": "0ca1838c-1f0a-4ecd-97c7-d1444cc743d0",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "$unit": {
                      "ID": "3ace686a-1f5a-4c30-bfdf-a42966342245",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$unit",
                      "Value": ""
                    }
                  }
                }
              }
            },
            "Ambient_0": {
              "ID": "4a6c9e24-9832-4f49-beb5-55a13b55a07b",
              "ElementType": 14,
              "Parent": "DSMonitor",
              "Name": "Ambient_0",
              "Attrs": {},
              "Props": {
                "Ambient_0": {
                  "ID": "a9f79661-824c-4ca8-adda-ca876bcb0a53",
                  "ElementType": 16,
                  "Parent": "Ambient_0",
                  "Name": "Ambient_0",
                  "Value": "",
                  "Attrs": {
                    "$name": {
                      "ID": "d2d020e8-ea10-44ae-a96d-71eee7ea444e",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Ambient_0"
                    },
                    "state": {
                      "ID": "3c06b5ef-0b17-4e98-9b20-32673e0e5b12",
                      "ElementType": 17,
                      "Parent": "state",
                      "Name": "state",
                      "Value": "OK"
                    },
                    "temperature": {
                      "ID": "734ec6a7-ada5-4714-8fb5-27294a23c366",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "temperature",
                      "Value": "73.06"
                    }
                  }
                }
              }
            },
            "Ambient_1": {
              "ID": "948e1626-fe9f-44ec-97e2-4addc0e9c819",
              "ElementType": 14,
              "Parent": "DSMonitor",
              "Name": "Ambient_1",
              "Attrs": {},
              "Props": {
                "Ambient_1": {
                  "ID": "7801c431-309d-4008-9067-43acb54af03d",
                  "ElementType": 16,
                  "Parent": "Ambient_1",
                  "Name": "Ambient_1",
                  "Value": "",
                  "Attrs": {
                    "$name": {
                      "ID": "59ccf14c-16d1-4589-b227-3db00950d8b5",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Ambient_1"
                    },
                    "state": {
                      "ID": "61005c3a-840d-4675-98b5-2bcfae2c8d41",
                      "ElementType": 17,
                      "Parent": "state",
                      "Name": "state",
                      "Value": "OK"
                    },
                    "temperature": {
                      "ID": "9b2ee6b4-4f24-463d-bea4-59ecd1bd8ce9",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "temperature",
                      "Value": "73.96"
                    }
                  }
                }
              }
            },
            "Ambient_2": {
              "ID": "774bbb09-e7c3-402e-b125-ed48f7cf50cb",
              "ElementType": 14,
              "Parent": "DSMonitor",
              "Name": "Ambient_2",
              "Attrs": {},
              "Props": {
                "Ambient_2": {
                  "ID": "f5b8a30c-67d0-4d20-82e7-ae3aee482af7",
                  "ElementType": 16,
                  "Parent": "Ambient_2",
                  "Name": "Ambient_2",
                  "Value": "",
                  "Attrs": {
                    "$name": {
                      "ID": "71a2ce48-1526-43b0-963e-25a755849b84",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Ambient_2"
                    },
                    "state": {
                      "ID": "fe85a281-caf2-4c66-b857-5b38768fc075",
                      "ElementType": 17,
                      "Parent": "state",
                      "Name": "state",
                      "Value": "OK"
                    },
                    "temperature": {
                      "ID": "9b7bc336-747d-468c-9d0c-eb45e11d1249",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "temperature",
                      "Value": "73.63"
                    }
                  }
                }
              }
            },
            "Ambient_3": {
              "ID": "6e761852-3dfc-4864-b929-11c8948c4c40",
              "ElementType": 14,
              "Parent": "DSMonitor",
              "Name": "Ambient_3",
              "Attrs": {},
              "Props": {
                "Ambient_3": {
                  "ID": "a49e5eaa-929b-43dd-a652-d353392ab073",
                  "ElementType": 16,
                  "Parent": "Ambient_3",
                  "Name": "Ambient_3",
                  "Value": "",
                  "Attrs": {
                    "$name": {
                      "ID": "5e5325a2-9bd9-40d5-9d49-90c761f9a6a5",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Ambient_3"
                    },
                    "state": {
                      "ID": "a6a29a7b-1d0e-4cbf-9ab0-5aae79b165ab",
                      "ElementType": 17,
                      "Parent": "state",
                      "Name": "state",
                      "Value": "OK"
                    },
                    "temperature": {
                      "ID": "b19a1b8a-0e00-4767-bd8b-67544a0b8983",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "temperature",
                      "Value": "72.39"
                    }
                  }
                }
              }
            }
          }
        },
        "FamilyRoom": {
          "ID": "2f63939498209cf4584b6e82954d2407",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "FamilyRoom",
          "Attrs": {
            "$fw": {
              "ID": "21570c29-ee1a-40e5-964b-e5cb45cee9a7",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "version": {
                  "ID": "4003ce2e-40fb-4da9-a09e-68c08953444a",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "2.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "18f7cb79-ff92-4132-a6eb-0eabbd3f044c",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "6e4c48f4-0d82-4a2d-a65f-78985b5a8304",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$implementation",
              "Value": "",
              "Props": {
                "version": {
                  "ID": "c8ef63f2-2999-4309-9820-bb07caba6fb0",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "02406709-936f-44d6-ba5c-4b1685d4c66d",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$localip",
              "Value": "10.100.1.166",
              "Props": {}
            },
            "$mac": {
              "ID": "09f47c14-c8af-4079-9f1b-7784a2b410ed",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$mac",
              "Value": "BC:DD:C2:E5:C4:24",
              "Props": {}
            },
            "$name": {
              "ID": "ba9f7588-bc22-4e13-90d9-44757c5e51d0",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$name",
              "Value": "FamilyRoom",
              "Props": {}
            },
            "$nodes": {
              "ID": "9dcda820-465e-49c8-b572-429dec560cec",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$nodes",
              "Value": "Ambient,Presence,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "83f9cb20-6c05-46dd-8b31-c1008d6f69b3",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "183f985f-7c4a-4d4d-8450-9b4fc71b0238",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$stats",
              "Value": "",
              "Props": {
                "uptime": {
                  "ID": "22d5fcdd-add8-4cf6-a9b3-ce08aaee9ef1",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "423819",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "Ambient": {
              "ID": "293d8b06-1983-4b22-a4fa-c39a661842ee",
              "ElementType": 14,
              "Parent": "FamilyRoom",
              "Name": "Ambient",
              "Attrs": {},
              "Props": {
                "Ambient": {
                  "ID": "81590b63-362e-4395-8052-ce9dfcb23f6f",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "6764204d-e359-4423-82f3-19ad115b650d",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "e3df68f7-3d4d-4a8b-862f-2ad9ecc16ce9",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Temperature and Humidity Sensor"
                    },
                    "$properties": {
                      "ID": "c031c19d-16a1-4c64-9638-33ab6b8674ac",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "humidity,temperature"
                    },
                    "$type": {
                      "ID": "9fe909d8-4bea-45a8-9eb9-b6fab7574b4d",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "$unit": {
                      "ID": "d0eb684f-f3a3-4fb8-a0ca-ef1579347dc9",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": ""
                    },
                    "humidity": {
                      "ID": "6668595b-0dc6-414e-bcea-4cf1d70029e8",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "humidity",
                      "Value": "20.50"
                    },
                    "temperature": {
                      "ID": "87675d10-3303-4705-97b4-4fe73512f5bb",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "temperature",
                      "Value": "67.64"
                    }
                  }
                }
              }
            },
            "Presence": {
              "ID": "510442dc-7a97-4524-9d75-bf0b1c71559c",
              "ElementType": 14,
              "Parent": "FamilyRoom",
              "Name": "Presence",
              "Attrs": {},
              "Props": {
                "Presence": {
                  "ID": "6a68b494-1072-4f33-abb5-eed6e661d967",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "d78f5e05-f995-4904-8775-40ff5c733e60",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$format": {
                      "ID": "37c27c2e-526f-485c-85dd-d70aa83a854d",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "1b9671de-aa13-4195-832e-cc07ac05a99f",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Motion Sensor"
                    },
                    "$properties": {
                      "ID": "90f2b947-3f08-4664-8759-ccdbad4bc177",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "motion"
                    },
                    "$type": {
                      "ID": "db0e6cc7-6e0e-41cb-b077-839d3367ce69",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "motion": {
                      "ID": "2e7a2a4f-4d81-48ed-9f72-bd7273011a7c",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "motion",
                      "Value": "OPEN"
                    }
                  }
                }
              }
            },
            "hardware": {
              "ID": "942291a3-ad58-4928-8174-2c5ec77a2bcb",
              "ElementType": 14,
              "Parent": "FamilyRoom",
              "Name": "hardware",
              "Attrs": {},
              "Props": {
                "hardware": {
                  "ID": "f4e8d958-2bd8-4b03-a5e7-0ae3c561e8a8",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "e5ff1c48-4499-43a6-8c18-8d6633830208",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "4056b1ff-a198-4f12-ac02-716a266dcbc4",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Device Info"
                    },
                    "$properties": {
                      "ID": "e177e5c2-6c6b-483e-80bf-55b5ccce328c",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "signal,mac,resetReason,voltage"
                    },
                    "$type": {
                      "ID": "7749b40f-aa8d-4a1f-94fe-ebb0cd2e89a5",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "$unit": {
                      "ID": "f69f3133-0daa-482e-b016-212f1de09c24",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": ""
                    },
                    "mac": {
                      "ID": "f02236e6-0de7-4399-ba3e-87a9bf1c695c",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "mac",
                      "Value": "BC:DD:C2:E5:C4:24"
                    },
                    "resetReason": {
                      "ID": "9c28ea02-21a9-4496-9765-4f70c16cb99e",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "resetReason",
                      "Value": "External System"
                    },
                    "signal": {
                      "ID": "8a47e06d-e654-4c66-92e6-d49b912ec838",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "signal",
                      "Value": "-62"
                    },
                    "voltage": {
                      "ID": "53effe0b-b216-4b09-91b8-4a5b03833da5",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "voltage",
                      "Value": "3.03"
                    }
                  }
                }
              }
            }
          }
        },
        "GarageMonitor": {
          "ID": "36ec4f3d3eedbe07b28e5fa6ab6ddb9b",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "GarageMonitor",
          "Attrs": {
            "$fw": {
              "ID": "b473142f-a670-4ec9-83c6-ef8b763c88fd",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "f18d471a-3c2f-4491-b51d-4c441bf0be0f",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "615fed382ab44bd43fe83508aecac682",
                  "Props": {}
                },
                "name": {
                  "ID": "6c0acf80-ecaf-428f-bb99-8e1168d3ef48",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "Monitor-SHT31-RCWL-Metrics",
                  "Props": {}
                },
                "version": {
                  "ID": "2003d05e-1a4e-4505-9637-7c385e4ae8a7",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "2.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "7fe80c37-3b5d-49ad-b319-ea8dffdaa80d",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "2b3b649e-f470-41ad-83dc-344a89889ff2",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$implementation",
              "Value": "",
              "Props": {
                "config": {
                  "ID": "88243d96-c3af-4b73-9248-53033d06203b",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"Garage Monitor\",\"device_id\":\"GarageMonitor\",\"device_stats_interval\":900,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"sensorInterval\":900,\"motionHoldInterval\":60}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "984fef25-e581-45ad-967f-f11223c0f545",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "ota": {
                      "ID": "5043ddfa-62e8-4e54-bec1-b93f22f296f6",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "07f20bf0-c888-468a-9e9c-f97a39348680",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "c051de1e-080a-4d7e-b7f7-8d9d72d7cff6",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$localip",
              "Value": "10.100.1.177",
              "Props": {}
            },
            "$mac": {
              "ID": "1690f46b-af64-4a0b-b7de-99c203121b23",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$mac",
              "Value": "B4:E6:2D:1B:5C:4D",
              "Props": {}
            },
            "$name": {
              "ID": "6c1debfa-7c26-4d8d-8bf0-0ccb1ec9f172",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$name",
              "Value": "Garage Monitor",
              "Props": {}
            },
            "$nodes": {
              "ID": "09b51723-c1b0-4154-94d2-a0a2d9efcb56",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$nodes",
              "Value": "Ambient,Presence,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "5e932417-831b-4e28-bc64-673c736fac90",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "d0c21e08-22cb-4d3a-b696-83875b949899",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$stats",
              "Value": "",
              "Props": {
                "interval": {
                  "ID": "e83926d5-7c83-4e30-ac21-b9ae80cbc55a",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "905",
                  "Props": {}
                },
                "signal": {
                  "ID": "75fbfc66-3d30-4cc7-bb7d-cb7e497908f0",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "54",
                  "Props": {}
                },
                "uptime": {
                  "ID": "cdd37489-6bb6-410d-af04-c97cfa848acd",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "593896",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "Ambient": {
              "ID": "b81bdbea-12a7-4404-bb17-4b0c883386ae",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "Ambient",
              "Attrs": {
                "$name": {
                  "ID": "38410dfc-27cc-45cb-8603-4f60c8e49c4b",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$name",
                  "Value": "Temperature and Humidity Sensor"
                },
                "$properties": {
                  "ID": "c1256e03-1410-4ffc-97dc-33c755c61ba6",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$properties",
                  "Value": "humidity,temperature"
                },
                "$type": {
                  "ID": "880d5d1c-4083-4d2b-bc61-6bd4c2236cec",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "Ambient": {
                  "ID": "6ca46cd3-e688-4c9c-818e-1bd67df94f21",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "81ca14bc-fd5d-4d83-812b-3276de7d4ea2",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "07c37f40-0dd9-447d-940d-ed4ca663f169",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Temperature and Humidity Sensor"
                    },
                    "$properties": {
                      "ID": "8bd91c34-3ebc-4b50-a161-8a95c7be9703",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "humidity,temperature"
                    },
                    "$type": {
                      "ID": "24dea291-205c-4029-877d-edf1cde2cb30",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "$unit": {
                      "ID": "7f245edd-f161-4143-9df6-5dc55bb98008",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": ""
                    },
                    "humidity": {
                      "ID": "da362aaf-b486-47b8-91fb-e1cf658dc157",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "humidity",
                      "Value": "57.18"
                    },
                    "temperature": {
                      "ID": "cd0b201f-a509-4d2f-a172-3fe3b6a8b6c9",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "temperature",
                      "Value": "26.18"
                    }
                  }
                },
                "humidity": {
                  "ID": "e65dd3e1-d3ec-4958-9b48-6c7243725535",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "humidity",
                  "Value": "57.18",
                  "Attrs": {
                    "$datatype": {
                      "ID": "6337c58e-960b-4852-9c62-221fd7f1aee3",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "297c103c-e5e5-4780-923e-410d6ed9cbb2",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$name",
                      "Value": "Humidity"
                    },
                    "$unit": {
                      "ID": "8bc8ed17-62ec-4bf9-bb67-5fc76c2b78b4",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": "%rH"
                    }
                  }
                },
                "temperature": {
                  "ID": "64ef1fdf-1d95-46d2-8431-11f319f3b10c",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "temperature",
                  "Value": "26.18",
                  "Attrs": {
                    "$datatype": {
                      "ID": "5c4f3c92-a846-46e8-9091-93fc9f987fd4",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "bcff30bd-5fc6-41b4-b295-5f0bdf0921cf",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$name",
                      "Value": "Temperature"
                    },
                    "$unit": {
                      "ID": "5dee2c4e-421b-4dbe-abd3-b56cdd7f73dd",
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
              "ID": "759bc87f-2cae-4933-b110-8cc3a45b784b",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "Presence",
              "Attrs": {
                "$name": {
                  "ID": "4fda92df-a620-48a7-9f1f-ad5337fccb75",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$name",
                  "Value": "Motion Sensor"
                },
                "$properties": {
                  "ID": "b801ce5e-9568-4815-a8b2-ce7d3ebe8b65",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$properties",
                  "Value": "motion"
                },
                "$type": {
                  "ID": "29388d18-a81c-4963-9883-8fb5019ee26f",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "Presence": {
                  "ID": "f3211ce3-7ead-4df9-b768-753b5cb3821c",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "9fbb77c6-c9ba-4562-af89-9134a639c9ca",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$format": {
                      "ID": "e819cf0a-df1b-4dbe-be7e-ff873a4bdab1",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "0cb19324-2782-4e47-892e-661da40e02c2",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Motion Sensor"
                    },
                    "$properties": {
                      "ID": "17f13245-c196-4b7a-89e9-d8adfc8e29db",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "motion"
                    },
                    "$type": {
                      "ID": "b72bb3f5-4868-42e9-acef-3845f50acc76",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "motion": {
                      "ID": "eeda3b4b-f8c2-4fb6-b0ef-d367a100fa1e",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "motion",
                      "Value": "OPEN"
                    }
                  }
                },
                "motion": {
                  "ID": "cf8b775d-9a39-422a-a026-b98f891ca8f1",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "motion",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "bd2fa666-b1b7-43a1-9371-31a094d6662f",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "ac624677-396e-4bb0-ab31-f2ebdb0a4503",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "0fad471d-3333-41fe-8519-19eccaf9ac06",
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
              "ID": "682ab211-890e-4fc9-bca0-e624f356dce1",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "hardware",
              "Attrs": {
                "$name": {
                  "ID": "a5330788-4c9c-41df-9313-bf33c5a12076",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$name",
                  "Value": "Device Info"
                },
                "$properties": {
                  "ID": "783de74f-7c0f-4f5b-9477-b545419743ae",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$properties",
                  "Value": "signal,mac,resetReason,voltage"
                },
                "$type": {
                  "ID": "ed505c57-e575-402f-922b-bbaaf9e0d1db",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "hardware": {
                  "ID": "6e718d43-da7c-4edd-990b-fc910f8f2be0",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "377f315e-54f0-41cb-ae73-f7714cc64278",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "631800af-7538-4241-8364-68bf93ca4e11",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Device Info"
                    },
                    "$properties": {
                      "ID": "ba785cb5-8dfe-42d8-8f11-79b9794dc584",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "signal,mac,resetReason,voltage"
                    },
                    "$type": {
                      "ID": "e1665044-0b86-478c-8d6b-3d32112b453b",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "$unit": {
                      "ID": "be025a49-02c3-4e22-a20a-ac4266eb22d3",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": ""
                    },
                    "mac": {
                      "ID": "21949b52-02dd-4598-821c-06d81f64d5b6",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "mac",
                      "Value": "B4:E6:2D:1B:5C:4D"
                    },
                    "resetReason": {
                      "ID": "72f54efe-a67b-47be-91d5-49f29335fe6d",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "resetReason",
                      "Value": "External System"
                    },
                    "signal": {
                      "ID": "87243b0f-08ca-4399-8022-b06dcad14626",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "signal",
                      "Value": "-73"
                    },
                    "voltage": {
                      "ID": "c001b1fb-b96c-4e84-a7cf-ba00fb8348b2",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "voltage",
                      "Value": "3.05"
                    }
                  }
                },
                "mac": {
                  "ID": "6e441555-f71a-4ca1-bec5-791f14f50288",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "mac",
                  "Value": "B4:E6:2D:1B:5C:4D",
                  "Attrs": {
                    "$datatype": {
                      "ID": "7e6c7a9e-2e72-4b8b-b934-05fbb2ebd92b",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": "sring"
                    },
                    "$name": {
                      "ID": "0cd6b4f8-9776-4047-b0ab-3575f436b2a8",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$name",
                      "Value": "mac"
                    }
                  }
                },
                "resetReason": {
                  "ID": "cb02134e-d5ad-412f-a4c6-2771c01b4f32",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "resetReason",
                  "Value": "External System",
                  "Attrs": {
                    "$datatype": {
                      "ID": "cc60a603-1343-4f94-ab40-7940418a9b90",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$name": {
                      "ID": "c0c40038-3361-4ec4-8cde-406ef7517b44",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$name",
                      "Value": "Last Reset Reason"
                    }
                  }
                },
                "signal": {
                  "ID": "a6719be7-9943-4f43-8685-bde986b2367f",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "signal",
                  "Value": "-73",
                  "Attrs": {
                    "$datatype": {
                      "ID": "8609876a-21d8-4d83-ad06-c608876f1496",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$name": {
                      "ID": "778c13ff-b8bc-42c3-ba44-a9dbd32be2c9",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$name",
                      "Value": "RSSI"
                    },
                    "$unit": {
                      "ID": "ead30875-ec13-4ea2-bdaa-7511cbea73eb",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": "dBm"
                    }
                  }
                },
                "voltage": {
                  "ID": "f98a546e-20b3-408f-9e42-858dd17ca7ff",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "voltage",
                  "Value": "3.05",
                  "Attrs": {
                    "$datatype": {
                      "ID": "58370d27-43c3-4c6d-867b-3c0ed9ce7216",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "f0080209-20e9-4c2e-b03f-0017645673b0",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$name",
                      "Value": "3.3V Supply"
                    },
                    "$unit": {
                      "ID": "fbd660dd-d204-4226-96bf-043a80db9e61",
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
        },
        "GarageProvider": {
          "ID": "71c6d9d39d15fe3565d22ab2ea689028",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "GarageProvider",
          "Attrs": {
            "$fw": {
              "ID": "0c4593d7-e8c1-47e1-9f51-d00d5c36b4c4",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "version": {
                  "ID": "2d5d150e-c8c8-4c18-9dbe-0df9073056bc",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "1.1.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "299399f6-906e-44ae-9b7d-2d6dc7f12d10",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "3d00e93c-dc5d-4cc8-9792-66e9b84149bf",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$implementation",
              "Value": "",
              "Props": {
                "version": {
                  "ID": "2d00a4aa-599d-462a-bfa7-d732b1242829",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "71877dc8-746f-4e3b-84f4-87d328ca6d21",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$localip",
              "Value": "10.100.1.218",
              "Props": {}
            },
            "$mac": {
              "ID": "659743a5-c67a-492d-87b3-17127b82b2ce",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$mac",
              "Value": "24:6F:28:97:63:B8",
              "Props": {}
            },
            "$name": {
              "ID": "d42d2c8a-f2f5-47c7-9702-a9a054259065",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$name",
              "Value": "Garage Provider",
              "Props": {}
            },
            "$nodes": {
              "ID": "fe470115-79d4-4fa4-b154-b8b9e9f2f1ed",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$nodes",
              "Value": "garageDoor,environmentMonitor",
              "Props": {}
            },
            "$state": {
              "ID": "43d10c2a-c966-4e88-9bc8-4b692812e906",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "1f89e345-bde9-4b11-a8ed-acbf2e1f3f28",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$stats",
              "Value": "",
              "Props": {
                "uptime": {
                  "ID": "90cd8012-1599-44c6-a51f-fe64704263d5",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "113713",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "environmentMonitor": {
              "ID": "c6357013-ff4a-4624-9116-c08864542133",
              "ElementType": 14,
              "Parent": "GarageProvider",
              "Name": "environmentMonitor",
              "Attrs": {},
              "Props": {
                "environmentMonitor": {
                  "ID": "2720a991-ee79-41b1-afda-cd68038bf4bc",
                  "ElementType": 16,
                  "Parent": "environmentMonitor",
                  "Name": "environmentMonitor",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "e3928058-341a-4813-9ee7-962ff5cdf1ec",
                      "ElementType": 17,
                      "Parent": "humdity",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$format": {
                      "ID": "8de58bbe-120f-4f3f-8bbc-58f0f217d308",
                      "ElementType": 17,
                      "Parent": "humdity",
                      "Name": "$format",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "364d7012-fff6-4300-9869-f46256047684",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Garage"
                    },
                    "$properties": {
                      "ID": "9abbd011-5d2f-4d47-8912-ef31cf750076",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "motion,temperature,humdity"
                    },
                    "$type": {
                      "ID": "80551f29-bea0-44fa-86a8-168da03751b5",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "SknSensors"
                    },
                    "$unit": {
                      "ID": "edf04e6d-a015-4a86-9537-503820b4aa14",
                      "ElementType": 17,
                      "Parent": "humdity",
                      "Name": "$unit",
                      "Value": ""
                    },
                    "humdity": {
                      "ID": "94b5958d-7d06-4df6-a12d-ebf44962a70d",
                      "ElementType": 17,
                      "Parent": "humdity",
                      "Name": "humdity",
                      "Value": "27.0"
                    },
                    "motion": {
                      "ID": "d202466c-c642-4ed4-b683-39fa82924e56",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "motion",
                      "Value": "ON"
                    },
                    "temperature": {
                      "ID": "9e1fd678-0d96-48b3-9470-1ebb75cb89c4",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "temperature",
                      "Value": "73.6"
                    }
                  }
                }
              }
            },
            "garageDoor": {
              "ID": "7196c3f0-e051-47ac-91fc-05d44f030e27",
              "ElementType": 14,
              "Parent": "GarageProvider",
              "Name": "garageDoor",
              "Attrs": {},
              "Props": {
                "garageDoor": {
                  "ID": "14d25a5c-17c8-40ed-b7dd-8a6926f4c03f",
                  "ElementType": 16,
                  "Parent": "garageDoor",
                  "Name": "garageDoor",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "b4639610-849a-4fc6-9f1f-65ec4e8c9581",
                      "ElementType": 17,
                      "Parent": "direction",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$format": {
                      "ID": "027ef38c-cb98-4387-8ba1-872c773cff99",
                      "ElementType": 17,
                      "Parent": "direction",
                      "Name": "$format",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "25291caa-0d33-4d1f-9107-f961a49e43d9",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Operations"
                    },
                    "$properties": {
                      "ID": "d39309cf-d142-4aa4-a2a1-1697555a3c88",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "operator,positon,direction"
                    },
                    "$settable": {
                      "ID": "d7d5df65-7aee-415d-af02-1f04b99e02ac",
                      "ElementType": 17,
                      "Parent": "operator",
                      "Name": "$settable",
                      "Value": ""
                    },
                    "$type": {
                      "ID": "11104911-2051-458a-9288-2f6c504fe2bf",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "SknSensors"
                    },
                    "$unit": {
                      "ID": "31221cea-03b8-4a4d-9363-896c39ff464b",
                      "ElementType": 17,
                      "Parent": "positon",
                      "Name": "$unit",
                      "Value": ""
                    },
                    "direction": {
                      "ID": "a4b6d357-7e00-426e-a0c8-d9305c4bde8a",
                      "ElementType": 17,
                      "Parent": "direction",
                      "Name": "direction",
                      "Value": "OPENING"
                    },
                    "positon": {
                      "ID": "c1544b64-b088-4ba1-97b2-46ac5e2bcc7f",
                      "ElementType": 17,
                      "Parent": "positon",
                      "Name": "positon",
                      "Value": "19"
                    }
                  }
                }
              }
            }
          }
        },
        "GuestRoom": {
          "ID": "bdab9513512a9d316939b5c0ac1ee099",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "GuestRoom",
          "Attrs": {
            "$fw": {
              "ID": "151b446e-4b9f-4800-9017-bb1946e971a3",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "df5241bc-7060-4325-aea4-301e07485d45",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "3ef8fbb48c5b23788a22a998c14a1a6d",
                  "Props": {}
                },
                "name": {
                  "ID": "885b64e9-15f6-48cc-a9ee-bfa52e9376ed",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "Monitor-DHT-RCWL-Metrics",
                  "Props": {}
                },
                "version": {
                  "ID": "8d27e30b-5753-4460-8c41-e68546d65d14",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "2.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "fb079a86-1bd2-48c7-9d47-e1ea867ec5ad",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "3287790c-03c9-4c59-a486-dd47f68cb8a8",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$implementation",
              "Value": "",
              "Props": {
                "config": {
                  "ID": "94730467-3d72-4508-b014-b076af114382",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"Guest Room\",\"device_id\":\"GuestRoom\",\"device_stats_interval\":180,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"sensorInterval\":900,\"motionHoldInterval\":60}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "f3cd85cf-4d61-405d-b315-a6636f36bf88",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "ota": {
                      "ID": "c8350db2-6aa1-406b-8765-f90fdcad4b32",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "36f4ea64-eb6c-4f81-baa5-89f8f84804f1",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "67ed302d-46de-43b4-a03a-13e200c903ff",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$localip",
              "Value": "10.100.1.178",
              "Props": {}
            },
            "$mac": {
              "ID": "cf365935-56eb-46b4-91e4-cb93ad0d12f2",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$mac",
              "Value": "BC:DD:C2:81:04:72",
              "Props": {}
            },
            "$name": {
              "ID": "ff262283-3c23-42cf-bf04-0831c6005fc6",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$name",
              "Value": "Guest Room",
              "Props": {}
            },
            "$nodes": {
              "ID": "89c900db-5ab8-48c7-9b6e-a1a1f2d6b6aa",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$nodes",
              "Value": "Ambient,Presence,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "9d1b79d3-139c-4aaf-a38a-315737893a21",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "14ae6303-b41d-48a1-b32c-bf84be7ac847",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$stats",
              "Value": "",
              "Props": {
                "interval": {
                  "ID": "48a76a13-92bc-4235-8c76-049bb832d8b7",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "185",
                  "Props": {}
                },
                "signal": {
                  "ID": "ca6a2879-a151-4071-995f-bdc372c4cc9f",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "72",
                  "Props": {}
                },
                "uptime": {
                  "ID": "0fc742d1-7d4a-4bee-91f0-010d1174001a",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "362358",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "Ambient": {
              "ID": "5b9e5ff5-1a10-4b2e-bff3-c901dcc80b36",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "Ambient",
              "Attrs": {
                "$name": {
                  "ID": "41441fa3-91ca-4c03-b493-a5628c4e4171",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$name",
                  "Value": "Temperature and Humidity Sensor"
                },
                "$properties": {
                  "ID": "676852a6-e97b-464b-ac59-fd3280663899",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$properties",
                  "Value": "humidity,temperature"
                },
                "$type": {
                  "ID": "cc37ce18-94a3-4a75-b6b7-ffc00cd5587c",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "Ambient": {
                  "ID": "2511c103-5d5f-4217-81fd-87bbe77d1613",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "26f206f2-89ab-4b81-a1ab-0587141ea388",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "8ca29bb5-b846-493f-9339-6375e2873a25",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Temperature and Humidity Sensor"
                    },
                    "$properties": {
                      "ID": "f7abf256-e36e-4be8-9113-30ffe865f9e0",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "humidity,temperature"
                    },
                    "$type": {
                      "ID": "ab2ca502-a922-4ee7-94f1-ade68c9a9381",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "$unit": {
                      "ID": "95f990ae-f96b-4f27-9457-bcfc3a8e8892",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": ""
                    },
                    "humidity": {
                      "ID": "1c089e64-5426-4317-8e96-2995667634c4",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "humidity",
                      "Value": "11.10"
                    },
                    "temperature": {
                      "ID": "6daf5380-7b05-4d94-8da2-1f4880ceb110",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "temperature",
                      "Value": "73.76"
                    }
                  }
                },
                "humidity": {
                  "ID": "c930e933-d76c-4229-a3e9-5ac3a335933a",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "humidity",
                  "Value": "11.10",
                  "Attrs": {
                    "$datatype": {
                      "ID": "2cdc110f-412b-4efd-9271-24013812ff21",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "0cfb2340-1b07-4082-9ed9-51e5e17a3521",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$name",
                      "Value": "Humidity"
                    },
                    "$unit": {
                      "ID": "233d1497-3c32-45b0-8109-c5fad7c5ff3d",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": "%rH"
                    }
                  }
                },
                "temperature": {
                  "ID": "d71302ab-1b9c-466b-a866-00c45cdd01a8",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "temperature",
                  "Value": "73.76",
                  "Attrs": {
                    "$datatype": {
                      "ID": "88c88759-07c8-4038-8f48-d9ab739894d0",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "675a01f6-cb22-4b48-af29-233639a73c80",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$name",
                      "Value": "Temperature"
                    },
                    "$unit": {
                      "ID": "7e7d5f4d-0faa-44c9-ae48-45c000201a0a",
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
              "ID": "8ce1308f-7b65-462f-a9eb-8b18c7785fde",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "Presence",
              "Attrs": {
                "$name": {
                  "ID": "5da87e62-3ae5-4871-8b6d-e346c285cba8",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$name",
                  "Value": "Motion Sensor"
                },
                "$properties": {
                  "ID": "4751ed78-ae6f-4da8-81bd-834b302e10c7",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$properties",
                  "Value": "motion"
                },
                "$type": {
                  "ID": "92d4ed53-6540-4247-9476-d57dc0673914",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "Presence": {
                  "ID": "e0c48e9d-f8c9-4196-a3ba-f2f5ab5d83f6",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "eafdfb08-ba67-4348-9c7a-ee6547c1cf07",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$format": {
                      "ID": "55a20332-6f27-4704-8463-c5879582703c",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "188c22a3-765a-486a-8be0-f9e6e47a4ad3",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Motion Sensor"
                    },
                    "$properties": {
                      "ID": "9072123d-af60-47c0-9e55-d131a9be2938",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "motion"
                    },
                    "$type": {
                      "ID": "bc7b650c-925b-4442-b30b-53f3fc288526",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "motion": {
                      "ID": "9e5c9cc8-92fb-4538-97aa-b43c15d8a777",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "motion",
                      "Value": "CLOSED"
                    }
                  }
                },
                "motion": {
                  "ID": "cccbbc90-3f8f-4c18-94d2-140078d55cc6",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "motion",
                  "Value": "CLOSED",
                  "Attrs": {
                    "$datatype": {
                      "ID": "2a842301-0fee-48d1-9bc1-60729f43a2be",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "27fdeb8e-2eb5-4fa2-817a-f2ec9b5e0496",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "c81e4f12-b872-4605-805e-fcae021f0646",
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
              "ID": "70d20307-488b-448c-9620-bb632837dc0f",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "hardware",
              "Attrs": {
                "$name": {
                  "ID": "a0349bd8-e8a4-4969-92a9-9c6822d3b007",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$name",
                  "Value": "Device Info"
                },
                "$properties": {
                  "ID": "ce6e71ff-de3f-4f9f-b4a5-cc12b97e7658",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$properties",
                  "Value": "signal,mac,resetReason,voltage"
                },
                "$type": {
                  "ID": "cd6442f4-f826-4a07-bb13-95f120c86fd6",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "hardware": {
                  "ID": "4728e110-8cc6-4e8b-ad01-e21dd43bc131",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "693fbca6-4ff6-4a10-9a93-48ebcc536649",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "5a49de53-2056-4158-8b41-9cb4248e0734",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Device Info"
                    },
                    "$properties": {
                      "ID": "35cac20f-7da1-46e7-99ab-75010534c38d",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "signal,mac,resetReason,voltage"
                    },
                    "$type": {
                      "ID": "6550ef76-27f0-4e51-bc0a-71e08f35685e",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "$unit": {
                      "ID": "debdd479-3d1d-49d8-a059-c229fd6e9c27",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": ""
                    },
                    "mac": {
                      "ID": "7cfb7197-89fb-4968-b15d-0944f92440db",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "mac",
                      "Value": "BC:DD:C2:81:04:72"
                    },
                    "resetReason": {
                      "ID": "6b4b908a-56db-4e2e-82ba-e41b56a77ffe",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "resetReason",
                      "Value": "External System"
                    },
                    "signal": {
                      "ID": "f277b14e-e794-4197-ba80-fa256f2dac65",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "signal",
                      "Value": "-64"
                    },
                    "voltage": {
                      "ID": "b9dad451-c472-44c7-982e-848220240040",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "voltage",
                      "Value": "2.98"
                    }
                  }
                },
                "mac": {
                  "ID": "229cfb03-e7f3-4ab1-b5aa-ebe38292d1b1",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "mac",
                  "Value": "BC:DD:C2:81:04:72",
                  "Attrs": {
                    "$datatype": {
                      "ID": "2dae1f64-a4f7-487a-95d3-e953698cb354",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": "sring"
                    },
                    "$name": {
                      "ID": "6289fc9d-c7d4-4cfa-a190-cc1dbd3f772a",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$name",
                      "Value": "mac"
                    }
                  }
                },
                "resetReason": {
                  "ID": "87caf14c-27fb-4633-981d-a55c292569b1",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "resetReason",
                  "Value": "External System",
                  "Attrs": {
                    "$datatype": {
                      "ID": "2bb976b9-1fc8-4d9c-89e6-8af9c793e63c",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$name": {
                      "ID": "ef6533f7-dcfb-44eb-956d-ff3b9c96b47b",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$name",
                      "Value": "Last Reset Reason"
                    }
                  }
                },
                "signal": {
                  "ID": "2e14a453-1d06-4e84-aaa6-f2e2bd93158b",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "signal",
                  "Value": "-64",
                  "Attrs": {
                    "$datatype": {
                      "ID": "c90e9cc3-c0eb-41fd-9abf-2cd8c69c7434",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$name": {
                      "ID": "cb952cdd-38d6-4cb9-ba5f-f54aaceeb387",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$name",
                      "Value": "RSSI"
                    },
                    "$unit": {
                      "ID": "36344740-d56d-4401-8d18-83ce32635d50",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": "dBm"
                    }
                  }
                },
                "voltage": {
                  "ID": "7f92fd24-a46c-497f-96c8-7edc6dee1d7b",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "voltage",
                  "Value": "2.98",
                  "Attrs": {
                    "$datatype": {
                      "ID": "2e98f35b-c154-4149-9c5f-01002793aae6",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "3f4e4823-3165-4310-ac3d-032d9f992326",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$name",
                      "Value": "3.3V Supply"
                    },
                    "$unit": {
                      "ID": "d65c9534-4b5d-41e3-9a37-18b453a37c45",
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
        },
        "HomeOffice": {
          "ID": "8c3705faef6bf02d8df1bb7b9516e182",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "HomeOffice",
          "Attrs": {
            "$fw": {
              "ID": "87d66026-5738-4064-8638-e2b6fac369bf",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "version": {
                  "ID": "af03d444-a8d6-4188-b333-68e6e92e7fe3",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "2.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "bb7630c4-dc40-462c-930b-0f5537fac292",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "a6854983-27a1-4a80-98e4-a870b7c82ee9",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$implementation",
              "Value": "",
              "Props": {
                "version": {
                  "ID": "4470c895-f0f5-4ceb-b46f-491e13154d1d",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "8d435843-1b04-4942-b7b5-f452d3914e25",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$localip",
              "Value": "10.100.1.161",
              "Props": {}
            },
            "$mac": {
              "ID": "7d9bf16d-04bf-4624-a441-0211cc6075ab",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$mac",
              "Value": "84:F3:EB:B7:55:D5",
              "Props": {}
            },
            "$name": {
              "ID": "b7fde697-5dad-4db9-864a-555d4dc74029",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$name",
              "Value": "Home Office",
              "Props": {}
            },
            "$nodes": {
              "ID": "f540b01c-492b-4513-a617-4191340bd8b3",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$nodes",
              "Value": "Ambient,Presence,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "7a2249cc-857b-4f72-a42f-9736076c3084",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "af0b9a3a-ed2f-44e2-af27-a2d319b701b6",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$stats",
              "Value": "",
              "Props": {
                "uptime": {
                  "ID": "78091a73-45a6-46d4-9760-d1b72bf0cae8",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "46625",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "Ambient": {
              "ID": "e37fd9b5-9f64-45fc-8d86-960cc279ea6c",
              "ElementType": 14,
              "Parent": "HomeOffice",
              "Name": "Ambient",
              "Attrs": {},
              "Props": {
                "Ambient": {
                  "ID": "292730de-a8d7-44ee-beea-c9be83661a63",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "b5c50643-4700-4daa-9ab5-7782702efa76",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "7b5cc725-beda-46cb-a67c-973a9f144aa4",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Temperature and Humidity Sensor"
                    },
                    "$properties": {
                      "ID": "b10e69f3-e2f1-40a3-8d85-efbdae736777",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "humidity,temperature"
                    },
                    "$type": {
                      "ID": "32b0f90a-5d64-431f-a5fd-4824df3a5bbe",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "$unit": {
                      "ID": "2fd721dc-2bd6-4d58-8980-3249ca8f0783",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": ""
                    },
                    "humidity": {
                      "ID": "d7e3fbbd-a266-4d2f-8844-d34a46924b0e",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "humidity",
                      "Value": "10.20"
                    },
                    "temperature": {
                      "ID": "b2254eaf-69e9-4a22-983a-410bedc4357e",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "temperature",
                      "Value": "75.20"
                    }
                  }
                }
              }
            },
            "Presence": {
              "ID": "03d9b898-9566-4489-8335-e55504113a9a",
              "ElementType": 14,
              "Parent": "HomeOffice",
              "Name": "Presence",
              "Attrs": {},
              "Props": {
                "Presence": {
                  "ID": "c13613c2-f73e-4f36-8de8-0e51578acde8",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "a4f13f0d-573d-487d-8b41-ed6394405990",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$format": {
                      "ID": "a1daf91c-81b6-4b18-a625-006d6680cea3",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "ee78175c-2577-4820-97af-290a726b7592",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Motion Sensor"
                    },
                    "$properties": {
                      "ID": "6ab60bb6-ca57-4281-9edd-c94ac287e40f",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "motion"
                    },
                    "$type": {
                      "ID": "cdab0e11-74c7-4b66-a885-1f41f95a24ed",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "motion": {
                      "ID": "00810f02-8d71-4c8c-8f16-c4e7bbe0e81e",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "motion",
                      "Value": "OPEN"
                    }
                  }
                }
              }
            },
            "hardware": {
              "ID": "69563621-84f8-4b15-8db6-7cb19ed13884",
              "ElementType": 14,
              "Parent": "HomeOffice",
              "Name": "hardware",
              "Attrs": {},
              "Props": {
                "hardware": {
                  "ID": "40b96962-2550-4eb5-b038-118cd7d625cf",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "42846b30-61a3-490e-b451-89ec7e759430",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "157af51a-d344-4b16-a2d3-e071f17edd38",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Device Info"
                    },
                    "$properties": {
                      "ID": "a4842abb-e5bb-487b-b004-a49e9547e0be",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "signal,mac,resetReason,voltage"
                    },
                    "$type": {
                      "ID": "ae7bbe3e-07b9-4f50-9a3f-7e2a7267e07e",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "$unit": {
                      "ID": "e86b77ca-f49f-49bf-bfe9-17c1a1a7593b",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": ""
                    },
                    "mac": {
                      "ID": "a7569d0d-9ef2-4754-bd73-739912810903",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "mac",
                      "Value": "84:F3:EB:B7:55:D5"
                    },
                    "resetReason": {
                      "ID": "f80660b5-a49a-45a2-869e-4a9c6585a870",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "resetReason",
                      "Value": "External System"
                    },
                    "signal": {
                      "ID": "feabec65-36c3-4b87-82f4-41c19799b353",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "signal",
                      "Value": "-65"
                    },
                    "voltage": {
                      "ID": "1ee5fcc7-cc26-43e8-b84c-89b0f22da1bd",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "voltage",
                      "Value": "2.98"
                    }
                  }
                }
              }
            }
          }
        },
        "MCPProvider": {
          "ID": "66e0038fbeb3c18a787483b03732a372",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "MCPProvider",
          "Attrs": {
            "$fw": {
              "ID": "84bd7c7d-533b-450c-b0d6-51be1264a934",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "version": {
                  "ID": "402c9682-2104-467c-8399-2717584ebc1e",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "1.4.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "562cfa21-1816-49de-a9d0-0a59e637a363",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "56db2005-dd1e-4cab-924b-6f44c4030bcd",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$implementation",
              "Value": "",
              "Props": {
                "version": {
                  "ID": "d77a5f8a-a310-45a4-b00f-b4e6119e9c63",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "d957d643-37f2-46a5-b0de-32fb8cfa2a18",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$localip",
              "Value": "10.100.1.222",
              "Props": {}
            },
            "$mac": {
              "ID": "071417a7-dd83-4f6d-b04c-a89e20afbf05",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$mac",
              "Value": "BC:DD:C2:24:B7:3C",
              "Props": {}
            },
            "$name": {
              "ID": "93c21dc8-8ee0-429d-9ce2-f33f7d4554ee",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$name",
              "Value": "Hard Wired Alarm Sensors",
              "Props": {}
            },
            "$nodes": {
              "ID": "1db99ff4-a09b-41c4-be91-b573e3afa48c",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$nodes",
              "Value": "wiredMonitor,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "5c02fb24-2c4f-43dc-97fd-91aac128fba5",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "55e35e4d-a02b-4f49-9f71-d69c94571fca",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$stats",
              "Value": "",
              "Props": {
                "uptime": {
                  "ID": "3237b1f5-68b7-4605-b4e7-05d51adce5ce",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "113410",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "hardware": {
              "ID": "5a292eb5-96e4-4754-9c3e-830cb8b0c310",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "hardware",
              "Attrs": {},
              "Props": {
                "hardware": {
                  "ID": "76f440d7-b10b-4f82-9d4b-e5374ed8df07",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "90f1112c-b630-4e91-a518-7f616fc75af4",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "31f6e63d-73de-4ab8-9e45-9052ae168e79",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Device Info"
                    },
                    "$properties": {
                      "ID": "d7696264-ef84-437e-8fe2-3f39e5e19f96",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "signal,mac,resetReason,voltage"
                    },
                    "$type": {
                      "ID": "acb407db-ae10-44a0-993e-7e069155274e",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "$unit": {
                      "ID": "37435f58-348b-4830-a5a4-b78f3afbc40a",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": ""
                    },
                    "mac": {
                      "ID": "e3e75c4e-a177-4939-8bae-8df38bd53d8f",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "mac",
                      "Value": "BC:DD:C2:24:B7:3C"
                    },
                    "resetReason": {
                      "ID": "dc398684-f0f9-420e-8a66-f82c5f685e91",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "resetReason",
                      "Value": "External System"
                    },
                    "signal": {
                      "ID": "1658a2d8-4a4f-41fd-90d1-f1c66f8279ed",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "signal",
                      "Value": "-70"
                    },
                    "voltage": {
                      "ID": "e8cd69a1-de91-46ff-ad63-7a262d273eee",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "voltage",
                      "Value": "2.99"
                    }
                  }
                }
              }
            },
            "wiredMonitor": {
              "ID": "94ef7a5c-ae0d-47d9-ab4f-d347859d31a7",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "wiredMonitor",
              "Attrs": {},
              "Props": {
                "wiredMonitor": {
                  "ID": "c3664187-822c-409f-8d1c-602d0278b248",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "wiredMonitor",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "9422ded4-8a74-4709-aae1-ecc4ba6ead26",
                      "ElementType": 17,
                      "Parent": "pin0",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$format": {
                      "ID": "078cf352-c203-4a5b-b96c-32e416b906c7",
                      "ElementType": 17,
                      "Parent": "pin0",
                      "Name": "$format",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "410a4c78-c922-405b-9786-deeaf210aeb8",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Wired Sensors"
                    },
                    "$properties": {
                      "ID": "172a0518-82e1-40a5-9044-c7a75fd3223d",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "pin0,pin1,pin2,pin3,pin4,pin5,pin6,pin7,pin8,pin9,pin10,pin11,pin12,pin13,pin14,pin15"
                    },
                    "$type": {
                      "ID": "a1468740-29f2-424d-9d4f-eef14a697f67",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "pin0": {
                      "ID": "ecc19900-3aa6-4973-8058-a36c105c2a29",
                      "ElementType": 17,
                      "Parent": "pin0",
                      "Name": "pin0",
                      "Value": "OPEN"
                    },
                    "pin1": {
                      "ID": "7849b842-1f63-45d0-91d6-007b9f0ae031",
                      "ElementType": 17,
                      "Parent": "pin1",
                      "Name": "pin1",
                      "Value": "OPEN"
                    },
                    "pin10": {
                      "ID": "0f3efe99-315b-4716-9b70-de0b9373fbab",
                      "ElementType": 17,
                      "Parent": "pin10",
                      "Name": "pin10",
                      "Value": "OPEN"
                    },
                    "pin11": {
                      "ID": "165ed15e-1cbf-4859-9938-8dfcc3f27dfd",
                      "ElementType": 17,
                      "Parent": "pin11",
                      "Name": "pin11",
                      "Value": "OPEN"
                    },
                    "pin12": {
                      "ID": "e9027853-8e09-48e3-99ca-a5f4183be935",
                      "ElementType": 17,
                      "Parent": "pin12",
                      "Name": "pin12",
                      "Value": "OPEN"
                    },
                    "pin13": {
                      "ID": "d43025ec-af83-4a0d-9bdd-9f9fb7c78a39",
                      "ElementType": 17,
                      "Parent": "pin13",
                      "Name": "pin13",
                      "Value": "CLOSED"
                    },
                    "pin14": {
                      "ID": "4650e42b-fc5b-4111-9f22-8f5510dde83c",
                      "ElementType": 17,
                      "Parent": "pin14",
                      "Name": "pin14",
                      "Value": "OPEN"
                    },
                    "pin15": {
                      "ID": "2c6b0b41-3c0a-4dbd-8615-2e7015e4d134",
                      "ElementType": 17,
                      "Parent": "pin15",
                      "Name": "pin15",
                      "Value": "OPEN"
                    },
                    "pin2": {
                      "ID": "eb9aa069-ae19-422d-a8c1-38ba52bdfc58",
                      "ElementType": 17,
                      "Parent": "pin2",
                      "Name": "pin2",
                      "Value": "OPEN"
                    },
                    "pin3": {
                      "ID": "1b0e9b9c-f11b-434f-9868-add59a805ffe",
                      "ElementType": 17,
                      "Parent": "pin3",
                      "Name": "pin3",
                      "Value": "OPEN"
                    },
                    "pin4": {
                      "ID": "0b6449d0-cecf-42c1-9b99-5ac00ca5af22",
                      "ElementType": 17,
                      "Parent": "pin4",
                      "Name": "pin4",
                      "Value": "OPEN"
                    },
                    "pin5": {
                      "ID": "11eda608-755d-4e8d-a8e5-ce9538b89d66",
                      "ElementType": 17,
                      "Parent": "pin5",
                      "Name": "pin5",
                      "Value": "OPEN"
                    },
                    "pin6": {
                      "ID": "b3b00c53-7890-4004-b40f-6ce5511cc8c4",
                      "ElementType": 17,
                      "Parent": "pin6",
                      "Name": "pin6",
                      "Value": "OPEN"
                    },
                    "pin7": {
                      "ID": "880fa241-b1c3-4fd8-915d-9aa8857c0bdb",
                      "ElementType": 17,
                      "Parent": "pin7",
                      "Name": "pin7",
                      "Value": "OPEN"
                    },
                    "pin8": {
                      "ID": "83165b39-4316-4f9e-af7f-a978206cd15d",
                      "ElementType": 17,
                      "Parent": "pin8",
                      "Name": "pin8",
                      "Value": "OPEN"
                    },
                    "pin9": {
                      "ID": "94f1f79a-8ae1-452f-a97b-a491cebd8e43",
                      "ElementType": 17,
                      "Parent": "pin9",
                      "Name": "pin9",
                      "Value": "OPEN"
                    }
                  }
                }
              }
            }
          }
        },
        "MediaRoom": {
          "ID": "d187e485b51c392d189d86bb61ce45a2",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "MediaRoom",
          "Attrs": {
            "$fw": {
              "ID": "610d127d-e976-40c8-a866-ab46706d091b",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "version": {
                  "ID": "f4dfd6e9-0964-4b28-8c40-b0aa823099e3",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "2.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "5f325f79-f053-4ee9-84d7-fbdb1c45c9e0",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "66100088-5e85-4857-a786-cc167c1e236d",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$implementation",
              "Value": "",
              "Props": {
                "version": {
                  "ID": "009f2066-3651-42bc-93a4-0bbd5504865f",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "79147309-1ff5-4c1b-8c1a-b21c43585153",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$localip",
              "Value": "10.100.1.180",
              "Props": {}
            },
            "$mac": {
              "ID": "3b3f1e3e-bfca-4551-8d83-2b964da19b52",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$mac",
              "Value": "B4:E6:2D:15:50:3A",
              "Props": {}
            },
            "$name": {
              "ID": "42b7d46f-6232-4e5a-b3ca-674374430bed",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$name",
              "Value": "Media Room",
              "Props": {}
            },
            "$nodes": {
              "ID": "dfe8bc80-7f92-45d6-a2b5-f9b9a52667b3",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$nodes",
              "Value": "Ambient,Presence,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "93cb06f5-be4c-4e4b-9d2f-1a5868de5201",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "0d84782b-6c6f-4540-a4cd-44e27b672c9f",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$stats",
              "Value": "",
              "Props": {
                "uptime": {
                  "ID": "40e99799-10f6-42cf-80cd-665758a0105d",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "102606",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "Ambient": {
              "ID": "b5f8edcd-6ea1-4e21-97f1-70c0cb160497",
              "ElementType": 14,
              "Parent": "MediaRoom",
              "Name": "Ambient",
              "Attrs": {},
              "Props": {
                "Ambient": {
                  "ID": "a9fb4f95-4a62-470c-bfca-aecba7fa628b",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "c60d29cd-321b-4d44-a28b-4bd4bfc25a00",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "c5641fe6-3a9f-401b-9cc6-f604ff8d06cd",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Temperature and Humidity Sensor"
                    },
                    "$properties": {
                      "ID": "6fb4402d-a07f-4c63-9f5f-47570a21c6c0",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "humidity,temperature"
                    },
                    "$type": {
                      "ID": "24aabef1-963e-41f4-b37f-5f56a0b69cb3",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "$unit": {
                      "ID": "5eda330e-aed6-49ea-a214-ec09a5082b2f",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": ""
                    },
                    "humidity": {
                      "ID": "680614bf-ea82-4677-9b18-06c602bc7caf",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "humidity",
                      "Value": "6.60"
                    },
                    "temperature": {
                      "ID": "189512ee-3fc5-45c6-889a-090f6d41e929",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "temperature",
                      "Value": "72.86"
                    }
                  }
                }
              }
            },
            "Presence": {
              "ID": "f494757f-d7c4-43dd-af35-5ea7be7bfc4e",
              "ElementType": 14,
              "Parent": "MediaRoom",
              "Name": "Presence",
              "Attrs": {},
              "Props": {
                "Presence": {
                  "ID": "1cd98e7d-dedc-49eb-9945-b677f20cd033",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "1689ac4f-83c9-417b-a12d-b3b3c35df0fb",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$format": {
                      "ID": "45ed03fd-64c2-4746-a27b-d92349d12da7",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "57485f69-8873-4cd8-a3cf-db2fff88135e",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Motion Sensor"
                    },
                    "$properties": {
                      "ID": "5d8b9d90-fff6-40e8-9fda-7844064a73e6",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "motion"
                    },
                    "$type": {
                      "ID": "cc6398c8-0c05-4cfc-9b30-d688914902ac",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "motion": {
                      "ID": "7ae5e75d-bc6c-4b81-9bc6-3362bbd21780",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "motion",
                      "Value": "OPEN"
                    }
                  }
                }
              }
            },
            "hardware": {
              "ID": "e4fccae9-c565-404b-8e8d-eec954436b07",
              "ElementType": 14,
              "Parent": "MediaRoom",
              "Name": "hardware",
              "Attrs": {},
              "Props": {
                "hardware": {
                  "ID": "29e85ca6-1538-432f-bc14-301f10d0ef29",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "a3f22e1d-c8e9-4c5d-9f1d-738866deb638",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "40fd4ac1-939b-4f09-be61-e3abc4e10a81",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Device Info"
                    },
                    "$properties": {
                      "ID": "75b47f52-da48-4302-9f86-c1df5a8d5bfa",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "signal,mac,resetReason,voltage"
                    },
                    "$type": {
                      "ID": "99d4b6e8-0dbc-4014-9dcc-f2a7504d7c63",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "$unit": {
                      "ID": "5496c18d-5e2c-4742-bd17-32a05b0623f5",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": ""
                    },
                    "mac": {
                      "ID": "35d4483c-a28b-451e-8ba0-3069fe45ef7f",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "mac",
                      "Value": "B4:E6:2D:15:50:3A"
                    },
                    "resetReason": {
                      "ID": "5df058fe-dc77-4045-bd38-d84858a8c6f8",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "resetReason",
                      "Value": "External System"
                    },
                    "signal": {
                      "ID": "036743e6-8b23-49f2-9213-6621e45aaec8",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "signal",
                      "Value": "-67"
                    },
                    "voltage": {
                      "ID": "dfe04fd2-c543-4f68-9eb6-ed1fc4852a05",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "voltage",
                      "Value": "3.12"
                    }
                  }
                }
              }
            }
          }
        },
        "OutsideMonitor": {
          "ID": "c6e49801761844b56e0a5604847e6214",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "OutsideMonitor",
          "Attrs": {
            "$fw": {
              "ID": "1c3351a0-bc55-4eea-8f3f-d23b874dd4ad",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "version": {
                  "ID": "b15ca46e-835d-43db-9266-5687e10ff063",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "2.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "f3f2a02f-9fa3-4100-a256-2cdea22a5012",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "c8adb82f-7dd9-4387-bf3b-323f33b34ed6",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$implementation",
              "Value": "",
              "Props": {
                "version": {
                  "ID": "949e4af3-b20b-484d-a47f-b980f6909f49",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "bab29332-3e2f-47f4-82f7-3a2c89663b24",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$localip",
              "Value": "10.100.1.182",
              "Props": {}
            },
            "$mac": {
              "ID": "7877f83d-3f26-4917-aaee-e90962bd6862",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$mac",
              "Value": "18:FE:34:FD:8C:1B",
              "Props": {}
            },
            "$name": {
              "ID": "e3f488f2-bebd-4390-805e-1474b1a30d63",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$name",
              "Value": "Outside Monitor",
              "Props": {}
            },
            "$nodes": {
              "ID": "21fd23ba-e689-47a1-b1ac-ebb62236430d",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$nodes",
              "Value": "Ambient,Presence,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "f9261a34-591d-4dcb-b9c2-0a9218cd09fb",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "a9941893-b0d9-4327-88bc-5e21de2e10a1",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$stats",
              "Value": "",
              "Props": {
                "uptime": {
                  "ID": "3bc620fc-9d06-44e3-9a1d-f1744405569e",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "396023",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "Ambient": {
              "ID": "a9635d75-9336-40a1-8694-fa830027fc11",
              "ElementType": 14,
              "Parent": "OutsideMonitor",
              "Name": "Ambient",
              "Attrs": {},
              "Props": {
                "Ambient": {
                  "ID": "4c04c5af-5df3-4699-ac85-5a5242e14df7",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "8fef1e3a-5a25-4d3d-9e62-e15b55af0d5c",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "738dd836-c119-47b4-8d9d-08f02188f49d",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Temperature and Humidity Sensor"
                    },
                    "$properties": {
                      "ID": "96e00030-d2b3-4132-a313-a3dbbbfee52b",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "humidity,temperature"
                    },
                    "$type": {
                      "ID": "809fa833-8a3b-48ff-8efb-1ab6a61ddd15",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "$unit": {
                      "ID": "12a9a517-51e5-4d40-a10f-a0c319d66fb1",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": ""
                    },
                    "humidity": {
                      "ID": "dcc247d9-d840-48a6-bfd3-2fdad39434d8",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "humidity",
                      "Value": "66.71"
                    },
                    "temperature": {
                      "ID": "d67e20d2-56de-4274-88f2-043d3f10939c",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "temperature",
                      "Value": "20.00"
                    }
                  }
                }
              }
            },
            "Presence": {
              "ID": "bcf18d6b-cf7d-464f-bb9f-4cc461653c80",
              "ElementType": 14,
              "Parent": "OutsideMonitor",
              "Name": "Presence",
              "Attrs": {},
              "Props": {
                "Presence": {
                  "ID": "d15ee33a-f422-46ab-afdb-061b7b53b648",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "a5e80896-0d3b-49a4-ade2-9a2e4933aad6",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$format": {
                      "ID": "4812bc04-ee6e-4050-b92f-27518f83381c",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "409787cc-c86e-426f-a6ed-5b6987fcb138",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Motion Sensor"
                    },
                    "$properties": {
                      "ID": "3deb5878-c8e9-4082-a529-23026b3033ce",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "motion"
                    },
                    "$type": {
                      "ID": "c51a713a-2a09-4c65-b8d0-8d609814b0f7",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "motion": {
                      "ID": "2ed6ba3d-8497-4595-b522-e148adebc638",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "motion",
                      "Value": "OPEN"
                    }
                  }
                }
              }
            },
            "hardware": {
              "ID": "cb765cf8-8c68-43c5-b98b-2e7b33ce3b12",
              "ElementType": 14,
              "Parent": "OutsideMonitor",
              "Name": "hardware",
              "Attrs": {},
              "Props": {
                "hardware": {
                  "ID": "bb83d2bb-1003-4ce0-9d49-1904aecf4262",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "576c26a1-d763-4366-9b66-d6f33b3410d7",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": ""
                    },
                    "$name": {
                      "ID": "dc343734-11fe-4742-b1f8-84ce58631635",
                      "ElementType": 17,
                      "Parent": "$name",
                      "Name": "$name",
                      "Value": "Device Info"
                    },
                    "$properties": {
                      "ID": "d7366c46-2ea7-465a-ae12-8313bbfc7916",
                      "ElementType": 17,
                      "Parent": "$properties",
                      "Name": "$properties",
                      "Value": "signal,mac,resetReason,voltage"
                    },
                    "$type": {
                      "ID": "2118dca0-ccb1-4365-8d82-6cf1eb4ba2b6",
                      "ElementType": 17,
                      "Parent": "$type",
                      "Name": "$type",
                      "Value": "sensor"
                    },
                    "$unit": {
                      "ID": "0ee138a3-ac58-462e-a16e-fae8d23d4e82",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": ""
                    },
                    "mac": {
                      "ID": "1ee3d0fd-f81f-47c4-ba9b-7527567cab04",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "mac",
                      "Value": "18:FE:34:FD:8C:1B"
                    },
                    "resetReason": {
                      "ID": "09d0c530-dd9e-4c74-8b49-51ce5d75d366",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "resetReason",
                      "Value": "External System"
                    },
                    "signal": {
                      "ID": "799b1e35-f3ff-4516-88c5-cf06004e1642",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "signal",
                      "Value": "-54"
                    },
                    "voltage": {
                      "ID": "cd5d4b0e-0020-4cbc-88e0-63c86335e82a",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "voltage",
                      "Value": "3.44"
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
```

### Console log from reload
```text
Configs.go init() called
Configs.go init() completed:   
event=calling handleCommandLineParams(), Config=demo-config, Environment=mqtt-config, debug=false 
level=info module=Homie-Service time:=2021-07-09T12:34:15.547073Z caller=service.go:141 pkg=services event="service starting"
level=info module=Homie-Service time:=2021-07-09T12:34:15.54724Z caller=service.go:185 method=doNetworkDiscovery event="demoRender() completed" limit=true
level=warn module=Homie-Service time:=2021-07-09T12:34:19.557664Z caller=usecase.go:141 pkg=deviceStorage Alert="LoadSchedules() Failed" error="no schedules available"
(0) X...... sknSensors 
(1) X/D...... sknSensors:D1R1MiniA 
(1)[1k] X/D/?...... Bucket[1]=D1R1MiniA, enum Key=$fw Value=[] 
(2)[a] X/D/A.... D1R1MiniA:$fw:checksum -->[] 
(3)[a] X/D/A/P.. D1R1MiniA:$fw:checksum:checksum -->[6867842dfce4674d4c724085467362c9]
(2)[a] X/D/A.... D1R1MiniA:$fw:name -->[] 
(3)[a] X/D/A/P.. D1R1MiniA:$fw:name:name -->[SknLiquids-128x32-oled]
(2)[a] X/D/A.... D1R1MiniA:$fw:version -->[] 
(3)[a] X/D/A/P.. D1R1MiniA:$fw:version:version -->[1.0.0]
(1)[1k] X/D/?...... Bucket[1]=D1R1MiniA, enum Key=$homie Value=[] 
(2)[a] X/D/A.... D1R1MiniA:$homie:$homie -->[3.0.1] 
(1)[1k] X/D/?...... Bucket[1]=D1R1MiniA, enum Key=$implementation Value=[] 
(2)[a] X/D/A.... D1R1MiniA:$implementation:$implementation -->[esp8266] 
(2)[a] X/D/A.... D1R1MiniA:$implementation:config -->[] 
(3)[a] X/D/A/P.. D1R1MiniA:$implementation:config:config -->[{"name":"Furnace","device_id":"D1R1MiniA","device_stats_interval":900,"wifi":{"ssid":"SFNSS1-24G"},"mqtt":{"host":"openhab.skoona.net","port":1883,"base_topic":"sknSensors/","auth":true},"ota":{"enabled":true}}]
(2)[a] X/D/A.... D1R1MiniA:$implementation:ota -->[] 
(3)[a] X/D/A/P.. D1R1MiniA:$implementation:ota:enabled -->[]
(4)[a] X/D/A/P/P D1R1MiniA:$implementation:ota:enabled:enabled -->[true] 
(2)[a] X/D/A.... D1R1MiniA:$implementation:version -->[] 
(3)[a] X/D/A/P.. D1R1MiniA:$implementation:version:version -->[3.0.0]
(1)[1k] X/D/?...... Bucket[1]=D1R1MiniA, enum Key=$localip Value=[] 
(2)[a] X/D/A.... D1R1MiniA:$localip:$localip -->[10.100.1.181] 
(1)[1k] X/D/?...... Bucket[1]=D1R1MiniA, enum Key=$mac Value=[] 
(2)[a] X/D/A.... D1R1MiniA:$mac:$mac -->[84:F3:EB:0C:38:6F] 
(1)[1k] X/D/?...... Bucket[1]=D1R1MiniA, enum Key=$name Value=[] 
(2)[a] X/D/A.... D1R1MiniA:$name:$name -->[Furnace] 
(1)[1k] X/D/?...... Bucket[1]=D1R1MiniA, enum Key=$nodes Value=[] 
(2)[a] X/D/A.... D1R1MiniA:$nodes:$nodes -->[Liquids] 
(1)[1k] X/D/?...... Bucket[1]=D1R1MiniA, enum Key=$state Value=[] 
(2)[a] X/D/A.... D1R1MiniA:$state:$state -->[ready] 
(1)[1k] X/D/?...... Bucket[1]=D1R1MiniA, enum Key=$stats Value=[] 
(2)[a] X/D/A.... D1R1MiniA:$stats:$stats -->[uptime] 
(2)[a] X/D/A.... D1R1MiniA:$stats:interval -->[] 
(3)[a] X/D/A/P.. D1R1MiniA:$stats:interval:interval -->[905]
(2)[a] X/D/A.... D1R1MiniA:$stats:signal -->[] 
(3)[a] X/D/A/P.. D1R1MiniA:$stats:signal:signal -->[88]
(2)[a] X/D/A.... D1R1MiniA:$stats:uptime -->[] 
(3)[a] X/D/A/P.. D1R1MiniA:$stats:uptime:uptime -->[114307]
(1)[1k] X/D/?...... Bucket[1]=D1R1MiniA, enum Key=Liquids Value=[] 
(2)[n] X/D/N.... D1R1MiniA:Liquids:$name -->[] 
(3)[n] X/D/N/A|P.. D1R1MiniA:Liquids:$name:$name -->[Water Sensor]
(2)[n] X/D/N.... D1R1MiniA:Liquids:$properties -->[] 
(3)[n] X/D/N/A|P.. D1R1MiniA:Liquids:$properties:$properties -->[level,volts]
(2)[n] X/D/N.... D1R1MiniA:Liquids:$type -->[] 
(3)[n] X/D/N/A|P.. D1R1MiniA:Liquids:$type:$type -->[sensor]
(2)[n] X/D/N.... D1R1MiniA:Liquids:level -->[] 
(3)[n] X/D/N/A|P.. D1R1MiniA:Liquids:level:$datatype -->[]
(4)[n] X/D/N/P/A D1R1MiniA:Liquids:level:$datatype:$datatype -->[integer] 
(3)[n] X/D/N/A|P.. D1R1MiniA:Liquids:level:$format -->[]
(4)[n] X/D/N/P/A D1R1MiniA:Liquids:level:$format:$format -->[0:1000] 
(3)[n] X/D/N/A|P.. D1R1MiniA:Liquids:level:$name -->[]
(4)[n] X/D/N/P/A D1R1MiniA:Liquids:level:$name:$name -->[Level] 
(3)[n] X/D/N/A|P.. D1R1MiniA:Liquids:level:$unit -->[]
(4)[n] X/D/N/P/A D1R1MiniA:Liquids:level:$unit:$unit -->[#] 
(3)[n] X/D/N/A|P.. D1R1MiniA:Liquids:level:level -->[7]
(2)[n] X/D/N.... D1R1MiniA:Liquids:volts -->[] 
(3)[n] X/D/N/A|P.. D1R1MiniA:Liquids:volts:$datatype -->[]
(4)[n] X/D/N/P/A D1R1MiniA:Liquids:volts:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. D1R1MiniA:Liquids:volts:$format -->[]
(4)[n] X/D/N/P/A D1R1MiniA:Liquids:volts:$format:$format -->[0:4] 
(3)[n] X/D/N/A|P.. D1R1MiniA:Liquids:volts:$name -->[]
(4)[n] X/D/N/P/A D1R1MiniA:Liquids:volts:$name:$name -->[Volts] 
(3)[n] X/D/N/A|P.. D1R1MiniA:Liquids:volts:$unit -->[]
(4)[n] X/D/N/P/A D1R1MiniA:Liquids:volts:$unit:$unit -->[V] 
(3)[n] X/D/N/A|P.. D1R1MiniA:Liquids:volts:volts -->[0.0]
(0) X...... sknSensors 
(1) X/D...... sknSensors:DSMonitor 
(1)[1k] X/D/?...... Bucket[1]=DSMonitor, enum Key=$fw Value=[] 
(2)[a] X/D/A.... DSMonitor:$fw:checksum -->[] 
(3)[a] X/D/A/P.. DSMonitor:$fw:checksum:checksum -->[93341625d6fa5d9533e6e315fbdde634]
(2)[a] X/D/A.... DSMonitor:$fw:name -->[] 
(3)[a] X/D/A/P.. DSMonitor:$fw:name:name -->[Environment-DS18B20]
(2)[a] X/D/A.... DSMonitor:$fw:version -->[] 
(3)[a] X/D/A/P.. DSMonitor:$fw:version:version -->[1.0.2]
(1)[1k] X/D/?...... Bucket[1]=DSMonitor, enum Key=$homie Value=[] 
(2)[a] X/D/A.... DSMonitor:$homie:$homie -->[3.0.1] 
(1)[1k] X/D/?...... Bucket[1]=DSMonitor, enum Key=$implementation Value=[] 
(2)[a] X/D/A.... DSMonitor:$implementation:$implementation -->[esp8266] 
(2)[a] X/D/A.... DSMonitor:$implementation:config -->[] 
(3)[a] X/D/A/P.. DSMonitor:$implementation:config:config -->[{"name":"Environment-DS18B20","device_id":"DSMonitor","device_stats_interval":900,"wifi":{"ssid":"SFNSS1-24G"},"mqtt":{"host":"openhab.skoona.net","port":1883,"base_topic":"sknSensors/","auth":true},"ota":{"enabled":true},"settings":{"sensorsInterval":900}}]
(2)[a] X/D/A.... DSMonitor:$implementation:ota -->[] 
(3)[a] X/D/A/P.. DSMonitor:$implementation:ota:enabled -->[]
(4)[a] X/D/A/P/P DSMonitor:$implementation:ota:enabled:enabled -->[true] 
(2)[a] X/D/A.... DSMonitor:$implementation:version -->[] 
(3)[a] X/D/A/P.. DSMonitor:$implementation:version:version -->[3.0.0]
(1)[1k] X/D/?...... Bucket[1]=DSMonitor, enum Key=$localip Value=[] 
(2)[a] X/D/A.... DSMonitor:$localip:$localip -->[10.100.1.164] 
(1)[1k] X/D/?...... Bucket[1]=DSMonitor, enum Key=$mac Value=[] 
(2)[a] X/D/A.... DSMonitor:$mac:$mac -->[DC:4F:22:3B:33:4F] 
(1)[1k] X/D/?...... Bucket[1]=DSMonitor, enum Key=$name Value=[] 
(2)[a] X/D/A.... DSMonitor:$name:$name -->[Environment-DS18B20] 
(1)[1k] X/D/?...... Bucket[1]=DSMonitor, enum Key=$nodes Value=[] 
(2)[a] X/D/A.... DSMonitor:$nodes:$nodes -->[Ambient[]] 
(1)[1k] X/D/?...... Bucket[1]=DSMonitor, enum Key=$state Value=[] 
(2)[a] X/D/A.... DSMonitor:$state:$state -->[ready] 
(1)[1k] X/D/?...... Bucket[1]=DSMonitor, enum Key=$stats Value=[] 
(2)[a] X/D/A.... DSMonitor:$stats:$stats -->[uptime] 
(2)[a] X/D/A.... DSMonitor:$stats:interval -->[] 
(3)[a] X/D/A/P.. DSMonitor:$stats:interval:interval -->[905]
(2)[a] X/D/A.... DSMonitor:$stats:signal -->[] 
(3)[a] X/D/A/P.. DSMonitor:$stats:signal:signal -->[80]
(2)[a] X/D/A.... DSMonitor:$stats:uptime -->[] 
(3)[a] X/D/A/P.. DSMonitor:$stats:uptime:uptime -->[5]
(1)[1k] X/D/?...... Bucket[1]=DSMonitor, enum Key=Ambient Value=[] 
(2)[n] X/D/N.... DSMonitor:Ambient:$array -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient:$array:$array -->[0-3]
(2)[n] X/D/N.... DSMonitor:Ambient:$name -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient:$name:$name -->[Environment Monitor]
(2)[n] X/D/N.... DSMonitor:Ambient:$properties -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient:$properties:$properties -->[state,temperature]
(2)[n] X/D/N.... DSMonitor:Ambient:$type -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient:$type:$type -->[sensor]
(2)[n] X/D/N.... DSMonitor:Ambient:state -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient:state:$datatype -->[]
(4)[n] X/D/N/P/A DSMonitor:Ambient:state:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient:state:$format -->[]
(4)[n] X/D/N/P/A DSMonitor:Ambient:state:$format:$format -->[OK,Error,InvalidAddress] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient:state:$name -->[]
(4)[n] X/D/N/P/A DSMonitor:Ambient:state:$name:$name -->[State] 
(2)[n] X/D/N.... DSMonitor:Ambient:temperature -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient:temperature:$datatype -->[]
(4)[n] X/D/N/P/A DSMonitor:Ambient:temperature:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient:temperature:$name -->[]
(4)[n] X/D/N/P/A DSMonitor:Ambient:temperature:$name:$name -->[Temperature] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient:temperature:$unit -->[]
(4)[n] X/D/N/P/A DSMonitor:Ambient:temperature:$unit:$unit -->[°F] 
(1)[1k] X/D/?...... Bucket[1]=DSMonitor, enum Key=Ambient_0 Value=[] 
(2)[n] X/D/N.... DSMonitor:Ambient_0:$name -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient_0:$name:$name -->[Ambient_0]
(2)[n] X/D/N.... DSMonitor:Ambient_0:state -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient_0:state:state -->[OK]
(2)[n] X/D/N.... DSMonitor:Ambient_0:temperature -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient_0:temperature:temperature -->[73.06]
(1)[1k] X/D/?...... Bucket[1]=DSMonitor, enum Key=Ambient_1 Value=[] 
(2)[n] X/D/N.... DSMonitor:Ambient_1:$name -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient_1:$name:$name -->[Ambient_1]
(2)[n] X/D/N.... DSMonitor:Ambient_1:state -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient_1:state:state -->[OK]
(2)[n] X/D/N.... DSMonitor:Ambient_1:temperature -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient_1:temperature:temperature -->[73.96]
(1)[1k] X/D/?...... Bucket[1]=DSMonitor, enum Key=Ambient_2 Value=[] 
(2)[n] X/D/N.... DSMonitor:Ambient_2:$name -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient_2:$name:$name -->[Ambient_2]
(2)[n] X/D/N.... DSMonitor:Ambient_2:state -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient_2:state:state -->[OK]
(2)[n] X/D/N.... DSMonitor:Ambient_2:temperature -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient_2:temperature:temperature -->[73.63]
(1)[1k] X/D/?...... Bucket[1]=DSMonitor, enum Key=Ambient_3 Value=[] 
(2)[n] X/D/N.... DSMonitor:Ambient_3:$name -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient_3:$name:$name -->[Ambient_3]
(2)[n] X/D/N.... DSMonitor:Ambient_3:state -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient_3:state:state -->[OK]
(2)[n] X/D/N.... DSMonitor:Ambient_3:temperature -->[] 
(3)[n] X/D/N/A|P.. DSMonitor:Ambient_3:temperature:temperature -->[72.39]
(0) X...... sknSensors 
(1) X/D...... sknSensors:FamilyRoom 
(1)[1k] X/D/?...... Bucket[1]=FamilyRoom, enum Key=$fw Value=[] 
(2)[a] X/D/A.... FamilyRoom:$fw:checksum -->[] 
(3)[a] X/D/A/P.. FamilyRoom:$fw:checksum:checksum -->[3ef8fbb48c5b23788a22a998c14a1a6d]
(2)[a] X/D/A.... FamilyRoom:$fw:name -->[] 
(3)[a] X/D/A/P.. FamilyRoom:$fw:name:name -->[Monitor-DHT-RCWL-Metrics]
(2)[a] X/D/A.... FamilyRoom:$fw:version -->[] 
(3)[a] X/D/A/P.. FamilyRoom:$fw:version:version -->[2.0.0]
(1)[1k] X/D/?...... Bucket[1]=FamilyRoom, enum Key=$homie Value=[] 
(2)[a] X/D/A.... FamilyRoom:$homie:$homie -->[3.0.1] 
(1)[1k] X/D/?...... Bucket[1]=FamilyRoom, enum Key=$implementation Value=[] 
(2)[a] X/D/A.... FamilyRoom:$implementation:$implementation -->[esp8266] 
(2)[a] X/D/A.... FamilyRoom:$implementation:config -->[] 
(3)[a] X/D/A/P.. FamilyRoom:$implementation:config:config -->[{"name":"FamilyRoom","device_id":"FamilyRoom","device_stats_interval":180,"wifi":{"ssid":"SFNSS1-24G"},"mqtt":{"host":"openhab.skoona.net","port":1883,"base_topic":"sknSensors/","auth":true},"ota":{"enabled":true},"settings":{"sensorInterval":900,"motionHoldInterval":60}}]
(2)[a] X/D/A.... FamilyRoom:$implementation:ota -->[] 
(3)[a] X/D/A/P.. FamilyRoom:$implementation:ota:enabled -->[]
(4)[a] X/D/A/P/P FamilyRoom:$implementation:ota:enabled:enabled -->[true] 
(2)[a] X/D/A.... FamilyRoom:$implementation:version -->[] 
(3)[a] X/D/A/P.. FamilyRoom:$implementation:version:version -->[3.0.0]
(1)[1k] X/D/?...... Bucket[1]=FamilyRoom, enum Key=$localip Value=[] 
(2)[a] X/D/A.... FamilyRoom:$localip:$localip -->[10.100.1.166] 
(1)[1k] X/D/?...... Bucket[1]=FamilyRoom, enum Key=$mac Value=[] 
(2)[a] X/D/A.... FamilyRoom:$mac:$mac -->[BC:DD:C2:E5:C4:24] 
(1)[1k] X/D/?...... Bucket[1]=FamilyRoom, enum Key=$name Value=[] 
(2)[a] X/D/A.... FamilyRoom:$name:$name -->[FamilyRoom] 
(1)[1k] X/D/?...... Bucket[1]=FamilyRoom, enum Key=$nodes Value=[] 
(2)[a] X/D/A.... FamilyRoom:$nodes:$nodes -->[Ambient,Presence,hardware] 
(1)[1k] X/D/?...... Bucket[1]=FamilyRoom, enum Key=$state Value=[] 
(2)[a] X/D/A.... FamilyRoom:$state:$state -->[ready] 
(1)[1k] X/D/?...... Bucket[1]=FamilyRoom, enum Key=$stats Value=[] 
(2)[a] X/D/A.... FamilyRoom:$stats:$stats -->[uptime] 
(2)[a] X/D/A.... FamilyRoom:$stats:interval -->[] 
(3)[a] X/D/A/P.. FamilyRoom:$stats:interval:interval -->[185]
(2)[a] X/D/A.... FamilyRoom:$stats:signal -->[] 
(3)[a] X/D/A/P.. FamilyRoom:$stats:signal:signal -->[76]
(2)[a] X/D/A.... FamilyRoom:$stats:uptime -->[] 
(3)[a] X/D/A/P.. FamilyRoom:$stats:uptime:uptime -->[423819]
(1)[1k] X/D/?...... Bucket[1]=FamilyRoom, enum Key=Ambient Value=[] 
(2)[n] X/D/N.... FamilyRoom:Ambient:$name -->[] 
(3)[n] X/D/N/A|P.. FamilyRoom:Ambient:$name:$name -->[Temperature and Humidity Sensor]
(2)[n] X/D/N.... FamilyRoom:Ambient:$properties -->[] 
(3)[n] X/D/N/A|P.. FamilyRoom:Ambient:$properties:$properties -->[humidity,temperature]
(2)[n] X/D/N.... FamilyRoom:Ambient:$type -->[] 
(3)[n] X/D/N/A|P.. FamilyRoom:Ambient:$type:$type -->[sensor]
(2)[n] X/D/N.... FamilyRoom:Ambient:humidity -->[] 
(3)[n] X/D/N/A|P.. FamilyRoom:Ambient:humidity:$datatype -->[]
(4)[n] X/D/N/P/A FamilyRoom:Ambient:humidity:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. FamilyRoom:Ambient:humidity:$name -->[]
(4)[n] X/D/N/P/A FamilyRoom:Ambient:humidity:$name:$name -->[Humidity] 
(3)[n] X/D/N/A|P.. FamilyRoom:Ambient:humidity:$unit -->[]
(4)[n] X/D/N/P/A FamilyRoom:Ambient:humidity:$unit:$unit -->[%rH] 
(3)[n] X/D/N/A|P.. FamilyRoom:Ambient:humidity:humidity -->[20.50]
(2)[n] X/D/N.... FamilyRoom:Ambient:temperature -->[] 
(3)[n] X/D/N/A|P.. FamilyRoom:Ambient:temperature:$datatype -->[]
(4)[n] X/D/N/P/A FamilyRoom:Ambient:temperature:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. FamilyRoom:Ambient:temperature:$name -->[]
(4)[n] X/D/N/P/A FamilyRoom:Ambient:temperature:$name:$name -->[Temperature] 
(3)[n] X/D/N/A|P.. FamilyRoom:Ambient:temperature:$unit -->[]
(4)[n] X/D/N/P/A FamilyRoom:Ambient:temperature:$unit:$unit -->[°F] 
(3)[n] X/D/N/A|P.. FamilyRoom:Ambient:temperature:temperature -->[67.64]
(1)[1k] X/D/?...... Bucket[1]=FamilyRoom, enum Key=Presence Value=[] 
(2)[n] X/D/N.... FamilyRoom:Presence:$name -->[] 
(3)[n] X/D/N/A|P.. FamilyRoom:Presence:$name:$name -->[Motion Sensor]
(2)[n] X/D/N.... FamilyRoom:Presence:$properties -->[] 
(3)[n] X/D/N/A|P.. FamilyRoom:Presence:$properties:$properties -->[motion]
(2)[n] X/D/N.... FamilyRoom:Presence:$type -->[] 
(3)[n] X/D/N/A|P.. FamilyRoom:Presence:$type:$type -->[sensor]
(2)[n] X/D/N.... FamilyRoom:Presence:motion -->[] 
(3)[n] X/D/N/A|P.. FamilyRoom:Presence:motion:$datatype -->[]
(4)[n] X/D/N/P/A FamilyRoom:Presence:motion:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. FamilyRoom:Presence:motion:$format -->[]
(4)[n] X/D/N/P/A FamilyRoom:Presence:motion:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. FamilyRoom:Presence:motion:$name -->[]
(4)[n] X/D/N/P/A FamilyRoom:Presence:motion:$name:$name -->[Motion] 
(3)[n] X/D/N/A|P.. FamilyRoom:Presence:motion:motion -->[OPEN]
(1)[1k] X/D/?...... Bucket[1]=FamilyRoom, enum Key=hardware Value=[] 
(2)[n] X/D/N.... FamilyRoom:hardware:$name -->[] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:$name:$name -->[Device Info]
(2)[n] X/D/N.... FamilyRoom:hardware:$properties -->[] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:$properties:$properties -->[signal,mac,resetReason,voltage]
(2)[n] X/D/N.... FamilyRoom:hardware:$type -->[] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:$type:$type -->[sensor]
(2)[n] X/D/N.... FamilyRoom:hardware:mac -->[] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:mac:$datatype -->[]
(4)[n] X/D/N/P/A FamilyRoom:hardware:mac:$datatype:$datatype -->[sring] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:mac:$name -->[]
(4)[n] X/D/N/P/A FamilyRoom:hardware:mac:$name:$name -->[mac] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:mac:mac -->[BC:DD:C2:E5:C4:24]
(2)[n] X/D/N.... FamilyRoom:hardware:resetReason -->[] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:resetReason:$datatype -->[]
(4)[n] X/D/N/P/A FamilyRoom:hardware:resetReason:$datatype:$datatype -->[string] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:resetReason:$name -->[]
(4)[n] X/D/N/P/A FamilyRoom:hardware:resetReason:$name:$name -->[Last Reset Reason] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:resetReason:resetReason -->[External System]
(2)[n] X/D/N.... FamilyRoom:hardware:signal -->[] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:signal:$datatype -->[]
(4)[n] X/D/N/P/A FamilyRoom:hardware:signal:$datatype:$datatype -->[integer] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:signal:$name -->[]
(4)[n] X/D/N/P/A FamilyRoom:hardware:signal:$name:$name -->[RSSI] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:signal:$unit -->[]
(4)[n] X/D/N/P/A FamilyRoom:hardware:signal:$unit:$unit -->[dBm] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:signal:signal -->[-62]
(2)[n] X/D/N.... FamilyRoom:hardware:voltage -->[] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:voltage:$datatype -->[]
(4)[n] X/D/N/P/A FamilyRoom:hardware:voltage:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:voltage:$name -->[]
(4)[n] X/D/N/P/A FamilyRoom:hardware:voltage:$name:$name -->[3.3V Supply] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:voltage:$unit -->[]
(4)[n] X/D/N/P/A FamilyRoom:hardware:voltage:$unit:$unit -->[V] 
(3)[n] X/D/N/A|P.. FamilyRoom:hardware:voltage:voltage -->[3.03]
(0) X...... sknSensors 
(1) X/D...... sknSensors:GarageMonitor 
(1)[1k] X/D/?...... Bucket[1]=GarageMonitor, enum Key=$fw Value=[] 
(2)[a] X/D/A.... GarageMonitor:$fw:checksum -->[] 
(3)[a] X/D/A/P.. GarageMonitor:$fw:checksum:checksum -->[615fed382ab44bd43fe83508aecac682]
(2)[a] X/D/A.... GarageMonitor:$fw:name -->[] 
(3)[a] X/D/A/P.. GarageMonitor:$fw:name:name -->[Monitor-SHT31-RCWL-Metrics]
(2)[a] X/D/A.... GarageMonitor:$fw:version -->[] 
(3)[a] X/D/A/P.. GarageMonitor:$fw:version:version -->[2.0.0]
(1)[1k] X/D/?...... Bucket[1]=GarageMonitor, enum Key=$homie Value=[] 
(2)[a] X/D/A.... GarageMonitor:$homie:$homie -->[3.0.1] 
(1)[1k] X/D/?...... Bucket[1]=GarageMonitor, enum Key=$implementation Value=[] 
(2)[a] X/D/A.... GarageMonitor:$implementation:$implementation -->[esp8266] 
(2)[a] X/D/A.... GarageMonitor:$implementation:config -->[] 
(3)[a] X/D/A/P.. GarageMonitor:$implementation:config:config -->[{"name":"Garage Monitor","device_id":"GarageMonitor","device_stats_interval":900,"wifi":{"ssid":"SFNSS1-24G"},"mqtt":{"host":"openhab.skoona.net","port":1883,"base_topic":"sknSensors/","auth":true},"ota":{"enabled":true},"settings":{"sensorInterval":900,"motionHoldInterval":60}}]
(2)[a] X/D/A.... GarageMonitor:$implementation:ota -->[] 
(3)[a] X/D/A/P.. GarageMonitor:$implementation:ota:enabled -->[]
(4)[a] X/D/A/P/P GarageMonitor:$implementation:ota:enabled:enabled -->[true] 
(2)[a] X/D/A.... GarageMonitor:$implementation:version -->[] 
(3)[a] X/D/A/P.. GarageMonitor:$implementation:version:version -->[3.0.0]
(1)[1k] X/D/?...... Bucket[1]=GarageMonitor, enum Key=$localip Value=[] 
(2)[a] X/D/A.... GarageMonitor:$localip:$localip -->[10.100.1.177] 
(1)[1k] X/D/?...... Bucket[1]=GarageMonitor, enum Key=$mac Value=[] 
(2)[a] X/D/A.... GarageMonitor:$mac:$mac -->[B4:E6:2D:1B:5C:4D] 
(1)[1k] X/D/?...... Bucket[1]=GarageMonitor, enum Key=$name Value=[] 
(2)[a] X/D/A.... GarageMonitor:$name:$name -->[Garage Monitor] 
(1)[1k] X/D/?...... Bucket[1]=GarageMonitor, enum Key=$nodes Value=[] 
(2)[a] X/D/A.... GarageMonitor:$nodes:$nodes -->[Ambient,Presence,hardware] 
(1)[1k] X/D/?...... Bucket[1]=GarageMonitor, enum Key=$state Value=[] 
(2)[a] X/D/A.... GarageMonitor:$state:$state -->[ready] 
(1)[1k] X/D/?...... Bucket[1]=GarageMonitor, enum Key=$stats Value=[] 
(2)[a] X/D/A.... GarageMonitor:$stats:$stats -->[uptime] 
(2)[a] X/D/A.... GarageMonitor:$stats:interval -->[] 
(3)[a] X/D/A/P.. GarageMonitor:$stats:interval:interval -->[905]
(2)[a] X/D/A.... GarageMonitor:$stats:signal -->[] 
(3)[a] X/D/A/P.. GarageMonitor:$stats:signal:signal -->[54]
(2)[a] X/D/A.... GarageMonitor:$stats:uptime -->[] 
(3)[a] X/D/A/P.. GarageMonitor:$stats:uptime:uptime -->[593896]
(1)[1k] X/D/?...... Bucket[1]=GarageMonitor, enum Key=Ambient Value=[] 
(2)[n] X/D/N.... GarageMonitor:Ambient:$name -->[] 
(3)[n] X/D/N/A|P.. GarageMonitor:Ambient:$name:$name -->[Temperature and Humidity Sensor]
(2)[n] X/D/N.... GarageMonitor:Ambient:$properties -->[] 
(3)[n] X/D/N/A|P.. GarageMonitor:Ambient:$properties:$properties -->[humidity,temperature]
(2)[n] X/D/N.... GarageMonitor:Ambient:$type -->[] 
(3)[n] X/D/N/A|P.. GarageMonitor:Ambient:$type:$type -->[sensor]
(2)[n] X/D/N.... GarageMonitor:Ambient:humidity -->[] 
(3)[n] X/D/N/A|P.. GarageMonitor:Ambient:humidity:$datatype -->[]
(4)[n] X/D/N/P/A GarageMonitor:Ambient:humidity:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. GarageMonitor:Ambient:humidity:$name -->[]
(4)[n] X/D/N/P/A GarageMonitor:Ambient:humidity:$name:$name -->[Humidity] 
(3)[n] X/D/N/A|P.. GarageMonitor:Ambient:humidity:$unit -->[]
(4)[n] X/D/N/P/A GarageMonitor:Ambient:humidity:$unit:$unit -->[%rH] 
(3)[n] X/D/N/A|P.. GarageMonitor:Ambient:humidity:humidity -->[57.18]
(2)[n] X/D/N.... GarageMonitor:Ambient:temperature -->[] 
(3)[n] X/D/N/A|P.. GarageMonitor:Ambient:temperature:$datatype -->[]
(4)[n] X/D/N/P/A GarageMonitor:Ambient:temperature:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. GarageMonitor:Ambient:temperature:$name -->[]
(4)[n] X/D/N/P/A GarageMonitor:Ambient:temperature:$name:$name -->[Temperature] 
(3)[n] X/D/N/A|P.. GarageMonitor:Ambient:temperature:$unit -->[]
(4)[n] X/D/N/P/A GarageMonitor:Ambient:temperature:$unit:$unit -->[°F] 
(3)[n] X/D/N/A|P.. GarageMonitor:Ambient:temperature:temperature -->[26.18]
(1)[1k] X/D/?...... Bucket[1]=GarageMonitor, enum Key=Presence Value=[] 
(2)[n] X/D/N.... GarageMonitor:Presence:$name -->[] 
(3)[n] X/D/N/A|P.. GarageMonitor:Presence:$name:$name -->[Motion Sensor]
(2)[n] X/D/N.... GarageMonitor:Presence:$properties -->[] 
(3)[n] X/D/N/A|P.. GarageMonitor:Presence:$properties:$properties -->[motion]
(2)[n] X/D/N.... GarageMonitor:Presence:$type -->[] 
(3)[n] X/D/N/A|P.. GarageMonitor:Presence:$type:$type -->[sensor]
(2)[n] X/D/N.... GarageMonitor:Presence:motion -->[] 
(3)[n] X/D/N/A|P.. GarageMonitor:Presence:motion:$datatype -->[]
(4)[n] X/D/N/P/A GarageMonitor:Presence:motion:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. GarageMonitor:Presence:motion:$format -->[]
(4)[n] X/D/N/P/A GarageMonitor:Presence:motion:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. GarageMonitor:Presence:motion:$name -->[]
(4)[n] X/D/N/P/A GarageMonitor:Presence:motion:$name:$name -->[Motion] 
(3)[n] X/D/N/A|P.. GarageMonitor:Presence:motion:motion -->[OPEN]
(1)[1k] X/D/?...... Bucket[1]=GarageMonitor, enum Key=hardware Value=[] 
(2)[n] X/D/N.... GarageMonitor:hardware:$name -->[] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:$name:$name -->[Device Info]
(2)[n] X/D/N.... GarageMonitor:hardware:$properties -->[] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:$properties:$properties -->[signal,mac,resetReason,voltage]
(2)[n] X/D/N.... GarageMonitor:hardware:$type -->[] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:$type:$type -->[sensor]
(2)[n] X/D/N.... GarageMonitor:hardware:mac -->[] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:mac:$datatype -->[]
(4)[n] X/D/N/P/A GarageMonitor:hardware:mac:$datatype:$datatype -->[sring] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:mac:$name -->[]
(4)[n] X/D/N/P/A GarageMonitor:hardware:mac:$name:$name -->[mac] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:mac:mac -->[B4:E6:2D:1B:5C:4D]
(2)[n] X/D/N.... GarageMonitor:hardware:resetReason -->[] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:resetReason:$datatype -->[]
(4)[n] X/D/N/P/A GarageMonitor:hardware:resetReason:$datatype:$datatype -->[string] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:resetReason:$name -->[]
(4)[n] X/D/N/P/A GarageMonitor:hardware:resetReason:$name:$name -->[Last Reset Reason] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:resetReason:resetReason -->[External System]
(2)[n] X/D/N.... GarageMonitor:hardware:signal -->[] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:signal:$datatype -->[]
(4)[n] X/D/N/P/A GarageMonitor:hardware:signal:$datatype:$datatype -->[integer] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:signal:$name -->[]
(4)[n] X/D/N/P/A GarageMonitor:hardware:signal:$name:$name -->[RSSI] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:signal:$unit -->[]
(4)[n] X/D/N/P/A GarageMonitor:hardware:signal:$unit:$unit -->[dBm] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:signal:signal -->[-73]
(2)[n] X/D/N.... GarageMonitor:hardware:voltage -->[] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:voltage:$datatype -->[]
(4)[n] X/D/N/P/A GarageMonitor:hardware:voltage:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:voltage:$name -->[]
(4)[n] X/D/N/P/A GarageMonitor:hardware:voltage:$name:$name -->[3.3V Supply] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:voltage:$unit -->[]
(4)[n] X/D/N/P/A GarageMonitor:hardware:voltage:$unit:$unit -->[V] 
(3)[n] X/D/N/A|P.. GarageMonitor:hardware:voltage:voltage -->[3.05]
(0) X...... sknSensors 
(1) X/D...... sknSensors:GarageProvider 
(1)[1k] X/D/?...... Bucket[1]=GarageProvider, enum Key=$fw Value=[] 
(2)[a] X/D/A.... GarageProvider:$fw:checksum -->[] 
(3)[a] X/D/A/P.. GarageProvider:$fw:checksum:checksum -->[39286ba639a9ef1d5395e1c0e45d13fa]
(2)[a] X/D/A.... GarageProvider:$fw:name -->[] 
(3)[a] X/D/A/P.. GarageProvider:$fw:name:name -->[GarageProvider]
(2)[a] X/D/A.... GarageProvider:$fw:version -->[] 
(3)[a] X/D/A/P.. GarageProvider:$fw:version:version -->[1.1.0]
(1)[1k] X/D/?...... Bucket[1]=GarageProvider, enum Key=$homie Value=[] 
(2)[a] X/D/A.... GarageProvider:$homie:$homie -->[3.0.1] 
(1)[1k] X/D/?...... Bucket[1]=GarageProvider, enum Key=$implementation Value=[] 
(2)[a] X/D/A.... GarageProvider:$implementation:$implementation -->[esp32] 
(2)[a] X/D/A.... GarageProvider:$implementation:config -->[] 
(3)[a] X/D/A/P.. GarageProvider:$implementation:config:config -->[{"name":"Garage Provider","device_id":"GarageProvider","device_stats_interval":900,"wifi":{"ssid":"SFNSS1-24G"},"mqtt":{"host":"openhab.skoona.net","port":1883,"base_topic":"sknSensors/","auth":true},"ota":{"enabled":true},"settings":{"sensorsInterval":900}}]
(2)[a] X/D/A.... GarageProvider:$implementation:ota -->[] 
(3)[a] X/D/A/P.. GarageProvider:$implementation:ota:enabled -->[]
(4)[a] X/D/A/P/P GarageProvider:$implementation:ota:enabled:enabled -->[true] 
(2)[a] X/D/A.... GarageProvider:$implementation:version -->[] 
(3)[a] X/D/A/P.. GarageProvider:$implementation:version:version -->[3.0.0]
(1)[1k] X/D/?...... Bucket[1]=GarageProvider, enum Key=$localip Value=[] 
(2)[a] X/D/A.... GarageProvider:$localip:$localip -->[10.100.1.218] 
(1)[1k] X/D/?...... Bucket[1]=GarageProvider, enum Key=$mac Value=[] 
(2)[a] X/D/A.... GarageProvider:$mac:$mac -->[24:6F:28:97:63:B8] 
(1)[1k] X/D/?...... Bucket[1]=GarageProvider, enum Key=$name Value=[] 
(2)[a] X/D/A.... GarageProvider:$name:$name -->[Garage Provider] 
(1)[1k] X/D/?...... Bucket[1]=GarageProvider, enum Key=$nodes Value=[] 
(2)[a] X/D/A.... GarageProvider:$nodes:$nodes -->[garageDoor,environmentMonitor] 
(1)[1k] X/D/?...... Bucket[1]=GarageProvider, enum Key=$state Value=[] 
(2)[a] X/D/A.... GarageProvider:$state:$state -->[ready] 
(1)[1k] X/D/?...... Bucket[1]=GarageProvider, enum Key=$stats Value=[] 
(2)[a] X/D/A.... GarageProvider:$stats:$stats -->[uptime] 
(2)[a] X/D/A.... GarageProvider:$stats:interval -->[] 
(3)[a] X/D/A/P.. GarageProvider:$stats:interval:interval -->[905]
(2)[a] X/D/A.... GarageProvider:$stats:signal -->[] 
(3)[a] X/D/A/P.. GarageProvider:$stats:signal:signal -->[72]
(2)[a] X/D/A.... GarageProvider:$stats:uptime -->[] 
(3)[a] X/D/A/P.. GarageProvider:$stats:uptime:uptime -->[113713]
(1)[1k] X/D/?...... Bucket[1]=GarageProvider, enum Key=environmentMonitor Value=[] 
(2)[n] X/D/N.... GarageProvider:environmentMonitor:$name -->[] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:$name:$name -->[Garage]
(2)[n] X/D/N.... GarageProvider:environmentMonitor:$properties -->[] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:$properties:$properties -->[motion,temperature,humdity]
(2)[n] X/D/N.... GarageProvider:environmentMonitor:$type -->[] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:$type:$type -->[SknSensors]
(2)[n] X/D/N.... GarageProvider:environmentMonitor:humdity -->[] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:humdity:$datatype -->[]
(4)[n] X/D/N/P/A GarageProvider:environmentMonitor:humdity:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:humdity:$format -->[]
(4)[n] X/D/N/P/A GarageProvider:environmentMonitor:humdity:$format:$format -->[%.1f] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:humdity:$name -->[]
(4)[n] X/D/N/P/A GarageProvider:environmentMonitor:humdity:$name:$name -->[Humidity] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:humdity:$unit -->[]
(4)[n] X/D/N/P/A GarageProvider:environmentMonitor:humdity:$unit:$unit -->[%] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:humdity:humdity -->[27.0]
(2)[n] X/D/N.... GarageProvider:environmentMonitor:motion -->[] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:motion:$datatype -->[]
(4)[n] X/D/N/P/A GarageProvider:environmentMonitor:motion:$datatype:$datatype -->[string] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:motion:$format -->[]
(4)[n] X/D/N/P/A GarageProvider:environmentMonitor:motion:$format:$format -->[%s] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:motion:$name -->[]
(4)[n] X/D/N/P/A GarageProvider:environmentMonitor:motion:$name:$name -->[Motion] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:motion:motion -->[ON]
(2)[n] X/D/N.... GarageProvider:environmentMonitor:temperature -->[] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:temperature:$datatype -->[]
(4)[n] X/D/N/P/A GarageProvider:environmentMonitor:temperature:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:temperature:$format -->[]
(4)[n] X/D/N/P/A GarageProvider:environmentMonitor:temperature:$format:$format -->[%.1f] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:temperature:$name -->[]
(4)[n] X/D/N/P/A GarageProvider:environmentMonitor:temperature:$name:$name -->[Temperature] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:temperature:$unit -->[]
(4)[n] X/D/N/P/A GarageProvider:environmentMonitor:temperature:$unit:$unit -->[ºF] 
(3)[n] X/D/N/A|P.. GarageProvider:environmentMonitor:temperature:temperature -->[73.6]
(1)[1k] X/D/?...... Bucket[1]=GarageProvider, enum Key=garageDoor Value=[] 
(2)[n] X/D/N.... GarageProvider:garageDoor:$name -->[] 
(3)[n] X/D/N/A|P.. GarageProvider:garageDoor:$name:$name -->[Operations]
(2)[n] X/D/N.... GarageProvider:garageDoor:$properties -->[] 
(3)[n] X/D/N/A|P.. GarageProvider:garageDoor:$properties:$properties -->[operator,positon,direction]
(2)[n] X/D/N.... GarageProvider:garageDoor:$type -->[] 
(3)[n] X/D/N/A|P.. GarageProvider:garageDoor:$type:$type -->[SknSensors]
(2)[n] X/D/N.... GarageProvider:garageDoor:direction -->[] 
(3)[n] X/D/N/A|P.. GarageProvider:garageDoor:direction:$datatype -->[]
(4)[n] X/D/N/P/A GarageProvider:garageDoor:direction:$datatype:$datatype -->[string] 
(3)[n] X/D/N/A|P.. GarageProvider:garageDoor:direction:$format -->[]
(4)[n] X/D/N/P/A GarageProvider:garageDoor:direction:$format:$format -->[%s] 
(3)[n] X/D/N/A|P.. GarageProvider:garageDoor:direction:$name -->[]
(4)[n] X/D/N/P/A GarageProvider:garageDoor:direction:$name:$name -->[Travel Direction] 
(3)[n] X/D/N/A|P.. GarageProvider:garageDoor:direction:direction -->[OPENING]
(2)[n] X/D/N.... GarageProvider:garageDoor:operator -->[] 
(3)[n] X/D/N/A|P.. GarageProvider:garageDoor:operator:$datatype -->[]
(4)[n] X/D/N/P/A GarageProvider:garageDoor:operator:$datatype:$datatype -->[boolean] 
(3)[n] X/D/N/A|P.. GarageProvider:garageDoor:operator:$format -->[]
(4)[n] X/D/N/P/A GarageProvider:garageDoor:operator:$format:$format -->[%s] 
(3)[n] X/D/N/A|P.. GarageProvider:garageDoor:operator:$name -->[]
(4)[n] X/D/N/P/A GarageProvider:garageDoor:operator:$name:$name -->[Operator] 
(3)[n] X/D/N/A|P.. GarageProvider:garageDoor:operator:$settable -->[]
(4)[n] X/D/N/P/A GarageProvider:garageDoor:operator:$settable:$settable -->[true] 
(2)[n] X/D/N.... GarageProvider:garageDoor:positon -->[] 
(3)[n] X/D/N/A|P.. GarageProvider:garageDoor:positon:$datatype -->[]
(4)[n] X/D/N/P/A GarageProvider:garageDoor:positon:$datatype:$datatype -->[integer] 
(3)[n] X/D/N/A|P.. GarageProvider:garageDoor:positon:$format -->[]
(4)[n] X/D/N/P/A GarageProvider:garageDoor:positon:$format:$format -->[%d] 
(3)[n] X/D/N/A|P.. GarageProvider:garageDoor:positon:$name -->[]
(4)[n] X/D/N/P/A GarageProvider:garageDoor:positon:$name:$name -->[Position MM] 
(3)[n] X/D/N/A|P.. GarageProvider:garageDoor:positon:$unit -->[]
(4)[n] X/D/N/P/A GarageProvider:garageDoor:positon:$unit:$unit -->[mm] 
(3)[n] X/D/N/A|P.. GarageProvider:garageDoor:positon:positon -->[19]
(0) X...... sknSensors 
(1) X/D...... sknSensors:GuestRoom 
(1)[1k] X/D/?...... Bucket[1]=GuestRoom, enum Key=$fw Value=[] 
(2)[a] X/D/A.... GuestRoom:$fw:checksum -->[] 
(3)[a] X/D/A/P.. GuestRoom:$fw:checksum:checksum -->[3ef8fbb48c5b23788a22a998c14a1a6d]
(2)[a] X/D/A.... GuestRoom:$fw:name -->[] 
(3)[a] X/D/A/P.. GuestRoom:$fw:name:name -->[Monitor-DHT-RCWL-Metrics]
(2)[a] X/D/A.... GuestRoom:$fw:version -->[] 
(3)[a] X/D/A/P.. GuestRoom:$fw:version:version -->[2.0.0]
(1)[1k] X/D/?...... Bucket[1]=GuestRoom, enum Key=$homie Value=[] 
(2)[a] X/D/A.... GuestRoom:$homie:$homie -->[3.0.1] 
(1)[1k] X/D/?...... Bucket[1]=GuestRoom, enum Key=$implementation Value=[] 
(2)[a] X/D/A.... GuestRoom:$implementation:$implementation -->[esp8266] 
(2)[a] X/D/A.... GuestRoom:$implementation:config -->[] 
(3)[a] X/D/A/P.. GuestRoom:$implementation:config:config -->[{"name":"Guest Room","device_id":"GuestRoom","device_stats_interval":180,"wifi":{"ssid":"SFNSS1-24G"},"mqtt":{"host":"openhab.skoona.net","port":1883,"base_topic":"sknSensors/","auth":true},"ota":{"enabled":true},"settings":{"sensorInterval":900,"motionHoldInterval":60}}]
(2)[a] X/D/A.... GuestRoom:$implementation:ota -->[] 
(3)[a] X/D/A/P.. GuestRoom:$implementation:ota:enabled -->[]
(4)[a] X/D/A/P/P GuestRoom:$implementation:ota:enabled:enabled -->[true] 
(2)[a] X/D/A.... GuestRoom:$implementation:version -->[] 
(3)[a] X/D/A/P.. GuestRoom:$implementation:version:version -->[3.0.0]
(1)[1k] X/D/?...... Bucket[1]=GuestRoom, enum Key=$localip Value=[] 
(2)[a] X/D/A.... GuestRoom:$localip:$localip -->[10.100.1.178] 
(1)[1k] X/D/?...... Bucket[1]=GuestRoom, enum Key=$mac Value=[] 
(2)[a] X/D/A.... GuestRoom:$mac:$mac -->[BC:DD:C2:81:04:72] 
(1)[1k] X/D/?...... Bucket[1]=GuestRoom, enum Key=$name Value=[] 
(2)[a] X/D/A.... GuestRoom:$name:$name -->[Guest Room] 
(1)[1k] X/D/?...... Bucket[1]=GuestRoom, enum Key=$nodes Value=[] 
(2)[a] X/D/A.... GuestRoom:$nodes:$nodes -->[Ambient,Presence,hardware] 
(1)[1k] X/D/?...... Bucket[1]=GuestRoom, enum Key=$state Value=[] 
(2)[a] X/D/A.... GuestRoom:$state:$state -->[ready] 
(1)[1k] X/D/?...... Bucket[1]=GuestRoom, enum Key=$stats Value=[] 
(2)[a] X/D/A.... GuestRoom:$stats:$stats -->[uptime] 
(2)[a] X/D/A.... GuestRoom:$stats:interval -->[] 
(3)[a] X/D/A/P.. GuestRoom:$stats:interval:interval -->[185]
(2)[a] X/D/A.... GuestRoom:$stats:signal -->[] 
(3)[a] X/D/A/P.. GuestRoom:$stats:signal:signal -->[72]
(2)[a] X/D/A.... GuestRoom:$stats:uptime -->[] 
(3)[a] X/D/A/P.. GuestRoom:$stats:uptime:uptime -->[362358]
(1)[1k] X/D/?...... Bucket[1]=GuestRoom, enum Key=Ambient Value=[] 
(2)[n] X/D/N.... GuestRoom:Ambient:$name -->[] 
(3)[n] X/D/N/A|P.. GuestRoom:Ambient:$name:$name -->[Temperature and Humidity Sensor]
(2)[n] X/D/N.... GuestRoom:Ambient:$properties -->[] 
(3)[n] X/D/N/A|P.. GuestRoom:Ambient:$properties:$properties -->[humidity,temperature]
(2)[n] X/D/N.... GuestRoom:Ambient:$type -->[] 
(3)[n] X/D/N/A|P.. GuestRoom:Ambient:$type:$type -->[sensor]
(2)[n] X/D/N.... GuestRoom:Ambient:humidity -->[] 
(3)[n] X/D/N/A|P.. GuestRoom:Ambient:humidity:$datatype -->[]
(4)[n] X/D/N/P/A GuestRoom:Ambient:humidity:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. GuestRoom:Ambient:humidity:$name -->[]
(4)[n] X/D/N/P/A GuestRoom:Ambient:humidity:$name:$name -->[Humidity] 
(3)[n] X/D/N/A|P.. GuestRoom:Ambient:humidity:$unit -->[]
(4)[n] X/D/N/P/A GuestRoom:Ambient:humidity:$unit:$unit -->[%rH] 
(3)[n] X/D/N/A|P.. GuestRoom:Ambient:humidity:humidity -->[11.10]
(2)[n] X/D/N.... GuestRoom:Ambient:temperature -->[] 
(3)[n] X/D/N/A|P.. GuestRoom:Ambient:temperature:$datatype -->[]
(4)[n] X/D/N/P/A GuestRoom:Ambient:temperature:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. GuestRoom:Ambient:temperature:$name -->[]
(4)[n] X/D/N/P/A GuestRoom:Ambient:temperature:$name:$name -->[Temperature] 
(3)[n] X/D/N/A|P.. GuestRoom:Ambient:temperature:$unit -->[]
(4)[n] X/D/N/P/A GuestRoom:Ambient:temperature:$unit:$unit -->[°F] 
(3)[n] X/D/N/A|P.. GuestRoom:Ambient:temperature:temperature -->[73.76]
(1)[1k] X/D/?...... Bucket[1]=GuestRoom, enum Key=Presence Value=[] 
(2)[n] X/D/N.... GuestRoom:Presence:$name -->[] 
(3)[n] X/D/N/A|P.. GuestRoom:Presence:$name:$name -->[Motion Sensor]
(2)[n] X/D/N.... GuestRoom:Presence:$properties -->[] 
(3)[n] X/D/N/A|P.. GuestRoom:Presence:$properties:$properties -->[motion]
(2)[n] X/D/N.... GuestRoom:Presence:$type -->[] 
(3)[n] X/D/N/A|P.. GuestRoom:Presence:$type:$type -->[sensor]
(2)[n] X/D/N.... GuestRoom:Presence:motion -->[] 
(3)[n] X/D/N/A|P.. GuestRoom:Presence:motion:$datatype -->[]
(4)[n] X/D/N/P/A GuestRoom:Presence:motion:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. GuestRoom:Presence:motion:$format -->[]
(4)[n] X/D/N/P/A GuestRoom:Presence:motion:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. GuestRoom:Presence:motion:$name -->[]
(4)[n] X/D/N/P/A GuestRoom:Presence:motion:$name:$name -->[Motion] 
(3)[n] X/D/N/A|P.. GuestRoom:Presence:motion:motion -->[CLOSED]
(1)[1k] X/D/?...... Bucket[1]=GuestRoom, enum Key=hardware Value=[] 
(2)[n] X/D/N.... GuestRoom:hardware:$name -->[] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:$name:$name -->[Device Info]
(2)[n] X/D/N.... GuestRoom:hardware:$properties -->[] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:$properties:$properties -->[signal,mac,resetReason,voltage]
(2)[n] X/D/N.... GuestRoom:hardware:$type -->[] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:$type:$type -->[sensor]
(2)[n] X/D/N.... GuestRoom:hardware:mac -->[] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:mac:$datatype -->[]
(4)[n] X/D/N/P/A GuestRoom:hardware:mac:$datatype:$datatype -->[sring] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:mac:$name -->[]
(4)[n] X/D/N/P/A GuestRoom:hardware:mac:$name:$name -->[mac] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:mac:mac -->[BC:DD:C2:81:04:72]
(2)[n] X/D/N.... GuestRoom:hardware:resetReason -->[] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:resetReason:$datatype -->[]
(4)[n] X/D/N/P/A GuestRoom:hardware:resetReason:$datatype:$datatype -->[string] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:resetReason:$name -->[]
(4)[n] X/D/N/P/A GuestRoom:hardware:resetReason:$name:$name -->[Last Reset Reason] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:resetReason:resetReason -->[External System]
(2)[n] X/D/N.... GuestRoom:hardware:signal -->[] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:signal:$datatype -->[]
(4)[n] X/D/N/P/A GuestRoom:hardware:signal:$datatype:$datatype -->[integer] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:signal:$name -->[]
(4)[n] X/D/N/P/A GuestRoom:hardware:signal:$name:$name -->[RSSI] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:signal:$unit -->[]
(4)[n] X/D/N/P/A GuestRoom:hardware:signal:$unit:$unit -->[dBm] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:signal:signal -->[-64]
(2)[n] X/D/N.... GuestRoom:hardware:voltage -->[] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:voltage:$datatype -->[]
(4)[n] X/D/N/P/A GuestRoom:hardware:voltage:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:voltage:$name -->[]
(4)[n] X/D/N/P/A GuestRoom:hardware:voltage:$name:$name -->[3.3V Supply] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:voltage:$unit -->[]
(4)[n] X/D/N/P/A GuestRoom:hardware:voltage:$unit:$unit -->[V] 
(3)[n] X/D/N/A|P.. GuestRoom:hardware:voltage:voltage -->[2.98]
(0) X...... sknSensors 
(1) X/D...... sknSensors:HomeOffice 
(1)[1k] X/D/?...... Bucket[1]=HomeOffice, enum Key=$fw Value=[] 
(2)[a] X/D/A.... HomeOffice:$fw:checksum -->[] 
(3)[a] X/D/A/P.. HomeOffice:$fw:checksum:checksum -->[3ef8fbb48c5b23788a22a998c14a1a6d]
(2)[a] X/D/A.... HomeOffice:$fw:name -->[] 
(3)[a] X/D/A/P.. HomeOffice:$fw:name:name -->[Monitor-DHT-RCWL-Metrics]
(2)[a] X/D/A.... HomeOffice:$fw:version -->[] 
(3)[a] X/D/A/P.. HomeOffice:$fw:version:version -->[2.0.0]
(1)[1k] X/D/?...... Bucket[1]=HomeOffice, enum Key=$homie Value=[] 
(2)[a] X/D/A.... HomeOffice:$homie:$homie -->[3.0.1] 
(1)[1k] X/D/?...... Bucket[1]=HomeOffice, enum Key=$implementation Value=[] 
(2)[a] X/D/A.... HomeOffice:$implementation:$implementation -->[esp8266] 
(2)[a] X/D/A.... HomeOffice:$implementation:config -->[] 
(3)[a] X/D/A/P.. HomeOffice:$implementation:config:config -->[{"name":"Home Office","device_id":"HomeOffice","device_stats_interval":180,"wifi":{"ssid":"SFNSS1-24G"},"mqtt":{"host":"openhab.skoona.net","port":1883,"base_topic":"sknSensors/","auth":true},"ota":{"enabled":true},"settings":{"sensorInterval":900,"motionHoldInterval":60}}]
(2)[a] X/D/A.... HomeOffice:$implementation:ota -->[] 
(3)[a] X/D/A/P.. HomeOffice:$implementation:ota:enabled -->[]
(4)[a] X/D/A/P/P HomeOffice:$implementation:ota:enabled:enabled -->[true] 
(2)[a] X/D/A.... HomeOffice:$implementation:version -->[] 
(3)[a] X/D/A/P.. HomeOffice:$implementation:version:version -->[3.0.0]
(1)[1k] X/D/?...... Bucket[1]=HomeOffice, enum Key=$localip Value=[] 
(2)[a] X/D/A.... HomeOffice:$localip:$localip -->[10.100.1.161] 
(1)[1k] X/D/?...... Bucket[1]=HomeOffice, enum Key=$mac Value=[] 
(2)[a] X/D/A.... HomeOffice:$mac:$mac -->[84:F3:EB:B7:55:D5] 
(1)[1k] X/D/?...... Bucket[1]=HomeOffice, enum Key=$name Value=[] 
(2)[a] X/D/A.... HomeOffice:$name:$name -->[Home Office] 
(1)[1k] X/D/?...... Bucket[1]=HomeOffice, enum Key=$nodes Value=[] 
(2)[a] X/D/A.... HomeOffice:$nodes:$nodes -->[Ambient,Presence,hardware] 
(1)[1k] X/D/?...... Bucket[1]=HomeOffice, enum Key=$state Value=[] 
(2)[a] X/D/A.... HomeOffice:$state:$state -->[ready] 
(1)[1k] X/D/?...... Bucket[1]=HomeOffice, enum Key=$stats Value=[] 
(2)[a] X/D/A.... HomeOffice:$stats:$stats -->[uptime] 
(2)[a] X/D/A.... HomeOffice:$stats:interval -->[] 
(3)[a] X/D/A/P.. HomeOffice:$stats:interval:interval -->[185]
(2)[a] X/D/A.... HomeOffice:$stats:signal -->[] 
(3)[a] X/D/A/P.. HomeOffice:$stats:signal:signal -->[70]
(2)[a] X/D/A.... HomeOffice:$stats:uptime -->[] 
(3)[a] X/D/A/P.. HomeOffice:$stats:uptime:uptime -->[46625]
(1)[1k] X/D/?...... Bucket[1]=HomeOffice, enum Key=Ambient Value=[] 
(2)[n] X/D/N.... HomeOffice:Ambient:$name -->[] 
(3)[n] X/D/N/A|P.. HomeOffice:Ambient:$name:$name -->[Temperature and Humidity Sensor]
(2)[n] X/D/N.... HomeOffice:Ambient:$properties -->[] 
(3)[n] X/D/N/A|P.. HomeOffice:Ambient:$properties:$properties -->[humidity,temperature]
(2)[n] X/D/N.... HomeOffice:Ambient:$type -->[] 
(3)[n] X/D/N/A|P.. HomeOffice:Ambient:$type:$type -->[sensor]
(2)[n] X/D/N.... HomeOffice:Ambient:humidity -->[] 
(3)[n] X/D/N/A|P.. HomeOffice:Ambient:humidity:$datatype -->[]
(4)[n] X/D/N/P/A HomeOffice:Ambient:humidity:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. HomeOffice:Ambient:humidity:$name -->[]
(4)[n] X/D/N/P/A HomeOffice:Ambient:humidity:$name:$name -->[Humidity] 
(3)[n] X/D/N/A|P.. HomeOffice:Ambient:humidity:$unit -->[]
(4)[n] X/D/N/P/A HomeOffice:Ambient:humidity:$unit:$unit -->[%rH] 
(3)[n] X/D/N/A|P.. HomeOffice:Ambient:humidity:humidity -->[10.20]
(2)[n] X/D/N.... HomeOffice:Ambient:temperature -->[] 
(3)[n] X/D/N/A|P.. HomeOffice:Ambient:temperature:$datatype -->[]
(4)[n] X/D/N/P/A HomeOffice:Ambient:temperature:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. HomeOffice:Ambient:temperature:$name -->[]
(4)[n] X/D/N/P/A HomeOffice:Ambient:temperature:$name:$name -->[Temperature] 
(3)[n] X/D/N/A|P.. HomeOffice:Ambient:temperature:$unit -->[]
(4)[n] X/D/N/P/A HomeOffice:Ambient:temperature:$unit:$unit -->[°F] 
(3)[n] X/D/N/A|P.. HomeOffice:Ambient:temperature:temperature -->[75.20]
(1)[1k] X/D/?...... Bucket[1]=HomeOffice, enum Key=Presence Value=[] 
(2)[n] X/D/N.... HomeOffice:Presence:$name -->[] 
(3)[n] X/D/N/A|P.. HomeOffice:Presence:$name:$name -->[Motion Sensor]
(2)[n] X/D/N.... HomeOffice:Presence:$properties -->[] 
(3)[n] X/D/N/A|P.. HomeOffice:Presence:$properties:$properties -->[motion]
(2)[n] X/D/N.... HomeOffice:Presence:$type -->[] 
(3)[n] X/D/N/A|P.. HomeOffice:Presence:$type:$type -->[sensor]
(2)[n] X/D/N.... HomeOffice:Presence:motion -->[] 
(3)[n] X/D/N/A|P.. HomeOffice:Presence:motion:$datatype -->[]
(4)[n] X/D/N/P/A HomeOffice:Presence:motion:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. HomeOffice:Presence:motion:$format -->[]
(4)[n] X/D/N/P/A HomeOffice:Presence:motion:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. HomeOffice:Presence:motion:$name -->[]
(4)[n] X/D/N/P/A HomeOffice:Presence:motion:$name:$name -->[Motion] 
(3)[n] X/D/N/A|P.. HomeOffice:Presence:motion:motion -->[OPEN]
(1)[1k] X/D/?...... Bucket[1]=HomeOffice, enum Key=hardware Value=[] 
(2)[n] X/D/N.... HomeOffice:hardware:$name -->[] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:$name:$name -->[Device Info]
(2)[n] X/D/N.... HomeOffice:hardware:$properties -->[] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:$properties:$properties -->[signal,mac,resetReason,voltage]
(2)[n] X/D/N.... HomeOffice:hardware:$type -->[] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:$type:$type -->[sensor]
(2)[n] X/D/N.... HomeOffice:hardware:mac -->[] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:mac:$datatype -->[]
(4)[n] X/D/N/P/A HomeOffice:hardware:mac:$datatype:$datatype -->[sring] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:mac:$name -->[]
(4)[n] X/D/N/P/A HomeOffice:hardware:mac:$name:$name -->[mac] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:mac:mac -->[84:F3:EB:B7:55:D5]
(2)[n] X/D/N.... HomeOffice:hardware:resetReason -->[] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:resetReason:$datatype -->[]
(4)[n] X/D/N/P/A HomeOffice:hardware:resetReason:$datatype:$datatype -->[string] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:resetReason:$name -->[]
(4)[n] X/D/N/P/A HomeOffice:hardware:resetReason:$name:$name -->[Last Reset Reason] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:resetReason:resetReason -->[External System]
(2)[n] X/D/N.... HomeOffice:hardware:signal -->[] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:signal:$datatype -->[]
(4)[n] X/D/N/P/A HomeOffice:hardware:signal:$datatype:$datatype -->[integer] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:signal:$name -->[]
(4)[n] X/D/N/P/A HomeOffice:hardware:signal:$name:$name -->[RSSI] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:signal:$unit -->[]
(4)[n] X/D/N/P/A HomeOffice:hardware:signal:$unit:$unit -->[dBm] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:signal:signal -->[-65]
(2)[n] X/D/N.... HomeOffice:hardware:voltage -->[] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:voltage:$datatype -->[]
(4)[n] X/D/N/P/A HomeOffice:hardware:voltage:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:voltage:$name -->[]
(4)[n] X/D/N/P/A HomeOffice:hardware:voltage:$name:$name -->[3.3V Supply] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:voltage:$unit -->[]
(4)[n] X/D/N/P/A HomeOffice:hardware:voltage:$unit:$unit -->[V] 
(3)[n] X/D/N/A|P.. HomeOffice:hardware:voltage:voltage -->[2.98]
(0) X...... sknSensors 
(1) X/D...... sknSensors:MCPProvider 
(1)[1k] X/D/?...... Bucket[1]=MCPProvider, enum Key=$fw Value=[] 
(2)[a] X/D/A.... MCPProvider:$fw:checksum -->[] 
(3)[a] X/D/A/P.. MCPProvider:$fw:checksum:checksum -->[b92b40fead22529097629c0597186c05]
(2)[a] X/D/A.... MCPProvider:$fw:name -->[] 
(3)[a] X/D/A/P.. MCPProvider:$fw:name:name -->[WiredProvider]
(2)[a] X/D/A.... MCPProvider:$fw:version -->[] 
(3)[a] X/D/A/P.. MCPProvider:$fw:version:version -->[1.4.0]
(1)[1k] X/D/?...... Bucket[1]=MCPProvider, enum Key=$homie Value=[] 
(2)[a] X/D/A.... MCPProvider:$homie:$homie -->[3.0.1] 
(1)[1k] X/D/?...... Bucket[1]=MCPProvider, enum Key=$implementation Value=[] 
(2)[a] X/D/A.... MCPProvider:$implementation:$implementation -->[esp8266] 
(2)[a] X/D/A.... MCPProvider:$implementation:config -->[] 
(3)[a] X/D/A/P.. MCPProvider:$implementation:config:config -->[{"name":"Hard Wired Alarm Sensors","device_id":"MCPProvider","device_stats_interval":300,"wifi":{"ssid":"SFNSS1-24G"},"mqtt":{"host":"openhab.skoona.net","port":1883,"base_topic":"sknSensors/","auth":true},"ota":{"enabled":true},"settings":{"ipolA":255,"ipolB":255}}]
(2)[a] X/D/A.... MCPProvider:$implementation:ota -->[] 
(3)[a] X/D/A/P.. MCPProvider:$implementation:ota:enabled -->[]
(4)[a] X/D/A/P/P MCPProvider:$implementation:ota:enabled:enabled -->[true] 
(2)[a] X/D/A.... MCPProvider:$implementation:version -->[] 
(3)[a] X/D/A/P.. MCPProvider:$implementation:version:version -->[3.0.0]
(1)[1k] X/D/?...... Bucket[1]=MCPProvider, enum Key=$localip Value=[] 
(2)[a] X/D/A.... MCPProvider:$localip:$localip -->[10.100.1.222] 
(1)[1k] X/D/?...... Bucket[1]=MCPProvider, enum Key=$mac Value=[] 
(2)[a] X/D/A.... MCPProvider:$mac:$mac -->[BC:DD:C2:24:B7:3C] 
(1)[1k] X/D/?...... Bucket[1]=MCPProvider, enum Key=$name Value=[] 
(2)[a] X/D/A.... MCPProvider:$name:$name -->[Hard Wired Alarm Sensors] 
(1)[1k] X/D/?...... Bucket[1]=MCPProvider, enum Key=$nodes Value=[] 
(2)[a] X/D/A.... MCPProvider:$nodes:$nodes -->[wiredMonitor,hardware] 
(1)[1k] X/D/?...... Bucket[1]=MCPProvider, enum Key=$state Value=[] 
(2)[a] X/D/A.... MCPProvider:$state:$state -->[ready] 
(1)[1k] X/D/?...... Bucket[1]=MCPProvider, enum Key=$stats Value=[] 
(2)[a] X/D/A.... MCPProvider:$stats:$stats -->[uptime] 
(2)[a] X/D/A.... MCPProvider:$stats:interval -->[] 
(3)[a] X/D/A/P.. MCPProvider:$stats:interval:interval -->[305]
(2)[a] X/D/A.... MCPProvider:$stats:signal -->[] 
(3)[a] X/D/A/P.. MCPProvider:$stats:signal:signal -->[58]
(2)[a] X/D/A.... MCPProvider:$stats:uptime -->[] 
(3)[a] X/D/A/P.. MCPProvider:$stats:uptime:uptime -->[113410]
(1)[1k] X/D/?...... Bucket[1]=MCPProvider, enum Key=hardware Value=[] 
(2)[n] X/D/N.... MCPProvider:hardware:$name -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:$name:$name -->[Device Info]
(2)[n] X/D/N.... MCPProvider:hardware:$properties -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:$properties:$properties -->[signal,mac,resetReason,voltage]
(2)[n] X/D/N.... MCPProvider:hardware:$type -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:$type:$type -->[sensor]
(2)[n] X/D/N.... MCPProvider:hardware:mac -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:mac:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:hardware:mac:$datatype:$datatype -->[sring] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:mac:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:hardware:mac:$name:$name -->[mac] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:mac:mac -->[BC:DD:C2:24:B7:3C]
(2)[n] X/D/N.... MCPProvider:hardware:resetReason -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:resetReason:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:hardware:resetReason:$datatype:$datatype -->[string] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:resetReason:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:hardware:resetReason:$name:$name -->[Last Reset Reason] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:resetReason:resetReason -->[External System]
(2)[n] X/D/N.... MCPProvider:hardware:signal -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:signal:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:hardware:signal:$datatype:$datatype -->[integer] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:signal:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:hardware:signal:$name:$name -->[RSSI] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:signal:$unit -->[]
(4)[n] X/D/N/P/A MCPProvider:hardware:signal:$unit:$unit -->[dBm] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:signal:signal -->[-70]
(2)[n] X/D/N.... MCPProvider:hardware:voltage -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:voltage:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:hardware:voltage:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:voltage:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:hardware:voltage:$name:$name -->[3.3V Supply] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:voltage:$unit -->[]
(4)[n] X/D/N/P/A MCPProvider:hardware:voltage:$unit:$unit -->[V] 
(3)[n] X/D/N/A|P.. MCPProvider:hardware:voltage:voltage -->[2.99]
(1)[1k] X/D/?...... Bucket[1]=MCPProvider, enum Key=wiredMonitor Value=[] 
(2)[n] X/D/N.... MCPProvider:wiredMonitor:$name -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:$name:$name -->[Wired Sensors]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:$properties -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:$properties:$properties -->[pin0,pin1,pin2,pin3,pin4,pin5,pin6,pin7,pin8,pin9,pin10,pin11,pin12,pin13,pin14,pin15]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:$type -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:$type:$type -->[sensor]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:pin0 -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin0:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin0:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin0:$format -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin0:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin0:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin0:$name:$name -->[Pin 0] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin0:pin0 -->[OPEN]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:pin1 -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin1:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin1:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin1:$format -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin1:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin1:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin1:$name:$name -->[Pin 1] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin1:pin1 -->[OPEN]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:pin10 -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin10:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin10:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin10:$format -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin10:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin10:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin10:$name:$name -->[Pin 10] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin10:pin10 -->[OPEN]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:pin11 -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin11:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin11:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin11:$format -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin11:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin11:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin11:$name:$name -->[Pin 11] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin11:pin11 -->[OPEN]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:pin12 -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin12:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin12:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin12:$format -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin12:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin12:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin12:$name:$name -->[Pin 12] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin12:pin12 -->[OPEN]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:pin13 -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin13:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin13:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin13:$format -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin13:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin13:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin13:$name:$name -->[Pin 13] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin13:pin13 -->[CLOSED]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:pin14 -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin14:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin14:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin14:$format -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin14:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin14:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin14:$name:$name -->[Pin 14] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin14:pin14 -->[OPEN]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:pin15 -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin15:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin15:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin15:$format -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin15:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin15:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin15:$name:$name -->[Pin 15] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin15:pin15 -->[OPEN]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:pin2 -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin2:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin2:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin2:$format -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin2:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin2:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin2:$name:$name -->[Pin 2] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin2:pin2 -->[OPEN]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:pin3 -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin3:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin3:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin3:$format -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin3:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin3:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin3:$name:$name -->[Pin 3] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin3:pin3 -->[OPEN]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:pin4 -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin4:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin4:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin4:$format -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin4:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin4:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin4:$name:$name -->[Pin 4] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin4:pin4 -->[OPEN]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:pin5 -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin5:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin5:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin5:$format -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin5:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin5:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin5:$name:$name -->[Pin 5] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin5:pin5 -->[OPEN]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:pin6 -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin6:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin6:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin6:$format -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin6:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin6:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin6:$name:$name -->[Pin 6] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin6:pin6 -->[OPEN]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:pin7 -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin7:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin7:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin7:$format -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin7:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin7:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin7:$name:$name -->[Pin 7] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin7:pin7 -->[OPEN]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:pin8 -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin8:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin8:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin8:$format -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin8:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin8:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin8:$name:$name -->[Pin 8] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin8:pin8 -->[OPEN]
(2)[n] X/D/N.... MCPProvider:wiredMonitor:pin9 -->[] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin9:$datatype -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin9:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin9:$format -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin9:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin9:$name -->[]
(4)[n] X/D/N/P/A MCPProvider:wiredMonitor:pin9:$name:$name -->[Pin 9] 
(3)[n] X/D/N/A|P.. MCPProvider:wiredMonitor:pin9:pin9 -->[OPEN]
(0) X...... sknSensors 
(1) X/D...... sknSensors:MediaRoom 
(1)[1k] X/D/?...... Bucket[1]=MediaRoom, enum Key=$fw Value=[] 
(2)[a] X/D/A.... MediaRoom:$fw:checksum -->[] 
(3)[a] X/D/A/P.. MediaRoom:$fw:checksum:checksum -->[3ef8fbb48c5b23788a22a998c14a1a6d]
(2)[a] X/D/A.... MediaRoom:$fw:name -->[] 
(3)[a] X/D/A/P.. MediaRoom:$fw:name:name -->[Monitor-DHT-RCWL-Metrics]
(2)[a] X/D/A.... MediaRoom:$fw:version -->[] 
(3)[a] X/D/A/P.. MediaRoom:$fw:version:version -->[2.0.0]
(1)[1k] X/D/?...... Bucket[1]=MediaRoom, enum Key=$homie Value=[] 
(2)[a] X/D/A.... MediaRoom:$homie:$homie -->[3.0.1] 
(1)[1k] X/D/?...... Bucket[1]=MediaRoom, enum Key=$implementation Value=[] 
(2)[a] X/D/A.... MediaRoom:$implementation:$implementation -->[esp8266] 
(2)[a] X/D/A.... MediaRoom:$implementation:config -->[] 
(3)[a] X/D/A/P.. MediaRoom:$implementation:config:config -->[{"name":"Media Room","device_id":"MediaRoom","device_stats_interval":180,"wifi":{"ssid":"SFNSS1-24G"},"mqtt":{"host":"openhab.skoona.net","port":1883,"base_topic":"sknSensors/","auth":true},"ota":{"enabled":true},"settings":{"sensorInterval":900,"motionHoldInterval":60}}]
(2)[a] X/D/A.... MediaRoom:$implementation:ota -->[] 
(3)[a] X/D/A/P.. MediaRoom:$implementation:ota:enabled -->[]
(4)[a] X/D/A/P/P MediaRoom:$implementation:ota:enabled:enabled -->[true] 
(2)[a] X/D/A.... MediaRoom:$implementation:version -->[] 
(3)[a] X/D/A/P.. MediaRoom:$implementation:version:version -->[3.0.0]
(1)[1k] X/D/?...... Bucket[1]=MediaRoom, enum Key=$localip Value=[] 
(2)[a] X/D/A.... MediaRoom:$localip:$localip -->[10.100.1.180] 
(1)[1k] X/D/?...... Bucket[1]=MediaRoom, enum Key=$mac Value=[] 
(2)[a] X/D/A.... MediaRoom:$mac:$mac -->[B4:E6:2D:15:50:3A] 
(1)[1k] X/D/?...... Bucket[1]=MediaRoom, enum Key=$name Value=[] 
(2)[a] X/D/A.... MediaRoom:$name:$name -->[Media Room] 
(1)[1k] X/D/?...... Bucket[1]=MediaRoom, enum Key=$nodes Value=[] 
(2)[a] X/D/A.... MediaRoom:$nodes:$nodes -->[Ambient,Presence,hardware] 
(1)[1k] X/D/?...... Bucket[1]=MediaRoom, enum Key=$state Value=[] 
(2)[a] X/D/A.... MediaRoom:$state:$state -->[ready] 
(1)[1k] X/D/?...... Bucket[1]=MediaRoom, enum Key=$stats Value=[] 
(2)[a] X/D/A.... MediaRoom:$stats:$stats -->[uptime] 
(2)[a] X/D/A.... MediaRoom:$stats:interval -->[] 
(3)[a] X/D/A/P.. MediaRoom:$stats:interval:interval -->[185]
(2)[a] X/D/A.... MediaRoom:$stats:signal -->[] 
(3)[a] X/D/A/P.. MediaRoom:$stats:signal:signal -->[66]
(2)[a] X/D/A.... MediaRoom:$stats:uptime -->[] 
(3)[a] X/D/A/P.. MediaRoom:$stats:uptime:uptime -->[102606]
(1)[1k] X/D/?...... Bucket[1]=MediaRoom, enum Key=Ambient Value=[] 
(2)[n] X/D/N.... MediaRoom:Ambient:$name -->[] 
(3)[n] X/D/N/A|P.. MediaRoom:Ambient:$name:$name -->[Temperature and Humidity Sensor]
(2)[n] X/D/N.... MediaRoom:Ambient:$properties -->[] 
(3)[n] X/D/N/A|P.. MediaRoom:Ambient:$properties:$properties -->[humidity,temperature]
(2)[n] X/D/N.... MediaRoom:Ambient:$type -->[] 
(3)[n] X/D/N/A|P.. MediaRoom:Ambient:$type:$type -->[sensor]
(2)[n] X/D/N.... MediaRoom:Ambient:humidity -->[] 
(3)[n] X/D/N/A|P.. MediaRoom:Ambient:humidity:$datatype -->[]
(4)[n] X/D/N/P/A MediaRoom:Ambient:humidity:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. MediaRoom:Ambient:humidity:$name -->[]
(4)[n] X/D/N/P/A MediaRoom:Ambient:humidity:$name:$name -->[Humidity] 
(3)[n] X/D/N/A|P.. MediaRoom:Ambient:humidity:$unit -->[]
(4)[n] X/D/N/P/A MediaRoom:Ambient:humidity:$unit:$unit -->[%rH] 
(3)[n] X/D/N/A|P.. MediaRoom:Ambient:humidity:humidity -->[6.60]
(2)[n] X/D/N.... MediaRoom:Ambient:temperature -->[] 
(3)[n] X/D/N/A|P.. MediaRoom:Ambient:temperature:$datatype -->[]
(4)[n] X/D/N/P/A MediaRoom:Ambient:temperature:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. MediaRoom:Ambient:temperature:$name -->[]
(4)[n] X/D/N/P/A MediaRoom:Ambient:temperature:$name:$name -->[Temperature] 
(3)[n] X/D/N/A|P.. MediaRoom:Ambient:temperature:$unit -->[]
(4)[n] X/D/N/P/A MediaRoom:Ambient:temperature:$unit:$unit -->[°F] 
(3)[n] X/D/N/A|P.. MediaRoom:Ambient:temperature:temperature -->[72.86]
(1)[1k] X/D/?...... Bucket[1]=MediaRoom, enum Key=Presence Value=[] 
(2)[n] X/D/N.... MediaRoom:Presence:$name -->[] 
(3)[n] X/D/N/A|P.. MediaRoom:Presence:$name:$name -->[Motion Sensor]
(2)[n] X/D/N.... MediaRoom:Presence:$properties -->[] 
(3)[n] X/D/N/A|P.. MediaRoom:Presence:$properties:$properties -->[motion]
(2)[n] X/D/N.... MediaRoom:Presence:$type -->[] 
(3)[n] X/D/N/A|P.. MediaRoom:Presence:$type:$type -->[sensor]
(2)[n] X/D/N.... MediaRoom:Presence:motion -->[] 
(3)[n] X/D/N/A|P.. MediaRoom:Presence:motion:$datatype -->[]
(4)[n] X/D/N/P/A MediaRoom:Presence:motion:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. MediaRoom:Presence:motion:$format -->[]
(4)[n] X/D/N/P/A MediaRoom:Presence:motion:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. MediaRoom:Presence:motion:$name -->[]
(4)[n] X/D/N/P/A MediaRoom:Presence:motion:$name:$name -->[Motion] 
(3)[n] X/D/N/A|P.. MediaRoom:Presence:motion:motion -->[OPEN]
(1)[1k] X/D/?...... Bucket[1]=MediaRoom, enum Key=hardware Value=[] 
(2)[n] X/D/N.... MediaRoom:hardware:$name -->[] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:$name:$name -->[Device Info]
(2)[n] X/D/N.... MediaRoom:hardware:$properties -->[] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:$properties:$properties -->[signal,mac,resetReason,voltage]
(2)[n] X/D/N.... MediaRoom:hardware:$type -->[] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:$type:$type -->[sensor]
(2)[n] X/D/N.... MediaRoom:hardware:mac -->[] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:mac:$datatype -->[]
(4)[n] X/D/N/P/A MediaRoom:hardware:mac:$datatype:$datatype -->[sring] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:mac:$name -->[]
(4)[n] X/D/N/P/A MediaRoom:hardware:mac:$name:$name -->[mac] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:mac:mac -->[B4:E6:2D:15:50:3A]
(2)[n] X/D/N.... MediaRoom:hardware:resetReason -->[] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:resetReason:$datatype -->[]
(4)[n] X/D/N/P/A MediaRoom:hardware:resetReason:$datatype:$datatype -->[string] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:resetReason:$name -->[]
(4)[n] X/D/N/P/A MediaRoom:hardware:resetReason:$name:$name -->[Last Reset Reason] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:resetReason:resetReason -->[External System]
(2)[n] X/D/N.... MediaRoom:hardware:signal -->[] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:signal:$datatype -->[]
(4)[n] X/D/N/P/A MediaRoom:hardware:signal:$datatype:$datatype -->[integer] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:signal:$name -->[]
(4)[n] X/D/N/P/A MediaRoom:hardware:signal:$name:$name -->[RSSI] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:signal:$unit -->[]
(4)[n] X/D/N/P/A MediaRoom:hardware:signal:$unit:$unit -->[dBm] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:signal:signal -->[-67]
(2)[n] X/D/N.... MediaRoom:hardware:voltage -->[] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:voltage:$datatype -->[]
(4)[n] X/D/N/P/A MediaRoom:hardware:voltage:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:voltage:$name -->[]
(4)[n] X/D/N/P/A MediaRoom:hardware:voltage:$name:$name -->[3.3V Supply] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:voltage:$unit -->[]
(4)[n] X/D/N/P/A MediaRoom:hardware:voltage:$unit:$unit -->[V] 
(3)[n] X/D/N/A|P.. MediaRoom:hardware:voltage:voltage -->[3.12]
(0) X...... sknSensors 
(1) X/D...... sknSensors:OutsideMonitor 
(1)[1k] X/D/?...... Bucket[1]=OutsideMonitor, enum Key=$fw Value=[] 
(2)[a] X/D/A.... OutsideMonitor:$fw:checksum -->[] 
(3)[a] X/D/A/P.. OutsideMonitor:$fw:checksum:checksum -->[615fed382ab44bd43fe83508aecac682]
(2)[a] X/D/A.... OutsideMonitor:$fw:name -->[] 
(3)[a] X/D/A/P.. OutsideMonitor:$fw:name:name -->[Monitor-SHT31-RCWL-Metrics]
(2)[a] X/D/A.... OutsideMonitor:$fw:version -->[] 
(3)[a] X/D/A/P.. OutsideMonitor:$fw:version:version -->[2.0.0]
(1)[1k] X/D/?...... Bucket[1]=OutsideMonitor, enum Key=$homie Value=[] 
(2)[a] X/D/A.... OutsideMonitor:$homie:$homie -->[3.0.1] 
(1)[1k] X/D/?...... Bucket[1]=OutsideMonitor, enum Key=$implementation Value=[] 
(2)[a] X/D/A.... OutsideMonitor:$implementation:$implementation -->[esp8266] 
(2)[a] X/D/A.... OutsideMonitor:$implementation:config -->[] 
(3)[a] X/D/A/P.. OutsideMonitor:$implementation:config:config -->[{"name":"Outside Monitor","device_id":"OutsideMonitor","device_stats_interval":900,"wifi":{"ssid":"SFNSS1-24G"},"mqtt":{"host":"openhab.skoona.net","port":1883,"base_topic":"sknSensors/","auth":true},"ota":{"enabled":true},"settings":{"sensorsInterval":900,"motionHoldInterval":60}}]
(2)[a] X/D/A.... OutsideMonitor:$implementation:ota -->[] 
(3)[a] X/D/A/P.. OutsideMonitor:$implementation:ota:enabled -->[]
(4)[a] X/D/A/P/P OutsideMonitor:$implementation:ota:enabled:enabled -->[true] 
(2)[a] X/D/A.... OutsideMonitor:$implementation:version -->[] 
(3)[a] X/D/A/P.. OutsideMonitor:$implementation:version:version -->[3.0.0]
(1)[1k] X/D/?...... Bucket[1]=OutsideMonitor, enum Key=$localip Value=[] 
(2)[a] X/D/A.... OutsideMonitor:$localip:$localip -->[10.100.1.182] 
(1)[1k] X/D/?...... Bucket[1]=OutsideMonitor, enum Key=$mac Value=[] 
(2)[a] X/D/A.... OutsideMonitor:$mac:$mac -->[18:FE:34:FD:8C:1B] 
(1)[1k] X/D/?...... Bucket[1]=OutsideMonitor, enum Key=$name Value=[] 
(2)[a] X/D/A.... OutsideMonitor:$name:$name -->[Outside Monitor] 
(1)[1k] X/D/?...... Bucket[1]=OutsideMonitor, enum Key=$nodes Value=[] 
(2)[a] X/D/A.... OutsideMonitor:$nodes:$nodes -->[Ambient,Presence,hardware] 
(1)[1k] X/D/?...... Bucket[1]=OutsideMonitor, enum Key=$state Value=[] 
(2)[a] X/D/A.... OutsideMonitor:$state:$state -->[ready] 
(1)[1k] X/D/?...... Bucket[1]=OutsideMonitor, enum Key=$stats Value=[] 
(2)[a] X/D/A.... OutsideMonitor:$stats:$stats -->[uptime] 
(2)[a] X/D/A.... OutsideMonitor:$stats:interval -->[] 
(3)[a] X/D/A/P.. OutsideMonitor:$stats:interval:interval -->[905]
(2)[a] X/D/A.... OutsideMonitor:$stats:signal -->[] 
(3)[a] X/D/A/P.. OutsideMonitor:$stats:signal:signal -->[92]
(2)[a] X/D/A.... OutsideMonitor:$stats:uptime -->[] 
(3)[a] X/D/A/P.. OutsideMonitor:$stats:uptime:uptime -->[396023]
(1)[1k] X/D/?...... Bucket[1]=OutsideMonitor, enum Key=Ambient Value=[] 
(2)[n] X/D/N.... OutsideMonitor:Ambient:$name -->[] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Ambient:$name:$name -->[Temperature and Humidity Sensor]
(2)[n] X/D/N.... OutsideMonitor:Ambient:$properties -->[] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Ambient:$properties:$properties -->[humidity,temperature]
(2)[n] X/D/N.... OutsideMonitor:Ambient:$type -->[] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Ambient:$type:$type -->[sensor]
(2)[n] X/D/N.... OutsideMonitor:Ambient:humidity -->[] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Ambient:humidity:$datatype -->[]
(4)[n] X/D/N/P/A OutsideMonitor:Ambient:humidity:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Ambient:humidity:$name -->[]
(4)[n] X/D/N/P/A OutsideMonitor:Ambient:humidity:$name:$name -->[Humidity] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Ambient:humidity:$unit -->[]
(4)[n] X/D/N/P/A OutsideMonitor:Ambient:humidity:$unit:$unit -->[%rH] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Ambient:humidity:humidity -->[66.71]
(2)[n] X/D/N.... OutsideMonitor:Ambient:temperature -->[] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Ambient:temperature:$datatype -->[]
(4)[n] X/D/N/P/A OutsideMonitor:Ambient:temperature:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Ambient:temperature:$name -->[]
(4)[n] X/D/N/P/A OutsideMonitor:Ambient:temperature:$name:$name -->[Temperature] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Ambient:temperature:$unit -->[]
(4)[n] X/D/N/P/A OutsideMonitor:Ambient:temperature:$unit:$unit -->[°F] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Ambient:temperature:temperature -->[20.00]
(1)[1k] X/D/?...... Bucket[1]=OutsideMonitor, enum Key=Presence Value=[] 
(2)[n] X/D/N.... OutsideMonitor:Presence:$name -->[] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Presence:$name:$name -->[Motion Sensor]
(2)[n] X/D/N.... OutsideMonitor:Presence:$properties -->[] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Presence:$properties:$properties -->[motion]
(2)[n] X/D/N.... OutsideMonitor:Presence:$type -->[] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Presence:$type:$type -->[sensor]
(2)[n] X/D/N.... OutsideMonitor:Presence:motion -->[] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Presence:motion:$datatype -->[]
(4)[n] X/D/N/P/A OutsideMonitor:Presence:motion:$datatype:$datatype -->[enum] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Presence:motion:$format -->[]
(4)[n] X/D/N/P/A OutsideMonitor:Presence:motion:$format:$format -->[OPEN,CLOSED] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Presence:motion:$name -->[]
(4)[n] X/D/N/P/A OutsideMonitor:Presence:motion:$name:$name -->[Motion] 
(3)[n] X/D/N/A|P.. OutsideMonitor:Presence:motion:motion -->[OPEN]
(1)[1k] X/D/?...... Bucket[1]=OutsideMonitor, enum Key=hardware Value=[] 
(2)[n] X/D/N.... OutsideMonitor:hardware:$name -->[] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:$name:$name -->[Device Info]
(2)[n] X/D/N.... OutsideMonitor:hardware:$properties -->[] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:$properties:$properties -->[signal,mac,resetReason,voltage]
(2)[n] X/D/N.... OutsideMonitor:hardware:$type -->[] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:$type:$type -->[sensor]
(2)[n] X/D/N.... OutsideMonitor:hardware:mac -->[] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:mac:$datatype -->[]
(4)[n] X/D/N/P/A OutsideMonitor:hardware:mac:$datatype:$datatype -->[sring] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:mac:$name -->[]
(4)[n] X/D/N/P/A OutsideMonitor:hardware:mac:$name:$name -->[mac] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:mac:mac -->[18:FE:34:FD:8C:1B]
(2)[n] X/D/N.... OutsideMonitor:hardware:resetReason -->[] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:resetReason:$datatype -->[]
(4)[n] X/D/N/P/A OutsideMonitor:hardware:resetReason:$datatype:$datatype -->[string] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:resetReason:$name -->[]
(4)[n] X/D/N/P/A OutsideMonitor:hardware:resetReason:$name:$name -->[Last Reset Reason] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:resetReason:resetReason -->[External System]
(2)[n] X/D/N.... OutsideMonitor:hardware:signal -->[] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:signal:$datatype -->[]
(4)[n] X/D/N/P/A OutsideMonitor:hardware:signal:$datatype:$datatype -->[integer] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:signal:$name -->[]
(4)[n] X/D/N/P/A OutsideMonitor:hardware:signal:$name:$name -->[RSSI] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:signal:$unit -->[]
(4)[n] X/D/N/P/A OutsideMonitor:hardware:signal:$unit:$unit -->[dBm] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:signal:signal -->[-54]
(2)[n] X/D/N.... OutsideMonitor:hardware:voltage -->[] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:voltage:$datatype -->[]
(4)[n] X/D/N/P/A OutsideMonitor:hardware:voltage:$datatype:$datatype -->[float] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:voltage:$name -->[]
(4)[n] X/D/N/P/A OutsideMonitor:hardware:voltage:$name:$name -->[3.3V Supply] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:voltage:$unit -->[]
(4)[n] X/D/N/P/A OutsideMonitor:hardware:voltage:$unit:$unit -->[V] 
(3)[n] X/D/N/A|P.. OutsideMonitor:hardware:voltage:voltage -->[3.44]
level=info module=Homie-Service time:=2021-07-09T12:34:19.561511Z caller=service.go:142 pkg=deviceCore service=coreService event="restore networks" network=sknSensors
level=info module=Homie-Service time:=2021-07-09T12:34:19.561613Z caller=service.go:115 pkg=demoProvider service=StreamProvider event="calling trafficGenerator()" sourcefile=./dataDB/demoData/mosquitto.log
level=info module=Homie-Service time:=2021-07-09T12:34:19.56163Z caller=service.go:125 pkg=demoProvider service=StreamProvider event="trafficGenerator() active"
level=info module=Homie-Service time:=2021-07-09T12:34:19.566998Z caller=service.go:185 pkg=demoProvider service=StreamProvider event="demoRender() completed" limit=false
level=info module=Homie-Service time:=2021-07-09T12:34:23.562177Z caller=service.go:177 event="service completed"
module=Homie-Service time:=2021-07-09T12:34:23.683988Z caller=main.go:101 ui=base Loading="Main Page"
module=Homie-Service time:=2021-07-09T12:34:23.684953Z caller=networks.go:28 ui=base component=ViewProvider selectednetwork=sknSensors
module=Homie-Service time:=2021-07-09T12:34:23.684965Z caller=networks.go:36 ui=base component=ViewProvider completeddeviceloading=GarageMonitor
module=Homie-Service time:=2021-07-09T12:34:23.68497Z caller=networks.go:36 ui=base component=ViewProvider completeddeviceloading=MediaRoom
module=Homie-Service time:=2021-07-09T12:34:23.684974Z caller=networks.go:36 ui=base component=ViewProvider completeddeviceloading=MCPProvider
module=Homie-Service time:=2021-07-09T12:34:23.684978Z caller=networks.go:36 ui=base component=ViewProvider completeddeviceloading=OutsideMonitor
module=Homie-Service time:=2021-07-09T12:34:23.684981Z caller=networks.go:36 ui=base component=ViewProvider completeddeviceloading=D1R1MiniA
module=Homie-Service time:=2021-07-09T12:34:23.684985Z caller=networks.go:36 ui=base component=ViewProvider completeddeviceloading=FamilyRoom
module=Homie-Service time:=2021-07-09T12:34:23.684988Z caller=networks.go:36 ui=base component=ViewProvider completeddeviceloading=GarageProvider
module=Homie-Service time:=2021-07-09T12:34:23.684992Z caller=networks.go:36 ui=base component=ViewProvider completeddeviceloading=HomeOffice
module=Homie-Service time:=2021-07-09T12:34:23.684996Z caller=networks.go:36 ui=base component=ViewProvider completeddeviceloading=DSMonitor
module=Homie-Service time:=2021-07-09T12:34:23.685Z caller=networks.go:36 ui=base component=ViewProvider completeddeviceloading=GuestRoom
module=Homie-Service time:=2021-07-09T12:34:23.841101Z caller=views.go:118 ui=base component=ViewProvider tab=Home event="selector called" value=sknSensors status=
module=Homie-Service time:=2021-07-09T12:35:12.082304Z caller=views.go:101 ui=base component=ViewProvider tab=Sites event="refresh called"
level=info module=Homie-Service time:=2021-07-09T12:38:06.252733Z caller=main.go:107 ui=base event="shutdown requested" cause="Fyne GUI Shutdown"
level=info module=Homie-Service time:=2021-07-09T12:38:09.254405Z caller=service.go:166 event="service ended"
                     
```