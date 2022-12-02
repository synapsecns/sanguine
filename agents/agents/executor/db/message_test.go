package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"github.com/synapsecns/sanguine/agents/agents/executor/db/datastore/sql/base"
	"github.com/synapsecns/sanguine/agents/agents/executor/types"
	"math/big"
)

func (t *DBSuite) TestStoreRetrieveMessage() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		chainIDA := gofakeit.Uint32()
		nonceA := gofakeit.Uint32()
		rootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		messageA := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		leafA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		blockNumberA := gofakeit.Uint64()
		dbMessageA := types.DBMessage{
			ChainID:     &chainIDA,
			Nonce:       &nonceA,
			Root:        &rootA,
			Message:     &messageA,
			Leaf:        &leafA,
			BlockNumber: &blockNumberA,
		}
		err := testDB.StoreMessage(t.GetTestContext(), dbMessageA)
		Nil(t.T(), err)

		chainIDB := gofakeit.Uint32()
		nonceB := gofakeit.Uint32()
		rootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		messageB := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		leafB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		blockNumberB := gofakeit.Uint64()
		dbMessageB := types.DBMessage{
			ChainID:     &chainIDB,
			Nonce:       &nonceB,
			Root:        &rootB,
			Message:     &messageB,
			Leaf:        &leafB,
			BlockNumber: &blockNumberB,
		}
		err = testDB.StoreMessage(t.GetTestContext(), dbMessageB)
		Nil(t.T(), err)

		messageAMask := types.DBMessage{
			ChainID:     dbMessageA.ChainID,
			Nonce:       dbMessageA.Nonce,
			BlockNumber: dbMessageA.BlockNumber,
		}
		retrievedMessageA, err := testDB.GetMessage(t.GetTestContext(), messageAMask)
		Nil(t.T(), err)
		Equal(t.T(), dbMessageA, *retrievedMessageA)

		messageBMask := types.DBMessage{
			Root:    dbMessageB.Root,
			Message: dbMessageB.Message,
			Leaf:    dbMessageB.Leaf,
		}
		retrievedMessageB, err := testDB.GetMessage(t.GetTestContext(), messageBMask)
		Nil(t.T(), err)
		Equal(t.T(), dbMessageB, *retrievedMessageB)
	})
}

func (t *DBSuite) TestGetLastBlockNumber() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		chainID := gofakeit.Uint32()
		nonceA := gofakeit.Uint32()
		nonceB := nonceA + 1
		rootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		rootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		messageA := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		messageB := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		leafA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		leafB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		blockNumberA := gofakeit.Uint64()
		blockNumberB := blockNumberA + 1

		dbMessageA := types.DBMessage{
			ChainID:     &chainID,
			Nonce:       &nonceA,
			Root:        &rootA,
			Message:     &messageA,
			Leaf:        &leafA,
			BlockNumber: &blockNumberA,
		}
		dbMessageB := types.DBMessage{
			ChainID:     &chainID,
			Nonce:       &nonceB,
			Root:        &rootB,
			Message:     &messageB,
			Leaf:        &leafB,
			BlockNumber: &blockNumberB,
		}

		err := testDB.StoreMessage(t.GetTestContext(), dbMessageA)
		Nil(t.T(), err)
		err = testDB.StoreMessage(t.GetTestContext(), dbMessageB)
		Nil(t.T(), err)

		lastBlockNumber, err := testDB.GetLastBlockNumber(t.GetTestContext(), chainID)
		Nil(t.T(), err)

		Equal(t.T(), blockNumberB, lastBlockNumber)
	})
}

func (t *DBSuite) TestMessageDBMessageParity() {
	chainID := gofakeit.Uint32()
	nonce := gofakeit.Uint32()
	root := common.BigToHash(big.NewInt(gofakeit.Int64()))
	message := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
	leaf := common.BigToHash(big.NewInt(gofakeit.Int64()))
	blockNumber := gofakeit.Uint64()
	initialDBMessage := types.DBMessage{
		ChainID:     &chainID,
		Nonce:       &nonce,
		Root:        &root,
		Message:     &message,
		Leaf:        &leaf,
		BlockNumber: &blockNumber,
	}

	initialMessage := base.DBMessageToMessage(initialDBMessage)

	finalDBMessage := base.MessageToDBMessage(initialMessage)

	finalMessage := base.DBMessageToMessage(finalDBMessage)

	Equal(t.T(), initialDBMessage, finalDBMessage)
	Equal(t.T(), initialMessage, finalMessage)
}
