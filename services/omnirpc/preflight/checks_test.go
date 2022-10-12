package preflight_test

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/services/omnirpc/preflight"
	types2 "github.com/synapsecns/sanguine/services/omnirpc/types"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
)

// TODO: deduplicate w/ confirmable, this needs to be moved into a testutil
func (p *PreflightSuite) checkRequest(makeReq func(client *ethclient.Client), checkReq func(body []byte)) {
	doneChan := make(chan bool)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		Nil(p.T(), err)

		checkReq(body)

		go func() {
			doneChan <- true
		}()
	}))

	defer server.Close()

	client, err := ethclient.DialContext(p.GetTestContext(), server.URL)
	Nil(p.T(), err)

	makeReq(client)

	<-doneChan
}

// TestCheckTransactionDontCheck tests a transaction that doesn't need to be checked
// we do this by sending funds to the 0x0000 address
func (p *PreflightSuite) TestCheckTransactionDontCheck() {
	testChainID := big.NewInt(1)
	backend := geth.NewEmbeddedBackendForChainID(p.GetTestContext(), p.T(), testChainID)
	backend.GetFundedAccount(p.GetTestContext(), big.NewInt(1000000000000000000))

	res := mocks.MockAddress()
	signedTx := backend.FaucetSignTx(types.NewTx(&types.LegacyTx{
		To:       &res,
		Value:    big.NewInt(params.GWei),
		Gas:      210000,
		GasPrice: big.NewInt(params.GWei * 15),
	}))

	//signedTx := backend.FaucetSignTx(newTx)

	p.checkRequest(func(client *ethclient.Client) {
		_ = client.SendTransaction(p.GetTestContext(), signedTx)
	}, func(body []byte) {
		req, err := types2.ParseRPCPayload(body)
		Nil(p.T(), err)

		ok, err := preflight.CheckTransaction(req, testChainID.Uint64())
		Nil(p.T(), err)
		True(p.T(), ok)
	})

}
