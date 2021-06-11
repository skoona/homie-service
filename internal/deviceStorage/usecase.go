package deviceStorage

/*
  deviceStorage/usecase.go:

  DeviceCore Repository implementation primarily for DeviceSource
  Utilities to List current Database
*/

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	bolt "go.etcd.io/bbolt"
	"strings"
)

/**
 * LoadSiteNetwork()
 * Reconstitute the Site Network Object
 */
//goland:noinspection VacuumLines
func (dbR *dbRepo) LoadNetwork(networkName string) dc.Network {
	nw  := dc.NewNetwork("Restored", networkName)

	devices := deviceList(dbR.db, networkName)
	for _, deviceName := range devices {
		if strings.HasPrefix(deviceName, "$broad") {
			continue
		}

		device, err := buildNetworkDevice(dbR.db, networkName, deviceName)
		if nil != err {
			level.Error(dbR.logger).Log("method", "LoadNetwork()", "error", err.Error())
			continue
		}

		nw.Devices[deviceName] = device
	}

	return nw
}

/**
 * StoreSchedule(schedule Schedule)
 */
func (dbR *dbRepo) StoreSchedule(d dc.Schedule) error {
	level.Debug(dbR.logger).Log("event", "Calling StoreSchedule()", "dm", d.String())
	var err error

	/*
	* Create Schedules */
	err = dbR.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Schedules"))
		if err != nil {
			return fmt.Errorf("[WARN] Schedules cannot be created or found!: %s, error: %s", "Schedules", err.Error())
		}

		b, err = b.CreateBucketIfNotExists([]byte(d.ID))
		if err != nil {
			return fmt.Errorf("[WARN] Schedule ID cannot be created or found!: %s, error: %s", d.ID, err.Error())
		}

		// save as a json blob
		value, _ := json.Marshal(d)
		err = b.Put([]byte(d.ID), value)
		return err
	})

	if err != nil {
		level.Error(dbR.logger).Log("Alert", "StoreSchedule() Failed", "error", err.Error(), "schedule", d.String())
	} else {
		level.Debug(dbR.logger).Log("event", "StoreSchedule() complete")
	}

	return err
}

/**
 * RemoveSchedule(schedule Schedule)
 */
func (dbR *dbRepo) RemoveSchedule(d dc.Schedule) error {
	level.Debug(dbR.logger).Log("event", "Calling RemoveSchedule()", "schedule", d.String())

	/*
	* Delete Schedules */
	err := dbR.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Schedules"))
		if err != nil {
			return fmt.Errorf("[WARN] Schedules cannot be created or found!: %s, error: %s", "Schedules", err.Error())
		}

		err = b.DeleteBucket([]byte(d.ID))
		return err
	})

	if err != nil {
		level.Error(dbR.logger).Log("Alert", "RemoveSchedule() Failed", "error", err.Error(), "schedule", d.String())
	} else {
		level.Debug(dbR.logger).Log("event", "RemoveSchedule() complete")
	}

	return err
}

/**
 * LoadSchedules()
 */
func (dbR *dbRepo) LoadSchedules() map[dc.EID]dc.Schedule {
	level.Debug(dbR.logger).Log("event", "Calling LoadSchedules()")
	schedMap := map[dc.EID]dc.Schedule{}

	/*
	 * Load Schedules */
	err := dbR.db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte("Schedules"))
		if b == nil {
			return fmt.Errorf("No schedules %s", "available")
		}
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var schedule dc.Schedule
			err = json.Unmarshal(v, &schedule)
			schedMap[dc.EID(string(k))] = schedule
			fmt.Printf("schedule key=%s, value=%s\n", k, v)
		}
		return err
	})

	if err != nil {
		level.Warn(dbR.logger).Log("Alert", "LoadSchedules() Failed", "error", err.Error())
	} else {
		level.Debug(dbR.logger).Log("event", "LoadSchedules() complete", "retrieved", len(schedMap))
	}

	return schedMap
}

/**
 * Remove()
 *
 * Repository Implementation
 */
