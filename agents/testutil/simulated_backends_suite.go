package testutil

import (
	"math/big"
	"testing"

	"github.com/Flaque/filet"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	executorsqllite "github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/contracts/test/destinationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/pingpongclient"
	"github.com/synapsecns/sanguine/agents/contracts/test/summitharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/testclient"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/preset"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	scribesqlite "github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
)

// SimulatedBackendsTestSuite can be used as the base for any test needing simulated backends
// that have an origin, destination and attestation collector and a guard and notary
// added to each.
// TODO (joe): For tests that do not need all 3 simulated backends, allow them to pass in
// flags indicating the subset of backends desired. Some tests might only want
// an attestation collector, others might only want an origin and an attestation collector,
// others might want just a destination, etc.
type SimulatedBackendsTestSuite struct {
	*testsuite.TestSuite
	OriginContract                      *originharness.OriginHarnessRef
	OriginContractMetadata              contracts.DeployedContract
	DestinationContractOnOrigin         *destinationharness.DestinationHarnessRef
	DestinationContractMetadataOnOrigin contracts.DeployedContract
	TestClientOnOrigin                  *testclient.TestClientRef
	TestClientMetadataOnOrigin          contracts.DeployedContract
	PingPongClientOnOrigin              *pingpongclient.PingPongClientRef
	PingPongClientMetadataOnOrigin      contracts.DeployedContract
	DestinationContract                 *destinationharness.DestinationHarnessRef
	DestinationContractMetadata         contracts.DeployedContract
	OriginContractOnDestination         *originharness.OriginHarnessRef
	OriginContractMetadataOnDestination contracts.DeployedContract
	TestClientOnDestination             *testclient.TestClientRef
	TestClientMetadataOnDestination     contracts.DeployedContract
	PingPongClientOnDestination         *pingpongclient.PingPongClientRef
	PingPongClientMetadataOnDestination contracts.DeployedContract
	SummitContract                      *summitharness.SummitHarnessRef
	SummitMetadata                      contracts.DeployedContract
	TestBackendOrigin                   backends.SimulatedTestBackend
	TestBackendDestination              backends.SimulatedTestBackend
	TestBackendSummit                   backends.SimulatedTestBackend
	NotaryBondedWallet                  wallet.Wallet
	NotaryOnOriginBondedWallet          wallet.Wallet
	GuardBondedWallet                   wallet.Wallet
	NotaryBondedSigner                  signer.Signer
	NotaryOnOriginBondedSigner          signer.Signer
	GuardBondedSigner                   signer.Signer
	NotaryUnbondedWallet                wallet.Wallet
	NotaryUnbondedSigner                signer.Signer
	NotaryOnOriginUnbondedWallet        wallet.Wallet
	NotaryOnOriginUnbondedSigner        signer.Signer
	GuardUnbondedWallet                 wallet.Wallet
	GuardUnbondedSigner                 signer.Signer
	ExecutorUnbondedWallet              wallet.Wallet
	ExecutorUnbondedSigner              signer.Signer
	OriginDomainClient                  domains.DomainClient
	SummitDomainClient                  domains.DomainClient
	DestinationDomainClient             domains.DomainClient
	TestDeployManager                   *DeployManager
	ScribeTestDB                        scribedb.EventDB
	DBPath                              string
	ExecutorTestDB                      db.ExecutorDB
}

