package config

// ParseRPCMap exports parseRPCMap for testing.
func ParseRPCMap(rawData []byte) (m RPCConfig, err error) {
	return parseRPCMap(rawData)
}
