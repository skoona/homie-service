package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/components"
	cc "github.com/skoona/homie-service/pkg/utils"
)

// main page
func (vp *viewProvider) MainPage() fyne.CanvasObject {
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), vp.HomeTab()),
		container.NewTabItemWithIcon("Networks", theme.ComputerIcon(), vp.NetworksTab()),
		container.NewTabItemWithIcon("Schedules", cc.SknSelectThemedResource("timeLapse_o"), vp.SchedulesTab()),
		container.NewTabItemWithIcon("Firmwares", theme.StorageIcon(), vp.FirmwaresTab()),
		container.NewTabItemWithIcon("Preferences", theme.SettingsIcon(), components.MakeTableTab()),
	)
	tabs.SetTabLocation(container.TabLocationTop)
	statusText := widget.NewLabel("Status")

	// Status Line
	bar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ContentAddIcon(), vp.ToolBarAddActionCb),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), vp.ToolBarRemoveActionCb),
		widget.NewToolbarSpacer(),
		widget.NewToolbarSeparator(),
	)
	nBar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), vp.ToolBarRefreshActionCb),
		widget.NewToolbarSpacer(),
		widget.NewToolbarSeparator(),
	)

	selector := widget.NewSelect(*vp.networks, vp.OnNetworkSelectionChangedCb)

	vp.pageTabs = tabs
	vp.statLine = statusText
	vp.statusActions = bar
	vp.netselect = selector
	vp.netSelectedStr = (*vp.networks)[0]   // maybe from a saved preference

	statusLine := container.NewHBox(nBar, bar, selector, statusText)
	page := container.NewBorder(nil, statusLine,nil, nil, tabs)

	selector.SetSelected(vp.netSelectedStr)
	selector.Refresh()
	selector.Hide()  // initial state
	bar.Hide() // Initial State
	nBar.Show() // on all pages
	vp.tabStatus["Preferences"] = "GOLang is Wonderful!!!"

	tabs.OnChanged = vp.OnMainTabsChangedCb

	return page
}

