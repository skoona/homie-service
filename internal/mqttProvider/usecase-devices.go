package mqttProvider

import "github.com/go-kit/kit/log/level"

/**
  mqttProvider/usecase-devices.go:

  Handle Interaction with deviceSource
  - Send network state messages
  - Listen for network delete messages
  - store/update network traffic to repository
*/

func establishProviderChannels() {
	var err error
	// Initialize a Message Channel
	// one use is to delete devices based on ui input
	fromDMService, err = s.GetProviderResponseChannel()
	if err != nil {
		level.Error(logger).Log("event", "provider response channel offline", "error", err.Error())
		client.Disconnect(250)
		panic(err.Error())
	}
}
