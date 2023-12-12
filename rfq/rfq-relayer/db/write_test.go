package db_test

import (
	"encoding/hex"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/utils"

	"math/big"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/bindings"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db/model"
)

func (t *DBSuite) TestStoreOriginBridgeEvent() {
	t.RunOnAllDBs(func(testDB db.TestDB) {
		ctx := t.GetTestContext()
		chainID := uint32(42161)
		var transactionID [32]byte
		copy(transactionID[:], common.FromHex("0xfcc1f7f7cc74717594f51e4e4359462632ea2488f9cd9624721f0f0b19dddb75"))
		request := common.FromHex("0x000000000000000000000000000000000000000000000000000000000000a4b10000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000040000000000000000000000002e234dae75c793f67a35089c9d99245e1c58470b000000000000000000000000f62849f9a0b5bf2913b396098f7c7019b51a820a0000000000000000000000000000000000000000000000000000000000a7d8c00000000000000000000000000000000000000000000000000000000000a763900000000000000000000000000000000000000000000000000000000000000e110000000000000000000000000000000000000000000000000000000000000000")

		// Create a dummy log and event data to insert
		log := &types.Log{
			Address:     common.HexToAddress("0x123"),
			Topics:      []common.Hash{common.HexToHash("0x456")},
			Data:        []byte{1, 2, 3},
			BlockNumber: 42,
			TxHash:      common.HexToHash("0x789"),
			BlockHash:   common.HexToHash("0xabc"),
		}
		OriginBridgeEvent := bindings.FastBridgeBridgeRequested{
			TransactionId: transactionID,
			Sender:        common.HexToAddress("0x0000000000000000000000000000000000000004"),
			Request:       request,
			Raw:           *log,
		}

		// Store the bridge event
		err := testDB.StoreOriginBridgeEvent(ctx, chainID, log, &OriginBridgeEvent)
		Nil(t.T(), err)

		// Retrieve the stored event and validate it
		var retrievedEvent model.OriginBridgeEvent
		err = testDB.UNSAFE_DB().WithContext(ctx).
			Where(&model.OriginBridgeEvent{TransactionID: hex.EncodeToString(OriginBridgeEvent.TransactionId[:])}).
			First(&retrievedEvent).Error
		Nil(t.T(), err)

		// Validate retrieved data
		Equal(t.T(), log.BlockNumber, retrievedEvent.BlockNumber)
		Equal(t.T(), log.TxHash.Hex(), retrievedEvent.TxHash)
		Equal(t.T(), log.TxIndex, retrievedEvent.TxIndex)
		Equal(t.T(), log.BlockHash.Hex(), retrievedEvent.BlockHash)
		Equal(t.T(), log.Index, retrievedEvent.LogIndex)
		Equal(t.T(), log.Removed, retrievedEvent.Removed)

		// Test chainID != bridgeTransaction.OriginChainID case
		wrongChainID := chainID + 1
		err = testDB.StoreOriginBridgeEvent(ctx, wrongChainID, log, &OriginBridgeEvent)
		NotNil(t.T(), err)

		// Test reinserting an event with the same id
		err = testDB.StoreOriginBridgeEvent(ctx, chainID, log, &OriginBridgeEvent)
		Nil(t.T(), err)

		// Retrieve the stored event and validate it
		var retrievedEvents []model.OriginBridgeEvent
		err = testDB.UNSAFE_DB().WithContext(ctx).
			Where(&model.OriginBridgeEvent{TransactionID: hex.EncodeToString(OriginBridgeEvent.TransactionId[:])}).
			Find(&retrievedEvents).Error
		Nil(t.T(), err)
		Equal(t.T(), 1, len(retrievedEvents)) // Should still only be one event with this ID
	})
}

