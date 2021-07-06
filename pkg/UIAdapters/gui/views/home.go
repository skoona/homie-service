package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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

	cards := container.New(layout.NewGridLayout(1))
	scroller := container.NewVScroll(cards)
	page := container.NewBorder(top,nil, nil, nil, scroller)

	if len(vp.siteNetworks.Broadcasts) == 0 {
		cards.Add(widget.NewLabel("no broadcasts available"))
	} else {
		for _, bc := range vp.siteNetworks.Broadcasts {
			card := components.SknNewBroadcastCards(bc)
			cards.Add(card)
		}
	}
	vp.homeCards = cards // retain ref for later updates

	return page
}
