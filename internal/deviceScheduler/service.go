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
	sch       *schedulerProvider
)

/*
 * Start()
 *
 * Initialize this service
 */
func Start(dfg cc.Config, s OTAInteractor, r dc.Repository) dc.SchedulerProvider {
	cfg = dfg
	otaStream = s

	NewSchedulerProvider(dfg, s, r)
	logger = sch.logger
	level.Debug(logger).Log("event", "Calling Start()")

	level.Debug(logger).Log("event", "Start() completed")

	return sch
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")

	level.Debug(logger).Log("event", "Stop() completed")
}
