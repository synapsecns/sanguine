package stiprelayer_test

import (
	"fmt"
	"testing"

	"math/big"

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
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	omnirpcHelper "github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/mockerc20"
	"github.com/synapsecns/sanguine/services/stiprelayer"
	"github.com/synapsecns/sanguine/services/stiprelayer/db"
	"github.com/synapsecns/sanguine/services/stiprelayer/db/sql"
	"github.com/synapsecns/sanguine/services/stiprelayer/stipconfig"
)

type STIPRelayerSuite struct {
	*testsuite.TestSuite
	omniRPCClient            omniClient.RPCClient
	omniRPCTestBackends      []backends.SimulatedTestBackend
	arbitrumSimulatedBackend backends.SimulatedTestBackend
	database                 db.STIPDB
	cfg                      stipconfig.Config
	testWallet               wallet.Wallet
	handler                  metrics.Handler
	arbERC20Address          common.Address
	stipRelayer              *stiprelayer.STIPRelayer
}

func NewSTIPRelayerSuite(tb testing.TB) *STIPRelayerSuite {
	tb.Helper()
	return &STIPRelayerSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (c *STIPRelayerSuite) SetupTest() {
	c.TestSuite.SetupTest()

	testOmnirpc := omnirpcHelper.NewOmnirpcServer(c.GetTestContext(), c.T(), c.omniRPCTestBackends...)
	omniRPCClient := omniClient.NewOmnirpcClient(testOmnirpc, c.handler, omniClient.WithCaptureReqRes())
	c.omniRPCClient = omniRPCClient

	stipRelayerInstance, err := stiprelayer.NewSTIPRelayer(c.GetTestContext(), c.cfg, c.handler, c.omniRPCClient, c.database)
	c.Require().NoError(err)
	c.stipRelayer = stipRelayerInstance
}

func (c *STIPRelayerSuite) SetupSuite() {
	c.TestSuite.SetupSuite()

	arbChainID := uint64(42161)
	backend := geth.NewEmbeddedBackendForChainID(c.GetSuiteContext(), c.T(), new(big.Int).SetUint64(arbChainID))

	testWallet, err := wallet.FromRandom()
	c.testWallet = testWallet
	c.Require().NoError(err)
	// Fund with relayer gas to deploy arb address for testing
	backend.FundAccount(c.GetSuiteContext(), c.testWallet.Address(), *big.NewInt(params.Ether))

	c.arbitrumSimulatedBackend = backend
	c.omniRPCTestBackends = append(c.omniRPCTestBackends, c.arbitrumSimulatedBackend)
	fmt.Println(c.arbitrumSimulatedBackend.RPCAddress())
	fmt.Println(c.omniRPCTestBackends[0].RPCAddress())
	c.handler = metrics.NewNullHandler()

	// Create an auth to interact with the blockchain
	arbChainIDBigInt := big.NewInt(int64(arbChainID))
	auth, err := bind.NewKeyedTransactorWithChainID(c.testWallet.PrivateKey(), arbChainIDBigInt)
	c.Require().NoError(err)

	mockErc20Address, tx, _, err := mockerc20.DeployMockERC20(auth, backend, "Arbitrum", 18)
	c.Require().NoError(err)
	backend.WaitForConfirmation(c.GetSuiteContext(), tx)

	c.arbERC20Address = mockErc20Address

	arbERC20Instance, err := mockerc20.NewMockERC20(c.arbERC20Address, backend)
	c.Require().NoError(err)
	//  Mint 1e18
	ether := big.NewInt(params.Ether)
	tx, err = arbERC20Instance.Mint(auth, c.testWallet.Address(), ether.Mul(ether, big.NewInt(10000)))
	c.Require().NoError(err)
	backend.WaitForConfirmation(c.GetSuiteContext(), tx)

	signerConfig := signerConfig.SignerConfig{
		Type: signerConfig.FileType.String(),
		File: filet.TmpFile(c.T(), "", c.testWallet.PrivateKeyHex()).Name(),
	}

	dbType, err := dbcommon.DBTypeFromString("sqlite")
	c.Require().NoError(err)

	// TODO use temp file / in memory sqlite3 to not create in directory files
	testDB, _ := sql.Connect(c.GetSuiteContext(), dbType, filet.TmpDir(c.T(), ""), c.handler)
	c.database = testDB

	c.cfg = stipconfig.Config{
		Signer:          signerConfig,
		SubmitterConfig: submitterConfig.Config{},
		ArbAddress:      c.arbERC20Address.Hex(),
		ArbChainID:      arbChainID,
	}
}

// TestConfigSuite runs the integration test suite.
func TestSTIPRelayerSuite(t *testing.T) {
	suite.Run(t, NewSTIPRelayerSuite(t))
}
