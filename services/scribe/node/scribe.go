package node

import (
	"context"
	"fmt"
	lru "github.com/hashicorp/golang-lru"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/util"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"go.opentelemetry.io/otel/attribute"
	otelMetrics "go.opentelemetry.io/otel/metric"

	"golang.org/x/sync/errgroup"

	"math/big"
	"time"
)

// Scribe is a live scribe that logs all event data.
type Scribe struct {
	// eventDB is the database to store event data in.
	eventDB db.EventDB
	// clients is a mapping of chain IDs -> clients.
	clients map[uint32][]backfill.ScribeBackend
	// scribeBackfiller is the backfiller for the scribe.
	scribeBackfiller *backfill.ScribeBackfiller
	// config is the config for the scribe.
	config config.Config
	// handler is the metrics handler for the scribe.
	handler metrics.Handler
	// reorgMeters holds a otel counter meter for reorgs for each chain
	reorgMeters map[uint32]otelMetrics.Int64Counter
}

// NewScribe creates a new scribe.
func NewScribe(eventDB db.EventDB, clients map[uint32][]backfill.ScribeBackend, config config.Config, handler metrics.Handler) (*Scribe, error) {
	scribeBackfiller, err := backfill.NewScribeBackfiller(eventDB, clients, config, handler)
	if err != nil {
		return nil, fmt.Errorf("could not create scribe backfiller: %w", err)
	}

	return &Scribe{
		eventDB:          eventDB,
		clients:          clients,
		scribeBackfiller: scribeBackfiller,
		config:           config,
		handler:          handler,
		reorgMeters:      make(map[uint32]otelMetrics.Int64Counter),
	}, nil
}

// Start starts the scribe. This works by starting a backfill and recording what the
// current block, which it will backfill to. Then, each chain will listen for new block
// heights and backfill to that height.
//
//nolint:cyclop
func (s Scribe) Start(ctx context.Context) error {
	g, groupCtx := errgroup.WithContext(ctx)

	for i := range s.config.Chains {
		chainConfig := s.config.Chains[i]
		chainID := chainConfig.ChainID
		reorgMeter, err := s.handler.Meter().NewCounter(fmt.Sprintf("scribe_reorg_meter_%d", chainID), "reorg_counter", "a reorg meter", "reorg events")
		if err != nil {
			return fmt.Errorf("error creating otel counter %w", err)
		}
		s.reorgMeters[chainID] = reorgMeter
		// Set default confirmation values
		if chainConfig.ConfirmationConfig.RequiredConfirmations == 0 ||
			chainConfig.ConfirmationConfig.ConfirmationThreshold == 0 ||
			chainConfig.ConfirmationConfig.ConfirmationRefreshRate == 0 {
			chainConfig.ConfirmationConfig = config.ConfirmationConfig{
				RequiredConfirmations:   250,
				ConfirmationThreshold:   100,
				ConfirmationRefreshRate: 5,
			}
		}
		confirmationRefreshRateTime := time.Duration(chainConfig.ConfirmationConfig.ConfirmationRefreshRate) * time.Second

		// Livefill the chains
		g.Go(func() error {
			err := s.scribeBackfiller.ChainBackfillers[chainID].Backfill(ctx, nil, true)
			if err != nil {
				return fmt.Errorf("could not backfill: %w", err)
			}
			return nil
		})

		// Check confirmations
		g.Go(func() error {
			b := &backoff.Backoff{
				Factor: 2,
				Jitter: true,
				Min:    1 * time.Second,
				Max:    10 * time.Second,
			}
			timeout := confirmationRefreshRateTime
			for {
				select {
				case <-groupCtx.Done():
					logger.Warnf("scribe for chain %d shutting down", chainConfig.ChainID)
					return nil
				case <-time.After(timeout):
					err := s.confirmBlocks(groupCtx, chainConfig)
					if err != nil {
						timeout = b.Duration()
						logger.Warnf("could not confirm blocks on chain %d, retrying: %v", chainConfig.ChainID, err)

						continue
					}

					// Set the timeout to the confirmation refresh rate.
					timeout = confirmationRefreshRateTime
					logger.Infof("processed blocks chain %d, continuing to confirm blocks", chainConfig.ChainID)
					b.Reset()
				}
			}
		})
	}
	if err := g.Wait(); err != nil {
		return fmt.Errorf("livefill failed: %w", err)
	}

	return nil
}

