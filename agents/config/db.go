package config

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/dbcommon"
)

// DBConfig is used to configure a database.
type DBConfig struct {
	// Type is the type of database. This can be either "sqlite" or "mysql".
	Type string `yaml:"type"`
	// Source is the source of the database. This can be either a path to a sqlite database or a mysql database url.
	Source string `yaml:"source"`
}

// IsValid asserts the database connection is valid.
func (d *DBConfig) IsValid(_ context.Context) (ok bool, err error) {
	if d.Type != dbcommon.Sqlite.String() && d.Type != dbcommon.Mysql.String() {
		return false, fmt.Errorf("invalid database type: %s", d.Type)
	}

	return true, nil
}
