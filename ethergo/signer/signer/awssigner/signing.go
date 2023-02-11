package awssigner

import (
	"bytes"
	"context"
	"encoding/asn1"
	"encoding/hex"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	kmsTypes "github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/pkg/errors"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"math/big"
)

var (
	secp256k1N     = crypto.S256().Params().N
	secp256k1HalfN = new(big.Int).Div(secp256k1N, big.NewInt(2))
)

// GetTransactor creates a kms transactor.
func (signingHandler *Signer) GetTransactor(ctx context.Context, chainID *big.Int) (*bind.TransactOpts, error) {
	pubKeyBytes := secp256k1.S256().Marshal(signingHandler.pubKeyData.ecdsaKey.X, signingHandler.pubKeyData.ecdsaKey.Y)
	latestSigner := types.LatestSignerForChainID(chainID)

	signerFn := func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
		if address != signingHandler.pubKeyData.address {
			return nil, bind.ErrNotAuthorized
		}

		txHashBytes := latestSigner.Hash(tx).Bytes()

		rBytes, sBytes, err := signingHandler.getSignatureFromKMS(ctx, txHashBytes)
		if err != nil {
			return nil, fmt.Errorf("could not get signature: %w", err)
		}

		// Adjust S value from signature according to Ethereum standard
		// see yellow paper: appendix f https://ethereum.github.io/yellowpaper/paper.pdf
		sBigInt := new(big.Int).SetBytes(sBytes)
		if sBigInt.Cmp(secp256k1HalfN) > 0 {
			sBytes = new(big.Int).Sub(secp256k1N, sBigInt).Bytes()
		}

		signature, err := signingHandler.getEthereumSignature(pubKeyBytes, txHashBytes, rBytes, sBytes)
		if err != nil {
			return nil, fmt.Errorf("could not derive signature: %w", err)
		}

		//nolint: wrapcheck
		return tx.WithSignature(latestSigner, signature)
	}

	return &bind.TransactOpts{
		From:   signingHandler.pubKeyData.address,
		Signer: signerFn,
	}, nil
}

func (signingHandler *Signer) getSignatureFromKMS(
	ctx context.Context, txHashBytes []byte,
) ([]byte, []byte, error) {
	signInput := &kms.SignInput{
		KeyId:            aws.String(signingHandler.keyID),
		SigningAlgorithm: kmsTypes.SigningAlgorithmSpecEcdsaSha256,
		MessageType:      kmsTypes.MessageTypeDigest,
		Message:          txHashBytes,
	}

	signOutput, err := signingHandler.client.Sign(ctx, signInput)
	if err != nil {
		return nil, nil, fmt.Errorf("could not get signature from kms: %w", err)
	}

	var sigAsn1 asn1EcSig
	_, err = asn1.Unmarshal(signOutput.Signature, &sigAsn1)
	if err != nil {
		return nil, nil, fmt.Errorf("could not unmarshall asn: %w", err)
	}

	return sigAsn1.R.Bytes, sigAsn1.S.Bytes, nil
}

// SignMessage signs a hashed message.
func (signingHandler *Signer) SignMessage(ctx context.Context, message []byte, hash bool) (signer.Signature, error) {
	if hash {
		message = crypto.Keccak256(message)
	}

	rBytes, sBytes, err := signingHandler.getSignatureFromKMS(ctx, message)
	if err != nil {
		return nil, fmt.Errorf("could not sign: %w", err)
	}

	sigBytes, err := signingHandler.getEthereumSignature(signingHandler.pubKeyData.rawPubKey, message, rBytes, sBytes)
	if err != nil {
		return nil, fmt.Errorf("could not derive ethereum signature: %w", err)
	}

	return signer.DecodeSignature(sigBytes), nil
}

func (signingHandler *Signer) getEthereumSignature(expectedPublicKeyBytes []byte, txHash []byte, r []byte, s []byte) ([]byte, error) {
	rsSignature := append(adjustSignatureLength(r), adjustSignatureLength(s)...)
	signature := append(rsSignature, []byte{0}...)

	recoveredPublicKeyBytes, err := crypto.Ecrecover(txHash, signature)
	if err != nil {
		return nil, fmt.Errorf("could not get key: %w", err)
	}

	if hex.EncodeToString(recoveredPublicKeyBytes) != hex.EncodeToString(expectedPublicKeyBytes) {
		signature = append(rsSignature, []byte{1}...)
		recoveredPublicKeyBytes, err = crypto.Ecrecover(txHash, signature)
		if err != nil {
			return nil, fmt.Errorf("could not recover signature: %w", err)
		}

		if hex.EncodeToString(recoveredPublicKeyBytes) != hex.EncodeToString(expectedPublicKeyBytes) {
			return nil, errors.New("can not reconstruct public key from sig")
		}
	}

	return signature, nil
}

func adjustSignatureLength(buffer []byte) []byte {
	buffer = bytes.TrimLeft(buffer, "\x00")
	for len(buffer) < 32 {
		zeroBuf := []byte{0}
		buffer = append(zeroBuf, buffer...)
	}
	return buffer
}

// asn1EcSig is the asn1 signature for a digest.
type asn1EcSig struct {
	// R is the r-value of the signature. This cannot be resused
	R asn1.RawValue
	// S is the s-value of the signature
	S asn1.RawValue
}
