package quoter_test

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (s *QuoterSuite) TestGenerateQuotes() {
	s.T().Skip("Needs to be fixed or removed")
	quotesReturned, err := s.manager.GenerateQuotes(1, common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"), big.NewInt(100))
	s.Require().NoError(err)
	s.Require().Len(quotesReturned, 1)
}
