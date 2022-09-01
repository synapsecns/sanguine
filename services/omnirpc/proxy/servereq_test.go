package proxy_test

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/sdk/freeport"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/services/omnirpc/config"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
	"github.com/synapsecns/sanguine/services/omnirpc/rpcmap"
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
