package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	cc "github.com/skoona/homie-service/pkg/utils"
	"os"
)

var _ = Describe("Utils Configs", func() {
	var oldArgs []string

	BeforeEach(func () {
		oldArgs = os.Args
		os.Args = []string{oldArgs[0], "--config", ""}  // force clearing of prior value
		defer func() { os.Args = oldArgs }()
		os.Unsetenv("HOMIE_SERVICE_CONFIG_FILE")
	})

	Context("Defaults to 'mqtt-config / live' when no commandline or env is given ", func() {
		It("should be a valid config", func() {
			os.Args = []string{oldArgs[0], "--config", ""}
			acfg, err := cc.BuildRuntimeConfig("Homie-Service-Test")
			Expect(err).To(BeNil(), "Configuration must be provided")

			Expect(acfg).ToNot(BeNil())
			Expect(acfg.RunMode).To(Equal("live"))
		})
	})

	Context("Environment is used when no commandline is given ", func() {
		It("should be a valid config", func() {
			os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "test-config")
			os.Args = []string{oldArgs[0], "--config", ""}

			acfg, err := cc.BuildRuntimeConfig("Homie-Service-Test")
			Expect(err).To(BeNil(), "Configuration must be provided")
			Expect(acfg).ToNot(BeNil())
			Expect(acfg.RunMode).To(Equal("test"))
		})
		It("Runmode equals live", func() {
			os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "mqtt-config")

			bcfg, err := cc.BuildRuntimeConfig("Homie-Service-Test")
			Expect(err).To(BeNil(), "Configuration must be provided")
			Expect(bcfg.RunMode).To(Equal("live"))
		})
		It("Runmode equals demo", func() {
			os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "demo-config")
			os.Args = []string{oldArgs[0], "--config", ""}

			ccfg, err := cc.BuildRuntimeConfig("Homie-Service-Test")
			Expect(err).To(BeNil(), "Configuration must be provided")
			Expect(ccfg.RunMode).To(Equal("demo"))
		})
	})

	Context("Command line options are honored ", func() {
		It("should be a valid config", func() {
			os.Args = []string{oldArgs[0], "--config", "test-config"}

			acfg, err := cc.BuildRuntimeConfig("Homie-Service-Test")
			Expect(err).To(BeNil(), "Configuration must be provided")
			Expect(acfg).ToNot(BeNil())
			Expect(acfg.RunMode).To(Equal("test"))
		})
		It("Runmode equals live", func() {
			os.Args = []string{oldArgs[0], "--config", "mqtt-config"}

			bcfg, err := cc.BuildRuntimeConfig("Homie-Service-Test")
			Expect(err).To(BeNil(), "Configuration must be provided")
			Expect(bcfg.RunMode).To(Equal("live"))
		})
		It("Runmode equals demo", func() {
			os.Args = []string{oldArgs[0], "--config", "demo-config"}

			ccfg, err := cc.BuildRuntimeConfig("Homie-Service-Test")
			Expect(err).To(BeNil(), "Configuration must be provided")
			Expect(ccfg.RunMode).To(Equal("demo"))
		})
	})

	Context("Command line options OVERRIDE environment ", func() {
		It("should be a valid config", func() {
			os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "demo-config")
			os.Args = []string{oldArgs[0], "--config", "test-config"}

			acfg, err := cc.BuildRuntimeConfig("Homie-Service-Test")
			Expect(err).To(BeNil(), "Configuration must be provided")
			Expect(acfg).ToNot(BeNil())
			Expect(acfg.RunMode).To(Equal("test"))
		})
		It("Runmode equals live", func() {
			os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "test-config")
			os.Args = []string{oldArgs[0], "--config", "mqtt-config"}

			bcfg, err := cc.BuildRuntimeConfig("Homie-Service-Test")
			Expect(err).To(BeNil(), "Configuration must be provided")
			Expect(bcfg.RunMode).To(Equal("live"))
		})
		It("Runmode equals demo", func() {
			os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "test-config")
			os.Args = []string{oldArgs[0], "--config", "demo-config"}

			ccfg, err := cc.BuildRuntimeConfig("Homie-Service-Test")
			Expect(err).To(BeNil(), "Configuration must be provided")
			Expect(ccfg.RunMode).To(Equal("demo"))
		})
	})
})