// NewSimulatedBackendsTestSuite creates an end-to-end test suite with simulated
// backends set up.
func NewSimulatedBackendsTestSuite(tb testing.TB) *SimulatedBackendsTestSuite {
	tb.Helper()
	return &SimulatedBackendsTestSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

// SetupOrigin sets up the backend that will have the origin contract deployed on it.
//
//nolint:dupl
func (a *SimulatedBackendsTestSuite) SetupOrigin(deployManager *DeployManager) {
	a.OriginContractMetadata, a.OriginContract = deployManager.GetOriginHarness(a.GetTestContext(), a.TestBackendOrigin)
	a.DestinationContractMetadataOnOrigin, a.DestinationContractOnOrigin = deployManager.GetDestinationHarness(a.GetTestContext(), a.TestBackendOrigin)
	a.TestClientMetadataOnOrigin, a.TestClientOnOrigin = deployManager.GetTestClient(a.GetTestContext(), a.TestBackendOrigin)
	a.PingPongClientMetadataOnOrigin, a.PingPongClientOnOrigin = deployManager.GetPingPongClient(a.GetTestContext(), a.TestBackendOrigin)

	originOwnerPtr, err := a.OriginContract.OriginHarnessCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}
	originOwnerAuth := a.TestBackendOrigin.GetTxContext(a.GetTestContext(), &originOwnerPtr)

	txOriginNotaryAdd, err := a.OriginContract.AddAgent(originOwnerAuth.TransactOpts, uint32(a.TestBackendDestination.GetChainID()), a.NotaryBondedSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendOrigin.WaitForConfirmation(a.GetTestContext(), txOriginNotaryAdd)
	txOriginGuardAdd, err := a.OriginContract.AddAgent(originOwnerAuth.TransactOpts, uint32(0), a.GuardBondedSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendOrigin.WaitForConfirmation(a.GetTestContext(), txOriginGuardAdd)

	destinationOwnerPtr, err := a.DestinationContractOnOrigin.DestinationHarnessCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}
	destinationOwnerAuth := a.TestBackendOrigin.GetTxContext(a.GetTestContext(), &destinationOwnerPtr)

	txDestinationNotaryAdd, err := a.DestinationContractOnOrigin.AddAgent(destinationOwnerAuth.TransactOpts, uint32(a.TestBackendOrigin.GetChainID()), a.NotaryOnOriginBondedSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendOrigin.WaitForConfirmation(a.GetTestContext(), txDestinationNotaryAdd)
	txDestinationGuardAdd, err := a.DestinationContractOnOrigin.AddAgent(destinationOwnerAuth.TransactOpts, uint32(0), a.GuardBondedSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendOrigin.WaitForConfirmation(a.GetTestContext(), txDestinationGuardAdd)

	a.OriginDomainClient, err = evm.NewEVM(a.GetTestContext(), "origin_client", config.DomainConfig{
		DomainID:           uint32(a.TestBackendOrigin.GetBigChainID().Uint64()),
		Type:               types.EVM.String(),
		OriginAddress:      a.OriginContract.Address().String(),
		DestinationAddress: a.DestinationContractOnOrigin.Address().String(),
		RPCUrl:             a.TestBackendOrigin.RPCAddress(),
	})
	if err != nil {
		a.T().Fatal(err)
	}

	a.TestBackendOrigin.FundAccount(a.GetTestContext(), a.NotaryUnbondedSigner.Address(), *big.NewInt(params.Ether))
	a.TestBackendOrigin.FundAccount(a.GetTestContext(), a.NotaryOnOriginUnbondedSigner.Address(), *big.NewInt(params.Ether))
	a.TestBackendOrigin.FundAccount(a.GetTestContext(), a.GuardUnbondedSigner.Address(), *big.NewInt(params.Ether))
	a.TestBackendOrigin.FundAccount(a.GetTestContext(), a.ExecutorUnbondedSigner.Address(), *big.NewInt(params.Ether))
}

// SetupDestination sets up the backend that will have the destination contract deployed on it.
//
//nolint:dupl
func (a *SimulatedBackendsTestSuite) SetupDestination(deployManager *DeployManager) {
	a.DestinationContractMetadata, a.DestinationContract = deployManager.GetDestinationHarness(a.GetTestContext(), a.TestBackendDestination)
	a.OriginContractMetadataOnDestination, a.OriginContractOnDestination = deployManager.GetOriginHarness(a.GetTestContext(), a.TestBackendDestination)
	a.TestClientMetadataOnDestination, a.TestClientOnDestination = deployManager.GetTestClient(a.GetTestContext(), a.TestBackendDestination)
	a.PingPongClientMetadataOnDestination, a.PingPongClientOnDestination = deployManager.GetPingPongClient(a.GetTestContext(), a.TestBackendDestination)

	destOwnerPtr, err := a.DestinationContract.DestinationHarnessCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}

	destOwnerAuth := a.TestBackendDestination.GetTxContext(a.GetTestContext(), &destOwnerPtr)
	txDestinationNotaryAdd, err := a.DestinationContract.AddAgent(destOwnerAuth.TransactOpts, uint32(a.TestBackendDestination.GetChainID()), a.NotaryBondedSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendDestination.WaitForConfirmation(a.GetTestContext(), txDestinationNotaryAdd)
	txDestinationGuardAdd, err := a.DestinationContract.AddAgent(destOwnerAuth.TransactOpts, uint32(0), a.GuardBondedSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendDestination.WaitForConfirmation(a.GetTestContext(), txDestinationGuardAdd)

	originOwnerPtr, err := a.OriginContractOnDestination.OriginHarnessCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}

	originOwnerAuth := a.TestBackendDestination.GetTxContext(a.GetTestContext(), &originOwnerPtr)
	txOriginNotaryAdd, err := a.OriginContractOnDestination.AddAgent(originOwnerAuth.TransactOpts, uint32(a.TestBackendOrigin.GetChainID()), a.NotaryOnOriginBondedSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendDestination.WaitForConfirmation(a.GetTestContext(), txOriginNotaryAdd)
	txOriginGuardAdd, err := a.OriginContractOnDestination.AddAgent(originOwnerAuth.TransactOpts, uint32(0), a.GuardBondedSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendDestination.WaitForConfirmation(a.GetTestContext(), txOriginGuardAdd)

	a.DestinationDomainClient, err = evm.NewEVM(a.GetTestContext(), "destination_client", config.DomainConfig{
		DomainID:           uint32(a.TestBackendDestination.GetBigChainID().Uint64()),
		Type:               types.EVM.String(),
		OriginAddress:      a.OriginContractOnDestination.Address().String(),
		DestinationAddress: a.DestinationContract.Address().String(),
		RPCUrl:             a.TestBackendDestination.RPCAddress(),
	})
	if err != nil {
		a.T().Fatal(err)
	}

	a.TestBackendDestination.FundAccount(a.GetTestContext(), a.NotaryUnbondedSigner.Address(), *big.NewInt(params.Ether))
	a.TestBackendDestination.FundAccount(a.GetTestContext(), a.NotaryOnOriginUnbondedSigner.Address(), *big.NewInt(params.Ether))
	a.TestBackendDestination.FundAccount(a.GetTestContext(), a.GuardUnbondedSigner.Address(), *big.NewInt(params.Ether))
	a.TestBackendDestination.FundAccount(a.GetTestContext(), a.ExecutorUnbondedSigner.Address(), *big.NewInt(params.Ether))
}

// SetupSummit sets up the backend that will have the summit contract deployed on it.
func (a *SimulatedBackendsTestSuite) SetupSummit(deployManager *DeployManager) {
	a.SummitMetadata, a.SummitContract = deployManager.GetSummitHarness(a.GetTestContext(), a.TestBackendSummit)

	summitOwnerPtr, err := a.SummitContract.SummitHarnessCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}
	summitOwnerAuth := a.TestBackendSummit.GetTxContext(a.GetTestContext(), &summitOwnerPtr)

	txAddNotary, err := a.SummitContract.AddAgent(summitOwnerAuth.TransactOpts, uint32(a.TestBackendDestination.GetChainID()), a.NotaryBondedSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendSummit.WaitForConfirmation(a.GetTestContext(), txAddNotary)
	txAddNotaryOnOrigin, err := a.SummitContract.AddAgent(summitOwnerAuth.TransactOpts, uint32(a.TestBackendOrigin.GetChainID()), a.NotaryOnOriginBondedSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendSummit.WaitForConfirmation(a.GetTestContext(), txAddNotaryOnOrigin)
	txAddGuard, err := a.SummitContract.AddAgent(summitOwnerAuth.TransactOpts, uint32(0), a.GuardBondedSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendSummit.WaitForConfirmation(a.GetTestContext(), txAddGuard)

	a.SummitDomainClient, err = evm.NewEVM(a.GetTestContext(), "summit_client", config.DomainConfig{
		DomainID:      uint32(a.TestBackendSummit.GetBigChainID().Uint64()),
		Type:          types.EVM.String(),
		SummitAddress: a.SummitContract.Address().String(),
		RPCUrl:        a.TestBackendSummit.RPCAddress(),
	})
	if err != nil {
		a.T().Fatal(err)
	}

	a.TestBackendSummit.FundAccount(a.GetTestContext(), a.NotaryUnbondedSigner.Address(), *big.NewInt(params.Ether))
	a.TestBackendSummit.FundAccount(a.GetTestContext(), a.NotaryOnOriginUnbondedSigner.Address(), *big.NewInt(params.Ether))
	a.TestBackendSummit.FundAccount(a.GetTestContext(), a.GuardUnbondedSigner.Address(), *big.NewInt(params.Ether))
	a.TestBackendSummit.FundAccount(a.GetTestContext(), a.ExecutorUnbondedSigner.Address(), *big.NewInt(params.Ether))
}

// SetupGuard sets up the Guard agent.
func (a *SimulatedBackendsTestSuite) SetupGuard() {
	var err error
	a.GuardBondedWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}
	a.GuardBondedSigner = localsigner.NewSigner(a.GuardBondedWallet.PrivateKey())

	a.GuardUnbondedWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}
	a.GuardUnbondedSigner = localsigner.NewSigner(a.GuardUnbondedWallet.PrivateKey())
}

// SetupNotary sets up the Notary agent.
func (a *SimulatedBackendsTestSuite) SetupNotary() {
	var err error
	a.NotaryBondedWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}
	a.NotaryBondedSigner = localsigner.NewSigner(a.NotaryBondedWallet.PrivateKey())

	a.NotaryUnbondedWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}
	a.NotaryUnbondedSigner = localsigner.NewSigner(a.NotaryUnbondedWallet.PrivateKey())
}

// SetupNotaryOnOrigin sets up the Notary agent on the origin chain.
func (a *SimulatedBackendsTestSuite) SetupNotaryOnOrigin() {
	var err error
	a.NotaryOnOriginBondedWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}
	a.NotaryOnOriginBondedSigner = localsigner.NewSigner(a.NotaryOnOriginBondedWallet.PrivateKey())

	a.NotaryOnOriginUnbondedWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}
	a.NotaryOnOriginUnbondedSigner = localsigner.NewSigner(a.NotaryOnOriginUnbondedWallet.PrivateKey())
}

// SetupExecutor sets up the Executor agent.
func (a *SimulatedBackendsTestSuite) SetupExecutor() {
	var err error
	a.ExecutorUnbondedWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}
	a.ExecutorUnbondedSigner = localsigner.NewSigner(a.ExecutorUnbondedWallet.PrivateKey())
}

