package logger

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-log"
	scribeTypes "github.com/synapsecns/sanguine/services/scribe/types"
)

var logger = log.Logger("scribe")

const (
	// ContextCancelled is returned when the context is cancelled
	ContextCancelled ErrorType = iota
	// LivefillIndexerError is returned when the livefill indexer encounters an error.
	LivefillIndexerError
	// BackfillIndexerError is returned when an indexer backfilling a contract to the head encounters an error.
	BackfillIndexerError
	// GetLogsError is returned when the logs cannot be retrieved
	GetLogsError
	// GetTxError is returned when the tx cannot be retrieved
	GetTxError
	// CouldNotGetReceiptError is returned when the receipt cannot be retrieved
	CouldNotGetReceiptError
	// GetBlockError is returned when the block cannot be retrieved
	GetBlockError
	// BlockByNumberError is returned when the block cannot be retrieved
	BlockByNumberError
	// StoreError is returned when data cannot be inserted into the database
	StoreError
	// ReadError is returned when data cannot be read from the database
	ReadError
)

const (
	// InitiatingLivefill is returned when a contract backfills and is moving to livefill.
	InitiatingLivefill StatusType = iota
)

type ErrorType int
type StatusType int

func ReportIndexerError(err error, indexerData scribeTypes.IndexerConfig, errorType ErrorType) {
	switch errorType {
	case ContextCancelled:
		logger.Errorf("Context cancelled for indexer. Error: %v\n%v", err, indexerData)
	case LivefillIndexerError:
		logger.Errorf("Livefill indexer failed. Error: %v\n%v", err, indexerData)
	case GetLogsError:
		logger.Errorf("Could not get logs. Error: %v\n%v", err, indexerData)
	case GetTxError:
		logger.Errorf("Could not get tx. Error: %v\n%v", err, indexerData)
	case CouldNotGetReceiptError:
		logger.Errorf("Could not get receipt. Error: %v\n%v", err, indexerData)
	case GetBlockError:
		logger.Errorf("Could not get head block. Error: %v\n%v", err, indexerData)
	case BlockByNumberError:
		logger.Errorf("Could not get block header. Error: %v\n%v", err, indexerData)
	case StoreError:
		logger.Errorf("Could not store data into database. Error: %v\n%v", err, indexerData)
	case ReadError:
		logger.Errorf("Could not read data from database. Error: %v\n%v", err, indexerData)
	default:
		logger.Errorf("Error: %v\n%v", err, indexerData)
	}
}

func ReportScribeError(err error, chainID uint32, errorType ErrorType) {
	switch errorType {
	case ContextCancelled:
		logger.Errorf("Context cancelled for scribe on chain %d. Error: %v", chainID, err)
	case GetBlockError:
		logger.Errorf("Could not get head block on chain %d. Error: %v", chainID, err)
	default:
		logger.Errorf("Error on chain %d: %v", chainID, err)
	}
}

func ReportScribeState(chainID uint32, block uint64, contract common.Address, statusType StatusType) {
	switch statusType {
	case InitiatingLivefill:
		logger.Warnf("Initiating livefill on chain %d on block %d while interacting with contract %s", chainID, block, contract.String())
	default:
		logger.Warnf("Event on chain %d on block %d while interacting with contract %s", chainID, block, contract.String())
	}
}
