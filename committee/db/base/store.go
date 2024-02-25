package base

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-datastore"
	sqlds "github.com/ipfs/go-ds-sql"
	"github.com/ipfs/go-ds-sql/sqlite"
	"github.com/synapsecns/sanguine/committee/db"
	"github.com/synapsecns/sanguine/committee/db/mysql/util"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"
	"github.com/synapsecns/sanguine/ethergo/submitter/db/txdb"
	"gorm.io/gorm"
)

// Store implements the service.
type Store struct {
	db             *gorm.DB
	submitterStore submitterDB.Service
}

// NewStore creates a new store.
func NewStore(db *gorm.DB, metrics metrics.Handler) *Store {
	txDB := txdb.NewTXStore(db, metrics)

	return &Store{
		db:             db,
		submitterStore: txDB,
	}
}

// DB gets the database object for mutation outside of the lib.
func (s Store) DB() *gorm.DB {
	return s.db
}

// SubmitterDB gets the submitter database object for mutation outside of the lib.
func (s Store) SubmitterDB() submitterDB.Service {
	return s.submitterStore
}

// DatastoreForSigner gets the datastore for a given signer.
func (s Store) DatastoreForSigner(address common.Address) (datastore.Batching, error) {
	return s.makeDatastore(fmt.Sprintf("kvs_%s", address.String()))
}

// GlobalDatastore gets the global datastore.
func (s Store) GlobalDatastore() (datastore.Batching, error) {
	return s.makeDatastore("kvs_global")
}

func (s Store) makeDatastore(name string) (datastore.Batching, error) {
	// s.DB() gets gorm db. s.DB().DB() gets the underlying db.
	underlyingDB, err := s.DB().DB()
	if err != nil {
		return nil, fmt.Errorf("could not get underlying db: %w", err)
	}

	switch s.db.Dialector.Name() {
	case dbcommon.Sqlite.String():
		if _, err := underlyingDB.Exec(fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				key TEXT PRIMARY KEY,
				data BLOB
			) WITHOUT ROWID;
		`, name)); err != nil {
			return nil, fmt.Errorf("could not ensure table exists: %w", err)
		}

		return sqlds.NewDatastore(underlyingDB, sqlite.NewQueries(name)), nil
	case dbcommon.Mysql.String():
		name = util.NamingStrategy.TableName(name)

		if _, err := underlyingDB.Exec(fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				%s VARCHAR(500) PRIMARY KEY,
				data BLOB
			);
		`, name, "`key`")); err != nil {
			return nil, fmt.Errorf("could not ensure table exists: %w", err)
		}

		return sqlds.NewDatastore(underlyingDB, util.NewQueries(name)), nil
	default:
		panic("unsupported database")
	}
}

// GetAllModels gets all models to migrate
// see: https://medium.com/@SaifAbid/slice-interfaces-8c78f8b6345d for an explanation of why we can't do this at initialization time
func GetAllModels() (allModels []interface{}) {
	allModels = append(txdb.GetAllModels(), &LastIndexed{}, &VerificationRequest{})
	return allModels
}

var _ db.Service = &Store{}
