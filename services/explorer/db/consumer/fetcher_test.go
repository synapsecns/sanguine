package consumer_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
	"math/big"
)

func (c *ConsumerSuite) TestFetchLogsInRange() {
	contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	chainID := gofakeit.Uint32()

	// Store 10 logs
	for blockNumber := 0; blockNumber < 10; blockNumber++ {
		storeLog := testutil.BuildLog(contractAddress, uint64(blockNumber), &c.logIndex)
		err := c.eventDB.StoreLog(c.GetTestContext(), storeLog, chainID)
		Nil(c.T(), err)
	}

	// Fetch logs from 4 to 8.
	fetcher := consumer.NewFetcher(c.gqlClient)
	logs, err := fetcher.FetchLogsInRange(c.GetTestContext(), chainID, 4, 8)
	Nil(c.T(), err)
	Equal(c.T(), 5, len(logs))
}

func (c *ConsumerSuite) TestToken() {
	fetcher, err := consumer.NewBridgeConfigFetcher(c.bridgeConfigContract.Address(), c.testBackend)
	Nil(c.T(), err)

	curentBlockNumber, err := c.testBackend.BlockNumber(c.GetTestContext())
	Nil(c.T(), err)

	for _, testToken := range testTokens {
		tokenID, err := fetcher.GetTokenID(c.GetTestContext(), uint32(testToken.ChainId.Uint64()), uint32(curentBlockNumber), common.HexToAddress(testToken.TokenAddress))

		Nil(c.T(), err)
		Equal(c.T(), *tokenID, testToken.tokenID)
		token, err := fetcher.GetToken(c.GetTestContext(), uint32(testToken.ChainId.Uint64()), uint32(curentBlockNumber), *tokenID)
		Nil(c.T(), err)
		tokenOut := *token
		Equal(c.T(), common.HexToAddress(testToken.TokenAddress).String(), common.HexToAddress(tokenOut.TokenAddress).String())
		Equal(c.T(), testToken.SwapFee, token.SwapFee)
		Equal(c.T(), testToken.IsUnderlying, token.IsUnderlying)
	}
}
