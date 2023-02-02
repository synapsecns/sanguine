package testmetaswap

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
	"strings"
)

func init() {
	var err error
	parsedSwap, err := abi.JSON(strings.NewReader(TestMetaSwapMetaData.ABI))
	if err != nil {
		panic(err)
	}

	// we do this here to throw a compile error if the event is not found
	TokenSwapUnderlyingTopic = parsedSwap.Events["TokenSwapUnderlying"].ID
}

// TokenSwapUnderlyingTopic is the topic used for token swap underlying.
var TokenSwapUnderlyingTopic common.Hash

// TopicMap maps events to topics.
// this is returned as a function to assert immutability.
func TopicMap() map[swap.EventType]common.Hash {
	return map[swap.EventType]common.Hash{
		swap.TokenSwapUnderlyingEvent: TokenSwapUnderlyingTopic,
	}
}

// EventTypeFromTopic gets the event type from the topic
// returns nil if the topic is not found.
func EventTypeFromTopic(ogTopic common.Hash) *swap.EventType {
	for eventType, topic := range TopicMap() {
		if bytes.Equal(ogTopic.Bytes(), topic.Bytes()) {
			return &eventType
		}
	}
	return nil
}

// Topic gets the topic from the event type.
func Topic(eventType swap.EventType) common.Hash {
	topicHash, ok := TopicMap()[swap.EventType(eventType.Int())]
	if !ok {
		panic("unknown event")
	}
	return topicHash
}
