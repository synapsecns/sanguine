package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"github.com/synapsecns/sanguine/services/sinner/db/model"
	graphqlModel "github.com/synapsecns/sanguine/services/sinner/graphql/server/graph/model"

	"github.com/synapsecns/sanguine/services/sinner/types"
	"math/big"
)

func (t *DBSuite) TestRetrieveLastStoredBlock() {
	t.RunOnAllDBs(func(testDB db.TestEventDB) {
		address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
		chainID := gofakeit.Uint32()
		blockNumber1 := gofakeit.Uint64()

		// Store the block using StoreLastIndexed
		err := testDB.StoreLastIndexed(t.GetTestContext(), address, chainID, blockNumber1)
		Nil(t.T(), err)

		// Retrieve the last stored block and validate
		retrievedBlock, err := testDB.RetrieveLastStoredBlock(t.GetTestContext(), chainID, address)
		Nil(t.T(), err)
		Equal(t.T(), blockNumber1, retrievedBlock)

		// Update the block number and store it using StoreLastIndexed
		blockNumber2 := blockNumber1 + 1
		err = testDB.StoreLastIndexed(t.GetTestContext(), address, chainID, blockNumber2)
		Nil(t.T(), err)

		// Retrieve the updated block and validate
		retrievedUpdatedBlock, err := testDB.RetrieveLastStoredBlock(t.GetTestContext(), chainID, address)
		Nil(t.T(), err)
		Equal(t.T(), blockNumber2, retrievedUpdatedBlock)
	})
}

func (t *DBSuite) TestRetrieveMessageStatus() {
	t.RunOnAllDBs(func(testDB db.TestEventDB) {
		messageHash1 := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
		originTxHash1 := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
		destinationTxHash1 := common.BigToHash(big.NewInt(gofakeit.Int64())).String()

		// Insert origin only record
		err := testDB.StoreOrUpdateMessageStatus(t.GetTestContext(), originTxHash1, messageHash1, types.Origin)

		Nil(t.T(), err)

		// Test: Retrieve and validate the origin only record
		retrievedStatus1, err := testDB.RetrieveMessageStatus(t.GetTestContext(), messageHash1)
		Nil(t.T(), err)
		Equal(t.T(), originTxHash1, *retrievedStatus1.OriginTxHash)
		Equal(t.T(), graphqlModel.MessageStateLastSeenOrigin, *retrievedStatus1.LastSeen)
		Equal(t.T(), "", *retrievedStatus1.DestinationTxHash)

		// Update the record to have both origin and destination
		err = testDB.StoreOrUpdateMessageStatus(t.GetTestContext(), destinationTxHash1, messageHash1, types.Destination)
		Nil(t.T(), err)

		// Test: Retrieve and validate the updated record
		retrievedStatus2, err := testDB.RetrieveMessageStatus(t.GetTestContext(), messageHash1)
		Nil(t.T(), err)
		Equal(t.T(), originTxHash1, *retrievedStatus2.OriginTxHash)
		Equal(t.T(), destinationTxHash1, *retrievedStatus2.DestinationTxHash)
		Equal(t.T(), graphqlModel.MessageStateLastSeenDestination, *retrievedStatus2.LastSeen)
	})
}

func (t *DBSuite) TestRetrieveOriginSent() {
	t.RunOnAllDBs(func(testDB db.TestEventDB) {
		// Setup: Insert some dummy data
		chainID := gofakeit.Uint32()
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
		messageHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()

		originSent := &model.OriginSent{
			ContractAddress: common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
			BlockNumber:     gofakeit.Uint64(),
			TxHash:          txHash,
			MessageHash:     messageHash,
			ChainID:         chainID,
		}

		err := testDB.StoreOriginSent(t.GetTestContext(), originSent)
		Nil(t.T(), err)

		filter := model.OriginSent{MessageHash: messageHash}
		retrievedRecord, err := testDB.RetrieveOriginSent(t.GetTestContext(), filter)
		Nil(t.T(), err)
		Equal(t.T(), txHash, retrievedRecord[0].TxHash)
		Equal(t.T(), chainID, retrievedRecord[0].ChainID)
		Equal(t.T(), messageHash, retrievedRecord[0].MessageHash)

		filter = model.OriginSent{ChainID: chainID, TxHash: txHash}
		retrievedRecords, err := testDB.RetrieveOriginSent(t.GetTestContext(), filter)
		Nil(t.T(), err)
		Equal(t.T(), txHash, retrievedRecords[0].TxHash)
		Equal(t.T(), chainID, retrievedRecords[0].ChainID)
		Equal(t.T(), messageHash, retrievedRecords[0].MessageHash)
	})
}

func (t *DBSuite) TestRetrieveExecuted() {
	t.RunOnAllDBs(func(testDB db.TestEventDB) {
		// Setup: Insert some dummy data
		chainID := gofakeit.Uint32()
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
		messageHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()

		executed := &model.Executed{
			ContractAddress: common.BigToAddress(big.NewInt(gofakeit.Int64())).String(),
			BlockNumber:     gofakeit.Uint64(),
			TxHash:          txHash,
			MessageHash:     messageHash,
			ChainID:         chainID,
			RemoteDomain:    gofakeit.Uint32(),
			Success:         gofakeit.Bool(),
		}

		err := testDB.StoreExecuted(t.GetTestContext(), executed)
		Nil(t.T(), err)

		filter := model.Executed{MessageHash: messageHash}
		retrievedRecord, err := testDB.RetrieveExecuted(t.GetTestContext(), filter)
		Nil(t.T(), err)
		Equal(t.T(), txHash, retrievedRecord[0].TxHash)
		Equal(t.T(), chainID, retrievedRecord[0].ChainID)
		Equal(t.T(), messageHash, retrievedRecord[0].MessageHash)

		filter = model.Executed{ChainID: chainID, TxHash: txHash}
		retrievedRecords, err := testDB.RetrieveExecuted(t.GetTestContext(), filter)
		Nil(t.T(), err)
		Equal(t.T(), txHash, retrievedRecords[0].TxHash)
		Equal(t.T(), chainID, retrievedRecords[0].ChainID)
		Equal(t.T(), messageHash, retrievedRecords[0].MessageHash)
	})
}
