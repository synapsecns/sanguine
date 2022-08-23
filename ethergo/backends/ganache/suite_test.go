package ganache_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type GanacheSuite struct {
	*testsuite.TestSuite
}

// NewGanacheSuite creates a end-to-end test suite.
func NewGanacheSuite(tb testing.TB) *GanacheSuite {
	tb.Helper()
	return &GanacheSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestTestUtilSuite(t *testing.T) {
	suite.Run(t, NewGanacheSuite(t))
}
