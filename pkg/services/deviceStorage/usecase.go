package deviceStorage

/*
  deviceStorage/usecase.go:

  DeviceSource Service Implementation
  Utilities to List current Database
*/

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	bolt "github.com/boltdb/bolt"
)

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
 * ListHomieDBCollection
 * List the Devices Found/Recorded
 */
func listHomieDBCollection(db *bolt.DB) {
	time.Sleep(1 * time.Second) // slow the pace
	if networks := networkList(db); len(networks) > 0 {
		for nets, network := range networks {
			fmt.Printf("[%d] Network: %s\n", nets, network)
			if devices := deviceList(db, network); len(devices) > 0 {
				for idx, device := range devices {
					fmt.Printf("\t[%d] Device: %s\n", idx, device)
					if detail, orderedKeys := deviceDetails(db, network, device); len(detail) > 0 {
						for _, k := range orderedKeys {
							v := detail[k]
							fmt.Printf("\t\t %s --> %s\n", k, v)
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
