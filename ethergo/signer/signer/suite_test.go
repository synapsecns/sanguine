package signer_test

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/awssigner"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/awssigner/kmsmock"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/gcpsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/gcpsigner/gcpmock"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"testing"
)

type SignerSuite struct {
	*testsuite.TestSuite
	testSigners []TestSigner
}

type TestSigner struct {
	signer.Signer
	signerType config.SignerType
}

func NewSignerSuite(tb testing.TB) *SignerSuite {
	tb.Helper()
	return &SignerSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestSignerSuite(t *testing.T) {
	suite.Run(t, NewSignerSuite(t))
}

func (s *SignerSuite) SetupTest() {
	s.TestSuite.SetupTest()
	// add aws signer
	s.addSigner(config.AWSType, NewSignerFromMockKMS(s.GetTestContext(), s.T()))

	// add gcp signer.
	gcpMock, err := gcpmock.NewMockClient()
	s.NoError(err, "should make mock client")

	gcpSigner, err := gcpsigner.NewManagedKey(s.GetTestContext(), gcpMock, "test")
	s.NoError(err, "should create managed key")

	s.addSigner(config.GCPType, gcpSigner)

	// add local signer
	newWallet, err := wallet.FromRandom()
	s.NoError(err, "should create wallet")
	s.addSigner(config.FileType, localsigner.NewSigner(newWallet.PrivateKey()))
}

// UNSAFENewSignerFromMockKMS creates a new kms signer from a mock kms.
//
// It is exported for testing. It should not be used in production code under any circumstances
// TODO: consider instead using UnsafeSetKeyID and UnsafeSetPubKeyData
// instead. This has the advantage of allowing you to create a testutils/ module
// so we do not have to import or compile kmsmock in production code.
func NewSignerFromMockKMS(ctx context.Context, tb testing.TB) awssigner.UnsafeSigner {
	tb.Helper()

	kmsMocker := kmsmock.NewMockKMS(ctx, tb)
	awsSigner := awssigner.MakeUnsafeSigner(kmsMocker.Client())

	testKey, err := kmsMocker.Client().CreateKey(ctx, &kms.CreateKeyInput{
		CustomerMasterKeySpec: types.CustomerMasterKeySpecEccSecgP256k1,
		Description:           aws.String("this is a test key"),
		KeyUsage:              types.KeyUsageTypeSignVerify,
		MultiRegion:           aws.Bool(false),
	})
	require.Nil(tb, err)

	awsSigner.UnsafeSetKeyID(*testKey.KeyMetadata.KeyId)
	pkdata, err := awsSigner.UnsafeMakePublicKeyData(ctx)
	require.Nil(tb, err)

	awsSigner.UnsafeSetPubKeyData(pkdata)
	return awsSigner
}

func (s *SignerSuite) addSigner(signerType config.SignerType, signer signer.Signer) {
	s.testSigners = append(s.testSigners, TestSigner{
		Signer:     signer,
		signerType: signerType,
	})
}

func (s *SignerSuite) RunOnAllSigners(f func(signer.Signer)) {
	for _, ts := range s.testSigners {
		s.T().Run(ts.signerType.String(), func(t *testing.T) {
			f(ts.Signer)
		})
	}
}
