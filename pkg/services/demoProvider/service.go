package demoProvider

/**
  demoProvider/service.go:

  Generate Demo DeviceMessages

*/

import (
	"bufio"
	"context"
	"os"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dss "github.com/skoona/homie-service/pkg/services/deviceSource"
	cc "github.com/skoona/homie-service/pkg/utils"
)

var ctx context.Context
var deviceService dss.Service
var logger log.Logger

/**
 * ProduceDeviceMessages
 * Generate DeviceMessages from a demo logfile
 * by converting them to DeviceMessages
 * outputs to device channels
 */
func produceDeviceMessages(demoFile string, svc dss.Service, logger log.Logger) (int, error) {
	level.Debug(logger).Log("called", "produceDeviceMessages()")
	/*
	 * Create a Go Routine for the MQTT Channel to
	 * convert msgs to DeviceMessages and output to dvcSyncChannels
	 */
	// go func(dsChan chan dss.DeviceMessage) {
	var err error
	var file *os.File
	var idx uint16 = 0
	var dm dss.DeviceMessage

	file, err = os.OpenFile(demoFile, os.O_RDONLY, 0666)
	if err != nil {
		return 0, err
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
		// dm, err := buildDeviceMessage(topic, []byte(payload), idx, false, 1)
		dm, err = dss.NewDeviceMessage(topic, []byte(payload), idx, false, 1)
		if err != nil {
			level.Error(logger).Log("error", err.Error())
		} else {
			// dsChan <- dm // send it to channel
			svc.ApplyCoreEvent(&dm)
		}
		time.Sleep(100 * time.Millisecond) // slow the pace
	}

	// }(dvcSyncChannel, config)

	return int(idx), nil
}

/*
 * Start()
 *
 * Initialize this service
 */
func Start(contx context.Context, svc dss.Service) error {
	var err error
	ctx = contx
	deviceService = svc
	demoFile := ctx.Value(cc.DbConfig).(cc.DBConfig).DemoSource
	logger = log.With(ctx.Value(cc.AppConfig).(cc.Config).Logger, "pkg", "demoProvider")
	level.Debug(logger).Log("msg", "Calling Start()", "demoFile", demoFile)

	demoCount, err := produceDeviceMessages(demoFile, svc, logger)

	level.Debug(logger).Log("sent", demoCount)
	return err
}

/*
 * Stop
 * Cleans up this service
 */
func Stop() {
	level.Debug(logger).Log("msg", "Calling Stop()")

}
