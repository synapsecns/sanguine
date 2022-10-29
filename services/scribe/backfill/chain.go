package backfill

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"golang.org/x/sync/errgroup"
)

// ChainBackfiller is a backfiller that fetches logs for a chain. It aggregates logs
// from a slice of ContractBackfillers.
type ChainBackfiller struct {
	// chainID is the chain ID of the chain
	chainID uint32
	// eventDB is the database to store event data in
	eventDB db.EventDB
	// client is the client for filtering
	client ScribeBackend
	// contractBackfillers is the list of contract backfillers
	contractBackfillers []*ContractBackfiller
	// startHeights is a map from address -> start height
	startHeights map[string]uint64
	// minBlockHeight is the minimum block height to store block time for
	minBlockHeight uint64
	// chainConfig is the config for the backfiller
	chainConfig config.ChainConfig
}

// NewChainBackfiller creates a new backfiller for a chain.
func NewChainBackfiller(chainID uint32, eventDB db.EventDB, client ScribeBackend, chainConfig config.ChainConfig) (*ChainBackfiller, error) {
	// initialize the list of contract backfillers
	contractBackfillers := []*ContractBackfiller{}
	// initialize each contract backfiller and start heights
	startHeights := make(map[string]uint64)

	// start with max uint64
	minBlockHeight := uint64(math.MaxUint64)
	for _, contract := range chainConfig.Contracts {
		contractBackfiller, err := NewContractBackfiller(chainConfig.ChainID, contract.Address, eventDB, client)
		if err != nil {
			return nil, fmt.Errorf("could not create contract backfiller: %w", err)
		}
		contractBackfillers = append(contractBackfillers, contractBackfiller)
		startHeights[contract.Address] = contract.StartBlock

		// Compare if current minBlockHeight is greater than current StartBlock set in the yaml
		if minBlockHeight > contract.StartBlock {
			minBlockHeight = contract.StartBlock
		}
	}

	return &ChainBackfiller{
		chainID:             chainID,
		eventDB:             eventDB,
		client:              client,
		contractBackfillers: contractBackfillers,
		startHeights:        startHeights,
		minBlockHeight:      minBlockHeight,
		chainConfig:         chainConfig,
	}, nil
}

// Backfill iterates over each contract backfiller and calls Backfill concurrently on each one.
// If `onlyOneBlock` is true, the backfiller will only backfill the block at `endHeight`.
//
//nolint:gocognit,cyclop
func (c ChainBackfiller) Backfill(ctx context.Context, onlyOneBlock bool) error {
	// initialize the errgroup
	g, groupCtx := errgroup.WithContext(ctx)
	// backoff in the case of an error
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    30 * time.Second,
	}

	// Starting with 0 time out.
	timeout := time.Duration(0)

	// Init endHeight
	var endHeight uint64
	var err error

	// Retry until block height for the current chain is retrieved.
	for {
		select {
		case <-groupCtx.Done():
			return fmt.Errorf("context canceled: %w", groupCtx.Err())
		case <-time.After(timeout):
			// get the end height for the backfill
			endHeight, err = c.client.BlockNumber(groupCtx)
			if err != nil {
				timeout = b.Duration()
				logger.Warnf("could not get block number, bad connection to rpc likely: %v", err)
				continue
			}
		}
		if b.Attempt() > 1 {
			c.eventDB.
		}
		b.Reset()
		break
	}

	// iterate over each contract backfiller
	for i := range c.contractBackfillers {
		// capture func literal
		contractBackfiller := c.contractBackfillers[i]
		// get the start height for the backfill
		startHeight := c.startHeights[contractBackfiller.address]
		// call Backfill concurrently
		g.Go(func() error {
			// timeout should always be 0 on the first attempt
			timeout = time.Duration(0)
			for {
				select {
				case <-groupCtx.Done():
					return fmt.Errorf("context canceled: %w", groupCtx.Err())
				case <-time.After(timeout):
					if onlyOneBlock {
						startHeight = endHeight
					}
					err = contractBackfiller.Backfill(groupCtx, startHeight, endHeight)
					if err != nil {
						timeout = b.Duration()
						logger.Warnf("could not backfill data: %w", err)
						continue
					}
					return nil
				}
			}
		})
	}

	// Backfill the block times
	g.Go(func() error {
		// Init backoff for backfilling block times
		bBlockNum := &backoff.Backoff{
			Factor: 2,
			Jitter: true,
			Min:    1 * time.Second,
			Max:    30 * time.Second,
		}

		// timeout should always be 0 on the first attempt
		timeoutBlockNum := time.Duration(0)

		// Set the start height to the minimum block height of all contracts
		startHeight := c.minBlockHeight
		if err != nil {
			return fmt.Errorf("could not get start height for block time: %w", err)
		}

		// Start at the block before the minimum block height
		if startHeight != 0 {
			startHeight--
		}

		// Current block
		blockNum := startHeight
		for {
			select {
			case <-groupCtx.Done():
				return fmt.Errorf("context canceled: %w", groupCtx.Err())
			case <-time.After(timeoutBlockNum):
				// Check if the current block's already exists in database.
				_, err := c.eventDB.RetrieveBlockTime(ctx, c.chainID, blockNum)
				if err == nil {
					logger.Warnf("skipping storing blocktime for block %s, blocktime for this block already stored", big.NewInt(int64(blockNum)).String())
					blockNum++
					continue
				}

				// Get information on the current block for further processing.
				rawBlock, err := c.client.HeaderByNumber(ctx, big.NewInt(int64(blockNum)))
				if err != nil {
					timeoutBlockNum = bBlockNum.Duration()
					logger.Warnf("could not get block time at block %s: %v", big.NewInt(int64(blockNum)).String(), err)
					continue
				}

				// Store the block time with the block retrieved above.
				err = c.eventDB.StoreBlockTime(groupCtx, c.chainID, blockNum, rawBlock.Time)
				if err != nil {
					timeoutBlockNum = bBlockNum.Duration()
					logger.Warnf("could not store block time: %v", err)
					continue
				}

				// store the last block time
				err = c.eventDB.StoreLastBlockTime(groupCtx, c.chainID, blockNum)
				if err != nil {
					timeoutBlockNum = bBlockNum.Duration()
					logger.Warnf("could not store last block time: %v", err)
					continue
				}

				// Move on to the next block.
				blockNum++

				// Reset the backoff after successful block parse run to prevent bloated back offs.
				bBlockNum.Reset()

				// If done with the range, exit go routine.
				if blockNum > endHeight {
					return nil
				}
			}
		}
	})

	// wait for all the backfillers to finish
	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not backfill: %w", err)
	}
	return nil
}
