package sql

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm/schema"

	"github.com/synapsecns/sanguine/agents/db"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql/mysql"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql/sqlite"
	"github.com/synapsecns/sanguine/core/dbcommon"
)

// NewStoreFromConfig creates a new datastore from a config file.
//
//nolint:wrapcheck
func NewStoreFromConfig(ctx context.Context, dbType dbcommon.DBType, connString string, prefix string) (db.SynapseDB, error) {
	switch dbType {
	case dbcommon.Mysql:
		if prefix != "" {
			mysql.NamingStrategy = schema.NamingStrategy{
				TablePrefix: fmt.Sprintf("%s_", prefix),
			}
		}

		return mysql.NewMysqlStore(ctx, connString)
	case dbcommon.Sqlite:
		return sqlite.NewSqliteStore(ctx, connString)
	case dbcommon.Clickhouse:
		return nil, ErrNoSuchDriver
	default:
		return nil, ErrNoSuchDriver
	}
}

// ErrNoSuchDriver indicates the driver does not exist.
var ErrNoSuchDriver = errors.New("no such db driver")
