package gcpsigner

import (
	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"context"
	"github.com/googleapis/gax-go/v2"
)

// KeyClient defines the interface for a GCP Key Management Service client.
type KeyClient interface {
	// GetPublicKey gets the public key for the given key name.
	GetPublicKey(context.Context, *kmspb.GetPublicKeyRequest, ...gax.CallOption) (*kmspb.PublicKey, error)
	// AsymmetricSign signs the given data using the given key name.
	AsymmetricSign(context.Context, *kmspb.AsymmetricSignRequest, ...gax.CallOption) (*kmspb.AsymmetricSignResponse, error)
}

var _ KeyClient = &kms.KeyManagementClient{}
