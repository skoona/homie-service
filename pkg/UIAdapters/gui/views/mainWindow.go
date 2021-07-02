package views

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/go-kit/kit/log"
)

// main page
func MainPage(w *fyne.Window, logger log.Logger) fyne.CanvasObject {
	scheduleRes := SknCanvasSVGImageFromPath("./docs/timeLapse-mbo-24px.svg")
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), HomeTab()),
		container.NewTabItemWithIcon("Networks", theme.ComputerIcon(), NetworksTab()),
		container.NewTabItemWithIcon("Schedules", scheduleRes.Resource, SchedulesTab()),
		container.NewTabItemWithIcon("Firmwares", theme.StorageIcon(), FirmwaresTab()),
		container.NewTabItemWithIcon("Preferences", theme.SettingsIcon(), widget.NewLabel("Preferences!")),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	// Status Line
	statusText := widget.NewLabel("Status")
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			sText := tabs.CurrentTab().Text
			statusText.SetText(fmt.Sprintf("%s Add Selected", sText))
			logger.Log("tab", sText, "event", "add called")
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func(){
			sText := tabs.CurrentTab().Text
			statusText.SetText(fmt.Sprintf("%s Remove Selected", sText))
			logger.Log("tab", sText, "event", "remove called")
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarSeparator(),
	)
	selector := widget.NewSelect([]string{"sknSensors", "homie"}, func(s string) {
		sText := tabs.CurrentTab().Text
		statusText.SetText(fmt.Sprintf("[%s] was Selected on tab: %s", s, sText))
		logger.Log("tab", sText, "event", "selector called", "value", s)
	})

	statusLine := container.NewHBox(bar, selector, statusText)
	page := container.NewBorder(nil, statusLine,nil, nil, tabs)

	selector.Hide()  // initial state
	bar.Hide() // Initial State

	tabs.OnChanged = func(tab *container.TabItem) {
		statusText.SetText(tab.Text)
		switch statusText.Text {
		case "Networks":
			selector.PlaceHolder = tab.Text
			if selector.Visible() {
				selector.Refresh()
			} else {
				selector.Show()
			}
			if bar.Visible() {
				bar.Hide()
			}
		case "Schedules", "Firmwares":
			if selector.Visible() {
				selector.Hide()
			}
			if bar.Visible() {
				bar.Refresh()
			} else {
				bar.Show()
			}
		default:
			bar.Hide()
			selector.Hide()
		}

	}
	return page
}

