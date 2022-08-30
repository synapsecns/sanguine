package livefill

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	lru "github.com/hashicorp/golang-lru"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/synapse-node/pkg/evm/client"
	"golang.org/x/sync/errgroup"
)

// ContractLivefiller is a livefiller that listens for logs from a contract.
type ContractLivefiller struct {
	// chainID is the chainID of the chain the contract is deployed on
	chainID uint32
	// address is the contract address to get logs from
	address string
	// eventDB is the database to store event data in
	eventDB db.EventDB
	// client is the client for filtering
	client client.EVMClient
	// cache is a cache for txHashes
	cache *lru.Cache
}

// NewContractLivefiller creates a new livefiller for a contract.
func NewContractLivefiller(chainID uint32, address string, eventDB db.EventDB, client client.EVMClient) (*ContractLivefiller, error) {
	// initialize the cache for the txHashes
	cache, err := lru.New(500)
	if err != nil {
		return nil, err
	}

	return &ContractLivefiller{
		chainID: chainID,
		address: address,
		eventDB: eventDB,
		client:  client,
		cache:   cache,
	}, nil
}

// Livefill listens for logs from a contract and stores them in the EventDB.
func (c ContractLivefiller) Livefill(ctx context.Context) error {
	// initialize the channel for the logs
	logs := make(chan types.Log)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(c.address)},
	}

	// subscribe to the logs
	sub, err := c.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return fmt.Errorf("could not subscribe to logs: %w", err)
	}

	// listen for logs
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		// backoff in case of an error
		b := &backoff.Backoff{
			Factor: 2,
			Jitter: true,
			Min:    1 * time.Second,
			Max:    30 * time.Second,
		}
		// timeout should always be 0 on the first attempt
		timeout := time.Duration(0)
		for {
			select {
			case err := <-sub.Err():
				timeout = b.Duration()
				logger.Warnf("error from subscription: %w", err)
				continue
			case log := <-logs:
				// TODO: add a notification for failure to store
				// wait the timeout (will be 0 on first attempt)
				time.Sleep(timeout)
				// check if the txHash has already been stored in the cache
				if _, ok := c.cache.Get(log.TxHash); ok {
					continue
				}
				err = c.Store(ctx, log)
				if err != nil {
					timeout = b.Duration()
					logger.Warnf("error storing log: %w", err)
					continue
				}
				// if everything works properly, restore timeout to 0
				timeout = time.Duration(0)
				b.Reset()
			}
		}
	})

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("error listening for logs: %w", err)
	}

	return nil
}

// Store stores the logs, receipts, and transactions for a tx hash.
//
//nolint:cyclop
func (c ContractLivefiller) Store(ctx context.Context, log types.Log) error {
	receipt, err := c.client.TransactionReceipt(ctx, log.TxHash)
	if err != nil {
		return fmt.Errorf("could not get transaction receipt for txHash: %w", err)
	}

	// parallelize storing logs, receipts, and transactions
	g, groupCtx := errgroup.WithContext(ctx)
	if err != nil {
		return fmt.Errorf("could not create errgroup: %w", err)
	}

	g.Go(func() error {
		// get the logs from the receipt and store them in the db
		for _, log := range receipt.Logs {
			if log == nil {
				return fmt.Errorf("log is nil")
			}
			err = c.eventDB.StoreLog(groupCtx, *log, c.chainID)
			if err != nil {
				return fmt.Errorf("could not store log: %w", err)
			}
		}
		return nil
	})

	g.Go(func() error {
		// store the receipt in the db
		err = c.eventDB.StoreReceipt(groupCtx, *receipt, c.chainID)
		if err != nil {
			return fmt.Errorf("could not store receipt: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		// store the transaction in the db
		txn, isPending, err := c.client.TransactionByHash(groupCtx, receipt.TxHash)
		if err != nil {
			return fmt.Errorf("could not get transaction by hash: %w", err)
		}
		if isPending {
			return fmt.Errorf("transaction is pending")
		}
		err = c.eventDB.StoreEthTx(groupCtx, txn, c.chainID)
		if err != nil {
			return fmt.Errorf("could not store transaction: %w", err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("could not store data: %w", err)
	}

	// store the last indexed block in the db
	err = c.eventDB.StoreLastIndexed(ctx, common.HexToAddress(c.address), c.chainID, receipt.BlockNumber.Uint64())
	if err != nil {
		return fmt.Errorf("could not store last indexed block: %w", err)
	}

	// store the txHash in the cache
	c.cache.Add(log.TxHash, true)

	return nil
}
