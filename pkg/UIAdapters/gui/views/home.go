package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	cc "github.com/skoona/homie-service/pkg/utils"
	"net/url"
	"time"
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

	cards := container.New(layout.NewGridLayout(2))
	scroller := container.NewVScroll(cards)
	page := container.NewBorder(top,nil, nil, nil, scroller)

	for _, bc := range (*vp.dSvc).AllBroadcasts() {
		title := canvas.NewText(bc.Level, theme.DefaultTheme().Color(theme.ColorNamePrimary, 2))
		title.SetMinSize(fyne.NewSize(2 * theme.CaptionTextSize(), 2 * theme.CaptionTextSize()))
		title.TextStyle = fyne.TextStyle{Bold: true}
		title.Alignment = fyne.TextAlignLeading

		subTitle := canvas.NewText(bc.Value, theme.DefaultTheme().Color(theme.ColorNamePrimary, 2))
		subTitle.SetMinSize(fyne.NewSize(theme.CaptionTextSize(),theme.CaptionTextSize()))
		subTitle.TextStyle = fyne.TextStyle{Bold: true}
		subTitle.Alignment = fyne.TextAlignLeading

		body := canvas.NewText(bc.Received.Local().Format(time.RFC822), cc.Blue)
		body.SetMinSize(fyne.NewSize(theme.TextSize(), theme.TextSize()))
		body.TextStyle = fyne.TextStyle{Monospace: true}
		body.Alignment = fyne.TextAlignLeading

		img := cc.SknSelectThemedImage("notificationAlert_o")
		img.SetMinSize(fyne.NewSize(64,64))
		//pad := container.NewPadded( img )

		pkg := container.NewBorder(nil, nil, img, nil, container.NewVBox(title, subTitle, body))
		card := container.NewVBox(pkg)

		cards.Add(card)
	}

	vp.homeCards = cards // retain ref for later updates

	return page
}
