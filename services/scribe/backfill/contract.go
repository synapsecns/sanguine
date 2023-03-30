package backfill

import (
	"context"
	"errors"
	"fmt"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/core/mapmutex"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	lru "github.com/hashicorp/golang-lru"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"golang.org/x/sync/errgroup"
)

// ContractBackfiller is a backfiller that fetches logs for a specific contract.
type ContractBackfiller struct {
	// chainConfig is the chain config for the chain that the contract is on.
	chainConfig config.ChainConfig
	// address is the contract address to get logs for.
	address string
	// eventDB is the database to store event data in.
	eventDB db.EventDB
	// client is the client for filtering.
	client []ScribeBackend
	// cache is a cache for txHashes.
	cache *lru.Cache
	// mux is the mutex used to prevent double inserting logs from the same tx
	mux mapmutex.StringerMapMutex
	// handler is the metrics handler for the scribe.
	handler metrics.Handler
}

// retryTolerance is the number of times to retry a failed operation before rerunning the entire Backfill function.
const retryTolerance = 20

// txNotSupportedError is for handling the legacy Arbitrum tx type.
const txNotSupportedError = "transaction type not supported"

// invalidTxVRSError is for handling Aurora VRS error.
const invalidTxVRSError = "invalid transaction v, r, s values"

// txNotFoundError is for handling omniRPC errors for BSC.
const txNotFoundError = "not found"

// NewContractBackfiller creates a new backfiller for a contract.
func NewContractBackfiller(chainConfig config.ChainConfig, address string, eventDB db.EventDB, client []ScribeBackend, handler metrics.Handler) (*ContractBackfiller, error) {
	cache, err := lru.New(500)
	if err != nil {
		return nil, fmt.Errorf("could not initialize cache: %w", err)
	}

	return &ContractBackfiller{
		chainConfig: chainConfig,
		address:     address,
		eventDB:     eventDB,
		client:      client,
		cache:       cache,
		mux:         mapmutex.NewStringerMapMutex(),
		handler:     handler,
	}, nil
}

// Backfill retrieves logs, receipts, and transactions for a contract from a given range and does so in the following manner.
// 1. Get logs for the contract in chunks of batch requests.
// 2. Iterate through each log's Tx Hash and performs the following
//   - Get the receipt for each log and store it and all of its logs.
//   - Get the transaction for each log and store it.
//
//nolint:gocognit, cyclop
func (c *ContractBackfiller) Backfill(parentCtx context.Context, givenStart uint64, endHeight uint64) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "contract.Backfill", trace.WithAttributes(
		attribute.Int("chain", int(c.chainConfig.ChainID)),
		attribute.String("address", c.address),
		attribute.Int("start", int(givenStart)),
		attribute.Int("end", int(endHeight)),
	))

	defer func() {
		go func() {
			if err != nil {
				span.RecordError(err)
			}
			span.End()
		}()
	}()

	g, groupCtx := errgroup.WithContext(ctx)
	startHeight := givenStart
	lastBlockIndexed, err := c.eventDB.RetrieveLastIndexed(groupCtx, common.HexToAddress(c.address), c.chainConfig.ChainID)
	if err != nil {
		LogEvent(WarnLevel, "Could not get last indexed", LogData{"cid": c.chainConfig.ChainID, "sh": startHeight, "eh": endHeight, "e": err.Error()})

		return fmt.Errorf("could not get last indexed: %w", err)
	}
	if lastBlockIndexed > startHeight {
		LogEvent(WarnLevel, "Using last indexed block (lastIndexBlock > startHeight)", LogData{"cid": c.chainConfig.ChainID, "sh": startHeight, "eh": endHeight})
		startHeight = lastBlockIndexed + 1
	}
	LogEvent(InfoLevel, "Beginning to backfill contract ", LogData{"cid": c.chainConfig.ChainID, "sh": startHeight, "eh": endHeight})

	// logsChain and doneChan are used to pass logs from rangeFilter onto the next stage of the backfill process.
	logsChan, doneChan := c.getLogs(groupCtx, startHeight, endHeight)

	// Reads from the local logsChan and stores the logs and associated receipts / txs.
	g.Go(func() error {
		concurrentCalls := 0
		gS, storeCtx := errgroup.WithContext(ctx)
		for {
			select {
			case <-groupCtx.Done():
				LogEvent(ErrorLevel, "Context canceled while storing and retrieving logs", LogData{"cid": c.chainConfig.ChainID, "ca": c.address})

				return fmt.Errorf("context canceled while storing and retrieving logs: %w", groupCtx.Err())
			case log := <-logsChan:
				concurrentCalls++
				gS.Go(func() error {
					// another goroutine is already storing this receipt
					locker, ok := c.mux.TryLock(log.TxHash)
					if !ok {
						return nil
					}
					defer locker.Unlock()

					// Check if the txHash has already been stored in the cache.
					if _, ok := c.cache.Get(log.TxHash); ok {
						return nil
					}

					err := c.store(storeCtx, log)
					if err != nil {
						LogEvent(ErrorLevel, "Could not store log", LogData{"cid": c.chainConfig.ChainID, "ca": c.address, "e": err.Error()})

						return fmt.Errorf("could not store log: %w", err)
					}

					return nil
				})

				// Stop spawning store threads and wait
				if concurrentCalls >= c.chainConfig.StoreConcurrency || endHeight-log.BlockNumber < c.chainConfig.StoreConcurrencyThreshold {
					if err = gS.Wait(); err != nil {
						return fmt.Errorf("error waiting for go routines: %w", err)
					}

					// Reset context TODO make this better
					gS, storeCtx = errgroup.WithContext(ctx)
					concurrentCalls = 0
					err = c.eventDB.StoreLastIndexed(ctx, common.HexToAddress(c.address), c.chainConfig.ChainID, log.BlockNumber)
					if err != nil {
						LogEvent(ErrorLevel, "Could not store last indexed block", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

						return fmt.Errorf("could not store last indexed block: %w", err)
					}
				}

			case doneFlag := <-doneChan:

				if doneFlag {
					LogEvent(InfoLevel, "Received doneChan", LogData{"cid": c.chainConfig.ChainID, "ca": c.address})

					return nil
				}
				return fmt.Errorf("doneChan returned false, context cancel while storing and retrieving logs")
			}
		}
	})

	err = g.Wait()

	if err != nil {
		return fmt.Errorf("could not backfill contract: %w \nChain: %d\nLog 's Contract Address: %s\nContract Address: %s", err, c.chainConfig.ChainID, c.address, c.address)
	}

	LogEvent(InfoLevel, "Finished backfilling contract", LogData{"cid": c.chainConfig.ChainID, "ca": c.address})

	return nil
}

