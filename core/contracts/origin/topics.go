package origin

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

func init() {
	// set topics
	var err error

	parsedOrigin, err := abi.JSON(strings.NewReader(OriginMetaData.ABI))
	if err != nil {
		panic(err)
	}

	DispatchTopic = parsedOrigin.Events["Dispatch"].ID
}

// DispatchTopic is the topic that gets emitted when the dispatch event is called.
var DispatchTopic common.Hash

// topicMap maps events to topics.
// this is returned as a function to assert immutability.
func topicMap() map[EventType]common.Hash {
	return map[EventType]common.Hash{
		DispatchEvent: DispatchTopic,
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
