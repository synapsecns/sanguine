package db_test

import (
	"math/big"

	"github.com/synapsecns/sanguine/services/scribe/db"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
)

func (t *DBSuite) TestStoreRetrieveLog() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		txHashRandom := gofakeit.Int64()
		chainID := gofakeit.Uint32()

		// Store two logs with the same txHash, and one with a different txHash.
		txHashA := common.BigToHash(big.NewInt(txHashRandom))
		logA := t.MakeRandomLog(txHashA)
		logA.BlockNumber = 3
		err := testDB.StoreLogs(t.GetTestContext(), chainID, logA)
		Nil(t.T(), err)

		logB := t.MakeRandomLog(txHashA)
		logB.BlockNumber = 2
		err = testDB.StoreLogs(t.GetTestContext(), chainID, logB)

		Nil(t.T(), err)

		txHashC := common.BigToHash(big.NewInt(txHashRandom + 1))
		logC := t.MakeRandomLog(txHashC)
		logC.BlockNumber = 1
		err = testDB.StoreLogs(t.GetTestContext(), chainID+1, logC)

		Nil(t.T(), err)

		// Ensure the logs from the database match the ones stored.
		// Check the logs for the two with the same txHash.
		txHashFilter := db.LogFilter{
			TxHash:  txHashA.String(),
			ChainID: chainID,
		}
		retrievedLogSame, err := testDB.RetrieveLogsWithFilter(t.GetTestContext(), txHashFilter, 1)
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
		txHashFilter = db.LogFilter{
			TxHash:  txHashC.String(),
			ChainID: chainID + 1,
		}
		retrievedLog, err := testDB.RetrieveLogsWithFilter(t.GetTestContext(), txHashFilter, 1)
		Nil(t.T(), err)

		resA, err = logC.MarshalJSON()
		Nil(t.T(), err)
		resB, err = retrievedLog[0].MarshalJSON()
		Nil(t.T(), err)
		Equal(t.T(), resA, resB)
	})
}

func (t *DBSuite) TestDeleteLogsForBlockHash() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		chainID := gofakeit.Uint32()

		// Store a log.
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
		log := t.MakeRandomLog(txHash)
		log.BlockHash = common.BigToHash(big.NewInt(5))
		err := testDB.StoreLogs(t.GetTestContext(), chainID, log)
		Nil(t.T(), err)

		// Ensure the log is in the database.
		logFilter := db.LogFilter{
			ChainID:   chainID,
			BlockHash: log.BlockHash.String(),
		}
		retrievedLogs, err := testDB.RetrieveLogsWithFilter(t.GetTestContext(), logFilter, 1)
		Nil(t.T(), err)
		Equal(t.T(), 1, len(retrievedLogs))

		// Delete the log.
		err = testDB.DeleteLogsForBlockHash(t.GetTestContext(), log.BlockHash, chainID)
		Nil(t.T(), err)

		// Make sure the log is not in the database.
		retrievedLogs, err = testDB.RetrieveLogsWithFilter(t.GetTestContext(), logFilter, 1)
		Nil(t.T(), err)
		Equal(t.T(), 0, len(retrievedLogs))
	})
}

func (t *DBSuite) TestLogCount() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		chainID := gofakeit.Uint32()
		contractAddressA := common.BigToAddress(big.NewInt(gofakeit.Int64()))
		contractAddressB := common.BigToAddress(big.NewInt(gofakeit.Int64()))

		// create and store logs, receipts, and txs
		var log types.Log

		var err error
		for blockNumber := 0; blockNumber < 5; blockNumber++ {
			// create and store logs
			log = t.buildLog(contractAddressA, uint64(blockNumber))
			err = testDB.StoreLogs(t.GetTestContext(), chainID, log)
			Nil(t.T(), err)
			log = t.buildLog(contractAddressB, uint64(blockNumber))
			err = testDB.StoreLogs(t.GetTestContext(), chainID, log)
			Nil(t.T(), err)
		}

		// test get logs and get logs with testDB
		logCountA, err := testDB.RetrieveLogCountForContract(t.GetTestContext(), contractAddressA, chainID)
		Nil(t.T(), err)
		Equal(t.T(), int64(5), logCountA)
		// store last indexed
		logCountB, err := testDB.RetrieveLogCountForContract(t.GetTestContext(), contractAddressB, chainID)
		Nil(t.T(), err)
		Equal(t.T(), int64(5), logCountB)
	})
}
func (t *DBSuite) MakeRandomLog(txHash common.Hash) types.Log {
	currentIndex := t.logIndex.Load()
	// increment next index
	t.logIndex.Add(1)
	return types.Log{
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
