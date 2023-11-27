package parser

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/scribe"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/token"
	"math/big"
	"time"

	"github.com/synapsecns/sanguine/services/explorer/consumer/fetchers/price"

	"golang.org/x/sync/errgroup"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/bridgev1"
	"github.com/synapsecns/sanguine/services/explorer/db"
	model "github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/static"
	bridgeTypes "github.com/synapsecns/sanguine/services/explorer/types/bridge"
)

// BridgeParser parses events from the bridge contract.
type BridgeParser struct {
	// consumerDB is the database to store parsed data in.
	consumerDB db.ConsumerDB
	// Filterer is the bridge Filterer we use to parse events.
	Filterer *bridge.SynapseBridgeFilterer
	// Filterer is the bridge Filterer we use to parse events.
	FiltererV1 *bridgev1.SynapseBridgeFilterer
	// bridgeAddress is the address of the bridge.
	bridgeAddress common.Address
	// tokenDataService contains the token data service/cache
	tokenDataService token.ITokenFetcher
	// tokenPriceService contains the token price service/cache
	tokenPriceService price.IPriceFetcher
	// consumerFetcher is the ScribeFetcher for sender and timestamp.
	consumerFetcher scribe.IScribeFetcher
	// coinGeckoIDs is the mapping of token id to coin gecko ID
	coinGeckoIDs map[string]string
	// fromAPI is true if the parser is being called from the API.
	fromAPI bool
}

const noTokenID = "NO_TOKEN"
const noPrice = "NO_PRICE"

// TODO these parsers need a custom struct with config with the services.

// NewBridgeParser creates a new parser for a given bridge.
func NewBridgeParser(consumerDB db.ConsumerDB, bridgeAddress common.Address, tokenDataService token.ITokenFetcher, consumerFetcher scribe.IScribeFetcher, tokenPriceService price.IPriceFetcher, fromAPI bool) (*BridgeParser, error) {
	filterer, err := bridge.NewSynapseBridgeFilterer(bridgeAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", bridge.SynapseBridgeFilterer{}, err)
	}

	// Old bridge contract to filter all events across all times.
	filtererV1, err := bridgev1.NewSynapseBridgeFilterer(bridgeAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", bridgev1.SynapseBridgeFilterer{}, err)
	}

	idCoinGeckoIDs, err := ParseYaml(static.GetTokenIDToCoingekoConfig())
	if err != nil {
		return nil, fmt.Errorf("could not open yaml file: %w", err)
	}

	return &BridgeParser{
		consumerDB:        consumerDB,
		Filterer:          filterer,
		FiltererV1:        filtererV1,
		bridgeAddress:     bridgeAddress,
		tokenDataService:  tokenDataService,
		tokenPriceService: tokenPriceService,
		consumerFetcher:   consumerFetcher,
		coinGeckoIDs:      idCoinGeckoIDs,
		fromAPI:           fromAPI,
	}, nil
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

	// Return an unknown event to avoid cases where user failed to check the event type.
	return bridgeTypes.EventType(len(bridgeTypes.AllEventTypes()) + 2), false
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
	var destinationKappa string
	if event.GetDestinationChainID() != nil {
		destinationChainID = big.NewInt(int64(event.GetDestinationChainID().Uint64()))
		destinationKappa = crypto.Keccak256Hash([]byte(event.GetTxHash().String())).String()[2:]
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
		DestinationKappa:   destinationKappa,
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

		// Placeholders for further data maturation of this event.
		TimeStamp:    nil,
		AmountUSD:    nil,
		FeeUSD:       nil,
		TokenDecimal: nil,
		TokenSymbol:  sql.NullString{},
	}
}

// ParserType returns the type of parser.
func (p *BridgeParser) ParserType() string {
	return "bridge"
}

// Parse parses the bridge logs and returns a model that can be stored.
func (p *BridgeParser) Parse(ctx context.Context, log ethTypes.Log, chainID uint32) (interface{}, error) {
	bridgeEvent, iFace, err := p.ParseLog(log, chainID)
	if err != nil {
		return nil, err
	}
	bridgeEventInterface, err := p.MatureLogs(ctx, bridgeEvent, iFace, chainID)
	if err != nil {
		return nil, err
	}
	return bridgeEventInterface, nil
}

