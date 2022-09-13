package bridge

//nolint: dupl

import (
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/types"
	"math/big"
)

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SynapseBridgeTokenMint) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetToken gets the destination chain id from the mint.
func (s SynapseBridgeTokenMint) GetToken() string {
	return s.Token.String()
}

// GetAmount gets the token amount.
func (s SynapseBridgeTokenMint) GetAmount() *big.Int {
	return s.Amount
}

// GetEventType gets the type of the redeem event.
func (s SynapseBridgeTokenMint) GetEventType() types.EventType {
	return types.MintEvent
}

// GetIdentifier gets the unique identifier (txhash) for the mint.
func (s SynapseBridgeTokenMint) GetIdentifier() string {
	return s.Raw.TxHash.String()
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenMint) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenMint) GetContractAddress() string {
	return s.Raw.Address.String()
}

// GetFee gets the fee for the token mint.
func (s SynapseBridgeTokenMint) GetFee() *big.Int {
	return s.Fee
}

// GetKappa gets the gappa for the token mint.
func (s SynapseBridgeTokenMint) GetKappa() [32]byte {
	return s.Kappa
}

var _ types.CrossChainBridgeEventLog = &SynapseBridgeTokenMint{}

// GetToken gets the token for the mint and remove operation.
func (s SynapseBridgeTokenMintAndSwap) GetToken() string {
	return s.Token.String()
}

// GetAmount gets the amount fo the mint.
func (s SynapseBridgeTokenMintAndSwap) GetAmount() *big.Int {
	return s.Amount
}

// GetEventType gets the mint and remove event type.
func (s SynapseBridgeTokenMintAndSwap) GetEventType() types.EventType {
	return types.MintAndSwap
}

// GetRaw gets the raw logs.
func (s SynapseBridgeTokenMintAndSwap) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetFee gets the fee for the bridge.
func (s SynapseBridgeTokenMintAndSwap) GetFee() *big.Int {
	return s.Fee
}

// GetKappa gets the kappa value for the event.
func (s SynapseBridgeTokenMintAndSwap) GetKappa() [32]byte {
	return s.Kappa
}

// GetIdentifier gets the unique identifier (txhash) for the mint.
func (s SynapseBridgeTokenMintAndSwap) GetIdentifier() string {
	return s.Raw.TxHash.String()
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenMintAndSwap) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenMintAndSwap) GetContractAddress() string {
	return s.Raw.Address.String()
}

var _ types.CrossChainBridgeEventLog = &SynapseBridgeTokenMintAndSwap{}
