package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/testutil"
	"math/big"
)

func (t *DBSuite) TestFetchLogsInRange() {
	contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	chainID := gofakeit.Uint32()

	// Store 10 logs
	for blockNumber := 0; blockNumber < 10; blockNumber++ {
		storeLog := testutil.BuildLog(contractAddress, uint64(blockNumber), &t.logIndex)
		err := t.eventDB.StoreLog(t.GetTestContext(), storeLog, chainID)
		Nil(t.T(), err)
	}

	// Fetch logs from 4 to 8.
	fetcher := consumer.NewFetcher(t.gqlClient)
	logs, err := fetcher.FetchLogsInRange(t.GetTestContext(), chainID, 4, 8)
	Nil(t.T(), err)
	Equal(t.T(), 5, len(logs))
}

func (t *DBSuite) TestToken() {
	_, bridgeConfigContract := t.deployManager.GetBridgeConfigV3(t.GetTestContext(), t.testBackend)

	fetcher, err := bridgeconfig.NewFetcher(bridgeConfigContract.Address(), t.testBackend)
	Nil(t.T(), err)

	curentBlockNumber, err := t.testBackend.BlockNumber(t.GetTestContext())
	Nil(t.T(), err)

	for _, testToken := range testTokens {
		tokenID, err := fetcher.GetTokenID(t.GetTestContext(), uint32(testToken.ChainId.Uint64()), uint32(curentBlockNumber), common.HexToAddress(testToken.TokenAddress))

		Nil(t.T(), err)
		Equal(t.T(), *tokenID, testToken.tokenID)
		token, err := fetcher.GetToken(t.GetTestContext(), uint32(testToken.ChainId.Uint64()), uint32(curentBlockNumber), *tokenID)
		Nil(t.T(), err)
		Equal(t.T(), common.HexToAddress(testToken.TokenAddress).String(), common.HexToAddress(token.TokenAddress).String())
		Equal(t.T(), testToken.SwapFee, token.SwapFee)
		Equal(t.T(), testToken.IsUnderlying, token.IsUnderlying)
	}
}
