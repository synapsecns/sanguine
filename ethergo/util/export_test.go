package util

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func MakeOptions(options ...CopyOption) TestCopyOptions {
	return makeOptions(options...)
}

// TestCopyOptions exports the TestCopyOptions interface for testing.
// this is used to export unexported fields from copyOptions.
type TestCopyOptions interface {
	Nonce() *uint64
	GasPrice() *big.Int
	GasFeeCap() *big.Int
	GasTipCap() *big.Int
	TxType() *uint8
}

var _ TestCopyOptions = copyOptions{}

func (c copyOptions) Nonce() *uint64 {
	return c.nonce
}

func (c copyOptions) GasPrice() *big.Int {
	return c.gasPrice
}

func (c copyOptions) GasFeeCap() *big.Int {
	return c.gasFeeCap
}

func (c copyOptions) GasTipCap() *big.Int {
	return c.gasTipCap
}

func (c copyOptions) TxType() *uint8 {
	return c.txType
}

// NullFieldAttribute is a constant used to test the null field attribute.
// it exports the underlying constant for testing.
const NullFieldAttribute = nullFieldAttribute

func AddressPtrToString(address *common.Address) string {
	return addressPtrToString(address)
}

const (
	// HashAttr exports hashAttr for testing.
	HashAttr = hashAttr
	// FromAttr exports fromAttr for testing.
	FromAttr = fromAttr
	// ToAttr exports toAttr for testing.
	ToAttr = toAttr
	// DataAttr exports dataAttr for testing.
	DataAttr = dataAttr
	// ValueAttr exports valueAttr for testing.
	ValueAttr = valueAttr
	// NonceAttr exports nonceAttr for testing.
	NonceAttr = nonceAttr
	// GasLimitAttr exports gasLimitAttr for testing.
	GasLimitAttr = gasLimitAttr
	// ChainIDAttr exports chainIDAttr for testing.
	ChainIDAttr = chainIDAttr
	// GasPriceAttr exports gasPriceAttr for testing.
	GasPriceAttr = gasPriceAttr
	// GasFeeCapAttr exports gasFeeCapAttr for testing.
	GasFeeCapAttr = gasFeeCapAttr
	// GasTipCapAttr exports gasTipCapAttr for testing.
	GasTipCapAttr = gasTipCapAttr
)
