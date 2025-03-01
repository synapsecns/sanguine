package pricer_test

import (
	"fmt"
	"math/big"

	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	clientMocks "github.com/synapsecns/sanguine/ethergo/client/mocks"
	fetcherMocks "github.com/synapsecns/sanguine/ethergo/submitter/mocks"
	"github.com/synapsecns/sanguine/services/rfq/relayer/pricer"
	priceMocks "github.com/synapsecns/sanguine/services/rfq/relayer/pricer/mocks"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
)

var defaultPrices = map[string]float64{"ETH": 2000., "USDC": 1., "MATIC": 0.5, "BNB": 600, "BTC": 95000}

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

func (s *PricerSuite) TestPricePairs() {
	// Build a new FeePricer with a mocked client for fetching gas price and token price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	client := new(clientMocks.EVM)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, mock.Anything).Twice().Return(client, nil)
	priceFetcher := getPriceFetcher(nil)
	priceFetcher.On(testsuite.GetFunctionName(priceFetcher.GetPrice), mock.Anything, "USDC").Return(1., nil)
	priceFetcher.On(testsuite.GetFunctionName(priceFetcher.GetPrice), mock.Anything, "DirectUSD").Return(1., nil)
	priceFetcher.On(testsuite.GetFunctionName(priceFetcher.GetPrice), mock.Anything, "ETH").Return(2000., nil)
	priceFetcher.On(testsuite.GetFunctionName(priceFetcher.GetPrice), mock.Anything, "MATIC").Return(0.5, nil)
	priceFetcher.On(testsuite.GetFunctionName(priceFetcher.GetPrice), mock.Anything, "BTC").Return(95000., nil)
	priceFetcher.On(testsuite.GetFunctionName(priceFetcher.GetPrice), mock.Anything, "BNB").Return(600., nil)
	priceFetcher.On(testsuite.GetFunctionName(priceFetcher.GetPrice), mock.Anything, "HYPE").Return(15., nil)
	feePricer := pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	// 1 ETH priced into various values
	baseValueWei := big.NewInt(1e18)
	fee, err := feePricer.PricePair(s.GetTestContext(), "Case 1 - ETH>USDC", s.origin, s.destination, "ETH", "USDC", *baseValueWei)
	s.NoError(err)
	expectedWei := new(big.Int).Mul(big.NewInt(2000), big.NewInt(1e6))
	s.Equal(expectedWei, fee.PricedToken.Wei)

	fee, err = feePricer.PricePair(s.GetTestContext(), "Case 2 - ETH>MATIC", s.origin, s.destination, "ETH", "MATIC", *baseValueWei)
	s.NoError(err)
	// $2000 of ETH priced into Matic at $0.50 = 4000 WEI
	expectedWei = new(big.Int).Mul(big.NewInt(4000), big.NewInt(1e18))
	s.Equal(expectedWei, fee.PricedToken.Wei)

	fee, err = feePricer.PricePair(s.GetTestContext(), "Case 3 - ETH>DirectUSD", s.origin, s.destination, "ETH", "DirectUSD", *baseValueWei)
	s.NoError(err)
	expectedWei = new(big.Int).Mul(big.NewInt(2000), big.NewInt(1e5))
	s.Equal(expectedWei, fee.PricedToken.Wei)

	// $2000 of ETH priced into BNB at $600 = 3333333333333333333 WEI  (3.333~ BNB)
	fee, err = feePricer.PricePair(s.GetTestContext(), "Case 4 - ETH>BNB", s.origin, s.destination, "ETH", "BNB", *baseValueWei)
	s.NoError(err)
	expectedWei = big.NewInt(3333333333333333333)
	s.Equal(expectedWei, fee.PricedToken.Wei)

	// 3.333~~~ of BNB priced back into ETH at $2000 = 1 ETH
	fee, err = feePricer.PricePair(s.GetTestContext(), "Case 5 - BNB>ETH", s.destination, s.origin, "BNB", "ETH", *fee.PricedToken.Wei)
	s.NoError(err)
	expectedWei = big.NewInt(999999999999999999) // 0.999~ ETH in reality due to precision loss of pricing
	s.Equal(expectedWei, fee.PricedToken.Wei)

	// $2000 of ETH priced into BTC at $95,000.00 = 2105263 WEI  (0.02105263 BTC)
	fee, err = feePricer.PricePair(s.GetTestContext(), "Case 6 - ETH>BTC", s.origin, s.destination, "ETH", "BTC", *baseValueWei)
	s.NoError(err)
	expectedWei = big.NewInt(2105263)
	s.Equal(expectedWei, fee.PricedToken.Wei)

	// 0.02105263 of BTC priced back into ETH at $2000 = 1 ETH
	fee, err = feePricer.PricePair(s.GetTestContext(), "Case 7 - BTC>ETH", s.destination, s.origin, "BTC", "ETH", *fee.PricedToken.Wei)
	s.NoError(err)
	expectedWei = big.NewInt(999999924999999999) // 0.999~ ETH in reality due to precision loss of pricing from BTC's 8 decimals
	s.Equal(expectedWei, fee.PricedToken.Wei)

	// test w/ random unusually high amount of eth (15734.985734985734530000)
	baseValueWei = new(big.Int).Mul(big.NewInt(1573498573498573453), big.NewInt(1e4))

	fee, err = feePricer.PricePair(s.GetTestContext(), "Case 8 - ETH>DirectUSD", s.origin, s.destination, "ETH", "DirectUSD", *baseValueWei)
	s.NoError(err)
	expectedWei = big.NewInt(3146997146997) // $31,469,971.46997 with ETH at $2000
	s.Equal(expectedWei, fee.PricedToken.Wei)
}

