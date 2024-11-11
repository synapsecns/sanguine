package quoter_test

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/testsuite"
	fetcherMocks "github.com/synapsecns/sanguine/ethergo/submitter/mocks"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	inventoryMocks "github.com/synapsecns/sanguine/services/rfq/relayer/inventory/mocks"
	"github.com/synapsecns/sanguine/services/rfq/relayer/pricer"
	priceMocks "github.com/synapsecns/sanguine/services/rfq/relayer/pricer/mocks"
	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/util"
)

func (s *QuoterSuite) TestGenerateQuotes() {
	// Generate quotes for USDC on the destination chain.
	balance := big.NewInt(1000_000_000) // 1000 USDC
	inv := map[int]map[common.Address]*big.Int{}
	quotes, err := s.manager.GenerateQuotes(s.GetTestContext(), int(s.destination), common.HexToAddress("0x0b2c639c533813f4aa9d7837caf62653d097ff85"), balance, inv)
	s.Require().NoError(err)

	// Verify the quotes are generated as expected.
	expectedQuotes := []model.PutRelayerQuoteRequest{
		{
			OriginChainID:           int(s.origin),
			OriginTokenAddr:         "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			DestChainID:             int(s.destination),
			DestTokenAddr:           "0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85",
			DestAmount:              balance.String(),
			MaxOriginAmount:         balance.String(),
			FixedFee:                "100050000",
			OriginFastBridgeAddress: common.HexToAddress("0x123").Hex(),
			DestFastBridgeAddress:   common.HexToAddress("0x456").Hex(),
		},
	}
	s.Equal(expectedQuotes, quotes)
}

func (s *QuoterSuite) TestGenerateQuotesForNativeToken() {
	// Generate quotes for ETH on the destination chain.
	balance, _ := new(big.Int).SetString("1000000000000000000", 10) // 1 ETH
	inv := map[int]map[common.Address]*big.Int{}
	quotes, err := s.manager.GenerateQuotes(s.GetTestContext(), int(s.destinationEth), util.EthAddress, balance, inv)
	s.Require().NoError(err)

	minGasToken, err := s.config.GetMinGasToken(int(s.destination))
	s.NoError(err)
	expectedQuoteAmount := new(big.Int).Sub(balance, minGasToken)

	// Verify the quotes are generated as expected.
	expectedQuotes := []model.PutRelayerQuoteRequest{
		{
			OriginChainID:           int(s.origin),
			OriginTokenAddr:         util.EthAddress.String(),
			DestChainID:             int(s.destinationEth),
			DestTokenAddr:           util.EthAddress.String(),
			DestAmount:              expectedQuoteAmount.String(),
			MaxOriginAmount:         expectedQuoteAmount.String(),
			FixedFee:                "150000000000000000", // (500k gas + 1m gas) * 100 gwei
			OriginFastBridgeAddress: common.HexToAddress("0x123").Hex(),
			DestFastBridgeAddress:   common.HexToAddress("0x789").Hex(),
		},
	}
	s.Equal(expectedQuotes, quotes)

	// Set MinGasToken and make sure it is accounted for in the DestAmount.
	s.config.BaseChainConfig.MinGasToken = "100000000000000000" // 0.1 ETH
	s.manager.SetConfig(s.config)

	quotes, err = s.manager.GenerateQuotes(s.GetTestContext(), int(s.destinationEth), util.EthAddress, balance, inv)
	s.Require().NoError(err)

	minGasToken, err = s.config.GetMinGasToken(int(s.destination))
	s.NoError(err)
	expectedQuoteAmount = new(big.Int).Sub(balance, minGasToken)

	// Verify the quotes are generated as expected.
	expectedQuotes = []model.PutRelayerQuoteRequest{
		{
			OriginChainID:           int(s.origin),
			OriginTokenAddr:         util.EthAddress.String(),
			DestChainID:             int(s.destinationEth),
			DestTokenAddr:           util.EthAddress.String(),
			DestAmount:              expectedQuoteAmount.String(),
			MaxOriginAmount:         expectedQuoteAmount.String(),
			FixedFee:                "150000000000000000", // (500k gas + 1m gas) * 100 gwei
			OriginFastBridgeAddress: common.HexToAddress("0x123").Hex(),
			DestFastBridgeAddress:   common.HexToAddress("0x789").Hex(),
		},
	}
	s.Equal(expectedQuotes, quotes)

	// Set MinGasToken to balance and make sure no quotes are generated.
	s.config.BaseChainConfig.MinGasToken = "1000000000000000001" // 0.1 ETH
	s.manager.SetConfig(s.config)

	quotes, err = s.manager.GenerateQuotes(s.GetTestContext(), int(s.destinationEth), util.EthAddress, balance, inv)
	s.NoError(err)
	s.Equal(quotes[0].DestAmount, "0")
	s.Equal(quotes[0].MaxOriginAmount, "0")
}

