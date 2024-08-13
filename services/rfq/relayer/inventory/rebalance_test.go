package inventory_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
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
}