func (s *PricerSuite) TestGetOriginFee() {
	// Build a new FeePricer with a mocked client for fetching gas price and token price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	client := new(clientMocks.EVM)
	currentHeader := big.NewInt(100_000_000_000) // 100 gwei
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(currentHeader, nil)
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
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(nil, fmt.Errorf("could not fetch header"))
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
	currentHeader := big.NewInt(100_000_000_000) // 100 gwei
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Return(currentHeader, nil)
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
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(nil, fmt.Errorf("could not fetch header"))
	fee, err = feePricer.GetOriginFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)
	s.Equal(expectedFee, fee)
}

func (s *PricerSuite) TestGetDestinationFee() {
	// Build a new FeePricer with a mocked client for fetching gas price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	client := new(clientMocks.EVM)
	currentHeader := big.NewInt(500_000_000_000) // 500 gwei
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(currentHeader, nil)
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
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(nil, fmt.Errorf("could not fetch header"))
	fee, err = feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)
	s.Equal(expectedFee, fee)
}

func (s *PricerSuite) TestGetDestinationFeeUSDC() {
	// Build a new FeePricer with a mocked client for fetching gas price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	client := new(clientMocks.EVM)
	currentHeader := big.NewInt(500_000_000_000) // 500 gwei produces a ~25 cent destination fee on our spoofed dest chain of Polygon
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(currentHeader, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, mock.Anything).Twice().Return(client, nil)
	priceFetcher := getPriceFetcher(nil)
	feePricer := pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	// Calculate the destination fee.
	fee, err := feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)

	expectedFee := big.NewInt(0.25 * 1e6) // USDC price = $1.00   .. expected cost = 0.25 USDC units =  0.25 usd
	s.Equal(expectedFee, fee)

	// Ensure that the fee has been cached.
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(nil, fmt.Errorf("could not fetch header"))
	fee, err = feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)
	s.Equal(expectedFee, fee)
}

