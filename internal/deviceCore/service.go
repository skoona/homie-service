package deviceCore

/*
  	deviceCore/service.go:

	Build Device Network Inventory
	Provide Interaction to CLI
	Provide Interaction to HTTP
	Drive Scheduler Feature
*/

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/google/uuid"
	cc "github.com/skoona/homie-service/internal/utils"
)

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
		RestoreNetworkFromDB(networkName string) Network
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
func NewDeviceMessage(topic string, payload []byte, idCounter uint16, retained bool, qos byte, plog log.Logger) (DeviceMessage, error) {
	return buildDeviceMessage(topic, payload, idCounter, retained, qos, plog)
}

/**
 * NewEventMessage()
 *
 *  Create a New DeviceMessage and initializes it.
 *  Called outside of the service
 */
func NewQueueMessage(msg QueueMessage, plog log.Logger) (DeviceMessage, error) {
	return NewDeviceMessage(msg.Topic(), msg.Payload(), msg.MessageID(), msg.Retained(), msg.Qos(), plog)
}

//NewID create a new entity ID
func NewEID() EID {
	uuid := uuid.New()
	return EID(uuid.String())
}

/*
decode device into topics
- -/D/A/P/P
- -/D/A/P
- -/D/A
- -/D/N/P/A
- -/D/N/P
- -/D/N/A
*/
func TopicsFromDevice(dv Device) []string {
	bundle := []string{}
	var ele string

	// unpacn device attrs
	for n, v := range dv.Attrs {  // x/d/a
		if len(v.Props) > 0 {
			for nn, vv := range v.Props { // x/d/a/p
				if len(vv.Props) > 0 {
					for nnn, _ := range vv.Props { // x/d/a/p/p
						ele = fmt.Sprintf("%s/%s/%s/%s/%s", dv.Parent, dv.Name, n, nn, nnn)
						bundle = append(bundle, ele)
					}
				} else {
					ele = fmt.Sprintf("%s/%s/%s/%s", dv.Parent, dv.Name, n, nn)
					bundle = append(bundle, ele)
				}
			}
		} else {
			ele = fmt.Sprintf("%s/%s/%s", dv.Parent, dv.Name, n)
			bundle = append(bundle, ele)
		}
	}
	// unpacn device nodes
	for n, v := range dv.Nodes {  // x/d/n
		for nn, vv := range v.Props { // x/d/n/p
			if len(vv.Attrs) > 0 {
				for nnn, _ := range vv.Attrs { // x/d/n/p/a
					ele = fmt.Sprintf("%s/%s/%s/%s/%s", dv.Parent, dv.Name, n, nn, nnn)
					bundle = append(bundle, ele)
				}
			} else {
				ele = fmt.Sprintf("%s/%s/%s/%s", dv.Parent, dv.Name, n, nn)
				bundle = append(bundle, ele)
			}
		}
		for nn, _ := range v.Attrs { // x/d/n/a
			ele = fmt.Sprintf("%s/%s/%s/%s", dv.Parent, dv.Name, n, nn)
			bundle = append(bundle, ele)
		}
	}

	return bundle
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
