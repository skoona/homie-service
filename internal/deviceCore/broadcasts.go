package deviceCore

/**
	deviceCore/broadcasts.go

The design goal for this file is:
	* record received broadcast messages

*/

import (
	"crypto/md5"
	"fmt"
	"time"
)

// Broadcast Alert Messages on  Network
type Broadcast struct {
	ID          string
	ElementType CoreType
	Parent      string
	Topic       string
	Level       string
	Value       string
	Received    time.Time
}

// newBroadcast Creates Component
func NewBroadcast(parent, topic, level, value string) Broadcast {
	bc := Broadcast{
		ID:          fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s.%s.%s", parent, topic, level)))),
		ElementType: CoreTypeBroadcast,
		Parent:      parent,
		Topic:       topic,
		Level:       level,
		Value:       value,
		Received:    time.Now(),
	}

	return bc
}
