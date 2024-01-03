package parser

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser/tokendata"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher/tokenprice"
	"github.com/synapsecns/sanguine/services/explorer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/explorer/db"
	model "github.com/synapsecns/sanguine/services/explorer/db/sql"
	bridgeTypes "github.com/synapsecns/sanguine/services/explorer/types/bridge"
	cctpTypes "github.com/synapsecns/sanguine/services/explorer/types/cctp"
)

// CCTPParser parses cctp logs.
type CCTPParser struct {
	// consumerDB is the database to store parsed data in
	consumerDB db.ConsumerDB
	// Filterer is the message Filterer we use to parse events
	Filterer *cctp.SynapseCCTPFilterer
	// messageAddress is the address of the message
	cctpAddress common.Address
	// consumerFetcher is the Fetcher for sender and timestamp
	consumerFetcher fetcher.ScribeFetcher
	// cctpService is the cctp service for getting token symbol information
	cctpService fetcher.CCTPService
	// tokenDataService contains the token data service/cache
	tokenDataService tokendata.Service
	// tokenPriceService contains the token price service/cache
	tokenPriceService tokenprice.Service
	// fromAPI is true if the parser is being called from the API.
	fromAPI bool
}

const usdcCoinGeckoID = "usd-coin"
const usdcDecimals = 6

// NewCCTPParser creates a new parser for a cctp event.
func NewCCTPParser(consumerDB db.ConsumerDB, cctpAddress common.Address, consumerFetcher fetcher.ScribeFetcher, cctpService fetcher.CCTPService, tokenDataService tokendata.Service, tokenPriceService tokenprice.Service, fromAPI bool) (*CCTPParser, error) {
	filterer, err := cctp.NewSynapseCCTPFilterer(cctpAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", cctp.SynapseCCTPFilterer{}, err)
	}
	return &CCTPParser{consumerDB, filterer, cctpAddress, consumerFetcher, cctpService, tokenDataService, tokenPriceService, fromAPI}, nil
}

// ParserType returns the type of parser.
func (c *CCTPParser) ParserType() string {
	return "cctp"
}

