package sql

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
)

// StoreSwapEvent stores a deposit event.
func (s *Store) StoreSwapEvent(ctx context.Context, data swap.EventLog, chainID uint32) error {

	var provider *string
	if data.GetBuyer() != nil {
		r := data.GetProvider().String()
		provider = &r
	}

	var buyer *string
	if data.GetBuyer() != nil {
		r := data.GetBuyer().String()
		buyer = &r
	}

	dbTx := s.DB().WithContext(ctx).
		Create(&SwapEvent{
			ContractAddress: data.GetContractAddress().String(),
			ChainID:         chainID,
			EventType:       data.GetEventType().Int(),
			BlockNumber:     data.GetBlockNumber(),
			TxHash:          data.GetTxHash().String(),
			Buyer:           buyer,
			TokensSold:      data.GetTokensSold(),
			TokensBought:    data.GetTokensBought(),
			SoldID:          data.GetSoldId(),
			BoughtID:        data.GetBoughtId(),
			Provider:        provider,
			TokenAmounts:    data.GetTokenAmounts(),
			Fees:            data.GetFees(),
			Invariant:       data.GetInvariant(),
			LPTokenSupply:   data.GetLPTokenSupply(),
			LPTokenAmount:   data.GetLPTokenAmount(),
			NewAdminFee:     data.GetNewAdminFee(),
			NewSwapFee:      data.GetNewSwapFee(),
			TokenIndex:      data.GetTokenIndex(),
			Amount:          data.GetAmount(),
			AmountFee:       data.GetAmountFee(),
			ProtocolFee:     data.GetProtocolFee(),
			OldA:            data.GetOldA(),
			NewA:            data.GetNewA(),
			InitialTime:     data.GetInitialTime(),
			FutureTime:      data.GetFutureTime(),
			CurrentA:        data.GetCurrentA(),
			Time:            data.GetTime(),
		})

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store deposit: %w", dbTx.Error)
	}

	return nil
}