// confirmBlocks checks for reorgs with data stored in the scribe database
// 0. Every few seconds (depending on ConfirmationRefreshRate in the config), the confirmBlocks function is called.
// 1. First, the head block and the last confirmed block is retrieved. There is no "backfill" capability for reorgs.
// it is suggested you reindex that range you want to confirm if it is far from the head.
// 2. Block hashes for the blocks since the last confirmed block up to (latest block - the confirmation threshold)
// are batch requested.
// 3. These hashes are used to query the databases for receipts that do not have those block hashes.
// 4. The returned receipts have their blocks deleted and re-backfilled.
// 5. The entire range of blocks is then confirmed and last confirmed is updated.
//
//nolint:gocognit, cyclop
func (s Scribe) confirmBlocks(ctx context.Context, chainConfig config.ChainConfig) error {
	chainID := chainConfig.ChainID
	requiredConfirmations := chainConfig.ConfirmationConfig.RequiredConfirmations
	getBlockBatchAmount := chainConfig.GetBlockBatchAmount
	if getBlockBatchAmount == 0 {
		getBlockBatchAmount = 25
	}
	confirmationThreshold := chainConfig.ConfirmationConfig.ConfirmationThreshold

	latestBlock, err := s.clients[chainID][0].BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("could not get current block number: %w", err)
	}
	lastConfirmedBlock, err := s.eventDB.RetrieveLastConfirmedBlock(ctx, chainID)
	if err != nil {
		return fmt.Errorf("could not retrieve last confirmed block: %w", err)
	}

	// If not enough blocks have passed since the last confirmed block, the function will terminate.
	if confirmationThreshold > latestBlock-lastConfirmedBlock {
		return nil
	}

	// TODO  add option to backfill confirmations
	// To prevent getting confirmations for anything more than 1000 blocks in the past (preventing backfilling
	// confirmations AKA checking every single hash for reorg)
	if latestBlock-lastConfirmedBlock > 1000 {
		lastConfirmedBlock = latestBlock - 1000
	}

	confirmTo := latestBlock - confirmationThreshold
	confirmFrom := lastConfirmedBlock + 1
	blockHashes, err := GetBlockHashes(ctx, s.clients[chainID][0], confirmFrom, confirmTo, getBlockBatchAmount)
	if err != nil {
		return fmt.Errorf("could not get blockHashes: %w", err)
	}
	// get receipts emitted on invalid block hashes
	invalidReceipts, err := s.eventDB.RetrieveReceiptsWithStaleBlockHash(ctx, chainID, blockHashes, confirmFrom, confirmTo)
	if err != nil {
		return fmt.Errorf("could not get invalid receipts from db: %w", err)
	}

	// A cache for receipts to prevent multiple rebackfills.
	cache, err := lru.New(int(confirmTo - confirmFrom))
	if err != nil {
		return fmt.Errorf("could not access cache: %w", err)
	}

	for i := range invalidReceipts {
		receipt := invalidReceipts[i]
		cacheKey := fmt.Sprintf("%s_%d", receipt.BlockHash, receipt.BlockNumber)

		// Skip this receipt if it is part of a block that already has been re-backfilled
		if _, ok := cache.Get(cacheKey); ok {
			continue
		}
		g, groupCtx := errgroup.WithContext(ctx)

		g.Go(func() error {
			err := s.eventDB.DeleteLogsForBlockHash(groupCtx, receipt.BlockHash, chainID)

			if err != nil {
				logger.Errorf(" [LIVEFILL] could not delete logs %d chain: %d,  %v", receipt.BlockHash, chainID, err)

				return fmt.Errorf("could not delete logs: %w", err)
			}

			return nil
		})

		g.Go(func() error {
			err := s.eventDB.DeleteReceiptsForBlockHash(groupCtx, chainID, receipt.BlockHash)
			if err != nil {
				logger.Errorf(" [LIVEFILL] could not delete receipts %d chain: %d, %v", receipt.BlockHash, chainID, err)

				return fmt.Errorf("could not delete receipts: %w", err)
			}

			return nil
		})

		g.Go(func() error {
			err := s.eventDB.DeleteEthTxsForBlockHash(groupCtx, receipt.BlockHash, chainID)
			if err != nil {
				logger.Errorf(" [LIVEFILL] could not delete eth txs %d chain: %d, %v", receipt.BlockHash, chainID, err)

				return fmt.Errorf("could not delete eth txs: %w", err)
			}

			return nil
		})

		if err := g.Wait(); err != nil {
			logger.Errorf(" [LIVEFILL] could not delete block %d chain: %d, block: %d, %v", latestBlock-uint64(requiredConfirmations), chainID, i, err)

			return fmt.Errorf("could not delete block: %w", err)
		}
		blockNumber := receipt.BlockNumber.Uint64()
		err = s.scribeBackfiller.ChainBackfillers[chainID].Backfill(ctx, &blockNumber, false)
		if err != nil {
			logger.Errorf(" [LIVEFILL] could not backfill %d chain: %d, block: %d, %v", latestBlock-uint64(requiredConfirmations), chainID, i, err)

			return fmt.Errorf("could not backfill: %w", err)
		}

		cache.Add(cacheKey, true)

		// Add to meter
		s.reorgMeters[chainID].Add(ctx, 1, otelMetrics.WithAttributeSet(
			attribute.NewSet(attribute.Int64("block_number", int64(blockNumber)), attribute.Int64("chain_id", int64(chainID)))),
		)
	}

	// update items in the database as confirmed
	err = s.confirmToBlockNumber(ctx, chainID, confirmFrom, confirmTo)
	if err != nil {
		return fmt.Errorf("could not confirm items in database after backfilling %w", err)
	}
	return nil
}

