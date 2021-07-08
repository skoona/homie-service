package deviceStorage

/*
  deviceStorage/usecase.go:

  DeviceCore Repository implementation primarily for DeviceSource
  Utilities to List current Database
*/

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	bolt "go.etcd.io/bbolt"
	"strings"
)

// RestoreNetworkFromDB Reconstitute the Site NETWORK Object
func (r *dbRepo) RestoreNetworkFromDB(networkName string) dc.Network {
	nw  := dc.NewNetwork("Restored", networkName)

	devices := deviceList(r.db, networkName)
	for _, deviceName := range devices {
		if strings.HasPrefix(deviceName, "$broad") {
			continue
		}
		device, err := buildNetworkDevice(r.db, networkName, deviceName)
		if nil != err {
			level.Error(r.logger).Log("method", "()", "error", err.Error())
			continue
		}
		nw.Devices[deviceName] = device
	}

	return nw
}

// StoreSchedule (schedule Schedule)
func (r *dbRepo) StoreSchedule(d dc.Schedule) error {
	level.Debug(r.logger).Log("event", "Calling StoreSchedule()", "dm", d.String())
	if d.ElementType != dc.CoreTypeSchedule {
		return fmt.Errorf("warn invalid schedule type %d", d.ElementType)
	}

	/*
	* Create Schedules */
	err := r.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Schedules"))
		if err != nil {
			return fmt.Errorf("[WARN] Schedules cannot be created or found!: %s, error: %s", "Schedules", err.Error())
		}

		b, err = b.CreateBucketIfNotExists([]byte(d.ID))
		if err != nil {
			return fmt.Errorf("[WARN] Schedule ID cannot be created or found!: %s, error: %s", d.ID, err.Error())
		}

		// save as a json blob
		value, err := json.Marshal(d)
		if err != nil {
			level.Error(r.logger).Log("Alert", "StoreSchedule(a) Failed", "error", err.Error(), "schedule", d.String())
			return err
		}
		return b.Put([]byte(d.ID), value)
	})

	if err != nil {
		level.Error(r.logger).Log("Alert", "StoreSchedule(b) Failed", "error", err.Error(), "schedule", d.String())
	}

	return err
}

// RemoveSchedule(schedule Schedule)
//
func (r *dbRepo) RemoveSchedule(scheduleID string) error {
	level.Debug(r.logger).Log("event", "Calling RemoveSchedule()", "schedule", scheduleID)

	/*
	* Delete Schedules */
	err := r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Schedules"))
		if b == nil {
			return fmt.Errorf("[WARN] Schedules cannot be created or found!: %s", "Schedules")
		}

		err := b.DeleteBucket([]byte(scheduleID))
		return err
	})
	if err != nil {
		level.Error(r.logger).Log("Alert", "RemoveSchedule() Failed", "error", err.Error(), "schedule", scheduleID)
		return err
	}
	level.Debug(r.logger).Log("event", "RemoveSchedule() complete")

	return err
}


// LoadSchedules()
func (r *dbRepo) LoadSchedules() map[string]dc.Schedule {
	_ = level.Debug(r.logger).Log("event", "Calling LoadSchedules()")
	schedMap := map[string]dc.Schedule{}

	/*
	 * Load Schedules */
	err := r.db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte("Schedules"))
		if b == nil {
			return errors.New("no schedules available")
		}

		err = b.ForEach( func(k, v []byte) error {
			if string(v) == "" {
				c := b.Bucket(k)
				if c == nil {
					return nil
				}
				d := c.Cursor()
				kk, vv := d.First()
				schedule := dc.Schedule{}

				err := json.Unmarshal(vv, &schedule)
				if err == nil {
					schedMap[schedule.ID] = schedule
					_ = level.Info(r.logger).Log("Schedule", schedule.String())
				} else {
					_ = level.Error(r.logger).Log("error", err.Error())
				}
				_ = level.Info(r.logger).Log( "method","LoadSchedules()", "schedule key", kk,"value", vv)
			}
			return err
		})
		return err
	})

	if err != nil {
		_ = level.Warn(r.logger).Log("Alert", "LoadSchedules() Failed", "error", err.Error())
	} else {
		_ = level.Debug(r.logger).Log("event", "LoadSchedules() complete", "retrieved", len(schedMap))
	}

	return schedMap
}


