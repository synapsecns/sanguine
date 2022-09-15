package sql

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/dbcommon"
	gormClickhouse "gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
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

// NamingStrategy is exported here for testing.
var NamingStrategy = schema.NamingStrategy{}

func OpenGormClickhouse(ctx context.Context, address string) (*Store, error) {
	fmt.Println("SOME", gormClickhouse.Open(address))
	clickhouseDB, err := gorm.Open(gormClickhouse.Open(address), &gorm.Config{
		Logger:               dbcommon.GetGormLogger(logger),
		FullSaveAssociations: true,
		NamingStrategy:       NamingStrategy,
		NowFunc:              time.Now,
	})

	if err != nil {
		return nil, fmt.Errorf("could not create clickhouse connection: %w", err)
	}

	//testing just swap
	err = clickhouseDB.WithContext(ctx).AutoMigrate(GetAllModels()...)
	//err = clickhouseDB.AutoMigrate(&SwapEvent{})
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

// var _ db.Service = &Store{}
