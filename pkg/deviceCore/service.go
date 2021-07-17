package deviceCore

/*
  	deviceCore/service.go:

	Build Device Network Inventory
	Provide Interaction to CLI
	Provide Interaction to HTTP
	Drive Scheduler Feature
*/

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	cc "github.com/skoona/homie-service/pkg/utils"
)

type (
	/*
	 * Interactions with UI */
	CoreService interface {
		AllNetworks() SiteNetworks
		PrivateSiteNetworks() *SiteNetworks
		NetworkDevices(networkName string) (map[string]Device, error)
		NetworkByName(networkName string) Network
		DeviceByName(deviceName, networkName string) (Device, error)
		DeviceByID(deviceID string, networkName string) (Device, error)
		RemoveDeviceByID(deviceID string, networkName string) error

		PublishNetworkMessage(dm DeviceMessage)

		AllSchedules() []Schedule
		CreateSchedule(networkName string, deviceID string, transport OTATransport, firmwareID EID) (string, error)
		RemoveSchedule(scheduleID string) error
		ScheduleByID(scheduleID string) Schedule
		ScheduleByDeviceID(deviceID string) *Schedule

		AllFirmwares() []Firmware
		CreateFirmware(srcFile, dstFile string) (EID, error)
		RemoveFirmwareByID(firmwareID EID)
		FirmwareByID(firmwareID EID) (Firmware, error)
		FirmwareByName(firmwareName string) (Firmware, error)

		AllBroadcasts() []Broadcast
		RemoveBroadcastByID(broadcastID string)
		BroadcastByID(broadcastID string) (Broadcast, error)
	}

	/*
	 * Interactions with device Scheduling Service */
	SchedulerProvider interface {
		ActivateStreamProvider()
		ApplySiteNetworks(sn *SiteNetworks)
		Schedules() []Schedule
		BuildScheduleCatalog() map[string]Schedule
		FindScheduleByDeviceID(deviceID string) *Schedule
		CreateSchedule(networkName string, deviceID string, transport OTATransport, firmwareID EID) (string, error)
		DeleteSchedule(scheduleID string) error
		BuildFirmwareCatalog() []Firmware
		Firmwares() []Firmware
		GetFirmware(id EID) (Firmware, error)
		CreateFirmware(srcFile, dstFile string) (EID, error)
		DeleteFirmware(id EID) error
	}

	/*
	 * Interactions with DeviceSource */
	DeviceEventProvider interface {
		ActivateStreamProvider()
		ApplyDeviceEvent(dm DeviceMessage)
		PublishToStreamProvider(dm DeviceMessage)
		ConsumeDeviceStream(dm DeviceMessage) error
	}

	// Device Source Storage Repository
	Repository interface {
		Store(d DeviceMessage) error
		Remove(d DeviceMessage) error
		RestoreNetworkFromDB(networkName string) Network
		LoadSchedules() map[string]Schedule
		StoreSchedule(d Schedule) error
		RemoveSchedule(scheduleID string) error
		LoadBroadcasts(networkName string) []Broadcast
		RemoveAllBroadcasts(networkName string) error
	}

	// Service Implementation
	coreService struct {
		cfg    cc.Config
		dsp    DeviceEventProvider
		scp    SchedulerProvider
		repo    Repository
		logger log.Logger
	}
)

var (
	logger log.Logger
)

/**
 * NewCoreService()
 *
 *  Create a New NewCoreService and initializes it.
 */
func NewCoreService(dfg cc.Config, sp DeviceEventProvider, sscp SchedulerProvider, repo Repository) CoreService {
	em = &coreService{
		cfg:    dfg,
		dsp:    sp,
		scp:    sscp,
		repo:	repo,
		logger: log.With(dfg.Logger, "pkg", "deviceCore", "service", "coreService"),
	}
	return em
}


/*
 * Start()
 *
 * Initialize this service
 */
func Start(dfg cc.Config, sp DeviceEventProvider, sscp SchedulerProvider, repo Repository, discoveredNetworks []string) (CoreService, *SiteNetworks) {

	svc := NewCoreService(dfg, sp, sscp, repo)

	logger = em.logger

	level.Debug(em.logger).Log("event", "Calling Start()")

	// Initialze networks
	sites := NewSiteNetworks("Skoona Consulting",
		"Homie Monitor (GOLANG)",
		discoveredNetworks,
		sscp.BuildFirmwareCatalog(),
		sscp.BuildScheduleCatalog()	)

	for _, netStr := range discoveredNetworks {
		sites.DeviceNetworks[netStr] = repo.RestoreNetworkFromDB(netStr)
		sites.Broadcasts = repo.LoadBroadcasts(netStr)
		level.Info(em.logger).Log("event", "restore networks", "network", netStr)
	}

	/* start the message flow from stream providers */
	if sscp != nil {
		sscp.ApplySiteNetworks(sites)
		sscp.ActivateStreamProvider()
	}
	sp.ActivateStreamProvider()

	level.Debug(em.logger).Log("event", "Start() completed")
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
