package backfill

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	lru "github.com/hashicorp/golang-lru"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"golang.org/x/sync/errgroup"
)

// ContractBackfiller is a backfiller that fetches logs for a specific contract.
type ContractBackfiller struct {
	// chainID is the chainID of the chain the contract is deployed on
	chainID uint32
	// address is the contract address to get logs for
	address string
	// eventDB is the database to store event data in
	eventDB db.EventDB
	// client is the client for filtering
	client []ScribeBackend
	// cache is a cache for txHashes
	cache *lru.Cache
}

// ---- Specific Backfiller Errors ----
// txNotSupportedError is for handling the legacy Arbitrum tx type.
const txNotSupportedError = "transaction type not supported"

// invalidTxVRSError is for handling Aurora VRS error.
const invalidTxVRSError = "invalid transaction v, r, s values"

// txNotFoundError is for handling omniRPC errors for BSC.
const txNotFoundError = "not found"

// NewContractBackfiller creates a new backfiller for a contract.
func NewContractBackfiller(chainID uint32, address string, eventDB db.EventDB, client []ScribeBackend) (*ContractBackfiller, error) {
	// initialize the cache for the txHashes
	cache, err := lru.New(500)
	if err != nil {
		return nil, fmt.Errorf("could not initialize cache: %w", err)
	}

	return &ContractBackfiller{
		chainID: chainID,
		address: address,
		eventDB: eventDB,
		client:  client,
		cache:   cache,
	}, nil
}

