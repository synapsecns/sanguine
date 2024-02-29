// not how this is always scoped outside of the package itself
package relayer_test

import (
	"fmt"
	"math/big"
	"net/url"
	"strconv"
	"testing"

	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/cctp-relayer/db/sql/base"
	"github.com/synapsecns/sanguine/services/cctp-relayer/db/sql/sqlite"

	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/services/cctp-relayer/config"
	"github.com/synapsecns/sanguine/services/cctp-relayer/metadata"
	cctpTest "github.com/synapsecns/sanguine/services/cctp-relayer/testutil"
	omnirpcHelper "github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	scribeClient "github.com/synapsecns/sanguine/services/scribe/client"
	scribeHelper "github.com/synapsecns/sanguine/services/scribe/testhelper"
	"golang.org/x/sync/errgroup"
)

// TestHelperSuite defines the basic test suite.
type CCTPRelayerSuite struct {
	*testsuite.TestSuite
	// testBackends contains a list of all test backends
	testBackends []backends.SimulatedTestBackend
	// we'll use this later
	deployManager *cctpTest.DeployManager
	// testScribeURL setup in SetupTest
	testScribe string
	// testOmnirpc setup in SetupTest
	testOmnirpc string
	// metricsHandler is the metrics handler for the test
	metricsHandler metrics.Handler
	// testStore is the test store for the test
	testStore *base.Store
	// testWallet is the test wallet for the test
	testWallet wallet.Wallet
}

// NewTestSuite creates a new test suite.
func NewTestSuite(tb testing.TB) *CCTPRelayerSuite {
	tb.Helper()
	return &CCTPRelayerSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (s *CCTPRelayerSuite) SetupSuite() {
	s.TestSuite.SetupSuite()
	// for tracing
	localmetrics.SetupTestJaeger(s.GetSuiteContext(), s.T())
	// let's create 2 mock chains
	chainIDs := []uint64{1, 43114}

	// preallocate a slice for testbackends to the size of chainIDs
	// this way we can avoid non-deterministic order + needing to acquire/release a lock
	s.testBackends = make([]backends.SimulatedTestBackend, len(chainIDs))

	g, _ := errgroup.WithContext(s.GetSuiteContext())
	for i, chainID := range chainIDs {
		pos := i           // get position of chain id in array
		chainID := chainID // capture func literal
		g.Go(func() error {
			// we need to use the embedded backend here, because the simulated backend doesn't support rpcs required by scribe
			backend := geth.NewEmbeddedBackendForChainID(s.GetSuiteContext(), s.T(), new(big.Int).SetUint64(chainID))

			// add the backend to the list of backends
			s.testBackends[pos] = backend
			return nil
		})
	}

	// wait for all backends to be ready
	if err := g.Wait(); err != nil {
		s.T().Fatal(err)
	}

	// fund test wallet with ether
	var err error
	s.testWallet, err = wallet.FromRandom()
	s.Nil(err)
	for _, backend := range s.testBackends {
		backend.FundAccount(s.GetSuiteContext(), s.testWallet.Address(), *big.NewInt(params.Ether))
	}
}

// TODO: there should be a way to do this in the deployer, this probably involves making the concept of a contract-registry
// multi-chain (possibly by wrapping the registry). This would allow the use of setting remotes in the deployer itself rather than here.
func (s *CCTPRelayerSuite) registerRemoteDeployments() {
	for _, backend := range s.testBackends {
		cctpContract, cctpHandle := s.deployManager.GetSynapseCCTP(s.GetTestContext(), backend)
		_, tokenMessengeHandle := s.deployManager.GetMockTokenMessengerType(s.GetTestContext(), backend)

		// on the above contract, set the remote for each backend
		for _, backendToSetFrom := range s.testBackends {
			// we don't need to set the backends own remote!
			if backendToSetFrom.GetChainID() == backend.GetChainID() {
				continue
			}

			remoteCCTP, _ := s.deployManager.GetSynapseCCTP(s.GetTestContext(), backendToSetFrom)
			remoteMessenger, _ := s.deployManager.GetMockTokenMessengerType(s.GetTestContext(), backendToSetFrom)

			txOpts := backend.GetTxContext(s.GetTestContext(), cctpContract.OwnerPtr())
			// set the remote cctp contract on this cctp contract
			// TODO: verify chainID / domain are correct
			remoteDomain := cctpTest.ChainIDDomainMap[uint32(remoteCCTP.ChainID().Int64())]

			tx, err := cctpHandle.SetRemoteDomainConfig(txOpts.TransactOpts,
				big.NewInt(remoteCCTP.ChainID().Int64()), remoteDomain, remoteCCTP.Address())
			s.Require().NoError(err)
			backend.WaitForConfirmation(s.GetTestContext(), tx)

			// register the remote token messenger on the tokenMessenger contract
			_, err = tokenMessengeHandle.SetRemoteTokenMessenger(txOpts.TransactOpts, uint32(backendToSetFrom.GetChainID()), addressToBytes32(remoteMessenger.Address()))
			s.Nil(err)
		}
	}
}

// CCTPPrefix is the prefix for all CCTP tokens.
const CCTPPrefix = "CCTP."

func (s *CCTPRelayerSuite) registerTokens() {
	for _, backend := range s.testBackends {
		_, tokenHandle := s.deployManager.GetMockMintBurnTokenType(s.GetTestContext(), backend)
		cctpContract, cctpHandle := s.deployManager.GetSynapseCCTP(s.GetTestContext(), backend)

		txOpts := backend.GetTxContext(s.GetTestContext(), cctpContract.OwnerPtr())
		tx, err := cctpHandle.AddToken(txOpts.TransactOpts, fmt.Sprintf("%sUSDC", CCTPPrefix), tokenHandle.Address(), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))
		s.Require().NoError(err)
		backend.WaitForConfirmation(s.GetTestContext(), tx)
	}
}

