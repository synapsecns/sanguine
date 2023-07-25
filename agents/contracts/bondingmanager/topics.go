package bondingmanager

import (
	"bytes"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	// set topics
	var err error

	parsedBondingManager, err := abi.JSON(strings.NewReader(BondingManagerMetaData.ABI))
	if err != nil {
		panic(err)
	}

	StatusUpdatedTopic = parsedBondingManager.Events["StatusUpdated"].ID

	if StatusUpdatedTopic == (common.Hash{}) {
		panic("StatusUpdatedTopic is nil")
	}
}

// StatusUpdatedTopic is the topic that gets emitted when the StatusUpdated event is called.
var StatusUpdatedTopic common.Hash

// topicMap maps events to topics.
// this is returned as a function to assert immutability.
func topicMap() map[EventType]common.Hash {
	return map[EventType]common.Hash{
		StatusUpdatedEvent: StatusUpdatedTopic,
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
