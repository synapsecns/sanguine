package types

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// SignEncoder encodes a type, and signs it with the given salt.
func signEncoder(ctx context.Context, signer signer.Signer, encoder Encoder, salt string) (signer.Signature, []byte, common.Hash, error) {
	// Encode the given type.
	encoded, err := encoder.Encode()
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("could not encode: %w", err)
	}

	// Hash the encoded type, and concatenate with hashed salt.
	hashedEncoded := crypto.Keccak256Hash(encoded).Bytes()
	toSign := append(crypto.Keccak256Hash([]byte(salt)).Bytes(), hashedEncoded...)
	hashedDigest, err := HashRawBytes(toSign)
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("could not hash: %w", err)
	}

	// Sign the message.
	sig, err := signer.SignMessage(ctx, core.BytesToSlice(hashedDigest), false)
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("could not sign: %w", err)
	}

	sig = NewSignature(new(big.Int).Add(big.NewInt(27), sig.V()), sig.R(), sig.S())

	return sig, encoded, hashedDigest, nil
}
