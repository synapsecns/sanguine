package multibackend_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/synapse-node/testutils"
	"testing"
)

type MultiBackendSuite struct {
	*testutils.TestSuite
}

// NewMultiBackendSuite creates a end-to-end test suite.
func NewMultiBackendSuite(tb testing.TB) *MultiBackendSuite {
	tb.Helper()
	return &MultiBackendSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func TestMultiBackendSuite(t *testing.T) {
	suite.Run(t, NewMultiBackendSuite(t))
}
