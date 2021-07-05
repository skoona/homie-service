package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/go-kit/kit/log/level"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/components"
)

// NetworksTab ui page which includes networks
// the is to be used on main page
func (vp *viewProvider) NetworksTab() fyne.CanvasObject {
	form := components.SknSideForm()
	side := container.New(layout.NewPaddedLayout(), form)

	cards := container.NewGridWrap(fyne.NewSize( CardWidth, CardHeight))

	nMaps, err := (*vp.dSvc).NetworkDevices(vp.netSelectedStr) // "sknSensors")
	if err != nil {
		level.Error(vp.logger).Log("selected network", vp.netSelectedStr, "error", err.Error())
		cards.Add(widget.NewLabel(err.Error()))
	} else {
		vp.logger.Log("selected network", vp.netSelectedStr)
		for _, device := range nMaps {
			card := components.SknNewDeviceBadge(&device, vp.logger)
			cards.Add(card)
			vp.logger.Log("completed device loading", device.Name)
		}
	}
	scroller := container.NewVScroll(cards)

	page := container.NewBorder(nil,nil, side, nil, scroller)

	vp.networkSide = &form // retain for later update
	vp.networkCards = cards  // retain for later update

	return page
}
