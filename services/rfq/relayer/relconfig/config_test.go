package relconfig

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetters(t *testing.T) {
	cfgWithBase := Config{
		Chains: map[int]ChainConfig{
			1: {
				DeadlineBufferSeconds:  10,
				OriginGasEstimate:      10000,
				DestGasEstimate:        20000,
				L1FeeChainID:           10,
				L1FeeOriginGasEstimate: 30000,
				L1FeeDestGasEstimate:   40000,
				MinGasToken:            "1000",
				QuotePct:               50,
				QuoteOffsetBps:         10,
				FixedFeeMultiplier:     1.1,
			},
		},
		BaseChainConfig: ChainConfig{
			DeadlineBufferSeconds:  11,
			OriginGasEstimate:      10001,
			DestGasEstimate:        20001,
			L1FeeChainID:           11,
			L1FeeOriginGasEstimate: 30001,
			L1FeeDestGasEstimate:   40001,
			MinGasToken:            "1001",
			QuotePct:               51,
			QuoteOffsetBps:         11,
			FixedFeeMultiplier:     1.2,
		},
	}
	cfg := Config{
		Chains: map[int]ChainConfig{
			1: {
				DeadlineBufferSeconds:  10,
				OriginGasEstimate:      10000,
				DestGasEstimate:        20000,
				L1FeeChainID:           10,
				L1FeeOriginGasEstimate: 30000,
				L1FeeDestGasEstimate:   40000,
				MinGasToken:            "1000",
				QuotePct:               50,
				QuoteOffsetBps:         10,
				FixedFeeMultiplier:     1.1,
			},
		},
	}

	t.Run("GetOriginGasEstimate", func(t *testing.T) {
		defaultVal, err := cfg.GetOriginGasEstimate(2)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, defaultChainConfig.OriginGasEstimate)

		baseVal, err := cfgWithBase.GetOriginGasEstimate(2)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, cfgWithBase.BaseChainConfig.OriginGasEstimate)

		chainVal, err := cfgWithBase.GetOriginGasEstimate(1)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, cfgWithBase.Chains[1].OriginGasEstimate)
	})
}
