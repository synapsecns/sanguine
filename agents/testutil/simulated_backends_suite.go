package testutil

import (
	"math/big"
	"testing"

	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/scribe/backfill"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/agents/config"
	"github.com/synapsecns/sanguine/agents/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/agents/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/destinationharness"
	"github.com/synapsecns/sanguine/agents/contracts/test/originharness"
	"github.com/synapsecns/sanguine/agents/domains"
	"github.com/synapsecns/sanguine/agents/domains/evm"
	"github.com/synapsecns/sanguine/agents/testutil/agentstestcontract"
	"github.com/synapsecns/sanguine/agents/types"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/backends/preset"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
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
	OriginContract                               *originharness.OriginHarnessRef
	OriginContractMetadata                       contracts.DeployedContract
	OriginOwnerPtr                               common.Address
	OriginOwnerAuth                              backends.AuthType
	DestinationContract                          *destinationharness.DestinationHarnessRef
	DestinationContractMetadata                  contracts.DeployedContract
	DestinationOwnerPtr                          common.Address
	DestinationOwnerAuth                         backends.AuthType
	AttestationHarness                           *attestationharness.AttestationHarnessRef
	AttestationContract                          *attestationcollector.AttestationCollectorRef
	AttestationOwnerPtr                          common.Address
	AttestationOwnerAuth                         backends.AuthType
	AttestationContractMetadata                  contracts.DeployedContract
	AgentsTestContractOnOriginChain              *agentstestcontract.AgentsTestContractRef
	AgentsTestContractMetadataOnOriginChain      contracts.DeployedContract
	AgentsTestContractOnDestinationChain         *agentstestcontract.AgentsTestContractRef
	AgentsTestContractMetadataOnDestinationChain contracts.DeployedContract
	TestBackendOrigin                            backends.SimulatedTestBackend
	OriginSimulatedChain                         *geth.Backend
	OriginSimulatedClient                        backfill.ScribeBackend
	TestBackendDestination                       backends.SimulatedTestBackend
	DestinationSimulatedChain                    *geth.Backend
	DestinationSimulatedClient                   backfill.ScribeBackend
	TestBackendAttestation                       backends.SimulatedTestBackend
	AttestationSimulatedChain                    *geth.Backend
	AttestationSimulatedClient                   backfill.ScribeBackend
	NotaryWallet                                 wallet.Wallet
	GuardWallet                                  wallet.Wallet
	NotarySigner                                 signer.Signer
	GuardSigner                                  signer.Signer
	OriginWallet                                 wallet.Wallet
	DestinationWallet                            wallet.Wallet
	AttestationWallet                            wallet.Wallet
	OriginSigner                                 signer.Signer
	DestinationSigner                            signer.Signer
	AttestationSigner                            signer.Signer
	OriginDomainClient                           domains.DomainClient
	AttestationDomainClient                      domains.DomainClient
	DestinationDomainClient                      domains.DomainClient
	TestDeployManager                            *DeployManager
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
	a.TestBackendOrigin = preset.GetRinkeby().Geth(a.GetTestContext(), a.T())
	a.OriginContractMetadata, a.OriginContract = deployManager.GetOriginHarness(a.GetTestContext(), a.TestBackendOrigin)
	var err error
	a.OriginOwnerPtr, err = a.OriginContract.OriginHarnessCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}
	a.OriginOwnerAuth = a.TestBackendOrigin.GetTxContext(a.GetTestContext(), &a.OriginOwnerPtr)

	txOriginNotaryAdd, err := a.OriginContract.AddAgent(a.OriginOwnerAuth.TransactOpts, uint32(a.TestBackendDestination.GetChainID()), a.NotarySigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendOrigin.WaitForConfirmation(a.GetTestContext(), txOriginNotaryAdd)
	txOriginGuardAdd, err := a.OriginContract.AddAgent(a.OriginOwnerAuth.TransactOpts, uint32(0), a.GuardSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendOrigin.WaitForConfirmation(a.GetTestContext(), txOriginGuardAdd)

	a.OriginDomainClient, err = evm.NewEVM(a.GetTestContext(), "origin_client", config.DomainConfig{
		DomainID:      uint32(a.TestBackendOrigin.GetBigChainID().Uint64()),
		Type:          types.EVM.String(),
		OriginAddress: a.OriginContract.Address().String(),
		RPCUrl:        a.TestBackendOrigin.RPCAddress(),
	})
	if err != nil {
		a.T().Fatal(err)
	}

	a.OriginWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}

	a.OriginSigner = localsigner.NewSigner(a.OriginWallet.PrivateKey())

	a.TestBackendOrigin.FundAccount(a.GetTestContext(), a.OriginSigner.Address(), *big.NewInt(params.Ether))

	a.OriginSimulatedChain = geth.NewEmbeddedBackendForChainID(a.GetTestContext(), a.T(), big.NewInt(a.TestBackendOrigin.GetBigChainID().Int64()))
	a.OriginSimulatedClient, err = backfill.DialBackend(a.GetTestContext(), a.OriginSimulatedChain.RPCAddress())
	if err != nil {
		a.T().Fatal(err)
	}
	a.OriginSimulatedChain.FundAccount(a.GetTestContext(), a.OriginWallet.Address(), *big.NewInt(params.Ether))

	a.AgentsTestContractMetadataOnOriginChain, a.AgentsTestContractOnOriginChain = deployManager.GetAgentsTestContract(a.GetTestContext(), a.OriginSimulatedChain)

	a.DestinationOwnerAuth = a.TestBackendDestination.GetTxContext(a.GetTestContext(), &a.DestinationOwnerPtr)
}

