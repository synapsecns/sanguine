package mysql

import (
	"context"
	"fmt"
	common_base "github.com/synapsecns/sanguine/core/dbcommon"
	"time"

	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/base"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Store is the sqlite store. It extends the base store for sqlite specific queries.
type Store struct {
	*base.Store
}

// MaxIdleConns is exported here for testing. Tests execute too slowly with a reconnect each time.
var MaxIdleConns = 10

// NamingStrategy is exported here for testing.
var NamingStrategy = schema.NamingStrategy{
	TablePrefix: "v3_",
}

// NewMysqlStore creates a new mysql store for a given data store.
func NewMysqlStore(ctx context.Context, dbURL string) (*Store, error) {
	logger.Debug("creating mysql store")

	gdb, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{
		Logger:                 common_base.GetGormLogger(logger),
		FullSaveAssociations:   true,
		NamingStrategy:         NamingStrategy,
		NowFunc:                time.Now,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return nil, fmt.Errorf("could not create mysql connection: %w", err)
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		return nil, fmt.Errorf("could not get sql db: %w", err)
	}

	// fixes a timeout issue https://stackoverflow.com/a/42146536
	sqlDB.SetMaxIdleConns(MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = gdb.WithContext(ctx).AutoMigrate(base.GetAllModels()...)

	if err != nil {
		return nil, fmt.Errorf("could not migrate on mysql: %w", err)
	}
	return &Store{base.NewStore(gdb)}, nil
}

// var _ db.Service = &Store{}
