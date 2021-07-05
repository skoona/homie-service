package components

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-kit/kit/log"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	cc "github.com/skoona/homie-service/pkg/utils"
	"strings"
)

// containers for GridLayout

func SknNewDeviceBadge(dv *dc.Device, logger log.Logger) fyne.CanvasObject {
	var sensor *canvas.Image

	if dv.Attrs["$state"].Value == "ready" {
		sensor = cc.SknSelectThemedImage("sensorOn_r")
	} else {
		sensor = cc.SknSelectThemedImage("sensorOff_r")
	}
	sensor.SetMinSize(fyne.NewSize(96,96))
	sensor.Refresh()

	wTitle := widget.NewLabel(dv.Name)
	wTitle.Alignment = fyne.TextAlignCenter
	wSubTitle := widget.NewLabel(dv.Parent)
	wSubTitle.Alignment = fyne.TextAlignCenter

	box := deviceTable(dv, logger)
	//box := SknDeviceTableBuilder(dv, logger)

	content := widget.NewCard(dv.Name, dv.Parent, box)
	content.SetImage( sensor)

	return content
}

func SknDeviceTableBuilder(dv *dc.Device, logger log.Logger) *container.Scroll {
	devList := unWrapDevice(dv)
	mapRows := len(devList)

	tableGrid := container.NewGridWithColumns(2)
	key := ""
	val := ""
	for row := 0; row < mapRows; row++ {
		key = fmt.Sprintf("%d:0")
		val = fmt.Sprintf("%d:1")
		wKey := widget.NewLabel(devList[key])
		wVal := widget.NewLabel(devList[val])
		//if strings.Contains(devList[val], "Attribute") ||
		//	strings.Contains(devList[val], "Property") {
		//	wKey.Color = *cc.Purple
		//	wVal.Color = *cc.Purple
		//}
		tableGrid.Add(wKey)
		tableGrid.Add(wVal)
	}

	scroller := container.NewVScroll( tableGrid )
	scroller.Refresh()
	return scroller
}


func deviceTable(dv *dc.Device, logger log.Logger) *widget.Table {
	devList := unWrapDevice(dv)
	mapRows := len(devList)

	table := widget.NewTable(
		func() (int, int) {
			rows := mapRows
			_ = logger.Log("callback","CountRequest", "Rows", rows, "Cols", 2, "from", dv.Name)
			return rows, 2 // answer row and col count
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("dev attributes") // object of same size
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			dList := devList
			// i.Row, i.Col
			value := dList[fmt.Sprintf("%d:%d", i.Row, i.Col)]
			_ = logger.Log("callback", "DataRequest", "Row",i.Row ,"Col",i.Col, "from", dv.Name, "value", value)
			o.(*widget.Label).SetText( value ) // set value of row / col
		})

	return table
}

func MakeTableTab() fyne.CanvasObject {
	t := widget.NewTable(
		func() (int, int) { return 8, 2 },
		func() fyne.CanvasObject {
			return widget.NewLabel("Cell 000, 000")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			switch id.Col {
			case 0:
				label.SetText(fmt.Sprintf("%d", id.Row+1))
			case 1:
				label.SetText("A longer cell")
			default:
				label.SetText(fmt.Sprintf("Cell %d, %d", id.Row+1, id.Col+1))
			}
		})
	t.SetColumnWidth(0, 34)
	t.SetColumnWidth(1, 102)
	return t
}

// unWrapDevice unwrap device into k/v map
func unWrapDevice(d *dc.Device) map[string]string {
	var dlist = map[string]string{}
	row := 0
	dlist[fmt.Sprintf("%d:0", row)] = d.Name
	dlist[fmt.Sprintf("%d:1", row)] = d.Parent
	row += 1
	dlist[fmt.Sprintf("%d:0", row)] = "Device"
	dlist[fmt.Sprintf("%d:1", row)] = "Device Attributes"
	row += 1
	for _, v := range d.Attrs {
		dlist[fmt.Sprintf("%d:0", row)] = v.Name
		if len(v.Props) == 0 {
			dlist[fmt.Sprintf("%d:1", row)] = v.Value
			row += 1
		} else {
			dlist[fmt.Sprintf("%d:1", row)] = v.Value
			row += 1
			dlist[fmt.Sprintf("%d:0", row)] = d.Name
			dlist[fmt.Sprintf("%d:1", row)] = "Device Attr Properties"
			row += 1
			for _, p := range v.Props {
				dlist[fmt.Sprintf("%d:0", row)] = p.Name
				if len(p.Props) == 0 {
					dlist[fmt.Sprintf("%d:1", row)] = p.Value
					row += 1
				} else {
					dlist[fmt.Sprintf("%d:1", row)] = p.Value
					row += 1
					dlist[fmt.Sprintf("%d:0", row)] = d.Name
					dlist[fmt.Sprintf("%d:1", row)] = "Device Attr Prop Properties"
					row += 1
					for _, pp := range p.Props {
						dlist[fmt.Sprintf("%d:0", row)] = pp.Name
						dlist[fmt.Sprintf("%d:1", row)] = pp.Value
						row += 1
					}
				}
			}
		}
	}
	nodes := strings.Split(d.Attrs["$nodes"].Value, ",")
	for _, node := range nodes {
		dlist[fmt.Sprintf("%d:0", row)] = node
		dlist[fmt.Sprintf("%d:1", row)] = "Node Attributes"
		row += 1
		for _, v := range d.Nodes[node].Attrs {
			dlist[fmt.Sprintf("%d:0", row)] = v.Name
			dlist[fmt.Sprintf("%d:1", row)] = v.Value
			row += 1
		}
		dlist[fmt.Sprintf("%d:0", row)] = node
		dlist[fmt.Sprintf("%d:1", row)] = "Node Properties"
		row += 1
		for _, v := range d.Nodes[node].Props {
			dlist[fmt.Sprintf("%d:0", row)] = v.Name
			if len(v.Attrs) == 0 {
				dlist[fmt.Sprintf("%d:1", row)] = v.Value
				row += 1
			} else {
				dlist[fmt.Sprintf("%d:1", row)] = v.Value
				row += 1
				dlist[fmt.Sprintf("%d:0", row)] = node
				dlist[fmt.Sprintf("%d:1", row)] = "Node Prop Attributes"
				row += 1
				for _, pa := range v.Attrs {
					dlist[fmt.Sprintf("%d:0", row)] = pa.Name
					dlist[fmt.Sprintf("%d:1", row)] = pa.Value
					row += 1
				}
			}
		}
	}

	return dlist
}
