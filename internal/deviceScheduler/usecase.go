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
func (s *schedulerProvider) ActivateStreamProvider() {
	// Initialize a Message Channel
	consumeFromOTAStreamProvider(s.otaStream.EnableTriggers(), s.logger)
}
func (s *schedulerProvider) ApplySiteNetworks(sn *dc.SiteNetworks) {
	level.Debug(s.logger).Log("event", "ApplySiteNetworks() called")
	s.snwk = sn
}
func (s *schedulerProvider) BuildScheduleCatalog() map[dc.EID]dc.Schedule {
	level.Debug(s.logger).Log("event", "BuildFirmwareCatalog() called")
	sch.snwk.Schedules = buildScheduleCatalog()
	return sch.snwk.Schedules
}
func (s *schedulerProvider) Schedules() []dc.Schedule {
	level.Debug(s.logger).Log("event", "Calling Schedules()")
	values := []dc.Schedule{}
	for _, val := range s.snwk.Schedules {
		values = append(values, val)
	}
	return values
}
func (s *schedulerProvider) FindSchedulesByDeviceEID(deviceEID dc.EID) []dc.Schedule {
	level.Debug(s.logger).Log("event", "Calling FindSchedulesByDeviceID()")
	return []dc.Schedule{}
}
func (s *schedulerProvider) CreateSchedule(networkName string, deviceID dc.EID, transport dc.OTATransport, firmware *dc.Firmware) (dc.EID, error) {
	level.Debug(s.logger).Log("event", "Calling CreateSchedule()")
	// s.repo.StoreSchedule(schedule)
	return dc.NewEID(), nil
}
func (s *schedulerProvider) DeleteSchedule(scheduleID dc.EID) error {
	level.Debug(s.logger).Log("event", "Calling DeleteSchedule()")
	// s.repo.RemoveSchedule(schedule)
	return nil
}
func (s *schedulerProvider) BuildFirmwareCatalog() []dc.Firmware {
	level.Debug(s.logger).Log("event", "BuildFirmwareCatalog() called")
	sch.snwk.Firmwares = buildFirmwareCatalog()
	return sch.snwk.Firmwares
}
func (s *schedulerProvider) Firmwares() []dc.Firmware {
	level.Debug(s.logger).Log("event", "Calling Firmwares()")
	return sch.snwk.Firmwares
}
func (s *schedulerProvider) GetFirmware(id dc.EID) (dc.Firmware, error) {
	level.Debug(s.logger).Log("event", "Calling GetFirmware()")
	return dc.Firmware{}, nil
}
func (s *schedulerProvider) CreateFirmware(path string) error {
	level.Debug(s.logger).Log("event", "Calling CreateFirmware()")
	fw, err := NewFirmware(path)
	s.snwk.Firmwares = append(s.snwk.Firmwares, fw)
	return err
}
func (s *schedulerProvider) DeleteFirmware(id dc.EID) error {
	level.Debug(s.logger).Log("event", "Calling DeleteFirmware()")
	return nil
}
