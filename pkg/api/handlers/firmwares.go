package handlers

import (
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"net/http"
)


// AllFirmwares a list of  []Firmware
func (c *Controller) AllFirmwares(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "AllFirmwares() called")
	rw.Header().Add("Content-Type", "application/json")

	body := c.service.AllFirmwares()

	rw.WriteHeader(http.StatusOK)
	err := ToJSON(body, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		level.Error(c.logger).Log( "error", err.Error())
	}
}

// CreateFirmwareRequest upload a new firmware
type CreateFirmwareRequest struct {
	SrcFile string `json:"srcFile" validate:"required"`
	DstFile string `json:"dstFile" validate:"required"`
}
// CreateFirmwareResponse id of newly created firmware
type CreateFirmwareResponse struct {
	FirmwareID string `json:"firmwareID"`
}

// CreateFirmware (srcFile, dstFile string) (EID, error)
func (c *Controller) CreateFirmware(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "CreateFirmware() called")
	rw.Header().Add("Content-Type", "application/json")


	cfr := CreateFirmwareRequest{}
	err := FromJSON(&cfr, r.Body)
	if err != nil {
		level.Error(c.logger).Log( "error", err.Error())
		rw.WriteHeader(http.StatusUnprocessableEntity)
		ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	// validate the product
	errs := c.validator.Validate(cfr)
	if len(errs) != 0 {
		c.logger.Log("validation", errs)

		// return the validation messages as an array
		rw.WriteHeader(http.StatusUnprocessableEntity)
		ToJSON(&ValidationErrorMessage{Messages: errs.Errors()}, rw)
		return
	}

	body, err := c.service.CreateFirmware(cfr.SrcFile, cfr.DstFile)
	if err == nil {
		rw.WriteHeader(http.StatusOK)
		err := ToJSON(CreateFirmwareResponse{string(body)}, rw)
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

// FirmwareByID (firmwareID EID) (Firmware, error)
func (c *Controller) FirmwareByID(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "FirmwareByID() called")
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	body, err := c.service.FirmwareByID(dc.EID(vars["firmwareID"]))
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

// FirmwareByName (firmwareName string) (Firmware, error)
func (c *Controller) FirmwareByName(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "FirmwareByName() called")
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	body, err := c.service.FirmwareByName(vars["firmwareName"])
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

// RemoveFirmwareID (firmwareID EID)
func (c *Controller) RemoveFirmwareID(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "RemoveFirmwareByID() called")
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	c.service.RemoveFirmwareByID(dc.EID(vars["firmwareID"]))
	rw.WriteHeader(http.StatusNoContent)
}
