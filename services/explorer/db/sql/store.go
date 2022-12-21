package sql

import (
	"context"
	"fmt"
	gormClickhouse "gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

// Store is the clickhouse store. It extends the base store for sqlite specific queries.
type Store struct {
	db *gorm.DB
}

// UNSAFE_DB gets the underlying gorm db.
//
//nolint:golint,revive,stylecheck
func (s *Store) UNSAFE_DB() *gorm.DB {
	return s.db
}

// OpenGormClickhouse opens a gorm connection to clickhouse.
func OpenGormClickhouse(ctx context.Context, address string, readOnly bool) (*Store, error) {
	clickhouseDB, err := gorm.Open(gormClickhouse.New(gormClickhouse.Config{
		DSN: address,
	}), &gorm.Config{
		Logger:               gormLogger.Default.LogMode(gormLogger.Silent),
		FullSaveAssociations: true,
		NowFunc:              time.Now,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open gorm clickhouse: %w", err)
	}

	if !readOnly {
		// load all models
		err = clickhouseDB.WithContext(ctx).Set("gorm:table_options", "ENGINE=ReplacingMergeTree(insert_time) ORDER BY (event_index, block_number, event_type, tx_hash, chain_id, contract_address)").AutoMigrate(&SwapEvent{}, &BridgeEvent{}, MessageBusEvent{})

		if err != nil {
			return nil, fmt.Errorf("could not migrate on clickhouse: %w", err)
		}
		err = clickhouseDB.WithContext(ctx).Set("gorm:table_options", "ENGINE=MergeTree ORDER BY (chain_id, contract_address)").AutoMigrate(&LastBlock{})
		if err != nil {
			return nil, fmt.Errorf("could not migrate last block number on clickhouse: %w", err)
		}
		err = clickhouseDB.WithContext(ctx).Set("gorm:table_options", "ENGINE=MergeTree ORDER BY (chain_id, token_index, contract_address)").AutoMigrate(&TokenIndex{})
		if err != nil {
			return nil, fmt.Errorf("could not migrate token indexes on clickhouse: %w", err)
		}
	}
	db, err := clickhouseDB.DB()

	if err != nil {
		return nil, fmt.Errorf("failed to get clickhouse db: %w", err)
	}

	db.SetConnMaxIdleTime(300 * time.Second)
	db.SetConnMaxLifetime(300 * time.Second)
	return &Store{clickhouseDB}, nil
}
