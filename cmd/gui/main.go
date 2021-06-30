// cmd/gui/main.go
// GUI Main using tView

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	views2 "github.com/skoona/homie-service/pkg/UIAdapters/gui/views"
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
	myApp := app.New()
	myApp.Settings().SetTheme(&hsTheme{})

	myWindow := myApp.NewWindow("Homie Service")

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Networks", theme.ComputerIcon(), views2.DefaultContent()),
		container.NewTabItemWithIcon("Schedules", theme.DocumentCreateIcon(), views2.DefaultContent()),
		container.NewTabItemWithIcon("Firmwares", theme.StorageIcon(), views2.DefaultContent()),
		container.NewTabItemWithIcon("Broadcasts", theme.InfoIcon(), views2.DefaultContent()),
		container.NewTabItemWithIcon("Preferences", theme.SettingsIcon(), widget.NewLabel("Preferences!")),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	statusLine := widget.NewLabel("Status")

	statusWindow := container.NewBorder(nil, statusLine,nil, nil, tabs) //  NewVSplit(split, widget.NewLabel("Status"))
	//statusWindow.Offset = 0.95

	myWindow.SetContent(statusWindow)
	//myWindow.Resize(fyne.NewSize(600, 800))


/*
 * Prepare for clean exit
	errs := make(chan error, 1)
	go func(shutdown chan error) {
		systemSignalChannel := make(chan os.Signal, 1)
		signal.Notify(systemSignalChannel, syscall.SIGINT, syscall.SIGTERM)
		sig := <-systemSignalChannel // wait on ctrl-c
		shutdown <- fmt.Errorf("%s", sig)
	}(errs)
*/

	myWindow.ShowAndRun()
	//errs <- errors.New("Fyne GUI Shutdown")


	level.Info(logger).Log("event", "shutdown requested", "cause", "Fyne GUI Shutdown") // <-errs)

	Shutdown()

	os.Exit(0)
}
