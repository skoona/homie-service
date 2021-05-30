package deviceCore

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/google/uuid"
	cc "github.com/skoona/homie-service/internal/utils"
)

/*
  	deviceCore/service.go:

	Build Device Network Inventory
	Provide Interaction to CLI
	Provide Interaction to HTTP
	Consider Scheduler Feature
*/

type (
	/*
	 * Interactions with UI */
	Service interface {
		AllNetworks() SiteNetworks
		NetworkByName(networkName string) Network
		DeviceByNameFromNetwork(deviceName, networkName string) (Device, error)
		DeviceByEIDFromNetwork(deviceEID EID, networkName string) (Device, error)
		RemoveDeviceByEIDFromNetwork(deviceEID EID, networkName string) error

		AllFirmwares() []Firmware
		AddFirmware(firmware Firmware)
		RemoveFirmwareByEID(firmwareEID EID)
		FirmwareByEID(firmwareEID EID) (Firmware, error)
		FirmwareByName(firmwareName string) (Firmware, error)

		AllSchedules() []Schedule
		AddSchedule(schedule Schedule)
		RemoveSchedule(scheduleID EID)
		ScheduleByEID(scheduleID EID) Schedule
		ScheduleByDeviceName(deviceName string) Schedule

		AllBroadcasts() []Broadcast
		AddBroadcast(broadcast Broadcast)
		RemoveBroadcastByEID(broadcastEID EID)
		BroadcastByEID(broadcastID EID) (Broadcast, error)
	}

	/*
	 * Interactions with DeviceSource */
	DeviceSourceInteractor interface {
		CreateDemoDeviceMessage(topic string, payload []byte, idCounter uint16, retained bool, qos byte) (DeviceMessage, error)
		CreateQueueDeviceMessage(qmsg QueueMessage) (DeviceMessage, error)
		GetCoreRequestChannel() (chan DeviceMessage, error)
		GetCoreResponseChannel() (chan DeviceMessage, error)
		FromDeviceSource(dm DeviceMessage) error
	}

	// Service Implementation
	coreService struct {
		cfg    cc.Config
		logger log.Logger
	}

	coreDeviceSourceService struct {
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
	em = &coreService{
		cfg:    dfg,
		logger: log.With(dfg.Logger, "pkg", "deviceCore", "service", "coreService"),
	}
	return em
}

/**
 * NewCoreDeviceSourceService()
 *
 *  Create a New NewCoreService and initializes it.
 */
func NewCoreDeviceSourceService(dfg cc.Config) DeviceSourceInteractor {
	cdss = &coreDeviceSourceService{
		cfg:    dfg,
		logger: log.With(dfg.Logger, "pkg", "deviceCore", "service", "coreDeviceSourceService"),
	}
	return cdss
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
func Start(dfg cc.Config, discoveredNetworks []string) (DeviceSourceInteractor, error) {
	var err error

	// svc := NewCoreService(dfg)
	svc := NewCoreDeviceSourceService(dfg)

	level.Debug(cdss.logger).Log("event", "Calling Start()")

	// Initialze networks
	NewSiteNetworks("Skoona Consulting",
		"Homie Monitor (GOLANG)",
		discoveredNetworks,
		[]Firmware{},
		map[string]Schedule{})

	level.Debug(cdss.logger).Log("event", "Start() completed")

	return svc, err
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(cdss.logger).Log("event", "Calling Stop()")

	level.Debug(cdss.logger).Log("event", "Stop() completed")
}
