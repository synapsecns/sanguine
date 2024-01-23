package config

import "time"

// Config is the configuration for the screener.
type Config struct {
	// TRMKey is the api key for trmlabs
	TRMKey string `yaml:"trm-key"`
	// Rules of [caller_type]->risk_type
	Rulesets map[string]RulesetConfig `yaml:"rulesets"`
	// BlacklistURL is the url to the blacklist file
	// this is appplied to all rules and cannot be overridden
	BlacklistURL string `yaml:"blacklist-url"`
	// CacheTime is the time to cache results for (in seconds)
	// can be overridden per rulesets
	CacheTime int `yaml:"cache-time"`
	// Port is the port to listen on
	Port int `yaml:"port"`
	// Database is the database configuration
	Database DatabaseConfig `yaml:"database"`
}

// GetCacheTime gets how long to use the cache for a given ruleset.
func (c Config) GetCacheTime(rulset string) time.Duration {
	ruleset, hasRuleset := c.Rulesets[rulset]
	if !hasRuleset {
		return time.Duration(c.CacheTime) * time.Second
	}

	if ruleset.CacheTime != nil {
		return time.Duration(*ruleset.CacheTime) * time.Second
	}

	return time.Duration(c.CacheTime) * time.Second
}

// RulesetConfig is the config for each given ruleset.
type RulesetConfig struct {
	// Filename is the filename of the ruleset
	Filename string `json:"filename"`
	// CacheTime (in seconds)
	CacheTime *int `json:"cache-time"`
}

// DatabaseConfig represents the configuration for the database.
type DatabaseConfig struct {
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"` // Data Source Name
}
