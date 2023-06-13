package base

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/services/cctp-relayer/types"
)

// GetLastBlockNumber gets the last block number that had a message in the database.
func (s Store) GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error) {
	var message types.Message
	var lastBlockNumber uint64
	var numMessages int64

	// Get the total amount of messages stored in the database.
	dbTx := s.DB().WithContext(ctx).
		Model(&message).
		Where(&types.Message{OriginChainID: chainID}).
		Count(&numMessages)
	if dbTx.Error != nil {
		return 0, fmt.Errorf("failed to get number of messages: %w", dbTx.Error)
	}
	if numMessages == 0 {
		return 0, nil
	}

	dbTx = s.DB().WithContext(ctx).
		Model(&message).
		Where(fmt.Sprintf("%s = ?", OriginChainIDFieldName), chainID).
		Select(fmt.Sprintf("MAX(%s)", BlockNumberFieldName)).
		Find(&lastBlockNumber)
	if dbTx.Error != nil {
		return 0, fmt.Errorf("failed to get last block number: %w", dbTx.Error)
	}

	return lastBlockNumber, nil
}

func (s Store) StoreMessage(ctx context.Context, msg types.Message) error {
	dbTx := s.DB().WithContext(ctx).Create(&msg)
	if dbTx.Error != nil {
		return fmt.Errorf("failed to store message: %w", dbTx.Error)
	}
	return nil
}
