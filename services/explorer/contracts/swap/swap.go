package swap

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
	"math/big"
)

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SwapFlashLoanTokenSwap) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the withdraw.
func (s SwapFlashLoanTokenSwap) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SwapFlashLoanTokenSwap) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SwapFlashLoanTokenSwap) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the redeem event.
func (s SwapFlashLoanTokenSwap) GetEventType() swap.EventType {
	return swap.TokenSwapEvent
}

// GetBuyer gets the buyer
func (s SwapFlashLoanTokenSwap) GetBuyer() common.Address {
	return s.Buyer
}

// GetTokensSold gets the tokens sold.
func (s SwapFlashLoanTokenSwap) GetTokensSold() *big.Int {
	return s.TokensSold
}

// GetTokensBought gets the tokens bought.
func (s SwapFlashLoanTokenSwap) GetTokensBought() *big.Int {
	return s.TokensBought
}

// GetSoldId gets the solid id
func (s SwapFlashLoanTokenSwap) GetSoldId() *big.Int {
	return s.SoldId
}

// GetBoughtId gets the bought id.
func (s SwapFlashLoanTokenSwap) GetBoughtId() *big.Int {
	return s.BoughtId
}

var _ swap.TokenSwapLog = &SwapFlashLoanTokenSwap{}
