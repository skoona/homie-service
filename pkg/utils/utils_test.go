package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	cc "github.com/skoona/homie-service/pkg/utils"
)

var _ = Describe("Utils Utils", func() {

	Context("LocateFile() Operations", func() {
		It("returns filename when file exists and is not a directory", func() {
			Expect(cc.LocateFile("utils_suite_test.go")).To(Equal("utils_suite_test.go"))
		})
		It("returns filename when file exists but is in current dir (i.e. ../../file)", func() {
			Expect(cc.LocateFile("../../utils_suite_test.go")).To(Equal("./utils_suite_test.go"))
		})
		It("returns empty string when is a directory vs a file", func() {
			Expect(cc.LocateFile("../utils")).To(BeEmpty())
		})
		It("returns empty string when file is not found", func() {
			Expect(cc.LocateFile("utils_suites_test.go")).To(BeEmpty())
		})
	})

	Context("RemoveIndexFromSlice() Operations", func() {
		It("removes indexed item as expected", func() {
			slice := []int{1,2,3,4}
			expected := []int{1,2,4}
			Expect(cc.RemoveIndexFromSlice(slice, 2)).To(Equal(expected))
		})
		It("handles index out of bounds", func() {
			slice := []int{1,2,3,4}
			expected := []int{1,2,3}
			Expect(cc.RemoveIndexFromSlice(slice, 6)).To(Equal(expected))
		})
	})
	Context("HumanFormatFromBytes() Operations", func() {
		It("Handles 120 B", func() {
			Expect(cc.HumanFormatFromBytes(120)).To(Equal("120 B"))
		})
		It("Handles 1 KB", func() {
			Expect(cc.HumanFormatFromBytes(1200)).To(Equal("1.2 KB"))
		})
		It("Handles 100 KB", func() {
			Expect(cc.HumanFormatFromBytes(120000)).To(Equal("117.2 KB"))
		})
		It("Handles 1 MB", func() {
			Expect(cc.HumanFormatFromBytes(1200000)).To(Equal("1.1 MB"))
		})
		It("Handles 1 GB", func() {
			Expect(cc.HumanFormatFromBytes(1200000000)).To(Equal("1.1 GB"))
		})
		It("Handles 1 TB", func() {
			Expect(cc.HumanFormatFromBytes(1200000000000)).To(Equal("1.1 TB"))
		})
	})

})
