package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// FirmwaresTab ui page which includes firmwares
// the is to be used on main page
func FirmwaresTab() fyne.CanvasObject {
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func(){}),
	)
	list := container.NewVBox()
	list.Add(widget.NewLabel("label 1"))
	list.Add(widget.NewLabel("label 2"))
	side := container.New(layout.NewBorderLayout(bar, nil, nil, nil), bar, list)

	card := widget.NewCard("HomeOffice", "80%", widget.NewLabel("motion"))
	card.SetImage( canvas.NewImageFromResource(theme.DocumentIcon()))
	card1 := widget.NewCard("MediaRoom", "60%", widget.NewLabel("motion"))
	card1.SetImage( canvas.NewImageFromResource(theme.DocumentIcon()))
	card2 := widget.NewCard("GuestRoom", "49%", widget.NewLabel("motion"))
	card2.SetImage( canvas.NewImageFromResource(theme.DocumentIcon()))
	card3 := widget.NewCard("Garage", "49%", widget.NewLabel("motion"))
	card3.SetImage( canvas.NewImageFromResource(theme.DocumentIcon()))
	card4 := widget.NewCard("FrontDoor", "49%", widget.NewLabel("motion"))
	card4.SetImage( canvas.NewImageFromResource(theme.DocumentIcon()))
	card5 := widget.NewCard("PatioDoor", "49%", widget.NewLabel("motion"))
	card5.SetImage( canvas.NewImageFromResource(theme.DocumentIcon()))
	card6 := widget.NewCard("FamilyRom", "49%", widget.NewLabel("motion"))
	card6.SetImage( canvas.NewImageFromResource(theme.DocumentIcon()))
	card7 := widget.NewCard("MechanicalRoom", "49%", widget.NewLabel("motion"))
	card7.SetImage( canvas.NewImageFromResource(theme.DocumentIcon()))
	content := container.New(layout.NewGridLayout(5), card, card1, card2, card3, card4, card5, card6, card7)

	scroller := container.NewVScroll(content)
	split := container.NewHSplit(side, scroller)
	split.Offset = 0.2

	return split
}
