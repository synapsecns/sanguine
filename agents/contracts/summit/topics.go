package summit

import (
	"bytes"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	// set topics
	var err error

	parsedsummit, err := abi.JSON(strings.NewReader(SummitMetaData.ABI))
	if err != nil {
		panic(err)
	}

	AttestationSavedTopic = parsedsummit.Events["AttestationSaved"].ID

	if AttestationSavedTopic == (common.Hash{}) {
		panic("AttestationSavedTopic is nil")
	}
}

// AttestationSavedTopic is the topic that gets emitted when the AttestationSaved event is called.
var AttestationSavedTopic common.Hash

// topicMap maps events to topics.
// this is returned as a function to assert immutability.
func topicMap() map[EventType]common.Hash {
	return map[EventType]common.Hash{
		AttestationSavedEvent: AttestationSavedTopic,
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
