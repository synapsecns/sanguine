package prom_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// MetricsSuite is the Metrics test suite.
type MetricsSuite struct {
	*testsuite.TestSuite
}

// NewMetricsSuite creates a end-to-end test suite.
func NewMetricsSuite(tb testing.TB) *MetricsSuite {
	tb.Helper()
	return &MetricsSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

// TestMetricsSuite runs the integration test suite.
func TestMetricsSuite(t *testing.T) {
	suite.Run(t, NewMetricsSuite(t))
}
