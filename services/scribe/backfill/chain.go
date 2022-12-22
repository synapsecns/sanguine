package backfill

import (
	"context"
	"fmt"
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
}

// Used for handling logging of various context types.
type contextKey int

const (
	chainContextKey contextKey = iota
)

// NewChainBackfiller creates a new backfiller for a chain. This is done by passing through all the function parameters
// into the ChainBackfiller struct, as well as iterating through all the contracts in the chain config and creating
// ContractBackfillers for each contract.
func NewChainBackfiller(chainID uint32, eventDB db.EventDB, client []ScribeBackend, chainConfig config.ChainConfig) (*ChainBackfiller, error) {
	var contractBackfillers []*ContractBackfiller

	startHeights := make(map[string]uint64)

	if chainConfig.BlockTimeChunkCount == 0 {
		chainConfig.BlockTimeChunkCount = 40
	}

	if chainConfig.BlockTimeChunkSize == 0 {
		chainConfig.BlockTimeChunkSize = 50
	}

	if chainConfig.ContractSubChunkSize == 0 {
		chainConfig.ContractSubChunkSize = 600
	}

	if chainConfig.ContractChunkSize == 0 {
		chainConfig.ContractChunkSize = 30000
	}

	if chainConfig.StoreConcurrency == 0 {
		chainConfig.StoreConcurrency = 20
	}

	if chainConfig.StoreConcurrencyThreshold == 0 {
		chainConfig.StoreConcurrencyThreshold = 500
	}
	minBlockHeight := uint64(math.MaxUint64)

	for _, contract := range chainConfig.Contracts {
		contractBackfiller, err := NewContractBackfiller(chainConfig, contract.Address, eventDB, client)

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
// If `onlyOneBlock` is true, the backfiller will only backfill the block at `currentBlock`.
//
//nolint:gocognit,cyclop
func (c ChainBackfiller) Backfill(ctx context.Context, onlyOneBlock *uint64) error {
	var currentBlock uint64
	var err error

	// Create a new context for the chain so all chains don't halt when backfilling is completed.
	chainCtx := context.WithValue(ctx, chainContextKey, fmt.Sprintf("%d-%d", c.chainID, c.minBlockHeight))
	backfillGroup, backfillCtx := errgroup.WithContext(chainCtx)

	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    10 * time.Second,
	}

	timeout := time.Duration(0)
	startTime := time.Now()

	if onlyOneBlock == nil {
		// Retry until block height for the current chain is retrieved.
		for {
			select {
			case <-backfillCtx.Done():
				LogEvent(ErrorLevel, "Context canceled", LogData{"cid": c.chainID, "bn": currentBlock, "bd": b.Duration(), "a": b.Attempt(), "e": backfillCtx.Err(), "bt": true})

				return fmt.Errorf("%s context canceled: %w", backfillCtx.Value(chainContextKey), backfillCtx.Err())
			case <-time.After(timeout):
				currentBlock, err = c.client[0].BlockNumber(backfillCtx)

				if err != nil {
					timeout = b.Duration()
					LogEvent(InfoLevel, "Could not get block number, bad connection to rpc likely", LogData{"cid": c.chainID, "e": err.Error()})
					continue
				}
			}
			b.Reset()

			break
		}
	} else {
		c.minBlockHeight = *onlyOneBlock
		currentBlock = *onlyOneBlock + 1
	}

	for i := range c.contractBackfillers {
		contractBackfiller := c.contractBackfillers[i]
		startHeight := c.startHeights[contractBackfiller.address]

		if onlyOneBlock != nil {
			startHeight = *onlyOneBlock
			currentBlock = *onlyOneBlock
		}

		LogEvent(InfoLevel, "Starting backfilling contracts", LogData{"cid": c.chainID, "bn": currentBlock})
		backfillGroup.Go(func() error {
			timeout = time.Duration(0)
			for {
				select {
				case <-backfillCtx.Done():
					LogEvent(ErrorLevel, "Could not backfill data, context canceled", LogData{"cid": c.chainID, "ca": contractBackfiller.address, "bn": currentBlock, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt(), "e": backfillCtx.Err()})

					return fmt.Errorf("%s chain context canceled: %w", backfillCtx.Value(chainContextKey), backfillCtx.Err())
				case <-time.After(timeout):
					err = contractBackfiller.Backfill(backfillCtx, startHeight, currentBlock)

					if err != nil {
						timeout = b.Duration()
						LogEvent(WarnLevel, "Could not backfill contract, retrying", LogData{"cid": c.chainID, "ca": contractBackfiller.address, "bn": currentBlock, "sh": startHeight, "bd": b.Duration(), "a": b.Attempt(), "e": err.Error()})

						continue
					}

					return nil
				}
			}
		})
	}

	if err := backfillGroup.Wait(); err != nil {
		LogEvent(ErrorLevel, "Could not backfill with error group", LogData{"cid": c.chainID, "bn": currentBlock, "bd": b.Duration(), "a": b.Attempt(), "e": err.Error(), "bt": true})

		return fmt.Errorf("could not backfill: %w", err)
	}
	LogEvent(WarnLevel, "Finished backfilling blocktimes and contracts", LogData{"cid": c.chainID, "eh": currentBlock, "t": time.Since(startTime).Hours()})

	return nil
}
