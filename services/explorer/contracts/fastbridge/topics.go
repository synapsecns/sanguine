package fastbridge

import (
	"bytes"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/types/fastbridge"
)

func init() {
	var err error
	parsedRFQEvent, err := abi.JSON(strings.NewReader(FastBridgeMetaData.ABI))
	if err != nil {
		panic(err)
	}

	BridgeRequestedTopic = parsedRFQEvent.Events["BridgeRequested"].ID

	BridgeRelayedTopic = parsedRFQEvent.Events["BridgeRelayed"].ID
}

// CircleRequestSentTopic is when a Circle token is sent with an attached action request.
var BridgeRequestedTopic common.Hash

// CircleRequestFulfilledTopic is when a Circle token is received with an attached action request.
var BridgeRelayedTopic common.Hash

// TopicMap maps events to topics.
// this is returned as a function to assert immutability.
func TopicMap() map[fastbridge.EventType]common.Hash {
	return map[fastbridge.EventType]common.Hash{
		fastbridge.BridgeRelayedEvent:   BridgeRelayedTopic,
		fastbridge.BridgeRequestedEvent: BridgeRequestedTopic,
	}
}

// EventTypeFromTopic gets the event type from the topic
// returns nil if the topic is not found.
func EventTypeFromTopic(ogTopic common.Hash) *fastbridge.EventType {
	for eventType, topic := range TopicMap() {
		if bytes.Equal(ogTopic.Bytes(), topic.Bytes()) {
			return &eventType
		}
	}
	return nil
}

// Topic gets the topic from the event type.
func Topic(eventType fastbridge.EventType) common.Hash {
	topicHash, ok := TopicMap()[fastbridge.EventType(eventType.Int())]
	if !ok {
		panic("unknown event")
	}
	return topicHash
}
