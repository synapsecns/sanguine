package l1gateway

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	var err error

	parsedL1GatewayRouter, err := abi.JSON(strings.NewReader(L1GatewayRouterMetaData.ABI))
	if err != nil {
		panic(err)
	}

	DepositETHTopic = parsedL1GatewayRouter.Events["DepositETH"].ID
	DepositERC20Topic = parsedL1GatewayRouter.Events["DepositERC20"].ID
	FinalizeWithdrawETHTopic = parsedL1GatewayRouter.Events["FinalizeWithdrawETH"].ID
	FinalizeWithdrawERC20Topic = parsedL1GatewayRouter.Events["FinalizeWithdrawERC20"].ID

	if DepositETHTopic == (common.Hash{}) {
		panic("topic is nil")
	}
	if DepositERC20Topic == (common.Hash{}) {
		panic("topic is nil")
	}
	if FinalizeWithdrawETHTopic == (common.Hash{}) {
		panic("topic is nil")
	}
	if FinalizeWithdrawERC20Topic == (common.Hash{}) {
		panic("topic is nil")
	}
}

// DepositETHTopic is the topic that gets emitted when the depositForBurn event is called.
var DepositETHTopic common.Hash

// DepositERC20Topic is the topic that gets emitted when the depositForBurn event is called.
var DepositERC20Topic common.Hash

// FinalizeWithdrawETHTopic is the topic that gets emitted when the depositForBurn event is called.
var FinalizeWithdrawETHTopic common.Hash

// FinalizeWithdrawERC20Topic is the topic that gets emitted when the depositForBurn event is called.
var FinalizeWithdrawERC20Topic common.Hash