// LoadBroadcasts()
func (r *dbRepo) LoadBroadcasts(networkName string) []dc.Broadcast {
	var bc dc.Broadcast
	broads := []dc.Broadcast{}

	r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(networkName)) // (1) Network Level
		if b == nil {
			return fmt.Errorf("network not found!: %s", networkName)
		}
		b = b.Bucket([]byte("$broadcast")) // (2) Broadcast
		if b == nil {
			return fmt.Errorf("$broadcast Not Found!: %s", "top level")
		}
		// x/b/t/i v
		b.ForEach(func(topic, v []byte) error {
			c := b.Bucket(topic) // (3) Topic
			if c == nil {
				return nil
			}
			c.ForEach(func(item, vv []byte) error {
				d := c.Bucket(item) // (4) item
				if d == nil {
					return nil
				}
				d.ForEach(func(kkk, vvv []byte) error {
					bc = dc.NewBroadcast(networkName, string(topic), string(kkk), string(vvv))
					broads = append(broads, bc)
					return nil
				})
				return nil
			})
			return nil
		})
		return nil
	})
	return broads
}

// RemoveAllBroadcasts()
func (r *dbRepo) RemoveAllBroadcasts(networkName string) error {
	_ = r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(networkName))
		if b == nil {
			return fmt.Errorf("[WARN] Network Not Found!: %s", networkName)
		}
		return b.DeleteBucket([]byte("$broadcast"))
	})
	return nil
}

// Remove()
// Repository Implementation
func (r *dbRepo) Remove(d dc.DeviceMessage) error {
	level.Debug(r.logger).Log("event", "Calling Remove()", "dm", d.String())
	err := r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(d.NetworkID)
		if b == nil {
			return fmt.Errorf("[WARN] Network Not Found!: %s", d.NetworkID)
		}
		return b.DeleteBucket(d.DeviceID)
	})
	if err != nil {
		level.Error(r.logger).Log("Alert", "Remove() Update Failed", "error", err.Error(), "dm", d.String())
		return err
	}
	level.Debug(r.logger).Log("event", "Remove() complete", "device", d.DeviceID)

	return err
}

/**
 * Store()
 *
 * Repository Implementation
 */
