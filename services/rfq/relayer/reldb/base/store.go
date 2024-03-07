package base

import (
	"github.com/synapsecns/sanguine/core/metrics"
	listenerDB "github.com/synapsecns/sanguine/ethergo/listener/db"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/ethergo/submitter/db/txdb"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"gorm.io/gorm"
)

// Store implements the service.
type Store struct {
	listenerDB.ChainListenerDB
	db             *gorm.DB
	submitterStore submitterDB.Service
}

// NewStore creates a new store.
func NewStore(db *gorm.DB, metrics metrics.Handler) *Store {
	txDB := txdb.NewTXStore(db, metrics)

	return &Store{ChainListenerDB: listenerDB.NewChainListenerStore(db, metrics), db: db, submitterStore: txDB}
}

// DB gets the database object for mutation outside of the lib.
func (s Store) DB() *gorm.DB {
	return s.db
}

// SubmitterDB gets the submitter database object for mutation outside of the lib.
func (s Store) SubmitterDB() submitterDB.Service {
	return s.submitterStore
}

// GetAllModels gets all models to migrate
// see: https://medium.com/@SaifAbid/slice-interfaces-8c78f8b6345d for an explanation of why we can't do this at initialization time
func GetAllModels() (allModels []interface{}) {
	allModels = append(txdb.GetAllModels(), &RequestForQuote{}, &Rebalance{})
	allModels = append(allModels, listenerDB.GetAllModels()...)
	return allModels
}

var _ reldb.Service = &Store{}
