package attestationcollector_test

import (
	"math/big"
	"testing"

	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/preset"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

// AttestationCollectorSuite is the attestation collector test suite.
type AttestationCollectorSuite struct {
	*testsuite.TestSuite
	originContract              *origin.OriginRef
	destinationContract         *destination.DestinationRef
	destinationContractMetadata contracts.DeployedContract
	attestationHarness          *attestationharness.AttestationHarnessRef
	attestationContract         *attestationcollector.AttestationCollectorRef
	attestationContractMetadata contracts.DeployedContract
	testBackendOrigin           backends.SimulatedTestBackend
	testBackendDestination      backends.SimulatedTestBackend
	testBackendAttestation      backends.SimulatedTestBackend
	notaryWallet                wallet.Wallet
	guardWallet                 wallet.Wallet
	notarySigner                signer.Signer
	guardSigner                 signer.Signer
}

// NewAttestationCollectorSuite creates an end-to-end test suite.
func NewAttestationCollectorSuite(tb testing.TB) *AttestationCollectorSuite {
	tb.Helper()
	return &AttestationCollectorSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

// SetupTest sets up the test.
func (a *AttestationCollectorSuite) SetupTest() {
	a.TestSuite.SetupTest()

	a.testBackendOrigin = preset.GetRinkeby().Geth(a.GetTestContext(), a.T())
	a.testBackendDestination = preset.GetBSCTestnet().Geth(a.GetTestContext(), a.T())
	a.testBackendAttestation = preset.GetMaticMumbai().Geth(a.GetTestContext(), a.T())
	deployManager := testutil.NewDeployManager(a.T())

	_, a.originContract = deployManager.GetOrigin(a.GetTestContext(), a.testBackendOrigin)
	_, a.attestationHarness = deployManager.GetAttestationHarness(a.GetTestContext(), a.testBackendOrigin)
	a.attestationContractMetadata, a.attestationContract = deployManager.GetAttestationCollector(a.GetTestContext(), a.testBackendAttestation)
	a.destinationContractMetadata, a.destinationContract = deployManager.GetDestination(a.GetTestContext(), a.testBackendDestination)

	var err error
	a.notaryWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}
	a.guardWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}

	a.notarySigner = localsigner.NewSigner(a.notaryWallet.PrivateKey())
	a.testBackendOrigin.FundAccount(a.GetTestContext(), a.notarySigner.Address(), *big.NewInt(params.Ether))
	a.testBackendDestination.FundAccount(a.GetTestContext(), a.notarySigner.Address(), *big.NewInt(params.Ether))
	a.testBackendAttestation.FundAccount(a.GetTestContext(), a.notarySigner.Address(), *big.NewInt(params.Ether))
	a.guardSigner = localsigner.NewSigner(a.guardWallet.PrivateKey())
	a.testBackendOrigin.FundAccount(a.GetTestContext(), a.guardSigner.Address(), *big.NewInt(params.Ether))
	a.testBackendDestination.FundAccount(a.GetTestContext(), a.guardSigner.Address(), *big.NewInt(params.Ether))
	a.testBackendAttestation.FundAccount(a.GetTestContext(), a.guardSigner.Address(), *big.NewInt(params.Ether))

	destOwnerPtr, err := a.destinationContract.DestinationCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}

	destOwnerAuth := a.testBackendDestination.GetTxContext(a.GetTestContext(), &destOwnerPtr)
	txDestinationNotaryAdd, err := a.destinationContract.AddNotary(destOwnerAuth.TransactOpts, uint32(a.testBackendDestination.GetChainID()), a.notarySigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.testBackendDestination.WaitForConfirmation(a.GetTestContext(), txDestinationNotaryAdd)
	txDestinationGuardAdd, err := a.destinationContract.AddGuard(destOwnerAuth.TransactOpts, a.guardSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.testBackendDestination.WaitForConfirmation(a.GetTestContext(), txDestinationGuardAdd)

	originOwnerPtr, err := a.originContract.OriginCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}
	originOwnerAuth := a.testBackendOrigin.GetTxContext(a.GetTestContext(), &originOwnerPtr)

	txOriginNotaryAdd, err := a.originContract.AddNotary(originOwnerAuth.TransactOpts, uint32(a.testBackendDestination.GetChainID()), a.notarySigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.testBackendOrigin.WaitForConfirmation(a.GetTestContext(), txOriginNotaryAdd)
	txOriginGuardAdd, err := a.originContract.AddGuard(originOwnerAuth.TransactOpts, a.guardSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.testBackendOrigin.WaitForConfirmation(a.GetTestContext(), txOriginGuardAdd)

	attestOwnerPtr, err := a.attestationContract.AttestationCollectorCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}
	attestOwnerAuth := a.testBackendAttestation.GetTxContext(a.GetTestContext(), &attestOwnerPtr)

	txAddNotary, err := a.attestationContract.AddNotary(attestOwnerAuth.TransactOpts, uint32(a.testBackendDestination.GetChainID()), a.notarySigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.testBackendAttestation.WaitForConfirmation(a.GetTestContext(), txAddNotary)
	txAddGuard, err := a.attestationContract.AddGuard(attestOwnerAuth.TransactOpts, a.guardSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.testBackendAttestation.WaitForConfirmation(a.GetTestContext(), txAddGuard)
}

// TestAttestationCollectorSuite runs the integration test suite.
func TestAttestationCollectorSuite(t *testing.T) {
	suite.Run(t, NewAttestationCollectorSuite(t))
}
