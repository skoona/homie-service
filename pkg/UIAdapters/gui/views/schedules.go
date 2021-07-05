package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/components"
	cc "github.com/skoona/homie-service/pkg/utils"
)

// SchedulesTab ui page which includes schedules
// the is to be used on main page
func (vp *viewProvider) SchedulesTab() fyne.CanvasObject {
	form := components.SknScheduleSideForm()
	side := container.New(layout.NewPaddedLayout(), form)

	card := widget.NewCard("HomeOffice", "80%", widget.NewLabel("motion"))
		card.SetImage( cc.SknSelectThemedImage("wallClock"))
	card1 := widget.NewCard("MediaRoom", "60%", widget.NewLabel("motion"))
		card1.SetImage(cc.SknSelectThemedImage("wallClock"))

	cards := container.NewGridWrap(fyne.NewSize( CardWidth, CardHeight), card, card1)

	scroller := container.NewVScroll(cards)

	page := container.NewBorder(nil,nil, side, nil, scroller)

	vp.scheduleSide = &form
	vp.scheduleCards = cards

	return page
}
