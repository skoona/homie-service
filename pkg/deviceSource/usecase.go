package deviceSource

/*
  deviceSource/usecase.go:

  DeviceSource Service Implementation
*/

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type (
	deviceSource struct {
		repository Repositiory
		logger     log.Logger
	}
)

// Retained DeviceSource Service, once created
var dvService deviceSource

func (s *deviceSource) ApplyCoreEvent(dm *DeviceMessage) error {
	var err error
	logger := log.With(s.logger, "method", "ApplyCoreEvent")
	level.Error(logger).Log("err", err)
	level.Debug(logger).Log("key", "value")
	logger.Log("DeviceID ", dm.DeviceID)

	return err
}

func (s *deviceSource) ApplyOTAEvent(dm *DeviceMessage) error {
	var err error
	logger := log.With(s.logger, "method", "ApplyOTAEvent")
	level.Error(logger).Log("err", err)
	level.Debug(logger).Log("key", "value")
	logger.Log("DeviceID ", dm.DeviceID)

	return err
}

func (s *deviceSource) ApplyDiscoveryEvent(dm *DeviceMessage) error {
	var err error
	logger := log.With(s.logger, "method", "ApplyDiscoveryEvent")
	level.Error(logger).Log("err", err)
	level.Debug(logger).Log("key", "value")
	logger.Log("DeviceID ", dm.DeviceID)

	return err
}
