package anvil_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type AnvilSuite struct {
	*testsuite.TestSuite
}

// NewAnvilSuite creates a end-to-end test suite.
func NewAnvilSuite(tb testing.TB) *AnvilSuite {
	tb.Helper()
	return &AnvilSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestTestUtilSuite(t *testing.T) {
	suite.Run(t, NewAnvilSuite(t))
}
