package deviceStorage

/*
	deviceStorage/repository.go

	Implement dataDB/LevelDB Connection to DeviceSource
*/

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	bolt "github.com/boltdb/bolt"
	"github.com/go-kit/kit/log"
	ds "github.com/skoona/homie-service/pkg/deviceSource"
)

var dbs dbRepo

type dbRepo struct {
	db     *bolt.DB
	logger log.Logger
}

func NewRepo(db *bolt.DB, logger log.Logger) ds.Repositiory {
	dbs = dbRepo{
		db:     db,
		logger: log.With(logger, "pkg", "deviceStorage"),
	}

	return &dbs
}

/**
 * Store()
 *
 * Save the DeviceMessage to the Store using Devicename/Bucket -> (Topic:value)
 * Save if value > nil,
 * Delete topic if present
 * Deletebucket if no topic exists
 */
func (dbR *dbRepo) Store(d *ds.DeviceMessage) error {
	var err error
	//     0         1           2        3              4
	// sknSensors/device/node/version: 3.0.0
	//  * homie / device ID / $device-attribute
	//              NetworkBucket(Bucket(topic:value))
	//  * homie / device ID / node ID / $node-attribute
	//				NetworkBucket(Bucket(Bucket(topic:value)))
	//  * homie / device ID / node ID / property ID / $property-attribute
	//				NetworkBucket(Bucket(Bucket(Bucket(topic:value))))

	//     0         1           2        3              4                         5
	// sknSensors/device/node/version: 3.0.0
	//  * homie / device ID / $device-attribute / device-attribute-Property / device-attribute-Property-Property
	//				Bucket(Bucket(Bucket(property:value)))
	//  * homie / device ID / $device-attribute / device-attribute-Property
	//				Bucket(Bucket(property:value))
	//  * homie / device ID / $device-attribute
	//              Bucket(attribute:value)
	//  * homie / device ID / node ID / property ID / $property-attribute
	//				Bucket(Bucket(Bucket(attribute:value)))
	//  * homie / device ID / node ID / $node-attribute
	//				Bucket(Bucket(attribute:value))
	//  * homie / device ID / node ID / property ID
	//				Bucket(Bucket(Bucket(property:value)))

	if nil == d.Value || len(d.Value) == 0 { // Delete Actions
		err = dbR.db.Update(func(tx *bolt.Tx) error {
			var err error

			b := tx.Bucket(d.NetworkID)
			if b == nil {
				return fmt.Errorf("[WARN] Network Not Found!: %s", d.NetworkID)
			}

			b = b.Bucket(d.DeviceID)
			if b == nil {
				return fmt.Errorf("[WARN] Device Not Found!: %s", d.DeviceID)
			}

			if len(d.NodeID) == 0 && len(d.PropertyID) == 0 { // device only action
				var keyCount, buckets int
				err = b.ForEach(func(k, v []byte) error {
					if v != nil && len(v) > 0 {
						keyCount++
					}
					if v == nil {
						buckets++
					}
					return nil
				})
				if keyCount > 0 {
					return b.Delete(d.AttributeID)
				} else if buckets == 0 { // no more attrs or child buckets
					return tx.DeleteBucket(d.DeviceID)
				}
			}

			if len(d.NodeID) > 0 && len(d.PropertyID) > 0 && len(d.AttributeID) > 0 {
				b = b.Bucket(d.NodeID)
				if b == nil {
					return fmt.Errorf("[WARN] Node Not Found!: %s", d.NodeID)
				}

				b = b.Bucket(d.PropertyID)
				if b == nil {
					return fmt.Errorf("[WARN] Property Not Found!: %s", d.PropertyID)
				}

				return b.Delete(d.AttributeID)
			}

			if len(d.NodeID) > 0 && len(d.PropertyID) == 0 && len(d.AttributeID) > 0 {
				b = b.Bucket(d.NodeID)
				if b == nil {
					return fmt.Errorf("[WARN] Node Not Found!: %s", d.NodeID)
				}

				return b.Delete(d.AttributeID)
			}

			if len(d.NodeID) > 0 && len(d.PropertyID) > 0 && len(d.AttributeID) == 0 {
				b = b.Bucket(d.NodeID)
				if b == nil {
					return fmt.Errorf("[WARN] Node Not Found!: %s", d.NodeID)
				}

				return b.DeleteBucket(d.PropertyID)
			}

			return err
		})

	} else {
		//  * homie / device ID / $device-attribute
		//              NetworkBucket(Bucket(attribute:value))
		//  * homie / device ID / node ID / $node-attribute
		//				NetworkBucket(Bucket(Bucket(attribute:value)))
		//  * homie / device ID / node ID / property ID / $property-attribute
		//				NetworkBucket(Bucket(Bucket(Bucket(attribute:value))))
		//  * homie / device ID / node ID / property ID
		//				NetworkBucket(Bucket(Bucket(Bucket(property:value))))

		err = dbR.db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists(d.NetworkID)
			if err != nil {
				return fmt.Errorf("[WARN] Device cannot be created or found!: %s", d.DeviceID)
			}

			b, err = b.CreateBucketIfNotExists(d.DeviceID)
			if err != nil {
				return fmt.Errorf("[WARN] Device cannot be created or found!: %s", d.DeviceID)
			}

			if len(d.NodeID) == 0 && len(d.PropertyID) == 0 { // device attrs
				err = b.Put(d.AttributeID, d.Value)
				return err
			}

			if len(d.NodeID) > 0 {
				b, err = b.CreateBucketIfNotExists(d.NodeID)
				if err != nil {
					return fmt.Errorf("[WARN] Node cannot be created or found!: %s", d.NodeID)
				}
			}

			if len(d.PropertyID) == 0 && len(d.AttributeID) > 0 { // Node Attr
				err = b.Put(d.AttributeID, d.Value)
				return err
			}

			if len(d.PropertyID) > 0 && len(d.AttributeID) == 0 { // Node Property
				b, err = b.CreateBucketIfNotExists(d.PropertyID)
				if err != nil {
					return fmt.Errorf("[WARN] property cannot be created or found!: %s", d.PropertyID)
				}

				return b.Put(d.PropertyID, d.Value)
			}

			if len(d.PropertyID) > 0 && len(d.AttributeID) > 0 { // node property attr
				b, err = b.CreateBucketIfNotExists(d.PropertyID)
				if err != nil {
					return fmt.Errorf("[WARN] property cannot be created or found!: %s", d.PropertyID)
				}

				return b.Put(d.AttributeID, d.Value)
			}

			return err
		})
	}

	if err != nil {
		log.Printf("[ERROR] Store() --> Update Failed %v, %s\n", err.Error(), d.String())
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
		log.Printf("[WARN] Network rendering Failed: %v", err.Error())
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
		log.Printf("[WARN] Network::Device rendering Failed: %v", err.Error())
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
 * ListHomieDBCollection
 * List the Devices Found/Recorded
 */
func listHomieDBCollection(db *bolt.DB) {
	time.Sleep(1 * time.Second) // slow the pace
	if networks := networkList(db); len(networks) > 0 {
		for nets, network := range networks {
			log.Printf("[%d] Network: %s\n", nets, network)
			if devices := deviceList(db, network); len(devices) > 0 {
				for idx, device := range devices {
					log.Printf("\t[%d] Device: %s\n", idx, device)
					if detail, orderedKeys := deviceDetails(db, network, device); len(detail) > 0 {
						for _, k := range orderedKeys {
							v := detail[k]
							log.Printf("\t\t %s --> %s\n", k, v)
						}
					}
				}
			}
		}
	}
}

/*
 * DBStatsAsJSON
 * Outputs boltDB Database Stats
 */
func DBStatsAsJSON(db *bolt.DB) string {
	dbs := db.Stats()
	var res string

	json, _ := json.MarshalIndent(dbs, "", "    ")
	res = string(json)

	return res
}
