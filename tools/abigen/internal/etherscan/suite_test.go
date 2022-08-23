package etherscan_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/synapse-node/testutils"
	"testing"
)

type EtherscanSuite struct {
	*testutils.TestSuite
}

// NewEtherscanSuite creates a end-to-end test suite.
func NewEtherscanSuite(tb testing.TB) *EtherscanSuite {
	tb.Helper()
	return &EtherscanSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func TestEtherscanSuite(t *testing.T) {
	suite.Run(t, NewEtherscanSuite(t))
}
