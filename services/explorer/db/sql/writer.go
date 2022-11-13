package sql

import (
	"context"
	"fmt"
)

//// StoreEvent stores a generic event that has the proper fields set by `eventToBridgeEvent`.
//func (s *Store) StoreEvent(ctx context.Context, bridgeEvent *BridgeEvent, swapEvent *SwapEvent) error {
//	if bridgeEvent != nil {
//		dbTx := s.UNSAFE_DB().WithContext(ctx).Create(*bridgeEvent)
//		if dbTx.Error != nil {
//			return fmt.Errorf("failed to store bridge event: %w", dbTx.Error)
//		}
//	}
//
//	if swapEvent != nil {
//		dbTx := s.UNSAFE_DB().WithContext(ctx).Create(*swapEvent)
//		if dbTx.Error != nil {
//			return fmt.Errorf("failed to store swap event: %w", dbTx.Error)
//		}
//	}
//
//	return nil
//}

// StoreEvent stores a generic event that has the proper fields set by `eventToBridgeEvent`.
func (s *Store) StoreEvent(ctx context.Context, event interface{}) error {
	switch conv := event.(type) {
	case *BridgeEvent:
		dbTx := s.UNSAFE_DB().WithContext(ctx).Create(conv)
		if dbTx.Error != nil {
			return fmt.Errorf("failed to store bridge event: %w", dbTx.Error)
		}
	case *SwapEvent:
		dbTx := s.UNSAFE_DB().WithContext(ctx).Create(conv)
		if dbTx.Error != nil {
			return fmt.Errorf("failed to store swap event: %w", dbTx.Error)
		}
	case *MessageEvent:
		dbTx := s.UNSAFE_DB().WithContext(ctx).Create(conv)
		if dbTx.Error != nil {
			return fmt.Errorf("failed to store message event: %w", dbTx.Error)
		}
	}

	return nil
}

// StoreLastBlock stores the last block number that has been backfilled for a given chain.
func (s *Store) StoreLastBlock(ctx context.Context, chainID uint32, blockNumber uint64) error {
	entry := LastBlock{}
	dbTx := s.db.WithContext(ctx).
		Model(&LastBlock{}).
		Where(&LastBlock{
			ChainID: chainID,
		}).
		Scan(&entry)
	if dbTx.Error != nil {
		return fmt.Errorf("could not retrieve last block: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		dbTx = s.db.WithContext(ctx).
			Model(&LastBlock{}).
			Create(&LastBlock{
				ChainID:     chainID,
				BlockNumber: blockNumber,
			})
		if dbTx.Error != nil {
			return fmt.Errorf("could not store last block: %w", dbTx.Error)
		}

		return nil
	}

	if blockNumber > entry.BlockNumber {
		dbTx = s.db.WithContext(ctx).
			Model(&LastBlock{}).
			Create(&LastBlock{
				ChainID:     chainID,
				BlockNumber: blockNumber,
			})
		if dbTx.Error != nil {
			return fmt.Errorf("could not store last block: %w", dbTx.Error)
		}

		s.db.WithContext(ctx).Exec(fmt.Sprintf("ALTER TABLE last_blocks UPDATE %s=%d WHERE %s = %d ", BlockNumberFieldName, blockNumber, ChainIDFieldName, chainID))
	}

	return nil
}
