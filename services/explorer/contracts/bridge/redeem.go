//nolint:golint,revive
package bridge

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
	"math/big"
)

// GetRaw gets the raw event logs from the redeem and swap event.
func (s SynapseBridgeTokenRedeemAndSwap) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetDestinationChainID gets the destination chain id from the redeem and swap log.
func (s SynapseBridgeTokenRedeemAndSwap) GetDestinationChainID() *big.Int {
	return core.CopyBigInt(s.ChainId)
}

// GetToken gets the destination chain id from the redeem and swap log.
func (s SynapseBridgeTokenRedeemAndSwap) GetToken() common.Address {
	return s.Token
}

// GetAmount gets the destination chain id from the redeem and swap log.
func (s SynapseBridgeTokenRedeemAndSwap) GetAmount() *big.Int {
	return core.CopyBigInt(s.Amount)
}

// GetEventType gets the type of the redeem and swap event.
func (s SynapseBridgeTokenRedeemAndSwap) GetEventType() bridge.EventType {
	return bridge.RedeemAndSwapEvent
}

// GetTxHash gets the unique identifier (txhash) for the redeem and swap.
func (s SynapseBridgeTokenRedeemAndSwap) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenRedeemAndSwap) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenRedeemAndSwap) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetRecipient gets the recipient address.
func (s SynapseBridgeTokenRedeemAndSwap) GetRecipient() *common.Address {
	return &s.To
}

// GetTokenIndexFrom gets the index of the `from` token.
func (s SynapseBridgeTokenRedeemAndSwap) GetTokenIndexFrom() *uint8 {
	return &s.TokenIndexFrom
}

// GetTokenIndexTo gets the index of the `to` token.
func (s SynapseBridgeTokenRedeemAndSwap) GetTokenIndexTo() *uint8 {
	return &s.TokenIndexTo
}

// GetMinDy gets the minimum amount to receive.
func (s SynapseBridgeTokenRedeemAndSwap) GetMinDy() *big.Int {
	return core.CopyBigInt(s.MinDy)
}

// GetDeadline gets the deadline for the redeem.
func (s SynapseBridgeTokenRedeemAndSwap) GetDeadline() *big.Int {
	return core.CopyBigInt(s.Deadline)
}

func (s SynapseBridgeTokenRedeemAndSwap) GetSwapTokenIndex() *uint8 {
	return nil
}

func (s SynapseBridgeTokenRedeemAndSwap) GetSwapMinAmount() *big.Int {
	return nil
}

func (s SynapseBridgeTokenRedeemAndSwap) GetSwapDeadline() *big.Int {
	return nil
}

func (s SynapseBridgeTokenRedeemAndSwap) GetRecipientBytes() *[32]byte {
	return nil
}

func (s SynapseBridgeTokenRedeemAndSwap) GetFee() *big.Int {
	return nil
}

func (s SynapseBridgeTokenRedeemAndSwap) GetKappa() *[32]byte {
	return nil
}

func (s SynapseBridgeTokenRedeemAndSwap) GetSwapSuccess() *bool {
	return nil
}

func (s SynapseBridgeTokenRedeemAndSwap) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ bridge.EventLog = &SynapseBridgeTokenRedeemAndSwap{}

// GetRaw gets the raw event logs from the redeem event.
func (s SynapseBridgeTokenRedeem) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetDestinationChainID gets the destination chain id from the redeem log.
func (s SynapseBridgeTokenRedeem) GetDestinationChainID() *big.Int {
	return core.CopyBigInt(s.ChainId)
}

// GetToken gets the destination chain id from the redeem log.
func (s SynapseBridgeTokenRedeem) GetToken() common.Address {
	return s.Token
}

// GetAmount gets the token amount.
func (s SynapseBridgeTokenRedeem) GetAmount() *big.Int {
	return core.CopyBigInt(s.Amount)
}

// GetEventType gets the type of the redeem event.
func (s SynapseBridgeTokenRedeem) GetEventType() bridge.EventType {
	return bridge.RedeemEvent
}

// GetTxHash gets the unique identifier (txhash) for the redeem.
func (s SynapseBridgeTokenRedeem) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenRedeem) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenRedeem) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetRecipient gets the recipient address.
func (s SynapseBridgeTokenRedeem) GetRecipient() *common.Address {
	return &s.To
}

func (s SynapseBridgeTokenRedeem) GetTokenIndexFrom() *uint8 {
	return nil
}

func (s SynapseBridgeTokenRedeem) GetTokenIndexTo() *uint8 {
	return nil
}

func (s SynapseBridgeTokenRedeem) GetMinDy() *big.Int {
	return nil
}

func (s SynapseBridgeTokenRedeem) GetDeadline() *big.Int {
	return nil
}

func (s SynapseBridgeTokenRedeem) GetSwapTokenIndex() *uint8 {
	return nil
}

func (s SynapseBridgeTokenRedeem) GetSwapMinAmount() *big.Int {
	return nil
}

func (s SynapseBridgeTokenRedeem) GetSwapDeadline() *big.Int {
	return nil
}

func (s SynapseBridgeTokenRedeem) GetRecipientBytes() *[32]byte {
	return nil
}

func (s SynapseBridgeTokenRedeem) GetFee() *big.Int {
	return nil
}

func (s SynapseBridgeTokenRedeem) GetKappa() *[32]byte {
	return nil
}

func (s SynapseBridgeTokenRedeem) GetSwapSuccess() *bool {
	return nil
}

