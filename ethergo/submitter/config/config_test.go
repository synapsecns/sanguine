package config_test

import (
	"math/big"
	"testing"

	"github.com/synapsecns/sanguine/ethergo/submitter/config"
	"gopkg.in/yaml.v2"

	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"
)

func TestGetters(t *testing.T) {
	cfg := config.Config{
		ChainConfig: config.ChainConfig{
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

func TestGlobalConfig(t *testing.T) {
	cfgStr := `max_gas_price: 250000000000
bump_interval_seconds: 60
gas_bump_percentage: 10
gas_estimate: 1000
is_l2: true
dynamic_gas_estimate: true
supports_eip_1559: true`
	var cfg config.Config
	err := yaml.Unmarshal([]byte(cfgStr), &cfg)
	assert.NoError(t, err)
	assert.Equal(t, big.NewInt(250000000000), cfg.MaxGasPrice)
	assert.Equal(t, 60, cfg.BumpIntervalSeconds)
	assert.Equal(t, 10, cfg.GasBumpPercentage)
	assert.Equal(t, uint64(1000), cfg.GasEstimate)
	assert.Equal(t, true, cfg.IsL2(0))
	assert.Equal(t, true, cfg.DynamicGasEstimate)
	assert.Equal(t, true, cfg.SupportsEIP1559(0))
}
