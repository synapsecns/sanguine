package proxy_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"

	"github.com/lmittmann/w3/module/eth"
	ethClient "github.com/synapsecns/sanguine/ethergo/chain/client"
	"github.com/synapsecns/sanguine/ethergo/parser/rpc"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-http-utils/headers"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/ethergo/backends/geth"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
	"github.com/valyala/fasthttp"
)

// captureResponse captures the response from geth so we can use it for testing.
func (p *ProxySuite) captureResponse(backendURL string, makeReq func(client ethClient.EVMClient), checkResp func(req []rpc.Request, response proxy.JSONRPCMessage, rawResponse []byte)) {
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

		rpcReq, err := rpc.ParseRPCPayload(requestBody)
		Nil(p.T(), err)

		var rpcMessage proxy.JSONRPCMessage
		err = json.Unmarshal(fullResp, &rpcMessage)
		Nil(p.T(), err)

		if err != nil {
			fmt.Println(err)
		}

		checkResp(rpcReq, rpcMessage, fullResp)
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

	client, err := ethClient.NewClientFromChainID(p.GetTestContext(), server.URL, params.AllCliqueProtocolChanges.ChainID)
	Nil(p.T(), err)

	makeReq(client)

	<-doneChan
}

// readResponseBodyNoMutate reads a response body (decompressing if the content-encoding header)
// so specifies and then resets the reader. The response is returned.
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
// to instruct the caller not to continue proxying the http request.
func attemptAddGetBody(req *http.Request, w http.ResponseWriter) (shouldContinue bool) {
	// make a copy of the body we can re-read to get the method name
	reqBody, err := io.ReadAll(req.Body)
	// catch an error in case resquest body can't be read w/o hanging
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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

func (p *ProxySuite) TestChainID() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
		_, err := client.ChainID(p.GetTestContext())
		Nil(p.T(), err)
	}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
		standardizedResponse, err := proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
		Nil(p.T(), err)

		JSONEq(p.T(), string(standardizedResponse), string(response.Result))
	})
}

func (p *ProxySuite) TestStorageAt() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
		_, err := client.StorageAt(p.GetTestContext(), common.Address{}, common.Hash{}, nil)
		Nil(p.T(), err)
	}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
		standardizedResponse, err := proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
		Nil(p.T(), err)

		JSONEq(p.T(), string(standardizedResponse), string(response.Result))
	})
}

func (p *ProxySuite) TestCodeAt() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
		_, err := client.CodeAt(p.GetTestContext(), common.Address{}, nil)
		Nil(p.T(), err)
	}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
		standardizedResponse, err := proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
		Nil(p.T(), err)

		JSONEq(p.T(), string(standardizedResponse), string(response.Result))
	})
}

func (p *ProxySuite) TestNonceAt() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
		_, err := client.NonceAt(p.GetTestContext(), common.Address{}, nil)
		Nil(p.T(), err)
	}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
		standardizedResponse, err := proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
		Nil(p.T(), err)

		JSONEq(p.T(), string(standardizedResponse), string(response.Result))
	})
}

func (p *ProxySuite) TestEstimateGas() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	fundedAccount := backend.GetFundedAccount(p.GetTestContext(), new(big.Int).SetUint64(params.Ether))

	p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
		_, err := client.EstimateGas(p.GetTestContext(), ethereum.CallMsg{
			From:  fundedAccount.Address,
			To:    &common.Address{},
			Value: new(big.Int).SetUint64(params.Ether),
		})
		Nil(p.T(), err)
	}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
		standardizedResponse, err := proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
		Nil(p.T(), err)

		JSONEq(p.T(), string(standardizedResponse), string(response.Result))
	})
}

func (p *ProxySuite) TestGasPrice() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
		_, err := client.SuggestGasPrice(p.GetTestContext())
		Nil(p.T(), err)
	}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
		standardizedResponse, err := proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
		Nil(p.T(), err)

		JSONEq(p.T(), string(standardizedResponse), string(response.Result))
	})
}

func (p *ProxySuite) TestMaxPriority() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
		_, err := client.SuggestGasTipCap(p.GetTestContext())
		Nil(p.T(), err)
	}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
		standardizedResponse, err := proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
		Nil(p.T(), err)

		JSONEq(p.T(), string(standardizedResponse), string(response.Result))
	})
}

func (p *ProxySuite) TestBalanceAt() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	fundedAccount := backend.GetFundedAccount(p.GetTestContext(), big.NewInt(params.Ether))

	p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
		_, err := client.BalanceAt(p.GetTestContext(), fundedAccount.Address, nil)
		Nil(p.T(), err)
	}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
		standardizedResponse, err := proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
		Nil(p.T(), err)

		JSONEq(p.T(), string(standardizedResponse), string(response.Result))
	})
}

