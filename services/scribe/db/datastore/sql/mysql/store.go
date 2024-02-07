package mysql

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	scribeLogger "github.com/synapsecns/sanguine/services/scribe/logger"
	gormLogger "gorm.io/gorm/logger"

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
var MaxIdleConns = 1048

// MaxOpenConns is exported here for testing. Tests execute too slowly with a reconnect each time.
var MaxOpenConns = 1048

// NamingStrategy is exported here for testing.
var NamingStrategy = schema.NamingStrategy{
	TablePrefix: "v3_",
}

// NewMysqlStore creates a new mysql store for a given data store.
func NewMysqlStore(parentCtx context.Context, dbURL string, handler metrics.Handler, skipMigrations bool) (_ *Store, err error) {
	logger.Debug("creating mysql store")
	scribeLogger.ReportScribeState(0, 0, nil, scribeLogger.CreatingSQLStore)
	ctx, span := handler.Tracer().Start(parentCtx, "start-mysql")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	gdb, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{
		Logger:                 gormLogger.Default.LogMode(gormLogger.Silent),
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
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetMaxOpenConns(MaxOpenConns)

	handler.AddGormCallbacks(gdb)

	if !skipMigrations {
		// migrate in a transaction since we skip this by default
		err = gdb.Transaction(func(tx *gorm.DB) error {
			//nolint: wrapcheck
			return gdb.WithContext(ctx).AutoMigrate(base.GetAllModels()...)
		})
	}

	if err != nil {
		return nil, fmt.Errorf("could not migrate on mysql: %w", err)
	}
	return &Store{base.NewStore(gdb, handler)}, nil
}

// var _ db.Service = &Store{}