func (t *DBSuite) TestStoreDestinationBridgeEvent() {
	t.RunOnAllDBs(func(testDB db.TestDB) {
		ctx := t.GetTestContext()

		// Create a dummy log and event data to insert
		log := &types.Log{
			BlockNumber: 100,
			TxHash:      common.HexToHash("0x123"),
			TxIndex:     1,
			BlockHash:   common.HexToHash("0x456"),
			Index:       2,
			Removed:     false,
		}
		originBridgeEvent := &model.OriginBridgeEvent{
			TransactionID: "transactionId1",
			Request:       "request1",
			OriginChainID: 42161,
			DestChainID:   1,
			BlockNumber:   42,
			TxHash:        "0x789",
			TxIndex:       0,
			BlockHash:     "0xabc",
			LogIndex:      0,
			Removed:       false,
		}

		// Store the bridge event
		err := testDB.StoreDestinationBridgeEvent(ctx, log, originBridgeEvent)
		Nil(t.T(), err)

		// Retrieve the stored event and validate it
		var retrievedEvent model.DestinationBridgeEvent
		err = testDB.UNSAFE_DB().WithContext(ctx).
			Where(&model.DestinationBridgeEvent{TransactionID: originBridgeEvent.TransactionID}).
			First(&retrievedEvent).Error
		Nil(t.T(), err)

		// Validate retrieved data
		Equal(t.T(), originBridgeEvent.TransactionID, retrievedEvent.TransactionID)
		Equal(t.T(), originBridgeEvent.Request, retrievedEvent.Request)
		Equal(t.T(), originBridgeEvent.OriginChainID, retrievedEvent.OriginChainID)
		Equal(t.T(), originBridgeEvent.DestChainID, retrievedEvent.DestChainID)
		Equal(t.T(), log.BlockNumber, retrievedEvent.BlockNumber)
		Equal(t.T(), log.TxHash.Hex(), retrievedEvent.TxHash)
		Equal(t.T(), log.TxIndex, retrievedEvent.TxIndex)
		Equal(t.T(), log.BlockHash.Hex(), retrievedEvent.BlockHash)
		Equal(t.T(), log.Index, retrievedEvent.LogIndex)
		Equal(t.T(), log.Removed, retrievedEvent.Removed)

		// Test reinserting an event with the same id
		err = testDB.StoreDestinationBridgeEvent(ctx, log, originBridgeEvent)
		Nil(t.T(), err)

		// Retrieve the stored event and validate it
		var retrievedEvents []model.DestinationBridgeEvent
		err = testDB.UNSAFE_DB().WithContext(ctx).
			Where(&model.DestinationBridgeEvent{TransactionID: originBridgeEvent.TransactionID}).
			Find(&retrievedEvents).Error
		Nil(t.T(), err)
		Equal(t.T(), 1, len(retrievedEvents)) // Should still only be one event with this ID
	})
}

func (t *DBSuite) TestStoreLastIndexed() {
	t.RunOnAllDBs(func(testDB db.TestDB) {
		// Randomly generate values for the fields of the LastIndexedInfo type
		contractAddressRandom := common.HexToAddress(big.NewInt(gofakeit.Int64()).String())
		chainID := gofakeit.Uint32()
		blockNumber := uint64(gofakeit.Uint8())

		// Store the last indexed info
		err := testDB.StoreLastIndexed(t.GetTestContext(), contractAddressRandom, chainID, blockNumber)
		Nil(t.T(), err)

		// Retrieve the stored last indexed info and validate it
		var retrievedInfo model.LastIndexed
		err = testDB.UNSAFE_DB().WithContext(t.GetTestContext()).
			Where(&model.LastIndexed{
				ContractAddress: contractAddressRandom.String(),
				ChainID:         chainID,
			}).First(&retrievedInfo).Error
		Nil(t.T(), err)

		// Check if the retrieved info matches the stored info
		Equal(t.T(), contractAddressRandom.String(), retrievedInfo.ContractAddress)
		Equal(t.T(), chainID, retrievedInfo.ChainID)
		Equal(t.T(), blockNumber, retrievedInfo.BlockNumber)

		// Test updating the block number
		newBlockNumber := blockNumber + 1
		err = testDB.StoreLastIndexed(t.GetTestContext(), contractAddressRandom, chainID, newBlockNumber)
		Nil(t.T(), err)

		// Retrieve the updated info and validate it
		err = testDB.UNSAFE_DB().WithContext(t.GetTestContext()).
			Where(&model.LastIndexed{
				ContractAddress: contractAddressRandom.String(),
				ChainID:         chainID,
			}).First(&retrievedInfo).Error
		Nil(t.T(), err)

		// Check if the retrieved info matches the updated info
		Equal(t.T(), newBlockNumber, retrievedInfo.BlockNumber)
	})
}

