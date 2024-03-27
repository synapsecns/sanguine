package tracely

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DenrianWeiss/tracely/model"
	"github.com/bcicen/jstream"
)

func GetTxResult(rpc, txid string) ([]model.TraceStep, error) {
	result := make([]model.TraceStep, 0)
	req := model.NewRequest(txid)
	// enable storage
	req.Params = req.Params[:len(req.Params)-1]
	// enable memory
	req.Params = append(req.Params, map[string]bool{
		"enableMemory":     true,
		"enableReturnData": true,
	})

	reqJson, _ := json.Marshal(req)
	reqStream := bytes.NewReader(reqJson)
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	post, err := client.Post(rpc, "application/json", reqStream)
	if err != nil {
		log.Println("error", err.Error())
		return nil, fmt.Errorf("error could not ")
	}
	if post.StatusCode != http.StatusOK {
		log.Println("Failed to fetch")
	}
	resultStream := post.Body

	decoder := jstream.NewDecoder(resultStream, 3)
	for mv := range decoder.Stream() {
		v := mv.Value.(map[string]interface{})
		OpCode := v["op"].(string)
		if OpCodeFocus[OpCode] {
			r := model.TraceStep{}
			if v["depth"] != nil {
				r.Depth = int(v["depth"].(float64))
			}
			r.Op = OpCode
			if v["gas"] != nil {
				r.Gas = int(v["gas"].(float64))
			}
			if v["stack"] != nil {
				r.Stack = v["stack"].([]interface{})
			}
			if v["memory"] != nil {
				r.Memory = v["memory"].([]interface{})
			}
			result = append(result, r)
		}
	}
	return result, nil
}
