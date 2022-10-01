package sql

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
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

func (s *Store) BridgeCountByChainID(ctx context.Context, chainID uint32, address common.Address, directionIn bool, hours uint) (chId uint32, count uint64, err error) {
	var res int64
	if directionIn {
		dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
			`SELECT COUNT(DISTINCT tx_hash) FROM bridge_events WHERE destination_chain_id = %d`,
			chainID,
			// `SELECT count(*),
			//	argMax(destination_chain_id, insert_time) as destination_chain_id,
			//	argMax(address, insert_time) as address,
			//	FROM bridge_events WHERE destination_chain_id = ? AND address = ?;`,
		)).Find(&res)
		// Count(&res)
		if dbTx.Error != nil {
			return 0, 0, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
		}
	} else {
		dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
			`SELECT COUNT(DISTINCT tx_hash) FROM bridge_events WHERE chain_id = %d`,
			chainID,
		)).Find(&res)
		if dbTx.Error != nil {
			return 0, 0, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
		}
	}
	return chainID, uint64(res), nil
}

// SELECT count(*), argMax(tx_hash, insert_time) AS tx_hash FROM bridge_events WHERE chain_id = 1337 AND block_number >= 5;
