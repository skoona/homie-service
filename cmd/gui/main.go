// cmd/gui/main.go
// GUI Main using tView

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	wv "github.com/skoona/homie-service/pkg/UIAdapters/gui/views"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"github.com/skoona/homie-service/pkg/services"
	cc "github.com/skoona/homie-service/pkg/utils"
	"image/color"
	"os"
)

// Fyne Theme Implementation
type hsTheme struct {
}
func (m *hsTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		if v == theme.VariantLight {
			return &color.NRGBA{0xcf, 0xd8, 0xdc, 0xff}
		}
		return &color.NRGBA{0x45, 0x5A, 0x64, 0xff}
	case theme.ColorNameFocus:
		return &color.NRGBA{0xff, 0xc1, 0x07, 0xff}
	case theme.ColorNameForeground:
		return &color.NRGBA{44, 69, 232, 255}
	}

	return theme.DefaultTheme().Color(n, v)
}

func (t *hsTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}

func (t *hsTheme) Font(fyne.TextStyle) fyne.Resource {
	return theme.DefaultTextMonospaceFont()
}

func (t *hsTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}
// end theme implementation

func StartUp() (dc.CoreService, cc.Config) {
	ca, cb := services.Service()
	return *ca, *cb
}

func Shutdown() {
	services.Shutdown()
}

func main() {

	//coreSvc, cfg := StartUp()
	_, cfg := StartUp()
	logger := log.With(cfg.Logger, "ui", "base")

	// implement gui here and in ./pkg/gui/{controllers,views,actions}
	myApp := app.NewWithID("net.skoona.projects.homie-service")
	myApp.Settings().SetTheme(&hsTheme{})

	myWindow := myApp.NewWindow("Homie Service, GUI by Fyne")

	myWindow.SetContent( wv.MainPage(&myWindow, logger) )
	//myWindow.Resize(fyne.NewSize(600, 800))

	myWindow.ShowAndRun()

	level.Info(logger).Log("event", "shutdown requested", "cause", "Fyne GUI Shutdown") // <-errs)

	Shutdown()

	os.Exit(0)
}
