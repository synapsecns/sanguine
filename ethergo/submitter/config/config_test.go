package config_test

import (
	"github.com/synapsecns/sanguine/ethergo/submitter/config"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
)

func TestGetters(t *testing.T) {
	cfg := config.Config{
		GlobalConfig: config.ChainConfig{
			MaxBatchSize: 5,
			DoNotBatch:   false,
			MaxGasPrice:  big.NewInt(250 * params.GWei),
		},
		Chains: map[int]config.ChainConfig{
			1: {
				MaxBatchSize: 8,
				DoNotBatch:   true,
				MaxGasPrice:  big.NewInt(300 * params.GWei),
			},
			2: {
				MaxBatchSize: 0, // Should use global config value
			},
			3: {
				DoNotBatch: false, // Should use global config value
			},
		},
	}

	t.Run("GetMaxBatchSize", func(t *testing.T) {
		assert.Equal(t, 8, cfg.GetMaxBatchSize(1))
		assert.Equal(t, 5, cfg.GetMaxBatchSize(2))
		assert.Equal(t, 5, cfg.GetMaxBatchSize(3))
		assert.Equal(t, 5, cfg.GetMaxBatchSize(4)) // Nonexistent chain, should use global config value
	})

	t.Run("GetBatch", func(t *testing.T) {
		assert.Equal(t, false, cfg.GetBatch(1))
		assert.Equal(t, true, cfg.GetBatch(2))
		assert.Equal(t, true, cfg.GetBatch(3))
		assert.Equal(t, true, cfg.GetBatch(4)) // Nonexistent chain, should use global config value
	})

	t.Run("GetMaxGasPrice", func(t *testing.T) {
		assert.Equal(t, big.NewInt(300*params.GWei), cfg.GetMaxGasPrice(1))
		assert.Equal(t, big.NewInt(250*params.GWei), cfg.GetMaxGasPrice(2))
		assert.Equal(t, big.NewInt(250*params.GWei), cfg.GetMaxGasPrice(3))
		assert.Equal(t, big.NewInt(250*params.GWei), cfg.GetMaxGasPrice(4)) // Nonexistent chain, should use global config value
	})
}
