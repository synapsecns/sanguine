package bridge

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
	"strings"
)

func init() {
	var err error
	parsedBridge, err := abi.JSON(strings.NewReader(SynapseBridgeMetaData.ABI))
	if err != nil {
		panic(err)
	}

	// we do this here to throw a compile error if the event is not found
	DepositTopic = parsedBridge.Events["TokenDeposit"].ID

	RedeemTopic = parsedBridge.Events["TokenRedeem"].ID

	WithdrawTopic = parsedBridge.Events["TokenWithdraw"].ID

	MintTopic = parsedBridge.Events["TokenMint"].ID

	DepositAndSwapTopic = parsedBridge.Events["TokenDepositAndSwap"].ID

	RedeemAndSwapTopic = parsedBridge.Events["TokenRedeemAndSwap"].ID

	RedeemAndRemoveTopic = parsedBridge.Events["TokenRedeemAndRemove"].ID

	MintAndSwapTopic = parsedBridge.Events["TokenMintAndSwap"].ID

	WithdrawAndRemoveTopic = parsedBridge.Events["TokenWithdrawAndRemove"].ID

	RedeemV2Topic = parsedBridge.Events["TokenRedeemV2"].ID
}

// DepositTopic is the topic used for token deposits.
var DepositTopic common.Hash

// RedeemTopic is the topic used for token redeems.
var RedeemTopic common.Hash

// WithdrawTopic is the topic used for token withdraws (called by bridge).
var WithdrawTopic common.Hash

// MintTopic is the topic used for token mints (called by bridge).
var MintTopic common.Hash

// DepositAndSwapTopic is the topic used for token deposits->swaps.
var DepositAndSwapTopic common.Hash

// RedeemAndSwapTopic is the topic used for redeems->swaps.
var RedeemAndSwapTopic common.Hash

// RedeemAndRemoveTopic is the topic used for redeems->swaps/burn.
var RedeemAndRemoveTopic common.Hash

// MintAndSwapTopic is the topic used for mint and swaps (called by bridge).
var MintAndSwapTopic common.Hash

// WithdrawAndRemoveTopic is the topic used for withdraw and removes (called by bridge).
var WithdrawAndRemoveTopic common.Hash

// RedeemV2Topic is the topic used for redeems to a non-evm chain.
var RedeemV2Topic common.Hash

// TopicMap maps events to topics.
// this is returned as a function to assert immutability.
func TopicMap() map[bridge.EventType]common.Hash {
	return map[bridge.EventType]common.Hash{
		bridge.DepositEvent:           DepositTopic,
		bridge.RedeemEvent:            RedeemTopic,
		bridge.WithdrawEvent:          WithdrawTopic,
		bridge.MintEvent:              MintTopic,
		bridge.DepositAndSwapEvent:    DepositAndSwapTopic,
		bridge.MintAndSwapEvent:       MintAndSwapTopic,
		bridge.RedeemAndSwapEvent:     RedeemAndSwapTopic,
		bridge.RedeemAndRemoveEvent:   RedeemAndRemoveTopic,
		bridge.WithdrawAndRemoveEvent: WithdrawAndRemoveTopic,
		bridge.RedeemV2Event:          RedeemV2Topic,
	}
}

// EventTypeFromTopic gets the event type from the topic
// returns nil if the topic is not found.
func EventTypeFromTopic(ogTopic common.Hash) *bridge.EventType {
	for eventType, topic := range TopicMap() {
		if bytes.Equal(ogTopic.Bytes(), topic.Bytes()) {
			return &eventType
		}
	}
	return nil
}

// Topic gets the topic from the event type.
func Topic(eventType bridge.EventType) common.Hash {
	topicHash, ok := TopicMap()[bridge.EventType(eventType.Int())]
	if !ok {
		panic("unknown event")
	}
	return topicHash
}
