package api

import dc "github.com/skoona/homie-service/pkg/deviceCore"

// Requests object for Core Service

// AllNetworksRequest request all network on site
type AllNetworksRequest struct {
}
// AllNetworksResponse returns site network structure
type AllNetworksResponse struct {
	Sites dc.SiteNetworks `json:"sites"`
}

// NetworkByNameRequest request an individual network by name
type NetworkByNameRequest struct {
	NetworkName string `json:"networkName"`
}
// NetworkByNameResponse individual network structure
type NetworkByNameResponse struct {
	Network dc.Network `json:"network"`
}

// DeviceByNameRequest request homie device by name
type DeviceByNameRequest struct {
	DeviceName string `json:"deviceName"`
	NetworkName string `json:"networkName"`
}
// DeviceByNameResponse homie device structure
type DeviceByNameResponse struct {
	Device dc.Device `json:"device"`
	EMessage error `json:"eMessage"`
}

// DeviceByIDRequest request homie device by ID on a certain network
type DeviceByIDRequest struct {
	DeviceID string `json:"deviceID"`
	NetworkName string `json:"networkName"`
}
// DeviceByIDResponse homie device structure
type DeviceByIDResponse struct {
	DeviceID string `json:"deviceID"`
	EMessage error `json:"eMessage"`
}

// RemoveDeviceByIDRequest remove a device from the network
type RemoveDeviceByIDRequest struct {
	DeviceID string `json:"deviceID"`
	NetworkName string `json:"networkName"`
}
// RemoveDeviceByIDResponse result of removal attempt
type RemoveDeviceByIDResponse struct {
	EMessage error `json:"eMessage"`
}

// PublishNetworkMessageRequest send any message over MQTT
type PublishNetworkMessageRequest struct {
	Topic string `json:"topic"`
	Qos int `json:"qos"`
	Retained bool `json:"retained"`
	Value string `json:"value"`
}
// PublishNetworkMessageResponse no response expected
type PublishNetworkMessageResponse struct {
}

// CreateScheduleRequest  create a new firmware OTA schedule
type CreateScheduleRequest struct {
	NetworkName string `json:"networkName"`
	DeviceID string `json:"deviceID"`
	Transport dc.OTATransport `json:"transportType"`
	FirmwareID string `json:"firmwareID"`
}
// CreateScheduleResponse returns id of created schedule
type CreateScheduleResponse struct {
	ScheduleID string `json:"scheduleID"`
	EMessage error `json:"eMessage"`
}

// RemoveScheduleRequest remove or cancel a schedule
type RemoveScheduleRequest struct {
	ScheduleID string `json:"scheduleID"`
}
// RemoveScheduleResponse no response expected
type RemoveScheduleResponse struct {
}

// ScheduleByIDRequest get a schedule by its id
type ScheduleByIDRequest struct {
	ScheduleID string `json:"scheduleID"`
}
// ScheduleByIDResponse requested schedule if found
type ScheduleByIDResponse struct {
	Schedule dc.Schedule `json:"schedule"`
}

// ScheduleByDeviceIDRequest get schedule for a certain device
type ScheduleByDeviceIDRequest struct {
	DeviceID string `json:"deviceID"`
}
// ScheduleByDeviceIDResponse schedule available for requested device
type ScheduleByDeviceIDResponse struct {
	Schedule dc.Schedule `json:"schedule"`
}

// AllFirmwaresRequest request a list of firmware on file
type AllFirmwaresRequest struct {
}
// AllFirmwaresResponse details of each firmware on file
type AllFirmwaresResponse struct {
	Firmware []dc.Firmware `json:"firmware"`
}

// CreateFirmwareRequest upload a new firmware
type CreateFirmwareRequest struct {
	FirmwareID string `json:"firmwareID"`
}
// CreateFirmwareResponse id of newly created firmeware
type CreateFirmwareResponse struct {
	FirmwareID string `json:"firmwareID"`
	EMessage error `json:"eMessage"`
}

// RemoveFirmwareByIDRequest delete existing firmware
type RemoveFirmwareByIDRequest struct {
	FirmwareID string `json:"firmwareID"`
}
// RemoveFirmwareByIDResponse no result expected
type RemoveFirmwareByIDResponse struct {
}

// FirmwareByIDRequest get a firmware by its ID
type FirmwareByIDRequest struct {
	FirmwareID string `json:"firmwareID"`
}
// FirmwareByIDResponse details of requested firmware
type FirmwareByIDResponse struct {
	Firmware dc.Firmware `json:"firmware"`
	EMessage error `json:"eMessage"`
}

// FirmwareByNameRequest get firmware by its filename
type FirmwareByNameRequest struct {
	FirmwareName string `json:"firmwareName"`
}
// FirmwareByNameResponse details of requested firmware
type FirmwareByNameResponse struct {
	Firmware dc.Firmware `json:"firmware"`
	EMessage error `json:"eMessage"`
}

// AllBroadcastsRequest retrieve all known broadcasts
type AllBroadcastsRequest struct {
}
// AllBroadcastsResponse list of known broadcasts
type AllBroadcastsResponse struct {
	Broadcasts []dc.Broadcast `json:"broadcasts"`
}

// RemoveBroadcastByIDRequest remove a particular broadcast
type RemoveBroadcastByIDRequest struct {
	BroadcastID string `json:"broadcastID"`
}
// RemoveBroadcastByIDResponse no reposne expected
type RemoveBroadcastByIDResponse struct {
}

// BroadcastByIDRequest get a particular broadcast by its ID
type BroadcastByIDRequest struct {
	BroadcastID string `json:"broadcastID"`
}
// BroadcastByIDResponse details of requested broadcast
type BroadcastByIDResponse struct {
	Broadcast dc.Broadcast `json:"broadcast"`
	EMessage error `json:"eMessage"`
}
