package demoProvider_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	dp "github.com/skoona/homie-service/internal/demoProvider"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	cc "github.com/skoona/homie-service/internal/utils"
	"os"
	"sync"
)
var cfg cc.Config
var oldArgs []string
var err error

var _ = Describe("DemoProvider inactive service", func() {

	BeforeEach(func() {
		oldArgs = os.Args
		os.Args = []string{oldArgs[0], "--config", ""} // force clearing of prior value
		defer func() { os.Args = oldArgs }()
		os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "test-config")
		os.Args = []string{oldArgs[0], "--config", ""}
		cfg, err = cc.BuildRuntimeConfig("Homie-Service-Test")
	})

	Context("Initializes properly ", func() {
		It("Start() errors out if called before initialize()", func() {
			err := dp.Start()
			Expect(err).ToNot(BeNil(), "Service did not Fail")
		})


		It("#Initialize() returns a valid StreamProvider service", func() {
			sch, sp, aStr, err := dp.Initialize(cfg)
			Expect(err).To(BeNil())
			Expect(aStr[0]).To(Equal(cfg.Dbc.DemoNetworks[0]))
			Expect(sp).ToNot(BeNil())
			Expect(sch).ToNot(BeNil())
			dp.Stop()
		})

		It("StreamProvider supports a Publish() Channel", func() {
			_, sp, _, err := dp.Initialize(cfg)
			Expect(err).To(BeNil())
			err = dp.Start()
			Expect(err).To(BeNil())
			Expect(sp.GetPublishChannel()).ToNot(BeNil())
			dp.Stop()
		})
	})

})
var _ = Describe("DemoProvider active service", func() {
	BeforeEach(func() {
		oldArgs = os.Args
		os.Args = []string{oldArgs[0], "--config", "test-config"} // force clearing of prior value
		defer func() { os.Args = oldArgs }()
		cfg, err = cc.BuildRuntimeConfig("Homie-Service-Test")
	})

	Context("Produces demo messages ", func() {
		It("StreamProvider produces messages on notify channel", func() {
			var wg sync.WaitGroup
			_, sp, _, _ := dp.Initialize(cfg)
			err = dp.Start()
			Expect(err).To(BeNil())
			var counter int = 0
			wg.Add(1)
			go func(mChan chan dc.DeviceMessage) {
				defer wg.Done()
				for dm := range mChan {
					fmt.Printf("DeviceMessage: %s\n", dm.String())
					counter += 1
					if counter >= 106 { // number of lines in test log
						break
					}
				}
				fmt.Printf("Test drain completed, counter=%d\n", counter)
			}(sp.ActivateNotifications())
			wg.Wait()
			dp.Stop()
			Expect(counter).To(Equal(106))
		})
	})
})
