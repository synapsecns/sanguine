package proxy_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type ProxySuite struct {
	*testsuite.TestSuite
}

// NewProxySuite creates a end-to-end test suite.
func NewProxySuite(tb testing.TB) *ProxySuite {
	tb.Helper()
	return &ProxySuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestProxySuite(t *testing.T) {
	suite.Run(t, NewProxySuite(t))
}
