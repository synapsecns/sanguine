package backfill

import (
	"context"
	"fmt"
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
}

var logsChanLenFrequency = 1000

// txNotSupportedError is for handling the legacy Arbitrum tx type.
const txNotSupportedError = "transaction type not supported"

// invalidTxVRSError is for handling Aurora VRS error.
const invalidTxVRSError = "invalid transaction v, r, s values"

// txNotFoundError is for handling omniRPC errors for BSC.
const txNotFoundError = "not found"

// retryTolerance is the number of times to retry a failed operation.
const retryTolerance = 9

// NewContractBackfiller creates a new backfiller for a contract.
func NewContractBackfiller(chainConfig config.ChainConfig, address string, eventDB db.EventDB, client []ScribeBackend) (*ContractBackfiller, error) {
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
	}, nil
}

// Backfill takes in a channel of logs, uses each log to get the receipt from its txHash,
// gets all the logs from the receipt, then stores the receipt, the logs from the
// receipt, and the last indexed block for hte contract in the EventDB.
//
//nolint:gocognit, cyclop
func (c *ContractBackfiller) Backfill(ctx context.Context, givenStart uint64, endHeight uint64) error {
	startHeight := givenStart
	lastBlockIndexed, _ := c.eventDB.RetrieveLastIndexed(ctx, common.HexToAddress(c.address), c.chainConfig.ChainID)

	if lastBlockIndexed > startHeight {
		LogEvent(WarnLevel, "Using last indexed block (lastIndexBlock > startHeight)", LogData{"cid": c.chainConfig.ChainID, "sh": startHeight, "eh": endHeight})
		startHeight = lastBlockIndexed
	}

	// In the case of a failed backfill, we want to start from the last indexed block - 1.
	if startHeight != 0 {
		startHeight--
	}

	g, groupCtx := errgroup.WithContext(ctx)
	LogEvent(InfoLevel, "Beginning to backfill contract ", LogData{"cid": c.chainConfig.ChainID, "sh": startHeight, "eh": endHeight})
	logsChan, doneChan := c.getLogs(groupCtx, startHeight, endHeight)

	// Concurrently get logs from the logsChan and store them in the EventDB.
	g.Go(func() error {
		for {
			select {
			case <-groupCtx.Done():
				LogEvent(InfoLevel, "Context cancelled while storing and retrieving logs", LogData{"cid": c.chainConfig.ChainID, "ca": c.address})

				return fmt.Errorf("context cancelled while storing and retrieving logs: %w", groupCtx.Err())
			case log := <-logsChan:
				// Check if the txHash has already been stored in the cache.
				if _, ok := c.cache.Get(log.TxHash); ok {
					continue
				}

				if len(logsChan)%logsChanLenFrequency == 0 && len(logsChan) != 0 {
					LogEvent(InfoLevel, "logsChan length", LogData{"lc": len(logsChan)})
				}

				err := c.store(groupCtx, log)
				if err != nil {
					return fmt.Errorf("could not store log: %w \nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nLog 's Contract Address: %s\nContract Address: %s", err, c.chainConfig.ChainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
				}
			case doneFlag := <-doneChan:
				if doneFlag {
					return nil
				}
				return fmt.Errorf("doneChan returned false, context cancel while storing and retrieving logs")
			}
		}
	})

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("could not backfill contract: %w \nChain: %d\nLog 's Contract Address: %s\nContract Address: %s", err, c.chainConfig.ChainID, c.address, c.address)
	}
	LogEvent(InfoLevel, "Finished backfilling contract", LogData{"cid": c.chainConfig.ChainID, "ca": c.address})
	return nil
}

