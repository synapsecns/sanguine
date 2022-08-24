package base

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"gorm.io/gorm"
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
	dbTx := s.DB().WithContext(ctx).Create(&Log{
		Address:      log.Address.String(),
		ChainID:      chainID,
		PrimaryTopic: topics[0],
		TopicA:       topics[1],
		TopicB:       topics[2],
		TopicC:       topics[3],
		Data:         log.Data,
		BlockNumber:  log.BlockNumber,
		TxHash:       log.TxHash.String(),
		TxIndex:      uint64(log.TxIndex),
		BlockHash:    log.BlockHash.String(),
		Index:        uint64(log.Index),
		Removed:      log.Removed,
	})

	if dbTx.Error != nil {
		return fmt.Errorf("could not store log: %w", dbTx.Error)
	}

	return nil
}

// RetrieveLogsByTxHash retrieves all logs that match a tx hash.
func (s Store) RetrieveLogsByTxHash(ctx context.Context, txHash common.Hash) (logs []*types.Log, err error) {
	dbLogs := []Log{}
	dbTx := s.DB().WithContext(ctx).Model(&Log{}).Where(&Log{
		TxHash: txHash.String(),
	}).Find(&dbLogs)

	if dbTx.Error != nil {
		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return []*types.Log{}, fmt.Errorf("could not find logs with tx hash %s: %w", txHash.String(), dbcommon.ErrNotFound)
		}
		return []*types.Log{}, fmt.Errorf("could not store log: %w", dbTx.Error)
	}

	// Format the topics list, only including existing topics.
	for _, dbLog := range dbLogs {
		topics := []common.Hash{}
		if dbLog.PrimaryTopic.Valid {
			topics = append(topics, common.HexToHash(dbLog.PrimaryTopic.String))
		}
		if dbLog.TopicA.Valid {
			topics = append(topics, common.HexToHash(dbLog.TopicA.String))
		}
		if dbLog.TopicB.Valid {
			topics = append(topics, common.HexToHash(dbLog.TopicB.String))
		}
		if dbLog.TopicC.Valid {
			topics = append(topics, common.HexToHash(dbLog.TopicC.String))
		}

		parsedLog := &types.Log{
			Address:     common.HexToAddress(dbLog.Address),
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

	return logs, nil
}
