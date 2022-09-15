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

func (p *Parser) ParseAndStore(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	for _, logTopic := range log.Topics {
		switch logTopic {
		case bridge.Topic(bridgeTypes.DepositEvent):
			err := p.parseAndStoreDeposit(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not parse and store deposit: %w", err)
			}
		case bridge.Topic(bridgeTypes.RedeemEvent):
			err := p.parseAndStoreRedeem(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not parse and store redeem: %w", err)
			}
		case bridge.Topic(bridgeTypes.WithdrawEvent):
			err := p.parseAndStoreWithdraw(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not parse and store withdraw: %w", err)
			}
		case bridge.Topic(bridgeTypes.MintEvent):
			err := p.parseAndStoreMint(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not parse and store mint: %w", err)
			}
		case bridge.Topic(bridgeTypes.DepositAndSwapEvent):
			err := p.parseAndStoreDepositAndSwap(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not parse and store deposit and swap: %w", err)
			}
		case bridge.Topic(bridgeTypes.MintAndSwapEvent):
			err := p.parseAndStoreMintAndSwap(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not parse and store mint and swap: %w", err)
			}
		case bridge.Topic(bridgeTypes.RedeemAndSwapEvent):
			err := p.parseAndStoreRedeemAndSwap(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not parse and store redeem and swap: %w", err)
			}
		case bridge.Topic(bridgeTypes.RedeemAndRemoveEvent):
			err := p.parseAndStoreRedeemAndRemove(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not parse and store redeem and remove: %w", err)
			}
		case bridge.Topic(bridgeTypes.WithdrawAndRemoveEvent):
			err := p.parseAndStoreWithdrawAndRemove(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not parse and store withdraw and remove: %w", err)
			}
		case bridge.Topic(bridgeTypes.RedeemV2Event):
			err := p.parseAndStoreRedeemV2(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not parse and store redeem v2: %w", err)
			}
		}
	}

	return nil
}

func (p *Parser) parseAndStoreDeposit(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenDeposit(log)
	if err != nil {
		return fmt.Errorf("could not parse token deposit: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID)
	if err != nil {
		return fmt.Errorf("could not store deposit: %w", err)
	}
	return nil
}

func (p *Parser) parseAndStoreRedeem(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenRedeem(log)
	if err != nil {
		return fmt.Errorf("could not parse token redeem: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID)
	if err != nil {
		return fmt.Errorf("could not store redeem: %w", err)
	}
	return nil
}

func (p *Parser) parseAndStoreWithdraw(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenWithdraw(log)
	if err != nil {
		return fmt.Errorf("could not parse token withdraw: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID)
	if err != nil {
		return fmt.Errorf("could not store withdraw: %w", err)
	}
	return nil
}

func (p *Parser) parseAndStoreMint(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenMint(log)
	if err != nil {
		return fmt.Errorf("could not parse token mint: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID)
	if err != nil {
		return fmt.Errorf("could not store mint: %w", err)
	}
	return nil
}

func (p *Parser) parseAndStoreDepositAndSwap(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenDepositAndSwap(log)
	if err != nil {
		return fmt.Errorf("could not parse token deposit and swap: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID)
	if err != nil {
		return fmt.Errorf("could not store deposit and swap: %w", err)
	}
	return nil
}

func (p *Parser) parseAndStoreMintAndSwap(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenMintAndSwap(log)
	if err != nil {
		return fmt.Errorf("could not parse token mint and swap: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID)
	if err != nil {
		return fmt.Errorf("could not store mint and swap: %w", err)
	}
	return nil
}

func (p *Parser) parseAndStoreRedeemAndSwap(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenRedeemAndSwap(log)
	if err != nil {
		return fmt.Errorf("could not parse token redeem and swap: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID)
	if err != nil {
		return fmt.Errorf("could not store redeem and swap: %w", err)
	}
	return nil
}

func (p *Parser) parseAndStoreRedeemAndRemove(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenRedeemAndRemove(log)
	if err != nil {
		return fmt.Errorf("could not parse token redeem and remove: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID)
	if err != nil {
		return fmt.Errorf("could not store redeem and remove: %w", err)
	}
	return nil
}

func (p *Parser) parseAndStoreWithdrawAndRemove(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenWithdrawAndRemove(log)
	if err != nil {
		return fmt.Errorf("could not parse token withdraw and remove: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID)
	if err != nil {
		return fmt.Errorf("could not store withdraw and remove: %w", err)
	}
	return nil
}

func (p *Parser) parseAndStoreRedeemV2(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenRedeemV2(log)
	if err != nil {
		return fmt.Errorf("could not parse token redeem v2: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID)
	if err != nil {
		return fmt.Errorf("could not store redeem v2: %w", err)
	}
	return nil
}
