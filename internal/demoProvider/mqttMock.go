package demoProvider

/*
 Mock Implementation of MQTT / PAHO

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
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"time"
)


type mockToken struct {
	ready bool
	err   error
	complete chan struct{}
}

func NewToken(tf bool) *mockToken {
	return &mockToken{
		ready: tf,
		err: nil,
		complete: make(chan struct{}),
	}
}

func (mt *mockToken) Done() <-chan struct{} {
	close(mt.complete)
	return mt.complete
}
func (mt *mockToken) Wait() bool {
	return mt.ready
}
func (mt *mockToken) WaitTimeout(time.Duration) bool {
	return mt.ready
}
func (mt *mockToken) flowComplete() {
	mt.ready = true
	return
}
func (mt *mockToken) Error() error {
	return mt.err
}
func (mt *mockToken) SetError(err error) {
	close(mt.complete)
	mt.err = err
}

type subscription struct {
	topics	[]string
	qos		byte
	callback mqtt.MessageHandler
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

type mockClient struct {
	isConnected bool
	subscriptions map[string]subscription
	token       *mockToken
	cbDefault   mqtt.MessageHandler // func(client mqtt.Client, msg mqtt.Message) {}
	cbOnConnectHandler  mqtt.OnConnectHandler // func(client mqtt.Client){}
	cbConnectionLostHandler  mqtt.ConnectionLostHandler //  func(client mqtt.Client, err error) {}
}


func NewMockClient() mqtt.Client {

	mClient = &mockClient{
		isConnected:   true,
		token:         NewToken(false),
		subscriptions: map[string]subscription{},
	}
	return mClient
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
	return mc.isConnected
}
func (mc *mockClient) Connect() mqtt.Token {
	mc.token.Done()
	mc.token = NewToken(true)
	mc.isConnected = true
	mc.cbOnConnectHandler(mc)
	mc.token.Done()
	return mc.token
}
func (mc *mockClient) Disconnect(quiesce uint) {
	mc.isConnected = false
	mc.token.Done()
	mc.token = NewToken(false)

	mc.cbConnectionLostHandler(mc, mc.token.Error())
}
func (mc *mockClient) Publish(topic string, qos byte, retained bool, payload interface{}) mqtt.Token {
	mc.token.Done()
	return mc.token
}
func (mc *mockClient) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) mqtt.Token {
	k, s := mc.addSubscription(topic, qos, callback)
	mc.subscriptions["key|" + k] = s
	return mc.token
}
func (mc *mockClient) SubscribeMultiple(filters map[string]byte, callback mqtt.MessageHandler) mqtt.Token {
	key, sub := mc.addSubscriptions(filters,  callback)
	mc.subscriptions[key] = sub
	return mc.token
}
func (mc *mockClient) Unsubscribe(topics ...string) mqtt.Token {
	var key string = "key"
	for _, k := range topics {
		key = fmt.Sprintf("%s|%s", key, k)
	}
	delete(mc.subscriptions, key)
	mc.token.Done()
	return mc.token
}
func (mc *mockClient) AddRoute(topic string, callback mqtt.MessageHandler) {
	mc.addSubscription(topic, 1, callback)
}
func (mc *mockClient) OptionsReader() mqtt.ClientOptionsReader {
	panic("implement me")
}

var (
	mClient    *mockClient
)
