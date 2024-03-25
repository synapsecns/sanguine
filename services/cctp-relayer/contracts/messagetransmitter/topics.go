package messagetransmitter

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	var err error

	parsedMessageTransmitter, err := abi.JSON(strings.NewReader(MessageTransmitterMetaData.ABI))
	if err != nil {
		panic(err)
	}

	MessageSentTopic = parsedMessageTransmitter.Events["MessageSent"].ID

	if MessageSentTopic == (common.Hash{}) {
		panic("topic is nil")
	}

	MessageReceivedTopic = parsedMessageTransmitter.Events["MessageReceived"].ID

	if MessageReceivedTopic == (common.Hash{}) {
		panic("topic is nil")
	}
}

// MessageSentTopic is the topic that gets emitted when the sent event is called.
var MessageSentTopic common.Hash

// MessageReceivedTopic is the topic that gets emitted when the received event is called.
var MessageReceivedTopic common.Hash
