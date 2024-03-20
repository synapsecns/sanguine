package tracely

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DenrianWeiss/tracely/service"
	"log"
	"net/http"
	"time"

	"github.com/DenrianWeiss/tracely/model"
	"github.com/bcicen/jstream"
)

// GetTxResult fetches the result of a transaction
func GetTxResult(rpc, txid string) ([]model.TraceStep, error) {
	result := make([]model.TraceStep, 0)
	req := model.NewRequest(txid)
	reqJson, _ := json.Marshal(req)
	reqStream := bytes.NewReader(reqJson)
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	// TODO: use request w/ context
	post, err := client.Post(rpc, "application/json", reqStream)
	if err != nil {
		log.Println("error", err.Error())
		return nil, fmt.Errorf("error could not fetch: %w", err)
	}
	if post.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error could not fetch: %w", err)
	}
	resultStream := post.Body
	decoder := jstream.NewDecoder(resultStream, 3)
	for mv := range decoder.Stream() {
		v, ok := mv.Value.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("could not cast %T to map[string]interface{}", mv.Value)
		}
		OpCode, ok := v["op"].(string)
		if !ok {
			return nil, fmt.Errorf("could not cast %T to string", v["op"])
		}
		if service.OpCodeFocus[OpCode] {
			r := model.TraceStep{}
			depth, ok := v["depth"].(float64)
			if !ok {
				return nil, fmt.Errorf("could not cast %T to float64", v["depth"])
			}
			r.Depth = int(depth)
			r.Op = OpCode
			gas, ok := v["gas"].(float64)
			if !ok {
				return nil, fmt.Errorf("could not cast %T to float64", v["gas"])
			}
			r.Gas = int(gas)
			r.Stack, ok = v["stack"].([]interface{})
			if !ok {
				return nil, fmt.Errorf("could not cast %T to []interface{}", v["stack"])
			}
			r.Memory, ok = v["memory"].([]interface{})
			if !ok {
				return nil, fmt.Errorf("could not cast %T to []interface{}", v["memory"])
			}
			result = append(result, r)
		}

	}
	return result, nil
}
