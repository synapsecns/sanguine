// not how this is always scoped outside of the package itself
package relayer_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	cctpTest "github.com/synapsecns/sanguine/services/cctp-relayer/testutil"
	omnirpcHelper "github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	scribeHelper "github.com/synapsecns/sanguine/services/scribe/testhelper"
	"golang.org/x/sync/errgroup"
)

// TestHelperSuite defines the basic test suite.
type CCTPRelayerSuite struct {
	*testsuite.TestSuite
	// testBackends contins a list of all test backends
	testBackends []backends.SimulatedTestBackend
	// we'll use this later
	deployManager *cctpTest.DeployManager
	// testScribeURL setup in SetupTest
	testScribe string
	// testOmnirpc setup in SetupTest
	testOmnirpc string
}

// NewTestSuite creates a new test suite.
func NewTestSuite(tb testing.TB) *CCTPRelayerSuite {
	tb.Helper()
	return &CCTPRelayerSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (s *CCTPRelayerSuite) SetupSuite() {
	// for tracing
	localmetrics.SetupTestJaeger(s.GetSuiteContext(), s.T())
	// let's create 2 mock chains
	chainIDs := []uint64{1, 43114}

	// preallocate a slice for testbackends to the size of chainIDs
	// this way we can avoid non-deterministic order + needing to acquire/release a lock
	s.testBackends = make([]backends.SimulatedTestBackend, len(chainIDs))

	g, _ := errgroup.WithContext(s.GetSuiteContext())
	for i, chainID := range chainIDs {
		pos := i           // get position of chain id in array
		chainID := chainID // capture func literal
		g.Go(func() error {
			// we need to use the embedded backend here, because the simulated backend doesn't support rpcs required by scribe
			backend := geth.NewEmbeddedBackendForChainID(s.GetSuiteContext(), s.T(), new(big.Int).SetUint64(chainID))

			// add the backend to the list of backends
			s.testBackends[pos] = backend
			return nil
		})
	}

	// wait for all backends to be ready
	if err := g.Wait(); err != nil {
		s.T().Fatal(err)
	}

}

func (s *CCTPRelayerSuite) SetupTest() {
	s.TestSuite.SetupTest()

	s.deployManager = cctpTest.NewDeployManager(s.T())
	// deploy the contract to all backends
	// note: since we haven't gone over contract generation, we'll do this a bit later
	// s.deployManager.BulkDeploy(s.GetTestContext(), s.testBackends, )

	// create the test scribe backend
	s.testScribe = scribeHelper.NewTestScribe(s.GetTestContext(), s.T(), s.deployManager.GetDeployedContracts(), s.testBackends...)
	// create the test omnirpc backend
	s.testOmnirpc = omnirpcHelper.NewOmnirpcServer(s.GetTestContext(), s.T(), s.testBackends...)

	// deploy the contract to all backends
	s.deployManager.BulkDeploy(s.GetTestContext(), s.testBackends, cctpTest.SynapseCCTPType, cctpTest.MockMintBurnTokenType)
}

func TestCCTPRelayerSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
