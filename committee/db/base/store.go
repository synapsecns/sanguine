package base

import (
	"context"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"github.com/synapsecns/sanguine/committee/db"
	"github.com/synapsecns/sanguine/core/metrics"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/ethergo/submitter/db/txdb"
	"gorm.io/gorm"
)

// Store implements the service.
type Store struct {
	db             *gorm.DB
	submitterStore submitterDB.Service
	rawTxDecoder   RawTransactionDecoder
}

func NewStore(db *gorm.DB, metrics metrics.Handler, rawTxDecoder RawTransactionDecoder) *Store {
	txDB := txdb.NewTXStore(db, metrics)

	return &Store{db: db, submitterStore: txDB}
}

type RawTransactionDecoder func(ctx context.Context, data []byte) (synapsemodule.InterchainInterchainTransaction, error)

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
	allModels = append(txdb.GetAllModels(), &LastIndexed{}, &SignRequest{}, &Signature{})
	return allModels
}

var _ db.Service = &Store{}