// store stores the logs, receipts, and transactions for a tx hash.
//
//nolint:cyclop, gocognit
func (c *ContractBackfiller) store(ctx context.Context, log types.Log) error {
	var returnedReceipt types.Receipt
	startTime := time.Now()

	// Parallelize storing logs, receipts, and transactions.
	g, groupCtx := errgroup.WithContext(ctx)

	// Get the receipt for the log and notify the other goroutines that depend on it that
	// it has been retrieved via doneChan.
	g.Go(func() error {

		b := &backoff.Backoff{
			Factor: 2,
			Jitter: true,
			Min:    1 * time.Second,
			Max:    10 * time.Second,
		}
		timeout := time.Duration(0)
		tryCount := 0
	RETRY:
		tryCount++
		if tryCount > retryTolerance {
			return fmt.Errorf("retry tolerance exceeded")
		}

		select {
		case <-groupCtx.Done():
			LogEvent(ErrorLevel, "Context canceled while storing logs/receipts", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": groupCtx.Err()})

			return fmt.Errorf("context canceled while storing logs/receipts: %w", groupCtx.Err())

		case <-time.After(timeout):
			receipt, err := c.client[0].TransactionReceipt(ctx, log.TxHash)
			if err != nil {
				switch err.Error() {
				// txNotFoundError handles a null return from omnirpc, re-queries with a client with >1 confirmations.
				case txNotFoundError:
					// Try with client(s) with additional confirmations.
					for i := range c.client[1:] {
						client := c.client[i]
						receipt, err = client.TransactionReceipt(ctx, log.TxHash)
						if err == nil {
							break
						}
					}
					if err != nil {
						timeout = b.Duration()
						LogEvent(ErrorLevel, "Could not get transaction receipt for txHash, retrying", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

						goto RETRY
					}
				default:
					return fmt.Errorf("could not get transaction receipt for txHash: %w\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nLog 's Contract Address: %s", err, c.chainConfig.ChainID, log.TxHash.String(), log.BlockNumber, log.Address.String())
				}
			}

			returnedReceipt = *receipt

			// Store receipt in the EventDB.
			err = c.eventDB.StoreReceipt(groupCtx, returnedReceipt, c.chainConfig.ChainID)
			if err != nil {
				timeout = b.Duration()
				LogEvent(ErrorLevel, "Could not store receipt, retrying", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

				goto RETRY
			}

			// Store the logs in the EventDB.
			for _, log := range returnedReceipt.Logs {
				if log == nil {
					LogEvent(ErrorLevel, "log is nil", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address})

					return fmt.Errorf("log is nil\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nLog 's Contract Address: %s\nContract Address: %s", c.chainConfig.ChainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
				}

				err := c.eventDB.StoreLog(groupCtx, *log, c.chainConfig.ChainID)
				if err != nil {
					timeout = b.Duration()
					LogEvent(ErrorLevel, "Could not store log, retrying", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

					goto RETRY
				}
			}

			return nil
		}

	})

	// Store the transaction in the EventDB, while checking for and handling common errors.
	g.Go(func() error {
		b := &backoff.Backoff{
			Factor: 2,
			Jitter: true,
			Min:    1 * time.Second,
			Max:    10 * time.Second,
		}
		timeout := time.Duration(0)
		tryCount := 0
	RETRY:
		tryCount++
		if tryCount > retryTolerance {
			return fmt.Errorf("retry tolerance exceeded")
		}
		select {
		case <-groupCtx.Done():
			LogEvent(ErrorLevel, "Context canceled while storing logs/receipts", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": groupCtx.Err()})

			return fmt.Errorf("context canceled while storing logs/receipts: %w", groupCtx.Err())

		case <-time.After(timeout):
			txn, isPending, err := c.client[0].TransactionByHash(groupCtx, log.TxHash)
			if err != nil {
				switch err.Error() {
				// txNotSupportedError handles an unsupported tx (typically legacy arbitrum txs), skips the tx entirely.
				case txNotSupportedError:
					LogEvent(InfoLevel, "Invalid tx", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})
					return nil

				// invalidTxVRSError handles an invalid tx (typically aurora v,r,s not found txs), skips the tx entirely.
				case invalidTxVRSError:
					LogEvent(InfoLevel, "Invalid tx", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})
					return nil

				// txNotFoundError handles a null return from omnirpc, re-queries with a client with >1 confirmations.
				case txNotFoundError:
					LogEvent(InfoLevel, "Could not get tx for txHash, attempting with additional confirmations", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})
					// Try with client(s) with additional confirmations.
					for i := range c.client[1:] {
						LogEvent(InfoLevel, "Could not get tx for txHash, attempting with additional confirmations", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "cn": i, "e": err.Error()})

						client := c.client[i]
						txn, isPending, err = client.TransactionByHash(groupCtx, log.TxHash)
						if err == nil {
							break
						}
					}
					if err != nil {
						timeout = b.Duration()
						LogEvent(ErrorLevel, "Could not get tx for txHash", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

						goto RETRY
					}

				default:
					return fmt.Errorf("could not get transaction by hash: %w\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nLog 's Contract Address: %s\nContract Address: %s", err, c.chainConfig.ChainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
				}
			}

			if isPending {
				return fmt.Errorf("transaction is pending")
			}

			err = c.eventDB.StoreEthTx(groupCtx, txn, c.chainConfig.ChainID, log.BlockHash, log.BlockNumber, uint64(log.TxIndex))
			if err != nil {
				timeout = b.Duration()
				LogEvent(ErrorLevel, "Could not store eth tx", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

				goto RETRY
			}

			return nil
		}
	})

	err := g.Wait()
	if err != nil {
		LogEvent(ErrorLevel, "Could not store data", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

		return fmt.Errorf("could not store data: %w\n%s on chain %d from %d to %s", err, c.address, c.chainConfig.ChainID, log.BlockNumber, log.TxHash.String())
	}

	err = c.eventDB.StoreLastIndexed(ctx, common.HexToAddress(c.address), c.chainConfig.ChainID, log.BlockNumber)
	if err != nil {
		LogEvent(ErrorLevel, "Could not store last indexed block", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

		return fmt.Errorf("could not store last indexed block: %w", err)
	}

	c.cache.Add(log.TxHash, true)
	LogEvent(InfoLevel, "Log, Receipt, and Tx stored", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "ts": time.Since(startTime).Seconds()})

	return nil
}

// getLogs gets all logs for the contract through channels constructed and populated by the rangeFilter.
func (c ContractBackfiller) getLogs(ctx context.Context, startHeight, endHeight uint64) (<-chan types.Log, <-chan bool) {
	rangeFilter := NewRangeFilter(common.HexToAddress(c.address), c.client[0], big.NewInt(int64(startHeight)), big.NewInt(int64(endHeight)), c.chainConfig.ContractChunkSize, true, c.chainConfig.ContractSubChunkSize)
	logsChan := make(chan types.Log)
	doneChan := make(chan bool)

	go func() error {
		err := rangeFilter.Start(ctx)
		if err != nil {
			return fmt.Errorf("could not filter range: %w \nChain: %d\nstart height: %d, end: %d\nContract Address: %s", err, c.chainConfig.ChainID, startHeight, endHeight, c.address)
		}

		return nil
	}()

	// Concurrently read from the range filter and send to the logsChan.
	go func() error {
		for {
			select {
			case <-ctx.Done():
				LogEvent(ErrorLevel, "Context canceled while getting log", LogData{"cid": c.chainConfig.ChainID, "sh": startHeight, "eh": endHeight})
				doneChan <- false
				return nil
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

					return nil
				}
			}
		}
	}()

	return logsChan, doneChan
}
