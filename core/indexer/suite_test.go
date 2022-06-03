package indexer_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"github.com/synapsecns/sanguine/core/testutil"
	"github.com/synapsecns/synapse-node/testutils"
	"github.com/synapsecns/synapse-node/testutils/backends"
	"github.com/synapsecns/synapse-node/testutils/backends/simulated"
	"testing"
)

type IndexerSuite struct {
	*testutils.TestSuite
	homeContract *home.HomeRef
	testBackend  backends.SimulatedTestBackend
}

func NewIndexerSuite(tb testing.TB) *IndexerSuite {
	tb.Helper()
	return &IndexerSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func (i *IndexerSuite) SetupTest() {
	i.TestSuite.SetupTest()

	deployManager := testutil.NewDeployManager(i.T())
	i.testBackend = simulated.NewSimulatedBackend(i.GetTestContext(), i.T())

	_, i.homeContract = deployManager.GetHome(i.GetTestContext(), i.testBackend)
}

func TestIndexerSuite(t *testing.T) {
	suite.Run(t, NewIndexerSuite(t))
}
