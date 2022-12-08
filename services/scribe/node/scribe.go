package node

import (
	"context"
	"fmt"
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
}

// NewScribe creates a new scribe.
func NewScribe(eventDB db.EventDB, clients map[uint32][]backfill.ScribeBackend, config config.Config) (*Scribe, error) {
	scribeBackfiller, err := backfill.NewScribeBackfiller(eventDB, clients, config)
	if err != nil {
		return nil, fmt.Errorf("could not create scribe backfiller: %w", err)
	}

	return &Scribe{
		eventDB:          eventDB,
		clients:          clients,
		scribeBackfiller: scribeBackfiller,
		config:           config,
	}, nil
}

// Start starts the scribe. This works by starting a backfill and recording what the
// current block, which it will backfill to. Then, each chain will listen for new block
// heights and backfill to that height.
//
//nolint:cyclop
func (s Scribe) Start(ctx context.Context) error {
	refreshRate := s.config.RefreshRate

	if refreshRate == 0 {
		refreshRate = 1
	}

	for i := range s.config.Chains {
		chainConfig := s.config.Chains[i]

		go func() {
			b := &backoff.Backoff{
				Factor: 2,
				Jitter: true,
				Min:    1 * time.Second,
				Max:    30 * time.Second,
			}

			timeout := time.Duration(0)

			for {
				select {
				case <-ctx.Done():
					logger.Warnf("scribe for chain %d shutting down", chainConfig.ChainID)
					return
				case <-time.After(timeout):
					err := s.processRange(ctx, chainConfig.ChainID, chainConfig.RequiredConfirmations)
					if err != nil {
						timeout = b.Duration()
						logger.Warnf("could not livefill chain %d: %v", chainConfig.ChainID, err)

						continue
					}

					b.Reset()
					timeout = time.Duration(refreshRate) * time.Second
					logger.Infof("processed range for chain %d, continuing to livefill", chainConfig.ChainID)
				}
			}
		}()
	}

	return nil
}

//nolint:gocognit, cyclop
func (s Scribe) processRange(ctx context.Context, chainID uint32, requiredConfirmations uint32) error {
	newBlock, err := s.clients[chainID][0].BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("could not get current block number: %w", err)
	}

	err = s.scribeBackfiller.ChainBackfillers[chainID].Backfill(ctx, false)
	if err != nil {
		return fmt.Errorf("could not backfill: %w", err)
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
		if len(receipts) == 0 {
			logger.Errorf(" [LIVEFILL] no receipts found for block %d chain: %d, block: %d, %v", newBlock-uint64(requiredConfirmations), chainID, i, err)

			return fmt.Errorf("no receipts found for block %d", i)
		}

		// If the block hash is not the same, then the block is invalid. Otherwise, mark the block as valid.
		//nolint:nestif
		if block.Hash() != receipts[0].BlockHash {
			g, groupCtx := errgroup.WithContext(ctx)

			g.Go(func() error {
				err := s.eventDB.DeleteLogsForBlockHash(groupCtx, receipts[0].BlockHash, chainID)
				if err != nil {
					logger.Errorf(" [LIVEFILL] could not delete logs %d chain: %d, block: %d, %v", newBlock-uint64(requiredConfirmations), chainID, i, err)

					return fmt.Errorf("could not delete logs: %w", err)
				}

				return nil
			})

			g.Go(func() error {
				err := s.eventDB.DeleteReceiptsForBlockHash(groupCtx, receipts[0].BlockHash, chainID)
				if err != nil {
					logger.Errorf(" [LIVEFILL] could not delete receipts %d chain: %d, block: %d, %v", newBlock-uint64(requiredConfirmations), chainID, i, err)

					return fmt.Errorf("could not delete receipts: %w", err)
				}

				return nil
			})

			g.Go(func() error {
				err := s.eventDB.DeleteEthTxsForBlockHash(groupCtx, receipts[0].BlockHash, chainID)
				if err != nil {
					logger.Errorf(" [LIVEFILL] could not delete eth txs %d chain: %d, block: %d, %v", newBlock-uint64(requiredConfirmations), chainID, i, err)

					return fmt.Errorf("could not delete eth txs: %w", err)
				}

				return nil
			})

			if err := g.Wait(); err != nil {
				logger.Errorf(" [LIVEFILL] could not delete block %d chain: %d, block: %d, %v", newBlock-uint64(requiredConfirmations), chainID, i, err)

				return fmt.Errorf("could not delete block: %w", err)
			}

			err = s.scribeBackfiller.ChainBackfillers[chainID].Backfill(ctx, true)
			if err != nil {
				logger.Errorf(" [LIVEFILL] could not backfill %d chain: %d, block: %d, %v", newBlock-uint64(requiredConfirmations), chainID, i, err)

				return fmt.Errorf("could not backfill: %w", err)
			}
		} else {
			// Mark each receipt, log, and transaction belonging to block i to be confirmed.
			g, groupCtx := errgroup.WithContext(ctx)

			g.Go(func() error {
				err := s.eventDB.ConfirmLogsForBlockHash(groupCtx, block.Hash(), chainID)
				if err != nil {
					logger.Errorf(" [LIVEFILL] could not confirm log %d chain: %d, block: %d, %v", newBlock-uint64(requiredConfirmations), chainID, i, err)

					return fmt.Errorf("could not confirm log: %w", err)
				}

				return nil
			})
			g.Go(func() error {
				err := s.eventDB.ConfirmReceiptsForBlockHash(groupCtx, block.Hash(), chainID)
				if err != nil {
					logger.Errorf(" [LIVEFILL] could not confirm transaction %d chain: %d, block: %d, %v", newBlock-uint64(requiredConfirmations), chainID, i, err)

					return fmt.Errorf("could not confirm transaction: %w", err)
				}

				return nil
			})
			g.Go(func() error {
				err := s.eventDB.ConfirmEthTxsForBlockHash(groupCtx, block.Hash(), chainID)
				if err != nil {
					logger.Errorf(" [LIVEFILL] could not confirm transaction %d chain: %d, block: %d, %v", newBlock-uint64(requiredConfirmations), chainID, i, err)

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
