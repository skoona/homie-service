package deviceCore_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDeviceCore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DeviceCore Suite")
}
