package origin

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/types/origin"
	"strings"
)

func init() {
	var err error
	parsedOrigin, err := abi.JSON(strings.NewReader(OriginMetaData.ABI))
	if err != nil {
		panic(err)
	}

	// we do this here to throw a compile error if the event is not found
	SentTopic = parsedOrigin.Events["Sent"].ID

}

// SentTopic is the topic a sent topic.
var SentTopic common.Hash

// OwnershipTransferredTopic is the topic used for adding liquidity.
var OwnershipTransferredTopic common.Hash

// TopicMap maps events to topics.
// this is returned as a function to assert immutability.
func TopicMap() map[origin.EventType]common.Hash {
	return map[origin.EventType]common.Hash{
		origin.SentEvent: SentTopic}
}

// EventTypeFromTopic gets the event type from the topic
// returns nil if the topic is not found.
func EventTypeFromTopic(ogTopic common.Hash) *origin.EventType {
	for eventType, topic := range TopicMap() {
		if bytes.Equal(ogTopic.Bytes(), topic.Bytes()) {
			return &eventType
		}
	}
	return nil
}

// Topic gets the topic from the event type.
func Topic(eventType origin.EventType) common.Hash {
	topicHash, ok := TopicMap()[origin.EventType(eventType.Int())]
	if !ok {
		panic("unknown event")
	}
	return topicHash
}
