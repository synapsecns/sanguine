package ganache_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/synapse-node/testutils"
	"testing"
)

type GanacheSuite struct {
	*testutils.TestSuite
}

// NewGanacheSuite creates a end-to-end test suite.
func NewGanacheSuite(tb testing.TB) *GanacheSuite {
	tb.Helper()
	return &GanacheSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func TestTestUtilSuite(t *testing.T) {
	suite.Run(t, NewGanacheSuite(t))
}
