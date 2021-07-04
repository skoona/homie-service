package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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

	cards := container.NewGridWithColumns(3)
	vp.homeCards = cards // retain ref for later updates

	for _, bc := range (*vp.dSvc).AllBroadcasts() {
		card := widget.NewCard(bc.Level, bc.Value, widget.NewLabel(bc.Received.String()))
		card.SetImage( cc.SknSelectThemedImage("notificationAlert_o"))
		cards.Add(card)
	}

	image := cc.SknSelectImage("worksWithHomie")
	image.SetMinSize(fyne.NewSize(160, 120))
	uri, _ := url.Parse("https://homieiot.github.io/")
	specLink := widget.NewHyperlink("Homie for ESP8266/ESP32", uri)
	specLink.Alignment = fyne.TextAlignCenter
	banner := container.NewVBox(image, specLink)

	scroller := container.NewVScroll(cards)
	page := container.NewBorder(banner,nil, nil, nil, scroller)

	return page
}
