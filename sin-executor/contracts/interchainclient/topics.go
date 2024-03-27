package interchainclient

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

var (
	// InterchainTransactionSentTopic is the event topic for a sent transaction.
	InterchainTransactionSentTopic common.Hash
)

// static checks to make sure topics actually exist.
func init() {
	var err error

	parsedABI, err := abi.JSON(strings.NewReader(InterchainClientV1MetaData.ABI))
	if err != nil {
		panic(err)
	}

	InterchainTransactionSentTopic = parsedABI.Events["InterchainTransactionSent"].ID

	_, err = parsedABI.EventByID(InterchainTransactionSentTopic)
	if err != nil {
		panic(err)
	}
}

// topicMap maps events to topics.
// this is returned as a function to assert immutability.
func topicMap() map[EventType]common.Hash {
	return map[EventType]common.Hash{
		InterchainTransactionSentEvent: InterchainTransactionSentTopic,
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
