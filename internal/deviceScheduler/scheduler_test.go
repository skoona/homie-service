package deviceScheduler_test

import (
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
		os.Args = []string{oldArgs[0], "--config", ""}
		cfg, err = cc.BuildRuntimeConfig("Homie-Service-Test")
		if err != nil {
			Fail(err.Error())
		}

		// Start DemoProvider
		otap, dsp, networks, err = dp.Initialize(cfg)
		if err != nil {
			Fail(err.Error())
		}
		sn = dc.NewSiteNetworks("ginkoo testing","Scheduler suite",
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
			Expect(len(cat)).To(Equal(0), "no schedules yet")
			sch.Stop()
		})
		It("Builds catalog of Firmware", func() {
			sched = sch.Start(cfg, otap, repo)
			Expect(sched).ToNot(BeNil(), "No Service")
			sched.ApplySiteNetworks(sn)

			cat := sched.BuildFirmwareCatalog()
			Expect(len(cat)).To(Equal(3), "should be three in package")
			sch.Stop()
		})
	})

})
