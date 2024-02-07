package proxy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/omnirpc/metadata"
)

type ProxySuite struct {
	*testsuite.TestSuite
	metrics metrics.Handler
}

// NewProxySuite creates a end-to-end test suite.
func NewProxySuite(tb testing.TB) *ProxySuite {
	tb.Helper()
	return &ProxySuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (p *ProxySuite) SetupSuite() {
	p.TestSuite.SetupSuite()

	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(p.GetSuiteContext(), p.T())
		metricsHandler = metrics.Jaeger
	}

	var err error
	p.metrics, err = metrics.NewByType(p.GetSuiteContext(), metadata.BuildInfo(), metricsHandler)
	assert.Nil(p.T(), err)
}

func TestProxySuite(t *testing.T) {
	suite.Run(t, NewProxySuite(t))
}
