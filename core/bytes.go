package core

import (
	"encoding/json"
	"fmt"
)

// BytesToSlice converts a 32 bit array to a slice slice.
func BytesToSlice(bytes [32]byte) []byte {
	rawBytes := make([]byte, len(bytes))
	copy(rawBytes, bytes[:])
	return rawBytes
}

// BytesToJSONString converts a 32 bit array to a JSON string without escapes, newlines, etc.
func BytesToJSONString(bz []byte) (string, error) {
	var jsonData map[string]interface{}
	if err := json.Unmarshal(bz, &jsonData); err != nil {
		return "", fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	formattedJSON, err := json.Marshal(jsonData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}

	return string(formattedJSON), nil
}
func BytesToArray(bz []byte) ([32]byte, error) {
	var bytes [32]byte
	if len(bz) != 32 {
		return bytes, fmt.Errorf("invalid length of bytes: %d", len(bz))
	}
	copy(bytes[:], bz)
	return bytes, nil
}
