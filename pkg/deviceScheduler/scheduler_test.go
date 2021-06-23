package deviceScheduler_test

import (
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
)


var _ = Describe("Scheduler Service", func() {

	Context("Initializes properly ", func() {
		It("Start() returns valid service", func() {
			Expect(sched).NotTo(BeNil(), "No Service")
		})
	})

	Context("Scheduler Operations ", func() {
		It("Creates new Schedule and finds it by deviceID", func() {
			var device dc.Device
			fw := sn.Firmwares[0]
			for _, dv := range sn.DeviceNetworks["sknSensors"].Devices {
				device = dv
				break
			}
			out, _ := json.MarshalIndent(device, "", "  ")
			Expect(device.Name).ToNot(BeEmpty(), string(out))

			scheduleID, err := sched.CreateSchedule(device.Parent, string(device.ID), dc.Base64Strict, fw.ID)
			Expect(err).To(BeNil(), "must create new")

			schedule := sched.FindScheduleByDeviceID(string(device.ID))
			out, _ = json.MarshalIndent(sn.Schedules, "", "  ")
			Expect(schedule.ID).To(Equal(scheduleID), string(out))
		})
		It("Delete Schedule by scheduleID", func() {
			count := len(sn.Schedules)

			var device dc.Device
			fw := sn.Firmwares[0]
			for _, dv := range sn.DeviceNetworks["sknSensors"].Devices {
				device = dv
				//break
			}
			out, _ := json.MarshalIndent(device, "", "  ")
			Expect(device.Name).ToNot(BeEmpty(), string(out))

			schedule := sched.FindScheduleByDeviceID(string(device.ID))
			if schedule.ElementType == dc.CoreTypeSchedule {
				err = sched.DeleteSchedule(schedule.ID)
				out, _ = json.MarshalIndent(schedule, "", "  ")
				Expect(err).To(BeNil(), string(out))
			}

			count = len(sn.Schedules)

			scheduleID, err := sched.CreateSchedule(device.Parent, string(device.ID), dc.Base64Strict, fw.ID)
			Expect(err).To(BeNil(), "must create new")
			Expect(len(sn.Schedules)).To(Equal(count + 1), "should be one")

			schedule = sched.FindScheduleByDeviceID(string(device.ID))
			out, _ = json.MarshalIndent(sn.Schedules, "", "  ")
			Expect(schedule.ID).To(Equal(scheduleID), string(out))

			err = sched.DeleteSchedule(schedule.ID)
			Expect(err).To(BeNil(), "must delete schedule")

			Expect(len(sn.Schedules)).To(Equal(count), "should be zero")
		})
	})

	Context("Firmware Operations ", func() {
		It("Retrieves firmwares by firmwareID", func() {
			Expect(len(sn.Firmwares)).To(Equal(3), "should be three in package")

			fw, err := sched.GetFirmware(sn.Firmwares[0].ID)
			out, _ := json.MarshalIndent(sn.Firmwares[0], "", "  ")
			Expect(err).To(BeNil(), string(out))

			out, _ = json.MarshalIndent(fw, "", "  ")
			Expect(fw.ID).NotTo(BeEmpty(), string(out))
		})
		It("Retrieves all firmwares", func() {
			count := len(sn.Firmwares)
			Expect(len(sched.Firmwares())).To(Equal(count), "empty list of Firmwares")
		})
		It("Deletes firmware based on ID", func() {
			count := len(sn.Firmwares)
			Expect(sched.DeleteFirmware(sn.Firmwares[0].ID)).To(BeNil(), "delete without error")
			out, _ := json.MarshalIndent(sn.Firmwares, "", "  ")
			Expect(len(sn.Firmwares)).To(Equal(count - 1), string(out))
		})
	})
})
