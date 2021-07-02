package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// SchedulesTab ui page which includes schedules
// the is to be used on main page
func SchedulesTab() fyne.CanvasObject {
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func(){}),
	)
	list := container.NewVBox()
	list.Add(widget.NewLabel("label 1"))
	list.Add(widget.NewLabel("label 2"))
	side := container.New(layout.NewBorderLayout(bar, nil, nil, nil), bar, list)

	card := widget.NewCard("HomeOffice", "80%", widget.NewLabel("motion"))
		card.SetImage( SknCanvasSVGImageFromPath("./docs/toggleOff-mbr-24px.svg"))
	card1 := widget.NewCard("MediaRoom", "60%", widget.NewLabel("motion"))
		card1.SetImage(SknCanvasSVGImageFromPath("./docs/toggleOn-mbr-24px.svg"))
	card2 := widget.NewCard("GuestRoom", "49%", widget.NewLabel("motion"))
		card2.SetImage(SknCanvasSVGImageFromPath("./docs/toggleOff-mbo-24px.svg"))
	card3 := widget.NewCard("Garage", "49%", widget.NewLabel("motion"))
		card3.SetImage(SknCanvasSVGImageFromPath("./docs/toggleOn-mbo-24px.svg"))
	card4 := widget.NewCard("FrontDoor", "49%", widget.NewLabel("motion"))
		card4.SetImage(SknCanvasSVGImageFromPath("./docs/timeLapse-mbr-24px.svg"))
	card5 := widget.NewCard("PatioDoor", "49%", widget.NewLabel("motion"))
		card5.SetImage(SknCanvasSVGImageFromPath("./docs/timeLapse-mbo-24px.svg"))
	card6 := widget.NewCard("FamilyRoom", "49%", widget.NewLabel("motion"))
		card6.SetImage(SknCanvasSVGImageFromPath("./docs/timeLapse-mbo-24px.svg"))
	card7 := widget.NewCard("MechanicalRoom", "49%", widget.NewLabel("motion"))
		card7.SetImage(SknCanvasSVGImageFromPath("./docs/timeLapse-mbo-24px.svg"))
	content := container.New(layout.NewGridLayout(5), card, card1, card2, card3, card4, card5, card6, card7)

	scroller := container.NewVScroll(content)
	split := container.NewHSplit(side, scroller)
	split.Offset = 0.2

	return split
}
