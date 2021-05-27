package deviceSource

/*
  deviceSource/usecase.go:

  DeviceSource Service Implementation
*/

import (
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	sch "github.com/skoona/homie-service/internal/deviceScheduler"
	cc "github.com/skoona/homie-service/internal/utils"
)

type (
	// Core Service Implementation
	deviceSource struct {
		repository Repository
		coreSvc    dc.DeviceSourceInteractor
		cfg        cc.Config
		logger     log.Logger
	}
	deviceMessageHandler struct {
		cfg    cc.Config
		sched  sch.SchedulerService
		logger log.Logger
	}
)

// Retained DeviceSource Service, once created
var (
	dmHandler *deviceMessageHandler
	dvService *deviceSource
	logger    log.Logger
)

func (s *deviceSource) ApplyDMEvent(dm dc.DeviceMessage) error {
	logger := log.With(s.logger, "method", "ApplyDMEvent")

	err := s.repository.Store(dm)
	if err != nil {
		level.Error(logger).Log("error", err)
		return err
	}

	// also sent it to core
	s.coreSvc.FromDeviceSource(dm)
	if err != nil {
		level.Error(logger).Log("error", err)
		return err
	}

	level.Debug(logger).Log("DeviceID ", dm.DeviceID)

	return err
}

func (s *deviceSource) ApplyOTAEvent(dm dc.DeviceMessage) error {
	var err error
	logg := log.With(s.logger, "method", "ApplyOTAEvent")

	level.Debug(logg).Log("DeviceID ", dm.DeviceID)

	return err
}

// handle incoming core events
func (s *deviceSource) ApplyCoreEvent(dm dc.DeviceMessage) error {
	var err error
	logg := log.With(s.logger, "method", "ApplyCoreEvent")

	err = s.repository.Store(dm)
	if err != nil {
		level.Error(logg).Log("error", err)
		return err
	}

	level.Debug(logg).Log("DeviceID ", dm.DeviceID)

	return err
}

func (s *deviceSource) ApplySchedulerEvent(dm dc.DeviceMessage) error {
	var err error
	logg := log.With(s.logger, "method", "ApplySchedulerEvent")

	level.Debug(logg).Log("DeviceID ", dm.DeviceID)

	return err
}

/*
 * DeviceSourceInteractor
 */
/**
 * ConsumeFromProviders
 * Handles incoming channel DM message
 */
func consumeFromProviders(consumer chan dc.DeviceMessage) error {
	/*
	 * Create a Go Routine for the Providers Channel to
	 */
	go func(dmChan chan dc.DeviceMessage) {
		level.Debug(logger).Log("event", "ConsumeFromDMProviders(gofunc) called")
		for msg := range dmChan { // read until closed
			err := dvService.ApplyDMEvent(msg)
			if err != nil {
				level.Error(logger).Log("method", "ConsumeFromDMProviders(gofunc)", "error", err.Error())
			}
			level.Debug(logger).Log("method", "ConsumeFromDMProviders(gofunc)", "queue depth", len(dmChan))
		}
		level.Debug(logger).Log("method", "ConsumeFromDMProviders(gofunc)", "event", "Completed")
	}(consumer)

	return nil
}

/**
 * ConsumeFromOTAProviders
 * routes from OTA to Scheduler
 */
func consumeFromOTAProviders(consumer chan dc.DeviceMessage, publisher chan dc.DeviceMessage) error {
	/*
	 * Create a Go Routine for the Providers Channel to
	 */
	go func(consumeChan chan dc.DeviceMessage, publishChan chan dc.DeviceMessage) {
		level.Debug(logger).Log("event", "ConsumeFromOTAProviders(gofunc) called")
		for msg := range consumeChan { // read until closed

			err := dvService.ApplyOTAEvent(msg)
			if err != nil {
				level.Error(logger).Log("method", "ConsumeFromOTAProviders(gofunc)", "error", err.Error())
			}

			if nil != publishChan {
				publishChan <- msg
			}

			level.Debug(logger).Log("method", "ConsumeFromOTAProviders(gofunc)", "consume queue depth", len(consumeChan), "publish queue depth", len(publishChan))
		}
		level.Debug(logger).Log("method", "ConsumeFromOTAProviders()", "event", "Completed")
	}(consumer, publisher)

	return nil
}

