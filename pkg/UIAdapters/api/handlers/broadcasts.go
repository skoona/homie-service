package handlers

import (
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	"net/http"
)



// swagger:route GET /broadcasts Broadcast-Operations  broadcasts
// List all broadcasts
// responses:
//	202: broadcastsResponse
//  500: genericError

// Broadcasts () []Broadcast
func (c *Controller) Broadcasts(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "AllBroadcasts() called")

	body := c.service.AllBroadcasts()

	rw.WriteHeader(http.StatusAccepted)
	err := ToJSON(body, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		level.Error(c.logger).Log( "error", err.Error())
	}
}


// swagger:route GET /broadcastById/{broadcastID} Broadcast-Operations  broadcastById
// Get a specific broadcast
// responses:
//	202: broadcastResponse
//  404: genericError
//  406: validationError
//  500: genericError

// BroadcastByID (broadcastID string) (Broadcast, error)
func (c *Controller) BroadcastByID(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "BroadcastByID() called")

	vars := mux.Vars(r)
	emsg := c.validation.ValidateParam(vars["broadcastID"], "eid")
	if len(emsg.Messages) != 0 {
		rw.WriteHeader(http.StatusNotAcceptable)
		level.Error(c.logger).Log( "validation", emsg.Messages)
		ToJSON(&emsg, rw)
		return
	}

	body, err := c.service.BroadcastByID(vars["broadcastID"])
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


// swagger:route DELETE /removeBroadcastId/{broadcastID} Broadcast-Operations  removeBroadcastId
// Remove a specific broadcast
// responses:
//	204: noContentResponse
//  406: validationError
//  500: genericError

// RemoveBroadcastID (broadcastID string)
func (c *Controller) RemoveBroadcastID(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "RemoveBroadcastID() called")

	vars := mux.Vars(r)
	emsg := c.validation.ValidateParam(vars["broadcastID"], "eid")
	if len(emsg.Messages) != 0 {
		rw.WriteHeader(http.StatusNotAcceptable)
		level.Error(c.logger).Log( "validation", emsg.Messages)
		ToJSON(&emsg, rw)
		return
	}

	c.service.RemoveBroadcastByID(vars["broadcastID"])
	rw.WriteHeader(http.StatusNoContent)
}
