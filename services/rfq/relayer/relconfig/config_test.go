package relconfig_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
)

//nolint:maintidx
func TestChainGetters(t *testing.T) {
	chainID := 1
	badChainID := 2
	usdcAddr := "0x123"
	cfgWithBase := relconfig.Config{
		Chains: map[int]relconfig.ChainConfig{
			chainID: {
				RFQAddress:              "0x123",
				SynapseCCTPAddress:      "0x456",
				TokenMessengerAddress:   "0x789",
				Confirmations:           1,
				NativeToken:             "MATIC",
				DeadlineBufferSeconds:   10,
				OriginGasEstimate:       10000,
				DestGasEstimate:         20000,
				L1FeeChainID:            10,
				L1FeeOriginGasEstimate:  30000,
				L1FeeDestGasEstimate:    40000,
				MinGasToken:             "1000",
				QuotePct:                relconfig.NewFloat64Pointer(0),
				QuoteWidthBps:           10,
				QuoteFixedFeeMultiplier: relconfig.NewFloat64Pointer(1.1),
			},
		},
		BaseChainConfig: relconfig.ChainConfig{
			RFQAddress:              "0x1234",
			SynapseCCTPAddress:      "0x456",
			TokenMessengerAddress:   "0x789",
			Confirmations:           2,
			NativeToken:             "ARB",
			DeadlineBufferSeconds:   11,
			OriginGasEstimate:       10001,
			DestGasEstimate:         20001,
			L1FeeChainID:            11,
			L1FeeOriginGasEstimate:  30001,
			L1FeeDestGasEstimate:    40001,
			MinGasToken:             "1001",
			QuotePct:                relconfig.NewFloat64Pointer(51),
			QuoteWidthBps:           11,
			QuoteFixedFeeMultiplier: relconfig.NewFloat64Pointer(1.2),
		},
	}
	cfg := relconfig.Config{
		Chains: map[int]relconfig.ChainConfig{
			chainID: {
				RFQAddress:              "0x123",
				SynapseCCTPAddress:      "0x456",
				TokenMessengerAddress:   "0x789",
				Confirmations:           1,
				NativeToken:             "MATIC",
				DeadlineBufferSeconds:   10,
				OriginGasEstimate:       10000,
				DestGasEstimate:         20000,
				L1FeeChainID:            10,
				L1FeeOriginGasEstimate:  30000,
				L1FeeDestGasEstimate:    40000,
				MinGasToken:             "1000",
				QuotePct:                relconfig.NewFloat64Pointer(50),
				QuoteWidthBps:           10,
				QuoteFixedFeeMultiplier: relconfig.NewFloat64Pointer(1.1),
				Tokens: map[string]relconfig.TokenConfig{
					"USDC": {
						Address:            usdcAddr,
						Decimals:           6,
						MaxRebalanceAmount: "1000",
					},
				},
			},
		},
	}

	t.Run("GetRFQAddress", func(t *testing.T) {
		defaultVal, err := cfg.GetRFQAddress(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, relconfig.DefaultChainConfig.RFQAddress)

		baseVal, err := cfgWithBase.GetRFQAddress(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, cfgWithBase.BaseChainConfig.RFQAddress)

		chainVal, err := cfgWithBase.GetRFQAddress(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, cfgWithBase.Chains[chainID].RFQAddress)
	})

	t.Run("GetSynapseCCTPAddress", func(t *testing.T) {
		defaultVal, err := cfg.GetSynapseCCTPAddress(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, relconfig.DefaultChainConfig.SynapseCCTPAddress)

		baseVal, err := cfgWithBase.GetSynapseCCTPAddress(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, cfgWithBase.BaseChainConfig.SynapseCCTPAddress)

		chainVal, err := cfgWithBase.GetSynapseCCTPAddress(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, cfgWithBase.Chains[chainID].SynapseCCTPAddress)
	})

	t.Run("GetTokenMessengerAddress", func(t *testing.T) {
		defaultVal, err := cfg.GetTokenMessengerAddress(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, relconfig.DefaultChainConfig.TokenMessengerAddress)

		baseVal, err := cfgWithBase.GetTokenMessengerAddress(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, cfgWithBase.BaseChainConfig.TokenMessengerAddress)

		chainVal, err := cfgWithBase.GetTokenMessengerAddress(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, cfgWithBase.Chains[chainID].TokenMessengerAddress)
	})

	t.Run("GetConfirmations", func(t *testing.T) {
		defaultVal, err := cfg.GetConfirmations(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, relconfig.DefaultChainConfig.Confirmations)

		baseVal, err := cfgWithBase.GetConfirmations(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, cfgWithBase.BaseChainConfig.Confirmations)

		chainVal, err := cfgWithBase.GetConfirmations(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, cfgWithBase.Chains[chainID].Confirmations)
	})

	t.Run("GetNativeToken", func(t *testing.T) {
		defaultVal, err := cfg.GetNativeToken(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, relconfig.DefaultChainConfig.NativeToken)

		baseVal, err := cfgWithBase.GetNativeToken(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, cfgWithBase.BaseChainConfig.NativeToken)

		chainVal, err := cfgWithBase.GetNativeToken(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, cfgWithBase.Chains[chainID].NativeToken)
	})

	t.Run("GetDeadlineBuffer", func(t *testing.T) {
		defaultVal, err := cfg.GetDeadlineBuffer(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, time.Duration(relconfig.DefaultChainConfig.DeadlineBufferSeconds)*time.Second)

		baseVal, err := cfgWithBase.GetDeadlineBuffer(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, time.Duration(cfgWithBase.BaseChainConfig.DeadlineBufferSeconds)*time.Second)

		chainVal, err := cfgWithBase.GetDeadlineBuffer(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, time.Duration(cfgWithBase.Chains[chainID].DeadlineBufferSeconds)*time.Second)
	})

	t.Run("GetOriginGasEstimate", func(t *testing.T) {
		defaultVal, err := cfg.GetOriginGasEstimate(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, relconfig.DefaultChainConfig.OriginGasEstimate)

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
		assert.Equal(t, defaultVal, relconfig.DefaultChainConfig.DestGasEstimate)

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
		assert.Equal(t, defaultVal, relconfig.DefaultChainConfig.L1FeeChainID)

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
		assert.Equal(t, defaultVal, relconfig.DefaultChainConfig.L1FeeOriginGasEstimate)

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
		assert.Equal(t, defaultVal, relconfig.DefaultChainConfig.L1FeeDestGasEstimate)

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
		assert.Equal(t, defaultVal.String(), relconfig.DefaultChainConfig.MinGasToken)

		baseVal, err := cfgWithBase.GetMinGasToken(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal.String(), cfgWithBase.BaseChainConfig.MinGasToken)

		chainVal, err := cfgWithBase.GetMinGasToken(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal.String(), cfgWithBase.Chains[chainID].MinGasToken)
	})

	t.Run("GetQuotePct", func(t *testing.T) {
		defaultVal, err := cfg.GetQuotePct(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, 100.)

		baseVal, err := cfgWithBase.GetQuotePct(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, 51.)

		chainVal, err := cfgWithBase.GetQuotePct(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, 0.)
	})

	t.Run("GetQuoteWidthBps", func(t *testing.T) {
		defaultVal, err := cfg.GetQuoteWidthBps(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, relconfig.DefaultChainConfig.QuoteWidthBps)

		baseVal, err := cfgWithBase.GetQuoteWidthBps(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, cfgWithBase.BaseChainConfig.QuoteWidthBps)

		chainVal, err := cfgWithBase.GetQuoteWidthBps(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, cfgWithBase.Chains[chainID].QuoteWidthBps)
	})

	t.Run("GetQuoteFixedFeeMultiplier", func(t *testing.T) {
		defaultVal, err := cfg.GetQuoteFixedFeeMultiplier(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, defaultVal, *relconfig.DefaultChainConfig.QuoteFixedFeeMultiplier)

		baseVal, err := cfgWithBase.GetQuoteFixedFeeMultiplier(badChainID)
		assert.NoError(t, err)
		assert.Equal(t, baseVal, *cfgWithBase.BaseChainConfig.QuoteFixedFeeMultiplier)

		chainVal, err := cfgWithBase.GetQuoteFixedFeeMultiplier(chainID)
		assert.NoError(t, err)
		assert.Equal(t, chainVal, *cfgWithBase.Chains[chainID].QuoteFixedFeeMultiplier)
	})

	t.Run("GetMaxRebalanceAmount", func(t *testing.T) {
		defaultVal := cfg.GetMaxRebalanceAmount(badChainID, common.HexToAddress(usdcAddr))
		assert.Equal(t, defaultVal.String(), abi.MaxInt256.String())

		chainVal := cfg.GetMaxRebalanceAmount(chainID, common.HexToAddress(usdcAddr))
		assert.Equal(t, chainVal.String(), "1000000000")
	})
}

