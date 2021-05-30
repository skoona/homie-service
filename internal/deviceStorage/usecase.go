package deviceStorage

/*
  deviceStorage/usecase-stream.go:

  DeviceSource Service Implementation
  Utilities to List current Database
*/

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	bolt "github.com/boltdb/bolt"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
)

/**
 * ScheduleStore(schedule Schedule)
 */
func (dbR *dbRepo) ScheduleStore(d dc.Schedule) map[dc.EID]dc.Schedule {
	level.Debug(dbR.logger).Log("event", "Calling ScheduleStore()", "dm", d.String())
	// when input is present, store new
	// when input is empty or nil, render
	schedMap := make(map[dc.EID]dc.Schedule, 0)
	return schedMap
}

/**
 * Store()
 *
 * Repository Implementation
 *
 * Save the DeviceMessage to the Store using Devicename/Bucket -> (Topic:value)
 * Save if value > nil,
 * Delete topic if present
 * Deletebucket if no topic exists
 */
func (dbR *dbRepo) Store(d dc.DeviceMessage) error {
	var err error

	//     0         1           2        3              4                         5
	// sknSensors/device/node/version: 3.0.0
	//  * homie / device ID / $device-attribute / device-attribute-Property / device-attribute-Property-Property
	//				Bucket(Bucket(Bucket(bucket(property:value))))
	//  * homie / device ID / $device-attribute / device-attribute-Property
	//				Bucket(Bucket(bucket(property:value)))
	//  * homie / device ID / $device-attribute
	//              Bucket(attribute:value)
	//  * homie / device ID / node ID / property ID / $property-attribute
	//				Bucket(Bucket(Bucket(attribute:value)))
	//  * homie / device ID / node ID / $node-attribute
	//				Bucket(Bucket(attribute:value))
	//  * homie / device ID / node ID / property ID
	//				Bucket(Bucket(Bucket(property:value)))

	level.Debug(dbR.logger).Log("event", "Calling Store()", "dm", d.String())

	if nil == d.Value || len(d.Value) == 0 { // Delete actions
		err = dbR.db.Update(func(tx *bolt.Tx) error {
			var err error

			b := tx.Bucket(d.NetworkID)
			if b == nil {
				return fmt.Errorf("[WARN] Network Not Found!: %s, error: %s", d.NetworkID, err.Error())
			}

			b = b.Bucket(d.DeviceID)
			if b == nil {
				return fmt.Errorf("[WARN] Device Not Found!: %s, error: %s", d.DeviceID, err.Error())
			}

			/* Devices */

			// X/D/A/P/P
			if len(d.NodeID) == 0 && len(d.PPropertyID) > 0 {
				b = b.Bucket(d.AttributeID)
				if b == nil {
					return fmt.Errorf("[WARN] Attribute Not Found!: %s, error: %s", d.AttributeID, err.Error())
				}
				b = b.Bucket(d.PropertyID)
				if b == nil {
					return fmt.Errorf("[WARN] Property Not Found!: %s, error: %s", d.PropertyID, err.Error())
				}

				return b.DeleteBucket(d.PPropertyID)
			}

			// X/D/A/P
			if len(d.NodeID) == 0 && len(d.PPropertyID) == 0 {
				b = b.Bucket(d.AttributeID)
				if b == nil {
					return fmt.Errorf("[WARN] Atribute Not Found!: %s, error: %s", d.AttributeID, err.Error())
				}
				return b.DeleteBucket(d.PropertyID)
			}

			// X/D/A
			if len(d.NodeID) == 0 && len(d.PropertyID) == 0 { // X/D/A
				return b.Delete(d.AttributeID)
			}

			/* Nodes */
			b = b.Bucket(d.NodeID)
			if b == nil {
				return fmt.Errorf("[WARN] Node Not Found!: %s, error: %s", d.NodeID, err.Error())
			}

			// X/D/N/P/A
			if len(d.NodeID) > 0 && len(d.PropertyID) > 0 && len(d.AttributeID) > 0 {
				b = b.Bucket(d.PropertyID)
				if b == nil {
					return fmt.Errorf("[WARN] Property Not Found!: %s, error: %s", d.PropertyID, err.Error())
				}

				return b.Delete(d.AttributeID)
			}

			// X/D/N/A
			if len(d.NodeID) > 0 && len(d.AttributeID) > 0 {
				return b.Delete(d.AttributeID)
			}

			// X/D/N/P
			if len(d.NodeID) > 0 && len(d.PropertyID) > 0 {
				return b.DeleteBucket(d.PropertyID)
			}

			return err
		})

	} else { // Create Actions

		err = dbR.db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists(d.NetworkID)
			if err != nil {
				return fmt.Errorf("[WARN] Network cannot be created or found!: %s, error: %s", d.NetworkID, err.Error())
			}

			b, err = b.CreateBucketIfNotExists(d.DeviceID)
			if err != nil {
				return fmt.Errorf("[WARN] Device cannot be created or found!: %s, error: %s", d.DeviceID, err.Error())
			}

			if len(d.NodeID) == 0 && len(d.PPropertyID) > 0 { // device attrs-property-property
				b, err = b.CreateBucketIfNotExists(d.AttributeID)
				if err != nil {
					return fmt.Errorf("[WARN] Attribute(x/d/a/p/p) cannot be created or found!: %s, error: %s", d.AttributeID, err.Error())
				}

				b, err = b.CreateBucketIfNotExists(d.PropertyID)
				if err != nil {
					return fmt.Errorf("[WARN] Property(x/d/a/p/p) cannot be created or found!: %s, error: %s", d.PropertyID, err.Error())
				}

				b, err = b.CreateBucketIfNotExists(d.PPropertyID)
				if err != nil {
					return fmt.Errorf("[WARN] PropertyProperty(x/d/a/p/p) cannot be created or found!: %s, error: %s", d.PPropertyID, err.Error())
				}

				err = b.Put(d.PPropertyID, d.Value)
				return err
			}

			if len(d.NodeID) == 0 && len(d.PropertyID) > 0 { // device attr-property
				b, err = b.CreateBucketIfNotExists(d.AttributeID)
				if err != nil {
					return fmt.Errorf("[WARN] Attribute(x/d/a/p) cannot be created or found!: %s, error: %s", d.AttributeID, err.Error())
				}

				b, err = b.CreateBucketIfNotExists(d.PropertyID)
				if err != nil {
					return fmt.Errorf("[WARN] Property(x/d/a/p) cannot be created or found!: %s, error: %s", d.PropertyID, err.Error())
				}

				err = b.Put(d.PropertyID, d.Value)
				return err
			}

			if len(d.NodeID) == 0 && len(d.AttributeID) > 0 { // device attrs
				err = b.Put(d.AttributeID, d.Value)
				return err
			}

			// x/d/n
			if len(d.NodeID) > 0 {
				b, err = b.CreateBucketIfNotExists(d.NodeID)
				if err != nil {
					return fmt.Errorf("[WARN] Node cannot be created or found!: %s, error: %s", d.NodeID, err.Error())
				}
			} else {
				return fmt.Errorf("ALERT Unknown(node) Message!: %s, error: %s", d.String(), err.Error())
			}

			// x/d/n/p/a
			if len(d.PropertyID) > 0 && len(d.AttributeID) > 0 { // node property attr
				b, err = b.CreateBucketIfNotExists(d.PropertyID)
				if err != nil {
					return fmt.Errorf("[WARN] Property(x/d/n/p/a) cannot be created or found!: %s, error: %s", d.PropertyID, err.Error())
				}

				return b.Put(d.AttributeID, d.Value)
			}

			// x/d/n/p
			if len(d.PropertyID) > 0 && len(d.AttributeID) == 0 { // Node Property
				b, err = b.CreateBucketIfNotExists(d.PropertyID)
				if err != nil {
					return fmt.Errorf("[WARN] Property(x/d/n/p) cannot be created or found!: %s, error: %s", d.PropertyID, err.Error())
				}

				return b.Put(d.PropertyID, d.Value)
			}

			// x/d/n/a
			if len(d.PropertyID) == 0 && len(d.AttributeID) > 0 { // Node Attr
				return b.Put(d.AttributeID, d.Value)
			}

			return err
		})
	}

	if err != nil {
		level.Error(dbR.logger).Log("Alert", "Store() Update Failed", "error", err.Error(), "dm", d.String())
	}

	return err
}

