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
	client ScribeBackend
	// cache is a cache for txHashes
	cache *lru.Cache
}

// NewContractBackfiller creates a new backfiller for a contract.
func NewContractBackfiller(chainID uint32, address string, eventDB db.EventDB, client ScribeBackend) (*ContractBackfiller, error) {
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
	// initialize the channel for the logs
	startHeight, err := c.startHeightForBackfill(ctx, givenStart)
	if err != nil {
		return fmt.Errorf("could not get start height: %w", err)
	}
	// in the case of a failed backfill, we want to start from the last indexed block - 1
	if startHeight != 0 {
		startHeight--
	}
	// start listening for logs
	g, groupCtx := errgroup.WithContext(ctx)

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
				err = c.store(groupCtx, log)
				if err != nil {
					return fmt.Errorf("could not store log: %w", err)
				}
			case <-doneChan:
				return nil
			}
		}
	})

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("could not backfill contract: %w", err)
	}
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
		receipt, err := c.client.TransactionReceipt(ctx, log.TxHash)
		if err != nil {
			return fmt.Errorf("could not get transaction receipt for txHash: %w", err)
		}

		returnedReceipt = *receipt
		doneChan <- true
		doneChan <- true
		return nil
	})

	g.Go(func() error {
		select {
		case <-groupCtx.Done():
			return fmt.Errorf("context canceled")
		case <-doneChan:
			// get the logs from the receipt and store them in the db
			for _, log := range returnedReceipt.Logs {
				if log == nil {
					return fmt.Errorf("log is nil")
				}
				err := c.eventDB.StoreLog(groupCtx, *log, c.chainID)
				if err != nil {
					return fmt.Errorf("could not store log: %w", err)
				}
			}
			return nil
		}
	})

	g.Go(func() error {
		select {
		case <-groupCtx.Done():
			return fmt.Errorf("context canceled")
		case <-doneChan:
			// store the receipt in the db
			err := c.eventDB.StoreReceipt(groupCtx, returnedReceipt, c.chainID)
			if err != nil {
				return fmt.Errorf("could not store receipt: %w", err)
			}
			return nil
		}
	})

	g.Go(func() error {
		// store the transaction in the db
		txn, isPending, err := c.client.TransactionByHash(groupCtx, log.TxHash)
		if err != nil {
			return fmt.Errorf("could not get transaction by hash: %w", err)
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
		return fmt.Errorf("could not store data: %w", err)
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
	rangeFilter := NewRangeFilter(common.HexToAddress(c.address), c.client, big.NewInt(int64(startHeight)), big.NewInt(int64(endHeight)), chunkSize, true)
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		// start the range filterer, return any errors to an error channel
		err := rangeFilter.Start(ctx)
		if err != nil {
			return fmt.Errorf("could not filter range: %w", err)
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

// startHeightForBackfill gets the startHeight for backfilling. This is the maximum
// of the most recent block for the contract and the startHeight given in the config.
func (c ContractBackfiller) startHeightForBackfill(ctx context.Context, givenStart uint64) (startHeight uint64, err error) {
	lastBlock, err := c.eventDB.RetrieveLastIndexed(ctx, common.HexToAddress(c.address), c.chainID)
	if err != nil {
		return 0, fmt.Errorf("could not retrieve last indexed block for contract: %w", err)
	}

	if lastBlock > givenStart {
		return lastBlock, nil
	}
	return givenStart, nil
}
