package handlers

import (
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	"net/http"
)



// AllBroadcasts () []Broadcast
func (c *Controller) AllBroadcasts(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "AllBroadcasts() called")
	rw.Header().Add("Content-Type", "application/json")

	body := c.service.AllBroadcasts()

	rw.WriteHeader(http.StatusOK)
	err := ToJSON(body, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		level.Error(c.logger).Log( "error", err.Error())
	}
}

// BroadcastByID (broadcastID string) (Broadcast, error)
func (c *Controller) BroadcastByID(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "BroadcastByID() called")
	rw.Header().Add("Content-Type", "application/json")

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

// RemoveBroadcastID (broadcastID string)
func (c *Controller) RemoveBroadcastID(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "RemoveBroadcastID() called")
	rw.Header().Add("Content-Type", "application/json")

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