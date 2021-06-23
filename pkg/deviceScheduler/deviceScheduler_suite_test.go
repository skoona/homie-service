package deviceScheduler_test

import (
	dp "github.com/skoona/homie-service/pkg/demoProvider"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	sch "github.com/skoona/homie-service/pkg/deviceScheduler"
	dss "github.com/skoona/homie-service/pkg/deviceSource"
	dds "github.com/skoona/homie-service/pkg/deviceStorage"
	cc "github.com/skoona/homie-service/pkg/utils"
	"os"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDeviceScheduler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DeviceScheduler Suite")
}

var (
	networks     []string
	dep          dc.DeviceEventProvider
	sched        dc.SchedulerProvider
	repo         dc.Repository
	otap         sch.OTAInteractor
	dsp          dss.StreamProvider
	err 		 error
	cfg			cc.Config
	oldArgs		[]string
	sn			*dc.SiteNetworks
)

var _ = BeforeSuite(func() {
	oldArgs = os.Args
	os.Args = []string{oldArgs[0], "--config", ""} // force clearing of prior value
	defer func() { os.Args = oldArgs }()
	os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "test-config")
	os.Args = []string{oldArgs[0], "--config", "test-config"}

	cfg, err = cc.BuildRuntimeConfig("Homie-Service-Test")
	if err != nil {
		Fail(err.Error())
	}

	// Start DemoProvider
	otap, dsp, networks, err = dp.Initialize(cfg)
	if err != nil {
		Fail(err.Error())
	}

	// Start Repository
	repo, err = dds.Start(cfg)
	if err != nil {
		Fail(err.Error())
	}

	// start device message handler
	dep, err = dss.Start(cfg, repo, dsp)
	if err != nil {
		Fail(err.Error())
	}

	//sn.DeviceNetworks["sknSensors"] = repo.RestoreNetworkFromDB("sknSensors")

	sched = sch.Start(cfg, otap, repo)
	//sched.ApplySiteNetworks(sn)

	// network logic -- may need scheduler
	_, sn = dc.Start(cfg, dep, sched, repo, networks)

	// Start traffic
	err = dp.Start()
	if err != nil {
		Fail(err.Error())
	}

	time.Sleep(5 * time.Second)  // allow network to populate

})

var _ = AfterSuite(func(){
	dp.Stop()
	dss.Stop()
	sch.Stop()
	dc.Stop()
	dds.Stop()
})
