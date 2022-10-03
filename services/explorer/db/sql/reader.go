package sql

import (
	"context"
	"fmt"
	"math/big"
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

// func (s *Store) ReadBridgeTransactions(ctx context.Context, chainID uint32, address common.Address, txHash, kappa common.Hash, includePending bool, page int, tokenAddress common.Address) BridgeEvent {
//	var bridgeEvent BridgeEvent
//	s.db.Raw(
//		`SELECT `,
//	).Find(&bridgeEvent)
//}

// GetAllChainIDs gets all chain IDs that have been used in bridge events.
func (s *Store) GetAllChainIDs(ctx context.Context) ([]uint32, error) {
	var resOut []uint32
	dbTx := s.db.WithContext(ctx).Raw(`SELECT DISTINCT chain_id FROM bridge_events`).Find(&resOut)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}
	var resIn []big.Int
	dbTx = s.db.WithContext(ctx).Raw(`SELECT DISTINCT destination_chain_id FROM bridge_events`).Find(&resIn)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}
	uniqueChainIDs := make(map[uint32]bool)
	for _, chainID := range resOut {
		uniqueChainIDs[chainID] = true
	}
	for _, chainID := range resIn {
		uniqueChainIDs[uint32(chainID.Uint64())] = true
	}
	var res []uint32
	for chainID := range uniqueChainIDs {
		res = append(res, chainID)
	}
	return res, nil
}

// BridgeCountByChainID returns the number of bridge events for a given chain ID.
func (s *Store) BridgeCountByChainID(ctx context.Context, chainID uint32, address *string, directionIn bool, firstBlock uint64) (count uint64, err error) {
	var res int64
	var addressSpecifier string
	if address != nil {
		addressSpecifier = fmt.Sprintf(" AND contract_address = '%s'", *address)
	}

	if directionIn {
		dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
			`SELECT COUNT(DISTINCT (tx_hash, event_index)) FROM bridge_events WHERE destination_chain_id = %d AND block_number >= %d%s`,
			chainID, firstBlock, addressSpecifier,
		)).Find(&res)
		if dbTx.Error != nil {
			return 0, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
		}
	} else {
		dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
			`SELECT COUNT(DISTINCT tx_hash, event_index) FROM bridge_events WHERE chain_id = %d AND block_number >= %d%s`,
			chainID, firstBlock, addressSpecifier,
		)).Find(&res)
		if dbTx.Error != nil {
			return 0, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
		}
	}
	return uint64(res), nil
}
