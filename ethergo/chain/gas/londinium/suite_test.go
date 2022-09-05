package londinium_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type LondoniumSuite struct {
	*testsuite.TestSuite
}

// NewLondoniumSuite creates a end-to-end test suite.
func NewLondoniumSuite(tb testing.TB) *LondoniumSuite {
	tb.Helper()
	return &LondoniumSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestLondoniumSuite(t *testing.T) {
	suite.Run(t, NewLondoniumSuite(t))
}