/**
 * ConsumeFromScheduler
 * routes from Scheduler to dm provider
 */
func consumeFromScheduler(consumer chan dc.DeviceMessage, publisher chan dc.DeviceMessage) error {
	/*
	 * Create a Go Routine for the Providers Channel to
	 */
	go func(dmChan chan dc.DeviceMessage, otaChan chan dc.DeviceMessage) {
		level.Debug(logger).Log("event", "ConsumeFromScheduler(gofunc) called")
		for msg := range dmChan { // read until closed
			err := dvService.ApplySchedulerEvent(msg)
			if err != nil {
				level.Error(logger).Log("method", "ConsumeFromScheduler(gofunc)", "error", err.Error())
			}
			if nil != otaChan {
				otaChan <- msg
			}

			level.Debug(logger).Log("method", "ConsumeFromScheduler(gofunc)", "dm queue depth", len(dmChan), "ota queue depth", len(otaChan))
		}
		level.Debug(logger).Log("method", "ConsumeFromScheduler()", "event", "Completed")
	}(consumer, publisher)

	return nil
}

// Inject Scheduler Service after Start or New...
func (dmHandler *deviceMessageHandler) BuildFirmwareCatalog(svc sch.SchedulerService) []dc.Firmware {
	dmHandler.sched = svc
	return svc.BuildFirmwareCatalog()
}

func (dmHandler *deviceMessageHandler) CreateDemoDeviceMessage(topic string, payload []byte, idCounter uint16, retained bool, qos byte) (dc.DeviceMessage, error) {
	level.Debug(dmHandler.logger).Log("method", "CreateDemoDeviceMessage() called")
	return dc.NewDeviceMessage(topic, payload, idCounter, retained, qos)
}

func (dmHandler *deviceMessageHandler) CreateQueueDeviceMessage(qmsg dc.QueueMessage) (dc.DeviceMessage, error) {
	level.Debug(dmHandler.logger).Log("method", "CreateQueueDeviceMessage() called")
	return dc.NewQueueMessage(qmsg)
}

func (dmHandler *deviceMessageHandler) GetProviderRequestChannel() (chan dc.DeviceMessage, error) {
	level.Debug(dmHandler.logger).Log("method", "GetProviderRequestChannel() called")
	var err error
	if fromProvider == nil {
		fromProvider = make(chan dc.DeviceMessage, 256) // averages 120 on startup
		if fromProvider != nil {
			err = consumeFromProviders(fromProvider) // start receiver
		}
	}

	if nil == fromProvider {
		err = errors.New("create publishing channel failed")
		level.Error(dmHandler.logger).Log("error", err.Error())
		return nil, err
	}

	return fromProvider, err
}

func (dmHandler *deviceMessageHandler) GetProviderResponseChannel() (chan dc.DeviceMessage, error) {
	level.Debug(dmHandler.logger).Log("method", "GetProviderResponseChannel() called")
	var err error
	if toProvider == nil {
		toProvider = make(chan dc.DeviceMessage, 256) // averages 120 on startup
	}

	if nil == toProvider {
		err = errors.New("create subscribing channel failed")
		level.Error(dmHandler.logger).Log("error", err.Error())
		return nil, err
	}
	return toProvider, err
}

func (dmHandler *deviceMessageHandler) GetOTARequestChannel() (chan dc.DeviceMessage, error) {
	level.Debug(dmHandler.logger).Log("method", "GetOTARequestChannel() called")
	var err error
	if fromOTAProvider == nil {
		fromOTAProvider = make(chan dc.DeviceMessage, 256) // averages 120 on startup
		if fromOTAProvider != nil {
			err = consumeFromOTAProviders(fromOTAProvider, toScheduler) // Route Messages to Scheduler
		}
	}

	if nil == fromOTAProvider {
		err = errors.New("create publishing channel failed")
		level.Error(dmHandler.logger).Log("error", err.Error())
		return nil, err
	}

	return fromOTAProvider, err
}

