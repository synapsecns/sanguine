package destination_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/contracts/destination"
	"github.com/synapsecns/sanguine/core/contracts/origin"
	"github.com/synapsecns/sanguine/core/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/core/testutil"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/synapse-node/testutils"
	"github.com/synapsecns/synapse-node/testutils/backends"
	"github.com/synapsecns/synapse-node/testutils/backends/preset"
)

// DestinationSuite is the destination test suite.
type DestinationSuite struct {
	*testutils.TestSuite
	originContract              *origin.OriginRef
	destinationContract         *destination.DestinationRef
	destinationContractMetadata backends.DeployedContract
	attestationHarness          *attestationharness.AttestationHarnessRef
	testBackendOrigin           backends.SimulatedTestBackend
	testBackendDestination      backends.SimulatedTestBackend
	wallet                      wallet.Wallet
	signer                      signer.Signer
}

// NewDestinationSuite creates a end-to-end test suite.
func NewDestinationSuite(tb testing.TB) *DestinationSuite {
	tb.Helper()
	return &DestinationSuite{
		TestSuite: testutils.NewTestSuite(tb),
	}
}

func (d *DestinationSuite) SetupTest() {
	d.TestSuite.SetupTest()

	d.testBackendOrigin = preset.GetRinkeby().Geth(d.GetTestContext(), d.T())
	d.testBackendDestination = preset.GetBSCTestnet().Geth(d.GetTestContext(), d.T())
	deployManager := testutil.NewDeployManager(d.T())

	_, d.originContract = deployManager.GetOrigin(d.GetTestContext(), d.testBackendOrigin)
	_, d.attestationHarness = deployManager.GetAttestationHarness(d.GetTestContext(), d.testBackendOrigin)
	d.destinationContractMetadata, d.destinationContract = deployManager.GetDestination(d.GetTestContext(), d.testBackendDestination)

	var err error
	d.wallet, err = wallet.FromRandom()
	if err != nil {
		d.T().Fatal(err)
	}

	_, notaryManager := deployManager.GetNotaryManager(d.GetTestContext(), d.testBackendOrigin)
	owner, err := notaryManager.Owner(&bind.CallOpts{Context: d.GetTestContext()})
	if err != nil {
		d.T().Fatal(err)
	}

	d.signer = localsigner.NewSigner(d.wallet.PrivateKey())
	d.testBackendOrigin.FundAccount(d.GetTestContext(), d.signer.Address(), *big.NewInt(params.Ether))
	d.testBackendDestination.FundAccount(d.GetTestContext(), d.signer.Address(), *big.NewInt(params.Ether))

	transactOpts := d.testBackendOrigin.GetTxContext(d.GetTestContext(), &owner)

	tx, err := notaryManager.SetNotary(transactOpts.TransactOpts, d.signer.Address())
	if err != nil {
		d.T().Fatal(err)
	}

	d.testBackendOrigin.WaitForConfirmation(d.GetTestContext(), tx)
}

// TestDestinationSuite runs the integration test suite.
func TestDestinationSuite(t *testing.T) {
	suite.Run(t, NewDestinationSuite(t))
}
