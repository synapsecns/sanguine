package nonce_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// NonceSuite is the nonce suite.
type NonceSuite struct {
	*testsuite.TestSuite
}

// NewNonceSuite creates a end-to-end test suite.
func NewNonceSuite(tb testing.TB) *NonceSuite {
	tb.Helper()
	return &NonceSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestNonceSuite(t *testing.T) {
	suite.Run(t, NewNonceSuite(t))
}