// SetupDestination sets up the backend that will have the destination contract deployed on it.
//
//nolint:dupl
func (a *SimulatedBackendsTestSuite) SetupDestination(deployManager *DeployManager) {
	a.TestBackendDestination = preset.GetBSCTestnet().Geth(a.GetTestContext(), a.T())
	a.DestinationContractMetadata, a.DestinationContract = deployManager.GetDestinationHarness(a.GetTestContext(), a.TestBackendDestination)

	var err error
	a.DestinationOwnerPtr, err = a.DestinationContract.DestinationHarnessCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}

	a.DestinationOwnerAuth = a.TestBackendDestination.GetTxContext(a.GetTestContext(), &a.DestinationOwnerPtr)
	txDestinationNotaryAdd, err := a.DestinationContract.AddAgent(a.DestinationOwnerAuth.TransactOpts, uint32(a.TestBackendDestination.GetChainID()), a.NotarySigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendDestination.WaitForConfirmation(a.GetTestContext(), txDestinationNotaryAdd)
	txDestinationGuardAdd, err := a.DestinationContract.AddAgent(a.DestinationOwnerAuth.TransactOpts, uint32(0), a.GuardSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendDestination.WaitForConfirmation(a.GetTestContext(), txDestinationGuardAdd)

	a.DestinationDomainClient, err = evm.NewEVM(a.GetTestContext(), "destination_client", config.DomainConfig{
		DomainID:           uint32(a.TestBackendDestination.GetBigChainID().Uint64()),
		Type:               types.EVM.String(),
		DestinationAddress: a.DestinationContract.Address().String(),
		RPCUrl:             a.TestBackendDestination.RPCAddress(),
	})
	if err != nil {
		a.T().Fatal(err)
	}

	a.DestinationWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}

	a.DestinationSigner = localsigner.NewSigner(a.DestinationWallet.PrivateKey())

	a.TestBackendDestination.FundAccount(a.GetTestContext(), a.DestinationSigner.Address(), *big.NewInt(params.Ether))

	a.DestinationSimulatedChain = geth.NewEmbeddedBackendForChainID(a.GetTestContext(), a.T(), big.NewInt(a.TestBackendDestination.GetBigChainID().Int64()))
	a.DestinationSimulatedClient, err = backfill.DialBackend(a.GetTestContext(), a.DestinationSimulatedChain.RPCAddress())
	if err != nil {
		a.T().Fatal(err)
	}
	a.DestinationSimulatedChain.FundAccount(a.GetTestContext(), a.DestinationWallet.Address(), *big.NewInt(params.Ether))

	a.AgentsTestContractMetadataOnDestinationChain, a.AgentsTestContractOnDestinationChain = deployManager.GetAgentsTestContract(a.GetTestContext(), a.DestinationSimulatedChain)
}