func (dmHandler *deviceMessageHandler) GetOTAResponseChannel() (chan dc.DeviceMessage, error) {
	level.Debug(dmHandler.logger).Log("method", "GetOTAResponseChannel() called")
	var err error
	if toOTAProvider == nil {
		toOTAProvider = make(chan dc.DeviceMessage, 120) // averages 120 on startup
	}

	if nil == toOTAProvider {
		err = errors.New("create subscribing channel failed")
		level.Error(dmHandler.logger).Log("error", err.Error())
		return nil, err
	}
	return toOTAProvider, err
}

func (dmHandler *deviceMessageHandler) GetSchedulerRequestChannel() (chan dc.DeviceMessage, error) {
	level.Debug(dmHandler.logger).Log("method", "GetSchedulerRequestChannel() called")
	var err error
	if fromScheduler == nil {
		fromScheduler = make(chan dc.DeviceMessage, 120) // averages 120 on startup
		if fromScheduler != nil {
			err = consumeFromScheduler(fromScheduler, toOTAProvider) // Route to OTA command
		}
	}

	if nil == fromScheduler {
		err = errors.New("create publishing channel failed")
		level.Error(dmHandler.logger).Log("error", err.Error())
		return nil, err
	}

	return fromScheduler, err
}

func (dmHandler *deviceMessageHandler) GetSchedulerResponseChannel() (chan dc.DeviceMessage, error) {
	level.Debug(dmHandler.logger).Log("method", "GetSchedulerResponseChannel() called")
	var err error
	if toScheduler == nil {
		toScheduler = make(chan dc.DeviceMessage, 120) // averages 120 on startup
	}

	if nil == toScheduler {
		err = errors.New("create subscribing channel failed")
		level.Error(dmHandler.logger).Log("error", err.Error())
		return nil, err
	}
	return toScheduler, err
}

func (dmHandler *deviceMessageHandler) FromMQTTProvider(qmsg dc.QueueMessage) error {
	level.Debug(dmHandler.logger).Log("method", "FromMQTTProvider() called", "msgID", qmsg.MessageID())
	dm, err := dmHandler.CreateQueueDeviceMessage(qmsg)
	if err != nil {
		level.Error(dmHandler.logger).Log("error", err.Error())
	} else {
		ch, err := dmHandler.GetProviderRequestChannel()
		if err == nil {
			ch <- dm // send it to channel
		}
	}
	return err
}
func (dmHandler *deviceMessageHandler) FromDemoProvider(topic string, payload []byte, idCounter uint16, retained bool, qos byte) error {
	level.Debug(dmHandler.logger).Log("method", "FromDemoProvider() called", "msgID", idCounter)
	dm, err := dmHandler.CreateDemoDeviceMessage(topic, payload, idCounter, retained, qos)
	if err != nil {
		level.Error(dmHandler.logger).Log("error", err.Error())
	} else {
		ch, err := dmHandler.GetProviderRequestChannel()
		if err == nil {
			ch <- dm // send it to channel
		}
	}
	return err
}
func (dmHandler *deviceMessageHandler) FromOTAProvider(qmsg dc.QueueMessage) error {
	level.Debug(dmHandler.logger).Log("method", "FromOTAProvider() called", "msgID", qmsg.MessageID())
	dm, err := dmHandler.CreateQueueDeviceMessage(qmsg)
	if err != nil {
		level.Error(dmHandler.logger).Log("error", err.Error())
	} else {
		ch, err := dmHandler.GetOTARequestChannel()
		if err == nil {
			ch <- dm // send it to channel
		}
	}
	return err
}
func (dmHandler *deviceMessageHandler) FromScheduler(qmsg dc.QueueMessage) error {
	level.Debug(dmHandler.logger).Log("method", "FromScheduler() called", "msgID", qmsg.MessageID())
	dm, err := dmHandler.CreateQueueDeviceMessage(qmsg)
	if err != nil {
		level.Error(dmHandler.logger).Log("error", err.Error())
	} else {
		ch, err := dmHandler.GetSchedulerRequestChannel()
		if err == nil {
			ch <- dm // send it to channel
		}
	}
	return err
}

// Receive/send DM from Channel
// Receive/send OTA from Channel
// Receive/Send to Scheduler
// Receive/Send to Core

// Route OTA to/from Scheduler
// Route DM to Repository and Core Service