func (s *QuoterSuite) TestShouldProcess() {
	// Should process a valid quote.
	balance := big.NewInt(1000_000_000) // 1000 USDC
	fee := big.NewInt(100_050_000)      // 100.05 USDC
	quote := reldb.QuoteRequest{
		BlockNumber:         1,
		OriginTokenDecimals: 6,
		DestTokenDecimals:   6,
		Transaction: fastbridge.IFastBridgeBridgeTransaction{
			OriginChainId: s.origin,
			DestChainId:   s.destination,
			OriginToken:   common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
			DestToken:     common.HexToAddress("0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85"),
			OriginAmount:  balance,
			DestAmount:    new(big.Int).Sub(balance, fee),
		},
	}
	s.True(s.manager.ShouldProcess(s.GetTestContext(), quote))

	// Set different numbers of decimals for origin / dest tokens; should never process this.
	badQuote := quote
	badQuote.DestTokenDecimals = 18
	s.False(s.manager.ShouldProcess(s.GetTestContext(), badQuote))

	// Toggle relayPaused
	s.manager.SetRelayPaused(true)
	s.False(s.manager.ShouldProcess(s.GetTestContext(), quote))
	s.manager.SetRelayPaused(false)
	s.True(s.manager.ShouldProcess(s.GetTestContext(), quote))

	// Set max relay amount
	originTokenCfg := s.config.Chains[int(s.origin)].Tokens["USDC"]
	originTokenCfg.MaxRelayAmount = "900" // less than balance
	s.config.Chains[int(s.origin)].Tokens["USDC"] = originTokenCfg
	s.manager.SetConfig(s.config)
	s.False(s.manager.ShouldProcess(s.GetTestContext(), quote))
}

func (s *QuoterSuite) TestIsProfitable() {
	// Set fee to breakeven; i.e. destAmount = originAmount - fee.
	balance := big.NewInt(1000_000_000) // 1000 USDC
	fee := big.NewInt(100_050_000)      // 100.05 USDC
	quote := reldb.QuoteRequest{
		BlockNumber:         1,
		OriginTokenDecimals: 6,
		DestTokenDecimals:   6,
		Transaction: fastbridge.IFastBridgeBridgeTransaction{
			OriginChainId: s.origin,
			DestChainId:   s.destination,
			OriginToken:   common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
			DestToken:     common.HexToAddress("0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85"),
			OriginAmount:  balance,
			DestAmount:    new(big.Int).Sub(balance, fee),
		},
	}
	s.True(s.manager.IsProfitable(s.GetTestContext(), quote))

	// Set fee to greater than breakeven; i.e. destAmount > originAmount - fee.
	quote.Transaction.DestAmount = new(big.Int).Sub(balance, new(big.Int).Mul(fee, big.NewInt(2)))
	s.True(s.manager.IsProfitable(s.GetTestContext(), quote))

	// Set fee to less than breakeven; i.e. destAmount < originAmount - fee.
	quote.Transaction.DestAmount = balance
	s.False(s.manager.IsProfitable(s.GetTestContext(), quote))

	origin := int(s.origin)
	dest := int(s.destination)
	setQuoteOffsets := func(originOffset, destOffset float64) {
		originTokenCfg := s.config.Chains[origin].Tokens["USDC"]
		originTokenCfg.QuoteOffsetBps = originOffset
		s.config.Chains[origin].Tokens["USDC"] = originTokenCfg
		destTokenCfg := s.config.Chains[dest].Tokens["USDC"]
		destTokenCfg.QuoteOffsetBps = destOffset
		s.config.Chains[dest].Tokens["USDC"] = destTokenCfg
		s.manager.SetConfig(s.config)
	}
	quote.Transaction.DestAmount = new(big.Int).Sub(balance, fee)

	// Set dest offset to 20%; we send a token that is more valuable -> not profitable
	setQuoteOffsets(0, 2000)
	s.False(s.manager.IsProfitable(s.GetTestContext(), quote))

	// Set dest offset to -20%; we send a token that is less valuable -> profitable
	setQuoteOffsets(0, -2000)
	s.True(s.manager.IsProfitable(s.GetTestContext(), quote))

	// Set origin offset to 20%; we get a token that is more valuable -> not profitable
	setQuoteOffsets(2000, 0)
	s.True(s.manager.IsProfitable(s.GetTestContext(), quote))

	// Set origin offset to -20%; we send a token that is less valuable -> not profitable
	setQuoteOffsets(-2000, 0)
	s.False(s.manager.IsProfitable(s.GetTestContext(), quote))
}

