package sql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
	"math/big"
	"time"
)

// EventType is an enum for event types.
type EventType int8

const (
	// Bridge - SynapseBridge event.
	Bridge int8 = iota
	// Swap - SwapFlashLoan event.
	Swap
)

// BoolToUint8 is a helper function to handle bool to uint8 conversion for clickhouse.
func BoolToUint8(input *bool) *uint8 {
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
	var recipient sql.NullString
	if event.GetRecipient() != nil {
		recipient.Valid = true
		recipient.String = event.GetRecipient().String()
	} else {
		recipient.Valid = false
	}
	var recipientBytes sql.NullString
	if event.GetRecipientBytes() != nil {
		recipientBytes.Valid = true
		recipientBytes.String = common.Bytes2Hex(event.GetRecipientBytes()[:])
	} else {
		recipientBytes.Valid = false
	}
	var destinationChainID *big.Int
	if event.GetDestinationChainID() != nil {
		destinationChainID = big.NewInt(int64(event.GetDestinationChainID().Uint64()))
	}
	var tokenIndexFrom *big.Int
	if event.GetTokenIndexFrom() != nil {
		tokenIndexFrom = big.NewInt(int64(*event.GetTokenIndexFrom()))
	}
	var tokenIndexTo *big.Int
	if event.GetTokenIndexTo() != nil {
		tokenIndexTo = big.NewInt(int64(*event.GetTokenIndexTo()))
	}
	var swapSuccess *big.Int
	if event.GetSwapSuccess() != nil {
		swapSuccess = big.NewInt(int64(*BoolToUint8(event.GetSwapSuccess())))
	}
	var swapTokenIndex *big.Int
	if event.GetSwapTokenIndex() != nil {
		swapTokenIndex = big.NewInt(int64(*event.GetSwapTokenIndex()))
	}
	var kappa sql.NullString
	if event.GetKappa() != nil {
		kappa.Valid = true
		kappa.String = common.Bytes2Hex(event.GetKappa()[:])
	} else {
		kappa.Valid = false
	}
	var tokID sql.NullString
	if tokenID != nil {
		tokID.Valid = true
		tokID.String = *tokenID
	} else {
		tokID.Valid = false
	}

	return BridgeEvent{
		InsertTime:         uint64(time.Now().UnixNano()),
		ContractAddress:    event.GetContractAddress().String(),
		ChainID:            chainID,
		EventType:          event.GetEventType().Int(),
		BlockNumber:        event.GetBlockNumber(),
		TxHash:             event.GetTxHash().String(),
		Amount:             event.GetAmount(),
		EventIndex:         event.GetEventIndex(),
		Recipient:          recipient,
		RecipientBytes:     recipientBytes,
		DestinationChainID: destinationChainID,
		Token:              event.GetToken().String(),
		Fee:                event.GetFee(),
		Kappa:              kappa,
		TokenIndexFrom:     tokenIndexFrom,
		TokenIndexTo:       tokenIndexTo,
		MinDy:              event.GetMinDy(),
		Deadline:           event.GetDeadline(),
		SwapSuccess:        swapSuccess,
		SwapTokenIndex:     swapTokenIndex,
		SwapMinAmount:      event.GetSwapMinAmount(),
		SwapDeadline:       event.GetSwapDeadline(),
		TokenID:            tokID,
	}
}

// eventToSwapEvent stores a swap event.
func (s *Store) eventToSwapEvent(event swap.EventLog, chainID uint32, tokenID *string) SwapEvent {
	var buyer sql.NullString
	if event.GetBuyer() != nil {
		buyer.Valid = true
		buyer.String = event.GetBuyer().String()
	} else {
		buyer.Valid = false
	}
	var provider sql.NullString
	if event.GetProvider() != nil {
		provider.Valid = true
		provider.String = event.GetProvider().String()
	} else {
		provider.Valid = false
	}
	var tokenIndex *big.Int
	if event.GetTokenIndex() != nil {
		tokenIndex = big.NewInt(int64(*event.GetTokenIndex()))
	}
	var receiver sql.NullString
	if event.GetReceiver() != nil {
		receiver.Valid = true
		receiver.String = event.GetReceiver().String()
	} else {
		receiver.Valid = false
	}
	var tokID sql.NullString
	if tokenID != nil {
		tokID.Valid = true
		tokID.String = *tokenID
	} else {
		tokID.Valid = false
	}

	return SwapEvent{
		InsertTime:      uint64(time.Now().UnixNano()),
		ContractAddress: event.GetContractAddress().String(),
		ChainID:         chainID,
		EventType:       event.GetEventType().Int(),
		BlockNumber:     event.GetBlockNumber(),
		TxHash:          event.GetTxHash().String(),
		EventIndex:      event.GetEventIndex(),
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
		TokenIndex:      tokenIndex,
		Amount:          event.GetAmount(),
		AmountFee:       event.GetAmountFee(),
		ProtocolFee:     event.GetProtocolFee(),
		OldA:            event.GetOldA(),
		NewA:            event.GetNewA(),
		InitialTime:     event.GetInitialTime(),
		FutureTime:      event.GetFutureTime(),
		CurrentA:        event.GetCurrentA(),
		Time:            event.GetTime(),
		Receiver:        receiver,
		TokenID:         tokID,
	}
}
