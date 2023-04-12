package proxy_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/omnirpc/cmd"
	"testing"
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

	localmetrics.SetupTestJaeger(p.GetSuiteContext(), p.T())

	var err error
	p.metrics, err = metrics.NewByType(p.GetSuiteContext(), cmd.BuildInfo(), metrics.Jaeger)
	assert.Nil(p.T(), err)
}

func TestProxySuite(t *testing.T) {
	suite.Run(t, NewProxySuite(t))
}
