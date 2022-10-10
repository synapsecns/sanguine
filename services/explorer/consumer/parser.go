package consumer

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
	"github.com/synapsecns/sanguine/services/explorer/db"
	model "github.com/synapsecns/sanguine/services/explorer/db/sql"
	bridgeTypes "github.com/synapsecns/sanguine/services/explorer/types/bridge"
	swapTypes "github.com/synapsecns/sanguine/services/explorer/types/swap"
	"math/big"
	"time"
)

// BoolToUint8 is a helper function to handle bool to uint8 conversion for clickhouse.
func BoolToUint8(input *bool) *uint8 {
	if input == nil {
		return nil
	}
	if *input {
		one := uint8(1)
		return &one
	}
	zero := uint8(0)
	return &zero
}

// ToNullString is a helper function to convert values to null string.
func ToNullString(str *string) sql.NullString {
	var newNullStr sql.NullString
	if str != nil {
		newNullStr.Valid = true
		newNullStr.String = *str
	} else {
		newNullStr.Valid = false
	}
	return newNullStr
}

// Parser parses events and stores them.
type Parser interface {
	// ParseAndStore parses the logs and stores them in the database.
	ParseAndStore(ctx context.Context, log ethTypes.Log, chainID uint32) error
}

// BridgeParser parses events from the bridge contract.
type BridgeParser struct {
	// consumerDB is the database to store parsed data in
	consumerDB db.ConsumerDB
	// Filterer is the bridge Filterer we use to parse events
	Filterer *bridge.SynapseBridgeFilterer
	// bridgeAddress is the address of the bridge
	bridgeAddress common.Address
	// fetcher is a Bridge Config Fetcher
	fetcher BridgeConfigFetcher
	// consumerFetcher is the Fetcher for sender and timestamp
	consumerFetcher *Fetcher
}

// NewBridgeParser creates a new parser for a given bridge.
func NewBridgeParser(consumerDB db.ConsumerDB, bridgeAddress common.Address, bridgeConfigFetcher BridgeConfigFetcher, consumerFetcher *Fetcher) (*BridgeParser, error) {
	filterer, err := bridge.NewSynapseBridgeFilterer(bridgeAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", bridge.SynapseBridgeFilterer{}, err)
	}
	return &BridgeParser{consumerDB, filterer, bridgeAddress, bridgeConfigFetcher, consumerFetcher}, nil
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
	// swap is the address of the bridge
	swapAddress common.Address
	// Filterer is the swap Filterer we use to parse events
	Filterer *swap.SwapFlashLoanFilterer
	// consumerFetcher is the Fetcher for sender and timestamp
	consumerFetcher *Fetcher
	// swapFetcher is the fetcher for token data from swaps.
	swapFetcher SwapFetcher
}

