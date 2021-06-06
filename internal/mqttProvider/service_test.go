package mqttProvider_test


import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	mq "github.com/skoona/homie-service/internal/mqttProvider"
	cc "github.com/skoona/homie-service/internal/utils"
	"os"
)
var cfg cc.Config
var oldArgs []string

// otap, dsp, networks, _ = mq.Initialize(cfg)                 // message stream

var _ = Describe("Service", func() {
	BeforeEach(func() {
		oldArgs = os.Args
		defer func() { os.Args = oldArgs }()
		os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "test-config")
	})

	Context("Initializes properly ", func() {
		It("Initialize() without broker should err out", func() {
			//otap, dsp, networks, err := mq.Initialize(cfg)

			os.Args = []string{oldArgs[0], "--config", ""} // force clearing of prior value
			os.Unsetenv("MQTT_BROKER")
			cfg, err := cc.BuildRuntimeConfig("Homie-Service-Test")
			Expect(err).To(BeNil(), "Configuration must be provided")

			_, _, _, err = mq.Initialize(cfg)
			Expect(err).ToNot(BeNil(), err.Error())
		})
	})

})
