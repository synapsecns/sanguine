package consumer

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
	"github.com/synapsecns/sanguine/services/explorer/db"
	bridgeTypes "github.com/synapsecns/sanguine/services/explorer/types/bridge"
	swapTypes "github.com/synapsecns/sanguine/services/explorer/types/swap"
)

// Parser parses events and stores them.
type Parser interface {
	// ParseAndStore parses the logs and stores them in the database.
	ParseAndStore(ctx context.Context, log ethTypes.Log, chainID uint32) error
}

// BridgeParser parses events from the bridge contract.
type BridgeParser struct {
	// consumerDB is the database to store parsed data in
	consumerDB db.ConsumerDB
	// filterer is the bridge filterer we use to parse events
	filterer *bridge.SynapseBridgeFilterer
	// bridgeAddress is the address of the bridge
	bridgeAddress common.Address
	// fetcher is a Bridge Config Fetcher
	fetcher BridgeConfigFetcher
}

// NewBridgeParser creates a new parser for a given bridge.
func NewBridgeParser(consumerDB db.ConsumerDB, bridgeAddress common.Address, bridgeConfigFetcher BridgeConfigFetcher) (*BridgeParser, error) {
	filterer, err := bridge.NewSynapseBridgeFilterer(bridgeAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", bridge.SynapseBridgeFilterer{}, err)
	}
	return &BridgeParser{consumerDB, filterer, bridgeAddress, bridgeConfigFetcher}, nil
}

// EventType returns the event type of a bridge log.
func (p *BridgeParser) EventType(log ethTypes.Log) (_ bridgeTypes.EventType, ok bool) {
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

// SwapParser parses events from the swap contract.
type SwapParser struct {
	// consumerDB is the database to store parsed data in
	consumerDB db.ConsumerDB
	// filterer is the swap filterer we use to parse events
	filterer *swap.SwapFlashLoanFilterer
}

// NewSwapParser creates a new parser for a given bridge.
func NewSwapParser(consumerDB db.ConsumerDB, swapAddress common.Address) (*SwapParser, error) {
	filterer, err := swap.NewSwapFlashLoanFilterer(swapAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", bridge.SynapseBridgeFilterer{}, err)
	}
	return &SwapParser{consumerDB, filterer}, nil
}

// EventType returns the event type of a swap log.
func (p *SwapParser) EventType(log ethTypes.Log) (_ swapTypes.EventType, ok bool) {
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

// ParseAndStore parses the bridge logs and stores them in the database.
//
// nolint:gocognit,cyclop
func (p *BridgeParser) ParseAndStore(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	//nolint:dupl
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

func (p *BridgeParser) parseAndStoreDeposit(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenDeposit(log)
	if err != nil {
		return fmt.Errorf("could not parse token deposit: %w", err)
	}

	// get BridgeConfig data
	tokenID, err := p.fetcher.GetTokenID(ctx, chainID, uint32(iface.GetBlockNumber()), iface.GetToken())
	if err != nil {
		return fmt.Errorf("could not parse get token from bridge config event: %w", err)
	}

	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID, tokenID)
	if err != nil {
		return fmt.Errorf("could not store deposit: %w", err)
	}
	return nil
}

func (p *BridgeParser) parseAndStoreRedeem(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenRedeem(log)
	if err != nil {
		return fmt.Errorf("could not parse token redeem: %w", err)
	}

	// get BridgeConfig data
	tokenID, err := p.fetcher.GetTokenID(ctx, chainID, uint32(iface.GetBlockNumber()), iface.GetToken())
	if err != nil {
		return fmt.Errorf("could not parse get token from bridge config event: %w", err)
	}

	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID, tokenID)
	if err != nil {
		return fmt.Errorf("could not store redeem: %w", err)
	}
	return nil
}

func (p *BridgeParser) parseAndStoreWithdraw(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenWithdraw(log)
	if err != nil {
		return fmt.Errorf("could not parse token withdraw: %w", err)
	}

	// get BridgeConfig data
	tokenID, err := p.fetcher.GetTokenID(ctx, chainID, uint32(iface.GetBlockNumber()), iface.GetToken())
	if err != nil {
		return fmt.Errorf("could not parse get token from bridge config event: %w", err)
	}

	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID, tokenID)
	if err != nil {
		return fmt.Errorf("could not store withdraw: %w", err)
	}
	return nil
}

