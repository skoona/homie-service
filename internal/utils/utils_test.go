package configs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	cc "github.com/skoona/homie-service/internal/utils"
	"os"
)
var config cc.Config

var _ = Describe("Utils", func() {

	Context("Command line options are honored ", func() {
		It("should be a valid config", func() {
			os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "test-config")
			//os.Args = nil
			//os.Args = append(os.Args, "--config")
			//os.Args = append(os.Args, "test-config")
			config = cc.BuildRuntimeConfig("Homie-Service-Test")
			Expect(config).ToNot(BeNil())
			Expect(config.RunMode).To(Equal("test"))
		})
	})
})

var _ = Describe("Utils", func() {
	It("Runmode equals demo", func() {
		os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "demo-config")
		//os.Args = nil
		//os.Args = append(os.Args, "--config")
		//os.Args = append(os.Args, "demo-config")
		config = cc.BuildRuntimeConfig("Homie-Service-Test")
		Expect(config.RunMode).To(Equal("demo"))
	})
})

var _ = Describe("Utils", func() {
	It("Runmode equals live", func() {
		os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "mqtt-config")
		//os.Args = nil
		//os.Args = append(os.Args, "--config")
		//os.Args = append(os.Args, "mqtt-config")
		config = cc.BuildRuntimeConfig("Homie-Service-Test")
		Expect(config.RunMode).To(Equal("live"))
	})
})
