package demoProvider

/**
  demoProvider/service.go:

  Generate Demo DeviceMessages

*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dss "github.com/skoona/homie-service/internal/deviceSource"
	cc "github.com/skoona/homie-service/internal/utils"
)

var (
	cfg    cc.Config
	logger log.Logger
	dmh    dss.DeviceMessageHandler
)

/**
 * ProduceDeviceMessages
 * Generate DeviceMessages from a demo logfile
 * by converting them to DeviceMessages
 * outputs to device channels
 */
func produceDeviceMessages(demoFile string, logger log.Logger) {
	level.Debug(logger).Log("event", "calling produceDeviceMessages()")
	/*
	 * Create a Go Routine for the MQTT Channel to
	 * convert msgs to DeviceMessages and output to dvcSyncChannels
	 */
	go func(filepath string, logger log.Logger) {
		var err error
		var file *os.File
		var idx uint16 = 0

		file, err = os.OpenFile(filepath, os.O_RDONLY, 0666)
		if err != nil {
			level.Error(logger).Log("error", err.Error())
			panic(err.Error())
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if err = scanner.Err(); err != nil {
				continue
			}

			parts := strings.Split(line, " ")
			topic := parts[0]
			payload := strings.Join(parts[1:], " ")

			idx++
			err = dmh.FromDemoProvider(topic, []byte(payload), idx, false, 1)
			if err != nil {
				level.Error(logger).Log("error", err.Error())
			}
			time.Sleep(100 * time.Millisecond) // slow the pace
		}

		level.Debug(logger).Log("event", "produceDeviceMessages(gofunc()) completed")
	}(demoFile, logger)

	level.Debug(logger).Log("event", "produceDeviceMessages() active")
}

/*
 * Start()
 *
 * Initialize this service
 */
func Start(s dss.DeviceMessageHandler) error {
	var err error
	// ensure Initialize() is called first
	if logger == nil {
		panic(fmt.Errorf("you must call Initialize() in this package before calling Start()"))
	}

	dmh = s

	demoFile := cfg.Dbc.DemoSource
	level.Debug(logger).Log("event", "Calling Start()", "demoFile", demoFile, "demoNetworks", strings.Join(cfg.Dbc.DemoNetworks, ","))

	produceDeviceMessages(demoFile, logger)

	level.Debug(logger).Log("event", "Start() completed")

	return err
}

/*
 * Initialize()
 *
 * Initialize this service
 */
func Initialize(dfg cc.Config) ([]string, error) {
	var err error
	cfg = dfg
	logger = log.With(cfg.Logger, "pkg", "demoProvider")
	level.Debug(logger).Log("event", "calling Initialize()")

	level.Debug(logger).Log("event", "Initialize() completed", "networks discovered", strings.Join(dfg.Dbc.DemoNetworks, ","))
	return dfg.Dbc.DemoNetworks, err
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")

	level.Debug(logger).Log("event", "Stop() completed")
}
