package util

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core"
)

// CopyTXWithNonce copies a transaction and sets the nonce to the given value.
func CopyTXWithNonce(unsignedTx *types.Transaction, nonce uint64) (*types.Transaction, error) {
	// tx is immutable except within the confines of type. Here we manually copy over the inner values

	// these need to be overwritten, but copied over anyway for parity
	v, r, s := unsignedTx.RawSignatureValues()

	switch unsignedTx.Type() {
	case types.LegacyTxType:
		return types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			GasPrice: core.CopyBigInt(unsignedTx.GasPrice()),
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
			ChainID:    unsignedTx.ChainId(),
			Nonce:      nonce,
			GasTipCap:  core.CopyBigInt(unsignedTx.GasTipCap()),
			GasFeeCap:  core.CopyBigInt(unsignedTx.GasFeeCap()),
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
