//go:generate fyne bundle -o ./svgImages.go ../../svgResources


// cmd/gui/main.go
// GUI Main using tView

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/providers"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"github.com/skoona/homie-service/pkg/services"
	cc "github.com/skoona/homie-service/pkg/utils"
	"os"
)

// StartUp core services
func StartUp() (dc.CoreService, cc.Config, []string) {
	ca, cb, networks := services.Service()
	return *ca, *cb, networks
}

// Shutdown coreServices
func Shutdown() {
	services.Shutdown()
}

func main() {

	coreSvc, cfg, networks := StartUp()
	logger := log.With(cfg.Logger, "ui", "base")

	myApp := app.NewWithID("net.skoona.projects.homie-service")
	myWindow := myApp.NewWindow("Homie Service, GUI by Fyne")
	provider := providers.NewGuiController(&cfg, &myWindow, &coreSvc, &networks, &logger)
	myApp.Settings().SetTheme(provider.HomieTheme())

	myWindow.SetContent( provider.MainPage() )
	myWindow.Resize(fyne.NewSize(560, 400))

	myWindow.ShowAndRun()

	level.Info(logger).Log("event", "shutdown requested", "cause", "Fyne GUI Shutdown") // <-errs)

	Shutdown()

	os.Exit(0)
}
