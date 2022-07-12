package kmsmock_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/synapse-node/testutils"
	"testing"
)

type KMSSuite struct {
	*testutils.TestSuite
}

// NewKMSSuite creates a end-to-end test suite.
func NewKMSSuite(tb testing.TB) *KMSSuite {
	tb.Helper()
	return &KMSSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func TestKMSSuite(t *testing.T) {
	suite.Run(t, NewKMSSuite(t))
}
