package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// BroadcastsTab ui page which includes broadcasts
// this is to be used as the main page
func BroadcastsTab() fyne.CanvasObject {
	badge := SknLoadImageFromPath("./docs/homie-banner.png").(*canvas.Image)

	card := widget.NewCard("LWT", "Homie Monitor Offline", widget.NewLabel("Received: 2021-06-30T18:43:15.067335"))
		card.SetImage(badge)

	card1 := widget.NewCard("Alert", "Go Lang is wonderful", widget.NewLabel("Received: 2021-06-30T18:44:48.306133"))
		card1.SetImage(badge)

	scroller := container.NewVScroll(
		container.New(layout.NewGridLayout(5), card, card1))


	image := SknLoadImageFromPath( "./docs/WeMos_D1_Esp_Wroom-Mini_WiFi_Nodemcu.png")
	image.Resize(fyne.NewSize(360,120))
	banner := container.NewPadded(image)

	content := container.NewBorder(banner,nil, nil, nil, scroller)

	return content
}
