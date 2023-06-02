package destination

import (
	"bytes"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	// set topics
	var err error

	parsedDestination, err := abi.JSON(strings.NewReader(DestinationMetaData.ABI))
	if err != nil {
		panic(err)
	}

	ExecutedTopic = parsedDestination.Events["Executed"].ID

	if ExecutedTopic == (common.Hash{}) {
		panic("ExecutedTopic is nil")
	}
}

// AttestationAcceptedTopic is the topic that gets emitted when the AttestationAccepted event is called.
var AttestationAcceptedTopic common.Hash

// ExecutedTopic is the topic that gets emitted when the Executed event is called.
var ExecutedTopic common.Hash

// topicMap maps events to topics.
// this is returned as a function to assert immutability.
func topicMap() map[EventType]common.Hash {
	return map[EventType]common.Hash{
		ExecutedEvent: ExecutedTopic,
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
