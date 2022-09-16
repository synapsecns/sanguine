package sql

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
)

// Helper function to handle bool to uint8 conversion for clickhouse
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

// StoreEvent stores a generic event that has the proper fields set by `eventToBridgeEvent`.
func (s *Store) StoreEvent(ctx context.Context, bridgeEvent bridge.EventLog, swapEvent swap.EventLog, chainID uint32, tokenId *string) error {

	if bridgeEvent != nil {
		dbTx := s.DB().WithContext(ctx).
			Create(s.eventToBridgeEvent(bridgeEvent, chainID, tokenId))
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
func (s *Store) eventToBridgeEvent(event bridge.EventLog, chainID uint32, tokenId *string) BridgeEvent {
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
		SwapSuccess:        boolToUint8(event.GetSwapSuccess()), //clickhouse stores boolean values as an uint8
		SwapTokenIndex:     event.GetSwapTokenIndex(),
		SwapMinAmount:      event.GetSwapMinAmount(),
		SwapDeadline:       event.GetSwapDeadline(),
		TokenID:            tokenId,
	}

}

// eventToSwapEvent stores a swap event.
func (s *Store) eventToSwapEvent(event swap.EventLog, chainID uint32, tokenId *string) SwapEvent {

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
		TokenID:         nil,
	}
}