// NewSwapParser creates a new parser for a given bridge.
func NewSwapParser(consumerDB db.ConsumerDB, swapAddress common.Address, swapFetcher SwapFetcher, consumerFetcher *Fetcher) (*SwapParser, error) {
	filterer, err := swap.NewSwapFlashLoanFilterer(swapAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", bridge.SynapseBridgeFilterer{}, err)
	}
	return &SwapParser{consumerDB, swapAddress, filterer, consumerFetcher, swapFetcher}, nil
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

// eventToBridgeEvent stores a bridge event.
func eventToBridgeEvent(event bridgeTypes.EventLog, chainID uint32) model.BridgeEvent {
	var recipient sql.NullString
	if event.GetRecipient() != nil {
		recipient.Valid = true
		recipient.String = event.GetRecipient().String()
	} else {
		recipient.Valid = false
	}
	var recipientBytes sql.NullString
	if event.GetRecipientBytes() != nil {
		recipientBytes.Valid = true
		recipientBytes.String = common.Bytes2Hex(event.GetRecipientBytes()[:])
	} else {
		recipientBytes.Valid = false
	}
	var destinationChainID *big.Int
	if event.GetDestinationChainID() != nil {
		destinationChainID = big.NewInt(int64(event.GetDestinationChainID().Uint64()))
	}
	var tokenIndexFrom *big.Int
	if event.GetTokenIndexFrom() != nil {
		tokenIndexFrom = big.NewInt(int64(*event.GetTokenIndexFrom()))
	}
	var tokenIndexTo *big.Int
	if event.GetTokenIndexTo() != nil {
		tokenIndexTo = big.NewInt(int64(*event.GetTokenIndexTo()))
	}
	var swapSuccess *big.Int
	if event.GetSwapSuccess() != nil {
		swapSuccess = big.NewInt(int64(*BoolToUint8(event.GetSwapSuccess())))
	}
	var swapTokenIndex *big.Int
	if event.GetSwapTokenIndex() != nil {
		swapTokenIndex = big.NewInt(int64(*event.GetSwapTokenIndex()))
	}
	var kappa sql.NullString
	if event.GetKappa() != nil {
		kappa.Valid = true
		kappa.String = common.Bytes2Hex(event.GetKappa()[:])
	} else {
		kappa.Valid = false
	}

	return model.BridgeEvent{
		InsertTime:         uint64(time.Now().UnixNano()),
		ContractAddress:    event.GetContractAddress().String(),
		ChainID:            chainID,
		EventType:          event.GetEventType().Int(),
		BlockNumber:        event.GetBlockNumber(),
		TxHash:             event.GetTxHash().String(),
		Amount:             event.GetAmount(),
		EventIndex:         event.GetEventIndex(),
		DestinationKappa:   crypto.Keccak256Hash(event.GetTxHash().Bytes()).String(),
		Sender:             "",
		Recipient:          recipient,
		RecipientBytes:     recipientBytes,
		DestinationChainID: destinationChainID,
		Token:              event.GetToken().String(),
		Fee:                event.GetFee(),
		Kappa:              kappa,
		TokenIndexFrom:     tokenIndexFrom,
		TokenIndexTo:       tokenIndexTo,
		MinDy:              event.GetMinDy(),
		Deadline:           event.GetDeadline(),
		SwapSuccess:        swapSuccess,
		SwapTokenIndex:     swapTokenIndex,
		SwapMinAmount:      event.GetSwapMinAmount(),
		SwapDeadline:       event.GetSwapDeadline(),

		// placeholders for further data maturation of this event.
		TokenID:      sql.NullString{},
		TimeStamp:    nil,
		AmountUSD:    nil,
		FeeAmountUSD: nil,
		TokenDecimal: nil,
		TokenSymbol:  sql.NullString{},
	}
}

// ParseAndStore parses the bridge logs and stores them in the database.
//
// nolint:gocognit,cyclop,dupl
func (p *BridgeParser) ParseAndStore(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	logTopic := log.Topics[0]
	iFace, err := func(log ethTypes.Log) (bridgeTypes.EventLog, error) {
		switch logTopic {
		case bridge.Topic(bridgeTypes.DepositEvent):
			iFace, err := p.Filterer.ParseTokenDeposit(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse deposit: %w", err)
			}
			return iFace, nil
		case bridge.Topic(bridgeTypes.RedeemEvent):
			iFace, err := p.Filterer.ParseTokenRedeem(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse redeem: %w", err)
			}
			return iFace, nil
		case bridge.Topic(bridgeTypes.WithdrawEvent):
			iFace, err := p.Filterer.ParseTokenWithdraw(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse withdraw: %w", err)
			}
			return iFace, nil
		case bridge.Topic(bridgeTypes.MintEvent):
			iFace, err := p.Filterer.ParseTokenMint(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse mint: %w", err)
			}
			return iFace, nil
		case bridge.Topic(bridgeTypes.DepositAndSwapEvent):
			iFace, err := p.Filterer.ParseTokenDepositAndSwap(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse deposit and swap: %w", err)
			}
			return iFace, nil
		case bridge.Topic(bridgeTypes.MintAndSwapEvent):
			iFace, err := p.Filterer.ParseTokenMintAndSwap(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse mint and swap: %w", err)
			}
			return iFace, nil
		case bridge.Topic(bridgeTypes.RedeemAndSwapEvent):
			iFace, err := p.Filterer.ParseTokenRedeemAndSwap(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse redeem and swap: %w", err)
			}
			return iFace, nil
		case bridge.Topic(bridgeTypes.RedeemAndRemoveEvent):
			iFace, err := p.Filterer.ParseTokenRedeemAndRemove(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse redeem and remove: %w", err)
			}
			return iFace, nil
		case bridge.Topic(bridgeTypes.WithdrawAndRemoveEvent):
			iFace, err := p.Filterer.ParseTokenWithdrawAndRemove(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse withdraw and remove: %w", err)
			}
			return iFace, nil
		case bridge.Topic(bridgeTypes.RedeemV2Event):
			iFace, err := p.Filterer.ParseTokenRedeemV2(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse redeem v2: %w", err)
			}
			return iFace, nil
		default:
			return nil, fmt.Errorf("unknown topic: %s", logTopic.Hex())
		}
	}(log)

	if err != nil {
		// Switch failed.
		return err
	}

	// get TokenID from BridgeConfig data
	tokenID, err := p.fetcher.GetTokenID(ctx, chainID, iFace.GetToken())
	if err != nil {
		return fmt.Errorf("could not parse get token from bridge config event: %w", err)
	}

	// get Token from BridgeConfig data (for getting token decimal but use this for anything else).
	token, err := p.fetcher.GetToken(ctx, chainID, tokenID)
	if err != nil {
		return fmt.Errorf("could not parse get token from bridge config event: %w", err)
	}

	// populate bridge event type so following operations can mature the event data.
	bridgeEvent := eventToBridgeEvent(iFace, chainID)
	// Add TokenID to bridgeEvent
	bridgeEvent.TokenID = ToNullString(tokenID)
	// Add TokenDecimal to bridgeEvent
	bridgeEvent.TokenDecimal = &token.TokenDecimals

	// Get timestamp from consumer
	timeStamp, err := p.consumerFetcher.FetchClient.GetBlockTime(ctx, int(chainID), int(iFace.GetBlockNumber()))
	// If we have a timestamp, populate the following attributes of bridgeEvent.
	if err == nil {
		timeStampBig := uint64(*timeStamp.Response)
		bridgeEvent.TimeStamp = &timeStampBig
		// Add the price of the token at the block the event occurred using coin gecko (to bridgeEvent).
		tokenPrice, symbol := GetTokenMetadataWithTokenID(ctx, *timeStamp.Response, tokenID)
		if tokenPrice != nil {
			// Add AmountUSD to bridgeEvent (if price is not nil)
			bridgeEvent.AmountUSD = GetAmountUSD(iFace.GetAmount(), token.TokenDecimals, tokenPrice)
			// Add FeeAmountUSD to bridgeEvent (if price is not nil)
			bridgeEvent.FeeAmountUSD = GetAmountUSD(iFace.GetFee(), token.TokenDecimals, tokenPrice)
			// Add TokenSymbol to bridgeEvent
			bridgeEvent.TokenSymbol = ToNullString(symbol)
		}
	}

	// sender, err := p.consumerFetcher.FetchClient.GetTxSender(ctx, int(chainID), iFace.GetTxHash().String())
	// if err != nil || sender == nil {
	//	fmt.Println("could not get tx sender: %w", err)
	//	bridgeEvent.Sender = "FAKE_SENDER"
	// } else {
	//	bridgeEvent.Sender = *sender.Response
	//}

	err = p.consumerDB.StoreEvent(ctx, &bridgeEvent, nil)
	if err != nil {
		return fmt.Errorf("could not store event: %w", err)
	}
	return nil
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

// ParseAndStore parses and stores the swap logs.
//
//nolint:gocognit,cyclop,dupl
func (p *SwapParser) ParseAndStore(ctx context.Context, log ethTypes.Log, chainID uint32) error {
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
			return nil, fmt.Errorf("unknown topic: %s", logTopic.Hex())
		}
	}(log)
	if err != nil {
		// Switch failed.
		return err
	}

	// populate swap event type so following operations can mature the event data.
	swapEvent := eventToSwapEvent(iFace, chainID)

	if swapEvent.Amount != nil {
		var tokenPrices map[uint8]float64
		var tokenDecimals map[uint8]uint8
		var tokenSymbols map[uint8]string

		// Get metadata for each token amount
		for tokenIndex, _ := range swapEvent.Amount {
			// get token symbol and decimals from the erc20 contract associated to the token.
			symbol, decimals := p.swapFetcher.GetTokenMetaData(ctx, tokenIndex)
			if symbol != nil && decimals != nil {
				tokenSymbols[tokenIndex] = *symbol
				tokenDecimals[tokenIndex] = *decimals

				// get timestamp of the block where the event occurred.
				timeStamp, err := p.consumerFetcher.FetchClient.GetBlockTime(ctx, int(chainID), int(iFace.GetBlockNumber()))
				if err != nil {
					return fmt.Errorf("could not get timestamp: %w", err)
				}

				// get the token price from the defi llama
				tokenPrice, _ := GetTokenMetadataWithTokenSymbol(ctx, *timeStamp.Response, symbol)
				tokenPrices[tokenIndex] = *tokenPrice
			}
			swapEvent.TokenPrices = tokenPrices
			swapEvent.TokenDecimal = tokenDecimals
			swapEvent.TokenSymbol = tokenSymbols

		}

	}

	// Store bridgeEvent
	err = p.consumerDB.StoreEvent(ctx, nil, &swapEvent)
	if err != nil {
		return fmt.Errorf("could not store event: %w", err)
	}
	return nil
}
