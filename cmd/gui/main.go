// cmd/gui/main.go
// GUI Main using tView

package main

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"github.com/skoona/homie-service/pkg/gui/views"
	"os"
	"os/signal"
	"syscall"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"github.com/skoona/homie-service/pkg/services"
	cc "github.com/skoona/homie-service/pkg/utils"
	//"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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

	// implement gui here and in ./pkg/gui/{controllers,views,actions}
	myApp := app.New()
	myWindow := myApp.NewWindow("Homie Service")

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Networks", theme.ComputerIcon(), views.DefaultContent()),
		container.NewTabItemWithIcon("Schedules", theme.DocumentCreateIcon(), views.DefaultContent()),
		container.NewTabItemWithIcon("Firmwares", theme.StorageIcon(), views.DefaultContent()),
		container.NewTabItemWithIcon("Broadcasts", theme.InfoIcon(), views.DefaultContent()),
		container.NewTabItemWithIcon("Preferences", theme.SettingsIcon(), widget.NewLabel("Preferences!")),
	)
	tabs.SetTabLocation(container.TabLocationTop)
	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(600, 800))


	/*
	 * Prepare for clean exit
	*/
	errs := make(chan error, 1)
	go func(shutdown chan error) {
		systemSignalChannel := make(chan os.Signal, 1)
		signal.Notify(systemSignalChannel, syscall.SIGINT, syscall.SIGTERM)
		sig := <-systemSignalChannel // wait on ctrl-c
		shutdown <- fmt.Errorf("%s", sig)
	}(errs)

	myWindow.ShowAndRun()
	errs <- errors.New("Fyne GUI Shutdown")


	level.Info(logger).Log("event", "shutdown requested", "cause", <-errs)

	Shutdown()

	os.Exit(0)
}
