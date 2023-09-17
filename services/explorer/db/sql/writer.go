package sql

import (
	"context"
	"fmt"
)

// StoreEvent stores a generic event that has the proper fields set by `eventToBridgeEvent`.
func (s *Store) StoreEvent(ctx context.Context, event interface{}) error {
	switch conv := event.(type) {
	case *BridgeEvent:
		dbTx := s.db.WithContext(ctx).Create(conv)
		if dbTx.Error != nil {
			return fmt.Errorf("failed to store bridge event: %w", dbTx.Error)
		}
	case *SwapEvent:
		dbTx := s.db.WithContext(ctx).Create(conv)
		if dbTx.Error != nil {
			return fmt.Errorf("failed to store swap event: %w", dbTx.Error)
		}
	case *MessageBusEvent:
		dbTx := s.db.WithContext(ctx).Create(conv)
		if dbTx.Error != nil {
			return fmt.Errorf("failed to store message event: %w", dbTx.Error)
		}
	case *CCTPEvent:
		dbTx := s.db.WithContext(ctx).Create(conv)
		if dbTx.Error != nil {
			return fmt.Errorf("failed to store cctp event: %w", dbTx.Error)
		}
	}
	return nil
}

// StoreEvents stores a list of events in batches.
//
//nolint:cyclop
func (s *Store) StoreEvents(ctx context.Context, events []interface{}) error {
	var bridgeEvents []*BridgeEvent
	var swapEvents []SwapEvent
	var messageBusEvents []MessageBusEvent
	var cctpEvents []*CCTPEvent
	for _, event := range events {
		switch conv := event.(type) {
		case *BridgeEvent:
			bridgeEvents = append(bridgeEvents, conv)
		case SwapEvent:
			swapEvents = append(swapEvents, conv)
		case MessageBusEvent:
			messageBusEvents = append(messageBusEvents, conv)
		case *CCTPEvent:
			cctpEvents = append(cctpEvents, conv)
		}
	}

	// TODO: maybe switch this out with a generic
	if len(bridgeEvents) > 0 {
		dbTx := s.db.WithContext(ctx).Create(&bridgeEvents)
		if dbTx.Error != nil {
			return fmt.Errorf("failed to store bridge events: %w", dbTx.Error)
		}
	}

	if len(swapEvents) > 0 {
		dbTx := s.db.WithContext(ctx).Create(&swapEvents)
		if dbTx.Error != nil {
			return fmt.Errorf("failed to store swap events: %w", dbTx.Error)
		}
	}

	if len(messageBusEvents) > 0 {
		dbTx := s.db.WithContext(ctx).Create(&messageBusEvents)
		if dbTx.Error != nil {
			return fmt.Errorf("failed to store message event: %w", dbTx.Error)
		}
	}

	if len(cctpEvents) > 0 {
		dbTx := s.db.WithContext(ctx).Create(&cctpEvents)
		if dbTx.Error != nil {
			return fmt.Errorf("failed to store cctp event: %w", dbTx.Error)
		}
	}

	return nil
}

// StoreLastBlock stores the last block number that has been backfilled for a given chain.
func (s *Store) StoreLastBlock(ctx context.Context, chainID uint32, blockNumber uint64, contractAddress string) error {
	var lastBlock LastBlock
	dbTx := s.db.WithContext(ctx).
		Where(&LastBlock{
			ChainID:         chainID,
			ContractAddress: contractAddress,
		}).
		Order("block_number DESC").
		Limit(1).
		Find(&lastBlock)
	if dbTx.Error == nil && blockNumber > lastBlock.BlockNumber {
		lastBlock.ChainID = chainID
		lastBlock.BlockNumber = blockNumber
		lastBlock.ContractAddress = contractAddress
		dbTx = s.db.WithContext(ctx).Create(&lastBlock)
		if dbTx.Error != nil {
			return fmt.Errorf("could not store last block: %w", dbTx.Error)
		}
	}

	return nil
}

// StoreTokenIndex stores the token index data.
func (s *Store) StoreTokenIndex(ctx context.Context, chainID uint32, tokenIndex uint8, tokenAddress string, contractAddress string) error {
	entry := TokenIndex{}
	dbTx := s.db.WithContext(ctx).
		Model(&TokenIndex{}).
		Where(&TokenIndex{
			ChainID:         chainID,
			TokenIndex:      tokenIndex,
			ContractAddress: contractAddress,
		}).
		Limit(1).
		Find(&entry)
	if dbTx.Error != nil {
		return fmt.Errorf("could not retrieve last block: %w", dbTx.Error)
	}
	if tokenAddress != entry.TokenAddress {
		dbTx = s.db.WithContext(ctx).
			Model(&TokenIndex{}).
			Create(&TokenIndex{
				ChainID:         chainID,
				TokenIndex:      tokenIndex,
				ContractAddress: contractAddress,
				TokenAddress:    tokenAddress,
			})
		if dbTx.Error != nil {
			return fmt.Errorf("could not store last block: %w", dbTx.Error)
		}

		return nil
	}
	return nil
}

// StoreSwapFee stores the swap fee.
func (s *Store) StoreSwapFee(ctx context.Context, chainID uint32, blockNumber uint64, contractAddress string, fee uint64, feeType string) error {
	entry := SwapFees{}
	dbTx := s.db.WithContext(ctx).
		Model(&SwapFees{}).
		Where(&SwapFees{
			ChainID:         chainID,
			BlockNumber:     blockNumber,
			ContractAddress: contractAddress,
			FeeType:         feeType,
		}).
		Limit(1).
		Find(&entry)
	if dbTx.Error != nil {
		return fmt.Errorf("could not retrieve last block: %w", dbTx.Error)
	}
	if blockNumber != entry.BlockNumber {
		dbTx = s.db.WithContext(ctx).
			Model(&SwapFees{}).
			Create(&SwapFees{
				ChainID:         chainID,
				ContractAddress: contractAddress,
				BlockNumber:     blockNumber,
				FeeType:         feeType,
				Fee:             fee,
			})
		if dbTx.Error != nil {
			return fmt.Errorf("could not store last block: %w", dbTx.Error)
		}

		return nil
	}
	return nil
}
