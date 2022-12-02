package attestationcollector

import (
	"bytes"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	// set topics
	var err error

	parsedAttestationCollector, err := abi.JSON(strings.NewReader(AttestationCollectorABI))
	if err != nil {
		panic(err)
	}

	AttestationAcceptedTopic = parsedAttestationCollector.Events["AttestationAccepted"].ID
}

// AttestationAcceptedTopic is the topic that gets emitted
// when the AttestationSubmitted event is called.
var AttestationAcceptedTopic common.Hash

// topicMap maps events to topics.
// this is returned as a function to assert immutability.
func topicMap() map[EventType]common.Hash {
	return map[EventType]common.Hash{
		AttestationAcceptedEvent: AttestationAcceptedTopic,
	}
}

// eventTypeFromTopic gets the event type from the topic
// returns nil if the topic is not found.
func eventTypeFromTopic(ogTopic common.Hash) *EventType {
	for eventType, topic := range topicMap() {
		if bytes.Equal(ogTopic.Bytes(), topic.Bytes()) {
			return &eventType
		}
	}
	return nil
}
