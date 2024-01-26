package quoter_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

func (s *QuoterSuite) TestGenerateQuotes() {
	// Generate quotes for USDC on the destination chain.
	balance := big.NewInt(1000_000_000) // 1000 USDC
	quotes, err := s.manager.GenerateQuotes(s.GetTestContext(), int(s.destination), common.HexToAddress("0x0b2c639c533813f4aa9d7837caf62653d097ff85"), balance)
	s.Require().NoError(err)

	// Verify the qutoes are generated as expected.
	expectedQuotes := []model.PutQuoteRequest{
		{
			OriginChainID:   int(s.origin),
			OriginTokenAddr: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
			DestChainID:     int(s.destination),
			DestTokenAddr:   "0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85",
			DestAmount:      balance.String(),
			MaxOriginAmount: balance.String(),
			FixedFee:        "100050000",
		},
	}
	s.Equal(expectedQuotes, quotes)
}

func (s *QuoterSuite) TestShouldProcess() {
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
	s.True(s.manager.ShouldProcess(s.GetTestContext(), quote))

	// Set fee to greater than breakeven; i.e. destAmount > originAmount - fee.
	quote.Transaction.DestAmount = new(big.Int).Sub(balance, new(big.Int).Mul(fee, big.NewInt(2)))
	s.True(s.manager.ShouldProcess(s.GetTestContext(), quote))

	// Set fee to less than breakeven; i.e. destAmount < originAmount - fee.
	quote.Transaction.DestAmount = balance
	s.False(s.manager.ShouldProcess(s.GetTestContext(), quote))

	// Set different numbers of decimals for origin / dest tokens; should never process this.
	quote = reldb.QuoteRequest{
		BlockNumber:         1,
		OriginTokenDecimals: 6,
		DestTokenDecimals:   18,
		Transaction: fastbridge.IFastBridgeBridgeTransaction{
			OriginChainId: s.origin,
			DestChainId:   s.destination,
			OriginToken:   common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
			DestToken:     common.HexToAddress("0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85"),
			OriginAmount:  balance,
			DestAmount:    new(big.Int).Sub(balance, fee),
		},
	}
	s.False(s.manager.ShouldProcess(s.GetTestContext(), quote))
}

func (s *QuoterSuite) TestGetQuoteAmount() {
	chainID := int(s.origin)
	address := common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	balance := big.NewInt(1000_000_000) // 1000 USDC

	setQuoteParams := func(quotePct float64, minQuoteAmount string) {
		s.config.QuotePct = quotePct
		tokenCfg := s.config.Chains[chainID].Tokens["USDC"]
		tokenCfg.MinQuoteAmount = minQuoteAmount
		s.config.Chains[chainID].Tokens["USDC"] = tokenCfg
		s.manager.SetConfig(s.config)
	}

	// Set default quote params; should return the balance.
	quoteAmount := s.manager.GetQuoteAmount(s.GetTestContext(), chainID, address, balance)
	expectedAmount := balance
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 50 with MinQuoteAmount of 0; should be 50% of balance.
	setQuoteParams(50, "0")
	quoteAmount = s.manager.GetQuoteAmount(s.GetTestContext(), chainID, address, balance)
	expectedAmount = big.NewInt(500_000_000)
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 50 with MinQuoteAmount of 500; should be 50% of balance.
	setQuoteParams(50, "500")
	quoteAmount = s.manager.GetQuoteAmount(s.GetTestContext(), chainID, address, balance)
	expectedAmount = big.NewInt(500_000_000)
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 25 with MinQuoteAmount of 500; should be 50% of balance.
	setQuoteParams(25, "500")
	quoteAmount = s.manager.GetQuoteAmount(s.GetTestContext(), chainID, address, balance)
	expectedAmount = big.NewInt(500_000_000)
	s.Equal(expectedAmount, quoteAmount)

	// Set QuotePct to 25 with MinQuoteAmount of 1500; should be total balance.
	setQuoteParams(25, "1500")
	quoteAmount = s.manager.GetQuoteAmount(s.GetTestContext(), chainID, address, balance)
	expectedAmount = big.NewInt(1000_000_000)
	s.Equal(expectedAmount, quoteAmount)
}

func (s *QuoterSuite) TestGetDestAmount() {
	balance := big.NewInt(1000_000_000) // 1000 USDC

	setQuoteParams := func(quoteOffsetBps int) {
		s.config.QuoteOffsetBps = quoteOffsetBps
		s.manager.SetConfig(s.config)
	}

	// Set default quote params; should return the balance.
	destAmount := s.manager.GetDestAmount(s.GetTestContext(), balance)
	expectedAmount := balance
	s.Equal(expectedAmount, destAmount)

	// Set QuoteOffsetBps to 100, should return 99% of balance.
	setQuoteParams(100)
	destAmount = s.manager.GetDestAmount(s.GetTestContext(), balance)
	expectedAmount = big.NewInt(990_000_000)
	s.Equal(expectedAmount, destAmount)

	// Set QuoteOffsetBps to 500, should return 95% of balance.
	setQuoteParams(500)
	destAmount = s.manager.GetDestAmount(s.GetTestContext(), balance)
	expectedAmount = big.NewInt(950_000_000)
	s.Equal(expectedAmount, destAmount)

	// Set QuoteOffsetBps to -100, should default to balance.
	setQuoteParams(-100)
	destAmount = s.manager.GetDestAmount(s.GetTestContext(), balance)
	expectedAmount = balance
	s.Equal(expectedAmount, destAmount)
}
