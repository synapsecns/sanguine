// Package signer provides a common interface for signing
package signer

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/synapse-node/pkg/common"
	"math/big"
)

// Signer provides a common interface for signing/transacting.
//go:generate go run github.com/vektra/mockery/v2 --name Signer --output ./mocks --case=underscore
type Signer interface {
	// SignMessage signs a message
	SignMessage(ctx context.Context, message []byte, hash bool) (Signature, error)
	// GetTransactor gets the transactor for a tx manager.
	// TODO: this doesn't support pre-london txes yet
	GetTransactor(chainID *big.Int) (*bind.TransactOpts, error)
	// Address gets the address of the address of the signer
	Address() ethCommon.Address
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

func (sg signatureImpl) V() *big.Int {
	return common.CopyBigInt(sg.v)
}

func (sg signatureImpl) R() *big.Int {
	return common.CopyBigInt(sg.r)
}

func (sg signatureImpl) S() *big.Int {
	return common.CopyBigInt(sg.s)
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

var _ Signature = signatureImpl{}
