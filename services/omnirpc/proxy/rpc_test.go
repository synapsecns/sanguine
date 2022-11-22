package proxy_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
	"testing"
)

func TestIsBatch(t *testing.T) {
	const batchTx = "[{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":\"0x539\"},{\"jsonrpc\":\"2.0\",\"id\":2,\"result\":\"0x0\"}]\n"
	True(t, proxy.IsBatch([]byte(batchTx)))
}
