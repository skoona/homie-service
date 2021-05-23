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
		repository Repositiory
		cfg        cc.Config
		logger     log.Logger
	}
)

// Retained DeviceSource Service, once created
var dvService deviceSource

func (s *deviceSource) ApplyDMEvent(dm *dc.DeviceMessage) error {
	logger := log.With(s.logger, "method", "ApplyDMEvent")

	err := s.repository.Store(dm)
	if err != nil {
		level.Error(logger).Log("error", err)
	}

	level.Debug(logger).Log("DeviceID ", dm.DeviceID)

	return err
}

func (s *deviceSource) ApplyOTAEvent(dm *dc.DeviceMessage) error {
	var err error
	logger := log.With(s.logger, "method", "ApplyOTAEvent")

	level.Debug(logger).Log("DeviceID ", dm.DeviceID)

	return err
}

func (s *deviceSource) ApplyCoreEvent(dm *dc.DeviceMessage) error {
	var err error
	logger := log.With(s.logger, "method", "ApplyCoreEvent")

	level.Debug(logger).Log("DeviceID ", dm.DeviceID)

	return err
}

func (s *deviceSource) ApplySchedulerEvent(dm *dc.DeviceMessage) error {
	var err error
	logger := log.With(s.logger, "method", "ApplySchedulerEvent")

	level.Debug(logger).Log("DeviceID ", dm.DeviceID)

	return err
}

// Receive/send DM from Channel
// Receive/send OTA from Channel
// Receive/Send to Scheduler
// Receive/Send to Core

// Route OTA to/from Scheduler
// Route DM to Repository and Core Service
