package db_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/scribe/db"
)

func (t *DBSuite) TestPagination() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
		// Store 100 logs.
		for i := 101; i > 1; i-- {
			log := t.MakeRandomLog(txHash)
			log.BlockNumber = uint64(i)
			err := testDB.StoreLogs(t.GetTestContext(), 1, log)
			Nil(t.T(), err)
		}
		// Store another log that should be on the second page.
		log := t.MakeRandomLog(txHash)
		log.BlockNumber = 1
		err := testDB.StoreLogs(t.GetTestContext(), 1, log)
		Nil(t.T(), err)

		// Retrieve the log on the second page.
		retrievedLog, err := testDB.RetrieveLogsWithFilter(t.GetTestContext(), db.LogFilter{}, 2)
		Nil(t.T(), err)
		Equal(t.T(), 1, len(retrievedLog))

		// Check the value of the log.
		resA, err := log.MarshalJSON()
		Nil(t.T(), err)
		resB, err := retrievedLog[0].MarshalJSON()
		Nil(t.T(), err)
		Equal(t.T(), resA, resB)

		// Retrieve a receipt associated with all 101 logs.
		receipt := types.Receipt{
			Type:              gofakeit.Uint8(),
			PostState:         []byte(gofakeit.Sentence(10)),
			Status:            gofakeit.Uint64(),
			CumulativeGasUsed: gofakeit.Uint64(),
			Bloom:             types.BytesToBloom([]byte(gofakeit.Sentence(10))),
			TxHash:            txHash,
			ContractAddress:   common.BigToAddress(big.NewInt(gofakeit.Int64())),
			GasUsed:           gofakeit.Uint64(),
			BlockNumber:       big.NewInt(int64(gofakeit.Uint32())),
			TransactionIndex:  uint(gofakeit.Uint64()),
		}
		err = testDB.StoreReceipt(t.GetTestContext(), 1, receipt)
		Nil(t.T(), err)
		retrievedReceipt, err := testDB.RetrieveReceiptsWithFilter(t.GetTestContext(), db.ReceiptFilter{}, 1)
		Nil(t.T(), err)
		// Ensure the receipt has 101 logs.
		Equal(t.T(), 101, len(retrievedReceipt[0].Logs))
	})
}
