package base

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/synapsecns/sanguine/agents/agents/executor/db"
	"github.com/synapsecns/sanguine/agents/agents/executor/types"
	agentsTypes "github.com/synapsecns/sanguine/agents/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// StoreMessage stores a message in the database.
func (s Store) StoreMessage(ctx context.Context, message agentsTypes.Message, blockNumber uint64, minimumTimeSet bool, minimumTime uint64) error {
	dbMessage, err := AgentsTypesMessageToMessage(message, blockNumber, minimumTimeSet, minimumTime)
	if err != nil {
		return fmt.Errorf("failed to convert message: %w", err)
	}

	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: ChainIDFieldName}, {Name: DestinationFieldName}, {Name: NonceFieldName},
			},
			DoNothing: true,
		}).
		Create(&dbMessage)

	if dbTx.Error != nil {
		return fmt.Errorf("failed to store message: %w", dbTx.Error)
	}
	return nil
}

// ExecuteMessage marks a message as executed in the database.
// TODO: Make batch update.
func (s Store) ExecuteMessage(ctx context.Context, messageMask db.DBMessage) error {
	dbMessageMask := DBMessageToMessage(messageMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&Message{}).
		Where(&dbMessageMask).
		Update(ExecutedFieldName, true)
	if dbTx.Error != nil {
		return fmt.Errorf("failed to execute message: %w", dbTx.Error)
	}

	return nil
}

// SetMinimumTime sets the minimum time of a message.
func (s Store) SetMinimumTime(ctx context.Context, messageMask db.DBMessage, minimumTime uint64) error {
	dbMessageMask := DBMessageToMessage(messageMask)
	update := Message{MinimumTime: minimumTime, MinimumTimeSet: true}
	dbTx := s.DB().WithContext(ctx).
		Model(&Message{}).
		Where(&dbMessageMask).
		Updates(update)
	if dbTx.Error != nil {
		return fmt.Errorf("failed to set minimum time: %w", dbTx.Error)
	}

	return nil
}

