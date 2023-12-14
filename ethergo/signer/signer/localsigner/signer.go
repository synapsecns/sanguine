package localsigner

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// Signer is a new local signer.
type Signer struct {
	privateKey *ecdsa.PrivateKey
}

// NewSigner creates a new signer.
func NewSigner(key *ecdsa.PrivateKey) *Signer {
	return &Signer{
		privateKey: key,
	}
}

// GetTransactor gets the transcator.
func (s *Signer) GetTransactor(ctx context.Context, chainID *big.Int) (*bind.TransactOpts, error) {
	realSigner, err := bind.NewKeyedTransactorWithChainID(s.privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("could not create signer: %w", err)
	}
	return realSigner, nil
}

// SignMessage signs a message w/o eip-155.
func (s *Signer) SignMessage(_ context.Context, message []byte, hash bool) (signer.Signature, error) {
	if hash {
		message = crypto.Keccak256(message)
	}

	sig, err := crypto.Sign(message, s.privateKey)
	if err != nil {
		return nil, fmt.Errorf("could not sign: %w", err)
	}

	return decodeSignature(sig), nil
}

// Address gets the address of the signer.
func (s *Signer) Address() common.Address {
	return crypto.PubkeyToAddress(s.privateKey.PublicKey)
}

func decodeSignature(sig []byte) signer.Signature {
	// panic here should never happen, this is an additional sanity check and should be considered a static assertion
	if len(sig) != crypto.SignatureLength {
		panic(fmt.Sprintf("wrong size for signature: got %d, want %d", len(sig), crypto.SignatureLength))
	}
	v := new(big.Int).SetBytes([]byte{sig[64]})
	r := new(big.Int).SetBytes(sig[:32])
	s := new(big.Int).SetBytes(sig[32:64])

	return signer.NewSignature(v, r, s)
}
