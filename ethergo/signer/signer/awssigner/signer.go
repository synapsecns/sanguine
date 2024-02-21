package awssigner

import (
	"context"
	"fmt"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/common"
	libp2p "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/crypto/pb"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// NewKmsSigner creates a kms handler.
func NewKmsSigner(ctx context.Context, awsRegion, awsAccessKey, awsSecretAccessKey, keyID string) (_ signer.Signer, err error) {
	aws, err := awsConfig.LoadDefaultConfig(ctx, func(options *awsConfig.LoadOptions) error {
		options.Credentials = newCredentialProvider(awsAccessKey, awsSecretAccessKey)
		options.Region = awsRegion
		return nil
	})

	kmsHandler := Signer{
		awsRegion: awsRegion,
		keyID:     keyID,
		client:    kms.NewFromConfig(aws),
	}

	if err != nil {
		return nil, fmt.Errorf("could not get address: %w", err)
	}

	kmsHandler.pubKeyData, err = kmsHandler.getPublicKey(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get public key: %w", err)
	}

	return &kmsHandler, nil
}

// Signer is an aws signer.
type Signer struct {
	// awsRegion is the aws region
	awsRegion string
	// keyID is the kms key id
	keyID string
	// client is the kms client
	client *kms.Client
	// address is the address we're signing with
	pubKeyData *PubKeyData
}

// Equals assumes we're using *libp2p.Secp256k1PublicKey type for o.
func (signingHandler *Signer) Equals(o libp2p.Key) bool {
	return signingHandler.GetPublic().Equals(o)
}

func (signingHandler *Signer) Raw() ([]byte, error) {
	return signingHandler.GetPublic().Raw()
}

func (signingHandler *Signer) Type() pb.KeyType {
	return signingHandler.GetPublic().Type()
}

func (signingHandler *Signer) Sign(bytes []byte) ([]byte, error) {
	// TODO: we should figure out a way to respect context here. One possible solution
	sigBytes, err := signingHandler.SignMessage(context.Background(), bytes, false)
	if err != nil {
		return nil, fmt.Errorf("could not derive ethereum signature: %w", err)
	}

	return signer.Encode(sigBytes), nil
}

func (signingHandler *Signer) GetPublic() libp2p.PubKey {
	var x, y *secp256k1.FieldVal
	x = &secp256k1.FieldVal{}
	y = &secp256k1.FieldVal{}

	x.SetByteSlice(signingHandler.pubKeyData.ecdsaKey.X.Bytes())
	y.SetByteSlice(signingHandler.pubKeyData.ecdsaKey.Y.Bytes())

	pubkey := secp256k1.NewPublicKey(x, y)
	return (*libp2p.Secp256k1PublicKey)(pubkey)
}

func (signingHandler *Signer) PrivKey() libp2p.PrivKey {
	return signingHandler
}

// Address gets the address of the signing group.
func (signingHandler *Signer) Address() common.Address {
	return signingHandler.pubKeyData.address
}

var _ signer.Signer = &Signer{}
var _ libp2p.PrivKey = &Signer{}
