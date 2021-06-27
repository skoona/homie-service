package handlers

import (
	"encoding/json"
	"github.com/go-kit/kit/log"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"io"
)

// CtxKeyOne is a key used for the Product object in the context
type CtxKeyOne struct{}

// Controller handler for getting and updating API Items
type Controller struct {
	_ struct{}
	service dc.CoreService
	logger log.Logger
	validation *Validation
}

var (
	// Controller interface to Core Service
	ctrl *Controller
)

// NewApiController returns a new products handler with the given logger
func NewApiController(s *dc.CoreService, l *log.Logger, v *Validation) *Controller {
	ctrl = &Controller{
		service: *s,
		logger: *l,
		validation: v,
	}
	return ctrl
}


// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
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

