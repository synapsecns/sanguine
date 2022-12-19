package destination_test

import (
	"math/big"
	"testing"

	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/destinationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/preset"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

// DestinationSuite is the destination test suite.
type DestinationSuite struct {
	*testsuite.TestSuite
	originContract              *originharness.OriginHarnessRef
	destinationContract         *destinationharness.DestinationHarnessRef
	destinationContractMetadata contracts.DeployedContract
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
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (d *DestinationSuite) SetupTest() {
	d.TestSuite.SetupTest()

	d.testBackendOrigin = preset.GetRinkeby().Geth(d.GetTestContext(), d.T())
	d.testBackendDestination = preset.GetBSCTestnet().Geth(d.GetTestContext(), d.T())
	deployManager := testutil.NewDeployManager(d.T())

	_, d.originContract = deployManager.GetOriginHarness(d.GetTestContext(), d.testBackendOrigin)
	_, d.attestationHarness = deployManager.GetAttestationHarness(d.GetTestContext(), d.testBackendOrigin)
	d.destinationContractMetadata, d.destinationContract = deployManager.GetDestinationHarness(d.GetTestContext(), d.testBackendDestination)

	var err error
	d.wallet, err = wallet.FromRandom()
	if err != nil {
		d.T().Fatal(err)
	}

	d.signer = localsigner.NewSigner(d.wallet.PrivateKey())
	d.testBackendOrigin.FundAccount(d.GetTestContext(), d.signer.Address(), *big.NewInt(params.Ether))
	d.testBackendDestination.FundAccount(d.GetTestContext(), d.signer.Address(), *big.NewInt(params.Ether))

	destOwnerPtr, err := d.destinationContract.DestinationHarnessCaller.Owner(&bind.CallOpts{Context: d.GetTestContext()})
	if err != nil {
		d.T().Fatal(err)
	}

	testContext := d.GetTestContext()
	destOwnerAuth := d.testBackendDestination.GetTxContext(testContext, &destOwnerPtr)
	_, err = d.destinationContract.AddAgent(destOwnerAuth.TransactOpts, uint32(d.testBackendDestination.GetChainID()), d.signer.Address())
	if err != nil {
		d.T().Fatal(err)
	}
}

// TestDestinationSuite runs the integration test suite.
func TestDestinationSuite(t *testing.T) {
	suite.Run(t, NewDestinationSuite(t))
}
