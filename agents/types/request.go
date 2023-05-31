package types

import (
	"math/big"
)

const (
	// VersionOffset is the offset of the version.
	VersionOffset = 0
	// GasLimitOffset if the offset of the gas limit.
	GasLimitOffset = 4
	// GasDropOffset is the gas drop offset.
	GasDropOffset = 12
	// RequestSize is the size of the request.
	RequestSize = 24
)

// Request is an interface that contains the request from the base message.
//
//nolint:interfacebloat
type Request interface {
	// Version is the base message version to pass to the recipient.
	Version() uint32
	// GasLimit is the minimum amount of gas units to supply for execution.
	GasLimit() uint64
	// GasDrop is the minimum amount of gas token to drop to the recipient.
	GasDrop() *big.Int
}

// requestImpl implements a request. It is used for testutils. Real requests are emitted by the contract.
type requestImpl struct {
	version  uint32
	gasLimit uint64
	gasDrop  *big.Int
}

// NewRequest creates a new request from fields passed in.
func NewRequest(version uint32, gasLimit uint64, gasDrop *big.Int) Request {
	return &requestImpl{
		version:  version,
		gasLimit: gasLimit,
		gasDrop:  gasDrop,
	}
}

func (m requestImpl) Version() uint32 {
	return m.version
}

func (m requestImpl) GasLimit() uint64 {
	return m.gasLimit
}

func (m requestImpl) GasDrop() *big.Int {
	return m.gasDrop
}
