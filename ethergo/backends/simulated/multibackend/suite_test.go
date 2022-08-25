package multibackend_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type MultiBackendSuite struct {
	*testsuite.TestSuite
}

// NewMultiBackendSuite creates a end-to-end test suite.
func NewMultiBackendSuite(tb testing.TB) *MultiBackendSuite {
	tb.Helper()
	return &MultiBackendSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestMultiBackendSuite(t *testing.T) {
	suite.Run(t, NewMultiBackendSuite(t))
}
