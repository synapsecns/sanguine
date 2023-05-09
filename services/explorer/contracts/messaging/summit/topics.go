package summit

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/types/summit"
	"strings"
)

func init() {
	var err error
	summitOrigin, err := abi.JSON(strings.NewReader(SummitMetaData.ABI))
	if err != nil {
		panic(err)
	}

	// we do this here to throw a compile error if the event is not found
	ReceiptAcceptedTopic = summitOrigin.Events["ReceiptAccepted"].ID
	SnapshotAcceptedTopic = summitOrigin.Events["SnapshotAccepted"].ID
	ReceiptConfirmedTopic = summitOrigin.Events["ReceiptConfirmed"].ID
	TipAwardedTopic = summitOrigin.Events["TipAwarded"].ID

}

var ReceiptAcceptedTopic common.Hash
var SnapshotAcceptedTopic common.Hash
var ReceiptConfirmedTopic common.Hash
var TipAwardedTopic common.Hash

// TopicMap maps events to topics.
// this is returned as a function to assert immutability.
func TopicMap() map[summit.EventType]common.Hash {
	return map[summit.EventType]common.Hash{
		summit.ReceiptAcceptedEvent:  ReceiptAcceptedTopic,
		summit.SnapshotAcceptedEvent: SnapshotAcceptedTopic,
		summit.ReceiptConfirmedEvent: ReceiptConfirmedTopic,
		summit.TipAwardedEvent:       TipAwardedTopic}
}

// EventTypeFromTopic gets the event type from the topic
// returns nil if the topic is not found.
func EventTypeFromTopic(ogTopic common.Hash) *summit.EventType {
	for eventType, topic := range TopicMap() {
		if bytes.Equal(ogTopic.Bytes(), topic.Bytes()) {
			return &eventType
		}
	}
	return nil
}

// Topic gets the topic from the event type.
func Topic(eventType summit.EventType) common.Hash {
	topicHash, ok := TopicMap()[summit.EventType(eventType.Int())]
	if !ok {
		panic("unknown event")
	}
	return topicHash
}
