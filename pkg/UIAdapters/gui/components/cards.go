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
	"image/color"
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

	textGrid := MakeDeviceTextGrid(dv)
	box := container.NewVScroll(textGrid)

	content := container.NewVBox(sensor, wTitle, wSubTitle, box)

	return content
}

func scratchTable(dv dc.Device, logger log.Logger) {
	devList := unWrapDeviceForTables(dv)

	table := widget.NewTable(
		func() (int, int) {
			dList := devList
			_ = logger.Log("callback","CountRequest", "Rows",len(dList), "Cols", 2, "from", dv.Name)
			return len(dList), 2 // answer row and col count
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("content") // object of same size
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			dList := devList
			// i.Row, i.Col
			_ = logger.Log("callback", "DataRequest", "Row",i.Row ,"Col",i.Col, "from", dv.Name)
			o.(*widget.Label).SetText( dList[fmt.Sprintf("%d:%d", i.Row, i.Col)]) // set value of row / col
		})
	//table.Resize(fyne.NewSize(240,460))
	table.Refresh()
}

// unWrapDeviceForTables unwrapp device into rc map
func unWrapDeviceForTables(d dc.Device) map[string]string {
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

// unWrapDeviceForTextGrid unwrap device into k/v map
func unWrapDeviceForTextGrid(d *dc.Device) map[string]string {
	var dlist = map[string]string{}
	fmt.Printf("[UnWrapping]{%s} from: %s ----------------------\n", d.Name, d.Parent)
	row := 0
	dlist[fmt.Sprintf("%d:0", row)] = d.Name
	dlist[fmt.Sprintf("%d:1", row)] = d.Parent
	row += 1
	dlist[fmt.Sprintf("%d:0", row)] = "Device"
	dlist[fmt.Sprintf("%d:1", row)] = "Device Attributes"
	row += 1
	for _, v := range d.Attrs {
		fmt.Printf("[DeviceAttrs]{%d} Key: %s, Value: %s\n", row, v.Name, v.Value)
		dlist[fmt.Sprintf("%d:0", row)] = v.Name
		if v.Value != "" {
			dlist[fmt.Sprintf("%d:1", row)] = v.Value
			row += 1
		} else {
			dlist[fmt.Sprintf("%d:1", row)] = "Device Attr Properties"
			row += 1
			for _, p := range v.Props {
				fmt.Printf("[Props]{%d} Key: %s, Value: %s\n", row, p.Name, p.Value)
				dlist[fmt.Sprintf("%d:0", row)] = p.Name
				if p.Value != "" {
					dlist[fmt.Sprintf("%d:1", row)] = p.Value
					row += 1
				} else {
					dlist[fmt.Sprintf("%d:1", row)] = "Device Attr Prop Properties"
					row += 1
					for _, pp := range p.Props {
						fmt.Printf("[PProps]{%d} Key: %s, Value: %s\n", row, pp.Name, pp.Value)
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
			fmt.Printf("[Attrs]{%d} Key: %s, Value: %s\n", row, v.Name, v.Value)
			dlist[fmt.Sprintf("%d:0", row)] = v.Name
			dlist[fmt.Sprintf("%d:1", row)] = v.Value
			row += 1
		}
		dlist[fmt.Sprintf("%d:0", row)] = node
		dlist[fmt.Sprintf("%d:1", row)] = "Node Properties"
		row += 1
		for _, v := range d.Nodes[node].Props {
			fmt.Printf("[Props]{%d} Key: %s, Value: %s\n", row, v.Name, v.Value)
			dlist[fmt.Sprintf("%d:0", row)] = v.Name
			if v.Value != "" {
				dlist[fmt.Sprintf("%d:1", row)] = v.Value
				row += 1
			} else {
				dlist[fmt.Sprintf("%d:1", row)] = "Node Prop Attributes"
				row += 1
				for _, pa := range v.Attrs {
					fmt.Printf("[PAttrs]{%d} Key: %s, Value: %s\n", row, pa.Name, pa.Value)
					dlist[fmt.Sprintf("%d:0", row)] = pa.Name
					dlist[fmt.Sprintf("%d:1", row)] = pa.Value
					row += 1
				}
			}
		}
	}

	return dlist
}


func makeTextGrid() *widget.TextGrid {
	grid := widget.NewTextGridFromString("TextGrid\n\tContent\nZebra")
	grid.SetStyleRange(0, 4, 0, 7,
		&widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 64, B: 192, A: 128}})
	grid.SetRowStyle(1, &widget.CustomTextGridStyle{BGColor: &color.NRGBA{R: 64, G: 192, B: 64, A: 128}})

	white := &widget.CustomTextGridStyle{FGColor: color.White, BGColor: color.Black}
	black := &widget.CustomTextGridStyle{FGColor: color.Black, BGColor: color.White}
	grid.Rows[2].Cells[0].Style = white
	grid.Rows[2].Cells[1].Style = black
	grid.Rows[2].Cells[2].Style = white
	grid.Rows[2].Cells[3].Style = black
	grid.Rows[2].Cells[4].Style = white

	grid.ShowLineNumbers = true
	grid.ShowWhitespace = true

	return grid
}


func MakeDeviceTextGrid(d *dc.Device) *widget.TextGrid {
	dList := unWrapDeviceForTextGrid(d)
	rowLength := len(dList)
	buffer := strings.Builder{}
	for row := 0; row < rowLength; row++ {
		fmt.Fprintf(&buffer, "%s\t%s\n",
			dList[fmt.Sprintf("%d:0", row)],
			dList[fmt.Sprintf("%d:1", row)])
	}
	grid := widget.NewTextGridFromString(buffer.String())

	fmt.Println("TextBuffer Contents", buffer.String())

	return grid
}
