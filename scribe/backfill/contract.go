package backfill

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/scribe/db"
	"github.com/synapsecns/synapse-node/pkg/evm/client"
	"golang.org/x/sync/errgroup"
)

type contractBackfiller struct {
	// contract is the contract to get logs for
	contract contracts.DeployedContract
	// eventDB is the database to store event data in
	eventDB db.EventDB
	// client is the client for filtering
	client client.EVMClient
}

// NewContractBackfiller creates a new backfiller for a contract.
func NewContractBackfiller(eventDB db.EventDB, contract contracts.DeployedContract, client client.EVMClient) *contractBackfiller {
	return &contractBackfiller{
		contract: contract,
		eventDB:  eventDB,
		client:   client,
	}
}

// chunkSize is how big to make the chunks when fetching
const chunkSize = 1024

// GetLogs gets all logs for the contract.
func (c contractBackfiller) GetLogs(ctx context.Context, startHeight, endHeight uint64) (logsChan <-chan types.Log, errsChan <-chan error, completeChan <-chan bool) {
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
func (c contractBackfiller) StartHeightForBackfill(ctx context.Context, useDB bool) (startHeight uint64, err error) {
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
				return err
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
