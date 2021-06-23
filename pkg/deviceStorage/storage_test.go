package deviceStorage_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	dds "github.com/skoona/homie-service/pkg/deviceStorage"
	cc "github.com/skoona/homie-service/pkg/utils"
	"os"
	"time"
)


var _ = Describe("Repository", func() {
	var (
		//networks     []string
		//dep          dc.DeviceEventProvider
		repo         dc.Repository
		//otap         sch.OTAInteractor
		//dsp          dss.StreamProvider
		err 		 error
		cfg			cc.Config
		oldArgs		[]string
		sn			*dc.SiteNetworks
	)

	schedule01 := dc.Schedule{
		ID: "5058f1af8388633f609cadb75a75dc91",
		ElementType: 22,
		DeviceID: "36ec4f3d3eedbe07b28e5fa6ab6ddb91",
		FirmwareID: "5a3a2c57b5084ee4243017beaa0df461",
		State: "pending",
		Status: "waiting",
		Retries: 0,
		Transport: 32,
		Scheduled: time.Now(),
		Completed: time.Now(),
	}
	schedule02 := dc.Schedule{
		ID: "5058f1af8388633f609cadb75a75dc92",
		ElementType: 22,
		DeviceID: "36ec4f3d3eedbe07b28e5fa6ab6ddb98",
		FirmwareID: "5a3a2c57b5084ee4243017beaa0df462",
		State: "pending",
		Status: "waiting",
		Retries: 0,
		Transport: 32,
		Scheduled: time.Now(),
		Completed: time.Now(),
	}
	schedule03 := dc.Schedule{
		ID: "1dbb8dcaa63aabc9c2739e775f1e94df",
		ElementType: 22,
		DeviceID: "36ec4f3d3eedbe07b28e5fa6ab6ddb9b",
		FirmwareID: "5a3a2c57b5084ee4243017beaa0df466",
		State: "complete",
		Status: "rejected",
		Retries: 0,
		Transport: 32,
		Scheduled: time.Now(),
		Completed: time.Now(),
	}

	//firmware := dc.Firmware{
	//	ID: "5a3a2c57b5084ee4243017beaa0df466",
	//	ElementType: 21,
	//	Name: "Environment-DS18B20",
	//	FileName: "../../dataDB/firmwares/Environment-DS18B20.bin",
	//	Version: "1.0.2",
	//	Path: "../../dataDB/firmwares/Environment-DS18B20.bin",
	//	Size: 561456,
	//	MD5Digest: "5a3a2c57b5084ee4243017beaa0df466",
	//	Brand: "SknSensors",
	//	Created: time.Now(),
	//}

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
		//_, _, networks, err = dp.Initialize(cfg)
		//if err != nil {
		//	Fail(err.Error())
		//}

		//sn = dc.NewSiteNetworks("ginkgo testing","Repository suite",
		//	networks, []dc.Firmware{},	map[string]dc.Schedule{})
		sn = dc.NewSiteNetworks("ginkgo testing","Repository suite",
			cfg.Dbc.DemoNetworks, []dc.Firmware{},	map[string]dc.Schedule{})

		// Start Repository
		repo, err = dds.Start(cfg)
		if err != nil {
			Fail(err.Error())
		}

		// Start traffic
		//err = dp.Start()
		//if err != nil {
		//	Fail(err.Error())
		//}
	})

	AfterEach(func(){
		//dp.Stop()
		dds.Stop()
	})

	Context("Initializes properly ", func() {
		It("Repository Exists...", func() {
			//var dbRepo dc.Repository
			Expect(repo).ToNot(BeNil(), "must be a valid repo")
		})
	})

	Context("Repository Operations ", func() {
		Context("Schedule Operations ", func() {
			It("Schedules Exists...", func() {
				sn.Schedules = repo.LoadSchedules()
				Expect(len(sn.Schedules)).ToNot(Equal(0), "starting set")
			})
			It("Stores Schedules...", func() {
				err = repo.StoreSchedule(schedule01)
				Expect(err).To(BeNil(), "stores one")
				err = repo.StoreSchedule(schedule02)
				Expect(err).To(BeNil(), "stores two")
				err = repo.StoreSchedule(schedule03)
				Expect(err).To(BeNil(), "stores three")
			})
			It("Stores Schedules and Removes Schedules...", func() {
				schedules := repo.LoadSchedules()
				count := len(schedules)

				err = repo.StoreSchedule(schedule01)
				Expect(err).To(BeNil(), "stores one")
				err = repo.StoreSchedule(schedule02)
				Expect(err).To(BeNil(), "stores two")
				err = repo.StoreSchedule(schedule03)
				Expect(err).To(BeNil(), "stores three")

				schedules = repo.LoadSchedules()
				Expect(len(schedules)).To(Equal((count - 3) + 3))

				_, ok := schedules[schedule02.ID]
				Expect(ok).To(BeTrue(), "failed to retrieve existing")

				_ = repo.RemoveSchedule(schedule01.ID)
				sn.Schedules = repo.LoadSchedules()
				Expect(len(sn.Schedules)).To(Equal((count-3)+2), "one less")
			})
		})
		Context("Broadcast Operations ", func() {
			It("Loads broadcasts...", func() {
				broadcasts := repo.LoadBroadcasts("sknSensors")
				Expect(len(broadcasts)).To(Equal(2))
			})
			It("Clears all broadcasts...", func() {
				err = repo.RemoveAllBroadcasts("sknSensors")
				Expect(err).To(BeNil(), "removes them all")
				broadcasts := repo.LoadBroadcasts("sknSensors")
				Expect(len(broadcasts)).To(Equal(0))
			})
		})
		Context("Network Operations ", func() {
		})
	})
})
