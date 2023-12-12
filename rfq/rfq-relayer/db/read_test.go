package db_test

import (
	"math/big"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/bindings"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db/model"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/utils"
)

func (t *DBSuite) TestGetOriginBridgeEvent() {
	t.RunOnAllDBs(func(testDB db.TestDB) {
		ctx := t.GetTestContext()
		chainID := uint32(42161)
		transactionIdStr := "0xfcc1f7f7cc74717594f51e4e4359462632ea2488f9cd9624721f0f0b19dddb75"
		var transactionId [32]byte
		copy(transactionId[:], common.FromHex(transactionIdStr))
		request := common.FromHex("0x000000000000000000000000000000000000000000000000000000000000a4b10000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000040000000000000000000000002e234dae75c793f67a35089c9d99245e1c58470b000000000000000000000000f62849f9a0b5bf2913b396098f7c7019b51a820a0000000000000000000000000000000000000000000000000000000000a7d8c00000000000000000000000000000000000000000000000000000000000a763900000000000000000000000000000000000000000000000000000000000000e110000000000000000000000000000000000000000000000000000000000000000")

		// Create a dummy log and event data to insert
		log := &types.Log{
			Address:     common.BigToAddress(big.NewInt(gofakeit.Int64())),
			Topics:      []common.Hash{common.HexToHash("0x456")},
			Data:        []byte{1, 2, 3},
			BlockNumber: 42,
			TxHash:      common.BigToHash(big.NewInt(gofakeit.Int64())),
			BlockHash:   common.BigToHash(big.NewInt(gofakeit.Int64())),
		}
		OriginBridgeEvent := bindings.FastBridgeBridgeRequested{
			TransactionId: transactionId,
			Sender:        common.HexToAddress("0x0000000000000000000000000000000000000004"),
			Request:       request,
			Raw:           *log,
		}

		// Store the bridge event
		err := testDB.StoreOriginBridgeEvent(ctx, chainID, log, &OriginBridgeEvent)
		Nil(t.T(), err)

		// Test retrieving the event using GetOriginBridgeEvent
		retrievedEvent, err := testDB.GetOriginBridgeEvent(t.GetTestContext(), transactionIdStr[2:])
		Nil(t.T(), err)
		NotNil(t.T(), retrievedEvent)

		// Validate retrieved events.
		Equal(t.T(), transactionIdStr[2:], retrievedEvent.TransactionID)
		Equal(t.T(), log.BlockNumber, retrievedEvent.BlockNumber)
		Equal(t.T(), log.TxHash.Hex(), retrievedEvent.TxHash)
		Equal(t.T(), log.TxIndex, retrievedEvent.TxIndex)
		Equal(t.T(), log.BlockHash.Hex(), retrievedEvent.BlockHash)
		Equal(t.T(), log.Index, retrievedEvent.LogIndex)
		Equal(t.T(), log.Removed, retrievedEvent.Removed)

		// Test Non-Existent Event
		_, err = testDB.GetOriginBridgeEvent(t.GetTestContext(), gofakeit.UUID())
		NotNil(t.T(), err) // Expect an error since the event should not exist
	})
}
func (t *DBSuite) TestGetDestinationBridgeEvent() {
	t.RunOnAllDBs(func(testDB db.TestDB) {
		ctx := t.GetTestContext()
		transactionIdStr := "0xfcc1f7f7cc74717594f51e4e4359462632ea2488f9cd9624721f0f0b19dddb75"

		// Create a dummy log and event data to insert
		log := &types.Log{
			Address:     common.BigToAddress(big.NewInt(gofakeit.Int64())),
			Topics:      []common.Hash{common.HexToHash("0x456")},
			Data:        []byte{1, 2, 3},
			BlockNumber: 42,
			TxHash:      common.BigToHash(big.NewInt(gofakeit.Int64())),
			BlockHash:   common.BigToHash(big.NewInt(gofakeit.Int64())),
		}
		originBridgeEvent := model.OriginBridgeEvent{
			TransactionID: transactionIdStr,
			Request:       "request1",
			OriginChainId: 42161,
			DestChainId:   1,
			BlockNumber:   42,
			TxHash:        log.TxHash.Hex(),
			TxIndex:       0,
			BlockHash:     log.BlockHash.Hex(),
			LogIndex:      0,
			Removed:       false,
		}

		// Store the event
		err := testDB.StoreDestinationBridgeEvent(ctx, log, &originBridgeEvent)
		Nil(t.T(), err)

		// Test retrieving the event
		retrievedEvent, err := testDB.GetDestinationBridgeEvent(ctx, transactionIdStr)
		Nil(t.T(), err)
		NotNil(t.T(), retrievedEvent)

		// Validate retrieved events.
		Equal(t.T(), transactionIdStr, retrievedEvent.TransactionID)
		Equal(t.T(), originBridgeEvent.BlockNumber, retrievedEvent.BlockNumber)
		Equal(t.T(), originBridgeEvent.TxHash, retrievedEvent.TxHash)
		Equal(t.T(), originBridgeEvent.TxIndex, retrievedEvent.TxIndex)
		Equal(t.T(), originBridgeEvent.BlockHash, retrievedEvent.BlockHash)
		Equal(t.T(), originBridgeEvent.LogIndex, retrievedEvent.LogIndex)
		Equal(t.T(), originBridgeEvent.Removed, retrievedEvent.Removed)

		// Test Non-Existent Event
		_, err = testDB.GetDestinationBridgeEvent(ctx, transactionIdStr+"1")
		NotNil(t.T(), err) // Expect an error since the event should not exist
	})
}

