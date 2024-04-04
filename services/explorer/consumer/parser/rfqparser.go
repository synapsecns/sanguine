package parser

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser/tokendata"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher/tokenprice"
	"github.com/synapsecns/sanguine/services/explorer/db"
	model "github.com/synapsecns/sanguine/services/explorer/db/sql"
	bridgeTypes "github.com/synapsecns/sanguine/services/explorer/types/bridge"
	rfqTypes "github.com/synapsecns/sanguine/services/explorer/types/rfq"
)

// RFQParser parsers rfq logs
type RFQParser struct {
	// consumerDB is the database to store parsed data in
	consumerDB db.ConsumerDB
	// Filterer is the message Filterer we use to parse events
	Filterer *rfq.SynapseRFQFilterer
	// messageAddress is the address of the message
	rfqAddress common.Address
	// consumerFetcher is the Fetcher for sender and timestamp
	consumerFetcher fetcher.ScribeFetcher
	// rfqService is the rfq service for getting token symbol information
	rfqService fetcher.RFQService
	// tokenDataService contains the token data service/cache
	tokenDataService tokendata.Service
	// tokenPriceService contains the token price service/cache
	tokenPriceService tokenprice.Service
	// fromAPI is true if the parser is being called from the API.
	fromAPI bool
}

const usdcAddress = "usd-coin"
const usdcDecimals = 6
const ethAddress = "eth"
const ethDecimals = 18

// NewRFQParser creates a new RFQParser
func NewRFQParser(consumerDB db.ConsumerDB, rfqAddress common.Address, consumerFetcher fetcher.ScribeFetcher, rfqService fetcher.RFQService, tokenDataService tokendata.Service, tokenPriceService tokenprice.Service, fromAPI bool) (*RFQParser, error) {
	filterer, err := rfq.NewSynapseRFQFilterer(rfqAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", rfq.SynapseRFQFilterer{}, err)
	}
	return &RFQParser{consumerDB, filterer, rfqAddress, consumerFetcher, rfqService, tokenDataService, tokenPriceService, fromAPI}, nil
}

// ParserType returns the type of parser
func (p *RFQParser) ParserType() string {
	return "rfq"
}

// ParseLog log converts an eth log to a rfq event type
func (p *RFQParser) ParseLog(log ethTypes.Log, chainID uint32) (*model.RFQEvent, rfqTypes.EventLog, error) {
	logTopic := log.Topics[0]
	iFace, err := func(log ethTypes.Log) (rfqTypes.EventLog, error) {
		switch logTopic {
		case rfq.Topic(rfqTypes.BridgeRequestedEvent):
			iFace, err := p.Filterer.ParseBridgeRequestSent(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse rfq event requested : %w", err)
			}
			return iFace, nil
		case rfq.Topic(rfqTypes.BridgeRelayedEvent):
			iFace, err := p.Filterer.ParseBridgeRelayed(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse rfq event relayed: %w", err)
			}
			return iFace, nil

		default:
			logger.Warnf("ErrUnknownTopic in rfq: %s %s chain: %d address: %s", log.TxHash, logTopic.String(), chainID, log.Address.Hex())

			return nil, fmt.Errorf(ErrUnknownTopic)
		}
	}(log)

	if err != nil {
		// Switch failed.

		return nil, nil, err
	}
	if iFace == nil {
		// Unknown topic.
		return nil, nil, fmt.Errorf("unknwn topic")
	}

	// Populate rfq event type so following operations can mature the event data.
	rfqEvent := eventToRFQEvent(iFace, chainID)
	return &rfqEvent, iFace, nil
}

// MatureLogs takes a rfq event and adds data to them.
func (p *RFQParser) MatureLogs(ctx context.Context, rfqEvent *model.RFQEvent, iFace rfqTypes.EventLog, chainID uint32) (interface{}, error) {
	// Get timestamp from consumer
	timeStamp, err := p.consumerFetcher.FetchBlockTime(ctx, int(chainID), int(iFace.GetBlockNumber()))
	if err != nil {
		return nil, fmt.Errorf("could not get block time: %w", err)
	}

	// If we have a timestamp, populate the following attributes of rfqEvent.
	timeStampBig := uint64(*timeStamp)
	rfqEvent.TimeStamp = &timeStampBig

	tokenData, err := p.tokenDataService.GetTokenData(ctx, chainID, common.HexToAddress(rfqEvent.OriginToken))
	if err != nil {
		logger.Errorf("could not get token data: %v", err)
		return nil, fmt.Errorf("could not get pool token data: %w", err)
	}
	decimals := tokenData.Decimals()
	rfqEvent.TokenSymbol = tokenData.TokenID()

	rfqEvent.TokenDecimal = &decimals
	p.applyPriceData(ctx, rfqEvent, usdcCoinGeckoID)

	// Would store into bridge database with a new goroutine but saw unreliable storage of events w/parent context cancellation.
	bridgeEvent := rfqEventToBridgeEvent(*rfqEvent)
	if p.fromAPI {
		return bridgeEvent, nil
	}
	err = p.storeBridgeEvent(ctx, bridgeEvent)
	if err != nil {
		logger.Errorf("could not store cctp event into bridge database: %v", err)
	}

	return rfqEvent, nil
}

