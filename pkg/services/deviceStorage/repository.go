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
	cc "github.com/skoona/homie-service/pkg/utils"
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

/*
 * Start
 * Initializes this service
 */
func Start(ctx context.Context) (dss.Repositiory, error) {
	var err error
	logger := log.With(ctx.Value(cc.AppConfig).(cc.Config).Logger, "pkg", "deviceStorage")
	level.Debug(logger).Log("msg", "Calling Start()")

	/*
	 * Open K/V Store
	 */
	pDB, err := bolt.Open(ctx.Value(cc.DbConfig).(cc.DBConfig).DataStorage, 0764, nil) // &bolt.Options{Timeout: 1 * time.Second, ReadOnly: false})
	if err != nil {
		level.Error(logger).Log("msg", "[ERROR] Main bBolt DB Open Failed")
		err = fmt.Errorf("Main bBolt DB Open Failed: %v", err.Error())
		return nil, err
	}

	repo := NewRepo(ctx, pDB, logger)

	// Initialize a Message Channel
	// dvcSyncChannel = make(chan DeviceMessage, 256)   // averages 120 on startup
	// coreLogicChannel = make(chan DeviceMessage, 256) // averages 120 on startup

	return repo, nil
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(dbs.logger).Log("msg", "Calling Stop()")

	// Close the Device Channel
	// close(dvcSyncChannel)
	// close(coreLogicChannel)

	// List the Devices Found/Recorded
	listHomieDBCollection(dbs.db)

	// Show bBolt DB Stats
	if stats := DBStatsAsJSON(dbs.db); len(stats) > 0 {
		level.Info(dbs.logger).Log("dbStats", stats)
	}

	dbs.db.Close()
}
