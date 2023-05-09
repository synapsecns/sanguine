package parser

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/contracts/messaging/summit"
	"github.com/synapsecns/sanguine/services/explorer/db"
	model "github.com/synapsecns/sanguine/services/explorer/db/sql"
	summitTypes "github.com/synapsecns/sanguine/services/explorer/types/summit"
	"golang.org/x/sync/errgroup"
)

// SummitParser parses events from the summit contract.
type SummitParser struct {
	// consumerDB is the database to store parsed data in.
	consumerDB db.ConsumerDB
	// Filterer is the bridge Filterer we use to parse events.
	Filterer *summit.SummitFilterer
	// summitAddress is the address of the summit contract.
	summitAddress common.Address
	// consumerFetcher is the ScribeFetcher for sender and timestamp.
	consumerFetcher *fetcher.ScribeFetcher
}

// NewSummitParser creates a new parser for a given Summit Contract.
func NewSummitParser(consumerDB db.ConsumerDB, summitAddress common.Address, consumerFetcher *fetcher.ScribeFetcher) (*SummitParser, error) {

	filterer, err := summit.NewSummitFilterer(summitAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", summit.SummitFilterer{}, err)
	}

	return &SummitParser{
		consumerDB:      consumerDB,
		Filterer:        filterer,
		summitAddress:   summitAddress,
		consumerFetcher: consumerFetcher,
	}, nil
}

// EventType returns the event type of a summit log.
func (p *SummitParser) EventType(log ethTypes.Log) (_ summitTypes.EventType, ok bool) {
	for _, logTopic := range log.Topics {
		eventType := summit.EventTypeFromTopic(logTopic)
		if eventType == nil {
			continue
		}

		return *eventType, true
	}

	// Return an unknown event to avoid cases where user failed to check the event type.
	return summitTypes.EventType(len(summitTypes.AllEventTypes())), false
}

// eventToSummitEvent stores a summit event.
func eventToSummitEvent(event summitTypes.EventLog, chainID uint32) model.SummitEvent {
	var agent sql.NullString

	if event.GetAgent() != nil {
		agent.Valid = true
		agent.String = event.GetAgent().String()
	} else {
		agent.Valid = false
	}
	return model.SummitEvent{
		InsertTime:      uint64(time.Now().UnixNano()),
		ContractAddress: event.GetContractAddress().String(),
		ChainID:         chainID,
		EventType:       event.GetEventType().Int(),
		BlockNumber:     event.GetBlockNumber(),
		TxHash:          event.GetTxHash().String(),
		EventIndex:      event.GetEventIndex(),
		Sender:          "",
		Domain:          event.GetDomain(),
		Agent:           agent,
		RcptPayload:     event.GetRcptPayload(),
		RcptSignature:   event.GetRcptSignature(),
		Snapshot:        event.GetSnapshot(),
		SnapSignature:   event.GetSnapSignature(),
		Tip:             event.GetTip(),
	}
}

// Parse parses the bridge logs and returns a model that can be stored
//
// nolint:gocognit,cyclop,dupl,maintidx
func (p *SummitParser) Parse(ctx context.Context, log ethTypes.Log, chainID uint32) (interface{}, error) {
	logTopic := log.Topics[0]

	iFace, err := func(log ethTypes.Log) (summitTypes.EventLog, error) {
		switch logTopic {
		case summit.Topic(summitTypes.ReceiptAcceptedEvent):
			iFace, err := p.Filterer.ParseReceiptAccepted(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse ReceiptAcceptedEvent: %w", err)
			}

			return iFace, nil
		case summit.Topic(summitTypes.SnapshotAcceptedEvent):
			iFace, err := p.Filterer.ParseSnapshotAccepted(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse SnapshotAcceptedEvent: %w", err)
			}

			return iFace, nil
		case summit.Topic(summitTypes.TipAwardedEvent):
			iFace, err := p.Filterer.ParseTipAwarded(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse TipAwardedEvent: %w", err)
			}

			return iFace, nil
		default:
			logger.Warnf("ErrUnknownTopic in summit: %s %s chain: %d address: %s", log.TxHash, logTopic.String(), chainID, log.Address.Hex())

			return nil, fmt.Errorf(ErrUnknownTopic)
		}
	}(log)

	if err != nil {
		// Switch failed.
		return nil, err
	}

	summitEvent := eventToSummitEvent(iFace, chainID)
	g, groupCtx := errgroup.WithContext(ctx)

	var sender *string
	var timeStamp *uint64
	g.Go(func() error {
		timeStamp, sender, err = p.consumerFetcher.FetchTx(groupCtx, iFace.GetTxHash().String(), int(chainID), int(summitEvent.BlockNumber))
		if err != nil {
			return fmt.Errorf("could not get timestamp, sender on chain %d and tx %s from tx %w", chainID, iFace.GetTxHash().String(), err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could not parse summit event: %w", err)
	}
	if *timeStamp == 0 {
		logger.Errorf("empty block time: chain: %d address %s", chainID, log.Address.Hex())
		return nil, fmt.Errorf("empty block time: chain: %d address %s", chainID, log.Address.Hex())
	}

	summitEvent.TimeStamp = timeStamp
	summitEvent.Sender = *sender

	return summitEvent, nil
}
