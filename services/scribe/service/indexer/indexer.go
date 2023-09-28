package indexer

import (
	"context"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	scribeTypes "github.com/synapsecns/sanguine/services/scribe/types"

	"github.com/synapsecns/sanguine/services/scribe/logger"
	"math/big"
	"time"

	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/core/mapmutex"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	otelMetrics "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	lru "github.com/hashicorp/golang-lru"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"golang.org/x/sync/errgroup"
)

// Indexer is a backfiller that fetches logs for a specific contract.
type Indexer struct {
	// indexerConfig holds all the metadata needed for logging and indexing.
	indexerConfig scribeTypes.IndexerConfig
	// eventDB is the database to store event data in.
	eventDB db.EventDB
	// client is the client for filtering.
	client []backend.ScribeBackend
	// cache is a cache for txHashes.
	cache *lru.Cache
	// mux is the mutex used to prevent double inserting logs from the same tx
	mux mapmutex.StringerMapMutex
	// handler is the metrics handler for the scribe.
	handler metrics.Handler
	// blockMeter is an otel historgram for doing metrics on block heights by chain
	blockMeter otelMetrics.Int64Histogram
	// refreshRate is the rate at which the indexer will refresh when livefilling.
	refreshRate uint64
	// toHead is a boolean signifying if the indexer is livefilling to the head.
	toHead bool
	// isBackfill is a boolean signifying if the indexer is backfilling (prevents last indexed from running)
	isBackfill bool
}

// retryTolerance is the number of times to retry a failed operation before rerunning the entire Backfill function.
const retryTolerance = 20

// txNotSupportedError is for handling the legacy Arbitrum tx type.
const txNotSupportedError = "transaction type not supported"

// invalidTxVRSError is for handling Aurora VRS error.
const invalidTxVRSError = "invalid transaction v, r, s values"

// txNotFoundError is for handling omniRPC errors for BSx.
const txNotFoundError = "not found"

// txData returns the transaction data for a given transaction hash.
type txData struct {
	receipt     types.Receipt
	transaction types.Transaction
	blockHeader types.Header
	success     bool
}

// errNoContinue indicates an error that is not recoverable, and should not be retried.
var errNoContinue = errors.New("encountered unreconcilable error, will not attempt to store tx")

// errNoTx indicates a tx cannot be parsed, this is only returned when the tx doesn't match our data model.
var errNoTx = errors.New("tx is not supported by the client")

// NewIndexer creates a new backfiller for a contract.
func NewIndexer(chainConfig config.ChainConfig, addresses []common.Address, eventDB db.EventDB, client []backend.ScribeBackend, handler metrics.Handler, blockMeter otelMetrics.Int64Histogram, toHead bool) (*Indexer, error) {
	cache, err := lru.New(500)
	if err != nil {
		return nil, fmt.Errorf("could not initialize cache: %w", err)
	}

	refreshRate := uint64(1)
	if len(addresses) > 1 || len(addresses) == 0 { // livefill settings
		chainConfig.GetLogsRange = chainConfig.LivefillRange
		chainConfig.GetLogsBatchAmount = 1
		chainConfig.StoreConcurrency = 1
		chainConfig.ConcurrencyThreshold = 10000
	} else {
		for i := range chainConfig.Contracts { // get the refresh rate for the contract
			contract := chainConfig.Contracts[i]
			// Refresh rate for more than one contract is 1 second, the refresh rate set in the config is used when it is the only contract.
			if contract.Address == addresses[0].String() && contract.RefreshRate > 0 {
				refreshRate = contract.RefreshRate
				break
			}
		}
	}

	indexerConfig := scribeTypes.IndexerConfig{
		Addresses:            addresses,
		GetLogsRange:         chainConfig.GetLogsRange,
		GetLogsBatchAmount:   chainConfig.GetLogsBatchAmount,
		StoreConcurrency:     chainConfig.StoreConcurrency,
		ChainID:              chainConfig.ChainID,
		ConcurrencyThreshold: chainConfig.ConcurrencyThreshold,
	}

	return &Indexer{
		indexerConfig: indexerConfig,
		eventDB:       eventDB,
		client:        client,
		cache:         cache,
		mux:           mapmutex.NewStringerMapMutex(),
		handler:       handler,
		blockMeter:    blockMeter,
		refreshRate:   refreshRate,
		toHead:        toHead,
		isBackfill:    false,
	}, nil
}

