package mqttProvider_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMqttProvider(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MqttProvider Suite")
}
