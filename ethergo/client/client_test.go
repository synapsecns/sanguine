package client_test

import (
	"context"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/instrumentation"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/mocks"
	"github.com/synapsecns/sanguine/ethergo/parser/rpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	_ "go.opentelemetry.io/otel/sdk/trace/tracetest"
)

// TestEVM is an EVM that can be used for testing.
type TestEVM interface {
	client.EVM
	GetMetrics() metrics.Handler
}

func (c *ClientSuite) TestParseCalls() {
	calls := make([]w3types.Caller, 4)
	chainID := new(uint64)
	calls[0] = eth.ChainID().Returns(chainID)
	maxHeight := new(big.Int)
	calls[1] = eth.BlockNumber().Returns(maxHeight)
	var mockHash common.Hash
	calls[2] = eth.StorageAt(mocks.MockAddress(), mocks.NewMockHash(c.T()), big.NewInt(1)).Returns(&mockHash)
	filter := ethereum.FilterQuery{
		Addresses: []common.Address{mocks.MockAddress()},
	}
	calls[3] = eth.Logs(filter).Returns(nil)
	res := client.ParseCalls(calls)
	Equal(c.T(), res.Value.AsStringSlice(), []string{"eth_chainId", "eth_blockNumber", "eth_getStorageAt", "eth_getLogs"})
}

// checkRequest is a helper method for checking requests.
func (c *ClientSuite) checkRequest(makeReq func(client TestEVM)) {
	ctx, cancel := context.WithCancel(c.GetTestContext())
	defer cancel()
	doneChan := make(chan bool)

	mockTracer := metrics.NewTestTracer(ctx, c.T())

	var rpcRequest rpc.Requests

	var body []byte
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		body, err = io.ReadAll(r.Body)
		c.Require().NoError(err, "failed to read request body")

		rpcRequest, err = rpc.ParseRPCPayload(body)
		c.Require().NoError(err, "failed to parse rpc payload")

		go func() {
			doneChan <- true
		}()
	}))

	defer server.Close()

	evmClient, err := client.DialBackend(ctx, server.URL, mockTracer, client.Capture(true))
	c.Require().NoError(err)

	castClient, ok := evmClient.(TestEVM)
	c.Require().True(ok, "failed to cast client to TestEVM")

	makeReq(castClient)

	<-doneChan

	spans := mockTracer.GetSpansByName(rpcRequest.Method())
	requestSpans := mockTracer.GetSpansByName(instrumentation.RequestSpanName)
	// make sure we got at most 1 span
	c.Require().Equal(len(spans), 1, "expected 1 span, got %d", len(spans))
	span := spans[0]

	c.Require().Equal(len(requestSpans), 1, "expected 1 request span, got %d", len(spans))
	requestSpan := requestSpans[0]

	// make sure the span has an exception
	c.Require().True(spanHasException(span), "expected exception event, got none")
	Equal(c.T(), spanAttributeByName(span, "endpoint").AsString(), server.URL)
	Equal(c.T(), spanEventByName(requestSpan, instrumentation.RequestEventName).AsString(), string(body))
}

func spanEventByName(stub tracetest.SpanStub, name string) *attribute.Value {
	for _, event := range stub.Events {
		if event.Name == name {
			return &event.Attributes[0].Value
		}
	}
	return nil
}

func spanAttributeByName(stub tracetest.SpanStub, name string) *attribute.Value {
	for _, attr := range stub.Attributes {
		if attr.Key == attribute.Key(name) {
			return &attr.Value
		}
	}
	return nil
}

func spanHasException(stub tracetest.SpanStub) bool {
	for _, event := range stub.Events {
		if event.Name == "exception" {
			return true
		}
	}
	return false
}

