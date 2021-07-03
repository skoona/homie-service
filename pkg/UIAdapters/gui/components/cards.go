package components

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	cc "github.com/skoona/homie-service/pkg/utils"
)

// containers for GridLayout

func NewDevice(dv dc.Device) fyne.CanvasObject {
	var sensor *canvas.Image

	if dv.Attrs["$state"].Value == "ready" {
		sensor = cc.SknSelectThemedImage("sensorOn_r")
	} else {
		sensor = cc.SknSelectThemedImage("sensorOff_r")
	}
	sensor.SetMinSize(fyne.NewSize(128,128))

	table := widget.NewTable(
		func() (int, int) {
			dList := unWrapDevice(dv)
			return len(dList), 2 // answer row and col count
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content") // object of same size
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			// i.Row, i.Col
			dList := unWrapDevice(dv)
			o.(*widget.Label).SetText( dList[fmt.Sprintf("%s:%s", i.Row, i.Col)]) // set value of row / col
		})
	table.Refresh()


	//wTitle := widget.NewLabel(title)
	//wTitle.Alignment = fyne.TextAlignCenter
	//wTitle.Refresh()
	//wSubTitle := widget.NewLabel(subTitle)
	//wSubTitle.Alignment = fyne.TextAlignCenter
	//wSubTitle.Refresh()
	content := widget.NewCard(dv.Name, dv.Attrs["$localip"].Value, table)
	content.SetImage(sensor)
	content.Refresh()
	//content := container.NewWithoutLayout(sensor,wTitle, wSubTitle, table)

	return content
}

func unWrapDevice(d dc.Device) map[string]string {
	var dlist = map[string]string{}

	for name, node := range d.Nodes {
		row := 0
		dlist[fmt.Sprintf("%s:%s", row, 0)] = name
		dlist[fmt.Sprintf("%s:%s", row, 1)] = "Attributes"
		row = 1
		for _, v := range node.Attrs {
			dlist[fmt.Sprintf("%s:%s", row, 0)] = v.Name
			dlist[fmt.Sprintf("%s:%s", row, 1)] = v.Value
			row += 1
		}
		dlist[fmt.Sprintf("%s:%s", row, 0)] = name
		dlist[fmt.Sprintf("%s:%s", row, 1)] = "Properties"
		row += 1
		for _, v := range node.Props {
			dlist[fmt.Sprintf("%s:%s", row, 0)] = v.Name
			dlist[fmt.Sprintf("%s:%s", row, 1)] = v.Value
			row += 1
		}
	}

	return dlist
}