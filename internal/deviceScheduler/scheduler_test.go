package deviceScheduler_test

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	dp "github.com/skoona/homie-service/internal/demoProvider"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	sch "github.com/skoona/homie-service/internal/deviceScheduler"
	dss "github.com/skoona/homie-service/internal/deviceSource"
	dds "github.com/skoona/homie-service/internal/deviceStorage"
	cc "github.com/skoona/homie-service/internal/utils"
	"os"
)

var (
	networks     []string
	//dep          dc.DeviceEventProvider
	sched        dc.SchedulerProvider
	repo         dc.Repository
	otap         sch.OTAInteractor
	dsp          dss.StreamProvider
	err 		 error
	cfg			cc.Config
	oldArgs		[]string
	sn			*dc.SiteNetworks
)

var _ = Describe("Scheduler Service", func() {
	BeforeEach(func() {
		oldArgs = os.Args
		os.Args = []string{oldArgs[0], "--config", ""} // force clearing of prior value
		defer func() { os.Args = oldArgs }()
		os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "test-config")
		os.Args = []string{oldArgs[0], "--debug", "true",  "--config", ""}
		cfg, err = cc.BuildRuntimeConfig("Homie-Service-Test")
		if err != nil {
			Fail(err.Error())
		}

		// Start DemoProvider
		otap, dsp, networks, err = dp.Initialize(cfg)
		if err != nil {
			Fail(err.Error())
		}
		sn = dc.NewSiteNetworks("ginkgo testing","Scheduler suite",
			 networks, []dc.Firmware{},	map[string]dc.Schedule{})

		// Start Repository
		repo, err = dds.Start(cfg)
		if err != nil {
			Fail(err.Error())
		}

		// Start traffic
		err = dp.Start()
		if err != nil {
			Fail(err.Error())
		}
	})

	AfterEach(func(){
		dp.Stop()
		dds.Stop()
	})

	Context("Initializes properly ", func() {
		It("Start() returns valid service", func() {
			sched = sch.Start(cfg, otap, repo)
			Expect(sched).ToNot(BeNil(), "No Service")
			sch.Stop()
		})
	})

	Context("Scheduler Operations ", func() {
		It("Builds catalog of Schedules", func() {
			sched = sch.Start(cfg, otap, repo)
			Expect(sched).ToNot(BeNil(), "No Service")
			sched.ApplySiteNetworks(sn)

			cat := sched.BuildScheduleCatalog()
			Expect(len(cat)).ToNot(Equal(0), "some may be there")
			sch.Stop()
		})
		It("Returns catalog of Schedules", func() {
			sched = sch.Start(cfg, otap, repo)
			Expect(sched).ToNot(BeNil(), "No Service")
			sched.ApplySiteNetworks(sn)

			cat := sched.BuildScheduleCatalog()
			Expect(len(cat)).ToNot(Equal(0), "no schedules yet")

			Expect(len(sched.Schedules())).To(Equal(len(cat)), "should be same count")

			sch.Stop()
		})
		It("Creates new Schedule and finds it by deviceID", func() {
			sched = sch.Start(cfg, otap, repo)
			Expect(sched).ToNot(BeNil(), "No Service")
			sched.ApplySiteNetworks(sn)

			schedules := repo.LoadSchedules()
			count := len(schedules)

			schedules = sched.BuildScheduleCatalog()
			Expect(len(schedules)).To(Equal(count), "same as count")

			firmwares := sched.BuildFirmwareCatalog()
			Expect(len(firmwares)).To(Equal(3), "should be three in package")

			var device dc.Device
			fw := firmwares[0]
			for _, dv := range sn.DeviceNetworks["snkSensors"].Devices {
				device = dv
			}
			scheduleID, err := sched.CreateSchedule(device.Parent, device.ID, dc.Base64Strict, fw.ID)
			Expect(err).To(BeNil(), "must create new")

			schedule := sched.FindScheduleByDeviceID(device.ID)
			out, _ := json.MarshalIndent(sn.Schedules, "", "  ")
			Expect(schedule.ID).To(Equal(scheduleID), string(out))

			sch.Stop()
		})
		It("Delete Schedule by scheduleID", func() {
			sched = sch.Start(cfg, otap, repo)
			Expect(sched).ToNot(BeNil(), "No Service")
			sched.ApplySiteNetworks(sn)

			schedules := repo.LoadSchedules()
			count := len(schedules)

			schedules = sched.BuildScheduleCatalog()
			Expect(len(schedules)).To(Equal(count), "no schedules yet")

			firmwares := sched.BuildFirmwareCatalog()
			Expect(len(firmwares)).To(Equal(3), "should be three in package")

			var device dc.Device
			fw := firmwares[0]
			for _, dv := range sn.DeviceNetworks["snkSensors"].Devices {
				device = dv
				break
			}
			scheduleID, err := sched.CreateSchedule(device.Parent, device.ID, dc.Base64Strict, fw.ID)
			Expect(err).To(BeNil(), "must create new")

			Expect(len(sn.Schedules)).To(Equal(count), "should be one")

			schedule := sched.FindScheduleByDeviceID(device.ID)
			out, _ := json.MarshalIndent(sn.Schedules, "", "  ")
			Expect(schedule.ID).To(Equal(scheduleID), string(out))

			err = sched.DeleteSchedule(schedule.ID)
			Expect(err).To(BeNil(), "must delete schedule")
			Expect(len(sn.Schedules)).To(Equal(count - 1), "should be zero")

			sch.Stop()
		})
	})

	Context("Firmware Operations ", func() {
		BeforeEach(func() {
			sched = sch.Start(cfg, otap, repo)
			Expect(sched).ToNot(BeNil(), "No Service")
			sched.ApplySiteNetworks(sn)
		})

		It("Builds catalog of Firmware", func() {
			cat := sched.BuildFirmwareCatalog()
			Expect(len(cat)).To(Equal(3), "should be three in package")
			sch.Stop()
		})
		It("Retrieves firmwares by firmwareID", func() {
			cat := sched.BuildFirmwareCatalog()
			Expect(len(cat)).To(Equal(3), "should be three in package")

			fw, err := sched.GetFirmware(cat[0].ID)
			out, _ := json.MarshalIndent(cat[0], "", "  ")
			Expect(err).To(BeNil(), string(out))

			out, _ = json.MarshalIndent(fw, "", "  ")
			Expect(fw.ID).NotTo(BeEmpty(), string(out))

			sch.Stop()
		})
		It("Retrieves all firmwares", func() {
			cat := sched.BuildFirmwareCatalog()
			Expect(len(cat)).To(Equal(3), "should be three in package")

			Expect(len(sched.Firmwares())).To(Equal(3), "empty list of Firmwares")
			sch.Stop()
		})
		It("Deletes firmware based on ID", func() {
			cat := sched.BuildFirmwareCatalog()
			Expect(len(cat)).To(Equal(3), "should be three in package")

			Expect(sched.DeleteFirmware(cat[0].ID)).To(BeNil(), "delete without error")

			out, _ := json.MarshalIndent(sn.Firmwares, "", "  ")
			Expect(len(sn.Firmwares)).To(Equal(2), string(out))

			sch.Stop()
		})
	})
})
