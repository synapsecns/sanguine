package preflight_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// PreflightSuite runs tests on the proxy package.
type PreflightSuite struct {
	*testsuite.TestSuite
}

func NewPreflightSuite(tb testing.TB) *PreflightSuite {
	tb.Helper()
	return &PreflightSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestPreflightSuite(t *testing.T) {
	suite.Run(t, NewPreflightSuite(t))
}
