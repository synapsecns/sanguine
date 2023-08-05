package base

import (
	"github.com/synapsecns/sanguine/core/metrics"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/ethergo/submitter/db/txdb"
	"github.com/synapsecns/sanguine/services/cctp-relayer/db"
	"github.com/synapsecns/sanguine/services/cctp-relayer/types"
	"gorm.io/gorm"
)

// Store is a store that implements an underlying gorm db.
type Store struct {
	db             *gorm.DB
	metrics        metrics.Handler
	submitterStore submitterDB.Service
}

// NewStore creates a new store.
func NewStore(db *gorm.DB, metrics metrics.Handler) *Store {
	txDB := txdb.NewTXStore(db, metrics)
	return &Store{db: db, metrics: metrics, submitterStore: txDB}
}

// DB gets the database object for mutation outside of the lib.
func (s Store) DB() *gorm.DB {
	return s.db
}

// SubmitterDB gets the submitter database object for mutation outside of the lib.
func (s Store) SubmitterDB() submitterDB.Service {
	return s.submitterStore
}

// GetAllModels gets all models to migrate.
// see: https://medium.com/@SaifAbid/slice-interfaces-8c78f8b6345d for an explanation of why we can't do this at initialization time
func GetAllModels() (allModels []interface{}) {
	allModels = append(allModels, txdb.GetAllModels()...)
	allModels = append(allModels, &types.Message{})
	return allModels
}

var _ db.CCTPRelayerDB = &Store{}
