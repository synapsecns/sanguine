package sql

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
)

// EventType is an enum for event types.
type EventType int8

const (
	// Bridge - SynapseBridge event.
	Bridge int8 = 0
	// Swap - SwapFlashLoan event.
	Swap int8 = iota
)

// Helper function to handle bool to uint8 conversion for clickhouse.
func boolToUint8(input *bool) *uint8 {
	if input == nil {
		return nil
	}
	if *input {
		one := uint8(1)
		return &one
	}
	zero := uint8(0)
	return &zero
}

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

// StoreEvent stores a generic event that has the proper fields set by `eventToBridgeEvent`.
func (s *Store) StoreEvent(ctx context.Context, bridgeEvent bridge.EventLog, swapEvent swap.EventLog, chainID uint32, tokenID *string) error {
	if bridgeEvent != nil {
		dbTx := s.DB().WithContext(ctx).
			Create(s.eventToBridgeEvent(bridgeEvent, chainID, tokenID))
		if dbTx.Error != nil {
			return fmt.Errorf("failed to store bridge event: %w", dbTx.Error)
		}
	}
	if swapEvent != nil {
		dbTx := s.DB().WithContext(ctx).
			Create(s.eventToSwapEvent(swapEvent, chainID, nil))
		if dbTx.Error != nil {
			return fmt.Errorf("failed to store swap event: %w", dbTx.Error)
		}
	}

	return nil
}

// eventToBridgeEvent stores a bridge event.
func (s *Store) eventToBridgeEvent(event bridge.EventLog, chainID uint32, tokenID *string) BridgeEvent {
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
		hash := common.BytesToHash(k[:]).String()
		kappa = &hash
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
		SwapSuccess:        boolToUint8(event.GetSwapSuccess()), // clickhouse stores boolean values as an uint8
		SwapTokenIndex:     event.GetSwapTokenIndex(),
		SwapMinAmount:      event.GetSwapMinAmount(),
		SwapDeadline:       event.GetSwapDeadline(),
		TokenID:            tokenID,
	}
}

// eventToSwapEvent stores a swap event.
func (s *Store) eventToSwapEvent(event swap.EventLog, chainID uint32, tokenID *string) SwapEvent {
	var provider *string
	if event.GetProvider() != nil {
		r := event.GetProvider().String()
		provider = &r
	}

	var buyer *string
	if event.GetBuyer() != nil {
		r := event.GetBuyer().String()
		buyer = &r
	}

	return SwapEvent{
		ContractAddress: event.GetContractAddress().String(),
		ChainID:         chainID,
		EventType:       event.GetEventType().Int(),
		BlockNumber:     event.GetBlockNumber(),
		TxHash:          event.GetTxHash().String(),
		Buyer:           buyer,
		TokensSold:      event.GetTokensSold(),
		TokensBought:    event.GetTokensBought(),
		SoldID:          event.GetSoldId(),
		BoughtID:        event.GetBoughtId(),
		Provider:        provider,
		TokenAmounts:    event.GetTokenAmounts(),
		Fees:            event.GetFees(),
		Invariant:       event.GetInvariant(),
		LPTokenSupply:   event.GetLPTokenSupply(),
		LPTokenAmount:   event.GetLPTokenAmount(),
		NewAdminFee:     event.GetNewAdminFee(),
		NewSwapFee:      event.GetNewSwapFee(),
		TokenIndex:      event.GetTokenIndex(),
		Amount:          event.GetAmount(),
		AmountFee:       event.GetAmountFee(),
		ProtocolFee:     event.GetProtocolFee(),
		OldA:            event.GetOldA(),
		NewA:            event.GetNewA(),
		InitialTime:     event.GetInitialTime(),
		FutureTime:      event.GetFutureTime(),
		CurrentA:        event.GetCurrentA(),
		Time:            event.GetTime(),
		TokenID:         tokenID,
	}
}
