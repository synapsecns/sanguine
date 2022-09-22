package dbcommon

import (
	"fmt"
	"strings"
)

// DBType is the database driver to use.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=DBType -linecomment
type DBType int

const (
	// Mysql is a mysql base db.
	Mysql DBType = 0 // mysql
	// Sqlite file based db.
	Sqlite DBType = iota // sqlite
	// Clickhouse performant db by yandex.
	Clickhouse DBType = iota // clickhouse
)

// DBTypeFromString parses a database type from a string.
func DBTypeFromString(str string) (DBType, error) {
	switch strings.ToLower(str) {
	case Mysql.String():
		return Mysql, nil
	case Sqlite.String():
		return Sqlite, nil
	case Clickhouse.String():
		return Clickhouse, nil
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
