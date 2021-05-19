package deviceStorage

/*
	deviceStorage/repository.go

	Implement dataDB/LevelDB Connection to DeviceSource
*/

import (
	"context"
	"fmt"

	bolt "github.com/boltdb/bolt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dss "github.com/skoona/homie-service/pkg/services/deviceSource"
)

var dbs dbRepo

type dbRepo struct {
	db     *bolt.DB
	ctx    context.Context
	logger log.Logger
}

func NewRepo(ctx context.Context, db *bolt.DB, log log.Logger) dss.Repositiory {
	dbs = dbRepo{
		db:     db,
		ctx:    ctx,
		logger: log,
	}

	return &dbs
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
func (dbR *dbRepo) Store(d *dss.DeviceMessage) error {
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

	level.Debug(dbR.logger).Log("msg", "Calling Store()", "dm", d.String())

	if nil == d.Value || len(d.Value) == 0 { // Create/Update Actions
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

	} else { // Delete Actions
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
		level.Error(dbR.logger).Log("Alert", "Store() Update Failed", "errorMsg", err.Error(), "dm", d.String())
	}

	return err
}
