package db_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/scribe/db"
)

func (t *DBSuite) TestStoreRetrieveReceipt() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		txHashRandom := gofakeit.Int64()
		txHashA := common.BigToHash(big.NewInt(txHashRandom))
		txHashB := common.BigToHash(big.NewInt(txHashRandom + 1))
		randomLogs := []types.Log{
			*MakeRandomLog(txHashA),
			*MakeRandomLog(txHashA),
			*MakeRandomLog(txHashB),
			*MakeRandomLog(txHashB),
		}

		// Store all random logs, since `RetrieveReceiptByTxHash` needs to query them to build the Receipt.
		for _, log := range randomLogs {
			err := testDB.StoreLog(t.GetTestContext(), log, gofakeit.Uint32())
			Nil(t.T(), err)
		}

		// Store two receipts with different tx hashes.
		receiptA := types.Receipt{
			Type:              gofakeit.Uint8(),
			PostState:         []byte(gofakeit.Sentence(10)),
			Status:            gofakeit.Uint64(),
			CumulativeGasUsed: gofakeit.Uint64(),
			Bloom:             types.BytesToBloom([]byte(gofakeit.Sentence(10))),
			Logs: []*types.Log{
				&randomLogs[0],
				&randomLogs[1],
			},
			TxHash:           txHashA,
			GasUsed:          gofakeit.Uint64(),
			BlockNumber:      big.NewInt(int64(gofakeit.Uint32())),
			TransactionIndex: uint(gofakeit.Uint64()),
		}
		err := testDB.StoreReceipt(t.GetTestContext(), receiptA)
		Nil(t.T(), err)

		receiptB := types.Receipt{
			Type:              gofakeit.Uint8(),
			PostState:         []byte(gofakeit.Sentence(10)),
			Status:            gofakeit.Uint64(),
			CumulativeGasUsed: gofakeit.Uint64(),
			Bloom:             types.BytesToBloom([]byte(gofakeit.Sentence(10))),
			Logs: []*types.Log{
				&randomLogs[2],
				&randomLogs[3],
			},
			TxHash:           txHashB,
			GasUsed:          gofakeit.Uint64(),
			BlockNumber:      big.NewInt(int64(gofakeit.Uint32())),
			TransactionIndex: uint(gofakeit.Uint64()),
		}
		err = testDB.StoreReceipt(t.GetTestContext(), receiptB)
		Nil(t.T(), err)

		// Ensure the receipts from the database match the ones stored.
		retrievedReceiptA, err := testDB.RetrieveReceiptByTxHash(t.GetTestContext(), txHashA)
		Nil(t.T(), err)
		Equal(t.T(), receiptA.Type, retrievedReceiptA.Type)
		Equal(t.T(), receiptA.PostState, retrievedReceiptA.PostState)
		Equal(t.T(), receiptA.Status, retrievedReceiptA.Status)
		Equal(t.T(), receiptA.CumulativeGasUsed, retrievedReceiptA.CumulativeGasUsed)
		Equal(t.T(), receiptA.Bloom, retrievedReceiptA.Bloom)
		Equal(t.T(), receiptA.Logs, retrievedReceiptA.Logs)
		Equal(t.T(), receiptA.TxHash, retrievedReceiptA.TxHash)
		Equal(t.T(), receiptA.GasUsed, retrievedReceiptA.GasUsed)
		Equal(t.T(), receiptA.BlockNumber, retrievedReceiptA.BlockNumber)
		Equal(t.T(), receiptA.TransactionIndex, retrievedReceiptA.TransactionIndex)

		retrievedReceiptB, err := testDB.RetrieveReceiptByTxHash(t.GetTestContext(), txHashB)
		Nil(t.T(), err)
		Equal(t.T(), receiptB.Type, retrievedReceiptB.Type)
		Equal(t.T(), receiptB.PostState, retrievedReceiptB.PostState)
		Equal(t.T(), receiptB.Status, retrievedReceiptB.Status)
		Equal(t.T(), receiptB.CumulativeGasUsed, retrievedReceiptB.CumulativeGasUsed)
		Equal(t.T(), receiptB.Bloom, retrievedReceiptB.Bloom)
		Equal(t.T(), receiptB.Logs, retrievedReceiptB.Logs)
		Equal(t.T(), receiptB.TxHash, retrievedReceiptB.TxHash)
		Equal(t.T(), receiptB.GasUsed, retrievedReceiptB.GasUsed)
		Equal(t.T(), receiptB.BlockNumber, retrievedReceiptB.BlockNumber)
		Equal(t.T(), receiptB.TransactionIndex, retrievedReceiptB.TransactionIndex)
	})
}
