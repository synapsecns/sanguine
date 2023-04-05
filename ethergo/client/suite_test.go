package client_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// ClientSuite defines the basic test suite.
type ClientSuite struct {
	*testsuite.TestSuite
}

// NewClientSuite creates a new test suite and performs some basic checks afterward.
// Every test suite in the synapse library should inherit from this suite and override where necessary.
func NewClientSuite(tb testing.TB) *ClientSuite {
	tb.Helper()
	return &ClientSuite{
		testsuite.NewTestSuite(tb),
	}
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, NewClientSuite(t))
}
