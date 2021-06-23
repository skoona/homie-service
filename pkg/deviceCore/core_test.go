package deviceCore_test

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
)

var _ = Describe("Core Service", func() {

	Context("Core Operations", func() {
		Context("Network Operations", func() {
			It("AllNetworks() returns the SiteNetwork", func() {
				var sn dc.SiteNetworks
				Expect(coreSvc.AllNetworks()).To(BeAssignableToTypeOf(sn))
			})
			It("AllNetworks() has known network", func() {
				sn := coreSvc.AllNetworks()
				Expect(len(sn.DeviceNetworks)).To(Equal(1))
				Expect(sn.DeviceNetworks["sknSensors"]).ToNot(BeNil())
			})
			It("AllNetworks() known network has Devices", func() {
				sn := coreSvc.AllNetworks()
				Expect(len(sn.DeviceNetworks["sknSensors"].Devices)).NotTo(Equal(0))
			})
		})
		Context("Broadcast Operations", func() {
			It("AllBroadcasts() returns the available broadcasts", func() {
				var bc []dc.Broadcast
				Expect(coreSvc.AllBroadcasts()).To(BeAssignableToTypeOf(bc))
			})
			It("AllBroadcasts() have known items", func() {
				bc := coreSvc.AllBroadcasts()
				Expect(len(bc)).ToNot(Equal(0))
				Expect(bc[0]).ToNot(BeNil())
			})
		})
		Context("Firmware Operations", func() {
			It("AllFirmwares() returns the current library", func() {
				var fw []dc.Firmware
				Expect(coreSvc.AllFirmwares()).To(BeAssignableToTypeOf(fw))
			})
			It("AllFirmwares() have known items", func() {
				var fw []dc.Firmware
				fws := coreSvc.AllFirmwares()
				Expect(len(fws)).ToNot(Equal(0))
				Expect(fw).To(BeAssignableToTypeOf(fw))
			})
		})
		Context("Schedule Operations", func() {
			It("AllSchedules() returns the available schedules", func() {
				var sc []dc.Schedule
				Expect(coreSvc.AllSchedules()).To(BeAssignableToTypeOf(sc))

				schedules := repo.LoadSchedules()
				count := len(schedules)

				out, _ := json.MarshalIndent(coreSvc.AllSchedules(), "", "  ")
				Expect(len(coreSvc.AllSchedules())).To(Equal(count), string(out))
			})
			It("CreateSchedule() returns a new schedules", func() {
				var fw dc.Firmware
				var dv dc.Device
				var sc dc.Schedule
				var scc dc.Schedule
				schedules := coreSvc.AllSchedules()
				count := len(schedules)

				dv, err = coreSvc.DeviceByName("GuestRoom", "sknSensors")
					//out, _ := json.MarshalIndent(dv, "", "  ")
					//fmt.Printf("\nDevice: %s\n", out)
				fw = coreSvc.AllFirmwares()[0]
					//out, _ = json.MarshalIndent(fw, "", "  ")
					//fmt.Printf("\nFirmware: %s\n", out)

				scID, err := coreSvc.CreateSchedule("sknSensors", string(dv.ID), dc.Base64Strict, fw.ID)
				Expect(err).To(BeNil())
				Expect(scID).NotTo(BeEmpty())

				sc = coreSvc.ScheduleByID(scID)
				Expect(sc).To(BeAssignableToTypeOf(scc))

				out, _ := json.MarshalIndent(coreSvc.AllSchedules(), "", "  ")
				Expect(len(coreSvc.AllSchedules())).To(Equal(count + 1), string(out))

				out, _ = json.MarshalIndent(sc, "", "  ")
					//fmt.Printf("\nSchedule: %s\n", out)
				Expect(sc.FirmwareID).NotTo(BeEmpty(), string(out))
			})
		})
	})

})
