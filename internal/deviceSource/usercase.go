package deviceSource

/*
  deviceSource/usecase.go:

  dc.DeviceEventProvider Service Implementation

  The design goal for this package:
	* Implement dc.DeviceEventProvider for network event interactions with deviceCore
*/

import (
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
func (s *deviceSource) PublishToStreamProvider(dm dc.DeviceMessage) {
	plog := log.With(s.logger, "method", "PublishToStreamProvider()")

	s.dStream.GetPublishChannel() <- dm

	level.Debug(plog).Log("DeviceID ", dm.DeviceID)
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
