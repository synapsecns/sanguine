package executor_test

import (
	"context"
	"errors"
	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"github.com/synapsecns/sanguine/sin-executor/config"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchainclient"
	"github.com/synapsecns/sanguine/sin-executor/executor"
	"github.com/synapsecns/sanguine/sin-executor/testutil"
	"math/big"
	"sync"
	"testing"
)

// InterchainSuite is a test suite for the interchain package.
type InterchainSuite struct {
	*testsuite.TestSuite
	metrics       metrics.Handler
	originChain   backends.SimulatedTestBackend
	destChain     backends.SimulatedTestBackend
	deployManager *testutil.DeployManager
	originModule  *interchainclient.InterchainClientRef
	destModule    *interchainclient.InterchainClientRef
	omnirpcURL    string
	executor      *executor.Executor
}

func NewInterchainSuite(tb testing.TB) *InterchainSuite {
	tb.Helper()
	return &InterchainSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		metrics:   metrics.NewNullHandler(),
	}
}

func (i *InterchainSuite) SetupTest() {
	i.TestSuite.SetupTest()
	i.deployManager = testutil.NewDeployManager(i.T())

	var originInfo, destInfo contracts.DeployedContract
	// create origin and destination chains
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		i.originChain = geth.NewEmbeddedBackendForChainID(i.GetTestContext(), i.T(), big.NewInt(1))
		originInfo, i.originModule = i.deployManager.GetInterchainClient(i.GetTestContext(), i.originChain)
	}()
	go func() {
		defer wg.Done()
		i.destChain = geth.NewEmbeddedBackendForChainID(i.GetTestContext(), i.T(), big.NewInt(2))
		destInfo, i.destModule = i.deployManager.GetInterchainClient(i.GetTestContext(), i.destChain)
	}()
	wg.Wait()

	i.omnirpcURL = testhelper.NewOmnirpcServer(i.GetTestContext(), i.T(), i.originChain, i.destChain)

	i.setClientConfigs(i.originChain, i.originModule, originInfo, destInfo, i.destChain)
	i.setClientConfigs(i.destChain, i.destModule, destInfo, originInfo, i.originChain)
	i.makeExecutor()
}

func (i *InterchainSuite) setClientConfigs(backend backends.SimulatedTestBackend, contract *interchainclient.InterchainClientRef, myContract, otherContract contracts.DeployedContract, otherBackend backends.SimulatedTestBackend) {
	auth := backend.GetTxContext(i.GetTestContext(), myContract.OwnerPtr())

	tx, err := contract.SetLinkedClient(auth.TransactOpts, otherContract.ChainID(), i.addressToBytes32(otherContract.Address()))
	i.Require().NoError(err)
	backend.WaitForConfirmation(i.GetTestContext(), tx)

	// set the receiving module on the app
	_, appMock := i.deployManager.GetInterchainAppMock(i.GetTestContext(), backend)

	chainIDS := []uint64{backend.GetBigChainID().Uint64(), otherBackend.GetBigChainID().Uint64()}
	linkedApps := []common.Address{i.deployManager.Get(i.GetTestContext(), backend, testutil.InterchainApp).Address(), i.deployManager.Get(i.GetTestContext(), otherBackend, testutil.InterchainApp).Address()}

	sendingModules, err := appMock.GetSendingModules0(&bind.CallOpts{Context: i.GetTestContext()})
	i.Require().NoError(err)

	receivingModules, err := appMock.GetReceivingModules(&bind.CallOpts{Context: i.GetTestContext()})
	i.Require().NoError(err)
	// same thing

	tx, err = appMock.SetAppConfig(auth.TransactOpts, chainIDS, linkedApps, sendingModules, receivingModules, big.NewInt(1), 0)
	i.Require().NoError(err)
	backend.WaitForConfirmation(i.GetTestContext(), tx)
}

func (i *InterchainSuite) addressToBytes32(addie common.Address) [32]byte {
	_, lib := i.deployManager.GetOptionsLib(i.GetTestContext(), i.originChain)
	res, err := lib.AddressToBytes32(&bind.CallOpts{Context: i.GetTestContext()}, addie)
	i.Require().NoError(err)

	return res
}

func (i *InterchainSuite) makeExecutor() {
	testWallet, err := wallet.FromRandom()
	i.NoError(err)

	i.originChain.FundAccount(i.GetTestContext(), testWallet.Address(), *new(big.Int).SetUint64(params.Ether))
	i.destChain.FundAccount(i.GetTestContext(), testWallet.Address(), *new(big.Int).SetUint64(params.Ether))

	cfg := config.Config{
		Chains: map[int]string{
			1: i.originModule.Address().String(),
			2: i.destModule.Address().String(),
		},
		OmnirpcURL: i.omnirpcURL,
		Database: config.DatabaseConfig{
			Type: dbcommon.Sqlite.String(),
			DSN:  filet.TmpDir(i.T(), ""),
		},
		Signer: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(i.T(), "", testWallet.PrivateKeyHex()).Name(),
		},
		SubmitterConfig: submitterConfig.Config{},
	}
	i.executor, err = executor.NewExecutor(i.GetTestContext(), i.metrics, cfg)
	i.Require().NoError(err)

	go func() {
		err = i.executor.Start(i.GetTestContext())
		if !errors.Is(err, context.Canceled) {
			i.Require().NoError(err)
		}
	}()
}

func TestNodeSuite(t *testing.T) {
	suite.Run(t, NewInterchainSuite(t))
}
