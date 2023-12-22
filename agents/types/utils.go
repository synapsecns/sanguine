package types

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
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

var EtherscanTxURLs = map[uint32]string{
	444:      "https://explorerl2-synapse-sepolia-testnet-1mdqkm651f.t.conduit.xyz/tx",
	421614:   "https://sepolia-explorer.arbitrum.io/tx",
	11155111: "https://sepolia.etherscan.io/tx",
	1115420:  "https://optimism-sepolia.blockscout.com/tx",
	534351:   "https://sepolia.scrollscan.dev/tx",
	80001:    "https://mumbai.polygonscan.com/tx",
}

func LogTx(agent, msg string, chainID uint32, tx *ethTypes.Transaction) {
	fmt.Printf("[AGENT_%s:%d] %s [%s]\n", agent, chainID, msg, GetTxLink(chainID, tx))
}

func GetTxLink(chainID uint32, tx *ethTypes.Transaction) string {
	var link string
	if tx != nil {
		link = tx.Hash().String()
	} else {
		return ""
	}
	url, ok := EtherscanTxURLs[chainID]
	if !ok {
		return link
	}
	link = fmt.Sprintf("%s/%s", url, tx.Hash().String())
	return link
}
