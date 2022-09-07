package base

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sort"

	"github.com/synapsecns/sanguine/services/scribe/db"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// StoreLog stores a log.
func (s Store) StoreLog(ctx context.Context, log types.Log, chainID uint32) error {
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
				{Name: ContractAddressFieldName}, {Name: ChainIDFieldName}, {Name: TxHashFieldName}, {Name: IndexFieldName},
			},
			DoNothing: true,
		}).
		Create(&Log{
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
			Index:           uint64(log.Index),
			Removed:         log.Removed,
		})

	if dbTx.Error != nil {
		return fmt.Errorf("could not store log: %w", dbTx.Error)
	}

	return nil
}

// logFilterToQuery takes in a LogFilter and converts it to a database-type Log.
// This is used to query with `WHERE` based on the filter.
func logFilterToQuery(logFilter db.LogFilter) Log {
	return Log{
		ContractAddress: logFilter.ContractAddress,
		ChainID:         logFilter.ChainID,
		BlockNumber:     logFilter.BlockNumber,
		TxHash:          logFilter.TxHash,
		TxIndex:         logFilter.TxIndex,
		BlockHash:       logFilter.BlockHash,
		Index:           logFilter.Index,
	}
}

// RetrieveLogsWithFilter retrieves all logs that match a filter.
func (s Store) RetrieveLogsWithFilter(ctx context.Context, logFilter db.LogFilter) (logs []*types.Log, err error) {
	dbLogs := []Log{}
	query := logFilterToQuery(logFilter)
	dbTx := s.DB().WithContext(ctx).
		Model(&Log{}).
		Where(&query).
		Find(&dbLogs)

	if dbTx.Error != nil {
		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return []*types.Log{}, fmt.Errorf("could not find logs with filter %v: %w", logFilter, db.ErrNotFound)
		}
		return []*types.Log{}, fmt.Errorf("could not store log: %w", dbTx.Error)
	}

	return buildLogsFromDBLogs(dbLogs), nil
}

// RetrieveLogsInRange retrieves all logs that match an inputted filter, and are within a range.
func (s Store) RetrieveLogsInRange(ctx context.Context, logFilter db.LogFilter, startBlock, endBlock uint64) (logs []*types.Log, err error) {
	dbLogs := []Log{}
	queryFilter := logFilterToQuery(logFilter)
	rangeQuery := fmt.Sprintf("%s BETWEEN ? AND ?", BlockNumberFieldName)
	dbTx := s.DB().WithContext(ctx).
		Model(&Log{}).
		Where(&queryFilter).
		Where(rangeQuery, startBlock, endBlock).
		Find(&dbLogs)

	if dbTx.Error != nil {
		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return []*types.Log{}, fmt.Errorf("could not find logs with filter %v, in range %v-%v: %w", logFilter, startBlock, endBlock, db.ErrNotFound)
		}
		return []*types.Log{}, fmt.Errorf("could not store log: %w", dbTx.Error)
	}

	return buildLogsFromDBLogs(dbLogs), nil
}

func buildLogsFromDBLogs(dbLogs []Log) []*types.Log {
	var logs []*types.Log
	for _, dbLog := range dbLogs {
		topics := buildTopics(dbLog)

		parsedLog := &types.Log{
			Address:     common.HexToAddress(dbLog.ContractAddress),
			Topics:      topics,
			Data:        dbLog.Data,
			BlockNumber: dbLog.BlockNumber,
			TxHash:      common.HexToHash(dbLog.TxHash),
			TxIndex:     uint(dbLog.TxIndex),
			BlockHash:   common.HexToHash(dbLog.BlockHash),
			Index:       uint(dbLog.Index),
			Removed:     dbLog.Removed,
		}

		logs = append(logs, parsedLog)
	}
	sort.Slice(logs, func(i, j int) bool {
		return logs[i].Index < logs[j].Index
	})
	return logs
}

func buildTopics(log Log) []common.Hash {
	topics := []common.Hash{}
	if log.PrimaryTopic.Valid {
		topics = append(topics, common.HexToHash(log.PrimaryTopic.String))
	}
	if log.TopicA.Valid {
		topics = append(topics, common.HexToHash(log.TopicA.String))
	}
	if log.TopicB.Valid {
		topics = append(topics, common.HexToHash(log.TopicB.String))
	}
	if log.TopicC.Valid {
		topics = append(topics, common.HexToHash(log.TopicC.String))
	}

	return topics
}
