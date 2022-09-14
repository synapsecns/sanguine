package swap

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
	"math/big"
)

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SwapFlashLoanRampA) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the withdraw.
func (s SwapFlashLoanRampA) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SwapFlashLoanRampA) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SwapFlashLoanRampA) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the redeem event.
func (s SwapFlashLoanRampA) GetEventType() swap.EventType {
	return swap.RampAEvent
}

// GetOldA gets the old A.
func (s SwapFlashLoanRampA) GetOldA() *big.Int {
	return s.OldA
}

// GetNewA gets the new A.
func (s SwapFlashLoanRampA) GetNewA() *big.Int {
	return s.NewA
}

// GetInitialTime gets the initial time.
func (s SwapFlashLoanRampA) GetInitialTime() *big.Int {
	return s.InitialTime
}

// GetFutureTime gets the future time.
func (s SwapFlashLoanRampA) GetFutureTime() *big.Int {
	return s.FutureTime
}

var _ swap.RampALog = &SwapFlashLoanRampA{}

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SwapFlashLoanStopRampA) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the withdraw.
func (s SwapFlashLoanStopRampA) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SwapFlashLoanStopRampA) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SwapFlashLoanStopRampA) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the redeem event.
func (s SwapFlashLoanStopRampA) GetEventType() swap.EventType {
	return swap.RampAEvent
}

// GetCurrentA gets the current A.
func (s SwapFlashLoanStopRampA) GetCurrentA() *big.Int {
	return s.CurrentA
}

// GetTime gets the time.
func (s SwapFlashLoanStopRampA) GetTime() *big.Int {
	return s.Time
}

var _ swap.StopRampALog = &SwapFlashLoanStopRampA{}
