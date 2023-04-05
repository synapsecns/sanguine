package backfill

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	"math"
	"time"

	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"golang.org/x/sync/errgroup"
)

// ChainBackfiller is a backfiller that fetches logs for a chain. It aggregates logs
// from a slice of ContractBackfillers.
type ChainBackfiller struct {
	// chainID is the chain ID of the chain.
	chainID uint32
	// eventDB is the database to store event data in.
	eventDB db.EventDB
	// client contains the clients used for backfilling.
	client []ScribeBackend
	// contractBackfillers is the list of contract backfillers.
	contractBackfillers []*ContractBackfiller
	// startHeights is a map from address -> start height.
	startHeights map[string]uint64
	// minBlockHeight is the minimum block height to store block time for.
	minBlockHeight uint64
	// chainConfig is the config for the backfiller.
	chainConfig config.ChainConfig
	// refreshRate is the rate at which the backfiller will refresh when livefilling.
	refreshRate int
	// handler is the metrics handler for the scribe.
	handler metrics.Handler
}

// Used for handling logging of various context types.
type contextKey int

const (
	chainContextKey contextKey = iota
)

// NewChainBackfiller creates a new backfiller for a chain. This is done by passing through all the function parameters
// into the ChainBackfiller struct, as well as iterating through all the contracts in the chain config and creating
// ContractBackfillers for each contract.
func NewChainBackfiller(eventDB db.EventDB, client []ScribeBackend, chainConfig config.ChainConfig, refreshRate int, handler metrics.Handler) (*ChainBackfiller, error) {
	var contractBackfillers []*ContractBackfiller

	startHeights := make(map[string]uint64)

	if chainConfig.ContractSubChunkSize == 0 {
		chainConfig.ContractSubChunkSize = 600
	}

	if chainConfig.ContractChunkSize == 0 {
		chainConfig.ContractChunkSize = 30000
	}

	if chainConfig.StoreConcurrency == 0 {
		chainConfig.StoreConcurrency = 20
	}

	if refreshRate == 0 {
		refreshRate = 1
	}

	if chainConfig.StoreConcurrencyThreshold == 0 {
		chainConfig.StoreConcurrencyThreshold = 500
	}
	minBlockHeight := uint64(math.MaxUint64)

	for _, contract := range chainConfig.Contracts {
		contractBackfiller, err := NewContractBackfiller(chainConfig, contract, eventDB, client, handler)
		if err != nil {
			return nil, fmt.Errorf("could not create contract backfiller: %w", err)
		}
		contractBackfillers = append(contractBackfillers, contractBackfiller)
		startHeights[contract.Address] = contract.StartBlock

		if minBlockHeight > contract.StartBlock {
			minBlockHeight = contract.StartBlock
		}
	}

	return &ChainBackfiller{
		chainID:             chainConfig.ChainID,
		eventDB:             eventDB,
		client:              client,
		contractBackfillers: contractBackfillers,
		startHeights:        startHeights,
		minBlockHeight:      minBlockHeight,
		chainConfig:         chainConfig,
		refreshRate:         refreshRate,
		handler:             handler,
	}, nil
}

// Backfill iterates over each contract backfiller and calls Backfill concurrently on each one.
// If `onlyOneBlock` is true, the backfiller will only backfill the block at `currentBlock`.
//
//nolint:gocognit,cyclop
func (c ChainBackfiller) Backfill(ctx context.Context, onlyOneBlock *uint64, livefill bool) error {
	// Create a new context for the chain so all chains don't halt when backfilling is completed.
	chainCtx := context.WithValue(ctx, chainContextKey, fmt.Sprintf("%d-%d", c.chainID, c.minBlockHeight))
	backfillGroup, backfillCtx := errgroup.WithContext(chainCtx)

	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    3 * time.Second,
	}

	timeout := time.Duration(0)
	startTime := time.Now()

	for i := range c.contractBackfillers {
		contractBackfiller := c.contractBackfillers[i]
		startHeight := c.startHeights[contractBackfiller.contractConfig.Address]

		LogEvent(InfoLevel, "Starting livefilling contracts", LogData{"cid": c.chainID})
		backfillGroup.Go(func() error {
			timeout = time.Duration(0)
			for {
				select {
				case <-backfillCtx.Done():
					LogEvent(ErrorLevel, "Couldn't livefill contract, context canceled", LogData{"cid": c.chainID, "ca": contractBackfiller.contractConfig.Address, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt(), "e": backfillCtx.Err()})

					return fmt.Errorf("%s chain context canceled: %w", backfillCtx.Value(chainContextKey), backfillCtx.Err())
				case <-time.After(timeout):
					var latestBlock *uint64
					var err error

					// onlyOneBlock is used for amending single blocks with a blockhash discrepancies or for testing.
					if onlyOneBlock != nil {
						startHeight = *onlyOneBlock
						latestBlock = onlyOneBlock
					} else {
						latestBlock, err = c.getLatestBlock(backfillCtx)
						if err != nil {
							return fmt.Errorf("could not get current block number while backfilling: %w", err)
						}
					}

					err = contractBackfiller.Backfill(backfillCtx, startHeight, *latestBlock)
					if err != nil {
						timeout = b.Duration()

						// If the contract has been given a specific refresh rate, then use that refresh rate for error handling.
						if contractBackfiller.contractConfig.RefreshRate > 1 {
							timeout = time.Duration(contractBackfiller.contractConfig.RefreshRate) * time.Second
						}
						LogEvent(WarnLevel, "Could not backfill contract, retrying", LogData{"cid": c.chainID, "ca": contractBackfiller.contractConfig.Address, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt(), "e": err.Error()})

						continue
					}

					if !livefill {
						return nil
					}

					timeout = time.Duration(contractBackfiller.contractConfig.RefreshRate) * time.Second
					LogEvent(InfoLevel, "Continuing to livefill contract", LogData{"t": timeout, "cid": c.chainID, "ca": contractBackfiller.contractConfig.Address, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt()})
				}
			}
		})
	}

	if err := backfillGroup.Wait(); err != nil {
		LogEvent(ErrorLevel, "Could not backfill with error group", LogData{"cid": c.chainID, "bd": b.Duration(), "a": b.Attempt(), "e": err.Error(), "bt": true})

		return fmt.Errorf("could not backfill: %w", err)
	}
	LogEvent(WarnLevel, "Finished backfilling blocktimes and contracts", LogData{"cid": c.chainID, "t": time.Since(startTime).Hours()})

	return nil
}

func (c *ChainBackfiller) getLatestBlock(ctx context.Context) (*uint64, error) {
	var currentBlock uint64
	var err error
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    10 * time.Second,
	}

	timeout := time.Duration(0)
	for {
		select {
		case <-ctx.Done():

			return nil, fmt.Errorf("%s context canceled: %w", ctx.Value(chainContextKey), ctx.Err())
		case <-time.After(timeout):
			currentBlock, err = c.client[0].BlockNumber(ctx)

			if err != nil {
				timeout = b.Duration()
				LogEvent(InfoLevel, "Could not get block number, bad connection to rpc likely", LogData{"cid": c.chainID, "e": err.Error()})
				continue
			}
		}

		return &currentBlock, nil
	}
}
