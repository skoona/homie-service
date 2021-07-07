package views

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-kit/kit/log"
	"github.com/skoona/homie-service/pkg/UIAdapters/gui/components"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
)

const (
	HomeTab     = "Home"
	NetworksTab = "Networks"
	FirmwareTab = "Firmwares"
	ScheduleTab = "Schedules"
	SitesTab    = "Sites"
	CardWidth   = 384
	CardHeight  = 448
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
		logger           log.Logger
		dSvc             *dc.CoreService
		networks         *[]string
		pageTabs         *container.AppTabs
		statLine         *widget.Label
		netSelect        *widget.Select
		statusActions    *widget.Toolbar
		netSelectedStr   string
		lastTabStr       string
		tabStatus        map[string]string
		inventoryMLEntry *widget.Entry
		homeCards        *widget.List
		networkCards     *widget.List
		networkSide      *fyne.Container
		scheduleCards    *widget.List
		scheduleSide     *fyne.CanvasObject
		firmwareCards    *widget.List
		firmwareSide     *fyne.CanvasObject
		siteNetworks     *dc.SiteNetworks
		devices          []dc.Device
		schedules        []dc.Schedule
		devSummary       [][]string
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
		devSummary: make([][]string, 24),
		siteNetworks: (*ds).PrivateSiteNetworks(),
		devices: []dc.Device{},
		schedules: []dc.Schedule{},
	}
	return vp
}

// ToolBarRefreshActionCb callback for refresh toolbar button
func (vp *viewProvider) ToolBarRefreshActionCb() {
	sText := vp.pageTabs.CurrentTab().Text
	switch sText {
	case HomeTab:
		vp.pageTabs.CurrentTab().Content = vp.HomeTab()
	case NetworksTab:
		vp.pageTabs.CurrentTab().Content = vp.NetworksTab()
	case ScheduleTab:
		vp.pageTabs.CurrentTab().Content = vp.SchedulesTab()
	case FirmwareTab:
		vp.pageTabs.CurrentTab().Content = vp.FirmwaresTab()
	case SitesTab:
		vp.pageTabs.CurrentTab().Content = vp.SitesTab()
	default:
		vp.pageTabs.CurrentTab().Content.Refresh()
	}
	vp.pageTabs.Refresh()
	vp.statLine.SetText(fmt.Sprintf("%s refresh completed", sText))
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
	case NetworksTab:
		if !vp.netSelect.Visible() {
			vp.netSelect.Show()
		}
		if vp.statusActions.Visible() {
			vp.statusActions.Hide()
		}
		tab.Content.Refresh()
	case ScheduleTab, FirmwareTab:
		if vp.netSelect.Visible() {
			vp.netSelect.Hide()
		}
		if !vp.statusActions.Visible() {
			vp.statusActions.Show()
		}
	default:
		vp.statusActions.Hide()
		vp.netSelect.Hide()
	}
	vp.statLine.Refresh()
}

func (vp *viewProvider) OnBroadcastSelected(id widget.ListItemID) {
	vp.statLine.SetText(fmt.Sprintf("%s Broadcast Selected", vp.siteNetworks.Broadcasts[id].Level))
	vp.logger.Log("broadcast id", id, "event", "selected","Broadcast", vp.siteNetworks.Broadcasts[id].Level)
}
func (vp *viewProvider) OnDeviceSelected(id widget.ListItemID) {
	vp.networkSide.Objects[0] = components.SknDeviceTreeSide(&vp.devices[id])
	vp.networkSide.Refresh()
	vp.statLine.SetText(fmt.Sprintf("%s Device Selected", vp.devices[id].Name))
	vp.logger.Log("Device index", id, "event", "selected","device", vp.devices[id].Name)
}
func (vp *viewProvider) OnFirmwareSelected(id widget.ListItemID) {
	vp.statLine.SetText(fmt.Sprintf("%s Firmware Selected", vp.siteNetworks.Firmwares[id].Name))
	vp.logger.Log("Firmware index", id, "event", "selected", "firmware", vp.siteNetworks.Firmwares[id].Name)
}
func (vp *viewProvider) OnScheduleSelected(id widget.ListItemID) {
	vp.statLine.SetText(fmt.Sprintf("%s Schedule Selected", vp.schedules[id].ID))
	vp.logger.Log("Schedule index", id, "event", "selected","scheduleId", vp.schedules[id].ID)
}


