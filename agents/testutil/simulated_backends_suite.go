package testutil

import (
	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	executorsqllite "github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/sqlite"
	executorMetadata "github.com/synapsecns/sanguine/agents/agents/executor/metadata"
	guardMetadata "github.com/synapsecns/sanguine/agents/agents/guard/metadata"
	notaryMetadata "github.com/synapsecns/sanguine/agents/agents/notary/metadata"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/contracts/inbox"
	"github.com/synapsecns/sanguine/agents/contracts/lightinbox"
	"github.com/synapsecns/sanguine/agents/contracts/test/bondingmanagerharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/destinationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/lightmanagerharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/pingpongclient"
	"github.com/synapsecns/sanguine/agents/contracts/test/summitharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/testclient"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/testutil/agentstestcontract"
	"github.com/synapsecns/sanguine/agents/types"
	coreConfig "github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/preset"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	scribedb "github.com/synapsecns/sanguine/services/scribe/db"
	scribesqlite "github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/sqlite"
	scribeMetadata "github.com/synapsecns/sanguine/services/scribe/metadata"
	"math/big"
	"sync"
	"testing"
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
	LightInboxOnOrigin                  *lightinbox.LightInboxRef
	LightInboxMetadataOnOrigin          contracts.DeployedContract
	LightManagerOnOrigin                *lightmanagerharness.LightManagerHarnessRef
	LightManagerMetadataOnOrigin        contracts.DeployedContract
	OriginContract                      *originharness.OriginHarnessRef
	OriginContractMetadata              contracts.DeployedContract
	DestinationContractOnOrigin         *destinationharness.DestinationHarnessRef
	DestinationContractMetadataOnOrigin contracts.DeployedContract
	TestContractOnOrigin                *agentstestcontract.AgentsTestContractRef
	TestContractMetadataOnOrigin        contracts.DeployedContract
	TestContractOnSummit                *agentstestcontract.AgentsTestContractRef
	TestContractMetadataOnSummit        contracts.DeployedContract
	TestContractOnDestination           *agentstestcontract.AgentsTestContractRef
	TestContractMetadataOnDestination   contracts.DeployedContract
	TestClientOnOrigin                  *testclient.TestClientRef
	TestClientMetadataOnOrigin          contracts.DeployedContract
	PingPongClientOnOrigin              *pingpongclient.PingPongClientRef
	PingPongClientMetadataOnOrigin      contracts.DeployedContract
	DestinationContract                 *destinationharness.DestinationHarnessRef
	DestinationContractMetadata         contracts.DeployedContract
	LightInboxOnDestination             *lightinbox.LightInboxRef
	LightInboxMetadataOnDestination     contracts.DeployedContract
	LightManagerOnDestination           *lightmanagerharness.LightManagerHarnessRef
	LightManagerMetadataOnDestination   contracts.DeployedContract
	OriginContractOnDestination         *originharness.OriginHarnessRef
	OriginContractMetadataOnDestination contracts.DeployedContract
	TestClientOnDestination             *testclient.TestClientRef
	TestClientMetadataOnDestination     contracts.DeployedContract
	PingPongClientOnDestination         *pingpongclient.PingPongClientRef
	PingPongClientMetadataOnDestination contracts.DeployedContract
	InboxOnSummit                       *inbox.InboxRef
	InboxMetadataOnSummit               contracts.DeployedContract
	BondingManagerOnSummit              *bondingmanagerharness.BondingManagerHarnessRef
	BondingManagerMetadataOnSummit      contracts.DeployedContract
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
	ScribeMetrics                       metrics.Handler
	ExecutorMetrics                     metrics.Handler
	NotaryMetrics                       metrics.Handler
	GuardMetrics                        metrics.Handler
	ContractMetrics                     metrics.Handler
}

