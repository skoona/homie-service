package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	cc "github.com/skoona/homie-service/pkg/utils"
	"strings"
	"time"
)

// containers for GridLayout


func SknNewSiteCards() fyne.CanvasObject {
	return nil
}


func SknNewDeviceCard(d dc.Device) *fyne.Container {
	var sensor *canvas.Image

	if strings.Contains(d.Attrs["$state"].Value, "ready") {
		sensor = cc.SknSelectThemedImage("sensorOn_r")
	} else {
		sensor = cc.SknSelectThemedImage("sensorOff_r")
	}
	sensor.SetMinSize(fyne.NewSize(64,64))

	title := canvas.NewText(d.Name, theme.DefaultTheme().Color(theme.ColorNamePrimary, CurrentVariant))
	title.SetMinSize(fyne.NewSize(3 * theme.CaptionTextSize(),3 * theme.CaptionTextSize()))
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	version := canvas.NewText("Signal: " + d.Attrs["$stats"].Props["signal"].Value, theme.DefaultTheme().Color(theme.ColorNamePrimary, CurrentVariant))
	version.SetMinSize(fyne.NewSize(2 * theme.CaptionTextSize(), 2 * theme.CaptionTextSize()))
	version.TextStyle = fyne.TextStyle{Bold: true}
	version.Alignment = fyne.TextAlignCenter

	subTitle := canvas.NewText("firmware: " + d.Attrs["$fw"].Props["name"].Value, theme.DefaultTheme().Color(theme.ColorNamePrimary, CurrentVariant))
	subTitle.SetMinSize(fyne.NewSize(theme.CaptionTextSize(),theme.CaptionTextSize()))
	subTitle.TextStyle = fyne.TextStyle{Bold: true}
	subTitle.Alignment = fyne.TextAlignCenter

	body := canvas.NewText("ip: " +d.Attrs["$localip"].Value, theme.DefaultTheme().Color(theme.ColorNamePrimary, CurrentVariant))
	body.SetMinSize(fyne.NewSize(theme.TextSize(), theme.TextSize()))
	body.TextStyle = fyne.TextStyle{Monospace: true}
	body.Alignment = fyne.TextAlignCenter

	return container.NewBorder(nil, nil, sensor, nil, container.NewVBox(title, version, subTitle, body))
}

func SknNewFirmwareCards(fw dc.Firmware) fyne.CanvasObject {
	title := canvas.NewText(fw.Name, theme.DefaultTheme().Color(theme.ColorNamePrimary, CurrentVariant))
	title.SetMinSize(fyne.NewSize(theme.CaptionTextSize(),theme.CaptionTextSize()))
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	version := canvas.NewText(fw.Version, theme.DefaultTheme().Color(theme.ColorNamePrimary, CurrentVariant))
	version.SetMinSize(fyne.NewSize(theme.CaptionTextSize(),theme.CaptionTextSize()))
	version.TextStyle = fyne.TextStyle{Bold: true}
	version.Alignment = fyne.TextAlignCenter

	subTitle := canvas.NewText(fw.Brand, theme.DefaultTheme().Color(theme.ColorNamePrimary, CurrentVariant))
	subTitle.SetMinSize(fyne.NewSize(theme.CaptionTextSize(),theme.CaptionTextSize()))
	subTitle.TextStyle = fyne.TextStyle{Bold: true}
	subTitle.Alignment = fyne.TextAlignCenter

	body := canvas.NewText(fw.Created.Local().Format(time.RFC822), theme.DefaultTheme().Color(theme.ColorNamePrimary, CurrentVariant))
	body.SetMinSize(fyne.NewSize(theme.TextSize(), theme.TextSize()))
	body.TextStyle = fyne.TextStyle{Monospace: true}
	body.Alignment = fyne.TextAlignCenter

	img := cc.SknSelectThemedImage("insertDriveFile_r")
	img.SetMinSize(fyne.NewSize(64,64))

	return container.NewBorder(nil, nil, img, nil, container.NewVBox(title, version, subTitle, body))
}

func SknNewScheduleCards(s dc.Schedule) fyne.CanvasObject {
	sensor := cc.SknSelectThemedImage("wallClock")
	sensor.SetMinSize(fyne.NewSize(64,64))

	title := canvas.NewText("Device" + s.DeviceID, theme.DefaultTheme().Color(theme.ColorNamePrimary, CurrentVariant))
	title.SetMinSize(fyne.NewSize(theme.CaptionTextSize(),theme.CaptionTextSize()))
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	version := canvas.NewText("firmware: " + string(s.FirmwareID), theme.DefaultTheme().Color(theme.ColorNamePrimary, CurrentVariant))
	version.SetMinSize(fyne.NewSize(theme.CaptionTextSize(),theme.CaptionTextSize()))
	version.TextStyle = fyne.TextStyle{Bold: true}
	version.Alignment = fyne.TextAlignCenter

	subTitle := canvas.NewText("Status: " + s.Status, theme.DefaultTheme().Color(theme.ColorNamePrimary, CurrentVariant))
	subTitle.SetMinSize(fyne.NewSize(theme.CaptionTextSize(),theme.CaptionTextSize()))
	subTitle.TextStyle = fyne.TextStyle{Bold: true}
	subTitle.Alignment = fyne.TextAlignCenter

	body := canvas.NewText("Completed: " +s.Scheduled.Local().Format(time.RFC822), theme.DefaultTheme().Color(theme.ColorNamePrimary, CurrentVariant))
	body.SetMinSize(fyne.NewSize(theme.TextSize(), theme.TextSize()))
	body.TextStyle = fyne.TextStyle{Monospace: true}
	body.Alignment = fyne.TextAlignCenter

	return container.NewBorder(nil, nil, sensor, nil, container.NewVBox(title, version, subTitle, body))
}

func SknNewBroadcastCards(bc dc.Broadcast) fyne.CanvasObject {
	title := canvas.NewText(bc.Level, theme.DefaultTheme().Color(theme.ColorNamePrimary,  CurrentVariant))
	title.SetMinSize(fyne.NewSize(2 * theme.CaptionTextSize(), 2 * theme.CaptionTextSize()))
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignLeading

	subTitle := canvas.NewText(bc.Value, theme.DefaultTheme().Color(theme.ColorNamePrimary, CurrentVariant))
	subTitle.SetMinSize(fyne.NewSize(theme.CaptionTextSize(),theme.CaptionTextSize()))
	subTitle.TextStyle = fyne.TextStyle{Bold: true}
	subTitle.Alignment = fyne.TextAlignLeading

	body := canvas.NewText(bc.Received.Local().Format(time.RFC822), cc.Blue)
	body.SetMinSize(fyne.NewSize(theme.TextSize(), theme.TextSize()))
	body.TextStyle = fyne.TextStyle{Monospace: true}
	body.Alignment = fyne.TextAlignLeading

	img := cc.SknSelectThemedImage("notificationAlert_r")
	img.SetMinSize(fyne.NewSize(64,64))

	return container.NewBorder(nil, nil, img, nil, container.NewVBox(title, subTitle, body))
}