func (p *ProxySuite) TestTransactionCount() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	block, err := backend.BlockByNumber(p.GetTestContext(), nil)
	Nil(p.T(), err)

	p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
		_, err := client.TransactionCount(p.GetTestContext(), block.Hash())
		Nil(p.T(), err)
	}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
		standardizedResponse, err := proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
		Nil(p.T(), err)

		JSONEq(p.T(), string(standardizedResponse), string(response.Result))
	})
}

func (p *ProxySuite) TestPendingTransactionCount() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
		_, err := client.PendingTransactionCount(p.GetTestContext())
		Nil(p.T(), err)
	}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
		standardizedResponse, err := proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
		Nil(p.T(), err)

		JSONEq(p.T(), string(standardizedResponse), string(response.Result))
	})
}

func (p *ProxySuite) TestBlockNumber() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
		_, err := client.BlockNumber(p.GetTestContext())
		Nil(p.T(), err)
	}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
		standardizedResponse, err := proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
		Nil(p.T(), err)

		JSONEq(p.T(), string(standardizedResponse), string(response.Result))
	})
}

func (p *ProxySuite) TestBlockByHash() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	latestBlock, err := backend.BlockByNumber(p.GetTestContext(), nil)
	Nil(p.T(), err)

	const respCount = 2

	resps := make([][]byte, respCount)

	for i := 0; i < respCount; i++ {
		i := i // capture func literal
		// TODO: we should probably test txes for this as well and mock some
		p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
			_, err := client.BlockByHash(p.GetTestContext(), latestBlock.Hash())
			Nil(p.T(), err)
		}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
			resps[i], err = proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
			Nil(p.T(), err)
		})
	}

	// ensure response parity after de/re-serialization
	Equal(p.T(), resps[0], resps[1])
}

// nolint:dupl
func (p *ProxySuite) TestBlockByNumber() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	latestNumber, err := backend.BlockNumber(p.GetTestContext())
	Nil(p.T(), err)

	const respCount = 2

	resps := make([][]byte, respCount)

	for i := 0; i < respCount; i++ {
		i := i
		// TODO: we should probably test txes for this as well and mock some
		p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
			_, err := client.BlockByNumber(p.GetTestContext(), new(big.Int).SetUint64(latestNumber))
			Nil(p.T(), err)
		}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
			resps[i], err = proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
			Nil(p.T(), err)
		})
	}

	// ensure response parity after de/re-serialization
	Equal(p.T(), resps[0], resps[1])
}

// nolint:dupl
func (p *ProxySuite) TestHeaderByNumber() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	latestNumber, err := backend.BlockNumber(p.GetTestContext())
	Nil(p.T(), err)

	const respCount = 3

	resps := make([][]byte, respCount)

	for i := 0; i < respCount; i++ {
		i := i
		p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
			// TODO: we should probably test txes for this as well and mock some
			_, err := client.HeaderByNumber(p.GetTestContext(), new(big.Int).SetUint64(latestNumber))
			Nil(p.T(), err)
		}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
			resps[i], err = proxy.StandardizeResponseFalseParams(p.GetTestContext(), req, fullResp)
			Nil(p.T(), err)
		})
	}

	// ensure response parity after de/re-serialization
	Equal(p.T(), resps[0], resps[1])
}

func (p *ProxySuite) TestTransactionByHash() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	// get gas price
	gasPrice, err := backend.SuggestGasPrice(p.GetTestContext())
	Nil(p.T(), err)

	// create a fake tx to send
	testTx := backend.FaucetSignTx(types.NewTx(&types.LegacyTx{
		To:       &common.Address{},
		Value:    big.NewInt(1),
		Gas:      21000,
		GasPrice: gasPrice,
	}))

	err = backend.SendTransaction(p.GetTestContext(), testTx)
	Nil(p.T(), err)

	backend.WaitForConfirmation(p.GetTestContext(), testTx)

	const respCount = 2

	resps := make([][]byte, respCount)

	for i := 0; i < respCount; i++ {
		i := i // capture func literal
		// TODO: we should probably test txes for this as well and mock some
		p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
			_, _, err := client.TransactionByHash(p.GetTestContext(), testTx.Hash())
			Nil(p.T(), err)
		}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
			resps[i], err = proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
			Nil(p.T(), err)
		})
	}

	// ensure response parity after de/re-serialization
	Equal(p.T(), resps[0], resps[1])
}

