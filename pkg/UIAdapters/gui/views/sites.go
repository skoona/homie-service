package views

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-kit/kit/log/level"
)


// [![works with MQTT Homie](https://homieiot.github.io/img/works-with-homie.svg "works with MQTT Homie")](https://homieiot.github.io/)
// [![works with MQTT Homie]
//       (https://homieiot.github.io/img/works-with-homie.svg "works with MQTT Homie")
//  ](https://homieiot.github.io/)

// SitesTab ui page which includes all SiteNetworks
// this is to be used as the main page
func (vp *viewProvider) SitesTab() fyne.CanvasObject {

	content := container.NewPadded()

	multiEntry := widget.NewMultiLineEntry()
	multiEntry.Wrapping = fyne.TextWrapWord
	multiEntry.TextStyle = fyne.TextStyle{Monospace: true}
	vp.siteMlEntry = multiEntry // retain ref for latter updates

	content.Add(multiEntry)

	// Dump the SiteNetwork and all nodes as JSON
	out, err := json.MarshalIndent((*vp.dSvc).AllNetworks(), "", "  ")
	if err != nil {
		level.Warn(vp.logger).Log("tab", "Preferences", "error", err.Error())
	} else {
		multiEntry.SetText(string(out))
	}

	return content
}