// SetupTest sets up the test.
func (a *SimulatedBackendsTestSuite) SetupTest() {
	a.TestSuite.SetupTest()
	a.TestSuite.DeferAfterSuite(a.cleanAfterTestSuite)

	a.SetupGuard()
	a.SetupNotary()
	a.SetupNotaryOnOrigin()
	a.SetupExecutor()

	a.TestDeployManager = NewDeployManager(a.T())

	a.TestBackendOrigin = preset.GetRinkeby().Geth(a.GetTestContext(), a.T())
	a.TestBackendDestination = preset.GetBSCTestnet().Geth(a.GetTestContext(), a.T())
	a.TestBackendSummit = preset.GetMaticMumbaiFakeSynDomain().Geth(a.GetTestContext(), a.T())

	a.SetupDestination(a.TestDeployManager)
	a.SetupOrigin(a.TestDeployManager)
	a.SetupSummit(a.TestDeployManager)

	a.DBPath = filet.TmpDir(a.T(), "")
	scribeSqliteStore, err := scribesqlite.NewSqliteStore(a.GetTestContext(), a.DBPath)
	if err != nil {
		a.T().Fatal(err)
	}
	a.ScribeTestDB = scribeSqliteStore
	sqliteStore, err := executorsqllite.NewSqliteStore(a.GetTestContext(), a.DBPath)
	if err != nil {
		a.T().Fatal(err)
	}
	a.ExecutorTestDB = sqliteStore
}

// cleanAfterTestSuite does cleanup after test suite is finished.
func (a *SimulatedBackendsTestSuite) cleanAfterTestSuite() {
	filet.CleanUp(a.T())
}
