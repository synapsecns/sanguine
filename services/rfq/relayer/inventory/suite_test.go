package inventory_test

import (
	"github.com/Flaque/filet"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/base"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb/sqlite"
	"github.com/synapsecns/sanguine/services/rfq/testutil"
	"math/big"
	"sync"
	"testing"
)

// InventoryTestSuite is the test suite used for testing the inventory manager.
type InventoryTestSuite struct {
	*testsuite.TestSuite
	backends   map[int]backends.SimulatedTestBackend
	manager    *testutil.DeployManager
	relayer    wallet.Wallet
	omnirpcURL string
	db         reldb.Service
}

// NewInventorySuite creates the inventory suite.
func NewInventorySuite(tb testing.TB) *InventoryTestSuite {
	return &InventoryTestSuite{
		TestSuite: testsuite.NewTestSuite(tb),
		backends:  map[int]backends.SimulatedTestBackend{},
	}
}

func TestInventorySuite(t *testing.T) {
	suite.Run(t, NewInventorySuite(t))
}

func (i *InventoryTestSuite) TearDownTest() {
	i.TestSuite.TearDownTest()

	i.backends = map[int]backends.SimulatedTestBackend{}
	i.manager = nil
	i.relayer = nil
}

func (i *InventoryTestSuite) SetupTest() {
	var err error
	i.TestSuite.SetupTest()
	i.manager = testutil.NewDeployManager(i.T())
	i.relayer, err = wallet.FromRandom()
	i.Require().NoError(err)

	// used for omnirpc client construction
	var allBackends []backends.SimulatedTestBackend
	var mux sync.Mutex
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		i.db, err = sqlite.NewSqliteStore(i.GetTestContext(), filet.TmpDir(i.T(), ""), metrics.Get(), false)
		i.NoError(err)
	}()

	for it := 1; it < 3; it++ {
		wg.Add(1)
		it := it // capture func literal
		go func() {
			defer wg.Done()
			backend := geth.NewEmbeddedBackendForChainID(i.GetTestContext(), i.T(), big.NewInt(int64(it)))
			mux.Lock()
			i.backends[it] = backend
			allBackends = append(allBackends, backend)
			mux.Unlock()

			backend.Store(base.WalletToKey(i.T(), i.relayer))
			backend.FundAccount(i.GetTestContext(), i.relayer.Address(), *new(big.Int).Mul(big.NewInt(params.Ether), big.NewInt(10)))
		}()
	}

	wg.Wait()

	i.omnirpcURL = testhelper.NewOmnirpcServer(i.GetTestContext(), i.T(), allBackends...)
}
