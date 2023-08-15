package db_test

import (
	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/scribe/db"
)

func (t *DBSuite) TestStoreRetrieveReceipt() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		txHashRandom := gofakeit.Int64()
		chainID := gofakeit.Uint32()
		txHashA := common.BigToHash(big.NewInt(txHashRandom))
		txHashB := common.BigToHash(big.NewInt(txHashRandom + 1))
		randomLogsA := []types.Log{
			t.MakeRandomLog(txHashA),
			t.MakeRandomLog(txHashA),
		}
		randomLogsA[0].BlockNumber = 4
		randomLogsA[1].BlockNumber = 3
		randomLogsB := []types.Log{
			t.MakeRandomLog(txHashB),
			t.MakeRandomLog(txHashB),
		}
		randomLogsB[0].BlockNumber = 2
		randomLogsB[1].BlockNumber = 1

		// Store all random logs, since `RetrieveReceipt` needs to query them to build the Receipt.
		for _, log := range randomLogsA {
			err := testDB.StoreLogs(t.GetTestContext(), chainID, log)
			Nil(t.T(), err)
		}
		for _, log := range randomLogsB {
			err := testDB.StoreLogs(t.GetTestContext(), chainID+1, log)
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
				&randomLogsA[0],
				&randomLogsA[1],
			},
			TxHash:           txHashA,
			ContractAddress:  common.BigToAddress(big.NewInt(gofakeit.Int64())),
			GasUsed:          gofakeit.Uint64(),
			BlockNumber:      big.NewInt(1),
			TransactionIndex: uint(gofakeit.Uint64()),
		}
		err := testDB.StoreReceipt(t.GetTestContext(), chainID, receiptA)
		Nil(t.T(), err)

		receiptB := types.Receipt{
			Type:              gofakeit.Uint8(),
			PostState:         []byte(gofakeit.Sentence(10)),
			Status:            gofakeit.Uint64(),
			CumulativeGasUsed: gofakeit.Uint64(),
			Bloom:             types.BytesToBloom([]byte(gofakeit.Sentence(10))),
			Logs: []*types.Log{
				&randomLogsB[0],
				&randomLogsB[1],
			},
			TxHash:           txHashB,
			ContractAddress:  common.BigToAddress(big.NewInt(gofakeit.Int64())),
			GasUsed:          gofakeit.Uint64(),
			BlockNumber:      big.NewInt(2),
			TransactionIndex: uint(gofakeit.Uint64()),
		}
		err = testDB.StoreReceipt(t.GetTestContext(), chainID+1, receiptB)
		Nil(t.T(), err)

		// Ensure the receipts from the database match the ones stored.
		receiptFilter := db.ReceiptFilter{
			TxHash:  txHashA.String(),
			ChainID: chainID,
		}
		retrievedReceiptA, err := testDB.RetrieveReceiptsWithFilter(t.GetTestContext(), receiptFilter, 1)
		Nil(t.T(), err)

		resA, err := receiptA.MarshalJSON()
		Nil(t.T(), err)
		resB, err := retrievedReceiptA[0].MarshalJSON()
		Nil(t.T(), err)
		Equal(t.T(), resA, resB)

		receiptFilter = db.ReceiptFilter{
			TxHash:  txHashB.String(),
			ChainID: chainID + 1,
		}
		retrievedReceiptB, err := testDB.RetrieveReceiptsWithFilter(t.GetTestContext(), receiptFilter, 1)
		Nil(t.T(), err)

		resA, err = receiptB.MarshalJSON()
		Nil(t.T(), err)
		resB, err = retrievedReceiptB[0].MarshalJSON()
		Nil(t.T(), err)
		Equal(t.T(), resA, resB)

		// Ensure RetrieveAllReceipts gets all receipts.
		allReceipts, err := testDB.RetrieveReceiptsWithFilter(t.GetTestContext(), db.ReceiptFilter{}, 1)
		Nil(t.T(), err)
		Equal(t.T(), 2, len(allReceipts))
	})
}

func (t *DBSuite) TestDeleteReceiptsForBlockHash() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		chainID := gofakeit.Uint32()

		// Store a receipt.
		receipt := t.MakeRandomReceipt(common.BigToHash(big.NewInt(gofakeit.Int64())))
		receipt.BlockHash = common.BigToHash(big.NewInt(5))
		err := testDB.StoreReceipt(t.GetTestContext(), chainID, receipt)
		Nil(t.T(), err)

		// Ensure the receipt is in the database.
		receiptFilter := db.ReceiptFilter{
			ChainID:   chainID,
			BlockHash: receipt.BlockHash.String(),
		}
		retrievedReceipts, err := testDB.RetrieveReceiptsWithFilter(t.GetTestContext(), receiptFilter, 1)
		Nil(t.T(), err)
		Equal(t.T(), 1, len(retrievedReceipts))

		// Delete the receipt.
		err = testDB.DeleteReceiptsForBlockHash(t.GetTestContext(), chainID, receipt.BlockHash)
		Nil(t.T(), err)

		// Ensure the receipt is not in the database.
		retrievedReceipts, err = testDB.RetrieveReceiptsWithFilter(t.GetTestContext(), receiptFilter, 1)
		Nil(t.T(), err)
		Equal(t.T(), 0, len(retrievedReceipts))
	})
}

func (t *DBSuite) MakeRandomReceipt(txHash common.Hash) types.Receipt {
	return types.Receipt{
		Type:              gofakeit.Uint8(),
		PostState:         []byte(gofakeit.Sentence(10)),
		Status:            gofakeit.Uint64(),
		CumulativeGasUsed: gofakeit.Uint64(),
		Bloom:             types.BytesToBloom([]byte(gofakeit.Sentence(10))),
		Logs:              []*types.Log{},
		TxHash:            txHash,
		ContractAddress:   common.BigToAddress(big.NewInt(gofakeit.Int64())),
		GasUsed:           gofakeit.Uint64(),
		BlockNumber:       big.NewInt(int64(gofakeit.Uint32())),
		TransactionIndex:  uint(gofakeit.Uint64()),
	}
}

func (t *DBSuite) TestRetrieveReceiptsWithStaleBlockHash() {
	t.RunOnAllDBs(func(testDB db.EventDB) {
		chainID := gofakeit.Uint32()
		startBlock := 2
		hash1 := common.BigToHash(big.NewInt(1))
		hash2 := common.BigToHash(big.NewInt(2))
		hash3 := common.BigToHash(big.NewInt(3))
		newHash2 := common.BigToHash(big.NewInt(22))
		blockHashes := []common.Hash{hash1, hash2, hash3}
		blockHashesStr := []string{hash1.String(), newHash2.String(), hash3.String()}

		for i := 0; i < 10; i++ {
			receipt := t.MakeRandomReceipt(common.BigToHash(big.NewInt(gofakeit.Int64())))
			receipt.BlockNumber = big.NewInt(int64(i))
			receipt.BlockHash = blockHashes[i%3]
			err := testDB.StoreReceipt(t.GetTestContext(), chainID, receipt)
			Nil(t.T(), err)
		}

		// Retrieve the receipts
		retrievedReceipts, err := testDB.RetrieveReceiptsWithStaleBlockHash(t.GetTestContext(), chainID, blockHashesStr, uint64(startBlock), 10)
		Nil(t.T(), err)

		// Ensure the correct receipts (all the ones with hash 2) were retrieved
		Equal(t.T(), 2, len(retrievedReceipts))
		Equal(t.T(), hash2, retrievedReceipts[0].BlockHash)
	})
}
