package balance_test

import (
	"sync/atomic"
	"testing"

	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/metadata"

	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
)

type BalanceSuite struct {
	*testsuite.TestSuite
	logIndex atomic.Int64 // For thread safety
	metrics  metrics.Handler
}

// NewBalanceSuiteSuite creates a new BalanceSuite.
func NewBalanceSuiteSuite(tb testing.TB) *BalanceSuite {
	tb.Helper()
	return &BalanceSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

// SetupSuite sets up the rest of the test suite.
func (t *BalanceSuite) SetupSuite() {
	t.TestSuite.SetupSuite()
	t.logIndex.Store(0)

	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(t.GetSuiteContext(), t.T())
		metricsHandler = metrics.Jaeger
	}
	var err error
	t.metrics, err = metrics.NewByType(t.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	Nil(t.T(), err)
}

// TestDBSuite tests the Balance suite.
func TestBalanceSuiteSuite(t *testing.T) {
	suite.Run(t, NewBalanceSuiteSuite(t))
}
