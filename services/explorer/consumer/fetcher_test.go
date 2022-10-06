package consumer_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/consumer"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
	"math/big"
)

func (c *ConsumerSuite) TestFetchLogsInRange() {
	defer c.cleanup()
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
	defer c.cleanup()
	fetcher, err := consumer.NewBridgeConfigFetcher(c.bridgeConfigContract.Address(), c.testBackend)
	Nil(c.T(), err)

	for _, testToken := range testTokens {
		tokenID, err := fetcher.GetTokenID(c.GetTestContext(), uint32(testToken.ChainId.Uint64()), common.HexToAddress(testToken.TokenAddress))

		Nil(c.T(), err)
		Equal(c.T(), *tokenID, testToken.tokenID)
		token, err := fetcher.GetToken(c.GetTestContext(), uint32(testToken.ChainId.Uint64()), tokenID)
		Nil(c.T(), err)
		tokenOut := *token
		Equal(c.T(), common.HexToAddress(testToken.TokenAddress).String(), common.HexToAddress(tokenOut.TokenAddress).String())
		Equal(c.T(), testToken.SwapFee, token.SwapFee)
		Equal(c.T(), testToken.IsUnderlying, token.IsUnderlying)
	}
}

func (c *ConsumerSuite) TestTimeToBlockNumber() {
	defer c.cleanup()
	fetcher := consumer.NewFetcher(c.gqlClient)
	chainID := gofakeit.Uint32()

	baseTime := uint64(0)

	// Store 10 block numbers and block times.
	for blockNumber := uint64(1); blockNumber <= 10; blockNumber++ {
		err := c.eventDB.StoreBlockTime(c.GetTestContext(), chainID, blockNumber, baseTime)
		Nil(c.T(), err)
		baseTime += uint64(gofakeit.Uint32())
	}

	targetTime := uint64(gofakeit.Uint32() * 5)

	blockNumber, err := fetcher.TimeToBlockNumber(c.GetTestContext(), chainID, 1, targetTime)
	Nil(c.T(), err)

	// Find the block number that is closest to the target time.
	var closestBlockNumber uint64
	closestBlockTime := ^uint64(0)
	for blockNumber := uint64(1); blockNumber <= 10; blockNumber++ {
		blockTime, err := c.eventDB.RetrieveBlockTime(c.GetTestContext(), chainID, blockNumber)
		Nil(c.T(), err)
		timeDiff := abs(int64(blockTime) - int64(targetTime))
		if closestBlockTime > timeDiff {
			closestBlockTime = timeDiff
			closestBlockNumber = blockNumber
		}
	}
	Equal(c.T(), closestBlockNumber, blockNumber)
}

func abs(a int64) uint64 {
	if a < 0 {
		return uint64(-a)
	}
	return uint64(a)
}
