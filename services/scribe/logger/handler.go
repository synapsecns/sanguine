package logger

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-log"
	scribeTypes "github.com/synapsecns/sanguine/services/scribe/types"
)

var logger = log.Logger("scribe")

const (
	// ContextCancelled is returned when the context is canceled.
	ContextCancelled ErrorType = iota
	// LivefillIndexerError is returned when the livefill indexer encounters an error.
	LivefillIndexerError
	// BackfillIndexerError is returned when an indexer backfilling a contract to the head encounters an error.
	BackfillIndexerError
	// GetLogsError is returned when the logs cannot be retrieved.
	GetLogsError
	// GetTxError is returned when the tx cannot be retrieved.
	GetTxError
	// CouldNotGetReceiptError is returned when the receipt cannot be retrieved.
	CouldNotGetReceiptError
	// GetBlockError is returned when the block cannot be retrieved.
	GetBlockError
	// BlockByNumberError is returned when the block cannot be retrieved.
	BlockByNumberError
	// StoreError is returned when data cannot be inserted into the database.
	StoreError
	// ReadError is returned when data cannot be read from the database.
	ReadError
	// TestError is returned when an error during a test occurs.
	TestError
	// EmptyGetLogsChunk is returned when a getLogs chunk is empty.
	EmptyGetLogsChunk
	// FatalScribeError is for when something goes wrong with scribe.
	FatalScribeError
	// ErroneousHeadBlock is returned when the head block is below the last indexed.
	ErroneousHeadBlock
)

const (
	// InitiatingLivefill is returned when a contract backfills and is moving to livefill.
	InitiatingLivefill StatusType = iota
	// ConcurrencyThresholdReached is returned when the concurrency threshold is reached.
	ConcurrencyThresholdReached
	// FlushingLivefillAtHead is returned when a livefill indexer is flushing at the head.
	FlushingLivefillAtHead
	// CreatingSQLStore is returned when a SQL store is being created.
	CreatingSQLStore
	// BackfillCompleted is returned when a backfill is completed.
	BackfillCompleted
	// BeginBackfillIndexing is returned when a backfill is beginning.
	BeginBackfillIndexing
	// StoringLogs is returned when logs are being stored.
	StoringLogs
)

// ErrorType is a type of error.
type ErrorType int

// StatusType is a type of status for a process in scribe.
type StatusType int

// ReportIndexerError reports an error that occurs in an indexer.
//
// nolint
func ReportIndexerError(err error, indexerData scribeTypes.IndexerConfig, errorType ErrorType) {
	// nolint:exhaustive
	if err == nil {
		logger.Errorf("Error, @DEV: NIL ERROR\n%s", unpackIndexerConfig(indexerData))
		return
	}

	errStr := err.Error()

	// Stop cloudflare error messages from nuking readablity of logs
	if len(errStr) > 1000 {
		errStr = errStr[:1000]
	}
	switch errorType {
	case ContextCancelled:
		logger.Errorf("Context canceled for indexer. Error: %v\n%s", errStr, unpackIndexerConfig(indexerData))
	case LivefillIndexerError:
		logger.Errorf("Livefill indexer failed. Error: %v\n%s", errStr, unpackIndexerConfig(indexerData))
	case GetLogsError:
		logger.Errorf("Could not get logs. Error: %v\n%s", errStr, unpackIndexerConfig(indexerData))
	case GetTxError:
		logger.Errorf("Could not get tx. Error: %v\n%s", errStr, unpackIndexerConfig(indexerData))
	case CouldNotGetReceiptError:
		logger.Errorf("Could not get receipt. Error: %v\n%s", errStr, unpackIndexerConfig(indexerData))
	case GetBlockError:
		logger.Errorf("Could not get head block. Error: %v\n%s", errStr, unpackIndexerConfig(indexerData))
	case BlockByNumberError:
		logger.Errorf("Could not get block header. Error: %v\n%s", errStr, unpackIndexerConfig(indexerData))
	case StoreError:
		logger.Errorf("Could not store data into database. Error: %v\n%s", errStr, unpackIndexerConfig(indexerData))
	case ReadError:
		logger.Errorf("Could not read data from database. Error: %v\n%s", errStr, unpackIndexerConfig(indexerData))
	case EmptyGetLogsChunk:
		logger.Warnf("Encountered empty getlogs chunk%s", unpackIndexerConfig(indexerData))
	case ErroneousHeadBlock:
		logger.Warnf("Head block is below last indexed block%s", unpackIndexerConfig(indexerData))
	default:
		logger.Errorf("Error: %v\n%s", errStr, unpackIndexerConfig(indexerData))
	}
}

// ReportScribeError reports an error that occurs anywhere in scribe.
//
// nolint:exhaustive
func ReportScribeError(err error, chainID uint32, errorType ErrorType) {
	switch errorType {
	case ContextCancelled:
		logger.Errorf("Context canceled for scribe on chain %d. Error: %v", chainID, err)
	case GetBlockError:
		logger.Errorf("Could not get head block on chain %d. Error: %v", chainID, err)
	case TestError:
		logger.Errorf("Test error on chain %d. Error: %v", chainID, err)

	default:

		logger.Errorf("Error on chain %d: %v", chainID, err)
	}
}

// ReportScribeState reports a state that occurs anywhere in scribe.
func ReportScribeState(chainID uint32, block uint64, addresses []common.Address, statusType StatusType) {
	// nolint:exhaustive
	switch statusType {
	case InitiatingLivefill:
		logger.Warnf("Initiating livefill on chain %d on block %d while interacting with contract %s", chainID, block, dumpAddresses(addresses))
	case BackfillCompleted:
		logger.Warnf("Backfill completed on chain %d on block %d while interacting with contract %s", chainID, block, dumpAddresses(addresses))
	case BeginBackfillIndexing:
		logger.Warnf("Backfill beginning on chain %d on block %d while interacting with contract %s", chainID, block, dumpAddresses(addresses))
	case ConcurrencyThresholdReached:
		logger.Warnf("Concurrency threshold reached on chain %d on block %d while interacting with contract %s", chainID, block, dumpAddresses(addresses))
	case FlushingLivefillAtHead:
		logger.Warnf("Flushing logs at head on chain %d", chainID)
	case CreatingSQLStore:
		logger.Warnf("Creating SQL store")

	default:
		logger.Warnf("Event on chain %d on block %d while interacting with contract %s", chainID, block, dumpAddresses(addresses))
	}
}

// LogEventLogStore records when a log has been seen and will be stored. Used for debugging BSC missing txs.
func LogEventLogStore(chainID uint32, block uint64, txHash string) {
	logger.Warnf("Log seen: ChainID: %d | Block %d | TxHash %s", chainID, block, txHash)
}

func unpackIndexerConfig(indexerData scribeTypes.IndexerConfig) string {
	return fmt.Sprintf("Contracts: %v, GetLogsRange: %d, GetLogsBatchAmount: %d, StoreConcurrency: %d, ChainID: %d, StartHeight: %d, EndHeight: %d, ConcurrencyThreshold: %d",
		indexerData.Addresses, indexerData.GetLogsRange, indexerData.GetLogsBatchAmount, indexerData.StoreConcurrency, indexerData.ChainID, indexerData.StartHeight, indexerData.EndHeight, indexerData.ConcurrencyThreshold)
}

func dumpAddresses(addresses []common.Address) string {
	addressesStr := ""
	for i := range addresses {
		if i == len(addresses)-1 {
			addressesStr += addresses[i].String()
		} else {
			addressesStr += addresses[i].String() + ", "
		}
	}
	return addressesStr
}
