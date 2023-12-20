package pricer_test

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/core/testsuite"
	clientMocks "github.com/synapsecns/sanguine/ethergo/client/mocks"
	fetcherMocks "github.com/synapsecns/sanguine/ethergo/submitter/mocks"
	"github.com/synapsecns/sanguine/services/rfq/relayer/pricer"
)

func (s *PricerSuite) TestGetOriginFee() {
	// Build a new FeePricer with a mocked client for fetching gas price.
	clientFetcher := new(fetcherMocks.ClientFetcher)
	client := new(clientMocks.EVM)
	currentHeader := &types.Header{BaseFee: big.NewInt(100_000_000_000)} // 100 gwei
	client.On(testsuite.GetFunctionName(client.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(currentHeader, nil)
	clientFetcher.On(testsuite.GetFunctionName(clientFetcher.GetClient), mock.Anything, mock.Anything).Twice().Return(client, nil)
	feePricer := pricer.NewFeePricer(s.feePricerConfig, s.chainConfigs, clientFetcher)

	// Calculate the origin fee.
	fee, err := feePricer.GetOriginFee(s.GetTestContext(), s.origin, s.destination, "USDC")
	s.NoError(err)

	/*
		The expected fee should be:
		fee_eth: gas_price * gas_estimate / native_decimals_factor
		fee_usd: fee_eth * eth_price_usd
		fee_usdc: fee_usd * usdc_price_usd
		fee_usdc_decimals: fee_usdc * usdc_decimals_factor
		fee_usdc_decimals = (((gas_price * gas_estimate / native_decimals_factor) * eth_price_usd) * usdc_price_usd) * usdc_decimals_factor
		So, with our numbers:
		fee_denom = (((100e9 * 500000 / 1e18) * 2000) * 1) * 1e6 = 100_000_000
	*/

	expectedFee := big.NewInt(100_000_000) // 100 usd
	s.Equal(expectedFee, fee)

	// Ensure that the fee has been cached.
	client.On(testsuite.GetFunctionName(client.HeaderByNumber), mock.Anything, mock.Anything).Once().Return(nil, fmt.Errorf("could not fetch header"))
	fee, err = feePricer.GetOriginFee(s.GetTestContext(), s.origin, s.destination, "USDC")
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
	feePricer := pricer.NewFeePricer(s.feePricerConfig, s.chainConfigs, clientFetcher)

	// Calculate the destination fee.
	fee, err := feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "USDC")
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
	fee, err = feePricer.GetDestinationFee(s.GetTestContext(), s.origin, s.destination, "USDC")
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
	feePricer := pricer.NewFeePricer(s.feePricerConfig, s.chainConfigs, clientFetcher)

	// Calculate the total fee.
	fee, err := feePricer.GetTotalFee(s.GetTestContext(), s.origin, s.destination, "USDC")
	s.NoError(err)

	// The expected fee should be the sum of the Origin and Destination fees, i.e. 100_250_000.
	expectedFee := big.NewInt(100_250_000) // 100.25 usd
	s.Equal(expectedFee, fee)
}
