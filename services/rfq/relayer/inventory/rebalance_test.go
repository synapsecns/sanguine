package inventory_test

import (
	"math/big"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
)

func TestGetOriginMetadatas(t *testing.T) {
	t.Run("IncompatibleRebalanceMethod", func(t *testing.T) {
		cfg := relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{
				1: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:          common.HexToAddress("0x1").Hex(),
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
		}
		originData, destData, method := inventory.GetRebalanceMetadatas(cfg, tokens, "USDC", []relconfig.RebalanceMethod{relconfig.RebalanceMethodCircleCCTP})
		assert.Nil(t, originData)
		assert.Nil(t, destData)
		assert.Equal(t, method, relconfig.RebalanceMethodNone)
	})

	t.Run("SingleMethod", func(t *testing.T) {
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
							RebalanceMethods: []string{"circlecctp"},
						},
					},
				},
				3: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:          common.HexToAddress("0x3").Hex(),
							RebalanceMethods: []string{"circlecctp"},
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
			3: {
				common.HexToAddress("0x3"): {
					Name:    "USDC",
					Balance: big.NewInt(100_000_000),
					ChainID: 3,
					Addr:    common.HexToAddress("0x3"),
				},
			},
		}
		originData, destData, method := inventory.GetRebalanceMetadatas(cfg, tokens, "USDC", []relconfig.RebalanceMethod{relconfig.RebalanceMethodCircleCCTP})
		assert.Equal(t, originData, tokens[3][common.HexToAddress("0x3")])
		assert.Equal(t, destData, tokens[1][common.HexToAddress("0x1")])
		assert.Equal(t, method, relconfig.RebalanceMethodCircleCCTP)
	})

	t.Run("FilterOnMethod", func(t *testing.T) {
		cfg := relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{
				1: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:          common.HexToAddress("0x1").Hex(),
							RebalanceMethods: []string{"circlecctp", "synapsecctp", "scroll"},
						},
					},
				},
				2: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:          common.HexToAddress("0x2").Hex(),
							RebalanceMethods: []string{"circlecctp", "synapsecctp"},
						},
					},
				},
				3: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:          common.HexToAddress("0x3").Hex(),
							RebalanceMethods: []string{"circlecctp"},
						},
					},
				},
				4: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:          common.HexToAddress("0x4").Hex(),
							RebalanceMethods: []string{"scroll"},
						},
					},
				},
			},
		}
		tokens := map[int]map[common.Address]*inventory.TokenMetadata{
			1: {
				common.HexToAddress("0x1"): {
					Name:    "USDC",
					Balance: big.NewInt(1_000_000_000),
					ChainID: 1,
					Addr:    common.HexToAddress("0x1"),
				},
			},
			2: {
				common.HexToAddress("0x2"): {
					Name:    "USDC",
					Balance: big.NewInt(100_000_000),
					ChainID: 2,
					Addr:    common.HexToAddress("0x2"),
				},
			},
			3: {
				common.HexToAddress("0x3"): {
					Name:    "USDC",
					Balance: big.NewInt(10_000_000),
					ChainID: 3,
					Addr:    common.HexToAddress("0x3"),
				},
			},
			4: {
				common.HexToAddress("0x4"): {
					Name:    "USDC",
					Balance: big.NewInt(1_000_000),
					ChainID: 4,
					Addr:    common.HexToAddress("0x4"),
				},
			},
		}
		// multiple methods are available, but only consider circlecctp
		originData, destData, method := inventory.GetRebalanceMetadatas(cfg, tokens, "USDC", []relconfig.RebalanceMethod{relconfig.RebalanceMethodCircleCCTP})
		assert.Equal(t, originData, tokens[1][common.HexToAddress("0x1")])
		assert.Equal(t, destData, tokens[3][common.HexToAddress("0x3")])
		assert.Equal(t, method, relconfig.RebalanceMethodCircleCCTP)
	})

	t.Run("ConsiderAllMethods", func(t *testing.T) {
		cfg := relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{
				1: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:          common.HexToAddress("0x1").Hex(),
							RebalanceMethods: []string{"circlecctp", "synapsecctp", "scroll"},
						},
					},
				},
				2: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:          common.HexToAddress("0x2").Hex(),
							RebalanceMethods: []string{"circlecctp", "synapsecctp"},
						},
					},
				},
				3: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:          common.HexToAddress("0x3").Hex(),
							RebalanceMethods: []string{"circlecctp"},
						},
					},
				},
				4: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							Address:          common.HexToAddress("0x4").Hex(),
							RebalanceMethods: []string{"scroll"},
						},
					},
				},
			},
		}
		tokens := map[int]map[common.Address]*inventory.TokenMetadata{
			1: {
				common.HexToAddress("0x1"): {
					Name:    "USDC",
					Balance: big.NewInt(1_000_000_000),
					ChainID: 1,
					Addr:    common.HexToAddress("0x1"),
				},
			},
			2: {
				common.HexToAddress("0x2"): {
					Name:    "USDC",
					Balance: big.NewInt(100_000_000),
					ChainID: 2,
					Addr:    common.HexToAddress("0x2"),
				},
			},
			3: {
				common.HexToAddress("0x3"): {
					Name:    "USDC",
					Balance: big.NewInt(10_000_000),
					ChainID: 3,
					Addr:    common.HexToAddress("0x3"),
				},
			},
			4: {
				common.HexToAddress("0x4"): {
					Name:    "USDC",
					Balance: big.NewInt(1_000_000),
					ChainID: 4,
					Addr:    common.HexToAddress("0x4"),
				},
			},
		}
		// take the path with the overall largest balance delta
		methods := []relconfig.RebalanceMethod{
			relconfig.RebalanceMethodSynapseCCTP,
			relconfig.RebalanceMethodCircleCCTP,
			relconfig.RebalanceMethodScroll,
		}
		originData, destData, method := inventory.GetRebalanceMetadatas(cfg, tokens, "USDC", methods)
		assert.Equal(t, originData, tokens[1][common.HexToAddress("0x1")])
		assert.Equal(t, destData, tokens[4][common.HexToAddress("0x4")])
		assert.Equal(t, method, relconfig.RebalanceMethodScroll)
	})
}
