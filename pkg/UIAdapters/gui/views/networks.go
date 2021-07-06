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
	tree := components.MakeTreeTab()
	side := container.New(layout.NewPaddedLayout(), tree)

	cards := container.New(layout.NewGridLayout(1))
	scroller := container.NewVScroll(cards)

	page := container.NewHSplit(side, scroller)
	page.SetOffset(0.4)

	nMaps := vp.siteNetworks.DeviceNetworks[vp.netSelectedStr].Devices // "sknSensors")
	if len(nMaps) == 0 {
		level.Error(vp.logger).Log("selected network", vp.netSelectedStr, "error", "empty")
		cards.Add(widget.NewLabel("no devices available on " + vp.netSelectedStr))
	} else {
		vp.logger.Log("selected network", vp.netSelectedStr)
		for _, device := range nMaps {
			card := components.SknNewDeviceCard(device)
			cards.Add(card)
			vp.logger.Log("completed device loading", device.Name)
		}
	}

	vp.networkSide  = side // retain for later update
	vp.networkCards = cards  // retain for later update

	return page
}
