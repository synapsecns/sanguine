package quoter_test

import (
	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
)

func (s *QuoterSuite) TestGetOriginFee() {
	feePricerConfig := relconfig.FeePricerConfig{
		GasPriceCacheTTL:       60,
		TokenPriceCacheTTL:     60,
		OriginGasEstimate:      100000,
		DestinationGasEstimate: 100000,
	}
	chainConfigs := map[int]relconfig.ChainConfig{
		42161: relconfig.ChainConfig{
			Tokens: map[string]relconfig.TokenConfig{
				"USDC": relconfig.TokenConfig{
					Address:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					PriceUSD: 1,
				},
				"ETH": relconfig.TokenConfig{
					Address:  "",
					PriceUSD: 2000,
				},
			},
		},
	}
	feePricer := quoter.NewFeePricer(feePricerConfig, chainConfigs, "http://localhost:8545", nil)
}
