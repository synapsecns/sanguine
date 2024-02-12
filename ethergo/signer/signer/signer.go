// Package signer provides a common interface for signing
package signer

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	libp2p "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/synapsecns/sanguine/core"
)

// Signer provides a common interface for signing/transacting.
//
//go:generate go run github.com/vektra/mockery/v2 --name Signer --output ./mocks --case=underscore
type Signer interface {
	// SignMessage signs a message
	SignMessage(ctx context.Context, message []byte, hash bool) (Signature, error)
	// GetTransactor gets the transactor for a tx manager.
	// TODO: this doesn't support pre-london txes yet
	GetTransactor(ctx context.Context, chainID *big.Int) (*bind.TransactOpts, error)
	// Address gets the address of the address of the signer
	Address() ethCommon.Address
	// PrivKey gets the libp2p priv key.
	// TODO: we should consider abstracting this otu of ethergo. It's an odd, synapse specific dependency
	// TODO: this method also needs a cross-implementation test similiar to RunOnAllDBs()
	// that does not fit into what etherog is supposed to be.
	PrivKey() libp2p.PrivKey
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

// IsEqual checks if two signatures are equal.
func IsEqual(sig1, sig2 Signature) bool {
	if sig1.R().Cmp(sig2.R()) != 0 {
		return false
	}
	if sig1.S().Cmp(sig2.S()) != 0 {
		return false
	}
	return true
}

type signatureImpl struct {
	v, r, s *big.Int
}

func (sg signatureImpl) V() *big.Int {
	return core.CopyBigInt(sg.v)
}

func (sg signatureImpl) R() *big.Int {
	return core.CopyBigInt(sg.r)
}

func (sg signatureImpl) S() *big.Int {
	return core.CopyBigInt(sg.s)
}

// Encode encodes a signature.
func Encode(sg Signature) []byte {
	r, s := sg.R().Bytes(), sg.S().Bytes()
	sig := make([]byte, crypto.SignatureLength)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	sig[64] = byte(sg.V().Uint64())

	return sig
}

// DecodeSignature decodes a signature.
func DecodeSignature(sig []byte) Signature {
	if len(sig) != crypto.SignatureLength {
		panic(fmt.Sprintf("wrong size for signature: got %d, want %d", len(sig), crypto.SignatureLength))
	}
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:64])
	v := new(big.Int).SetBytes([]byte{sig[64]})
	return NewSignature(v, r, s)
}

var _ Signature = signatureImpl{}
