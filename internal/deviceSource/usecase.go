package deviceSource

/*
  deviceSource/usecase.go:

  DeviceSource Service Implementation
*/

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	cc "github.com/skoona/homie-service/internal/utils"
)

type (
	// Service Implementation
	deviceSource struct {
		repository Repository
		coreSvc    dc.DeviceSourceInteractor
		cfg        cc.Config
		logger     log.Logger
	}
)

// Retained DeviceSource Service, once created
var (
	dvService deviceSource
	logger    log.Logger
)

func (s *deviceSource) ApplyDMEvent(dm dc.DeviceMessage) error {
	logger := log.With(s.logger, "method", "ApplyDMEvent")

	err := s.repository.Store(dm)
	if err != nil {
		level.Error(logger).Log("error", err)
		return err
	}

	// also sent it to core
	if toCore == nil {
		toCore, err = s.coreSvc.GetCoreRequestChannel()
		if err != nil {
			level.Error(logger).Log("error", err)
			return err
		}
	}
	toCore <- dm // send to Core

	level.Debug(logger).Log("DeviceID ", dm.DeviceID)

	return err
}

func (s *deviceSource) ApplyOTAEvent(dm dc.DeviceMessage) error {
	var err error
	logg := log.With(s.logger, "method", "ApplyOTAEvent")

	level.Debug(logg).Log("DeviceID ", dm.DeviceID)

	return err
}

// handle incoming core events
func (s *deviceSource) ApplyCoreEvent(dm dc.DeviceMessage) error {
	var err error
	logg := log.With(s.logger, "method", "ApplyCoreEvent")

	err = s.repository.Store(dm)
	if err != nil {
		level.Error(logg).Log("error", err)
		return err
	}

	level.Debug(logg).Log("DeviceID ", dm.DeviceID)

	return err
}

func (s *deviceSource) ApplySchedulerEvent(dm dc.DeviceMessage) error {
	var err error
	logg := log.With(s.logger, "method", "ApplySchedulerEvent")

	level.Debug(logg).Log("DeviceID ", dm.DeviceID)

	return err
}

// Receive/send DM from Channel
// Receive/send OTA from Channel
// Receive/Send to Scheduler
// Receive/Send to Core

// Route OTA to/from Scheduler
// Route DM to Repository and Core Service
