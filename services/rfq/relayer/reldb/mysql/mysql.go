// Package mysql provides a common interface for starting sql-lite databases
package mysql

import (
	"context"
	"fmt"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb/base"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var logger = log.Logger("mysql-logger")

// Store is the sqlite store. It extends the base store for sqlite specific queries.
type Store struct {
	*base.Store
}

// MaxIdleConns is exported here for testing. Tests execute too slowly with a reconnect each time.
var MaxIdleConns = 0

// NamingStrategy is used to exported here for testing.
var NamingStrategy = schema.NamingStrategy{}

// NewMysqlStore creates a new mysql store for a given data store.
func NewMysqlStore(ctx context.Context, dbURL string, handler metrics.Handler) (*Store, error) {
	logger.Debug("create mysql store")

	gdb, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{
		Logger:               dbcommon.GetGormLogger(logger),
		FullSaveAssociations: true,
		NamingStrategy:       NamingStrategy,
		NowFunc:              time.Now,
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

	handler.AddGormCallbacks(gdb)

	err = gdb.WithContext(ctx).AutoMigrate(base.GetAllModels()...)
	if err != nil {
		return nil, fmt.Errorf("could not migrate on mysql: %w", err)
	}

	return &Store{base.NewStore(gdb, handler)}, nil
}