func (p *ProxySuite) TestTransactionInBlock() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	// get gas price
	gasPrice, err := backend.SuggestGasPrice(p.GetTestContext())
	Nil(p.T(), err)

	// create a fake tx to send
	testTx := backend.FaucetSignTx(types.NewTx(&types.LegacyTx{
		To:       &common.Address{},
		Value:    big.NewInt(1),
		Gas:      21000,
		GasPrice: gasPrice,
	}))

	err = backend.SendTransaction(p.GetTestContext(), testTx)
	Nil(p.T(), err)

	backend.WaitForConfirmation(p.GetTestContext(), testTx)

	txReceipt, err := backend.TransactionReceipt(p.GetTestContext(), testTx.Hash())
	Nil(p.T(), err)

	const respCount = 2

	resps := make([][]byte, respCount)

	for i := 0; i < respCount; i++ {
		i := i // capture func literal
		// TODO: we should probably test txes for this as well and mock some
		p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
			_, err := client.TransactionInBlock(p.GetTestContext(), txReceipt.BlockHash, txReceipt.TransactionIndex)
			Nil(p.T(), err)
		}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
			resps[i], err = proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
			Nil(p.T(), err)
		})
	}

	// ensure response parity after de/re-serialization
	Equal(p.T(), resps[0], resps[1])
}

func (p *ProxySuite) TestTransactionReceipt() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	// get gas price
	gasPrice, err := backend.SuggestGasPrice(p.GetTestContext())
	Nil(p.T(), err)

	// create a fake tx to send
	testTx := backend.FaucetSignTx(types.NewTx(&types.LegacyTx{
		To:       &common.Address{},
		Value:    big.NewInt(1),
		Gas:      21000,
		GasPrice: gasPrice,
	}))

	err = backend.SendTransaction(p.GetTestContext(), testTx)
	Nil(p.T(), err)

	backend.WaitForConfirmation(p.GetTestContext(), testTx)

	const respCount = 2

	resps := make([][]byte, respCount)

	for i := 0; i < respCount; i++ {
		i := i // capture func literal
		// TODO: we should probably test txes for this as well and mock some
		p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
			_, err := client.TransactionReceipt(p.GetTestContext(), testTx.Hash())
			Nil(p.T(), err)
		}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
			resps[i], err = proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
			Nil(p.T(), err)
		})
	}

	// ensure response parity after de/re-serialization
	Equal(p.T(), resps[0], resps[1])
}

func (p *ProxySuite) TestSyncProgress() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	const respCount = 2

	resps := make([][]byte, respCount)

	for i := 0; i < respCount; i++ {
		i := i // capture func literal
		// TODO: we should probably test txes for this as well and mock some
		p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
			_, err := client.SyncProgress(p.GetTestContext())
			Nil(p.T(), err)
		}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
			var err error
			resps[i], err = proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
			Nil(p.T(), err)
		})
	}

	// ensure response parity after de/re-serialization
	Equal(p.T(), resps[0], resps[1])
}

func (p *ProxySuite) TestGetLogsMethod() {
	p.T().Skip("TODO: we're not going to touch this until we can do it properly w/ a testutil")
}

func (p *ProxySuite) TestFeeHistory() {
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	// get gas price
	gasPrice, err := backend.SuggestGasPrice(p.GetTestContext())
	Nil(p.T(), err)

	// create a fake tx to send to populate some data
	testTx := backend.FaucetSignTx(types.NewTx(&types.LegacyTx{
		To:       &common.Address{},
		Value:    big.NewInt(1),
		Gas:      21000,
		GasPrice: gasPrice,
	}))

	err = backend.SendTransaction(p.GetTestContext(), testTx)
	Nil(p.T(), err)

	backend.WaitForConfirmation(p.GetTestContext(), testTx)

	lastBlock, err := backend.BlockNumber(p.GetTestContext())
	Nil(p.T(), err)

	const respCount = 2

	resps := make([][]byte, respCount)

	for i := 0; i < respCount; i++ {
		i := i // capture func literal
		// TODO: we should probably test txes for this as well and mock some
		p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
			_, err := client.FeeHistory(p.GetTestContext(), lastBlock, new(big.Int).SetUint64(lastBlock), []float64{95, 99})
			Nil(p.T(), err)
		}, func(req []rpc.Request, response proxy.JSONRPCMessage, fullResp []byte) {
			resps[i], err = proxy.StandardizeResponse(p.GetTestContext(), req, fullResp)
			Nil(p.T(), err)
		})
	}

	// ensure response parity after de/re-serialization
	Equal(p.T(), resps[0], resps[1])
}

func (p *ProxySuite) TestBatch() {
	p.T().Skip("TODO: this works, we need to modify around it for the test to pass (our captureResponse method breaks). This is currently tested in scribe, but should be tested here.")
	backend := geth.NewEmbeddedBackend(p.GetTestContext(), p.T())

	var chainID uint64
	var blockNumber big.Int

	p.captureResponse(backend.HTTPEndpoint(), func(client ethClient.EVMClient) {
		err := client.BatchContext(p.GetTestContext(), eth.ChainID().Returns(&chainID), eth.BlockNumber().Returns(&blockNumber))
		Nil(p.T(), err)
	}, func(req []rpc.Request, response proxy.JSONRPCMessage, rawResponse []byte) {
		Greater(p.T(), blockNumber, 1)
		Equal(p.T(), chainID, params.AllCliqueProtocolChanges.ChainID.Uint64())
	})
}
