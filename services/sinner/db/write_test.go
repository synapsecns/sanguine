package db_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"github.com/synapsecns/sanguine/services/sinner/db/model"
	"github.com/synapsecns/sanguine/services/sinner/types"
	"math/big"
)

func (t *DBSuite) TestStoreRetrieveExecuted() {
	t.RunOnAllDBs(func(testDB db.TestEventDB) {
		// Randomly generate values for the fields of the Executed type
		contractAddressRandom := common.BigToHash(big.NewInt(gofakeit.Int64()))
		blockNumber := gofakeit.Uint64()
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
		txIndex := uint(gofakeit.Uint32())
		messageHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
		chainID := gofakeit.Uint32()
		remoteDomain := gofakeit.Uint32()
		success := gofakeit.Bool()

		executedEvent := &model.Executed{
			ContractAddress: contractAddressRandom.String(),
			BlockNumber:     blockNumber,
			TxHash:          txHash.String(),
			TxIndex:         txIndex,
			MessageHash:     messageHash.String(),
			ChainID:         chainID,
			RemoteDomain:    remoteDomain,
			Success:         success,
		}

		// Store the executed event
		err := testDB.StoreExecuted(t.GetTestContext(), executedEvent)
		Nil(t.T(), err)

		// Retrieve the stored executed event and validate it
		var retrievedEvent model.Executed
		err = testDB.UNSAFE_DB().WithContext(t.GetTestContext()).Find(&model.Executed{}).First(&retrievedEvent).Error
		Nil(t.T(), err)

		// Check if the retrieved event matches the stored event
		Equal(t.T(), executedEvent.ContractAddress, retrievedEvent.ContractAddress)
		Equal(t.T(), executedEvent.BlockNumber, retrievedEvent.BlockNumber)
		Equal(t.T(), executedEvent.TxHash, retrievedEvent.TxHash)
		Equal(t.T(), executedEvent.TxIndex, retrievedEvent.TxIndex)
		Equal(t.T(), executedEvent.MessageHash, retrievedEvent.MessageHash)
		Equal(t.T(), executedEvent.ChainID, retrievedEvent.ChainID)
		Equal(t.T(), executedEvent.RemoteDomain, retrievedEvent.RemoteDomain)
		Equal(t.T(), executedEvent.Success, retrievedEvent.Success)
	})
}

func (t *DBSuite) TestStoreRetrieveOriginSent() {
	t.RunOnAllDBs(func(testDB db.TestEventDB) {
		// Randomly generate values for the fields of the OriginSent type
		contractAddressRandom := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
		blockNumber := gofakeit.Uint64()
		txHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
		txIndex := uint(gofakeit.Uint32())
		sender := gofakeit.Address().Address
		recipient := gofakeit.Address().Address
		messageLeaf := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
		messageID := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
		messageHash := common.BigToHash(big.NewInt(gofakeit.Int64())).String()
		chainID := gofakeit.Uint32()
		destinationChainID := gofakeit.Uint32()
		nonce := gofakeit.Uint32()
		message := gofakeit.Sentence(10)
		optimisticSeconds := gofakeit.Uint32()
		messageFlag := uint8(gofakeit.Uint32())
		summitTip := gofakeit.Word()
		attestationTip := gofakeit.Word()
		executionTip := gofakeit.Word()
		deliveryTip := gofakeit.Word()
		version := gofakeit.Uint32()
		gasLimit := gofakeit.Uint64()
		gasDrop := gofakeit.Word()

		originSentEvent := &model.OriginSent{
			ContractAddress:    contractAddressRandom,
			BlockNumber:        blockNumber,
			TxHash:             txHash,
			TxIndex:            txIndex,
			Sender:             sender,
			Recipient:          recipient,
			MessageLeaf:        messageLeaf,
			MessageID:          messageID,
			MessageHash:        messageHash,
			ChainID:            chainID,
			DestinationChainID: destinationChainID,
			Nonce:              nonce,
			Message:            message,
			OptimisticSeconds:  optimisticSeconds,
			MessageFlag:        messageFlag,
			SummitTip:          summitTip,
			AttestationTip:     attestationTip,
			ExecutionTip:       executionTip,
			DeliveryTip:        deliveryTip,
			Version:            version,
			GasLimit:           gasLimit,
			GasDrop:            gasDrop,
		}

		// Store the originSent event
		err := testDB.StoreOriginSent(t.GetTestContext(), originSentEvent)
		Nil(t.T(), err)

		// Retrieve the stored originSent event and validate it
		var retrievedEvent model.OriginSent
		err = testDB.UNSAFE_DB().WithContext(t.GetTestContext()).Find(&model.OriginSent{}).First(&retrievedEvent).Error
		Nil(t.T(), err)

		// Check if the retrieved event matches the stored event
		Equal(t.T(), originSentEvent.ContractAddress, retrievedEvent.ContractAddress)
		Equal(t.T(), originSentEvent.BlockNumber, retrievedEvent.BlockNumber)
		Equal(t.T(), originSentEvent.TxHash, retrievedEvent.TxHash)
	})
}

