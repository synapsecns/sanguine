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

func (i *InventoryTestSuite) TestGetRebalance() {
	origin := 1
	dest := 2
	usdcDataOrigin := inventory.TokenMetadata{
		Name:     "USDC",
		Decimals: 6,
		ChainID:  origin,
		Addr:     common.HexToAddress("0x0000000000000000000000000000000000000123"),
	}
	usdcDataDest := inventory.TokenMetadata{
		Name:     "USDC",
		Decimals: 6,
		ChainID:  dest,
		Addr:     common.HexToAddress("0x0000000000000000000000000000000000000456"),
	}
	tokens := map[int]map[common.Address]*inventory.TokenMetadata{
		origin: {
			usdcDataOrigin.Addr: &usdcDataOrigin,
		},
		dest: {
			usdcDataDest.Addr: &usdcDataDest,
		},
	}
	cfg := relconfig.Config{
		Chains: map[int]relconfig.ChainConfig{
			origin: {
				Tokens: map[string]relconfig.TokenConfig{
					"USDC": {
						Address:               usdcDataOrigin.Addr.Hex(),
						MaintenanceBalancePct: 10,
						InitialBalancePct:     30,
					},
				},
			},
			dest: {
				Tokens: map[string]relconfig.TokenConfig{
					"USDC": {
						Address:               usdcDataDest.Addr.Hex(),
						MaintenanceBalancePct: 10,
						InitialBalancePct:     30,
					},
				},
			},
		},
	}

	// 10 USDC on both chains; no rebalance needed
	usdcDataOrigin.Balance = big.NewInt(1e7)
	usdcDataDest.Balance = big.NewInt(1e7)
	rebalance, err := inventory.GetRebalance(cfg, tokens, origin, usdcDataOrigin.Addr)
	i.NoError(err)
	i.Nil(rebalance)
}
