package api

import dc "github.com/skoona/homie-service/pkg/deviceCore"

// Requests object for Core Service

// noContent no response content expected
type noContent struct {
}
// ErrorResponse message describing the cause of the error
type ErrorResponse struct {
	Message string `json:"message"`
}

// AllNetworksRequest request all network on site
type AllNetworksRequest struct {
}
// AllNetworksResponse returns site network structure
type AllNetworksResponse struct {
	Body struct {
		Sites dc.SiteNetworks `json:"sites"`
	}
}

// NetworkByNameRequest request an individual network by name
type NetworkByNameRequest struct {
	Body struct {
		NetworkName string `json:"networkName"`
	}
}
// NetworkByNameResponse individual network structure
type NetworkByNameResponse struct {
	Body struct {
		Network dc.Network `json:"network"`
	}
}

// DeviceByNameRequest request homie device by name
type DeviceByNameRequest struct {
	DeviceName string `json:"deviceName"`
	NetworkName string `json:"networkName"`
}
// DeviceByNameResponse homie device structure
type DeviceByNameResponse struct {
	Body struct {
		Device dc.Device `json:"device"`
	}
}

// DeviceByIDRequest request homie device by ID on a certain network
type DeviceByIDRequest struct {
	DeviceID string `json:"deviceID"`
	NetworkName string `json:"networkName"`
}
// DeviceByIDResponse homie device structure
type DeviceByIDResponse struct {
	Body struct {
		DeviceID string `json:"deviceID"`
	}
}

// RemoveDeviceByIDRequest remove a device from the network
type RemoveDeviceByIDRequest struct {
	DeviceID string `json:"deviceID"`
	NetworkName string `json:"networkName"`
}

// PublishNetworkMessageRequest send any message over MQTT
type PublishNetworkMessageRequest struct {
	Body struct {
		Topic    string `json:"topic"`
		Qos      int    `json:"qos"`
		Retained bool   `json:"retained"`
		Value    string `json:"value"`
	}
}

// CreateScheduleRequest  create a new firmware OTA schedule
type CreateScheduleRequest struct {
	Body struct {
		NetworkName string          `json:"networkName"`
		DeviceID    string          `json:"deviceID"`
		Transport   dc.OTATransport `json:"transportType"`
		FirmwareID  string          `json:"firmwareID"`
	}
}
// CreateScheduleResponse returns id of created schedule
type CreateScheduleResponse struct {
	Body struct {
		ScheduleID string `json:"scheduleID"`
	}
}

// RemoveScheduleRequest remove or cancel a schedule
type RemoveScheduleRequest struct {
	ScheduleID string `json:"scheduleID"`
}

// ScheduleByIDRequest get a schedule by its id
type ScheduleByIDRequest struct {
	ScheduleID string `json:"scheduleID"`
}
// ScheduleByIDResponse requested schedule if found
type ScheduleByIDResponse struct {
	Body struct {
		Schedule dc.Schedule `json:"schedule"`
	}
}

// ScheduleByDeviceIDRequest get schedule for a certain device
type ScheduleByDeviceIDRequest struct {
	DeviceID string `json:"deviceID"`
}
// ScheduleByDeviceIDResponse schedule available for requested device
type ScheduleByDeviceIDResponse struct {
	Body struct {
		Schedule dc.Schedule `json:"schedule"`
	}
}

// AllFirmwaresRequest request a list of firmware on file
type AllFirmwaresRequest struct {
}
// AllFirmwaresResponse details of each firmware on file
type AllFirmwaresResponse struct {
	Body struct {
		Firmware []dc.Firmware `json:"firmware"`
	}
}

// CreateFirmwareRequest upload a new firmware
type CreateFirmwareRequest struct {
	FirmwareID string `json:"firmwareID"`
}
// CreateFirmwareResponse id of newly created firmeware
type CreateFirmwareResponse struct {
	Body struct {
		FirmwareID string `json:"firmwareID"`
	}
}

// RemoveFirmwareByIDRequest delete existing firmware
type RemoveFirmwareByIDRequest struct {
	FirmwareID string `json:"firmwareID"`
}

// FirmwareByIDRequest get a firmware by its ID
type FirmwareByIDRequest struct {
	FirmwareID string `json:"firmwareID"`
}
// FirmwareByIDResponse details of requested firmware
type FirmwareByIDResponse struct {
	Body struct {
		Firmware dc.Firmware `json:"firmware"`
	}
}

// FirmwareByNameRequest get firmware by its filename
type FirmwareByNameRequest struct {
	FirmwareName string `json:"firmwareName"`
}
// FirmwareByNameResponse details of requested firmware
type FirmwareByNameResponse struct {
	Body struct {
		Firmware dc.Firmware `json:"firmware"`
	}
}

// AllBroadcastsRequest retrieve all known broadcasts
type AllBroadcastsRequest struct {
}
// AllBroadcastsResponse list of known broadcasts
type AllBroadcastsResponse struct {
	Body struct {
		Broadcasts []dc.Broadcast `json:"broadcasts"`
	}
}

// RemoveBroadcastByIDRequest remove a particular broadcast
type RemoveBroadcastByIDRequest struct {
	BroadcastID string `json:"broadcastID"`
}

// BroadcastByIDRequest get a particular broadcast by its ID
type BroadcastByIDRequest struct {
	BroadcastID string `json:"broadcastID"`
}
// BroadcastByIDResponse details of requested broadcast
type BroadcastByIDResponse struct {
	Body struct {
		Broadcast dc.Broadcast `json:"broadcast"`
	}
}
