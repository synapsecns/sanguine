package service_test

import (
	"math/big"
	"sync"
	"testing"
	"time"

	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	omnirpcHelper "github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/fastbridgemock"
	quoterMock "github.com/synapsecns/sanguine/services/rfq/relayer/quoter/mocks"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
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

	originContract, _ := r.manager.GetMockFastBridge(r.GetTestContext(), r.originBackend)
	destContract, _ := r.manager.GetMockFastBridge(r.GetTestContext(), r.destBackend)
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
		OmniRPCURL:  serverURL,
		BlockWindow: 5,      // 5 blocks cannot surpass $10k relay volume
		VolumeLimit: 10_000, // $10k usd
	}
}

func (r *RelayerTestSuite) TestStore() {
	r.T().Skip("TODO, test storage")

	rel, err := service.NewRelayer(r.GetTestContext(), r.metrics, r.cfg)
	r.NoError(err)

	go func() {
		r.NoError(rel.StartChainParser(r.GetTestContext()))
	}()

	_, oc := r.manager.GetMockFastBridge(r.GetTestContext(), r.originBackend)

	auth := r.originBackend.GetTxContext(r.GetTestContext(), nil)

	_, originToken := r.manager.GetMockERC20(r.GetTestContext(), r.originBackend)
	r.NoError(err)

	_, destToken := r.manager.GetMockERC20(r.GetTestContext(), r.destBackend)
	r.NoError(err)

	//nolint: typecheck
	tx, err := oc.MockBridgeRequest(auth.TransactOpts, [32]byte(crypto.Keccak256([]byte("3"))), mocks.MockAddress(), fastbridgemock.IFastBridgeBridgeParams{
		DstChainId:   uint32(r.destBackend.GetChainID()),
		To:           mocks.MockAddress(),
		OriginToken:  originToken.Address(),
		DestToken:    destToken.Address(),
		OriginAmount: big.NewInt(1),
		DestAmount:   big.NewInt(2),
		Deadline:     big.NewInt(3),
	})
	r.originBackend.WaitForConfirmation(r.GetTestContext(), tx)

	r.T().Skip("TODO, test storage")
	// TODO: check db
	time.Sleep(time.Second * 1000)
}

func (r *RelayerTestSuite) TestCommit() {
	r.T().Skip("TODO, test storage")

	rel, err := service.NewRelayer(r.GetTestContext(), r.metrics, r.cfg)
	r.NoError(err)

	go func() {
		r.NoError(rel.StartChainParser(r.GetTestContext()))
	}()

	_, oc := r.manager.GetMockFastBridge(r.GetTestContext(), r.originBackend)

	auth := r.originBackend.GetTxContext(r.GetTestContext(), nil)

	_, originToken := r.manager.GetMockERC20(r.GetTestContext(), r.originBackend)
	r.NoError(err)

	_, destToken := r.manager.GetMockERC20(r.GetTestContext(), r.destBackend)
	r.NoError(err)

	//nolint: typecheck
	tx, err := oc.MockBridgeRequest(auth.TransactOpts, [32]byte(crypto.Keccak256([]byte("3"))), mocks.MockAddress(), fastbridgemock.IFastBridgeBridgeParams{
		DstChainId:   uint32(r.destBackend.GetChainID()),
		To:           mocks.MockAddress(),
		OriginToken:  originToken.Address(),
		DestToken:    destToken.Address(),
		OriginAmount: big.NewInt(1),
		DestAmount:   big.NewInt(2),
		Deadline:     big.NewInt(3),
	})
	r.originBackend.WaitForConfirmation(r.GetTestContext(), tx)

	r.T().Skip("TODO, test storage")
	// TODO: check db
	time.Sleep(time.Second * 100000)

}

func (r *RelayerTestSuite) TestRateLimit() {
	rel, err := service.NewRelayer(r.GetTestContext(), r.metrics, r.cfg)
	r.NoError(err)

	// set up quoter
	quoter := new(quoterMock.Quoter)
	rel.SetQuoter(quoter)
	quoter.On("GetPrice", mock.Anything, mock.Anything).Return(10_001, nil)

	// deploy some ERC20s
	_, originToken := r.manager.GetMockERC20(r.GetTestContext(), r.originBackend)
	r.NoError(err)

	_, destToken := r.manager.GetMockERC20(r.GetTestContext(), r.destBackend)
	r.NoError(err)

	// set up the bridge and the auth signer
	_, oc := r.manager.GetMockFastBridge(r.GetTestContext(), r.originBackend)
	auth := r.originBackend.GetTxContext(r.GetTestContext(), nil)

	// start listening for transactions
	go func() {
		r.NoError(rel.StartChainParser(r.GetTestContext()))
	}()

	addy := mocks.MockAddress()

	// send the bridge request that should get rate limited
	tx, err := oc.MockBridgeRequest(
		auth.TransactOpts,
		[32]byte(crypto.Keccak256([]byte("3"))),
		addy,
		fastbridgemock.IFastBridgeBridgeParams{
			DstChainId:   uint32(r.destBackend.GetChainID()),
			To:           mocks.MockAddress(),
			OriginToken:  originToken.Address(),
			DestToken:    destToken.Address(),
			OriginAmount: big.NewInt(1),
			DestAmount:   big.NewInt(2),
			Deadline:     big.NewInt(3),
		})
	r.NoError(err)

	// get the current block
	currentBlock, err := r.originBackend.BlockByNumber(r.GetTestContext(), nil)
	r.NoError(err)
	r.originBackend.WaitForConfirmation(r.GetTestContext(), tx)
	receipt, err := r.originBackend.TransactionReceipt(r.GetTestContext(), tx.Hash())
	r.NoError(err)

	// should hagve waited.
	r.Greater(receipt.BlockNumber.Uint64(), currentBlock.Number)

	r.Eventually(
		func() bool {
			txs, err := rel.DB().SubmitterDB().GetAllTXAttemptByStatus(
				r.GetTestContext(),
				auth.From,
				r.originBackend.ChainConfig().ChainID,
				db.WithStatuses(db.Confirmed),
			)
			if err != nil {
				return false
			}
			return len(txs) == 1
		},
	)
}
