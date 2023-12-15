package inventory_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/backends"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"math/big"
	"sync"
)

func (i *InventoryTestSuite) TestInventoryBootAndRefresh() {
	// setup a mux to keep track of how much we're actually minting.
	localTokens := map[int]map[common.Address]*big.Int{}
	var mux sync.Mutex
	_ = localTokens
	_ = mux

	var wg sync.WaitGroup
	wg.Add(len(i.backends))
	for _, backend := range i.backends {
		go func(backend backends.SimulatedTestBackend) {
			defer wg.Done()

			// fund the relayer
			mintMulAmount := big.NewInt(int64(gofakeit.Number(0, 6)))
			mintAmount := new(big.Int).Mul(mintMulAmount, big.NewInt(params.Ether))

			metadata, usdt := i.manager.GetUSDT(i.GetTestContext(), backend)
			_ = metadata
			_ = usdt
			_ = mintAmount
		}(backend)
	}
	wg.Wait()

	cfg := relconfig.Config{
		Tokens: map[int][]string{},
	}

	for _, backend := range i.backends {
		handle, _ := i.manager.GetMockERC20(i.GetTestContext(), backend)
		cfg.Tokens[int(backend.GetChainID())] = []string{handle.Address().String()}
	}

	im, err := inventory.NewInventoryManager(i.GetTestContext(), omnirpcClient.NewOmnirpcClient(i.omnirpcURL, metrics.Get()), metrics.Get(), cfg, i.relayer.Address(), i.db)
	i.Require().NoError(err)

	_ = im
}