// ParseLog parses the bridge logs and returns a model that can be stored.
//
// nolint:gocognit,cyclop
func (p *BridgeParser) ParseLog(log ethTypes.Log, chainID uint32) (*model.BridgeEvent, bridgeTypes.EventLog, error) {
	logTopic := log.Topics[0]

	iFace, err := func(log ethTypes.Log) (bridgeTypes.EventLog, error) {
		switch logTopic {
		case bridge.Topic(bridgeTypes.DepositEvent):
			iFace, err := p.Filterer.ParseTokenDeposit(log)
			if err != nil {
				iFaceV1, err := p.FiltererV1.ParseTokenDeposit(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse deposit: %w", err)
				}

				logger.Warnf("used v1 bridge contract to parse deposit")

				return iFaceV1, nil
			}

			return iFace, nil
		case bridge.Topic(bridgeTypes.RedeemEvent):
			iFace, err := p.Filterer.ParseTokenRedeem(log)
			if err != nil {
				iFaceV1, err := p.FiltererV1.ParseTokenRedeem(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse redeem: %w", err)
				}

				logger.Warnf("used v1 bridge contract to parse redeem")

				return iFaceV1, nil
			}

			return iFace, nil
		case bridge.Topic(bridgeTypes.WithdrawEvent):
			iFace, err := p.Filterer.ParseTokenWithdraw(log)
			if err != nil {
				iFaceV1, err := p.FiltererV1.ParseTokenWithdraw(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse withdraw: %w", err)
				}

				logger.Warnf("used v1 bridge contract to parse withdraw")

				return iFaceV1, nil
			}

			return iFace, nil
		case bridge.Topic(bridgeTypes.MintEvent):
			iFace, err := p.Filterer.ParseTokenMint(log)
			if err != nil {
				iFaceV1, err := p.FiltererV1.ParseTokenMint(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse mint: %w", err)
				}

				logger.Warnf("used v1 bridge contract to parse mint")

				return iFaceV1, nil
			}

			return iFace, nil
		case bridge.Topic(bridgeTypes.DepositAndSwapEvent):
			iFace, err := p.Filterer.ParseTokenDepositAndSwap(log)
			if err != nil {
				iFaceV1, err := p.FiltererV1.ParseTokenDepositAndSwap(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse deposit and swap: %w", err)
				}

				logger.Warnf("used v1 bridge contract to parse deposit and swap")

				return iFaceV1, nil
			}

			return iFace, nil
		case bridge.Topic(bridgeTypes.MintAndSwapEvent):
			iFace, err := p.Filterer.ParseTokenMintAndSwap(log)
			if err != nil {
				iFaceV1, err := p.FiltererV1.ParseTokenMintAndSwap(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse mint and swap: %w", err)
				}

				logger.Warnf("used v1 bridge contract to parse mint and swap")

				return iFaceV1, nil
			}

			return iFace, nil
		case bridge.Topic(bridgeTypes.RedeemAndSwapEvent):
			iFace, err := p.Filterer.ParseTokenRedeemAndSwap(log)
			if err != nil {
				iFaceV1, err := p.FiltererV1.ParseTokenRedeemAndSwap(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse redeem and swap: %w", err)
				}

				logger.Warnf("used v1 bridge contract to parse redeem and swap")

				return iFaceV1, nil
			}

			return iFace, nil
		case bridge.Topic(bridgeTypes.RedeemAndRemoveEvent):
			iFace, err := p.Filterer.ParseTokenRedeemAndRemove(log)
			if err != nil {
				iFaceV1, err := p.FiltererV1.ParseTokenRedeemAndRemove(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse redeem and remove: %w", err)
				}

				logger.Warnf("used v1 bridge contract to parse redeem and remove")

				return iFaceV1, nil
			}

			return iFace, nil
		case bridge.Topic(bridgeTypes.WithdrawAndRemoveEvent):
			iFace, err := p.Filterer.ParseTokenWithdrawAndRemove(log)
			if err != nil {
				iFaceV1, err := p.FiltererV1.ParseTokenWithdrawAndRemove(log)
				if err != nil {
					return nil, fmt.Errorf("could not parse withdraw and remove: %w", err)
				}

				logger.Warnf("used v1 bridge contract to parse withdraw and remove")

				return iFaceV1, nil
			}

			return iFace, nil
		case bridge.Topic(bridgeTypes.RedeemV2Event):
			iFace, err := p.Filterer.ParseTokenRedeemV2(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse redeem v2: %w", err)
			}

			return iFace, nil
		default:
			logger.Warnf("ErrUnknownTopic in bridge: %s %s chain: %d address: %s", log.TxHash, logTopic.String(), chainID, log.Address.Hex())

			return nil, fmt.Errorf(ErrUnknownTopic)
		}
	}(log)

	if err != nil {
		// Switch failed.
		return nil, nil, err
	}
	bridgeEvent := eventToBridgeEvent(iFace, chainID)

	return &bridgeEvent, iFace, nil
}

