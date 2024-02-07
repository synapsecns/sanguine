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
		messageA := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		blockNumberA := gofakeit.Uint64()
		minimumTimeSetA := gofakeit.Bool()
		minimumTimeA := gofakeit.Uint64()

		headerA := agentsTypes.NewHeader(agentsTypes.MessageFlagManager, chainIDA, nonceA, destinationA, gofakeit.Uint32())
		typesMessageA := agentsTypes.NewMessage(headerA, nil, messageA)

		err := testDB.StoreMessage(t.GetTestContext(), typesMessageA, blockNumberA, minimumTimeSetA, minimumTimeA)
		Nil(t.T(), err)

		chainIDB := gofakeit.Uint32()
		destinationB := gofakeit.Uint32()
		nonceB := gofakeit.Uint32()
		messageB := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		blockNumberB := gofakeit.Uint64()
		minimumTimeSetB := gofakeit.Bool()
		minimumTimeB := gofakeit.Uint64()

		headerB := agentsTypes.NewHeader(agentsTypes.MessageFlagManager, chainIDB, nonceB, destinationB, gofakeit.Uint32())
		typesMessageB := agentsTypes.NewMessage(headerB, nil, messageB)

		err = testDB.StoreMessage(t.GetTestContext(), typesMessageB, blockNumberB, minimumTimeSetB, minimumTimeB)
		Nil(t.T(), err)

		messageAMask := db.DBMessage{
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

		messageBMask := db.DBMessage{
			Nonce:          &nonceB,
			MinimumTimeSet: &minimumTimeSetB,
			MinimumTime:    &minimumTimeB,
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
		destination := gofakeit.Uint32()
		nonce := gofakeit.Uint32()
		optimisticSeconds := gofakeit.Uint32()
		messageBody := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		snapshotRoot := common.BigToHash(big.NewInt(gofakeit.Int64()))
		stateRoot := common.BigToHash(big.NewInt(gofakeit.Int64()))
		dataHash := common.BigToHash(big.NewInt(gofakeit.Int64()))
		stateBlockNumber := big.NewInt(gofakeit.Int64())
		summitBlockNumber := big.NewInt(gofakeit.Int64())
		summitTimestamp := big.NewInt(gofakeit.Int64())
		stateTimestamp := big.NewInt(gofakeit.Int64())
		destinationTimestamp := gofakeit.Uint64()
		proof := [][]byte{common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()}
		stateIndex := gofakeit.Uint32()
		gasPrice := gofakeit.Uint16()
		dataPrice := gofakeit.Uint16()
		execBuffer := gofakeit.Uint16()
		amortAttCost := gofakeit.Uint16()
		etherPrice := gofakeit.Uint16()
		markup := gofakeit.Uint16()

		messageBlockNumber := gofakeit.Uint64()
		attestationDestinationBlockNumber := gofakeit.Uint64()
		stateSnapshotBlockNumber := gofakeit.Uint64()

		// Create and store a message.
		header := agentsTypes.NewHeader(agentsTypes.MessageFlagManager, chainID, nonce, destination, optimisticSeconds)
		message, err := agentsTypes.NewMessageFromManagerMessage(header, messageBody)
		Nil(t.T(), err)
		err = testDB.StoreMessage(t.GetTestContext(), message, messageBlockNumber, false, 0)
		Nil(t.T(), err)

		// Create and store an attestation.
		attestation := agentsTypes.NewAttestation(snapshotRoot, dataHash, nonce, summitBlockNumber, summitTimestamp)
		err = testDB.StoreAttestation(t.GetTestContext(), attestation, destination, attestationDestinationBlockNumber, destinationTimestamp)
		Nil(t.T(), err)

		// Create and store a state.
		gasData := agentsTypes.NewGasData(gasPrice, dataPrice, execBuffer, amortAttCost, etherPrice, markup)
		state := agentsTypes.NewState(stateRoot, chainID, nonce, stateBlockNumber, stateTimestamp, gasData)
		err = testDB.StoreState(t.GetTestContext(), state, snapshotRoot, proof, stateIndex, stateSnapshotBlockNumber)
		Nil(t.T(), err)

		// Get the last block number for each type.
		lastMessageBlockNumber, err := testDB.GetLastBlockNumber(t.GetTestContext(), chainID, types.OriginContract)
		Nil(t.T(), err)
		Equal(t.T(), messageBlockNumber, lastMessageBlockNumber)

		lastAttestationBlockNumber, err := testDB.GetLastBlockNumber(t.GetTestContext(), destination, types.LightInboxContract)
		Nil(t.T(), err)
		Equal(t.T(), attestationDestinationBlockNumber, lastAttestationBlockNumber)

		lastStateBlockNumber, err := testDB.GetLastBlockNumber(t.GetTestContext(), chainID, types.InboxContract)
		Nil(t.T(), err)
		Equal(t.T(), stateSnapshotBlockNumber, lastStateBlockNumber)
	})
}

