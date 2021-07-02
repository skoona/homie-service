package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"net/url"
)


// [![works with MQTT Homie](https://homieiot.github.io/img/works-with-homie.svg "works with MQTT Homie")](https://homieiot.github.io/)
// [![works with MQTT Homie]
//       (https://homieiot.github.io/img/works-with-homie.svg "works with MQTT Homie")
//  ](https://homieiot.github.io/)

// HomeTab ui page which includes broadcasts
// this is to be used as the main page
func (vp *viewProvider) HomeTab() fyne.CanvasObject {

	card := widget.NewCard("LWT", "Homie Monitor Offline", widget.NewLabel("Received: 2021-06-30T18:43:15.067335"))
		card.SetImage(SknCanvasSVGImageFromPath("./docs/notificationAlert-mbo-24px.svg"))

	card1 := widget.NewCard("Alert", "Go Lang is wonderful", widget.NewLabel("Received: 2021-06-30T18:44:48.306133"))
		card1.SetImage(SknCanvasSVGImageFromPath("./docs/notificationAlert-mbr-24px.svg"))

	scroller := container.NewVScroll(
		container.New(layout.NewGridLayout(3), card, card1))


	image := SknLoadImageFromPath("./works-with-homie.svg")
	uri, _ := url.Parse("https://homieiot.github.io/")
	banner := container.NewPadded(image, widget.NewHyperlink("Homie for ESP8266/ESP32", uri))

	content := container.NewBorder(banner,nil, nil, nil, scroller)

	return content
}
