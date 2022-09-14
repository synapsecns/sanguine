package sql

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
)

// StoreDeposit stores a deposit event.
func (s *Store) StoreDeposit(ctx context.Context, deposit bridge.DepositLog, chainID uint32) error {
	dbTx := s.DB().WithContext(ctx).
		Create(&BridgeEvent{
			ContractAddress:    deposit.GetContractAddress().String(),
			ChainID:            chainID,
			EventType:          deposit.GetEventType().Int(),
			Recipient:          deposit.GetRecipient().String(),
			BlockNumber:        deposit.GetBlockNumber(),
			TxHash:             deposit.GetTxHash().String(),
			Token:              deposit.GetToken().String(),
			Amount:             *deposit.GetAmount(),
			DestinationChainID: uint32(deposit.GetDestinationChainID().Uint64()),
		})

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store deposit: %w", dbTx.Error)
	}

	return nil
}
