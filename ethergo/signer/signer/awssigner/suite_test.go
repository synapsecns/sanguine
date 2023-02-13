package awssigner_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type KMSSuite struct {
	*testsuite.TestSuite
}

// NewKMSSuite creates a end-to-end test suite.
func NewKMSSuite(tb testing.TB) *KMSSuite {
	tb.Helper()
	return &KMSSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestKMSSuite(t *testing.T) {
	suite.Run(t, NewKMSSuite(t))
}
