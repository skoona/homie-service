package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/components"
	"github.com/skoona/homie-service/pkg/deviceCore"
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

	listBox := SknBroadcastList(&vp.siteNetworks.Broadcasts)
	listBox.OnSelected = vp.OnBroadcastSelected
	page := container.NewBorder(top,nil, nil, nil, listBox)
	vp.homeCards = listBox // retain ref for later updates

	return page
}

func SknBroadcastList(data *[]deviceCore.Broadcast) *widget.List {
	list := widget.NewList(
	func() int {
		return len(*data)
	},
	func() fyne.CanvasObject {
		return  components.SknNewBroadcastCards((*data)[0])
	},
	func(i widget.ListItemID, o fyne.CanvasObject) {
		var title *canvas.Text
		var subTitle *canvas.Text
		var body *canvas.Text
		var img  *canvas.Image
		var bc = (*data)[i]

		//                   VBox                         Border                  Image / VBox                 //Text
		img      = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[1].(*canvas.Image)
		title    = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*canvas.Text)
		subTitle = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*fyne.Container).Objects[1].(*canvas.Text)
		body     = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*fyne.Container).Objects[2].(*canvas.Text)

		img = cc.SknSelectThemedImage("notificationAlert_o")
		img.SetMinSize(fyne.NewSize(64,64))

		title.Text = bc.Level
		subTitle.Text = bc.Value
		body.Text = bc.Received.Local().Format(time.RFC822)
			//SetText((*data)[i].Level)
	})

	return list
}