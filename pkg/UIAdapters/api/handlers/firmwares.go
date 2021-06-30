package handlers

import (
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"net/http"
)


// swagger:route GET /firmwares Firmware-Operations firmwares
// List all firmwares on file
// responses:
//	202: firmwaresResponse
//  404: genericError
//  406: validationError
//  500: genericError

// AllFirmwares a list of  []Firmware
func (c *Controller) AllFirmwares(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "AllFirmwares() called")

	body := c.service.AllFirmwares()

	rw.WriteHeader(http.StatusAccepted)
	err := ToJSON(body, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		level.Error(c.logger).Log( "error", err.Error())
	}
}


// swagger:route POST /createFirmware Firmware-Operations createFirmware
// Add a new homie firmware to the collection
// responses:
//	202: firmwareIdResponse
//  404: genericError
//  406: validationError
//  500: genericError

// CreateFirmwareRequest upload a new firmware
type CreateFirmwareRequest struct {
	SrcFile string `json:"srcFile" validate:"required,fwfn"`
	DstFile string `json:"dstFile" validate:"required,fwfn"`
}
// CreateFirmwareResponse id of newly created firmware
type CreateFirmwareResponse struct {
	FirmwareID string `json:"firmwareID"`
}

// CreateFirmware (srcFile, dstFile string) (EID, error)
func (c *Controller) CreateFirmware(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "CreateFirmware() called")

	cfr := CreateFirmwareRequest{}
	err := FromJSON(&cfr, r.Body)
	if err != nil {
		level.Error(c.logger).Log( "error", err.Error())
		rw.WriteHeader(http.StatusUnprocessableEntity)
		ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	// validate the product
	errs := c.validation.ValidateStruct(cfr)
	if len(errs.Messages) != 0 {
		c.logger.Log("validation", errs)

		// return the validation messages as an array
		rw.WriteHeader(http.StatusNotAcceptable)
		ToJSON(&errs, rw)
		return
	}

	body, err := c.service.CreateFirmware(cfr.SrcFile, cfr.DstFile)
	if err == nil {
		rw.WriteHeader(http.StatusAccepted)
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


// swagger:route GET /firmwareById/{firmwareID} Firmware-Operations firmwareById
// Get specific firmware details
// responses:
//	202: firmwareResponse
//  404: genericError
//  406: validationError
//  500: genericError

// FirmwareByID (firmwareID EID) (Firmware, error)
func (c *Controller) FirmwareByID(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "FirmwareByID() called")

	vars := mux.Vars(r)
	emsg := c.validation.ValidateParam(vars["firmwareID"], "eid")
	if len(emsg.Messages) != 0 {
		rw.WriteHeader(http.StatusNotAcceptable)
		level.Error(c.logger).Log( "validation", emsg.Messages)
		ToJSON(&emsg, rw)
		return
	}

	body, err := c.service.FirmwareByID(dc.EID(vars["firmwareID"]))
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


// swagger:route GET /firmwareByName/{firmwareName} Firmware-Operations firmwareByName
// Get specific firmware details
// responses:
//	202: firmwareResponse
//  404: genericError
//  500: genericError

// FirmwareByName (firmwareName string) (Firmware, error)
func (c *Controller) FirmwareByName(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "FirmwareByName() called")

	vars := mux.Vars(r)
	body, err := c.service.FirmwareByName(vars["firmwareName"])
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


// swagger:route DELETE /removeFirmwareId/{firmwareID} Firmware-Operations removeFirmwareId
// delete a specific firmware and filesystem file
// responses:
//	204: noContentResponse
//  406: validationError

// RemoveFirmwareID (firmwareID EID)
func (c *Controller) RemoveFirmwareID(rw http.ResponseWriter, r *http.Request) {
	level.Debug(c.logger).Log( "api-method", "RemoveFirmwareByID() called")

	vars := mux.Vars(r)
	emsg := c.validation.ValidateParam(vars["firmwareID"], "eid")
	if len(emsg.Messages) != 0 {
		rw.WriteHeader(http.StatusNotAcceptable)
		level.Error(c.logger).Log( "validation", emsg.Messages)
		ToJSON(&emsg, rw)
		return
	}

	c.service.RemoveFirmwareByID(dc.EID(vars["firmwareID"]))
	rw.WriteHeader(http.StatusNoContent)
}
