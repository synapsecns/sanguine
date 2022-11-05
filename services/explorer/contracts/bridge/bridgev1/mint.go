//nolint:golint,revive,dupl
package bridgev1

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
	"math/big"
)

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SynapseBridgeTokenMint) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetToken gets the destination chain id from the mint.
func (s SynapseBridgeTokenMint) GetToken() common.Address {
	return s.Token
}

// GetAmount gets the token amount.
func (s SynapseBridgeTokenMint) GetAmount() *big.Int {
	return core.CopyBigInt(s.Amount)
}

// GetEventType gets the type of the redeem event.
func (s SynapseBridgeTokenMint) GetEventType() bridge.EventType {
	return bridge.MintEvent
}

// GetTxHash gets the unique identifier (txhash) for the mint.
func (s SynapseBridgeTokenMint) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenMint) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenMint) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetFee gets the fee for the token mint.
func (s SynapseBridgeTokenMint) GetFee() *big.Int {
	return core.CopyBigInt(s.Fee)
}

// GetKappa gets the gappa for the token mint.
func (s SynapseBridgeTokenMint) GetKappa() *[32]byte {
	return &s.Kappa
}

// GetRecipient gets the recipient for the mint.
func (s SynapseBridgeTokenMint) GetRecipient() *common.Address {
	return &s.To
}

func (s SynapseBridgeTokenMint) GetDestinationChainID() *big.Int {
	return nil
}

func (s SynapseBridgeTokenMint) GetTokenIndexFrom() *uint8 {
	return nil
}

func (s SynapseBridgeTokenMint) GetTokenIndexTo() *uint8 {
	return nil
}

func (s SynapseBridgeTokenMint) GetMinDy() *big.Int {
	return nil
}

func (s SynapseBridgeTokenMint) GetDeadline() *big.Int {
	return nil
}

func (s SynapseBridgeTokenMint) GetSwapTokenIndex() *uint8 {
	return nil
}

func (s SynapseBridgeTokenMint) GetSwapMinAmount() *big.Int {
	return nil
}

func (s SynapseBridgeTokenMint) GetSwapDeadline() *big.Int {
	return nil
}

func (s SynapseBridgeTokenMint) GetRecipientBytes() *[32]byte {
	return nil
}

func (s SynapseBridgeTokenMint) GetSwapSuccess() *bool {
	return nil
}

func (s SynapseBridgeTokenMint) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ bridge.EventLog = &SynapseBridgeTokenMint{}

// GetToken gets the token for the mint and remove operation.
func (s SynapseBridgeTokenMintAndSwap) GetToken() common.Address {
	return s.Token
}

// GetAmount gets the amount fo the mint.
func (s SynapseBridgeTokenMintAndSwap) GetAmount() *big.Int {
	return core.CopyBigInt(s.Amount)
}

// GetEventType gets the mint and remove event type.
func (s SynapseBridgeTokenMintAndSwap) GetEventType() bridge.EventType {
	return bridge.MintAndSwapEvent
}

// GetRaw gets the raw logs.
func (s SynapseBridgeTokenMintAndSwap) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetFee gets the fee for the bridge.
func (s SynapseBridgeTokenMintAndSwap) GetFee() *big.Int {
	return core.CopyBigInt(s.Fee)
}

// GetKappa gets the kappa value for the event.
func (s SynapseBridgeTokenMintAndSwap) GetKappa() *[32]byte {
	return &s.Kappa
}

// GetTxHash gets the unique identifier (txhash) for the mint.
func (s SynapseBridgeTokenMintAndSwap) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenMintAndSwap) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenMintAndSwap) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetRecipient gets the recipient for the mint and swap.
func (s SynapseBridgeTokenMintAndSwap) GetRecipient() *common.Address {
	return &s.To
}

// GetTokenIndexFrom gets the index of the `from` token.
func (s SynapseBridgeTokenMintAndSwap) GetTokenIndexFrom() *uint8 {
	return &s.TokenIndexFrom
}

// GetTokenIndexTo gets the index of the `to` token.
func (s SynapseBridgeTokenMintAndSwap) GetTokenIndexTo() *uint8 {
	return &s.TokenIndexTo
}

// GetMinDy gets the minimum amount to receive.
func (s SynapseBridgeTokenMintAndSwap) GetMinDy() *big.Int {
	return core.CopyBigInt(s.MinDy)
}

// GetDeadline gets the deadline for the swap.
func (s SynapseBridgeTokenMintAndSwap) GetDeadline() *big.Int {
	return core.CopyBigInt(s.Deadline)
}

// GetSwapSuccess gets if the swap was successful.
func (s SynapseBridgeTokenMintAndSwap) GetSwapSuccess() *bool {
	return &s.SwapSuccess
}

func (s SynapseBridgeTokenMintAndSwap) GetDestinationChainID() *big.Int {
	return nil
}

func (s SynapseBridgeTokenMintAndSwap) GetSwapTokenIndex() *uint8 {
	return nil
}

func (s SynapseBridgeTokenMintAndSwap) GetSwapMinAmount() *big.Int {
	return nil
}

func (s SynapseBridgeTokenMintAndSwap) GetSwapDeadline() *big.Int {
	return nil
}

func (s SynapseBridgeTokenMintAndSwap) GetRecipientBytes() *[32]byte {
	return nil
}

func (s SynapseBridgeTokenMintAndSwap) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ bridge.EventLog = &SynapseBridgeTokenMintAndSwap{}
