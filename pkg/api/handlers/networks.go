package handlers

import (
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"net/http"
	"strings"
)

// AllNetworks  SiteNetworks
func (c *Controller) AllNetworks(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "AllNetworks() called")
	rw.Header().Add("Content-Type", "application/json")

	body := c.service.AllNetworks()

	rw.WriteHeader(http.StatusOK)
	err := ToJSON(body, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		level.Error(c.logger).Log( "error", err.Error())
	}
}

// NetworkByName (networkName string) Network
func (c *Controller) NetworkByName(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "NetworkByName() called")
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	body := c.service.NetworkByName(vars["networkName"])

	rw.WriteHeader(http.StatusOK)
	err := ToJSON(body, rw)
	if err != nil {
		level.Error(c.logger).Log( "error", err.Error())
		rw.WriteHeader(http.StatusNotFound)
		ToJSON(&GenericError{Message: err.Error()}, rw)
	}
}

// DeviceByName (deviceName, networkName string) (Device, error)
func (c *Controller) DeviceByName(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "DeviceByName() called")
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	body, err := c.service.DeviceByName(vars["deviceName"], vars["networkName"])
	if err == nil {
		rw.WriteHeader(http.StatusOK)
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


// DeviceByID (deviceID string, networkName string) (Device, error)
func (c *Controller) DeviceByID(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "DeviceByID() called")
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	body, err := c.service.DeviceByID(vars["deviceID"], vars["networkName"])
	if err == nil {
		rw.WriteHeader(http.StatusOK)
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

// RemoveDeviceID (deviceID string, networkName string) error
func (c *Controller) RemoveDeviceID(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "RemoveDeviceID() called")
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

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
	errs := c.validator.Validate(nmr)
	if len(errs) != 0 {
		c.logger.Log("validation", errs)

		// return the validation messages as an array
		rw.WriteHeader(http.StatusUnprocessableEntity)
		ToJSON(&ValidationErrorMessage{Messages: errs.Errors()}, rw)
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
