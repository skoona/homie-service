package demoProvider

/**
  demoProvider/service.go:

  Generate Demo DeviceMessages

*/

import (
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/pkg/services/deviceCore"
	dss "github.com/skoona/homie-service/pkg/services/deviceSource"
	cc "github.com/skoona/homie-service/pkg/utils"
)

var (
	cfg           cc.Config
	deviceService dss.Service
	logger        log.Logger
	fromDMService chan dc.DeviceMessage // in
	toDMService   chan dc.DeviceMessage // out
)

/**
 * ProduceDeviceMessages
 * Generate DeviceMessages from a demo logfile
 * by converting them to DeviceMessages
 * outputs to device channels
 */
func produceDeviceMessages(demoFile string, consumer chan dc.DeviceMessage, logger log.Logger) {
	level.Debug(logger).Log("event", "calling produceDeviceMessages()")
	/*
	 * Create a Go Routine for the MQTT Channel to
	 * convert msgs to DeviceMessages and output to dvcSyncChannels
	 */
	go func(dsChan chan dc.DeviceMessage, filepath string, logger log.Logger) {
		var err error
		var file *os.File
		var idx uint16 = 0
		var dm dc.DeviceMessage

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
			dm, err = dc.NewDeviceMessage(topic, []byte(payload), idx, false, 1)
			if err != nil {
				level.Error(logger).Log("error", err.Error())
			} else {
				dsChan <- dm // send it to channel
			}
			time.Sleep(100 * time.Millisecond) // slow the pace
		}

		level.Debug(logger).Log("event", "produceDeviceMessages(gofunc()) completed")
	}(consumer, demoFile, logger)

	level.Debug(logger).Log("event", "produceDeviceMessages() active")
}

/*
 * Start()
 *
 * Initialize this service
 */
func Start(dfg cc.Config, svc dss.Service) ([]string, error) {
	var err error
	cfg = dfg
	deviceService = svc
	demoFile := cfg.Dbc.DemoSource
	logger = log.With(cfg.Logger, "pkg", "demoProvider")
	level.Debug(logger).Log("event", "Calling Start()", "demoFile", demoFile, "demoNetworks", dfg.Dbc.DemoNetworks)

	// Initialize a Message Channel
	fromDMService, toDMService, err = dss.ChannelsForDMProviders()
	if err != nil {
		level.Error(logger).Log("event", "Channels offline", "error", err.Error())
		panic(err.Error())
	}

	produceDeviceMessages(demoFile, toDMService, logger)

	level.Debug(logger).Log("event", "Start() completed")

	return dfg.Dbc.DemoNetworks, err
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")
	if nil != toDMService {
		close(toDMService) // only the send should shutdown channels
	}

	level.Debug(logger).Log("event", "Stop() completed")
}
