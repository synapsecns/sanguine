package types

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// SignEncoder encodes a type, and signs it with the given salt.
func SignEncoder(ctx context.Context, signer signer.Signer, encoder Encoder, salt []byte) (signer.Signature, []byte, common.Hash, error) {
	// Encode the given type.
	encoded, err := encoder.Encode()
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("could not encode receipt: %w", err)
	}

	// Hash the encoded type, and concatenate with hashed salt.
	hashedEncoded := crypto.Keccak256Hash(encoded).Bytes()
	toSign := append(crypto.Keccak256Hash(salt).Bytes(), hashedEncoded...)
	hashedDigest, err := HashRawBytes(toSign)
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("could not hash receipt: %w", err)
	}

	// Sign the message.
	signature, err := signer.SignMessage(ctx, core.BytesToSlice(hashedDigest), false)
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("could not sign: %w", err)
	}
	return signature, encoded, hashedDigest, nil
}
