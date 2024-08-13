package inventory_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/util"
)

func TestGetRebalances(t *testing.T) {
	t.Run("NoRebalanceMethod", func(t *testing.T) {
		cfg := relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{
				1: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address: common.HexToAddress("0x1").Hex(),
						},
					},
				},
				2: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address: common.HexToAddress("0x2").Hex(),
						},
					},
				},
			},
		}
		tokens := map[int]map[common.Address]*inventory.TokenMetadata{
			1: {
				common.HexToAddress("0x1"): {
					Name:    "USDC",
					Balance: big.NewInt(1_000_000),
					ChainID: 1,
					Addr:    common.HexToAddress("0x1"),
				},
			},
			2: {
				common.HexToAddress("0x2"): {
					Name:    "USDC",
					Balance: big.NewInt(10_000_000),
					ChainID: 2,
					Addr:    common.HexToAddress("0x2"),
				},
			},
		}
		rebalances, err := inventory.GetRebalances(context.Background(), cfg, tokens)
		assert.Nil(t, err)
		assert.Empty(t, rebalances)
	})

	t.Run("IncompatibleRebalanceMethods", func(t *testing.T) {
		cfg := relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{
				1: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:          common.HexToAddress("0x1").Hex(),
							RebalanceMethods: []string{"circlecctp"},
						},
					},
				},
				2: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:          common.HexToAddress("0x2").Hex(),
							RebalanceMethods: []string{"synapsecctp"},
						},
					},
				},
			},
		}
		tokens := map[int]map[common.Address]*inventory.TokenMetadata{
			1: {
				common.HexToAddress("0x1"): {
					Name:    "USDC",
					Balance: big.NewInt(1_000_000),
					ChainID: 1,
					Addr:    common.HexToAddress("0x1"),
				},
			},
			2: {
				common.HexToAddress("0x2"): {
					Name:    "USDC",
					Balance: big.NewInt(10_000_000),
					ChainID: 2,
					Addr:    common.HexToAddress("0x2"),
				},
			},
		}
		rebalances, err := inventory.GetRebalances(context.Background(), cfg, tokens)
		assert.Nil(t, err)
		assert.Nil(t, rebalances["USDC"])
	})

	t.Run("RebalanceSingleMethod", func(t *testing.T) {
		cfg := relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{
				1: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               common.HexToAddress("0x1").Hex(),
							RebalanceMethods:      []string{"circlecctp"},
							MaintenanceBalancePct: 10,
							InitialBalancePct:     50,
						},
					},
				},
				2: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               common.HexToAddress("0x2").Hex(),
							RebalanceMethods:      []string{"circlecctp"},
							MaintenanceBalancePct: 10,
							InitialBalancePct:     50,
						},
					},
				},
			},
		}
		tokens := map[int]map[common.Address]*inventory.TokenMetadata{
			1: {
				common.HexToAddress("0x1"): {
					Name:    "USDC",
					Balance: big.NewInt(1_000_000),
					ChainID: 1,
					Addr:    common.HexToAddress("0x1"),
				},
			},
			2: {
				common.HexToAddress("0x2"): {
					Name:    "USDC",
					Balance: big.NewInt(9_000_000),
					ChainID: 2,
					Addr:    common.HexToAddress("0x2"),
				},
			},
		}
		rebalances, err := inventory.GetRebalances(context.Background(), cfg, tokens)
		assert.Nil(t, err)
		expectedRebalance := &inventory.RebalanceData{
			Method:         relconfig.RebalanceMethodCircleCCTP,
			OriginMetadata: tokens[2][common.HexToAddress("0x2")],
			DestMetadata:   tokens[1][common.HexToAddress("0x1")],
			Amount:         big.NewInt(4_000_000),
		}
		assert.Equal(t, rebalances["USDC"], expectedRebalance)
	})

	t.Run("RebalanceSingleMethodAboveMaintenance", func(t *testing.T) {
		cfg := relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{
				1: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               common.HexToAddress("0x1").Hex(),
							RebalanceMethods:      []string{"circlecctp"},
							MaintenanceBalancePct: 10,
							InitialBalancePct:     50,
						},
					},
				},
				2: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               common.HexToAddress("0x2").Hex(),
							RebalanceMethods:      []string{"circlecctp"},
							MaintenanceBalancePct: 10,
							InitialBalancePct:     50,
						},
					},
				},
			},
		}
		tokens := map[int]map[common.Address]*inventory.TokenMetadata{
			1: {
				common.HexToAddress("0x1"): {
					Name:    "USDC",
					Balance: big.NewInt(2_000_000),
					ChainID: 1,
					Addr:    common.HexToAddress("0x1"),
				},
			},
			2: {
				common.HexToAddress("0x2"): {
					Name:    "USDC",
					Balance: big.NewInt(8_000_000),
					ChainID: 2,
					Addr:    common.HexToAddress("0x2"),
				},
			},
		}
		rebalances, err := inventory.GetRebalances(context.Background(), cfg, tokens)
		assert.Nil(t, err)
		assert.Nil(t, rebalances["USDC"])
	})

	t.Run("RebalanceMultipleMethods", func(t *testing.T) {
		cfg := relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{
				1: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               common.HexToAddress("0x1").Hex(),
							RebalanceMethods:      []string{"circlecctp", "scroll"},
							MaintenanceBalancePct: 10,
							InitialBalancePct:     50,
						},
					},
				},
				2: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               common.HexToAddress("0x2").Hex(),
							RebalanceMethods:      []string{"circlecctp"},
							MaintenanceBalancePct: 10,
							InitialBalancePct:     50,
						},
					},
				},
				3: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               common.HexToAddress("0x3").Hex(),
							RebalanceMethods:      []string{"scroll"},
							MaintenanceBalancePct: 10,
							InitialBalancePct:     50,
						},
					},
				},
			},
		}
		tokens := map[int]map[common.Address]*inventory.TokenMetadata{
			1: {
				common.HexToAddress("0x1"): {
					Name:    "USDC",
					Balance: big.NewInt(9_000_000),
					ChainID: 1,
					Addr:    common.HexToAddress("0x1"),
				},
			},
			2: {
				common.HexToAddress("0x2"): {
					Name:    "USDC",
					Balance: big.NewInt(100_000),
					ChainID: 2,
					Addr:    common.HexToAddress("0x2"),
				},
			},
			3: {
				common.HexToAddress("0x3"): {
					Name:    "USDC",
					Balance: big.NewInt(900_000),
					ChainID: 3,
					Addr:    common.HexToAddress("0x3"),
				},
			},
		}
		rebalances, err := inventory.GetRebalances(context.Background(), cfg, tokens)
		assert.Nil(t, err)
		expectedRebalance := &inventory.RebalanceData{
			Method:         relconfig.RebalanceMethodCircleCCTP,
			OriginMetadata: tokens[1][common.HexToAddress("0x1")],
			DestMetadata:   tokens[2][common.HexToAddress("0x2")],
			Amount:         big.NewInt(4_450_000),
		}
		assert.Equal(t, rebalances["USDC"], expectedRebalance)
	})

	t.Run("RebalanceMultipleMethodsAndTokens", func(t *testing.T) {
		cfg := relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{
				1: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               common.HexToAddress("0x1").Hex(),
							RebalanceMethods:      []string{"circlecctp", "scroll"},
							MaintenanceBalancePct: 10,
							InitialBalancePct:     50,
						},
						"ETH": {
							Address:               util.EthAddress.Hex(),
							RebalanceMethods:      []string{"circlecctp", "scroll"},
							MaintenanceBalancePct: 10,
							InitialBalancePct:     50,
						},
					},
				},
				2: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               common.HexToAddress("0x2").Hex(),
							RebalanceMethods:      []string{"circlecctp"},
							MaintenanceBalancePct: 10,
							InitialBalancePct:     50,
						},
						"ETH": {
							Address:               util.EthAddress.Hex(),
							RebalanceMethods:      []string{"circlecctp"},
							MaintenanceBalancePct: 10,
							InitialBalancePct:     50,
						},
					},
				},
				3: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:               common.HexToAddress("0x3").Hex(),
							RebalanceMethods:      []string{"scroll"},
							MaintenanceBalancePct: 10,
							InitialBalancePct:     50,
						},
						"ETH": {
							Address:               util.EthAddress.Hex(),
							RebalanceMethods:      []string{"scroll"},
							MaintenanceBalancePct: 10,
							InitialBalancePct:     50,
						},
					},
				},
			},
		}
		tokens := map[int]map[common.Address]*inventory.TokenMetadata{
			1: {
				common.HexToAddress("0x1"): {
					Name:    "USDC",
					Balance: big.NewInt(9_000_000),
					ChainID: 1,
					Addr:    common.HexToAddress("0x1"),
				},
				util.EthAddress: {
					Name:    "ETH",
					Balance: big.NewInt(9e18),
					ChainID: 1,
					Addr:    util.EthAddress,
				},
			},
			2: {
				common.HexToAddress("0x2"): {
					Name:    "USDC",
					Balance: big.NewInt(100_000),
					ChainID: 2,
					Addr:    common.HexToAddress("0x2"),
				},
				util.EthAddress: {
					Name:    "ETH",
					Balance: big.NewInt(9e17),
					ChainID: 2,
					Addr:    util.EthAddress,
				},
			},
			3: {
				common.HexToAddress("0x3"): {
					Name:    "USDC",
					Balance: big.NewInt(900_000),
					ChainID: 3,
					Addr:    common.HexToAddress("0x3"),
				},
				util.EthAddress: {
					Name:    "ETH",
					Balance: big.NewInt(1e17),
					ChainID: 3,
					Addr:    util.EthAddress,
				},
			},
		}
		rebalances, err := inventory.GetRebalances(context.Background(), cfg, tokens)
		assert.Nil(t, err)
		expectedRebalance := &inventory.RebalanceData{
			Method:         relconfig.RebalanceMethodCircleCCTP,
			OriginMetadata: tokens[1][common.HexToAddress("0x1")],
			DestMetadata:   tokens[2][common.HexToAddress("0x2")],
			Amount:         big.NewInt(4_450_000),
		}
		assert.Equal(t, rebalances["USDC"], expectedRebalance)
		expectedRebalance = &inventory.RebalanceData{
			Method:         relconfig.RebalanceMethodScroll,
			OriginMetadata: tokens[1][util.EthAddress],
			DestMetadata:   tokens[3][util.EthAddress],
			Amount:         big.NewInt(445e16),
		}
		assert.Equal(t, rebalances["ETH"], expectedRebalance)
	})
}
