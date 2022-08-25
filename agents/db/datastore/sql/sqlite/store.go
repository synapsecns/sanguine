package sqlite

import (
	"context"
	"fmt"
	common_base "github.com/synapsecns/sanguine/core/dbcommon"
	"os"

	"github.com/synapsecns/sanguine/agents/db/datastore/sql/base"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Store is the sqlite store. It extends the base store for sqlite specific queries.
type Store struct {
	*base.Store
}

// NewSqliteStore creates a new sqlite data store.
func NewSqliteStore(ctx context.Context, dbPath string) (*Store, error) {
	logger.Debugf("creating sqlite store at %s", dbPath)

	// creat the directory to the store if it doesn't exist
	err := os.MkdirAll(dbPath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("could not create sqlite store")
	}

	gdb, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s/%s", dbPath, "synapse.db")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   common_base.GetGormLogger(logger),
		FullSaveAssociations:                     true,
	})
	if err != nil {
		return nil, fmt.Errorf("could not connect to db %s: %w", dbPath, err)
	}

	err = gdb.WithContext(ctx).AutoMigrate(base.GetAllModels()...)
	if err != nil {
		return nil, fmt.Errorf("could not migrate models: %w", err)
	}
	return &Store{base.NewStore(gdb)}, nil
}

// var _ db.TxQueueDb = &Store{}
