package util

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core"
	"go.opentelemetry.io/otel/attribute"
	"math/big"
)

// CopyTransactOpts creates a deep copy of the given TransactOpts struct
// with big ints wrapped in core.CopyBigInt().
func CopyTransactOpts(opts *bind.TransactOpts) *bind.TransactOpts {
	copyOpts := &bind.TransactOpts{
		From:      opts.From,
		Nonce:     core.CopyBigInt(opts.Nonce),
		Signer:    opts.Signer,
		Value:     core.CopyBigInt(opts.Value),
		GasPrice:  core.CopyBigInt(opts.GasPrice),
		GasFeeCap: core.CopyBigInt(opts.GasFeeCap),
		GasTipCap: core.CopyBigInt(opts.GasTipCap),
		GasLimit:  opts.GasLimit,
		Context:   opts.Context,
		NoSend:    opts.NoSend,
	}
	return copyOpts
}

const (
	hashAttr      = "tx.Hash"
	fromAttr      = "tx.From"
	toAttr        = "tx.To"
	dataAttr      = "tx.Data"
	valueAttr     = "tx.Value"
	nonceAttr     = "tx.Nonce"
	gasLimitAttr  = "tx.GasLimit"
	chainIDAttr   = "tx.ChainID"
	gasPriceAttr  = "tx.GasPrice"
	gasFeeCapAttr = "tx.GasFeeCap"
	gasTipCapAttr = "tx.GasTipCap"
)

// TxToAttributes converts a transaction to a slice of attribute.KeyValue.
func TxToAttributes(transaction *types.Transaction) []attribute.KeyValue {
	var from string
	call, err := TxToCall(transaction)
	if err != nil {
		from = fmt.Sprintf("could not be detected: %v", err)
	} else {
		from = call.From.Hex()
	}
	var attributes = []attribute.KeyValue{
		attribute.String(hashAttr, transaction.Hash().Hex()),
		attribute.String(fromAttr, from),
		attribute.String(toAttr, addressPtrToString(transaction.To())),
		attribute.String(dataAttr, fmt.Sprintf("%x", transaction.Data())),
		attribute.String(valueAttr, BigPtrToString(transaction.Value())),
		// TODO: this could be downcast to int64, but it's unclear how we should handle overflows.
		// since this is only for tracing, we can probably ignore it for now.
		attribute.Int64(nonceAttr, int64(transaction.Nonce())),
		attribute.Int64(gasLimitAttr, int64(transaction.Gas())),
		attribute.String(chainIDAttr, BigPtrToString(transaction.ChainId())),
	}

	if transaction.Type() == types.LegacyTxType && transaction.GasPrice() != nil {
		attributes = append(attributes, attribute.String(gasPriceAttr, BigPtrToString(transaction.GasPrice())))
	}

	if transaction.Type() == types.DynamicFeeTxType && transaction.GasFeeCap() != nil {
		attributes = append(attributes, attribute.String(gasFeeCapAttr, BigPtrToString(transaction.GasFeeCap())))
	}

	if transaction.Type() == types.DynamicFeeTxType && transaction.GasTipCap() != nil {
		attributes = append(attributes, attribute.String(gasTipCapAttr, BigPtrToString(transaction.GasTipCap())))
	}

	return attributes
}

const nullFieldAttribute = "null"

func addressPtrToString(address *common.Address) string {
	if address == nil {
		return nullFieldAttribute
	}
	return address.Hex()
}

// BigPtrToString converts a big.Int pointer to a string.
func BigPtrToString(num *big.Int) string {
	if num == nil {
		return nullFieldAttribute
	}
	return num.String()
}
