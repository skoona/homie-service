package deviceStorage_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	dds "github.com/skoona/homie-service/internal/deviceStorage"
	cc "github.com/skoona/homie-service/internal/utils"
	"os"
)


var _ = Describe("Repository", func() {

	var (
		 cfg cc.Config
		 oldArgs []string
		 err error
		 repo dc.Repository
	)

	BeforeEach(func() {
		oldArgs = os.Args
		defer func() { os.Args = oldArgs }()
		os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "test-config")
		os.Args = []string{oldArgs[0], "--debug", "true", "--config", "test-config"} // force clearing of prior value
		os.Unsetenv("MQTT_BROKER")
		cfg, err = cc.BuildRuntimeConfig("Homie-Service-Test")
		if err != nil {
			Fail(err.Error())
		}
	})

	Context("Initializes properly ", func() {
		It("Start() returns a valid repo", func() {

			repo, err = dds.Start(cfg)               // message db

			Expect(err).To(BeNil(), "should not produce an error on startup")
			Expect(repo).NotTo(BeNil(), "must be a valid repo")
			//Expect(aRepo).To(BeEquivalentTo(bRepo))
			dds.Stop()
		})
	})
})
