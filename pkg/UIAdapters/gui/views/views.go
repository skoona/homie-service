package views

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-kit/kit/log"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
)

type (
	// ViewProvider interface provides GUI features
	ViewProvider interface {
		HomeTab() fyne.CanvasObject
		NetworksTab() fyne.CanvasObject
		SchedulesTab() fyne.CanvasObject
		FirmwaresTab() fyne.CanvasObject
		MainPage() fyne.CanvasObject
		ToolBarAddActionCb()
		ToolBarRemoveActionCb()
		OnNetworkSelectionChangedCb(s string)
		OnMainTabsChangedCb(tab *container.TabItem)
	}

	// ViewProvider struct contains GUI state information
	viewProvider struct {
		logger log.Logger
		dSvc   *dc.CoreService
		networks *[]string
		pageTabs      *container.AppTabs
		statLine      *widget.Label
		netselect     *widget.Select
		statusActions *widget.Toolbar
		netSelectedStr string
		lastTabStr     string
		tabStatus   map[string]string
	}
)

var vp    *viewProvider

// NewViewProvider manages Tab level pages and their callbacks
func NewViewProvider(ds *dc.CoreService, nets *[]string, logger *log.Logger) ViewProvider {
	vp = &viewProvider{
		logger: log.With(*logger, "component", "ViewProvider"),
		dSvc: ds,
		networks: nets,
		tabStatus: map[string]string{},
		netSelectedStr: (*nets)[0],
	}
	return vp
}

// ToolBarRefreshActionCb callback for refresh toolbar button
func (vp *viewProvider) ToolBarRefreshActionCb() {
	sText := vp.pageTabs.CurrentTab().Text
	vp.statLine.SetText(fmt.Sprintf("%s refresh Selected", sText))
	vp.logger.Log("tab", sText, "event", "refresh called")
}
// ToolBarAddActionCb callback for statusline toolbar Add button
func (vp *viewProvider) ToolBarAddActionCb() {
	sText := vp.pageTabs.CurrentTab().Text
	vp.statLine.SetText(fmt.Sprintf("%s Add Selected", sText))
	vp.logger.Log("tab", sText, "event", "add called")
}
// ToolBarRemoveActionCb callback for statusline toolbar Remove button
func (vp *viewProvider) ToolBarRemoveActionCb() {
	sText := vp.pageTabs.CurrentTab().Text
	vp.statLine.SetText(fmt.Sprintf("%s Remove Selected", sText))
	vp.logger.Log("tab", sText, "event", "remove called")
}
// OnNetworkSelectionChangedCb callback for OnSelect Network selection
func (vp *viewProvider) OnNetworkSelectionChangedCb(s string) {
	vp.netSelectedStr = s
	vp.logger.Log("tab", vp.pageTabs.CurrentTab().Text,
		"event", "selector called",
		"value", s,
		"status", vp.tabStatus[vp.pageTabs.CurrentTab().Text])
}
// OnMainTabsChangedCb callback from OnChange Application Tabs
func (vp *viewProvider) OnMainTabsChangedCb(tab *container.TabItem) {
	// capture any changes before restoring
	_, found := vp.tabStatus[vp.lastTabStr]
	if found {
		vp.tabStatus[vp.lastTabStr] = vp.statLine.Text
	}
	vp.lastTabStr = tab.Text

	// restore prior status
	newStatusText, found := vp.tabStatus[tab.Text]
	if found {
		vp.statLine.SetText(newStatusText)
	} else {
		vp.statLine.SetText("Status Line...")
	}
	vp.tabStatus[tab.Text] = vp.statLine.Text

	// manage stat line objects
	switch tab.Text {
	case "Networks":
		if !vp.netselect.Visible() {
			vp.netselect.Show()
		}
		if vp.statusActions.Visible() {
			vp.statusActions.Hide()
		}
		tab.Content.Refresh()
	case "Schedules", "Firmwares":
		if vp.netselect.Visible() {
			vp.netselect.Hide()
		}
		if !vp.statusActions.Visible() {
			vp.statusActions.Show()
		}
	default:
		vp.statusActions.Hide()
		vp.netselect.Hide()
	}
	vp.statLine.Refresh()
}
