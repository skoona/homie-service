package deviceSource

/**
  deviceSource/service.go:

  Handle input streams from MQTT or Demo Logs

*/

import (
	"github.com/go-kit/kit/log"
)

// Message Interface for MQTT and Demo
type (
	//Incoming Messages
	Service interface {
		ApplyCoreEvent(dm *DeviceMessage) error
		ApplyOTAEvent(dm *DeviceMessage) error
		ApplyDiscoveryEvent(dm *DeviceMessage) error
	}

	// Device Source Storage Repository
	Repositiory interface {
		Store(d *DeviceMessage) error
	}
)

/**
 * NewDeviceSourceService()
 *
 *  Create a New NewDeviceSourceService and initializes it.
 */
func NewDeviceSourceService(repo Repositiory, logger log.Logger) Service {
	dvService = deviceSource{
		repository: repo,
		logger:     log.With(logger, "pkg", "deviceSource"),
	}
	return &dvService
}

/**
 * NewDeviceMessage()
 *
 *  Create a New DeviceMessage and initializes it.
 */
func NewDeviceMessage(topic string, payload []byte, idCounter uint16, retained bool, qos byte) (DeviceMessage, error) {
	return buildDeviceMessage(topic, payload, idCounter, retained, qos)
}
