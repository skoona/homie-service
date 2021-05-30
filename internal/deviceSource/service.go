package deviceSource

/**
  deviceSource/service.go:

  Handle input streams from MQTT or Demo Logs

*/

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	cc "github.com/skoona/homie-service/internal/utils"
)

// Message Interface for MQTT and Demo
type (
	StreamProvider interface {
		ActivateNotifications() chan dc.DeviceMessage
		GetPublishChannel() chan dc.DeviceMessage
		GetNotifyChannel() chan dc.DeviceMessage
		CreateQueueDeviceMessage(qmsg dc.QueueMessage) dc.DeviceMessage
		CreateDemoDeviceMessage(topic string, payload []byte, idCounter uint16, retained bool, qos byte) dc.DeviceMessage
	}

	//Incoming Messages
	Service interface {
		ActivateStreamProvider()
		PublishCoreEvent(dm dc.DeviceMessage)
		ConsumeCoreEvent(dm dc.DeviceMessage) error
		PublishToStreamProvider(dm dc.DeviceMessage)
		ConsumeDeviceStream(dm dc.DeviceMessage) error
	}

	// Device Source Storage Repository
	Repository interface {
		Store(d dc.DeviceMessage) error
	}

	// Core Service Implementation
	deviceSource struct {
		cfg        cc.Config
		repository Repository
		dStream    StreamProvider
		coreSvc    dc.DeviceSourceInteractor
		logger     log.Logger
	}
)

// Retained DeviceSource Service, once created
var (
	dvService *deviceSource
	logger    log.Logger
)

/**
 * NewDeviceSourceService()
 *
 *  Create a New NewDeviceSourceService and initializes it.
 */
func NewDeviceSourceService(cfg cc.Config, repo Repository, stream StreamProvider, plog log.Logger) Service {
	dvService = &deviceSource{
		repository: repo,
		dStream:    stream,
		cfg:        cfg,
		logger:     log.With(plog, "service", "Service"),
	}
	return dvService
}

/*
 * ActivateStreamProvider
 */
func (s *deviceSource) ActivateStreamProvider() {
	ConsumeFromStreamProvider(s.dStream.ActivateNotifications(), s.logger)
}

/*
 * Start()
 *
 * Initialize this service
 */
func Start(dfg cc.Config, repo Repository, dProvider StreamProvider) (Service, error) {
	var err error
	logger = log.With(dfg.Logger, "pkg", "deviceSource")
	level.Debug(logger).Log("event", "Calling Start()")

	NewDeviceSourceService(dfg, repo, dProvider, logger)

	dvService.ActivateStreamProvider() // start things moving

	level.Debug(logger).Log("event", "Start() Completed")

	return dvService, err
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")

	level.Debug(logger).Log("event", "Stop() Completed")
}
