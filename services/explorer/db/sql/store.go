package sql

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	"time"

	gormClickhouse "gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
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
//
//nolint:cyclop
func OpenGormClickhouse(ctx context.Context, address string, readOnly bool, handler metrics.Handler) (*Store, error) {
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

	// nolint
	if !readOnly {
		// load all models, check if table exists before doing so
		if (!clickhouseDB.WithContext(ctx).Migrator().HasTable(&TokenIndex{})) {
			err = clickhouseDB.WithContext(ctx).Set("gorm:table_options", "ENGINE=MergeTree ORDER BY (chain_id, token_index, contract_address)").AutoMigrate(&TokenIndex{})
			if err != nil {
				return nil, fmt.Errorf("could not migrate token indexes on clickhouse: %w", err)
			}
		}
		if (!clickhouseDB.WithContext(ctx).Migrator().HasTable(&SwapEvent{}) || !clickhouseDB.WithContext(ctx).Migrator().HasTable(&BridgeEvent{}) || !clickhouseDB.WithContext(ctx).Migrator().HasTable(&MessageBusEvent{})) {
			err = clickhouseDB.WithContext(ctx).Set("gorm:table_options", "ENGINE=ReplacingMergeTree(insert_time) ORDER BY (event_index, block_number, event_type, tx_hash, chain_id, contract_address)").AutoMigrate(&SwapEvent{}, &BridgeEvent{}, MessageBusEvent{})
			if err != nil {
				return nil, fmt.Errorf("could not migrate on clickhouse: %w", err)
			}
		}
		if (!clickhouseDB.WithContext(ctx).Migrator().HasTable(&SwapFees{})) {
			err = clickhouseDB.WithContext(ctx).Set("gorm:table_options", "ENGINE=MergeTree ORDER BY (chain_id, contract_address, block_number, fee_type)").AutoMigrate(&SwapFees{})
			if err != nil {
				return nil, fmt.Errorf("could not migrate token indexes on clickhouse: %w", err)
			}
		}
		if (!clickhouseDB.WithContext(ctx).Migrator().HasTable(&LastBlock{})) {
			err = clickhouseDB.WithContext(ctx).Set("gorm:table_options", "ENGINE=MergeTree ORDER BY (chain_id, contract_address)").AutoMigrate(&LastBlock{})
			if err != nil {
				return nil, fmt.Errorf("could not migrate last block number on clickhouse: %w", err)
			}
		}
		if (!clickhouseDB.WithContext(ctx).Migrator().HasTable(&CCTPEvent{})) {
			err = clickhouseDB.WithContext(ctx).Set("gorm:table_options", "ENGINE=ReplacingMergeTree(insert_time) ORDER BY (tx_hash, contract_address, block_number, event_type, request_id)").AutoMigrate(&CCTPEvent{})
			if err != nil {
				return nil, fmt.Errorf("could not migrate last block number on clickhouse: %w", err)
			}
		}
	}
	db, err := clickhouseDB.DB()

	if err != nil {
		return nil, fmt.Errorf("failed to get clickhouse db: %w", err)
	}
	db.SetConnMaxIdleTime(1 * time.Second)
	db.SetConnMaxLifetime(30 * time.Second)
	handler.AddGormCallbacks(clickhouseDB)

	return &Store{clickhouseDB}, nil
}
