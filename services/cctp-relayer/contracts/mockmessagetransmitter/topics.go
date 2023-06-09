package mockmessagetransmitter

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	var err error

	parsedMessageTransmitter, err := abi.JSON(strings.NewReader(MockMessageTransmitterMetaData.ABI))
	if err != nil {
		panic(err)
	}

	MessageSentTopic = parsedMessageTransmitter.Events["MessageSent"].ID

	if MessageSentTopic == (common.Hash{}) {
		panic("topic is nil")
	}
}

// MessageSentTopic is the topic that gets emitted when the sent event is called.
var MessageSentTopic common.Hash
