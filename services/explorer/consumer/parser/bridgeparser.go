package parser

import (
	"context"
	"database/sql"
	"fmt"
	"math/big"
	"time"

	"github.com/synapsecns/sanguine/services/explorer/consumer/parser/tokendata"
	"golang.org/x/sync/errgroup"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
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
	// fetcher is a Bridge Config ScribeFetcher.
	fetcher fetcher.Service
	// tokenDataService contains the token data service/cache
	tokenDataService tokendata.Service
	// consumerFetcher is the ScribeFetcher for sender and timestamp.
	consumerFetcher *fetcher.ScribeFetcher
	// coinGeckoIDs is the mapping of token id to coin gecko ID
	coinGeckoIDs map[string]string
}

// NewBridgeParser creates a new parser for a given bridge.
func NewBridgeParser(consumerDB db.ConsumerDB, bridgeAddress common.Address, bridgeConfigFetcher fetcher.Service, consumerFetcher *fetcher.ScribeFetcher) (*BridgeParser, error) {
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

	tokenDataService, err := tokendata.NewTokenDataService(bridgeConfigFetcher)
	if err != nil {
		return nil, fmt.Errorf("could not create token data service: %w", err)
	}

	return &BridgeParser{
		consumerDB:       consumerDB,
		Filterer:         filterer,
		FiltererV1:       filtererV1,
		bridgeAddress:    bridgeAddress,
		fetcher:          bridgeConfigFetcher,
		tokenDataService: tokenDataService,
		consumerFetcher:  consumerFetcher,
		coinGeckoIDs:     idCoinGeckoIDs,
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
		TokenID:      sql.NullString{},
		TimeStamp:    nil,
		AmountUSD:    nil,
		FeeAmountUSD: nil,
		TokenDecimal: nil,
		TokenSymbol:  sql.NullString{},
	}
}

// ParseAndStore parses the bridge logs and returns a model that can be stored
// Deprecated: use Parse and store separately.
func (p *BridgeParser) ParseAndStore(ctx context.Context, log ethTypes.Log, chainID uint32) error {
	bridgeEvent, err := p.Parse(ctx, log, chainID)
	if err != nil {
		return fmt.Errorf("could not parse event: %w", err)
	}
	err = p.consumerDB.StoreEvent(ctx, &bridgeEvent)

	if err != nil {
		return fmt.Errorf("could not store event: %w chain: %d address %s", err, chainID, log.Address.String())
	}
	return nil
}

// Parse parses the bridge logs and returns a model that can be stored
//
// nolint:gocognit,cyclop,dupl,maintidx
func (p *BridgeParser) Parse(ctx context.Context, log ethTypes.Log, chainID uint32) (interface{}, error) {
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
			logger.Errorf("ErrUnknownTopic in bridge: %s %s chain: %d address: %s", log.TxHash, logTopic.String(), chainID, log.Address.Hex())

			return nil, fmt.Errorf(ErrUnknownTopic)
		}
	}(log)

	if err != nil {
		// Switch failed.
		return nil, err
	}

	bridgeEvent := eventToBridgeEvent(iFace, chainID)
	g, groupCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		bridgeEvent.Sender, err = p.consumerFetcher.FetchTxSender(groupCtx, chainID, iFace.GetTxHash().String())
		if err != nil {
			logger.Errorf("could not get tx sender: %v", err)
		}
		return nil
	})

	var timeStampBig uint64
	var timeStamp *int
	g.Go(func() error {
		timeStamp, err = p.consumerFetcher.FetchBlockTime(groupCtx, int(chainID), int(iFace.GetBlockNumber()))
		if err != nil {
			return fmt.Errorf("could unknownTopic get block time: %w, %d, %d", err, int(chainID), int(iFace.GetBlockNumber()))
		}
		timeStampBig = uint64(*timeStamp)
		return nil
	})

	var tokenData tokendata.ImmutableTokenData
	g.Go(func() error {
		// Get Token from BridgeConfig data (for getting token decimal but use this for anything else).
		tokenData, err = p.tokenDataService.GetTokenData(ctx, chainID, iFace.GetToken())
		if err != nil {
			return fmt.Errorf("could not parse get token from bridge config event: %w", err)
		}

		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could not parse bridge event: %w", err)
	}
	if *timeStamp == 0 {
		logger.Errorf("empty block time: chain: %d address %s", chainID, log.Address.Hex())
		return nil, fmt.Errorf("empty block time: chain: %d address %s", chainID, log.Address.Hex())
	}

	bridgeEvent.TimeStamp = &timeStampBig

	if tokenData.TokenID() == fetcher.NoTokenID {
		// handle an inauthentic token.
		return &bridgeEvent, nil
	}

	realDecimals := tokenData.Decimals()
	realID := tokenData.TokenID()
	bridgeEvent.TokenDecimal = &realDecimals

	// Add the price of the token at the block the event occurred using coin gecko (to bridgeEvent).
	coinGeckoID := p.coinGeckoIDs[tokenData.TokenID()]

	// Add TokenSymbol to bridgeEvent.
	bridgeEvent.TokenSymbol = ToNullString(&realID)
	var tokenPrice *float64
	if !(coinGeckoID == "xjewel" && *timeStamp < 1649030400) && !(coinGeckoID == "synapse-2" && *timeStamp < 1619827200) && !(coinGeckoID == "governance-ohm" && *timeStamp < 1668646800) {
		tokenPrice, _ = fetcher.GetDefiLlamaData(ctx, *timeStamp, coinGeckoID)
	}
	if tokenPrice != nil {
		// Add AmountUSD to bridgeEvent (if price is not nil).
		bridgeEvent.AmountUSD = GetAmountUSD(iFace.GetAmount(), tokenData.Decimals(), tokenPrice)

		// Add FeeAmountUSD to bridgeEvent (if price is not nil).
		if iFace.GetFee() != nil {
			bridgeEvent.FeeAmountUSD = GetAmountUSD(iFace.GetFee(), tokenData.Decimals(), tokenPrice)
		}
	}

	return bridgeEvent, nil
}