func (t *DBSuite) TestStoreLastIndexed() {
	t.RunOnAllDBs(func(testDB db.TestEventDB) {
		// Randomly generate values for the fields of the LastIndexedInfo type
		contractAddressRandom := common.HexToAddress(big.NewInt(gofakeit.Int64()).String())
		chainID := gofakeit.Uint32()
		blockNumber := gofakeit.Uint64()

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

func (t *DBSuite) TestStoreOrUpdateMessageStatus() {
	t.RunOnAllDBs(func(testDB db.TestEventDB) {
		// Test Case 1: Create message status record with Origin
		messageID1 := common.HexToAddress(big.NewInt(gofakeit.Int64()).String()).String()
		txHash1 := common.HexToHash(big.NewInt(gofakeit.Int64()).String()).String()
		err := testDB.StoreOrUpdateMessageStatus(t.GetTestContext(), txHash1, messageID1, types.Origin)
		Nil(t.T(), err)

		var status1 model.MessageStatus
		err = testDB.UNSAFE_DB().WithContext(t.GetTestContext()).Where("message_hash = ?", messageID1).First(&status1).Error
		Nil(t.T(), err)
		Equal(t.T(), txHash1, status1.OriginTxHash)

		// Test Case 2: Create message status record with Destination
		messageID2 := common.HexToAddress(big.NewInt(gofakeit.Int64()).String()).String()
		txHash2 := common.HexToHash(big.NewInt(gofakeit.Int64()).String()).String()
		err = testDB.StoreOrUpdateMessageStatus(t.GetTestContext(), txHash2, messageID2, types.Destination)
		Nil(t.T(), err)

		var status2 model.MessageStatus
		err = testDB.UNSAFE_DB().WithContext(t.GetTestContext()).Where("message_hash = ?", messageID2).First(&status2).Error
		Nil(t.T(), err)
		Equal(t.T(), txHash2, status2.DestinationTxHash)

		// Test Case 3: Update message status record with Origin
		newTxHash1 := common.HexToHash(big.NewInt(gofakeit.Int64()).String()).String()
		err = testDB.StoreOrUpdateMessageStatus(t.GetTestContext(), newTxHash1, messageID1, types.Origin)
		Nil(t.T(), err)

		var updatedStatus1 model.MessageStatus
		err = testDB.UNSAFE_DB().WithContext(t.GetTestContext()).Where("message_hash = ?", messageID1).First(&updatedStatus1).Error
		Nil(t.T(), err)
		Equal(t.T(), newTxHash1, updatedStatus1.OriginTxHash)

		// Test Case 4: Update message status record with Destination
		newTxHash2 := common.HexToHash(big.NewInt(gofakeit.Int64()).String()).String()
		err = testDB.StoreOrUpdateMessageStatus(t.GetTestContext(), newTxHash2, messageID2, types.Destination)
		Nil(t.T(), err)

		var updatedStatus2 model.MessageStatus
		err = testDB.UNSAFE_DB().WithContext(t.GetTestContext()).Where("message_hash = ?", messageID2).First(&updatedStatus2).Error
		Nil(t.T(), err)
		Equal(t.T(), newTxHash2, updatedStatus2.DestinationTxHash)
	})
}
