package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	cc "github.com/skoona/homie-service/pkg/utils"
)

// main page
func (vp *viewProvider) MainPage() fyne.CanvasObject {
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon(HomeTab, theme.HomeIcon(), vp.HomeTab()),
		container.NewTabItemWithIcon(NetworksTab, theme.ComputerIcon(), vp.NetworksTab()),
		container.NewTabItemWithIcon(ScheduleTab, cc.SknSelectThemedResource("timeLapse_o"), vp.SchedulesTab()),
		container.NewTabItemWithIcon(FirmwareTab, theme.StorageIcon(), vp.FirmwaresTab()),
		container.NewTabItemWithIcon(SitesTab, theme.SettingsIcon(), vp.SitesTab()),
	)
	tabs.SetTabLocation(container.TabLocationTop)
	statusText := widget.NewLabel("Status ...")
	statusText.Wrapping = fyne.TextWrapWord

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
	statusLine := container.NewHBox(nBar, bar, selector, statusText)
	page := container.NewBorder(nil, statusLine,nil, nil, tabs)

	vp.pageTabs = tabs
	tabs.OnChanged = vp.OnMainTabsChangedCb
	vp.statLine = statusText
	vp.statusActions = bar
	vp.netSelect = selector
	vp.netSelectedStr = (*vp.networks)[0]   // maybe from a saved preference
	vp.tabStatus[SitesTab] = "GOLang is Wonderful!!!"

	selector.SetSelected(vp.netSelectedStr)
	selector.Refresh()
	selector.Hide()  // initial state
	bar.Hide() // Initial State
	nBar.Show() // on all pages


	return page
}

