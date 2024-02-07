package parser

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher/tokenprice"
	"golang.org/x/sync/errgroup"

	"github.com/synapsecns/sanguine/services/explorer/consumer/parser/tokendata"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser/tokenpool"
	"github.com/synapsecns/sanguine/services/explorer/contracts/metaswap"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
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
	consumerFetcher fetcher.ScribeFetcher
	// tokenDataService contains the token data service/cache
	tokenDataService tokendata.Service
	// poolTokenDataService get the token address from the token index
	poolTokenDataService tokenpool.Service
	// tokenPriceService get the token price from the coingecko id
	tokenPriceService tokenprice.Service
	// swapService is the swap service
	swapService fetcher.SwapService
	// FiltererMetaSwap is the meta swap Filterer we use to parse events.
	FiltererMetaSwap *metaswap.MetaSwapFilterer
	// coinGeckoIDs is a mapping from coin token symbol to coin gecko ID
	coinGeckoIDs map[string]string
}

// NewSwapParser creates a new parser for a given bridge.
func NewSwapParser(consumerDB db.ConsumerDB, swapAddress common.Address, metaSwap bool, consumerFetcher fetcher.ScribeFetcher, swapService fetcher.SwapService, tokenDataService tokendata.Service, tokenPriceService tokenprice.Service) (*SwapParser, error) {
	var filterer *swap.SwapFlashLoanFilterer
	var filtererMetaSwap *metaswap.MetaSwapFilterer
	var err error
	if metaSwap {
		filtererMetaSwap, err = metaswap.NewMetaSwapFilterer(swapAddress, nil)
		if err != nil {
			return nil, fmt.Errorf("could not create %T: %w", metaswap.MetaSwapFilterer{}, err)
		}
		filterer = nil
	} else {
		filterer, err = swap.NewSwapFlashLoanFilterer(swapAddress, nil)
		if err != nil {
			return nil, fmt.Errorf("could not create %T: %w", swap.SwapFlashLoanFilterer{}, err)
		}
		filtererMetaSwap = nil
	}

	poolTokenDataService, err := tokenpool.NewPoolTokenDataService(swapService, consumerDB)
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
		tokenPriceService:    tokenPriceService,
		swapService:          swapService,
		FiltererMetaSwap:     filtererMetaSwap,
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
		SoldID:          event.GetSoldID(),
		BoughtID:        event.GetBoughtID(),
		Provider:        provider,

		Invariant:     event.GetInvariant(),
		LPTokenSupply: event.GetLPTokenSupply(),
		LPTokenAmount: event.GetLPTokenAmount(),
		NewAdminFee:   event.GetNewAdminFee(),
		NewSwapFee:    event.GetNewSwapFee(),
		Amount:        event.GetAmount(),
		Fee:           event.GetAmountFee(),
		ProtocolFee:   event.GetProtocolFee(),
		OldA:          event.GetOldA(),
		NewA:          event.GetNewA(),
		InitialTime:   event.GetInitialTime(),
		FutureTime:    event.GetFutureTime(),
		CurrentA:      event.GetCurrentA(),
		Time:          event.GetTime(),
		Receiver:      receiver,

		TimeStamp:    nil,
		TokenPrice:   nil,
		TokenSymbol:  nil,
		TokenDecimal: nil,
		AdminFee:     nil,
		FeeUSD:       nil,
		AdminFeeUSD:  nil,
		AmountUSD:    nil,
	}
}

// ParserType returns the type of parser.
func (p *SwapParser) ParserType() string {
	return "swap"
}

