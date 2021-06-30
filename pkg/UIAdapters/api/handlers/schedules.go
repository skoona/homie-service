package handlers

import (
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"net/http"
)


// swagger:route GET /schedules Schedule-Operations schedules
// List all device schedules
// responses:
//	202: schedulesResponse
//  500: genericError

// AllSchedules  []Schedule
func (c *Controller) AllSchedules(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "AllSchedules() called")

	body := c.service.AllSchedules()

	rw.WriteHeader(http.StatusAccepted)
	err := ToJSON(body, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		level.Error(c.logger).Log( "error", err.Error())
	}
}

// swagger:route GET /scheduleByDeviceId/{deviceID} Schedule-Operations  scheduleByDeviceId
// Get OTA schedule for a specific device
// responses:
//	202: scheduleResponse
//  404: genericError
//  406: validationError
//  500: genericError

// ScheduleByDeviceID (deviceID string) Schedule
func (c *Controller) ScheduleByDeviceID(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "ScheduleByDeviceID() called")

	vars := mux.Vars(r)
	emsg := c.validation.ValidateParam(vars["deviceID"], "eid")
	if len(emsg.Messages) != 0 {
		rw.WriteHeader(http.StatusNotAcceptable)
		level.Error(c.logger).Log( "validation", emsg.Messages)
		ToJSON(&emsg, rw)
		return
	}

	body := c.service.ScheduleByDeviceID(vars["deviceID"])
	if body.ElementType == dc.CoreTypeSchedule {
		rw.WriteHeader(http.StatusAccepted)
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

// swagger:route GET /scheduleById/{scheduleID} Schedule-Operations  scheduleById
// Get a specific schedule using its unique id
// responses:
//	202: scheduleResponse
//  404: genericError
//  406: validationError
//  500: genericError

// ScheduleByID (scheduleID string) Schedule
func (c *Controller) ScheduleByID(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "ScheduleByID() called")

	vars := mux.Vars(r)
	emsg := c.validation.ValidateParam(vars["scheduleID"], "eid")
	if len(emsg.Messages) != 0 {
		rw.WriteHeader(http.StatusNotAcceptable)
		level.Error(c.logger).Log( "validation", emsg.Messages)
		ToJSON(&emsg, rw)
		return
	}

	body := c.service.ScheduleByID(vars["scheduleID"])
	if body.ElementType == dc.CoreTypeSchedule {
		rw.WriteHeader(http.StatusAccepted)
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
	DeviceID    string          `json:"deviceID" validate:"required,eid"`
	Transport   int             `json:"transportType" validate:"required"`
	FirmwareID  string          `json:"firmwareID" validate:"required,eid"`
}
// CreateScheduleResponse returns id of created schedule
type CreateScheduleResponse struct {
	ScheduleID string `json:"scheduleID"`
}

// swagger:route POST /createSchedule Schedule-Operations  createSchedule
// Create a new device OTA schedule
// responses:
//	202: scheduleIdResponse
//  404: genericError
//  406: validationError
//  422: genericError
//  500: genericError

// CreateSchedule (networkName string, deviceID string, transport OTATransport, firmwareID EID) (string, error)
func (c *Controller) CreateSchedule(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "CreateSchedule() called")

	csr := CreateScheduleRequest{}
	err := FromJSON(&csr, r.Body)
	if err != nil {
		level.Error(c.logger).Log( "error", err.Error())
		rw.WriteHeader(http.StatusUnprocessableEntity)
		ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	// validate the product
	errs := c.validation.ValidateStruct(csr)
	if len(errs.Messages) != 0 {
		c.logger.Log("validation", errs)

		// return the validation messages as an array
		rw.WriteHeader(http.StatusNotAcceptable)
		ToJSON(&errs, rw)
		return
	}

	body, err := c.service.CreateSchedule(csr.NetworkName, csr.DeviceID, dc.OTATransport(csr.Transport), dc.EID(csr.FirmwareID))
	if err == nil {
		rw.WriteHeader(http.StatusAccepted)
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

// swagger:route DELETE /removeScheduleId/{scheduleID} Schedule-Operations removeScheduleId
// Remove a device from the named network
// responses:
//	204: noContentResponse
//  404: genericError
//  406: validationError

// RemoveSchedule (scheduleID string)
func (c *Controller) RemoveSchedule(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "RemoveSchedule() called")

	vars := mux.Vars(r)
	emsg := c.validation.ValidateParam(vars["scheduleID"], "eid")
	if len(emsg.Messages) != 0 {
		rw.WriteHeader(http.StatusNotAcceptable)
		level.Error(c.logger).Log( "validation", emsg.Messages)
		ToJSON(&emsg, rw)
		return
	}

	err := c.service.RemoveSchedule(vars["scheduleID"])
	if err == nil {
		rw.WriteHeader(http.StatusNoContent)
	} else {
		level.Error(c.logger).Log( "error", err.Error())
		rw.WriteHeader(http.StatusNotFound)
		ToJSON(&GenericError{Message: err.Error()}, rw)
	}
}
