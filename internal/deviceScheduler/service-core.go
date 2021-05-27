package deviceScheduler

/**
  deviceScheduler/service-core.go:

  Interactions with Network Engine
*/

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	cc "github.com/skoona/homie-service/internal/utils"
)

type (
	SchedulerService interface {
		BuildFirmwareCatalog() []dc.Firmware
		ApplySchedule(scheduleEID, deviceEID dc.EID)
		StartSchedule(scheduleEID dc.EID)
		DeleteSchedule(scheduleEID dc.EID)
		UpdateScheduler(dm dc.DeviceMessage)

		ApplyFirmwareUpload(blob []byte) error
		DeleteFirmware(firmwareEID dc.EID)
		GetFirmware(firmwareEID dc.EID) (dc.Firmware, error)
		GetFirmwareMeta(firmwareEID dc.EID)
	}

	schedulerService struct {
		sites  *dc.SiteNetworks
		cfg    cc.Config
		logger log.Logger
	}
)

func NewSchedulerService(dfg cc.Config, siteNetworks *dc.SiteNetworks) SchedulerService {
	sch = &schedulerService{
		sites:  siteNetworks,
		cfg:    dfg,
		logger: log.With(dfg.Logger, "pkg", "deviceScheduler", "service", "SchedulerService"),
	}
	return sch
}

func (s *schedulerService) BuildFirmwareCatalog() []dc.Firmware {
	level.Debug(s.logger).Log("event", "BuildFirmwareCatalog() called")
	return buildFirmwareCatalog()
}
func (s *schedulerService) ApplySchedule(scheduleEID, deviceEID dc.EID) {
	level.Debug(s.logger).Log("event", "ApplySchedule() called")
}
func (s *schedulerService) StartSchedule(scheduleEID dc.EID) {
	level.Debug(s.logger).Log("event", "StartSchedule() called")
}
func (s *schedulerService) DeleteSchedule(scheduleEID dc.EID) {
	level.Debug(s.logger).Log("event", "DeleteSchedule() called")
}
func (s *schedulerService) UpdateScheduler(dm dc.DeviceMessage) {
	level.Debug(s.logger).Log("event", "UpdateScheduler() called")
}

func (s *schedulerService) ApplyFirmwareUpload(blob []byte) error {
	var err error
	level.Debug(s.logger).Log("event", "ApplyFirmwareUpload() called")
	return err
}
func (s *schedulerService) DeleteFirmware(firmwareEID dc.EID) {
	level.Debug(s.logger).Log("event", "DeleteFirmware() called")
}
func (s *schedulerService) GetFirmware(firmwareEID dc.EID) (dc.Firmware, error) {
	var err error
	level.Debug(s.logger).Log("event", "GetFirmware() called")
	return dc.Firmware{}, err
}
func (s *schedulerService) GetFirmwareMeta(firmwareEID dc.EID) {
	level.Debug(s.logger).Log("event", "GetFirmwareMeta() called")
}
