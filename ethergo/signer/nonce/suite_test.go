package nonce_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/synapse-node/testutils"
	"testing"
)

// NonceSuite is the nonce suite.
type NonceSuite struct {
	*testutils.TestSuite
}

// NewNonceSuite creates a end-to-end test suite.
func NewNonceSuite(tb testing.TB) *NonceSuite {
	tb.Helper()
	return &NonceSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func TestNonceSuite(t *testing.T) {
	suite.Run(t, NewNonceSuite(t))
}
