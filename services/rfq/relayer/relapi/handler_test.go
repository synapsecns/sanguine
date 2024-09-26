package relapi_test

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relapi"
	"testing"
)

func TestWithdrawRequestJSON(t *testing.T) {
	original := relapi.WithdrawRequest{
		ChainID:      1,
		Amount:       "100",
		TokenAddress: common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
		To:           common.HexToAddress("0xDeaDbeefdEAdbeefdEadbEEFdeadbeEFdEaDbeeF"),
	}

	// Marshal to JSON
	data, err := json.Marshal(original)
	assert.NoError(t, err)

	// Unmarshal back to struct
	var unmarshalled relapi.WithdrawRequest
	err = json.Unmarshal(data, &unmarshalled)
	assert.NoError(t, err)

	// Check if the original and unmarshalled structs are the same
	assert.Equal(t, original, unmarshalled)

	// Check the JSON string explicitly
	expectedJSON := `{"chain_id":1,"amount":"100","token_address":"0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee","to":"0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"}`
	assert.JSONEq(t, expectedJSON, string(data))
}