// ParseLog log converts an eth log to a cctp event type.
func (c *CCTPParser) ParseLog(log ethTypes.Log, chainID uint32) (*model.CCTPEvent, cctpTypes.EventLog, error) {
	logTopic := log.Topics[0]
	iFace, err := func(log ethTypes.Log) (cctpTypes.EventLog, error) {
		switch logTopic {
		case cctp.Topic(cctpTypes.CircleRequestSentEvent):
			iFace, err := c.Filterer.ParseCircleRequestSent(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse circle request sent : %w", err)
			}
			return iFace, nil
		case cctp.Topic(cctpTypes.CircleRequestFulfilledEvent):
			iFace, err := c.Filterer.ParseCircleRequestFulfilled(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse circle request fulfilled: %w", err)
			}
			return iFace, nil

		default:
			logger.Warnf("ErrUnknownTopic in cctp: %s %s chain: %d address: %s", log.TxHash, logTopic.String(), chainID, log.Address.Hex())

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

	// Populate cctp event type so following operations can mature the event data.
	cctpEvent := eventToCCTPEvent(iFace, chainID)
	return &cctpEvent, iFace, nil
}

// MatureLogs takes a cctp event and adds data to them.
func (c *CCTPParser) MatureLogs(ctx context.Context, cctpEvent *model.CCTPEvent, iFace cctpTypes.EventLog, chainID uint32) (interface{}, error) {
	// Get timestamp from consumer
	timeStamp, err := c.consumerFetcher.FetchBlockTime(ctx, int(chainID), int(iFace.GetBlockNumber()))
	if err != nil {
		return nil, fmt.Errorf("could not get block time: %w", err)
	}

	// If we have a timestamp, populate the following attributes of cctpEvent.
	timeStampBig := uint64(*timeStamp)
	cctpEvent.TimeStamp = &timeStampBig

	tokenData, err := c.tokenDataService.GetCCTPTokenData(ctx, chainID, common.HexToAddress(cctpEvent.Token), c.cctpService)
	if err != nil {
		logger.Errorf("could not get token data: %v", err)
		return nil, fmt.Errorf("could not get pool token data: %w", err)
	}
	decimals := uint8(usdcDecimals)
	// Hotfix
	if chainID == 8453 &&(cctpEvent.Token == "0x417Ac0e078398C154EdFadD9Ef675d30Be60Af93" || cctpEvent.Token == "0x50c5725949A6F0c72E6C4a641F24049A917DB0Cb") {
		decimals = 18
	}
	cctpEvent.TokenSymbol = tokenData.TokenID()
	if (cctpEvent.Token == "0xDA10009cBd5D07dd0CeCc66161FC93D7c9000da1") {
		decimals = 18
		cctpEvent.TokenSymbol = "DAI"
	}
	if chainID == 10 && (cctpEvent.Token == "0x8c6f28f2F1A3C87F0f938b96d27520d9751ec8d9") {
		decimals = 18
		cctpEvent.TokenSymbol = "sUSD"
	}
	cctpEvent.TokenDecimal = &decimals
	c.applyPriceData(ctx, cctpEvent, usdcCoinGeckoID)

	// Would store into bridge database with a new goroutine but saw unreliable storage of events w/parent context cancellation.
	bridgeEvent := cctpEventToBridgeEvent(*cctpEvent)
	if c.fromAPI {
		return bridgeEvent, nil
	}
	err = c.storeBridgeEvent(ctx, bridgeEvent)
	if err != nil {
		logger.Errorf("could not store cctp event into bridge database: %v", err)
	}

	return cctpEvent, nil
}

// Parse parses the cctp logs.
//
// nolint:gocognit,cyclop,dupl
func (c *CCTPParser) Parse(ctx context.Context, log ethTypes.Log, chainID uint32) (interface{}, error) {
	cctpEvent, iFace, err := c.ParseLog(log, chainID)
	if err != nil {
		return nil, fmt.Errorf("could not parse cctp event: %w", err)
	}
	bridgeEventInterface, err := c.MatureLogs(ctx, cctpEvent, iFace, chainID)
	if err != nil {
		return nil, fmt.Errorf("could not mature cctp event: %w", err)
	}
	return bridgeEventInterface, nil
}

// applyPriceData applies price data to the cctp event, setting USD values.
func (c *CCTPParser) applyPriceData(ctx context.Context, cctpEvent *model.CCTPEvent, coinGeckoID string) {
	tokenPrice := c.tokenPriceService.GetPriceData(ctx, int(*cctpEvent.TimeStamp), coinGeckoID)
	if tokenPrice == nil {
		logger.Warnf("CCTP could not get token price for coingeckotoken; assuming price of 1:  %s txhash %s %d", coinGeckoID, cctpEvent.TxHash, cctpEvent.TimeStamp)
		one := 1.0
		tokenPrice = &one
	}

	if cctpEvent.Amount != nil {
		amountUSD := GetAmountUSD(cctpEvent.Amount, *cctpEvent.TokenDecimal, tokenPrice)
		if amountUSD != nil {
			cctpEvent.AmountUSD = *amountUSD
		}
	}
	if cctpEvent.Fee != nil {
		cctpEvent.FeeUSD = GetAmountUSD(cctpEvent.Fee, *cctpEvent.TokenDecimal, tokenPrice)
	}
}

// eventToCCTPEvent stores a message event.
func eventToCCTPEvent(event cctpTypes.EventLog, chainID uint32) model.CCTPEvent {
	requestID := event.GetRequestID()

	var formattedRequest sql.NullString
	if event.GetFormattedRequest() != nil {
		formattedRequest.Valid = true
		formattedRequest.String = common.Bytes2Hex(*event.GetFormattedRequest())
	} else {
		formattedRequest.Valid = false
	}

	return model.CCTPEvent{
		InsertTime:      uint64(time.Now().UnixNano()),
		ChainID:         chainID,
		TxHash:          event.GetTxHash().String(),
		ContractAddress: event.GetContractAddress().String(),
		BlockNumber:     event.GetBlockNumber(),
		EventType:       event.GetEventType().Int(),
		RequestID:       common.Bytes2Hex(requestID[:]),

		Token:              event.GetToken(),
		Amount:             event.GetAmount(),
		EventIndex:         event.GetEventIndex(),
		OriginChainID:      event.GetOriginChainID(),
		DestinationChainID: event.GetDestinationChainID(),
		Sender:             ToNullString(event.GetSender()),
		Nonce:              ToNullInt64(event.GetNonce()),
		MintToken:          ToNullString(event.GetMintToken()),
		RequestVersion:     ToNullInt32(event.GetRequestVersion()),
		FormattedRequest:   formattedRequest,
		Recipient:          ToNullString(event.GetRecipient()),
		Fee:                event.GetFee(),
	}
}

func cctpEventToBridgeEvent(cctpEvent model.CCTPEvent) model.BridgeEvent {
	bridgeType := bridgeTypes.CircleRequestSentEvent

	destinationKappa := cctpEvent.RequestID
	var kappa *string
	if cctpEvent.EventType == cctpTypes.CircleRequestFulfilledEvent.Int() {
		bridgeType = bridgeTypes.CircleRequestFulfilledEvent
		destinationKappa = ""
		kappa = &cctpEvent.RequestID
	}
	return model.BridgeEvent{
		InsertTime:       cctpEvent.InsertTime,
		ContractAddress:  cctpEvent.ContractAddress,
		ChainID:          cctpEvent.ChainID,
		EventType:        bridgeType.Int(),
		BlockNumber:      cctpEvent.BlockNumber,
		TxHash:           cctpEvent.TxHash,
		Token:            cctpEvent.Token,
		Amount:           cctpEvent.Amount,
		EventIndex:       cctpEvent.EventIndex,
		DestinationKappa: destinationKappa,
		Sender:           cctpEvent.Sender.String,

		Recipient:          cctpEvent.Recipient,
		RecipientBytes:     sql.NullString{},
		DestinationChainID: cctpEvent.DestinationChainID,
		Fee:                cctpEvent.Fee,
		Kappa:              ToNullString(kappa),
		TokenIndexFrom:     nil,
		TokenIndexTo:       nil,
		MinDy:              nil,
		Deadline:           nil,

		SwapSuccess:    nil,
		SwapTokenIndex: nil,
		SwapMinAmount:  nil,
		SwapDeadline:   nil,
		AmountUSD:      &cctpEvent.AmountUSD,
		FeeUSD:         cctpEvent.FeeUSD,
		TokenDecimal:   cctpEvent.TokenDecimal,
		TokenSymbol:    ToNullString(&cctpEvent.TokenSymbol),
		TimeStamp:      cctpEvent.TimeStamp,
	}
}

func (c *CCTPParser) storeBridgeEvent(ctx context.Context, bridgeEvent model.BridgeEvent) error {
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
			return fmt.Errorf("%w while retrying store cctp converted bridge event", ctx.Err())
		case <-time.After(timeout):
			err := c.consumerDB.StoreEvent(ctx, &bridgeEvent)
			if err != nil {
				timeout = b.Duration()
				continue
			}
			return nil
		}
	}
}
