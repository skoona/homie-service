package handlers

import (
	"encoding/json"
	"github.com/go-kit/kit/log"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"io"
	"net/http"
	"time"
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

func NewAPIServer(bindAddress string, router http.Handler, read, write, idle time.Duration) *http.Server {
// create a new server
	return &http.Server{
		Addr:         bindAddress, // configure the bind address
		Handler:      router,                  // set the default handler
		ReadTimeout:  read * time.Second,     // max time to read request from the client
		WriteTimeout: write * time.Second,    // max time to write response to the client
		IdleTimeout:  idle * time.Second,   // max time for connections using TCP Keep-Alive
	}
}

// NewApiController returns a new products handler with the given logger
func NewApiController(s *dc.CoreService, l *log.Logger, v *Validation) *Controller {
	ctrl = &Controller{
		service: *s,
		logger: *l,
		validation: v,
	}
	return ctrl
}

// NoContent No content is expected
type NoContent struct {
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
