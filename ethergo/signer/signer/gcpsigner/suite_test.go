package gcpsigner_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type GCPSignerSuite struct {
	*testsuite.TestSuite
}

// NewGCPSignerSuite creates a end-to-end test suite.
func NewGCPSignerSuite(tb testing.TB) *GCPSignerSuite {
	tb.Helper()
	return &GCPSignerSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestGCPSignerSuite(t *testing.T) {
	suite.Run(t, NewGCPSignerSuite(t))
}
