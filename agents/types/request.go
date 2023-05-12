package types

import (
	"math/big"
)

const (
	// GasLimitOffset if the offset of the gas limit.
	GasLimitOffset = 0
	// GasDropOffset is the gas drop offset.
	GasDropOffset = 8
	// RequestSize is the size of the request.
	RequestSize = 20
)

// Request is an interface that contains the request from the base message.
//
//nolint:interfacebloat
type Request interface {
	// GasLimit is the minimum amount of gas units to supply for execution.
	GasLimit() uint64
	// GasDrop is the minimum amount of gas token to drop to the recipient.
	GasDrop() *big.Int
}

// requestImpl implements a request. It is used for testutils. Real requests are emitted by the contract.
type requestImpl struct {
	gasLimit uint64
	gasDrop  *big.Int
}

// NewRequest creates a new request from fields passed in.
func NewRequest(gasLimit uint64, gasDrop *big.Int) Request {
	return &requestImpl{
		gasLimit: gasLimit,
		gasDrop:  gasDrop,
	}
}

func (m requestImpl) GasLimit() uint64 {
	return m.gasLimit
}

func (m requestImpl) GasDrop() *big.Int {
	return m.gasDrop
}
