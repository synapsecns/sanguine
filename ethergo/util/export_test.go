package util

import "math/big"

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
