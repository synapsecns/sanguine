package sql

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/services/explorer/db"
	gormClickhouse "gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"time"
)

// Store is the clickhouse store. It extends the base store for sqlite specific queries.
type Store struct {
	db *gorm.DB
}

// DB gets the underlying gorm db.
func (s Store) DB() *gorm.DB {
	return s.db
}

// OpenGormClickhouse opens a gorm connection to clickhouse.
func OpenGormClickhouse(ctx context.Context, address string) (*Store, error) {
	clickhouseDB, err := gorm.Open(gormClickhouse.New(gormClickhouse.Config{
		DSN: address,
	}), &gorm.Config{
		Logger:               dbcommon.GetGormLogger(logger),
		FullSaveAssociations: true,
		NowFunc:              time.Now,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open gorm clickhouse: %w", err)
	}

	// load all models
	err = clickhouseDB.WithContext(ctx).
		// TODO: add ReplacingEngineTree
		// Set("gorm:table_options", "ENGINE=ReplacingMergeTree(insert_time) ORDER BY block_number").
		AutoMigrate(&SwapEvent{}, &BridgeEvent{})
	if err != nil {
		return nil, fmt.Errorf("could not migrate on clickhouse: %w", err)
	}
	return &Store{clickhouseDB}, nil
}

var _ db.ConsumerDB = &Store{}
