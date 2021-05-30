package deviceScheduler

/**
  deviceScheduler/service.go:

  Consume OTA Incoming messages and manage the schedule of firmware downloads
*/

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	cc "github.com/skoona/homie-service/internal/utils"
)

var (
	cfg       cc.Config
	logger    log.Logger
	otaStream OTAInteractor
	sch       *schedulerService
)

/*
 * Start()
 *
 * Initialize this service
 */
func Start(dfg cc.Config, s OTAInteractor) (SchedulerService, error) {
	var err error
	cfg = dfg
	otaStream = s

	NewSchedulerService(dfg, s)
	logger = sch.logger
	level.Debug(logger).Log("event", "Calling Start()")

	// Initialize a Message Channel
	consumeFromOTAStreamProvider(sch.otaStream.EnableTriggers(), logger)

	level.Debug(logger).Log("event", "Start() completed")

	return sch, err
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")

	level.Debug(logger).Log("event", "Stop() completed")
}
