package base

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/scribe/db"
	"github.com/synapsecns/sanguine/scribe/types"
	"gorm.io/gorm"
)

// StoreLog stores a log.
func (s Store) StoreLog(ctx context.Context, log ethTypes.Log, chainID uint32) error {
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

// RetrieveLogByTxHash retrieves a log by tx hash.
func (s Store) RetrieveLogByTxHash(ctx context.Context, txHash common.Hash) (log types.Log, err error) {
	dbLog := Log{}
	dbTx := s.DB().WithContext(ctx).Model(&Log{}).Where(&Log{
		TxHash: txHash.String(),
	}).First(&dbLog)

	if dbTx.Error != nil {
		if errors.Is(dbTx.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("could not find log with tx hash %s: %w", txHash.String(), db.ErrNotFound)
		}
		return nil, fmt.Errorf("could not store log: %w", dbTx.Error)
	}

	parsedLog := types.NewLog(
		common.HexToAddress(dbLog.Address),
		dbLog.ChainID,
		nullStringToHash(dbLog.PrimaryTopic),
		nullStringToHash(dbLog.TopicA),
		nullStringToHash(dbLog.TopicB),
		nullStringToHash(dbLog.TopicC),
		dbLog.Data,
		dbLog.BlockNumber,
		common.HexToHash(dbLog.TxHash),
		dbLog.TxIndex,
		common.HexToHash(dbLog.BlockHash),
		dbLog.Index,
		dbLog.Removed,
	)

	return parsedLog, nil
}

// nullStringToHash converts a null string to a hash.
func nullStringToHash(ns sql.NullString) common.Hash {
	if ns.Valid {
		return common.HexToHash(ns.String)
	}
	// If the nullString is not valid, return an empty hash.
	return common.Hash{}
}
