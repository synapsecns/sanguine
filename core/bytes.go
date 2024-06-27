package core

import (
	"encoding/json"
)

// BytesToSlice converts a 32 bit array to a slice slice.
func BytesToSlice(bytes [32]byte) []byte {
	rawBytes := make([]byte, len(bytes))
	copy(rawBytes, bytes[:])
	return rawBytes
}

// BytesToJSONString converts a 32 bit array to a JSON string without escapes, newlines, etc.
func BytesToJSONString(bz []byte) string {
	var jsonData map[string]interface{}
	if err := json.Unmarshal(bz, &jsonData); err != nil {
		return ""
	}
	formattedJson, err := json.Marshal(jsonData)
	if err != nil {
		return ""
	}

	return string(formattedJson)
}
