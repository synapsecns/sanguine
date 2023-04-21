package node

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/util"
	"math/big"
	"time"

	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"golang.org/x/sync/errgroup"
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
	}, nil
}

// Start starts the scribe. This works by starting a backfill and recording what the
// current block, which it will backfill to. Then, each chain will listen for new block
// heights and backfill to that height.
//
//nolint:cyclop
func (s Scribe) Start(ctx context.Context) error {
	confirmationRefreshRate := s.config.ConfirmationRefreshRate

	if confirmationRefreshRate == 0 {
		confirmationRefreshRate = 1000
	}
	confirmationRefreshRateTime := time.Duration(confirmationRefreshRate) * time.Second
	g, groupCtx := errgroup.WithContext(ctx)

	for i := range s.config.Chains {
		chainConfig := s.config.Chains[i]
		chainID := chainConfig.ChainID

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
				Min:    30 * time.Millisecond,
				Max:    2 * time.Second,
			}
			timeout := confirmationRefreshRateTime
			for {
				select {
				case <-groupCtx.Done():
					logger.Warnf("scribe for chain %d shutting down", chainConfig.ChainID)
					return nil
				case <-time.After(timeout):
					err := s.confirmBlocks(groupCtx, chainConfig.ChainID, chainConfig.RequiredConfirmations)
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

//nolint:gocognit, cyclop
func (s Scribe) confirmBlocks(ctx context.Context, chainID uint32, requiredConfirmations uint32) error {
	logger.Infof("[LIVEFILL] start livefilling chain: %d", chainID)
	newBlock, err := s.clients[chainID][0].BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("could not get current block number: %w", err)
	}

	// TODO: backfill last confirmed block to current block - required confirmations
	const blockBatchSize = 30
	chunkStart := newBlock - uint64(requiredConfirmations)
	chunkEnd := newBlock
	iterator := util.NewChunkIterator(big.NewInt(int64(chunkStart)), big.NewInt(int64(chunkEnd)), blockBatchSize-1, true)
	for subChunk := iterator.NextChunk(); subChunk != nil; subChunk = iterator.NextChunk() {
		blockHashes, err := backfill.BlockHashesInRange(ctx, s.clients[chainID][0], subChunk.StartBlock.Uint64(), subChunk.EndBlock.Uint64())
		if err != nil {
			return fmt.Errorf("could not get block hashes: %w", err)
		}
		itr := blockHashes.Iterator()
		for !itr.Done() {
			select {
			case <-ctx.Done():
				return fmt.Errorf("could not finish checking confirmations: %w", ctx.Err())
			default:
				blockNumber, hash, _ := itr.Next()
				if hash == "" {
					logger.Errorf("[LIVEFILL] could not get block HASH EMPTY by number block: %d chain: %d, bloc %v", blockNumber, chainID, err)
					continue
				}
				receiptFilter := db.ReceiptFilter{
					ChainID:     chainID,
					BlockNumber: blockNumber,
				}

				receipts, err := s.eventDB.RetrieveReceiptsWithFilter(ctx, receiptFilter, 1)
				if err != nil {
					logger.Errorf("[LIVEFILL] could not retrieve receipts with filter %d chain: %d, block: %d, %v", newBlock-uint64(requiredConfirmations), chainID, i, err)

					return fmt.Errorf("could not retrieve receipts with filter: %w", err)
				}

				// No receipts for this block, so we can't confirm it.
				if len(receipts) == 0 {
					logger.Infof("[LIVEFILL] no receipts found for block %d chain: %d, block: %d, with filter %v", newBlock-uint64(requiredConfirmations), chainID, i, receiptFilter)

					continue
				}

				// If the block hash is not the same, then the block is invalid. Otherwise, mark the block as valid.
				//nolint:nestif
				if hash != receipts[0].BlockHash.String() {
					logger.Errorf(" [LIVEFILL] incorrect blockhash, deleting blockhash %s (block: %s) on chain %d. correct block hash: %s (block: %s)", receipts[0].BlockHash.String(), receipts[0].BlockNumber.String(), chainID, block.Hash().String(), block.Number().String())

					g, groupCtx := errgroup.WithContext(ctx)

					g.Go(func() error {
						err := s.eventDB.DeleteLogsForBlockHash(groupCtx, receipts[0].BlockHash, chainID)
						if err != nil {
							logger.Errorf(" [LIVEFILL] could not delete logs %d chain: %d,  %v", receipts[0].BlockHash, chainID, err)

							return fmt.Errorf("could not delete logs: %w", err)
						}

						return nil
					})

					g.Go(func() error {
						err := s.eventDB.DeleteReceiptsForBlockHash(groupCtx, chainID, receipts[0].BlockHash)
						if err != nil {
							logger.Errorf(" [LIVEFILL] could not delete receipts %d chain: %d, %v", receipts[0].BlockHash, chainID, err)

							return fmt.Errorf("could not delete receipts: %w", err)
						}

						return nil
					})

					g.Go(func() error {
						err := s.eventDB.DeleteEthTxsForBlockHash(groupCtx, receipts[0].BlockHash, chainID)
						if err != nil {
							logger.Errorf(" [LIVEFILL] could not delete eth txs %d chain: %d, %v", receipts[0].BlockHash, chainID, err)

							return fmt.Errorf("could not delete eth txs: %w", err)
						}

						return nil
					})

					if err := g.Wait(); err != nil {
						logger.Errorf(" [LIVEFILL] could not delete block %d chain: %d, block: %d, %v", newBlock-uint64(requiredConfirmations), chainID, i, err)

						return fmt.Errorf("could not delete block: %w", err)
					}

					blockNumber := hash
					err = s.scribeBackfiller.ChainBackfillers[chainID].Backfill(ctx, &blockNumber, false)
					if err != nil {
						logger.Errorf(" [LIVEFILL] could not backfill %d chain: %d, block: %d, %v", newBlock-uint64(requiredConfirmations), chainID, i, err)

						return fmt.Errorf("could not backfill: %w", err)
					}
				} else {
					// Mark each receipt, log, and transaction belonging to block i to be confirmed.
					g, groupCtx := errgroup.WithContext(ctx)

					g.Go(func() error {
						err := s.eventDB.ConfirmLogsForBlockHash(groupCtx, chainID, block.Hash())
						if err != nil {
							logger.Errorf(" [LIVEFILL] could not confirm log %d chain: %d,  %v", block.Hash(), chainID, err)

							return fmt.Errorf("could not confirm log: %w", err)
						}

						return nil
					})
					g.Go(func() error {
						err := s.eventDB.ConfirmReceiptsForBlockHash(groupCtx, chainID, block.Hash())
						if err != nil {
							logger.Errorf(" [LIVEFILL] could not confirm transaction %d chain: %d,  %v", block.Hash(), chainID, err)

							return fmt.Errorf("could not confirm transaction: %w", err)
						}

						return nil
					})
					g.Go(func() error {
						err := s.eventDB.ConfirmEthTxsForBlockHash(groupCtx, block.Hash(), chainID)
						if err != nil {
							logger.Errorf(" [LIVEFILL] could not confirm transaction %d chain: %d,  %v", block.Hash(), chainID, err)

							return fmt.Errorf("could not confirm transaction: %w", err)
						}

						return nil
					})

					if err := g.Wait(); err != nil {
						logger.Errorf(" [LIVEFILL] could not confirm block %d chain: %d, block: %d, %v", newBlock-uint64(requiredConfirmations), chainID, i, err)

						return fmt.Errorf("could not confirm block: %w", err)
					}
				}

				// update the last confirmed block number
				err = s.eventDB.StoreLastConfirmedBlock(ctx, chainID, i)
				if err != nil {
					logger.Errorf(" [LIVEFILL] could not store last confirmed block %d chain: %d, block: %d, %v", newBlock-uint64(requiredConfirmations), chainID, i, err)

					return fmt.Errorf("could not store last confirmed block: %w", err)
				}
				logger.Warnf("Confirmed block %d chainID: %d", i, chainID)
			}

			return nil
		}
	}
}

func getBlockByHeader(ctx context.Context) (*uint64, error) {
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

func (s Scribe) confirmToBlockNumber(ctx context.Context, blockNumber uint64, chainID uint32) error {
	g, groupCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		err := s.eventDB.ConfirmLogsInRange(groupCtx, 0, blockNumber, chainID)
		if err != nil {
			logger.Errorf(" [LIVEFILL] confirmToBlockNumber() could not confirm log %d chain: %d, %v", blockNumber, chainID, err)

			return fmt.Errorf("could not confirm log: %w", err)
		}

		return nil
	})
	g.Go(func() error {
		err := s.eventDB.ConfirmReceiptsInRange(groupCtx, 0, blockNumber, chainID)

		if err != nil {
			logger.Errorf(" [LIVEFILL] confirmToBlockNumber() could not confirm receipt %d chain: %d, %v", blockNumber, chainID, err)

			return fmt.Errorf("could not confirm receipt: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		err := s.eventDB.ConfirmEthTxsInRange(groupCtx, 0, blockNumber, chainID)
		if err != nil {
			logger.Errorf(" [LIVEFILL] confirmToBlockNumber() could not confirm tx %d chain: %d, %v", blockNumber, chainID, err)

			return fmt.Errorf("could not confirm transaction: %w", err)
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		logger.Errorf(" [LIVEFILL] confirmToBlockNumber() could not confirm blocks %d chain: %d, %v", blockNumber, chainID, err)

		return fmt.Errorf("could not confirm blocks: %w", err)
	}

	err := s.eventDB.StoreLastConfirmedBlock(ctx, chainID, blockNumber)
	if err != nil {
		logger.Errorf(" [LIVEFILL] confirmToBlockNumber() could not store last confirmed blocks %d chain: %d, %v", blockNumber, chainID, err)

		return fmt.Errorf("could not store last confirmed block: %w", err)
	}

	return nil
}
