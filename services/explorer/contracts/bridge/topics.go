package bridge

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/types"
	"strings"
)

func init() {
	var err error
	parsedBridge, err := abi.JSON(strings.NewReader(SynapseBridgeMetaData.ABI))
	if err != nil {
		panic(err)
	}

	// we do this here to throw a compile error if the event is not found
	TokenDepositTopic = parsedBridge.Events["TokenDeposit"].ID

	TokenRedeemTopic = parsedBridge.Events["TokenRedeem"].ID

	TokenWithdrawTopic = parsedBridge.Events["TokenWithdraw"].ID

	TokenMintTopic = parsedBridge.Events["TokenMint"].ID

	TokenDepositAndSwap = parsedBridge.Events["TokenDepositAndSwap"].ID

	TokenRedeemAndSwapTopic = parsedBridge.Events["TokenRedeemAndSwap"].ID

	TokenRedeemAndRemoveTopic = parsedBridge.Events["TokenRedeemAndRemove"].ID

	TokenMintAndSwapTopic = parsedBridge.Events["TokenMintAndSwap"].ID

	TokenWithdrawAndRemoveTopic = parsedBridge.Events["TokenWithdrawAndRemove"].ID

	TokenRedeemV2Topic = parsedBridge.Events["TokenRedeemV2"].ID
}

// TokenDepositTopic is the topic used for token deposits.
var TokenDepositTopic common.Hash

// TokenRedeemTopic is the topic used for token redeems.
var TokenRedeemTopic common.Hash

// TokenWithdrawTopic is the topic used for token withdraws (called by bridge).
var TokenWithdrawTopic common.Hash

// TokenMintTopic is the topic used for token mints (called by bridge).
var TokenMintTopic common.Hash

// TokenDepositAndSwap is the topic used for token deposits->swaps.
var TokenDepositAndSwap common.Hash

// TokenRedeemAndSwapTopic is the topic used for redeems->swaps.
var TokenRedeemAndSwapTopic common.Hash

// TokenRedeemAndRemoveTopic is the topic used for redeems->swaps/burn.
var TokenRedeemAndRemoveTopic common.Hash

// TokenMintAndSwapTopic is the topic used for mint and swaps (called by bridge).
var TokenMintAndSwapTopic common.Hash

// TokenWithdrawAndRemoveTopic is the topic used for withdraw and removes (called by bridge).
var TokenWithdrawAndRemoveTopic common.Hash

// TokenRedeemV2Topic is the topic used for redeems to a non-evm chain.
var TokenRedeemV2Topic common.Hash

// topicMap maps events to topics.
// this is returned as a function to assert immutability.
func topicMap() map[types.EventType]common.Hash {
	return map[types.EventType]common.Hash{
		types.DepositEvent:         TokenDepositTopic,
		types.RedeemEvent:          TokenRedeemTopic,
		types.WithdrawEvent:        TokenWithdrawTopic,
		types.MintEvent:            TokenMintTopic,
		types.DepositAndSwapEvent:  TokenDepositAndSwap,
		types.RedeemAndSwapEvent:   TokenRedeemAndSwapTopic,
		types.RedeemAndRemoveEvent: TokenRedeemAndRemoveTopic,
		types.MintAndSwap:          TokenMintAndSwapTopic,
		types.WithdrawAndRemove:    TokenWithdrawAndRemoveTopic,
		types.RedeemV2Event:        TokenRedeemV2Topic,
	}
}

// eventTypeFromTopic gets the event type from the topic
// returns nil if the topic is not found.
func eventTypeFromTopic(ogTopic common.Hash) *types.EventType {
	for eventType, topic := range topicMap() {
		if bytes.Equal(ogTopic.Bytes(), topic.Bytes()) {
			return &eventType
		}
	}
	return nil
}

// Topic gets the topic from the event type.
func Topic(eventType types.EventType) common.Hash {
	topicHash, ok := topicMap()[types.EventType(eventType.Int())]
	if !ok {
		panic("unknown event")
	}
	return topicHash
}
