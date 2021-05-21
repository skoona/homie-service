package deviceSource

/**
  deviceSource/service.go:

  Handle input streams from MQTT or Demo Logs

*/

import (
	"errors"

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

var (
	toScheduler   chan DeviceMessage // out
	fromScheduler chan DeviceMessage // in

	toProvider   chan DeviceMessage // out
	fromProvider chan DeviceMessage // in

	toOTAProvider   chan DeviceMessage // out
	fromOTAProvider chan DeviceMessage // in

	toCore   chan DeviceMessage // out
	fromCore chan DeviceMessage // in
)

func ChannelsForScheduler() (chan DeviceMessage, chan DeviceMessage, error) {
	var err error
	if toScheduler == nil {
		toScheduler = make(chan DeviceMessage, 256) // averages 120 on startup
	}
	if fromScheduler == nil {
		fromScheduler = make(chan DeviceMessage, 256) // averages 120 on startup
	}

	if nil == toScheduler || nil == fromScheduler {
		err = errors.New("create scheduler channels failed")
		level.Error(dvService.logger).Log("error", err.Error())
		return nil, nil, err
	}

	return toScheduler, fromScheduler, err
}
func ChannelsForOTAProviders() (chan DeviceMessage, chan DeviceMessage, error) {
	var err error
	if toOTAProvider == nil {
		toOTAProvider = make(chan DeviceMessage, 256) // averages 120 on startup
	}
	if fromOTAProvider == nil {
		fromOTAProvider = make(chan DeviceMessage, 256) // averages 120 on startup
	}

	if nil == toOTAProvider || nil == fromOTAProvider {
		err = errors.New("create ota channels failed")
		level.Error(dvService.logger).Log("error", err.Error())
		return nil, nil, err
	}

	return toOTAProvider, fromOTAProvider, err
}
func ChannelsForDMProviders() (chan DeviceMessage, chan DeviceMessage, error) {
	var err error
	if toProvider == nil {
		toProvider = make(chan DeviceMessage, 256)
	}
	if fromProvider == nil {
		fromProvider = make(chan DeviceMessage, 256)
	}

	if nil == toProvider || nil == fromProvider {
		err = errors.New("create provider channels failed")
		level.Error(dvService.logger).Log("error", err.Error())
		return nil, nil, err
	}

	return toProvider, fromProvider, err
}
func ChannelsForCore() (chan DeviceMessage, chan DeviceMessage, error) {
	var err error
	if toCore == nil {
		toCore = make(chan DeviceMessage, 256) // averages 120 on startup
	}

	if fromCore == nil {
		fromCore = make(chan DeviceMessage, 256) // averages 120 on startup
	}

	if nil == toCore || nil == fromCore {
		err = errors.New("create core channels failed")
		level.Error(dvService.logger).Log("error", err.Error())
		return nil, nil, err
	}

	return toCore, fromCore, err
}

/*
 * Start()
 *
 * Initialize this service
 */
func Start(cfg cc.Config, repo Repositiory) (Service, error) {
	var err error
	logger := log.With(cfg.Logger, "pkg", "deviceSource")

	level.Debug(logger).Log("event", "Calling Start()")

	svc := NewDeviceSourceService(cfg, repo, logger)

	level.Debug(dvService.logger).Log("event", "Start() Completed")

	return svc, err
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(dvService.logger).Log("event", "Calling Stop()")

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

	// close(fromCore)
	// close(fromScheduler)
	// close(fromProvider)
	// close(fromOTAProvider)
	level.Debug(dvService.logger).Log("event", "Stop() Completed")
}
