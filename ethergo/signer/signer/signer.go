// Package signer provides a common interface for signing
package signer

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/core"
	"math/big"
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

// EncodeHex encodes a signature as a hex string.
func EncodeHex(sg Signature) string {
	return hexutil.Encode(Encode(sg))
}

// DecodeSignature decodes a signature.
func DecodeSignature(sig []byte) Signature {
	if len(sig) != crypto.SignatureLength {
		panic(fmt.Sprintf("wrong size for signature: got %d, want %d", len(sig), crypto.SignatureLength))
	}
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:64])
	v := new(big.Int).SetBytes([]byte{sig[64] + 27})
	return NewSignature(v, r, s)
}

var _ Signature = signatureImpl{}

// SignHash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calculated as
//
//	keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
// TODO: when ethereum is updated, import.
// this comes from: https://github.com/ewasm/go-ethereum/blob/v1.8.10/signer/core/api.go#L451
// TODO: nonetheless, we should test independently.Z
func SignHash(data []byte) ([]byte, string) {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg)), msg
}
