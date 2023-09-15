package inbox

import (
	"bytes"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	// set topics
	var err error

	parsedInbox, err := abi.JSON(strings.NewReader(InboxMetaData.ABI))
	if err != nil {
		panic(err)
	}

	SnapshotAcceptedTopic = parsedInbox.Events["SnapshotAccepted"].ID

	ReceiptAcceptedTopic = parsedInbox.Events["ReceiptAccepted"].ID

	if SnapshotAcceptedTopic == (common.Hash{}) {
		panic("SnapshotAcceptedTopic is nil")
	}
	if ReceiptAcceptedTopic == (common.Hash{}) {
		panic("ReceiptAcceptedTopic is nil")
	}
}

// SnapshotAcceptedTopic is the topic that gets emitted
// when the SnapshotAccepted event is called.
var SnapshotAcceptedTopic common.Hash

// ReceiptAcceptedTopic is the topic that gets emitted
// when the ReceiptAccepted event is called.
var ReceiptAcceptedTopic common.Hash

// topicMap maps events to topics.
// this is returned as a function to assert immutability.
func topicMap() map[EventType]common.Hash {
	return map[EventType]common.Hash{
		SnapshotAcceptedEvent: SnapshotAcceptedTopic,
		ReceiptAcceptedEvent:  ReceiptAcceptedTopic,
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
