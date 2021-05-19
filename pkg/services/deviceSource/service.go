package deviceSource

/**
  deviceSource/service.go:

  Handle input streams from MQTT or Demo Logs

*/

import (
	"context"

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
func NewDeviceSourceService(ctx context.Context, repo Repositiory, logger log.Logger) Service {
	dvService = deviceSource{
		repository: repo,
		ctx:        ctx,
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
func Start(ctx context.Context, repo Repositiory) (Service, error) {
	var err error
	logger := log.With(ctx.Value(cc.AppConfig).(cc.Config).Logger, "pkg", "deviceSource")

	level.Debug(logger).Log("msg", "Calling Start()")

	svc := NewDeviceSourceService(ctx, repo, logger)

	return svc, err
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(dvService.logger).Log("msg", "Calling Stop()")

}
