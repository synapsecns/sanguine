package destination

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/types/destination"
	"strings"
)

func init() {
	var err error
	parsedDestination, err := abi.JSON(strings.NewReader(DestinationMetaData.ABI))
	if err != nil {
		panic(err)
	}
	AttestationAcceptedTopic = parsedDestination.Events["AttestationAccepted"].ID
	AgentRootAcceptedTopic = parsedDestination.Events["AgentRootAccepted"].ID
}

var AttestationAcceptedTopic common.Hash
var AgentRootAcceptedTopic common.Hash

// TopicMap maps events to topics.
// this is returned as a function to assert immutability.
func TopicMap() map[destination.EventType]common.Hash {
	return map[destination.EventType]common.Hash{
		destination.AttestationAcceptedEvent: AttestationAcceptedTopic,
		destination.AgentRootAcceptedEvent:   AgentRootAcceptedTopic}
}

// EventTypeFromTopic gets the event type from the topic
// returns nil if the topic is not found.
func EventTypeFromTopic(ogTopic common.Hash) *destination.EventType {
	for eventType, topic := range TopicMap() {
		if bytes.Equal(ogTopic.Bytes(), topic.Bytes()) {
			return &eventType
		}
	}
	return nil
}

// Topic gets the topic from the event type.
func Topic(eventType destination.EventType) common.Hash {
	topicHash, ok := TopicMap()[destination.EventType(eventType.Int())]
	if !ok {
		panic("unknown event")
	}
	return topicHash
}