// Parse parses the rfq logs.
//
// nolint:gocognit,cyclop,dupl
func (p *RFQParser) Parse(ctx context.Context, log ethTypes.Log, chainID uint32) (interface{}, error) {
	rfqEvent, iFace, err := p.ParseLog(log, chainID)
	if err != nil {
		return nil, fmt.Errorf("could not parse rfq event: %w", err)
	}
	bridgeEventInterface, err := p.MatureLogs(ctx, rfqEvent, iFace, chainID)
	if err != nil {
		return nil, fmt.Errorf("could not mature rfq event: %w", err)
	}
	return bridgeEventInterface, nil
}

// applyPriceData applies price data to the rfq event, setting USD values.
func (p *RFQParser) applyPriceData(ctx context.Context, rfqEvent *model.RFQEvent, coinGeckoID string) {
	tokenPrice := p.tokenPriceService.GetPriceData(ctx, int(*rfqEvent.TimeStamp), coinGeckoID)
	if tokenPrice == nil {
		logger.Warnf("RFQ could not get token price for coingeckotoken; assuming price of 1:  %s txhash %s %d", coinGeckoID, rfqEvent.TxHash, rfqEvent.TimeStamp)
		one := 1.0
		tokenPrice = &one
	}

	if rfqEvent.OriginAmount != nil {
		amountUSD := GetAmountUSD(rfqEvent.OriginAmount, *rfqEvent.TokenDecimal, tokenPrice)
		if amountUSD != nil {
			rfqEvent.AmountUSD = *amountUSD
		}
	}

}

// eventToRFQEvent stores a message event.
func eventToRFQEvent(event rfqTypes.EventLog, chainID uint32) model.RFQEvent {
	transactionID := event.GetTransactionID()

	return model.RFQEvent{
		InsertTime:      uint64(time.Now().UnixNano()),
		ChainID:         chainID,
		TxHash:          event.GetTxHash().String(),
		ContractAddress: event.GetContractAddress().String(),
		BlockNumber:     event.GetBlockNumber(),
		EventType:       event.GetEventType().Int(),
		TransactionID:   common.Bytes2Hex(transactionID[:]),

		OriginToken:        event.GetOriginToken(),
		OriginAmount:       event.GetOriginAmount(),
		DestinationToken:   event.GetDestinationToken(),
		DestinationAmount:  event.GetDestinationAmount(),
		EventIndex:         event.GetEventIndex(),
		OriginChainID:      event.GetOriginChainID(),
		DestinationChainID: event.GetDestinationChainID(),
		Sender:             ToNullString(event.GetSender()),
		Relayer:            ToNullString(event.GetRelayer()),
		Request:            event.GetRequest(),
		Recipient:          ToNullString(event.GetRecipient()),
	}
}

func rfqEventToBridgeEvent(rfqEvent model.RFQEvent) model.BridgeEvent {
	bridgeType := bridgeTypes.BridgeRequestedEvent

	destinationKappa := rfqEvent.TransactionID
	var kappa *string
	if rfqEvent.EventType == rfqTypes.BridgeRelayedEvent.Int() {
		bridgeType = bridgeTypes.BridgeRelayedEvent
		destinationKappa = ""
		kappa = &rfqEvent.TransactionID
	}
	return model.BridgeEvent{
		InsertTime:       rfqEvent.InsertTime,
		ContractAddress:  rfqEvent.ContractAddress,
		ChainID:          rfqEvent.ChainID,
		EventType:        bridgeType.Int(),
		BlockNumber:      rfqEvent.BlockNumber,
		TxHash:           rfqEvent.TxHash,
		Token:            rfqEvent.OriginToken,
		Amount:           rfqEvent.OriginAmount,
		EventIndex:       rfqEvent.EventIndex,
		DestinationKappa: destinationKappa,
		Sender:           rfqEvent.Sender.String,

		Recipient:          rfqEvent.Recipient,
		RecipientBytes:     sql.NullString{},
		DestinationChainID: rfqEvent.DestinationChainID,
		Fee:                nil,
		Kappa:              ToNullString(kappa),
		TokenIndexFrom:     nil,
		TokenIndexTo:       nil,
		MinDy:              nil,
		Deadline:           nil,

		SwapSuccess:    nil,
		SwapTokenIndex: nil,
		SwapMinAmount:  nil,
		SwapDeadline:   nil,
		AmountUSD:      &rfqEvent.AmountUSD,
		FeeUSD:         nil,
		TokenDecimal:   rfqEvent.TokenDecimal,
		TokenSymbol:    ToNullString(&rfqEvent.TokenSymbol),
		TimeStamp:      rfqEvent.TimeStamp,
	}
}

func (p *RFQParser) storeBridgeEvent(ctx context.Context, bridgeEvent model.BridgeEvent) error {
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    300 * time.Second,
	}

	timeout := time.Duration(0)
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("%w while retrying store rfq converted bridge event", ctx.Err())
		case <-time.After(timeout):
			err := p.consumerDB.StoreEvent(ctx, &bridgeEvent)
			if err != nil {
				timeout = b.Duration()
				continue
			}
			return nil
		}
	}
}
