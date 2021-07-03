package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	cc "github.com/skoona/homie-service/pkg/utils"
)

// FirmwaresTab ui page which includes firmwares
// the is to be used on main page
func (vp *viewProvider) FirmwaresTab() fyne.CanvasObject{
	clk := cc.SknSelectImage("wallClock")
	clk.Resize(fyne.NewSize(48,32))
	rect := canvas.NewRectangle(cc.Green)
	rect.SetMinSize(fyne.NewSize(64,48))

	side := container.NewPadded(rect, clk)
	content := container.NewGridWithColumns(3)

	for _, fw := range (*vp.dSvc).AllFirmwares() {
		card := widget.NewCard(fw.Name, fw.Version, widget.NewLabel(fw.Created.String()))
		card.SetImage( cc.SknSelectThemedImage("insertDriveFile_r"))
		content.Add(card)
	}

	scroller := container.NewVScroll(content)
	page := container.NewBorder(nil,nil, side, nil, scroller)

	return page
}
