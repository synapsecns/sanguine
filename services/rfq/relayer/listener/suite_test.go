package listener_test

import (
	"math/big"
	"testing"

	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/listener"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb/sqlite"
	"github.com/synapsecns/sanguine/services/rfq/testutil"
)

const chainID = 10

type ListenerTestSuite struct {
	*testsuite.TestSuite
	manager            *testutil.DeployManager
	backend            backends.SimulatedTestBackend
	store              reldb.Service
	metrics            metrics.Handler
	fastBridge         *fastbridge.FastBridgeRef
	fastBridgeMetadata contracts.DeployedContract
}

func NewListenerSuite(tb testing.TB) *ListenerTestSuite {
	return &ListenerTestSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func TestListenerSuite(t *testing.T) {
	suite.Run(t, NewListenerSuite(t))
}

func (l *ListenerTestSuite) SetupTest() {
	l.TestSuite.SetupTest()

	l.manager = testutil.NewDeployManager(l.T())
	l.backend = geth.NewEmbeddedBackendForChainID(l.GetTestContext(), l.T(), big.NewInt(chainID))
	var err error
	l.metrics = metrics.NewNullHandler()
	l.store, err = sqlite.NewSqliteStore(l.GetTestContext(), filet.TmpDir(l.T(), ""), l.metrics)
	l.Require().NoError(err)

	l.fastBridgeMetadata, l.fastBridge = l.manager.GetFastBridge(l.GetTestContext(), l.backend)
}

func (l *ListenerTestSuite) TestGetMetadataNoStore() {
	deployBlock, err := l.fastBridge.DeployBlock(&bind.CallOpts{Context: l.GetTestContext()})
	l.NoError(err)

	// nothing stored, should use start block
	cl := listener.NewTestChainListener(listener.TestChainListenerArgs{
		Address:      l.fastBridge.Address(),
		InitialBlock: deployBlock.Uint64(),
		Client:       l.backend,
		Store:        l.store,
		Handler:      l.metrics,
	})

	startBlock, myChainID, err := cl.GetMetadata(l.GetTestContext())
	l.NoError(err)
	l.Equal(myChainID, uint64(chainID))
	l.Equal(startBlock, deployBlock.Uint64())
}

func (l *ListenerTestSuite) TestStartBlock() {
	cl := listener.NewTestChainListener(listener.TestChainListenerArgs{
		Address: l.fastBridge.Address(),
		Client:  l.backend,
		Store:   l.store,
		Handler: l.metrics,
	})

	deployBlock, err := l.fastBridge.DeployBlock(&bind.CallOpts{Context: l.GetTestContext()})
	l.NoError(err)

	expectedLastIndexed := deployBlock.Uint64() + 10
	err = l.store.PutLatestBlock(l.GetTestContext(), chainID, expectedLastIndexed)
	l.NoError(err)

	startBlock, cid, err := cl.GetMetadata(l.GetTestContext())
	l.Equal(cid, uint64(chainID))
	l.Equal(startBlock, expectedLastIndexed)
}

func (l *ListenerTestSuite) TestListen() {

}