// GetBlockHashes gets an array of block hashes from a range of blocks.
func GetBlockHashes(ctx context.Context, backend backfill.ScribeBackend, startBlock, endBlock uint64, getBlockBatchAmount int) ([]string, error) {
	iterator := util.NewChunkIterator(big.NewInt(int64(startBlock)), big.NewInt(int64(endBlock)), getBlockBatchAmount-1, true)
	blockRange := iterator.NextChunk()
	var hashes []string
	for blockRange != nil {
		blockHashes, err := backfill.BlockHashesInRange(ctx, backend, blockRange.StartBlock.Uint64(), blockRange.EndBlock.Uint64())
		if err != nil {
			logger.Errorf("[LIVEFILL] could not get block hashes in range %d to %d,  %v", startBlock, endBlock, err)
			// TODO potentially add a retry here
			return nil, fmt.Errorf("could not get block hashes in batch: %w", err)
		}
		itr := blockHashes.Iterator()
		for !itr.Done() {
			_, hash, _ := itr.Next()
			hashes = append(hashes, hash)
		}
		blockRange = iterator.NextChunk()
	}
	return hashes, nil
}

func (s Scribe) confirmToBlockNumber(ctx context.Context, chainID uint32, fromBlock uint64, toBlock uint64) error {
	g, groupCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		err := s.eventDB.ConfirmLogsInRange(groupCtx, fromBlock, toBlock, chainID)
		if err != nil {
			logger.Errorf(" [LIVEFILL] confirmToBlockNumber() could not confirm logs fromBlock: %d toBlock: %d chain: %d, %v", fromBlock, toBlock, chainID, err)

			return fmt.Errorf("could not confirm log: %w", err)
		}

		return nil
	})
	g.Go(func() error {
		err := s.eventDB.ConfirmReceiptsInRange(groupCtx, fromBlock, toBlock, chainID)

		if err != nil {
			logger.Errorf(" [LIVEFILL] confirmToBlockNumber() could not confirm receipts fromBlock: %d toBlock: %d chain: %d, %v", fromBlock, toBlock, chainID, err)

			return fmt.Errorf("could not confirm receipt: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		err := s.eventDB.ConfirmEthTxsInRange(groupCtx, fromBlock, toBlock, chainID)
		if err != nil {
			logger.Errorf(" [LIVEFILL] confirmToBlockNumber() could not confirm txs fromBlock: %d toBlock: %d chain: %d, %v", fromBlock, toBlock, chainID, err)

			return fmt.Errorf("could not confirm transaction: %w", err)
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		logger.Errorf(" [LIVEFILL] confirmToBlockNumber() could not confirm fromBlock: %d toBlock: %d chain: %d, %v", fromBlock, toBlock, chainID, err)

		return fmt.Errorf("could not confirm blocks: %w", err)
	}

	err := s.eventDB.StoreLastConfirmedBlock(ctx, chainID, toBlock)
	if err != nil {
		logger.Errorf(" [LIVEFILL] confirmToBlockNumber() could not store last confirmed fromBlock: %d toBlock: %d chain: %d, %v", fromBlock, toBlock, chainID, err)

		return fmt.Errorf("could not store last confirmed block: %w", err)
	}

	return nil
}
