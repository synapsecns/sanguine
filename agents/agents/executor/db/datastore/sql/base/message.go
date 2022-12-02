package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/agents/executor/types"
	"gorm.io/gorm/clause"
)

// StoreMessage stores a message in the database.
func (s Store) StoreMessage(ctx context.Context, message types.DBMessage) error {
	dbMessage := DBMessageToMessage(message)
	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: ChainIDFieldName}, {Name: NonceFieldName}, {Name: RootFieldName},
			},
			DoNothing: true,
		}).
		Create(&dbMessage)

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store message: %w", dbTx.Error)
	}

	return nil
}

// GetMessage gets a message from the database.
func (s Store) GetMessage(ctx context.Context, messageMask types.DBMessage) (*types.DBMessage, error) {
	var message Message

	dbMessageMask := DBMessageToMessage(messageMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&message).
		Where(&dbMessageMask).
		Scan(&message)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get message: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		//nolint:nilnil
		return nil, nil
	}

	returnMessage := MessageToDBMessage(message)

	return &returnMessage, nil
}

// GetLastBlockNumber gets the last block number that had a message in the database.
func (s Store) GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error) {
	var message Message

	dbTx := s.DB().WithContext(ctx).
		Model(&message).
		Where(fmt.Sprintf("%s = ?", ChainIDFieldName), chainID).
		Order(fmt.Sprintf("%s DESC", BlockNumberFieldName)).
		Scan(&message)
	if dbTx.Error != nil {
		return 0, fmt.Errorf("failed to get last block number: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		return 0, nil
	}

	return message.BlockNumber, nil
}

// DBMessageToMessage converts a DBMessage to a Message.
func DBMessageToMessage(dbMessage types.DBMessage) Message {
	var message Message

	if dbMessage.ChainID != nil {
		message.ChainID = *dbMessage.ChainID
	}

	if dbMessage.Nonce != nil {
		message.Nonce = *dbMessage.Nonce
	}

	if dbMessage.Root != nil {
		message.Root = dbMessage.Root.String()
	}

	if dbMessage.Message != nil {
		message.Message = common.Bytes2Hex(*dbMessage.Message)
	}

	if dbMessage.Leaf != nil {
		message.Leaf = dbMessage.Leaf.String()
	}

	if dbMessage.BlockNumber != nil {
		message.BlockNumber = *dbMessage.BlockNumber
	}

	return message
}

// MessageToDBMessage converts a Message to a DBMessage.
func MessageToDBMessage(message Message) types.DBMessage {
	chainID := message.ChainID
	nonce := message.Nonce
	root := common.HexToHash(message.Root)
	messageBytes := common.HexToHash(message.Message).Bytes()
	leaf := common.HexToHash(message.Leaf)
	blockNumber := message.BlockNumber

	return types.DBMessage{
		ChainID:     &chainID,
		Nonce:       &nonce,
		Root:        &root,
		Message:     &messageBytes,
		Leaf:        &leaf,
		BlockNumber: &blockNumber,
	}
}
