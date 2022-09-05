package client_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	"testing"
)

type ClientSuite struct {
	*testsuite.TestSuite
}

// NewClientSuite creates a end-to-end test suite.
func NewClientSuite(tb testing.TB) *ClientSuite {
	tb.Helper()
	return &ClientSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, NewClientSuite(t))
}

func (c ClientSuite) SetupTest() {
	c.TestSuite.SetupTest()
	client.SetResetTimeout(client.GetDefaultResetTimeout())
}
