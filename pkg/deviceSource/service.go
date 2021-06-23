package deviceSource

/**
  deviceSource/service.go:

  The design goal for this package:
	* Define the StreamProvider interface for network device events
	* Implement dc.DeviceEventProvider for network event interactions with deviceCore
	* Define Repository interface for storage of network device events & schedules
*/

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	cc "github.com/skoona/homie-service/pkg/utils"
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

	// Core Service Implementation
	deviceSource struct {
		cfg        cc.Config
		repository dc.Repository
		dStream    StreamProvider
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
func NewDeviceSourceService(cfg cc.Config, repo dc.Repository, stream StreamProvider, plog log.Logger) dc.DeviceEventProvider {
	dvService = &deviceSource{
		repository: repo,
		dStream:    stream,
		cfg:        cfg,
		logger:     log.With(plog, "service", "Service"),
	}
	return dvService
}

/*
 * Start()
 *
 * Initialize this service
 */
func Start(dfg cc.Config, repo dc.Repository, dProvider StreamProvider) (dc.DeviceEventProvider, error) {
	var err error
	logger = log.With(dfg.Logger, "pkg", "deviceSource")
	level.Debug(logger).Log("event", "Calling Start()")

	s := NewDeviceSourceService(dfg, repo, dProvider, logger)

	level.Debug(logger).Log("event", "Start() Completed")

	return s, err
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")

	level.Debug(logger).Log("event", "Stop() Completed")
}
