// cmd/gui/main.go
// GUI Main using tView

package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/components"
	 "github.com/skoona/homie-service/pkg/UIAdapters/gui/views"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"github.com/skoona/homie-service/pkg/services"
	cc "github.com/skoona/homie-service/pkg/utils"
	"os"
)

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

	myApp := app.NewWithID("net.skoona.projects.homie-service")
	myApp.Settings().SetTheme(&components.HsTheme{})

	myWindow := myApp.NewWindow("Homie Service, GUI by Fyne")
	myWindow.SetContent( views.MainPage(&myWindow, logger) )
	//myWindow.Resize(fyne.NewSize(600, 800))

	myWindow.ShowAndRun()

	level.Info(logger).Log("event", "shutdown requested", "cause", "Fyne GUI Shutdown") // <-errs)

	Shutdown()

	os.Exit(0)
}
