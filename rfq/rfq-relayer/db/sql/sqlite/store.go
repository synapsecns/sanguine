package sqlite

import (
	"context"
	"fmt"
	"os"

	common_base "github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db/model"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db/sql/base"
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

	ctx, span := handler.Tracer().Start(parentCtx, "start-sqlite")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// create the directory to the store if it doesn't exist
	err = os.MkdirAll(dbPath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("could not create sqlite store")
	}

	logger.Infof("executor database is at %s", fmt.Sprintf("%s/%s", dbPath, "synapse.db"))

	gdb, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s/%s", dbPath, "synapse.db")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		FullSaveAssociations:                     true,
		SkipDefaultTransaction:                   true,
		Logger:                                   common_base.GetGormLogger(logger),
	})
	if err != nil {
		return nil, fmt.Errorf("could not connect to db %s: %w", dbPath, err)
	}

	handler.AddGormCallbacks(gdb)

	if !skipMigrations {
		err = gdb.WithContext(ctx).AutoMigrate(model.GetAllModels()...)
		if err != nil {
			return nil, fmt.Errorf("could not migrate models: %w", err)
		}
	}

	return &Store{base.NewStore(gdb, handler)}, nil
}
