package views

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/go-kit/kit/log"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	cc "github.com/skoona/homie-service/pkg/utils"
	"image/color"

)



type (
	ViewProvider interface {
		HomeTab() fyne.CanvasObject
		NetworksTab() fyne.CanvasObject
		SchedulesTab() fyne.CanvasObject
		FirmwaresTab() fyne.CanvasObject
		MainPage() fyne.CanvasObject
		ToolBarAddActionCb()
		ToolBarRemoveActionCb()
		OnNetworkSelectionChangedCb(s string)
		OnMainTabsChangedCb(tab *container.TabItem)
	}

	viewProvider struct {
		logger log.Logger
		dSvc   *dc.CoreService
		networks *[]string
		pageTabs      *container.AppTabs
		statLine      *widget.Label
		netselect     *widget.Select
		statusActions *widget.Toolbar
		netSelectedStr string
		lastTabStr     string
		tabStatus   map[string]string
	}
)


func NewViewProvider(ds *dc.CoreService, nets *[]string, logger *log.Logger) ViewProvider {
	vp = &viewProvider{
		logger: log.With(*logger, "component", "ViewProvider"),
		dSvc: ds,
		networks: nets,
		tabStatus: map[string]string{},
	}
	return vp
}

var (
	Green color.Color = color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	vp  *viewProvider
)
//text1 := canvas.NewText("Hello", green)

// SknCanvasSVGImageFromPath loads as an canvas.Image
func SknCanvasSVGImageFromPath(filename string) *canvas.Image  {
	image := SknLoadImageFromURI(storage.NewFileURI( cc.LocateFile(filename))).(*canvas.Image)
	image.Resource = theme.NewThemedResource(image.Resource)
	image.FillMode = canvas.ImageFillContain
	return image
}

// SknLoadImageFromPath returns a CanvasObject
func SknLoadImageFromPath(path string) fyne.CanvasObject {
	return SknLoadImageFromURI(storage.NewFileURI(path))
}

// SknLoadImageFromURI returns a CanvasObject
func SknLoadImageFromURI(u fyne.URI) fyne.CanvasObject {
	read, err := storage.Reader(u)
	if err != nil {
		fmt.Println("Error opening image", err)
		return canvas.NewRectangle(color.Black)
	}
	res, err :=	storage.LoadResourceFromURI(read.URI())
	if err != nil {
		fmt.Println("Error reading image", err)
		return canvas.NewRectangle(color.Black)
	}
	img := canvas.NewImageFromResource(res)
	//img.FillMode = canvas.ImageFillContain
	img.FillMode = canvas.ImageFillOriginal
	return img
}

func (vp *viewProvider) ToolBarAddActionCb() {
	sText := vp.pageTabs.CurrentTab().Text
	vp.statLine.SetText(fmt.Sprintf("%s Add Selected", sText))
	vp.logger.Log("tab", sText, "event", "add called")
}
func (vp *viewProvider) ToolBarRemoveActionCb() {
	sText := vp.pageTabs.CurrentTab().Text
	vp.statLine.SetText(fmt.Sprintf("%s Remove Selected", sText))
	vp.logger.Log("tab", sText, "event", "remove called")
}
func (vp *viewProvider) OnNetworkSelectionChangedCb(s string) {
	vp.netSelectedStr = s
	vp.logger.Log("tab", vp.pageTabs.CurrentTab().Text,
		"event", "selector called",
		"value", s,
		"status", vp.tabStatus[vp.pageTabs.CurrentTab().Text])
}
func (vp *viewProvider) OnMainTabsChangedCb(tab *container.TabItem) {
	// capture any changes before restoring
	_, found := vp.tabStatus[vp.lastTabStr]
	if found {
		vp.tabStatus[vp.lastTabStr] = vp.statLine.Text
	}
	vp.lastTabStr = tab.Text

	// restore prior status
	newStatusText, found := vp.tabStatus[tab.Text]
	if found {
		vp.statLine.SetText(newStatusText)
	} else {
		vp.statLine.SetText("Status Line...")
	}
	vp.tabStatus[tab.Text] = vp.statLine.Text

	// manage stat line objects
	switch tab.Text {
	case "Networks":
		if !vp.netselect.Visible() {
			vp.netselect.Show()
		}
		if vp.statusActions.Visible() {
			vp.statusActions.Hide()
		}
	case "Schedules", "Firmwares":
		if vp.netselect.Visible() {
			vp.netselect.Hide()
		}
		if !vp.statusActions.Visible() {
			vp.statusActions.Show()
		}
	default:
		vp.statusActions.Hide()
		vp.netselect.Hide()
	}
	vp.statLine.Refresh()
}
