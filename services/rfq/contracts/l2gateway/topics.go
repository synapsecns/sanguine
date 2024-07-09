package l2gateway

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	var err error

	parsedL2GatewayRouter, err := abi.JSON(strings.NewReader(L2GatewayRouterMetaData.ABI))
	if err != nil {
		panic(err)
	}

	WithdrawETHTopic = parsedL2GatewayRouter.Events["WithdrawETH"].ID
	WithdrawERC20Topic = parsedL2GatewayRouter.Events["WithdrawERC20"].ID
	FinalizeDepositETHTopic = parsedL2GatewayRouter.Events["FinalizeDepositETH"].ID
	FinalizeDepositERC20Topic = parsedL2GatewayRouter.Events["FinalizeDepositERC20"].ID

	if WithdrawETHTopic == (common.Hash{}) {
		panic("topic is nil")
	}
	if WithdrawERC20Topic == (common.Hash{}) {
		panic("topic is nil")
	}
	if FinalizeDepositETHTopic == (common.Hash{}) {
		panic("topic is nil")
	}
	if FinalizeDepositERC20Topic == (common.Hash{}) {
		panic("topic is nil")
	}
}

// WithdrawETHTopic is the topic that gets emitted when the depositForBurn event is called.
var WithdrawETHTopic common.Hash

// WithdrawERC20Topic is the topic that gets emitted when the depositForBurn event is called.
var WithdrawERC20Topic common.Hash

// FinalizeDepositETHTopic is the topic that gets emitted when the depositForBurn event is called.
var FinalizeDepositETHTopic common.Hash

// FinalizeDepositERC20Topic is the topic that gets emitted when the depositForBurn event is called.
var FinalizeDepositERC20Topic common.Hash