func (t *DBSuite) TestExecuteMessage() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		chainID := gofakeit.Uint32()
		destination := gofakeit.Uint32()
		nonce := gofakeit.Uint32()
		message := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		blockNumber := gofakeit.Uint64()

		header := agentsTypes.NewHeader(agentsTypes.MessageFlagManager, chainID, nonce, destination, gofakeit.Uint32())
		typesMessage := agentsTypes.NewMessage(header, nil, message)

		err := testDB.StoreMessage(t.GetTestContext(), typesMessage, blockNumber, true, 5)
		Nil(t.T(), err)

		messageMask := db.DBMessage{
			ChainID: &chainID,
		}
		messages, err := testDB.GetExecutableMessages(t.GetTestContext(), messageMask, 10, 1)
		Nil(t.T(), err)

		Equal(t.T(), 1, len(messages))

		messageMask = db.DBMessage{
			ChainID:     &chainID,
			Destination: &destination,
			Nonce:       &nonce,
		}
		err = testDB.ExecuteMessage(t.GetTestContext(), messageMask)
		Nil(t.T(), err)

		messageMask = db.DBMessage{
			ChainID: &chainID,
		}
		messages, err = testDB.GetExecutableMessages(t.GetTestContext(), messageMask, 0, 1)
		Nil(t.T(), err)

		Equal(t.T(), 0, len(messages))
	})
}

func (t *DBSuite) TestGetExecutableMessages() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		chainID := gofakeit.Uint32()
		destination := gofakeit.Uint32()
		nonce := gofakeit.Uint32()
		message := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		blockNumber := gofakeit.Uint64()

		header := agentsTypes.NewHeader(agentsTypes.MessageFlagManager, chainID, nonce, destination, gofakeit.Uint32())
		typesMessage := agentsTypes.NewMessage(header, nil, message)

		err := testDB.StoreMessage(t.GetTestContext(), typesMessage, blockNumber, false, 10)
		Nil(t.T(), err)

		messageMask := db.DBMessage{
			ChainID: &chainID,
		}
		// Check when the current time is after the minimum time, but minimum time is set to false.
		messages, err := testDB.GetExecutableMessages(t.GetTestContext(), messageMask, 15, 1)
		Nil(t.T(), err)

		Equal(t.T(), 0, len(messages))

		chainID = gofakeit.Uint32()
		destination = gofakeit.Uint32()
		nonce = gofakeit.Uint32()
		message = common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		blockNumber = gofakeit.Uint64()

		header = agentsTypes.NewHeader(agentsTypes.MessageFlagManager, chainID, nonce, destination, gofakeit.Uint32())
		typesMessage = agentsTypes.NewMessage(header, nil, message)

		err = testDB.StoreMessage(t.GetTestContext(), typesMessage, blockNumber, true, 20)
		Nil(t.T(), err)

		// Check when the current time is after the minimum time, and minimum time is set to true.
		messages, err = testDB.GetExecutableMessages(t.GetTestContext(), messageMask, 25, 1)
		Nil(t.T(), err)

		Equal(t.T(), 1, len(messages))

		messageMask = db.DBMessage{
			ChainID:     &chainID,
			Destination: &destination,
			Nonce:       &nonce,
		}
		err = testDB.ExecuteMessage(t.GetTestContext(), messageMask)
		Nil(t.T(), err)

		messageMask = db.DBMessage{
			ChainID: &chainID,
		}
		// Check when a message has the correct current time, has its minimum time set, but has already been executed.
		messages, err = testDB.GetExecutableMessages(t.GetTestContext(), messageMask, 15, 1)
		Nil(t.T(), err)

		Equal(t.T(), 0, len(messages))
	})
}

