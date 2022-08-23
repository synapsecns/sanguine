package geth_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type GethSuite struct {
	*testsuite.TestSuite
}

// NewGethSuite creates a end-to-end test suite.
func NewGethSuite(tb testing.TB) *GethSuite {
	tb.Helper()
	return &GethSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestGethSuite(t *testing.T) {
	suite.Run(t, NewGethSuite(t))
}
