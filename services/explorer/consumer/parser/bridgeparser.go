package parser

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge/bridgev1"
	"github.com/synapsecns/sanguine/services/explorer/db"
	model "github.com/synapsecns/sanguine/services/explorer/db/sql"
	bridgeTypes "github.com/synapsecns/sanguine/services/explorer/types/bridge"
	"math/big"
	"path/filepath"
	"time"
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
	fetcher fetcher.BridgeConfigFetcher
	// consumerFetcher is the ScribeFetcher for sender and timestamp.
	consumerFetcher *fetcher.ScribeFetcher
	// coinGeckoIDs is the mapping of token id to coin gecko ID
	coinGeckoIDs map[string]string
}

// NewBridgeParser creates a new parser for a given bridge.
func NewBridgeParser(consumerDB db.ConsumerDB, bridgeAddress common.Address, bridgeConfigFetcher fetcher.BridgeConfigFetcher, consumerFetcher *fetcher.ScribeFetcher) (*BridgeParser, error) {
	filterer, err := bridge.NewSynapseBridgeFilterer(bridgeAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", bridge.SynapseBridgeFilterer{}, err)
	}

	// Old bridge contract to filter all events across all times.
	filtererV1, err := bridgev1.NewSynapseBridgeFilterer(bridgeAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", bridgev1.SynapseBridgeFilterer{}, err)
	}
	idPath := filepath.Clean("../static/tokenIDToCoinGeckoID.yaml")
	if err != nil {
		return nil, fmt.Errorf("could find path to yaml file: %w", err)
	}
	idCoinGeckoIDs, err := OpenYaml(idPath)
	if err != nil {
		return nil, fmt.Errorf("could not open yaml file: %w", err)
	}

	return &BridgeParser{consumerDB, filterer, filtererV1, bridgeAddress, bridgeConfigFetcher, consumerFetcher, idCoinGeckoIDs}, nil
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
		Recipient:          ToNullString(event.GetRecipient()),
		RecipientBytes:     ToNullString(event.GetRecipientBytes()),
		DestinationChainID: event.GetDestinationChainID(),
		Token:              event.GetToken().String(),
		Fee:                event.GetFee(),
		Kappa:              ToNullString(event.GetKappa()),
		TokenIndexFrom:     event.GetTokenIndexFrom(),
		TokenIndexTo:       event.GetTokenIndexTo(),
		MinDy:              event.GetMinDy(),
		Deadline:           event.GetDeadline(),
		SwapSuccess:        BoolToUint8(event.GetSwapSuccess()),
		SwapTokenIndex:     event.GetSwapTokenIndex(),
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

// ParseAndStore parses the bridge logs and stores them in the database.
//
// nolint:gocognit,cyclop,dupl,maintidx
func (p *BridgeParser) ParseAndStore(ctx context.Context, log ethTypes.Log, chainID uint32) error {
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
			return nil, fmt.Errorf("unknown topic: %s %s", logTopic.Hex(), logTopic.String())
		}
	}(log)
	if err != nil {
		// Switch failed.
		return err
	}

	// Get TokenID from BridgeConfig data.
	tokenID, err := p.fetcher.GetTokenID(ctx, big.NewInt(int64(chainID)), iFace.GetToken())
	if err != nil {
		return fmt.Errorf("could not parse get token from bridge config event: %w", err)
	}

	// Get Token from BridgeConfig data (for getting token decimal but use this for anything else).
	token, err := p.fetcher.GetToken(ctx, chainID, tokenID, uint32(iFace.GetBlockNumber()))
	if err != nil {
		return fmt.Errorf("could not parse get token from bridge config event: %w", err)
	}

	bridgeEvent := eventToBridgeEvent(iFace, chainID)
	bridgeEvent.TokenID = ToNullString(tokenID)
	bridgeEvent.TokenDecimal = &token.TokenDecimals
	timeStamp, err := p.consumerFetcher.FetchClient.GetBlockTime(ctx, int(chainID), int(iFace.GetBlockNumber()))
	if err != nil {
		return fmt.Errorf("could not get block time: %w", err)
	}

	timeStampBig := uint64(*timeStamp.Response)
	bridgeEvent.TimeStamp = &timeStampBig

	// Add the price of the token at the block the event occurred using coin gecko (to bridgeEvent).
	coinGeckoID := p.coinGeckoIDs[*tokenID]
	tokenPrice, symbol := fetcher.GetDefiLlamaData(ctx, *timeStamp.Response, coinGeckoID)
	if tokenPrice != nil {
		// Add AmountUSD to bridgeEvent (if price is not nil).
		bridgeEvent.AmountUSD = GetAmountUSD(iFace.GetAmount(), token.TokenDecimals, tokenPrice)

		// Add FeeAmountUSD to bridgeEvent (if price is not nil).
		if iFace.GetFee() != nil {
			bridgeEvent.FeeAmountUSD = GetAmountUSD(iFace.GetFee(), token.TokenDecimals, tokenPrice)
		}

		// Add TokenSymbol to bridgeEvent.
		bridgeEvent.TokenSymbol = ToNullString(symbol)
	}

	sender, err := p.consumerFetcher.FetchTxSender(ctx, chainID, iFace.GetTxHash().String())
	if err != nil {
		logger.Errorf("could not get tx sender: %v", err)
	}

	bridgeEvent.Sender = sender
	err = p.consumerDB.StoreEvent(ctx, &bridgeEvent, nil)
	if err != nil {
		return fmt.Errorf("could not store event: %w", err)
	}

	return nil
}