func (dbR *dbRepo) Remove(d dc.DeviceMessage) error {

	level.Debug(dbR.logger).Log("event", "Calling Remove()", "dm", d.String())

	err := dbR.db.Update(func(tx *bolt.Tx) error {
		var err error

		b := tx.Bucket(d.NetworkID)
		if b == nil {
			return fmt.Errorf("[WARN] Network Not Found!: %s, error: %s", d.NetworkID, err.Error())
		}

		return b.DeleteBucket(d.DeviceID)
	})

	if err != nil {
		level.Error(dbR.logger).Log("Alert", "Remove() Update Failed", "error", err.Error(), "dm", d.String())
	} else {
		level.Debug(dbR.logger).Log("event", "Remove() complete", "device", d.DeviceID)
	}

	return err
}

/**
 * Store()
 *
 * Repository Implementation
 */
func (dbR *dbRepo) Store(d dc.DeviceMessage) error {
	//     0         1           2        3              4                         5
	// sknSensors/device/node/version: 3.0.0
	//  * homie / device ID / $device-attribute / device-attribute-Property / device-attribute-Property-Property
	//				Bucket(Bucket(Bucket(bucket(bucket(property:value)))))
	//  * homie / device ID / $device-attribute / device-attribute-Property
	//				Bucket(Bucket(bucket(bucket(property:value))))
	//  * homie / device ID / $device-attribute
	//              Bucket(bucket(bucket(attribute:value)))
	//  * homie / device ID / node ID / property ID / $property-attribute
	//				Bucket(Bucket(Bucket(bucket(bucket(attribute:value)))))
	//  * homie / device ID / node ID / $node-attribute
	//				Bucket(Bucket(bucket(bucket(attribute:value))))
	//  * homie / device ID / node ID / property ID
	//				Bucket(Bucket(Bucket(bucket(property:value))))

	level.Debug(dbR.logger).Log("event", "Calling Store()", "dm", d.String())

	err := dbR.db.Update(func(tx *bolt.Tx) error {
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
				return fmt.Errorf("[WARN] PProperty(x/d/a/p/p) cannot be created or found!: %s, error: %s", d.PPropertyID, err.Error())
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
			b, err = b.CreateBucketIfNotExists(d.AttributeID)
			if err != nil {
				return fmt.Errorf("[WARN] Attribute(x/d/a) cannot be created or found!: %s, error: %s", d.AttributeID, err.Error())
			}
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

			b, err = b.CreateBucketIfNotExists(d.AttributeID)
			if err != nil {
				return fmt.Errorf("[WARN] Attribute(x/d/n/p/a) cannot be created or found!: %s, error: %s", d.AttributeID, err.Error())
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
			b, err = b.CreateBucketIfNotExists(d.AttributeID)
			if err != nil {
				return fmt.Errorf("[WARN] Attribute(x/d/n/a) cannot be created or found!: %s, error: %s", d.AttributeID, err.Error())
			}

			return b.Put(d.AttributeID, d.Value)
		}

		return err
	})

	if err != nil {
		level.Error(dbR.logger).Log("Alert", "Store() Failed", "error", err.Error(), "dm", d.String())
	} else {
		level.Debug(dbR.logger).Log("event", "Store() complete", "device", d.DeviceID)
	}

	return err
}

/*
 * Returns string array a device's properties (topic:value) */
