package awssigner

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
)

type UnsafeSigner interface {
	signer.Signer
	// UnsafeSetKeyID sets the key id.
	UnsafeSetKeyID(keyID string)
	// UnsafeSetPubKeyData sets the public key data.
	UnsafeSetPubKeyData(pubKeyData *PubKeyData)
	// UnsafeMakePublicKeyData creates the public key data.
	UnsafeMakePublicKeyData(ctx context.Context) (*PubKeyData, error)
}

// MakeUnsafeSigner creates a new unsafe signer.
//
// / this is only to be used for testing. It should not be used in production code under any circumstances
func MakeUnsafeSigner(client *kms.Client) UnsafeSigner {
	return &Signer{
		client: client,
	}
}

func (signingHandler *Signer) UnsafeSetKeyID(keyID string) {
	signingHandler.keyID = keyID
}

func (signingHandler *Signer) UnsafeSetPubKeyData(pubKeyData *PubKeyData) {
	signingHandler.pubKeyData = pubKeyData
}

// UnsafeMakePublicKeyData creates the public key data.
func (signingHandler *Signer) UnsafeMakePublicKeyData(ctx context.Context) (*PubKeyData, error) {
	return signingHandler.getPublicKey(ctx)
}
