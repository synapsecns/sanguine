package replicamanager_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/contracts/home"
	"github.com/synapsecns/sanguine/core/contracts/replicamanager"
	"github.com/synapsecns/sanguine/core/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/core/testutil"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/synapse-node/testutils"
	"github.com/synapsecns/synapse-node/testutils/backends"
	"github.com/synapsecns/synapse-node/testutils/backends/preset"
)

// ReplicaManagerSuite is the home test suite.
type ReplicaManagerSuite struct {
	*testutils.TestSuite
	homeContract            *home.HomeRef
	replicaContract         *replicamanager.ReplicaManagerRef
	replicaContractMetadata backends.DeployedContract
	attestationHarness      *attestationharness.AttestationHarnessRef
	testBackendHome         backends.SimulatedTestBackend
	testBackendReplica      backends.SimulatedTestBackend
	wallet                  wallet.Wallet
	signer                  signer.Signer
}

// NewHomeSuite creates a end-to-end test suite.
func NewHomeSuite(tb testing.TB) *ReplicaManagerSuite {
	tb.Helper()
	return &ReplicaManagerSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func (r *ReplicaManagerSuite) SetupTest() {
	r.TestSuite.SetupTest()

	r.testBackendHome = preset.GetRinkeby().Geth(r.GetTestContext(), r.T())
	r.testBackendReplica = preset.GetBSCTestnet().Geth(r.GetTestContext(), r.T())
	deployManager := testutil.NewDeployManager(r.T())

	_, r.homeContract = deployManager.GetHome(r.GetTestContext(), r.testBackendHome)
	_, r.attestationHarness = deployManager.GetAttestationHarness(r.GetTestContext(), r.testBackendHome)
	r.replicaContractMetadata, r.replicaContract = deployManager.GetReplicaManager(r.GetTestContext(), r.testBackendReplica)

	var err error
	r.wallet, err = wallet.FromRandom()
	if err != nil {
		r.T().Fatal(err)
	}

	_, updaterManager := deployManager.GetUpdaterManager(r.GetTestContext(), r.testBackendHome)
	owner, err := updaterManager.Owner(&bind.CallOpts{Context: r.GetTestContext()})
	if err != nil {
		r.T().Fatal(err)
	}

	r.signer = localsigner.NewSigner(r.wallet.PrivateKey())
	r.testBackendHome.FundAccount(r.GetTestContext(), r.signer.Address(), *big.NewInt(params.Ether))
	r.testBackendReplica.FundAccount(r.GetTestContext(), r.signer.Address(), *big.NewInt(params.Ether))

	transactOpts := r.testBackendHome.GetTxContext(r.GetTestContext(), &owner)

	tx, err := updaterManager.SetUpdater(transactOpts.TransactOpts, r.signer.Address())
	if err != nil {
		r.T().Fatal(err)
	}

	r.testBackendHome.WaitForConfirmation(r.GetTestContext(), tx)
}

// TestHomeSuite runs the integration test suite.
func TestHomeSuite(t *testing.T) {
	suite.Run(t, NewHomeSuite(t))
}
