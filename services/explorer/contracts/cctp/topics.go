package cctp

import (
	"bytes"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/types/cctp"
)

func init() {
	var err error
	parsedCCTPEvent, err := abi.JSON(strings.NewReader(SynapseCCTPMetaData.ABI))
	if err != nil {
		panic(err)
	}

	CircleRequestSentTopic = parsedCCTPEvent.Events["CircleRequestSent"].ID

	CircleRequestFulfilledTopic = parsedCCTPEvent.Events["CircleRequestFulfilled"].ID
}

// CircleRequestSentTopic is when a Circle token is sent with an attached action request.
var CircleRequestSentTopic common.Hash

// CircleRequestFulfilledTopic is when a Circle token is received with an attached action request.
var CircleRequestFulfilledTopic common.Hash

// TopicMap maps events to topics.
// this is returned as a function to assert immutability.
func TopicMap() map[cctp.EventType]common.Hash {
	return map[cctp.EventType]common.Hash{
		cctp.CircleRequestSentEvent:      CircleRequestSentTopic,
		cctp.CircleRequestFulfilledEvent: CircleRequestFulfilledTopic,
	}
}

// EventTypeFromTopic gets the event type from the topic
// returns nil if the topic is not found.
func EventTypeFromTopic(ogTopic common.Hash) *cctp.EventType {
	for eventType, topic := range TopicMap() {
		if bytes.Equal(ogTopic.Bytes(), topic.Bytes()) {
			return &eventType
		}
	}
	return nil
}

// Topic gets the topic from the event type.
func Topic(eventType cctp.EventType) common.Hash {
	topicHash, ok := TopicMap()[cctp.EventType(eventType.Int())]
	if !ok {
		panic("unknown event")
	}
	return topicHash
}
