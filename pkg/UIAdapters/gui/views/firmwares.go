package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// FirmwaresTab ui page which includes firmwares
// the is to be used on main page
func (vp *viewProvider) FirmwaresTab() fyne.CanvasObject{
	clk := SknLoadImageFromPath("./docs/wall-clock-svgrepo-com.svg")
	clk.Resize(fyne.NewSize(48,32))
	rect := canvas.NewRectangle(Green)
	rect.SetMinSize(fyne.NewSize(64,48))

	side := container.New(layout.NewPaddedLayout(), rect, clk)

	card := widget.NewCard("HomeOffice", "80%", widget.NewLabel("motion"))
		card.SetImage( SknCanvasSVGImageFromPath("./docs/insertDriveFile-mdr-24px.svg")) // canvas.NewImageFromResource(theme.StorageIcon()))
	card1 := widget.NewCard("MediaRoom", "60%", widget.NewLabel("motion"))
		card1.SetImage( SknCanvasSVGImageFromPath("./docs/insertDriveFile-mdr-24px.svg")) // canvas.NewImageFromResource(theme.StorageIcon()))
	card2 := widget.NewCard("GuestRoom", "49%", widget.NewLabel("motion"))
		card2.SetImage( SknCanvasSVGImageFromPath("./docs/insertDriveFile-mdr-24px.svg")) // canvas.NewImageFromResource(theme.StorageIcon()))
	card3 := widget.NewCard("Garage", "49%", widget.NewLabel("motion"))
		card3.SetImage( SknCanvasSVGImageFromPath("./docs/insertDriveFile-mdr-24px.svg")) // canvas.NewImageFromResource(theme.StorageIcon()))
	card4 := widget.NewCard("FrontDoor", "49%", widget.NewLabel("motion"))
		card4.SetImage( SknCanvasSVGImageFromPath("./docs/insertDriveFile-mdr-24px.svg")) // canvas.NewImageFromResource(theme.StorageIcon()))
	card5 := widget.NewCard("PatioDoor", "49%", widget.NewLabel("motion"))
		card5.SetImage( SknCanvasSVGImageFromPath("./docs/insertDriveFile-mdr-24px.svg")) // canvas.NewImageFromResource(theme.StorageIcon()))
	card6 := widget.NewCard("FamilyRom", "49%", widget.NewLabel("motion"))
		card6.SetImage( SknCanvasSVGImageFromPath("./docs/insertDriveFile-mdr-24px.svg")) // canvas.NewImageFromResource(theme.StorageIcon()))
	card7 := widget.NewCard("MechanicalRoom", "49%", widget.NewLabel("motion"))
		card7.SetImage( SknCanvasSVGImageFromPath("./docs/insertDriveFile-mdr-24px.svg")) // canvas.NewImageFromResource(theme.StorageIcon()))
	content := container.New(layout.NewGridLayout(3), card, card1, card2, card3, card4, card5, card6, card7)

	scroller := container.NewVScroll(content)
	page := container.NewBorder(nil,nil, side, nil, scroller)

	return page
}
