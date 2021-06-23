package demoProvider_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDemoProvider(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DemoProvider Suite")
}
