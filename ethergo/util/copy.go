package util

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core"
	"math/big"
)

// CopyTX copies a transaction and sets the nonce to the given value.
func CopyTX(unsignedTx *types.Transaction, options ...CopyOption) (*types.Transaction, error) {
	// tx is immutable except within the confines of type. Here we manually copy over the inner values

	// these need to be overwritten, but copied over anyway for parity
	v, r, s := unsignedTx.RawSignatureValues()

	newNonce := unsignedTx.Nonce()
	newGasPrice := unsignedTx.GasPrice()
	newGasFeeCap := unsignedTx.GasFeeCap()
	newGasTipCap := unsignedTx.GasTipCap()
	toChange := makeOptions(options...)
	if toChange.nonce != nil {
		newNonce = *toChange.nonce
	}
	if toChange.gasPrice != nil {
		newGasPrice = toChange.gasPrice
	}
	if toChange.gasFeeCap != nil {
		newGasFeeCap = toChange.gasFeeCap
	}
	if toChange.gasTipCap != nil {
		newGasTipCap = toChange.gasTipCap
	}

	switch unsignedTx.Type() {
	case types.LegacyTxType:
		return types.NewTx(&types.LegacyTx{
			Nonce:    newNonce,
			GasPrice: core.CopyBigInt(newGasPrice),
			Gas:      unsignedTx.Gas(),
			To:       unsignedTx.To(),
			Value:    core.CopyBigInt(unsignedTx.Value()),
			Data:     unsignedTx.Data(),
			V:        core.CopyBigInt(v),
			R:        core.CopyBigInt(r),
			S:        core.CopyBigInt(s),
		}), nil
	case types.AccessListTxType:
		return nil, fmt.Errorf("unsupported tx type %d", types.AccessListTxType)
	case types.DynamicFeeTxType:
		return types.NewTx(&types.DynamicFeeTx{
			ChainID:    core.CopyBigInt(unsignedTx.ChainId()),
			Nonce:      newNonce,
			GasTipCap:  core.CopyBigInt(newGasTipCap),
			GasFeeCap:  core.CopyBigInt(newGasFeeCap),
			Gas:        unsignedTx.Gas(),
			To:         unsignedTx.To(),
			Value:      core.CopyBigInt(unsignedTx.Value()),
			Data:       unsignedTx.Data(),
			AccessList: unsignedTx.AccessList(),
			V:          core.CopyBigInt(v),
			R:          core.CopyBigInt(r),
			S:          core.CopyBigInt(s),
		}), nil
	}
	return nil, errors.New("an unexpected error occurred")
}

// copyOptions is a struct that holds the options for copying a transaction
type copyOptions struct {
	nonce     *uint64
	gasPrice  *big.Int
	gasFeeCap *big.Int
	gasTipCap *big.Int
}

// CopyOption is a function that sets a copy option
// Certain options are not supported for certain transaction types on purpose
// please exercise caution before adding new options
type CopyOption func(*copyOptions)

// WithNonce sets the nonce for the copy
func WithNonce(nonce uint64) CopyOption {
	return func(options *copyOptions) {
		options.nonce = &nonce
	}
}

// WithGasPrice sets the gas price for the copy
func WithGasPrice(gasPrice *big.Int) CopyOption {
	return func(options *copyOptions) {
		options.gasPrice = gasPrice
	}
}

// WithGasFeeCap sets the gas fee cap for the copy
func WithGasFeeCap(gasFeeCap *big.Int) CopyOption {
	return func(options *copyOptions) {
		options.gasFeeCap = gasFeeCap
	}
}

// WithGasTipCap sets the gas tip cap for the copy
func WithGasTipCap(gasTipCap *big.Int) CopyOption {
	return func(options *copyOptions) {
		options.gasTipCap = gasTipCap
	}
}

func makeOptions(options ...CopyOption) *copyOptions {
	opts := &copyOptions{}
	for _, option := range options {
		option(opts)
	}
	return opts
}
