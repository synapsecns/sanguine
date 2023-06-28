package executor

import (
	"context"
	"fmt"
)

// DBConfig is used to configure a database.
type DBConfig struct {
	// Type is the type of database. This can be either "sqlite" or "mysql".
	Type string `yaml:"type"`
	// Source is the source of the database. This can be either a path to a sqlite database or a mysql database url.
	Source string `yaml:"source,omitempty"`
}

// IsValid asserts the database connection is valid.
func (d *DBConfig) IsValid(_ context.Context) (ok bool, err error) {
	if d.Type != "sqlite" && d.Type != "mysql" {
		return false, fmt.Errorf("invalid database type: %s", d.Type)
	}

	return true, nil
}