func (s *PricerSuite) TestGetDestinationFeeMATIC() {
	// Build a new FeePricer with a mocked client for fetching gas price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	client := new(clientMocks.EVM)
	currentHeader := big.NewInt(500_000_000_000) // 500 gwei produces a ~25 cent destination fee on our spoofed dest chain of Polygon
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(currentHeader, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, mock.Anything).Twice().Return(client, nil)
	priceFetcher := getPriceFetcher(nil)
	feePricer := pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	// Calculate the destination fee.
	fee, err := feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "MATIC", true)
	s.NoError(err)

	expectedFee := big.NewInt(0.5 * 1e18) // spoofed MATIC price = $0.5   .. expected cost = 0.5 MATIC units =  0.25 usd
	s.Equal(expectedFee, fee)

	// Ensure that the fee has been cached.
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(nil, fmt.Errorf("could not fetch header"))
	fee, err = feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "MATIC", true)
	s.NoError(err)
	s.Equal(expectedFee, fee)
}

func (s *PricerSuite) TestGetDestinationFeeETH() {
	// Build a new FeePricer with a mocked client for fetching gas price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	client := new(clientMocks.EVM)
	currentHeader := big.NewInt(500_000_000_000) // 500 gwei produces a ~25 cent destination fee on our spoofed dest chain of Polygon
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(currentHeader, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, mock.Anything).Twice().Return(client, nil)
	priceFetcher := getPriceFetcher(nil)
	feePricer := pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	// Calculate the destination fee.
	fee, err := feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "ETH", true)
	s.NoError(err)

	expectedFee := big.NewInt(124999999999999) // spoofed ETH price = $2000   .. expected cost = 0.0001249~~ ETH units =  0.25 usd
	s.Equal(expectedFee, fee)

	// Ensure that the fee has been cached.
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(nil, fmt.Errorf("could not fetch header"))
	fee, err = feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "ETH", true)
	s.NoError(err)
	s.Equal(expectedFee, fee)
}

func (s *PricerSuite) TestGetDestinationFeeBNB() {
	// Build a new FeePricer with a mocked client for fetching gas price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	client := new(clientMocks.EVM)
	currentHeader := big.NewInt(500_000_000_000) // 500 gwei produces a ~25 cent destination fee on our spoofed dest chain of Polygon
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(currentHeader, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, mock.Anything).Twice().Return(client, nil)
	priceFetcher := getPriceFetcher(nil)
	feePricer := pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	// Calculate the destination fee.
	fee, err := feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "BNB", true)
	s.NoError(err)

	expectedFee := big.NewInt(416666666666666) // spoofed BNB price = $600   .. expected cost = 0.00041666~~ BNB units =  0.25 usd
	s.Equal(expectedFee, fee)

	// Ensure that the fee has been cached.
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(nil, fmt.Errorf("could not fetch header"))
	fee, err = feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "BNB", true)
	s.NoError(err)
	s.Equal(expectedFee, fee)
}

func (s *PricerSuite) TestGetDestinationFeeBTC() {
	// Build a new FeePricer with a mocked client for fetching gas price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	client := new(clientMocks.EVM)
	currentHeader := big.NewInt(500_000_000_000) // 500 gwei produces a ~25 cent destination fee on our spoofed dest chain of Polygon
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(currentHeader, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, mock.Anything).Twice().Return(client, nil)
	priceFetcher := getPriceFetcher(nil)
	feePricer := pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	// Calculate the destination fee.
	fee, err := feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "BTC", true)
	s.NoError(err)

	expectedFee := big.NewInt(263) // spoofed BTC price = $95,000   .. expected cost = 0.00000263 BTC units =  0.25 usd
	s.Equal(expectedFee, fee)

	// Ensure that the fee has been cached.
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(nil, fmt.Errorf("could not fetch header"))
	fee, err = feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "BTC", true)
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
	currentHeader := big.NewInt(500_000_000_000) // 500 gwei
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Return(currentHeader, nil)
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
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(nil, fmt.Errorf("could not fetch header"))
	fee, err = feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)
	s.Equal(expectedFee, fee)
}

