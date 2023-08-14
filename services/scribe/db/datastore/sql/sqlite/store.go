package sqlite

import (
	"context"
	"fmt"
	scribeLogger "github.com/synapsecns/sanguine/services/scribe/logger"
	gormLogger "gorm.io/gorm/logger"

	"github.com/synapsecns/sanguine/core/metrics"
	"os"

	"github.com/synapsecns/sanguine/services/scribe/db/datastore/sql/base"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Store is the sqlite store. It extends the base store for sqlite specific queries.
type Store struct {
	*base.Store
}

// NewSqliteStore creates a new sqlite data store.
func NewSqliteStore(parentCtx context.Context, dbPath string, handler metrics.Handler, skipMigrations bool) (_ *Store, err error) {
	logger.Debugf("creating sqlite store at %s", dbPath)
	scribeLogger.ReportScribeState(0, 0, nil, scribeLogger.CreatingSQLStore)

	ctx, span := handler.Tracer().Start(parentCtx, "start-sqlite")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// create the directory to the store if it doesn't exist
	err = os.MkdirAll(dbPath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("could not create sqlite store")
	}

	logger.Warnf("database is at %s/synapse.db", dbPath)

	gdb, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s/%s", dbPath, "synapse.db")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		//Logger:                                   common_base.GetGormLogger(logger),
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),

		FullSaveAssociations:   true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, fmt.Errorf("could not connect to db %s: %w", dbPath, err)
	}

	handler.AddGormCallbacks(gdb)

	if !skipMigrations {
		err = gdb.WithContext(ctx).AutoMigrate(base.GetAllModels()...)
		if err != nil {
			return nil, fmt.Errorf("could not migrate models: %w", err)
		}
	}
	return &Store{base.NewStore(gdb, handler)}, nil
}

// var _ db.TxQueueDb = &Store{}
