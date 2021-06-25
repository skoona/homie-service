package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"io"
	"net/http"
	"strconv"
)

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

// Controller handler for getting and updating API Items
type Controller struct {
	_ struct{}
	service dc.CoreService
	logger log.Logger
	validator *Validation
}

var (
	// Controller interface to Core Service
	ctrl *Controller

	// ErrInvalidPath is an error message when the product path is not valid
	ErrInvalidPath = fmt.Errorf("Invalid Path, path should be /products/[id]")
)

// NewApiController returns a new products handler with the given logger
func NewApiController(s *dc.CoreService, l *log.Logger, v *Validation) *Controller {
	ctrl = &Controller{
		service: *s,
		logger: *l,
		validator: v,
	}
	return ctrl
}


// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// getIDParam returns the ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getIDParam(r *http.Request) int {
	// parse the  id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}


// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// FromJSON deserializes the object from JSON string
// in an io.Reader to the given interface
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
