package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/components"
)

// SchedulesTab ui page which includes schedules
// the is to be used on main page
func (vp *viewProvider) SchedulesTab() fyne.CanvasObject {
	form := components.SknScheduleSideForm()
	side := container.New(layout.NewPaddedLayout(), form)

	cards := container.New(layout.NewGridLayout(1))
	scroller := container.NewVScroll(cards)
	page := container.NewBorder(nil,nil, side, nil, scroller)

	if len(vp.siteNetworks.Schedules) == 0 {
		cards.Add(widget.NewLabel("no schedules available"))
	} else {
		for _, s := range vp.siteNetworks.Schedules {
			card := components.SknNewScheduleCards(s)
			cards.Add(card)
		}
	}

	vp.scheduleSide = &form
	vp.scheduleCards = cards

	return page
}
