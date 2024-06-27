package config

// Config is the configuration for the screener.
type Config struct {
	// AppSecret is the app secret
	AppSecret string `yaml:"app-secret"`
	// AppID is the app id
	AppID string `yaml:"app-id"`
	// ChainalysisKey is the api key for chainalysis
	ChainalysisKey string `yaml:"chainalysis-key"`
	// ChainalysisURL is the url for chainalysis
	ChainalysisURL string `yaml:"chainalysis-url"`
	// BlacklistURL is the url to the blacklist file
	// this is applied to all rules and cannot be overridden
	BlacklistURL string `yaml:"blacklist-url"`
	// Port is the port to listen on
	Port int `yaml:"port"`
	// Database is the database configuration
	Database DatabaseConfig `yaml:"database"`
	// Severities are the severity levels for each address we want to screen
	RiskLevels []string `yaml:"risk-levels"`
	// Whitelist is a list of addresses to whitelist
	Whitelist []string `yaml:"whitelist"`
}

// DatabaseConfig represents the configuration for the database.
type DatabaseConfig struct {
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"` // Data Source Name
}
