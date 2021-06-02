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
	CoreService interface {
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
		ScheduleByDeviceEID(deviceID EID) Schedule

		AllBroadcasts() []Broadcast
		AddBroadcast(broadcast Broadcast)
		RemoveBroadcastByEID(broadcastEID EID)
		BroadcastByEID(broadcastID EID) (Broadcast, error)
	}

	/*
	 * Interactions with device Scheduling Service */
	SchedulerProvider interface {
		ActivateStreamProvider()
		ApplySiteNetworks(sn *SiteNetworks)
		BuildFirmwareCatalog() []Firmware
		Firmwares() []Firmware
		GetFirmware(id EID) (Firmware, error)
		CreateFirmware(path string) error
		DeleteFirmware(id EID) error
		BuildScheduleCatalog() map[EID]Schedule
		Schedules() []Schedule
		FindSchedulesByDeviceEID(deviceID EID) []Schedule
		CreateSchedule(networkName string, deviceID EID, transport OTATransport, firmware *Firmware) (EID, error)
		DeleteSchedule(scheduleID EID) error
	}

	/*
	 * Interactions with DeviceSource */
	DeviceEventProvider interface {
		ActivateStreamProvider()
		ApplyDeviceEvent(dm DeviceMessage)
		HandleCoreEvent(dv Device) error
		PublishToStreamProvider(dv Device)
		ConsumeDeviceStream(dm DeviceMessage) error
	}

	// Device Source Storage Repository
	Repository interface {
		Store(d DeviceMessage) error
		Remove(d DeviceMessage) error
		LoadSchedules() map[EID]Schedule
		StoreSchedule(d Schedule) error
		RemoveSchedule(d Schedule) error
	}

	// Service Implementation
	coreService struct {
		cfg    cc.Config
		dsp    DeviceEventProvider
		scp    SchedulerProvider
		logger log.Logger
	}

	//EID entity EID
	EID string
)

var (
	logger log.Logger
)

/**
 * NewCoreService()
 *
 *  Create a New NewCoreService and initializes it.
 */
func NewCoreService(dfg cc.Config, sp DeviceEventProvider, sscp SchedulerProvider) CoreService {
	em = &coreService{
		cfg:    dfg,
		dsp:    sp,
		scp:    sscp,
		logger: log.With(dfg.Logger, "pkg", "deviceCore", "service", "coreService"),
	}
	return em
}

/**
 * NewDeviceMessage()
 *
 *  Create a New DeviceMessage and initializes it.
 *  Called outside of the service
 */
func NewDeviceMessage(topic string, payload []byte, idCounter uint16, retained bool, qos byte) (DeviceMessage, error) {
	return buildDeviceMessage(topic, payload, idCounter, retained, qos)
}

/**
 * NewEventMessage()
 *
 *  Create a New DeviceMessage and initializes it.
 *  Called outside of the service
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
func Start(dfg cc.Config, sp DeviceEventProvider, sscp SchedulerProvider, discoveredNetworks []string) (CoreService, *SiteNetworks) {

	svc := NewCoreService(dfg, sp, sscp)

	logger = em.logger

	level.Debug(em.logger).Log("event", "Calling Start()")

	// Initialze networks
	sites := NewSiteNetworks("Skoona Consulting",
		"Homie Monitor (GOLANG)",
		discoveredNetworks,
		[]Firmware{},
		map[EID]Schedule{})

	level.Debug(em.logger).Log("event", "Start() completed")

	/* start the message flow from stream providers */
	sp.ActivateStreamProvider()

	if sscp != nil {
		sscp.ApplySiteNetworks(sites)
		sscp.BuildFirmwareCatalog()
		sscp.BuildScheduleCatalog()
		sscp.ActivateStreamProvider()
	}

	return svc, sites
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(em.logger).Log("event", "Calling Stop()")

	level.Debug(em.logger).Log("event", "Stop() completed")
}