// Parse parses the swap logs.
//
//nolint:gocognit,cyclop,dupl,maintidx
func (p *SwapParser) Parse(ctx context.Context, log ethTypes.Log, chainID uint32) (interface{}, error) {
	logTopic := log.Topics[0]

	iFace, err := func(log ethTypes.Log) (swapTypes.EventLog, error) {
		// nolint:nestif
		if p.FiltererMetaSwap != nil {
			switch logTopic {
			case swap.Topic(swapTypes.TokenSwapUnderlyingEvent):

				iFace, err := p.FiltererMetaSwap.ParseTokenSwapUnderlying(log)
				if err != nil {
					return nil, fmt.Errorf("could not store token swap underlying: %w", err)
				}

				return iFace, nil
			default:
				logger.Warnf("ErrUnknownTopic in meta swap: %s %s chain: %d address: %s", log.TxHash, logTopic.String(), chainID, log.Address.Hex())

				return nil, fmt.Errorf(ErrUnknownTopic)
			}
		} else {
			switch logTopic {
			case swap.Topic(swapTypes.TokenSwapEvent):

				iFace, err := p.Filterer.ParseTokenSwap(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse token swap: %w", err)
				}

				return iFace, nil
			case swap.Topic(swapTypes.AddLiquidityEvent):
				iFace, err := p.Filterer.ParseAddLiquidity(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse add liquidity: %w", err)
				}

				return iFace, nil
			case swap.Topic(swapTypes.RemoveLiquidityEvent):
				iFace, err := p.Filterer.ParseRemoveLiquidity(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse remove liquidity: %w", err)
				}

				return iFace, nil
			case swap.Topic(swapTypes.RemoveLiquidityOneEvent):
				iFace, err := p.Filterer.ParseRemoveLiquidityOne(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse remove liquidity one: %w", err)
				}

				return iFace, nil
			case swap.Topic(swapTypes.RemoveLiquidityImbalanceEvent):
				iFace, err := p.Filterer.ParseRemoveLiquidityImbalance(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse remove liquidity imbalance: %w", err)
				}

				return iFace, nil
			case swap.Topic(swapTypes.NewAdminFeeEvent):
				iFace, err := p.Filterer.ParseNewAdminFee(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse new admin fee: %w", err)
				}
				err = p.consumerDB.StoreSwapFee(ctx, chainID, log.BlockNumber, log.Address.String(), iFace.NewAdminFee.Uint64(), "admin")
				if err != nil {
					return nil, fmt.Errorf("could not store new admin fee : %w", err)
				}
				return iFace, nil
			case swap.Topic(swapTypes.NewSwapFeeEvent):
				iFace, err := p.Filterer.ParseNewSwapFee(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse new swap fee: %w", err)
				}
				err = p.consumerDB.StoreSwapFee(ctx, chainID, log.BlockNumber, log.Address.String(), iFace.NewSwapFee.Uint64(), "swap")
				if err != nil {
					return nil, fmt.Errorf("could not store new admin fee : %w", err)
				}
				return iFace, nil
			case swap.Topic(swapTypes.RampAEvent):
				iFace, err := p.Filterer.ParseRampA(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse ramp a: %w", err)
				}

				return iFace, nil
			case swap.Topic(swapTypes.StopRampAEvent):
				iFace, err := p.Filterer.ParseStopRampA(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse stop ramp a: %w", err)
				}

				return iFace, nil
			case swap.Topic(swapTypes.FlashLoanEvent):
				iFace, err := p.Filterer.ParseFlashLoan(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse flash loan: %w", err)
				}

				return iFace, nil

			default:
				logger.Warnf("ErrUnknownTopic in swap: %s %s chain: %d address: %s", log.TxHash, logTopic.String(), chainID, log.Address.Hex())

				return nil, fmt.Errorf(ErrUnknownTopic)
			}
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

	maxIndex := uint8(0)
	totalTokenIndexRange := make(map[uint8]bool)
	if swapEvent.Amount != nil {
		for k := range swapEvent.Amount {
			totalTokenIndexRange[k] = true
			if k > maxIndex {
				maxIndex = k
			}
		}
	}
	if swapEvent.Fee != nil {
		for k := range swapEvent.Fee {
			totalTokenIndexRange[k] = true
			if k > maxIndex {
				maxIndex = k
			}
		}
	}

	adminFee := uint64(6000000000)
	swapFee := uint64(4000000)

	tokenPricesArr := make([]float64, maxIndex+1)
	tokenDecimalsArr := make([]uint8, maxIndex+1)
	tokenSymbolsArr := make([]string, maxIndex+1)
	tokenCoinGeckoIDsArr := make([]string, maxIndex+1)

	tokenPrices := make(map[uint8]float64, len(totalTokenIndexRange))
	tokenDecimals := make(map[uint8]uint8, len(totalTokenIndexRange))
	tokenSymbols := make(map[uint8]string, len(totalTokenIndexRange))
	tokenCoinGeckoIDs := make(map[uint8]string, len(totalTokenIndexRange))
	g, groupCtx := errgroup.WithContext(ctx)

	// TODO: need to deploy the test swap contracts with token indexes that match the test token address
	// nolint:nestif
	if !core.IsTest() {
		for i := range totalTokenIndexRange {
			tokenIndex := i
			g.Go(func() error {
				var tokenData tokendata.ImmutableTokenData
				// Get token symbol and decimals from the erc20 contract associated to the token.
				tokenAddress, err := p.poolTokenDataService.GetTokenAddress(groupCtx, chainID, tokenIndex, swapEvent.ContractAddress)
				if err != nil {
					logger.Errorf("token with index %d not in pool: %v, %d, %s, %v %s, %d, %v", tokenIndex, err, chainID, swapEvent.ContractAddress, swapEvent.Amount, swapEvent.TxHash, swapEvent.EventType, p.FiltererMetaSwap)
					return fmt.Errorf("token with index %d not in pool: %w, %d, %s, %v %s, %d, %v", tokenIndex, err, chainID, swapEvent.ContractAddress, swapEvent.Amount, swapEvent.TxHash, swapEvent.EventType, p.FiltererMetaSwap)
				}
				tokenData, err = p.tokenDataService.GetPoolTokenData(groupCtx, chainID, *tokenAddress, p.swapService)
				if err != nil {
					logger.Errorf("could not get token data: %v", err)
					return fmt.Errorf("could not get pool token data: %w", err)
				}
				tokenSymbolsArr[tokenIndex] = tokenData.TokenID()
				tokenDecimalsArr[tokenIndex] = tokenData.Decimals()
				coinGeckoID := p.coinGeckoIDs[tokenData.TokenID()]
				tokenCoinGeckoIDsArr[tokenIndex] = coinGeckoID

				if !(coinGeckoID == "xjewel" && *timeStamp < 1649030400) && !(coinGeckoID == "synapse-2" && *timeStamp < 1630281600) && !(coinGeckoID == "governance-ohm" && *timeStamp < 1638316800) && !(coinGeckoID == "highstreet" && *timeStamp < 1634263200) {
					tokenPrice := p.tokenPriceService.GetPriceData(groupCtx, int(*swapEvent.TimeStamp), coinGeckoID)
					if (tokenPrice == nil) && coinGeckoID != noTokenID && coinGeckoID != noPrice {
						return fmt.Errorf("SWAP could not get token price for coingeckotoken:  %s chain: %d txhash %s %d", coinGeckoID, chainID, swapEvent.TxHash, swapEvent.TimeStamp)
					}
					tokenPricesArr[tokenIndex] = *tokenPrice
				}

				// TODO DELETE
				if tokenPricesArr[tokenIndex] == 0 {
					logger.Warnf("SWAP - TOKEN PRICE IS ZERO tokenPricesArr[tokenIndex]: s%s, chainID: %d, tokenIndex: %d, tokenAddress: %s", tokenPricesArr[tokenIndex], chainID, tokenIndex, tokenAddress)
				}
				return nil
			})
		}
	}
	g.Go(func() error {
		// Check for all fee emitting event types
		if swapEvent.EventType == 0 || swapEvent.EventType == 1 || swapEvent.EventType == 4 || swapEvent.EventType == 9 || swapEvent.EventType == 10 {
			dbAdminFee, dbSwapFee, err := p.GetCorrectSwapFee(ctx, swapEvent)
			if err != nil {
				return fmt.Errorf("could not process swap event: %w", err)
			}
			if dbAdminFee > 0 {
				logger.Infof("USING DIFFERENT ADMIN FEE: %d", dbAdminFee)
				adminFee = dbAdminFee
			}
			if dbSwapFee > 0 {
				logger.Infof("USING DIFFERENT SWAP FEE: %d", dbSwapFee)
				swapFee = dbSwapFee
			}
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could get all swap data from go routines: %w", err)
	}

	for i := range totalTokenIndexRange {
		tokenPrices[i] = tokenPricesArr[i]
		tokenDecimals[i] = tokenDecimalsArr[i]
		tokenSymbols[i] = tokenSymbolsArr[i]
		tokenCoinGeckoIDs[i] = tokenCoinGeckoIDsArr[i]
	}

	swapEvent.TokenPrice = tokenPrices
	swapEvent.TokenDecimal = tokenDecimals
	swapEvent.TokenSymbol = tokenSymbols
	swapEvent.TokenCoinGeckoID = tokenCoinGeckoIDs

	amountResults := make(map[uint8]float64, len(totalTokenIndexRange))
	feeResults := make(map[uint8]float64, len(totalTokenIndexRange))
	adminFeeResults := make(map[uint8]float64, len(totalTokenIndexRange))
	adminFeeAmountResults := make(map[uint8]string, len(totalTokenIndexRange))

	if swapEvent.EventType == 0 || swapEvent.EventType == 10 {
		fee, err := convertFee(swapEvent.TokensSold, swapEvent.TokenDecimal[uint8(swapEvent.SoldID.Uint64())], swapFee)
		if err != nil {
			return nil, fmt.Errorf("could not convert fee: %w %d %d %d %d", err, swapEvent.TokensSold, uint8(swapEvent.SoldID.Uint64()), swapFee, adminFee)
		}
		swapEvent.Fee[uint8(swapEvent.SoldID.Uint64())] = fee
	}

	for i := range swapEvent.Amount {
		n := new(big.Int)
		n, ok := n.SetString(swapEvent.Amount[i], 10)
		if !ok {
			return nil, fmt.Errorf("error in parsing amount %s", swapEvent.Amount[i])
		}
		price := swapEvent.TokenPrice[i]
		amountResults[i] = *GetAmountUSD(n, swapEvent.TokenDecimal[i], &price)
	}

	for i := range swapEvent.Fee {
		n := new(big.Int)
		n, ok := n.SetString(swapEvent.Fee[i], 10)
		if !ok {
			return nil, fmt.Errorf("error in parsing fee amount %s", swapEvent.Fee[i])
		}

		price := swapEvent.TokenPrice[i]
		feeResults[i] = *GetAmountUSD(n, swapEvent.TokenDecimal[i], &price)
		adminFeeAmountResults[i], err = convertFee(n, swapEvent.TokenDecimal[i], adminFee)
		if err != nil {
			return nil, fmt.Errorf("could not convert fee: %w %d %d %d %d", err, swapEvent.TokensSold, uint8(swapEvent.SoldID.Uint64()), swapFee, adminFee)
		}

		adminFeeResults[i] = feeResults[i] * getAdjustedFee(adminFee, 10)
	}

	swapEvent.AmountUSD = amountResults
	swapEvent.FeeUSD = feeResults
	swapEvent.AdminFee = adminFeeAmountResults
	swapEvent.AdminFeeUSD = adminFeeResults

	return swapEvent, nil
}

// convertFee gets the fee amount.
func convertFee(amount *big.Int, decimal uint8, feeAmount uint64) (string, error) {
	adjustedAmount := GetAdjustedAmount(amount, decimal)
	if adjustedAmount == nil {
		return "", fmt.Errorf("SWAP - adjusted amount IS NIL %d", adjustedAmount)
	}

	fee := big.NewFloat(*adjustedAmount * getAdjustedFee(feeAmount, 10))
	feeDecimals := big.NewFloat(math.Pow(10, float64(decimal)))
	fee.Mul(fee, feeDecimals)

	result := new(big.Int)
	fee.Int(result)

	return fee.Text('f', 0), nil
}

func getAdjustedFee(fee uint64, decimal uint8) float64 {
	return float64(fee) / math.Pow(10, float64(decimal))
}

// TODO make more dynamic

// GetCorrectSwapFee returns the correct swap fee for the given pool contract.
func (p *SwapParser) GetCorrectSwapFee(ctx context.Context, swapEvent model.SwapEvent) (uint64, uint64, error) {
	var dbAdminFee uint64
	var dbSwapFee uint64
	var err error
	g, groupCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		dbAdminFee, err = p.consumerDB.GetUint64(groupCtx, fmt.Sprintf("SELECT fee FROM swap_fees WHERE chain_id = %d AND contract_address = '%s' AND fee_type = '%s' AND block_number <= %d ORDER BY block_number DESC LIMIT 1", swapEvent.ChainID, swapEvent.ContractAddress, "admin", swapEvent.BlockNumber))
		if err != nil {
			return fmt.Errorf("could not get admin fee: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		dbSwapFee, err = p.consumerDB.GetUint64(groupCtx, fmt.Sprintf("SELECT fee FROM swap_fees WHERE chain_id = %d AND contract_address = '%s' AND fee_type = '%s' AND block_number <= %d ORDER BY block_number DESC LIMIT 1", swapEvent.ChainID, swapEvent.ContractAddress, "swap", swapEvent.BlockNumber))
		if err != nil {
			return fmt.Errorf("could not get swap fee: %w", err)
		}
		return nil
	})
	err = g.Wait()
	if err != nil {
		return 0, 0, fmt.Errorf("could notget newest swap fees: %w", err)
	}
	return dbAdminFee, dbSwapFee, nil
}
