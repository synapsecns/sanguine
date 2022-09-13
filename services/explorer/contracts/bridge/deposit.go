package bridge

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
	"math/big"
)

// GetRaw gets the raw event logs from the deposit event.
func (s SynapseBridgeTokenDeposit) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetDestinationChainID gets the destination chain id from the deposit log.
func (s SynapseBridgeTokenDeposit) GetDestinationChainID() *big.Int {
	return s.ChainId
}

// GetToken gets the destination chain id from the deposit and swap log.
func (s SynapseBridgeTokenDeposit) GetToken() common.Address {
	return s.Token
}

// GetAmount gets the destination chain id from the deposit log.
func (s SynapseBridgeTokenDeposit) GetAmount() *big.Int {
	return s.Amount
}

// GetEventType gets the type of the tokend eposit event.
func (s SynapseBridgeTokenDeposit) GetEventType() bridge.EventType {
	return bridge.DepositEvent
}

// GetTxHash gets the unique identifier (txhash) for the deposit.
func (s SynapseBridgeTokenDeposit) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenDeposit) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenDeposit) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetRecipient gets the recipient of a deposit.
func (s SynapseBridgeTokenDeposit) GetRecipient() common.Address {
	return s.To
}

var _ bridge.DepositLog = &SynapseBridgeTokenDeposit{}

// GetRaw gets the raw event logs from the deposit and swap event.
func (s SynapseBridgeTokenDepositAndSwap) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetDestinationChainID gets the destination chain id from the deposit and swap log.
func (s SynapseBridgeTokenDepositAndSwap) GetDestinationChainID() *big.Int {
	return s.ChainId
}

// GetToken gets the destination chain id from the deposit and swap log.
func (s SynapseBridgeTokenDepositAndSwap) GetToken() common.Address {
	return s.Token
}

// GetAmount gets the destination chain id from the deposit and swap log.
func (s SynapseBridgeTokenDepositAndSwap) GetAmount() *big.Int {
	return s.Amount
}

// GetEventType gets the type of the deposit and swap event.
func (s SynapseBridgeTokenDepositAndSwap) GetEventType() bridge.EventType {
	return bridge.DepositAndSwapEvent
}

// GetTxHash gets the unique identifier (txhash) for the deposit.
func (s SynapseBridgeTokenDepositAndSwap) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenDepositAndSwap) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenDepositAndSwap) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetRecipient gets the recipient of a deposit.
func (s SynapseBridgeTokenDepositAndSwap) GetRecipient() common.Address {
	return s.To
}

// GetTokenIndexFrom gets the token index of the `from` token.
func (s SynapseBridgeTokenDepositAndSwap) GetTokenIndexFrom() uint8 {
	return s.TokenIndexFrom
}

// GetTokenIndexTo gets the token index of the `to` token.
func (s SynapseBridgeTokenDepositAndSwap) GetTokenIndexTo() uint8 {
	return s.TokenIndexTo
}

// GetMinDy gets the minimum amount to receive.
func (s SynapseBridgeTokenDepositAndSwap) GetMinDy() *big.Int {
	return s.MinDy
}

// GetDeadline gets the deadline for the swap.
func (s SynapseBridgeTokenDepositAndSwap) GetDeadline() *big.Int {
	return s.Deadline
}

var _ bridge.DepositAndSwapLog = &SynapseBridgeTokenDepositAndSwap{}
