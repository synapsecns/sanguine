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

// retryTolerance is the number of times to retry a failed operation before rerunning the entire Backfill function.
const retryTolerance = 1

// txNotSupportedError is for handling the legacy Arbitrum tx type.
const txNotSupportedError = "transaction type not supported"

// invalidTxVRSError is for handling Aurora VRS error.
const invalidTxVRSError = "invalid transaction v, r, s values"

// txNotFoundError is for handling omniRPC errors for BSC.
const txNotFoundError = "not found"

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

// Backfill retrieves logs, receipts, and transactions for a contract from a given range and does so in the following manner.
// 1. Get logs for the contract in chunks of batch requests.
// 2. Iterate through each log's Tx Hash and performs the following
//   - Get the receipt for each log and store it and all of its logs.
//   - Get the transaction for each log and store it.
//
//nolint:gocognit, cyclop
func (c *ContractBackfiller) Backfill(ctx context.Context, givenStart uint64, endHeight uint64) error {
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
		for {
			select {
			case <-groupCtx.Done():
				LogEvent(ErrorLevel, "Context canceled while storing and retrieving logs", LogData{"cid": c.chainConfig.ChainID, "ca": c.address})

				return fmt.Errorf("context canceled while storing and retrieving logs: %w", groupCtx.Err())
			case log := <-logsChan:
				// Check if the txHash has already been stored in the cache.
				if _, ok := c.cache.Get(log.TxHash); ok {
					continue
				}
				err := c.store(ctx, log)
				if err != nil {
					LogEvent(ErrorLevel, "Could not store log", LogData{"cid": c.chainConfig.ChainID, "ca": c.address, "e": err.Error()})

					return fmt.Errorf("could not store log: %w", err)
				}

				err = c.eventDB.StoreLastIndexed(ctx, common.HexToAddress(c.address), c.chainConfig.ChainID, log.BlockNumber)
				if err != nil {
					LogEvent(ErrorLevel, "Could not store last indexed block", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

					return fmt.Errorf("could not store last indexed block: %w", err)
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
//nolint:cyclop, gocognit, maintidx
func (c *ContractBackfiller) store(ctx context.Context, log types.Log) error {
	var returnedReceipt types.Receipt
	startTime := time.Now()

	// Parallelize storing logs, receipts, and transactions.
	g, groupCtx := errgroup.WithContext(ctx)

	// Get the log's receipt and store its receipts and logs
	g.Go(func() error {
		b := &backoff.Backoff{
			Factor: 2,
			Jitter: true,
			Min:    3 * time.Millisecond,
			Max:    2 * time.Second,
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
			receipt, err := c.client[0].TransactionReceipt(groupCtx, log.TxHash)
			if err != nil {
				switch err.Error() {
				// txNotFoundError handles a null return from omnirpc, re-queries with a client with >1 confirmations.
				case txNotFoundError:
					// Try with client(s) with additional confirmations.
					for i := range c.client[1:] {
						client := c.client[i]
						receipt, err = client.TransactionReceipt(groupCtx, log.TxHash)
						if err == nil {
							break
						}
					}
					if err != nil {
						timeout = b.Duration()
						LogEvent(ErrorLevel, "Could not get transaction receipt for txHash with extra confirmations, retrying", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

						goto RETRY
					}
				default:
					LogEvent(ErrorLevel, "Could not get transaction receipt for txHash, retrying", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

					goto RETRY
				}
			}

			returnedReceipt = *receipt

			// Parallelize storing logs, receipts, and transactions.
			gInner, groupInnerCtx := errgroup.WithContext(groupCtx)
			gInner.Go(func() error {
				// Store receipt in the EventDB.
				err = c.eventDB.StoreReceipt(groupInnerCtx, returnedReceipt, c.chainConfig.ChainID)
				if err != nil {
					timeout = b.Duration()
					LogEvent(ErrorLevel, "Could not store receipt, retrying", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

					return fmt.Errorf("could not store receipt: %w", err)
				}
				return nil
			})
			// Store the logs in the EventDB.
			for i := range returnedReceipt.Logs {
				log := returnedReceipt.Logs[i]
				if log == nil {
					LogEvent(ErrorLevel, "log is nil", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address})

					return fmt.Errorf("log is nil\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nLog 's Contract Address: %s\nContract Address: %s", c.chainConfig.ChainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
				}
				gInner.Go(func() error {
					err := c.eventDB.StoreLog(groupCtx, *log, c.chainConfig.ChainID)
					if err != nil {
						timeout = b.Duration()
						LogEvent(ErrorLevel, "Could not store log, retrying", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

						return fmt.Errorf("could not store log: %w", err)
					}
					return nil
				})
			}
			err = gInner.Wait()
			if err != nil {
				LogEvent(ErrorLevel, "Could not store data", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

				goto RETRY
			}
			return nil
		}
	})

	// Get the log's tx and store it.
	g.Go(func() error {
		b := &backoff.Backoff{
			Factor: 2,
			Jitter: true,
			Min:    3 * time.Millisecond,
			Max:    2 * time.Second,
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

	c.cache.Add(log.TxHash, true)
	LogEvent(InfoLevel, "Log, Receipt, and Tx stored", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "ts": time.Since(startTime).Seconds()})

	return nil
}

func (c ContractBackfiller) getLogs(ctx context.Context, startHeight, endHeight uint64) (<-chan types.Log, <-chan bool) {
	// rangeFilter generates filter type that will retrives logs from omnirpc in chunks of batch requests specified in the config.
	rangeFilter := NewRangeFilter(common.HexToAddress(c.address), c.client[0], big.NewInt(int64(startHeight)), big.NewInt(int64(endHeight)), c.chainConfig.ContractChunkSize, true, c.chainConfig.ContractSubChunkSize)
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
