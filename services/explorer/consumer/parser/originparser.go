package parser

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/contracts/messaging/origin"
	"github.com/synapsecns/sanguine/services/explorer/db"
	model "github.com/synapsecns/sanguine/services/explorer/db/sql"
	originTypes "github.com/synapsecns/sanguine/services/explorer/types/origin"
	"golang.org/x/sync/errgroup"
	"time"
)

// OriginParser parses events from the origin contract.
type OriginParser struct {
	// consumerDB is the database to store parsed data in.
	consumerDB db.ConsumerDB
	// Filterer is the bridge Filterer we use to parse events.
	Filterer *origin.OriginFilterer
	// originAddress is the address of the origin contract.
	originAddress common.Address
	// consumerFetcher is the ScribeFetcher for sender and timestamp.
	consumerFetcher *fetcher.ScribeFetcher
}

// NewOriginParser creates a new parser for a given Origin Contract.
func NewOriginParser(consumerDB db.ConsumerDB, originAddress common.Address, consumerFetcher *fetcher.ScribeFetcher) (*OriginParser, error) {

	filterer, err := origin.NewOriginFilterer(originAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", origin.OriginFilterer{}, err)
	}

	return &OriginParser{
		consumerDB:      consumerDB,
		Filterer:        filterer,
		originAddress:   originAddress,
		consumerFetcher: consumerFetcher,
	}, nil
}

// EventType returns the event type of a origin log.
func (p *OriginParser) EventType(log ethTypes.Log) (_ originTypes.EventType, ok bool) {
	for _, logTopic := range log.Topics {
		eventType := origin.EventTypeFromTopic(logTopic)
		if eventType == nil {
			continue
		}

		return *eventType, true
	}

	// Return an unknown event to avoid cases where user failed to check the event type.
	return originTypes.EventType(len(originTypes.AllEventTypes())), false
}

// eventToOriginEvent stores a origin event.
func eventToOriginEvent(event originTypes.EventLog, chainID uint32) model.OriginEvent {
	return model.OriginEvent{
		InsertTime:      uint64(time.Now().UnixNano()),
		ContractAddress: event.GetContractAddress().String(),
		ChainID:         chainID,
		EventType:       event.GetEventType().Int(),
		BlockNumber:     event.GetBlockNumber(),
		TxHash:          event.GetTxHash().String(),
		EventIndex:      event.GetEventIndex(),
		Sender:          "",
		MessageHash:     event.GetMessageHash(),
		Nonce:           event.GetNonce(),
		Destination:     event.GetDestination(),
		Message:         event.GetMessage(),
	}
}

// Parse parses the bridge logs and returns a model that can be stored
//
// nolint:gocognit,cyclop,dupl,maintidx
func (p *OriginParser) Parse(ctx context.Context, log ethTypes.Log, chainID uint32) (interface{}, error) {
	logTopic := log.Topics[0]

	iFace, err := func(log ethTypes.Log) (originTypes.EventLog, error) {
		switch logTopic {
		case origin.Topic(originTypes.SentEvent):
			iFace, err := p.Filterer.ParseSent(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse sent event: %w", err)
			}

			return iFace, nil
		default:
			logger.Warnf("ErrUnknownTopic in origin: %s %s chain: %d address: %s", log.TxHash, logTopic.String(), chainID, log.Address.Hex())

			return nil, fmt.Errorf(ErrUnknownTopic)
		}
	}(log)

	if err != nil {
		// Switch failed.
		return nil, err
	}

	originEvent := eventToOriginEvent(iFace, chainID)
	g, groupCtx := errgroup.WithContext(ctx)

	var sender *string
	var timeStamp *uint64
	g.Go(func() error {
		timeStamp, sender, err = p.consumerFetcher.FetchTx(groupCtx, iFace.GetTxHash().String(), int(chainID), int(originEvent.BlockNumber))
		if err != nil {
			return fmt.Errorf("could not get timestamp, sender on chain %d and tx %s from tx %w", chainID, iFace.GetTxHash().String(), err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could not parse origin event: %w", err)
	}
	if *timeStamp == 0 {
		logger.Errorf("empty block time: chain: %d address %s", chainID, log.Address.Hex())
		return nil, fmt.Errorf("empty block time: chain: %d address %s", chainID, log.Address.Hex())
	}

	originEvent.TimeStamp = timeStamp
	originEvent.Sender = *sender

	return originEvent, nil
}