// TODO split two goroutines into sep functions for maintainability
// store stores the logs, receipts, and transactions for a tx hash.
//
//nolint:cyclop,gocognit,maintidx
func (c *ContractBackfiller) store(parentCtx context.Context, log types.Log) (err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "store", trace.WithAttributes(
		attribute.String("contract", c.address),
		attribute.String("tx", log.TxHash.Hex()),
		attribute.String("block", fmt.Sprintf("%d", log.BlockNumber)),
	))

	defer func() {
		go func() {
			if err != nil {
				span.RecordError(err)
			}
			span.End()
		}()
	}()

	startTime := time.Now()

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
			LogEvent(ErrorLevel, "Context canceled while storing logs/receipts", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": ctx.Err()})

			return fmt.Errorf("context canceled while storing logs/receipts: %w", ctx.Err())
		case <-time.After(timeout):
			tryCount++

			tx, err = c.fetchTx(ctx, log.TxHash, log.BlockNumber)
			if err != nil {
				if errors.Is(err, errNoContinue) {
					return nil
				}

				if errors.Is(err, errNoTx) {
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

	// TODO: this will all be handled in the store function
	g, groupCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		// Store receipt in the EventDB.
		err = c.eventDB.StoreReceipt(groupCtx, c.chainConfig.ChainID, tx.receipt)
		if err != nil {
			LogEvent(ErrorLevel, "Could not store receipt, retrying", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

			return fmt.Errorf("could not store receipt: %w", err)
		}
		return nil
	})

	if hasTX {
		g.Go(func() error {
			err = c.eventDB.StoreEthTx(groupCtx, &tx.transaction, c.chainConfig.ChainID, log.BlockHash, log.BlockNumber, uint64(log.TxIndex))
			if err != nil {
				return fmt.Errorf("could not store tx: %w", err)
			}
			return nil
		})
	}

	g.Go(func() error {
		logs, err := c.prunedReceiptLogs(tx.receipt)
		if err != nil {
			return err
		}

		err = c.eventDB.StoreLogs(groupCtx, c.chainConfig.ChainID, logs...)
		if err != nil {
			return fmt.Errorf("could not store receipt logs: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		err := c.eventDB.StoreBlockTime(groupCtx, c.chainConfig.ChainID, tx.blockHeader.Number.Uint64(), tx.blockHeader.Time)
		if err != nil {
			return fmt.Errorf("could not store receipt logs: %w", err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		LogEvent(ErrorLevel, "Could not store data", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

		return fmt.Errorf("could not store data: %w\n%s on chain %d from %d to %s", err, c.address, c.chainConfig.ChainID, log.BlockNumber, log.TxHash.String())
	}

	c.cache.Add(log.TxHash, true)
	LogEvent(InfoLevel, "Log, Receipt, and Tx stored", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "ts": time.Since(startTime).Seconds()})

	return nil
}

func (c *ContractBackfiller) getLogs(parentCtx context.Context, startHeight, endHeight uint64) (<-chan types.Log, <-chan bool) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "getLogs")

	defer func() {
		go func() {
			span.End()
		}()
	}()

	// rangeFilter generates filter type that will retrieve logs from omnirpc in chunks of batch requests specified in the config.
	rangeFilter := NewRangeFilter(common.HexToAddress(c.address), c.client[0], big.NewInt(int64(startHeight)), big.NewInt(int64(endHeight)), c.chainConfig.ContractChunkSize, true, c.chainConfig.ContractSubChunkSize, c.chainConfig.ChainID)
	logsChan := make(chan types.Log)
	doneChan := make(chan bool)
	// This go routine is responsible for running the range filter and collect logs from omnirpc and put it into it's logChan (see filter.go).
	go func() {
		err := rangeFilter.Start(ctx)
		if err != nil {
			doneChan <- false
			return
		}
	}()

	// Reads from the range filter's logsChan and puts the logs into the local logsChan until completion.
	go func() {
		for {
			select {
			case <-ctx.Done():
				LogEvent(ErrorLevel, "Context canceled while getting log", LogData{"cid": c.chainConfig.ChainID, "sh": startHeight, "eh": endHeight, "e": ctx.Err()})
				doneChan <- false
				return
			case logInfos := <-rangeFilter.GetLogChan():
				for _, log := range logInfos.logs {
					logsChan <- log
				}

			default:
				if rangeFilter.Done() {
					finLogs, _ := rangeFilter.Drain(ctx)

					for _, log := range finLogs {
						logsChan <- log
					}
					doneChan <- true
					return
				}
			}
		}
	}()

	return logsChan, doneChan
}

// prunedReceiptLogs gets all logs from a receipt and prunes null logs.
func (c *ContractBackfiller) prunedReceiptLogs(receipt types.Receipt) (logs []types.Log, err error) {
	for i := range receipt.Logs {
		log := receipt.Logs[i]
		if log == nil {
			LogEvent(ErrorLevel, "log is nil", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address})

			return nil, fmt.Errorf("log is nil\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nLog 's Contract Address: %s\nContract Address: %s", c.chainConfig.ChainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
		}
		logs = append(logs, *log)
	}
	return logs, nil
}

// txData returns the transaction data for a given transaction hash.
type txData struct {
	receipt     types.Receipt
	transaction types.Transaction
	blockHeader types.Header
	success     bool
}

var errNoContinue = errors.New("encountered unreconcilable error, will not attempt to store tx")

// errNoTx indicates a tx cannot be parsed, this is only returned when the tx doesn't match our data model.
var errNoTx = errors.New("tx is not supported by the client")

// fetchTx tries to fetch a transaction from the cache, if it's not there it tries to fetch it from the database.
// nolint: cyclop
func (c *ContractBackfiller) fetchTx(parentCtx context.Context, txhash common.Hash, blockNumber uint64) (tx *txData, err error) {
	ctx, span := c.handler.Tracer().Start(parentCtx, "fetchTx", trace.WithAttributes(
		attribute.String("tx", txhash.Hex()),
		attribute.String("block", fmt.Sprintf("%d", blockNumber)),
	))

	defer func() {
		go func() {
			if err != nil {
				span.RecordError(err)
			}
			span.End()
		}()
	}()

OUTER:
	// increasing this across more clients puts too much load on the server, results in failed requests. TODO investigate
	for i := range c.client[0:1] {
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
		if err := c.client[i].Batch(ctx, calls...); err != nil {
			//nolint: errorlint
			callErr, ok := err.(w3.CallErrors)
			if !ok {
				return nil, fmt.Errorf("could not parse errors: %w", err)
			}

			if callErr[receiptIndex] != nil {
				if callErr[receiptIndex].Error() == txNotFoundError {
					LogEvent(InfoLevel, "Could not get tx for txHash, attempting with additional confirmations", LogData{"cid": c.chainConfig.ChainID, "tx": txhash, "ca": c.address, "e": err.Error()})
					continue OUTER
				}
			}

			if callErr[txIndex] != nil {
				switch callErr[txIndex].Error() {
				case txNotSupportedError:
					LogEvent(InfoLevel, "Invalid tx", LogData{"cid": c.chainConfig.ChainID, "tx": txhash, "ca": c.address, "e": err.Error()})
					return tx, errNoTx
				case invalidTxVRSError:
					LogEvent(InfoLevel, "Could not get tx for txHash, attempting with additional confirmations", LogData{"cid": c.chainConfig.ChainID, "tx": txhash, "ca": c.address, "e": err.Error()})
					return tx, errNoTx
				case txNotFoundError:
					LogEvent(InfoLevel, "Could not get tx for txHash, attempting with additional confirmations", LogData{"cid": c.chainConfig.ChainID, "tx": txhash, "ca": c.address, "e": err.Error()})
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