// GetMessage gets a message from the database.
func (s Store) GetMessage(ctx context.Context, messageMask db.DBMessage) (*agentsTypes.Message, error) {
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
func (s Store) GetMessages(ctx context.Context, messageMask db.DBMessage, page int) ([]agentsTypes.Message, error) {
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

// GetBlockNumber gets the block number of a message from the database.
func (s Store) GetBlockNumber(ctx context.Context, messageMask db.DBMessage) (uint64, error) {
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
func (s Store) GetLastBlockNumber(ctx context.Context, chainID uint32, contractType types.ContractType) (uint64, error) {
	var lastBlockNumber sql.NullInt64

	preDBTx := s.DB().WithContext(ctx)
	var dbTx *gorm.DB

	//nolint:exhaustive
	switch contractType {
	case types.OriginContract:
		dbTx = preDBTx.Model(&Message{}).
			Where(fmt.Sprintf("%s = ?", ChainIDFieldName), chainID).
			Select(fmt.Sprintf("MAX(%s)", BlockNumberFieldName)).
			Find(&lastBlockNumber)
	case types.LightInboxContract, types.SummitContract:
		dbTx = preDBTx.Model(&Attestation{}).
			Where(fmt.Sprintf("%s = ?", DestinationFieldName), chainID).
			Select(fmt.Sprintf("MAX(%s)", DestinationBlockNumberFieldName)).
			Find(&lastBlockNumber)
	case types.InboxContract:
		// note: this makes the assumption there is one summit contract. If these are switched between chains without a state copy
		// you may receive erroneous results from this function
		dbTx = preDBTx.Model(&State{}).
			Select(fmt.Sprintf("MAX(%s)", BlockNumberFieldName)).
			Find(&lastBlockNumber)
	default:
		return 0, fmt.Errorf("unknown contract type: %v", contractType)
	}

	if dbTx.Error != nil {
		return 0, fmt.Errorf("failed to get last block number: %w", dbTx.Error)
	}

	// No messages are in the database yet.
	if !lastBlockNumber.Valid {
		return 0, nil
	}

	return uint64(lastBlockNumber.Int64), nil
}

// GetExecutableMessages gets executable messages from the database.
func (s Store) GetExecutableMessages(ctx context.Context, messageMask db.DBMessage, currentTime uint64, page int) ([]agentsTypes.Message, error) {
	if page < 1 {
		page = 1
	}

	var messages []Message

	dbMessageMask := DBMessageToMessage(messageMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&messages).
		Where(&dbMessageMask).
		Where(fmt.Sprintf("%s < ?", MinimumTimeFieldName), currentTime).
		Where(fmt.Sprintf("%s = ?", MinimumTimeSetFieldName), true).
		Where(fmt.Sprintf("%s = ?", ExecutedFieldName), false).
		Order(fmt.Sprintf("%s ASC", MinimumTimeFieldName)).
		Limit(PageSize).
		Offset((page - 1) * PageSize).
		Scan(&messages)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get messages: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		//nolint:nilnil
		return nil, nil
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

// GetUnsetMinimumTimeMessages gets messages from the database that have not had their minimum time set.
func (s Store) GetUnsetMinimumTimeMessages(ctx context.Context, messageMask db.DBMessage, page int) ([]agentsTypes.Message, error) {
	if page < 1 {
		page = 1
	}

	var messages []Message

	dbMessageMask := DBMessageToMessage(messageMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&messages).
		Where(&dbMessageMask).
		Where(fmt.Sprintf("%s = ?", MinimumTimeSetFieldName), false).
		Order(fmt.Sprintf("%s ASC", NonceFieldName)).
		Offset((page - 1) * PageSize).
		Limit(PageSize).
		Scan(&messages)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get messages: %w", dbTx.Error)
	}
	if dbTx.RowsAffected == 0 {
		//nolint:nilnil
		return nil, nil
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

// GetMessageMinimumTime gets the minimum time for a message to be executed.
func (s Store) GetMessageMinimumTime(ctx context.Context, messageMask db.DBMessage) (*uint64, error) {
	var message Message

	dbMessageMask := DBMessageToMessage(messageMask)
	dbTx := s.DB().WithContext(ctx).
		Model(&message).
		Where(&dbMessageMask).
		Scan(&message)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get message: %w", dbTx.Error)
	}
	if dbTx.RowsAffected > 1 {
		return nil, fmt.Errorf("multiple messages found with the same mask")
	}
	if !message.MinimumTimeSet || dbTx.RowsAffected != 1 {
		//nolint:nilnil
		return nil, nil
	}

	return &message.MinimumTime, nil
}

// DBMessageToMessage converts a DBMessage to a Message.
func DBMessageToMessage(dbMessage db.DBMessage) Message {
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

	if dbMessage.Message != nil {
		message.Message = *dbMessage.Message
	}

	if dbMessage.BlockNumber != nil {
		message.BlockNumber = *dbMessage.BlockNumber
	}

	if dbMessage.Executed != nil {
		message.Executed = *dbMessage.Executed
	}

	if dbMessage.MinimumTimeSet != nil {
		message.MinimumTimeSet = *dbMessage.MinimumTimeSet
	}

	if dbMessage.MinimumTime != nil {
		message.MinimumTime = *dbMessage.MinimumTime
	}

	return message
}

// MessageToDBMessage converts a Message to a DBMessage.
func MessageToDBMessage(message Message) db.DBMessage {
	chainID := message.ChainID
	destination := message.Destination
	nonce := message.Nonce
	messageBytes := message.Message
	blockNumber := message.BlockNumber
	executed := message.Executed
	minimumTimeSet := message.MinimumTimeSet
	minimumTime := message.MinimumTime

	return db.DBMessage{
		ChainID:        &chainID,
		Destination:    &destination,
		Nonce:          &nonce,
		Message:        &messageBytes,
		BlockNumber:    &blockNumber,
		Executed:       &executed,
		MinimumTimeSet: &minimumTimeSet,
		MinimumTime:    &minimumTime,
	}
}

// AgentsTypesMessageToMessage converts an agentsTypes.Message to a Message.
func AgentsTypesMessageToMessage(message agentsTypes.Message, blockNumber uint64, minimumTimeSet bool, minimumTime uint64) (Message, error) {
	rawMessage, err := agentsTypes.EncodeMessage(message)
	if err != nil {
		return Message{}, fmt.Errorf("failed to encode message: %w", err)
	}
	return Message{
		ChainID:        message.OriginDomain(),
		Destination:    message.DestinationDomain(),
		Nonce:          message.Nonce(),
		Message:        rawMessage,
		BlockNumber:    blockNumber,
		Executed:       false,
		MinimumTimeSet: minimumTimeSet,
		MinimumTime:    minimumTime,
	}, nil
}
