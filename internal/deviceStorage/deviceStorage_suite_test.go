package deviceStorage_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDeviceStorage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Device Storage (Repository) Suite")
}
