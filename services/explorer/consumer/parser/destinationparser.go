package parser

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/contracts/messaging/destination"
	"github.com/synapsecns/sanguine/services/explorer/db"
	model "github.com/synapsecns/sanguine/services/explorer/db/sql"
	destinationTypes "github.com/synapsecns/sanguine/services/explorer/types/destination"
	"golang.org/x/sync/errgroup"
)

// DestinationParser parses events from the destination contract.
type DestinationParser struct {
	// consumerDB is the database to store parsed data in.
	consumerDB db.ConsumerDB
	// Filterer is the bridge Filterer we use to parse events.
	Filterer *destination.DestinationFilterer
	// destinationAddress is the address of the destination contract.
	destinationAddress common.Address
	// consumerFetcher is the ScribeFetcher for sender and timestamp.
	consumerFetcher *fetcher.ScribeFetcher
}

// NewDestinationParser creates a new parser for a given Destination Contract.
func NewDestinationParser(consumerDB db.ConsumerDB, destinationAddress common.Address, consumerFetcher *fetcher.ScribeFetcher) (*DestinationParser, error) {

	filterer, err := destination.NewDestinationFilterer(destinationAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", destination.DestinationFilterer{}, err)
	}

	return &DestinationParser{
		consumerDB:         consumerDB,
		Filterer:           filterer,
		destinationAddress: destinationAddress,
		consumerFetcher:    consumerFetcher,
	}, nil
}

// EventType returns the event type of a destination log.
func (p *DestinationParser) EventType(log ethTypes.Log) (_ destinationTypes.EventType, ok bool) {
	for _, logTopic := range log.Topics {
		eventType := destination.EventTypeFromTopic(logTopic)
		if eventType == nil {
			continue
		}

		return *eventType, true
	}

	// Return an unknown event to avoid cases where user failed to check the event type.
	return destinationTypes.EventType(len(destinationTypes.AllEventTypes())), false
}

// eventToDestinationEvent stores a destination event.
func eventToDestinationEvent(event destinationTypes.EventLog, chainID uint32) model.DestinationEvent {
	var agent sql.NullString

	if event.GetAgent() != nil {
		agent.Valid = true
		agent.String = event.GetAgent().String()
	} else {
		agent.Valid = false
	}
	return model.DestinationEvent{
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
		Attestation:     event.GetAttestation(),
		AttSignature:    event.GetAttSignature(),
		AgentRoot:       event.GetAgentRoot(),
	}
}

// Parse parses the bridge logs and returns a model that can be stored
//
// nolint:gocognit,cyclop,dupl,maintidx
func (p *DestinationParser) Parse(ctx context.Context, log ethTypes.Log, chainID uint32) (interface{}, error) {
	logTopic := log.Topics[0]

	iFace, err := func(log ethTypes.Log) (destinationTypes.EventLog, error) {
		switch logTopic {
		case destination.Topic(destinationTypes.AttestationAcceptedEvent):
			iFace, err := p.Filterer.ParseAttestationAccepted(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse AttestationAcceptedEvent: %w", err)
			}

			return iFace, nil
		case destination.Topic(destinationTypes.AgentRootAcceptedEvent):
			iFace, err := p.Filterer.ParseAgentRootAccepted(log)
			if err != nil {
				return nil, fmt.Errorf("could not parse AgentRootAcceptedEvent: %w", err)
			}

			return iFace, nil
		default:
			logger.Warnf("ErrUnknownTopic in destination: %s %s chain: %d address: %s", log.TxHash, logTopic.String(), chainID, log.Address.Hex())

			return nil, fmt.Errorf(ErrUnknownTopic)
		}
	}(log)

	if err != nil {
		// Switch failed.
		return nil, err
	}

	destinationEvent := eventToDestinationEvent(iFace, chainID)
	g, groupCtx := errgroup.WithContext(ctx)

	var sender *string
	var timeStamp *uint64
	g.Go(func() error {
		timeStamp, sender, err = p.consumerFetcher.FetchTx(groupCtx, iFace.GetTxHash().String(), int(chainID), int(destinationEvent.BlockNumber))
		if err != nil {
			return fmt.Errorf("could not get timestamp, sender on chain %d and tx %s from tx %w", chainID, iFace.GetTxHash().String(), err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could not parse destination event: %w", err)
	}
	if *timeStamp == 0 {
		logger.Errorf("empty block time: chain: %d address %s", chainID, log.Address.Hex())
		return nil, fmt.Errorf("empty block time: chain: %d address %s", chainID, log.Address.Hex())
	}

	destinationEvent.TimeStamp = timeStamp
	destinationEvent.Sender = *sender

	return destinationEvent, nil
}
