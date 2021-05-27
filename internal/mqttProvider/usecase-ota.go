package mqttProvider

/**
  mqttProvider/usecase-ota.go:

  Handle Interaction with Scheduler
  - Send OTA Triggers
  - Execute OTA to device
  - Watch OTA Progress
  - Send Status
*/

/**
 * ConsumeOTAMessages
 * Handles OTA Streaming responses
 */
func ConsumeOTAMessages(network, name string, enabled bool) error {
	var err error
	if enabled {
		err = WatchOTAProgress(network, name)
	} else {
		err = UnWatchOTAProgress(network, name)
	}

	if nil != err {
		level.Error(logger).Log("event", "Subscribe Failed", "error", err.Error())
	}
	level.Debug(logger).Log("ConsumeOTAMessages Completed", enabled, "network", network, "name", name)

	return err
}

// WatchOTAProgress for Scheduler
func WatchOTAProgress(network, device string) error {
	topic := fmt.Sprintf("%s/%s/$implementation/ota/status", network, device)
	token := subWithHandler(topic, otaResponses) // OTA Responses
	return token.Error()
}

// UnWatchOTAProgress for Scheduler
func UnWatchOTAProgress(network, device string) error {
	topic := fmt.Sprintf("%s/%s/$implementation/ota/status", network, device)
	token := client.Unsubscribe(topic)
	token.Wait()
	return token.Error()
}

func establishOTAChannels() {
	var err error
	// one use is to start an OTA, or cancel an OTA
	fromOTAService, err = s.GetOTAResponseChannel()
	if err != nil {
		level.Error(logger).Log("event", "OTA respnse channel offline", "error", err.Error())
		client.Disconnect(250)
		panic(err.Error())
	}
}