// SetupAttestation sets up the backend that will have the attestation collector contract deployed on it.
func (a *SimulatedBackendsTestSuite) SetupAttestation(deployManager *DeployManager) {
	a.TestBackendAttestation = preset.GetMaticMumbai().Geth(a.GetTestContext(), a.T())
	_, a.AttestationHarness = deployManager.GetAttestationHarness(a.GetTestContext(), a.TestBackendAttestation)
	a.AttestationContractMetadata, a.AttestationContract = deployManager.GetAttestationCollector(a.GetTestContext(), a.TestBackendAttestation)

	var err error
	a.AttestationOwnerPtr, err = a.AttestationContract.AttestationCollectorCaller.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}
	a.AttestationOwnerAuth = a.TestBackendAttestation.GetTxContext(a.GetTestContext(), &a.AttestationOwnerPtr)

	txAddNotary, err := a.AttestationContract.AddAgent(a.AttestationOwnerAuth.TransactOpts, uint32(a.TestBackendDestination.GetChainID()), a.NotarySigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendAttestation.WaitForConfirmation(a.GetTestContext(), txAddNotary)
	txAddGuard, err := a.AttestationContract.AddAgent(a.AttestationOwnerAuth.TransactOpts, uint32(0), a.GuardSigner.Address())
	if err != nil {
		a.T().Fatal(err)
	}
	a.TestBackendAttestation.WaitForConfirmation(a.GetTestContext(), txAddGuard)

	a.AttestationDomainClient, err = evm.NewEVM(a.GetTestContext(), "attestation_client", config.DomainConfig{
		DomainID:                    uint32(a.TestBackendAttestation.GetBigChainID().Uint64()),
		Type:                        types.EVM.String(),
		AttestationCollectorAddress: a.AttestationContract.Address().String(),
		RPCUrl:                      a.TestBackendAttestation.RPCAddress(),
	})
	if err != nil {
		a.T().Fatal(err)
	}

	a.AttestationWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}
	a.AttestationSigner = localsigner.NewSigner(a.AttestationWallet.PrivateKey())

	a.TestBackendAttestation.FundAccount(a.GetTestContext(), a.AttestationSigner.Address(), *big.NewInt(params.Ether))

	a.AttestationSimulatedChain = geth.NewEmbeddedBackendForChainID(a.GetTestContext(), a.T(), big.NewInt(a.TestBackendAttestation.GetBigChainID().Int64()))
	a.AttestationSimulatedClient, err = backfill.DialBackend(a.GetTestContext(), a.AttestationSimulatedChain.RPCAddress())
	if err != nil {
		a.T().Fatal(err)
	}

	a.AttestationSimulatedChain.FundAccount(a.GetTestContext(), a.AttestationWallet.Address(), *big.NewInt(params.Ether))
}

// SetupGuard sets up the Guard agent.
func (a *SimulatedBackendsTestSuite) SetupGuard() {
	var err error
	a.GuardWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}
	a.GuardSigner = localsigner.NewSigner(a.GuardWallet.PrivateKey())
}

// SetupNotary sets up the Notary agent.
func (a *SimulatedBackendsTestSuite) SetupNotary() {
	var err error
	a.NotaryWallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}
	a.NotarySigner = localsigner.NewSigner(a.NotaryWallet.PrivateKey())
}

// SetupTest sets up the test.
func (a *SimulatedBackendsTestSuite) SetupTest() {
	a.TestSuite.SetupTest()

	a.SetupGuard()
	a.SetupNotary()

	a.TestDeployManager = NewDeployManager(a.T())
	a.SetupDestination(a.TestDeployManager)
	a.SetupOrigin(a.TestDeployManager)
	a.SetupAttestation(a.TestDeployManager)
}
