package base

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/agents/executor/types"
	agentsTypes "github.com/synapsecns/sanguine/agents/types"
	"gorm.io/gorm/clause"
)

// StoreMessage stores a message in the database.
func (s Store) StoreMessage(ctx context.Context, message agentsTypes.Message, root common.Hash, blockNumber uint64) error {
	dbMessage, err := AgentsTypesMessageToMessage(message, root, blockNumber)
	if err != nil {
		return fmt.Errorf("failed to convert message: %w", err)
	}

	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: ChainIDFieldName}, {Name: DestinationFieldName}, {Name: NonceFieldName}, {Name: RootFieldName},
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
func (s Store) GetMessage(ctx context.Context, messageMask types.DBMessage) (*agentsTypes.Message, error) {
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

	decodedMessage, err := agentsTypes.DecodeMessage(message.Message)
	if err != nil {
		return nil, fmt.Errorf("failed to decode message: %w", err)
	}

	return &decodedMessage, nil
}

// GetMessages gets messages from the database, paginated and ordered in ascending order by nonce.
func (s Store) GetMessages(ctx context.Context, messageMask types.DBMessage, page int) ([]agentsTypes.Message, error) {
	if page < 1 {
		page = 1
	}

	var messages []Message

	dbMessageMask := DBMessageToMessage(messageMask)

	dbTx := s.DB().WithContext(ctx).
		Model(&messages).
		Where(&dbMessageMask).
		Order(fmt.Sprintf("%s ASC", NonceFieldName)).
		Offset((page - 1) * PageSize).
		Limit(PageSize).
		Scan(&messages)

	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get messages: %w", dbTx.Error)
	}

	decodedMessages := make([]agentsTypes.Message, len(messages))
	for i, message := range messages {
		decodedMessage, err := agentsTypes.DecodeMessage(message.Message)
		if err != nil {
			return nil, fmt.Errorf("failed to decode message: %w", err)
		}
		decodedMessages[i] = decodedMessage
	}

	return decodedMessages, nil
}

// GetRoot gets the root of a message from the database.
func (s Store) GetRoot(ctx context.Context, messageMask types.DBMessage) (common.Hash, error) {
	var message Message

	dbMessageMask := DBMessageToMessage(messageMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&message).
		Where(&dbMessageMask).
		Scan(&message)
	if dbTx.Error != nil {
		return common.Hash{}, fmt.Errorf("failed to get message: %w", dbTx.Error)
	}

	return common.HexToHash(message.Root), nil
}

// GetBlockNumber gets the block number of a message from the database.
func (s Store) GetBlockNumber(ctx context.Context, messageMask types.DBMessage) (uint64, error) {
	var message Message

	dbMessageMask := DBMessageToMessage(messageMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&message).
		Where(&dbMessageMask).
		First(&message)
	if dbTx.Error != nil {
		return 0, fmt.Errorf("failed to get message: %w", dbTx.Error)
	}

	return message.BlockNumber, nil
}

// GetLastBlockNumber gets the last block number that had a message in the database.
func (s Store) GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error) {
	var message Message
	var lastBlockNumber uint64
	var numMessages int64

	// Get the total amount of messages stored in the database.
	dbTx := s.DB().WithContext(ctx).
		Model(&message).
		Where(&Message{ChainID: chainID}).
		Count(&numMessages)
	if dbTx.Error != nil {
		return 0, fmt.Errorf("failed to get number of messages: %w", dbTx.Error)
	}
	if numMessages == 0 {
		return 0, nil
	}

	dbTx = s.DB().WithContext(ctx).
		Model(&message).
		Where(fmt.Sprintf("%s = ?", ChainIDFieldName), chainID).
		Select(fmt.Sprintf("MAX(%s)", BlockNumberFieldName)).
		Find(&lastBlockNumber)
	if dbTx.Error != nil {
		return 0, fmt.Errorf("failed to get last block number: %w", dbTx.Error)
	}

	return lastBlockNumber, nil
}

// DBMessageToMessage converts a DBMessage to a Message.
func DBMessageToMessage(dbMessage types.DBMessage) Message {
	var message Message

	if dbMessage.ChainID != nil {
		message.ChainID = *dbMessage.ChainID
	}

	if dbMessage.Destination != nil {
		message.Destination = *dbMessage.Destination
	}

	if dbMessage.Nonce != nil {
		message.Nonce = *dbMessage.Nonce
	}

	if dbMessage.Root != nil {
		message.Root = dbMessage.Root.String()
	}

	if dbMessage.Message != nil {
		message.Message = *dbMessage.Message
	}

	if dbMessage.BlockNumber != nil {
		message.BlockNumber = *dbMessage.BlockNumber
	}

	return message
}

// MessageToDBMessage converts a Message to a DBMessage.
func MessageToDBMessage(message Message) types.DBMessage {
	chainID := message.ChainID
	destination := message.Destination
	nonce := message.Nonce
	root := common.HexToHash(message.Root)
	messageBytes := message.Message
	blockNumber := message.BlockNumber

	return types.DBMessage{
		ChainID:     &chainID,
		Destination: &destination,
		Nonce:       &nonce,
		Root:        &root,
		Message:     &messageBytes,
		BlockNumber: &blockNumber,
	}
}

// AgentsTypesMessageToMessage converts an agentsTypes.Message to a Message.
func AgentsTypesMessageToMessage(message agentsTypes.Message, root common.Hash, blockNumber uint64) (Message, error) {
	rawMessage, err := agentsTypes.EncodeMessage(message)
	if err != nil {
		return Message{}, fmt.Errorf("failed to encode message: %w", err)
	}
	return Message{
		ChainID:     message.OriginDomain(),
		Destination: message.DestinationDomain(),
		Nonce:       message.Nonce(),
		Root:        root.String(),
		Message:     rawMessage,
		BlockNumber: blockNumber,
	}, nil
}
