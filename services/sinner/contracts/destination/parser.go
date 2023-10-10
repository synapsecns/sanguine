package destination

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"github.com/synapsecns/sanguine/services/sinner/db/model"
	"github.com/synapsecns/sanguine/services/sinner/logger"
	"github.com/synapsecns/sanguine/services/sinner/types"
)

type ParserImpl struct {
	filterer *destination.DestinationFilterer
	// parser is the parser interface.
	parser destination.Parser
	// db is the database
	db db.EventDB
	// chainID is the chain ID
	chainID uint32
	// txMap is a map of tx hashes to tx data
	txMap map[string]types.TxSupplementalInfo
}

// NewParser creates a new parser for the origin contract.
func NewParser(destinationAddress common.Address, db db.EventDB, chainID uint32) (*ParserImpl, error) {
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

	parser := &ParserImpl{
		filterer: filter,
		parser:   agentsParser,
		db:       db,
		chainID:  chainID,
	}
	return parser, nil
}

func (p ParserImpl) UpdateTxMap(txMap map[string]types.TxSupplementalInfo) {
	p.txMap = txMap
}
func (p ParserImpl) ParseAndStore(ctx context.Context, log ethTypes.Log) error {
	eventType, ok := p.parser.EventType(log)
	if !ok {
		logger.ReportSinnerError(fmt.Errorf("unknown destination log topic"), 0, logger.UnknownTopic)
	}
	switch eventType {
	case destination.ExecutedEvent:
		executedEvent, err := p.ParseExecuted(log)
		if err != nil {
			return fmt.Errorf("error while parsing origin sent event. Err: %w", err)
		}

		// TODO go func this
		err = p.db.StoreOrUpdateMessageStatus(ctx, executedEvent.TxHash, executedEvent.MessageHash, types.Destination)
		if err != nil {
			return fmt.Errorf("error while storing origin sent event. Err: %w", err)
		}

		err = p.db.StoreExecuted(ctx, executedEvent)
		if err != nil {
			return fmt.Errorf("error while storing origin sent event. Err: %w", err)
		}

	}
	return nil
}

func (p ParserImpl) ParseExecuted(log ethTypes.Log) (*model.Executed, error) {

	iFace, err := p.filterer.ParseExecuted(log)
	if err != nil {
		return nil, fmt.Errorf("could not parse sent log. err: %w", err)
	}
	parsedEvent := model.Executed{

		ContractAddress: iFace.Raw.Address.String(),
		BlockNumber:     iFace.Raw.BlockNumber,
		TxHash:          iFace.Raw.TxHash.String(),
		TxIndex:         iFace.Raw.TxIndex,
		MessageHash:     string(iFace.MessageHash[:]),
		RemoteDomain:    iFace.RemoteDomain,
		Success:         iFace.Success,
		ChainID:         p.chainID,
	}
	return &parsedEvent, nil
}
