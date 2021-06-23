package deviceCore_test

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dp "github.com/skoona/homie-service/pkg/demoProvider"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	sch "github.com/skoona/homie-service/pkg/deviceScheduler"
	dss "github.com/skoona/homie-service/pkg/deviceSource"
	dds "github.com/skoona/homie-service/pkg/deviceStorage"
	mq "github.com/skoona/homie-service/pkg/mqttProvider"
	cc "github.com/skoona/homie-service/pkg/utils"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDeviceCore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DeviceCore Suite")
}

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

var _ = BeforeSuite(func() {
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

var _ = AfterSuite(func(){
	Shutdown(cfg)
})
