package evm_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/testutil"
)

// RPCSuite defines a suite where we need live rpc endpoints (as opposed to a simulated backend) to test.
type RPCSuite struct {
	*testutil.SimulatedBackendsTestSuite
}

// NewRPCSuite creates a new chain testing suite.
func NewRPCSuite(tb testing.TB) *RPCSuite {
	tb.Helper()
	return &RPCSuite{SimulatedBackendsTestSuite: testutil.NewSimulatedBackendsTestSuite(tb)}
}

func (e *RPCSuite) SetupTest() {
	evm.SetMinBackoff(time.Duration(0))
	evm.SetMaxBackoff(time.Duration(0))

	e.SimulatedBackendsTestSuite.SetupTest()
}

func TestEVMSuite(t *testing.T) {
	suite.Run(t, NewRPCSuite(t))
}

// ContractSuite defines a suite for testing contracts. This uses the simulated backend.
type ContractSuite struct {
	*testutil.SimulatedBackendsTestSuite
}

func NewContractSuite(tb testing.TB) *ContractSuite {
	tb.Helper()
	return &ContractSuite{
		SimulatedBackendsTestSuite: testutil.NewSimulatedBackendsTestSuite(tb),
	}
}

func (i *ContractSuite) SetupTest() {
	i.SimulatedBackendsTestSuite.SetupTest()
}

func TestContractSuite(t *testing.T) {
	suite.Run(t, NewContractSuite(t))
}

// TxQueueSuite tests out the transaction queue.
type TxQueueSuite struct {
	*testutil.SimulatedBackendsTestSuite
}

// NewQueueSuite creates the queue.
func NewQueueSuite(tb testing.TB) *TxQueueSuite {
	tb.Helper()

	return &TxQueueSuite{
		SimulatedBackendsTestSuite: testutil.NewSimulatedBackendsTestSuite(tb),
	}
}

func (t *TxQueueSuite) SetupTest() {
	t.SimulatedBackendsTestSuite.SetupTest()
}

func TestQueueSuite(t *testing.T) {
	suite.Run(t, NewQueueSuite(t))
}