func TestGetQuoteOffset(t *testing.T) {
	chainID := 1
	usdcAddr := "0x123"
	cfg := relconfig.Config{
		Chains: map[int]relconfig.ChainConfig{
			chainID: {
				RFQAddress:              "0x123",
				SynapseCCTPAddress:      "0x456",
				TokenMessengerAddress:   "0x789",
				Confirmations:           1,
				NativeToken:             "MATIC",
				DeadlineBufferSeconds:   10,
				OriginGasEstimate:       10000,
				DestGasEstimate:         20000,
				L1FeeChainID:            10,
				L1FeeOriginGasEstimate:  30000,
				L1FeeDestGasEstimate:    40000,
				MinGasToken:             "1000",
				QuotePct:                relconfig.NewFloat64Pointer(50),
				QuoteWidthBps:           10,
				QuoteFixedFeeMultiplier: relconfig.NewFloat64Pointer(1.1),
				Tokens: map[string]relconfig.TokenConfig{
					"USDC": {
						Address:            usdcAddr,
						Decimals:           6,
						MaxRebalanceAmount: "1000",
						QuoteOffsetBps:     100,
					},
				},
			},
		},
	}

	t.Run("GetQuoteOffsetForOrigin", func(t *testing.T) {
		val, err := cfg.GetQuoteOffsetBps(chainID, "USDC", true)
		assert.NoError(t, err)
		assert.Equal(t, -100., val)
	})

	t.Run("GetQuoteOffsetForDest", func(t *testing.T) {
		val, err := cfg.GetQuoteOffsetBps(chainID, "USDC", false)
		assert.NoError(t, err)
		assert.Equal(t, 100., val)
	})
}

