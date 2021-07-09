
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
  "ID": "681d4ab0-1629-4501-9446-7269f8460850",
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
      "Received": "2021-07-09T08:25:51.970753-04:00"
    },
    {
      "ID": "8490e6cbf39eeadd991cb2a346580e51",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "alert",
      "Level": "door",
      "Value": "Door Open",
      "Received": "2021-07-09T08:25:51.970756-04:00"
    },
    {
      "ID": "c6c659175e0f9ca804664b3c6e3fda37",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "alert",
      "Level": "item",
      "Value": "Garage Door Open",
      "Received": "2021-07-09T08:25:51.970758-04:00"
    },
    {
      "ID": "df9b91ec197689e2c566aea2add68ad5",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "notice",
      "Level": "item",
      "Value": "openHAB3 Online",
      "Received": "2021-07-09T08:25:51.970761-04:00"
    },
    {
      "ID": "c254b24c700dbc2ecba237869f9b493b",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "$broadcast",
      "Level": "alert",
      "Value": "openHAB3 Offline",
      "Received": "2021-07-09T08:25:56.014997-04:00"
    },
    {
      "ID": "c254b24c700dbc2ecba237869f9b493b",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "$broadcast",
      "Level": "alert",
      "Value": "Door Open",
      "Received": "2021-07-09T08:25:56.052386-04:00"
    },
    {
      "ID": "80ea5ef8b530e884349d967e9c09c28a",
      "ElementType": 20,
      "Parent": "sknSensors",
      "Topic": "$broadcast",
      "Level": "notice",
      "Value": "openHAB3 Online",
      "Received": "2021-07-09T08:25:56.091997-04:00"
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
              "ID": "3ae33bd2-a0d3-43f9-8858-12e38f034916",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "1c169c9a-8d6d-49b0-afb7-c378db933fe8",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "6867842dfce4674d4c724085467362c9",
                  "Props": {}
                },
                "name": {
                  "ID": "a909dca1-0b9b-405f-a62f-071ace7bdc28",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "SknLiquids-128x32-oled",
                  "Props": {}
                },
                "version": {
                  "ID": "f36e1d3a-4f49-43a1-9d0c-07def5745fbd",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "1.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "b075babf-cce9-43ba-80cc-9d9c66463c7d",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "624fd7ce-bd86-4ea5-a506-24469d9f64e4",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$implementation",
              "Value": "esp8266",
              "Props": {
                "config": {
                  "ID": "3899e163-5467-4bb1-8884-977c9b352789",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"Furnace\",\"device_id\":\"D1R1MiniA\",\"device_stats_interval\":900,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "de49c7d3-84e1-41b0-8b8c-dc23d7faa9ea",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "enabled": {
                      "ID": "3e96880b-6fed-4c95-9538-8931baeda96d",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "321c7fa1-65f6-4ae6-9227-a94939c445ce",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "8dc443cc-0bf9-4eff-913b-47b6fcbb4798",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$localip",
              "Value": "10.100.1.181",
              "Props": {}
            },
            "$mac": {
              "ID": "8d3c1ae6-af20-447b-9eea-8eb182e50065",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$mac",
              "Value": "84:F3:EB:0C:38:6F",
              "Props": {}
            },
            "$name": {
              "ID": "539bb43d-fedc-457f-a3ce-5c431de5301b",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$name",
              "Value": "Furnace",
              "Props": {}
            },
            "$nodes": {
              "ID": "92552340-f90a-4fa4-833c-10781f9356c8",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$nodes",
              "Value": "Liquids",
              "Props": {}
            },
            "$state": {
              "ID": "a7a8dd43-0888-4894-9873-6cc275d9f07d",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "65c7be0c-2e5d-41d4-b4fa-19a561f82926",
              "ElementType": 11,
              "Parent": "D1R1MiniA",
              "Name": "$stats",
              "Value": "uptime",
              "Props": {
                "interval": {
                  "ID": "8178acbd-0008-41e2-b599-bd37c0d357ac",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "905",
                  "Props": {}
                },
                "signal": {
                  "ID": "f6813e61-45ab-461a-ac6e-4c9a5ffe8f3f",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "88",
                  "Props": {}
                },
                "uptime": {
                  "ID": "c85e0bbe-d66d-4319-8b98-55c7bea62e2e",
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
            "$name": {
              "ID": "097413a9-1f9c-4435-b47a-3853b254351f",
              "ElementType": 14,
              "Parent": "D1R1MiniA",
              "Name": "$name",
              "Attrs": {
                "Liquids": {
                  "ID": "73ef6ddc-6f90-4a22-bb7e-175f74f46ffa",
                  "ElementType": 15,
                  "Parent": "Liquids",
                  "Name": "Liquids",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$properties": {
              "ID": "2a39e970-869c-4f04-b8f3-3f05fa498f07",
              "ElementType": 14,
              "Parent": "D1R1MiniA",
              "Name": "$properties",
              "Attrs": {
                "Liquids": {
                  "ID": "7a85e19b-7cf3-44fd-9337-bbf9f8c78863",
                  "ElementType": 15,
                  "Parent": "Liquids",
                  "Name": "Liquids",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$type": {
              "ID": "f6d0daf4-abc0-4f7e-a11d-bfac6e5400c3",
              "ElementType": 14,
              "Parent": "D1R1MiniA",
              "Name": "$type",
              "Attrs": {
                "Liquids": {
                  "ID": "00b1ae9b-3a2a-4b12-a1cb-0ad3ce3f7a2a",
                  "ElementType": 15,
                  "Parent": "Liquids",
                  "Name": "Liquids",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "level": {
              "ID": "0e3d7e8d-8b39-47c9-8a9c-35476b7c8bf0",
              "ElementType": 14,
              "Parent": "D1R1MiniA",
              "Name": "level",
              "Attrs": {},
              "Props": {
                "level": {
                  "ID": "9e7cd4d3-4022-4ea4-8868-9c396c71346b",
                  "ElementType": 16,
                  "Parent": "Liquids",
                  "Name": "level",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "92b4767f-ca69-4925-a70b-a907a2d500c0",
                      "ElementType": 17,
                      "Parent": "level",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$format": {
                      "ID": "7ef8d39e-559f-42fb-ba48-eb3c5b899a29",
                      "ElementType": 17,
                      "Parent": "level",
                      "Name": "$format",
                      "Value": "0:1000"
                    },
                    "$name": {
                      "ID": "435cc2c7-0b8f-4d4b-98c6-5d54d7d88739",
                      "ElementType": 17,
                      "Parent": "level",
                      "Name": "$name",
                      "Value": "Level"
                    },
                    "$unit": {
                      "ID": "9a167699-00f2-45e4-b8f2-1e2ca6511dbe",
                      "ElementType": 17,
                      "Parent": "level",
                      "Name": "$unit",
                      "Value": "#"
                    }
                  }
                }
              }
            },
            "volts": {
              "ID": "befbedfe-175e-481b-aad7-4092a09b205b",
              "ElementType": 14,
              "Parent": "D1R1MiniA",
              "Name": "volts",
              "Attrs": {},
              "Props": {
                "volts": {
                  "ID": "b7dbc6cd-fa69-4fba-a519-41561f2795f4",
                  "ElementType": 16,
                  "Parent": "Liquids",
                  "Name": "volts",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "ca135876-b0e2-492f-834d-f952c4c88a94",
                      "ElementType": 17,
                      "Parent": "volts",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$format": {
                      "ID": "30828c7a-94d0-489b-bab8-d75bc435ef8c",
                      "ElementType": 17,
                      "Parent": "volts",
                      "Name": "$format",
                      "Value": "0:4"
                    },
                    "$name": {
                      "ID": "d543c342-c1d6-4542-93cc-1efc51c5b5a3",
                      "ElementType": 17,
                      "Parent": "volts",
                      "Name": "$name",
                      "Value": "Volts"
                    },
                    "$unit": {
                      "ID": "3a4f2a7e-bb35-4a9b-a2be-81471da9e867",
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
              "ID": "26c2f40c-3956-4588-af7b-2900b09bf3a5",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "87b71864-d7e3-44a4-8de9-488779d47b68",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "3ef8fbb48c5b23788a22a998c14a1a6d",
                  "Props": {}
                },
                "name": {
                  "ID": "c52b0710-ba50-46ae-99b0-1a99a580795c",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "Monitor-DHT-RCWL-Metrics",
                  "Props": {}
                },
                "version": {
                  "ID": "d7c558b5-191d-43ce-bb46-a02a700fbb31",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "2.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "03d132c9-91a0-4ab8-bf88-e82993c8c027",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "41be8d2b-ea98-44be-8722-0cf379d18ce5",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$implementation",
              "Value": "esp8266",
              "Props": {
                "config": {
                  "ID": "6a7b7517-fecd-4c7a-927f-c001c55add5b",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"FamilyRoom\",\"device_id\":\"FamilyRoom\",\"device_stats_interval\":180,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"sensorInterval\":900,\"motionHoldInterval\":60}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "6df1add2-6c53-4f91-9f99-69ae628d9d53",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "enabled": {
                      "ID": "3dbd0dbc-2f00-4a83-9928-1f2c92f75fa9",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "d6ec3b7a-9f04-41ec-b0df-34c23694d895",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "c65d38cc-9bcc-4ba8-ada0-674cd070eb79",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$localip",
              "Value": "10.100.1.166",
              "Props": {}
            },
            "$mac": {
              "ID": "dc09bb61-20c3-4ac8-a8a0-fad6b109f767",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$mac",
              "Value": "BC:DD:C2:E5:C4:24",
              "Props": {}
            },
            "$name": {
              "ID": "ea8e35d3-5569-4295-b23b-c59f1e5d5481",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$name",
              "Value": "FamilyRoom",
              "Props": {}
            },
            "$nodes": {
              "ID": "e23c22cc-ed95-4979-875e-4e06452652f4",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$nodes",
              "Value": "Ambient,Presence,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "e068e90c-e371-4463-87a4-d2553d1c41f1",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "d1f0e469-2a1b-4a3a-86f5-39e4147f7d52",
              "ElementType": 11,
              "Parent": "FamilyRoom",
              "Name": "$stats",
              "Value": "uptime",
              "Props": {
                "interval": {
                  "ID": "2f6d5aba-4e6e-4392-934e-d38a302f97b9",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "185",
                  "Props": {}
                },
                "signal": {
                  "ID": "31d2c7a7-ae15-47d7-9d8c-7f620080e913",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "78",
                  "Props": {}
                },
                "uptime": {
                  "ID": "7105351f-30c4-4c43-8d03-b5896a941d16",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "423639",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "$name": {
              "ID": "48893f5d-20d8-4dd5-a461-abf19636c338",
              "ElementType": 14,
              "Parent": "FamilyRoom",
              "Name": "$name",
              "Attrs": {
                "Ambient": {
                  "ID": "429b2e99-6746-4a8f-9086-06f61fa42b51",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "529b9902-2eaf-42b2-90fc-71819e4cdfc3",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "cc00ff2c-c6ed-46e0-a59d-b8428bf90568",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$properties": {
              "ID": "38b8ee7a-d71e-49c7-955c-3acabb2823f5",
              "ElementType": 14,
              "Parent": "FamilyRoom",
              "Name": "$properties",
              "Attrs": {
                "Ambient": {
                  "ID": "273e9718-8635-42f9-8d3a-6a95b8acf2a1",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "a471278c-db67-4686-858b-e4a73ba6ca2a",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "73843bdd-c1e0-4b15-b668-fcaf62ae8a55",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$type": {
              "ID": "69907992-f015-40fa-b1df-184bf3d3d35c",
              "ElementType": 14,
              "Parent": "FamilyRoom",
              "Name": "$type",
              "Attrs": {
                "Ambient": {
                  "ID": "88423ee2-82c8-42d9-a212-c66308cedef1",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "d995789e-6990-44d9-9458-4a027a269066",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "7f1f6aae-91ed-4b0d-9b74-5aa486ead59b",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "humidity": {
              "ID": "27776b95-7526-43d4-ae63-7d0baee8d442",
              "ElementType": 14,
              "Parent": "FamilyRoom",
              "Name": "humidity",
              "Attrs": {},
              "Props": {
                "humidity": {
                  "ID": "73324cf2-05f3-4342-8536-6e605321f7a5",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "humidity",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "323f1872-d744-4a6e-8478-bc83b3cd10e1",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "8021a33d-ca33-4225-a3b3-7de33ebf67b7",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$name",
                      "Value": "Humidity"
                    },
                    "$unit": {
                      "ID": "5aab74b6-45e3-485b-a119-e013b5acf017",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": "%rH"
                    }
                  }
                }
              }
            },
            "mac": {
              "ID": "1af1fb69-7ebe-4b54-9f90-4b1302bd9b60",
              "ElementType": 14,
              "Parent": "FamilyRoom",
              "Name": "mac",
              "Attrs": {},
              "Props": {
                "mac": {
                  "ID": "fad6d6a0-6a7c-4e47-89ae-5531f566451e",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "mac",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "7aed40c4-e762-4717-aa8d-daa2fa50ff46",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": "sring"
                    },
                    "$name": {
                      "ID": "e621c0af-c767-4f03-8127-c4f9239fc361",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$name",
                      "Value": "mac"
                    }
                  }
                }
              }
            },
            "motion": {
              "ID": "dd9ff59a-cb64-4d63-b644-021bf6c36309",
              "ElementType": 14,
              "Parent": "FamilyRoom",
              "Name": "motion",
              "Attrs": {},
              "Props": {
                "motion": {
                  "ID": "f12f7ea8-9fe5-4e3c-bdd0-67e16373dff9",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "motion",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "fba6ab9f-0939-4b61-bd49-0573ce108abe",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "1852ad9f-8f0c-43cd-9324-527df3e89f7a",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "23bf30a5-20af-49c1-8105-51defa49524a",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$name",
                      "Value": "Motion"
                    }
                  }
                }
              }
            },
            "resetReason": {
              "ID": "c5ca06a7-3548-416c-a2b9-ba1118051044",
              "ElementType": 14,
              "Parent": "FamilyRoom",
              "Name": "resetReason",
              "Attrs": {},
              "Props": {
                "resetReason": {
                  "ID": "63b4084b-25c0-4fc7-8c88-057720e9c3b1",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "resetReason",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "6af6af4e-4c16-4ba8-9263-247555fc7332",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$name": {
                      "ID": "ef7b54c9-ce8b-4a14-b482-9106713ad698",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$name",
                      "Value": "Last Reset Reason"
                    }
                  }
                }
              }
            },
            "signal": {
              "ID": "b030c5bb-124b-45e4-ba79-411a1f1b5c60",
              "ElementType": 14,
              "Parent": "FamilyRoom",
              "Name": "signal",
              "Attrs": {},
              "Props": {
                "signal": {
                  "ID": "8c84ccd2-1cfb-427d-9192-ec2356dbf950",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "signal",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "44b56c7f-8b77-413a-a7f0-411eb634f96b",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$name": {
                      "ID": "0fc02240-e1e5-4f8d-a5a5-7d41ea9a20ad",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$name",
                      "Value": "RSSI"
                    },
                    "$unit": {
                      "ID": "499d19c8-0137-4c94-b6dc-a77a84a9fbed",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": "dBm"
                    }
                  }
                }
              }
            },
            "temperature": {
              "ID": "9995ca00-5e94-4672-94b8-89e0046f999b",
              "ElementType": 14,
              "Parent": "FamilyRoom",
              "Name": "temperature",
              "Attrs": {},
              "Props": {
                "temperature": {
                  "ID": "a52664d3-6b77-470a-8fae-c06d2efbb9ef",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "temperature",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "827168a6-cef9-48c7-8219-02efa8be25bb",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "0210c06a-a6a7-4785-aa86-3fc98c6f2796",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$name",
                      "Value": "Temperature"
                    },
                    "$unit": {
                      "ID": "5ddb14df-26f0-46ab-b669-3f03c2c09868",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$unit",
                      "Value": "°F"
                    }
                  }
                }
              }
            },
            "voltage": {
              "ID": "b03d84ec-15b9-4e8e-a016-adc6882a5517",
              "ElementType": 14,
              "Parent": "FamilyRoom",
              "Name": "voltage",
              "Attrs": {},
              "Props": {
                "voltage": {
                  "ID": "bb638a7f-7ed5-4834-9c35-656abfcf29c8",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "voltage",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "de9fad23-2b19-498d-955a-2af14c017fab",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "848bab7e-ab1e-4f4b-ae0f-5d00eaa9ead6",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$name",
                      "Value": "3.3V Supply"
                    },
                    "$unit": {
                      "ID": "be34869e-31dc-48e4-a058-071562204cad",
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
              "ID": "c1f61acf-9756-4e53-9cd5-4065cbe55e54",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "6e8aaec4-8623-4d38-9d8a-744e8c3f6088",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "615fed382ab44bd43fe83508aecac682",
                  "Props": {}
                },
                "name": {
                  "ID": "ba31a77d-dac1-41da-8a60-e982bb9f7c71",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "Monitor-SHT31-RCWL-Metrics",
                  "Props": {}
                },
                "version": {
                  "ID": "5362c8c1-899d-4497-bed2-ca5d6fc4f241",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "2.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "ee929f0a-d9b8-4f97-bdf8-781f2d7f5c32",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "eb96855e-89e9-4a74-9af0-66927feb41b5",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$implementation",
              "Value": "esp8266",
              "Props": {
                "config": {
                  "ID": "807798f6-410c-469c-b2b9-b4cfe518a151",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"Garage Monitor\",\"device_id\":\"GarageMonitor\",\"device_stats_interval\":900,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"sensorInterval\":900,\"motionHoldInterval\":60}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "d85b49bb-8c7c-4ffa-8afe-4e6498498f64",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "enabled": {
                      "ID": "18d4fb46-fd27-46d7-9415-622059b9bf82",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "627a4bc8-1800-4650-b514-596db01fcc36",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "0090a9b8-76db-47f3-ad5c-51173618bca6",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$localip",
              "Value": "10.100.1.177",
              "Props": {}
            },
            "$mac": {
              "ID": "0bc9c042-3898-4183-9396-2a05298bdf2d",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$mac",
              "Value": "B4:E6:2D:1B:5C:4D",
              "Props": {}
            },
            "$name": {
              "ID": "5ef74b6f-ce3c-4cee-91b7-ea9189c47c95",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$name",
              "Value": "Garage Monitor",
              "Props": {}
            },
            "$nodes": {
              "ID": "e0fdd7e5-43a8-4263-8aee-2066aa3bda15",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$nodes",
              "Value": "Ambient,Presence,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "ba9bf506-2fbf-438d-a6bb-ff588212be4a",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "b22bf3df-ec5a-4f2f-938d-39e55252e866",
              "ElementType": 11,
              "Parent": "GarageMonitor",
              "Name": "$stats",
              "Value": "uptime",
              "Props": {
                "interval": {
                  "ID": "afc53958-4999-4774-beb5-8536f7a9c414",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "905",
                  "Props": {}
                },
                "signal": {
                  "ID": "7e8997bf-2966-4e03-a323-5e1e15709f7e",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "54",
                  "Props": {}
                },
                "uptime": {
                  "ID": "2d762788-7799-48a8-abc7-d7c6f29e8b7f",
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
            "$name": {
              "ID": "a3f05453-1d04-426f-ad77-7258a56ea753",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "$name",
              "Attrs": {
                "Ambient": {
                  "ID": "2614dbdb-4a05-4249-a870-fd2c3704ac0d",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "3b7ce033-3ac7-471b-9a52-63049291e268",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "0fe35746-5778-4517-8668-a1ea21374a10",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$properties": {
              "ID": "f609b867-5e90-4a38-9b4b-237a260b8d8f",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "$properties",
              "Attrs": {
                "Ambient": {
                  "ID": "570e907b-ef22-4257-b8ff-6f2c65991aec",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "c734451a-b457-40c9-b88f-548aae274f5b",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "44edffa6-04eb-4e93-bae9-4f512924d571",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$type": {
              "ID": "6fa0f9ba-ac5c-4c39-bfbb-b3e308c4d4e3",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "$type",
              "Attrs": {
                "Ambient": {
                  "ID": "b4fe336f-528e-4779-a711-a9ff811881fc",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "5dbea210-2dc9-4c34-b1b5-c6b9e8c599a6",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "4ba21cc4-41f8-4f04-8b4d-83ca30cfa6dd",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "Ambient": {
              "ID": "d5c41768-b41f-460b-9b9a-a1267870aa6c",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "Ambient",
              "Attrs": {
                "$name": {
                  "ID": "16e22355-380c-4c20-a6ae-14c51724997f",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$name",
                  "Value": "Temperature and Humidity Sensor"
                },
                "$properties": {
                  "ID": "fce1458f-6bb1-4cbb-9a40-d38acb596bd4",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$properties",
                  "Value": "humidity,temperature"
                },
                "$type": {
                  "ID": "216d1ffb-7afe-4d3a-b543-0cb77921243f",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "humidity": {
                  "ID": "86ac2d40-5a23-427b-b024-07edf3c0ef7b",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "humidity",
                  "Value": "57.18",
                  "Attrs": {
                    "$datatype": {
                      "ID": "92e25342-f6cf-46f9-bf42-206fca2d42f9",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "99baa084-3727-44bb-8621-e120c98bdc13",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$name",
                      "Value": "Humidity"
                    },
                    "$unit": {
                      "ID": "f3954df9-3463-4bdc-a1ae-f5c965ac1a54",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": "%rH"
                    }
                  }
                },
                "temperature": {
                  "ID": "e6f85d0b-60a4-450c-b910-dc554a8c6fa8",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "temperature",
                  "Value": "26.18",
                  "Attrs": {
                    "$datatype": {
                      "ID": "5776b80b-155b-4cb7-9d2c-3d85b57b8bc9",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "9fb16609-3aa1-41f1-9b5e-43eb3bf95e14",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$name",
                      "Value": "Temperature"
                    },
                    "$unit": {
                      "ID": "538f88bf-f010-45f8-8125-50c3e70cf2e0",
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
              "ID": "ff58947d-7b62-4846-a51e-fde3731097da",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "Presence",
              "Attrs": {
                "$name": {
                  "ID": "8aae11d8-1d75-4e46-bfbb-e21a045a755e",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$name",
                  "Value": "Motion Sensor"
                },
                "$properties": {
                  "ID": "0e01d233-a588-440b-9b68-382d333b369b",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$properties",
                  "Value": "motion"
                },
                "$type": {
                  "ID": "18cec021-95c6-4b4c-af1c-0308f22142f8",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "motion": {
                  "ID": "f937ad95-3b25-4d10-8042-a8c78b133a5c",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "motion",
                  "Value": "OPEN",
                  "Attrs": {
                    "$datatype": {
                      "ID": "ee8b7c05-8110-46db-a62f-697018cd5422",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "2d84bbbf-ea1b-4667-8760-07c8a9519771",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "f1231a3e-e014-4458-9db6-86e368a8d9cc",
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
              "ID": "f8f7914d-c1dd-4533-a6f6-79801e712c35",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "hardware",
              "Attrs": {
                "$name": {
                  "ID": "bb88e738-3b4b-4f15-80b8-b8597ae9e399",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$name",
                  "Value": "Device Info"
                },
                "$properties": {
                  "ID": "c61b9cf0-83ae-4e88-8b04-30cdc8ed9b5b",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$properties",
                  "Value": "signal,mac,resetReason,voltage"
                },
                "$type": {
                  "ID": "e19e7ef0-8a22-4860-87ef-a584dbc93cf5",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "mac": {
                  "ID": "2ec863c7-9f38-4d18-93e1-9391df739811",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "mac",
                  "Value": "B4:E6:2D:1B:5C:4D",
                  "Attrs": {
                    "$datatype": {
                      "ID": "b6fe9064-30c0-408b-bbde-8e4fca6d486c",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": "sring"
                    },
                    "$name": {
                      "ID": "2c01a6ec-72ff-4ba7-aee2-18884231cb62",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$name",
                      "Value": "mac"
                    }
                  }
                },
                "resetReason": {
                  "ID": "a5d947fe-dff5-4483-8bc6-e2f2a72777b0",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "resetReason",
                  "Value": "External System",
                  "Attrs": {
                    "$datatype": {
                      "ID": "681e95e0-4108-4ef0-b7a2-4baca354e345",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$name": {
                      "ID": "a2f0bf60-df8d-4cf8-90dc-154ec95bc4e9",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$name",
                      "Value": "Last Reset Reason"
                    }
                  }
                },
                "signal": {
                  "ID": "a5bb7720-6816-4b7b-b068-1f338dcd542f",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "signal",
                  "Value": "-73",
                  "Attrs": {
                    "$datatype": {
                      "ID": "a5206d99-4d74-405d-85e9-4d74125bb9b3",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$name": {
                      "ID": "b7b8ebd6-fcb0-4f34-9be8-f15fbaa47b9f",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$name",
                      "Value": "RSSI"
                    },
                    "$unit": {
                      "ID": "0e5a5b36-a269-424f-b242-21a08884a41e",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": "dBm"
                    }
                  }
                },
                "voltage": {
                  "ID": "a67938b7-bb14-4711-bad4-b902fb7c8b9e",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "voltage",
                  "Value": "3.05",
                  "Attrs": {
                    "$datatype": {
                      "ID": "15f962f8-e1f8-44d1-b342-ba11a7ede2a9",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "be539d10-e77a-4b4a-b209-429828d3998e",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$name",
                      "Value": "3.3V Supply"
                    },
                    "$unit": {
                      "ID": "7a571beb-2eb0-458c-acb6-ecd040c50416",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$unit",
                      "Value": "V"
                    }
                  }
                }
              }
            },
            "humidity": {
              "ID": "f6e4f6f0-eb96-4e3c-930c-cfda41fb15f2",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "humidity",
              "Attrs": {},
              "Props": {
                "humidity": {
                  "ID": "15d696fb-d138-4ae8-b9f4-a9e9698bf0d2",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "humidity",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "b9ec34f5-e2c1-47d1-a2f7-9039b412dbe7",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "1eced269-f83b-4325-a7b4-a3e35934cff6",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$name",
                      "Value": "Humidity"
                    },
                    "$unit": {
                      "ID": "be0c4d69-87fa-47db-b687-7d2d4c8ad923",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": "%rH"
                    }
                  }
                }
              }
            },
            "mac": {
              "ID": "77169d75-d21a-4922-882e-bb05d3dc7295",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "mac",
              "Attrs": {},
              "Props": {
                "mac": {
                  "ID": "33634ddb-4d55-4b1f-bfee-7af0345eae2d",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "mac",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "12a06ccc-3989-4c93-ac8a-da5011829bdb",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": "sring"
                    },
                    "$name": {
                      "ID": "42ac9836-b46a-4fcf-a981-648100e6803b",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$name",
                      "Value": "mac"
                    }
                  }
                }
              }
            },
            "motion": {
              "ID": "f5cb2b8a-3761-4bb5-9f97-6c191a8d9941",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "motion",
              "Attrs": {},
              "Props": {
                "motion": {
                  "ID": "0ced5bd6-3643-41fd-a533-f77b21ad6019",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "motion",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "36314c50-f972-4798-b219-1a07d6c7d982",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "a8baed50-8715-446d-9e06-09c25e953cf3",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "2083f50a-6948-4a05-bd19-33d3d6b05dd8",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$name",
                      "Value": "Motion"
                    }
                  }
                }
              }
            },
            "resetReason": {
              "ID": "dae0aea2-d4c9-4e38-957c-c22e0dc4f1d5",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "resetReason",
              "Attrs": {},
              "Props": {
                "resetReason": {
                  "ID": "ff925e99-3827-4c1a-b15a-c9ccf6c2cfd1",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "resetReason",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "f9ad586d-6919-413c-b9fd-4d8af2135b31",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$name": {
                      "ID": "975b3a85-b435-4edb-b40e-fad10950e1b2",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$name",
                      "Value": "Last Reset Reason"
                    }
                  }
                }
              }
            },
            "signal": {
              "ID": "e798ca85-ab99-4902-93d3-f371574bf879",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "signal",
              "Attrs": {},
              "Props": {
                "signal": {
                  "ID": "3473324a-69ba-4b13-9fb9-ca197c96ee43",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "signal",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "6ba64498-380b-4f3f-8d3c-79bb941a6d62",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$name": {
                      "ID": "2b30e144-2fff-4aee-a52d-ddc106cd0491",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$name",
                      "Value": "RSSI"
                    },
                    "$unit": {
                      "ID": "fe8a05b8-0e3a-49e0-ab38-bf3c9ef7d769",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": "dBm"
                    }
                  }
                }
              }
            },
            "temperature": {
              "ID": "59a17ff1-acb2-48d0-8045-1f115a02a9c2",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "temperature",
              "Attrs": {},
              "Props": {
                "temperature": {
                  "ID": "9717ec63-d9af-4626-9735-b132c394e7ca",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "temperature",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "9f6cecc3-dd5f-4a1c-939b-de1cc4cb04ef",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "03a99179-3f0e-4dec-b0a4-b6dac5891bec",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$name",
                      "Value": "Temperature"
                    },
                    "$unit": {
                      "ID": "de6bd467-56b9-418f-b710-8990044fe3a7",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$unit",
                      "Value": "°F"
                    }
                  }
                }
              }
            },
            "voltage": {
              "ID": "42974d37-9f80-4e57-8b46-8098e4bebe84",
              "ElementType": 14,
              "Parent": "GarageMonitor",
              "Name": "voltage",
              "Attrs": {},
              "Props": {
                "voltage": {
                  "ID": "6ecc44ba-e059-4c93-acef-cdfd51682569",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "voltage",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "023e4995-6771-497c-8ea0-f7737ebb2248",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "71caa279-ed03-4835-8091-f648bd96570b",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$name",
                      "Value": "3.3V Supply"
                    },
                    "$unit": {
                      "ID": "9a0ad92d-59eb-46b8-bd8c-35c34bf1ae75",
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
              "ID": "81645ec3-3c0d-4c42-b5c5-901ffbf04b92",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "f5e6fae0-f4d8-4d84-8058-92cb4e3ee475",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "39286ba639a9ef1d5395e1c0e45d13fa",
                  "Props": {}
                },
                "name": {
                  "ID": "8b25bd85-83b1-4aaf-ac15-f5fd21ff78ee",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "GarageProvider",
                  "Props": {}
                },
                "version": {
                  "ID": "416f956e-4399-40fb-9d48-9f8abbd0f397",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "1.1.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "0f9f6855-20a7-482c-ba89-5716b80e9c4f",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "c68c3008-a41b-4835-bfd9-91cfcf2d7f41",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$implementation",
              "Value": "esp32",
              "Props": {
                "config": {
                  "ID": "456aa56a-0951-43d4-98a3-4757fd84a1af",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"Garage Provider\",\"device_id\":\"GarageProvider\",\"device_stats_interval\":900,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"sensorsInterval\":900}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "a537c68a-5cbd-4376-8e99-e33322f60788",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "enabled": {
                      "ID": "8e816f73-0b8d-40e3-a258-333e93f63e6c",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "d227c038-9ac5-498d-b5a6-b8b56c1171f4",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "44e8b7c9-ef9d-4728-9f5f-2dea60428cb6",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$localip",
              "Value": "10.100.1.218",
              "Props": {}
            },
            "$mac": {
              "ID": "6cc42f28-9ea9-4ab7-9d0c-40c98b9f965a",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$mac",
              "Value": "24:6F:28:97:63:B8",
              "Props": {}
            },
            "$name": {
              "ID": "c7ea31dd-ccb8-4c30-89da-5c61d905c267",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$name",
              "Value": "Garage Provider",
              "Props": {}
            },
            "$nodes": {
              "ID": "491973c8-41fc-474d-a8cc-656e0fea40fc",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$nodes",
              "Value": "garageDoor,environmentMonitor",
              "Props": {}
            },
            "$state": {
              "ID": "72db9212-fa1c-4b5f-8221-1f440e4e5cde",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "b1a302e1-3f84-48ea-8ddf-515216be4dcf",
              "ElementType": 11,
              "Parent": "GarageProvider",
              "Name": "$stats",
              "Value": "uptime",
              "Props": {
                "interval": {
                  "ID": "6b9595d1-0ec8-483b-b19a-4848c02a95ac",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "905",
                  "Props": {}
                },
                "signal": {
                  "ID": "80416dc0-e13b-44df-995d-a06ddb0f0cb4",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "72",
                  "Props": {}
                },
                "uptime": {
                  "ID": "01abfbf9-807f-4990-93e9-ae410376402b",
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
            "$name": {
              "ID": "9745e9fc-78a8-49ce-b85b-4a436898b42d",
              "ElementType": 14,
              "Parent": "GarageProvider",
              "Name": "$name",
              "Attrs": {
                "environmentMonitor": {
                  "ID": "8ca9501e-0447-451c-845e-ff937bf6619d",
                  "ElementType": 15,
                  "Parent": "environmentMonitor",
                  "Name": "environmentMonitor",
                  "Value": ""
                },
                "garageDoor": {
                  "ID": "6329ce21-9aef-491f-bb38-47337985659a",
                  "ElementType": 15,
                  "Parent": "garageDoor",
                  "Name": "garageDoor",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$properties": {
              "ID": "086d6bd8-5cca-453c-8473-dcb500423b7b",
              "ElementType": 14,
              "Parent": "GarageProvider",
              "Name": "$properties",
              "Attrs": {
                "environmentMonitor": {
                  "ID": "7984ee73-0994-4a4a-9129-dc6a9bafc8c1",
                  "ElementType": 15,
                  "Parent": "environmentMonitor",
                  "Name": "environmentMonitor",
                  "Value": ""
                },
                "garageDoor": {
                  "ID": "08731525-3d32-41c7-aa5b-a0f39a343730",
                  "ElementType": 15,
                  "Parent": "garageDoor",
                  "Name": "garageDoor",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$type": {
              "ID": "61b853bd-2040-4a39-9f07-f48fea5cb1f3",
              "ElementType": 14,
              "Parent": "GarageProvider",
              "Name": "$type",
              "Attrs": {
                "environmentMonitor": {
                  "ID": "89cfd95d-e232-4be5-9ccd-f834e2bd9db2",
                  "ElementType": 15,
                  "Parent": "environmentMonitor",
                  "Name": "environmentMonitor",
                  "Value": ""
                },
                "garageDoor": {
                  "ID": "d37edcbd-bdb6-4b22-ba2d-808e84e14def",
                  "ElementType": 15,
                  "Parent": "garageDoor",
                  "Name": "garageDoor",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "direction": {
              "ID": "6492f6b1-dfb1-451e-9c63-465fcaec698a",
              "ElementType": 14,
              "Parent": "GarageProvider",
              "Name": "direction",
              "Attrs": {},
              "Props": {
                "direction": {
                  "ID": "7e311323-ce32-4e58-b4c6-09de27ba1d4d",
                  "ElementType": 16,
                  "Parent": "garageDoor",
                  "Name": "direction",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "31093d78-6afb-44a5-aeba-d5bce5c7c3d5",
                      "ElementType": 17,
                      "Parent": "direction",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$format": {
                      "ID": "25de39ae-bfa9-44ca-ba3b-c4b38d1f6f83",
                      "ElementType": 17,
                      "Parent": "direction",
                      "Name": "$format",
                      "Value": "%s"
                    },
                    "$name": {
                      "ID": "d3876407-8998-400b-afd0-0fbd9c20c6a3",
                      "ElementType": 17,
                      "Parent": "direction",
                      "Name": "$name",
                      "Value": "Travel Direction"
                    }
                  }
                }
              }
            },
            "humdity": {
              "ID": "07dd76f5-51b2-4237-9cc2-22bff5feceb3",
              "ElementType": 14,
              "Parent": "GarageProvider",
              "Name": "humdity",
              "Attrs": {},
              "Props": {
                "humdity": {
                  "ID": "7b6b70e3-4788-4252-9644-d771aed5a816",
                  "ElementType": 16,
                  "Parent": "environmentMonitor",
                  "Name": "humdity",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "7129e559-4854-4f95-b374-77e4528e6695",
                      "ElementType": 17,
                      "Parent": "humdity",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$format": {
                      "ID": "5e8eac3c-d446-4c00-ab27-3d9fba3843f9",
                      "ElementType": 17,
                      "Parent": "humdity",
                      "Name": "$format",
                      "Value": "%.1f"
                    },
                    "$name": {
                      "ID": "88714a12-1f99-4dc9-8cce-14b305ebbfa5",
                      "ElementType": 17,
                      "Parent": "humdity",
                      "Name": "$name",
                      "Value": "Humidity"
                    },
                    "$unit": {
                      "ID": "c4536a05-5014-40fc-8748-a113da88dc63",
                      "ElementType": 17,
                      "Parent": "humdity",
                      "Name": "$unit",
                      "Value": "%"
                    }
                  }
                }
              }
            },
            "motion": {
              "ID": "b54921b1-338c-49f6-9fa5-789cdb511c92",
              "ElementType": 14,
              "Parent": "GarageProvider",
              "Name": "motion",
              "Attrs": {},
              "Props": {
                "motion": {
                  "ID": "d5e09993-0472-47ca-97df-985f006ee6d6",
                  "ElementType": 16,
                  "Parent": "environmentMonitor",
                  "Name": "motion",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "4c1745a9-dbe5-4349-a65b-8d65b462c433",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$format": {
                      "ID": "18463cc2-60e8-44c6-add2-eac36987c3da",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": "%s"
                    },
                    "$name": {
                      "ID": "d571f0ef-05c5-4ca2-932f-521981f4c484",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$name",
                      "Value": "Motion"
                    }
                  }
                }
              }
            },
            "operator": {
              "ID": "009a49f5-eb51-4adc-899c-8e1bae6b76ab",
              "ElementType": 14,
              "Parent": "GarageProvider",
              "Name": "operator",
              "Attrs": {},
              "Props": {
                "operator": {
                  "ID": "27f8ab98-1cc8-4e53-a35c-14441c39b32e",
                  "ElementType": 16,
                  "Parent": "garageDoor",
                  "Name": "operator",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "1e523081-7fe6-4066-ae8d-2c0c38299835",
                      "ElementType": 17,
                      "Parent": "operator",
                      "Name": "$datatype",
                      "Value": "boolean"
                    },
                    "$format": {
                      "ID": "beb6ed74-5f90-4e64-83ad-52665105f3e8",
                      "ElementType": 17,
                      "Parent": "operator",
                      "Name": "$format",
                      "Value": "%s"
                    },
                    "$name": {
                      "ID": "cfcfacb1-83f5-47d5-823b-c03d328d4941",
                      "ElementType": 17,
                      "Parent": "operator",
                      "Name": "$name",
                      "Value": "Operator"
                    },
                    "$settable": {
                      "ID": "7ae1ab4b-7fa0-4ab8-99a4-4a08874f39cb",
                      "ElementType": 17,
                      "Parent": "operator",
                      "Name": "$settable",
                      "Value": "true"
                    }
                  }
                }
              }
            },
            "positon": {
              "ID": "7562e25f-1fdc-4d4d-833e-6a135b657fa0",
              "ElementType": 14,
              "Parent": "GarageProvider",
              "Name": "positon",
              "Attrs": {},
              "Props": {
                "positon": {
                  "ID": "9722ff3e-fa89-4abf-8ffe-e0a0a538906f",
                  "ElementType": 16,
                  "Parent": "garageDoor",
                  "Name": "positon",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "6f7b8139-0cf0-4203-92a9-b7896b2ff37b",
                      "ElementType": 17,
                      "Parent": "positon",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$format": {
                      "ID": "251eb25f-bfd2-42d4-afd6-d9c844271c18",
                      "ElementType": 17,
                      "Parent": "positon",
                      "Name": "$format",
                      "Value": "%d"
                    },
                    "$name": {
                      "ID": "eba44e44-78e7-4a53-b665-896c3e584d96",
                      "ElementType": 17,
                      "Parent": "positon",
                      "Name": "$name",
                      "Value": "Position MM"
                    },
                    "$unit": {
                      "ID": "84f68637-9bfe-4fba-a5d1-88e5c939e5d7",
                      "ElementType": 17,
                      "Parent": "positon",
                      "Name": "$unit",
                      "Value": "mm"
                    }
                  }
                }
              }
            },
            "temperature": {
              "ID": "047ab8e6-d75a-40db-b77b-17543ce3d6c6",
              "ElementType": 14,
              "Parent": "GarageProvider",
              "Name": "temperature",
              "Attrs": {},
              "Props": {
                "temperature": {
                  "ID": "41d7eeaa-7539-4746-a1d0-4302bbd31144",
                  "ElementType": 16,
                  "Parent": "environmentMonitor",
                  "Name": "temperature",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "343c7672-5867-4e52-bd05-578fb3a26ac7",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$format": {
                      "ID": "e125dad5-9f86-417f-97e3-0352c9cbf4bf",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$format",
                      "Value": "%.1f"
                    },
                    "$name": {
                      "ID": "5353265d-6885-450c-8b81-fd980d875c0a",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$name",
                      "Value": "Temperature"
                    },
                    "$unit": {
                      "ID": "fba15074-dc49-4719-b78a-be51324a4de7",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$unit",
                      "Value": "ºF"
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
              "ID": "48919bf3-22fe-48fb-a832-2a8a018bb93a",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "64b2d8c3-1b59-4cfa-bfbf-5ef0568c68f9",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "3ef8fbb48c5b23788a22a998c14a1a6d",
                  "Props": {}
                },
                "name": {
                  "ID": "653d763d-23c7-47c9-806c-2451e6684953",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "Monitor-DHT-RCWL-Metrics",
                  "Props": {}
                },
                "version": {
                  "ID": "bf0a5c35-78ff-4444-8a05-cb000d624f7c",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "2.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "833368f7-e778-4783-86b2-9458e767d7ab",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "50e401f7-6409-41cf-9cc6-9b8090505ab5",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$implementation",
              "Value": "esp8266",
              "Props": {
                "config": {
                  "ID": "9d346890-13ec-4a57-a34c-825884ce4306",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"Guest Room\",\"device_id\":\"GuestRoom\",\"device_stats_interval\":180,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"sensorInterval\":900,\"motionHoldInterval\":60}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "7e1ea46f-8e5d-4fd7-a5db-130c3f35c015",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "enabled": {
                      "ID": "4ce2963d-d2cb-457e-b680-34feefcc8dd9",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "2268479e-c303-437c-9774-48b0d0e1d20f",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "e455ce89-5568-4b14-b885-2dd8a1346c77",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$localip",
              "Value": "10.100.1.178",
              "Props": {}
            },
            "$mac": {
              "ID": "54b38117-5b7d-437e-aef0-1777939124c8",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$mac",
              "Value": "BC:DD:C2:81:04:72",
              "Props": {}
            },
            "$name": {
              "ID": "73173608-c9a4-4276-9012-22f235f4364f",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$name",
              "Value": "Guest Room",
              "Props": {}
            },
            "$nodes": {
              "ID": "71e92a26-a15d-4f4f-87c9-98169a20a2ba",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$nodes",
              "Value": "Ambient,Presence,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "ed01a7ff-a421-458c-a957-0590532e7506",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "7c97ad5e-498a-486e-a269-c9cde7e8df7b",
              "ElementType": 11,
              "Parent": "GuestRoom",
              "Name": "$stats",
              "Value": "uptime",
              "Props": {
                "interval": {
                  "ID": "b4efef9d-08c3-4431-9a91-6fee30e70466",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "185",
                  "Props": {}
                },
                "signal": {
                  "ID": "1fb9107c-c826-48c9-944f-12e968446fa4",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "72",
                  "Props": {}
                },
                "uptime": {
                  "ID": "8f97844c-f787-4602-a734-064cb02dc9bc",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "uptime",
                  "Value": "362178",
                  "Props": {}
                }
              }
            }
          },
          "Nodes": {
            "$name": {
              "ID": "61da6e48-a4e8-4740-a0d4-feac61fb7e16",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "$name",
              "Attrs": {
                "Ambient": {
                  "ID": "c9631c39-d7d8-4bbe-9e72-4f2469e6cef7",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "86c41c7f-94cd-412c-abe2-4fd69095a7ae",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "36776cc3-38af-4f3a-81be-4d3346965eef",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$properties": {
              "ID": "140cb0d4-69a9-4c21-bc6a-98911607a784",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "$properties",
              "Attrs": {
                "Ambient": {
                  "ID": "0d12408f-ce82-4046-852e-21384bb661f2",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "d57734d7-4514-4a42-b952-b44d10306c53",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "a053a191-5eaf-40ec-87d8-1147a6872f33",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$type": {
              "ID": "0f67a58d-cfab-4565-a6ed-24f58c6356ff",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "$type",
              "Attrs": {
                "Ambient": {
                  "ID": "bfcc4931-e381-458f-9a08-675ded319b31",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "96b69005-60cd-4038-8beb-faa7b7b0fc51",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "d5181d5d-c3db-4f84-a80c-bc0e7ce20272",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "Ambient": {
              "ID": "ac4c3ff7-e290-42de-914f-2888c6a02f3f",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "Ambient",
              "Attrs": {
                "$name": {
                  "ID": "1dc6cb5b-0525-4bda-bd0f-a2e6e4458f9d",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$name",
                  "Value": "Temperature and Humidity Sensor"
                },
                "$properties": {
                  "ID": "25eb42dd-e6c9-4d96-bc48-0a8d8c52406a",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$properties",
                  "Value": "humidity,temperature"
                },
                "$type": {
                  "ID": "436c2aad-6d33-4d3f-aa08-09b75789b91a",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "humidity": {
                  "ID": "273f69d5-2f38-4151-91db-c34582daca8d",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "humidity",
                  "Value": "11.10",
                  "Attrs": {
                    "$datatype": {
                      "ID": "30caf034-0473-4a15-98bd-affcc9bbe8fa",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "1ddad8ef-6777-41db-8f13-938d3cf4b51e",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$name",
                      "Value": "Humidity"
                    },
                    "$unit": {
                      "ID": "a0977806-4347-43f2-a497-34af310f47f9",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": "%rH"
                    }
                  }
                },
                "temperature": {
                  "ID": "2d2914de-a911-4c56-b2e8-bfb6e66573df",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "temperature",
                  "Value": "73.76",
                  "Attrs": {
                    "$datatype": {
                      "ID": "cb444bc0-d0e0-43f5-8e93-71b4859ac812",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "be01cd48-4817-4cb1-8d7b-69da0cb00af9",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$name",
                      "Value": "Temperature"
                    },
                    "$unit": {
                      "ID": "639d3135-2e79-4e36-b711-109b6610533f",
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
              "ID": "783e1af9-c8f3-4f7b-bcd5-fcbb1842958d",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "Presence",
              "Attrs": {
                "$name": {
                  "ID": "e96795b2-c663-49b3-ba96-78ae424da9c4",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$name",
                  "Value": "Motion Sensor"
                },
                "$properties": {
                  "ID": "2b4f8cc5-2055-4b13-a8b0-f9d4642b27d5",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$properties",
                  "Value": "motion"
                },
                "$type": {
                  "ID": "50aebd1b-2423-4236-8519-7202721b2a8f",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "motion": {
                  "ID": "c63c827f-087b-48b9-ac13-543d09f5a18e",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "motion",
                  "Value": "CLOSED",
                  "Attrs": {
                    "$datatype": {
                      "ID": "154cbbda-7fb0-4845-9fab-23b2fd111827",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "12460c7d-6bf1-4e81-b8a2-8ec65ca1bbc0",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "d1f1ea66-56d6-45db-8a68-00d040fca2c9",
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
              "ID": "21950a51-6329-4d6a-a3a2-d0230fd4fca1",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "hardware",
              "Attrs": {
                "$name": {
                  "ID": "ea302196-4066-4888-99c6-3d765a1a5fec",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$name",
                  "Value": "Device Info"
                },
                "$properties": {
                  "ID": "d11b37cf-99a9-4835-87cc-7d16e15aafa6",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$properties",
                  "Value": "signal,mac,resetReason,voltage"
                },
                "$type": {
                  "ID": "712e5b1c-be6c-4442-840c-4012540ab3c0",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "$type",
                  "Value": "sensor"
                }
              },
              "Props": {
                "mac": {
                  "ID": "c9403a75-362e-4c37-b1b0-3d9591627ab1",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "mac",
                  "Value": "BC:DD:C2:81:04:72",
                  "Attrs": {
                    "$datatype": {
                      "ID": "cc0a5f96-f2bf-45fe-8d8d-69e43d9c65ee",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": "sring"
                    },
                    "$name": {
                      "ID": "41996315-370c-4e6e-9e54-5b77fc241879",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$name",
                      "Value": "mac"
                    }
                  }
                },
                "resetReason": {
                  "ID": "95d99c8b-229e-424e-b32f-26bc8837df1a",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "resetReason",
                  "Value": "External System",
                  "Attrs": {
                    "$datatype": {
                      "ID": "0049bd55-71a7-4a2b-96dc-e3d52870d342",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$name": {
                      "ID": "2c33e9fe-99e0-4849-bbb2-ef2b03a64d5a",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$name",
                      "Value": "Last Reset Reason"
                    }
                  }
                },
                "signal": {
                  "ID": "9d24a8ff-0f6a-4c07-9fa2-87f5fb22fa3a",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "signal",
                  "Value": "-64",
                  "Attrs": {
                    "$datatype": {
                      "ID": "932e135a-0923-4532-88c8-a9c31b83d431",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$name": {
                      "ID": "4f7711f0-c95c-4caf-8bf0-a82b41897e8a",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$name",
                      "Value": "RSSI"
                    },
                    "$unit": {
                      "ID": "bb18ec4e-21f0-49c4-894b-3395f38d8613",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": "dBm"
                    }
                  }
                },
                "voltage": {
                  "ID": "0a1e8a3e-11e5-4c5e-b9fd-15621e515d1c",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "voltage",
                  "Value": "2.98",
                  "Attrs": {
                    "$datatype": {
                      "ID": "133cc3ac-993d-4e0a-9a4d-d57d074cb0bb",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "288078e4-e22f-42cd-a6f6-ccfb6ca2a00a",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$name",
                      "Value": "3.3V Supply"
                    },
                    "$unit": {
                      "ID": "8b1f3967-8a66-4647-82f5-f1f53eb3a4a3",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$unit",
                      "Value": "V"
                    }
                  }
                }
              }
            },
            "humidity": {
              "ID": "edaf17e1-ba04-4d1a-a112-83653a1efe42",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "humidity",
              "Attrs": {},
              "Props": {
                "humidity": {
                  "ID": "df5a38aa-5fad-4eda-ae40-100f6875164d",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "humidity",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "7d0a2f9a-c9a2-4fad-bc7f-0091043c7def",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "ccb0161b-38ae-4ffe-b565-96740b526875",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$name",
                      "Value": "Humidity"
                    },
                    "$unit": {
                      "ID": "487d4756-c5fe-41d4-913e-75c5294b6ddc",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": "%rH"
                    }
                  }
                }
              }
            },
            "mac": {
              "ID": "bd52a83a-e8a9-4593-b580-f81e2637106a",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "mac",
              "Attrs": {},
              "Props": {
                "mac": {
                  "ID": "326cea65-076a-416a-a32e-1639fec92c8c",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "mac",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "6e1f354c-e3d1-48fe-a6ee-c39059dd570a",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": "sring"
                    },
                    "$name": {
                      "ID": "f98b4e75-ab14-43ad-bde0-5d6d2ccabcee",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$name",
                      "Value": "mac"
                    }
                  }
                }
              }
            },
            "motion": {
              "ID": "02cbe174-bc2b-465e-b2a9-2562377148cd",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "motion",
              "Attrs": {},
              "Props": {
                "motion": {
                  "ID": "dfaa19ba-18fe-4865-a997-19b0ac0c9be4",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "motion",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "e11a27dd-bc1f-4590-ac1d-393fcd9b7598",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "6091a743-445d-4fd9-be02-fd3d76355472",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "ffbb7bd8-6a88-4b21-935d-0dbeaafc341e",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$name",
                      "Value": "Motion"
                    }
                  }
                }
              }
            },
            "resetReason": {
              "ID": "af748739-314a-4ae7-9b74-ff9555386285",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "resetReason",
              "Attrs": {},
              "Props": {
                "resetReason": {
                  "ID": "8a5dcdee-3b26-4541-9c4c-2a6dbf2acf7b",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "resetReason",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "7288646f-09e5-467b-ad0b-38f4ce13fae7",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$name": {
                      "ID": "6a572b1b-4aa1-4229-8462-ad0cf0682453",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$name",
                      "Value": "Last Reset Reason"
                    }
                  }
                }
              }
            },
            "signal": {
              "ID": "0c19749b-cbd0-42f5-9cbc-5f3a6e7fa827",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "signal",
              "Attrs": {},
              "Props": {
                "signal": {
                  "ID": "44ba3169-090b-4a17-9ba7-929ea8452690",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "signal",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "3867e441-06e5-4116-8eb1-7bb485a6d9e2",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$name": {
                      "ID": "4182da3f-0f2d-4ea8-8698-4e0dcab02b7c",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$name",
                      "Value": "RSSI"
                    },
                    "$unit": {
                      "ID": "67b62e75-3ada-40fd-9b6b-aa6616b81b12",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": "dBm"
                    }
                  }
                }
              }
            },
            "temperature": {
              "ID": "e06b3833-c63e-4dcd-b717-69c57d88c1af",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "temperature",
              "Attrs": {},
              "Props": {
                "temperature": {
                  "ID": "55333543-e1cb-4f12-882b-f5025c19c5ac",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "temperature",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "278c9998-5cb4-455e-8891-deda75043358",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "f66735d5-a1dc-45b5-91a2-3aad63623719",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$name",
                      "Value": "Temperature"
                    },
                    "$unit": {
                      "ID": "810b56c3-c7cb-4a32-8fce-7b8a0f2c1ff8",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$unit",
                      "Value": "°F"
                    }
                  }
                }
              }
            },
            "voltage": {
              "ID": "bb4bd407-79a6-4386-b20a-18fe75daaf33",
              "ElementType": 14,
              "Parent": "GuestRoom",
              "Name": "voltage",
              "Attrs": {},
              "Props": {
                "voltage": {
                  "ID": "f02f03d5-5918-4267-b615-3fae498457d9",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "voltage",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "235b7d7d-67d1-4bff-b862-25e7c0cbc00a",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "08aa6862-f30d-43ab-a31b-5a7eeeec873b",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$name",
                      "Value": "3.3V Supply"
                    },
                    "$unit": {
                      "ID": "89819807-cd2b-4401-8e9f-31cc00e8d44c",
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
              "ID": "cb244e9e-3c31-4120-86f7-ee6f8711835d",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "885522cf-f1cd-4fb0-b860-4f1406be28ee",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "3ef8fbb48c5b23788a22a998c14a1a6d",
                  "Props": {}
                },
                "name": {
                  "ID": "631c5eb7-feaa-422b-a009-51d9639a57b9",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "Monitor-DHT-RCWL-Metrics",
                  "Props": {}
                },
                "version": {
                  "ID": "7127aadb-735f-4bf4-9330-c2c59c300202",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "2.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "616c504e-761d-4e00-a7f3-4e5bbf0d0760",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "4b5ede3e-3b52-44a7-b989-cc69744ba87d",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$implementation",
              "Value": "esp8266",
              "Props": {
                "config": {
                  "ID": "d3342d7f-f9d1-4f06-be61-0f17833d9587",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"Home Office\",\"device_id\":\"HomeOffice\",\"device_stats_interval\":180,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"sensorInterval\":900,\"motionHoldInterval\":60}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "991ddd3a-f942-4f94-9a34-53be3e95c0a5",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "enabled": {
                      "ID": "6a204181-6e69-425b-a358-edc06a18f8e2",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "6f9bf083-933e-49e9-a7c3-01ed2914d128",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "0570c936-0244-47cb-910f-ec973b5bf668",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$localip",
              "Value": "10.100.1.161",
              "Props": {}
            },
            "$mac": {
              "ID": "c445e11c-a85a-4ad8-9188-a4ccf45f84d4",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$mac",
              "Value": "84:F3:EB:B7:55:D5",
              "Props": {}
            },
            "$name": {
              "ID": "37413265-8667-45f5-b698-8c22ffb445da",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$name",
              "Value": "Home Office",
              "Props": {}
            },
            "$nodes": {
              "ID": "0d83f5a8-9dd5-44c7-ab8f-c7ae2c84a2b2",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$nodes",
              "Value": "Ambient,Presence,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "684280a0-df42-4dd0-94d1-538f77960fbf",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "3f48ed9b-be6d-4fb7-b7d7-a6bdafb9fc22",
              "ElementType": 11,
              "Parent": "HomeOffice",
              "Name": "$stats",
              "Value": "uptime",
              "Props": {
                "interval": {
                  "ID": "688a322c-38b7-45f7-811e-a9498b6090ea",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "185",
                  "Props": {}
                },
                "signal": {
                  "ID": "fdf01b6d-8b01-4276-8335-5f1ca69df4b7",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "70",
                  "Props": {}
                },
                "uptime": {
                  "ID": "0e0d6b10-f98e-44f5-a3df-68c6d6268773",
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
            "$name": {
              "ID": "e5298ad3-6622-44c3-ac43-53ac13f682f0",
              "ElementType": 14,
              "Parent": "HomeOffice",
              "Name": "$name",
              "Attrs": {
                "Ambient": {
                  "ID": "e0f4b6c9-d17c-4b15-9a94-c1c969fd3a65",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "3692b785-63c1-4ffa-b30e-9f7ff0c1a768",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "5c228097-881f-419e-959b-364871a21acc",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$properties": {
              "ID": "075dbb7c-5d9b-48ed-98a3-9442f36d787a",
              "ElementType": 14,
              "Parent": "HomeOffice",
              "Name": "$properties",
              "Attrs": {
                "Ambient": {
                  "ID": "9e1054a0-a263-4cdd-9e7e-b28df1faa8b5",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "125fd5cd-0b81-4f8a-a0c3-5742f5fbfe13",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "68df104d-f094-421a-a7db-02483896ae9a",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$type": {
              "ID": "03e037d6-e4e1-4f5c-9af3-787d3384dc28",
              "ElementType": 14,
              "Parent": "HomeOffice",
              "Name": "$type",
              "Attrs": {
                "Ambient": {
                  "ID": "dd0f9740-84e4-433f-8e61-528c5e8f1fba",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "a5029a7e-2d3d-4bbd-bfd4-00e95c5cfcfc",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "8146ea77-5396-4024-8ce0-4e6a1ff60d3b",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "humidity": {
              "ID": "f85a6376-a715-4bc9-8eb7-974e8f1c0047",
              "ElementType": 14,
              "Parent": "HomeOffice",
              "Name": "humidity",
              "Attrs": {},
              "Props": {
                "humidity": {
                  "ID": "d655969c-2468-4ffe-b63d-470ea78cc7e6",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "humidity",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "2d09f170-02ec-4d05-8595-018b05ed95eb",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "4ae88e4a-8483-468e-b4c1-869b99dcab1c",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$name",
                      "Value": "Humidity"
                    },
                    "$unit": {
                      "ID": "a214d587-36ff-43d1-834b-fe0311595257",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": "%rH"
                    }
                  }
                }
              }
            },
            "mac": {
              "ID": "0141b949-c58a-4a78-a188-2e42e5e71e9f",
              "ElementType": 14,
              "Parent": "HomeOffice",
              "Name": "mac",
              "Attrs": {},
              "Props": {
                "mac": {
                  "ID": "3eb1f5dc-34ad-41f2-a8e0-8a94c575dac2",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "mac",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "e01805bf-1f16-4d03-b912-11fadcb0ba8c",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": "sring"
                    },
                    "$name": {
                      "ID": "a6530d43-6d0e-466f-adfe-ceb354caf8eb",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$name",
                      "Value": "mac"
                    }
                  }
                }
              }
            },
            "motion": {
              "ID": "e0bcef9d-a7d1-4208-95cf-73c807c2b97c",
              "ElementType": 14,
              "Parent": "HomeOffice",
              "Name": "motion",
              "Attrs": {},
              "Props": {
                "motion": {
                  "ID": "99ace8b9-5939-4ef6-9b90-5eaf3b239488",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "motion",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "bb538230-25da-4796-9c66-df13b1311f52",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "76fecab2-8c06-4514-beec-ae9fc2f4bd21",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "b641e720-8fbc-416c-a997-358883bf874e",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$name",
                      "Value": "Motion"
                    }
                  }
                }
              }
            },
            "resetReason": {
              "ID": "ba81f742-ba27-4396-af43-2e78d7fcf232",
              "ElementType": 14,
              "Parent": "HomeOffice",
              "Name": "resetReason",
              "Attrs": {},
              "Props": {
                "resetReason": {
                  "ID": "60aaae45-a933-4a20-bfe6-e154f2ab694c",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "resetReason",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "f762d982-6950-4471-a032-e8f42bd3e1ec",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$name": {
                      "ID": "fc23b213-b415-4907-98d3-b3f2f2122587",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$name",
                      "Value": "Last Reset Reason"
                    }
                  }
                }
              }
            },
            "signal": {
              "ID": "b8125d69-667b-40b6-a06b-a541b26cbb89",
              "ElementType": 14,
              "Parent": "HomeOffice",
              "Name": "signal",
              "Attrs": {},
              "Props": {
                "signal": {
                  "ID": "403d14cd-dd07-48bd-9856-d4fe3cc2223b",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "signal",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "3c88eb77-e3aa-422a-8858-d6dceba082e1",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$name": {
                      "ID": "227fa05e-f5fd-4e74-b92f-dbe839d31d60",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$name",
                      "Value": "RSSI"
                    },
                    "$unit": {
                      "ID": "77ace5e8-88ad-46ea-be41-540676d84a42",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": "dBm"
                    }
                  }
                }
              }
            },
            "temperature": {
              "ID": "af965cad-1fff-47b4-ad0b-5f5bacd880da",
              "ElementType": 14,
              "Parent": "HomeOffice",
              "Name": "temperature",
              "Attrs": {},
              "Props": {
                "temperature": {
                  "ID": "3fe14983-e1e3-47d3-851c-f5f03da5597f",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "temperature",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "dbe58d11-a20c-4f66-90f9-97bb88f60fb4",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "7d5871c9-adf6-4f88-8508-bc798d6f0aa0",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$name",
                      "Value": "Temperature"
                    },
                    "$unit": {
                      "ID": "963dc43a-2d03-4d49-8d18-d0648baa72bf",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$unit",
                      "Value": "°F"
                    }
                  }
                }
              }
            },
            "voltage": {
              "ID": "251ec16e-2e26-4ddc-81ce-49edbab59b27",
              "ElementType": 14,
              "Parent": "HomeOffice",
              "Name": "voltage",
              "Attrs": {},
              "Props": {
                "voltage": {
                  "ID": "bffa08d4-779e-4d11-a718-5d3faad0d03c",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "voltage",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "80602e13-8011-46ee-9b33-3511c26e396d",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "cc6b7daa-0839-4db3-a53a-eb7d9b3e0a6d",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$name",
                      "Value": "3.3V Supply"
                    },
                    "$unit": {
                      "ID": "42e37b0e-030e-4fd2-b75e-6dbf6e4daea0",
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
              "ID": "1581b325-a9f8-4385-8c64-76de646dcd24",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "c316d663-d153-4cc8-9fc1-4411db356bed",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "b92b40fead22529097629c0597186c05",
                  "Props": {}
                },
                "name": {
                  "ID": "91723bab-b455-4234-b5df-6e229d6f7ae8",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "WiredProvider",
                  "Props": {}
                },
                "version": {
                  "ID": "aa8dee92-8111-4e5f-8002-bda1bfde2da1",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "1.4.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "9655e98b-5fdb-4236-a817-e1ce86e20e11",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "a2034d20-7390-450d-92ae-a3fc31f8df32",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$implementation",
              "Value": "esp8266",
              "Props": {
                "config": {
                  "ID": "05203d90-e5b7-47e6-abfd-cd83ae45f940",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"Hard Wired Alarm Sensors\",\"device_id\":\"MCPProvider\",\"device_stats_interval\":300,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"ipolA\":255,\"ipolB\":255}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "9f19fd6d-3b8f-47b2-9103-404f1a6ee65e",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "enabled": {
                      "ID": "b192e0b6-3072-4c8d-88b2-808c5fa41272",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "9ac597ae-1d2f-4439-8638-ddb746e3e77c",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "a9bff277-3a03-4947-bcfd-436b2e25ee94",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$localip",
              "Value": "10.100.1.222",
              "Props": {}
            },
            "$mac": {
              "ID": "cbe4e48e-abd9-447d-9e16-30eea49a6c9e",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$mac",
              "Value": "BC:DD:C2:24:B7:3C",
              "Props": {}
            },
            "$name": {
              "ID": "103b45f5-8973-4278-aef0-452ad8d9e13f",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$name",
              "Value": "Hard Wired Alarm Sensors",
              "Props": {}
            },
            "$nodes": {
              "ID": "903ca5aa-b961-46ef-9789-572db6f3262d",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$nodes",
              "Value": "wiredMonitor,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "a7b76a64-d28c-4184-a8de-4bb8bd6e7e89",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "9b32942a-28b2-48c4-91e6-a7cc5c346a77",
              "ElementType": 11,
              "Parent": "MCPProvider",
              "Name": "$stats",
              "Value": "uptime",
              "Props": {
                "interval": {
                  "ID": "2bdd9c20-3d0f-47da-9f88-2b007521fe96",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "305",
                  "Props": {}
                },
                "signal": {
                  "ID": "bd52a460-3492-4b8d-aba4-f79d42a6ad54",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "58",
                  "Props": {}
                },
                "uptime": {
                  "ID": "27e40c50-a924-4975-9184-47c9e8b42c44",
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
            "$name": {
              "ID": "6b18e190-f7f5-414e-9609-ce827f671377",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "$name",
              "Attrs": {
                "hardware": {
                  "ID": "e48d185b-74a4-4444-8093-d095f85aaa24",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                },
                "wiredMonitor": {
                  "ID": "d76912a4-4179-4770-b11d-7d917c72aa77",
                  "ElementType": 15,
                  "Parent": "wiredMonitor",
                  "Name": "wiredMonitor",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$properties": {
              "ID": "cf74eb4f-f5eb-43a4-bf69-11f7279530b3",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "$properties",
              "Attrs": {
                "hardware": {
                  "ID": "ef252100-0e29-4af4-8c36-d352f81ceac0",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                },
                "wiredMonitor": {
                  "ID": "972cda5e-23a0-4d7c-87a8-dee18607dd4b",
                  "ElementType": 15,
                  "Parent": "wiredMonitor",
                  "Name": "wiredMonitor",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$type": {
              "ID": "23f0ac81-6dc2-4148-80b3-be2a01ffeb72",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "$type",
              "Attrs": {
                "hardware": {
                  "ID": "ec2321d1-1ca5-4fe6-ab68-93beaf059edb",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                },
                "wiredMonitor": {
                  "ID": "4155f1c1-1041-4ef0-8931-23373678fc93",
                  "ElementType": 15,
                  "Parent": "wiredMonitor",
                  "Name": "wiredMonitor",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "mac": {
              "ID": "fd9f971b-48ba-4594-b1f7-90526d974259",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "mac",
              "Attrs": {},
              "Props": {
                "mac": {
                  "ID": "e0715751-a467-4029-a035-32162e390f04",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "mac",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "97d32e38-72b0-454f-8ba4-d88a6a3f00e1",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": "sring"
                    },
                    "$name": {
                      "ID": "3f878a60-11d8-473e-bf83-13af0eb8b334",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$name",
                      "Value": "mac"
                    }
                  }
                }
              }
            },
            "pin0": {
              "ID": "7dc1d177-1c44-4b68-9cfa-1772c6cc78ba",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "pin0",
              "Attrs": {},
              "Props": {
                "pin0": {
                  "ID": "8f599fce-f8e3-4be5-ad51-d95e84cef928",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin0",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "91f1ba97-31de-4a7f-bffd-a379a4f9bc30",
                      "ElementType": 17,
                      "Parent": "pin0",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "3f2a44e7-75a3-42f7-8fa3-3cd7af5fb7e9",
                      "ElementType": 17,
                      "Parent": "pin0",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "24e87691-40cc-402e-8d06-4fa59b29cb3d",
                      "ElementType": 17,
                      "Parent": "pin0",
                      "Name": "$name",
                      "Value": "Pin 0"
                    }
                  }
                }
              }
            },
            "pin1": {
              "ID": "147860a2-726e-4c06-ae3b-35eb2177c475",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "pin1",
              "Attrs": {},
              "Props": {
                "pin1": {
                  "ID": "ba57d2dd-c0ed-4779-8181-6cc9a694e40f",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin1",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "d79a4c0b-1465-4e07-bbf0-f9945ba1e25b",
                      "ElementType": 17,
                      "Parent": "pin1",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "bdbeb0df-52e6-48bc-a43d-08fb13a0e6c1",
                      "ElementType": 17,
                      "Parent": "pin1",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "cd570b16-3f4e-406f-bb8e-db1e13f5f51f",
                      "ElementType": 17,
                      "Parent": "pin1",
                      "Name": "$name",
                      "Value": "Pin 1"
                    }
                  }
                }
              }
            },
            "pin10": {
              "ID": "ca1d0103-b94b-4b4e-868a-ea2c2b8ac403",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "pin10",
              "Attrs": {},
              "Props": {
                "pin10": {
                  "ID": "65e5f581-7b01-4353-8993-51e7def1fc8c",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin10",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "2025be84-4a7d-40a8-a6f1-56bcf5620a4e",
                      "ElementType": 17,
                      "Parent": "pin10",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "33a0ccb1-5645-4f6d-8708-4005fba98386",
                      "ElementType": 17,
                      "Parent": "pin10",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "832ab2ac-2989-4056-a43b-649a91b51f8e",
                      "ElementType": 17,
                      "Parent": "pin10",
                      "Name": "$name",
                      "Value": "Pin 10"
                    }
                  }
                }
              }
            },
            "pin11": {
              "ID": "c807e824-5605-458f-9bca-09ddf57282a8",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "pin11",
              "Attrs": {},
              "Props": {
                "pin11": {
                  "ID": "02af1a66-dc8d-4746-818e-57658d4ad317",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin11",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "2201ea62-a9fe-4ad5-89bb-825db53f99e9",
                      "ElementType": 17,
                      "Parent": "pin11",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "b43527ac-564c-4118-b8c2-458f26fc7489",
                      "ElementType": 17,
                      "Parent": "pin11",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "f2a84b88-800a-41ae-ae8e-61caf34ab93a",
                      "ElementType": 17,
                      "Parent": "pin11",
                      "Name": "$name",
                      "Value": "Pin 11"
                    }
                  }
                }
              }
            },
            "pin12": {
              "ID": "f1e2ecf5-7e01-4e5a-81e3-0a054d46c2a5",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "pin12",
              "Attrs": {},
              "Props": {
                "pin12": {
                  "ID": "41fb6f84-5b7e-44e3-99f7-b496678da4a5",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin12",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "c5dafe34-09d3-49cf-91a4-27d198c98079",
                      "ElementType": 17,
                      "Parent": "pin12",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "8429b1d7-4bac-4f0f-83d1-b44cb561144f",
                      "ElementType": 17,
                      "Parent": "pin12",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "b8ca5c59-e9ef-4dac-9835-f163e43a2aa7",
                      "ElementType": 17,
                      "Parent": "pin12",
                      "Name": "$name",
                      "Value": "Pin 12"
                    }
                  }
                }
              }
            },
            "pin13": {
              "ID": "d80a9ad2-b446-4bc8-b35b-e5d4cf254d64",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "pin13",
              "Attrs": {},
              "Props": {
                "pin13": {
                  "ID": "e4e517b3-f52f-4dc2-b06d-e96f44203fa7",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin13",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "a07cf9cc-5173-4d83-91ce-0e1aee4cdb1c",
                      "ElementType": 17,
                      "Parent": "pin13",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "0cbce8c6-9da5-448a-9e36-a542670a0a76",
                      "ElementType": 17,
                      "Parent": "pin13",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "ba7c55e3-75ca-4869-85ab-9e2cb84da6dc",
                      "ElementType": 17,
                      "Parent": "pin13",
                      "Name": "$name",
                      "Value": "Pin 13"
                    }
                  }
                }
              }
            },
            "pin14": {
              "ID": "669ec139-4a88-4f73-b12c-ccde611902cc",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "pin14",
              "Attrs": {},
              "Props": {
                "pin14": {
                  "ID": "f4acd27e-6da0-452a-8a6f-ed5544ba1c41",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin14",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "ecff7c4b-9d5f-4448-a5ff-2d6b0a367b10",
                      "ElementType": 17,
                      "Parent": "pin14",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "f0cb2372-d4e3-46a4-838e-1be509cd7826",
                      "ElementType": 17,
                      "Parent": "pin14",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "00050b6d-6a8c-4c7a-bd49-82bbba5a5e19",
                      "ElementType": 17,
                      "Parent": "pin14",
                      "Name": "$name",
                      "Value": "Pin 14"
                    }
                  }
                }
              }
            },
            "pin15": {
              "ID": "c58e8fac-0445-44da-87c0-e31a65f01d2c",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "pin15",
              "Attrs": {},
              "Props": {
                "pin15": {
                  "ID": "e1680a8d-e368-4ff9-a8b5-457c7f04a1ac",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin15",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "549a5db0-ebf5-4351-bbb1-113502d5babc",
                      "ElementType": 17,
                      "Parent": "pin15",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "53452a99-9c18-4d29-892d-7ff392066961",
                      "ElementType": 17,
                      "Parent": "pin15",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "29273def-158e-4da7-a6ed-e03bec6275ed",
                      "ElementType": 17,
                      "Parent": "pin15",
                      "Name": "$name",
                      "Value": "Pin 15"
                    }
                  }
                }
              }
            },
            "pin2": {
              "ID": "b2c19a71-a042-4fb5-b480-4a0978546607",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "pin2",
              "Attrs": {},
              "Props": {
                "pin2": {
                  "ID": "6777207f-8620-4b4a-a648-a79b0f381482",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin2",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "b723726e-6245-455e-9465-fac917c53f93",
                      "ElementType": 17,
                      "Parent": "pin2",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "eb1ed3fb-e156-4a00-a8d2-5683910551ff",
                      "ElementType": 17,
                      "Parent": "pin2",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "0caa87e9-5a8a-491b-b1cc-f37ba1a04c66",
                      "ElementType": 17,
                      "Parent": "pin2",
                      "Name": "$name",
                      "Value": "Pin 2"
                    }
                  }
                }
              }
            },
            "pin3": {
              "ID": "065a134f-b228-400d-843d-610f844b13a0",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "pin3",
              "Attrs": {},
              "Props": {
                "pin3": {
                  "ID": "4843aca8-1d26-49ad-beab-f11ee39ee877",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin3",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "f071e10e-26b8-4e01-8201-4273378458b6",
                      "ElementType": 17,
                      "Parent": "pin3",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "c3106ef3-b859-402a-94c0-e8ab9909ad86",
                      "ElementType": 17,
                      "Parent": "pin3",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "aab3a26b-fb5a-414b-9c22-5bf953ab4585",
                      "ElementType": 17,
                      "Parent": "pin3",
                      "Name": "$name",
                      "Value": "Pin 3"
                    }
                  }
                }
              }
            },
            "pin4": {
              "ID": "c7a0f235-ab94-41e2-a355-ab1000cf959b",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "pin4",
              "Attrs": {},
              "Props": {
                "pin4": {
                  "ID": "5ddab75f-3bc4-43a2-8649-98ee1288e365",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin4",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "185b84ff-6f83-4713-822f-d5ad7d4f5c0b",
                      "ElementType": 17,
                      "Parent": "pin4",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "8cf471c7-fc90-472a-bbfb-1a7ce4ccb38e",
                      "ElementType": 17,
                      "Parent": "pin4",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "9ffe5d0e-64c4-48ff-ba1a-8f8f704cb2d7",
                      "ElementType": 17,
                      "Parent": "pin4",
                      "Name": "$name",
                      "Value": "Pin 4"
                    }
                  }
                }
              }
            },
            "pin5": {
              "ID": "adb387ee-880d-4900-965a-d35e2e8f41d1",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "pin5",
              "Attrs": {},
              "Props": {
                "pin5": {
                  "ID": "13466437-0428-4f8d-963e-65c97bc525cf",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin5",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "17f08805-cde4-4208-a950-86dac9161ba6",
                      "ElementType": 17,
                      "Parent": "pin5",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "ade37ff0-71a4-4624-b936-d438563fe1b1",
                      "ElementType": 17,
                      "Parent": "pin5",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "98e4d446-376e-443f-95c6-884d1b87f25e",
                      "ElementType": 17,
                      "Parent": "pin5",
                      "Name": "$name",
                      "Value": "Pin 5"
                    }
                  }
                }
              }
            },
            "pin6": {
              "ID": "9ab07179-4e41-4dd2-9efb-ed890953577e",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "pin6",
              "Attrs": {},
              "Props": {
                "pin6": {
                  "ID": "f5db577e-9a79-48fb-8514-df4823ab238e",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin6",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "954e9612-7f3e-49cb-9b10-2df96e10baa5",
                      "ElementType": 17,
                      "Parent": "pin6",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "3c14f046-5ab6-48e4-b734-a4ad211d19bb",
                      "ElementType": 17,
                      "Parent": "pin6",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "fe9d664b-1f35-46a6-a8f1-a1a4620b484b",
                      "ElementType": 17,
                      "Parent": "pin6",
                      "Name": "$name",
                      "Value": "Pin 6"
                    }
                  }
                }
              }
            },
            "pin7": {
              "ID": "0d2717d8-628f-43c6-9d14-c62da1ef184b",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "pin7",
              "Attrs": {},
              "Props": {
                "pin7": {
                  "ID": "cc1b31e0-3e5d-4b93-a368-66eaaf544109",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin7",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "5992fe71-572f-4732-9be1-dc71ed7f0e3c",
                      "ElementType": 17,
                      "Parent": "pin7",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "496284c0-0ce1-454c-b9f1-5a16d5895819",
                      "ElementType": 17,
                      "Parent": "pin7",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "2ce50731-75ea-4ed6-863d-519cfd0cb724",
                      "ElementType": 17,
                      "Parent": "pin7",
                      "Name": "$name",
                      "Value": "Pin 7"
                    }
                  }
                }
              }
            },
            "pin8": {
              "ID": "7aab5380-459f-42bf-ab06-e3c794b25e2f",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "pin8",
              "Attrs": {},
              "Props": {
                "pin8": {
                  "ID": "3a2b08d2-5b81-4079-bd7d-3a556cbf8404",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin8",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "a999d73b-a998-4b84-8f61-8dca86ee9b8e",
                      "ElementType": 17,
                      "Parent": "pin8",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "85dc3501-4b0f-4c70-af1c-ec043900cdc7",
                      "ElementType": 17,
                      "Parent": "pin8",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "fddacc32-8eb6-4a34-8b9d-8024f26ac0e2",
                      "ElementType": 17,
                      "Parent": "pin8",
                      "Name": "$name",
                      "Value": "Pin 8"
                    }
                  }
                }
              }
            },
            "pin9": {
              "ID": "d5217d34-af1b-4fb6-bdd5-ade521a81edf",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "pin9",
              "Attrs": {},
              "Props": {
                "pin9": {
                  "ID": "b2e9eec8-bf69-4ed9-a4c7-b06af6100b37",
                  "ElementType": 16,
                  "Parent": "wiredMonitor",
                  "Name": "pin9",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "9bce39be-1829-4322-a46c-d3b4d0d43efb",
                      "ElementType": 17,
                      "Parent": "pin9",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "53ce690c-294e-4977-bd2f-7260a1350aeb",
                      "ElementType": 17,
                      "Parent": "pin9",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "3c87cd53-f36a-4c5d-9dd5-9c85476e902f",
                      "ElementType": 17,
                      "Parent": "pin9",
                      "Name": "$name",
                      "Value": "Pin 9"
                    }
                  }
                }
              }
            },
            "resetReason": {
              "ID": "4c195389-66ee-4e57-b924-d6b0574828b7",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "resetReason",
              "Attrs": {},
              "Props": {
                "resetReason": {
                  "ID": "06845d97-5bcd-4606-92f0-ecf3c86f4319",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "resetReason",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "418bb39e-d5b2-4133-a41b-0613028cbf4e",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$name": {
                      "ID": "f9b9b743-ade8-46eb-b4e4-e37de1d1bb3c",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$name",
                      "Value": "Last Reset Reason"
                    }
                  }
                }
              }
            },
            "signal": {
              "ID": "152ce614-e0ff-4c93-a00e-86e9db8d3cdb",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "signal",
              "Attrs": {},
              "Props": {
                "signal": {
                  "ID": "bbb29133-b93c-4321-b949-5e16094b6cd6",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "signal",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "3535ce64-31c0-4167-ab25-f2c998ded393",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$name": {
                      "ID": "d812fff9-0620-4024-9328-a5b588f6312e",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$name",
                      "Value": "RSSI"
                    },
                    "$unit": {
                      "ID": "f0ffa7c0-107f-45f3-8226-5dd4745c38a3",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": "dBm"
                    }
                  }
                }
              }
            },
            "voltage": {
              "ID": "bdca6ed9-f35d-4198-801c-093459965381",
              "ElementType": 14,
              "Parent": "MCPProvider",
              "Name": "voltage",
              "Attrs": {},
              "Props": {
                "voltage": {
                  "ID": "5a08cc9d-9bbf-4cf3-8eff-b63d326f7d5c",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "voltage",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "1b5e98c8-1509-4e21-b04d-74f03240e322",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "bf1fe812-ed44-4c59-a194-26bdf224689c",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$name",
                      "Value": "3.3V Supply"
                    },
                    "$unit": {
                      "ID": "10dac851-4fcb-43e3-8851-57d65189496a",
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
        "MediaRoom": {
          "ID": "d187e485b51c392d189d86bb61ce45a2",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "MediaRoom",
          "Attrs": {
            "$fw": {
              "ID": "30e80b28-eb81-4de8-8509-98d36e93c6ff",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "69e7c8a9-c7b1-4196-95f2-c4c652a4e9a7",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "3ef8fbb48c5b23788a22a998c14a1a6d",
                  "Props": {}
                },
                "name": {
                  "ID": "dbf3c348-ee2e-4d73-a597-a3faeb562777",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "Monitor-DHT-RCWL-Metrics",
                  "Props": {}
                },
                "version": {
                  "ID": "2c2cfc77-f787-4d0a-96e3-46075f8db44b",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "2.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "ab3ebde1-1d17-4462-a35f-066d2fe5b13a",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "083f73e6-01cd-45ed-be51-10914a89eb70",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$implementation",
              "Value": "esp8266",
              "Props": {
                "config": {
                  "ID": "bdba9217-6226-4e75-8a46-cdba79f3612f",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"Media Room\",\"device_id\":\"MediaRoom\",\"device_stats_interval\":180,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"sensorInterval\":900,\"motionHoldInterval\":60}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "58938252-a368-484a-880b-4b682f585891",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "enabled": {
                      "ID": "2e1565de-359b-4fad-9b9c-32876049a8f0",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "b387ce08-f4f9-4027-8950-8fd90bd8aa9b",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "79681b6d-f880-4a4b-bfe1-16a963b11d0e",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$localip",
              "Value": "10.100.1.180",
              "Props": {}
            },
            "$mac": {
              "ID": "ce8b2844-fb04-4544-82a0-52efe75672a4",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$mac",
              "Value": "B4:E6:2D:15:50:3A",
              "Props": {}
            },
            "$name": {
              "ID": "80f6c603-e146-4fcc-bb89-b630049e05c1",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$name",
              "Value": "Media Room",
              "Props": {}
            },
            "$nodes": {
              "ID": "e8046748-a820-4cc3-aaf5-9107f4976c94",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$nodes",
              "Value": "Ambient,Presence,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "10e66555-55ea-40ae-a883-2eb68a7858b3",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "c00f3c6c-19bc-40d5-890e-d6d4a947433c",
              "ElementType": 11,
              "Parent": "MediaRoom",
              "Name": "$stats",
              "Value": "uptime",
              "Props": {
                "interval": {
                  "ID": "99b9b646-6ea7-4222-a7dd-ac063fe24ab9",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "185",
                  "Props": {}
                },
                "signal": {
                  "ID": "9f0eb341-4b4b-4dd8-9a29-40dbbc93f4ff",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "66",
                  "Props": {}
                },
                "uptime": {
                  "ID": "61c358f6-1945-491f-a8d8-a4c2354932a9",
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
            "$name": {
              "ID": "ab8efd15-8611-4d80-bf3a-c0ff0c2a3581",
              "ElementType": 14,
              "Parent": "MediaRoom",
              "Name": "$name",
              "Attrs": {
                "Ambient": {
                  "ID": "521d83f3-2b53-459f-a20c-40c89452e451",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "0132c159-c8d0-4382-b122-8339fd2f25a4",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "06ca7a27-d578-45ed-870e-e22db9c03d17",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$properties": {
              "ID": "042c1bca-5afa-4259-96a5-451136c0c06c",
              "ElementType": 14,
              "Parent": "MediaRoom",
              "Name": "$properties",
              "Attrs": {
                "Ambient": {
                  "ID": "6751fc09-2bdd-43ab-a2de-c3df89c95852",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "663d3c28-ad33-494c-844f-e631d7a24866",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "7b1882be-4270-4643-b2ec-3f3c10805ea1",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$type": {
              "ID": "9fe59f58-8678-430e-a92e-11b9830efa13",
              "ElementType": 14,
              "Parent": "MediaRoom",
              "Name": "$type",
              "Attrs": {
                "Ambient": {
                  "ID": "0ad18d6f-b976-457e-a5aa-b64af6dab4db",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "689b063c-c48e-46b1-9ba8-d0e08a00f450",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "d19efca7-3732-4123-8509-9cb96a4b4e1d",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "humidity": {
              "ID": "70a46f7b-00a8-46f7-9497-017bf5a45406",
              "ElementType": 14,
              "Parent": "MediaRoom",
              "Name": "humidity",
              "Attrs": {},
              "Props": {
                "humidity": {
                  "ID": "abf08d83-c242-47d9-9808-61b57361c22c",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "humidity",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "e24b44f9-423f-4ccf-b57f-c9940db82a40",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "36542be6-67cc-49fa-8f0c-bd6ef4ac9589",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$name",
                      "Value": "Humidity"
                    },
                    "$unit": {
                      "ID": "6db61105-eb6f-496b-9646-e36320f305cd",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": "%rH"
                    }
                  }
                }
              }
            },
            "mac": {
              "ID": "90684a22-be38-4bf3-b72c-68b7a7fd8db3",
              "ElementType": 14,
              "Parent": "MediaRoom",
              "Name": "mac",
              "Attrs": {},
              "Props": {
                "mac": {
                  "ID": "33d706e3-34c5-4b26-a611-3b0d0135bd96",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "mac",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "c1806f8e-6edd-4ab7-a352-dccf55399163",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": "sring"
                    },
                    "$name": {
                      "ID": "36d8219d-de70-4411-8827-80632bfc639b",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$name",
                      "Value": "mac"
                    }
                  }
                }
              }
            },
            "motion": {
              "ID": "0c8b2d38-b757-4829-b225-24559445b903",
              "ElementType": 14,
              "Parent": "MediaRoom",
              "Name": "motion",
              "Attrs": {},
              "Props": {
                "motion": {
                  "ID": "e2946b55-f6e6-49e3-a5b1-9db8b72c3f45",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "motion",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "a7120d8d-2bcd-4373-8dac-7c1b58ed1b1c",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "d2dbb848-cb87-4df6-be51-cbd7ba5a2c65",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "1dfee8ec-0125-4ba7-9448-00b9601d42ed",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$name",
                      "Value": "Motion"
                    }
                  }
                }
              }
            },
            "resetReason": {
              "ID": "4d593c2e-3ee9-4515-8d34-6fa9d5814341",
              "ElementType": 14,
              "Parent": "MediaRoom",
              "Name": "resetReason",
              "Attrs": {},
              "Props": {
                "resetReason": {
                  "ID": "fc186eb2-cec2-4c38-a5e4-12c4f954aa93",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "resetReason",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "a3bb5287-1137-44c9-8954-035f4b9019d7",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$name": {
                      "ID": "6f414de3-afe0-4428-97af-0af0e62c3079",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$name",
                      "Value": "Last Reset Reason"
                    }
                  }
                }
              }
            },
            "signal": {
              "ID": "3bbb8948-54d2-4768-bad5-7824927c50b6",
              "ElementType": 14,
              "Parent": "MediaRoom",
              "Name": "signal",
              "Attrs": {},
              "Props": {
                "signal": {
                  "ID": "1bc86a51-dcf6-4eee-a2aa-2bff56736303",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "signal",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "28e6f647-046b-443c-8a8a-709bbe9a83a5",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$name": {
                      "ID": "073fd6dd-ab06-40db-a39a-ec8f0b9ab4e7",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$name",
                      "Value": "RSSI"
                    },
                    "$unit": {
                      "ID": "a8e83ff9-6c87-4749-8a82-609fe59b1ff3",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": "dBm"
                    }
                  }
                }
              }
            },
            "temperature": {
              "ID": "1e5f2172-4f04-46e9-88bc-fe4da0e2c587",
              "ElementType": 14,
              "Parent": "MediaRoom",
              "Name": "temperature",
              "Attrs": {},
              "Props": {
                "temperature": {
                  "ID": "9864798b-d0be-495e-a3e0-1c2422857130",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "temperature",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "b837f666-e17b-4b95-ae4c-ae75054c9d8a",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "b4ad9ae6-2630-4358-be66-24544b2f5427",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$name",
                      "Value": "Temperature"
                    },
                    "$unit": {
                      "ID": "761a006e-7cbb-423e-9f3f-79395f07e894",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$unit",
                      "Value": "°F"
                    }
                  }
                }
              }
            },
            "voltage": {
              "ID": "1c33d82d-4f5c-4a2e-899d-2097221269ca",
              "ElementType": 14,
              "Parent": "MediaRoom",
              "Name": "voltage",
              "Attrs": {},
              "Props": {
                "voltage": {
                  "ID": "b930ff14-0ccd-494a-80fc-ffb0b2f68ff4",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "voltage",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "d56335cc-7623-469a-85eb-3dba553fe233",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "e79ddb1f-a851-49c6-96b8-fabe6a635887",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$name",
                      "Value": "3.3V Supply"
                    },
                    "$unit": {
                      "ID": "29936ad8-fcb3-47d7-86e7-98e3dc66c942",
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
        "OutsideMonitor": {
          "ID": "c6e49801761844b56e0a5604847e6214",
          "Title": "",
          "ElementType": 10,
          "OTAEnabled": true,
          "Parent": "sknSensors",
          "Name": "OutsideMonitor",
          "Attrs": {
            "$fw": {
              "ID": "4eb892a0-2fa5-4eec-9fdf-7e8fdb628566",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$fw",
              "Value": "",
              "Props": {
                "checksum": {
                  "ID": "04a5bf57-fb0d-4b30-ad0e-f784b82d2f7f",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "checksum",
                  "Value": "615fed382ab44bd43fe83508aecac682",
                  "Props": {}
                },
                "name": {
                  "ID": "1e2191f4-f092-4544-a8a4-e08e1f98ff99",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "name",
                  "Value": "Monitor-SHT31-RCWL-Metrics",
                  "Props": {}
                },
                "version": {
                  "ID": "ca5ca678-4742-4d2b-b204-44ae5db3548d",
                  "ElementType": 12,
                  "Parent": "$fw",
                  "Name": "version",
                  "Value": "2.0.0",
                  "Props": {}
                }
              }
            },
            "$homie": {
              "ID": "a39a904e-b86d-4137-9f79-79e0dad17a82",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$homie",
              "Value": "3.0.1",
              "Props": {}
            },
            "$implementation": {
              "ID": "871ee7b5-312a-4319-96bb-645965b06f4b",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$implementation",
              "Value": "esp8266",
              "Props": {
                "config": {
                  "ID": "aa63adce-c460-4c99-94fe-a3ebf9cbe654",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "config",
                  "Value": "{\"name\":\"Outside Monitor\",\"device_id\":\"OutsideMonitor\",\"device_stats_interval\":900,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"sensorsInterval\":900,\"motionHoldInterval\":60}}",
                  "Props": {}
                },
                "ota": {
                  "ID": "c00990f6-1ecb-4cd1-aa1b-849c6da89ade",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "ota",
                  "Value": "",
                  "Props": {
                    "enabled": {
                      "ID": "36e0ddb5-9bb8-44db-ba5c-e28d465d2677",
                      "ElementType": 13,
                      "Parent": "ota",
                      "Name": "enabled",
                      "Value": "true"
                    }
                  }
                },
                "version": {
                  "ID": "228cb13b-b906-4f85-8f51-d05e9e92eaf0",
                  "ElementType": 12,
                  "Parent": "$implementation",
                  "Name": "version",
                  "Value": "3.0.0",
                  "Props": {}
                }
              }
            },
            "$localip": {
              "ID": "fb4e6518-5bbe-41d1-b1b6-c58d1940fc3b",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$localip",
              "Value": "10.100.1.182",
              "Props": {}
            },
            "$mac": {
              "ID": "c8cd1f39-b268-4690-83b1-4dca42db6200",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$mac",
              "Value": "18:FE:34:FD:8C:1B",
              "Props": {}
            },
            "$name": {
              "ID": "81e5986b-f5bc-4501-aa98-0bd99944105c",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$name",
              "Value": "Outside Monitor",
              "Props": {}
            },
            "$nodes": {
              "ID": "0b0dc076-c0bf-41e5-9754-57402f332c2b",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$nodes",
              "Value": "Ambient,Presence,hardware",
              "Props": {}
            },
            "$state": {
              "ID": "bdcd82bb-0866-4118-9748-198ea282d215",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$state",
              "Value": "ready",
              "Props": {}
            },
            "$stats": {
              "ID": "8a063bb6-b1dd-4430-b709-c6b1b019a041",
              "ElementType": 11,
              "Parent": "OutsideMonitor",
              "Name": "$stats",
              "Value": "uptime",
              "Props": {
                "interval": {
                  "ID": "cc83f98f-a974-422f-b6fe-4bc8f7648b53",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "interval",
                  "Value": "905",
                  "Props": {}
                },
                "signal": {
                  "ID": "99418977-29cf-48a5-826c-565259009451",
                  "ElementType": 12,
                  "Parent": "$stats",
                  "Name": "signal",
                  "Value": "92",
                  "Props": {}
                },
                "uptime": {
                  "ID": "7a9485b8-488d-4332-ac65-4fe65d5d6c06",
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
            "$name": {
              "ID": "a8ef24c2-8652-469d-b4ae-a245ca8bb80e",
              "ElementType": 14,
              "Parent": "OutsideMonitor",
              "Name": "$name",
              "Attrs": {
                "Ambient": {
                  "ID": "7aa3befd-e072-41e8-b846-6141a4b0b17e",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "9089d4da-b410-4a86-81a7-b1c4f6684e40",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "0965d58a-11f1-4438-82e0-97f75e88a29e",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$properties": {
              "ID": "09627b69-105c-46f3-bdda-90146f65e82d",
              "ElementType": 14,
              "Parent": "OutsideMonitor",
              "Name": "$properties",
              "Attrs": {
                "Ambient": {
                  "ID": "a26c870e-961f-4448-9975-2894651c8065",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "957c9786-0a73-4101-b01f-01ef9b3b450b",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "83838bcf-109d-4ce0-b5f2-62c4ab9bff6b",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "$type": {
              "ID": "d9ad823b-76d3-446e-a1db-a4f1ed33c907",
              "ElementType": 14,
              "Parent": "OutsideMonitor",
              "Name": "$type",
              "Attrs": {
                "Ambient": {
                  "ID": "4fce8ff4-8aa2-4b13-a5a5-c87a0a03bfb2",
                  "ElementType": 15,
                  "Parent": "Ambient",
                  "Name": "Ambient",
                  "Value": ""
                },
                "Presence": {
                  "ID": "e4c5f1e1-e6d3-4ad4-925b-4c7f4b999e4d",
                  "ElementType": 15,
                  "Parent": "Presence",
                  "Name": "Presence",
                  "Value": ""
                },
                "hardware": {
                  "ID": "303603eb-8848-4cc5-b29c-3abca8ebce78",
                  "ElementType": 15,
                  "Parent": "hardware",
                  "Name": "hardware",
                  "Value": ""
                }
              },
              "Props": {}
            },
            "humidity": {
              "ID": "bc1b88b9-7e07-4d91-a671-aa225737f4ed",
              "ElementType": 14,
              "Parent": "OutsideMonitor",
              "Name": "humidity",
              "Attrs": {},
              "Props": {
                "humidity": {
                  "ID": "ba05fb46-20da-4fcf-acb2-130d82fd546f",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "humidity",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "dc8c7c8b-ba89-4716-b9b8-293b7081a4f8",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "a44cbef8-01c4-4726-a77d-3ef38482b235",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$name",
                      "Value": "Humidity"
                    },
                    "$unit": {
                      "ID": "ffb0552f-32e4-4dd1-9e86-b5cd6bb0cf70",
                      "ElementType": 17,
                      "Parent": "humidity",
                      "Name": "$unit",
                      "Value": "%rH"
                    }
                  }
                }
              }
            },
            "mac": {
              "ID": "8e6da9bb-b2b7-4995-a52f-b815a1fd6523",
              "ElementType": 14,
              "Parent": "OutsideMonitor",
              "Name": "mac",
              "Attrs": {},
              "Props": {
                "mac": {
                  "ID": "77499644-c52a-4978-bb25-0c859489dbd3",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "mac",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "9c08467e-8a00-491b-b053-2eacad792461",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$datatype",
                      "Value": "sring"
                    },
                    "$name": {
                      "ID": "3245e11c-7fc9-4d53-990f-71c9e717d803",
                      "ElementType": 17,
                      "Parent": "mac",
                      "Name": "$name",
                      "Value": "mac"
                    }
                  }
                }
              }
            },
            "motion": {
              "ID": "17abd834-c2fa-44f5-aa56-d28fc723717e",
              "ElementType": 14,
              "Parent": "OutsideMonitor",
              "Name": "motion",
              "Attrs": {},
              "Props": {
                "motion": {
                  "ID": "4ca8deaf-48a7-4597-9424-499618e583c6",
                  "ElementType": 16,
                  "Parent": "Presence",
                  "Name": "motion",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "fa9bd8a8-561d-49e2-aaf3-33cda2d3b486",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$datatype",
                      "Value": "enum"
                    },
                    "$format": {
                      "ID": "c8d56e63-7b64-4317-8adc-c6f244fda0a0",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$format",
                      "Value": "OPEN,CLOSED"
                    },
                    "$name": {
                      "ID": "80620b02-4b65-4451-8c6e-d305754f002e",
                      "ElementType": 17,
                      "Parent": "motion",
                      "Name": "$name",
                      "Value": "Motion"
                    }
                  }
                }
              }
            },
            "resetReason": {
              "ID": "01927b29-c05c-4839-b5f0-42593a1fcedb",
              "ElementType": 14,
              "Parent": "OutsideMonitor",
              "Name": "resetReason",
              "Attrs": {},
              "Props": {
                "resetReason": {
                  "ID": "8bc7a679-8947-4e2c-9590-8c1a72d0b894",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "resetReason",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "a9e6a102-1fcc-4a14-8f46-75cc5861ecd5",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$datatype",
                      "Value": "string"
                    },
                    "$name": {
                      "ID": "511dc0d8-1542-49fd-86ea-6a9912ed3628",
                      "ElementType": 17,
                      "Parent": "resetReason",
                      "Name": "$name",
                      "Value": "Last Reset Reason"
                    }
                  }
                }
              }
            },
            "signal": {
              "ID": "19e72bf6-adee-44c3-9cdf-2cfb437277ae",
              "ElementType": 14,
              "Parent": "OutsideMonitor",
              "Name": "signal",
              "Attrs": {},
              "Props": {
                "signal": {
                  "ID": "460d1015-4b12-4803-ab27-55621ae20f51",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "signal",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "8ab7f77b-6682-45bf-8675-4a19d3f76bac",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$datatype",
                      "Value": "integer"
                    },
                    "$name": {
                      "ID": "647593a5-80df-4da3-b397-675e0eacf6ff",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$name",
                      "Value": "RSSI"
                    },
                    "$unit": {
                      "ID": "5ed6f14d-a3fe-454c-8c32-ebac83d3ad62",
                      "ElementType": 17,
                      "Parent": "signal",
                      "Name": "$unit",
                      "Value": "dBm"
                    }
                  }
                }
              }
            },
            "temperature": {
              "ID": "e3b5beb9-96f2-4342-937c-80cf473d51d2",
              "ElementType": 14,
              "Parent": "OutsideMonitor",
              "Name": "temperature",
              "Attrs": {},
              "Props": {
                "temperature": {
                  "ID": "e9adb405-8719-47a8-a0a0-1537728b5ebc",
                  "ElementType": 16,
                  "Parent": "Ambient",
                  "Name": "temperature",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "f6754c4d-8d68-4cf5-b59e-fab16095e3e3",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "7aecc986-7a2b-4adc-8def-13f392a205e7",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$name",
                      "Value": "Temperature"
                    },
                    "$unit": {
                      "ID": "666bd6ad-3d3e-4d4a-a907-d4bc7d60bb6f",
                      "ElementType": 17,
                      "Parent": "temperature",
                      "Name": "$unit",
                      "Value": "°F"
                    }
                  }
                }
              }
            },
            "voltage": {
              "ID": "a46f2a7d-61dc-43e4-9419-f4e7faa24b6b",
              "ElementType": 14,
              "Parent": "OutsideMonitor",
              "Name": "voltage",
              "Attrs": {},
              "Props": {
                "voltage": {
                  "ID": "b117e966-a9a5-4546-a0e7-1831fee3c5d3",
                  "ElementType": 16,
                  "Parent": "hardware",
                  "Name": "voltage",
                  "Value": "",
                  "Attrs": {
                    "$datatype": {
                      "ID": "64a1d2b6-00eb-4ed1-b2ec-b71bb90293cf",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$datatype",
                      "Value": "float"
                    },
                    "$name": {
                      "ID": "3a7554f5-2d62-4e78-8446-a113fe0c5454",
                      "ElementType": 17,
                      "Parent": "voltage",
                      "Name": "$name",
                      "Value": "3.3V Supply"
                    },
                    "$unit": {
                      "ID": "b5f9e2db-98db-4134-a638-54c9ae796779",
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