package chain_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"testing"
)

// ChainSuite defines the basic chain suite.
type ChainSuite struct {
	*testsuite.TestSuite
}

// NewChainSuite creates a new chain testing suite.
func NewChainSuite(tb testing.TB) *ChainSuite {
	tb.Helper()
	return &ChainSuite{testsuite.NewTestSuite(tb)}
}

func TestChainSuite(t *testing.T) {
	suite.Run(t, NewChainSuite(t))
}
