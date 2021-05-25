package deviceSource

/**
  deviceSource/service.go:

  Handle input streams from MQTT or Demo Logs

*/

import (
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	cc "github.com/skoona/homie-service/internal/utils"
)

// Message Interface for MQTT and Demo
type (
	//Incoming Messages
	Service interface {
		ApplyCoreEvent(dm dc.DeviceMessage) error
		ApplyDMEvent(dm dc.DeviceMessage) error
		ApplyOTAEvent(dm dc.DeviceMessage) error
		ApplySchedulerEvent(dm dc.DeviceMessage) error
	}

	// Device Source Storage Repository
	Repository interface {
		Store(d dc.DeviceMessage) error
	}

	//Incoming Messages
	DeviceMessageHandler interface {
		CreateDemoDeviceMessage(topic string, payload []byte, idCounter uint16, retained bool, qos byte) (dc.DeviceMessage, error)
		CreateQueueDeviceMessage(qmsg dc.QueueMessage) (dc.DeviceMessage, error)
		GetProviderRequestChannel() (chan dc.DeviceMessage, error)
		GetProviderResponseChannel() (chan dc.DeviceMessage, error)
		GetOTARequestChannel() (chan dc.DeviceMessage, error)
		GetOTAResponseChannel() (chan dc.DeviceMessage, error)
		GetSchedulerRequestChannel() (chan dc.DeviceMessage, error)
		GetSchedulerResponseChannel() (chan dc.DeviceMessage, error)
		FromMQTTProvider(qmsg dc.QueueMessage) error
		FromDemoProvider(topic string, payload []byte, idCounter uint16, retained bool, qos byte) error
		FromOTAProvider(qmsg dc.QueueMessage) error
		FromScheduler(qmsg dc.QueueMessage) error
	}
)

/**
 * NewDeviceSourceService()
 *
 *  Create a New NewDeviceSourceService and initializes it.
 */
func NewDeviceSourceService(cfg cc.Config, repo Repository, coreSvc dc.DeviceSourceInteractor, logger log.Logger) Service {
	dvService = &deviceSource{
		repository: repo,
		coreSvc:    coreSvc,
		cfg:        cfg,
		logger:     logger,
	}
	return dvService
}

/**
 * NewDeviceMessageService()
 *
 *  Create a New NewDeviceSourceService and initializes it.
 */
func NewDeviceMessageHandler(cfg cc.Config, logger log.Logger) DeviceMessageHandler {
	dmHandler = &deviceMessageHandler{
		cfg:    cfg,
		logger: logger,
	}
	return dmHandler
}

var (
	toScheduler   chan dc.DeviceMessage // out
	fromScheduler chan dc.DeviceMessage // in

	toProvider   chan dc.DeviceMessage // out
	fromProvider chan dc.DeviceMessage // in

	toOTAProvider   chan dc.DeviceMessage // out
	fromOTAProvider chan dc.DeviceMessage // in

	toCore   chan dc.DeviceMessage // out
	fromCore chan dc.DeviceMessage // in
)

/**
 * ConsumeFromCore
 * Handles incoming channel DM message
 */
func ConsumeFromCore(consumer chan dc.DeviceMessage) error {
	/*
	 * Create a Go Routine for the Providers Channel to
	 */
	go func(dmChan chan dc.DeviceMessage) {
		level.Debug(dmHandler.logger).Log("event", "ConsumeFromCore(gofunc) called")
		for msg := range dmChan { // read until closed
			err := dvService.ApplyCoreEvent(msg)
			if err != nil {
				level.Error(dmHandler.logger).Log("method", "ConsumeFromCore(gofunc)", "error", err.Error())
			}

			level.Debug(dmHandler.logger).Log("method", "ConsumeFromCore(gofunc)", "queue depth", len(dmChan))
		}
		level.Debug(dmHandler.logger).Log("method", "ConsumeFromCore()", "event", "Completed")
	}(consumer)

	return nil
}

func ChannelsForCore() (chan dc.DeviceMessage, chan dc.DeviceMessage, error) {
	var err error
	if toCore == nil {
		toCore, err = dvService.coreSvc.GetCoreRequestChannel() // averages 120 on startup
	}

	if fromCore == nil {
		fromCore, err = dvService.coreSvc.GetCoreResponseChannel() // averages 120 on startup
		if err == nil {
			err = ConsumeFromCore(fromCore)
		}
	}

	if nil == toCore || nil == fromCore {
		err = errors.New("create core channels failed")
		level.Error(logger).Log("error", err.Error())
		return nil, nil, err
	}

	return toCore, fromCore, err
}

/*
 * Start()
 *
 * Initialize this service
 */
func Start(cfg cc.Config, repo Repository, coreSvc dc.DeviceSourceInteractor) (DeviceMessageHandler, error) {
	var err error
	logger = log.With(cfg.Logger, "pkg", "deviceSource")

	level.Debug(logger).Log("event", "Calling Start()")

	NewDeviceSourceService(cfg, repo, coreSvc, logger)
	dmh := NewDeviceMessageHandler(cfg, cfg.Logger)

	level.Debug(logger).Log("event", "Start() Completed")

	return dmh, err
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")

	if nil != toScheduler {
		close(toScheduler) // only the send should shutdown channels
	}
	if nil != toOTAProvider {
		close(toOTAProvider)
	}
	if nil != toProvider {
		close(toProvider)
	}
	if nil != toCore {
		close(toCore)
	}

	if nil != fromProvider {
		close(fromProvider) // only the send should shutdown channels
	}
	if nil != fromOTAProvider {
		close(fromOTAProvider) // only the send should shutdown channels
	}
	if nil != fromScheduler {
		close(fromScheduler) // only the send should shutdown channels
	}

	// close(fromCore)
	level.Debug(logger).Log("event", "Stop() Completed")
}
