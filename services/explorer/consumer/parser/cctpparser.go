package parser

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher/tokenprice"
	"github.com/synapsecns/sanguine/services/explorer/contracts/cctp"
	"github.com/synapsecns/sanguine/services/explorer/db"
	cctpTypes "github.com/synapsecns/sanguine/services/explorer/types/cctp"
	messageBusTypes "github.com/synapsecns/sanguine/services/explorer/types/messagebus"
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
	iFace, err := func(log ethTypes.Log) (messageBusTypes.EventLog, error) {
		switch logTopic {
		case cctp.Topic(cctpTypes.CircleRequestSentEvent):
			iFace, err := m.Filterer.ParseCircleRequestSent(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse token : %w", err)
			}
			return iFace, nil
		case cctp.Topic(cctpTypes.CircleRequestFulfilledEvent):
			iFace, err := m.Filterer.ParseCircleRequestFulfilled(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse sent message: %w", err)
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

	// populate cctp event type so following operations can mature the event data.

	// TODO make eventToCCTPEvent. This function pulls all log data (using the EventLog interface) into the CCTPEvent type for database insertion
	messageEvent := eventToCCTPEvent(iFace, chainID)

	// Get timestamp from consumer
	timeStamp, err := m.consumerFetcher.FetchBlockTime(ctx, int(chainID), int(iFace.GetBlockNumber()))
	if err != nil {
		return nil, fmt.Errorf("could not get block time: %w", err)
	}

	// If we have a timestamp, populate the following attributes of cctpEvent.
	timeStampBig := uint64(*timeStamp)
	cctpEvent.TimeStamp = &timeStampBig

	coinGeckoID := "usd-coin"
	usdcTokenPrice := c.tokenPriceService.GetPriceData(ctx, int(timeStampBig), coinGeckoID)
	if (usdcTokenPrice == nil) && coinGeckoID != noTokenID && coinGeckoID != noPrice {
		// TODO exit
	}

	// TODO get usd values for amount and fee (Add conditional for only when these values exist, aka on destination)
	feePrice := GetAmountUSD(messageEvent.Fee, 18, tokenPrice)
	amountPrice := GetAmountUSD(messageEvent.Fee, 18, tokenPrice)

	messageEvent.FeeUSD = feeValue

	return messageEvent, nil
}