// MatureLogs takes a bridge event and matures it by fetching the sender and timestamp from the API and more.
//
// nolint:gocognit,cyclop
func (p *BridgeParser) MatureLogs(ctx context.Context, bridgeEvent *model.BridgeEvent, iFace bridgeTypes.EventLog, chainID uint32) (interface{}, error) {
	g, groupCtx := errgroup.WithContext(ctx)
	var err error
	var sender *string
	var timeStamp *uint64
	g.Go(func() error {
		if p.fromAPI {
			rawTimeStamp, err := p.consumerFetcher.FetchBlockTime(groupCtx, int(chainID), int(bridgeEvent.BlockNumber))
			if err != nil {
				return fmt.Errorf("could not get timestamp, sender on chain %d and tx %s from tx %w", chainID, iFace.GetTxHash().String(), err)
			}
			uint64TimeStamp := uint64(*rawTimeStamp)
			timeStamp = &uint64TimeStamp
			senderStr := "" // empty for bridge watcher/api parser
			sender = &senderStr
			return nil
		}
		timeStamp, sender, err = p.consumerFetcher.FetchTx(groupCtx, iFace.GetTxHash().String(), int(chainID), int(bridgeEvent.BlockNumber))
		if err != nil {
			return fmt.Errorf("could not get timestamp, sender on chain %d and tx %s from tx %w", chainID, iFace.GetTxHash().String(), err)
		}
		return nil
	})

	var tokenData token.ImmutableTokenData
	fmt.Println("tokenData", tokenData)

	g.Go(func() error {
		// Get Token from BridgeConfig data (for getting token decimal but use this for anything else).
		tokenData, err = p.tokenDataService.GetBridgeTokenData(groupCtx, chainID, iFace.GetToken())
		if err != nil {
			return fmt.Errorf("could not parse get token from bridge config event: %w", err)
		}

		return nil
	})

	err = g.Wait()
	fmt.Println("tokenData", tokenData)
	if err != nil {
		return nil, fmt.Errorf("could not parse bridge event: %w", err)
	}
	if *timeStamp == 0 {
		logger.Errorf("empty block time: chain: %d address %s", chainID, bridgeEvent.ContractAddress)
		return nil, fmt.Errorf("empty block time: chain: %d address %s", chainID, bridgeEvent.ContractAddress)
	}

	bridgeEvent.TimeStamp = timeStamp
	bridgeEvent.Sender = *sender

	if tokenData.TokenID() == token.NoTokenID {
		logger.Errorf("could not get token data token id chain: %d address %s", chainID, bridgeEvent.ContractAddress)
		// handle an inauthentic token.
		return bridgeEvent, nil
	}

	realDecimals := tokenData.Decimals()
	realID := tokenData.TokenID()
	bridgeEvent.TokenDecimal = &realDecimals

	// Add the price of the token at the block the event occurred using coin gecko (to bridgeEvent).
	coinGeckoID := p.coinGeckoIDs[tokenData.TokenID()]
	if coinGeckoID == "" {
		logger.Warnf("BRIDGE - EMPTY TOKEN ID: %s, TokenID: %s", p.coinGeckoIDs[tokenData.TokenID()], tokenData.TokenID())
	}

	// Add TokenSymbol to bridgeEvent.
	bridgeEvent.TokenSymbol = ToNullString(&realID)
	var tokenPrice *float64
	// takes into account an empty bridge token id and for tokens that were bridged before price trackers (coin gecko) had price data.
	if coinGeckoID != "" && !(coinGeckoID == "xjewel" && *timeStamp < 1649030400) && !(coinGeckoID == "synapse-2" && *timeStamp < 1630281600) && !(coinGeckoID == "governance-ohm" && *timeStamp < 1638316800) && !(coinGeckoID == "highstreet" && *timeStamp < 1634263200) {
		tokenPrice = p.tokenPriceService.GetPriceData(ctx, int(*timeStamp), coinGeckoID)
		if tokenPrice == nil && coinGeckoID != noTokenID && coinGeckoID != noPrice {
			return nil, fmt.Errorf("BRIDGE could not get token price for coingeckotoken:  %s chain: %d txhash %s %d", coinGeckoID, chainID, bridgeEvent.TxHash, bridgeEvent.TimeStamp)
		}
	}

	if tokenPrice != nil {
		bridgeEvent.AmountUSD = GetAmountUSD(bridgeEvent.Amount, tokenData.Decimals(), tokenPrice)
		if iFace.GetFee() != nil {
			bridgeEvent.FeeUSD = GetAmountUSD(bridgeEvent.Fee, tokenData.Decimals(), tokenPrice)
		}
	}
	return bridgeEvent, nil
}
