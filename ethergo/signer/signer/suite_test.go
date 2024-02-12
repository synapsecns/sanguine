package signer_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/awssigner"
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
	s.addSigner(config.AWSType, awssigner.UNSAFENewSignerFromMockKMS(s.GetTestContext(), s.T()))

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
