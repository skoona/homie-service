package handlers

import (
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"net/http"
	"strconv"
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

// CreateSchedule (networkName string, deviceID string, transport OTATransport, firmwareID EID) (string, error)
// TODO write method
func (c *Controller) CreateSchedule(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "CreateSchedule() called")
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	trans, _ := strconv.Atoi(vars["transport"])

	body, err := c.service.CreateSchedule(
		vars["networkName"],
		vars["deviceID"],
		dc.OTATransport(trans),
		dc.EID(vars["firmwareID"]),
		)
	if err == nil {
		rw.WriteHeader(http.StatusOK)
		err := ToJSON(body, rw)
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

	c.service.RemoveSchedule(vars["scheduleID"])
	rw.WriteHeader(http.StatusNoContent)
}
