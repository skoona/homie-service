package deviceSource

/**
  deviceSource/service.go:

  Handle input streams from MQTT or Demo Logs

*/

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	cc "github.com/skoona/homie-service/pkg/utils"
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
func NewDeviceSourceService(cfg cc.Config, repo Repositiory, logger log.Logger) Service {
	dvService = deviceSource{
		repository: repo,
		cfg:        cfg,
		logger:     logger,
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

/*
 * Start()
 *
 * Initialize this service
 */
func Start(cfg cc.Config, repo Repositiory) (Service, error) {
	var err error
	logger := log.With(cfg.Logger, "pkg", "deviceSource")

	level.Debug(logger).Log("msg", "Calling Start()")

	svc := NewDeviceSourceService(cfg, repo, logger)

	return svc, err
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(dvService.logger).Log("msg", "Calling Stop()")

}