func (s *PricerSuite) TestGetTotalFee() {
	// Build a new FeePricer with a mocked client for fetching gas price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	clientOrigin := new(clientMocks.EVM)
	clientDestination := new(clientMocks.EVM)
	headerOrigin := big.NewInt(100_000_000_000)      // 100 gwei
	headerDestination := big.NewInt(500_000_000_000) // 500 gwei
	clientOrigin.On(testsuite.GetFunctionName(clientOrigin.SuggestGasPrice), mock.Anything).Once().Return(headerOrigin, nil)
	clientDestination.On(testsuite.GetFunctionName(clientDestination.SuggestGasPrice), mock.Anything).Once().Return(headerDestination, nil)
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
	currentHeader := big.NewInt(100_000_000_000) // 100 gwei
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Once().Return(currentHeader, nil)
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
	currentHeader = big.NewInt(200_000_000_000) // 200 gwei
	client.On(testsuite.GetFunctionName(client.SuggestGasPrice), mock.Anything).Return(currentHeader, nil)
	s.Eventually(func() bool {
		gasPrice, err = feePricer.GetGasPrice(s.GetTestContext(), s.origin)
		s.NoError(err)
		expectedGasPrice = big.NewInt(200_000_000_000) // 200 gwei
		return expectedGasPrice.String() == gasPrice.String()
	})
}

func (s *PricerSuite) TestGetTotalFeeWithMultiplier() {
	// Override the fixed fee multiplier to greater than 1.
	s.config.BaseChainConfig.QuoteFixedFeeMultiplier = relconfig.NewFloatPtr(2)
	s.config.BaseChainConfig.RelayFixedFeeMultiplier = relconfig.NewFloatPtr(4)

	// Build a new FeePricer with a mocked client for fetching gas price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	clientOrigin := new(clientMocks.EVM)
	clientDestination := new(clientMocks.EVM)
	headerOrigin := big.NewInt(100_000_000_000)      // 100 gwei
	headerDestination := big.NewInt(500_000_000_000) // 500 gwei
	clientOrigin.On(testsuite.GetFunctionName(clientOrigin.SuggestGasPrice), mock.Anything).Once().Return(headerOrigin, nil)
	clientDestination.On(testsuite.GetFunctionName(clientDestination.SuggestGasPrice), mock.Anything).Once().Return(headerDestination, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, big.NewInt(int64(s.origin))).Once().Return(clientOrigin, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, big.NewInt(int64(s.destination))).Once().Return(clientDestination, nil)
	priceFetcher := getPriceFetcher(nil)
	feePricer := pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	go func() { feePricer.Start(s.GetTestContext()) }()

	// Calculate the total fee [quote].
	fee, err := feePricer.GetTotalFee(s.GetTestContext(), s.origin, s.destination, "USDC", true)
	s.NoError(err)

	// The expected fee should be the sum of the Origin and Destination fees, i.e. 200_500_000.
	expectedFee := big.NewInt(200_500_000) // 200.50 usd
	s.Equal(expectedFee, fee)

	// Calculate the total fee [relay].
	fee, err = feePricer.GetTotalFee(s.GetTestContext(), s.origin, s.destination, "USDC", false)
	s.NoError(err)

	// The expected fee should be the sum of the Origin and Destination fees, i.e. 401_000_000.
	expectedFee = big.NewInt(401_000_000) // 401 usd
	s.Equal(expectedFee, fee)

	// Override the fixed fee multiplier to less than 1; should default to 1.
	s.config.BaseChainConfig.QuoteFixedFeeMultiplier = relconfig.NewFloatPtr(-1)

	// Build a new FeePricer with a mocked client for fetching gas price.
	clientOrigin.On(testsuite.GetFunctionName(clientOrigin.SuggestGasPrice), mock.Anything).Once().Return(headerOrigin, nil)
	clientDestination.On(testsuite.GetFunctionName(clientDestination.SuggestGasPrice), mock.Anything).Once().Return(headerDestination, nil)
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
	s.config.BaseChainConfig.QuoteFixedFeeMultiplier = relconfig.NewFloatPtr(0)

	// Build a new FeePricer with a mocked client for fetching gas price.
	clientOrigin.On(testsuite.GetFunctionName(clientOrigin.SuggestGasPrice), mock.Anything).Once().Return(headerOrigin, nil)
	clientDestination.On(testsuite.GetFunctionName(clientDestination.SuggestGasPrice), mock.Anything).Once().Return(headerDestination, nil)
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
