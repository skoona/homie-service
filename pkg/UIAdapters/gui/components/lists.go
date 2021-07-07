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
			var bc = (*data)[i]

			//                   Border                  Image / VBox                 Text
			title    = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*canvas.Text)
			subTitle = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[1].(*canvas.Text)
			body     = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[2].(*canvas.Text)

			title.Text = "level: " +bc.Level
			subTitle.Text = "msg:" + bc.Value
			body.Text = "received: " + bc.Received.Local().Format(time.RFC822)
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

			img = o.(*fyne.Container).Objects[1].(*canvas.Image)
			img.SetMinSize(fyne.NewSize(96,96))
			if strings.Contains(d.Attrs["$state"].Value, "ready") {
				img.Resource = cc.SknSelectThemedImage("sensorOn_r").Resource
			} else {
				img.Resource = cc.SknSelectThemedImage("sensorOff_r").Resource
			}
			//                  Border                  Image / VBox                 //Text
			title    = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*canvas.Text)
			version  = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[1].(*canvas.Text)
			subTitle = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[2].(*canvas.Text)
			body     = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[3].(*canvas.Text)

			title.Text    = d.Name
			version.Text  = "Signal: " + d.Attrs["$stats"].Props["signal"].Value
			subTitle.Text = "fw: " + d.Attrs["$fw"].Props["name"].Value
			body.Text     = "ip: " +d.Attrs["$localip"].Value
		})

	return list
}

func SknFirmwareList(data *[]deviceCore.Firmware) *widget.List {
	list := widget.NewList(
		func() int {
			return len((*data))
		},
		func() fyne.CanvasObject {
			return  SknNewFirmwareCards((*data)[0])
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			var title *canvas.Text
			var version *canvas.Text
			var subTitle *canvas.Text
			var body *canvas.Text
			var fw = (*data)[i]

			//                  Border                  Image / VBox                 Text
			title    = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*canvas.Text)
			version  = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[1].(*canvas.Text)
			subTitle = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[2].(*canvas.Text)
			body     = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[3].(*canvas.Text)

			title.Text = fw.Name
			version.Text = "version: " + fw.Version
			subTitle.Text = cc.HumanFormatFromBytes( uint64(fw.Size))
			body.Text = fw.Created.Local().Format(time.RFC822)
		})

	return list
}

func SknScheduleList(data []deviceCore.Schedule) *widget.List {
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return  SknNewScheduleCards(data[0])
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			var title *canvas.Text
			var version *canvas.Text
			var subTitle *canvas.Text
			var body *canvas.Text
			var img  *canvas.Image
			var sc = data[i]


			//                  Border                  Image / VBox                 Text
			img      = o.(*fyne.Container).Objects[1].(*canvas.Image)
			title    = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[0].(*canvas.Text)
			version  = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[1].(*canvas.Text)
			subTitle = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[2].(*canvas.Text)
			body     = o.(*fyne.Container).Objects[0].(*fyne.Container).Objects[3].(*canvas.Text)

			if strings.Contains(sc.State, "pending") { // waiting  -- complete
				img.Resource = cc.SknSelectThemedImage("wallClock").Resource
			} else if strings.Contains(sc.State, "complete") {
				img.Resource = cc.SknSelectThemedImage("thumbsUp_o").Resource
			} else {
				img.Resource = cc.SknSelectThemedImage("timeLapse_r").Resource
			}

			title.Text = "Device: " + sc.DeviceID
			version.Text = "Firmware: " + string(sc.FirmwareID)
			subTitle.Text = "State: " + sc.State + "->" + sc.Status
			body.Text = "Scheduled: " + sc.Scheduled.Local().Format(time.RFC822)
		})

	return list
}
