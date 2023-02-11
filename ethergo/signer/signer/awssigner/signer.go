package awssigner

import (
	"context"
	"fmt"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

// NewKmsSigner creates a kms handler.
func NewKmsSigner(ctx context.Context, awsRegion, awsAccessKey, awsSecretAccessKey, keyID string) (_ *Signer, err error) {
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

// Address gets the address of the signing group.
func (signingHandler *Signer) Address() common.Address {
	return signingHandler.pubKeyData.address
}

var _ signer.Signer = &Signer{}
