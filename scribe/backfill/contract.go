package backfill

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	lru "github.com/hashicorp/golang-lru"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/scribe/db"
	"github.com/synapsecns/synapse-node/pkg/evm/client"
	"golang.org/x/sync/errgroup"
)

// ContractBackfiller is a backfiller that fetches logs for a specific contract.
type ContractBackfiller struct {
	// contract is the contract to get logs for
	contract contracts.DeployedContract
	// eventDB is the database to store event data in
	eventDB db.EventDB
	// client is the client for filtering
	client client.EVMClient
}

// NewContractBackfiller creates a new backfiller for a contract.
func NewContractBackfiller(contract contracts.DeployedContract, eventDB db.EventDB, client client.EVMClient) *ContractBackfiller {
	return &ContractBackfiller{
		contract: contract,
		eventDB:  eventDB,
		client:   client,
	}
}

// Backfill takes in a channel of logs, uses each log to get the receipt from its txHash,
// gets all of the logs from the receipt, then stores the receipt, the logs from the
// receipt, and the last indexed block for hte contract in the EventDB.
//
//nolint:gocognit, cyclop
func (c ContractBackfiller) Backfill(ctx context.Context, endHeight uint64) error {
	// initialize the cache for the txHashes
	cache, err := lru.New(500)
	if err != nil {
		return fmt.Errorf("could not initialize cache: %w", err)
	}
	// initialize the channel for the logs
	startHeight, err := c.StartHeightForBackfill(ctx, true)
	if err != nil {
		return fmt.Errorf("could not get start height: %w", err)
	}
	logChan, errChan, doneChan := c.GetLogs(ctx, startHeight, endHeight)
	// start listening for logs
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case log := <-logChan:
				// get the receipt that the log belongs to
				receipt, err := c.GetReceiptFromLog(ctx, log.TxHash, cache)
				if err != nil {
					return fmt.Errorf("could not get receipt from log: %w", err)
				}
				// if the receipt is nil, then the receipt has already been added to the db
				if receipt == nil {
					continue
				}
				// since `StoreReceipt` stores logs that it gets from the db, we need to
				// store the logs before we store the receipt

				// get the logs from the receipt and store them in the db
				for _, log := range receipt.Logs {
					err = c.eventDB.StoreLog(ctx, *log, uint32(c.contract.ChainID().Uint64()))
					if err != nil {
						return fmt.Errorf("could not store log: %w", err)
					}
				}
				// store the receipt in the db
				err = c.eventDB.StoreReceipt(ctx, *receipt, uint32(c.contract.ChainID().Uint64()))
				if err != nil {
					return fmt.Errorf("could not store receipt: %w", err)
				}
				// store the transaction in the db
				txn, isPending, err := c.client.TransactionByHash(ctx, receipt.TxHash)
				if err != nil {
					return fmt.Errorf("could not get transaction by hash: %w", err)
				}
				if isPending {
					return fmt.Errorf("transaction is pending")
				}
				err = c.eventDB.StoreEthTx(ctx, txn, uint32(c.contract.ChainID().Uint64()))
				if err != nil {
					return fmt.Errorf("could not store transaction: %w", err)
				}
				// store the last indexed block in the db
				err = c.eventDB.StoreLastIndexed(ctx, c.contract.Address(), uint32(c.contract.ChainID().Uint64()), receipt.BlockNumber.Uint64())
				if err != nil {
					return fmt.Errorf("could not store last indexed block: %w", err)
				}
			case err := <-errChan:
				return fmt.Errorf("could not get logs: %w", err)
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

// GetReceiptFromLog gets the receipt for a log if the receipt of a certain txHash
// does not exist in the cache of added txHashes.
func (c ContractBackfiller) GetReceiptFromLog(ctx context.Context, txHash common.Hash, cache *lru.Cache) (*types.Receipt, error) {
	// check the cache for the receipt
	if _, ok := cache.Get(txHash); ok {
		//nolint:nilnil
		return nil, nil
	}
	// get the receipt
	receipt, err := c.client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, fmt.Errorf("could not get transaction receipt for txHash: %w", err)
	}
	// add the receipt to the cache
	cache.Add(txHash, 1)
	return receipt, nil
}

// chunkSize is how big to make the chunks when fetching.
const chunkSize = 1024

// GetLogs gets all logs for the contract.
func (c ContractBackfiller) GetLogs(ctx context.Context, startHeight, endHeight uint64) (logsChan <-chan types.Log, errsChan <-chan error, completeChan <-chan bool) {
	// initialize the channel
	logChan := make(chan types.Log)
	errChan := make(chan error)
	doneChan := make(chan bool)

	// start the filterer. This filters the range and sends the logs to the logChan.
	rangeFilter := NewRangeFilter(c.contract.Address(), c.client, big.NewInt(int64(startHeight)), big.NewInt(int64(endHeight)), chunkSize, true)
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		// start the range filterer, return any errors to an error channel
		err := rangeFilter.Start(ctx)
		if err != nil {
			return fmt.Errorf("could not filter range: %w", err)
		}
		return nil
	})

	// take the logs and put them in the log channel
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case logInfos := <-rangeFilter.GetLogChan():
				for _, log := range logInfos.logs {
					logChan <- log
				}
				if rangeFilter.Done() {
					doneChan <- true
				}
			}
		}
	})

	// return errors to the channel when done filtering
	go func() {
		err := g.Wait()
		if err != nil {
			errChan <- err
		}
	}()

	return logChan, errChan, doneChan
}

