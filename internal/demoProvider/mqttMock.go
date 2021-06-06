package demoProvider

/*
 Mock Implementation of MQTT / PAHO
	-- use of tokens is not required and breaks

type Client interface {
	IsConnected() bool
	IsConnectionOpen() bool
	Connect() Token
	Disconnect(quiesce uint)
	Publish(topic string, qos byte, retained bool, payload interface{}) Token
	Subscribe(topic string, qos byte, callback MessageHandler) Token
	SubscribeMultiple(filters map[string]byte, callback MessageHandler) Token
	Unsubscribe(topics ...string) Token
	AddRoute(topic string, callback MessageHandler)
	OptionsReader() ClientOptionsReader
}

type Message interface {
    Duplicate() bool
    Qos() byte
    Retained() bool
    Topic() string
    MessageID() uint16
    Payload() []byte
    Ack()
}

type Token interface {
    Wait() bool
    WaitTimeout(time.Duration) bool
    Done() <-chan struct{}
    Error() error
}

MessageHandler = func(client mqtt.Client, msg mqtt.Message) {}
OnConnectHandler = func(client mqtt.Client){}
ConnectionLostHandler = func(client mqtt.Client, err error) {}

*/

import (
	"errors"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"sync"
)

type mockMessage struct {
	id       uint16
	topic    string
	qos      byte
	retained bool
	payload  []byte
}

type subscription struct {
	topics	[]string
	qos		byte
	callback mqtt.MessageHandler
}

type mockClient struct {
	sync.Mutex
	connected bool
	exit      chan bool

	subs map[string]subscription
	cbDefault   mqtt.MessageHandler // func(client mqtt.Client, msg mqtt.Message) {}
	cbOnConnectHandler  mqtt.OnConnectHandler // func(client mqtt.Client){}
	cbConnectionLostHandler  mqtt.ConnectionLostHandler //  func(client mqtt.Client, err error) {}
}

func NewMockClient() mqtt.Client {
	mClient = &mockClient{
		subs: map[string]subscription{},
	}
	return mClient
}

func newMockMessage(topic string, seqNum uint16, qos byte, retained bool, payload []byte) mqtt.Message {
	return &mockMessage{
		id:       seqNum,
		topic:    topic,
		qos:      qos,
		retained: retained,
		payload:  payload,
	}
}

func (m *mockMessage) Ack() {
	return
}

func (m *mockMessage) Duplicate() bool {
	return false
}

func (m *mockMessage) Qos() byte {
	return m.qos
}

func (m *mockMessage) Retained() bool {
	return m.retained
}

func (m *mockMessage) Topic() string {
	return m.topic
}

func (m *mockMessage) MessageID() uint16 {
	return m.id
}

func (m *mockMessage) Payload() []byte {
	return m.payload
}

func (mc *mockClient) addSubscription(t string, q byte, cb mqtt.MessageHandler) (string, subscription) {
	var topic string
	if strings.ContainsAny(t, "+#") {
		topic = strings.ReplaceAll(t, "+/", "")
		topic = strings.ReplaceAll(topic, "/#", "")
	} else {
		topic = t
	}
	return topic, subscription{
		topics: []string{topic},
		qos: q,
		callback: cb,
	}
}
func (mc *mockClient) addSubscriptions(filters map[string]byte, cb mqtt.MessageHandler) (string, subscription) {
	var topic string
	// filters are key=topic, byte=qos
	var key string = "key"
	var qos byte = 0
	tops := []string{}
	for k, v := range filters {
		if strings.ContainsAny(k, "+#") {
			topic = strings.ReplaceAll(k, "+/", "")
			topic = strings.ReplaceAll(topic, "/#", "")
		} else {
			topic = k
		}
		key = fmt.Sprintf("%s|%s", key, topic)
		tops = append(tops, topic)
		qos = v
	}
	return key, subscription{
		topics: tops,
		qos: qos,
		callback: cb,
	}
}

func (mc *mockClient) SetDefaultHandler( callback mqtt.MessageHandler) {
	mc.cbDefault = callback
}
func (mc *mockClient) SetConnHandler( callback mqtt.OnConnectHandler) {
	mc.cbOnConnectHandler = callback
}
func (mc *mockClient) SetConnLostHandler( callback mqtt.ConnectionLostHandler) {
	mc.cbConnectionLostHandler = callback
}
func (mc *mockClient) IsConnectionOpen() bool {
	return mc.IsConnected()
}
func (mc *mockClient) IsConnected() bool {
	mc.Lock()
	defer mc.Unlock()
	return mc.connected
}
func (mc *mockClient) Connect() mqtt.Token {
	mc.Lock()
	defer mc.Unlock()

	if mc.connected {
		return nil
	}

	mc.connected = true
	mc.exit = make(chan bool)

	mc.cbOnConnectHandler(mc)
	return &mqtt.ConnectToken{}
}
func (mc *mockClient) Disconnect(quiesce uint) {
	mc.Lock()
	defer mc.Unlock()

	if !mc.connected {
		return
	}

	mc.connected = false

	select {
	case <-mc.exit:
		return
	default:
		close(mc.exit)
	}

	mc.cbConnectionLostHandler(mc, errors.New("Unknown Value"))
}
func (mc *mockClient) Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
	mc.Lock()
	defer mc.Unlock()

	if !mc.connected {
		return nil
	}

	return &mqtt.PublishToken{}
}
func (mc *mockClient) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token {
	mc.Lock()
	defer mc.Unlock()

	if !mc.connected {
		return nil
	}

	k, s := mc.addSubscription(topic, qos, callback)
	mc.subs["key|" + k] = s

	return &mqtt.SubscribeToken{}
}
func (mc *mockClient) SubscribeMultiple(filters map[string]byte, callback mqtt.MessageHandler) mqtt.Token {
	mc.Lock()
	defer mc.Unlock()

	if !mc.connected {
		return nil
	}

	key, sub := mc.addSubscriptions(filters,  callback)
	mc.subs[key] = sub

	return &mqtt.SubscribeToken{}
}
func (mc *mockClient) Unsubscribe(topics ...string) mqtt.Token {
	var key string = "key"
	mc.Lock()
	defer mc.Unlock()
	if !mc.connected {
		return nil
	}

	for _, k := range topics {
		key = fmt.Sprintf("%s|%s", key, k)
	}
	delete(mc.subs, key)

	return &mqtt.UnsubscribeToken{}
}
func (mc *mockClient) AddRoute(topic string, callback mqtt.MessageHandler) {
	mc.Lock()
	defer mc.Unlock()
	mc.addSubscription(topic, 1, callback)
}
func (mc *mockClient) OptionsReader() mqtt.ClientOptionsReader {
	return mqtt.ClientOptionsReader{}
}

var (
	mClient    *mockClient
)
