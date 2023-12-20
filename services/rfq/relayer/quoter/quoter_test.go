package quoter_test

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	rfqAPIClient "github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

func (s *QuoterSuite) TestGenerateQuotes() {
	// Generate quotes for USDC on the destination chain.
	balance := big.NewInt(1000_000_000) // 1000 USDC
	quotes, err := s.manager.GenerateQuotes(s.GetTestContext(), int(s.destination), common.HexToAddress("0x0b2c639c533813f4aa9d7837caf62653d097ff85"), balance)
	s.Require().NoError(err)

	// Verify the qutoes are generated as expected.
	expectedQuotes := []rfqAPIClient.APIQuotePutRequest{
		{
			OriginChainID:   strconv.Itoa(int(s.origin)),
			OriginTokenAddr: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
			DestChainID:     strconv.Itoa(int(s.destination)),
			DestTokenAddr:   "0x0b2C639c533813f4Aa9D7837CAf62653d097Ff85",
			DestAmount:      balance.String(),
			MaxOriginAmount: balance.String(),
			FixedFee:        "100050000",
		},
	}
	s.Equal(quotes, expectedQuotes)
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
