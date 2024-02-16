package synapsemodule

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

var (
	// VerificationRequestedTopic is the event topic for a verification request.
	VerificationRequestedTopic common.Hash
	// EntryVerificationTopic is the topic emitted by a verification.
	EntryVerificationTopic common.Hash
)

// static checks to make sure topics actually exist.
func init() {
	var err error

	parsedABI, err := abi.JSON(strings.NewReader(SynapseModuleMetaData.ABI))
	if err != nil {
		panic(err)
	}

	VerificationRequestedTopic = parsedABI.Events["VerificationRequested"].ID
	EntryVerificationTopic = parsedABI.Events["EntryVerified"].ID

	_, err = parsedABI.EventByID(VerificationRequestedTopic)
	if err != nil {
		panic(err)
	}

	_, err = parsedABI.EventByID(EntryVerificationTopic)
	if err != nil {
		panic(err)
	}
}

// topicMap maps events to topics.
// this is returned as a function to assert immutability.
func topicMap() map[EventType]common.Hash {
	return map[EventType]common.Hash{
		VerificationRequestedEvent: VerificationRequestedTopic,
		EntryVerificationEvent:     EntryVerificationTopic,
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
