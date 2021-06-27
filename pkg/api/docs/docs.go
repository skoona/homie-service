// Package classification Homie-Service API
//
// ### Documentation for Homie-Service API.
//
// Rewrite of [HomieMonitor](https://github.com/skoona/homieMonitor) in Golang.
//
// API service implements the `Homie Controller`, or `Monitor`, supporting development of IOT/Devices
// using [Homie-esp8266](https://github.com/homieiot/homie-esp8266) repository; although any `Homie Device` specification should be supported.
//
// ### References:
// * [Homie: An MQTT Convention for IOT/M2M](https://homieiot.github.io/specification/)
//
// * [Homie-ESP8266 Example of RCWL-0516 Microwave Presence Detector and DHT22 Temperature and Humidity sensors](https://github.com/skoona/EnvironmentMonitor_DHT)
//
// ### Status
// * Support Mutliple Homie Networks (auto discovery); /homie/#, /custom/#, ...
//
// * [MQTT](https://github.com/eclipse/paho.mqtt.golang) Unsecured connection.
//
// * Produce MQTT messages decoded to internal Device Model.
//
// * Produce MQTT messages from MQTT logfile for Demo use.
//
// * [bBolt DB](https://github.com/etcd-io/bbolt) data storage for decoded device messages, and ota schedules
//
// * CoreLogic
//
// * * decode and transform data into Homie's Broadcast, Device, Node, Property, and Attribute collections.
//
// * * encode firmware information into the Firmware OTA Scheduling Service
//
// * * implement firmware schedule
//
// * * implement data retention on change service
//
// * Create Web UI with WedSocket Driven Components
//
// * Enable SSL/TLS Connection to MQTT
//
// * Enabled OTA Scheduling.
//
//	Schemes: http
//	BasePath: /api/v1
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package docs

import (
	_ "github.com/pdrum/swagger-automation/api"
	"github.com/skoona/homie-service/pkg/api/handlers"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
)


// swagger:parameters noParamsRequired
type noContentRequestWrapper struct {
	// No content is expected
	// In: path
}

// No content is expected.
// swagger:response noContentResponse
type noContentResponseWrapper struct {
	// In: body
	Body handlers.NoContent
}

// a collection of validation error messages.
// swagger:response validationError
type validationErrorMessageWrapper struct {
	// In: body
	Body handlers.ValidationErrorMessage
}

// generic error message returned by the server.
// swagger:response genericError
type genericErrorWrapper struct {
	// In: body
	Body handlers.GenericError
}

// indicates a specific network by name.
// swagger:parameters networkNameParam
type networkNameParamsWrapper struct {
	// pattern: [a-zA-Z]+
	// In: path
	NetworkName string `json:"networkName"`
}

// returns site network structure.
// swagger:response networksResponse
type networksResponseWrapper struct {
	// In: body
	Body dc.SiteNetworks
}

// Individual network structure.
// swagger:response networkResponse
type networkByNameResponseWrapper struct {
	// In: body
	Body  dc.Network
}

// params of network and device names
// swagger:parameters deviceByNameParams
type deviceByNameRequestWrapper struct {
	// name of network for device.
	// pattern: [a-zA-Z]+
	// In: path
		NetworkName string `json:"networkName"`
	// homie device name
	// pattern: [a-zA-Z]+
	// In: path
		DeviceName  string `json:"deviceName"`
}

// params of network name and device ID
// swagger:parameters deviceByIdParams
type deviceByIdRequestWrapper struct {
	// name of network for device.
	// pattern: [a-zA-Z]+
	// In: path
	NetworkName string `json:"networkName"`
	// 32 char identifier for device
	// pattern: [a-zA-Z0-9]+
	// In: path
	DeviceID  string `json:"deviceID"`
}

// device details
// swagger:response deviceResponse
type deviceByNameResponseWrapper struct {
	// homie device structure.
	// In: body
	Body dc.Device
}

// params  to send any message over MQTT.
// swagger:parameters networkMessageParams
type NetworkMessageRequestWrapper struct {
	// In: body
	Body handlers.NetworkMessageRequest
}

// params to create a new firmware OTA schedule.
// swagger:parameters createScheduleParams
type CreateScheduleRequestWrapper struct {
	// In: body
	Body handlers.CreateScheduleRequest
}

// ID of a schedule.
// swagger:parameters scheduleIdParam
type ScheduleIdRequestWrapper struct {
	// In: path
	ScheduleID string `json:"scheduleID"`
}

// ID of a schedule.
// swagger:response scheduleIdResponse
type ScheduleIdResponseWrapper struct {
	// In: body
	Body string `json:"scheduleID"`
}

// schedule declared for device.
// swagger:response scheduleResponse
type ScheduleResponseWrapper struct {
	// In: body
	Body dc.Schedule `json:"schedule"`
}

// details of every firmware on file.
// swagger:response firmwaresResponse
type FirmwaresResponseWrapper struct {
	// In: body
	Body []dc.Firmware
}

// params to upload a new firmware.
// swagger:parameters firmwareCreateParams
type CreateFirmwareRequestWrapper struct {
	// In: body
	Body handlers.CreateFirmwareRequest
}

// get a firmware by its ID.
// swagger:parameters firmwareIdParam
type FirmwareIdParamWrapper struct {
	// In: path
	FirmwareID string `json:"firmwareID"`
}

// get firmware by its filename.
// swagger:parameters firmwareNameParam
type FirmwareNameParamWrapper struct {
	// In: path
	FirmwareName string `json:"firmwareName"`
}

// ID of firmware.
// swagger:response firmwareIdResponse
type FirmwareIdResponseWrapper struct {
	// In: body
	Body string `json:"firmwareID"`
}

// details of requested firmware.
// swagger:response firmwareResponse
type FirmwareResponseWrapper struct {
	// In: body
	Body dc.Firmware `json:"firmware"`
}

// get a particular broadcast by its ID.
// swagger:parameters broadcastIdParam
type BroadcastIdParamWrapper struct {
	// In: path
	BroadcastID string `json:"broadcastID"`
}

// list of known broadcasts.
// swagger:response broadcastsResponse
type BroadcastsResponseWrapper struct {
	Body []dc.Broadcast `json:"broadcasts"`
}

// details of requested broadcast.
// swagger:response broadcastResponse
type BroadcastResponseWrapper struct {
	// In: body
	Body dc.Broadcast `json:"broadcast"`
}
