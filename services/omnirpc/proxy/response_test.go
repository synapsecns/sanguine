package proxy_test

import (
	"bytes"
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
)

// captureResponse captures the response from geth so we can use it for testing
func (p *ProxySuite) captureResponse(backendURL string, makeReq func(client *ethclient.Client), checkResp func(response []byte)) {
	doneChan := make(chan bool)

	parsedURL, err := url.Parse(backendURL)
	Nil(p.T(), err)

	rp := httputil.NewSingleHostReverseProxy(parsedURL)
	rp.ModifyResponse = func(r *http.Response) error {
		fullResp, err := io.ReadAll(r.Body)
		Nil(p.T(), err)

		r.Body = io.NopCloser(bytes.NewReader(fullResp))

		checkResp(fullResp)
		return nil
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rp.ServeHTTP(w, r)

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

func (p *ProxySuite) TestBlockNumber() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	p.captureResponse(backend.HTTPEndpoint(), func(client *ethclient.Client) {
		_, err := client.ChainID(p.GetTestContext())
		Nil(p.T(), err)

	}, func(response []byte) {
		proxy.StandardizeResponse(response)

	})
}
