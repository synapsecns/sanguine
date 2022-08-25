package sql

import (
	"context"
	"errors"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/scribe/db"
	"github.com/synapsecns/sanguine/scribe/db/datastore/sql/mysql"
	"github.com/synapsecns/sanguine/scribe/db/datastore/sql/sqlite"
)

// NewStoreFromConfig creates a new database from a config file.
func NewStoreFromConfig(ctx context.Context, dbType dbcommon.DBType, connString string) (db.EventDB, error) {
	switch dbType {
	case dbcommon.Mysql:
		//nolint:wrapcheck
		return mysql.NewMysqlStore(ctx, connString)
	case dbcommon.Sqlite:
		//nolint:wrapcheck
		return sqlite.NewSqliteStore(ctx, connString)
	default:
		return nil, ErrNoSuchDriver
	}
}

// ErrNoSuchDriver indicates that the driver is not supported.
var ErrNoSuchDriver = errors.New("no such db driver")
