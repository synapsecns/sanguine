package metaswap

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
	"math/big"
)

// GetRaw gets the raw event logs.
func (s MetaSwapTokenSwapUnderlying) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the event.
func (s MetaSwapTokenSwapUnderlying) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s MetaSwapTokenSwapUnderlying) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s MetaSwapTokenSwapUnderlying) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the swap event.
func (s MetaSwapTokenSwapUnderlying) GetEventType() swap.EventType {
	return swap.TokenSwapUnderlyingEvent
}

// GetBuyer gets the buyer.
func (s MetaSwapTokenSwapUnderlying) GetBuyer() *common.Address {
	return &s.Buyer
}

// GetTokensSold gets the tokens sold.
func (s MetaSwapTokenSwapUnderlying) GetTokensSold() *big.Int {
	return core.CopyBigInt(s.TokensSold)
}

// GetTokensBought gets the tokens bought.
func (s MetaSwapTokenSwapUnderlying) GetTokensBought() *big.Int {
	return core.CopyBigInt(s.TokensBought)
}

// GetSoldID gets the solid id.
func (s MetaSwapTokenSwapUnderlying) GetSoldID() *big.Int {
	return core.CopyBigInt(s.SoldId)
}

// GetBoughtID gets the bought id.
func (s MetaSwapTokenSwapUnderlying) GetBoughtID() *big.Int {
	return core.CopyBigInt(s.BoughtId)
}

// GetLPTokenAmount gets the LP token supply.
func (s MetaSwapTokenSwapUnderlying) GetLPTokenAmount() *big.Int {
	return nil
}

// GetAmount puts the amount in a map with it's associated token index.
func (s MetaSwapTokenSwapUnderlying) GetAmount() map[uint8]string {
	output := map[uint8]string{uint8(s.SoldId.Int64()): core.CopyBigInt(s.TokensSold).String(), uint8(s.BoughtId.Int64()): core.CopyBigInt(s.TokensBought).String()}
	return output
}

// GetAmountFee gets the amount.
func (s MetaSwapTokenSwapUnderlying) GetAmountFee() map[uint8]string {
	return map[uint8]string{}
}

// GetProtocolFee gets the protocol fee of the tx.
func (s MetaSwapTokenSwapUnderlying) GetProtocolFee() *big.Int {
	return nil
}

// GetProvider gets the provider removing liquidity.
func (s MetaSwapTokenSwapUnderlying) GetProvider() *common.Address {
	return nil
}

// GetInvariant gets the invariant.
func (s MetaSwapTokenSwapUnderlying) GetInvariant() *big.Int {
	return nil
}

// GetLPTokenSupply gets LP token supply.
func (s MetaSwapTokenSwapUnderlying) GetLPTokenSupply() *big.Int {
	return nil
}

// GetNewAdminFee gets the new admin fee.
func (s MetaSwapTokenSwapUnderlying) GetNewAdminFee() *big.Int {
	return nil
}

// GetNewSwapFee gets the new swap fee.
func (s MetaSwapTokenSwapUnderlying) GetNewSwapFee() *big.Int {
	return nil
}

// GetOldA gets the old A.
func (s MetaSwapTokenSwapUnderlying) GetOldA() *big.Int {
	return nil
}

// GetNewA gets the new A.
func (s MetaSwapTokenSwapUnderlying) GetNewA() *big.Int {
	return nil
}

// GetInitialTime gets the initial time.
func (s MetaSwapTokenSwapUnderlying) GetInitialTime() *big.Int {
	return nil
}

// GetFutureTime gets the future time.
func (s MetaSwapTokenSwapUnderlying) GetFutureTime() *big.Int {
	return nil
}

// GetCurrentA gets the current A.
func (s MetaSwapTokenSwapUnderlying) GetCurrentA() *big.Int {
	return nil
}

// GetTime gets the current time.
func (s MetaSwapTokenSwapUnderlying) GetTime() *big.Int {
	return nil
}

// GetReceiver gets the receiver.
func (s MetaSwapTokenSwapUnderlying) GetReceiver() *common.Address {
	return nil
}

// GetEventIndex gets the index of the event in the block.
func (s MetaSwapTokenSwapUnderlying) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ swap.EventLog = &MetaSwapTokenSwapUnderlying{}
