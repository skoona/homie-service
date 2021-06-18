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
		err = dp.Start()
		if err != nil {
			Fail(err.Error())
		}

		// Start Repository
		repo, err = dds.Start(cfg)
		if err != nil {
			Fail(err.Error())
		}
	})

	AfterEach(func(){
		dds.Stop()
		dp.Stop()
	})

	Context("Initializes properly ", func() {
		It("Start() returns valid service", func() {
			sched = sch.Start(cfg, otap, repo)
			Expect(sched).ToNot(BeNil(), "No Service")
			sch.Stop()
		})
	})

})
