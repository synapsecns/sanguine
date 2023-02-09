package gcpmock

import (
	"cloud.google.com/go/kms/apiv1/kmspb"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/googleapis/gax-go/v2"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	s256_pem "github.com/vanhallio/go-secp256k1-pem"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"hash/crc32"
)

// NewMockClient creates a mock client for testing purposes.
func NewMockClient() (*MockClient, error) {
	wall, err := wallet.FromRandom()
	if err != nil {
		return nil, fmt.Errorf("could not create wallet: %w", err)
	}

	return &MockClient{
		privateKey: wall.PrivateKey(),
	}, nil
}

// MockClient is a mock client for testing purposes.
type MockClient struct {
	privateKey *ecdsa.PrivateKey
}

func (m *MockClient) GetPublicKey(_ context.Context, req *kmspb.GetPublicKeyRequest, _ ...gax.CallOption) (*kmspb.PublicKey, error) {
	pubKey := &btcec.PublicKey{
		Curve: btcec.S256(),
		X:     core.CopyBigInt(m.privateKey.PublicKey.X),
		Y:     core.CopyBigInt(m.privateKey.PublicKey.Y),
	}
	rawKey, err := s256_pem.PublicKeyToPem(pubKey)
	if err != nil {
		return nil, fmt.Errorf("could not convert public key to pem: %w", err)
	}

	crchash := crc32.ChecksumIEEE(rawKey)
	return &kmspb.PublicKey{
		Pem:             string(rawKey),
		Algorithm:       kmspb.CryptoKeyVersion_EC_SIGN_SECP256K1_SHA256,
		PemCrc32C:       wrapperspb.Int64(int64(crchash)),
		Name:            req.Name,
		ProtectionLevel: kmspb.ProtectionLevel_HSM,
	}, nil
}

func (m *MockClient) AsymmetricSign(ctx context.Context, res *kmspb.AsymmetricSignRequest, options ...gax.CallOption) (*kmspb.AsymmetricSignResponse, error) {
	panic("implement me")
}
