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
	vp.firmwareSide = &form

	if len(vp.siteNetworks.Firmwares) == 0 {
		vp.firmwareCards = &widget.List{}
		return container.NewBorder(nil,nil, side, nil, widget.NewLabel("no firmware available"))

	} else {
		listBox := components.SknFirmwareList(&vp.siteNetworks.Firmwares)
		listBox.OnSelected = vp.OnFirmwareSelected
		vp.firmwareCards = listBox
		return container.NewBorder(nil,nil, side, nil, listBox)
	}

}
