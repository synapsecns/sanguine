package sql

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
)

// StoreEvent stores a generic event that has the proper fields set by `eventToBridgeEvent`.
func (s *Store) StoreEvent(ctx context.Context, event bridge.EventLog, chainID uint32) error {
	dbTx := s.DB().WithContext(ctx).
		Create(s.eventToBridgeEvent(event, chainID))

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store bridge event: %w", dbTx.Error)
	}

	return nil
}

func (s *Store) eventToBridgeEvent(event bridge.EventLog, chainID uint32) BridgeEvent {
	var recipient *string
	if event.GetRecipient() != nil {
		r := event.GetRecipient().String()
		recipient = &r
	}
	var destinationChainID *uint32
	if event.GetDestinationChainID() != nil {
		d := uint32(event.GetDestinationChainID().Uint64())
		destinationChainID = &d
	}
	var kappa *string
	if event.GetKappa() != nil {
		k := event.GetKappa()
		kHash := common.BytesToHash(k[:]).String()
		kappa = &kHash
	}
	return BridgeEvent{
		ContractAddress:    event.GetContractAddress().String(),
		ChainID:            chainID,
		EventType:          event.GetEventType().Int(),
		BlockNumber:        event.GetBlockNumber(),
		TxHash:             event.GetTxHash().String(),
		Amount:             event.GetAmount(),
		Recipient:          recipient,
		DestinationChainID: destinationChainID,
		Token:              event.GetToken().String(),
		Fee:                event.GetFee(),
		Kappa:              kappa,
		TokenIndexFrom:     event.GetTokenIndexFrom(),
		TokenIndexTo:       event.GetTokenIndexTo(),
		MinDy:              event.GetMinDy(),
		Deadline:           event.GetDeadline(),
		SwapSuccess:        event.GetSwapSuccess(),
		SwapTokenIndex:     event.GetSwapTokenIndex(),
		SwapMinAmount:      event.GetSwapMinAmount(),
		SwapDeadline:       event.GetSwapDeadline(),
	}
}
