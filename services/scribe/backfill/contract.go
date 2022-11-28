package backfill

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	lru "github.com/hashicorp/golang-lru"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"golang.org/x/sync/errgroup"
)

// ContractBackfiller is a backfiller that fetches logs for a specific contract.
type ContractBackfiller struct {
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
				return nil
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
			case <-doneChan:
				return nil
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

	// Parallelize storing logs, receipts, and transactions.
	g, groupCtx := errgroup.WithContext(ctx)

	doneChan := make(chan bool, 2)

	// Get the receipt for the log and notify the other goroutines that depend on it that
	// it has been retrieved via doneChan.
	g.Go(func() error {
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
					return fmt.Errorf("could not get transaction receipt for txHash: %w\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nLog 's Contract Address: %s", err, c.chainConfig.ChainID, log.TxHash.String(), log.BlockNumber, log.Address.String())
				}
			default:
				return fmt.Errorf("could not get transaction receipt for txHash: %w\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nLog 's Contract Address: %s", err, c.chainConfig.ChainID, log.TxHash.String(), log.BlockNumber, log.Address.String())
			}
		}

		returnedReceipt = *receipt
		doneChan <- true
		doneChan <- true

		return nil
	})

	// Get the logs from the receipt and store them in the EventDB.
	g.Go(func() error {
		select {
		case <-groupCtx.Done():
			LogEvent(ErrorLevel, "Context canceled", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address})

			return fmt.Errorf("context canceled\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nLog 's Contract Address: %s\nContract Address: %s", c.chainConfig.ChainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
		case <-doneChan:
			for _, log := range returnedReceipt.Logs {
				if log == nil {
					LogEvent(ErrorLevel, "log is nil", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address})

					return fmt.Errorf("log is nil\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nLog 's Contract Address: %s\nContract Address: %s", c.chainConfig.ChainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
				}

				err := c.eventDB.StoreLog(groupCtx, *log, c.chainConfig.ChainID)
				if err != nil {
					LogEvent(ErrorLevel, "Could not store log", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address})

					return fmt.Errorf("could not store log: %w\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nLog 's Contract Address: %s\nContract Address: %s", err, c.chainConfig.ChainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
				}
			}

			return nil
		}
	})

	// Get the transaction from the receipt and store it in the EventDB.
	g.Go(func() error {
		select {
		case <-groupCtx.Done():
			LogEvent(ErrorLevel, "Context canceled", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address})

			return fmt.Errorf("context canceled\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nLog 's Contract Address: %s\nContract Address: %s", c.chainConfig.ChainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
		case <-doneChan:
			err := c.eventDB.StoreReceipt(groupCtx, returnedReceipt, c.chainConfig.ChainID)
			if err != nil {
				return fmt.Errorf("could not store receipt: %w\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nLog 's Contract Address: %s\nContract Address: %s", err, c.chainConfig.ChainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
			}

			return nil
		}
	})

	// Store the transaction in the EventDB, while checking for and handling common errors.
	g.Go(func() error {
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
					LogEvent(ErrorLevel, "Could not get tx for txHash", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

					return fmt.Errorf("could not get transaction receipt for txHash")
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
			LogEvent(ErrorLevel, "Could not store transaction", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

			return fmt.Errorf("could not store transaction: %w", err)
		}

		return nil
	})

	err := g.Wait()
	if err != nil {
		LogEvent(ErrorLevel, "Could not store data", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

		return fmt.Errorf("could not store data: %w\n%s on chain %d from %d to %s", err, c.address, c.chainConfig.ChainID, log.BlockNumber, log.TxHash.String())
	}

	err = c.eventDB.StoreLastIndexed(ctx, common.HexToAddress(c.address), c.chainConfig.ChainID, returnedReceipt.BlockNumber.Uint64())
	if err != nil {
		LogEvent(ErrorLevel, "Could not store last indexed block", LogData{"cid": c.chainConfig.ChainID, "bn": log.BlockNumber, "tx": log.TxHash.Hex(), "la": log.Address.String(), "ca": c.address, "e": err.Error()})

		return fmt.Errorf("could not store last indexed block: %w", err)
	}

	c.cache.Add(log.TxHash, true)

	return nil
}

// getLogs gets all logs for the contract through channels constructed and populated by the rangeFilter.
func (c ContractBackfiller) getLogs(ctx context.Context, startHeight, endHeight uint64) (<-chan types.Log, <-chan bool) {
	rangeFilter := NewRangeFilter(common.HexToAddress(c.address), c.client[0], big.NewInt(int64(startHeight)), big.NewInt(int64(endHeight)), c.chainConfig.ContractChunkSize, true, c.chainConfig.ContractSubChunkSize)
	g, ctx := errgroup.WithContext(ctx)

	// Concurrently start the range filter.
	g.Go(func() error {
		err := rangeFilter.Start(ctx)
		if err != nil {
			return fmt.Errorf("could not filter range: %w \nChain: %d\nstart height: %d, end: %d\nContract Address: %s", err, c.chainConfig.ChainID, startHeight, endHeight, c.address)
		}

		return nil
	})

	logsChan := make(chan types.Log)
	doneChan := make(chan bool)

	// Concurrently read from the range filter and send to the logsChan.
	g.Go(func() error {
	OUTER:
		for {
			select {
			case <-ctx.Done():
				LogEvent(ErrorLevel, "Context canceled while getting log", LogData{"cid": c.chainConfig.ChainID, "sh": startHeight, "eh": endHeight})

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

					break OUTER
				}
			}
		}

		return nil
	})

	return logsChan, doneChan
}
