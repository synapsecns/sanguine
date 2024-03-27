package executionservice

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

var (
	// ExecutionRequestedTopic is the event topic for a sent transaction.
	ExecutionRequestedTopic common.Hash
)

// static checks to make sure topics actually exist.
func init() {
	var err error

	parsedABI, err := abi.JSON(strings.NewReader(ExecutionServiceMetaData.ABI))
	if err != nil {
		panic(err)
	}

	ExecutionRequestedTopic = parsedABI.Events["ExecutionRequested"].ID

	_, err = parsedABI.EventByID(ExecutionRequestedTopic)
	if err != nil {
		panic(err)
	}
}

// topicMap maps events to topics.
// this is returned as a function to assert immutability.
func topicMap() map[EventType]common.Hash {
	return map[EventType]common.Hash{
		ExecutionRequestedEvent: ExecutionRequestedTopic,
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
