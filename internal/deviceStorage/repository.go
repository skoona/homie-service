package deviceStorage

/*
	deviceStorage/repository.go

	Implement dataDB/LevelDB Connection to DeviceSource
*/

import (
	"fmt"

	bolt "github.com/boltdb/bolt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dss "github.com/skoona/homie-service/internal/deviceSource"
	cc "github.com/skoona/homie-service/internal/utils"
)

var dbs dbRepo

type dbRepo struct {
	db     *bolt.DB
	cfg    cc.Config
	logger log.Logger
}

func NewRepo(cfg cc.Config, db *bolt.DB, log log.Logger) dss.Repository {
	dbs = dbRepo{
		db:     db,
		cfg:    cfg,
		logger: log,
	}

	return &dbs
}

/*
 * Start
 * Initializes this service
 */
func Start(cfg cc.Config) (dss.Repository, error) {
	var err error
	var dataFile string
	logger := log.With(cfg.Logger, "pkg", "deviceStorage")
	level.Debug(logger).Log("event", "Calling Start()")

	/*
	 * Open K/V Store
	 */
	dataFile = cfg.Dbc.DataStorage
	pDB, err := bolt.Open(dataFile, 0764, nil) // &bolt.Options{Timeout: 1 * time.Second, ReadOnly: false})
	if err != nil {
		level.Error(logger).Log("event", "Main bBolt DB Open Failed", "datafile", dataFile)
		err = fmt.Errorf("main bBolt DB Open Failed: %v", err.Error())
		panic(err.Error())
	}

	repo := NewRepo(cfg, pDB, logger)

	level.Debug(logger).Log("event", "Start() completed")
	return repo, nil
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(dbs.logger).Log("event", "calling Stop()")

	// close the database
	dbs.db.Close()
	level.Debug(dbs.logger).Log("event", "Stop() completed")
}
