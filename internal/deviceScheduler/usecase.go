package deviceScheduler

import (
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
)

/*
  	deviceScheduler/usecase.go:

	Build Device Network Inventory
	Provide Interaction to CLI
	Provide Interaction fo HTTP
	Consider Scheduler Feature
*/

func (s *schedulerService) ApplySchedule(scheduleEID, deviceEID dc.EID) {
	level.Debug(s.logger).Log("event", "ApplySchedule() called")
}
func (s *schedulerService) StartSchedule(scheduleEID dc.EID) {
	level.Debug(s.ogger).Log("event", "StartSchedule() called")
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
