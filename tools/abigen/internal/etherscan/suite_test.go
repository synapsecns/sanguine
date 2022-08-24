package etherscan_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

type EtherscanSuite struct {
	*testsuite.TestSuite
}

// NewEtherscanSuite creates a end-to-end test suite.
func NewEtherscanSuite(tb testing.TB) *EtherscanSuite {
	tb.Helper()
	return &EtherscanSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestEtherscanSuite(t *testing.T) {
	suite.Run(t, NewEtherscanSuite(t))
}
