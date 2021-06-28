// cmd/gui/main.go
// GUI Main using tView

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/pkg/deviceCore"
	"github.com/skoona/homie-service/pkg/services"
	cc "github.com/skoona/homie-service/pkg/utils"
)


func StartUp() (dc.CoreService, cc.Config) {
	ca, cb := services.Service()
	return *ca, *cb
}

func Shutdown() {
	services.Shutdown()
}

func main() {

	coreSvc, cfg := StartUp()
	siteNetworks := coreSvc.PrivateSiteNetworks()
	logger := log.With(cfg.Logger, "ui", "base")

	// implement gui here and in ./pkg/gui/{controllers,views,actions}

	/*
	 * Prepare for clean exit
	*/
	errs := make(chan error, 1)
	go func(shutdown chan error) {
		systemSignalChannel := make(chan os.Signal, 1)
		signal.Notify(systemSignalChannel, syscall.SIGINT, syscall.SIGTERM)
		sig := <-systemSignalChannel // wait on ctrl-c
		shutdown <- fmt.Errorf("%s", sig)
	}(errs)
	level.Info(logger).Log("event", "shutdown requested", "cause", <-errs)

	Shutdown()

	// Dump the SiteNetwork and all nodes as JSON
	out, err := json.MarshalIndent(siteNetworks, "", "  ")
	if err != nil {
		level.Warn(logger).Log("action", err.Error())
	} else {
		fmt.Println("Site Network")
		fmt.Println(string(out))
	}

	os.Exit(0)
}
