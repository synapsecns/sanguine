package proxy

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func StandardizeResponse(method string, body []byte) (out []byte, err error) {

	switch method {
	case ChainIDMethod:
		var result hexutil.Big
		err := json.Unmarshal(body, &result)
		if err != nil {
			return nil, fmt.Errorf("could not parse")
		}
		out, err = json.Marshal(result)
	}

	return out, nil
}
