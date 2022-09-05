package near_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type NearSuite struct {
	*testsuite.TestSuite
}

// NewNearSuite creates a end-to-end test suite.
func NewNearSuite(tb testing.TB) *NearSuite {
	tb.Helper()
	return &NearSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestNearSuite(t *testing.T) {
	suite.Run(t, NewNearSuite(t))
}

func (c NearSuite) SetupTest() {
	c.TestSuite.SetupTest()
}