func (s *QuoterSuite) TestGetOriginAmountActiveQuotes() {
	origin := int(s.origin)
	dest := int(s.destination)
	address := common.HexToAddress("0x0b2c639c533813f4aa9d7837caf62653d097ff85")
	originAddr := common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	balance := big.NewInt(1000_000_000)     // 1000 USDC
	depositAmount := big.NewInt(50_000_000) // 50 USDC

	type quoteParams struct {
		quotePct       float64
		quoteOffset    float64
		minQuoteAmount string
		maxBalance     string
		maxQuoteAmount string
	}

	setQuoteParams := func(params quoteParams) {
		s.config.BaseChainConfig.QuotePct = &params.quotePct
		destTokenCfg := s.config.Chains[dest].Tokens["USDC"]
		destTokenCfg.MinQuoteAmount = params.minQuoteAmount
		destTokenCfg.MaxRelayAmount = params.maxQuoteAmount
		originTokenCfg := s.config.Chains[origin].Tokens["USDC"]
		originTokenCfg.QuoteOffsetBps = params.quoteOffset
		originTokenCfg.MaxBalance = &params.maxBalance
		originTokenCfg.MaxRelayAmount = params.maxQuoteAmount
		s.config.Chains[dest].Tokens["USDC"] = destTokenCfg
		s.config.Chains[origin].Tokens["USDC"] = originTokenCfg
		s.manager.SetConfig(s.config)
	}

	input := quoter.QuoteInput{
		OriginChainID:     origin,
		DestChainID:       dest,
		OriginTokenAddr:   originAddr,
		DestTokenAddr:     address,
		OriginBalance:     balance,
		DestBalance:       balance,
		OriginAmountExact: depositAmount,
	}

	// Set default quote params; should return the depositAmount.
	quoteAmount, err := s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount := depositAmount
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 50 with MinQuoteAmount of 0; should still be 100% of deposit amount
	// IE: we're only willing to use 50% of our balance, but deposit amount is well below that.
	setQuoteParams(quoteParams{
		quotePct:       50,
		quoteOffset:    0,
		minQuoteAmount: "0",
		maxBalance:     "0",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = depositAmount
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 1 with MinQuoteAmount of 0;
	// should return zero even though we can partially cover -- because deposit amount is all-or-nothing.
	// IE: we're only willing to use 1% of our balance, which does not cover deposit amount.
	setQuoteParams(quoteParams{
		quotePct:       1,
		quoteOffset:    0,
		minQuoteAmount: "0",
		maxBalance:     "0",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = big.NewInt(0)
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 5 with MinQuoteAmount of 0;
	// should the deposit amount because 5% of 1K is just enough to cover.
	setQuoteParams(quoteParams{
		quotePct:       5,
		quoteOffset:    0,
		minQuoteAmount: "0",
		maxBalance:     "0",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = depositAmount
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 5 with MinQuoteAmount of 0;
	// should return the deposit amount because 5% of 1K balance is just enough to cover.
	setQuoteParams(quoteParams{
		quotePct:       5,
		quoteOffset:    0,
		minQuoteAmount: "500",
		maxBalance:     "0",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = depositAmount
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 25 with MinQuoteAmount of 500;
	// minQuoteAmt ceiling = 500
	// quotePct 25% * 1000 = 250
	// output s/b depositAmount becase it does not exceed either
	setQuoteParams(quoteParams{
		quotePct:       25,
		quoteOffset:    0,
		minQuoteAmount: "500",
		maxBalance:     "0",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = depositAmount
	s.Equal(expectedAmount, quoteAmount)

	// test minQuoteAmount ceiling *overriding* quotePct ceiling
	// output s/b depositAmount even though quotePct amount wont cover it, due to override
	setQuoteParams(quoteParams{
		quotePct:       1,
		quoteOffset:    0,
		minQuoteAmount: "500",
		maxBalance:     "0",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = depositAmount
	s.Equal(expectedAmount, quoteAmount)

	// test quotePct ceiling *overriding* minQuoteAmount ceiling
	// output s/b depositAmount even though minQuoteAmount amount wont cover it, due to override
	setQuoteParams(quoteParams{
		quotePct:       25,
		quoteOffset:    0,
		minQuoteAmount: "5",
		maxBalance:     "0",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = depositAmount
	s.Equal(expectedAmount, quoteAmount)

	// test depositAmount below both the quotePct ceiling and minQuoteAmount ceiling
	// output s/b 0 as neither ceiling setting accommodates the deposit amount
	setQuoteParams(quoteParams{
		quotePct:       1,
		quoteOffset:    0,
		minQuoteAmount: "5",
		maxBalance:     "0",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = big.NewInt(0)
	s.Equal(expectedAmount, quoteAmount)

	// Ceiling params much higher than deposit amount, should still return deposit amount exactly
	setQuoteParams(quoteParams{
		quotePct:       100,
		quoteOffset:    0,
		minQuoteAmount: "0",
		maxBalance:     "0",
		maxQuoteAmount: "500",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = depositAmount
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 25 with MinQuoteAmount of 1500 and MaxBalance of 1200;
	// effective ceiling is 200 due to maxBalance only allowing 200 more units on origin.
	// Since depositAmt is below this, output s/b deposit amount
	setQuoteParams(quoteParams{
		quotePct:       25,
		quoteOffset:    0,
		minQuoteAmount: "1500",
		maxBalance:     "1200",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = depositAmount
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 25 with MinQuoteAmount of 1500 and MaxBalance of 1025;
	// effective ceiling is 200 due to maxBalance only allowing 25 more units on origin.
	// Since depositAmt is above this ceiling, output s/b zero
	setQuoteParams(quoteParams{
		quotePct:       25,
		quoteOffset:    0,
		minQuoteAmount: "1500",
		maxBalance:     "1025",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = big.NewInt(0)
	s.Equal(expectedAmount, quoteAmount)

	// Toggle insufficient gas; should be 0.
	s.setGasSufficiency(false)
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = big.NewInt(0)
	s.Equal(expectedAmount, quoteAmount)

	input.OriginAmountExact = new(big.Int).Add(balance, balance)
	// depositAmount beyond our balance.
	// output s/b zero even with generous limits
	setQuoteParams(quoteParams{
		quotePct:       100,
		quoteOffset:    0,
		minQuoteAmount: "1500",
		maxBalance:     "2000",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = big.NewInt(0)
	s.Equal(expectedAmount, quoteAmount)

	input.OriginAmountExact = big.NewInt(0)
	// depositAmount set to zero. output s/b zero
	setQuoteParams(quoteParams{
		quotePct:       100,
		quoteOffset:    0,
		minQuoteAmount: "1500",
		maxBalance:     "2000",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = big.NewInt(0)
	s.Equal(expectedAmount, quoteAmount)

	input.OriginAmountExact = big.NewInt(-55000000)
	// depositAmount set to negative. output s/b zero
	setQuoteParams(quoteParams{
		quotePct:       100,
		quoteOffset:    0,
		minQuoteAmount: "1500",
		maxBalance:     "2000",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = big.NewInt(0)
	s.Equal(expectedAmount, quoteAmount)
}

func (s *QuoterSuite) TestGetOriginAmount() {
	origin := int(s.origin)
	dest := int(s.destination)
	address := common.HexToAddress("0x0b2c639c533813f4aa9d7837caf62653d097ff85")
	originAddr := common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	balance := big.NewInt(1000_000_000) // 1000 USDC

	type quoteParams struct {
		quotePct       float64
		quoteOffset    float64
		minQuoteAmount string
		maxBalance     string
		maxQuoteAmount string
	}

	setQuoteParams := func(params quoteParams) {
		s.config.BaseChainConfig.QuotePct = &params.quotePct
		destTokenCfg := s.config.Chains[dest].Tokens["USDC"]
		destTokenCfg.MinQuoteAmount = params.minQuoteAmount
		destTokenCfg.MaxRelayAmount = params.maxQuoteAmount
		originTokenCfg := s.config.Chains[origin].Tokens["USDC"]
		originTokenCfg.QuoteOffsetBps = params.quoteOffset
		originTokenCfg.MaxBalance = &params.maxBalance
		originTokenCfg.MaxRelayAmount = params.maxQuoteAmount
		s.config.Chains[dest].Tokens["USDC"] = destTokenCfg
		s.config.Chains[origin].Tokens["USDC"] = originTokenCfg
		s.manager.SetConfig(s.config)
	}

	input := quoter.QuoteInput{
		OriginChainID:   origin,
		DestChainID:     dest,
		OriginTokenAddr: originAddr,
		DestTokenAddr:   address,
		OriginBalance:   balance,
		DestBalance:     balance,
	}

	// Set default quote params; should return the balance.
	quoteAmount, err := s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount := balance
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 50 with MinQuoteAmount of 0; should be 50% of balance.
	setQuoteParams(quoteParams{
		quotePct:       50,
		quoteOffset:    0,
		minQuoteAmount: "0",
		maxBalance:     "0",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = big.NewInt(500_000_000)
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 25 with MinQuoteAmount of 500; should be 50% of balance.
	setQuoteParams(quoteParams{
		quotePct:       25,
		quoteOffset:    0,
		minQuoteAmount: "500",
		maxBalance:     "0",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = big.NewInt(500_000_000)
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 25 with MinQuoteAmount of 500; should be 50% of balance.
	setQuoteParams(quoteParams{
		quotePct:       25,
		quoteOffset:    0,
		minQuoteAmount: "500",
		maxBalance:     "0",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = big.NewInt(500_000_000)
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 25 with MinQuoteAmount of 1500; should be total balance.
	setQuoteParams(quoteParams{
		quotePct:       25,
		quoteOffset:    0,
		minQuoteAmount: "1500",
		maxBalance:     "0",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = big.NewInt(1000_000_000)
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 100 with MinQuoteAmount of 0 and MaxRelayAmount of 500; should be 500.
	setQuoteParams(quoteParams{
		quotePct:       100,
		quoteOffset:    0,
		minQuoteAmount: "0",
		maxBalance:     "0",
		maxQuoteAmount: "500",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = big.NewInt(500_000_000)
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 25 with MinQuoteAmount of 1500 and MaxBalance of 1200; should be 200.
	setQuoteParams(quoteParams{
		quotePct:       25,
		quoteOffset:    0,
		minQuoteAmount: "1500",
		maxBalance:     "1200",
	})
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = big.NewInt(200_000_000)
	s.Equal(expectedAmount, quoteAmount)

	// Toggle insufficient gas; should be 0.
	s.setGasSufficiency(false)
	quoteAmount, err = s.manager.GetOriginAmount(s.GetTestContext(), input)
	s.NoError(err)
	expectedAmount = big.NewInt(0)
	s.Equal(expectedAmount, quoteAmount)
}

func (s *QuoterSuite) setGasSufficiency(sufficient bool) {
	clientFetcher := new(fetcherMocks.ClientFetcher)
	priceFetcher := new(priceMocks.CoingeckoPriceFetcher)
	priceFetcher.On(testsuite.GetFunctionName(priceFetcher.GetPrice), mock.Anything, mock.Anything).Return(0., fmt.Errorf("not using mocked price"))
	feePricer := pricer.NewFeePricer(s.config, clientFetcher, priceFetcher, metrics.NewNullHandler())
	inventoryManager := new(inventoryMocks.Manager)
	inventoryManager.On(testsuite.GetFunctionName(inventoryManager.HasSufficientGas), mock.Anything, mock.Anything, mock.Anything).Return(sufficient, nil)
	mgr, err := quoter.NewQuoterManager(s.config, metrics.NewNullHandler(), inventoryManager, nil, feePricer, nil)
	s.NoError(err)

	var ok bool
	s.manager, ok = mgr.(*quoter.Manager)
	s.True(ok)
}

func (s *QuoterSuite) TestGetDestAmount() {
	balance := big.NewInt(1000_000_000) // 1000 USDC

	origin := int(s.origin)
	dest := int(s.destination)
	input := quoter.QuoteInput{
		OriginChainID: int(s.origin),
		DestChainID:   int(s.destination),
		OriginBalance: balance,
		DestBalance:   balance,
	}
	setQuoteParams := func(originQuoteOffsetBps, destQuoteOffsetBps, quoteWidthBps float64) {
		tokenCfg := s.config.Chains[origin].Tokens["USDC"]
		tokenCfg.QuoteOffsetBps = originQuoteOffsetBps
		s.config.Chains[origin].Tokens["USDC"] = tokenCfg
		tokenCfg = s.config.Chains[dest].Tokens["USDC"]
		tokenCfg.QuoteOffsetBps = destQuoteOffsetBps
		tokenCfg.QuoteWidthBps = quoteWidthBps
		s.config.Chains[dest].Tokens["USDC"] = tokenCfg
		s.manager.SetConfig(s.config)
	}

	// Set default quote params; should return the balance.
	destAmount, err := s.manager.GetDestAmount(s.GetTestContext(), balance, "USDC", input)
	s.NoError(err)
	expectedAmount := balance
	s.Equal(expectedAmount, destAmount)

	// Set QuoteWidthBps to 100, should return 99% of balance.
	setQuoteParams(0, 0, 100)
	destAmount, err = s.manager.GetDestAmount(s.GetTestContext(), balance, "USDC", input)
	s.NoError(err)
	expectedAmount = big.NewInt(990_000_000)
	s.Equal(expectedAmount, destAmount)

	// Set QuoteWidthBps to 500, should return 95% of balance.
	setQuoteParams(0, 0, 500)
	destAmount, err = s.manager.GetDestAmount(s.GetTestContext(), balance, "USDC", input)
	s.NoError(err)
	expectedAmount = big.NewInt(950_000_000)
	s.Equal(expectedAmount, destAmount)

	// Set QuoteWidthBps to 500 and QuoteOffsetBps to 100, should return 94% of balance.
	setQuoteParams(0, 100, 500)
	destAmount, err = s.manager.GetDestAmount(s.GetTestContext(), balance, "USDC", input)
	s.NoError(err)
	expectedAmount = big.NewInt(940_000_000)
	s.Equal(expectedAmount, destAmount)

	// Set QuoteWidthBps to 500 and QuoteOffsetBps to -100, should return 96% of balance.
	setQuoteParams(0, -100, 500)
	destAmount, err = s.manager.GetDestAmount(s.GetTestContext(), balance, "USDC", input)
	s.NoError(err)
	expectedAmount = big.NewInt(960_000_000)
	s.Equal(expectedAmount, destAmount)

	// Set QuoteWidthBps to -100, should error.
	setQuoteParams(0, 0, -100)
	destAmount, err = s.manager.GetDestAmount(s.GetTestContext(), balance, "USDC", input)
	s.Error(err)
	s.Nil(destAmount)

	// Set origin offset to 100, should return 101% of balance.
	setQuoteParams(100, 0, 0)
	destAmount, err = s.manager.GetDestAmount(s.GetTestContext(), balance, "USDC", input)
	s.NoError(err)
	expectedAmount = big.NewInt(1_010_000_000)
	s.Equal(expectedAmount, destAmount)

	// Set origin offset to -100, should return 99% of balance.
	setQuoteParams(-100, 0, 0)
	destAmount, err = s.manager.GetDestAmount(s.GetTestContext(), balance, "USDC", input)
	s.NoError(err)
	expectedAmount = big.NewInt(990_000_000)
	s.Equal(expectedAmount, destAmount)

	// Set origin offset to 100, dest offset to 300, should return 98% of balance.
	setQuoteParams(100, 300, 0)
	destAmount, err = s.manager.GetDestAmount(s.GetTestContext(), balance, "USDC", input)
	s.NoError(err)
	expectedAmount = big.NewInt(980_000_000)
	s.Equal(expectedAmount, destAmount)
}
