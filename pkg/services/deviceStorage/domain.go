package deviceStorage

/*
	deviceStorage/domain.go

	Implement dataDB/LevelDB Connection to DeviceSource Service
*/

import (
	"fmt"

	bolt "github.com/boltdb/bolt"
	"github.com/go-kit/kit/log/level"
	dss "github.com/skoona/homie-service/pkg/services/deviceSource"
)

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
func (dbR *dbRepo) Store(d *dss.DeviceMessage) error {
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

	level.Debug(dbR.logger).Log("msg", "Calling Store()", "dm", d.String())

	if nil == d.Value || len(d.Value) == 0 { // Delete actions
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

			if len(d.NodeID) == 0 && len(d.AttributeID) > 0 { // device attribute
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

			if len(d.NodeID) > 0 && len(d.AttributeID) > 0 {
				b = b.Bucket(d.NodeID)
				if b == nil {
					return fmt.Errorf("[WARN] Node Not Found!: %s", d.NodeID)
				}

				return b.Delete(d.AttributeID)
			}

			if len(d.NodeID) > 0 && len(d.PropertyID) > 0 {
				b = b.Bucket(d.NodeID)
				if b == nil {
					return fmt.Errorf("[WARN] Node Not Found!: %s", d.NodeID)
				}

				return b.DeleteBucket(d.PropertyID)
			}

			return err
		})

	} else { // Create Actions

		err = dbR.db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists(d.NetworkID)
			if err != nil {
				return fmt.Errorf("[WARN] Device cannot be created or found!: %s", d.DeviceID)
			}

			b, err = b.CreateBucketIfNotExists(d.DeviceID)
			if err != nil {
				return fmt.Errorf("[WARN] Device cannot be created or found!: %s", d.DeviceID)
			}

			if len(d.NodeID) == 0 && len(d.PPropertyID) > 0 { // device attrs-property-property
				b, err = b.CreateBucketIfNotExists(d.AttributeID)
				if err != nil {
					return fmt.Errorf("[WARN] Attribute cannot be created or found!: %s", d.AttributeID)
				}

				b, err = b.CreateBucketIfNotExists(d.PropertyID)
				if err != nil {
					return fmt.Errorf("[WARN] Property cannot be created or found!: %s", d.PropertyID)
				}

				b, err = b.CreateBucketIfNotExists(d.PPropertyID)
				if err != nil {
					return fmt.Errorf("[WARN] PropertyProperty cannot be created or found!: %s", d.PPropertyID)
				}

				err = b.Put(d.PPropertyID, d.Value)
				return err
			}

			if len(d.NodeID) == 0 && len(d.PropertyID) > 0 { // device attr-property
				b, err = b.CreateBucketIfNotExists(d.AttributeID)
				if err != nil {
					return fmt.Errorf("[WARN] Attribute cannot be created or found!: %s", d.AttributeID)
				}

				b, err = b.CreateBucketIfNotExists(d.PropertyID)
				if err != nil {
					return fmt.Errorf("[WARN] Property cannot be created or found!: %s", d.PropertyID)
				}

				err = b.Put(d.PropertyID, d.Value)
				return err
			}

			if len(d.NodeID) == 0 && len(d.AttributeID) > 0 { // device attrs
				err = b.Put(d.AttributeID, d.Value)
				return err
			}

			if len(d.NodeID) > 0 {
				b, err = b.CreateBucketIfNotExists(d.NodeID)
				if err != nil {
					return fmt.Errorf("[WARN] Node cannot be created or found!: %s", d.NodeID)
				}
			}

			if len(d.PropertyID) > 0 && len(d.AttributeID) > 0 { // node property attr
				b, err = b.CreateBucketIfNotExists(d.PropertyID)
				if err != nil {
					return fmt.Errorf("[WARN] Property cannot be created or found!: %s", d.PropertyID)
				}

				return b.Put(d.AttributeID, d.Value)
			}

			if len(d.PropertyID) > 0 && len(d.AttributeID) == 0 { // Node Property
				b, err = b.CreateBucketIfNotExists(d.PropertyID)
				if err != nil {
					return fmt.Errorf("[WARN] property cannot be created or found!: %s", d.PropertyID)
				}

				return b.Put(d.PropertyID, d.Value)
			}

			if len(d.PropertyID) == 0 && len(d.AttributeID) > 0 { // Node Attr
				err = b.Put(d.AttributeID, d.Value)
				return err
			}

			return err
		})
	}

	if err != nil {
		level.Error(dbR.logger).Log("Alert", "Store() Update Failed", "error", err.Error(), "dm", d.String())
	}

	return err
}
