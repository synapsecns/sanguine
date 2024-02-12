package pricer_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
)

// Pricer suite is the main API server test suite.
type PricerSuite struct {
	*testsuite.TestSuite
	config      relconfig.Config
	origin      uint32
	destination uint32
	l1ChainID   uint32
	signer      signer.Signer
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
	c.origin = 42161
	c.destination = 137
	c.l1ChainID = 1
	wall, err := wallet.FromRandom()
	c.Require().NoError(err)
	c.signer = localsigner.NewSigner(wall.PrivateKey())
	c.config = relconfig.Config{
		BaseChainConfig: relconfig.ChainConfig{
			OriginGasEstimate: 500000,
			DestGasEstimate:   1000000,
		},
		FeePricer: relconfig.FeePricerConfig{
			GasPriceCacheTTLSeconds:   60,
			TokenPriceCacheTTLSeconds: 60,
		},
		Chains: map[int]relconfig.ChainConfig{
			int(c.origin): {
				Tokens: map[string]relconfig.TokenConfig{
					"USDC": {
						Address:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
						PriceUSD: 1,
						Decimals: 6,
					},
					"ETH": {
						Address:  "",
						PriceUSD: 2000,
						Decimals: 18,
					},
				},
				NativeToken:        "ETH",
				DynamicGasEstimate: true,
			},
			int(c.destination): {
				Tokens: map[string]relconfig.TokenConfig{
					"USDC": {
						Address:  "0x0b2c639c533813f4aa9d7837caf62653d097ff85",
						PriceUSD: 1,
						Decimals: 6,
					},
					"MATIC": {
						Address:  "",
						PriceUSD: 0.5,
						Decimals: 18,
					},
				},
				NativeToken: "MATIC",
			},
			int(c.l1ChainID): {
				Tokens: map[string]relconfig.TokenConfig{
					"ETH": {
						Address:  "",
						PriceUSD: 2000,
						Decimals: 18,
					},
				},
				NativeToken: "ETH",
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
