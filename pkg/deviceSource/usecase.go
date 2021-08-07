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
	dc "github.com/skoona/homie-service/pkg/deviceCore"
)


/*
 * ActivateStreamProvider
 */
func (s *deviceSource) ActivateStreamProvider() {
	StoreDeviceStream(s.sStream, s.logger)
	ConsumeFromStreamProvider(s.dStream.ActivateNotifications(), s.logger)
}

func (s *deviceSource) ApplyDeviceEvent(dm dc.DeviceMessage) {
	plog := log.With(s.logger, "method", "ApplyDeviceEvent()")

	dc.ConsumeFromDeviceSource(dm)

	level.Debug(plog).Log("DeviceID ", dm.DeviceID)
}

// PublishToStreamProvider send a message on the outgoing mqtt channel
func (s *deviceSource) PublishToStreamProvider(dm dc.DeviceMessage) {
	plog := log.With(s.logger, "method", "PublishToStreamProvider()")

	s.dStream.GetPublishChannel() <- dm

	level.Debug(plog).Log("DeviceID ", dm.DeviceID)
}
