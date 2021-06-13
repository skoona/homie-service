package deviceStorage

/**
  deviceStorage/usecase-utils.go:

  General database utilities for examining the db
*/

import (
	"encoding/json"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"sort"
	"strings"
	"time"
)


/*
 * Returns string array of device names (buckets)
 * which will include $broadcast as a device name
*/
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
 * Returns string array of networks names (top level buckets) */
func networkList(db *bolt.DB) []string {
	networks := []string{}

	err := db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			if strings.Contains(string(name), "Schedules") {
				return nil
			}
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
 * Returns string array a device's properties (topic:value) */
func deviceDetails(db *bolt.DB, networkName, deviceName string) (map[string]string, []string) {
	details := map[string]string{}

	// X/D/N/P/A
	// X/D/A/P/P

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
				c := b.Bucket(k) // Node/Attr Level
				if c == nil {
					return nil
				}

				err := c.ForEach(func(kk, vv []byte) error {
					if nil == vv {
						d := c.Bucket(kk) // Property
						if d == nil {
							return nil
						}

						err := d.ForEach(func(kkk, vvv []byte) error {
							if nil == vvv {
								e := d.Bucket(kkk) // PProperty/Attr
								if e == nil {
									return nil
								}

								err := e.ForEach(func(kkkk, vvvv []byte) error {
									if nil == vvvv {
										f := e.Bucket(kkkk) // PProperty
										if f == nil {
											return nil
										}

										err := f.ForEach(func(kkkkk, vvvvv []byte) error {
											details["["+string(k)+"]["+string(kk)+"]["+string(kkk)+"]["+string(kkkk)+"]"+string(kkkkk)] = string(vvvvv)
											return nil
										})
										return err
									}
									details["["+string(k)+"]["+string(kk)+"]["+string(kkk)+"]"+string(kkkk)] = string(vvvv)
									return nil
								})
								return err
							}
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
