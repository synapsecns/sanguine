package abiutil_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/parser/abiutil/internal"
	"testing"
)

// AbiSuite defines the basic test suite.
type AbiSuite struct {
	*testsuite.TestSuite
	backend      backends.SimulatedTestBackend
	testContract *internal.TestSignature
	metadata     *bind.MetaData
}

// NewTestSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewTestSuite(tb testing.TB) *AbiSuite {
	tb.Helper()
	return &AbiSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		metadata:  internal.TestSignatureMetaData,
	}
}

func (a *AbiSuite) SetupSuite() {
	a.TestSuite.SetupSuite()

	a.backend = simulated.NewSimulatedBackend(a.GetSuiteContext(), a.T())
	auth := a.backend.GetTxContext(a.GetSuiteContext(), nil)

	var err error
	_, _, a.testContract, err = internal.DeployTestSignature(auth.TransactOpts, a.backend)
	a.Require().NoError(err)
}

func TestAbiSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
