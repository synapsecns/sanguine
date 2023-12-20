package quoter_test

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	rfqAPIClient "github.com/synapsecns/sanguine/services/rfq/api/client"
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