// Backfill takes in a channel of logs, uses each log to get the receipt from its txHash,
// gets all of the logs from the receipt, then stores the receipt, the logs from the
// receipt, and the last indexed block for hte contract in the EventDB.
//
//nolint:gocognit, cyclop
func (c *ContractBackfiller) Backfill(ctx context.Context, givenStart uint64, endHeight uint64) error {
	// Get the correct start height
	startHeight := givenStart
	lastBlockIndexed, _ := c.eventDB.RetrieveLastIndexed(ctx, common.HexToAddress(c.address), c.chainID)

	if lastBlockIndexed > startHeight {
		logger.Warnf("last indexed block is greater than start height, using last indexed block: %d over the provided start block: %d\nc Address: %s", lastBlockIndexed, startHeight, c.address)
		startHeight = lastBlockIndexed
	}

	// in the case of a failed backfill, we want to start from the last indexed block - 1
	if startHeight != 0 {
		startHeight--
	}

	// start listening for logs
	g, groupCtx := errgroup.WithContext(ctx)
	logger.Infof("Backfilling contract %s on chain %d from %d to %d", c.address, c.chainID, startHeight, endHeight)
	logsChan, doneChan := c.getLogs(groupCtx, startHeight, endHeight)
	g.Go(func() error {
		for {
			select {
			case <-groupCtx.Done():
				return nil
			case log := <-logsChan:
				// check if the txHash has already been stored in the cache
				if _, ok := c.cache.Get(log.TxHash); ok {
					continue
				}
				err := c.store(groupCtx, log)
				if err != nil {
					return fmt.Errorf("could not store log: %w \nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nAddress: %s\nc Address: %s", err, c.chainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
				}
			case <-doneChan:
				return nil
			}
		}
	})

	err := g.Wait()

	if err != nil {
		return fmt.Errorf("could not backfill contract: %w \nChain: %d\nAddress: %s\nc Address: %s", err, c.chainID, c.address, c.address)
	}
	logger.Infof("Finished backfilling contract %s on chain %d from %d to %d", c.address, c.chainID, startHeight, endHeight)
	return nil
}

// store stores the logs, receipts, and transactions for a tx hash.
//
//nolint:cyclop, gocognit
func (c *ContractBackfiller) store(ctx context.Context, log types.Log) error {
	// parallelize storing logs, receipts, and transactions
	g, groupCtx := errgroup.WithContext(ctx)

	var returnedReceipt types.Receipt
	doneChan := make(chan bool, 2)
	g.Go(func() error {
		// make getting receipt a channel in parallel
		receipt, err := c.client[0].TransactionReceipt(ctx, log.TxHash)
		if err != nil {
			if err.Error() == txNotFoundError {
				// Try with client with additional confirmations
				receipt, err = c.client[1].TransactionReceipt(ctx, log.TxHash)
				if err != nil {
					return fmt.Errorf("could not get transaction receipt for txHash: %w\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nAddress: %s", err, c.chainID, log.TxHash.String(), log.BlockNumber, log.Address.String())
				}
			}
		}

		returnedReceipt = *receipt
		doneChan <- true
		doneChan <- true
		return nil
	})

	g.Go(func() error {
		select {
		case <-groupCtx.Done():
			return fmt.Errorf("context canceled\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nAddress: %s\nc Address: %s", c.chainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
		case <-doneChan:
			// get the logs from the receipt and store them in the db
			for _, log := range returnedReceipt.Logs {
				if log == nil {
					return fmt.Errorf("log is nil\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nAddress: %s\nc Address: %s", c.chainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
				}
				err := c.eventDB.StoreLog(groupCtx, *log, c.chainID)
				if err != nil {
					return fmt.Errorf("could not store log: %w\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nAddress: %s\nc Address: %s", err, c.chainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
				}
			}
			return nil
		}
	})

	g.Go(func() error {
		select {
		case <-groupCtx.Done():
			return fmt.Errorf("context canceled\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nAddress: %s\nc Address: %s", c.chainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
		case <-doneChan:
			// store the receipt in the db
			err := c.eventDB.StoreReceipt(groupCtx, returnedReceipt, c.chainID)
			if err != nil {
				return fmt.Errorf("could not store receipt: %w\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nAddress: %s\nc Address: %s", err, c.chainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
			}
			return nil
		}
	})

	g.Go(func() error {
		// store the transaction in the db
		txn, isPending, err := c.client[0].TransactionByHash(groupCtx, log.TxHash)
		if err != nil {
			if err.Error() == txNotSupportedError || err.Error() == invalidTxVRSError {
				logger.Warnf("Invalid tx: %s\n%s on chain id: %d\nLog BlockNumber: %d\nAddress: %s\nc Address: %s", err.Error(), log.TxHash.Hex(), c.chainID, log.BlockNumber, log.Address.String(), c.address)
				return nil
			}
			return fmt.Errorf("could not get transaction by hash: %w\nChain: %d\nTxHash: %s\nLog BlockNumber: %d\nAddress: %s\nc Address: %s", err, c.chainID, log.TxHash.String(), log.BlockNumber, log.Address.String(), c.address)
		}

		if isPending {
			return fmt.Errorf("transaction is pending")
		}

		err = c.eventDB.StoreEthTx(groupCtx, txn, c.chainID, log.BlockHash, log.BlockNumber, uint64(log.TxIndex))
		if err != nil {
			return fmt.Errorf("could not store transaction: %w", err)
		}

		return nil
	})

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("could not store data: %w\n%s on chain %d from %d to %s", err, c.address, c.chainID, log.BlockNumber, log.TxHash.String())
	}
	// store the last indexed block in the db
	err = c.eventDB.StoreLastIndexed(ctx, common.HexToAddress(c.address), c.chainID, returnedReceipt.BlockNumber.Uint64())
	if err != nil {
		return fmt.Errorf("could not store last indexed block: %w", err)
	}
	// store the txHash in the cache
	c.cache.Add(log.TxHash, true)

	return nil
}

// chunkSize is how big to make the chunks when fetching.
const chunkSize = 500

// getLogs gets all logs for the contract.
func (c ContractBackfiller) getLogs(ctx context.Context, startHeight, endHeight uint64) (<-chan types.Log, <-chan bool) {
	// start the filterer. This filters the range and sends the logs to the logChan.
	rangeFilter := NewRangeFilter(common.HexToAddress(c.address), c.client[0], big.NewInt(int64(startHeight)), big.NewInt(int64(endHeight)), chunkSize, true)
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		// start the range filterer, return any errors to an error channel
		err := rangeFilter.Start(ctx)
		if err != nil {
			return fmt.Errorf("could not filter range: %w \nChain: %d\nstart height: %d, end: %d\nc Address: %s", err, c.chainID, startHeight, endHeight, c.address)
		}
		return nil
	})

	logsChan := make(chan types.Log)
	doneChan := make(chan bool)

	// take the logs and put them in the log channel
	g.Go(func() error {
	OUTER:
		for {
			select {
			case <-ctx.Done():
				logger.Warnf("context canceled while getting logs\nChain: %d\nstart height: %d, end: %d\nAddress: %s\nc Address: %s", c.chainID, startHeight, endHeight, c.address, c.address)
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
