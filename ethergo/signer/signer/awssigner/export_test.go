package awssigner

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/awssigner/kmsmock"
	"testing"
)

// NewSignerFromMockKMS creates a new kms signer from a mock kms.
func NewSignerFromMockKMS(ctx context.Context, tb testing.TB) *Signer {
	tb.Helper()

	kmsMocker := kmsmock.NewMockKMS(ctx, tb)
	signer := &Signer{
		client: kmsMocker.Client(),
	}

	testKey, err := kmsMocker.Client().CreateKey(ctx, &kms.CreateKeyInput{
		CustomerMasterKeySpec: types.CustomerMasterKeySpecEccSecgP256k1,
		Description:           aws.String("this is a test key"),
		KeyUsage:              types.KeyUsageTypeSignVerify,
		MultiRegion:           aws.Bool(false),
	})
	Nil(tb, err)

	signer.keyID = *testKey.KeyMetadata.KeyId
	signer.pubKeyData, err = signer.getPublicKey(ctx)
	Nil(tb, err)
	return signer
}

// Client gets the client.
func (signingHandler *Signer) Client() *kms.Client {
	return signingHandler.client
}
