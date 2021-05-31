package deviceSource

/*
  deviceSource/usecase-core.go:

  DeviceSource Service Implementation
*/

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
)

func (s *deviceSource) ApplyDeviceEvent(dm dc.DeviceMessage) {
	logger := log.With(s.logger, "method", "ApplyDeviceEvent()")

	dc.ConsumeFromDeviceSource(dm)

	level.Debug(logger).Log("DeviceID ", dm.DeviceID)
}

// handle incoming core events
func (s *deviceSource) HandleCoreEvent(dm dc.DeviceMessage) error {
	var err error
	plog := log.With(s.logger, "method", "HandleCoreEvent()")

	// can only be a delete request
	if dm.Value == nil {
		err = s.repository.Remove(dm)
	} else {
		err = s.repository.Store(dm)
	}
	if err != nil {
		level.Error(plog).Log("error", err)
		return err
	}

	// TODO: we should unwrap and send deletes
	s.PublishToStreamProvider(dm)

	level.Debug(plog).Log("DeviceID ", dm.DeviceID)
	return err
}
