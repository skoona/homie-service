package handlers

import (
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"net/http"
	"strings"
)

// swagger:route GET /networks network-tag noParamsRequired
// Return a list of all onsite networks.
// ---
// parameters:
// - none
// responses:
//	202: networksResponse
//  500: genericError

// AllNetworks  List all SiteNetworks
func (c *Controller) AllNetworks(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "AllNetworks() called")
	rw.Header().Add("Content-Type", "application/json")

	body := c.service.AllNetworks()

	rw.WriteHeader(http.StatusAccepted)
	err := ToJSON(body, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		level.Error(c.logger).Log( "error", err.Error())
	}
}


// swagger:route GET /network/{networkName:[a-zA-Z]+} network-tag networkNameParam
// Return a specified named network.
// ---
// parameters:
//  - networkName
//    in: path
//    type: string
// responses:
//	202: networkResponse
//  404: genericError

// NetworkByName (networkName string) Network
func (c *Controller) NetworkByName(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "NetworkByName() called")
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	body := c.service.NetworkByName(vars["networkName"])

	rw.WriteHeader(http.StatusAccepted)
	err := ToJSON(body, rw)
	if err != nil {
		level.Error(c.logger).Log( "error", err.Error())
		rw.WriteHeader(http.StatusNotFound)
		ToJSON(&GenericError{Message: err.Error()}, rw)
	}
}

// swagger:route GET /deviceByName/{networkName:[a-zA-Z]+}/{deviceName:[a-zA-Z]+} network-tag deviceByNameParams
// Return a specified device on the named network.
// responses:
//	202: deviceResponse
//  404: genericError
//  500: genericError

// DeviceByName (deviceName, networkName string) (Device, error)
func (c *Controller) DeviceByName(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "DeviceByName() called")
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	body, err := c.service.DeviceByName(vars["deviceName"], vars["networkName"])
	if err == nil {
		rw.WriteHeader(http.StatusAccepted)
		err := ToJSON(body, rw)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			level.Error(c.logger).Log( "error", err.Error())
		}
	} else {
		level.Error(c.logger).Log( "error", err.Error())
		rw.WriteHeader(http.StatusNotFound)
		ToJSON(&GenericError{Message: err.Error()}, rw)
	}
}


// swagger:route GET /deviceById/{networkName:[a-zA-Z]+}/{deviceID:[a-zA-Z0-9]+} network-tag  deviceByIdParams
// Return a specified device using it unique deviceID the named network
// responses:
//	202: deviceResponse
//  404: genericError
//  406: validationError
//  500: genericError

// DeviceByID (deviceID string, networkName string) (Device, error)
func (c *Controller) DeviceByID(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "DeviceByID() called")
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	emsg := c.validation.ValidateParam(vars["deviceID"], "eid")
	if len(emsg.Messages) != 0 {
		rw.WriteHeader(http.StatusNotAcceptable)
		level.Error(c.logger).Log( "validation", emsg.Messages)
		ToJSON(&emsg, rw)
		return
	}

	body, err := c.service.DeviceByID(vars["deviceID"], vars["networkName"])
	if err == nil {
		rw.WriteHeader(http.StatusAccepted)
		err := ToJSON(body, rw)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			level.Error(c.logger).Log( "error", err.Error())
		}
	} else {
		level.Error(c.logger).Log( "error", err.Error())
		rw.WriteHeader(http.StatusNotFound)
		ToJSON(&GenericError{Message: err.Error()}, rw)
	}
}


// swagger:route DELETE /removeDeviceId/{networkName:[a-zA-Z]+}/{deviceID:[a-zA-Z0-9]+} network-tag deviceByIdParams2
// Removes a device from the specified network
// responses:
//	204: noContentResponse
//  404: genericError
//  406: validationError

// RemoveDeviceID (deviceID string, networkName string) error
func (c *Controller) RemoveDeviceID(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "RemoveDeviceID() called")
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	emsg := c.validation.ValidateParam(vars["deviceID"], "eid")
	if len(emsg.Messages) != 0 {
		rw.WriteHeader(http.StatusNotAcceptable)
		level.Error(c.logger).Log( "validation", emsg.Messages)
		ToJSON(&emsg, rw)
		return
	}

	err := c.service.RemoveDeviceByID(vars["deviceID"], vars["networkName"])
	if err == nil {
		rw.WriteHeader(http.StatusNoContent)
	} else {
		level.Error(c.logger).Log( "error", err.Error())
		rw.WriteHeader(http.StatusNotFound)
		ToJSON(&GenericError{Message: err.Error()}, rw)
	}
}

// NetworkMessageRequest send any message over MQTT
type NetworkMessageRequest struct {
	Topic    string `json:"topic" validate:"required"`
	Qos      int    `json:"qos" validate:"required"`
	Retained string   `json:"retained" validate:"required"`
	Value    string `json:"value"`
}

// swagger:route POST /publishNetworkMessage network-tag networkMessageParams
// Return a specified device on the named network
// responses:
//	204: noContentResponse
//  404: genericError
//  406: validationError
//  422: genericError

// PublishNetworkMessage (dm DeviceMessage) topic, payload, qos, bRetained
func (c *Controller) PublishNetworkMessage(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "RemoveDeviceID() called")
	rw.Header().Add("Content-Type", "application/json")

	nmr := NetworkMessageRequest{}
	err := FromJSON(&nmr, r.Body)
	if err != nil {
		level.Error(c.logger).Log( "error", err.Error())
		rw.WriteHeader(http.StatusUnprocessableEntity)
		ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	// validate the product
	errs := c.validation.ValidateStruct(nmr)
	if len(errs.Messages) != 0 {
		c.logger.Log("validation", errs)

		// return the validation messages as an array
		rw.WriteHeader(http.StatusNotAcceptable)
		ToJSON(&errs, rw)
		return
	}

	tParts := strings.Split(nmr.Topic, "/")

	dm := dc.DeviceMessage{
		NetworkID: []byte(tParts[0]),
		DeviceID: []byte(tParts[1]),
		HomieType: dc.CoreTypePublishMessage,
		TopicS: nmr.Topic,
		Value: []byte(nmr.Value),
		Qosb: byte(nmr.Qos),
		RetainedB: strings.EqualFold(nmr.Retained, "true"),
	}

	c.service.PublishNetworkMessage(dm)
	rw.WriteHeader(http.StatusNoContent)
}
