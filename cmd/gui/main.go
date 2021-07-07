//go:generate fyne bundle -o ./svgImages.go ../../svgResources


// cmd/gui/main.go
// GUI Main using tView

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/providers"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"github.com/skoona/homie-service/pkg/services"
	cc "github.com/skoona/homie-service/pkg/utils"
	"net/url"
	"os"
)


func shortcutFocused(s fyne.Shortcut, w fyne.Window) {
	if focused, ok := w.Canvas().Focused().(fyne.Shortcutable); ok {
		focused.TypedShortcut(s)
	}
}

func sknMenus(a fyne.App, w fyne.Window) {
	settingsItem := fyne.NewMenuItem("Settings", func() {
		w := a.NewWindow("HomieService Settings")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	})

	cutItem := fyne.NewMenuItem("Cut", func() {
		shortcutFocused(&fyne.ShortcutCut{
			Clipboard: w.Clipboard(),
		}, w)
	})
	copyItem := fyne.NewMenuItem("Copy", func() {
		shortcutFocused(&fyne.ShortcutCopy{
			Clipboard: w.Clipboard(),
		}, w)
	})
	pasteItem := fyne.NewMenuItem("Paste", func() {
		shortcutFocused(&fyne.ShortcutPaste{
			Clipboard: w.Clipboard(),
		}, w)
	})

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("Documentation", func() {
			u, _ := url.Parse("https://developer.fyne.io")
			_ = a.OpenURL(u)
		}),
		fyne.NewMenuItem("Support", func() {
			u, _ := url.Parse("https://fyne.io/support/")
			_ = a.OpenURL(u)
		}),
	)
	file := fyne.NewMenu("File")
	if !fyne.CurrentDevice().IsMobile() {
		file.Items = append(file.Items, fyne.NewMenuItemSeparator(), settingsItem)
	}
	mainMenu := fyne.NewMainMenu(
		// a quit item will be appended to our first menu
		file,
		fyne.NewMenu("Edit", cutItem, copyItem, pasteItem),
		helpMenu,
	)
	w.SetMainMenu(mainMenu)
	w.SetMaster()
}

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
	sknMenus(myApp, myWindow)
	provider := providers.NewGuiController(&cfg, &myApp, &myWindow, &coreSvc, &networks, &logger)
	//myApp.Settings().SetTheme(provider.HomieTheme())
	
	myWindow.SetContent( provider.MainPage() )

	myWindow.ShowAndRun()

	level.Info(logger).Log("event", "shutdown requested", "cause", "Fyne GUI Shutdown") // <-errs)

	Shutdown()

	os.Exit(0)
}
