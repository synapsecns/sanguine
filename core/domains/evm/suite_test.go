package evm_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"github.com/synapsecns/sanguine/core/domains/evm"
	"github.com/synapsecns/sanguine/core/testutil"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/synapse-node/testutils"
	"github.com/synapsecns/synapse-node/testutils/backends"
	"github.com/synapsecns/synapse-node/testutils/backends/preset"
	"github.com/synapsecns/synapse-node/testutils/backends/simulated"
	"math/big"
	"testing"
	"time"
)

// RPCSuite defines a suite where we need live rpc endpoints (as opposed to a simulated backend) to test.
type RPCSuite struct {
	*testutils.TestSuite
	testBackend   backends.TestBackend
	deployManager *testutil.DeployManager
}

// NewRPCSuite creates a new chain testing suite.
func NewRPCSuite(tb testing.TB) *RPCSuite {
	tb.Helper()
	return &RPCSuite{TestSuite: testutils.NewTestSuite(tb)}
}

func (e *RPCSuite) SetupTest() {
	evm.SetMinBackoff(time.Duration(0))
	evm.SetMaxBackoff(time.Duration(0))

	e.TestSuite.SetupTest()
	e.testBackend = preset.GetRinkeby().Geth(e.GetTestContext(), e.T())
	e.deployManager = testutil.NewDeployManager(e.T())
}

func TestEVMSuite(t *testing.T) {
	suite.Run(t, NewRPCSuite(t))
}

// ContractSuite defines a suite for testing contracts. This uses the simulated backend.
type ContractSuite struct {
	*testutils.TestSuite
	homeContract        *home.HomeRef
	attestationContract *attestationcollector.AttestationCollectorRef
	testBackend         backends.SimulatedTestBackend
	attestationBackend  backends.SimulatedTestBackend
	signer              signer.Signer
}

func NewContractSuite(tb testing.TB) *ContractSuite {
	tb.Helper()
	return &ContractSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

const attestationDomain = 4

func (i *ContractSuite) SetupTest() {
	i.TestSuite.SetupTest()

	deployManager := testutil.NewDeployManager(i.T())
	i.testBackend = simulated.NewSimulatedBackendWithChainID(i.GetTestContext(), i.T(), big.NewInt(1))
	i.attestationBackend = simulated.NewSimulatedBackendWithChainID(i.GetTestContext(), i.T(), big.NewInt(2))

	_, i.homeContract = deployManager.GetHome(i.GetTestContext(), i.testBackend)

	var attestationContract backends.DeployedContract
	attestationContract, i.attestationContract = deployManager.GetAttestationCollector(i.GetTestContext(), i.attestationBackend)

	wall, err := wallet.FromRandom()
	Nil(i.T(), err)

	i.signer = localsigner.NewSigner(wall.PrivateKey())
	i.testBackend.FundAccount(i.GetTestContext(), wall.Address(), *big.NewInt(params.Ether))
	i.attestationBackend.FundAccount(i.GetTestContext(), wall.Address(), *big.NewInt(params.Ether))

	// change the updater as defined by the update manager contract
	_, updaterManager := deployManager.GetUpdaterManager(i.GetTestContext(), i.testBackend)
	owner, err := updaterManager.Owner(&bind.CallOpts{Context: i.GetTestContext()})
	Nil(i.T(), err)

	transactOpts := i.testBackend.GetTxContext(i.GetTestContext(), &owner)

	// set the signer address to the updater
	tx, err := updaterManager.SetUpdater(transactOpts.TransactOpts, i.signer.Address())
	Nil(i.T(), err)
	i.testBackend.WaitForConfirmation(i.GetTestContext(), tx)

	// add the updater to attestation contract
	auth := i.attestationBackend.GetTxContext(i.GetTestContext(), attestationContract.OwnerPtr())

	tx, err = i.attestationContract.AddUpdater(auth.TransactOpts, attestationDomain, i.signer.Address())
	Nil(i.T(), err)
	i.attestationBackend.WaitForConfirmation(i.GetTestContext(), tx)
}

func TestContractSuite(t *testing.T) {
	suite.Run(t, NewContractSuite(t))
}

// TxQueueSuite tests out the transaction queue.
type TxQueueSuite struct {
	*testutils.TestSuite
}

// NewQueueSuite creates the queue.
func NewQueueSuite(tb testing.TB) *TxQueueSuite {
	tb.Helper()

	return &TxQueueSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func TestQueueSuite(t *testing.T) {
	suite.Run(t, NewQueueSuite(t))
}
