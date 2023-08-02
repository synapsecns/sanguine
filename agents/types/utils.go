package types

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)


func sign(ctx context.Context, signer signer.Signer, encoded, salt []byte) (signer.Signature, common.Hash, error) {
	hashedEncodedReceipt := crypto.Keccak256Hash(encoded).Bytes()
	toSign := append(crypto.Keccak256Hash(salt).Bytes(), hashedEncodedReceipt...)

	hashedDigest, err := HashRawBytes(toSign)
	if err != nil {
		return nil, common.Hash{}, fmt.Errorf("could not hash receipt: %w", err)
	}

	signature, err := signer.SignMessage(ctx, core.BytesToSlice(hashedDigest), false)
	if err != nil {
		return nil, common.Hash{}, fmt.Errorf("could not sign: %w", err)
	}
	return signature, hashedDigest, nil
}