func (p *BridgeParser) parseAndStoreMint(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenMint(log)
	if err != nil {
		return fmt.Errorf("could not parse token mint: %w", err)
	}

	// get BridgeConfig data
	tokenID, err := p.fetcher.GetTokenID(ctx, chainID, uint32(iface.GetBlockNumber()), iface.GetToken())
	if err != nil {
		return fmt.Errorf("could not parse get token from bridge config event: %w", err)
	}

	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID, tokenID)
	if err != nil {
		return fmt.Errorf("could not store mint: %w", err)
	}
	return nil
}

func (p *BridgeParser) parseAndStoreDepositAndSwap(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenDepositAndSwap(log)
	if err != nil {
		return fmt.Errorf("could not parse token deposit and swap: %w", err)
	}
	// get BridgeConfig data
	tokenID, err := p.fetcher.GetTokenID(ctx, chainID, uint32(iface.GetBlockNumber()), iface.GetToken())
	if err != nil {
		return fmt.Errorf("could not parse get token from bridge config event: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID, tokenID)
	if err != nil {
		return fmt.Errorf("could not store deposit and swap: %w", err)
	}
	return nil
}

func (p *BridgeParser) parseAndStoreMintAndSwap(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenMintAndSwap(log)
	if err != nil {
		return fmt.Errorf("could not parse token mint and swap: %w", err)
	}
	// get BridgeConfig data
	tokenID, err := p.fetcher.GetTokenID(ctx, chainID, uint32(iface.GetBlockNumber()), iface.GetToken())
	if err != nil {
		return fmt.Errorf("could not parse get token from bridge config event: %w", err)
	}

	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID, tokenID)
	if err != nil {
		return fmt.Errorf("could not store mint and swap: %w", err)
	}
	return nil
}

func (p *BridgeParser) parseAndStoreRedeemAndSwap(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenRedeemAndSwap(log)
	if err != nil {
		return fmt.Errorf("could not parse token redeem and swap: %w", err)
	}
	// get BridgeConfig data
	tokenID, err := p.fetcher.GetTokenID(ctx, chainID, uint32(iface.GetBlockNumber()), iface.GetToken())
	if err != nil {
		return fmt.Errorf("could not parse get token from bridge config event: %w", err)
	}

	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID, tokenID)
	if err != nil {
		return fmt.Errorf("could not store redeem and swap: %w", err)
	}
	return nil
}

func (p *BridgeParser) parseAndStoreRedeemAndRemove(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenRedeemAndRemove(log)
	if err != nil {
		return fmt.Errorf("could not parse token redeem and remove: %w", err)
	}
	// get BridgeConfig data
	tokenID, err := p.fetcher.GetTokenID(ctx, chainID, uint32(iface.GetBlockNumber()), iface.GetToken())
	if err != nil {
		return fmt.Errorf("could not parse get token from bridge config event: %w", err)
	}

	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID, tokenID)
	if err != nil {
		return fmt.Errorf("could not store redeem and remove: %w", err)
	}
	return nil
}

func (p *BridgeParser) parseAndStoreWithdrawAndRemove(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenWithdrawAndRemove(log)
	if err != nil {
		return fmt.Errorf("could not parse token withdraw and remove: %w", err)
	}
	// get BridgeConfig data
	tokenID, err := p.fetcher.GetTokenID(ctx, chainID, uint32(iface.GetBlockNumber()), iface.GetToken())
	if err != nil {
		return fmt.Errorf("could not parse get token from bridge config event: %w", err)
	}

	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID, tokenID)
	if err != nil {
		return fmt.Errorf("could not store withdraw and remove: %w", err)
	}
	return nil
}

func (p *BridgeParser) parseAndStoreRedeemV2(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenRedeemV2(log)
	if err != nil {
		return fmt.Errorf("could not parse token redeem v2: %w", err)
	}
	// get BridgeConfig data
	tokenID, err := p.fetcher.GetTokenID(ctx, chainID, uint32(iface.GetBlockNumber()), iface.GetToken())
	if err != nil {
		return fmt.Errorf("could not parse get token from bridge config event: %w", err)
	}

	err = p.consumerDB.StoreEvent(ctx, iface, nil, chainID, tokenID)
	if err != nil {
		return fmt.Errorf("could not store redeem v2: %w", err)
	}
	return nil
}

