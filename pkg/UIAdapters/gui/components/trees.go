package components

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"sort"
	"strings"
)


func SknDeviceTreeSide(dv *dc.Device, devTreeDetails **map[string][][]string) fyne.CanvasObject {
	data := TreeDataFromDevice(dv)  // map[string][children][kv]string
	*devTreeDetails = &data
	content := widget.NewTree(
		func(uid widget.TreeNodeID) []string { // childUIDs
			children := data
			sort.Strings(children[uid][0])
			return (children[uid][0])
		},
		func(uid widget.TreeNodeID) bool { // isBranch
			haveChildren := data
			return ( len(haveChildren[uid][0]) > 0 )
		},
		func(_ bool) fyne.CanvasObject { // create
			return keyValueTemplate("some kinda key", "some kinda value")
		},
		func(uid widget.TreeNodeID, _ bool, template fyne.CanvasObject) { // update
			source := data
			template.(*fyne.Container).Objects[0].(*widget.Label).SetText(source[uid][1][0])
			template.(*fyne.Container).Objects[1].(*widget.Label).SetText(source[uid][1][1])
		})

	return content
}

func keyValueTemplate(key, value string) fyne.CanvasObject {
	k := widget.NewLabel(key)
	v := widget.NewLabel(value)
	v.Alignment = fyne.TextAlignCenter
	return container.NewHBox(k, v)
}


// TreeDataFromDevice()
// decode device into topics
// -/D/A/P/P
// -/D/A/P
// -/D/A
// -/D/N/P/A
// -/D/N/P
// -/D/N/A
// map[string][]string
func TreeDataFromDevice(dv *dc.Device) map[string][][]string {
	var (
		ele string
		dapp = map[string][]string{}
		dap  = map[string][]string{}
		da   = map[string][]string{}
		dn   = map[string][]string{}
		dna  = map[string][]string{}
		dnp  = map[string][]string{}
		dnpa = map[string][]string{}
	)

	// unpack device attrs
	for n, v := range dv.Attrs {  // x/d/a
		for nn, vv := range v.Props { // x/d/a/p
			for nnn, vvv := range vv.Props { // x/d/a/p/p
				ele = fmt.Sprintf("%s/%s/%s/%s/%s", dv.Parent, dv.Name, n, nn, nnn)
				dapp[ele] = []string{ string(nnn), string(vvv.Value)}
			}
			ele = fmt.Sprintf("%s/%s/%s/%s", dv.Parent, dv.Name, n, nn)
			dap[ele] = []string{ string(nn), string(vv.Value)}
		}
		ele = fmt.Sprintf("%s/%s/%s", dv.Parent, dv.Name, n)
		da[ele] = []string{ string(n), string(v.Value)}
	}
	// unpack device nodes
	for n, v := range dv.Nodes {  // x/d/n
		dn[fmt.Sprintf("%s/%s/%s", dv.Parent, dv.Name, n)] = []string{string(n), "Node"}
		for nn, vv := range v.Props { // x/d/n/p
			for nnn, vvv := range vv.Attrs { // x/d/n/p/a
				ele = fmt.Sprintf("%s/%s/%s/%s/%s", dv.Parent, dv.Name, n, nn, nnn)
				dnpa[ele] = []string{ string(nnn), string(vvv.Value)}
			}
			ele = fmt.Sprintf("%s/%s/%s/%s", dv.Parent, dv.Name, n, nn)
			dnp[ele] = []string{ string(nn), string(vv.Value)}
		}
		for nn, vv := range v.Attrs { // x/d/n/a
			ele = fmt.Sprintf("%s/%s/%s/%s", dv.Parent, dv.Name, n, nn)
			dna[ele] = []string{ string(nn), string(vv.Value)}
		}
	}

	// device                      DA, DN(s)
	// device attrs                self, DAP
	// device attr props           self, DAPP
	// device attr prop props      self
	// device node                 DNA, DNP
	// device node attrs           self
	// device node props           self, DNPA
	// device node prop attrs      self
	// map[string][][]string         data model
	// map[<long.name>][0] = []string   -- children logn/name keys
	// map[<long.name>][1] = []         -- self  k,v

	tree := make(map[string][][]string, 32)
	tree[""] = make([][]string,2)

	tree[""][0] = append(tree[""][0], dv.Name)
	tree[""][1] = []string{dv.Parent, dv.Name}
	tree[dv.Name] = make([][]string,2)
	tree[dv.Name][1] = []string{dv.Name, dv.Parent}
	// root
	for k, v := range da {
		tree[k] = make([][]string,2)
		tree[dv.Name][0] = append(tree[dv.Name][0], k)
		tree[k][1] = v
		// da's  include DA, DAP
		for x, y := range dap {
			if strings.HasPrefix(string(x),string(k)) {
				tree[k][0] = append(tree[k][0], x)
				tree[x] = make([][]string,2)
				tree[x][1] = y
				for r, s := range dapp {
					if strings.HasPrefix(string(r),string(x)) {
						tree[x][0] = append(tree[x][0], r)
						tree[r] = make([][]string,2)
						tree[r][1] = s
					}
				}
			}
		}
	}
	// DN's
	for k, v := range dn {
		tree[dv.Name][0] = append(tree[dv.Name][0], k)
		tree[k] = make([][]string,2)
		tree[k][1] = v
		for x, y := range dna {
			if strings.HasPrefix(string(x),string(k)) {
				tree[k][0] = append(tree[k][0], x)
				tree[x] = make([][]string,2)
				tree[x][1] = y
			}
		}
		for x, y := range dnp {
			if strings.HasPrefix(x,k) {
				tree[k][0] = append(tree[k][0], x)
				tree[x] = make([][]string,2)
				tree[x][1] = y
				for a, b := range dnpa {
					if strings.HasPrefix(string(a),string(x)) {
						tree[x][0] = append(tree[x][0], a)
						tree[a] = make([][]string,2)
						tree[a][1] = b
					}
				}
			}
		}
	}
	//fmt.Println("TREE MAP", tree)

	return tree
}


