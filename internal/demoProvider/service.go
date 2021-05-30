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
	dc "github.com/skoona/homie-service/internal/deviceCore"
	dss "github.com/skoona/homie-service/internal/deviceSource"
	cc "github.com/skoona/homie-service/internal/utils"
)

type (
	deviceStream struct {
		notifyChannel chan dc.DeviceMessage // in
		logger        log.Logger
	}
)

var (
	cfg     cc.Config
	dStream *deviceStream
	logger  log.Logger
)

/*
 * DeviceStream Implementation */

func NewStreamProvider(plog log.Logger) dss.StreamProvider {
	dStream = &deviceStream{
		logger: log.With(plog, "service", "StreamProvider"),
	}
	return dStream
}
func (s *deviceStream) ActivateNotifications() chan dc.DeviceMessage {
	demoFile := cfg.Dbc.DemoSource
	enableNetworkTraffic(demoFile, s.logger, s.GetNotifyChannel())
	return s.notifyChannel
}
func (s *deviceStream) CreateDemoDeviceMessage(topic string, payload []byte, idCounter uint16, retained bool, qos byte) dc.DeviceMessage {
	level.Debug(s.logger).Log("method", "CreateDemoDeviceMessage() called")
	dm, _ := dc.NewDeviceMessage(topic, payload, idCounter, retained, qos)
	return dm
}

func (s *deviceStream) CreateQueueDeviceMessage(qmsg dc.QueueMessage) dc.DeviceMessage {
	level.Debug(s.logger).Log("method", "CreateQueueDeviceMessage() called")
	dm, _ := dc.NewQueueMessage(qmsg)
	return dm
}

func (s *deviceStream) GetPublishChannel() chan dc.DeviceMessage {
	level.Debug(s.logger).Log("method", "GetPublishChannel()")
	return nil
}
func (s *deviceStream) GetNotifyChannel() chan dc.DeviceMessage {
	level.Debug(s.logger).Log("method", "GetNotifyChannel()")
	if s.notifyChannel == nil {
		s.notifyChannel = make(chan dc.DeviceMessage, 200)
	}
	return s.notifyChannel
}

/**
 * enableNetworkTraffic()
 * Generate DeviceMessages from a demo logfile
 * by converting them to DeviceMessages
 * outputs to device channels
 */
func enableNetworkTraffic(demoFile string, plog log.Logger, notify chan dc.DeviceMessage) {
	level.Debug(plog).Log("event", "calling enableNetworkTraffic()")
	/*
	 * Create a Go Routine for the MQTT Channel to
	 * convert msgs to DeviceMessages and output to dvcSyncChannels
	 */
	go func(filepath string, tlog log.Logger, publish chan dc.DeviceMessage) {
		var err error
		var file *os.File
		var idx uint16 = 0

		file, err = os.OpenFile(filepath, os.O_RDONLY, 0666)
		if err != nil {
			level.Error(tlog).Log("error", err.Error())
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
			publish <- dStream.CreateDemoDeviceMessage(topic, []byte(payload), idx, false, 1)

			time.Sleep(100 * time.Millisecond) // slow the pace
		}

		level.Debug(tlog).Log("event", "enableNetworkTraffic(gofunc()) completed")
	}(demoFile, plog, notify)

	level.Debug(plog).Log("event", "enableNetworkTraffic() active")
}

/*
 * Start()
 *
 * Initialize this service
 */
func Start() error {
	var err error
	// ensure Initialize() is called first
	if logger == nil {
		panic(fmt.Errorf("you must call Initialize() in this package before calling Start()"))
	}

	demoFile := cfg.Dbc.DemoSource
	level.Debug(logger).Log("event", "Calling Start()", "demoFile", demoFile, "demoNetworks", strings.Join(cfg.Dbc.DemoNetworks, ","))

	// first call to Activate will start things off

	level.Debug(logger).Log("event", "Start() completed")

	return err
}

/*
 * Initialize()
 *
 * Initialize this service
 */
func Initialize(dfg cc.Config) (dss.StreamProvider, []string, error) {
	var err error
	cfg = dfg
	logger = log.With(cfg.Logger, "pkg", "demoProvider")
	level.Debug(logger).Log("event", "calling Initialize()")

	NewStreamProvider(logger) // creates stream provider

	level.Debug(logger).Log("event", "Initialize() completed", "networks discovered", strings.Join(dfg.Dbc.DemoNetworks, ","))
	return dStream, dfg.Dbc.DemoNetworks, err
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(logger).Log("event", "Calling Stop()")

	if dStream.notifyChannel != nil {
		close(dStream.notifyChannel) // we own chan, cleanup when done
		dStream.notifyChannel = nil
	}

	time.Sleep(time.Second)
	level.Debug(logger).Log("event", "Stop() completed")
}
