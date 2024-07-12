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
	extra := 3
	usdcDataOrigin := inventory.TokenMetadata{
		Name:     "USDC",
		Decimals: 6,
		ChainID:  origin,
		Addr:     common.HexToAddress("0x0000000000000000000000000000000000000123"),
		Balance:  big.NewInt(0),
	}
	usdcDataDest := inventory.TokenMetadata{
		Name:     "USDC",
		Decimals: 6,
		ChainID:  dest,
		Addr:     common.HexToAddress("0x0000000000000000000000000000000000000456"),
		Balance:  big.NewInt(0),
	}
	usdcDataExtra := inventory.TokenMetadata{
		Name:     "USDC",
		Decimals: 6,
		ChainID:  extra,
		Addr:     common.HexToAddress("0x0000000000000000000000000000000000000789"),
		Balance:  big.NewInt(0),
	}
	tokens := map[int]map[common.Address]*inventory.TokenMetadata{
		origin: {
			usdcDataOrigin.Addr: &usdcDataOrigin,
		},
		dest: {
			usdcDataDest.Addr: &usdcDataDest,
		},
	}
	tokensWithExtra := map[int]map[common.Address]*inventory.TokenMetadata{
		origin: {
			usdcDataOrigin.Addr: &usdcDataOrigin,
		},
		dest: {
			usdcDataDest.Addr: &usdcDataDest,
		},
		extra: {
			usdcDataExtra.Addr: &usdcDataExtra,
		},
	}
	getConfig := func(minRebalanceAmount, maxRebalanceAmount string, originMethod, destMethod relconfig.RebalanceMethod) relconfig.Config {
		return relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{
				origin: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               usdcDataOrigin.Addr.Hex(),
							Decimals:              6,
							MaintenanceBalancePct: 20,
							InitialBalancePct:     50,
							MinRebalanceAmount:    minRebalanceAmount,
							MaxRebalanceAmount:    maxRebalanceAmount,
							RebalanceMethods:      []string{originMethod.String()},
						},
					},
				},
				dest: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               usdcDataDest.Addr.Hex(),
							Decimals:              6,
							MaintenanceBalancePct: 20,
							InitialBalancePct:     50,
							MinRebalanceAmount:    minRebalanceAmount,
							MaxRebalanceAmount:    maxRebalanceAmount,
							RebalanceMethods:      []string{destMethod.String()},
						},
					},
				},
				extra: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               usdcDataExtra.Addr.Hex(),
							Decimals:              6,
							MaintenanceBalancePct: 0,
							InitialBalancePct:     0,
							MinRebalanceAmount:    minRebalanceAmount,
							MaxRebalanceAmount:    maxRebalanceAmount,
							RebalanceMethods:      []string{destMethod.String()},
						},
					},
				},
			},
		}
	}

	i.Run("EqualBalances", func() {
		// 10 USDC on both chains; no rebalance needed
		cfg := getConfig("", "", relconfig.RebalanceMethodSynapseCCTP, relconfig.RebalanceMethodSynapseCCTP)
		usdcDataOrigin.Balance = big.NewInt(1e7)
		usdcDataDest.Balance = big.NewInt(1e7)
		rebalance, err := inventory.GetRebalance(cfg, tokens, origin, usdcDataOrigin.Addr)
		i.NoError(err)
		i.Nil(rebalance)
	})

	i.Run("ZeroBalances", func() {
		// Set balances to zero
		cfg := getConfig("", "", relconfig.RebalanceMethodSynapseCCTP, relconfig.RebalanceMethodSynapseCCTP)
		usdcDataOrigin.Balance = big.NewInt(0)
		usdcDataDest.Balance = big.NewInt(0)
		rebalance, err := inventory.GetRebalance(cfg, tokens, origin, usdcDataOrigin.Addr)
		i.NoError(err)
		i.Nil(rebalance)
	})

	i.Run("BasicRebalance", func() {
		// Set dest balance below maintenance threshold; need rebalance
		cfg := getConfig("", "", relconfig.RebalanceMethodSynapseCCTP, relconfig.RebalanceMethodSynapseCCTP)
		usdcDataOrigin.Balance = big.NewInt(9e6)
		usdcDataDest.Balance = big.NewInt(1e6)
		rebalance, err := inventory.GetRebalance(cfg, tokens, dest, usdcDataDest.Addr)
		i.NoError(err)
		expected := &inventory.RebalanceData{
			OriginMetadata: &usdcDataOrigin,
			DestMetadata:   &usdcDataDest,
			Amount:         big.NewInt(4e6),
			Method:         relconfig.RebalanceMethodSynapseCCTP,
		}
		i.Equal(expected, rebalance)
	})

	i.Run("RebalanceMethodMismatch", func() {
		// Set rebalance methods to mismatch
		cfg := getConfig("", "", relconfig.RebalanceMethodCircleCCTP, relconfig.RebalanceMethodSynapseCCTP)
		rebalance, err := inventory.GetRebalance(cfg, tokens, dest, usdcDataDest.Addr)
		i.NoError(err)
		i.Nil(rebalance)
	})

	i.Run("OneRebalanceMethodNone", func() {
		// Set one rebalance method to None
		cfg := getConfig("", "", relconfig.RebalanceMethodNone, relconfig.RebalanceMethodSynapseCCTP)
		rebalance, err := inventory.GetRebalance(cfg, tokens, dest, usdcDataDest.Addr)
		i.NoError(err)
		i.Nil(rebalance)
	})

	i.Run("BelowMinRebalanceAmount", func() {
		// Set min rebalance amount
		cfgWithMax := getConfig("10", "1000000000", relconfig.RebalanceMethodSynapseCCTP, relconfig.RebalanceMethodSynapseCCTP)
		rebalance, err := inventory.GetRebalance(cfgWithMax, tokens, dest, usdcDataDest.Addr)
		i.NoError(err)
		i.Nil(rebalance)
	})

	i.Run("AboveMaxRebalanceAmount", func() {
		// Set max rebalance amount
		cfgWithMax := getConfig("0", "1.1", relconfig.RebalanceMethodSynapseCCTP, relconfig.RebalanceMethodSynapseCCTP)
		rebalance, err := inventory.GetRebalance(cfgWithMax, tokens, dest, usdcDataDest.Addr)
		i.NoError(err)
		expected := &inventory.RebalanceData{
			OriginMetadata: &usdcDataOrigin,
			DestMetadata:   &usdcDataDest,
			Amount:         big.NewInt(1.1e6),
			Method:         relconfig.RebalanceMethodSynapseCCTP,
		}
		i.Equal(expected, rebalance)
	})

	i.Run("BelowInitalThresholdOnOrigin", func() {
		// Increase initial threshold so that no rebalance can occur from origin
		cfg := getConfig("", "", relconfig.RebalanceMethodNone, relconfig.RebalanceMethodSynapseCCTP)
		usdcDataOrigin.Balance = big.NewInt(2e6)
		usdcDataDest.Balance = big.NewInt(1e6)
		usdcDataExtra.Balance = big.NewInt(7e6)
		rebalance, err := inventory.GetRebalance(cfg, tokens, dest, usdcDataDest.Addr)
		i.NoError(err)
		i.Nil(rebalance)
	})

	i.Run("SkipLowestBalanceWithMismatch", func() {
		// Set origin as lowest balance, but mismatched rebalance method, so next lowest balance
		// should be chosen
		cfg := relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{
				origin: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               usdcDataOrigin.Addr.Hex(),
							Decimals:              6,
							MaintenanceBalancePct: 20,
							InitialBalancePct:     40,
							MinRebalanceAmount:    "",
							MaxRebalanceAmount:    "",
							RebalanceMethods:      []string{"synapsecctp"},
						},
					},
				},
				dest: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               usdcDataDest.Addr.Hex(),
							Decimals:              6,
							MaintenanceBalancePct: 20,
							InitialBalancePct:     40,
							MinRebalanceAmount:    "",
							MaxRebalanceAmount:    "",
							RebalanceMethods:      []string{"circlecctp"},
						},
					},
				},
				extra: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               usdcDataExtra.Addr.Hex(),
							Decimals:              6,
							MaintenanceBalancePct: 20,
							InitialBalancePct:     40,
							MinRebalanceAmount:    "",
							MaxRebalanceAmount:    "",
							RebalanceMethods:      []string{"circlecctp"},
						},
					},
				},
			},
		}
		usdcDataOrigin.Balance = big.NewInt(0)
		usdcDataDest.Balance = big.NewInt(1e6)
		usdcDataExtra.Balance = big.NewInt(9e6)
		rebalance, err := inventory.GetRebalance(cfg, tokensWithExtra, dest, usdcDataDest.Addr)
		i.NoError(err)
		expected := &inventory.RebalanceData{
			OriginMetadata: &usdcDataExtra,
			DestMetadata:   &usdcDataDest,
			Amount:         big.NewInt(5e6),
			Method:         relconfig.RebalanceMethodCircleCCTP,
		}
		i.Equal(expected, rebalance)
	})

	i.Run("AllowMultipleRebalanceMethods", func() {
		// Origin has lowest balance and supports multiple rebalance methods, so it should be chosen
		cfg := relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{
				origin: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               usdcDataOrigin.Addr.Hex(),
							Decimals:              6,
							MaintenanceBalancePct: 20,
							InitialBalancePct:     40,
							MinRebalanceAmount:    "",
							MaxRebalanceAmount:    "",
							RebalanceMethods:      []string{"synapsecctp", "circlecctp"},
						},
					},
				},
				dest: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               usdcDataDest.Addr.Hex(),
							Decimals:              6,
							MaintenanceBalancePct: 20,
							InitialBalancePct:     40,
							MinRebalanceAmount:    "",
							MaxRebalanceAmount:    "",
							RebalanceMethods:      []string{"circlecctp"},
						},
					},
				},
				extra: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               usdcDataExtra.Addr.Hex(),
							Decimals:              6,
							MaintenanceBalancePct: 20,
							InitialBalancePct:     40,
							MinRebalanceAmount:    "",
							MaxRebalanceAmount:    "",
							RebalanceMethods:      []string{"circlecctp"},
						},
					},
				},
			},
		}
		usdcDataOrigin.Balance = big.NewInt(0)
		usdcDataDest.Balance = big.NewInt(1e6)
		usdcDataExtra.Balance = big.NewInt(9e6)
		rebalance, err := inventory.GetRebalance(cfg, tokensWithExtra, origin, usdcDataOrigin.Addr)
		i.NoError(err)
		expected := &inventory.RebalanceData{
			OriginMetadata: &usdcDataExtra,
			DestMetadata:   &usdcDataOrigin,
			Amount:         big.NewInt(5e6),
			Method:         relconfig.RebalanceMethodCircleCCTP,
		}
		i.Equal(expected, rebalance)
	})
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
