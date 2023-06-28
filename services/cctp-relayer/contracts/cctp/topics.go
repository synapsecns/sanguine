package cctp

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	// set topics
	var err error

	parsedCCTP, err := abi.JSON(strings.NewReader(SynapseCCTPMetaData.ABI))
	if err != nil {
		panic(err)
	}

	CircleRequestSentTopic = parsedCCTP.Events["CircleRequestSent"].ID

	CircleRequestFulfilledTopic = parsedCCTP.Events["CircleRequestFulfilled"].ID

	for _, topic := range []common.Hash{CircleRequestSentTopic, CircleRequestFulfilledTopic} {
		if topic == (common.Hash{}) {
			panic("topic is nil")
		}
	}
}

var (
	// CircleRequestSentTopic is the topic that gets emitted when the sent event is called.
	CircleRequestSentTopic common.Hash
	// CircleRequestFulfilledTopic is the topic that gets emitted when the sent event is called.
	CircleRequestFulfilledTopic common.Hash
)
