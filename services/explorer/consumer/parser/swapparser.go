package parser

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser/tokendata"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser/tokenpool"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
	"github.com/synapsecns/sanguine/services/explorer/db"
	model "github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/static"
	swapTypes "github.com/synapsecns/sanguine/services/explorer/types/swap"
)

// SwapParser parses events from the swap contract.
type SwapParser struct {
	// consumerDB is the database to store parsed data in.
	consumerDB db.ConsumerDB
	// swap is the address of the bridge.
	swapAddress common.Address
	// Filterer is the swap Filterer we use to parse events.
	Filterer *swap.SwapFlashLoanFilterer
	// consumerFetcher is the ScribeFetcher for sender and timestamp.
	consumerFetcher *fetcher.ScribeFetcher
	// tokenDataService contains the token data service/cache
	tokenDataService tokendata.Service
	// poolTokenDataService get the token address from the token index
	poolTokenDataService tokenpool.Service
	// swapService is the swap service
	swapService *fetcher.SwapService
	// coinGeckoIDs is a mapping from coin token symbol to coin gecko ID
	coinGeckoIDs map[string]string
}

// NewSwapParser creates a new parser for a given bridge.
func NewSwapParser(consumerDB db.ConsumerDB, swapAddress common.Address, consumerFetcher *fetcher.ScribeFetcher, swapService *fetcher.SwapService, tokenDataService tokendata.Service) (*SwapParser, error) {
	filterer, err := swap.NewSwapFlashLoanFilterer(swapAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", bridge.SynapseBridgeFilterer{}, err)
	}

	poolTokenDataService, err := tokenpool.NewPoolTokenDataService(*swapService, consumerDB)
	if err != nil {
		return nil, fmt.Errorf("could not create token data service: %w", err)
	}

	coinGeckoIDs, err := ParseYaml(static.GetTokenIDToCoingekoConfig())
	if err != nil {
		return nil, fmt.Errorf("could not open yaml file: %w", err)
	}

	return &SwapParser{
		consumerDB:           consumerDB,
		swapAddress:          swapAddress,
		Filterer:             filterer,
		consumerFetcher:      consumerFetcher,
		tokenDataService:     tokenDataService,
		poolTokenDataService: poolTokenDataService,
		swapService:          swapService,
		coinGeckoIDs:         coinGeckoIDs}, nil
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

	// Return an unknown event to avoid cases where user failed to check the event type.
	return swapTypes.EventType(len(swapTypes.AllEventTypes()) + 2), false
}

// eventToSwapEvent stores a swap event.
func eventToSwapEvent(event swapTypes.EventLog, chainID uint32) model.SwapEvent {
	var buyer sql.NullString

	if event.GetBuyer() != nil {
		buyer.Valid = true
		buyer.String = event.GetBuyer().String()
	} else {
		buyer.Valid = false
	}

	var provider sql.NullString

	if event.GetProvider() != nil {
		provider.Valid = true
		provider.String = event.GetProvider().String()
	} else {
		provider.Valid = false
	}

	var receiver sql.NullString

	if event.GetReceiver() != nil {
		receiver.Valid = true
		receiver.String = event.GetReceiver().String()
	} else {
		receiver.Valid = false
	}

	return model.SwapEvent{
		InsertTime:      uint64(time.Now().UnixNano()),
		ContractAddress: event.GetContractAddress().String(),
		ChainID:         chainID,
		EventType:       event.GetEventType().Int(),
		BlockNumber:     event.GetBlockNumber(),
		TxHash:          event.GetTxHash().String(),
		EventIndex:      event.GetEventIndex(),
		Sender:          "",
		Buyer:           buyer,
		TokensSold:      event.GetTokensSold(),
		TokensBought:    event.GetTokensBought(),
		SoldID:          event.GetSoldId(),
		BoughtID:        event.GetBoughtId(),
		Provider:        provider,

		Invariant:     event.GetInvariant(),
		LPTokenSupply: event.GetLPTokenSupply(),
		LPTokenAmount: event.GetLPTokenAmount(),
		NewAdminFee:   event.GetNewAdminFee(),
		NewSwapFee:    event.GetNewSwapFee(),
		Amount:        event.GetAmount(),
		AmountFee:     event.GetAmountFee(),
		ProtocolFee:   event.GetProtocolFee(),
		OldA:          event.GetOldA(),
		NewA:          event.GetNewA(),
		InitialTime:   event.GetInitialTime(),
		FutureTime:    event.GetFutureTime(),
		CurrentA:      event.GetCurrentA(),
		Time:          event.GetTime(),
		Receiver:      receiver,

		TimeStamp:    nil,
		TokenPrices:  nil,
		TokenSymbol:  nil,
		TokenDecimal: nil,
	}
}

