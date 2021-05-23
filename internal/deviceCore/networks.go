package deviceCore

/*
	Defines the site collection of network models
*/

// SiteNetworks contains all known/valid Collections
type SiteNetworks struct {
	ID             EID
	ElementType    CoreType // site name
	SiteName       string
	Title          string
	Names          []string // any
	DeviceNetworks map[string]Network
	Broadcasts     []Broadcast
	Firmwares      []Firmware
	Schedules      map[string]Schedule
}

// NewNetworks Creates Component
func NewSiteNetworks(siteName, siteTitle string, networks []string, firmwares []Firmware, schedules map[string]Schedule) (*SiteNetworks, error) {
	var err error

	siteNetworks = SiteNetworks{
		ID:             NewEID(),
		ElementType:    CoreTypeSiteNetworks,
		SiteName:       siteName,
		Title:          siteTitle,
		Names:          networks,
		DeviceNetworks: make(map[string]Network, len(networks)+1),
		Broadcasts:     []Broadcast{},
		Firmwares:      firmwares,
		Schedules:      schedules,
	}

	for _, nName := range networks {
		siteNetworks.DeviceNetworks[nName] = newNetwork(nName)
	}

	return &siteNetworks, err
}

// newNetwork Creates Component
func newNetwork(name string) Network {
	hn := Network{
		ID:          NewEID(),
		ElementType: CoreTypeNetwork,
		Name:        name,
		Devices:     make(map[string]Device, 16),
	}

	siteNetworks.DeviceNetworks[name] = hn

	return hn
}

/*
 * SiteNetworksRepository - Application Domain Rules for elements
 */
type SiteNetworksRepository interface {
	CreateNetworks(discoveredNetworks string) (*SiteNetworks, error)
	ChangeSiteNetworkTitle(siteNetworkTitle string) error
	ChangeNetworkTitle(networkName, networkTitle string) error
	ListNetworks() (*[]string, error)
	GetNetwork(networkName string) (*Network, bool)
}
