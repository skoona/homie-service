package deviceCore_test

import (
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	dp "github.com/skoona/homie-service/internal/demoProvider"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	sch "github.com/skoona/homie-service/internal/deviceScheduler"
	dss "github.com/skoona/homie-service/internal/deviceSource"
	dds "github.com/skoona/homie-service/internal/deviceStorage"
	mq "github.com/skoona/homie-service/internal/mqttProvider"
	cc "github.com/skoona/homie-service/internal/utils"
	"os"
)

var (
	oldArgs		[]string
	networks     []string
	logger       log.Logger
	siteNetworks *dc.SiteNetworks
	dep          dc.DeviceEventProvider
	sched        dc.SchedulerProvider
	coreSvc      dc.CoreService
	repo         dc.Repository
	otap         sch.OTAInteractor
	dsp          dss.StreamProvider
	err 		 error
	cfg			cc.Config
)


func ShutdownDemo() {
	dp.Stop()
	dss.Stop()
	dds.Stop()
	dc.Stop()
}

func ShutdownLive() {
	mq.Stop()
	sch.Stop()
	dss.Stop()
	dds.Stop()
	dc.Stop()
}

func RunDemo(cfg cc.Config) (dc.CoreService, error) {
	otap, dsp, networks, err = dp.Initialize(cfg)                     // message stream
	if err != nil {
		return nil, err
	}
	repo, err = dds.Start(cfg)                                        // message db
	if err != nil {
		return nil, err
	}
	dep, _ = dss.Start(cfg, repo, dsp)                                // message aggregation
	sched = sch.Start(cfg, otap, repo)                                // ota scheduler
	coreSvc, siteNetworks = dc.Start(cfg, dep, sched, repo, networks) // network logic
	err = dp.Start()                                                  // activate message stream
	if err != nil {
		return nil, err
	}
	return coreSvc, err
}

func StartUp() (dc.CoreService, cc.Config, error) {
	cfg, err = cc.BuildRuntimeConfig("Homie-Service Test")
	if err != nil {
		os.Exit(1)
	}
	logger = log.With(cfg.Logger, "ginkgo", "testing")

	// remove db file
	os.Remove(cfg.Dbc.DataStorage)
	level.Debug(logger).Log("removed DB", cfg.Dbc.DataStorage)

	// Run the App
	coreSvc, err = RunDemo(cfg)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		os.Exit(2)
	}

	return coreSvc, cfg, err
}

func Shutdown(cfg cc.Config) {
		ShutdownDemo()
}

var _ = Describe("Core Service", func() {
	BeforeSuite(func() {
		oldArgs = os.Args
		os.Args = []string{oldArgs[0], "--config", ""} // force clearing of prior value
		defer func() { os.Args = oldArgs }()
		os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "test-config")
		os.Args = []string{oldArgs[0], "--debug", "true", "--config", ""}

		coreSvc, cfg, err = StartUp()
		if err != nil {
			Fail(err.Error())
		}

	})

	AfterSuite(func(){
		Shutdown(cfg)
	})

	Context("Core Operations", func() {
		Context("Network Operations", func() {
			It("AllNetworks() returns the SiteNetwork", func() {
				var sn dc.SiteNetworks
				Expect(coreSvc.AllNetworks()).To(BeAssignableToTypeOf(sn))
			})
			It("AllNetworks() has known network", func() {
				sn := coreSvc.AllNetworks()
				Expect(len(sn.DeviceNetworks)).To(Equal(1))
				Expect(sn.DeviceNetworks["sknSensors"]).ToNot(BeNil())
			})
			It("AllNetworks() known network has Devices", func() {
				sn := coreSvc.AllNetworks()
				Expect(len(sn.DeviceNetworks["sknSensors"].Devices)).NotTo(Equal(0))
			})
		})
		Context("Broadcast Operations", func() {
			It("AllBroadcasts() returns the available broadcasts", func() {
				var bc []dc.Broadcast
				Expect(coreSvc.AllBroadcasts()).To(BeAssignableToTypeOf(bc))
			})
			It("AllBroadcasts() have known items", func() {
				bc := coreSvc.AllBroadcasts()
				Expect(len(bc)).ToNot(Equal(0))
				Expect(bc[0]).ToNot(BeNil())
			})
		})
		Context("Firmware Operations", func() {
			It("AllFirmwares() returns the current library", func() {
				var fw []dc.Firmware
				Expect(coreSvc.AllFirmwares()).To(BeAssignableToTypeOf(fw))
			})
			It("AllFirmwares() have known items", func() {
				var fw []dc.Firmware
				fws := coreSvc.AllFirmwares()
				Expect(len(fws)).ToNot(Equal(0))
				Expect(fw).To(BeAssignableToTypeOf(fw))
			})
		})
		Context("Schedule Operations", func() {
			It("AllSchedules() returns the available schedules", func() {
				var sc []dc.Schedule
				Expect(coreSvc.AllSchedules()).To(BeAssignableToTypeOf(sc))

				schedules := repo.LoadSchedules()
				count := len(schedules)

				out, _ := json.MarshalIndent(coreSvc.AllSchedules(), "", "  ")
				Expect(len(coreSvc.AllSchedules())).To(Equal(count), string(out))
			})
			It("CreateSchedule() returns a new schedules", func() {
				var fw dc.Firmware
				var dv dc.Device
				var sc dc.Schedule
				var scc dc.Schedule
				schedules := coreSvc.AllSchedules()
				count := len(schedules)

				dv, err = coreSvc.DeviceByName("GuestRoom", "sknSensors")
					//out, _ := json.MarshalIndent(dv, "", "  ")
					//fmt.Printf("\nDevice: %s\n", out)
				fw = coreSvc.AllFirmwares()[0]
					//out, _ = json.MarshalIndent(fw, "", "  ")
					//fmt.Printf("\nFirmware: %s\n", out)

				scID, err := coreSvc.CreateSchedule("sknSensors", dv.ID, dc.Base64Strict, fw.ID)
				Expect(err).To(BeNil())
				Expect(scID).NotTo(BeEmpty())

				sc = coreSvc.ScheduleByID(scID)
				Expect(sc).To(BeAssignableToTypeOf(scc))

				out, _ := json.MarshalIndent(coreSvc.AllSchedules(), "", "  ")
				Expect(len(coreSvc.AllSchedules())).To(Equal(count + 1), string(out))

				out, _ = json.MarshalIndent(sc, "", "  ")
					//fmt.Printf("\nSchedule: %s\n", out)
				Expect(sc.FirmwareID).NotTo(BeEmpty(), string(out))
			})
		})
	})

})
