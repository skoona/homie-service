package views

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	cc "github.com/skoona/homie-service/pkg/utils"
	"image/color"
)

/*

*/

// sized svg icon
type SizedSVGIcon struct {
	widget.BaseWidget
	ThumbsUp *canvas.Image
	ThumbsDown *canvas.Image
	UpDown bool
}

func NewSizedSVGIcon() *SizedSVGIcon {
	var err error

	c := &SizedSVGIcon{}

	c.ThumbsUp.Resource, err = fyne.LoadResourceFromPath("./docs/thumbs-up-svgrepo-com.svg")
	if err != nil {
		fmt.Println("Failed to load thumbs up resource")
		return nil
	}
	c.ThumbsUp.FillMode = canvas.ImageFillContain

	c.ThumbsDown.Resource, err = fyne.LoadResourceFromPath("./docs/thumb-down-svgrepo-com.svg")
	if err != nil {
		fmt.Println("Failed to load thumbs down resource")
		return nil
	}
	c.ThumbsDown.FillMode = canvas.ImageFillContain

	c.UpDown = false

	c.ExtendBaseWidget(c)

	return c
}

func (c *SizedSVGIcon) Tapped(_ *fyne.PointEvent) {
	// do something here
	c.UpDown = !c.UpDown
	c.Refresh()
}

func (c *SizedSVGIcon) CreateRenderer() fyne.WidgetRenderer {
	r := &sizedSVGIconRender{svgIcon: c, img: &canvas.Image{}}
	r.updateImage()
	return r
}

type sizedSVGIconRender struct {
	svgIcon *SizedSVGIcon
	img   *canvas.Image
}

func (t *sizedSVGIconRender) BackgroundColor() color.Color {
	return color.Transparent
}

func (t *sizedSVGIconRender) Destroy() {
}

func (t *sizedSVGIconRender) Layout(_ fyne.Size) {
	t.img.Resize(t.MinSize())
}

func (t *sizedSVGIconRender) MinSize() fyne.Size {
	return fyne.NewSize(theme.IconInlineSize(), theme.IconInlineSize())
}

func (t *sizedSVGIconRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{t.img}
}

func (t *sizedSVGIconRender) Refresh() {
	t.updateImage()
}

func (t *sizedSVGIconRender) updateImage() {
	defer t.img.Refresh()

	if t.svgIcon.UpDown {
		t.img.Resource = theme.NewThemedResource(t.svgIcon.ThumbsUp.Resource)
	} else {
		t.img.Resource = theme.NewThemedResource(t.svgIcon.ThumbsDown.Resource)
	}
		//t.img.Resource = theme.CheckButtonIcon()
}
// end SVGIcon

// DefaultContent is a set of placeholder content while
// the main design is investigated
func DefaultContent() fyne.CanvasObject {
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func(){}),
	)
	list := container.NewVBox()
	list.Add(widget.NewLabel("label 1"))
	list.Add(widget.NewLabel("label 2"))
	side := container.New(layout.NewBorderLayout(bar, nil, nil, nil), bar, list)

	card := widget.NewCard("HomeOffice", "80%", widget.NewLabel("motion"))
		card.SetImage(SvgImageFromFile("./docs/calendar-schedule-svgrepo-com.svg"))
	card1 := widget.NewCard("MediaRoom", "60%", widget.NewLabel("motion"))
		card1.SetImage(SvgImageFromFile("./docs/calendar-svgrepo-com.svg"))
	card2 := widget.NewCard("GuestRoom", "49%", widget.NewLabel("motion"))
		card2.SetImage(SvgImageFromFile("./docs/power-button-svgrepo-com-2.svg"))
	card3 := widget.NewCard("Garage", "49%", widget.NewLabel("motion"))
		card3.SetImage(SvgImageFromFile("./docs/power-symbol-svgrepo-com.svg"))
	card4 := widget.NewCard("FrontDoor", "49%", widget.NewLabel("motion"))
		card4.SetImage(SvgImageFromFile("./docs/thumbs-up-svgrepo-com.svg"))
	card5 := widget.NewCard("PatioDoor", "49%", widget.NewLabel("motion"))
		card5.SetImage(SvgImageFromFile("./docs/thumb-down-svgrepo-com.svg"))
	card6 := widget.NewCard("FamilyRom", "49%", widget.NewLabel("motion"))
		card6.SetImage(SvgImageFromFile("./docs/clock-svgrepo-com.svg"))
	card7 := widget.NewCard("MechanicalRoom", "49%", widget.NewLabel("motion"))
		card7.SetImage(SvgImageFromFile("./docs/clock-timer-svgrepo-com.svg"))
	content := container.New(layout.NewGridLayout(5), card, card1, card2, card3, card4, card5, card6, card7)

	scroller := container.NewVScroll(content)
	split := container.NewHSplit(side, scroller)
	split.Offset = 0.2

	return split
}

func SvgImageFromFile(filename string) *canvas.Image  {

	//res, err := fyne.LoadResourceFromPath("./docs/thumb-down-svgrepo-com.svg")
	//if err != nil {
	//	log.Println("Failed to load indeterminate resource")
	//	return
	//}

	image := canvas.NewImageFromFile(cc.LocateFile(filename))
	image.FillMode = canvas.ImageFillContain
	image.SetMinSize(fyne.NewSize(24,16))
	image.Resize(fyne.NewSize(64,48))

	return image
}