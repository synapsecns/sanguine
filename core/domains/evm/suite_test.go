package evm_test

import (
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/synapse-node/testutils"
	"github.com/synapsecns/synapse-node/testutils/backends"
	"github.com/synapsecns/synapse-node/testutils/backends/preset"
	"testing"
)

// EVMSuite defines the basic chain suite.
type EVMSuite struct {
	*testutils.TestSuite
	testBackend backends.TestBackend
}

// NewEVMSuite creates a new chain testing suite.
func NewEVMSuite(tb testing.TB) *EVMSuite {
	tb.Helper()
	return &EVMSuite{TestSuite: testutils.NewTestSuite(tb)}
}

func (e *EVMSuite) SetupTest() {
	e.testBackend = preset.GetRinkeby().Geth(e.GetTestContext(), e.T())
}

func TestEVMSuite(t *testing.T) {
	suite.Run(t, NewEVMSuite(t))
}
