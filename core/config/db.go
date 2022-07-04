package config

import "context"

// DBConfig is used for configuring the.
type DBConfig struct {
	// Type of the database to use for sql. This does not affect hte pebble db
	Type string `toml:"Type"`
	// DBPath is the db path used for the pebble db
	DBPath string `toml:"DBPath"`
	// ConnString is the connection string used for mysql
	ConnString string `toml:"ConnString"`
}

func (d *DBConfig) IsValid(ctx context.Context) (ok bool, err error) {
	// TODO: there's no sense implementing this until we have a way to
	//  discriminate between networked and non-networked checks

	return true, nil
}
