package deviceScheduler

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	cc "github.com/skoona/homie-service/pkg/utils"
	"os"
	"strings"
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
	return  buildScheduleCatalog(s.logger)
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
	level.Debug(s.logger).Log("event", "Calling FindSchedulesByDeviceID()", "device", deviceID)
	schedules := []*dc.Schedule{}
	for _, schedule := range s.snwk.Schedules {
		if strings.Contains(schedule.DeviceID, deviceID) {
			schedules = append(schedules, &schedule)
			break
		}
	}

	if len(schedules) >= 1 {
		return schedules[0] // return first one, shouldn't be more that one per device
	}
	return &dc.Schedule{}
}
func (s *schedulerProvider) CreateSchedule(networkName string, deviceID string, transport dc.OTATransport, firmwareID dc.EID) (string, error) {
	level.Debug(s.logger).Log("event", "Calling CreateSchedule()", "network", networkName, "deviceID", deviceID)
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
	return buildFirmwareCatalog()
}
func (s *schedulerProvider) Firmwares() []dc.Firmware {
	level.Debug(s.logger).Log("event", "Calling Firmwares()")
	return sch.snwk.Firmwares
}
func (s *schedulerProvider) GetFirmware(id dc.EID) (dc.Firmware, error) {
	level.Debug(s.logger).Log("event", "Calling GetFirmware()")
	var fw dc.Firmware
	for _, val := range s.snwk.Firmwares {
		if id == val.ID {
			fw = val
		}
	}
	if fw.ID == "" {
		return fw, fmt.Errorf("firmware with ID=%s, was not found", string(id))
	}

	return fw, nil
}

// CreateFirmware File Params
//{
//  "qquuid"=>"43636118-c89e-48bc-a354-c52799628757",  -- qq...
//  "qqfilename"=>"Butterfly_HubbleVargas_5075.jpg",
//  "qqtotalfilesize"=>"981840",
//  "qqfile"=>{
//      :filename=>"Butterfly_HubbleVargas_5075.jpg",
//      :type=>"image/jpeg",
//      :name=>"qqfile",
//      :tempfile=>#<Tempfile:/var/folders/gm/6hf5xns52cddrgnqz6zz48bxyht1rq/T/RackMultipart20190308-29136-w5rdv7.jpg>,
//      :head=>"Content-Disposition: form-data; name=\"qqfile\"; filename=\"Butterfly_HubbleVargas_5075.jpg\"\r\nContent-Type: image/jpeg\r\n"
//    }
//}
func (s *schedulerProvider) CreateFirmware(srcFile, dstFile string) (dc.EID, error) {
	level.Debug(s.logger).Log("event", "Calling CreateFirmware()")
	fw, err := NewFirmware(srcFile, dstFile)
	if err == nil {
		s.snwk.Firmwares = append(s.snwk.Firmwares, fw)
	}
	return fw.ID, err
}

func (s *schedulerProvider) DeleteFirmware(id dc.EID) error {
	// todo: should the filesystem file be removed here?
	// Todo: appears to be done in Core Service "RemoveFirmwareByID()"
	level.Debug(s.logger).Log("event", "Calling DeleteFirmware()", "firmwareID", id)
	fwc := []dc.Firmware{}
	var firmware dc.Firmware
	var err error
	for _, fw := range s.snwk.Firmwares {
		if id != fw.ID {
			fwc = append(fwc, fw)
		} else {
			firmware = fw
		}
	}
	if len(fwc) == 0 {
		return fmt.Errorf("firmware with ID=%s, was not found.", id)
	}
	
	// remove from collection
	s.snwk.Firmwares = fwc

	// remove from file systems, unless testing
	if _, err := os.Stat(firmware.Path); err == nil {
		if s.cfg.RunMode == "live" {
			err = os.Remove(firmware.Path)
			if err != nil {
				err = fmt.Errorf("firmware with name {%s} was  not found to remove(): error=%s", firmware.Name, err.Error())
				level.Error(s.logger).Log("error", err.Error())
			}
		} else {
			err = fmt.Errorf("delete request ignored in test and demo run modes: %s " , s.cfg.RunMode)
			level.Info(s.logger).Log("event", "ignored", "cause", err.Error())
		}

	} else if os.IsNotExist(err) {
		err = fmt.Errorf("firmware with name {%s} was  not found to remove(): error=%s", firmware.Name, err.Error())
		level.Error(s.logger).Log("error", err.Error())
	}
	return err
}