func (t *DBSuite) TestGetLastIndexed() {
	t.RunOnAllDBs(func(testDB db.TestDB) {
		address := common.BigToAddress(big.NewInt(gofakeit.Int64()))
		chainID := gofakeit.Uint32()
		blockNumber1 := uint64(gofakeit.Uint8())

		// Store the block using StoreLastIndexed
		err := testDB.StoreLastIndexed(t.GetTestContext(), address, chainID, blockNumber1)
		Nil(t.T(), err)

		// Retrieve the last stored block and validate
		retrievedBlock, err := testDB.GetLastIndexed(t.GetTestContext(), chainID, address)
		Nil(t.T(), err)
		Equal(t.T(), blockNumber1, retrievedBlock)

		// Update the block number and store it using StoreLastIndexed
		blockNumber2 := blockNumber1 + 1
		err = testDB.StoreLastIndexed(t.GetTestContext(), address, chainID, blockNumber2)
		Nil(t.T(), err)

		// Retrieve the updated block and validate
		retrievedUpdatedBlock, err := testDB.GetLastIndexed(t.GetTestContext(), chainID, address)
		Nil(t.T(), err)
		Equal(t.T(), blockNumber2, retrievedUpdatedBlock)

		// Ensure that lower block numbers will not be stored
		blockNumber3 := blockNumber1 - 1
		err = testDB.StoreLastIndexed(t.GetTestContext(), address, chainID, blockNumber3)
		Nil(t.T(), err)

		// Retrieve the updated block and validate
		retrievedNonUpdatedBlock, err := testDB.GetLastIndexed(t.GetTestContext(), chainID, address)
		Nil(t.T(), err)
		Equal(t.T(), blockNumber2, retrievedNonUpdatedBlock)
	})
}

func (t *DBSuite) TestGetToken() {
	t.RunOnAllDBs(func(testDB db.TestDB) {
		ctx := t.GetTestContext()
		chainID := gofakeit.Uint32()
		address := common.HexToAddress(big.NewInt(gofakeit.Int64()).String())
		tokenID := utils.GenerateTokenID(chainID, address)
		token := &model.Token{
			TokenID:  tokenID,
			Name:     gofakeit.Name(),
			Symbol:   gofakeit.LetterN(3),
			Decimals: 18,
			ChainID:  gofakeit.Uint32(),
		}

		// Store the token
		err := testDB.StoreToken(ctx, token)
		Nil(t.T(), err)

		// Test retrieving the event using GetToken
		storedToken, err := testDB.GetToken(t.GetTestContext(), tokenID)
		Nil(t.T(), err)
		NotNil(t.T(), storedToken)

		// Validate
		Equal(t.T(), token.TokenID, storedToken.TokenID)
		Equal(t.T(), token.Name, storedToken.Name)
		Equal(t.T(), token.Symbol, storedToken.Symbol)
		Equal(t.T(), token.Decimals, storedToken.Decimals)
		Equal(t.T(), token.ChainID, storedToken.ChainID)

		// Test Non-Existent Token
		_, err = testDB.GetOriginBridgeEvent(t.GetTestContext(), utils.GenerateTokenID(chainID+1, address))
		NotNil(t.T(), err) // Expect an error since the event should not exist
	})
}

func (t *DBSuite) TestGetDeadlineQueueEvents() {
	t.RunOnAllDBs(func(testDB db.TestDB) {
		ctx := t.GetTestContext()

		// Generate and store random number of events
		queueEventCount := int(gofakeit.Float32Range(3, 5))
		for i := 0; i < queueEventCount; i++ {
			event := &model.DeadlineQueue{
				Timestamp:     time.Now().Unix(),
				TransactionID: gofakeit.UUID(),
			}
			err := testDB.StoreDeadlineQueueEvent(ctx, event)
			Nil(t.T(), err)
		}

		// Test retrieving the events using GetDeadlineQueueEvents
		storedEvents, err := testDB.GetDeadlineQueueEvents(ctx)
		Nil(t.T(), err)

		// Validate the count of retrieved events
		Equal(t.T(), queueEventCount, len(storedEvents))
	})
}
