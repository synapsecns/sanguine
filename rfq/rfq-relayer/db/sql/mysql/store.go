package mysql

import (
	"context"
	"fmt"

	common_base "github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db/model"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db/sql/base"

	"time"

	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Store is the sqlite store. It extends the base store for sqlite specific queries.
type Store struct {
	*base.Store
}

// MaxIdleConns is exported here for testing. Tests execute too slowly with a reconnect each time.
var MaxIdleConns = 1048

// NamingStrategy is for table prefixes.
var NamingStrategy = schema.NamingStrategy{}

// MaxOpenConns is exported here for testing. Tests execute too slowly with a reconnect each time.
var MaxOpenConns = 1048

// NewMysqlStore creates a new mysql store for a given data store.
func NewMysqlStore(parentCtx context.Context, dbURL string, handler metrics.Handler, skipMigrations bool) (_ *Store, err error) {
	logger.Debug("creating mysql store")
	ctx, span := handler.Tracer().Start(parentCtx, "start-mysql")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	gdb, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{
		Logger:                                   common_base.GetGormLogger(logger),
		FullSaveAssociations:                     true,
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		NamingStrategy:                           NamingStrategy,
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
	sqlDB.SetMaxOpenConns(MaxOpenConns)

	handler.AddGormCallbacks(gdb)

	if !skipMigrations {
		// migrate in a transaction since we skip this by default
		err = gdb.Transaction(func(tx *gorm.DB) error {
			//nolint: wrapcheck
			return gdb.WithContext(ctx).AutoMigrate(model.GetAllModels()...)
		})
	}
	if err != nil {
		return nil, fmt.Errorf("could not migrate on mysql: %w", err)
	}
	return &Store{base.NewStore(gdb, handler)}, nil
}