func buildNetworkDevice(db *bolt.DB, networkName, deviceName string) (dc.Device, error) {
	var device dc.Device
	var deviceAttribute dc.DeviceAttribute
	var deviceAttributeProperty dc.DeviceAttributeProperty
	var deviceAttributePropertyProperty dc.DeviceAttributePropertyProperty
	var node dc.DeviceNode
	//var nodeProperty dc.DeviceNodeProperty
	//var nodeAttribute dc.DeviceNodeAttribute
	//var nodePropertyAttribute dc.DeviceNodePropertyAttribute

	// X/D/								bucket(bucket())
	// X/D/N/							bucket(bucket(bucket()))
	//									  1/     2/     3/

	// X/D/A/(k,v) 						bucket(bucket(bucket(k,v)))
	//									  1/     2/     3/

	// X/D/A/P/(k,v) 					bucket(bucket(bucket(bucket(k,v))))
	// X/D/N/P/(k,v) 					bucket(bucket(bucket(bucket(k,v))))
	// X/D/N/A/(k,v) 					bucket(bucket(bucket(bucket(k,v))))
	//									  1/     2/     3/     4/

	// X/D/A/P/P/(k,v) 					bucket(bucket(bucket(bucket(bucket(k,v)))))
	// X/D/N/P/A/(k,v) 					bucket(bucket(bucket(bucket(bucket(k,v)))))
	//									  1/     2/     3/     4/     5/

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(networkName)) // (1) Network Level
		if b == nil {
			return fmt.Errorf("Network Not Found!: %s", networkName)
		}
		b = b.Bucket([]byte(deviceName)) // (2) Device Level
		if b == nil {
			return fmt.Errorf("Network Not Found!: %s", deviceName)
		}
		device = dc.NewDevice(networkName, deviceName)

		/*
		 * examine device bucket for nodes or attributes */
		err := b.ForEach(func(k, v []byte) error {
			fmt.Printf("(2) X/D...... Each Bucket=%s Key=%s Value=[%s]%T \n", deviceName, string(k), string(v), string(v))
			if !strings.HasPrefix(string(k), "$") && (len(device.Nodes) == 0) {
				node = dc.NewDeviceNode(deviceName, string(k))
				device.Nodes[string(k)] = node
			}

			if string(v) != "" {
				return nil
			}
			c := b.Bucket(k) // (3) Node/Attr Level
			if c == nil {
				return nil
			}
			/*
			 * bucket is either a Node or a DeviceAttribute container
			 * if attr follow device attr path
			 * else follow nodepath
			*/
			if strings.HasPrefix(string(k), "$") {  // then device attr path
				// 3 attribute
				// 4 property
				// 5 propertyProperty
				err := c.ForEach(func(kk, vv []byte) error {
					fmt.Printf("(3) X/D/A.... Each Bucket=%s Key=%s Value=[%s] \n", deviceName, kk, vv)
					deviceAttribute = dc.NewDeviceAttribute(deviceName, string(kk), string(vv))
					device.Attrs[string(kk)] = deviceAttribute

					if string(vv) != "" {
						return nil
					}
					d := c.Bucket(kk) // (4) Property
					if d == nil {
						return nil
					}
					err := d.ForEach(func(kkk, vvv []byte) error {
						fmt.Printf("(4) X/D/A/P.. Each Bucket=%s Key=%s Value=[%s] \n", k, kkk, vvv)
						deviceAttributeProperty = dc.NewDeviceAttributeProperty(string(k), string(kkk), string(vvv))
						device.Attrs[string(kk)].Props[string(kkk)] = deviceAttributeProperty

						if string(vvv) != "" {
							return nil
						}
						e := d.Bucket(kkk) // (5) PProperty/Attr
						if e == nil {
							return nil
						}
						err := e.ForEach(func(kkkk, vvvv []byte) error {
							fmt.Printf("(5) X/D/A/P/P Each Bucket=%s Key=%s Value=[%s] \n", kk, kkkk, vvvv)
							deviceAttributePropertyProperty = dc.NewDeviceAttributePropertyProperty(string(kkk), string(kkkk), string(vvvv))
							device.Attrs[string(kk)].Props[string(kkk)].Props[string(kkkk)] = deviceAttributePropertyProperty
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
					fmt.Printf("(3) X/D/N.... Each Bucket=%s Key=%s Value=[%s] \n", deviceName, kk, vv)

					if string(vv) != "" {
						return nil
					}
					d := c.Bucket(kk) // (4) Attribute | Property
					if d == nil {
						return nil
					}
					err := d.ForEach(func(kkk, vvv []byte) error {
						if strings.HasPrefix(string(kkk), "$") {
							fmt.Printf("(4) X/D/N/A.. Each Bucket=%s Key=%s Value=[%s] \n", k, kkk, vvv)
						} else {
							fmt.Printf("(4) X/D/N/P.. Each Bucket=%s Key=%s Value=[%s] \n", k, kkk, vvv)
						}

						if string(vvv) != "" {
							return nil
						}
						e := d.Bucket(kkk) // (5) PPropertyAttribute
						if e == nil {
							return nil
						}
						err := e.ForEach(func(kkkk, vvvv []byte) error {
							fmt.Printf("(5) X/D/N/P/A Each Bucket=%s Key=%s Value=[%s] \n", kk, kkkk, vvvv)
							return nil
						})
						return err
					})
					return err
				})
				return err
			} // end if path
			return nil
		})
		return err
	})

	return device, err
}
