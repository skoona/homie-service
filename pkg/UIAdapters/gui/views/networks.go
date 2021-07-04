package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/go-kit/kit/log/level"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/components"
	cc "github.com/skoona/homie-service/pkg/utils"
)

// NetworksTab ui page which includes networks
// the is to be used on main page
func (vp *viewProvider) NetworksTab() fyne.CanvasObject {

	clk := cc.SknSelectImage("wallClock")
	clk.Resize(fyne.NewSize(48,32))
	rect := canvas.NewRectangle(cc.Green)
	rect.SetMinSize(fyne.NewSize(64,48))

	side := container.New(layout.NewPaddedLayout(), rect, clk)

	cards := container.NewGridWithColumns(3)

	nMaps, err := (*vp.dSvc).NetworkDevices(vp.netSelectedStr) // "sknSensors")
	if err != nil {
		level.Error(vp.logger).Log("selected network", vp.netSelectedStr, "error", err.Error())
	}
	vp.logger.Log("selected network", vp.netSelectedStr)
	for _, device := range nMaps {
		card := components.NewDevice( device, vp.logger )
		cards.Add(card)
		vp.logger.Log("completed device loading", device.Name)
	}

	scroller := container.NewVScroll(cards)

	page := container.NewBorder(nil,nil, side, nil, scroller)

	return page
}
