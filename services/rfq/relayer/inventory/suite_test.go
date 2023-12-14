package inventory_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/omnirpc/testhelper"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
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

	for it := 1; it < 5; it++ {
		wg.Add(1)
		it := it // capture func literal
		go func() {
			defer wg.Done()
			backend := geth.NewEmbeddedBackendForChainID(i.GetTestContext(), i.T(), big.NewInt(int64(it)))
			mux.Lock()
			i.backends[it] = backend
			allBackends = append(allBackends, backend)
			mux.Unlock()

			// mint differing amounts on each chain of 1 eth * chainID.
			dc, mockContract := i.manager.GetMockERC20(i.GetTestContext(), i.backends[it])
			txContext := backend.GetTxContext(i.GetTestContext(), dc.OwnerPtr())

			tx, err := mockContract.Mint(txContext.TransactOpts, i.relayer.Address(), new(big.Int).Mul(big.NewInt(int64(it)), big.NewInt(params.Ether)))
			i.Require().NoError(err)

			backend.WaitForConfirmation(i.GetTestContext(), tx)
		}()
	}

	wg.Wait()

	i.omnirpcURL = testhelper.NewOmnirpcServer(i.GetTestContext(), i.T(), allBackends...)
	fmt.Print("hi")
}

func (i *InventoryTestSuite) TestInventoryBoot() {
	i.NotPanics(func() {
		cfg := relconfig.Config{
			Tokens: map[int][]string{},
		}

		for _, backend := range i.backends {
			handle, _ := i.manager.GetMockERC20(i.GetTestContext(), backend)
			cfg.Tokens[int(backend.GetChainID())] = []string{handle.Address().String()}
		}

		_, err := inventory.NewInventoryManager(i.GetTestContext(), omnirpcClient.NewOmnirpcClient(i.omnirpcURL, metrics.Get()), metrics.Get(), cfg, i.relayer.Address())
		i.Require().NoError(err)
	})
}