// {
// "FamilyRoom": {
//          "ID": "2f63939498209cf4584b6e82954d2407",
//          "Title": "",
//          "ElementType": 10,
//          "OTAEnabled": true,
//          "Parent": "sknSensors",
//          "Name": "FamilyRoom",
//          "Attrs": {
//            "$fw": {
//              "ID": "dc9cdbe2-77e3-417d-8656-61b44afcf93c",
//              "ElementType": 11,
//              "Parent": "FamilyRoom",
//              "Name": "$fw",
//              "Value": "",
//              "Props": {
//                "checksum": {
//                  "ID": "14859fe7-a81d-4cd5-99d5-a545dbb4d14a",
//                  "ElementType": 12,
//                  "Parent": "$fw",
//                  "Name": "checksum",
//                  "Value": "5310fe4c8093a077952bd83a72ac83d9",
//                  "Props": {}
//                },
//                "name": {
//                  "ID": "68c54809-314a-43fc-a59c-8cc54bd5757a",
//                  "ElementType": 12,
//                  "Parent": "$fw",
//                  "Name": "name",
//                  "Value": "Monitor-DHT-RCWL-Metrics",
//                  "Props": {}
//                },
//                "version": {
//                  "ID": "541e6eda-a830-47bc-a3b0-1c430b11d217",
//                  "ElementType": 12,
//                  "Parent": "$fw",
//                  "Name": "version",
//                  "Value": "2.0.1",
//                  "Props": {}
//                }
//              }
//            },
//            "$homie": {
//              "ID": "b6b91fd3-e00b-4ca1-b7e2-2ad906af71d6",
//              "ElementType": 11,
//              "Parent": "FamilyRoom",
//              "Name": "$homie",
//              "Value": "3.0.1",
//              "Props": {}
//            },
//            "$implementation": {
//              "ID": "62f9442a-c3a5-4419-a41f-e0683a354c53",
//              "ElementType": 11,
//              "Parent": "FamilyRoom",
//              "Name": "$implementation",
//              "Value": "esp8266",
//              "Props": {
//                "config": {
//                  "ID": "45e5108d-1ee8-4b6a-ab27-3f0176d763f4",
//                  "ElementType": 12,
//                  "Parent": "$implementation",
//                  "Name": "config",
//                  "Value": "{\"name\":\"FamilyRoom\",\"device_id\":\"FamilyRoom\",\"device_stats_interval\":180,\"wifi\":{\"ssid\":\"SFNSS1-24G\"},\"mqtt\":{\"host\":\"openhab.skoona.net\",\"port\":1883,\"base_topic\":\"sknSensors/\",\"auth\":true},\"ota\":{\"enabled\":true},\"settings\":{\"sensorInterval\":900,\"motionHoldInterval\":60}}",
//                  "Props": {}
//                },
//                "ota": {
//                  "ID": "970fd897-f868-4ebc-bd07-1a9d3de5ad19",
//                  "ElementType": 12,
//                  "Parent": "$implementation",
//                  "Name": "ota",
//                  "Value": "",
//                  "Props": {
//                    "enabled": {
//                      "ID": "5d7d0586-52f2-421b-af7c-ea2be48b1d29",
//                      "ElementType": 13,
//                      "Parent": "ota",
//                      "Name": "enabled",
//                      "Value": "true"
//                    }
//                  }
//                },
//                "version": {
//                  "ID": "37fefd2b-9a84-4040-a054-72ee68d275ce",
//                  "ElementType": 12,
//                  "Parent": "$implementation",
//                  "Name": "version",
//                  "Value": "3.0.0",
//                  "Props": {}
//                }
//              }
//            },
//            "$localip": {
//              "ID": "efd87f93-6ca4-4e96-96d5-2be512bb831e",
//              "ElementType": 11,
//              "Parent": "FamilyRoom",
//              "Name": "$localip",
//              "Value": "10.100.1.166",
//              "Props": {}
//            },
//            "$mac": {
//              "ID": "4e50eda4-93ed-4ba2-9573-76fa3d4dde05",
//              "ElementType": 11,
//              "Parent": "FamilyRoom",
//              "Name": "$mac",
//              "Value": "BC:DD:C2:E5:C4:24",
//              "Props": {}
//            },
//            "$name": {
//              "ID": "567825ce-5fd5-4bab-8b2c-692f0248b2b9",
//              "ElementType": 11,
//              "Parent": "FamilyRoom",
//              "Name": "$name",
//              "Value": "FamilyRoom",
//              "Props": {}
//            },
//            "$nodes": {
//              "ID": "b74d62d4-53c9-4ff4-b71c-60d326938439",
//              "ElementType": 11,
//              "Parent": "FamilyRoom",
//              "Name": "$nodes",
//              "Value": "Ambient,Presence,hardware",
//              "Props": {}
//            },
//            "$state": {
//              "ID": "84f773d5-2502-47e3-aaeb-c5d590b75ae3",
//              "ElementType": 11,
//              "Parent": "FamilyRoom",
//              "Name": "$state",
//              "Value": "ready",
//              "Props": {}
//            },
//            "$stats": {
//              "ID": "4c535e50-c761-4774-b0ee-b4ab616f32dd",
//              "ElementType": 11,
//              "Parent": "FamilyRoom",
//              "Name": "$stats",
//              "Value": "uptime",
//              "Props": {
//                "interval": {
//                  "ID": "ab5e76d7-828a-4dbd-8a9d-c425363811ac",
//                  "ElementType": 12,
//                  "Parent": "$stats",
//                  "Name": "interval",
//                  "Value": "185",
//                  "Props": {}
//                },
//                "signal": {
//                  "ID": "bb35d937-69fe-4a0c-8049-a2f1ce2a1a95",
//                  "ElementType": 12,
//                  "Parent": "$stats",
//                  "Name": "signal",
//                  "Value": "90",
//                  "Props": {}
//                },
//                "uptime": {
//                  "ID": "92831ca1-cf3c-4a30-b877-664eaf7ff110",
//                  "ElementType": 12,
//                  "Parent": "$stats",
//                  "Name": "uptime",
//                  "Value": "10205199",
//                  "Props": {}
//                }
//              }
//            }
//          },
//          "Nodes": {
//            "Ambient": {
//              "ID": "030b738a-30c4-45cc-a27e-39bf860f5b4d",
//              "ElementType": 14,
//              "Parent": "FamilyRoom",
//              "Name": "Ambient",
//              "Attrs": {
//                "$name": {
//                  "ID": "ab4ced5f-b7cf-4dbd-a74a-0995639aff1c",
//                  "ElementType": 15,
//                  "Parent": "Ambient",
//                  "Name": "$name",
//                  "Value": "Temperature and Humidity Sensor"
//                },
//                "$properties": {
//                  "ID": "66467e6f-0c0c-4197-8c80-e2266f0a612c",
//                  "ElementType": 15,
//                  "Parent": "Ambient",
//                  "Name": "$properties",
//                  "Value": "humidity,temperature"
//                },
//                "$type": {
//                  "ID": "7e05a423-0a8b-49dd-888b-061de121613a",
//                  "ElementType": 15,
//                  "Parent": "Ambient",
//                  "Name": "$type",
//                  "Value": "sensor"
//                }
//              },
//              "Props": {
//                "humidity": {
//                  "ID": "38d27600-9cd6-4f88-91b2-2933117f8d4a",
//                  "ElementType": 16,
//                  "Parent": "Ambient",
//                  "Name": "humidity",
//                  "Value": "",
//                  "Attrs": {
//                    "$datatype": {
//                      "ID": "ec9e144f-dfdf-4b21-90b3-3d30cbe99b36",
//                      "ElementType": 17,
//                      "Parent": "humidity",
//                      "Name": "$datatype",
//                      "Value": "float"
//                    },
//                    "$name": {
//                      "ID": "fa7075be-7cc9-4056-b4f1-2a66eba890c2",
//                      "ElementType": 17,
//                      "Parent": "humidity",
//                      "Name": "$name",
//                      "Value": "Humidity"
//                    },
//                    "$unit": {
//                      "ID": "91a59ec3-780d-485d-92f6-e6b4ad0f1e3d",
//                      "ElementType": 17,
//                      "Parent": "humidity",
//                      "Name": "$unit",
//                      "Value": "%rH"
//                    }
//                  }
//                },
//                "temperature": {
//                  "ID": "02b91dd4-0a60-4dcf-a6b0-837be78bd969",
//                  "ElementType": 16,
//                  "Parent": "Ambient",
//                  "Name": "temperature",
//                  "Value": "",
//                  "Attrs": {
//                    "$datatype": {
//                      "ID": "2ccc7d53-424a-4a40-ab65-60434ead7c79",
//                      "ElementType": 17,
//                      "Parent": "temperature",
//                      "Name": "$datatype",
//                      "Value": "float"
//                    },
//                    "$name": {
//                      "ID": "9ee18d59-42fe-466a-8773-11de2ebdc117",
//                      "ElementType": 17,
//                      "Parent": "temperature",
//                      "Name": "$name",
//                      "Value": "Temperature"
//                    },
//                    "$unit": {
//                      "ID": "5e719ca6-c059-48c0-8eea-7f2a0d3fa175",
//                      "ElementType": 17,
//                      "Parent": "temperature",
//                      "Name": "$unit",
//                      "Value": "??F"
//                    }
//                  }
//                }
//              }
//            },
//            "Presence": {
//              "ID": "eaf342c5-7039-464e-9d60-cb103858bc79",
//              "ElementType": 14,
//              "Parent": "FamilyRoom",
//              "Name": "Presence",
//              "Attrs": {
//                "$name": {
//                  "ID": "dc042ecc-0120-43e8-8d51-9688ea3b3af0",
//                  "ElementType": 15,
//                  "Parent": "Presence",
//                  "Name": "$name",
//                  "Value": "Motion Sensor"
//                },
//                "$properties": {
//                  "ID": "9fcf6783-9dce-49fd-96ca-c263ca38bad9",
//                  "ElementType": 15,
//                  "Parent": "Presence",
//                  "Name": "$properties",
//                  "Value": "motion"
//                },
//                "$type": {
//                  "ID": "7e233c4d-21bd-43a8-9ebc-7bfc966cd04e",
//                  "ElementType": 15,
//                  "Parent": "Presence",
//                  "Name": "$type",
//                  "Value": "contact"
//                }
//              },
//              "Props": {
//                "motion": {
//                  "ID": "61a39845-cff1-48e6-97c6-17af1c209010",
//                  "ElementType": 16,
//                  "Parent": "Presence",
//                  "Name": "motion",
//                  "Value": "",
//                  "Attrs": {
//                    "$datatype": {
//                      "ID": "c1ef13cf-67df-4db8-8129-2a85e29c809b",
//                      "ElementType": 17,
//                      "Parent": "motion",
//                      "Name": "$datatype",
//                      "Value": "enum"
//                    },
//                    "$format": {
//                      "ID": "aabbf322-d4a0-4ea3-abb2-bd8d2baffe13",
//                      "ElementType": 17,
//                      "Parent": "motion",
//                      "Name": "$format",
//                      "Value": "OPEN,CLOSED"
//                    },
//                    "$name": {
//                      "ID": "ffec0016-6327-40f7-9124-67f6607e0c0f",
//                      "ElementType": 17,
//                      "Parent": "motion",
//                      "Name": "$name",
//                      "Value": "Motion"
//                    }
//                  }
//                }
//              }
//            },
//            "hardware": {
//              "ID": "ae28f6c6-e451-4af4-b36b-e88225c307ed",
//              "ElementType": 14,
//              "Parent": "FamilyRoom",
//              "Name": "hardware",
//              "Attrs": {
//                "$name": {
//                  "ID": "f6b9a688-d9cd-49bb-bc19-4dc5e3d095a4",
//                  "ElementType": 15,
//                  "Parent": "hardware",
//                  "Name": "$name",
//                  "Value": "Device Info"
//                },
//                "$properties": {
//                  "ID": "c6cb6956-77d7-4be3-a096-92dbec0ebae7",
//                  "ElementType": 15,
//                  "Parent": "hardware",
//                  "Name": "$properties",
//                  "Value": "signal,mac,resetReason,voltage"
//                },
//                "$type": {
//                  "ID": "46215873-933c-4061-8130-8c94589e89fa",
//                  "ElementType": 15,
//                  "Parent": "hardware",
//                  "Name": "$type",
//                  "Value": "sensor"
//                }
//              },
//              "Props": {
//                "mac": {
//                  "ID": "0a6301bc-dad4-4378-8ac4-97e9af302d81",
//                  "ElementType": 16,
//                  "Parent": "hardware",
//                  "Name": "mac",
//                  "Value": "",
//                  "Attrs": {
//                    "$datatype": {
//                      "ID": "2f8710a4-d1b0-4ec2-b1ae-60b1173d582e",
//                      "ElementType": 17,
//                      "Parent": "mac",
//                      "Name": "$datatype",
//                      "Value": "sring"
//                    },
//                    "$name": {
//                      "ID": "923a6c2c-e3cd-4483-8f58-b85d1a65257f",
//                      "ElementType": 17,
//                      "Parent": "mac",
//                      "Name": "$name",
//                      "Value": "mac"
//                    }
//                  }
//                },
//                "resetReason": {
//                  "ID": "d72d2750-62bc-48c0-bb89-085bc452405d",
//                  "ElementType": 16,
//                  "Parent": "hardware",
//                  "Name": "resetReason",
//                  "Value": "",
//                  "Attrs": {
//                    "$datatype": {
//                      "ID": "194a8ac3-8a5d-47d6-b340-b16944324504",
//                      "ElementType": 17,
//                      "Parent": "resetReason",
//                      "Name": "$datatype",
//                      "Value": "string"
//                    },
//                    "$name": {
//                      "ID": "d2b5fd36-7b8a-41c8-84bd-0f6bf196095a",
//                      "ElementType": 17,
//                      "Parent": "resetReason",
//                      "Name": "$name",
//                      "Value": "Last Reset Reason"
//                    }
//                  }
//                },
//                "signal": {
//                  "ID": "c7e43d70-955f-41ea-9277-cb36321afddb",
//                  "ElementType": 16,
//                  "Parent": "hardware",
//                  "Name": "signal",
//                  "Value": "",
//                  "Attrs": {
//                    "$datatype": {
//                      "ID": "0387801f-321a-4fa3-8749-a63c21983fbe",
//                      "ElementType": 17,
//                      "Parent": "signal",
//                      "Name": "$datatype",
//                      "Value": "integer"
//                    },
//                    "$name": {
//                      "ID": "fb6a4d8f-705a-4aa2-a6f2-7f65381702e7",
//                      "ElementType": 17,
//                      "Parent": "signal",
//                      "Name": "$name",
//                      "Value": "RSSI"
//                    },
//                    "$unit": {
//                      "ID": "8bad9f7f-0e51-4716-8e0d-cefdf8696d4e",
//                      "ElementType": 17,
//                      "Parent": "signal",
//                      "Name": "$unit",
//                      "Value": "dBm"
//                    }
//                  }
//                },
//                "voltage": {
//                  "ID": "066fe8c0-7a95-497c-b173-fb07b9953081",
//                  "ElementType": 16,
//                  "Parent": "hardware",
//                  "Name": "voltage",
//                  "Value": "",
//                  "Attrs": {
//                    "$datatype": {
//                      "ID": "de9a7552-1bbb-4de4-8433-704a5c3a819b",
//                      "ElementType": 17,
//                      "Parent": "voltage",
//                      "Name": "$datatype",
//                      "Value": "float"
//                    },
//                    "$name": {
//                      "ID": "26884c4a-7b77-4bff-a87c-814094264b33",
//                      "ElementType": 17,
//                      "Parent": "voltage",
//                      "Name": "$name",
//                      "Value": "3.3V Supply"
//                    },
//                    "$unit": {
//                      "ID": "67edaff6-8731-4ff9-95b8-2f53069cfd36",
//                      "ElementType": 17,
//                      "Parent": "voltage",
//                      "Name": "$unit",
//                      "Value": "V"
//                    }
//                  }
//                }
//              }
//            }
//          }
//        }
// }
//
//map[
//	"":[[GarageMonitor] [sknSensors GarageMonitor]]
// GarageMonitor:[[sknSensors/GarageMonitor/$name sknSensors/GarageMonitor/$homie sknSensors/GarageMonitor/$mac sknSensors/GarageMonitor/$nodes sknSensors/GarageMonitor/$state sknSensors/GarageMonitor/$localip sknSensors/GarageMonitor/hardware sknSensors/GarageMonitor/Ambient sknSensors/GarageMonitor/Presence] []]
// sknSensors/GarageMonitor/$fw/checksum:[[] [checksum 615fed382ab44bd43fe83508aecac682]]
// sknSensors/GarageMonitor/$fw/name:[[] [name Monitor-SHT31-RCWL-Metrics]]
// sknSensors/GarageMonitor/$fw/version:[[] [version 2.0.0]]
// sknSensors/GarageMonitor/$homie:[[] [$homie 3.0.1]]
// sknSensors/GarageMonitor/$implementation/config:[[] [config {"name":"Garage Monitor","device_id":"GarageMonitor","device_stats_interval":900,"wifi":{"ssid":"SFNSS1-24G"},"mqtt":{"host":"openhab.skoona.net","port":1883,"base_topic":"sknSensors/","auth":true},"ota":{"enabled":true},"settings":{"sensorInterval":900,"motionHoldInterval":60}}]]
// sknSensors/GarageMonitor/$implementation/ota/enabled:[[] [enabled true]]
// sknSensors/GarageMonitor/$implementation/version:[[] [version 3.0.0]]
// sknSensors/GarageMonitor/$localip:[[] [$localip 10.100.1.177]]
// sknSensors/GarageMonitor/$mac:[[] [$mac B4:E6:2D:1B:5C:4D]]
// sknSensors/GarageMonitor/$name:[[] [$name Garage Monitor]]
// sknSensors/GarageMonitor/$nodes:[[] [$nodes Ambient,Presence,hardware]]
// sknSensors/GarageMonitor/$state:[[] [$state ready]]
// sknSensors/GarageMonitor/$stats/interval:[[] [interval 905]]
// sknSensors/GarageMonitor/$stats/signal:[[] [signal 46]]
// sknSensors/GarageMonitor/$stats/uptime:[[] [uptime 12727347]]
// sknSensors/GarageMonitor/Ambient:[[sknSensors/GarageMonitor/Ambient/$type sknSensors/GarageMonitor/Ambient/$name sknSensors/GarageMonitor/Ambient/humidity sknSensors/GarageMonitor/Ambient/temperature] [Ambient]]
// sknSensors/GarageMonitor/Ambient/$name:[[] [$name Temperature and Humidity Sensor]]
// sknSensors/GarageMonitor/Ambient/$type:[[] [$type sensor]]
// sknSensors/GarageMonitor/Ambient/humidity:[[] [humidity 66.43]]
// sknSensors/GarageMonitor/Ambient/temperature:[[] [temperature 81.98]]
// sknSensors/GarageMonitor/Presence:[[sknSensors/GarageMonitor/Presence/motion] [Presence]]
// sknSensors/GarageMonitor/Presence/motion:[[] [motion CLOSED]]
// sknSensors/GarageMonitor/hardware:[[sknSensors/GarageMonitor/hardware/resetReason sknSensors/GarageMonitor/hardware/signal sknSensors/GarageMonitor/hardware/voltage sknSensors/GarageMonitor/hardware/mac] [hardware]]
// sknSensors/GarageMonitor/hardware/mac:[[] [mac B4:E6:2D:1B:5C:4D]]
// sknSensors/GarageMonitor/hardware/resetReason:[[] [resetReason External System]]
// sknSensors/GarageMonitor/hardware/signal:[[] [signal -76]]
// sknSensors/GarageMonitor/hardware/voltage:[[] [voltage 3.03]]
// ]
