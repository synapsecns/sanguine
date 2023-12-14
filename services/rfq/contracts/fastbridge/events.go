package fastbridge

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

// TODO: consider not exporting to avoid accidental mutation
var (
	// BridgeRequestedTopic is the event topic for a bridge request
	BridgeRequestedTopic common.Hash
	// BridgeRelayedTopic is the topic emitted by a bridge relay
	BridgeRelayedTopic common.Hash
)

// static checks to make sure topics actually exist
func init() {
	var err error

	parsedABI, err := abi.JSON(strings.NewReader(FastBridgeMetaData.ABI))
	if err != nil {
		panic(err)
	}

	BridgeRequestedTopic = parsedABI.Events["BridgeRequestedEvent"].ID
	BridgeRelayedTopic = parsedABI.Events["BridgeRelayedEvent"].ID

	_, err = parsedABI.EventByID(BridgeRequestedTopic)
	if err != nil {
		panic(err)
	}

	_, err = parsedABI.EventByID(BridgeRelayedTopic)
	if err != nil {
		panic(err)
	}

}

// topicMap maps events to topics.
// this is returned as a function to assert immutability.
func topicMap() map[EventType]common.Hash {
	return map[EventType]common.Hash{
		BridgeRequestedEvent: BridgeRequestedTopic,
		BridgeRelayedEvent:   BridgeRelayedTopic,
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
