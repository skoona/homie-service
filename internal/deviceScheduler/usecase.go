package deviceScheduler

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	cc "github.com/skoona/homie-service/internal/utils"
)

/*
  	deviceScheduler/usecase.go:

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
func (s *schedulerProvider) BuildScheduleCatalog() map[string]dc.Schedule {
	level.Debug(s.logger).Log("event", "BuildFirmwareCatalog() called")
	sch.snwk.Schedules = buildScheduleCatalog(s.logger)
	return sch.snwk.Schedules
}
func (s *schedulerProvider) Schedules() []dc.Schedule {
	level.Debug(s.logger).Log("event", "Calling Schedules()")
	values := []dc.Schedule{}
	if len(s.snwk.Schedules) > 0 {
		for _, val := range s.snwk.Schedules {
			values = append(values, val)
		}
	}
	return values
}
func (s *schedulerProvider) FindScheduleByDeviceID(deviceID string) *dc.Schedule {
	level.Debug(s.logger).Log("event", "Calling FindSchedulesByDeviceID()")
	schedules := []*dc.Schedule{}
	for _, schedule := range s.snwk.Schedules {
		if schedule.DeviceID == deviceID {
			schedules = append(schedules, &schedule)
		}
	}
	if len(schedules) >= 1 {
		return schedules[0]
	}
	return &dc.Schedule{}
}
func (s *schedulerProvider) CreateSchedule(networkName string, deviceID string, transport dc.OTATransport, firmwareID dc.EID) (string, error) {
	level.Debug(s.logger).Log("event", "Calling CreateSchedule()")
	schedule := NewSchedule(networkName, deviceID, transport, firmwareID)
	s.snwk.Schedules[schedule.ID] = schedule
	err := s.repo.StoreSchedule(schedule)
	return schedule.ID, err
}
func (s *schedulerProvider) DeleteSchedule(scheduleID string) error {
	level.Debug(s.logger).Log("event", "Calling DeleteSchedule()")
	delete(s.snwk.Schedules, scheduleID)
	err := s.repo.RemoveSchedule(scheduleID)
	return err
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
func (s *schedulerProvider) CreateFirmware(path string) (dc.EID, error) {
	level.Debug(s.logger).Log("event", "Calling CreateFirmware()")
	fw, err := NewFirmware(path)
	s.snwk.Firmwares = append(s.snwk.Firmwares, fw)
	return fw.ID, err
}
func (s *schedulerProvider) DeleteFirmware(id dc.EID) error {
	level.Debug(s.logger).Log("event", "Calling DeleteFirmware()")
	return nil
}
