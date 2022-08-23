package db_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/scribe/db"
)

func (t *DBSuite) TestStoreRetrieveReceipt() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		txHashRandom := gofakeit.Int64()
		txHashA := common.BigToHash(big.NewInt(txHashRandom))

		// Store two receipts with different tx hashes.
		receiptA := ethTypes.Receipt{
			Type:              gofakeit.Uint8(),
			PostState:         []byte(gofakeit.Sentence(10)),
			Status:            gofakeit.Uint64(),
			CumulativeGasUsed: gofakeit.Uint64(),
			Bloom:             ethTypes.BytesToBloom([]byte(gofakeit.Sentence(10))),
			Logs: []*ethTypes.Log{
				MakeRandomLog(txHashA),
				MakeRandomLog(txHashA),
			},
			TxHash:           txHashA,
			GasUsed:          gofakeit.Uint64(),
			BlockNumber:      big.NewInt(gofakeit.Int64()),
			TransactionIndex: uint(gofakeit.Uint64()),
		}
		err := testDB.StoreReceipt(t.GetTestContext(), receiptA)
		Nil(t.T(), err)

		txHashB := common.BigToHash(big.NewInt(txHashRandom + 1))
		receiptB := ethTypes.Receipt{
			Type:              gofakeit.Uint8(),
			PostState:         []byte(gofakeit.Sentence(10)),
			Status:            gofakeit.Uint64(),
			CumulativeGasUsed: gofakeit.Uint64(),
			Bloom:             ethTypes.BytesToBloom([]byte(gofakeit.Sentence(10))),
			Logs: []*ethTypes.Log{
				MakeRandomLog(txHashB),
				MakeRandomLog(txHashB),
			},
			TxHash:           txHashB,
			GasUsed:          gofakeit.Uint64(),
			BlockNumber:      big.NewInt(gofakeit.Int64()),
			TransactionIndex: uint(gofakeit.Uint64()),
		}
		err = testDB.StoreReceipt(t.GetTestContext(), receiptB)
		Nil(t.T(), err)

		// Ensure the receipts from the database match the ones stored.
		retrievedReceiptA, err := testDB.RetrieveReceiptByTxHash(t.GetTestContext(), txHashA)
		Nil(t.T(), err)
		Equal(t.T(), retrievedReceiptA.Status(), receiptA.Status)
		Equal(t.T(), retrievedReceiptA.CumulativeGasUsed(), receiptA.CumulativeGasUsed)
		Equal(t.T(), retrievedReceiptA.TxHash(), receiptA.TxHash)
		Equal(t.T(), retrievedReceiptA.GasUsed(), receiptA.GasUsed)
		Equal(t.T(), retrievedReceiptA.BlockHash(), receiptA.BlockHash)
		Equal(t.T(), retrievedReceiptA.BlockNumber(), receiptA.BlockNumber.Uint64())
		Equal(t.T(), uint(retrievedReceiptA.TransactionIndex()), receiptA.TransactionIndex)

		retrievedReceiptB, err := testDB.RetrieveReceiptByTxHash(t.GetTestContext(), txHashB)
		Nil(t.T(), err)
		Equal(t.T(), retrievedReceiptB.Status(), receiptB.Status)
		Equal(t.T(), retrievedReceiptB.CumulativeGasUsed(), receiptB.CumulativeGasUsed)
		Equal(t.T(), retrievedReceiptB.TxHash(), receiptB.TxHash)
		Equal(t.T(), retrievedReceiptB.GasUsed(), receiptB.GasUsed)
		Equal(t.T(), retrievedReceiptB.BlockHash(), receiptB.BlockHash)
		Equal(t.T(), retrievedReceiptB.BlockNumber(), receiptB.BlockNumber.Uint64())
		Equal(t.T(), uint(retrievedReceiptB.TransactionIndex()), receiptB.TransactionIndex)
	})
}
