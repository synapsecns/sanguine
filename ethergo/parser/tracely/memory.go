package tracely

import "encoding/hex"

func AccessMemory(m []byte, offset, size int) []byte {
	return m[offset : offset+size]
}

func GenerateMemory(m []interface{}) []byte {
	result := make([]byte, 0)
	for _, v := range m {
		l := v.(string)
		d, _ := hex.DecodeString(l)
		result = append(result, d...)
	}
	return result
}
