package swap

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
	"math/big"
)

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SwapFlashLoanFlashLoan) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the withdraw.
func (s SwapFlashLoanFlashLoan) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SwapFlashLoanFlashLoan) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SwapFlashLoanFlashLoan) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the redeem event.
func (s SwapFlashLoanFlashLoan) GetEventType() swap.EventType {
	return swap.FlashLoanLogEvent
}

// GetTokenIndex gets the token index.
func (s SwapFlashLoanFlashLoan) GetTokenIndex() uint8 {
	return s.TokenIndex
}

// GetAmount gets the amount.
func (s SwapFlashLoanFlashLoan) GetAmount() *big.Int {
	return s.Amount
}

// GetAmountFee gets the amount fee.
func (s SwapFlashLoanFlashLoan) GetAmountFee() *big.Int {
	return s.AmountFee
}

// GetProtocolFee gets the protocol fee.
func (s SwapFlashLoanFlashLoan) GetProtocolFee() *big.Int {
	return s.ProtocolFee
}

var _ swap.FlashLoanLog = &SwapFlashLoanFlashLoan{}
