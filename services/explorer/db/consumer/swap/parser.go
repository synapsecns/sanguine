package swap

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
	"github.com/synapsecns/sanguine/services/explorer/db"
	swapTypes "github.com/synapsecns/sanguine/services/explorer/types/swap"
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
	// filterer is the swap filterer we use to parse events
	filterer *swap.SwapFlashLoanFilterer
}

// NewParser creates a new parser for a given bridge.
func NewParser(consumerDB db.ConsumerDB, swapAddress common.Address) (*Parser, error) {
	filterer, err := swap.NewSwapFlashLoanFilterer(swapAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", bridge.SynapseBridgeFilterer{}, err)
	}
	return &Parser{consumerDB, filterer}, nil
}

func (p *Parser) EventType(log ethTypes.Log) (_ swapTypes.EventType, ok bool) {
	for _, logTopic := range log.Topics {
		eventType := swap.EventTypeFromTopic(logTopic)
		if eventType == nil {
			continue
		}

		return *eventType, true
	}
	// return an unknown event to avoid cases where user failed to check the event type
	return swapTypes.EventType(len(swapTypes.AllEventTypes()) + 2), false
}

func (p *Parser) GetSwapEvent(log ethTypes.Log) error {
	for _, logTopic := range log.Topics {
		switch logTopic {
		case swap.Topic(swapTypes.TokenSwapEvent):
			iface, err := p.filterer.ParseTokenSwap(log)
			if err != nil {
				return fmt.Errorf("could not parse: %w", err)
			}
			err = p.consumerDB.StoreSwapEvent(context.TODO(), iface, 1)
			if err != nil {
				return fmt.Errorf("could not store deposit: %w", err)
			}
			return nil
			//...
		}
	}
	return fmt.Errorf("did not find event type for log")
}
