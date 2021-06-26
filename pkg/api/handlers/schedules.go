package handlers

import (
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"net/http"
)


// AllSchedules  []Schedule
func (c *Controller) AllSchedules(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "AllSchedules() called")
	rw.Header().Add("Content-Type", "application/json")

	body := c.service.AllSchedules()

	rw.WriteHeader(http.StatusOK)
	err := ToJSON(body, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		level.Error(c.logger).Log( "error", err.Error())
	}
}

// ScheduleByDeviceID (deviceID string) Schedule
func (c *Controller) ScheduleByDeviceID(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "ScheduleByDeviceID() called")
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	body := c.service.ScheduleByDeviceID(vars["deviceID"])
	if body.ElementType == dc.CoreTypeSchedule {
		rw.WriteHeader(http.StatusOK)
		err := ToJSON(body, rw)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			level.Error(c.logger).Log( "error", err.Error())
		}
	} else {
		level.Error(c.logger).Log( "error", "schedule not found")
		rw.WriteHeader(http.StatusNotFound)
		ToJSON(&GenericError{Message: "schedule not found"}, rw)
	}
}

// ScheduleByID (scheduleID string) Schedule
func (c *Controller) ScheduleByID(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "ScheduleByID() called")
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	body := c.service.ScheduleByID(vars["scheduleID"])
	if body.ElementType == dc.CoreTypeSchedule {
		rw.WriteHeader(http.StatusOK)
		err := ToJSON(body, rw)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			level.Error(c.logger).Log( "error", err.Error())
		}
	} else {
		level.Error(c.logger).Log( "error", "schedule not found")
		rw.WriteHeader(http.StatusNotFound)
		ToJSON(&GenericError{Message: "schedule not found"}, rw)
	}
}

// CreateScheduleRequest  create a new firmware OTA schedule
type CreateScheduleRequest struct {
	NetworkName string          `json:"networkName" validate:"required"`
	DeviceID    string          `json:"deviceID" validate:"required"`
	Transport   int             `json:"transportType" validate:"required"`
	FirmwareID  string          `json:"firmwareID" validate:"required"`
}
// CreateScheduleResponse returns id of created schedule
type CreateScheduleResponse struct {
	ScheduleID string `json:"scheduleID"`
}
// CreateSchedule (networkName string, deviceID string, transport OTATransport, firmwareID EID) (string, error)
func (c *Controller) CreateSchedule(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "CreateSchedule() called")
	rw.Header().Add("Content-Type", "application/json")

	csr := CreateScheduleRequest{}
	err := FromJSON(&csr, r.Body)
	if err != nil {
		level.Error(c.logger).Log( "error", err.Error())
		rw.WriteHeader(http.StatusUnprocessableEntity)
		ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	// validate the product
	errs := c.validator.Validate(csr)
	if len(errs) != 0 {
		c.logger.Log("validation", errs)

		// return the validation messages as an array
		rw.WriteHeader(http.StatusUnprocessableEntity)
		ToJSON(&ValidationErrorMessage{Messages: errs.Errors()}, rw)
		return
	}

	body, err := c.service.CreateSchedule(csr.NetworkName, csr.DeviceID, dc.OTATransport(csr.Transport), dc.EID(csr.FirmwareID))
	if err == nil {
		rw.WriteHeader(http.StatusOK)
		err := ToJSON(CreateScheduleResponse{body}, rw)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			level.Error(c.logger).Log( "error", err.Error())
		}
	} else {
		level.Error(c.logger).Log( "error", err.Error())
		rw.WriteHeader(http.StatusUnprocessableEntity)
		ToJSON(&GenericError{Message: err.Error()}, rw)
	}
}


// RemoveSchedule (scheduleID string)
func (c *Controller) RemoveSchedule(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "RemoveSchedule() called")
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	err := c.service.RemoveSchedule(vars["scheduleID"])
	if err == nil {
		rw.WriteHeader(http.StatusNoContent)
	} else {
		level.Error(c.logger).Log( "error", err.Error())
		rw.WriteHeader(http.StatusNotFound)
		ToJSON(&GenericError{Message: err.Error()}, rw)
	}
}
