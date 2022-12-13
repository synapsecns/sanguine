package types

import (
	"math/big"

	"github.com/synapsecns/sanguine/core"
)

// SignatureLength for ecdsa signature.
const SignatureLength = int(65)

// Signature creates a new signature.
type Signature interface {
	// V gets the v value of the signature
	V() *big.Int
	// R is the r value of the signature
	R() *big.Int
	// S is the s value of the signature
	S() *big.Int
}

// signature contains an ecdsa signature
// one of the reasons we use interfaces here is to ensure the underlying data structures
// are not accidentally mutated. To ensure this, we copy big ints before returning.
type signature struct {
	v, r, s *big.Int
}

// NewSignature creates a new signature.
func NewSignature(v, r, s *big.Int) Signature {
	return signature{
		v: core.CopyBigInt(v),
		r: core.CopyBigInt(r),
		s: core.CopyBigInt(s),
	}
}

func (s signature) V() *big.Int {
	return core.CopyBigInt(s.v)
}

func (s signature) R() *big.Int {
	return core.CopyBigInt(s.r)
}

func (s signature) S() *big.Int {
	return core.CopyBigInt(s.s)
}

var _ Signature = signature{}