/*
 * Returns string array of networks names (top level buckets) */
func networkList(db *bolt.DB) []string {
	networks := []string{}

	err := db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			networks = append(networks, string(name))
			return nil
		})
	})
	if err != nil {
		fmt.Printf("[WARN] Network rendering Failed: %v", err.Error())
	}

	return networks
}

/*
 * Returns string array of device names (buckets) */
func deviceList(db *bolt.DB, networkName string) []string {
	devices := []string{}

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(networkName)) // Network Level
		if b == nil {
			return fmt.Errorf("Network Not Found!: %s", networkName)
		}
		return b.ForEach(func(k, v []byte) error {
			if v == nil { // nodes have nil values
				devices = append(devices, string(k))
			}
			return nil
		})
	})
	if err != nil {
		fmt.Printf("[WARN] Network::Device rendering Failed: %v", err.Error())
	}

	return devices
}

/*
 * Returns string array a device's properties (topic:value) */
func deviceDetails(db *bolt.DB, networkName, deviceName string) (map[string]string, []string) {
	details := map[string]string{}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(networkName)) // Network Level
		if b == nil {
			return fmt.Errorf("Network Not Found!: %s", networkName)
		}
		b = b.Bucket([]byte(deviceName)) // Device Level
		if b == nil {
			return fmt.Errorf("Network Not Found!: %s", deviceName)
		}

		err := b.ForEach(func(k, v []byte) error {
			if nil == v {
				n := b.Bucket(k) // Node Level
				if n == nil {
					return nil
				}

				err := n.ForEach(func(kk, vv []byte) error {
					if nil == vv {
						p := n.Bucket(kk) // Property
						if p == nil {
							return nil
						}

						err := p.ForEach(func(kkk, vvv []byte) error {
							details["["+string(k)+"]["+string(kk)+"]"+string(kkk)] = string(vvv)
							return nil
						})

						return err
					}

					details["["+string(k)+"]"+string(kk)] = string(vv)

					return nil
				})

				return err

			}

			details[string(k)] = string(v)

			return nil
		})
		if err != nil {
			err = fmt.Errorf("[WARN] Detailed rendering failed: %v", err.Error())
		}
		return err
	})

	// orderedKeys :=
	keys := make([]string, 0, len(details))
	for k, _ := range details {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return details, keys
}

/*
 * DBStatsAsJSON
 * Outputs boltDB Database Stats
 */
func dbStatsAsJSON() string {
	sts := dbs.db.Stats()
	var res string

	json, _ := json.MarshalIndent(sts, "", "    ")
	res = string(json)

	return res
}

/*
 * ListHomieDBCollection
 * List the Devices Found/Recorded
 */
func ListHomieDB() {
	time.Sleep(1 * time.Second) // slow the pace
	if networks := networkList(dbs.db); len(networks) > 0 {
		for nets, network := range networks {
			fmt.Printf("[%d] Network: %s\n", nets, network)
			if devices := deviceList(dbs.db, network); len(devices) > 0 {
				for idx, device := range devices {
					fmt.Printf("\t[%d] Device: %s\n", idx, device)
					if detail, orderedKeys := deviceDetails(dbs.db, network, device); len(detail) > 0 {
						for _, k := range orderedKeys {
							v := detail[k]
							fmt.Printf("\t\t %s --> %s\n", k, v)
						}
					}
				}
			}
		}
	}

	// Show bBolt DB Stats
	if stats := dbStatsAsJSON(); len(stats) > 0 {
		fmt.Printf("dbStats=%s\n", stats)
	}
}
