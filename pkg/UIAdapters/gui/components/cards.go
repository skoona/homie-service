package components

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/go-kit/kit/log"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	cc "github.com/skoona/homie-service/pkg/utils"
	"strings"
)

// containers for GridLayout

func NewDevice(dv dc.Device, logger log.Logger) fyne.CanvasObject {
	var sensor *canvas.Image
	devList := unWrapDevice(dv)

	if dv.Attrs["$state"].Value == "ready" {
		sensor = cc.SknSelectThemedImage("sensorOn_r")
	} else {
		sensor = cc.SknSelectThemedImage("sensorOff_r")
	}
	sensor.SetMinSize(fyne.NewSize(128,128))

	table := widget.NewTable(
		func() (int, int) {
			dList := devList
			_ = logger.Log("callback","CountRequest", "Rows",len(dList), "Cols", 2, "from", dv.Name)
			return len(dList), 2 // answer row and col count
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content") // object of same size
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			dList := devList
			// i.Row, i.Col
			_ = logger.Log("callback", "DataRequest", "Row",i.Row ,"Col",i.Col, "from", dv.Name)
			o.(*widget.Label).SetText( dList[fmt.Sprintf("%d:%d", i.Row, i.Col)]) // set value of row / col
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
	fmt.Printf("[UnWrapping]{%s} from: %s ----------------------\n", d.Name, d.Parent)
	var row =1
	for _, v := range d.Attrs {
		fmt.Printf("[DeviceAttrs]{%d} Key: %s, Value: %s\n", row, v.Name, v.Value)
		dlist[fmt.Sprintf("%d:0", row)] = v.Name
		dlist[fmt.Sprintf("%d:1", row)] = v.Value
		row += 1
	}
	nodes := strings.Split(d.Attrs["$nodes"].Value, ",")
	for _, node := range nodes {
		for _, v := range d.Nodes[node].Attrs {
			fmt.Printf("[Attrs]{%d} Key: %s, Value: %s\n", row, v.Name, v.Value)
			dlist[fmt.Sprintf("%d:0", row)] = v.Name
			dlist[fmt.Sprintf("%d:1", row)] = v.Value
			row += 1
		}
		for _, v := range d.Nodes[node].Props {
			fmt.Printf("[Props]{%d} Key: %s, Value: %s\n", row, v.Name, v.Value)
			dlist[fmt.Sprintf("%d:0", row)] = v.Name
			dlist[fmt.Sprintf("%d:1", row)] = v.Value
			row += 1
		}
	}

	return dlist
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