// addressToBytes32 converts an address to a bytes32.
func addressToBytes32(addr common.Address) [32]byte {
	var buf [32]byte
	copy(buf[:], addr[:])
	return buf
}

func (s *CCTPRelayerSuite) SetupTest() {
	s.TestSuite.SetupTest()

	s.deployManager = cctpTest.NewDeployManager(s.T())

	// create the test scribe backend
	s.testScribe = scribeHelper.NewTestScribe(s.GetTestContext(), s.T(), s.deployManager.GetDeployedContracts(), s.testBackends...)
	// create the test omnirpc backend
	s.testOmnirpc = omnirpcHelper.NewOmnirpcServer(s.GetTestContext(), s.T(), s.testBackends...)

	// create the test metrics handler
	var err error
	// don't use metrics on ci for integration tests
	isCI := core.GetEnvBool("CI", false)
	useMetrics := !isCI
	metricsHandler := metrics.Null

	if useMetrics {
		localmetrics.SetupTestJaeger(s.GetSuiteContext(), s.T())
		metricsHandler = metrics.Jaeger
	}
	s.metricsHandler, err = metrics.NewByType(s.GetTestContext(), metadata.BuildInfo(), metricsHandler)
	s.Require().NoError(err)

	// create the test store
	path := filet.TmpDir(s.T(), "")
	db, err := sqlite.NewSqliteStore(s.GetTestContext(), path, s.metricsHandler, false)
	s.Require().NoError(err)
	s.testStore = base.NewStore(db.DB(), s.metricsHandler)

	// deploy the contract to all backends
	s.deployManager.BulkDeploy(s.GetTestContext(), s.testBackends, cctpTest.SynapseCCTPType, cctpTest.MockMintBurnTokenType)

	s.registerRemoteDeployments()
	s.registerTokens()
}

func (s *CCTPRelayerSuite) GetTestConfig() config.Config {
	cfg := config.Config{
		BaseOmnirpcURL: s.testBackends[0].RPCAddress(),
		Signer: signerConfig.SignerConfig{
			Type: signerConfig.FileType.String(),
			File: filet.TmpFile(s.T(), "", s.testWallet.PrivateKeyHex()).Name(),
		},
	}
	chains := []config.ChainConfig{}
	for _, backend := range s.testBackends {
		_, handle := s.deployManager.GetSynapseCCTP(s.GetTestContext(), backend)
		chains = append(chains, config.ChainConfig{
			ChainID:            uint32(backend.GetChainID()),
			SynapseCCTPAddress: handle.Address().String(),
		})
	}
	cfg.Chains = chains
	return cfg
}

func (s *CCTPRelayerSuite) GetTestScribe() scribeClient.ScribeClient {
	parsedScribe, err := url.Parse(s.testScribe)
	s.Nil(err)
	port, err := strconv.Atoi(parsedScribe.Opaque)
	s.Nil(err)

	sc := scribeClient.NewRemoteScribe(uint16(port), parsedScribe.Host, s.metricsHandler)
	return sc.ScribeClient
}

func TestCCTPRelayerSuite(t *testing.T) {
	suite.Run(t, NewTestSuite(t))
}