// StartHeightForBackfill gets the creation height for the contract. This is the maximum
// of the most recent block for the contract and the block that the contract was deployed.
func (c ContractBackfiller) StartHeightForBackfill(ctx context.Context, useDB bool) (startHeight uint64, err error) {
	g, ctx := errgroup.WithContext(ctx)
	// maxHeight to return
	var maxHeight uint64
	// prevents race conditions when determining max height
	var maxHeightMux sync.Mutex

	// setMaxHeight sets the max height in a thread safe way
	setMaxHeight := func(height uint64) {
		maxHeightMux.Lock()
		defer maxHeightMux.Unlock()
		if maxHeight < height {
			maxHeight = height
		}
	}

	// Get the block number the contract was deployed in.
	g.Go(func() error {
		deployTxHash := c.contract.DeployTx().Hash()
		// get the block that the contract was created in
		receipt, err := c.client.TransactionReceipt(ctx, deployTxHash)
		if err != nil {
			return fmt.Errorf("could not get transaction receipt for contract: %w", err)
		}
		setMaxHeight(receipt.BlockNumber.Uint64())
		return nil
	})

	// If the useDB flag is set, get the most recent block for the contract from the database.
	g.Go(func() error {
		if useDB {
			// lastBlock will either be the last block that had stored indexed data for, or is 0 if no data is stored
			lastBlock, err := c.eventDB.RetrieveLastIndexed(ctx, c.contract.Address(), uint32(c.contract.ChainID().Uint64()))
			if err != nil {
				return fmt.Errorf("could not retrieve last indexed block for contract: %w", err)
			}
			setMaxHeight(lastBlock)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return 0, fmt.Errorf("error getting startHeight: %w", err)
	}

	return maxHeight, nil
}

// package backfill

// import (
// 	"context"
// 	"fmt"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum/core/types"
// 	lru "github.com/hashicorp/golang-lru"
// 	"github.com/synapsecns/sanguine/ethergo/contracts"
// 	"github.com/synapsecns/sanguine/scribe/db"
// 	"github.com/synapsecns/synapse-node/pkg/evm/client"
// 	"golang.org/x/sync/errgroup"
// )

// // ContractBackfiller is a backfiller that fetches logs for a specific contract.
// type ContractBackfiller struct {
// 	// contract is the contract to get logs for
// 	contract contracts.DeployedContract
// 	// eventDB is the database to store event data in
// 	eventDB db.EventDB
// 	// client is the client for filtering
// 	client client.EVMClient
// 	// cache is a cache for txHashes
// 	cache *lru.Cache
// }

// // NewContractBackfiller creates a new backfiller for a contract.
// func NewContractBackfiller(contract contracts.DeployedContract, eventDB db.EventDB, client client.EVMClient) (*ContractBackfiller, error) {
// 	// initialize the cache for the txHashes
// 	cache, err := lru.New(500)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not initialize cache: %w", err)
// 	}

// 	return &ContractBackfiller{
// 		contract: contract,
// 		eventDB:  eventDB,
// 		client:   client,
// 		cache:    cache,
// 	}, nil
// }

// // Backfill takes in a channel of logs, uses each log to get the receipt from its txHash,
// // gets all of the logs from the receipt, then stores the receipt, the logs from the
// // receipt, and the last indexed block for hte contract in the EventDB.
// //
// //nolint:gocognit, cyclop
// func (c ContractBackfiller) Backfill(ctx context.Context, givenStart uint64, endHeight uint64) error {
// 	// initialize the channel for the logs
// 	startHeight, err := c.StartHeightForBackfill(ctx, givenStart)
// 	if err != nil {
// 		return fmt.Errorf("could not get start height: %w", err)
// 	}
// 	// in the case of a failed backfill, we want to start from the last indexed block - 1
// 	if startHeight != 0 {
// 		startHeight--
// 	}
// 	logChan, errChan, doneChan := c.GetLogs(ctx, startHeight, endHeight)
// 	// start listening for logs
// 	g, ctx := errgroup.WithContext(ctx)
// 	g.Go(func() error {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return nil
// 			case log := <-logChan:
// 				// check if the txHash has already been stored in the cache
// 				if _, ok := c.cache.Get(log.TxHash); ok {
// 					continue
// 				}
// 				err = c.Store(ctx, log)
// 				if err != nil {
// 					return fmt.Errorf("could not store data: %w", err)
// 				}
// 			case err := <-errChan:
// 				return fmt.Errorf("could not get logs: %w", err)
// 			case <-doneChan:
// 				return nil
// 			}
// 		}
// 	})

// 	err = g.Wait()
// 	if err != nil {
// 		return fmt.Errorf("could not backfill contract: %w", err)
// 	}
// 	return nil
// }

// // Store stores the logs, receipts, and transactions for a tx hash.
// //
// //nolint:cyclop
// func (c ContractBackfiller) Store(ctx context.Context, log types.Log) error {
// 	receipt, err := c.client.TransactionReceipt(ctx, log.TxHash)
// 	if err != nil {
// 		return fmt.Errorf("could not get transaction receipt for txHash: %w", err)
// 	}

// 	// parallelize storing logs, receipts, and transactions
// 	g, ctx := errgroup.WithContext(ctx)
// 	if err != nil {
// 		return fmt.Errorf("could not create errgroup: %w", err)
// 	}

// 	g.Go(func() error {
// 		// get the logs from the receipt and store them in the db
// 		for _, log := range receipt.Logs {
// 			if log == nil {
// 				return fmt.Errorf("log is nil")
// 			}
// 			err = c.eventDB.StoreLog(ctx, *log, uint32(c.contract.ChainID().Uint64()))
// 			if err != nil {
// 				return fmt.Errorf("could not store log: %w", err)
// 			}
// 		}
// 		return nil
// 	})

// 	g.Go(func() error {
// 		// store the receipt in the db
// 		err = c.eventDB.StoreReceipt(ctx, *receipt, uint32(c.contract.ChainID().Uint64()))
// 		if err != nil {
// 			return fmt.Errorf("could not store receipt: %w", err)
// 		}
// 		return nil
// 	})

// 	g.Go(func() error {
// 		// store the transaction in the db
// 		txn, isPending, err := c.client.TransactionByHash(ctx, receipt.TxHash)
// 		if err != nil {
// 			return fmt.Errorf("could not get transaction by hash: %w", err)
// 		}
// 		if isPending {
// 			return fmt.Errorf("transaction is pending")
// 		}
// 		err = c.eventDB.StoreEthTx(ctx, txn, uint32(c.contract.ChainID().Uint64()))
// 		if err != nil {
// 			return fmt.Errorf("could not store transaction: %w", err)
// 		}
// 		return nil
// 	})

// 	err = g.Wait()
// 	if err != nil {
// 		return fmt.Errorf("could not store data: %w", err)
// 	}

// 	// store the last indexed block in the db
// 	err = c.eventDB.StoreLastIndexed(ctx, c.contract.Address(), uint32(c.contract.ChainID().Uint64()), receipt.BlockNumber.Uint64())
// 	if err != nil {
// 		return fmt.Errorf("could not store last indexed block: %w", err)
// 	}

// 	// store the txHash in the cache
// 	c.cache.Add(log.TxHash, true)

// 	return nil
// }

// // chunkSize is how big to make the chunks when fetching.
// const chunkSize = 1024

// // GetLogs gets all logs for the contract.
// func (c ContractBackfiller) GetLogs(ctx context.Context, startHeight, endHeight uint64) (logsChan <-chan types.Log, errsChan <-chan error, completeChan <-chan bool) {
// 	// initialize the channel
// 	logChan := make(chan types.Log)
// 	errChan := make(chan error)
// 	doneChan := make(chan bool)

// 	// start the filterer. This filters the range and sends the logs to the logChan.
// 	rangeFilter := NewRangeFilter(c.contract.Address(), c.client, big.NewInt(int64(startHeight)), big.NewInt(int64(endHeight)), chunkSize, true)
// 	g, ctx := errgroup.WithContext(ctx)
// 	g.Go(func() error {
// 		// start the range filterer, return any errors to an error channel
// 		err := rangeFilter.Start(ctx)
// 		if err != nil {
// 			return fmt.Errorf("could not filter range: %w", err)
// 		}
// 		return nil
// 	})

// 	// take the logs and put them in the log channel
// 	g.Go(func() error {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return nil
// 			case logInfos := <-rangeFilter.GetLogChan():
// 				for _, log := range logInfos.logs {
// 					logChan <- log
// 				}
// 				if rangeFilter.Done() {
// 					doneChan <- true
// 				}
// 			}
// 		}
// 	})

// 	// return errors to the channel when done filtering
// 	go func() {
// 		err := g.Wait()
// 		if err != nil {
// 			errChan <- err
// 		}
// 	}()

// 	return logChan, errChan, doneChan
// }

// // StartHeightForBackfill gets the startHeight for backfilling. This is the maximum
// // of the most recent block for the contract and the startHeight given in the config.
// func (c ContractBackfiller) StartHeightForBackfill(ctx context.Context, givenStart uint64) (startHeight uint64, err error) {
// 	lastBlock, err := c.eventDB.RetrieveLastIndexed(ctx, c.contract.Address(), uint32(c.contract.ChainID().Uint64()))
// 	if err != nil {
// 		return 0, fmt.Errorf("could not retrieve last indexed block for contract: %w", err)
// 	}

// 	if lastBlock > givenStart {
// 		return lastBlock, nil
// 	}
// 	return givenStart, nil
// }
