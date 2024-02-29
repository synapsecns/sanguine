package lightmanager

import (
	"bytes"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	// set topics
	var err error

	parsedLightManager, err := abi.JSON(strings.NewReader(LightManagerMetaData.ABI))
	if err != nil {
		panic(err)
	}

	DisputeOpenedTopic = parsedLightManager.Events["DisputeOpened"].ID
	RootUpdatedTopic = parsedLightManager.Events["RootUpdated"].ID

	if DisputeOpenedTopic == (common.Hash{}) {
		panic("DisputeOpenedTopic is nil")
	}
}

// DisputeOpenedTopic is the topic that gets emitted when the DisputeOpened event is called.
var DisputeOpenedTopic common.Hash

// RootUpdatedTopic is the topic that gets emitted when the RootUpdates event is called.
var RootUpdatedTopic common.Hash

// topicMap maps events to topics.
// this is returned as a function to assert immutability.
func topicMap() map[EventType]common.Hash {
	return map[EventType]common.Hash{
		DisputeOpenedEvent: DisputeOpenedTopic,
		RootUpdatedEvent:   RootUpdatedTopic,
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
