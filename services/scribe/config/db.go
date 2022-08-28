package config

import "context"

// DBConfig is used for configuring the.
type DBConfig struct {
	// Type of the database to use for sql.
	Type string `toml:"Type"`
	// ConnString is the connection string used for mysql
	ConnString string `toml:"ConnString"`
}

// IsValid asserts the database connection is valid.
func (d *DBConfig) IsValid(_ context.Context) (ok bool, err error) {
	// TODO: there's no sense implementing this until we have a way to
	//  discriminate between networked and non-networked checks

	return true, nil
}