func (t *DBSuite) TestGetUnsetMinimumTimeMessages() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		chainID := gofakeit.Uint32()
		destination := gofakeit.Uint32()
		nonce := gofakeit.Uint32()
		message := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		blockNumber := gofakeit.Uint64()

		header := agentsTypes.NewHeader(agentsTypes.MessageFlagManager, chainID, nonce, destination, gofakeit.Uint32())
		typesMessage := agentsTypes.NewMessage(header, nil, message)

		err := testDB.StoreMessage(t.GetTestContext(), typesMessage, blockNumber, false, 0)
		Nil(t.T(), err)

		messageMask := db.DBMessage{
			ChainID: &chainID,
		}
		messages, err := testDB.GetUnsetMinimumTimeMessages(t.GetTestContext(), messageMask, 1)
		Nil(t.T(), err)

		Equal(t.T(), 1, len(messages))

		destination = gofakeit.Uint32()
		nonce = gofakeit.Uint32()
		message = common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		blockNumber = gofakeit.Uint64()

		header = agentsTypes.NewHeader(agentsTypes.MessageFlagManager, chainID, nonce, destination, gofakeit.Uint32())
		typesMessage = agentsTypes.NewMessage(header, nil, message)

		err = testDB.StoreMessage(t.GetTestContext(), typesMessage, blockNumber, true, 0)
		Nil(t.T(), err)

		messages, err = testDB.GetUnsetMinimumTimeMessages(t.GetTestContext(), messageMask, 1)
		Nil(t.T(), err)

		Equal(t.T(), 1, len(messages))
	})
}

func (t *DBSuite) TestSetMinimumTime() {
	t.RunOnAllDBs(func(testDB db.ExecutorDB) {
		chainID := gofakeit.Uint32()
		destination := gofakeit.Uint32()
		nonce := gofakeit.Uint32()
		message := common.BigToHash(big.NewInt(gofakeit.Int64())).Bytes()
		blockNumber := gofakeit.Uint64()

		header := agentsTypes.NewHeader(agentsTypes.MessageFlagManager, chainID, nonce, destination, gofakeit.Uint32())
		typesMessage := agentsTypes.NewMessage(header, nil, message)

		err := testDB.StoreMessage(t.GetTestContext(), typesMessage, blockNumber, false, 0)
		Nil(t.T(), err)

		trueVal := true
		messageMask := db.DBMessage{
			ChainID:        &chainID,
			MinimumTimeSet: &trueVal,
		}

		messages, err := testDB.GetMessages(t.GetTestContext(), messageMask, 1)
		Nil(t.T(), err)

		Equal(t.T(), 0, len(messages))

		messageMask = db.DBMessage{
			ChainID: &chainID,
		}

		err = testDB.SetMinimumTime(t.GetTestContext(), messageMask, 10)
		Nil(t.T(), err)

		messageMask = db.DBMessage{
			ChainID:        &chainID,
			MinimumTimeSet: &trueVal,
		}

		messages, err = testDB.GetMessages(t.GetTestContext(), messageMask, 1)
		Nil(t.T(), err)

		Equal(t.T(), 1, len(messages))

		time, err := testDB.GetMessageMinimumTime(t.GetTestContext(), messageMask)
		Nil(t.T(), err)

		Equal(t.T(), uint64(10), *time)
	})
}
