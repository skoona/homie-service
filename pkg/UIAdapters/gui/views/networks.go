package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/go-kit/kit/log/level"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/components"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
)

// NetworksTab ui page which includes networks
// the is to be used on main page
func (vp *viewProvider) NetworksTab() fyne.CanvasObject {
	tree := components.MakeTreeTab()
	side := container.New(layout.NewPaddedLayout(), tree)
	vp.networkSide  = side // retain for later update

	nMaps := vp.siteNetworks.DeviceNetworks[vp.netSelectedStr].Devices // "sknSensors")
	if len(nMaps) == 0 {
		level.Error(vp.logger).Log("selected network", vp.netSelectedStr, "error", "empty")
		page := container.NewHSplit(side, widget.NewLabel("no devices available on " + vp.netSelectedStr))
		return page
	} else {
		vp.logger.Log("selected network", vp.netSelectedStr)
		var devices []dc.Device =  []dc.Device{}
		for _, device := range nMaps {
			devices = append(devices, device)
			vp.logger.Log("completed device loading", device.Name)
		}
		listBox := components.SknDeviceList(devices)
		listBox.OnSelected = vp.OnDeviceSelected
		page := container.NewHSplit(side, listBox)
		page.SetOffset(0.4)
		vp.networkCards = listBox  // retain for later update
		return page
	}

}
