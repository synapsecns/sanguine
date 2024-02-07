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

// TODO support more filtering options

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

// RetrieveLogsFromHeadRangeQuery retrieves all logs (including unconfirmed) for a given contract address, chain ID, and range.
func (s Store) RetrieveLogsFromHeadRangeQuery(ctx context.Context, logFilter db.LogFilter, startBlock uint64, endBlock uint64, page int) ([]*types.Log, error) {
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
	queryFilter := logFilterToQuery(logFilter)

	var dbLogs []Log
	subQuery1 := s.DB().WithContext(ctx).ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(Log{}).Select("*").Where("block_number BETWEEN ? AND ?", startBlock, lastIndexed).Where(queryFilter).Find(&[]Log{})
	})
	subQuery2 := s.DB().WithContext(ctx).ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(LogAtHead{}).Select(LogColumns).Where("block_number BETWEEN ? AND ?", lastIndexed+1, endBlock).Where(queryFilter).Find(&[]Log{})
	})
	query := fmt.Sprintf("SELECT * FROM (%s UNION %s) AS unionedTable ORDER BY %s DESC, %s DESC LIMIT %d OFFSET %d", subQuery1, subQuery2, BlockNumberFieldName, BlockIndexFieldName, PageSize, (page-1)*PageSize)
	dbTx := s.DB().WithContext(ctx).Raw(query).Scan(&dbLogs)

	if dbTx.Error != nil {
		return nil, fmt.Errorf("error getting newly confirmed data %w", dbTx.Error)
	}
	return buildLogsFromDBLogs(dbLogs), nil
}

// RetrieveReceiptsFromHeadRangeQuery retrieves all receipts (including unconfirmed) for a given contract address, chain ID, and range.
func (s Store) RetrieveReceiptsFromHeadRangeQuery(ctx context.Context, receiptFilter db.ReceiptFilter, startBlock uint64, endBlock uint64, page int) ([]types.Receipt, error) {
	if receiptFilter.ContractAddress == "" || receiptFilter.ChainID == 0 {
		return nil, fmt.Errorf("contract address and chain ID must be passed")
	}
	if page < 1 {
		page = 1
	}

	lastIndexed, err := s.RetrieveLastIndexed(ctx, common.HexToAddress(receiptFilter.ContractAddress), receiptFilter.ChainID, false)
	if err != nil {
		return nil, fmt.Errorf("could not get last block indexed: %w", err)
	}
	queryFilter := receiptFilterToQuery(receiptFilter)

	var dbReceipts []Receipt
	subQuery1 := s.DB().WithContext(ctx).ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(Receipt{}).Select("*").Where("block_number BETWEEN ? AND ?", startBlock, lastIndexed).Where(queryFilter).Find(&[]Receipt{})
	})
	subQuery2 := s.DB().WithContext(ctx).ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(ReceiptAtHead{}).Select(ReceiptColumns).Where("block_number BETWEEN ? AND ?", lastIndexed+1, endBlock).Where(queryFilter).Find(&[]Receipt{})
	})
	query := fmt.Sprintf("SELECT * FROM (%s UNION %s) AS unionedTable ORDER BY %s DESC, %s DESC LIMIT %d OFFSET %d", subQuery1, subQuery2, BlockNumberFieldName, TransactionIndexFieldName, PageSize, (page-1)*PageSize)
	dbTx := s.DB().WithContext(ctx).Raw(query).Scan(&dbReceipts)

	if dbTx.Error != nil {
		return nil, fmt.Errorf("error getting newly confirmed data %w", dbTx.Error)
	}
	receipts, err := s.buildReceiptsFromDBReceipts(ctx, dbReceipts, receiptFilter.ChainID)
	if err != nil {
		return nil, fmt.Errorf("error building receipts from db receipts: %w", err)
	}
	return receipts, nil
}

// TODO make a query for getting latest tx

// RetrieveUnconfirmedEthTxsFromHeadRangeQuery retrieves all unconfirmed ethTx for a given chain ID and range.
// lastIndexed is passed because the ethtx table does not have contract addresses, thus the last indexed for that contract
// cannot be determined for the join. Pass last indexed for the log that you are trying to mature with data.
func (s Store) RetrieveUnconfirmedEthTxsFromHeadRangeQuery(ctx context.Context, ethTxFilter db.EthTxFilter, startBlock uint64, endBlock uint64, lastIndexed uint64, page int) ([]db.TxWithBlockNumber, error) {
	if ethTxFilter.ChainID == 0 {
		return nil, fmt.Errorf("chain ID must be passed")
	}
	if page < 1 {
		page = 1
	}

	queryFilter := ethTxFilterToQuery(ethTxFilter)
	var dbEthTxs []EthTx
	subQuery1 := s.DB().WithContext(ctx).ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(EthTx{}).Select("*").Where("block_number BETWEEN ? AND ?", startBlock, lastIndexed).Where(queryFilter).Find(&[]EthTx{})
	})
	subQuery2 := s.DB().WithContext(ctx).ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Model(EthTxAtHead{}).Select(EthTxColumns).Where("block_number BETWEEN ? AND ?", lastIndexed+1, endBlock).Where(queryFilter).Find(&[]EthTx{})
	})
	query := fmt.Sprintf("SELECT * FROM (%s UNION %s) AS unionedTable ORDER BY %s DESC, %s DESC LIMIT %d OFFSET %d", subQuery1, subQuery2, BlockNumberFieldName, TransactionIndexFieldName, PageSize, (page-1)*PageSize)
	dbTx := s.DB().WithContext(ctx).Raw(query).Scan(&dbEthTxs)

	if dbTx.Error != nil {
		return nil, fmt.Errorf("error getting newly confirmed data %w", dbTx.Error)
	}
	txs, err := buildEthTxsFromDBEthTxs(dbEthTxs)
	if err != nil {
		return nil, fmt.Errorf("error building receipts from db receipts: %w", err)
	}
	return txs, nil
}

// FlushFromHeadTables deletes all logs, receipts, and txs from the head table that are older than the given time.
func (s Store) FlushFromHeadTables(ctx context.Context, time int64) error {
	err := s.DB().WithContext(ctx).Model(&LogAtHead{}).Where("insert_time < ?", time).Delete(&LogAtHead{}).Error
	if err != nil {
		return fmt.Errorf("error flushing logs from head: %w", err)
	}
	err = s.DB().WithContext(ctx).Model(&EthTxAtHead{}).Where("insert_time < ?", time).Delete(&EthTxAtHead{}).Error
	if err != nil {
		return fmt.Errorf("error flushing eth_txes from head: %w", err)
	}
	err = s.DB().WithContext(ctx).Model(&ReceiptAtHead{}).Where("insert_time < ?", time).Delete(&ReceiptAtHead{}).Error
	if err != nil {
		return fmt.Errorf("error flushing receipts from head: %w", err)
	}
	return nil
}
