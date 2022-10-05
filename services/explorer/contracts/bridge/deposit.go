//nolint:golint,revive
package bridge

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
	"math/big"
)

// GetRaw gets the raw event logs from the deposit event.
func (s SynapseBridgeTokenDeposit) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetDestinationChainID gets the destination chain id from the deposit log.
func (s SynapseBridgeTokenDeposit) GetDestinationChainID() *big.Int {
	return core.CopyBigInt(s.ChainId)
}

// GetToken gets the destination chain id from the deposit and swap log.
func (s SynapseBridgeTokenDeposit) GetToken() common.Address {
	return s.Token
}

// GetAmount gets the destination chain id from the deposit log.
func (s SynapseBridgeTokenDeposit) GetAmount() *big.Int {
	return core.CopyBigInt(s.Amount)
}

// GetEventType gets the type of the token deposit event.
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
func (s SynapseBridgeTokenDeposit) GetRecipient() *common.Address {
	return &s.To
}

func (s SynapseBridgeTokenDeposit) GetTokenIndexFrom() *uint8 {
	return nil
}

func (s SynapseBridgeTokenDeposit) GetTokenIndexTo() *uint8 {
	return nil
}

func (s SynapseBridgeTokenDeposit) GetMinDy() *big.Int {
	return nil
}

func (s SynapseBridgeTokenDeposit) GetDeadline() *big.Int {
	return nil
}

func (s SynapseBridgeTokenDeposit) GetSwapTokenIndex() *uint8 {
	return nil
}

func (s SynapseBridgeTokenDeposit) GetSwapMinAmount() *big.Int {
	return nil
}

func (s SynapseBridgeTokenDeposit) GetSwapDeadline() *big.Int {
	return nil
}

func (s SynapseBridgeTokenDeposit) GetRecipientBytes() *[32]byte {
	return nil
}

func (s SynapseBridgeTokenDeposit) GetFee() *big.Int {
	return nil
}

func (s SynapseBridgeTokenDeposit) GetKappa() *[32]byte {
	return nil
}

func (s SynapseBridgeTokenDeposit) GetSwapSuccess() *bool {
	return nil
}

func (s SynapseBridgeTokenDeposit) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ bridge.EventLog = &SynapseBridgeTokenDeposit{}

// GetRaw gets the raw event logs from the deposit and swap event.
func (s SynapseBridgeTokenDepositAndSwap) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetDestinationChainID gets the destination chain id from the deposit and swap log.
func (s SynapseBridgeTokenDepositAndSwap) GetDestinationChainID() *big.Int {
	return core.CopyBigInt(s.ChainId)
}

// GetToken gets the destination chain id from the deposit and swap log.
func (s SynapseBridgeTokenDepositAndSwap) GetToken() common.Address {
	return s.Token
}

// GetAmount gets the destination chain id from the deposit and swap log.
func (s SynapseBridgeTokenDepositAndSwap) GetAmount() *big.Int {
	return core.CopyBigInt(s.Amount)
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
func (s SynapseBridgeTokenDepositAndSwap) GetRecipient() *common.Address {
	return &s.To
}

// GetTokenIndexFrom gets the token index of the `from` token.
func (s SynapseBridgeTokenDepositAndSwap) GetTokenIndexFrom() *uint8 {
	return &s.TokenIndexFrom
}

// GetTokenIndexTo gets the token index of the `to` token.
func (s SynapseBridgeTokenDepositAndSwap) GetTokenIndexTo() *uint8 {
	return &s.TokenIndexTo
}

// GetMinDy gets the minimum amount to receive.
func (s SynapseBridgeTokenDepositAndSwap) GetMinDy() *big.Int {
	return core.CopyBigInt(s.MinDy)
}

// GetDeadline gets the deadline for the swap.
func (s SynapseBridgeTokenDepositAndSwap) GetDeadline() *big.Int {
	return core.CopyBigInt(s.Deadline)
}

func (s SynapseBridgeTokenDepositAndSwap) GetSwapTokenIndex() *uint8 {
	return nil
}

func (s SynapseBridgeTokenDepositAndSwap) GetSwapMinAmount() *big.Int {
	return nil
}

func (s SynapseBridgeTokenDepositAndSwap) GetSwapDeadline() *big.Int {
	return nil
}

func (s SynapseBridgeTokenDepositAndSwap) GetRecipientBytes() *[32]byte {
	return nil
}

func (s SynapseBridgeTokenDepositAndSwap) GetFee() *big.Int {
	return nil
}

func (s SynapseBridgeTokenDepositAndSwap) GetKappa() *[32]byte {
	return nil
}

func (s SynapseBridgeTokenDepositAndSwap) GetSwapSuccess() *bool {
	return nil
}

func (s SynapseBridgeTokenDepositAndSwap) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ bridge.EventLog = &SynapseBridgeTokenDepositAndSwap{}
