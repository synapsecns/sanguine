package config

// ParseConfig exports parseRPCMap for testing.
func ParseConfig(rawData []byte) (c Config, err error) {
	return parseConfig(rawData)
}
