package sql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// StoreFailedLog stores a failed log.
func (s Store) StoreFailedLog(ctx context.Context, log types.Log, chainID uint32) error {
	topics := []sql.NullString{}
	topicsLength := len(log.Topics)
	// Ethereum topics are always 3 long, excluding the primary topic.
	indexedTopics := 3
	// Loop through the topics and convert them to nullStrings.
	// If the topic is empty, we set Valid to false.
	// If the topic is not empty, provide its string value and set Valid to true.
	for index := 0; index <= indexedTopics+1; index++ {
		if index < topicsLength {
			topics = append(topics, sql.NullString{
				String: log.Topics[index].String(),
				Valid:  true,
			})
		} else {
			topics = append(topics, sql.NullString{
				Valid: false,
			})
		}
	}
	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: ContractAddressFieldName}, {Name: ChainIDFieldName}, {Name: TxHashFieldName}, {Name: BlockIndexFieldName},
			},
			DoNothing: true,
		}).
		Create(&FailedLog{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			PrimaryTopic:    topics[0],
			TopicA:          topics[1],
			TopicB:          topics[2],
			TopicC:          topics[3],
			Data:            log.Data,
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),
			TxIndex:         uint64(log.TxIndex),
			BlockHash:       log.BlockHash.String(),
			BlockIndex:      uint64(log.Index),
			Removed:         log.Removed,
			Confirmed:       false,
		})

	if dbTx.Error != nil {
		return fmt.Errorf("could not store log: %w", dbTx.Error)
	}

	return nil
}

// DeleteFailedLog deletes a failed log from the table.
func (s Store) DeleteFailedLog(ctx context.Context, log types.Log, chainID uint32) error {
	dbTx := s.DB().WithContext(ctx).
		Where(&FailedLog{
			ContractAddress: log.Address.String(),
			ChainID:         chainID,
			BlockNumber:     log.BlockNumber,
			TxHash:          log.TxHash.String(),
			TxIndex:         uint64(log.TxIndex),
			BlockHash:       log.BlockHash.String(),
			BlockIndex:      uint64(log.Index),
		}).
		Delete(&FailedLog{})

	if dbTx.Error != nil {
		return fmt.Errorf("could not delete failed log: %w", dbTx.Error)
	}

	return nil
}

// RetrieveFailedLogs stores a log that was failed to be parsed and stored.
func (s Store) RetrieveFailedLogs(ctx context.Context, chainID uint32) (logs []types.Log, err error) {
	dbLogs := []FailedLog{}
	dbTx := s.DB().WithContext(ctx).
		Model(&FailedLog{}).
		Where(&FailedLog{
			ChainID: chainID,
		}).
		Order(fmt.Sprintf("%s, %s", BlockNumberFieldName, BlockIndexFieldName)).
		Find(&dbLogs)

	if dbTx.Error != nil {
		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return []types.Log{}, fmt.Errorf("could not find failed logs: %w", dbTx.Error)
		}
		return []types.Log{}, fmt.Errorf("could not store log: %w", dbTx.Error)
	}

	return buildLogsFromDBLogs(dbLogs), nil
}

func buildLogsFromDBLogs(failedLogs []FailedLog) []types.Log {
	var logs []types.Log
	for _, dbLog := range failedLogs {
		topics := buildTopics(dbLog)

		parsedLog := types.Log{
			Address:     common.HexToAddress(dbLog.ContractAddress),
			Topics:      topics,
			Data:        dbLog.Data,
			BlockNumber: dbLog.BlockNumber,
			TxHash:      common.HexToHash(dbLog.TxHash),
			TxIndex:     uint(dbLog.TxIndex),
			BlockHash:   common.HexToHash(dbLog.BlockHash),
			Index:       uint(dbLog.BlockIndex),
			Removed:     dbLog.Removed,
		}

		logs = append(logs, parsedLog)
	}
	return logs
}

func buildTopics(failedLog FailedLog) []common.Hash {
	topics := []common.Hash{}
	if failedLog.PrimaryTopic.Valid {
		topics = append(topics, common.HexToHash(failedLog.PrimaryTopic.String))
	}
	if failedLog.TopicA.Valid {
		topics = append(topics, common.HexToHash(failedLog.TopicA.String))
	}
	if failedLog.TopicB.Valid {
		topics = append(topics, common.HexToHash(failedLog.TopicB.String))
	}
	if failedLog.TopicC.Valid {
		topics = append(topics, common.HexToHash(failedLog.TopicC.String))
	}

	return topics
}
