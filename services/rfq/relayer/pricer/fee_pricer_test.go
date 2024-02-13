package pricer_test

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	clientMocks "github.com/synapsecns/sanguine/ethergo/client/mocks"
	fetcherMocks "github.com/synapsecns/sanguine/ethergo/submitter/mocks"
	"github.com/synapsecns/sanguine/services/rfq/relayer/pricer"
	priceMocks "github.com/synapsecns/sanguine/services/rfq/relayer/pricer/mocks"
)

var defaultPrices = map[string]float64{"ETH": 2000., "USDC": 1., "MATIC": 0.5}

func getPriceFetcher(prices map[string]float64) *priceMocks.CoingeckoPriceFetcher {
	priceFetcher := new(priceMocks.CoingeckoPriceFetcher)
	for token, price := range defaultPrices {
		if prices != nil {
			providedPrice, ok := prices[token]
			if ok {
				price = providedPrice
			}
		}
		priceFetcher.On(testsuite.GetFunctionName(priceFetcher.GetPrice), mock.Anything, token).Return(price, nil)
	}
	return priceFetcher
}

func (s *PricerSuite) TestGetOriginFee() {
	// Build a new FeePricer with a mocked client for fetching gas price and token price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	client := new(clientMocks.EVM)
	currentHeader := &types.Header{BaseFee: big.NewInt(100_000_000_000)} // 100 gwei
	client.On(testsuite.GetFunctionName(client.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(currentHeader, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, mock.Anything).Twice().Return(client, nil)
	priceFetcher := getPriceFetcher(nil)
	feePricer := pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	// Calculate the origin fee.
	fee, err := feePricer.GetOriginFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)

	/*
		The expected fee should be:
		fee_eth: gas_price * gas_estimate / native_decimals_factor
		fee_usd: fee_eth * eth_price_usd
		fee_usdc: fee_usd * usdc_price_usd
		fee_usdc_decimals: fee_usdc * usdc_decimals_factor
		fee_usdc_decimals = (((gas_price * gas_estimate / native_decimals_factor) * eth_price_usd) * usdc_price_usd) * usdc_decimals_factor
		So, with our numbers:
		fee_denom = (((100e9 * 500000 / 1e18) * 1000) * 1) * 1e6 = 100_000_000
	*/

	expectedFee := big.NewInt(100_000_000) // 100 usd
	s.Equal(expectedFee, fee)

	// Ensure that the fee has been cached.
	client.On(testsuite.GetFunctionName(client.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(nil, fmt.Errorf("could not fetch header"))
	fee, err = feePricer.GetOriginFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)
	s.Equal(expectedFee, fee)
}

func (s *PricerSuite) TestGetOriginFeeWithOverrides() {
	// Set chain fee overrides.
	l1ChainID := uint32(1)
	s.config.BaseChainConfig.OriginGasEstimate = 5_000_000
	s.config.BaseChainConfig.DestGasEstimate = 10_000_000
	s.config.BaseChainConfig.L1FeeChainID = l1ChainID
	s.config.BaseChainConfig.L1FeeOriginGasEstimate = 1_000_000
	s.config.BaseChainConfig.L1FeeDestGasEstimate = 2_000_000

	// Build a new FeePricer with a mocked client for fetching gas price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	client := new(clientMocks.EVM)
	currentHeader := &types.Header{BaseFee: big.NewInt(100_000_000_000)} // 100 gwei
	client.On(testsuite.GetFunctionName(client.HeaderByNumber), mock.Anything, mock.Anything).Return(currentHeader, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, mock.Anything).Return(client, nil)
	priceFetcher := getPriceFetcher(map[string]float64{"ETH": 1000})
	feePricer := pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	// Calculate the origin fee.
	fee, err := feePricer.GetOriginFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)

	/*
		The expected fee should be:
		fee_eth: gas_price * gas_estimate / native_decimals_factor
		fee_usd: fee_eth * eth_price_usd
		fee_usdc: fee_usd * usdc_price_usd
		fee_usdc_decimals: fee_usdc * usdc_decimals_factor
		fee_usdc_decimals = (((gas_price * gas_estimate / native_decimals_factor) * eth_price_usd) * usdc_price_usd) * usdc_decimals_factor
		So, with our numbers:
		fee_denom = (((100e9 * 5000000 / 1e18) * 2000) * 1) * 1e6 = 1_000_000_000

		Then, add the l1 fee component:
		fee_denom = (((100e9 * 1000000 / 1e18) * 2000) * 1) * 1e6 = 200_000_000

		So, the total is: 600_000_000
	*/

	expectedFee := big.NewInt(600_000_000) // 600 usd
	s.Equal(expectedFee, fee)

	// Ensure that the fee has been cached.
	client.On(testsuite.GetFunctionName(client.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(nil, fmt.Errorf("could not fetch header"))
	fee, err = feePricer.GetOriginFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)
	s.Equal(expectedFee, fee)
}

func (s *PricerSuite) TestGetDestinationFee() {
	// Build a new FeePricer with a mocked client for fetching gas price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	client := new(clientMocks.EVM)
	currentHeader := &types.Header{BaseFee: big.NewInt(500_000_000_000)} // 500 gwei
	client.On(testsuite.GetFunctionName(client.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(currentHeader, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, mock.Anything).Twice().Return(client, nil)
	priceFetcher := getPriceFetcher(nil)
	feePricer := pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	// Calculate the destination fee.
	fee, err := feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)

	/*
		The expected fee should be:
		fee_matic: gas_price * gas_estimate / native_decimals_factor
		fee_usd: fee_matic * matic_price_usd
		fee_usdc: fee_usd * usdc_price_usd
		fee_usdc_decimals: fee_usdc * usdc_decimals_factor
		fee_usdc_decimals = (((gas_price * gas_estimate / native_decimals_factor) * matic_price_usd) * usdc_price_usd) * usdc_decimals_factor
		So, with our numbers:
		fee_denom = (((500e9 * 1000000 / 1e18) * 0.5) * 1) * 1e6 = 250_000
	*/

	expectedFee := big.NewInt(250_000) // 0.25 usd
	s.Equal(expectedFee, fee)

	// Ensure that the fee has been cached.
	client.On(testsuite.GetFunctionName(client.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(nil, fmt.Errorf("could not fetch header"))
	fee, err = feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)
	s.Equal(expectedFee, fee)
}

func (s *PricerSuite) TestGetDestinationFeeWithOverrides() {
	// Set chain fee overrides.
	l1ChainID := uint32(1)
	s.config.BaseChainConfig.OriginGasEstimate = 5_000_000
	s.config.BaseChainConfig.DestGasEstimate = 10_000_000
	s.config.BaseChainConfig.L1FeeChainID = l1ChainID
	s.config.BaseChainConfig.L1FeeOriginGasEstimate = 1_000_000
	s.config.BaseChainConfig.L1FeeDestGasEstimate = 2_000_000

	// Build a new FeePricer with a mocked client for fetching gas price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	client := new(clientMocks.EVM)
	currentHeader := &types.Header{BaseFee: big.NewInt(500_000_000_000)} // 500 gwei
	client.On(testsuite.GetFunctionName(client.HeaderByNumber), mock.Anything, mock.Anything).Return(currentHeader, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, mock.Anything).Return(client, nil)
	priceFetcher := getPriceFetcher(nil)
	feePricer := pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	// Calculate the destination fee.
	fee, err := feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)

	/*
		The expected fee should be:
		fee_matic: gas_price * gas_estimate / native_decimals_factor
		fee_usd: fee_matic * matic_price_usd
		fee_usdc: fee_usd * usdc_price_usd
		fee_usdc_decimals: fee_usdc * usdc_decimals_factor
		fee_usdc_decimals = (((gas_price * gas_estimate / native_decimals_factor) * matic_price_usd) * usdc_price_usd) * usdc_decimals_factor
		So, with our numbers:
		fee_denom = (((500e9 * 10_000_000 / 1e18) * 0.5) * 1) * 1e6 = 2_500_000

		Then, add the l1 fee component:
		fee_denom = (((500e9 * 2_000_000 / 1e18) * 2000) * 1) * 1e6 = 2_000_000_000

		So, the total is: 2_002_500_000
	*/

	expectedFee := big.NewInt(2_002_500_000) // 2002.5 usd
	s.Equal(expectedFee, fee)

	// Ensure that the fee has been cached.
	client.On(testsuite.GetFunctionName(client.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(nil, fmt.Errorf("could not fetch header"))
	fee, err = feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)
	s.Equal(expectedFee, fee)
}

func (s *PricerSuite) TestGetTotalFee() {
	// Build a new FeePricer with a mocked client for fetching gas price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	clientOrigin := new(clientMocks.EVM)
	clientDestination := new(clientMocks.EVM)
	headerOrigin := &types.Header{BaseFee: big.NewInt(100_000_000_000)}      // 100 gwei
	headerDestination := &types.Header{BaseFee: big.NewInt(500_000_000_000)} // 500 gwei
	clientOrigin.On(testsuite.GetFunctionName(clientOrigin.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(headerOrigin, nil)
	clientDestination.On(testsuite.GetFunctionName(clientDestination.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(headerDestination, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, big.NewInt(int64(s.origin))).Once().Return(clientOrigin, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, big.NewInt(int64(s.destination))).Once().Return(clientDestination, nil)
	priceFetcher := getPriceFetcher(nil)
	feePricer := pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	// Calculate the total fee.
	fee, err := feePricer.GetTotalFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)

	// The expected fee should be the sum of the Origin and Destination fees, i.e. 100_250_000.
	expectedFee := big.NewInt(100_250_000) // 100.25 usd
	s.Equal(expectedFee, fee)
}

func (s *PricerSuite) TestGetGasPrice() {
	// Build a new FeePricer with a mocked client for fetching gas price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	client := new(clientMocks.EVM)
	currentHeader := &types.Header{BaseFee: big.NewInt(100_000_000_000)} // 100 gwei
	client.On(testsuite.GetFunctionName(client.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(currentHeader, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, mock.Anything).Twice().Return(client, nil)
	// Override the gas price cache TTL to 1 second.
	s.config.FeePricer.GasPriceCacheTTLSeconds = 1
	priceFetcher := getPriceFetcher(nil)
	feePricer := pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	// Fetch the mocked gas price.
	gasPrice, err := feePricer.GetGasPrice(s.GetTestContext(), s.origin)
	s.NoError(err)
	expectedGasPrice := big.NewInt(100_000_000_000) // 100 gwei
	s.Equal(expectedGasPrice, gasPrice)

	// Check that the mocked gas price is cached.
	gasPrice, err = feePricer.GetGasPrice(s.GetTestContext(), s.origin)
	s.NoError(err)
	s.Equal(expectedGasPrice, gasPrice)

	// Check that the mocked gas price is eventually evicted from the cache,
	// and an updated gas price is fetched.
	currentHeader = &types.Header{BaseFee: big.NewInt(200_000_000_000)} // 200 gwei
	client.On(testsuite.GetFunctionName(client.HeaderByNumber), mock.Anything, mock.Anything).Return(currentHeader, nil)
	s.Eventually(func() bool {
		gasPrice, err = feePricer.GetGasPrice(s.GetTestContext(), s.origin)
		s.NoError(err)
		expectedGasPrice = big.NewInt(200_000_000_000) // 200 gwei
		return expectedGasPrice.String() == gasPrice.String()
	})
}

func (s *PricerSuite) TestGetTotalFeeWithMultiplier() {
	// Override the fixed fee multiplier to greater than 1.
	s.config.BaseChainConfig.FixedFeeMultiplier = 2

	// Build a new FeePricer with a mocked client for fetching gas price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	clientOrigin := new(clientMocks.EVM)
	clientDestination := new(clientMocks.EVM)
	headerOrigin := &types.Header{BaseFee: big.NewInt(100_000_000_000)}      // 100 gwei
	headerDestination := &types.Header{BaseFee: big.NewInt(500_000_000_000)} // 500 gwei
	clientOrigin.On(testsuite.GetFunctionName(clientOrigin.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(headerOrigin, nil)
	clientDestination.On(testsuite.GetFunctionName(clientDestination.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(headerDestination, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, big.NewInt(int64(s.origin))).Once().Return(clientOrigin, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, big.NewInt(int64(s.destination))).Once().Return(clientDestination, nil)
	priceFetcher := getPriceFetcher(nil)
	feePricer := pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	// Calculate the total fee.
	fee, err := feePricer.GetTotalFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)

	// The expected fee should be the sum of the Origin and Destination fees, i.e. 200_500_000.
	expectedFee := big.NewInt(200_500_000) // 200.50 usd
	s.Equal(expectedFee, fee)

	// Override the fixed fee multiplier to less than 1; should default to 1.
	s.config.BaseChainConfig.FixedFeeMultiplier = -1

	// Build a new FeePricer with a mocked client for fetching gas price.
	clientOrigin.On(testsuite.GetFunctionName(clientOrigin.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(headerOrigin, nil)
	clientDestination.On(testsuite.GetFunctionName(clientDestination.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(headerDestination, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, big.NewInt(int64(s.origin))).Once().Return(clientOrigin, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, big.NewInt(int64(s.destination))).Once().Return(clientDestination, nil)
	feePricer = pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	// Calculate the total fee.
	fee, err = feePricer.GetTotalFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)

	// The expected fee should be the sum of the Origin and Destination fees, i.e. 100_250_000.
	expectedFee = big.NewInt(100_250_000) // 100.25 usd
	s.Equal(expectedFee, fee)

	// Reset the fixed fee multiplier to zero; should default to 1
	s.config.BaseChainConfig.FixedFeeMultiplier = 0

	// Build a new FeePricer with a mocked client for fetching gas price.
	clientOrigin.On(testsuite.GetFunctionName(clientOrigin.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(headerOrigin, nil)
	clientDestination.On(testsuite.GetFunctionName(clientDestination.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(headerDestination, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, big.NewInt(int64(s.origin))).Once().Return(clientOrigin, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, big.NewInt(int64(s.destination))).Once().Return(clientDestination, nil)
	feePricer = pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	// Calculate the total fee.
	fee, err = feePricer.GetTotalFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)

	// The expected fee should be the sum of the Origin and Destination fees, i.e. 100_250_000.
	expectedFee = big.NewInt(100_250_000) // 100.25 usd
	s.Equal(expectedFee, fee)
}
