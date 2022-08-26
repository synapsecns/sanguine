package db_test

import (
	"math/big"

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
		logA := MakeRandomLog(txHashA)
		err := testDB.StoreLog(t.GetTestContext(), *logA, chainID)
		Nil(t.T(), err)

		logB := MakeRandomLog(txHashA)
		err = testDB.StoreLog(t.GetTestContext(), *logB, chainID)
		Nil(t.T(), err)

		txHashC := common.BigToHash(big.NewInt(txHashRandom + 1))
		logC := MakeRandomLog(txHashC)
		err = testDB.StoreLog(t.GetTestContext(), *logC, chainID+1)
		Nil(t.T(), err)

		// Ensure the logs from the database match the ones stored.
		// Check the logs for the two with the same txHash.
		retrievedLogSame, err := testDB.RetrieveLogs(t.GetTestContext(), txHashA, chainID)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLogSame[0].Address, logA.Address)
		Equal(t.T(), retrievedLogSame[0].Topics, logA.Topics)
		Equal(t.T(), retrievedLogSame[0].Data, logA.Data)
		Equal(t.T(), retrievedLogSame[0].BlockNumber, logA.BlockNumber)
		Equal(t.T(), retrievedLogSame[0].TxHash, logA.TxHash)
		Equal(t.T(), retrievedLogSame[0].TxIndex, logA.TxIndex)
		Equal(t.T(), retrievedLogSame[0].BlockHash, logA.BlockHash)
		Equal(t.T(), retrievedLogSame[0].Index, logA.Index)
		Equal(t.T(), retrievedLogSame[0].Removed, logA.Removed)

		Equal(t.T(), retrievedLogSame[1].Address, logB.Address)
		Equal(t.T(), retrievedLogSame[1].Topics, logB.Topics)
		Equal(t.T(), retrievedLogSame[1].Data, logB.Data)
		Equal(t.T(), retrievedLogSame[1].BlockNumber, logB.BlockNumber)
		Equal(t.T(), retrievedLogSame[1].TxHash, logB.TxHash)
		Equal(t.T(), retrievedLogSame[1].TxIndex, logB.TxIndex)
		Equal(t.T(), retrievedLogSame[1].BlockHash, logB.BlockHash)
		Equal(t.T(), retrievedLogSame[1].Index, logB.Index)
		Equal(t.T(), retrievedLogSame[1].Removed, logB.Removed)

		// Check the logs for the one with a different txHash.
		retrievedLog, err := testDB.RetrieveLogs(t.GetTestContext(), txHashC, chainID+1)
		Nil(t.T(), err)
		Equal(t.T(), retrievedLog[0].Address, logC.Address)
		Equal(t.T(), retrievedLog[0].Topics, logC.Topics)
		Equal(t.T(), retrievedLog[0].Data, logC.Data)
		Equal(t.T(), retrievedLog[0].BlockNumber, logC.BlockNumber)
		Equal(t.T(), retrievedLog[0].TxHash, logC.TxHash)
		Equal(t.T(), retrievedLog[0].TxIndex, logC.TxIndex)
		Equal(t.T(), retrievedLog[0].BlockHash, logC.BlockHash)
		Equal(t.T(), retrievedLog[0].Index, logC.Index)
		Equal(t.T(), retrievedLog[0].Removed, logC.Removed)

		// Check if `RetrieveAllLogs` returns all the logs.
		allLogs, err := testDB.RetrieveAllLogs_Test(t.GetTestContext())
		Nil(t.T(), err)
		Equal(t.T(), len(allLogs), 3)
	})
}

func MakeRandomLog(txHash common.Hash) *types.Log {
	return &types.Log{
		Address:     common.BigToAddress(big.NewInt(gofakeit.Int64())),
		Topics:      []common.Hash{common.BigToHash(big.NewInt(gofakeit.Int64())), common.BigToHash(big.NewInt(gofakeit.Int64())), common.BigToHash(big.NewInt(gofakeit.Int64()))},
		Data:        []byte(gofakeit.Sentence(10)),
		BlockNumber: gofakeit.Uint64(),
		TxHash:      txHash,
		TxIndex:     uint(gofakeit.Uint64()),
		BlockHash:   common.BigToHash(big.NewInt(gofakeit.Int64())),
		Index:       uint(gofakeit.Uint64()),
		Removed:     gofakeit.Bool(),
	}
}
