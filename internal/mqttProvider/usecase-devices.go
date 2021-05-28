package mqttProvider

/**
  mqttProvider/usecase-devices.go:

  Handle Interaction with deviceSource
  - Send network state messages
  - Listen for network delete messages
  - store/update network traffic to repository
*/
import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	cc "github.com/skoona/homie-service/internal/utils"
)

type (
	DeviceStream interface {
		GetConsumerRequestChannel() chan dc.QueueMessage
		GetConsumerNotificationChannel() chan dc.QueueMessage
	}

	deviceStream struct {
		cfg         cc.MQTTConfig
		consumerNotify chan dc.DeviceMessage // in
		consumerRequest chan dc.DeviceMessage // in
		logger         log.Logger
	}
)

var (
	dStr *deviceStream
)

/*
 * DeviceStream Implementation */

func NewDeviceStream(dfg cc.Config, plog log.Logger) DeviceStream {
	dStr = &deviceStream{
		cfg: dfg,
		logger: log.With(plog, "service", "DeviceStream"),
		consumerNotify: make(chan dc.QueueMessage, 256),
		consumerRequest: make(chan dc.QueueMessage, 120),
	}
	return dStr
}

func (s *deviceStream) GetConsumerRequestChannel() chan dc.QueueMessage {
	level.Debug(s.logger).Log("method", "GetConsumerRequestChannel()")
	return s.consumerRequest
}
func (s *deviceStream) GetConsumerNotificationChannel() chan dc.QueueMessage {
	level.Debug(s.logger).Log("method", "GetConsumerNotificationChannel()")
	return s.consumerNotify
}


func establishStreamChannels() {
	var err error
	// Initialize a Message Channel
	// one use is to delete devices based on ui input
	dStr.consumerNotify, err = s.GetProviderResponseChannel()
	if err != nil {
		level.Error(logger).Log("event", "provider response channel offline", "error", err.Error())
		client.Disconnect(250)
		panic(err.Error())
	}
}
