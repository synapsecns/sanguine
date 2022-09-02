package proxy_test

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/sdk/freeport"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/omnirpc/config"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
	"github.com/synapsecns/sanguine/services/omnirpc/rpcmap"
	"io"
	"net/http"
	"net/http/httptest"
)

func (p *ProxySuite) TestServeRequestNoRPCs() {
	rpcMap := rpcmap.NewRPCMap()
	cfg := config.Config{
		//nolint: staticcheck
		Port: uint16(freeport.GetT(p.T(), 1)[0]),
	}

	prxy := proxy.NewProxy(rpcMap, cfg)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	prxy.ServeRPCReq(c, 1)
	Equal(p.T(), w.Code, http.StatusBadRequest)
}

func (p *ProxySuite) MustMarshall(v any) []byte {
	res, err := json.Marshal(v)
	Nil(p.T(), err)

	return res
}

func (p *ProxySuite) TestParseRPCPayload() {
	doneChan := make(chan bool)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		Nil(p.T(), err)

		method, err := proxy.ParseRPCPayload(body)
		Nil(p.T(), err)

		Equal(p.T(), "eth_getBlockByNumber", method)
		go func() {
			doneChan <- true
		}()
	}))

	client, err := ethclient.DialContext(p.GetTestContext(), server.URL)
	Nil(p.T(), err)

	_, _ = client.HeaderByNumber(p.GetTestContext(), nil)

	<-doneChan
}
