package config

import "context"

// DBConfig is used for configuring the.
type DBConfig struct {
	// Type of the database to use for sql. This does not affect hte pebble db
	Type string `yaml:"type"`
	// DBPath is the db path used for the pebble db
	DBPath string `yaml:"db_path"`
	// ConnString is the connection string used for mysql
	ConnString string `yaml:"conn_string"`
}

// IsValid asserts the database connection is valid.
func (d *DBConfig) IsValid(_ context.Context) (ok bool, err error) {
	// TODO: there's no sense implementing this until we have a way to
	//  discriminate between networked and non-networked checks

	return true, nil
}
