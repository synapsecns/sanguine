package rfq

import (
	"bytes"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/types/rfq"
)

func init() {
	var err error
	parsedRFQEvent, err := abi.JSON(strings.NewReader(SynapseRFQMetaData.ABI))
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
func TopicMap() map[rfq.EventType]common.Hash {
	return map[rfq.EventType]common.Hash{
		rfq.BridgeRelayedEvent:   BridgeRelayedTopic,
		rfq.BridgeRequestedEvent: BridgeRequestedTopic,
	}
}

// EventTypeFromTopic gets the event type from the topic
// returns nil if the topic is not found.
func EventTypeFromTopic(ogTopic common.Hash) *rfq.EventType {
	for eventType, topic := range TopicMap() {
		if bytes.Equal(ogTopic.Bytes(), topic.Bytes()) {
			return &eventType
		}
	}
	return nil
}

// Topic gets the topic from the event type.
func Topic(eventType rfq.EventType) common.Hash {
	topicHash, ok := TopicMap()[rfq.EventType(eventType.Int())]
	if !ok {
		panic("unknown event")
	}
	return topicHash
}
