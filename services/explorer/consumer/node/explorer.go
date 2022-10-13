package node

import (
	"context"
	"fmt"
	"github.com/nats-io/nats-server/v2/logger"
	"math/big"
	"time"

	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"golang.org/x/sync/errgroup"
)

// Explorer is a live scribe that logs all event data.
type Explorer struct {
	// consumerDB is the database to store event data in
	consumerDB db.ConsumerDB
	// clients is a mapping of chain IDs -> clients
	clients map[uint32]backfill.ScribeBackend
	// scribeBackfiller is the backfiller for the scribe
	scribeBackfiller *backfill.ScribeBackfiller
	// config is the config for the scribe
	config config.Config
}

// NewScribe creates a new scribe.
func NewScribe(eventDB db.EventDB, clients map[uint32]backfill.ScribeBackend, config config.Config) (*Explorer, error) {
	// initialize the scribe backfiller
	scribeBackfiller, err := backfill.NewScribeBackfiller(eventDB, clients, config)
	if err != nil {
		return nil, fmt.Errorf("could not create scribe backfiller: %w", err)
	}

	return &Explorer{
		consumerDB:       eventDB,
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
func (s Explorer) Start(ctx context.Context) error {
	refreshRate := s.config.RefreshRate
	if refreshRate == 0 {
		refreshRate = 1
	}
	// backfill each chain
	g, groupCtx := errgroup.WithContext(ctx)
	for _, chainConfig := range s.config.Chains {
		// capture func literal
		chainConfig := chainConfig
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
				case <-groupCtx.Done():
					return fmt.Errorf("context finished: %w", groupCtx.Err())
				case <-time.After(timeout):
					err := s.processRange(groupCtx, chainConfig.ChainID, chainConfig.RequiredConfirmations)
					if err != nil {
						timeout = b.Duration()
						logger.Warnf("could not get current block number: %v", err)
						continue
					}
					b.Reset()
					timeout = time.Duration(refreshRate) * time.Second
				}
			}
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not backfill: %w", err)
	}

	return nil
}

