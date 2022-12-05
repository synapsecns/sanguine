package evm_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
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
	originContract      *origin.OriginRef
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

	_, i.originContract = deployManager.GetOrigin(i.GetTestContext(), i.testBackend)

	var attestationContract contracts.DeployedContract
	attestationContract, i.attestationContract = deployManager.GetAttestationCollector(i.GetTestContext(), i.attestationBackend)

	wall, err := wallet.FromRandom()
	Nil(i.T(), err)

	i.signer = localsigner.NewSigner(wall.PrivateKey())
	i.testBackend.FundAccount(i.GetTestContext(), wall.Address(), *big.NewInt(params.Ether))
	i.attestationBackend.FundAccount(i.GetTestContext(), wall.Address(), *big.NewInt(params.Ether))

	// add the notary to attestation contract
	auth := i.attestationBackend.GetTxContext(i.GetTestContext(), attestationContract.OwnerPtr())

	tx, err := i.attestationContract.AddNotary(auth.TransactOpts, testDestinationDomain, i.signer.Address())
	Nil(i.T(), err)
	i.attestationBackend.WaitForConfirmation(i.GetTestContext(), tx)
}

func TestContractSuite(t *testing.T) {
	suite.Run(t, NewContractSuite(t))
}

// TxQueueSuite tests out the transaction queue.
type TxQueueSuite struct {
	*testsuite.TestSuite
}

// NewQueueSuite creates the queue.
func NewQueueSuite(tb testing.TB) *TxQueueSuite {
	tb.Helper()

	return &TxQueueSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestQueueSuite(t *testing.T) {
	suite.Run(t, NewQueueSuite(t))
}
