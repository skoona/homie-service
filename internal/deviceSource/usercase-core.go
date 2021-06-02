package deviceSource

/*
  deviceSource/usecase-core.go:

  DeviceSource Service Implementation
*/

import (
	"fmt"
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
func (s *deviceSource) HandleCoreEvent(dv dc.Device) error {
	var err error
	plog := log.With(s.logger, "method", "HandleCoreEvent()")

	dm := dc.DeviceMessage{
		ID:        99,
		Value:     nil,
		DeviceID: []byte(dv.Name),
		NetworkID: []byte(dv.Parent),
		HomieType: dv.ElementType,
		TopicS: fmt.Sprintf("%s/%s/$state", dv.Parent, dv.Name),
	}

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
	s.PublishToStreamProvider(dv)

	level.Debug(plog).Log("DeviceID ", dv.ID)
	return err
}
