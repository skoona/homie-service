package providers

import (
	"fyne.io/fyne/v2"
	"github.com/go-kit/kit/log"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/components"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/views"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	cc "github.com/skoona/homie-service/pkg/utils"
	"time"
)

// GuiProvider
type (
	GuiProvider interface {
		MainPage() fyne.CanvasObject
		HomieTheme() fyne.Theme
	}

// guiProvider handler for getting and updating Pages
 	guiProvider struct {
		_ struct{}
		root 		*fyne.Window
		service 	*dc.CoreService
		logger 		log.Logger
		cfg         *cc.Config
		nets        *[]string
		theme 		fyne.Theme
		vp          views.ViewProvider
	}
)

var (
	gtrl *guiProvider
)

func NewGuiController(cfg *cc.Config, root *fyne.Window, svc *dc.CoreService, nets *[]string, logger *log.Logger) GuiProvider {
	gtrl = &guiProvider{
		root: root,
		cfg: cfg,
		service: svc,
		logger: log.With(*logger,"component","GuiProvider"),
		nets: nets,
		theme: components.NewHomieTheme(),
		vp: views.NewViewProvider(svc, nets, logger),
	}
	return gtrl
}
func (g *guiProvider) HomieTheme() fyne.Theme {
	return g.theme
}
func (g *guiProvider) MainPage() fyne.CanvasObject {
	time.Sleep(15 * time.Second) // Allow MQTT messages to update Broadcasts
	return g.vp.MainPage()
}

