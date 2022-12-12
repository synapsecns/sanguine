package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"github.com/synapsecns/sanguine/agents/agents/executor/types"
	agentsTypes "github.com/synapsecns/sanguine/agents/types"
	"math/big"
)

func (t *DBSuite) TestStoreRetrieveMessage() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		chainIDA := gofakeit.Uint32()
		destinationA := gofakeit.Uint32()
		nonceA := gofakeit.Uint32()
		rootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		messageA := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		blockNumberA := gofakeit.Uint64()

		headerA := agentsTypes.NewHeader(chainIDA, common.BigToHash(big.NewInt(gofakeit.Int64())), nonceA, destinationA, common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint32())
		tipsA := agentsTypes.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))
		typesMessageA := agentsTypes.NewMessage(headerA, tipsA, messageA)

		err := testDB.StoreMessage(t.GetTestContext(), typesMessageA, rootA, blockNumberA)
		Nil(t.T(), err)

		chainIDB := gofakeit.Uint32()
		destinationB := gofakeit.Uint32()
		nonceB := gofakeit.Uint32()
		rootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		messageB := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		blockNumberB := gofakeit.Uint64()

		headerB := agentsTypes.NewHeader(chainIDB, common.BigToHash(big.NewInt(gofakeit.Int64())), nonceB, destinationB, common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint32())
		tipsB := agentsTypes.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))
		typesMessageB := agentsTypes.NewMessage(headerB, tipsB, messageB)

		err = testDB.StoreMessage(t.GetTestContext(), typesMessageB, rootB, blockNumberB)
		Nil(t.T(), err)

		messageAMask := types.DBMessage{
			ChainID:     &chainIDA,
			Destination: &destinationA,
			Nonce:       &nonceA,
			BlockNumber: &blockNumberA,
		}
		retrievedMessageA, err := testDB.GetMessage(t.GetTestContext(), messageAMask)
		Nil(t.T(), err)

		encodeTypesMessageA, err := agentsTypes.EncodeMessage(typesMessageA)
		Nil(t.T(), err)
		encodeRetrievedMessageA, err := agentsTypes.EncodeMessage(*retrievedMessageA)
		Nil(t.T(), err)

		Equal(t.T(), encodeTypesMessageA, encodeRetrievedMessageA)

		messageBMask := types.DBMessage{
			Root: &rootB,
		}
		retrievedMessageB, err := testDB.GetMessage(t.GetTestContext(), messageBMask)
		Nil(t.T(), err)

		encodeTypesMessageB, err := agentsTypes.EncodeMessage(typesMessageB)
		Nil(t.T(), err)
		encodeRetrievedMessageB, err := agentsTypes.EncodeMessage(*retrievedMessageB)
		Nil(t.T(), err)

		Equal(t.T(), encodeTypesMessageB, encodeRetrievedMessageB)
	})
}

func (t *DBSuite) TestGetLastBlockNumber() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		chainID := gofakeit.Uint32()
		destinationA := gofakeit.Uint32()
		destinationB := destinationA + 1
		nonceA := gofakeit.Uint32()
		nonceB := nonceA + 1
		rootA := common.BigToHash(big.NewInt(gofakeit.Int64()))
		rootB := common.BigToHash(big.NewInt(gofakeit.Int64()))
		messageA := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		messageB := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		blockNumberA := gofakeit.Uint64()
		blockNumberB := blockNumberA + 1

		headerA := agentsTypes.NewHeader(chainID, common.BigToHash(big.NewInt(gofakeit.Int64())), nonceA, destinationA, common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint32())
		tipsA := agentsTypes.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))
		typesMessageA := agentsTypes.NewMessage(headerA, tipsA, messageA)

		err := testDB.StoreMessage(t.GetTestContext(), typesMessageA, rootA, blockNumberA)
		Nil(t.T(), err)

		headerB := agentsTypes.NewHeader(chainID, common.BigToHash(big.NewInt(gofakeit.Int64())), nonceB, destinationB, common.BigToHash(big.NewInt(gofakeit.Int64())), gofakeit.Uint32())
		tipsB := agentsTypes.NewTips(big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0))
		typesMessageB := agentsTypes.NewMessage(headerB, tipsB, messageB)

		err = testDB.StoreMessage(t.GetTestContext(), typesMessageB, rootB, blockNumberB)
		Nil(t.T(), err)

		lastBlockNumber, err := testDB.GetLastBlockNumber(t.GetTestContext(), chainID)
		Nil(t.T(), err)

		Equal(t.T(), blockNumberB, lastBlockNumber)
	})
}
