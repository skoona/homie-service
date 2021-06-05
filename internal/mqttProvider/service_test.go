package mqttProvider_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	mq "github.com/skoona/homie-service/internal/mqttProvider"
	cc "github.com/skoona/homie-service/internal/utils"
	"os"
)
var cfg cc.Config
var oldArgs []string

// otap, dsp, networks, _ = mq.Initialize(cfg)                 // message stream
// err := mq.Start()

var _ = Describe("Service", func() {
	BeforeEach(func() {
		oldArgs = os.Args
		defer func() { os.Args = oldArgs }()
		os.Setenv("HOMIE_SERVICE_CONFIG_FILE", "test-config")
	})

	Context("Initializes properly ", func() {
		It("Initialize() without broker should panic", func() {
			//otap, dsp, networks, err := mq.Initialize(cfg)

			var err interface{}
			var bPanic bool = false

			os.Args = []string{oldArgs[0], "--config", ""} // force clearing of prior value
			os.Unsetenv("MQTT_BROKER")
			cfg = cc.BuildRuntimeConfig("Homie-Service-Test")

			defer func() {
				// recover from panic if one occured. Set err to nil otherwise.
				if err := recover(); err != nil {
					//err = fmt.Errorf("panic'd")
					bPanic = true
				}
			}()
			_, _, _, err2 := mq.Initialize(cfg)

			Expect(bPanic).To(BeTrue())
			fmt.Printf("\n ************** ERROR STRING 1 ************* = %s \n", err.(error).Error())
			fmt.Printf("\n ************** ERROR STRING 2 ************* = %s \n", err2.Error())
		})
	})

})
