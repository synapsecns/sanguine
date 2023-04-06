package rpc_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/parser/rpc"
	"testing"
)

func TestIsBatch(t *testing.T) {
	const batchTx = "[{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"0x539\"},{\"jsonrpc\":\"2.0\",\"id\":2,\"result\":\"0x0\"}]\n"
	True(t, rpc.IsBatch([]byte(batchTx)))
}

func TestParseRPCPayload(t *testing.T) {
	t.Skip("note: these are currently covered by client tests")
}
