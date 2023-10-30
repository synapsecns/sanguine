// Package destination is the execution hub contract parser.
package destination

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"github.com/synapsecns/sanguine/services/sinner/db/model"
	"github.com/synapsecns/sanguine/services/sinner/logger"
	"github.com/synapsecns/sanguine/services/sinner/types"
)

// parserImpl is the parser for the origin contract.
type parserImpl struct {
	filterer *destination.DestinationFilterer
	// parser is the parser interface.
	parser destination.Parser
	// db is the database
	db db.EventDB
	// chainID is the chain ID
	chainID uint32
	// txMap is a map of tx hashes to tx data, exported for testing.
	txMap map[string]types.TxSupplementalInfo
}

// NewParser creates a new parser for the origin contract.
func NewParser(destinationAddress common.Address, db db.EventDB, chainID uint32) (types.EventParser, error) {
	// Get agents parser to utilize event type parsing.
	agentsParser, err := destination.NewParser(destinationAddress)
	if err != nil {
		return nil, fmt.Errorf("could not create agents parser: %w", err)
	}

	// Get filterer to get ABI IFace from abi.
	filter, err := destination.NewDestinationFilterer(destinationAddress, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create %T: %w", destination.DestinationFilterer{}, err)
	}

	parser := &parserImpl{
		filterer: filter,
		parser:   agentsParser,
		db:       db,
		chainID:  chainID,
	}
	return parser, nil
}

// UpdateTxMap updates the tx map so that scribe does not have to be requested for each log.
// This function is not concurrency safe, and is intended to be used before using ParseAndStore.
func (p *parserImpl) UpdateTxMap(txMap map[string]types.TxSupplementalInfo) {
	p.txMap = txMap
}

// UnsafeGetTXMap gets the tx map strictly for testing purposes.
func (p *parserImpl) UnsafeGetTXMap() map[string]types.TxSupplementalInfo {
	return p.txMap
}

// ParseAndStore parses and stores the log.
func (p *parserImpl) ParseAndStore(ctx context.Context, log ethTypes.Log) error {
	eventType, ok := p.parser.EventType(log)

	if !ok {
		logger.ReportSinnerError(fmt.Errorf("unknown execution hub log topic"), 0, logger.UnknownTopic)
	}
	if eventType == destination.ExecutedEvent {
		executedEvent, err := p.parseExecuted(log)
		if err != nil {
			return fmt.Errorf("error while parsing executed event. Err: %w", err)
		}

		g, storeCtx := errgroup.WithContext(ctx)
		g.Go(func() error {
			err := p.db.StoreOrUpdateMessageStatus(storeCtx, executedEvent.TxHash, executedEvent.MessageHash, types.Destination)
			if err != nil {
				return fmt.Errorf("error while storing executed event. Err: %w", err)
			}
			return nil
		})

		g.Go(func() error {
			err := p.db.StoreExecuted(storeCtx, executedEvent)
			if err != nil {
				return fmt.Errorf("error while storing executed event. Err: %w", err)
			}
			return nil
		})

		err = g.Wait()
		if err != nil {
			return fmt.Errorf("error while storing executed event. Err: %w", err)
		}
	}
	return nil
}

// ParseExecuted parses the sent event.
func (p *parserImpl) parseExecuted(log ethTypes.Log) (*model.Executed, error) {
	iFace, err := p.filterer.ParseExecuted(log)
	if err != nil {
		return nil, fmt.Errorf("could not parse executed log. err: %w", err)
	}
	parsedEvent := model.Executed{

		ContractAddress: iFace.Raw.Address.String(),
		BlockNumber:     iFace.Raw.BlockNumber,
		TxHash:          iFace.Raw.TxHash.String(),
		TxIndex:         iFace.Raw.TxIndex,
		MessageHash:     common.Bytes2Hex(iFace.MessageHash[:]),
		RemoteDomain:    iFace.RemoteDomain,
		Success:         iFace.Success,
		ChainID:         p.chainID,
		Timestamp:       uint64(p.txMap[iFace.Raw.TxHash.String()].Timestamp),
		Sender:          p.txMap[iFace.Raw.TxHash.String()].Sender,
	}
	return &parsedEvent, nil
}
