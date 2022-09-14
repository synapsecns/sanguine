package bridge

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/db"
	bridgeTypes "github.com/synapsecns/sanguine/services/explorer/types/bridge"
)

//// Parser parses bridge events
//// TODO: consider moving this into a separate package.
//type Parser interface {
//	// EventType determines if an event was initiated by the bridge or the user.
//	// note: this currently will not handle multiple topics that corespond to events (think a multicall)
//	// support for this will need to be added for multi-call style rollups, etc
//	EventType(log ethTypes.Log) (_ types.EventType, ok bool)
//	// GetCrossChainUserEvent parses a cross chain user event from a log.
//	GetCrossChainUserEvent(log ethTypes.Log) (_ types.CrossChainUserEventLog, err error)
//	// GetCrossChainBridgeEvent gets a bridge event log.
//	GetCrossChainBridgeEvent(logs ethTypes.Log) (_ types.CrossChainBridgeEventLog, err error)
//}

type Parser struct {
	// consumerDB is the database to store parsed data in
	consumerDB db.ConsumerDB
	// filterer is the bridge filterer we use to parse events
	filterer *bridge.SynapseBridgeFilterer
}

// NewParser creates a new parser for a given bridge.
func NewParser(consumerDB db.ConsumerDB, bridgeAddress common.Address) (*Parser, error) {
	filterer, err := bridge.NewSynapseBridgeFilterer(bridgeAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", bridge.SynapseBridgeFilterer{}, err)
	}
	return &Parser{consumerDB, filterer}, nil
}

func (p *Parser) EventType(log ethTypes.Log) (_ bridgeTypes.EventType, ok bool) {
	for _, logTopic := range log.Topics {
		eventType := bridge.EventTypeFromTopic(logTopic)
		if eventType == nil {
			continue
		}

		return *eventType, true
	}
	// return an unknown event to avoid cases where user failed to check the event type
	return bridgeTypes.EventType(len(bridgeTypes.AllEventTypes()) + 2), false
}

func (p *Parser) GetCrossChainUserEvent(log ethTypes.Log) error {
	for _, logTopic := range log.Topics {
		switch logTopic {
		case bridge.Topic(bridgeTypes.RedeemAndSwapEvent):
			p.filterer.ParseTokenRedeemAndSwap(log)
		case bridge.Topic(bridgeTypes.DepositAndSwapEvent):
			p.filterer.ParseTokenDepositAndSwap(log)
		case bridge.Topic(bridgeTypes.RedeemAndRemoveEvent):
			p.filterer.ParseTokenRedeemAndRemove(log)
		case bridge.Topic(bridgeTypes.RedeemEvent):
			p.filterer.ParseTokenRedeem(log)
		case bridge.Topic(bridgeTypes.DepositEvent):
			iface, err := p.filterer.ParseTokenDeposit(log)
			if err != nil {
				return fmt.Errorf("could not parse token deposit: %w", err)
			}
			err = p.consumerDB.StoreDeposit(context.TODO(), iface, 1)
			if err != nil {
				return fmt.Errorf("could not store deposit: %w", err)
			}
			return nil
		case bridge.Topic(bridgeTypes.RedeemV2Event):
			p.filterer.ParseTokenRedeemV2(log)
		}
	}
	return fmt.Errorf("did not find event type for log")
}

func (p *Parser) GetCrossChainBridgeEvent(log ethTypes.Log) error {
	for _, logTopic := range log.Topics {
		switch logTopic {
		case bridge.Topic(bridgeTypes.WithdrawEvent):
			p.filterer.ParseTokenWithdraw(log)
		case bridge.Topic(bridgeTypes.MintEvent):
			p.filterer.ParseTokenMint(log)
		case bridge.Topic(bridgeTypes.MintAndSwapEvent):
			p.filterer.ParseTokenMintAndSwap(log)
		case bridge.Topic(bridgeTypes.WithdrawAndRemoveEvent):
			p.filterer.ParseTokenWithdrawAndRemove(log)
		}
	}
	return fmt.Errorf("could not get cross chain event log: %w", err)
}
