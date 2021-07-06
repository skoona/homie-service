package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/skoona/homie-service/pkg/deviceCore"
	cc "github.com/skoona/homie-service/pkg/utils"
	"strings"
	"time"
)

func SknBroadcastList(data *[]deviceCore.Broadcast) *widget.List {
	list := widget.NewList(
		func() int {
			return len(*data)
		},
		func() fyne.CanvasObject {
			return  SknNewBroadcastCards((*data)[0])
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

func SknDeviceList(data []deviceCore.Device) *widget.List {
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return  SknNewDeviceCard((data)[0])
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			var title *canvas.Text
			var version *canvas.Text
			var subTitle *canvas.Text
			var body *canvas.Text
			var img  *canvas.Image
			var d = (data)[i]

			img = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[1].(*canvas.Image)
			img.SetMinSize(fyne.NewSize(64,64))
			if strings.Contains(d.Attrs["$state"].Value, "ready") {
				img.Resource = cc.SknSelectThemedImage("sensorOn_r").Resource
			} else {
				img.Resource = cc.SknSelectThemedImage("sensorOff_r").Resource
			}
			//                   VBox                         Border                  Image / VBox                 //Text
			title    = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*canvas.Text)
			version  = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*fyne.Container).Objects[1].(*canvas.Text)
			subTitle = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*fyne.Container).Objects[2].(*canvas.Text)
			body     = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*fyne.Container).Objects[3].(*canvas.Text)

			title.Text    = d.Name
			version.Text  = "Signal: " + d.Attrs["$stats"].Props["signal"].Value
			subTitle.Text = "firmware: " + d.Attrs["$fw"].Props["name"].Value
			body.Text     = "ip: " +d.Attrs["$localip"].Value
		})

	return list
}