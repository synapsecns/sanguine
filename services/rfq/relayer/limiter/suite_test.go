package limiter_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
)

// Server suite is the main API server test suite.
type LimiterSuite struct {
	*testsuite.TestSuite
}

// NewServerSuite creates a end-to-end test suite.
func NewLimiterSuite(tb testing.TB) *LimiterSuite {
	tb.Helper()
	return &LimiterSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (s *LimiterSuite) SetupTest() {
	s.TestSuite.SetupTest()
	// Setup

}

func (s *LimiterSuite) SetupSuite() {
	s.TestSuite.SetupSuite()
}

// TestConfigSuite runs the integration test suite.
func TestLimiterSuite(t *testing.T) {
	suite.Run(t, NewLimiterSuite(t))
}
