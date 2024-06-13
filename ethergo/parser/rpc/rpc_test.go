package rpc_test

import (
	"testing"

	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/parser/rpc"
)

func TestIsBatch(t *testing.T) {
	const batchTx = "[{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"0x539\"},{\"jsonrpc\":\"2.0\",\"id\":2,\"result\":\"0x0\"}]\n"
	True(t, rpc.IsBatch([]byte(batchTx)))
}

func TestIsNullResponse(t *testing.T) {
	const nullResponse = "{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":null}\n"
	True(t, rpc.IsNullResponse([]byte(nullResponse)))
	const nonNullResponse = "{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":{\"blockNumber\": \"0x\"}}\n"
	False(t, rpc.IsNullResponse([]byte(nonNullResponse)))
}

func TestParseRPCPayload(t *testing.T) {
	t.Skip("note: these are currently covered by client tests")
}