// NewSimulatedBackendsTestSuite creates an end-to-end test suite with simulated
// backends set up.
func NewSimulatedBackendsTestSuite(tb testing.TB) *SimulatedBackendsTestSuite {
	tb.Helper()
	return &SimulatedBackendsTestSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

// SetupSuite sets up the test suite.
func (a *SimulatedBackendsTestSuite) SetupSuite() {
	a.TestSuite.SetupSuite()
	a.TestSuite.LogDir = filet.TmpDir(a.T(), "")
	localmetrics.SetupTestJaeger(a.GetSuiteContext(), a.T())

	var err error
	a.ScribeMetrics, err = metrics.NewByType(a.GetSuiteContext(), scribeMetadata.BuildInfo(), metrics.Jaeger)
	a.Require().Nil(err)
	a.ExecutorMetrics, err = metrics.NewByType(a.GetSuiteContext(), executorMetadata.BuildInfo(), metrics.Jaeger)
	a.Require().Nil(err)
	a.NotaryMetrics, err = metrics.NewByType(a.GetSuiteContext(), notaryMetadata.BuildInfo(), metrics.Jaeger)
	a.Require().Nil(err)
	a.GuardMetrics, err = metrics.NewByType(a.GetSuiteContext(), guardMetadata.BuildInfo(), metrics.Jaeger)
	a.Require().Nil(err)
	a.ContractMetrics, err = metrics.NewByType(a.GetSuiteContext(), coreConfig.NewBuildInfo(
		coreConfig.DefaultVersion,
		coreConfig.DefaultCommit,
		"contract",
		coreConfig.DefaultDate,
	), metrics.Jaeger)
	a.Require().Nil(err)
}

// SetupOrigin sets up the backend that will have the origin contract deployed on it.
//
//nolint:dupl
func (a *SimulatedBackendsTestSuite) SetupOrigin(deployManager *DeployManager) {
	a.OriginContractMetadata, a.OriginContract = deployManager.GetOriginHarness(a.GetTestContext(), a.TestBackendOrigin)
	a.DestinationContractMetadataOnOrigin, a.DestinationContractOnOrigin = deployManager.GetDestinationHarness(a.GetTestContext(), a.TestBackendOrigin)
	a.TestClientMetadataOnOrigin, a.TestClientOnOrigin = deployManager.GetTestClient(a.GetTestContext(), a.TestBackendOrigin)
	a.PingPongClientMetadataOnOrigin, a.PingPongClientOnOrigin = deployManager.GetPingPongClient(a.GetTestContext(), a.TestBackendOrigin)
	a.LightInboxMetadataOnOrigin, a.LightInboxOnOrigin = deployManager.GetLightInbox(a.GetTestContext(), a.TestBackendOrigin)
	a.LightManagerMetadataOnOrigin, a.LightManagerOnOrigin = deployManager.GetLightManagerHarness(a.GetTestContext(), a.TestBackendOrigin)
	a.TestContractMetadataOnOrigin, a.TestContractOnOrigin = deployManager.GetAgentsTestContract(a.GetTestContext(), a.TestBackendOrigin)

	var err error
	a.OriginDomainClient, err = evm.NewEVM(a.GetTestContext(), "origin_client", config.DomainConfig{
		DomainID:            uint32(a.TestBackendOrigin.GetBigChainID().Uint64()),
		Type:                types.EVM.String(),
		OriginAddress:       a.OriginContract.Address().String(),
		DestinationAddress:  a.DestinationContractOnOrigin.Address().String(),
		LightManagerAddress: a.LightManagerOnOrigin.Address().String(),
		LightInboxAddress:   a.LightInboxOnOrigin.Address().String(),
	}, a.TestBackendOrigin.RPCAddress())
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
	a.LightInboxMetadataOnDestination, a.LightInboxOnDestination = deployManager.GetLightInbox(a.GetTestContext(), a.TestBackendDestination)
	a.LightManagerMetadataOnDestination, a.LightManagerOnDestination = deployManager.GetLightManagerHarness(a.GetTestContext(), a.TestBackendDestination)
	a.TestContractMetadataOnDestination, a.TestContractOnDestination = deployManager.GetAgentsTestContract(a.GetTestContext(), a.TestBackendDestination)

	var err error
	/*agentStatus, err := a.DestinationContract.AgentStatus(&bind.CallOpts{Context: a.GetTestContext()}, a.NotaryBondedSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}

	if agentStatus.Domain != uint32(a.TestBackendDestination.GetChainID()) {
		a.T().Fatal(err)
	}*/

	a.DestinationDomainClient, err = evm.NewEVM(a.GetTestContext(), "destination_client", config.DomainConfig{
		DomainID:            uint32(a.TestBackendDestination.GetBigChainID().Uint64()),
		Type:                types.EVM.String(),
		OriginAddress:       a.OriginContractOnDestination.Address().String(),
		DestinationAddress:  a.DestinationContract.Address().String(),
		LightManagerAddress: a.LightManagerOnDestination.Address().String(),
		LightInboxAddress:   a.LightInboxOnDestination.Address().String(),
	}, a.TestBackendDestination.RPCAddress())
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
	a.InboxMetadataOnSummit, a.InboxOnSummit = deployManager.GetInbox(a.GetTestContext(), a.TestBackendSummit)
	a.BondingManagerMetadataOnSummit, a.BondingManagerOnSummit = deployManager.GetBondingManagerHarness(a.GetTestContext(), a.TestBackendSummit)
	a.SummitMetadata, a.SummitContract = deployManager.GetSummitHarness(a.GetTestContext(), a.TestBackendSummit)
	a.TestContractMetadataOnSummit, a.TestContractOnSummit = deployManager.GetAgentsTestContract(a.GetTestContext(), a.TestBackendSummit)

	var err error
	a.SummitDomainClient, err = evm.NewEVM(a.GetTestContext(), "summit_client", config.DomainConfig{
		DomainID:              uint32(a.TestBackendSummit.GetBigChainID().Uint64()),
		Type:                  types.EVM.String(),
		SummitAddress:         a.SummitContract.Address().String(),
		BondingManagerAddress: a.BondingManagerOnSummit.Address().String(),
		InboxAddress:          a.InboxOnSummit.Address().String(),
	}, a.TestBackendSummit.RPCAddress())
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

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		a.TestBackendOrigin = preset.GetRinkeby().Geth(a.GetTestContext(), a.T())
	}()
	go func() {
		defer wg.Done()
		a.TestBackendDestination = preset.GetBSCTestnet().Geth(a.GetTestContext(), a.T())
	}()
	go func() {
		defer wg.Done()
		a.TestBackendSummit = preset.GetMaticMumbaiFakeSynDomain().Geth(a.GetTestContext(), a.T())
	}()
	wg.Wait()

	a.SetupSummit(a.TestDeployManager)
	a.SetupDestination(a.TestDeployManager)
	a.SetupOrigin(a.TestDeployManager)

	err := a.TestDeployManager.LoadHarnessContractsOnChains(
		a.GetTestContext(),
		a.TestBackendSummit,
		[]backends.SimulatedTestBackend{a.TestBackendOrigin, a.TestBackendDestination},
		[]common.Address{a.GuardBondedSigner.Address(), a.NotaryBondedSigner.Address(), a.NotaryOnOriginBondedSigner.Address()},
		[]uint32{uint32(0), uint32(a.TestBackendDestination.GetChainID()), uint32(a.TestBackendOrigin.GetChainID())})
	if err != nil {
		a.T().Fatal(err)
	}

	a.DBPath = filet.TmpDir(a.T(), "")
	scribeSqliteStore, err := scribesqlite.NewSqliteStore(a.GetTestContext(), a.DBPath, a.ScribeMetrics, false)
	if err != nil {
		a.T().Fatal(err)
	}
	a.ScribeTestDB = scribeSqliteStore
	sqliteStore, err := executorsqllite.NewSqliteStore(a.GetTestContext(), a.DBPath, a.ExecutorMetrics, false)
	if err != nil {
		a.T().Fatal(err)
	}
	a.ExecutorTestDB = sqliteStore
}

// cleanAfterTestSuite does cleanup after test suite is finished.
func (a *SimulatedBackendsTestSuite) cleanAfterTestSuite() {
	filet.CleanUp(a.T())
}
