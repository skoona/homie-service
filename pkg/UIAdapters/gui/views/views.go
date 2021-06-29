package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func DefaultContent() fyne.CanvasObject {
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func(){}),
	)
	list := container.NewVBox()
	list.Add(widget.NewLabel("label 1"))
	list.Add(widget.NewLabel("label 2"))
	side := container.New(layout.NewBorderLayout(bar, nil, nil, nil), bar, list)
	//content :=  container.New(layout.NewBorderLayout(bar, nil, nil, nil), bar, u.list)
	card := widget.NewCard("HomeOffice", "80%", widget.NewLabel("motion"))
		card.SetImage(&canvas.Image{Resource: theme.ComputerIcon()})
	card1 := widget.NewCard("MediaRoom", "60%", widget.NewLabel("motion"))
		card1.SetImage(&canvas.Image{Resource: theme.ComputerIcon()})
	card2 := widget.NewCard("GuestRoom", "49%", widget.NewLabel("motion"))
		card2.SetImage(&canvas.Image{Resource: theme.ComputerIcon()})
	content := container.New(layout.NewAdaptiveGridLayout(8), card, card1, card2)
	split := container.NewHSplit(side, content)
	split.Offset = 0.333

	status := container.NewVSplit(split, widget.NewLabel("Status"))
	status.Offset = 0.95

	return status
}
