package swap

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/contracts/metaswap"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
	"strings"
)

func init() {
	var err error
	parsedSwap, err := abi.JSON(strings.NewReader(SwapFlashLoanMetaData.ABI))
	if err != nil {
		panic(err)
	}
	parsedMetaSwap, err := abi.JSON(strings.NewReader(metaswap.MetaSwapMetaData.ABI))
	if err != nil {
		panic(err)
	}

	// we do this here to throw a compile error if the event is not found
	TokenSwapTopic = parsedSwap.Events["TokenSwap"].ID

	AddLiquidityTopic = parsedSwap.Events["AddLiquidity"].ID

	RemoveLiquidityTopic = parsedSwap.Events["RemoveLiquidity"].ID

	RemoveLiquidityOneTopic = parsedSwap.Events["RemoveLiquidityOne"].ID

	RemoveLiquidityImbalanceSwap = parsedSwap.Events["RemoveLiquidityImbalance"].ID

	NewAdminFeeTopic = parsedSwap.Events["NewAdminFee"].ID

	NewSwapFeeTopic = parsedSwap.Events["NewSwapFee"].ID

	RampATopic = parsedSwap.Events["RampA"].ID

	StopRampATopic = parsedSwap.Events["StopRampA"].ID

	FlashLoanTopic = parsedSwap.Events["FlashLoan"].ID

	TokenSwapUnderlyingTopic = parsedMetaSwap.Events["TokenSwapUnderlying"].ID
}

// TokenSwapTopic is the topic used for token swap.
var TokenSwapTopic common.Hash

// AddLiquidityTopic is the topic used for adding liquidity.
var AddLiquidityTopic common.Hash

// RemoveLiquidityTopic is the topic used for removing liquidity.
var RemoveLiquidityTopic common.Hash

// RemoveLiquidityOneTopic is the topic used for removing liquidity one.
var RemoveLiquidityOneTopic common.Hash

// RemoveLiquidityImbalanceSwap is the topic used for removing a liquidity imbalance.
var RemoveLiquidityImbalanceSwap common.Hash

// NewAdminFeeTopic is the topic used for a new admin fee.
var NewAdminFeeTopic common.Hash

// NewSwapFeeTopic is the topic used for performing a new swap.
var NewSwapFeeTopic common.Hash

// RampATopic is the topic used for ramp a.
var RampATopic common.Hash

// StopRampATopic is the topic used for stopping ramp a.
var StopRampATopic common.Hash

// FlashLoanTopic is the topic used for Flash Loans.
var FlashLoanTopic common.Hash

// TokenSwapUnderlyingTopic is the topic used for token swap underlying.
var TokenSwapUnderlyingTopic common.Hash

// TopicMap maps events to topics.
// this is returned as a function to assert immutability.
func TopicMap() map[swap.EventType]common.Hash {
	return map[swap.EventType]common.Hash{
		swap.TokenSwapEvent:                TokenSwapTopic,
		swap.AddLiquidityEvent:             AddLiquidityTopic,
		swap.RemoveLiquidityEvent:          RemoveLiquidityTopic,
		swap.RemoveLiquidityOneEvent:       RemoveLiquidityOneTopic,
		swap.RemoveLiquidityImbalanceEvent: RemoveLiquidityImbalanceSwap,
		swap.NewAdminFeeEvent:              NewAdminFeeTopic,
		swap.NewSwapFeeEvent:               NewSwapFeeTopic,
		swap.RampAEvent:                    RampATopic,
		swap.StopRampAEvent:                StopRampATopic,
		swap.FlashLoanEvent:                FlashLoanTopic,
		swap.TokenSwapUnderlyingEvent:      TokenSwapUnderlyingTopic,
	}
}

// EventTypeFromTopic gets the event type from the topic
// returns nil if the topic is not found.
func EventTypeFromTopic(ogTopic common.Hash) *swap.EventType {
	for eventType, topic := range TopicMap() {
		if bytes.Equal(ogTopic.Bytes(), topic.Bytes()) {
			return &eventType
		}
	}
	return nil
}

// Topic gets the topic from the event type.
func Topic(eventType swap.EventType) common.Hash {
	topicHash, ok := TopicMap()[swap.EventType(eventType.Int())]
	if !ok {
		panic("unknown event")
	}
	return topicHash
}
