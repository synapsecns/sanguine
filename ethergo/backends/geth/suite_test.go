package geth_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/synapse-node/testutils"
	"testing"
)

type GethSuite struct {
	*testutils.TestSuite
}

// NewGethSuite creates a end-to-end test suite.
func NewGethSuite(tb testing.TB) *GethSuite {
	tb.Helper()
	return &GethSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func TestGethSuite(t *testing.T) {
	suite.Run(t, NewGethSuite(t))
}
