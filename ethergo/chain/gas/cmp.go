package gas

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core"
	"math/big"
)

// CompareGas allows for gas comparisons between txes. In the case of an eip-1559 txOpts and a
// non-eip 1559 txOpts the eip-1559 txOpts we follow the mempool logic of setting the gasTipCap and the GasFeeCap to the same value
// legacy txes do not use effectivegastipcap since base fee is not present in non-eip 1559 chains
// This is done because the current mempool logic (https://github.com/ethereum/EIPs/blob/master/EIPS/eip-1559.md)
//
// for non bumps (checking confirmations, etc) a nil gas block can be added which will order by tip (according to geth)
//
// gas block here.
//
//	-1 if x <  y
//	 0 if x == y
//	+1 if x >  y
func CompareGas(x *types.Transaction, y *types.Transaction, baseFee *big.Int) int {
	if x.Type() == types.LegacyTxType && y.Type() == types.LegacyTxType {
		return x.GasPrice().Cmp(y.GasPrice())
	}

	if baseFee != nil {
		return x.EffectiveGasTipCmp(y, baseFee)
	}
	return x.EffectiveGasTipCmp(y, nil)
}

// FeeGreaterThan determines if gas for a transaction is greater than threshold. This is used to enforce limits.
// DynamicFeeTxes set gas price to FeeTipCap so this will work for both.
func FeeGreaterThan(tx *types.Transaction, threshold *big.Int) bool {
	return tx.GasPrice().Cmp(threshold) > 0
}

// IsDynamicTx determines if a transaction is dynamic based on the options passed in.
// this follows the same logic as the abi binder in determining what txOpts type to use.
// see: https://github.com/ethereum/go-ethereum/blob/b20bc5c0cae6901209610c5f53f01401b6f7974e/accounts/abi/bind/base.go#L310
func IsDynamicTx(opts *bind.TransactOpts) bool {
	return opts.GasFeeCap != nil
}

// OptsToComparableTx converts transaction options to a transaction. This should be used for comparisons only.
func OptsToComparableTx(opts *bind.TransactOpts) (comparableTx *types.Transaction) {
	if IsDynamicTx(opts) {
		comparableTx = types.NewTx(&types.DynamicFeeTx{
			GasTipCap: core.CopyBigInt(opts.GasTipCap),
			GasFeeCap: core.CopyBigInt(opts.GasFeeCap),
		})
	} else {
		comparableTx = types.NewTx(&types.LegacyTx{
			GasPrice: core.CopyBigInt(opts.GasPrice),
		})
	}
	return comparableTx
}

// BumpGasFees bumps the gas price by percent increase. In the case of legacy txes this bumps the gas price
// in the case of fees, this bumps both the tip and fee cap by percent. If the fee cap exceeds the percent bump
// but the fee cap doesn't and the fee cap is still below the tip cap the new fee cap is used without bumping the tip cap.
func BumpGasFees(opts *bind.TransactOpts, percentIncrease int, baseFee *big.Int, maxPrice *big.Int) {
	if IsDynamicTx(opts) {
		bumpDynamicTxFees(opts, percentIncrease, baseFee, maxPrice)
	} else {
		bumpLegacyTxFees(opts, percentIncrease, maxPrice)
	}
}

// bumpLegacyTxFees bumps a legacy txes fee.
func bumpLegacyTxFees(opts *bind.TransactOpts, percentIncrease int, maxPrice *big.Int) {
	// bump the gas price by percent increase
	newPrice := BumpByPercent(opts.GasPrice, percentIncrease)
	// if the current price exceeds max price skip
	if newPrice.Cmp(maxPrice) > 0 {
		logger.Warnf("gas price %s exceeds max price %s, not bumping", newPrice, maxPrice)
		return
	}
	// otherwise set it
	opts.GasPrice = newPrice
}

// bumpDynamicTxFees bumps a dynamicFeeTx fee.
func bumpDynamicTxFees(opts *bind.TransactOpts, percentIncrease int, baseFee, maxPrice *big.Int) {
	newTipCap := BumpByPercent(opts.GasTipCap, percentIncrease)

	// if new fee cap less than tip cap AND base (fee + fee cap) > tip cap
	if maxPrice.Cmp(newTipCap) > 0 && big.NewInt(0).Sub(maxPrice, baseFee).Cmp(newTipCap) > 0 {
		opts.GasTipCap = newTipCap
		logger.Warnf("new tip cap %s still less than fee cap %s + base fee %s, bumping tip not and not fee", newTipCap, maxPrice, baseFee)
	}

	opts.GasFeeCap = maxPrice
}

// BumpByPercent bumps a gas price by a percentage.
func BumpByPercent(gasPrice *big.Int, percentIncrease int) *big.Int {
	price := core.CopyBigInt(gasPrice)
	calculatedGasPrice := big.NewFloat(0).Mul(big.NewFloat(1+0.01*float64(percentIncrease)), big.NewFloat(0).SetInt(price))
	price, _ = calculatedGasPrice.Int(price)
	return price
}

// min gets the minimum of two big ints.
func min(a, b *big.Int) *big.Int {
	if a.Cmp(b) > 0 {
		return b
	}
	return a
}
