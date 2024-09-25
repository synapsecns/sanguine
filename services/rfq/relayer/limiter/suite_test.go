package limiter_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/services/rfq/relayer/limiter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/util"
)

// Server suite is the main API server test suite.
type LimiterSuite struct {
	*testsuite.TestSuite
	cfg     relconfig.Config
	metrics metrics.Handler
	limiter limiter.Limiter
}

// NewServerSuite creates a end-to-end test suite.
func NewLimiterSuite(tb testing.TB) *LimiterSuite {
	tb.Helper()
	return &LimiterSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (s *LimiterSuite) SetupTest() {
	s.TestSuite.SetupTest()
	// Setup
	s.cfg = relconfig.Config{
		Chains: map[int]relconfig.ChainConfig{
			10: {
				Tokens: map[string]relconfig.TokenConfig{
					"USDC": {
						Address:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
						PriceUSD: 1,
						Decimals: 6,
					},
					"ETH": {
						Address:  util.EthAddress.String(),
						PriceUSD: 2000,
						Decimals: 18,
					},
				},
				RPCConfirmations:   1,
				LimitConfirmations: 1,
				VolumeLimit:        1000, // 1k usd
				RFQAddress:         "0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E",
			},
			81457: {
				Tokens: map[string]relconfig.TokenConfig{
					"USDC": {
						Address:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
						PriceUSD: 1,
						Decimals: 6,
					},
					"ETH": {
						Address:  util.EthAddress.String(),
						PriceUSD: 2000,
						Decimals: 18,
					},
				},
				RPCConfirmations:   1,
				LimitConfirmations: 1,
				VolumeLimit:        10000, // 10k usd
				RFQAddress:         "0x5523D3c98809DdDB82C686E152F5C58B1B0fB59E",
			},
		},
	}
	s.metrics = metrics.NewNullHandler()
}

func (s *LimiterSuite) SetupSuite() {
	s.TestSuite.SetupSuite()
}

// TestConfigSuite runs the integration test suite.
func TestLimiterSuite(t *testing.T) {
	suite.Run(t, NewLimiterSuite(t))
}
