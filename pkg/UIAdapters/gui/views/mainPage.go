package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// main page
func (vp *viewProvider) MainPage() fyne.CanvasObject {
	scheduleRes := SknCanvasSVGImageFromPath("./docs/timeLapse-mbo-24px.svg")
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), vp.HomeTab()),
		container.NewTabItemWithIcon("Networks", theme.ComputerIcon(), vp.NetworksTab()),
		container.NewTabItemWithIcon("Schedules", scheduleRes.Resource, vp.SchedulesTab()),
		container.NewTabItemWithIcon("Firmwares", theme.StorageIcon(), vp.FirmwaresTab()),
		container.NewTabItemWithIcon("Preferences", theme.SettingsIcon(), widget.NewLabel("Preferences!")),
	)
	tabs.SetTabLocation(container.TabLocationTop)
	statusText := widget.NewLabel("Status")

	// Status Line
	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), vp.ToolBarAddActionCb),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), vp.ToolBarRemoveActionCb),
		widget.NewToolbarSpacer(),
		widget.NewToolbarSeparator(),
	)
	selector := widget.NewSelect(*vp.networks, vp.OnNetworkSelectionChangedCb)

	vp.pageTabs = tabs
	vp.statLine = statusText
	vp.statusActions = bar
	vp.netselect = selector
	vp.netSelectedStr = (*vp.networks)[0]   // maybe from a saved preference

	statusLine := container.NewHBox(bar, selector, statusText)
	page := container.NewBorder(nil, statusLine,nil, nil, tabs)

	selector.SetSelected(vp.netSelectedStr)
	selector.Refresh()
	selector.Hide()  // initial state
	bar.Hide() // Initial State
	vp.tabStatus["Preferences"] = "GOLang is Wonderful!!!"

	tabs.OnChanged = vp.OnMainTabsChangedCb

	return page
}

