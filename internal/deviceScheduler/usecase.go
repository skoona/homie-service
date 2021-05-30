package deviceScheduler

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	cc "github.com/skoona/homie-service/internal/utils"
)

/*
  	deviceScheduler/usecase-stream.go:

	Build Device Network Inventory
	Provide Interaction to CLI
	Provide Interaction fo HTTP
	Consider Scheduler Feature
*/

type (
	schedulerProvider struct {
		otaStream OTAInteractor
		repo      dc.Repository
		snwk      *dc.SiteNetworks
		cfg       cc.Config
		logger    log.Logger
	}
)

func NewSchedulerProvider(dfg cc.Config, otas OTAInteractor, r dc.Repository) dc.SchedulerProvider {
	slog := log.With(dfg.Logger, "pkg", "deviceScheduler", "service", "SchedulerProvider")
	sch = &schedulerProvider{
		otaStream: otas,
		repo:      r,
		cfg:       dfg,
		logger:    slog,
	}
	return sch
}

func (s *schedulerProvider) ApplySiteNetworks(sn *dc.SiteNetworks) {
	level.Debug(s.logger).Log("event", "ApplySiteNetworks() called")
	s.snwk = sn
}
func (s *schedulerProvider) BuildScheduleCatalog() map[dc.EID]dc.Schedule {
	level.Debug(s.logger).Log("event", "BuildFirmwareCatalog() called")
	return buildScheduleCatalog()
}
func (s *schedulerProvider) Schedules() []dc.Schedule {
	level.Debug(s.logger).Log("event", "Calling Schedules()")
	return []dc.Schedule{}
}
func (s *schedulerProvider) FindSchedulesByDeviceID(deviceID dc.EID) ([]dc.Schedule, error) {
	level.Debug(s.logger).Log("event", "Calling FindSchedulesByDeviceID()")
	return []dc.Schedule{}, nil
}
func (s *schedulerProvider) CreateSchedule(networkName, deviceName string, transport dc.OTATransport, firmware *dc.Firmware) (dc.EID, error) {
	level.Debug(s.logger).Log("event", "Calling CreateSchedule()")
	return dc.NewEID(), nil
}
func (s *schedulerProvider) DeleteSchedule(scheduleID dc.EID) error {
	level.Debug(s.logger).Log("event", "Calling DeleteSchedule()")
	return nil
}
func (s *schedulerProvider) BuildFirmwareCatalog() []dc.Firmware {
	level.Debug(s.logger).Log("event", "BuildFirmwareCatalog() called")
	return buildFirmwareCatalog()
}
func (s *schedulerProvider) Firmwares() []dc.Firmware {
	level.Debug(s.logger).Log("event", "Calling Firmwares()")
	return []dc.Firmware{}
}
func (s *schedulerProvider) GetFirmware(id dc.EID) (dc.Firmware, error) {
	level.Debug(s.logger).Log("event", "Calling GetFirmware()")
	return dc.Firmware{}, nil
}
func (s *schedulerProvider) CreateFirmware(path string) error {
	level.Debug(s.logger).Log("event", "Calling CreateFirmware()")
	return nil
}
func (s *schedulerProvider) DeleteFirmware(id dc.EID) error {
	level.Debug(s.logger).Log("event", "Calling DeleteFirmware()")
	return nil
}