// UpdateAddress updates the address arrays for the indexer.
func (x *Indexer) UpdateAddress(addresses []common.Address) {
	x.indexerConfig.Addresses = addresses
}

// SetToBackfill sets the indexer to backfill (will not update last indexed).
func (x *Indexer) SetToBackfill() {
	x.isBackfill = true
}

// GetIndexerConfig returns the indexer config.
func (x *Indexer) GetIndexerConfig() scribeTypes.IndexerConfig {
	return x.indexerConfig
}

// RefreshRate returns the refresh rate for the indexer.
func (x *Indexer) RefreshRate() uint64 {
	return x.refreshRate
}

// Index retrieves logs, receipts, and transactions for a contract from a given range and does so in the following manner.
// 1. Get logs for the contract in chunks of batch requests.
// 2. Iterate through each log's Tx Hash and performs the following
//   - Get the receipt for each log and store it and all of its logs.
//   - Get the transaction for each log and store it.
//
//nolint:gocognit, cyclop
func (x *Indexer) Index(parentCtx context.Context, startHeight uint64, endHeight uint64) (err error) {
	ctx, span := x.handler.Tracer().Start(parentCtx, "contract.Backfill", trace.WithAttributes(
		attribute.Int("chain", int(x.indexerConfig.ChainID)),
		attribute.String("address", x.addressesToString(x.indexerConfig.Addresses)),
		attribute.Int("start", int(startHeight)),
		attribute.Int("end", int(endHeight)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	g, groupCtx := errgroup.WithContext(ctx)

	// For logging
	x.indexerConfig.StartHeight = startHeight
	x.indexerConfig.EndHeight = endHeight

	// Start fetching logs
	logFetcher := NewLogFetcher(x.client[0], big.NewInt(int64(startHeight)), big.NewInt(int64(endHeight)), &x.indexerConfig, true)
	logsChan := *logFetcher.GetFetchedLogsChan()
	g.Go(func() error {
		return logFetcher.Start(groupCtx)
	})
	// Reads from the local logsChan and stores the logs and associated receipts / txs.
	g.Go(func() error {
		concurrentCalls := 0
		lastBlockSeen := uint64(0)
		gS, storeCtx := errgroup.WithContext(ctx)
		// could change this to for - range
		for {
			select {
			case <-groupCtx.Done():
				logger.ReportIndexerError(groupCtx.Err(), x.indexerConfig, logger.ContextCancelled)
				return fmt.Errorf("context canceled while storing and retrieving logs: %w", groupCtx.Err())
			case log, ok := <-logsChan: // empty log passed when ok is false.
				if !ok {
					return nil
				}
				concurrentCalls++
				gS.Go(func() error {
					// another goroutine is already storing this receipt
					locker, ok := x.mux.TryLock(log.TxHash)
					if !ok {
						return nil
					}
					defer locker.Unlock()

					// Check if the txHash has already been stored in the cache.
					if _, ok := x.cache.Get(log.TxHash); ok {
						return nil
					}

					err := x.store(storeCtx, log)
					if err != nil {
						logger.ReportIndexerError(err, x.indexerConfig, logger.StoreError)

						return fmt.Errorf("could not store log: %w", err)
					}

					return nil
				})

				// Checks if:
				// 1. The number of concurrent calls is greater than the concurrency threshold.
				// 2. The indexer's distance from the chaintip is within the concurrency ending threshold.
				// If so, all the go routines are waited on and the last indexed block is stored.
				if concurrentCalls >= x.indexerConfig.StoreConcurrency || x.indexerConfig.ConcurrencyThreshold > endHeight-log.BlockNumber {
					if err = gS.Wait(); err != nil {
						return fmt.Errorf("error waiting for go routines: %w", err)
					}

					// reset group context and concurrent calls
					gS, storeCtx = errgroup.WithContext(ctx)
					concurrentCalls = 0

					// Only update last indexed if all logs from the last block have been processed to prevent premature
					// updates of last indexed. Prevents having to lag a block behind on downstream dependencies (agents).
					if lastBlockSeen < log.BlockNumber {
						err = x.saveLastIndexed(storeCtx, lastBlockSeen)
						if err != nil {
							logger.ReportIndexerError(err, x.indexerConfig, logger.StoreError)
							return fmt.Errorf("could not store last indexed: %w", err)
						}
						lastBlockSeen = log.BlockNumber
					}

					x.blockMeter.Record(ctx, int64(log.BlockNumber), otelMetrics.WithAttributeSet(
						attribute.NewSet(attribute.Int64("start_block", int64(startHeight)), attribute.Int64("chain_id", int64(x.indexerConfig.ChainID)))),
					)
				}
			}
		}
	})

	err = g.Wait()

	if err != nil {
		return fmt.Errorf("could not backfill contract: %w \nChain: %d\nLog 's Contract Address: %s\n ", err, x.indexerConfig.ChainID, x.indexerConfig.Addresses)
	}

	err = x.saveLastIndexed(ctx, endHeight)
	if err != nil {
		logger.ReportIndexerError(err, x.indexerConfig, logger.StoreError)
		return fmt.Errorf("could not store last indexed: %w", err)
	}

	x.blockMeter.Record(ctx, int64(endHeight), otelMetrics.WithAttributeSet(
		attribute.NewSet(attribute.Int64("start_block", int64(startHeight)), attribute.Int64("chain_id", int64(x.indexerConfig.ChainID)))),
	)

	return nil
}

// TODO split two goroutines into sep functions for maintainability
// store stores the logs, receipts, and transactions for a tx hash.
//
//nolint:cyclop,gocognit,maintidx
func (x *Indexer) store(parentCtx context.Context, log types.Log) (err error) {
	ctx, span := x.handler.Tracer().Start(parentCtx, "store", trace.WithAttributes(
		attribute.String("contract", x.addressesToString(x.indexerConfig.Addresses)),
		attribute.String("tx", log.TxHash.Hex()),
		attribute.String("block", fmt.Sprintf("%d", log.BlockNumber)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    3 * time.Millisecond,
		Max:    2 * time.Second,
	}

	timeout := time.Duration(0)
	tryCount := 0

	var tx *txData
	hasTX := true

OUTER:
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context canceled while storing logs/receipts: %w", ctx.Err())
		case <-time.After(timeout):
			tryCount++

			tx, err = x.fetchEventData(ctx, log.TxHash, log.BlockNumber)
			if err != nil {
				if errors.Is(err, errNoContinue) {
					logger.ReportIndexerError(err, x.indexerConfig, logger.GetTxError)
					return nil
				}

				if errors.Is(err, errNoTx) {
					logger.ReportIndexerError(err, x.indexerConfig, logger.GetTxError)
					hasTX = false
					break OUTER
				}

				if tryCount > retryTolerance {
					return fmt.Errorf("retry tolerance exceeded: %w", err)
				}

				timeout = b.Duration()
				continue
			}

			break OUTER
		}
	}

	g, groupCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		// Store receipt in the EventDB.
		if x.toHead {
			err = x.eventDB.StoreReceiptAtHead(groupCtx, x.indexerConfig.ChainID, tx.receipt)
		} else {
			err = x.eventDB.StoreReceipt(groupCtx, x.indexerConfig.ChainID, tx.receipt)
		}
		if err != nil {
			return fmt.Errorf("could not store receipt: %w", err)
		}
		return nil
	})

	if hasTX {
		g.Go(func() error {
			if x.toHead {
				err = x.eventDB.StoreEthTxAtHead(groupCtx, &tx.transaction, x.indexerConfig.ChainID, log.BlockHash, log.BlockNumber, uint64(log.TxIndex))
			} else {
				err = x.eventDB.StoreEthTx(groupCtx, &tx.transaction, x.indexerConfig.ChainID, log.BlockHash, log.BlockNumber, uint64(log.TxIndex))
			}
			if err != nil {
				return fmt.Errorf("could not store tx: %w", err)
			}
			return nil
		})
	}

	g.Go(func() error {
		logs, err := x.prunedReceiptLogs(tx.receipt)
		if err != nil {
			return err
		}
		if x.toHead {
			err = x.eventDB.StoreLogsAtHead(groupCtx, x.indexerConfig.ChainID, logs...)
		} else {
			err = x.eventDB.StoreLogs(groupCtx, x.indexerConfig.ChainID, logs...)
		}
		if err != nil {
			return fmt.Errorf("could not store receipt logs: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		err := x.eventDB.StoreBlockTime(groupCtx, x.indexerConfig.ChainID, tx.blockHeader.Number.Uint64(), tx.blockHeader.Time)
		if err != nil {
			return fmt.Errorf("could not store receipt logs: %w", err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("could not store data: %w\n%s on chain %d from %d to %s", err, x.addressesToString(x.indexerConfig.Addresses), x.indexerConfig.ChainID, log.BlockNumber, log.TxHash.String())
	}

	x.cache.Add(log.TxHash, true)
	return nil
}

// prunedReceiptLogs gets all logs from a receipt and prunes null logs.
func (x *Indexer) prunedReceiptLogs(receipt types.Receipt) (logs []types.Log, err error) {
	for i := range receipt.Logs {
		log := receipt.Logs[i]
		if log == nil {
			return nil, fmt.Errorf("log is nil\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nLog 's Contract Address: %s\nContract Address: %s", x.indexerConfig.ChainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), x.addressesToString(x.indexerConfig.Addresses))
		}
		logs = append(logs, *log)
	}
	return logs, nil
}

// fetchEventData tries to fetch a transaction from the cache, if it's not there it tries to fetch it from the database.
// nolint: cyclop
func (x *Indexer) fetchEventData(parentCtx context.Context, txhash common.Hash, blockNumber uint64) (tx *txData, err error) {
	ctx, span := x.handler.Tracer().Start(parentCtx, "fetchEventData", trace.WithAttributes(
		attribute.String("tx", txhash.Hex()),
		attribute.String("block", fmt.Sprintf("%d", blockNumber)),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

OUTER:
	// increasing this across more clients puts too much load on the server, results in failed requests. TODO investigate
	for i := range x.client[0:1] {
		tx = &txData{}

		calls := make([]w3types.Caller, 3)

		// setup referencable indexes so we can access errors from the calls
		const (
			receiptIndex = 0
			txIndex      = 1
			headerIndex  = 2
		)

		// get the transaction receipt
		calls[receiptIndex] = eth.TxReceipt(txhash).Returns(&tx.receipt)

		// get the raw transaction
		calls[txIndex] = eth.Tx(txhash).Returns(&tx.transaction)

		// get the block number
		calls[headerIndex] = eth.HeaderByNumber(new(big.Int).SetUint64(blockNumber)).Returns(&tx.blockHeader)

		//nolint: nestif
		if err := x.client[i].BatchWithContext(ctx, calls...); err != nil {
			//nolint: errorlint
			callErr, ok := err.(w3.CallErrors)
			if !ok {
				return nil, fmt.Errorf("could not parse errors: %w", err)
			}

			if callErr[receiptIndex] != nil {
				if callErr[receiptIndex].Error() == txNotFoundError {
					logger.ReportIndexerError(fmt.Errorf(txNotFoundError), x.indexerConfig, logger.GetTxError)
					continue OUTER
				}
			}

			if callErr[txIndex] != nil {
				switch callErr[txIndex].Error() {
				case txNotSupportedError:
					logger.ReportIndexerError(fmt.Errorf(txNotSupportedError), x.indexerConfig, logger.GetTxError)
					return tx, errNoTx
				case invalidTxVRSError:
					logger.ReportIndexerError(fmt.Errorf(invalidTxVRSError), x.indexerConfig, logger.GetTxError)
					return tx, errNoTx
				case txNotFoundError:
					logger.ReportIndexerError(fmt.Errorf(txNotFoundError), x.indexerConfig, logger.GetTxError)
					continue OUTER
				}
			}

			return nil, fmt.Errorf("could not get tx receipt: %w", err)
		}

		tx.success = true
	}

	if tx == nil || !tx.success {
		return nil, fmt.Errorf("could not get tx data: %w", err)
	}

	return tx, nil
}

// addressesToString is a helper function for logging events.
func (x *Indexer) addressesToString(addresses []common.Address) string {
	var output string
	for i := range addresses {
		if i == 0 {
			output = addresses[i].String()
		} else {
			output = output + "," + addresses[i].String()
		}
	}
	return output
}

func (x *Indexer) saveLastIndexed(parentCtx context.Context, blockNumber uint64) error {
	if !x.isBackfill {
		var err error
		var errMessage string
		if x.toHead {
			err = x.eventDB.StoreLastIndexed(parentCtx, common.Address{}, x.indexerConfig.ChainID, blockNumber, scribeTypes.LivefillAtHead)
			errMessage = "could not store last indexed block while livefilling at head"
		} else {
			err = x.eventDB.StoreLastIndexedMultiple(parentCtx, x.indexerConfig.Addresses, x.indexerConfig.ChainID, blockNumber)
			errMessage = "could not store last indexed blocks"
		}
		if err != nil {
			logger.ReportIndexerError(err, x.indexerConfig, logger.StoreError)
			return fmt.Errorf("%s: %w", errMessage, err)
		}
	}
	return nil
}
