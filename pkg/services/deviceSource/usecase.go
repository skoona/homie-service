package deviceSource

/*
  deviceSource/usecase.go:

  DeviceSource Service Implementation
*/

import (
	"fmt"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	cc "github.com/skoona/homie-service/pkg/utils"
)

type (
	// Service Implementation
	deviceSource struct {
		repository Repositiory
		cfg        cc.Config
		logger     log.Logger
	}
)

// Retained DeviceSource Service, once created
var dvService deviceSource

func (s *deviceSource) ApplyCoreEvent(dm *DeviceMessage) error {
	var err error
	logger := log.With(s.logger, "method", "ApplyCoreEvent")
	err = s.repository.Store(dm)
	if err != nil {
		level.Error(logger).Log("error", err)
	}

	logger.Log("DeviceID ", dm.DeviceID)

	return err
}

func (s *deviceSource) ApplyOTAEvent(dm *DeviceMessage) error {
	var err error
	logger := log.With(s.logger, "method", "ApplyOTAEvent")

	err = s.repository.Store(dm)
	if err != nil {
		level.Error(logger).Log("error", err)
	}

	logger.Log("DeviceID ", dm.DeviceID)

	return err
}

func (s *deviceSource) ApplyDiscoveryEvent(dm *DeviceMessage) error {
	var err error
	logger := log.With(s.logger, "method", "ApplyDiscoveryEvent")

	err = s.repository.Store(dm)
	if err != nil {
		level.Error(logger).Log("error", err)
	}

	logger.Log("DeviceID ", dm.DeviceID)

	return err
}

/*
   String ToString
   - utility funciton
*/
func (dm *DeviceMessage) String() string {
	level.Debug(dvService.logger).Log("DeviceMessage", "String()")
	return fmt.Sprintf("id=%06d retained=%-5t device=%-16s node=%-16s property=%-16s pProperty=%-16s attr=%-16s value=%s",
		dm.ID, dm.Retained, dm.DeviceID, dm.NodeID, dm.PropertyID, dm.PPropertyID, dm.AttributeID, dm.Value)
}

// Schedulable()
func (dm *DeviceMessage) Schedulable() bool {
	level.Debug(dvService.logger).Log("DeviceMessage", "Schedulable()")
	res := false
	for _, keys := range []string{"$state", "$online", "$fw", "$implementation"} {
		if strings.Contains(dm.Topic, keys) {
			res = true
			break
		}
	}
	return res
}

// Settable() determine is property is settable
func (dm *DeviceMessage) Settable() bool {
	level.Debug(dvService.logger).Log("DeviceMessage", "Settable()")
	return strings.HasSuffix(dm.Topic, "set")
}

// OTAactive() determines if device is downloading firmware
func (dm *DeviceMessage) OTAactive() bool {
	level.Debug(dvService.logger).Log("DeviceMessage", "OTAactive()")
	return strings.Contains(dm.Topic, "$implementation/ota/status") &&
		strings.HasPrefix(string(dm.Value), "206")
}

// Broadcast() determines if this is a Homie Broadcast message
func (dm *DeviceMessage) Broadcast() bool {
	level.Debug(dvService.logger).Log("DeviceMessage", "Broadcast()")
	return dm.Parts()[1] == "$broadcast"
}

// Parts() returns the individual parts of the original MQTT message
func (dm *DeviceMessage) Parts() []string {
	level.Debug(dvService.logger).Log("DeviceMessage", "Parts()")
	return strings.Split(dm.Topic, "/")
}

// PartsLen() returns nuber of parts in Topic
func (dm *DeviceMessage) PartsLen() int {
	level.Debug(dvService.logger).Log("DeviceMessage", "PartsLen()")
	return len(dm.Parts())
}
