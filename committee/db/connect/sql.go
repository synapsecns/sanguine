// Package connect contains the database connection logic for the RFQ relayer.
// TODO: this is a dumb name for a package in a dumb place. Move it somewhere else.
package connect

import (
	"context"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/committee/db"
	"github.com/synapsecns/sanguine/committee/db/base"
	"github.com/synapsecns/sanguine/committee/db/mysql"
	"github.com/synapsecns/sanguine/committee/db/sqlite"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
)

// Connect connects to the database.
func Connect(ctx context.Context, dbType dbcommon.DBType, path string, metrics metrics.Handler, rawTXDecoder base.RawTransactionDecoder) (db.Service, error) {
	switch dbType {
	case dbcommon.Mysql:
		store, err := mysql.NewMysqlStore(ctx, path, metrics, rawTXDecoder)
		if err != nil {
			return nil, fmt.Errorf("could not create mysql store: %w", err)
		}

		return store, nil
	case dbcommon.Sqlite:
		store, err := sqlite.NewSqliteStore(ctx, path, metrics, rawTXDecoder)
		if err != nil {
			return nil, fmt.Errorf("could not create sqlite store: %w", err)
		}

		return store, nil
	case dbcommon.Clickhouse:
		return nil, errors.New("driver not supported")
	default:
		return nil, fmt.Errorf("unsupported driver: %s", dbType)
	}
}
