package awssigner

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/stretchr/testify/require"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/awssigner/kmsmock"
	"testing"
)

// UNSAFENewSignerFromMockKMS creates a new kms signer from a mock kms.
//
// It is exported for testing. It should not be used in production code under any circumstances
// TODO: consider instead using UnsafeSetKeyID and UnsafeSetPubKeyData
// instead. This has the advantage of allowing you to create a testutils/ module
// so we do not have to import or compile kmsmock in production code.
func UNSAFENewSignerFromMockKMS(ctx context.Context, tb testing.TB) *Signer {
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
	require.Nil(tb, err)

	signer.keyID = *testKey.KeyMetadata.KeyId
	signer.pubKeyData, err = signer.getPublicKey(ctx)
	require.Nil(tb, err)
	return signer
}
