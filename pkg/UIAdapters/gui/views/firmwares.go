package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/components"
	cc "github.com/skoona/homie-service/pkg/utils"
	"time"
)

// FirmwaresTab ui page which includes firmwares
// the is to be used on main page
func (vp *viewProvider) FirmwaresTab() fyne.CanvasObject{
	form := components.SknFirmwareSideForm()
	side := container.New(layout.NewPaddedLayout(), form)

	cards := container.New(layout.NewGridLayout(2))
	scroller := container.NewVScroll(cards)


	page := container.NewBorder(nil,nil, side, nil, scroller)

	for _, fw := range (*vp.dSvc).AllFirmwares() {
		title := canvas.NewText(fw.Name, theme.DefaultTheme().Color(theme.ColorNamePrimary, 2))
		title.SetMinSize(fyne.NewSize(theme.CaptionTextSize(),theme.CaptionTextSize()))
		title.TextStyle = fyne.TextStyle{Bold: true}
		title.Alignment = fyne.TextAlignCenter

		version := canvas.NewText(fw.Version, theme.DefaultTheme().Color(theme.ColorNamePrimary, 2))
		version.SetMinSize(fyne.NewSize(theme.CaptionTextSize(),theme.CaptionTextSize()))
		version.TextStyle = fyne.TextStyle{Bold: true}
		version.Alignment = fyne.TextAlignCenter

		subTitle := canvas.NewText(fw.Brand, theme.DefaultTheme().Color(theme.ColorNamePrimary, 2))
		subTitle.SetMinSize(fyne.NewSize(theme.CaptionTextSize(),theme.CaptionTextSize()))
		subTitle.TextStyle = fyne.TextStyle{Bold: true}
		subTitle.Alignment = fyne.TextAlignCenter

		body := canvas.NewText(fw.Created.Local().Format(time.RFC822), theme.DefaultTheme().Color(theme.ColorNamePrimary, 2))
		body.SetMinSize(fyne.NewSize(theme.TextSize(), theme.TextSize()))
		body.TextStyle = fyne.TextStyle{Monospace: true}
		body.Alignment = fyne.TextAlignLeading

		img := cc.SknSelectThemedImage("insertDriveFile_r")
		img.SetMinSize(fyne.NewSize(64,64))

		pkg := container.NewBorder(nil, nil, img, nil, container.NewVBox(title, version, subTitle, body))
		card := container.NewVBox(pkg)

		cards.Add(card)
	}

	vp.firmwareSide = &form
	vp.firmwareCards = cards

	return page
}