// ParseAndStore parses and stores the swap logs.
//
//nolint:gocognit,cyclop
func (p *SwapParser) ParseAndStore(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	//nolint:dupl
	for _, logTopic := range log.Topics {
		switch logTopic {
		case swap.Topic(swapTypes.TokenSwapEvent):
			err := p.parseTokenSwap(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store token swap: %w", err)
			}
		case swap.Topic(swapTypes.AddLiquidityEvent):
			err := p.parseAddLiquidity(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store add liquidity: %w", err)
			}
		case swap.Topic(swapTypes.RemoveLiquidityEvent):
			err := p.parseRemoveLiquidity(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store remove liquidity: %w", err)
			}
		case swap.Topic(swapTypes.RemoveLiquidityOneEvent):
			err := p.parseRemoveLiquidityOne(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store remove liquidity one: %w", err)
			}
		case swap.Topic(swapTypes.RemoveLiquidityImbalanceEvent):
			err := p.parseRemoveLiquidityImbalance(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store remove liquidity imbalance: %w", err)
			}
		case swap.Topic(swapTypes.NewAdminFeeEvent):
			err := p.parseNewAdminFee(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store new admin fee: %w", err)
			}
		case swap.Topic(swapTypes.NewSwapFeeEvent):
			err := p.parseNewSwapFee(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store new swap fee: %w", err)
			}
		case swap.Topic(swapTypes.RampAEvent):
			err := p.parseRampA(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store ramp a: %w", err)
			}
		case swap.Topic(swapTypes.StopRampAEvent):
			err := p.parseStopRampA(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store stop ramp a: %w", err)
			}
		case swap.Topic(swapTypes.FlashLoanEvent):
			err := p.parseFlashLoan(ctx, log, chainID)
			if err != nil {
				return fmt.Errorf("could not store flash loan: %w", err)
			}
		}
	}
	return fmt.Errorf("did not find event type for log")
}

func (p *SwapParser) parseTokenSwap(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseTokenSwap(log)
	if err != nil {
		return fmt.Errorf("could not parse token swap: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID, nil)
	if err != nil {
		return fmt.Errorf("could not store token swap: %w", err)
	}
	return nil
}

func (p *SwapParser) parseAddLiquidity(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseAddLiquidity(log)
	if err != nil {
		return fmt.Errorf("could not parse add liquidity: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID, nil)
	if err != nil {
		return fmt.Errorf("could not store add liquidity: %w", err)
	}
	return nil
}

func (p *SwapParser) parseRemoveLiquidity(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseRemoveLiquidity(log)
	if err != nil {
		return fmt.Errorf("could not parse remove liquidity: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID, nil)
	if err != nil {
		return fmt.Errorf("could not store remove liquidity: %w", err)
	}
	return nil
}

func (p *SwapParser) parseRemoveLiquidityOne(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseRemoveLiquidityOne(log)
	if err != nil {
		return fmt.Errorf("could not parse remove liquidity one: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID, nil)
	if err != nil {
		return fmt.Errorf("could not store remove liquidity one: %w", err)
	}
	return nil
}
func (p *SwapParser) parseRemoveLiquidityImbalance(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseRemoveLiquidityImbalance(log)
	if err != nil {
		return fmt.Errorf("could not parse remove liquidity imbalance: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID, nil)
	if err != nil {
		return fmt.Errorf("could not store remove liquidity imbalance: %w", err)
	}
	return nil
}

func (p *SwapParser) parseNewAdminFee(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseNewAdminFee(log)
	if err != nil {
		return fmt.Errorf("could not parse new admin fee: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID, nil)
	if err != nil {
		return fmt.Errorf("could not store new admin fee: %w", err)
	}
	return nil
}

func (p *SwapParser) parseNewSwapFee(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseNewSwapFee(log)
	if err != nil {
		return fmt.Errorf("could not parse new swap fee: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID, nil)
	if err != nil {
		return fmt.Errorf("could not store new swap fee: %w", err)
	}
	return nil
}

func (p *SwapParser) parseRampA(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseRampA(log)
	if err != nil {
		return fmt.Errorf("could not parse Ramp A: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID, nil)
	if err != nil {
		return fmt.Errorf("could not store Ramp A: %w", err)
	}
	return nil
}
func (p *SwapParser) parseStopRampA(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseStopRampA(log)
	if err != nil {
		return fmt.Errorf("could not parse stop Ramp A: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID, nil)
	if err != nil {
		return fmt.Errorf("could not store stop Ramp A: %w", err)
	}
	return nil
}

func (p *SwapParser) parseFlashLoan(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	iface, err := p.filterer.ParseFlashLoan(log)
	if err != nil {
		return fmt.Errorf("could not parse flash loan: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, nil, iface, chainID, nil)
	if err != nil {
		return fmt.Errorf("could not store flash loan: %w", err)
	}
	return nil
}
