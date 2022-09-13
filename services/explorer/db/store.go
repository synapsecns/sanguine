package db

import (
	"context"
	"errors"
	"github.com/synapsecns/sanguine/agents/db"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/services/explorer/db/sql/clickhouse"
)

// NewStoreFromConfig creates a new database from a config file.
func NewStoreFromConfig(ctx context.Context, dbType dbcommon.DBType, connString string) (db.EventDB, error) {
	switch dbType {
	case dbcommon.Clickhouse:
		//nolint:wrapcheck
		return clickhouse.NewClickhouseStore(ctx, connString)
	default:
		return nil, ErrNoSuchDriver
	}
}

// ErrNoSuchDriver indicates that the driver is not supported.
var ErrNoSuchDriver = errors.New("no such db driver")
