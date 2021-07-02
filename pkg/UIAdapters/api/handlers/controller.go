package handlers

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	cc "github.com/skoona/homie-service/pkg/utils"
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
	httpServer *http.Server
	cfg        *cc.Config
}

var (
	// Controller interface to Core Service
	ctrl *Controller
)

// NewApiProvider returns a fully configured http api engine
func NewApiProvider(cfg *cc.Config, s *dc.CoreService, l *log.Logger) *Controller {
	ctrl = &Controller{
		service: *s,
		logger: *l,
		cfg: cfg,
	}
	ctrl.validation = NewValidation()
	ctrl.establishHttpServer() // configure routes and server
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

// establishHttpServer configure routes and server
func (c *Controller) establishHttpServer() {

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()
	lmw := LoggingMiddleware(c.logger)
	sm.Use(lmw)
	sm.Use(CommonMiddleware)
	sm.Use( gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"})))

	// handlers for API
	getR := sm.PathPrefix("/api/v1").
		Methods(http.MethodGet).
		Subrouter()

	getR.HandleFunc("/networks", c.AllNetworks)
	getR.HandleFunc("/network/{networkName:[a-zA-Z]+}", c.NetworkByName)
	getR.HandleFunc("/devices/{networkName:[a-zA-Z]+}", c.NetworkDevices)
	getR.HandleFunc("/deviceByName/{networkName:[a-zA-Z]+}/{deviceName:[a-zA-Z]+}", c.DeviceByName)
	getR.HandleFunc("/deviceById/{networkName:[a-zA-Z]+}/{deviceID:[a-zA-Z0-9]+}", c.DeviceByID)

	getR.HandleFunc("/schedules", c.AllSchedules)
	getR.HandleFunc("/scheduleById/{scheduleID:[a-zA-Z0-9]+}", c.ScheduleByID)
	getR.HandleFunc("/scheduleByDeviceId/{deviceID:[a-zA-Z0-9]+}", c.ScheduleByDeviceID)

	getR.HandleFunc("/firmwares", c.AllFirmwares)
	getR.HandleFunc("/firmwareByName/{firmwareName:[a-zA-Z0-9]+}", c.FirmwareByName)
	getR.HandleFunc("/firmwareById/{firmwareID:[a-zA-Z0-9]+}", c.FirmwareByID)

	getR.HandleFunc("/broadcasts", c.Broadcasts)
	getR.HandleFunc("/broadcastById/{broadcastID:[a-zA-Z0-9]+}", c.BroadcastByID)

	delR := sm.PathPrefix("/api/v1").
		Methods(http.MethodDelete).
		Subrouter()
	delR.HandleFunc("/removeDeviceId/{networkName:[a-zA-Z]+}/{deviceID:[a-zA-Z0-9]+}", c.RemoveDeviceID)
	delR.HandleFunc("/removeFirmwareId/{firmwareID:[a-zA-Z0-9]+}", c.RemoveFirmwareID)
	delR.HandleFunc("/removeScheduleId/{scheduleID:[a-zA-Z0-9]+}", c.RemoveSchedule)
	delR.HandleFunc("/removeBroadcastId/{broadcastID:[a-zA-Z0-9]+}", c.RemoveBroadcastID)

	poR := sm.PathPrefix("/api/v1").
		Methods(http.MethodPost).
		Subrouter()
	poR.HandleFunc("/createFirmware", c.CreateFirmware) // has req/rsp
	poR.HandleFunc("/createSchedule", c.CreateSchedule) // has req/rsp
	poR.HandleFunc("/publishNetworkMessage", c.PublishNetworkMessage) // has req/noc

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	gDocs := sm.Methods(http.MethodGet).Subrouter()
	gDocs.Handle("/docs", sh)
	gDocs.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	c.httpServer = &http.Server{
		Addr:         c.cfg.Api.BindAddress, // configure the bind address
		Handler:      sm,                  // set the default handler
		ReadTimeout:  5 * time.Second,     // max time to read request from the client
		WriteTimeout: 10 * time.Second,    // max time to write response to the client
		IdleTimeout:  120 * time.Second,   // max time for connections using TCP Keep-Alive
	}
}
// Run start the httpserver running
func (c *Controller) Run() error {
	return c.httpServer.ListenAndServe()
}
// Shutdown stops the http server
func (c *Controller) Shutdown() error {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	return c.httpServer.Shutdown(ctx)
}
