package api_test

import (
	"testing"

	"github.com/Flaque/filet"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/cctp-relayer/db/sql/base"
	"github.com/synapsecns/sanguine/services/cctp-relayer/db/sql/sqlite"
	"github.com/synapsecns/sanguine/services/cctp-relayer/metadata"
)

// RelayerAPISuite defines the basic test suite.
type RelayerAPISuite struct {
	*testsuite.TestSuite
	// metricsHandler is the metrics handler for the test
	metricsHandler metrics.Handler
	// testStore is the test store for the test
	testStore *base.Store
}

// NewTestSuite creates a new test suite.
func NewTestSuite(tb testing.TB) *RelayerAPISuite {
	tb.Helper()
	return &RelayerAPISuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (s *RelayerAPISuite) SetupSuite() {
	s.TestSuite.SetupSuite()
	// for tracing
	localmetrics.SetupTestJaeger(s.GetSuiteContext(), s.T())
}

func (s *RelayerAPISuite) SetupTest() {
	s.TestSuite.SetupTest()

	// create the test metrics handler
	var err error
	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(s.GetSuiteContext(), s.T())
		metricsHandler = metrics.Jaeger
	}
	s.metricsHandler, err = metrics.NewByType(s.GetTestContext(), metadata.BuildInfo(), metricsHandler)
	s.Require().NoError(err)

	// create the test store
	path := filet.TmpDir(s.T(), "")
	db, err := sqlite.NewSqliteStore(s.GetTestContext(), path, s.metricsHandler, false)
	s.Require().NoError(err)
	s.testStore = base.NewStore(db.DB(), s.metricsHandler)
}

func TestRelayerAPISuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
