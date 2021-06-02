package configs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	cc "github.com/skoona/homie-service/internal/utils"
	"os"
)

var _ = Describe("Utils", func() {
	It("should be a valid config", func() {
		os.Args = append(os.Args, "--config")
		os.Args = append(os.Args, "test-config")
		config := cc.BuildRuntimeConfig("Homie-Service-Test")
		Expect(config).ToNot(BeNil())
	})
})
