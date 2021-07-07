package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/components"
	cc "github.com/skoona/homie-service/pkg/utils"
	"net/url"
)


// [![works with MQTT Homie](https://homieiot.github.io/img/works-with-homie.svg "works with MQTT Homie")](https://homieiot.github.io/)
// [![works with MQTT Homie]
//       (https://homieiot.github.io/img/works-with-homie.svg "works with MQTT Homie")
//  ](https://homieiot.github.io/)

// HomeTab ui page which includes broadcasts
// this is to be used as the main page
func (vp *viewProvider) HomeTab() fyne.CanvasObject {
	image := cc.SknSelectImage("worksWithHomie")
	image.SetMinSize(fyne.NewSize(192, 128))

	uri, _ := url.Parse("https://homieiot.github.io/")
	specLink := widget.NewHyperlink("Homie for ESP8266/ESP32", uri)
	specLink.Alignment = fyne.TextAlignCenter

	top := container.NewVBox(image, specLink)

	if len(vp.siteNetworks.Broadcasts) == 0 {
		vp.homeCards = &widget.List{}
		return container.NewBorder(top,nil, nil, nil, widget.NewLabel("no network boradcasts available"))

	} else {

		listBox := components.SknBroadcastList(&vp.siteNetworks.Broadcasts)
		listBox.OnSelected = vp.OnBroadcastSelected
		vp.homeCards = listBox // retain ref for later updates
		return container.NewBorder(top,nil, nil, nil, listBox)
	}

}
