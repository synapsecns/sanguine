package pricer_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
)

// Pricer suite is the main API server test suite.
type PricerSuite struct {
	*testsuite.TestSuite
	feePricerConfig relconfig.FeePricerConfig
	chainConfigs    map[int]relconfig.ChainConfig
}

// NewPricerSuite creates a end-to-end test suite.
func NewPricerSuite(tb testing.TB) *PricerSuite {
	tb.Helper()
	return &PricerSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (c *PricerSuite) SetupTest() {
	c.TestSuite.SetupTest()
	// Setup
	c.feePricerConfig = relconfig.FeePricerConfig{
		GasPriceCacheTTL:       60,
		TokenPriceCacheTTL:     60,
		OriginGasEstimate:      100000,
		DestinationGasEstimate: 100000,
	}
	c.chainConfigs = map[int]relconfig.ChainConfig{
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
}

func (c *PricerSuite) SetupSuite() {
	c.TestSuite.SetupSuite()
}

// TestPricerSuite runs the test suite.
func TestPricerSuite(t *testing.T) {
	suite.Run(t, NewPricerSuite(t))
}
