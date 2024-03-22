package tokenmessenger

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	var err error

	parsedTokenMessenger, err := abi.JSON(strings.NewReader(TokenMessengerMetaData.ABI))
	if err != nil {
		panic(err)
	}

	DepositForBurnTopic = parsedTokenMessenger.Events["DepositForBurn"].ID

	if DepositForBurnTopic == (common.Hash{}) {
		panic("topic is nil")
	}
}

// DepositForBurnTopic is the topic that gets emitted when the depositForBurn event is called.
var DepositForBurnTopic common.Hash
