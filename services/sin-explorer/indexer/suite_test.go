package explorer_test

import (
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/services/sin-explorer/metadata"
	"github.com/synapsecns/sanguine/services/sin-explorer/testutil"
	"testing"
)

// IndexerSuite is a test suite for the indexer.
type IndexerSuite struct {
	*testsuite.TestSuite
	metrics       metrics.Handler
	backend       backends.SimulatedTestBackend
	deployManager *testutil.DeployManager
}

// NewIndexerSuite creates a new IndexerSuite.
func NewIndexerSuite(tb testing.TB) *IndexerSuite {
	tb.Helper()
	return &IndexerSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (i *IndexerSuite) SetupSuite() {
	i.TestSuite.SetupSuite()

	var err error
	i.metrics, err = metrics.NewByType(i.GetSuiteContext(), metadata.BuildInfo(), metrics.Null)
	i.NoError(err)
}

func (i *IndexerSuite) SetupTest() {
	i.TestSuite.SetupTest()

	i.backend = geth.NewEmbeddedBackend(i.GetTestContext(), i.T())
	i.deployManager = testutil.NewDeployManager(i.T())
}

func (i *IndexerSuite) TestIndexer() {
	_, eventMock := i.deployManager.GetInterchainClientEventMock(i.GetTestContext(), i.backend)
	_ = eventMock
}
