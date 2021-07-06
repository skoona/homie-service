package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/components"
)

// FirmwaresTab ui page which includes firmwares
// the is to be used on main page
func (vp *viewProvider) FirmwaresTab() fyne.CanvasObject{
	form := components.SknFirmwareSideForm()
	side := container.New(layout.NewPaddedLayout(), form)

	cards := container.New(layout.NewGridLayout(1))
	scroller := container.NewVScroll(cards)


	page := container.NewBorder(nil,nil, side, nil, scroller)

	if len(vp.siteNetworks.Firmwares) == 0 {
		cards.Add(widget.NewLabel("no firmware available"))
	} else {
		for _, fw := range vp.siteNetworks.Firmwares {
			card := components.SknNewFirmwareCards(fw)
			cards.Add(card)
		}
	}
	vp.firmwareSide = &form
	vp.firmwareCards = cards

	return page
}
