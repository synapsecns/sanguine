package proxy_test

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-http-utils/headers"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
	"github.com/valyala/fasthttp"
	"io"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
)

// captureResponse captures the response from geth so we can use it for testing
func (p *ProxySuite) captureResponse(backendURL string, makeReq func(client *ethclient.Client), checkResp func(method string, response []byte)) {
	doneChan := make(chan bool)

	parsedURL, err := url.Parse(backendURL)
	Nil(p.T(), err)

	rp := httputil.NewSingleHostReverseProxy(parsedURL)

	rp.ModifyResponse = func(response *http.Response) error {
		fullResp, err := readResponseBodyNoMutate(response)
		Nil(p.T(), err)

		reqBodyReader, err := response.Request.GetBody()
		Nil(p.T(), err)

		requestBody, err := io.ReadAll(reqBodyReader)
		Nil(p.T(), err)

		rpcReq, err := proxy.ParseRPCPayload(requestBody)
		Nil(p.T(), err)

		checkResp(rpcReq.Method, fullResp)
		return nil
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		shouldProxy := attemptAddGetBody(r, w)
		if !shouldProxy {
			return
		}

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

// readResponseBodyNoMutate reads a response body (decompressing if the content-encoding header)
// so specifies and then resets the reader. The response is returned
func readResponseBodyNoMutate(response *http.Response) (res []byte, err error) {
	fullResp, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	response.Body = io.NopCloser(bytes.NewReader(fullResp))

	// use fasthttp hhere since go doesn't export the transport decompression methods
	decompressor := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(decompressor)
	decompressor.Header.SetContentEncoding(response.Header.Get(headers.ContentEncoding))
	decompressor.SetBodyRaw(fullResp)
	uncompressedBody, err := decompressor.BodyUncompressed()
	if err != nil {
		return nil, fmt.Errorf("could not decompress header: %w", err)
	}

	return uncompressedBody, nil
}

// attemptAddGetBody attempts to add a get body method to the request
// in the case that this fails, we return an error response and an error bool
// to instruct the caller not to continue proxying the http request
func attemptAddGetBody(req *http.Request, w http.ResponseWriter) (shouldContinue bool) {
	// make a copy of the body we can re-read to get the method name
	reqBody, err := io.ReadAll(req.Body)
	// catch an error in case resquest body can't be read w/o hanging
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("cannot read response"))
		return false
	}

	// create a new body on demand for testing
	req.GetBody = func() (io.ReadCloser, error) {
		return io.NopCloser(bytes.NewReader(reqBody)), nil
	}

	// reset the body reader, this is required by httputil
	req.Body = io.NopCloser(bytes.NewReader(reqBody))
	return true
}

func (p *ProxySuite) TestBlockNumber() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	p.captureResponse(backend.HTTPEndpoint(), func(client *ethclient.Client) {
		_, err := client.ChainID(p.GetTestContext())
		Nil(p.T(), err)

	}, func(method string, response []byte) {
		standardizedResponse, err := proxy.StandardizeResponse(method, response)
		Nil(p.T(), err)

		_ = standardizedResponse
		// TODO: compare standardizedResponse against the actual result

		//JSONEq(p.T(), string(standardizedResponse), string(response))

	})
}
