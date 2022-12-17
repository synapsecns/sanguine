package evm_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/preset"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

// RPCSuite defines a suite where we need live rpc endpoints (as opposed to a simulated backend) to test.
type RPCSuite struct {
	*testsuite.TestSuite
	testBackend   backends.TestBackend
	deployManager *testutil.DeployManager
}

// NewRPCSuite creates a new chain testing suite.
func NewRPCSuite(tb testing.TB) *RPCSuite {
	tb.Helper()
	return &RPCSuite{TestSuite: testsuite.NewTestSuite(tb)}
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
	*testsuite.TestSuite
	originContract      *originharness.OriginHarnessRef
	attestationContract *attestationcollector.AttestationCollectorRef
	testBackend         backends.SimulatedTestBackend
	attestationBackend  backends.SimulatedTestBackend
	signer              signer.Signer
}

func NewContractSuite(tb testing.TB) *ContractSuite {
	tb.Helper()
	return &ContractSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

const attestationDomain = uint32(4)
const testDestinationDomain = attestationDomain + 1

func (i *ContractSuite) SetupTest() {
	i.TestSuite.SetupTest()

	deployManager := testutil.NewDeployManager(i.T())
	i.testBackend = simulated.NewSimulatedBackendWithChainID(i.GetTestContext(), i.T(), big.NewInt(1))
	i.attestationBackend = simulated.NewSimulatedBackendWithChainID(i.GetTestContext(), i.T(), big.NewInt(2))

	_, i.originContract = deployManager.GetOriginHarness(i.GetTestContext(), i.testBackend)

	var attestationContract contracts.DeployedContract
	attestationContract, i.attestationContract = deployManager.GetAttestationCollector(i.GetTestContext(), i.attestationBackend)

	wall, err := wallet.FromRandom()
	Nil(i.T(), err)

	i.signer = localsigner.NewSigner(wall.PrivateKey())
	i.testBackend.FundAccount(i.GetTestContext(), wall.Address(), *big.NewInt(params.Ether))
	i.attestationBackend.FundAccount(i.GetTestContext(), wall.Address(), *big.NewInt(params.Ether))

	// add the notary to attestation contract
	auth := i.attestationBackend.GetTxContext(i.GetTestContext(), attestationContract.OwnerPtr())

	tx, err := i.attestationContract.AddAgent(auth.TransactOpts, testDestinationDomain, i.signer.Address())
	Nil(i.T(), err)
	i.attestationBackend.WaitForConfirmation(i.GetTestContext(), tx)

	ownerPtr, err := i.originContract.OriginHarnessCaller.Owner(&bind.CallOpts{Context: i.GetTestContext()})
	Nil(i.T(), err)

	originOwnerAuth := i.testBackend.GetTxContext(i.GetTestContext(), &ownerPtr)

	notaries, err := i.originContract.AllAgents(&bind.CallOpts{Context: i.GetTestContext()}, destinationID)
	Nil(i.T(), err)
	Len(i.T(), notaries, 0)

	tx, err = i.originContract.AddAgent(originOwnerAuth.TransactOpts, destinationID, i.signer.Address())
	Nil(i.T(), err)
	i.testBackend.WaitForConfirmation(i.GetTestContext(), tx)

	notaries, err = i.originContract.AllAgents(&bind.CallOpts{Context: i.GetTestContext()}, destinationID)
	Nil(i.T(), err)
	Len(i.T(), notaries, 1)
}

func TestContractSuite(t *testing.T) {
	suite.Run(t, NewContractSuite(t))
}

// TxQueueSuite tests out the transaction queue.
type TxQueueSuite struct {
	*testsuite.TestSuite
	chn            backends.SimulatedTestBackend
	originContract *originharness.OriginHarnessRef
	testTransactor *bind.TransactOpts
	destinationID  uint32
}

// NewQueueSuite creates the queue.
func NewQueueSuite(tb testing.TB) *TxQueueSuite {
	tb.Helper()

	return &TxQueueSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (t *TxQueueSuite) SetupTest() {
	t.TestSuite.SetupTest()

	// create a test chain
	t.chn = simulated.NewSimulatedBackend(t.GetTestContext(), t.T())
	manager := testutil.NewDeployManager(t.T())

	t.destinationID = uint32(1)

	/*originContract*/
	_, originContractRef := manager.GetOriginHarness(t.GetTestContext(), t.chn)
	t.originContract = originContractRef

	// create a test signer
	wllt, err := wallet.FromRandom()
	Nil(t.T(), err)

	msigner := localsigner.NewSigner(wllt.PrivateKey())
	testDB, err := sqlite.NewSqliteStore(t.GetTestContext(), filet.TmpDir(t.T(), ""))
	Nil(t.T(), err)

	testQueue := evm.NewTxQueue(msigner, testDB, t.chn)

	t.testTransactor, err = testQueue.GetTransactor(t.GetTestContext(), t.chn.GetBigChainID())
	Nil(t.T(), err)

	t.chn.FundAccount(t.GetTestContext(), msigner.Address(), *big.NewInt(params.Ether))

	ownerPtr, err := t.originContract.OriginHarnessCaller.Owner(&bind.CallOpts{Context: t.GetTestContext()})
	Nil(t.T(), err)

	originOwnerAuth := t.chn.GetTxContext(t.GetTestContext(), &ownerPtr)
	tx, err := t.originContract.AddAgent(originOwnerAuth.TransactOpts, destinationID, msigner.Address())
	Nil(t.T(), err)
	t.chn.WaitForConfirmation(t.GetTestContext(), tx)

	notaries, err := t.originContract.AllAgents(&bind.CallOpts{Context: t.GetTestContext()}, destinationID)
	Nil(t.T(), err)
	Len(t.T(), notaries, 1)
}

func TestQueueSuite(t *testing.T) {
	suite.Run(t, NewQueueSuite(t))
}
