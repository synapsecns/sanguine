// Package signer provides a common interface for signing
package signer

import (
	"context"
	"github.com/synapsecns/synapse-node/pkg/common"
	"math/big"
)

// Signer provides a common interface for signing/transactnig.
type Signer interface {
	// SignMessage signs a message
	SignMessage(_ context.Context, message []byte) (Signature, error)
}

// Signature is an ecdsa signature interface.
type Signature interface {
	V() *big.Int
	R() *big.Int
	S() *big.Int
}

// NewSignature creates a new signature using the v, r, and s params.
func NewSignature(v, r, s *big.Int) Signature {
	return signatureImpl{
		v: v,
		r: r,
		s: s,
	}
}

type signatureImpl struct {
	v, r, s *big.Int
}

func (s signatureImpl) V() *big.Int {
	return common.CopyBigInt(s.v)
}

func (s signatureImpl) R() *big.Int {
	return common.CopyBigInt(s.r)
}

func (s signatureImpl) S() *big.Int {
	return common.CopyBigInt(s.s)
}

var _ Signature = signatureImpl{}
