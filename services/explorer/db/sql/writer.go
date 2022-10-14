package sql

import (
	"context"
	"fmt"
)

// StoreEvent stores a generic event that has the proper fields set by `eventToBridgeEvent`.
func (s *Store) StoreEvent(ctx context.Context, bridgeEvent *BridgeEvent, swapEvent *SwapEvent, messageEvent *MessageEvent) error {
	if bridgeEvent != nil {
		dbTx := s.UNSAFE_DB().WithContext(ctx).Create(*bridgeEvent)
		if dbTx.Error != nil {
			return fmt.Errorf("failed to store bridge event: %w", dbTx.Error)
		}
	}
	if swapEvent != nil {
		dbTx := s.UNSAFE_DB().WithContext(ctx).Create(*swapEvent)
		if dbTx.Error != nil {
			return fmt.Errorf("failed to store swap event: %w", dbTx.Error)
		}
	}
	if messageEvent != nil {
		dbTx := s.UNSAFE_DB().WithContext(ctx).Create(*messageEvent)
		if dbTx.Error != nil {
			return fmt.Errorf("failed to store message event: %w", dbTx.Error)
		}
	}
	return nil
}
