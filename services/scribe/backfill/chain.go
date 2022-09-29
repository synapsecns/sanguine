package backfill

import (
	"context"
	"fmt"
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
	minBlockHeight := uint64(0)
	for _, contract := range chainConfig.Contracts {
		contractBackfiller, err := NewContractBackfiller(chainConfig.ChainID, contract.Address, eventDB, client)
		if err != nil {
			return nil, fmt.Errorf("could not create contract backfiller: %w", err)
		}
		contractBackfillers = append(contractBackfillers, contractBackfiller)
		startHeights[contract.Address] = contract.StartBlock
		if contract.StartBlock < minBlockHeight {
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
func (c ChainBackfiller) Backfill(ctx context.Context, endHeight uint64, onlyOneBlock bool) error {
	// initialize the errgroup
	g, groupCtx := errgroup.WithContext(ctx)
	// iterate over each contract backfiller
	for _, contractBackfiller := range c.contractBackfillers {
		// capture func literal
		contractBackfiller := contractBackfiller
		// get the start height for the backfill
		startHeight := c.startHeights[contractBackfiller.address]
		// call Backfill concurrently
		g.Go(func() error {
			// backoff in the case of an error
			b := &backoff.Backoff{
				Factor: 2,
				Jitter: true,
				Min:    1 * time.Second,
				Max:    30 * time.Second,
			}
			// timeout should always be 0 on the first attempt
			timeout := time.Duration(0)
			for {
				// TODO: add a notification for failure to store
				select {
				case <-groupCtx.Done():
					return fmt.Errorf("context canceled: %w", groupCtx.Err())
				case <-time.After(timeout):
					if onlyOneBlock {
						startHeight = endHeight
					}
					err := contractBackfiller.Backfill(groupCtx, startHeight, endHeight)
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
	// backfill the block times
	g.Go(func() error {
		// backoff in the case of an error
		b := &backoff.Backoff{
			Factor: 2,
			Jitter: true,
			Min:    1 * time.Second,
			Max:    30 * time.Second,
		}
		// timeout should always be 0 on the first attempt
		timeout := time.Duration(0)
	RETRY:
		// TODO: add a notification for failure to store block time
		select {
		case <-groupCtx.Done():
			return fmt.Errorf("context canceled: %w", groupCtx.Err())
		case <-time.After(timeout):
			startHeight, err := c.startHeightForBlockTime(groupCtx)
			if err != nil {
				return fmt.Errorf("could not get start height for block time: %w", err)
			}
			if startHeight != 0 {
				startHeight--
			}
			for blockNum := startHeight; blockNum <= endHeight; blockNum++ {
				header, err := c.client.HeaderByNumber(groupCtx, big.NewInt(int64(blockNum)))
				if err != nil {
					timeout = b.Duration()
					logger.Warnf("could not get block time: %w", err)
					goto RETRY
				}
				err = c.eventDB.StoreBlockTime(groupCtx, c.chainID, blockNum, header.Time)
				if err != nil {
					timeout = b.Duration()
					logger.Warnf("could not store block time: %w", err)
					goto RETRY
				}
				// store the last block time
				err = c.eventDB.StoreLastBlockTime(groupCtx, c.chainID, blockNum)
				if err != nil {
					timeout = b.Duration()
					logger.Warnf("could not store last block time: %w", err)
					goto RETRY
				}
			}
			return nil
		}
	})

	// wait for all of the backfillers to finish
	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not backfill: %w", err)
	}

	return nil
}

// startHeightForBlockTime returns the start height for backfilling block times. This is
// the maximum of the most recent block for the chain and the minBlockHeight.
func (c ChainBackfiller) startHeightForBlockTime(ctx context.Context) (startHeight uint64, err error) {
	lastBlock, err := c.eventDB.RetrieveLastBlockTime(ctx, c.chainID)
	if err != nil {
		return 0, fmt.Errorf("could not retrieve last block time: %w", err)
	}
	if lastBlock > c.minBlockHeight {
		return lastBlock, nil
	}
	return c.minBlockHeight, nil
}
