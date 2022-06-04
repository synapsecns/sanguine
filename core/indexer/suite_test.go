package indexer_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"github.com/synapsecns/sanguine/core/testutil"
	"github.com/synapsecns/synapse-node/testutils"
	"github.com/synapsecns/synapse-node/testutils/backends"
	"github.com/synapsecns/synapse-node/testutils/backends/preset"
	"testing"
)

// IndexerSuite tests the indexer.
type IndexerSuite struct {
	*testutils.TestSuite
	testBackend   backends.TestBackend
	deployManager *testutil.DeployManager
	homeContract  *home.HomeRef
}

// NewIndexerSuite creates a new indexer suite for testing.
func NewIndexerSuite(tb testing.TB) *IndexerSuite {
	tb.Helper()
	return &IndexerSuite{
		TestSuite:     testutils.NewTestSuite(tb),
		testBackend:   nil,
		deployManager: nil,
	}
}

// SetupTests sets up a test.
func (i *IndexerSuite) SetupTest() {
	i.TestSuite.SetupTest()

	i.testBackend = preset.GetRinkeby().Geth(i.GetTestContext(), i.T())
	i.deployManager = testutil.NewDeployManager(i.T())
	_, i.homeContract = i.deployManager.GetHome(i.GetTestContext(), i.testBackend)
}

func TestIndexerSuite(t *testing.T) {
	suite.Run(t, NewIndexerSuite(t))
}
