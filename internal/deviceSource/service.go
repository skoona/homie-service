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
		ApplyCoreEvent(dm *dc.DeviceMessage) error
		ApplyDMEvent(dm *dc.DeviceMessage) error
		ApplyOTAEvent(dm *dc.DeviceMessage) error
		ApplySchedulerEvent(dm *dc.DeviceMessage) error
	}

	// Device Source Storage Repository
	Repositiory interface {
		Store(d *dc.DeviceMessage) error
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
		level.Debug(dvService.logger).Log("event", "ConsumeFromCore(gofunc) called")
		for msg := range dmChan { // read until closed
			err := dvService.ApplyCoreEvent(&msg)
			if err != nil {
				level.Error(dvService.logger).Log("method", "ConsumeFromCore(gofunc)", "error", err.Error())
			}

			level.Debug(dvService.logger).Log("method", "ConsumeFromCore(gofunc)", "queue depth", len(dmChan))
		}
		level.Debug(dvService.logger).Log("method", "ConsumeFromCore()", "event", "Completed")
	}(consumer)

	return nil
}

/**
 * ConsumeFromProviders
 * Handles incoming channel DM message
 */
func ConsumeFromDMProviders(consumer chan dc.DeviceMessage) error {
	/*
	 * Create a Go Routine for the Providers Channel to
	 */
	go func(dmChan chan dc.DeviceMessage) {
		level.Debug(dvService.logger).Log("event", "ConsumeFromDMProviders(gofunc) called")
		for msg := range dmChan { // read until closed
			err := dvService.ApplyDMEvent(&msg)
			if err != nil {
				level.Error(dvService.logger).Log("method", "ConsumeFromDMProviders(gofunc)", "error", err.Error())
			}

			// SHOULD WE SEND TO CORE HERE

			level.Debug(dvService.logger).Log("method", "ConsumeFromDMProviders(gofunc)", "queue depth", len(dmChan))
		}
		level.Debug(dvService.logger).Log("method", "ConsumeFromDMProviders()", "event", "Completed")
	}(consumer)

	return nil
}

/**
 * ConsumeFromOTAProviders
 * Handles incoming channel DM message
 */
func ConsumeFromOTAProviders(consumer chan dc.DeviceMessage, publisher chan dc.DeviceMessage) error {
	/*
	 * Create a Go Routine for the Providers Channel to
	 */
	go func(consumeChan chan dc.DeviceMessage, publishChan chan dc.DeviceMessage) {
		level.Debug(dvService.logger).Log("event", "ConsumeFromOTAProviders(gofunc) called")
		for msg := range consumeChan { // read until closed
			err := dvService.ApplyOTAEvent(&msg)
			if err != nil {
				level.Error(dvService.logger).Log("method", "ConsumeFromOTAProviders(gofunc)", "error", err.Error())
			}

			if nil != publishChan {
				publishChan <- msg
			}

			level.Debug(dvService.logger).Log("method", "ConsumeFromOTAProviders(gofunc)", "consume queue depth", len(consumeChan), "publish queue depth", len(publishChan))
		}
		level.Debug(dvService.logger).Log("method", "ConsumeFromOTAProviders()", "event", "Completed")
	}(consumer, publisher)

	return nil
}

/**
 * ConsumeFromScheduler
 * Handles incoming channel DM message
 */
func ConsumeFromScheduler(consumer chan dc.DeviceMessage, publisher chan dc.DeviceMessage) error {
	/*
	 * Create a Go Routine for the Providers Channel to
	 */
	go func(dmChan chan dc.DeviceMessage, otaChan chan dc.DeviceMessage) {
		level.Debug(dvService.logger).Log("event", "ConsumeFromScheduler(gofunc) called")
		for msg := range dmChan { // read until closed
			err := dvService.ApplySchedulerEvent(&msg)
			if err != nil {
				level.Error(dvService.logger).Log("method", "ConsumeFromScheduler(gofunc)", "error", err.Error())
			}
			if nil != otaChan {
				otaChan <- msg
			}

			level.Debug(dvService.logger).Log("method", "ConsumeFromScheduler(gofunc)", "dm queue depth", len(dmChan), "ota queue depth", len(otaChan))
		}
		level.Debug(dvService.logger).Log("method", "ConsumeFromScheduler()", "event", "Completed")
	}(consumer, publisher)

	return nil
}

func ChannelsForScheduler() (chan dc.DeviceMessage, chan dc.DeviceMessage, error) {
	var err error
	if toScheduler == nil {
		toScheduler = make(chan dc.DeviceMessage, 256) // averages 120 on startup
	}
	if fromScheduler == nil {
		fromScheduler = make(chan dc.DeviceMessage, 256) // averages 120 on startup
		err = ConsumeFromScheduler(fromScheduler, toOTAProvider)
	}

	if nil == toScheduler || nil == fromScheduler {
		err = errors.New("create scheduler channels failed")
		level.Error(dvService.logger).Log("error", err.Error())
		return nil, nil, err
	}

	return toScheduler, fromScheduler, err
}
func ChannelsForOTAProviders() (chan dc.DeviceMessage, chan dc.DeviceMessage, error) {
	var err error
	if toOTAProvider == nil {
		toOTAProvider = make(chan dc.DeviceMessage, 256) // averages 120 on startup
	}
	if fromOTAProvider == nil {
		fromOTAProvider = make(chan dc.DeviceMessage, 256) // averages 120 on startup
		err = ConsumeFromOTAProviders(fromOTAProvider, toScheduler)
	}

	if nil == toOTAProvider || nil == fromOTAProvider {
		err = errors.New("create ota channels failed")
		level.Error(dvService.logger).Log("error", err.Error())
		return nil, nil, err
	}

	return toOTAProvider, fromOTAProvider, err
}
func ChannelsForDMProviders() (chan dc.DeviceMessage, chan dc.DeviceMessage, error) {
	var err error
	if toProvider == nil {
		toProvider = make(chan dc.DeviceMessage, 256)
	}
	if fromProvider == nil {
		fromProvider = make(chan dc.DeviceMessage, 256)
		err = ConsumeFromDMProviders(fromProvider) // start receiver
	}

	if nil == toProvider || nil == fromProvider {
		err = errors.New("create provider channels failed")
		level.Error(dvService.logger).Log("error", err.Error())
		return nil, nil, err
	}

	return toProvider, fromProvider, err
}
func ChannelsForCore() (chan dc.DeviceMessage, chan dc.DeviceMessage, error) {
	var err error
	if toCore == nil {
		toCore = make(chan dc.DeviceMessage, 256) // averages 120 on startup
	}

	if fromCore == nil {
		fromCore = make(chan dc.DeviceMessage, 256) // averages 120 on startup
		err = ConsumeFromCore(fromCore)
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
