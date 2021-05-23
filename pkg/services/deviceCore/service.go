package deviceCore

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/google/uuid"
	cc "github.com/skoona/homie-service/pkg/utils"
)

/*
  	deviceCore/service.go:

	Build Device Network Inventory
	Provide Interaction to CLI
	Provide Interaction to HTTP
	Consider Scheduler Feature
*/

type (
	Service interface {
		EventHandler
	}

	// Service Implementation
	coreService struct {
		cfg    cc.Config
		logger log.Logger
	}

	//EID entity EID
	EID string
)

/**
 * NewCoreService()
 *
 *  Create a New NewCoreService and initializes it.
 */
func NewCoreService(dfg cc.Config) Service {
	em = coreService{
		cfg:    dfg,
		logger: log.With(dfg.Logger, "pkg", "deviceCore", "service", "coreService"),
	}
	return &em
}

/**
 * NewDeviceMessage()
 *
 *  Create a New DeviceMessage and initializes it.
 */
func NewDeviceMessage(topic string, payload []byte, idCounter uint16, retained bool, qos byte) (DeviceMessage, error) {
	return buildDeviceMessage(topic, payload, idCounter, retained, qos)
}

/**
 * NewEventMessage()
 *
 *  Create a New DeviceMessage and initializes it.
 */
func NewQueueMessage(msg QueueMessage) (DeviceMessage, error) {
	return NewDeviceMessage(msg.Topic(), msg.Payload(), msg.MessageID(), msg.Retained(), msg.Qos())
}

//NewID create a new entity ID
func NewEID() EID {
	uuid := uuid.New()
	return EID(uuid.String())
}

/*
 * Start()
 *
 * Initialize this service
 */
func Start(dfg cc.Config) (*Service, error) {
	var err error

	svc := NewCoreService(dfg)

	level.Debug(em.logger).Log("event", "Calling Start()")

	// if err != nil {
	// 	level.Error(logger).Log("event", "Channels offline", "error", err.Error())
	// 	panic(err.Error())
	// }

	level.Debug(em.logger).Log("event", "Start() completed")

	return &svc, err
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(em.logger).Log("event", "Calling Stop()")

	level.Debug(em.logger).Log("event", "Stop() completed")
}
