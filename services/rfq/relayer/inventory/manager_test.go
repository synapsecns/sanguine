package inventory_test

import (
	"math/big"
	"sync"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/backends"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
)

func (i *InventoryTestSuite) TestInventoryBootAndRefresh() {
	// setup a mux to keep track of how much we're actually minting.
	localTokens := map[int]map[common.Address]*big.Int{}
	_ = localTokens

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
		Chains: map[int]relconfig.ChainConfig{},
	}

	for _, backend := range i.backends {
		handle, _ := i.manager.GetMockERC20(i.GetTestContext(), backend)
		cfg.Chains[int(backend.GetChainID())] = relconfig.ChainConfig{
			Tokens: map[string]relconfig.TokenConfig{
				"USDC": {
					Address: handle.Address().String(),
				},
			},
		}
	}

	im, err := inventory.NewInventoryManager(i.GetTestContext(), omnirpcClient.NewOmnirpcClient(i.omnirpcURL, metrics.Get()), metrics.Get(), cfg, i.relayer.Address(), nil, i.db)
	i.Require().NoError(err)

	_ = im
}

func (i *InventoryTestSuite) TestHasSufficientGas() {
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

	// TODO: these chain IDs are hardcoded; should probably be assigned in suite as fields
	origin := 1
	dest := 2

	getManager := func(gasThresholds []*big.Int) inventory.Manager {
		cfg := relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{},
		}
		for idx, chainID := range []int{origin, dest} {
			backend := i.backends[chainID]
			handle, _ := i.manager.GetMockERC20(i.GetTestContext(), backend)
			cfg.Chains[int(backend.GetChainID())] = relconfig.ChainConfig{
				MinGasToken: gasThresholds[idx].String(),
				Tokens: map[string]relconfig.TokenConfig{
					"USDC": {
						Address:  handle.Address().String(),
						Decimals: 6,
					},
				},
			}
		}

		im, err := inventory.NewInventoryManager(i.GetTestContext(), omnirpcClient.NewOmnirpcClient(i.omnirpcURL, metrics.Get()), metrics.Get(), cfg, i.relayer.Address(), nil, i.db)
		i.Require().NoError(err)
		return im
	}

	im := getManager([]*big.Int{big.NewInt(params.Ether), big.NewInt(params.Ether)})
	sufficient, err := im.HasSufficientGas(i.GetTestContext(), origin, nil)
	i.NoError(err)
	i.True(sufficient)
	sufficient, err = im.HasSufficientGas(i.GetTestContext(), dest, nil)
	i.NoError(err)
	i.True(sufficient)

	// test with nonzero gasValue
	gasValue := new(big.Int).Mul(big.NewInt(params.Ether), big.NewInt(10))
	sufficient, err = im.HasSufficientGas(i.GetTestContext(), origin, gasValue)
	i.NoError(err)
	i.False(sufficient)

	// multiply big int to avoid overflow
	largeBalance := new(big.Int).Mul(big.NewInt(params.Ether), big.NewInt(100))
	im = getManager([]*big.Int{largeBalance, big.NewInt(params.Ether)})
	sufficient, err = im.HasSufficientGas(i.GetTestContext(), origin, nil)
	i.NoError(err)
	i.False(sufficient)
	sufficient, err = im.HasSufficientGas(i.GetTestContext(), dest, nil)
	i.NoError(err)
	i.True(sufficient)

	im = getManager([]*big.Int{big.NewInt(params.Ether), largeBalance})
	sufficient, err = im.HasSufficientGas(i.GetTestContext(), origin, nil)
	i.NoError(err)
	i.True(sufficient)
	sufficient, err = im.HasSufficientGas(i.GetTestContext(), dest, nil)
	i.NoError(err)
	i.False(sufficient)
}
