package db

import "context"

const (
	// clickhouse type.
	Clickhouse = "clickhouse"
)

// Config is the configuration to use for the db.
type Config struct {
	// Type: mysql or sqlite
	Type string `toml:"Type"`
	// URL: url for mysql or path for sqlite
	URL string `toml:"URL"`
}

// IsValid makes sure the p2p config is valid
// TODO find out if there are any invalid conditions here.
func (n *Config) IsValid(_ context.Context) (ok bool, err error) {
	return true, nil
}
