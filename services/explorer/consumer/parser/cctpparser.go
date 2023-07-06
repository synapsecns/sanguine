package parser

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher/tokenprice"
	"github.com/synapsecns/sanguine/services/explorer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/explorer/db"
	model "github.com/synapsecns/sanguine/services/explorer/db/sql"
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
	consumerFetcher *fetcher.ScribeFetcher
	// tokenPriceService contains the token price service/cache
	tokenPriceService tokenprice.Service
}

const usdcCoinGeckoID = "usd-coin"
const usdcDecimals = 6

// NewCCTPParser creates a new parser for a cctp event.
func NewCCTPParser(consumerDB db.ConsumerDB, cctpAddress common.Address, consumerFetcher *fetcher.ScribeFetcher, tokenPriceService tokenprice.Service) (*CCTPParser, error) {
	filterer, err := cctp.NewSynapseCCTPFilterer(cctpAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", cctp.SynapseCCTPFilterer{}, err)
	}
	return &CCTPParser{consumerDB, filterer, cctpAddress, consumerFetcher, tokenPriceService}, nil
}

// Parse parses the cctp logs.
//
// nolint:gocognit,cyclop,dupl
func (c *CCTPParser) Parse(ctx context.Context, log ethTypes.Log, chainID uint32) (interface{}, error) {
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

		return nil, err
	}
	if iFace == nil {
		// Unknown topic.
		return nil, fmt.Errorf("unknwn topic")
	}

	// Populate cctp event type so following operations can mature the event data.
	cctpEvent := eventToCCTPEvent(iFace)

	// Get timestamp from consumer
	timeStamp, err := c.consumerFetcher.FetchBlockTime(ctx, int(chainID), int(iFace.GetBlockNumber()))
	if err != nil {
		return nil, fmt.Errorf("could not get block time: %w", err)
	}

	// If we have a timestamp, populate the following attributes of cctpEvent.
	timeStampBig := uint64(*timeStamp)
	cctpEvent.TimeStamp = &timeStampBig

	err = c.applyPriceData(ctx, &cctpEvent, usdcCoinGeckoID)
	if err != nil {
		return nil, fmt.Errorf("could not apply price data: %w", err)
	}

	return cctpEvent, nil
}

// applyPriceData applies price data to the cctp event, setting USD values.
func (c *CCTPParser) applyPriceData(ctx context.Context, cctpEvent *model.CCTPEvent, coinGeckoID string) error {
	tokenPrice := c.tokenPriceService.GetPriceData(ctx, int(*cctpEvent.TimeStamp), coinGeckoID)
	if (tokenPrice == nil) && coinGeckoID != noTokenID && coinGeckoID != noPrice {
		return fmt.Errorf("CCTP could not get token price for coingeckotoken:  %s txhash %s %d", coinGeckoID, cctpEvent.TxHash, cctpEvent.TimeStamp)
	}

	cctpEvent.SentAmountUSD = GetAmountUSD(cctpEvent.SentAmount, usdcDecimals, tokenPrice)
	cctpEvent.FeeUSD = GetAmountUSD(cctpEvent.Fee, usdcDecimals, tokenPrice)
	return nil
}

// eventToCCTPEvent stores a message event.
func eventToCCTPEvent(event cctpTypes.EventLog) model.CCTPEvent {
	return model.CCTPEvent{
		// TODO add event type to implementation of event log
		InsertTime:         uint64(time.Now().UnixNano()),
		TxHash:             event.GetTxHash().String(),
		ContractAddress:    event.GetContractAddress().String(),
		BlockNumber:        event.GetBlockNumber(),
		OriginChainID:      event.GetOriginChainID(),
		DestinationChainID: event.GetDestinationChainID(),
		Sender:             ToNullString(event.GetSender()),
		Nonce:              ToNullInt64(event.GetNonce()),
		BurnToken:          ToNullString(event.GetBurnToken()),
		MintToken:          ToNullString(event.GetMintToken()),
		SentAmount:         event.GetSentAmount(),
		ReceivedAmount:     event.GetReceivedAmount(),
		RequestVersion:     ToNullInt32(event.GetRequestVersion()),
		FormattedRequest:   event.GetFormattedRequest(),
		RequestID:          event.GetRequestID(),
		Recipient:          ToNullString(event.GetRecipient()),
		Fee:                event.GetFee(),
		Token:              ToNullString(event.GetToken()),
	}
}
