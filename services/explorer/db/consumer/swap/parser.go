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

func (p *Parser) ParseAndStore(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	for _, logTopic := range log.Topics {
		switch logTopic {
		case swap.Topic(swapTypes.TokenSwapEvent):
			err := p.parseTokenSwap(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store deposit: %w", err)
			}
		case swap.Topic(swapTypes.AddLiquidityEvent):
			err := p.parseAddLiquidity(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store deposit: %w", err)
			}
		case swap.Topic(swapTypes.RemoveLiquidityEvent):
			err := p.parseRemoveLiquidity(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store deposit: %w", err)
			}
		case swap.Topic(swapTypes.RemoveLiquidityOneEvent):
			err := p.parseRemoveLiquidityOne(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store deposit: %w", err)
			}
		case swap.Topic(swapTypes.RemoveLiquidityImbalanceEvent):
			err := p.parseRemoveLiquidityImbalance(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store deposit: %w", err)
			}
		case swap.Topic(swapTypes.NewAdminFeeEvent):
			err := p.parseNewAdminFee(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store deposit: %w", err)
			}
		case swap.Topic(swapTypes.NewSwapFeeEvent):
			err := p.parseNewSwapFee(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store deposit: %w", err)
			}
		case swap.Topic(swapTypes.RampAEvent):
			err := p.parseRampA(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store deposit: %w", err)
			}
		case swap.Topic(swapTypes.StopRampAEvent):
			err := p.parseTokenSwap(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store deposit: %w", err)
			}
		case swap.Topic(swapTypes.FlashLoanEvent):
			err := p.parseTokenSwap(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store deposit: %w", err)
			}

		}

	}
	return fmt.Errorf("did not find event type for log")
}

func (p *Parser) parseTokenSwap(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenSwap(log)
	if err != nil {
		return fmt.Errorf("could not parse token swap: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID)
	if err != nil {
		return fmt.Errorf("could not store token swap: %w", err)
	}
	return nil
}

func (p *Parser) parseAddLiquidity(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseAddLiquidity(log)
	if err != nil {
		return fmt.Errorf("could not parse add liquidity: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID)
	if err != nil {
		return fmt.Errorf("could not store add liquidity: %w", err)
	}
	return nil
}

func (p *Parser) parseRemoveLiquidity(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseRemoveLiquidity(log)
	if err != nil {
		return fmt.Errorf("could not parse remove liquidity: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID)
	if err != nil {
		return fmt.Errorf("could not store remove liquidity: %w", err)
	}
	return nil
}

func (p *Parser) parseRemoveLiquidityOne(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseRemoveLiquidityOne(log)
	if err != nil {
		return fmt.Errorf("could not parse remove liquidity one: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID)
	if err != nil {
		return fmt.Errorf("could not store remove liquidity one: %w", err)
	}
	return nil
}
func (p *Parser) parseRemoveLiquidityImbalance(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseRemoveLiquidityImbalance(log)
	if err != nil {
		return fmt.Errorf("could not parse remove liquidity imbalance: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID)
	if err != nil {
		return fmt.Errorf("could not store remove liquidity imbalance: %w", err)
	}
	return nil
}

func (p *Parser) parseNewAdminFee(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseNewAdminFee(log)
	if err != nil {
		return fmt.Errorf("could not parse new admin fee: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID)
	if err != nil {
		return fmt.Errorf("could not store new admin fee: %w", err)
	}
	return nil
}

func (p *Parser) parseNewSwapFee(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseNewSwapFee(log)
	if err != nil {
		return fmt.Errorf("could not parse new swap fee: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID)
	if err != nil {
		return fmt.Errorf("could not store new swap fee: %w", err)
	}
	return nil
}

func (p *Parser) parseRampA(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseRampA(log)
	if err != nil {
		return fmt.Errorf("could not parse Ramp A: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID)
	if err != nil {
		return fmt.Errorf("could not store Ramp A: %w", err)
	}
	return nil
}
func (p *Parser) parseStopRampA(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseStopRampA(log)
	if err != nil {
		return fmt.Errorf("could not parse stop Ramp A: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID)
	if err != nil {
		return fmt.Errorf("could not store stop Ramp A: %w", err)
	}
	return nil
}

func (p *Parser) parseFlashLoan(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseFlashLoan(log)
	if err != nil {
		return fmt.Errorf("could not parse flash loan: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID)
	if err != nil {
		return fmt.Errorf("could not store flash loan: %w", err)
	}
	return nil
}
