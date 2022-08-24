package sql

import (
	"context"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/agents/db"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql/mysql"
	"github.com/synapsecns/sanguine/agents/db/datastore/sql/sqlite"
	"strings"
)

// NewStoreFromConfig creates a new datastore from a config file.
// nolint: wrapcheck
func NewStoreFromConfig(ctx context.Context, dbType DBType, connString string) (db.SynapseDB, error) {
	switch dbType {
	case Mysql:
		return mysql.NewMysqlStore(ctx, connString)
	case Sqlite:
		return sqlite.NewSqliteStore(ctx, connString)
	default:
		return nil, ErrNoSuchDriver
	}
}

// DBType is the database driver to use.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=DBType -linecomment
type DBType int

const (
	// Mysql is a mysql base db.
	Mysql DBType = 0 // mysql
	// Sqlite file based db.
	Sqlite DBType = iota // sqlite
)

// DBTypeFromString parses a database type from a string.
func DBTypeFromString(str string) (DBType, error) {
	switch strings.ToLower(str) {
	case Mysql.String():
		return Mysql, nil
	case Sqlite.String():
		return Sqlite, nil
	default:
		return DBType(-1), fmt.Errorf("could not convert %s to %T, must be one of %s", str, DBType(-1), allDBTypesList())
	}
}

// AllDBTypes is a list of all contract types. Since we use stringer and this is a testing library, instead
// of manually copying all these out we pull the names out of stringer. In order to make sure stringer is updated, we panic on
// any method called where the index is higher than the stringer array length.
var AllDBTypes []DBType

// set all contact types.
func init() {
	for i := 0; i < len(_DBType_index)-1; i++ {
		dbType := DBType(i)
		AllDBTypes = append(AllDBTypes, dbType)

		// statically assert these are all lowercase
		if dbType.String() != strings.ToLower(dbType.String()) {
			panic(fmt.Errorf("db type %s is not lowercase", dbType))
		}
	}
}

// allDBTypesList prints a list of all db types. This is useful for returning errors.
func allDBTypesList() string {
	var res []string
	for _, signerType := range AllDBTypes {
		res = append(res, signerType.String())
	}

	return strings.Join(res, ",")
}

// ErrNoSuchDriver indicates the driver does not exist.
var ErrNoSuchDriver = errors.New("no such db driver")
