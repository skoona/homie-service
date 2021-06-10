package deviceSource

/*
  deviceSource/usecase-core.go:

  dc.DeviceEventProvider Service Implementation

  The design goal for this package:
	* Implement dc.DeviceEventProvider for network event interactions with deviceCore
*/

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
)


/*
 * ActivateStreamProvider
 */
func (s *deviceSource) ActivateStreamProvider() {
	ConsumeFromStreamProvider(s.dStream.ActivateNotifications(), s.logger)
}

func (s *deviceSource) ApplyDeviceEvent(dm dc.DeviceMessage) {
	plog := log.With(s.logger, "method", "ApplyDeviceEvent()")

	dc.ConsumeFromDeviceSource(dm)

	level.Debug(plog).Log("DeviceID ", dm.DeviceID)
}

/*
 * PublishToStreamProvider
 */
func (s *deviceSource) PublishToStreamProvider(dv dc.Device) {
	plog := log.With(s.logger, "method", "PublishDeviceStream()")

	s.dStream.GetPublishChannel() <- dv

	level.Debug(plog).Log("DeviceID ", dv.ID)
}

// handle incoming device stream events
func (s *deviceSource) ConsumeDeviceStream(dm dc.DeviceMessage) error {
	var err error
	plog := log.With(s.logger, "method", "ConsumeDeviceStream")

	err = s.repository.Store(dm)
	if err != nil {
		level.Error(plog).Log("error", err)
		return err
	}

	s.ApplyDeviceEvent(dm) // to Core

	level.Debug(plog).Log("DeviceID ", dm.DeviceID)

	return err
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
	err = s.repository.Remove(dm)
	if err != nil {
		level.Error(plog).Log("error", err)
		return err
	}

	s.PublishToStreamProvider(dv)

	level.Debug(plog).Log("DeviceID ", dv.ID)
	return err
}
