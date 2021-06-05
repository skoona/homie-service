package demoProvider_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	dp "github.com/skoona/homie-service/internal/demoProvider"
	cc "github.com/skoona/homie-service/internal/utils"
	"os"
)

var _ = Describe("DemoProvider service", func() {
	var cfg cc.Config
	var oldArgs []string

	BeforeEach(func () {
		oldArgs = os.Args
		os.Args = []string{oldArgs[0], "--config", ""}  // force clearing of prior value
		defer func() { os.Args = oldArgs }()
		cfg = cc.BuildRuntimeConfig("Homie-Service-Test")
		os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "test-config")
		os.Args = []string{oldArgs[0], "--config", ""}
	})

	Context("Initializes properly ", func() {
		It("should return a valid StreamProvider service", func() {
			sp, aStr, err := dp.Initialize(cfg)
			Expect(err).To(BeNil())
			Expect(aStr).To(Equal(cfg.Dbc.DemoNetworks))
			Expect(sp).ToNot(BeNil())
		})

		It("StreamProvider does not support a Publish() Channel", func() {
			sp, _, err := dp.Initialize(cfg)
			Expect(sp).ToNot(BeNil())
			Expect(err).To(BeNil())
			Expect(sp.GetPublishChannel()).To(BeNil())

			err = dp.Start()
			Expect(err).To(BeNil())
		})
	})


})
