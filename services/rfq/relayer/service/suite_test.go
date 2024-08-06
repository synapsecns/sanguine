package service_test

import (
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/Flaque/filet"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	ethConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	omnirpcHelper "github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/relayer/service"
	"github.com/synapsecns/sanguine/services/rfq/testutil"
)

type RelayerTestSuite struct {
	*testsuite.TestSuite
	originBackend backends.SimulatedTestBackend
	destBackend   backends.SimulatedTestBackend
	manager       *testutil.DeployManager
	metrics       metrics.Handler
	cfg           relconfig.Config
}

func NewRelayerTestSuite(tb testing.TB) *RelayerTestSuite {
	return &RelayerTestSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestRelayerTestSuite(t *testing.T) {
	suite.Run(t, NewRelayerTestSuite(t))
}

func (r *RelayerTestSuite) SetupTest() {
	r.TestSuite.SetupTest()
	r.manager = testutil.NewDeployManager(r.T())
	r.metrics = metrics.NewNullHandler()

	testWallet, err := wallet.FromRandom()
	r.NoError(err)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		r.originBackend = geth.NewEmbeddedBackendForChainID(r.GetTestContext(), r.T(), big.NewInt(1))
	}()
	go func() {
		defer wg.Done()
		r.destBackend = geth.NewEmbeddedBackendForChainID(r.GetTestContext(), r.T(), big.NewInt(2))
	}()
	wg.Wait()

	serverURL := omnirpcHelper.NewOmnirpcServer(r.GetTestContext(), r.T(), r.destBackend, r.originBackend)

	originContract, _ := r.manager.GetFastBridge(r.GetTestContext(), r.originBackend)
	destContract, _ := r.manager.GetFastBridge(r.GetTestContext(), r.destBackend)

	r.cfg = relconfig.Config{
		Database: relconfig.DatabaseConfig{
			Type: dbcommon.Sqlite.String(),
			DSN:  filet.TmpDir(r.T(), ""),
		},
		Chains: map[int]relconfig.ChainConfig{
			int(r.originBackend.GetChainID()): {
				RFQAddress: originContract.Address().String(),
			},
			int(r.destBackend.GetChainID()): {
				RFQAddress: destContract.Address().String(),
			},
		},
		Signer: ethConfig.SignerConfig{
			Type: ethConfig.FileType.String(),
			File: filet.TmpFile(r.T(), "", testWallet.PrivateKeyHex()).Name(),
		},
		OmniRPCURL:  serverURL,
		BlockWindow: 5,      // 5 blocks cannot surpass $10k relay volume
		VolumeLimit: 10_000, // $10k usd
	}
}

// For a singular transaction over $10k.
func (r *RelayerTestSuite) TestRateLimit() {
	rel, err := service.NewRelayer(r.GetTestContext(), r.metrics, r.cfg)
	r.NoError(err)

	// deploy some ERC20s
	_, originToken := r.manager.GetMockERC20(r.GetTestContext(), r.originBackend)
	r.NoError(err)

	_, destToken := r.manager.GetMockERC20(r.GetTestContext(), r.destBackend)
	r.NoError(err)

	// set up the bridge and the auth signer
	_, oc := r.manager.GetFastBridge(r.GetTestContext(), r.originBackend)
	auth := r.originBackend.GetTxContext(r.GetTestContext(), nil)

	// start listening for transactions
	go func() {
		r.NoError(rel.StartChainParser(r.GetTestContext()))
	}()

	tx, err := oc.Bridge(
		auth.TransactOpts,
		fastbridge.IFastBridgeBridgeParams{
			DstChainId:   uint32(r.destBackend.GetChainID()),
			To:           mocks.MockAddress(),
			OriginToken:  originToken.Address(),
			DestToken:    destToken.Address(),
			OriginAmount: big.NewInt(1),
			DestAmount:   big.NewInt(2),
			Deadline:     big.NewInt(time.Now().Unix()),
		})
	r.NoError(err)

	r.originBackend.WaitForConfirmation(r.GetTestContext(), tx)
	receipt, err := r.originBackend.TransactionReceipt(r.GetTestContext(), tx.Hash())
	r.NoError(err)

	var good bool
	for {
		tx, err := rel.DB().GetQuoteRequestByOriginTxHash(
			r.GetTestContext(),
			tx.Hash(),
		)
		if err != nil {
			r.T().Error(err)
		}
		if tx == nil {
			continue
		}

		currentBlock, err := r.originBackend.BlockByNumber(r.GetTestContext(), nil)
		if err != nil {
			r.T().Error(err)
		}

		if tx.Status == reldb.RelayCompleted && currentBlock.Number().Cmp(receipt.BlockNumber) == 1 {
			good = true
			break
		}

		time.Sleep(5 * time.Second)
	}

	r.True(good)

	// r.Eventually(
	// 	func() bool {
	// 		txs, err := rel.DB().GetStatusCounts(
	// 			r.GetTestContext(),
	// 			reldb.RelayCompleted,
	// 		)
	// 		if err != nil {
	// 			return false
	// 		}

	// 		currentBlock, err := r.originBackend.BlockByNumber(r.GetTestContext(), nil)
	// 		if err != nil {
	// 			return false
	// 		}

	// 		return txs[reldb.RelayCompleted] >= 1 && currentBlock.Number().Cmp(receipt.BlockNumber) == 1
	// 	},
	// )
}
