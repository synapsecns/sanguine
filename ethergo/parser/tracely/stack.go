package tracely

import "github.com/ethereum/go-ethereum/common/hexutil"

func AccessStack(s [][]byte, index int) []byte {
	length := len(s)
	result := s[length-1-index]
	return result
}

func GenerateStack(s []interface{}) [][]byte {
	result := make([][]byte, 0)
	for _, v := range s {
		l := v.(string)
		lBytes, _ := hexutil.Decode(l)
		result = append(result, lBytes)
	}
	return result
}
