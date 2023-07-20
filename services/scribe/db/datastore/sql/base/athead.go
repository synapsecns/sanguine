package base

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// StoreLogsAtHead stores a log at the Head of the chain.
func (s Store) StoreLogsAtHead(ctx context.Context, chainID uint32, logs ...types.Log) error {
	var storeLogs []LogAtHead
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

		newLog := LogAtHead{
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
			InsertTime:      uint64(time.Now().UnixNano()),
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

// StoreReceiptAtHead stores a receipt.
func (s Store) StoreReceiptAtHead(ctx context.Context, chainID uint32, receipt types.Receipt) error {
	dbTx := s.DB().WithContext(ctx)
	if s.DB().Dialector.Name() == dbcommon.Sqlite.String() {
		dbTx = dbTx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: TxHashFieldName}, {Name: ChainIDFieldName}},
			DoNothing: true,
		})
	} else {
		dbTx = dbTx.Clauses(clause.Insert{
			Modifier: "IGNORE",
		})
	}
	dbTx = dbTx.Create(&ReceiptAtHead{
		ChainID:           chainID,
		Type:              receipt.Type,
		PostState:         receipt.PostState,
		Status:            receipt.Status,
		CumulativeGasUsed: receipt.CumulativeGasUsed,
		Bloom:             receipt.Bloom.Bytes(),
		TxHash:            receipt.TxHash.String(),
		ContractAddress:   receipt.ContractAddress.String(),
		GasUsed:           receipt.GasUsed,
		BlockHash:         receipt.BlockHash.String(),
		BlockNumber:       receipt.BlockNumber.Uint64(),
		TransactionIndex:  uint64(receipt.TransactionIndex),
		Confirmed:         false,
		InsertTime:        uint64(time.Now().UnixNano()),
	})

	if dbTx.Error != nil {
		return fmt.Errorf("could not store receipt: %w", dbTx.Error)
	}

	return nil
}

// StoreEthTxAtHead stores a processed text at Head.
func (s Store) StoreEthTxAtHead(ctx context.Context, tx *types.Transaction, chainID uint32, blockHash common.Hash, blockNumber uint64, transactionIndex uint64) error {
	marshalledTx, err := tx.MarshalBinary()
	if err != nil {
		return fmt.Errorf("could not marshall tx to binary: %w", err)
	}
	dbTx := s.DB().WithContext(ctx)
	if s.DB().Dialector.Name() == dbcommon.Sqlite.String() {
		dbTx = dbTx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: TxHashFieldName}, {Name: ChainIDFieldName}},
			DoNothing: true,
		})
	} else {
		dbTx = dbTx.Clauses(clause.Insert{
			Modifier: "IGNORE",
		})
	}

	dbTx = dbTx.Create(&EthTxAtHead{
		TxHash:           tx.Hash().String(),
		ChainID:          chainID,
		BlockHash:        blockHash.String(),
		BlockNumber:      blockNumber,
		RawTx:            marshalledTx,
		GasFeeCap:        tx.GasFeeCap().Uint64(),
		Confirmed:        false,
		TransactionIndex: transactionIndex,
		InsertTime:       uint64(time.Now().UnixNano()),
	})

	if dbTx.Error != nil {
		return fmt.Errorf("could not create raw tx: %w", dbTx.Error)
	}

	return nil
}

// RetrieveLogsFromHeadRangeQuery retrieves logs all logs (including unconfirmed) for a given contract address and chain ID.
func (s Store) RetrieveLogsFromHeadRangeQuery(ctx context.Context, logFilter db.LogFilter, startBlock uint64, endBlock uint64, page int) (logs []*types.Log, err error) {
	if logFilter.ContractAddress == "" || logFilter.ChainID == 0 {
		return nil, fmt.Errorf("contract address and chain ID must be passed")
	}
	if page < 1 {
		page = 1
	}

	lastIndexed, err := s.RetrieveLastIndexed(ctx, common.HexToAddress(logFilter.ContractAddress), logFilter.ChainID, false)
	if err != nil {
		return nil, fmt.Errorf("could not get last block indexed: %w", err)
	}

	var dbLogs []Log
	subquery1 := s.DB().WithContext(ctx).ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(Log{}).Select("*, NULL AS insert_time").Where("block_number BETWEEN ? AND ?", startBlock, lastIndexed).Find(&[]Log{})
	})
	subquery2 := s.DB().WithContext(ctx).ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(LogAtHead{}).Select("*").Where("block_number BETWEEN ? AND ?", lastIndexed+1, endBlock).Find(&[]Log{})
	})
	dbTx := s.DB().WithContext(ctx).Raw(fmt.Sprintf("SELECT * FROM (%s UNION %s) ORDER BY %s DESC, %s DESC LIMIT ? OFFSET ?", subquery1, subquery2, BlockNumberFieldName, BlockIndexFieldName), PageSize, (page-1)*PageSize).Scan(&dbLogs)

	if dbTx.Error != nil {
		return nil, fmt.Errorf("error getting newly confirmed data %w", dbTx.Error)
	}
	return buildLogsFromDBLogs(dbLogs), nil
}

// FlushLogsFromHead deletes all logs from the head table that are older than the given time.
func (s Store) FlushLogsFromHead(ctx context.Context, time int64) error {
	return s.DB().WithContext(ctx).Model(&LogAtHead{}).Where("insert_time < ?", time).Delete(&LogAtHead{}).Error
}

//
// func (s Store) RetrieveEthTxsWithFilterAndCleanHead(ctx context.Context, ethTxFilter db.EthTxFilter, page int) ([]db.TxWithBlockNumber, error) {
//	if page < 1 {
//		page = 1
//	}
//	var ethTxs []EthTx
//
//	result := s.DB().Table("EthTx").
//		Joins("JOIN EthTxAtHead ON EthTx.TransactionHash = EthTxAtHead.TransactionHash AND EthTx.ChainId = EthTxAtHead.ChainId").
//		Where("EthTx.BlockHash <> EthTxAtHead.BlockHash").
//		Find(&ethTxs)
//
//	if result.Error != nil {
//		return nil, fmt.Errorf("error getting newly confirmed data %v", result.Error)
//	}
//
//	parsedEthTxs, err := buildEthTxsFromDBEthTxs(ethTxs)
//	if err != nil {
//		return []db.TxWithBlockNumber{}, fmt.Errorf("could not build eth txs: %w", err)
//	}
//
//	return parsedEthTxs, nil
//}
