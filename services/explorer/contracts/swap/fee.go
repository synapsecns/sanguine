package swap

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
	"math/big"
)

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SwapFlashLoanNewAdminFee) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the withdraw.
func (s SwapFlashLoanNewAdminFee) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SwapFlashLoanNewAdminFee) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SwapFlashLoanNewAdminFee) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the redeem event.
func (s SwapFlashLoanNewAdminFee) GetEventType() swap.EventType {
	return swap.NewAdminFeeEvent
}

// GetNewAdminFee gets the admin fee.
func (s SwapFlashLoanNewAdminFee) GetNewAdminFee() *big.Int {
	return s.NewAdminFee
}

var _ swap.NewAdminFeeLog = &SwapFlashLoanNewAdminFee{}

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SwapFlashLoanNewSwapFee) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the withdraw.
func (s SwapFlashLoanNewSwapFee) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SwapFlashLoanNewSwapFee) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SwapFlashLoanNewSwapFee) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the redeem event.
func (s SwapFlashLoanNewSwapFee) GetEventType() swap.EventType {
	return swap.NewAdminFeeEvent
}

// GetNewAdminFee gets the admin fee.
func (s SwapFlashLoanNewSwapFee) GetNewSwapFee() *big.Int {
	return s.NewSwapFee
}

var _ swap.NewSwapFeeLog = &SwapFlashLoanNewSwapFee{}