// ParseAndStore parses the swap logs and returns a model that can be stored
// Deprecated: use Parse and store separately.
func (p *SwapParser) ParseAndStore(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	swapEvent, err := p.Parse(ctx, log, chainID)
	if err != nil {
		return fmt.Errorf("could not parse event: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, &swapEvent)

	if err != nil {
		return fmt.Errorf("could not store event: %w chain: %d address %s", err, chainID, log.Address.String())
	}
	return nil
}

// Parse parses the swap logs.
//
//nolint:gocognit,cyclop,dupl
func (p *SwapParser) Parse(ctx context.Context, log ethTypes.Log, chainID uint32) (interface{}, error) {
	logTopic := log.Topics[0]

	iFace, err := func(log ethTypes.Log) (swapTypes.EventLog, error) {
		switch logTopic {
		case swap.Topic(swapTypes.TokenSwapEvent):
			iFace, err := p.Filterer.ParseTokenSwap(log)
			if err != nil {
				return nil, fmt.Errorf("could not store token swap: %w", err)
			}

			return iFace, nil
		case swap.Topic(swapTypes.AddLiquidityEvent):
			iFace, err := p.Filterer.ParseAddLiquidity(log)
			if err != nil {
				return nil, fmt.Errorf("could not store add liquidity: %w", err)
			}

			return iFace, nil
		case swap.Topic(swapTypes.RemoveLiquidityEvent):
			iFace, err := p.Filterer.ParseRemoveLiquidity(log)
			if err != nil {
				return nil, fmt.Errorf("could not store remove liquidity: %w", err)
			}

			return iFace, nil
		case swap.Topic(swapTypes.RemoveLiquidityOneEvent):
			iFace, err := p.Filterer.ParseRemoveLiquidityOne(log)
			if err != nil {
				return nil, fmt.Errorf("could not store remove liquidity one: %w", err)
			}

			return iFace, nil
		case swap.Topic(swapTypes.RemoveLiquidityImbalanceEvent):
			iFace, err := p.Filterer.ParseRemoveLiquidityImbalance(log)
			if err != nil {
				return nil, fmt.Errorf("could not store remove liquidity imbalance: %w", err)
			}

			return iFace, nil
		case swap.Topic(swapTypes.NewAdminFeeEvent):
			iFace, err := p.Filterer.ParseNewAdminFee(log)
			if err != nil {
				return nil, fmt.Errorf("could not store new admin fee: %w", err)
			}

			return iFace, nil
		case swap.Topic(swapTypes.NewSwapFeeEvent):
			iFace, err := p.Filterer.ParseNewSwapFee(log)
			if err != nil {
				return nil, fmt.Errorf("could not store new swap fee: %w", err)
			}

			return iFace, nil
		case swap.Topic(swapTypes.RampAEvent):
			iFace, err := p.Filterer.ParseRampA(log)
			if err != nil {
				return nil, fmt.Errorf("could not store ramp a: %w", err)
			}

			return iFace, nil
		case swap.Topic(swapTypes.StopRampAEvent):
			iFace, err := p.Filterer.ParseStopRampA(log)
			if err != nil {
				return nil, fmt.Errorf("could not store stop ramp a: %w", err)
			}

			return iFace, nil
		case swap.Topic(swapTypes.FlashLoanEvent):
			iFace, err := p.Filterer.ParseFlashLoan(log)
			if err != nil {
				return nil, fmt.Errorf("could not store flash loan: %w", err)
			}

			return iFace, nil
		default:
			logger.Errorf("ErrUnknownTopic in swap: %s %s chain: %d address: %s", log.TxHash, logTopic.String(), chainID, log.Address.Hex())

			return nil, fmt.Errorf(ErrUnknownTopic)
		}
	}(log)
	if err != nil {
		// Switch failed.
		return nil, err
	}
	swapEvent := eventToSwapEvent(iFace, chainID)

	var sender *string
	var timeStamp *uint64

	timeStamp, sender, err = p.consumerFetcher.FetchTx(ctx, iFace.GetTxHash().String(), int(chainID), int(swapEvent.BlockNumber))
	if err != nil {
		return nil, fmt.Errorf("could not get timestamp, sender on chain %d and tx %s from tx %w", chainID, iFace.GetTxHash().String(), err)
	}

	if *timeStamp == 0 {
		logger.Errorf("empty block time: chain: %d address %s", chainID, log.Address.Hex())
		return nil, fmt.Errorf("empty block time: chain: %d address %s", chainID, log.Address.Hex())
	}

	swapEvent.TimeStamp = timeStamp
	swapEvent.Sender = *sender

	// nolint:nestif
	if swapEvent.Amount != nil {
		tokenPrices := map[uint8]float64{}
		tokenDecimals := map[uint8]uint8{}
		tokenSymbols := map[uint8]string{}

		// Get metadata for each token amount.
		for tokenIndex := range swapEvent.Amount {
			var tokenData tokendata.ImmutableTokenData
			// Get token symbol and decimals from the erc20 contract associated to the token.
			tokenAddress, err := p.poolTokenDataService.GetTokenAddress(ctx, chainID, tokenIndex, swapEvent.ContractAddress)
			if err != nil {
				logger.Errorf("token with index %d not in pool: %v", tokenIndex, err)
				continue
				// return nil, fmt.Errorf("could not get token address: %w", err)
			}
			tokenData, err = p.tokenDataService.GetPoolTokenData(ctx, chainID, *tokenAddress, *p.swapService)
			if err != nil {
				logger.Errorf("could not get token data: %v", err)
				return nil, fmt.Errorf("could not get pool token data: %w", err)
			}
			tokenSymbols[tokenIndex] = tokenData.TokenID()
			tokenDecimals[tokenIndex] = tokenData.Decimals()
			timeStamp, err := p.consumerFetcher.FetchClient.GetBlockTime(ctx, int(chainID), int(iFace.GetBlockNumber()))
			if err != nil {
				return nil, fmt.Errorf("could not get timestamp: %w", err)
			}
			coinGeckoID := p.coinGeckoIDs[tokenData.TokenID()]
			if !(coinGeckoID == "xjewel" && *swapEvent.TimeStamp < 1649030400) && !(coinGeckoID == "synapse-2" && *timeStamp.Response < 1630281600) && !(coinGeckoID == "governance-ohm" && *timeStamp.Response < 1668646800) && !(coinGeckoID == "highstreet" && *timeStamp.Response < 1634263200) {
				tokenPrice, _ := fetcher.GetDefiLlamaData(ctx, *timeStamp.Response, coinGeckoID)
				if tokenPrice != nil {
					tokenPrices[tokenIndex] = *tokenPrice
				}
			}
			swapEvent.TokenPrices = tokenPrices
			swapEvent.TokenDecimal = tokenDecimals
			swapEvent.TokenSymbol = tokenSymbols
		}
	}
	return swapEvent, nil
}
