package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	cc "github.com/skoona/homie-service/pkg/utils"
)

// SchedulesTab ui page which includes schedules
// the is to be used on main page
func (vp *viewProvider) SchedulesTab() fyne.CanvasObject {
	clk := cc.SknSelectImage("wallClock")
	clk.Resize(fyne.NewSize(64,64))
	rect := canvas.NewRectangle(cc.Green)
	rect.SetMinSize(fyne.NewSize(64,64))
	side := container.New(layout.NewPaddedLayout(), rect, clk)

	card := widget.NewCard("HomeOffice", "80%", widget.NewLabel("motion"))
		card.SetImage( cc.SknSelectThemedImage("toggleOff_r"))
	card1 := widget.NewCard("MediaRoom", "60%", widget.NewLabel("motion"))
		card1.SetImage(cc.SknSelectThemedImage("toggleOn_r"))

	content := container.New(layout.NewGridLayout(3), card, card1)
	scroller := container.NewVScroll(content)

	page := container.NewBorder(nil,nil, side, nil, scroller)

	return page
}
