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

var _ = Describe("DemoProvider inactive service", func() {

	BeforeEach(func() {
		oldArgs = os.Args
		os.Args = []string{oldArgs[0], "--config", ""} // force clearing of prior value
		defer func() { os.Args = oldArgs }()
		os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "test-config")
		os.Args = []string{oldArgs[0], "--config", ""}
		cfg = cc.BuildRuntimeConfig("Homie-Service-Test")
	})

	Context("Initializes properly ", func() {
		It("Start() complains and panics if called before initialize()", func() {
			var err interface{}
			defer func() {
				// recover from panic if one occured. Set err to nil otherwise.
				err = recover()
			}()
			//err = dp.Start()
			//Expect( err.(error).Error() ).To(Equal("you must call Initialize() in this package before calling Start()"))
			//Expect(err.(error).Error()).To(Equal("why any string"))
			Expect(dp.Start()).To(Panic())
			fmt.Printf("\n ************** ERROR STRING ************* = %s \n", err.(error).Error())
		})


		It("#Initialize() returns a valid StreamProvider service", func() {
			sp, aStr, err := dp.Initialize(cfg)
			Expect(err).To(BeNil())
			Expect(aStr).To(Equal(cfg.Dbc.DemoNetworks))
			Expect(sp).ToNot(BeNil())
		})

		It("StreamProvider does not support a Publish() Channel", func() {
			sp, _, err := dp.Initialize(cfg)
			Expect(err).To(BeNil())
			err = dp.Start()
			Expect(err).To(BeNil())
			Expect(sp.GetPublishChannel()).To(BeNil())
		})
	})

})
var _ = Describe("DemoProvider active service", func() {
	BeforeEach(func() {
		oldArgs = os.Args
		os.Args = []string{oldArgs[0], "--config", "test-config"} // force clearing of prior value
		defer func() { os.Args = oldArgs }()
		cfg = cc.BuildRuntimeConfig("Homie-Service-Test")
	})

	Context("Produces demo messages ", func() {
		It("StreamProvider produces messages on notify channel", func() {
			var wg sync.WaitGroup
			sp, _, err := dp.Initialize(cfg)
			err = dp.Start()
			Expect(err).To(BeNil())
			var counter int = 0
			wg.Add(1)
			go func(mChan chan dc.DeviceMessage) {
				defer wg.Done()
				for range mChan {
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
