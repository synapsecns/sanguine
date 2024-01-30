package relconfig

import (
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetters(t *testing.T) {
	chainID := 1
	badChainID := 2
	cfgWithBase := Config{
		Chains: map[int]ChainConfig{
			chainID: {
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
			chainID: {
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
		defaultVal, err := cfg.GetOriginGasEstimate(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, defaultChainConfig.OriginGasEstimate)

		baseVal, err := cfgWithBase.GetOriginGasEstimate(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, cfgWithBase.BaseChainConfig.OriginGasEstimate)

		chainVal, err := cfgWithBase.GetOriginGasEstimate(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, cfgWithBase.Chains[chainID].OriginGasEstimate)
	})

	t.Run("GetDestGasEstimate", func(t *testing.T) {
		defaultVal, err := cfg.GetDestGasEstimate(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, defaultChainConfig.DestGasEstimate)

		baseVal, err := cfgWithBase.GetDestGasEstimate(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, cfgWithBase.BaseChainConfig.DestGasEstimate)

		chainVal, err := cfgWithBase.GetDestGasEstimate(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, cfgWithBase.Chains[chainID].DestGasEstimate)
	})

	t.Run("GetL1FeeChainID", func(t *testing.T) {
		defaultVal, err := cfg.GetL1FeeChainID(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, defaultChainConfig.L1FeeChainID)

		baseVal, err := cfgWithBase.GetL1FeeChainID(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, cfgWithBase.BaseChainConfig.L1FeeChainID)

		chainVal, err := cfgWithBase.GetL1FeeChainID(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, cfgWithBase.Chains[chainID].L1FeeChainID)
	})

	t.Run("GetL1FeeOriginGasEstimate", func(t *testing.T) {
		defaultVal, err := cfg.GetL1FeeOriginGasEstimate(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, defaultChainConfig.L1FeeOriginGasEstimate)

		baseVal, err := cfgWithBase.GetL1FeeOriginGasEstimate(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, cfgWithBase.BaseChainConfig.L1FeeOriginGasEstimate)

		chainVal, err := cfgWithBase.GetL1FeeOriginGasEstimate(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, cfgWithBase.Chains[chainID].L1FeeOriginGasEstimate)
	})

	t.Run("GetL1FeeDestGasEstimate", func(t *testing.T) {
		defaultVal, err := cfg.GetL1FeeDestGasEstimate(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, defaultChainConfig.L1FeeDestGasEstimate)

		baseVal, err := cfgWithBase.GetL1FeeDestGasEstimate(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, cfgWithBase.BaseChainConfig.L1FeeDestGasEstimate)

		chainVal, err := cfgWithBase.GetL1FeeDestGasEstimate(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, cfgWithBase.Chains[chainID].L1FeeDestGasEstimate)
	})

	t.Run("GetMinGasToken", func(t *testing.T) {
		defaultVal, err := cfg.GetMinGasToken(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, defaultChainConfig.MinGasToken)

		baseVal, err := cfgWithBase.GetMinGasToken(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, cfgWithBase.BaseChainConfig.MinGasToken)

		chainVal, err := cfgWithBase.GetMinGasToken(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, cfgWithBase.Chains[chainID].MinGasToken)
	})

	t.Run("GetQuotePct", func(t *testing.T) {
		defaultVal, err := cfg.GetQuotePct(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, defaultChainConfig.QuotePct)

		baseVal, err := cfgWithBase.GetQuotePct(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, cfgWithBase.BaseChainConfig.QuotePct)

		chainVal, err := cfgWithBase.GetQuotePct(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, cfgWithBase.Chains[chainID].QuotePct)
	})

	t.Run("GetQuoteOffsetBps", func(t *testing.T) {
		defaultVal, err := cfg.GetQuoteOffsetBps(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, defaultChainConfig.QuoteOffsetBps)

		baseVal, err := cfgWithBase.GetQuoteOffsetBps(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, cfgWithBase.BaseChainConfig.QuoteOffsetBps)

		chainVal, err := cfgWithBase.GetQuoteOffsetBps(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, cfgWithBase.Chains[chainID].QuoteOffsetBps)
	})

	t.Run("GetFixedFeeMultiplier", func(t *testing.T) {
		defaultVal, err := cfg.GetFixedFeeMultiplier(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, defaultChainConfig.FixedFeeMultiplier)

		baseVal, err := cfgWithBase.GetFixedFeeMultiplier(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, cfgWithBase.BaseChainConfig.FixedFeeMultiplier)

		chainVal, err := cfgWithBase.GetFixedFeeMultiplier(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, cfgWithBase.Chains[chainID].FixedFeeMultiplier)
	})
}
