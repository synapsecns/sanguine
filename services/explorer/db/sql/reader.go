package sql

import (
	"context"
	"fmt"
)

// ReadBlockNumberByChainID provides an easy-to-use interface to validate database
// data from a recent write event via chain id.
func (s *Store) ReadBlockNumberByChainID(ctx context.Context, eventType int8, chainID uint32) (*uint64, error) {
	// If reading a bridge event
	var blockNumber uint64
	switch eventType {
	case Bridge:
		var resp BridgeEvent
		dbTx := s.DB().WithContext(ctx).
			Find(&resp, "chain_id = ?", chainID)
		if dbTx.Error != nil {
			return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
		}
		blockNumber = resp.BlockNumber

	// If reading a swap event
	case Swap:
		var resp SwapEvent
		dbTx := s.DB().WithContext(ctx).
			Find(&resp, "chain_id = ?", chainID)
		if dbTx.Error != nil {
			return nil, fmt.Errorf("failed to store read event: %w", dbTx.Error)
		}
		blockNumber = resp.BlockNumber
	}
	return &blockNumber, nil
}
