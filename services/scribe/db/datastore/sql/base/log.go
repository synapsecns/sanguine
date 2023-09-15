package base

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/services/scribe/db"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// StoreLogs stores a log.
func (s Store) StoreLogs(ctx context.Context, chainID uint32, logs ...types.Log) error {
	var storeLogs []Log
	for _, log := range logs {
		var topics []sql.NullString

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

		newLog := Log{
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
			Confirmed:       true,
		}

		storeLogs = append(storeLogs, newLog)
	}

	dbTx := s.DB().WithContext(ctx)
	if s.db.Dialector.Name() == dbcommon.Sqlite.String() {
		dbTx = dbTx.Clauses(clause.OnConflict{
			Columns: []clause.Column{
				{Name: ContractAddressFieldName}, {Name: ChainIDFieldName}, {Name: TxHashFieldName}, {Name: BlockIndexFieldName},
			},
			DoNothing: true,
		}).CreateInBatches(&storeLogs, 10)
	} else {
		dbTx = dbTx.Clauses(clause.Insert{
			Modifier: "IGNORE",
		}).Create(&storeLogs)
	}

	if dbTx.Error != nil {
		return fmt.Errorf("could not store log: %w", dbTx.Error)
	}

	return nil
}

// ConfirmLogsForBlockHash confirms logs for a given block hash.
func (s Store) ConfirmLogsForBlockHash(ctx context.Context, chainID uint32, blockHash common.Hash) error {
	dbTx := s.DB().WithContext(ctx).
		Model(&Log{}).
		Where(&Log{BlockHash: blockHash.String(), ChainID: chainID}).
		Update("confirmed", true)

	if dbTx.Error != nil {
		return fmt.Errorf("could not confirm log: %w", dbTx.Error)
	}

	return nil
}

// DeleteLogsForBlockHash deletes logs with a given block hash.
func (s Store) DeleteLogsForBlockHash(ctx context.Context, blockHash common.Hash, chainID uint32) error {
	dbTx := s.DB().WithContext(ctx).
		Where(&Log{BlockHash: blockHash.String(), ChainID: chainID}).
		Delete(&Log{})

	if dbTx.Error != nil {
		return fmt.Errorf("could not delete logs: %w", dbTx.Error)
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
		BlockIndex:      logFilter.Index,
		Confirmed:       logFilter.Confirmed,
	}
}

// RetrieveLogsWithFilter retrieves all logs that match a filter given a page.
func (s Store) RetrieveLogsWithFilter(ctx context.Context, logFilter db.LogFilter, page int) (logs []*types.Log, err error) {
	if page < 1 {
		page = 1
	}
	dbLogs := []Log{}
	queryFilter := logFilterToQuery(logFilter)

	dbTx := s.DB().WithContext(ctx).
		Model(&Log{}).
		Where(&queryFilter).
		Order(fmt.Sprintf("%s desc, %s desc", BlockNumberFieldName, BlockIndexFieldName)).
		Offset((page - 1) * PageSize).
		Limit(PageSize).
		Find(&dbLogs)

	if dbTx.Error != nil {
		logger.Infof("error while getting logs with queryFilter %v, err: %v", queryFilter, dbTx.Error)

		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return []*types.Log{}, fmt.Errorf("could not find logs with filter %v: %w", logFilter, db.ErrNotFound)
		}
		return []*types.Log{}, fmt.Errorf("could not retrieve logs: %w", dbTx.Error)
	}
	// TODO DELETE
	logger.Infof("RetrieveLogsWithFilter query: %v, logs: %v", queryFilter, dbLogs)
	return buildLogsFromDBLogs(dbLogs), nil
}

// RetrieveLogCountForContract retrieves the count of logs per contract.
func (s Store) RetrieveLogCountForContract(ctx context.Context, contractAddress common.Address, chainID uint32) (int64, error) {
	var count int64
	dbTx := s.DB().WithContext(ctx).
		Model(&Log{}).
		Where(&Log{ChainID: chainID}).
		Where(&Log{ContractAddress: contractAddress.String()}).
		Count(&count)

	if dbTx.Error != nil {
		return 0, fmt.Errorf("could not count logs: %w", dbTx.Error)
	}

	return count, nil
}

// RetrieveLogsInRange retrieves all logs that match an inputted filter and are within a range given a page.
func (s Store) RetrieveLogsInRange(ctx context.Context, logFilter db.LogFilter, startBlock, endBlock uint64, page int) (logs []*types.Log, err error) {
	logs, err = s.retrieveLogsInRangeQuery(ctx, logFilter, startBlock, endBlock, page, "desc")
	if err != nil {
		return nil, err
	}

	return logs, nil
}

// RetrieveLogsInRangeAsc retrieves all logs that match an inputted filter and are within a range given a page - in ascending order.
func (s Store) RetrieveLogsInRangeAsc(ctx context.Context, logFilter db.LogFilter, startBlock, endBlock uint64, page int) (logs []*types.Log, err error) {
	logs, err = s.retrieveLogsInRangeQuery(ctx, logFilter, startBlock, endBlock, page, "asc")
	if err != nil {
		return nil, err
	}

	return logs, nil
}

func buildLogsFromDBLogs(dbLogs []Log) []*types.Log {
	logs := make([]*types.Log, len(dbLogs))
	for i, dbLog := range dbLogs {
		topics := buildTopics(dbLog)

		parsedLog := &types.Log{
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

		logs[i] = parsedLog
	}
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

func (s Store) retrieveLogsInRangeQuery(ctx context.Context, logFilter db.LogFilter, startBlock, endBlock uint64, page int, order string) (logs []*types.Log, err error) {
	if page < 1 {
		page = 1
	}
	dbLogs := []Log{}
	queryFilter := logFilterToQuery(logFilter)
	rangeQuery := fmt.Sprintf("%s BETWEEN ? AND ?", BlockNumberFieldName)
	dbTx := s.DB().WithContext(ctx).
		Model(&Log{}).
		Where(&queryFilter).
		Where(rangeQuery, startBlock, endBlock).
		Order(fmt.Sprintf("%s %s, %s %s", BlockNumberFieldName, order, BlockIndexFieldName, order)).
		Offset((page - 1) * PageSize).
		Limit(PageSize).
		Find(&dbLogs)

	if dbTx.Error != nil {
		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return []*types.Log{}, fmt.Errorf("could not find logs with filter %v, in range %v-%v: %w", logFilter, startBlock, endBlock, db.ErrNotFound)
		}
		return []*types.Log{}, fmt.Errorf("could not retrieve logs: %w", dbTx.Error)
	}

	return buildLogsFromDBLogs(dbLogs), nil
}
