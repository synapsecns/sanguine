package sql

import (
	"context"
	"fmt"
)

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
