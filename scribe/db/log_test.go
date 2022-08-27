package db_test

import (
	"math/big"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/scribe/db"
)

func (t *DBSuite) TestStoreRetrieveLog() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		txHashRandom := gofakeit.Int64()
		chainID := gofakeit.Uint32()

		// Store two logs with the same txHash, and one with a different txHash.
		txHashA := common.BigToHash(big.NewInt(txHashRandom))
		logA := t.MakeRandomLog(txHashA)
		err := testDB.StoreLog(t.GetTestContext(), *logA, chainID)
		Nil(t.T(), err)

		logB := t.MakeRandomLog(txHashA)
		err = testDB.StoreLog(t.GetTestContext(), *logB, chainID)
		Nil(t.T(), err)

		txHashC := common.BigToHash(big.NewInt(txHashRandom + 1))
		logC := t.MakeRandomLog(txHashC)
		err = testDB.StoreLog(t.GetTestContext(), *logC, chainID+1)
		Nil(t.T(), err)
		time.Sleep(1 * time.Second)

		// Ensure the logs from the database match the ones stored.
		// Check the logs for the two with the same txHash.
		retrievedLogSame, err := testDB.RetrieveLogs(t.GetTestContext(), txHashA, chainID)
		Nil(t.T(), err)

		resA, err := logA.MarshalJSON()
		Nil(t.T(), err)
		resB, err := retrievedLogSame[0].MarshalJSON()
		Nil(t.T(), err)
		Equal(t.T(), resA, resB)

		resA, err = logB.MarshalJSON()
		Nil(t.T(), err)
		resB, err = retrievedLogSame[1].MarshalJSON()
		Nil(t.T(), err)
		Equal(t.T(), resA, resB)

		// Check the logs for the one with a different txHash.
		retrievedLog, err := testDB.RetrieveLogs(t.GetTestContext(), txHashC, chainID+1)
		Nil(t.T(), err)

		resA, err = logC.MarshalJSON()
		Nil(t.T(), err)
		resB, err = retrievedLog[0].MarshalJSON()
		Nil(t.T(), err)
		Equal(t.T(), resA, resB)

		// Check if `RetrieveAllLogs` returns all the logs.
		allLogs, err := testDB.RetrieveAllLogs_Test(t.GetTestContext())
		Nil(t.T(), err)
		Equal(t.T(), len(allLogs), 3)
	})
}

func (t *DBSuite) MakeRandomLog(txHash common.Hash) *types.Log {
	currentIndex := t.logIndex.Load()
	// increment next index
	t.logIndex.Add(1)
	return &types.Log{
		Address:     common.BigToAddress(big.NewInt(gofakeit.Int64())),
		Topics:      []common.Hash{common.BigToHash(big.NewInt(gofakeit.Int64())), common.BigToHash(big.NewInt(gofakeit.Int64())), common.BigToHash(big.NewInt(gofakeit.Int64()))},
		Data:        []byte(gofakeit.Sentence(10)),
		BlockNumber: gofakeit.Uint64(),
		TxHash:      txHash,
		TxIndex:     uint(gofakeit.Uint64()),
		BlockHash:   common.BigToHash(big.NewInt(gofakeit.Int64())),
		Index:       uint(currentIndex),
		Removed:     gofakeit.Bool(),
	}
}
