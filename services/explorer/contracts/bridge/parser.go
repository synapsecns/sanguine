package bridge

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/types"
)

// Parser parses bridge events
// TODO: consider moving this into a separate package.
type Parser interface {
	// EventType determines if an event was initiated by the bridge or the user.
	// note: this currently will not handle multiple topics that corespond to events (think a multicall)
	// support for this will need to be added for multi-call style rollups, etc
	EventType(log ethTypes.Log) (_ types.EventType, ok bool)
	// GetCrossChainUserEvent parses a cross chain user event from a log.
	GetCrossChainUserEvent(log ethTypes.Log) (_ types.CrossChainUserEventLog, err error)
	// GetCrossChainBridgeEvent gets a bridge event log.
	GetCrossChainBridgeEvent(logs ethTypes.Log) (_ types.CrossChainBridgeEventLog, err error)
}

type parserImpl struct {
	// filterer is the bridge filterer we use to parse events
	filterer *SynapseBridgeFilterer
}

// NewParser creates a new parser for a given bridge.
func NewParser(bridgeAddress common.Address) (Parser, error) {
	filterer, err := NewSynapseBridgeFilterer(bridgeAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", SynapseBridgeFilterer{}, err)
	}
	return &parserImpl{filterer: filterer}, nil
}

func (p *parserImpl) EventType(log ethTypes.Log) (_ types.EventType, ok bool) {
	for _, logTopic := range log.Topics {
		eventType := eventTypeFromTopic(logTopic)
		if eventType == nil {
			continue
		}

		return types.EventType(*eventType), true
	}
	// return an unknown event to avoid cases where user failed to check the event type
	return types.EventType(len(types.AllEventTypes()) + 2), false
}

func (p *parserImpl) GetCrossChainUserEvent(log ethTypes.Log) (_ types.CrossChainUserEventLog, err error) {
	for _, logTopic := range log.Topics {
		switch logTopic {
		case Topic(types.RedeemAndSwapEvent):
			return p.filterer.ParseTokenRedeemAndSwap(log)
		case Topic(types.DepositAndSwapEvent):
			return p.filterer.ParseTokenDepositAndSwap(log)
		case Topic(types.RedeemAndRemoveEvent):
			return p.filterer.ParseTokenRedeemAndRemove(log)
		case Topic(types.RedeemEvent):
			return p.filterer.ParseTokenRedeem(log)
		case Topic(types.DepositEvent):
			return p.filterer.ParseTokenDeposit(log)
		case Topic(types.RedeemV2Event):
			return p.filterer.ParseTokenRedeemV2(log)
		}
	}
	return nil, fmt.Errorf("could not get cross chain event log: %w", err)
}

func (p *parserImpl) GetCrossChainBridgeEvent(log ethTypes.Log) (_ types.CrossChainBridgeEventLog, err error) {
	for _, logTopic := range log.Topics {
		switch logTopic {
		case Topic(types.WithdrawEvent):
			return p.filterer.ParseTokenWithdraw(log)
		case Topic(types.MintEvent):
			return p.filterer.ParseTokenMint(log)
		case Topic(types.MintAndSwap):
			return p.filterer.ParseTokenMintAndSwap(log)
		case Topic(types.WithdrawAndRemove):
			return p.filterer.ParseTokenWithdrawAndRemove(log)
		}
	}
	return nil, fmt.Errorf("could not get cross chain event log: %w", err)
}
