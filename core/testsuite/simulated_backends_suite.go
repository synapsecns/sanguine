package testsuite

import (
	"math/big"
	"testing"

	"github.com/synapsecns/sanguine/ethergo/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
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

// SimulatedBackendsSuite can be used as the base for any test needing simulated backends
// that have an origin, destination and attestation collector and a guard and notary
// added to each.
// TODO (joe): For tests that do not need all 3 simulated backends, allow them to pass in
// flags indicating the subset of backends desired. Some tests might only want
// an attestation collector, others might only want an origin and an attestation collector,
// others might want just a destination, etc.
type SimulatedBackendsSuite struct {
	*TestSuite
	OriginContract              *origin.OriginRef
	DestinationContract         *destination.DestinationRef
	DestinationContractMetadata contracts.DeployedContract
	AttestationHarness          *attestationharness.AttestationHarnessRef
	AttestationContract         *attestationcollector.AttestationCollectorRef
	AttestationContractMetadata contracts.DeployedContract
	TestBackendOrigin           backends.SimulatedTestBackend
	TestBackendDestination      backends.SimulatedTestBackend
	TestBackendAttestation      backends.SimulatedTestBackend
	NotaryWallet                wallet.Wallet
	GuardWallet                 wallet.Wallet
	NotarySigner                signer.Signer
	GuardSigner                 signer.Signer
}

// NewSimulatedBackendsSuite creates an end-to-end test suite with simulated
// backends set up.
func NewSimulatedBackendsSuite(tb testing.TB) *SimulatedBackendsSuite {
	tb.Helper()
	return &SimulatedBackendsSuite{
		TestSuite: NewTestSuite(tb),
	}
}

func (a *SimulatedBackendsSuite) SetupOrigin(deployManager *testutil.DeployManager) {
	_, a.OriginContract = deployManager.GetOrigin(a.GetTestContext(), a.TestBackendOrigin)
	a.TestBackendOrigin.FundAccount(a.GetTestContext(), a.NotarySigner.Address(), *big.NewInt(params.Ether))
	a.TestBackendOrigin.FundAccount(a.GetTestContext(), a.GuardSigner.Address(), *big.NewInt(params.Ether))
	originOwnerPtr, err := a.OriginContract.OriginCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}
	originOwnerAuth := a.TestBackendOrigin.GetTxContext(a.GetTestContext(), &originOwnerPtr)

	txOriginNotaryAdd, err := a.OriginContract.AddNotary(originOwnerAuth.TransactOpts, uint32(a.TestBackendDestination.GetChainID()), a.NotarySigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendOrigin.WaitForConfirmation(a.GetTestContext(), txOriginNotaryAdd)
	txOriginGuardAdd, err := a.OriginContract.AddGuard(originOwnerAuth.TransactOpts, a.GuardSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendOrigin.WaitForConfirmation(a.GetTestContext(), txOriginGuardAdd)
}

func (a *SimulatedBackendsSuite) SetupDestination(deployManager *testutil.DeployManager) {
	a.DestinationContractMetadata, a.DestinationContract = deployManager.GetDestination(a.GetTestContext(), a.TestBackendDestination)

	a.TestBackendDestination.FundAccount(a.GetTestContext(), a.NotarySigner.Address(), *big.NewInt(params.Ether))
	a.TestBackendDestination.FundAccount(a.GetTestContext(), a.GuardSigner.Address(), *big.NewInt(params.Ether))
	destOwnerPtr, err := a.DestinationContract.DestinationCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}

	destOwnerAuth := a.TestBackendDestination.GetTxContext(a.GetTestContext(), &destOwnerPtr)
	txDestinationNotaryAdd, err := a.DestinationContract.AddNotary(destOwnerAuth.TransactOpts, uint32(a.TestBackendDestination.GetChainID()), a.NotarySigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendDestination.WaitForConfirmation(a.GetTestContext(), txDestinationNotaryAdd)
	txDestinationGuardAdd, err := a.DestinationContract.AddGuard(destOwnerAuth.TransactOpts, a.GuardSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendDestination.WaitForConfirmation(a.GetTestContext(), txDestinationGuardAdd)
}

func (a *SimulatedBackendsSuite) SetupAttestation(deployManager *testutil.DeployManager) {
	_, a.AttestationHarness = deployManager.GetAttestationHarness(a.GetTestContext(), a.TestBackendAttestation)
	a.AttestationContractMetadata, a.AttestationContract = deployManager.GetAttestationCollector(a.GetTestContext(), a.TestBackendAttestation)

	a.TestBackendAttestation.FundAccount(a.GetTestContext(), a.NotarySigner.Address(), *big.NewInt(params.Ether))
	a.TestBackendAttestation.FundAccount(a.GetTestContext(), a.GuardSigner.Address(), *big.NewInt(params.Ether))

	attestOwnerPtr, err := a.AttestationContract.AttestationCollectorCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}
	attestOwnerAuth := a.TestBackendAttestation.GetTxContext(a.GetTestContext(), &attestOwnerPtr)

	txAddNotary, err := a.AttestationContract.AddNotary(attestOwnerAuth.TransactOpts, uint32(a.TestBackendDestination.GetChainID()), a.NotarySigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendAttestation.WaitForConfirmation(a.GetTestContext(), txAddNotary)
	txAddGuard, err := a.AttestationContract.AddGuard(attestOwnerAuth.TransactOpts, a.GuardSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendAttestation.WaitForConfirmation(a.GetTestContext(), txAddGuard)
}

// SetupTest sets up the test.
func (a *SimulatedBackendsSuite) SetupTest() {
	a.TestSuite.SetupTest()

	a.TestBackendOrigin = preset.GetRinkeby().Geth(a.GetTestContext(), a.T())
	a.TestBackendDestination = preset.GetBSCTestnet().Geth(a.GetTestContext(), a.T())
	a.TestBackendAttestation = preset.GetMaticMumbai().Geth(a.GetTestContext(), a.T())

	var err error
	a.NotaryWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}
	a.GuardWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}

	a.NotarySigner = localsigner.NewSigner(a.NotaryWallet.PrivateKey())

	a.TestBackendAttestation.FundAccount(a.GetTestContext(), a.NotarySigner.Address(), *big.NewInt(params.Ether))
	a.GuardSigner = localsigner.NewSigner(a.GuardWallet.PrivateKey())

	deployManager := testutil.NewDeployManager(a.T())
	a.SetupOrigin(deployManager)
	a.SetupDestination(deployManager)
	a.SetupAttestation(deployManager)
}
