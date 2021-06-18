package deviceScheduler_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDeviceScheduler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DeviceScheduler Suite")
}