func (t *DBSuite) TestStoreToken() {
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
			Address:  address.String(),
		}
		err := testDB.StoreToken(ctx, token)
		Nil(t.T(), err)

		// Verify the token is stored
		var storedToken model.Token
		err = testDB.UNSAFE_DB().WithContext(ctx).
			Where(&model.Token{TokenID: token.TokenID}).
			First(&storedToken).Error
		Nil(t.T(), err)

		// Validate
		Equal(t.T(), token.TokenID, storedToken.TokenID)
		Equal(t.T(), token.Name, storedToken.Name)
		Equal(t.T(), token.Symbol, storedToken.Symbol)
		Equal(t.T(), token.Decimals, storedToken.Decimals)
		Equal(t.T(), token.ChainID, storedToken.ChainID)

		// Test handling of token update
		newName := token.Name + "_new"
		newSymbol := token.Symbol + "_new"
		duplicateToken := &model.Token{
			TokenID:  token.TokenID,
			Name:     newName,
			Symbol:   newSymbol,
			Decimals: token.Decimals,
		}
		err = testDB.StoreToken(ctx, duplicateToken)
		Nil(t.T(), err)

		// Verify that no duplicate token was created
		var tokens []model.Token
		err = testDB.UNSAFE_DB().WithContext(ctx).
			Where(&model.Token{TokenID: token.TokenID}).
			Find(&tokens).Error
		Nil(t.T(), err)
		Equal(t.T(), 1, len(tokens)) // Should still only be one token with this ID

		// Verify that the token was updated
		var updatedToken model.Token
		err = testDB.UNSAFE_DB().WithContext(ctx).
			Where(&model.Token{TokenID: token.TokenID}).
			First(&updatedToken).Error
		Nil(t.T(), err)
		Equal(t.T(), newName, updatedToken.Name)
		Equal(t.T(), newSymbol, updatedToken.Symbol)
		Equal(t.T(), token.Decimals, updatedToken.Decimals)
	})
}

func (t *DBSuite) TestStoreAndRemoveDeadlineQueueEvent() {
	t.RunOnAllDBs(func(testDB db.TestDB) {
		ctx := t.GetTestContext()
		newEntry := &model.DeadlineQueue{
			TransactionID: gofakeit.UUID(),
			Timestamp:     time.Now().Unix(),
		}
		err := testDB.StoreDeadlineQueueEvent(ctx, newEntry)
		Nil(t.T(), err)

		// Verify the event is stored
		var storedEvent model.DeadlineQueue
		err = testDB.UNSAFE_DB().WithContext(ctx).
			Where(&model.DeadlineQueue{TransactionID: newEntry.TransactionID}).
			First(&storedEvent).Error
		Nil(t.T(), err)

		// Validate stored event data
		Equal(t.T(), newEntry.TransactionID, storedEvent.TransactionID)
		Equal(t.T(), newEntry.Timestamp, storedEvent.Timestamp)

		// Test handling of duplicate events
		err = testDB.StoreDeadlineQueueEvent(ctx, newEntry)
		Nil(t.T(), err)
		var events []model.DeadlineQueue
		err = testDB.UNSAFE_DB().WithContext(ctx).
			Where(&model.DeadlineQueue{TransactionID: newEntry.TransactionID}).
			Find(&events).Error
		Nil(t.T(), err)
		Equal(t.T(), 1, len(events)) // Should still only be one event with this ID

		// Test removing the event
		err = testDB.RemoveDeadlineQueueEvent(ctx, newEntry.TransactionID)
		Nil(t.T(), err)

		// Verify the event is removed
		var noEvents []model.DeadlineQueue
		err = testDB.UNSAFE_DB().WithContext(ctx).
			Where(&model.DeadlineQueue{TransactionID: newEntry.TransactionID}).
			Find(&noEvents).Error
		Nil(t.T(), err)
		Equal(t.T(), 0, len(noEvents)) // No event should be found with this ID
	})
}
