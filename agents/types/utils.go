package types

import (
	"context"
	"fmt"
	"strings"

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
	signature, err := signer.SignMessage(ctx, core.BytesToSlice(hashedDigest), false)
	if err != nil {
		return nil, nil, common.Hash{}, fmt.Errorf("could not sign: %w", err)
	}
	return signature, encoded, hashedDigest, nil
}

var etherscanTxURLs = map[uint32]string{
	444:      "https://explorerl2-synapse-sepolia-testnet-1mdqkm651f.t.conduit.xyz/tx",
	421614:   "https://sepolia-explorer.arbitrum.io/tx",
	11155111: "https://sepolia.etherscan.io/tx",
}

func LogTx(agent, msg string, chainID uint32, tx *ethTypes.Transaction) {
	fmt.Printf("[%s:%d] Submitted %s: %s\n", agent, chainID, msg, getTxLink(chainID, tx))
}

func getTxLink(chainID uint32, tx *ethTypes.Transaction) string {
	var link string
	if tx != nil {
		link = tx.Hash().String()
	} else {
		metadata := []string{}
		metadata = append(metadata, fmt.Sprintf("cost=%s", tx.Cost().String()))
		metadata = append(metadata, fmt.Sprintf("gasPrice=%s", tx.GasPrice().String()))
		metadata = append(metadata, fmt.Sprintf("gasFeeCap=%s", tx.GasFeeCap().String()))
		metadata = append(metadata, fmt.Sprintf("gasTipCap=%s", tx.GasTipCap().String()))
		return "Not submitted; metadata: " + strings.Join(metadata, ", ")
	}
	url, ok := etherscanTxURLs[chainID]
	if !ok {
		return link
	}
	link = fmt.Sprintf("%s/%s", url, tx.Hash().String())
	return link
}