func (s SynapseBridgeTokenRedeem) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ bridge.EventLog = &SynapseBridgeTokenRedeem{}

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SynapseBridgeTokenRedeemAndRemove) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetDestinationChainID gets the destination chain id from the redeem and remove log.
func (s SynapseBridgeTokenRedeemAndRemove) GetDestinationChainID() *big.Int {
	return core.CopyBigInt(s.ChainId)
}

// GetToken gets the destination chain id from the redeem and remove log.
func (s SynapseBridgeTokenRedeemAndRemove) GetToken() common.Address {
	return s.Token
}

// GetAmount gets the token amount.
func (s SynapseBridgeTokenRedeemAndRemove) GetAmount() *big.Int {
	return core.CopyBigInt(s.Amount)
}

// GetEventType gets the type of the redeem event.
func (s SynapseBridgeTokenRedeemAndRemove) GetEventType() bridge.EventType {
	return bridge.RedeemAndRemoveEvent
}

// GetTxHash gets the unique identifier (txhash) for the redeem.
func (s SynapseBridgeTokenRedeemAndRemove) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenRedeemAndRemove) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenRedeemAndRemove) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetRecipient gets the recipient address.
func (s SynapseBridgeTokenRedeemAndRemove) GetRecipient() *common.Address {
	return &s.To
}

// GetSwapTokenIndex gets the index of the swap token.
func (s SynapseBridgeTokenRedeemAndRemove) GetSwapTokenIndex() *uint8 {
	return &s.SwapTokenIndex
}

// GetSwapMinAmount gets the minimum amount to receive.
func (s SynapseBridgeTokenRedeemAndRemove) GetSwapMinAmount() *big.Int {
	return core.CopyBigInt(s.SwapMinAmount)
}

// GetSwapDeadline gets the deadline for the redeem.
func (s SynapseBridgeTokenRedeemAndRemove) GetSwapDeadline() *big.Int {
	return core.CopyBigInt(s.SwapDeadline)
}

func (s SynapseBridgeTokenRedeemAndRemove) GetTokenIndexFrom() *uint8 {
	return nil
}

func (s SynapseBridgeTokenRedeemAndRemove) GetTokenIndexTo() *uint8 {
	return nil
}

func (s SynapseBridgeTokenRedeemAndRemove) GetMinDy() *big.Int {
	return nil
}

func (s SynapseBridgeTokenRedeemAndRemove) GetDeadline() *big.Int {
	return nil
}

func (s SynapseBridgeTokenRedeemAndRemove) GetRecipientBytes() *[32]byte {
	return nil
}

func (s SynapseBridgeTokenRedeemAndRemove) GetFee() *big.Int {
	return nil
}

func (s SynapseBridgeTokenRedeemAndRemove) GetKappa() *[32]byte {
	return nil
}

func (s SynapseBridgeTokenRedeemAndRemove) GetSwapSuccess() *bool {
	return nil
}

func (s SynapseBridgeTokenRedeemAndRemove) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ bridge.EventLog = &SynapseBridgeTokenRedeemAndRemove{}

// GetRaw gets the raw event logs from the redeem event.
func (s SynapseBridgeTokenRedeemV2) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetDestinationChainID gets the destination chain id from the redeem log.
func (s SynapseBridgeTokenRedeemV2) GetDestinationChainID() *big.Int {
	return core.CopyBigInt(s.ChainId)
}

// GetToken gets the destination chain id from the redeem log.
func (s SynapseBridgeTokenRedeemV2) GetToken() common.Address {
	return s.Token
}

// GetAmount gets the token amount.
func (s SynapseBridgeTokenRedeemV2) GetAmount() *big.Int {
	return core.CopyBigInt(s.Amount)
}

// GetEventType gets the redeem event type.
func (s SynapseBridgeTokenRedeemV2) GetEventType() bridge.EventType {
	return bridge.RedeemV2Event
}

// GetTxHash gets the unique identifier (txhash) for the mint.
func (s SynapseBridgeTokenRedeemV2) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenRedeemV2) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenRedeemV2) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetRecipient gets the recipient address.
func (s SynapseBridgeTokenRedeemV2) GetRecipient() *common.Address {
	return nil
}

// GetRecipientBytes gets the recipient address.
func (s SynapseBridgeTokenRedeemV2) GetRecipientBytes() *[32]byte {
	return &s.To
}

func (s SynapseBridgeTokenRedeemV2) GetTokenIndexFrom() *uint8 {
	return nil
}

func (s SynapseBridgeTokenRedeemV2) GetTokenIndexTo() *uint8 {
	return nil
}

func (s SynapseBridgeTokenRedeemV2) GetMinDy() *big.Int {
	return nil
}

func (s SynapseBridgeTokenRedeemV2) GetDeadline() *big.Int {
	return nil
}

func (s SynapseBridgeTokenRedeemV2) GetSwapTokenIndex() *uint8 {
	return nil
}

func (s SynapseBridgeTokenRedeemV2) GetSwapMinAmount() *big.Int {
	return nil
}

func (s SynapseBridgeTokenRedeemV2) GetSwapDeadline() *big.Int {
	return nil
}

func (s SynapseBridgeTokenRedeemV2) GetFee() *big.Int {
	return nil
}

func (s SynapseBridgeTokenRedeemV2) GetKappa() *[32]byte {
	return nil
}

func (s SynapseBridgeTokenRedeemV2) GetSwapSuccess() *bool {
	return nil
}

func (s SynapseBridgeTokenRedeemV2) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ bridge.EventLog = &SynapseBridgeTokenRedeemV2{}
