package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/explorer/db/consumer"
	"math/big"
)

func (t *DBSuite) TestFetchLogsInRange() {
	contractAddress := common.BigToAddress(big.NewInt(gofakeit.Int64()))
	chainID := gofakeit.Uint32()

	// Store 10 logs
	for blockNumber := 0; blockNumber < 10; blockNumber++ {
		storeLog := t.buildLog(contractAddress, uint64(blockNumber))
		err := t.eventDB.StoreLog(t.GetTestContext(), storeLog, chainID)
		Nil(t.T(), err)
	}

	// Fetch logs from 4 to 8.
	fetcher := consumer.NewFetcher(t.gqlClient)
	logs, err := fetcher.FetchLogsInRange(t.GetTestContext(), chainID, 4, 8)
	Nil(t.T(), err)
	Equal(t.T(), 5, len(logs))
}

func (t *DBSuite) buildLog(contractAddress common.Address, blockNumber uint64) types.Log {
	currentIndex := t.logIndex.Load()
	// increment next index
	t.logIndex.Add(1)
	log := types.Log{
		Address:     contractAddress,
		Topics:      []common.Hash{common.BigToHash(big.NewInt(gofakeit.Int64())), common.BigToHash(big.NewInt(gofakeit.Int64())), common.BigToHash(big.NewInt(gofakeit.Int64()))},
		Data:        []byte(gofakeit.Sentence(10)),
		BlockNumber: blockNumber,
		TxHash:      common.BigToHash(big.NewInt(gofakeit.Int64())),
		TxIndex:     uint(gofakeit.Uint64()),
		BlockHash:   common.BigToHash(big.NewInt(gofakeit.Int64())),
		Index:       uint(currentIndex),
		Removed:     gofakeit.Bool(),
	}

	return log
}
