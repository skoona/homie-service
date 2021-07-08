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
	nMaps := vp.siteNetworks.DeviceNetworks[vp.netSelectedStr].Devices // "sknSensors")

	if len(nMaps) == 0 {
		tree := widget.NewLabel("no devices")
		side := container.New(layout.NewPaddedLayout(), tree)
		vp.networkSide  = side // retain for later update

		level.Error(vp.logger).Log("selected network", vp.netSelectedStr, "error", "empty")
		vp.networkCards = &widget.List{}
		return container.NewHSplit(side, widget.NewLabel("no devices available on " + vp.netSelectedStr))

	} else {
		vp.logger.Log("selected network", vp.netSelectedStr)
		devices :=  []dc.Device{}
		sideDevice := &dc.Device{}
		for _, device := range nMaps {
			if sideDevice.ElementType != dc.CoreTypeDevice {
				sideDevice = &device
			}
			devices = append(devices, device)
			vp.logger.Log("completed device loading", device.Name)
		}
		tree := components.SknDeviceTreeSide(sideDevice, &vp.devTreeDetails)
		tree.(*widget.Tree).OnSelected = vp.OnDeviceTreeSelected
		side := container.New(layout.NewPaddedLayout(), tree)

		vp.networkSide  = side // retain for later update
		vp.devices = devices

		listBox := components.SknDeviceList(devices)
		listBox.OnSelected = vp.OnDeviceSelected
		vp.networkCards = listBox  // retain for later update
		page := container.NewHSplit(side, listBox)
		page.SetOffset(0.4)
		return page
	}

}