func (r *dbRepo) Store(d dc.DeviceMessage) error {
	//      1         2           3                  4                           5
	// sknSensors/device/node/version: 3.0.0
	//  * homie / device ID / $device-attribute / device-attribute-Property / device-attribute-Property-Property
	//	  Bucket( Bucket(     Bucket(             Bucket(                     Bucket(k:v)))))
	//  * homie / device ID / $device-attribute / device-attribute-Property
	//	  Bucket( Bucket(     Bucket(             Bucket(k:v))))
	//  * homie / device ID / $device-attribute
	//    Bucket( Bucket(     Bucket(k:v)))
	//  * homie / device ID / node ID / property ID / $property-attribute
	//	  Bucket( Bucket(     Bucket(   Bucket(       Bucket(k:v)))))
	//  * homie / device ID / node ID / $node-attribute
	//	  Bucket( Bucket(     Bucket(   Bucket(k:v))))
	//  * homie / device ID / node ID / property ID
	//	  Bucket( Bucket(     Bucket(   Bucket(k:v))))

	level.Debug(r.logger).Log("event", "Calling Store()", "dm", d.String())

	err := r.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(d.NetworkID)
		if err != nil {
			return fmt.Errorf("[WARN] Network cannot be created or found!: %s, error: %s", d.NetworkID, err.Error())
		}

		/*
		 * handle special case of $broadcast
		*/
		if d.HomieType == dc.CoreTypeBroadcast {
			b, err = b.CreateBucketIfNotExists(d.DeviceID)
			if err != nil {
				return fmt.Errorf("[WARN] (x/b....) cannot be created or found!: %s, error: %s", d.DeviceID, err.Error())
			}

			defaultTopic := d.PropertyID
			if string(defaultTopic) == "" {
				defaultTopic  = []byte("item")
			}
			b, err = b.CreateBucketIfNotExists(defaultTopic)
			if err != nil {
				return fmt.Errorf("[WARN] (x/b/t..) cannot be created or found!: %s, error: %s", defaultTopic, err.Error())
			}

			defaultProperty := d.PPropertyID
			if string(defaultProperty) == "" {
				defaultProperty = []byte("item")
			}
			b, err = b.CreateBucketIfNotExists(defaultProperty)
			if err != nil {
				return fmt.Errorf("[WARN] (x/b/t/i) cannot be created or found!: %s, error: %s", defaultProperty, err.Error())
			}

			err = b.Put(defaultProperty, d.Value)
			return err
		}

		b, err = b.CreateBucketIfNotExists(d.DeviceID)
		if err != nil {
			return fmt.Errorf("[WARN] Device cannot be created or found!: %s, error: %s", d.DeviceID, err.Error())
		}

		if d.HomieType == dc.CoreTypeDeviceAttributePropertyProperty { // device attrs-property-property
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

		if d.HomieType == dc.CoreTypeDeviceAttributeProperty { // device attr-property
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

		if d.HomieType == dc.CoreTypeDeviceAttribute { // device attrs
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
			return fmt.Errorf("ALERT Unknown(node) Message!: %s", d.String())
		}

		// x/d/n/p/a
		if d.HomieType == dc.CoreTypeDeviceNodePropertyAttribute { // node property attr
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
		if d.HomieType == dc.CoreTypeDeviceNodeProperty { // Node Property
			b, err = b.CreateBucketIfNotExists(d.PropertyID)
			if err != nil {
				return fmt.Errorf("[WARN] Property(x/d/n/p) cannot be created or found!: %s, error: %s", d.PropertyID, err.Error())
			}

			return b.Put(d.PropertyID, d.Value)
		}

		// x/d/n/a
		if d.HomieType == dc.CoreTypeDeviceNodeAttribute { // Node Attr
			b, err = b.CreateBucketIfNotExists(d.AttributeID)
			if err != nil {
				return fmt.Errorf("[WARN] Attribute(x/d/n/a) cannot be created or found!: %s, error: %s", d.AttributeID, err.Error())
			}

			return b.Put(d.AttributeID, d.Value)
		}

		return err
	})

	if err != nil {
		level.Error(r.logger).Log("Alert", "Store() Failed", "error", err.Error(), "dm", d.String())
	} else {
		level.Debug(r.logger).Log("event", "Store() complete", "device", d.DeviceID)
	}

	return err
}

/*
 * Returns string array a device's properties (topic:value) */
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
		fmt.Printf("(0) X/D...... %s:%s \n", networkName, deviceName)

		b = b.Bucket([]byte(deviceName)) // (2) Device Level
		if b == nil {
			return fmt.Errorf("network not found!: %s", deviceName)
		}
		device = dc.NewDevice(networkName, deviceName)

		/*
		 * examine device bucket for nodes or attributes */
		err := b.ForEach(func(k, v []byte) error {
			fmt.Printf("(2)[1k] X/D...... Bucket[2]=%s, enum Key=%s Value=[%s] \n", deviceName, string(k), string(v))

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
					deviceAttribute, found = device.Attrs[string(k)]
					if !found {
						fmt.Printf("(3)[2ka] X/D/A.... %s:%s:%s -->[%s] \n", deviceName, k, kk, vv)
						deviceAttribute = dc.NewDeviceAttribute(deviceName, string(k), string(v))
						device.Attrs[string(k)] = deviceAttribute
					}
					d := c.Bucket(kk) // (4) Property
					if d == nil {
						return nil
					}
					err := d.ForEach(func(kkk, vvv []byte) error {

						deviceAttributeProperty, found = deviceAttribute.Props[string(kkk)]
						if !found {
							fmt.Printf("(4)[3ka] X/D/A/P.. %s:%s:%s:%s -->[%s]\n", deviceName, k, kk, kkk, vvv)
							deviceAttributeProperty = dc.NewDeviceAttributeProperty(string(kk), string(kkk), string(vvv))
							deviceAttribute.Props[string(kkk)] = deviceAttributeProperty
						}

						e := d.Bucket(kkk) // (5) PProperty
						if e == nil {
							return nil
						}
						err := e.ForEach(func(kkkk, vvvv []byte) error {
							//deviceAttribute, found = device.Attrs[string(kk)]
							//if !found {
							//	deviceAttribute = dc.NewDeviceAttribute(deviceName, string(kk), string(vv))
							//	device.Attrs[string(kk)] = deviceAttribute
							//}
							//deviceAttributeProperty, found = deviceAttribute.Props[string(kkk)]
							//if !found {
							//	deviceAttributeProperty = dc.NewDeviceAttributeProperty(string(kk), string(kkk), string(vvv))
							//	deviceAttribute.Props[string(kkk)] = deviceAttributeProperty
							//}
							deviceAttributePropertyProperty, found = deviceAttributeProperty.Props[string(kkkk)]
							if !found {
								fmt.Printf("(5)[4ka] X/D/A/P/P %s:%s:%s:%s:%s -->[%s] \n", deviceName, k, kk, kkk, kkkk, vvvv)
								deviceAttributePropertyProperty = dc.NewDeviceAttributePropertyProperty(string(kkk), string(kkkk), string(vvvv))
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
					node, found = device.Nodes[string(k)]
					if !found {
						fmt.Printf("(3)[2kb] X/D/N.... %s:%s:%s -->[%s] \n", deviceName, k, kk, vv)
						node = dc.NewDeviceNode(deviceName, string(k))
						device.Nodes[string(k)] = node
					}

					d := c.Bucket(kk) // (4) Attribute | Property
					if d == nil {
						return nil
					}
					err := d.ForEach(func(kkk, vvv []byte) error {


						if strings.HasPrefix(string(kkk), "$") {
							nodeAttribute, found = node.Attrs[string(kkk)]
							if !found {
								fmt.Printf("(4) X/D/N/A.. %s/%s/%s -->[%s] \n", deviceName, k, kkk, vvv)
								nodeAttribute = dc.NewDeviceNodeAttribute(string(kk), string(kkk), string(vvv))
								node.Attrs[string(kkk)] = nodeAttribute
							}

						} else {
							nodeProperty, found = node.Props[string(kkk)]
							if !found {
								fmt.Printf("(4) X/D/N/P.. %s/%s/%s -->[%s] \n", deviceName, k, kkk, vvv)
								nodeProperty = dc.NewDeviceNodeProperty(string(kk), string(kkk), string(vvv))
								node.Props[string(kkk)] = nodeProperty
							}

						}

						e := d.Bucket(kkk) // (5) PPropertyAttribute
						if e == nil {
							return nil
						}
						err := e.ForEach(func(kkkk, vvvv []byte) error {
							nodeProperty, found = node.Props[string(kkk)]
							if !found {
								fmt.Printf("(4) X/D/N/P.. %s/%s/%s -->[%s] \n", deviceName, k, kkk, vvv)
								nodeProperty = dc.NewDeviceNodeProperty(string(kk), string(kkk), string(vvv))
								node.Props[string(kkk)] = nodeProperty
							}
							nodePropertyAttribute, found = nodeProperty.Attrs[string(kkkk)]
							if !found {
								fmt.Printf("(5) X/D/N/P/A %s/%s/%s/%s -->[%s] \n", deviceName, k, kk, kkkk, vvvv)
								nodePropertyAttribute = dc.NewDeviceNodePropertyAttribute(string(kkk), string(kkkk), string(vvvv))
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
