package config

// Config is the configuration for the screener
type Config struct {
	// TRMKey is the api key for trmlabs
	TRMKey string `json:"trm-key"`
	// Rules of [caller_type]->risk_type
	Rules map[string]string `json:"rules"`
	// BlacklistURL is the url to the blacklist file
	// this is appplied to all rules and cannot be overridden
	BlacklistURL string `json:"blacklist-url"`
	// CacheTime is the time to cache results for (in minutes)
	CacheTime int `json:"cache-time"`
	// Port is the port to listen on
	Port int `json:"port"`
}
