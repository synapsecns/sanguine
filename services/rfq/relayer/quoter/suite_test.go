package quoter_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	clientMocks "github.com/synapsecns/sanguine/ethergo/client/mocks"
	fetcherMocks "github.com/synapsecns/sanguine/ethergo/submitter/mocks"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	inventoryMocks "github.com/synapsecns/sanguine/services/rfq/relayer/inventory/mocks"
	"github.com/synapsecns/sanguine/services/rfq/relayer/pricer"
	priceMocks "github.com/synapsecns/sanguine/services/rfq/relayer/pricer/mocks"
	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
)

// Server suite is the main API server test suite.
type QuoterSuite struct {
	*testsuite.TestSuite
	config         relconfig.Config
	manager        *quoter.Manager
	origin         uint32
	destination    uint32
	destinationEth uint32
}

// NewServerSuite creates a end-to-end test suite.
func NewQuoterSuite(tb testing.TB) *QuoterSuite {
	tb.Helper()
	return &QuoterSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (s *QuoterSuite) SetupTest() {
	s.TestSuite.SetupTest()
	// Setup
	s.origin = 42161
	s.destination = 137
	s.destinationEth = 1
	s.config = relconfig.Config{
		Chains: map[int]relconfig.ChainConfig{
			int(s.origin): {
				Tokens: map[string]relconfig.TokenConfig{
					"USDC": {
						Address:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
						PriceUSD: 1,
						Decimals: 6,
					},
					"ETH": {
						Address:  chain.EthAddress.String(),
						PriceUSD: 2000,
						Decimals: 18,
					},
				},
				NativeToken: "ETH",
			},
			int(s.destination): {
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
			int(s.destinationEth): {
				Tokens: map[string]relconfig.TokenConfig{
					"ETH": {
						Address:  chain.EthAddress.String(),
						PriceUSD: 2000,
						Decimals: 18,
					},
				},
				NativeToken: "ETH",
			},
		},
		BaseChainConfig: relconfig.ChainConfig{
			OriginGasEstimate: 500000,
			DestGasEstimate:   1000000,
			MinGasToken:       "1000000000000000", // 1e15
		},
		FeePricer: relconfig.FeePricerConfig{
			GasPriceCacheTTLSeconds:   60,
			TokenPriceCacheTTLSeconds: 60,
		},
		QuotableTokens: map[string][]string{
			"42161-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48": {"137-0x0b2c639c533813f4aa9d7837caf62653d097ff85", "10-0x0b2c639c533813f4aa9d7837caf62653d097ff85"},
			"42161-0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE": {"137-0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE", "1-0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"},
			// "1-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48":     {"42161-0xaf88d065e77c8cc2239327c5edb3a432268e5831", "10-0x0b2c639c533813f4aa9d7837caf62653d097ff85"},
			// "10-0x0b2c639c533813f4aa9d7837caf62653d097ff85":    {"1-0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "42161-0xaf88d065e77c8cc2239327c5edb3a432268e5831"},
		},
	}

	// Build a FeePricer with mock gas price and mock token price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	client := new(clientMocks.EVM)
	priceFetcher := new(priceMocks.CoingeckoPriceFetcher)
	currentHeader := &types.Header{BaseFee: big.NewInt(100_000_000_000)} // 100 gwei
	client.On(testsuite.GetFunctionName(client.HeaderByNumber), mock.Anything, mock.Anything).Return(currentHeader, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, mock.Anything).Twice().Return(client, nil)
	priceFetcher.On(testsuite.GetFunctionName(priceFetcher.GetPrice), mock.Anything, mock.Anything).Return(0., fmt.Errorf("not using mocked price"))
	feePricer := pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	inventoryManager := new(inventoryMocks.Manager)
	inventoryManager.On(testsuite.GetFunctionName(inventoryManager.HasSufficientGas), mock.Anything, mock.Anything, mock.Anything).Return(true, nil)
	mgr, err := quoter.NewQuoterManager(s.config, metrics.NewNullHandler(), inventoryManager, nil, feePricer)
	s.NoError(err)

	var ok bool
	s.manager, ok = mgr.(*quoter.Manager)
	s.True(ok)
}

func (s *QuoterSuite) SetupSuite() {
	s.TestSuite.SetupSuite()
}

// TestConfigSuite runs the integration test suite.
func TestQuoterSuite(t *testing.T) {
	suite.Run(t, NewQuoterSuite(t))
}
