package service_test

import (
	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	omnirpcHelper "github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"github.com/synapsecns/sanguine/services/rfq/contracts/testcontracts/fastbridgemock"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/service"
	"github.com/synapsecns/sanguine/services/rfq/testutil"
	"math/big"
	"sync"
	"testing"
	"time"
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
		Bridges: map[int]relconfig.ChainConfig{
			int(r.originBackend.GetChainID()): {
				Bridge: originContract.Address().String(),
			},
			int(r.destBackend.GetChainID()): {
				Bridge: destContract.Address().String(),
			},
		},
		OmniRPCURL: serverURL,
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
