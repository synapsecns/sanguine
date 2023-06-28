package node

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
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

// checkFinality checks if the block is final on the chain.
// and deletes irrelevant blocks.
const checkFinality = false

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
			if !checkFinality {
				return nil
			}
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

// TODO: Has issues with last confirmed data. Needs to be fixed.
//
//nolint:gocognit, cyclop
func (s Scribe) confirmBlocks(ctx context.Context, chainID uint32, requiredConfirmations uint32) error {
	logger.Infof("[LIVEFILL] start livefilling chain: %d", chainID)
	newBlock, err := s.clients[chainID][0].BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("could not get current block number: %w", err)
	}
	// In the range (last confirmed block number, current block number - required confirmations],
	// check the validity of the blocks, and modify the database accordingly.
	lastBlockNumber, err := s.eventDB.RetrieveLastConfirmedBlock(ctx, chainID)
	if err != nil {
		return fmt.Errorf("could not retrieve last confirmed block: %w", err)
	}

	// If the last block number is 0 and current block - required confirmations is greater than 0,
	// then set all blocks up to current block - required confirmations to confirmed.
	if lastBlockNumber == 0 && newBlock > uint64(requiredConfirmations) {
		err := s.confirmToBlockNumber(ctx, newBlock-uint64(requiredConfirmations), chainID)
		if err != nil {
			logger.Errorf("[LIVEFILL] could not confirm to block number %d chain: %d: %v", newBlock-uint64(requiredConfirmations), chainID, err)
			return fmt.Errorf("could not confirm blocks: %w", err)
		}

		lastBlockNumber, err = s.eventDB.RetrieveLastConfirmedBlock(ctx, chainID)
		if err != nil {
			logger.Errorf("[LIVEFILL] could not retrieve last confirmed block %d chain: %d: %v", newBlock-uint64(requiredConfirmations), chainID, err)

			return fmt.Errorf("could not retrieve last confirmed block: %w", err)
		}
	}

	for i := lastBlockNumber + 1; i <= newBlock-uint64(requiredConfirmations); i++ {
		block, err := s.clients[chainID][0].BlockByNumber(ctx, big.NewInt(int64(i)))
		if err != nil {
			logger.Errorf("[LIVEFILL] could not get block by number %d chain: %d, block: %d, %v", newBlock-uint64(requiredConfirmations), chainID, i, err)

			return fmt.Errorf("could not get block by number: %w", err)
		}

		receiptFilter := db.ReceiptFilter{
			ChainID:     chainID,
			BlockNumber: i,
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
		if block.Hash() != receipts[0].BlockHash {
			logger.Errorf(" [LIVEFILL] incorrect blockhash, deleting blockhash %s on chain %d. correct block hash: %s", receipts[0].BlockHash.String(), chainID, block.Hash().String())

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

			blockNumber := block.Number().Uint64()
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
