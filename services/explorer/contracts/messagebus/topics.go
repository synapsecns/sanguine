package messagebus

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/types/messagebus"
	"strings"
)

func init() {
	var err error
	parsedMessage, err := abi.JSON(strings.NewReader(MessageBusUpgradeableMetaData.ABI))
	if err != nil {
		panic(err)
	}

	ExecutedTopic = parsedMessage.Events["Executed"].ID

	MessageSentTopic = parsedMessage.Events["MessageSent"].ID

	CallRevertedTopic = parsedMessage.Events["CallReverted"].ID
}

// ExecutedTopic is the topic used for receiving messages.
var ExecutedTopic common.Hash

// MessageSentTopic is the topic used for sending messages.
var MessageSentTopic common.Hash

// CallRevertedTopic is the topic used for checking reverted calls.
var CallRevertedTopic common.Hash

// TopicMap maps events to topics.
// this is returned as a function to assert immutability.
func TopicMap() map[messagebus.EventType]common.Hash {
	return map[messagebus.EventType]common.Hash{
		messagebus.ExecutedEvent:     ExecutedTopic,
		messagebus.MessageSentEvent:  MessageSentTopic,
		messagebus.CallRevertedEvent: CallRevertedTopic,
	}
}

// EventTypeFromTopic gets the event type from the topic
// returns nil if the topic is not found.
func EventTypeFromTopic(ogTopic common.Hash) *messagebus.EventType {
	for eventType, topic := range TopicMap() {
		if bytes.Equal(ogTopic.Bytes(), topic.Bytes()) {
			return &eventType
		}
	}
	return nil
}

// Topic gets the topic from the event type.
func Topic(eventType messagebus.EventType) common.Hash {
	topicHash, ok := TopicMap()[messagebus.EventType(eventType.Int())]
	if !ok {
		panic("unknown event")
	}
	return topicHash
}
