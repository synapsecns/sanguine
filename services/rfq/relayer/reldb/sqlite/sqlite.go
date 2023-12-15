package sqlite

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb/base"
	"os"

	"github.com/ipfs/go-log"
	common_base "github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Store is the sqlite store. It extends the base store for sqlite specific queries.
type Store struct {
	*base.Store
}

var logger = log.Logger("rfq-sqlite")

// NewSqliteStore creates a new sqlite data store.
func NewSqliteStore(parentCtx context.Context, dbPath string, handler metrics.Handler) (_ *Store, err error) {
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

	logger.Warnf("rfq database is at %s/synapse.db", dbPath)

	gdb, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s/%s", dbPath, "synapse.db")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   common_base.GetGormLogger(logger),
		FullSaveAssociations:                     true,
		SkipDefaultTransaction:                   true,
	})
	if err != nil {
		return nil, fmt.Errorf("could not connect to db %s: %w", dbPath, err)
	}

	handler.AddGormCallbacks(gdb)

	err = gdb.WithContext(ctx).AutoMigrate(base.GetAllModels()...)
	if err != nil {
		return nil, fmt.Errorf("could not migrate models: %w", err)
	}
	return &Store{base.NewStore(gdb, handler)}, nil
}

var _ reldb.Service = &Store{}