//nolint:gocognit, cyclop
func (s Explorer) processRange(ctx context.Context, chainID uint32, requiredConfirmations uint32) error {
	newBlock, err := s.clients[chainID].BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("could not get current block number: %w", err)
	}

	err = s.scribeBackfiller.ChainBackfillers[chainID].Backfill(ctx, false)
	if err != nil {
		return fmt.Errorf("could not backfill: %w", err)
	}

	// in the range (last confirmed block number, current block number - required confirmations],
	// check the validity of the blocks, and modify the database accordingly
	lastBlockNumber, err := s.consumerDB.RetrieveLastConfirmedBlock(ctx, chainID)
	if err != nil {
		return fmt.Errorf("could not retrieve last confirmed block: %w", err)
	}

	// if the last block number is 0 and current block - required confirmations is greater than 0,
	// then set all blocks up to current block - required confirmations to confirmed
	if lastBlockNumber == 0 && newBlock > uint64(requiredConfirmations) {
		err := s.confirmToBlockNumber(ctx, newBlock-uint64(requiredConfirmations), chainID)
		if err != nil {
			return fmt.Errorf("could not confirm blocks: %w", err)
		}
		lastBlockNumber, err = s.consumerDB.RetrieveLastConfirmedBlock(ctx, chainID)
		if err != nil {
			return fmt.Errorf("could not retrieve last confirmed block: %w", err)
		}
	}

	for i := lastBlockNumber + 1; i <= newBlock-uint64(requiredConfirmations); i++ {
		// check the validity of the block
		block, err := s.clients[chainID].BlockByNumber(ctx, big.NewInt(int64(i)))
		if err != nil {
			return fmt.Errorf("could not get block by number: %w", err)
		}
		// get the block hash of the stored block, using a receipt
		receiptFilter := db.ReceiptFilter{
			ChainID:     chainID,
			BlockNumber: i,
		}
		receipts, err := s.consumerDB.RetrieveReceiptsWithFilter(ctx, receiptFilter, 1)
		if err != nil {
			return fmt.Errorf("could not retrieve receipts with filter: %w", err)
		}
		if len(receipts) == 0 {
			return fmt.Errorf("no receipts found for block %d", i)
		}

		// if the block hash is not the same, then the block is invalid. otherwise, mark the block as valid
		//nolint:nestif
		if block.Hash() != receipts[0].BlockHash {
			g, groupCtx := errgroup.WithContext(ctx)
			g.Go(func() error {
				err := s.consumerDB.DeleteLogsForBlockHash(groupCtx, receipts[0].BlockHash, chainID)
				if err != nil {
					return fmt.Errorf("could not delete logs: %w", err)
				}
				return nil
			})
			g.Go(func() error {
				err := s.consumerDB.DeleteReceiptsForBlockHash(groupCtx, receipts[0].BlockHash, chainID)
				if err != nil {
					return fmt.Errorf("could not delete receipts: %w", err)
				}
				return nil
			})
			g.Go(func() error {
				err := s.consumerDB.DeleteEthTxsForBlockHash(groupCtx, receipts[0].BlockHash, chainID)
				if err != nil {
					return fmt.Errorf("could not delete eth txs: %w", err)
				}
				return nil
			})
			if err := g.Wait(); err != nil {
				return fmt.Errorf("could not delete block: %w", err)
			}

			// get the data for the block and backfill
			err = s.scribeBackfiller.ChainBackfillers[chainID].Backfill(ctx, true)
			if err != nil {
				return fmt.Errorf("could not backfill: %w", err)
			}
		} else {
			g, groupCtx := errgroup.WithContext(ctx)
			// mark each receipt, log, and transaction belonging to block i to be confirmed
			g.Go(func() error {
				err := s.consumerDB.ConfirmLogsForBlockHash(groupCtx, block.Hash(), chainID)
				if err != nil {
					return fmt.Errorf("could not confirm log: %w", err)
				}
				return nil
			})
			g.Go(func() error {
				err := s.consumerDB.ConfirmReceiptsForBlockHash(groupCtx, block.Hash(), chainID)
				if err != nil {
					return fmt.Errorf("could not confirm transaction: %w", err)
				}
				return nil
			})
			g.Go(func() error {
				err := s.consumerDB.ConfirmEthTxsForBlockHash(groupCtx, block.Hash(), chainID)
				if err != nil {
					return fmt.Errorf("could not confirm transaction: %w", err)
				}
				return nil
			})

			if err := g.Wait(); err != nil {
				return fmt.Errorf("could not confirm block: %w", err)
			}
		}

		// update the last confirmed block number
		err = s.consumerDB.StoreLastConfirmedBlock(ctx, chainID, i)
		if err != nil {
			return fmt.Errorf("could not store last confirmed block: %w", err)
		}
	}

	return nil
}

func (s Explorer) confirmToBlockNumber(ctx context.Context, blockNumber uint64, chainID uint32) error {
	g, groupCtx := errgroup.WithContext(ctx)
	// mark each receipt, log, and transaction to confirmed up to block `blockNumber`
	g.Go(func() error {
		err := s.consumerDB.ConfirmLogsInRange(groupCtx, 0, blockNumber, chainID)
		if err != nil {
			return fmt.Errorf("could not confirm log: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		err := s.consumerDB.ConfirmReceiptsInRange(groupCtx, 0, blockNumber, chainID)
		if err != nil {
			return fmt.Errorf("could not confirm transaction: %w", err)
		}
		return nil
	})
	g.Go(func() error {
		err := s.consumerDB.ConfirmEthTxsInRange(groupCtx, 0, blockNumber, chainID)
		if err != nil {
			return fmt.Errorf("could not confirm transaction: %w", err)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not confirm block: %w", err)
	}

	// update the last confirmed block number
	err := s.consumerDB.StoreLastConfirmedBlock(ctx, chainID, blockNumber)
	if err != nil {
		return fmt.Errorf("could not store last confirmed block: %w", err)
	}

	return nil
}
