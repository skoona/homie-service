package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/components"
	cc "github.com/skoona/homie-service/pkg/utils"
)

// FirmwaresTab ui page which includes firmwares
// the is to be used on main page
func (vp *viewProvider) FirmwaresTab() fyne.CanvasObject{
	form := components.MakeFormSide()
	side := container.New(layout.NewPaddedLayout(), form)

	cards := container.NewGridWrap(fyne.NewSize(300, 400))

	for _, fw := range (*vp.dSvc).AllFirmwares() {
		rec := widget.NewLabel(fw.Created.String())
		rec.Wrapping = fyne.TextWrapWord
		card := widget.NewCard(fw.Name, fw.Version, rec)
		card.SetImage( cc.SknSelectThemedImage("insertDriveFile_r"))
		cards.Add(card)
	}

	scroller := container.NewVScroll(cards)
	page := container.NewBorder(nil,nil, side, nil, scroller)

	vp.firmwareSide = &form
	vp.firmwareCards = cards

	return page
}