func TestValidation(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		cfg := relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{
				1: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							InitialBalancePct:     50,
							MaintenanceBalancePct: 25,
							RebalanceMethod:       "synapsecctp",
						},
					},
				},
				2: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							InitialBalancePct:     50,
							MaintenanceBalancePct: 25,
							RebalanceMethod:       "synapsecctp",
						},
					},
				},
			},
		}
		err := cfg.Validate()
		assert.Nil(t, err)
	})

	t.Run("InvalidInitialPct", func(t *testing.T) {
		cfg := relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{
				1: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							InitialBalancePct:     51,
							MaintenanceBalancePct: 50,
							RebalanceMethod:       "synapsecctp",
						},
					},
				},
				2: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							InitialBalancePct:     50,
							MaintenanceBalancePct: 50,
							RebalanceMethod:       "synapsecctp",
						},
					},
				},
			},
		}
		err := cfg.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, "total initial percent does not total 100 for USDC: 101.000000", err.Error())
	})

	t.Run("InvalidMaintenancePct", func(t *testing.T) {
		cfg := relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{
				1: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							InitialBalancePct:     50,
							MaintenanceBalancePct: 50,
							RebalanceMethod:       "synapsecctp",
						},
					},
				},
				2: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							InitialBalancePct:     50,
							MaintenanceBalancePct: 50.1,
							RebalanceMethod:       "synapsecctp",
						},
					},
				},
			},
		}
		err := cfg.Validate()
		assert.NotNil(t, err)
		assert.Equal(t, "total maintenance percent exceeds 100 for USDC: 100.100000", err.Error())
	})

	t.Run("ValidWithNoRebalanceMethod", func(t *testing.T) {
		cfg := relconfig.Config{
			Chains: map[int]relconfig.ChainConfig{
				1: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							InitialBalancePct:     50,
							MaintenanceBalancePct: 20,
						},
					},
				},
				2: {
					Tokens: map[string]relconfig.TokenConfig{
						"USDC": {
							InitialBalancePct:     50,
							MaintenanceBalancePct: 20,
						},
					},
				},
			},
		}
		err := cfg.Validate()
		assert.Nil(t, err)
	})
}

func TestDecodeTokenID(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		wantChain int
		wantAddr  common.Address
		wantErr   bool
	}{
		{
			name:      "valid token ID",
			id:        "1-0x1234567890abcdef1234567890abcdef12345678",
			wantChain: 1,
			wantAddr:  common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678"),
			wantErr:   false,
		},
		{
			name:    "invalid token ID format",
			id:      "1_0x1234567890abcdef1234567890abcdef12345678",
			wantErr: true,
		},
		{
			name:    "invalid chain ID",
			id:      "x-0x1234567890abcdef1234567890abcdef12345678",
			wantErr: true,
		},
		{
			name:    "invalid address",
			id:      "1-0x12345",
			wantErr: true,
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			gotChain, gotAddr, err := relconfig.DecodeTokenID(tt.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantChain, gotChain)
				assert.Equal(t, tt.wantAddr, gotAddr)
			}
		})
	}
}
