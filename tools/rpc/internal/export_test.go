package internal

// ParseRPCMap exports parseRPCMap for testing.
func ParseRPCMap(rawData []byte) (m RPCMap, err error) {
	return parseRPCMap(rawData)
}