// Test parsing.
//
//nolint:maintidx
func (c *ClientSuite) TestParseRPCPayload() {
	/*
	  CHECK BLOCKS
	*/

	// check latest block, should not be confirmable since
	// rpcs might be on different  latest heights
	c.checkRequest(func(client TestEVM) {
		_, _ = client.BlockByNumber(c.GetTestContext(), nil)
	})

	// check non-latest block, should be confirmable
	c.checkRequest(func(client TestEVM) {
		_, _ = client.BlockByNumber(c.GetTestContext(), new(big.Int).SetUint64(gofakeit.Uint64()))
	})

	/*
	  CHECK HEADERS
	*/

	// check non-latest block, should be confirmable
	c.checkRequest(func(client TestEVM) {
		_, _ = client.HeaderByNumber(c.GetTestContext(), new(big.Int).SetUint64(gofakeit.Uint64()))
	})

	// eth_blockNumber should not be confirmable, can differ based on rpc
	c.checkRequest(func(client TestEVM) {
		_, _ = client.BlockNumber(c.GetTestContext())
	})

	/*
	  CHECK Transaction Methods
	*/
	c.checkRequest(func(client TestEVM) {
		_, _, _ = client.TransactionByHash(c.GetTestContext(), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())))
	})

	c.checkRequest(func(client TestEVM) {
		_, _ = client.TransactionCount(c.GetTestContext(), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())))
	})

	c.checkRequest(func(client TestEVM) {
		_, _ = client.PendingTransactionCount(c.GetTestContext())
	})

	c.checkRequest(func(client TestEVM) {
		_, _ = client.TransactionInBlock(c.GetTestContext(), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())), 1)
	})

	c.checkRequest(func(client TestEVM) {
		_, _ = client.TransactionReceipt(c.GetTestContext(), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())))
	})

	/*
	  Sync Methods
	*/

	c.checkRequest(func(client TestEVM) {
		_, _ = client.SyncProgress(c.GetTestContext())
	})

	c.checkRequest(func(client TestEVM) {
		_, _ = client.NetworkID(c.GetTestContext())
	})

	/*
	  Accessor Methods
	*/

	// latest block, should not be confirmable
	c.checkRequest(func(client TestEVM) {
		_, _ = client.BalanceAt(c.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), nil)
	})

	// pending
	c.checkRequest(func(client TestEVM) {
		_, _ = client.PendingBalanceAt(c.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())))
	})

	// non-latest block
	c.checkRequest(func(client TestEVM) {
		_, _ = client.BalanceAt(c.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), new(big.Int).SetUint64(gofakeit.Uint64()))
	})

	// latest block, should not be confirmable
	c.checkRequest(func(client TestEVM) {
		_, _ = client.StorageAt(c.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())), nil)
	})

	c.checkRequest(func(client TestEVM) {
		_, _ = client.PendingStorageAt(c.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())))
	})

	// non-latest block
	c.checkRequest(func(client TestEVM) {
		_, _ = client.StorageAt(c.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64())), new(big.Int).SetUint64(gofakeit.Uint64()))
	})

	c.checkRequest(func(client TestEVM) {
		_, _ = client.CodeAt(c.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), nil)
	})

	c.checkRequest(func(client TestEVM) {
		_, _ = client.PendingCodeAt(c.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())))
	})

	// non-latest block
	c.checkRequest(func(client TestEVM) {
		_, _ = client.CodeAt(c.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), new(big.Int).SetUint64(gofakeit.Uint64()))
	})

	// storage:
	c.checkRequest(func(client TestEVM) {
		_, _ = client.NonceAt(c.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), nil)
	})

	c.checkRequest(func(client TestEVM) {
		_, _ = client.PendingNonceAt(c.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())))
	})

	// non-latest block
	c.checkRequest(func(client TestEVM) {
		_, _ = client.NonceAt(c.GetTestContext(), common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64())), new(big.Int).SetUint64(gofakeit.Uint64()))
	})

	c.checkRequest(func(client TestEVM) {
		_, _ = client.FilterLogs(c.GetTestContext(), ethereum.FilterQuery{
			FromBlock: nil,
			ToBlock:   nil,
			Addresses: []common.Address{common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64()))},
			Topics:    [][]common.Hash{{common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))}},
		})
	})

	// set only block hash
	c.checkRequest(func(client TestEVM) {
		bhash := common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))
		_, _ = client.FilterLogs(c.GetTestContext(), ethereum.FilterQuery{
			BlockHash: &bhash,
			Addresses: []common.Address{common.BigToAddress(new(big.Int).SetUint64(gofakeit.Uint64()))},
			Topics:    [][]common.Hash{{common.BigToHash(new(big.Int).SetUint64(gofakeit.Uint64()))}},
		})
	})

	/*
	   Call Contract
	*/
	c.checkRequest(func(client TestEVM) {
		_, _ = client.CallContract(c.GetTestContext(), ethereum.CallMsg{}, nil)
	})

	c.checkRequest(func(client TestEVM) {
		_, _ = client.CallContract(c.GetTestContext(), ethereum.CallMsg{}, big.NewInt(1))
	})

	c.checkRequest(func(client TestEVM) {
		_, _ = client.PendingCallContract(c.GetTestContext(), ethereum.CallMsg{})
	})

	/*
	   Gas Pricing
	*/
	c.checkRequest(func(client TestEVM) {
		_, _ = client.SuggestGasPrice(c.GetTestContext())
	})

	c.checkRequest(func(client TestEVM) {
		_, _ = client.SuggestGasTipCap(c.GetTestContext())
	})

	c.checkRequest(func(client TestEVM) {
		_, _ = client.EstimateGas(c.GetTestContext(), ethereum.CallMsg{})
	})

	// note: single batch requests will fail checkRequest since it is a helper function that expects more than one call, this could be checked manually
	// if so desired
	c.checkRequest(func(client TestEVM) {
		_ = client.BatchWithContext(c.GetTestContext(), eth.ChainID().Returns(nil), eth.ChainID().Returns(nil))
	})
}
