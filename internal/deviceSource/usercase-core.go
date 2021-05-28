package deviceSource

/*
  deviceSource/usecase-core.go:

  DeviceSource Service Implementation
*/

import (
	"errors"
	"github.com/go-kit/kit/log/level"
	dc "github.com/skoona/homie-service/internal/deviceCore"
	sch "github.com/skoona/homie-service/internal/deviceScheduler"
)

type (
	//Incoming Messages
	DeviceMessageHandler interface {
		CreateDemoDeviceMessage(topic string, payload []byte, idCounter uint16, retained bool, qos byte) (dc.DeviceMessage, error)
		CreateQueueDeviceMessage(qmsg dc.QueueMessage) (dc.DeviceMessage, error)
		GetProviderRequestChannel() (chan dc.DeviceMessage, error)
		GetProviderResponseChannel() (chan dc.DeviceMessage, error)
		FromMQTTProvider(qmsg dc.QueueMessage) error
		FromDemoProvider(topic string, payload []byte, idCounter uint16, retained bool, qos byte) error
	}
)


func (dmHandler *deviceMessageHandler) CreateDemoDeviceMessage(topic string, payload []byte, idCounter uint16, retained bool, qos byte) (dc.DeviceMessage, error) {
	level.Debug(dmHandler.logger).Log("method", "CreateDemoDeviceMessage() called")
	return dc.NewDeviceMessage(topic, payload, idCounter, retained, qos)
}

func (dmHandler *deviceMessageHandler) CreateQueueDeviceMessage(qmsg dc.QueueMessage) (dc.DeviceMessage, error) {
	level.Debug(dmHandler.logger).Log("method", "CreateQueueDeviceMessage() called")
	return dc.NewQueueMessage(qmsg)
}

func (dmHandler *deviceMessageHandler) FromMQTTProvider(qmsg dc.QueueMessage) error {
	level.Debug(dmHandler.logger).Log("method", "FromMQTTProvider() called", "msgID", qmsg.MessageID())
	dm, err := dmHandler.CreateQueueDeviceMessage(qmsg)
	if err != nil {
		level.Error(dmHandler.logger).Log("error", err.Error())
	} else {
		ch, err := dmHandler.GetProviderRequestChannel()
		if err == nil {
			ch <- dm // send it to channel
		}
	}
	return err
}
func (dmHandler *deviceMessageHandler) FromDemoProvider(topic string, payload []byte, idCounter uint16, retained bool, qos byte) error {
	level.Debug(dmHandler.logger).Log("method", "FromDemoProvider() called", "msgID", idCounter)
	dm, err := dmHandler.CreateDemoDeviceMessage(topic, payload, idCounter, retained, qos)
	if err != nil {
		level.Error(dmHandler.logger).Log("error", err.Error())
	} else {
		ch, err := dmHandler.GetProviderRequestChannel()
		if err == nil {
			ch <- dm // send it to channel
		}
	}
	return err
}
