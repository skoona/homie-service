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
	sch "github.com/skoona/homie-service/internal/deviceScheduler"
	cc "github.com/skoona/homie-service/internal/utils"
)

// Message Interface for MQTT and Demo
type (
	//Incoming Messages
	Service interface {
		ApplyStreamEvent(dm dc.DeviceMessage) error
		ApplyIncomingCoreEvent(dm dc.DeviceMessage) error
	}

	// Core Service Implementation
	deviceSource struct {
		cfg        cc.Config
		repository Repository
		dStream StreamProvider
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


func (s *deviceSource) ApplyStreamEvent(dm dc.DeviceMessage) error {
	logger := log.With(s.logger, "method", "ApplyStreamEvent")

	err := s.repository.Store(dm)
	if err != nil {
		level.Error(logger).Log("error", err)
		return err
	}

	// also sent it to core
	// s.coreSvc.FromDeviceSource(dm)
	if err != nil {
		level.Error(logger).Log("error", err)
		return err
	}

	level.Debug(logger).Log("DeviceID ", dm.DeviceID)

	return err
}


// handle incoming core events
func (s *deviceSource) ApplyIncomingCoreEvent(dm dc.DeviceMessage) error {
	var err error
	plog := log.With(s.logger, "method", "ApplyIncomingCoreEvent")

	// can only be a delete request
	err = s.repository.Store(dm)
	if err != nil {
		level.Error(plog).Log("error", err)
		return err
	}
	// explode the dm into stream messages
	// with no value to cause them to be deleted

	level.Debug(plog).Log("DeviceID ", dm.DeviceID)

	return err
}


var (
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
			err := dvService.ApplyIncomingCoreEvent(msg)
			if err != nil {
				level.Error(dmHandler.logger).Log("method", "ConsumeFromCore(gofunc)", "error", err.Error())
			}

			level.Debug(dmHandler.logger).Log("method", "ConsumeFromCore(gofunc)", "queue depth", len(dmChan))
		}
		level.Debug(dmHandler.logger).Log("method", "ConsumeFromCore()", "event", "Completed")
	}(consumer)

	return nil
}

func ChannelsForCore() (chan dc.DeviceMessage, error) {
	var err error
	// FromDeviceSource(dm DeviceMessage) error

	if fromCore == nil {
		fromCore, err = dvService.coreSvc.GetCoreResponseChannel() // averages 120 on startup
		if err == nil {
			err = ConsumeFromCore(fromCore)
		}
	}

	if nil == fromCore {
		err = errors.New("create core channels failed")
		level.Error(logger).Log("error", err.Error())
		return nil, err
	}

	return fromCore, err
}

/*
 * Start()
 *
 * Initialize this service
 */
func Start(dfg cc.Config, repo Repository, dProvider StreamProvider) (DeviceMessageHandler, error) {
	var err error
	logger = log.With(dfg.Logger, "pkg", "deviceSource")
	level.Debug(logger).Log("event", "Calling Start()")

	NewDeviceSourceService(dfg, repo, dProvider, logger)
	NewDeviceMessageHandler(dfg, dProvider, logger)

	level.Debug(logger).Log("event", "Start() Completed")

	return dmHandler, err
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
