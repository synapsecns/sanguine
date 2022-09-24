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
	clickhouseDB, err := gorm.Open(gormClickhouse.Open(address), &gorm.Config{
		Logger:               dbcommon.GetGormLogger(logger),
		FullSaveAssociations: true,
		NowFunc:              time.Now,
	})

	if err != nil {
		return nil, fmt.Errorf("could not create clickhouse connection: %w", err)
	}

	// load all models
	err = clickhouseDB.WithContext(ctx).AutoMigrate(GetAllModels()...)
	if err != nil {
		return nil, fmt.Errorf("could not migrate on clickhouse: %w", err)
	}
	return &Store{clickhouseDB}, nil
}

// GetAllModels gets all models to migrate.
func GetAllModels() (allModels []interface{}) {
	allModels = append(allModels,
		&SwapEvent{}, &BridgeEvent{},
	)
	return allModels
}

var _ db.ConsumerDB = &Store{}
