package near_test

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/jarcoal/httpmock"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/chain/client/near"
	"math/big"
	"net/http"
)

type testRPCClient struct {
	*rpc.Client
}

func (t testRPCClient) TransactionByHash(ctx context.Context, txHash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	panic("not required for this test")
}

func (t testRPCClient) BlockNumber(ctx context.Context) (uint64, error) {
	panic("not required for this test")
}

func (c NearSuite) newTestRPCClient() testRPCClient {
	testClient, err := rpc.DialContext(c.GetTestContext(), "https://mainnet.aurora.dev/")
	Nil(c.T(), err)

	return testRPCClient{testClient}
}

// TestGetBlock tests the get block method.
func (c NearSuite) TestGetBlockByNumber() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// return the output of
	// curl --location --request POST 'https://mainnet.aurora.dev/' \
	// --header 'Content-Type: application/json' \
	// --data-raw '{
	//  "jsonrpc":"2.0",
	// 	"method":"eth_getBlockByNumber",
	// 	"params":[
	//	"0x1b4",
	// 	true
	// ],
	// "id":1
	// }'
	httpmock.RegisterResponder(http.MethodPost, "https://mainnet.aurora.dev/",
		httpmock.NewStringResponder(http.StatusOK, "{\"jsonrpc\":\"2.0\",\"id\":1,\"result\":{\"number\":\"0x1b4\",\"hash\":\"0x69e68ca2a78332dfb0e0b861b44879b6f32bbd37570723054c53e8d8fb2d2000\",\"parentHash\":\"0x61b46b9615c20594bf55d240f4c936d4a8b2235cccb22e7fb65249d517bb1f40\",\"nonce\":\"0x0000000000000000\",\"sha3Uncles\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"logsBloom\":\"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000\",\"transactionsRoot\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"stateRoot\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"receiptsRoot\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"miner\":\"0x0000000000000000000000000000000000000000\",\"difficulty\":\"0x0\",\"totalDifficulty\":\"0x0\",\"extraData\":\"0x\",\"size\":\"0x0\",\"gasLimit\":\"0x0\",\"gasUsed\":\"0x0\",\"timestamp\":\"0x0\",\"uncles\":[],\"transactions\":[]}}"))

	block, err := near.BlockByNumber(c.GetTestContext(), c.newTestRPCClient(), big.NewInt(436))
	Nil(c.T(), err)

	_ = block
}
