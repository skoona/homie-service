package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/components"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"github.com/skoona/homie-service/pkg/deviceScheduler"
)

// SchedulesTab ui page which includes schedules
// the is to be used on main page
func (vp *viewProvider) SchedulesTab() fyne.CanvasObject {
	form := components.SknScheduleSideForm()
	side := container.New(layout.NewPaddedLayout(), form)
	vp.scheduleSide = &form

	if len(vp.siteNetworks.Schedules) != 0 {      // todo equal equal when real data
		vp.scheduleCards = &widget.List{}
		return container.NewBorder(nil,nil, side, nil, widget.NewLabel("no schedules available"))

	} else {
		schedules := []dc.Schedule{
			deviceScheduler.NewSchedule("sknSensors","2f63939498209cf4584b6e82954d2407", dc.Base64Strict, "3ef8fbb48c5b23788a22a998c14a1a6d"),
			deviceScheduler.NewSchedule("sknSensors","36ec4f3d3eedbe07b28e5fa6ab6ddb9b", dc.Base64Strict, "3ef8fbb48c5b23788a22a998c14a1a6d"),
			deviceScheduler.NewSchedule("sknSensors","bdab9513512a9d316939b5c0ac1ee099", dc.Base64Strict, "3ef8fbb48c5b23788a22a998c14a1a6d"),
			deviceScheduler.NewSchedule("sknSensors","8c3705faef6bf02d8df1bb7b9516e182", dc.Base64Strict, "3ef8fbb48c5b23788a22a998c14a1a6d"),
		}
		schedules[0].State = "complete"
		schedules[0].Status = "current"

		schedules[2].Status = "initializing"
		schedules[2].State = "active"

		for _, schedule := range vp.siteNetworks.Schedules {
			schedules = append(schedules, schedule)
		}
		vp.schedules = schedules
		listBox := components.SknScheduleList(schedules)
		listBox.OnSelected = vp.OnScheduleSelected
		vp.scheduleCards = listBox
		return container.NewBorder(nil,nil, side, nil, listBox)
	}
}
