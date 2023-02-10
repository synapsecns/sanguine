// Package gcpmock provides a mock client for testing purposes.
package gcpmock

import (
	"cloud.google.com/go/kms/apiv1/kmspb"
	"context"
	"crypto/ecdsa"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/googleapis/gax-go/v2"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"hash/crc32"
	"math/big"
)

// NewMockClient creates a mock client for testing purposes.
func NewMockClient() (*MockClient, error) {
	wall, err := wallet.FromRandom()
	if err != nil {
		return nil, fmt.Errorf("could not create wallet: %w", err)
	}

	return &MockClient{
		privateKey: wall.PrivateKey(),
		wall:       wall,
	}, nil
}

// MockClient is a mock client for testing purposes.
type MockClient struct {
	privateKey *ecdsa.PrivateKey
	wall       wallet.Wallet
}

// Address gets the address of the mock client.
func (m *MockClient) Address() common.Address {
	return m.wall.Address()
}

// GetPublicKey gets the public key of the mock client.
func (m *MockClient) GetPublicKey(_ context.Context, req *kmspb.GetPublicKeyRequest, _ ...gax.CallOption) (*kmspb.PublicKey, error) {
	pubKey := &ecdsa.PublicKey{
		Curve: btcec.S256(),
		X:     core.CopyBigInt(m.privateKey.PublicKey.X),
		Y:     core.CopyBigInt(m.privateKey.PublicKey.Y),
	}

	type pkixPublicKey struct {
		Algo      pkix.AlgorithmIdentifier
		BitString asn1.BitString
	}

	algo := asn1.ObjectIdentifier{1, 2, 840, 10045, 2, 1}
	publicKeyBytes := crypto.FromECDSAPub(pubKey)

	pkxPubKey, err := asn1.Marshal(pkixPublicKey{
		Algo: pkix.AlgorithmIdentifier{
			Algorithm: algo,
			Parameters: asn1.RawValue{
				Class:      0,
				Tag:        6,
				IsCompound: false,
				Bytes:      []byte{0x2b, 0x81, 0x04, 0x00, 0x0a},
				FullBytes:  []byte{0x06, 0x05, 0x2b, 0x81, 0x04, 0x00, 0x0a},
			},
		},
		BitString: asn1.BitString{
			Bytes:     publicKeyBytes,
			BitLength: 8 * len(publicKeyBytes),
		},
	})

	if err != nil {
		return nil, fmt.Errorf("could not marshal public key: %w", err)
	}

	encodedPem := pem.EncodeToMemory(&pem.Block{
		Type:  "EC PUBLIC KEY",
		Bytes: pkxPubKey,
	})

	crchash := crc32.ChecksumIEEE(pkxPubKey)
	return &kmspb.PublicKey{
		Pem:             string(encodedPem),
		Algorithm:       kmspb.CryptoKeyVersion_EC_SIGN_SECP256K1_SHA256,
		PemCrc32C:       wrapperspb.Int64(int64(crchash)),
		Name:            req.Name,
		ProtectionLevel: kmspb.ProtectionLevel_HSM,
	}, nil
}

// AsymmetricSign signs a message with the mock client.
func (m *MockClient) AsymmetricSign(ctx context.Context, res *kmspb.AsymmetricSignRequest, options ...gax.CallOption) (*kmspb.AsymmetricSignResponse, error) {
	sigBytes, err := crypto.Sign(res.Digest.GetSha256(), m.privateKey)
	if err != nil {
		return nil, fmt.Errorf("could not sign message: %w", err)
	}

	r := new(big.Int).SetBytes(sigBytes[:32])
	s := new(big.Int).SetBytes(sigBytes[32:64])
	asSig, err := asn1.Marshal(asn1EcSig{
		R: r,
		S: s,
	})

	if err != nil {
		return nil, fmt.Errorf("could not marshal signature: %w", err)
	}

	return &kmspb.AsymmetricSignResponse{
		Name:      res.Name,
		Signature: asSig,
	}, nil
}

type asn1EcSig struct {
	R *big.Int
	S *big.Int
}
