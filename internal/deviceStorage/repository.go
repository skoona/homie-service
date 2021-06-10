package deviceStorage

/*
	deviceStorage/repository.go

	Implement dataDB/LevelDB Connection to DeviceSource
	ref: https://github.com/etcd-io/bbolt
*/

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	cc "github.com/skoona/homie-service/internal/utils"
	bolt "go.etcd.io/bbolt" // bolt "github.com/boltdb/bolt"
	"time"
)

var dbs *dbRepo

type dbRepo struct {
	db     *bolt.DB
	cfg    cc.Config
	logger log.Logger
}

func NewRepo(cfg cc.Config, db *bolt.DB, log log.Logger) dc.Repository {
	dbs = &dbRepo{
		db:     db,
		cfg:    cfg,
		logger: log,
	}

	return dbs
}

/*
 * Start
 * Initializes this service
 */
func Start(cfg cc.Config) (dc.Repository, error) {
	var err error
	var dataFile string
	logger := log.With(cfg.Logger, "pkg", "deviceStorage")
	level.Debug(logger).Log("event", "Calling Start()")

	/*
	 * Open K/V Store
	 */
	dataFile = cc.LocateFile( cfg.Dbc.DataStorage )
	level.Debug(logger).Log("event", "get datafile", "dataFile", dataFile)
	if dataFile == "" {
		dataFile = cfg.Dbc.DataStorage // it will be created
	}

	pDB, err := bolt.Open(dataFile, 0764, &bolt.Options{Timeout: 5 * time.Second, ReadOnly: false})
	if err != nil {
		level.Error(logger).Log("event", "Main bBolt DB Open Failed", "datafile", dataFile)
		err = fmt.Errorf("main bBolt DB Open Failed: %v", err.Error())
		return nil, err
	}

	repo := NewRepo(cfg, pDB, logger)

	level.Debug(logger).Log("event", "Start() completed")
	return repo, err
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
