package indexer_test

import (
	"testing"
	"time"

	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/services/scribe/metadata"

	"github.com/Flaque/filet"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/services/scribe/testutil"
)

type IndexerSuite struct {
	*testsuite.TestSuite
	testDB  db.EventDB
	manager *testutil.DeployManager
	wallet  wallet.Wallet
	signer  *localsigner.Signer
	metrics metrics.Handler
}

// NewIndexerSuite creates a new indexer test suite.
func NewIndexerSuite(tb testing.TB) *IndexerSuite {
	tb.Helper()
	return &IndexerSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

// SetupTest sets up the test suite.
func (x *IndexerSuite) SetupTest() {
	x.TestSuite.SetupTest()
	x.SetTestTimeout(time.Minute * 10)
	sqliteStore, err := sqlite.NewSqliteStore(x.GetTestContext(), filet.TmpDir(x.T(), ""), x.metrics, false)
	Nil(x.T(), err)
	x.testDB = sqliteStore
	x.manager = testutil.NewDeployManager(x.T())
	x.wallet, err = wallet.FromRandom()
	Nil(x.T(), err)
	x.signer = localsigner.NewSigner(x.wallet.PrivateKey())
}

func (x *IndexerSuite) SetupSuite() {
	x.TestSuite.SetupSuite()

	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(x.GetSuiteContext(), x.T())
		metricsHandler = metrics.Jaeger
	}

	var err error
	x.metrics, err = metrics.NewByType(x.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	Nil(x.T(), err)
}

// TestIndexerSuite tests the indexer suite.
func TestIndexerSuite(t *testing.T) {
	suite.Run(t, NewIndexerSuite(t))
}
