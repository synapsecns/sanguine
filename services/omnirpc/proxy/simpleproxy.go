package proxy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/parser/rpc"
	"github.com/synapsecns/sanguine/services/omnirpc/collection"
	omniHTTP "github.com/synapsecns/sanguine/services/omnirpc/http"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"io"
	"k8s.io/apimachinery/pkg/util/sets"
	"math/big"
	"net/http"
	"strings"
	"sync"
)

// SimpleProxy handles simple rxoy requests to omnirpc
type SimpleProxy struct {
	tracer trace.Tracer
	// port is the port the server is run on
	port uint16
	// client contains the http client
	client omniHTTP.Client
	// handler is the metrics handler
	handler metrics.Handler
	// proxyURL is the proxy url to proxy to
	proxyURL string
}

// NewSimpleProxy creates a new simply proxy
func NewSimpleProxy(proxyURL string, handler metrics.Handler, port int) *SimpleProxy {
	return &SimpleProxy{
		proxyURL: proxyURL,
		handler:  handler,
		port:     uint16(port),
		client:   omniHTTP.NewRestyClient(),
		tracer:   handler.Tracer(),
	}
}

const maxAttempts = 15

func (r *SimpleProxy) Run(ctx context.Context) error {
	router := ginhelper.New(logger)
	router.Use(r.handler.Gin())

	router.POST("/", func(c *gin.Context) {
		var err error
		//for i := 0; i < 15; i++ {
		err = r.ProxyRequest(c)
		if err != nil {
			_ = c.Error(err)
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		}
		//}
	})

	router.GET("/collection.json", func(c *gin.Context) {
		res, err := collection.CreateCollection()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("could not parse collection: %v", err),
			})
		}
		c.Data(http.StatusOK, gin.MIMEJSON, res)
	})

	logger.Infof("running on port %d", r.port)
	err := router.Run(fmt.Sprintf("0.0.0.0:%d", r.port))
	if err != nil {
		return fmt.Errorf("could not run: %w", err)
	}
	return nil
}

var batchErr = errors.New("simple proxy batch requests are not supported")

// ProxyRequest proxies a request to the proxyURL
func (r *SimpleProxy) ProxyRequest(c *gin.Context) (err error) {
	ctx, span := r.tracer.Start(c, "ProxyRequest",
		trace.WithAttributes(attribute.String("endpoint", r.proxyURL)),
	)

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	requestID := []byte(c.GetHeader(omniHTTP.XRequestIDString))

	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return fmt.Errorf("could not read request body: %w", err)
	}

	// make sure it's not a batch
	if rpc.IsBatch(rawBody) {
		err = c.Error(batchErr)
		return err
	}

	rpcRequests, err := rpc.ParseRPCPayload(rawBody)
	if err != nil {
		return fmt.Errorf("could not parse payload: %w", err)
	}

	rpcRequest := rpcRequests[0]

	span.SetAttributes(attribute.String("original-body", string(rawBody)))
	customHandler, rawResp, err := r.verifyHarmonyRequest(ctx, rpcRequest, rawBody)
	if err != nil {
		return fmt.Errorf("could not verify harmony request: %w", err)
	}
	if customHandler {
		c.Data(http.StatusOK, gin.MIMEJSON, rawResp)
		return nil
	}

	body, err := json.Marshal(rpcRequest)
	if err != nil {
		return fmt.Errorf("could not marshal request")
	}

	req := r.client.NewRequest()
	resp, err := req.
		SetContext(ctx).
		SetRequestURI(r.proxyURL).
		SetBody(body).
		SetHeaderBytes(omniHTTP.XRequestID, requestID).
		SetHeaderBytes(omniHTTP.XForwardedFor, []byte(r.proxyURL)).
		SetHeaderBytes(omniHTTP.ContentType, omniHTTP.JSONType).
		SetHeaderBytes(omniHTTP.Accept, omniHTTP.JSONType).
		Do()
	if err != nil {
		return fmt.Errorf("could not get response from %s: %w", r.proxyURL, err)
	}

	// TODO: caste to rpc response

	c.Data(resp.StatusCode(), gin.MIMEJSON, resp.Body())
	return nil
}

func (r *SimpleProxy) verifyHarmonyRequest(ctx context.Context, req rpc.Request, rawBody []byte) (willHandle bool, resp []byte, err error) {
	switch client.RPCMethod(req.Method) {
	case client.GetLogsMethod:
		if len(req.Params) != 1 {
			return true, resp, fmt.Errorf("expected 1 param, got %d", len(req.Params))
		}

		params := req.Params[0]
		var fq filters.FilterCriteria
		rawJson, err := params.MarshalJSON()
		if err != nil {
			return true, resp, fmt.Errorf("could not marshal params: %w", err)
		}
		err = json.Unmarshal(rawJson, &fq)
		if err != nil {
			return true, resp, fmt.Errorf("could not unmarshal params: %w", err)
		}

		// according to godoc, this is the same as ethereum.FitlerQuery w/ an unmarshal method, so well convert ehre
		query := ethereum.FilterQuery{
			BlockHash: fq.BlockHash,
			FromBlock: fq.FromBlock,
			ToBlock:   fq.ToBlock,
			Addresses: fq.Addresses,
			Topics:    fq.Topics,
		}

		resp, err = r.getLogsHarmonyVerify(ctx, query, rawBody)
		if err != nil {
			return true, resp, fmt.Errorf("could not get logs: %w", err)
		}

		return true, resp, nil
	case client.TransactionReceiptByHashMethod:
		if len(req.Params) != 1 {
			return true, resp, fmt.Errorf("expected 1 param, got %d", len(req.Params))
		}

		params := req.Params[0]
		txHash := common.HexToHash(strings.Trim(string(params), "\""))

		resp, err = r.getHarmonyReceiptVerify(ctx, txHash, rawBody, true)
		if err != nil {
			return true, resp, fmt.Errorf("could not get receipt: %w", err)
		}

		return true, resp, nil

	}
	return false, []byte{}, nil
}

func (r *SimpleProxy) makeReq(parentCtx context.Context, body []byte) (_ []byte, err error) {
	ctx, span := r.tracer.Start(parentCtx, "makeReq")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()
	span.AddEvent("http.request", trace.WithAttributes(attribute.String("body", string(body))))

	req := r.client.NewRequest()
	resp, err := req.
		SetContext(ctx).
		SetRequestURI(r.proxyURL).
		SetBody(body).
		SetHeaderBytes(omniHTTP.XForwardedFor, []byte(r.proxyURL)).
		SetHeaderBytes(omniHTTP.ContentType, omniHTTP.JSONType).
		SetHeaderBytes(omniHTTP.Accept, omniHTTP.JSONType).
		Do()
	if err != nil {
		return nil, fmt.Errorf("could not get response from %s: %w", r.proxyURL, err)
	}

	respBody := resp.Body()
	span.AddEvent("http.response", trace.WithAttributes(attribute.String("body", string(respBody))))

	return respBody, nil
}

const expectedVersion = "Harmony (C) 2023. harmony, version v8197-v2023.4.2-1-g40a2374d"

func (r *SimpleProxy) getHarmonyReceiptVerify(parentCtx context.Context, txHash common.Hash, rawBody []byte, checkVersion bool) (_ []byte, err error) {
	ctx, span := r.tracer.Start(parentCtx, "getHarmonyReceiptVerify")

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	hmyClient, err := client.DialHarmonyBackend(ctx, r.proxyURL, r.handler, client.Capture(true))
	if err != nil {
		return nil, fmt.Errorf("could not dial harmony backend: %w", err)
	}

	var harmonyReceipt, ethReceipt *types.Receipt
	var rawResp []byte
	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		harmonyReceipt, err = hmyClient.HarmonyTransactionReceipt(gCtx, txHash)
		if err != nil {
			return fmt.Errorf("could not get harmony receipt: %w", err)
		}
		return nil
	})

	var rpcMessage JSONRPCMessage

	g.Go(func() error {
		/// no need to double up on this check when doing receipts
		if checkVersion {
			web3Version, err := hmyClient.Web3Version(gCtx)
			if err != nil {
				return fmt.Errorf("could not get web3 version: %w", err)
			}

			if !strings.Contains(web3Version, expectedVersion) {
				return fmt.Errorf("expected version %s, got %s", expectedVersion, web3Version)
			}
		}
		return nil
	})

	g.Go(func() error {
		rawResp, err = r.makeReq(ctx, rawBody)
		if err != nil {
			return fmt.Errorf("could not make req: %w", err)
		}

		err = json.Unmarshal(rawResp, &rpcMessage)
		if err != nil {
			return fmt.Errorf("could not unmarshal: %w", err)
		}

		err = json.Unmarshal(rpcMessage.Result, &ethReceipt)
		if err != nil {
			return fmt.Errorf("could not unmarshal eth receipt: %w", err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could not get receipts: %w", err)
	}

	if harmonyReceipt.BlockHash != ethReceipt.BlockHash {
		return nil, fmt.Errorf("expected block hash %s, got %s", harmonyReceipt.BlockHash, ethReceipt.BlockHash)
	}

	if harmonyReceipt.TxHash == ethReceipt.TxHash {
		return nil, fmt.Errorf("expected different tx hashes %s, got %s", harmonyReceipt.TxHash, ethReceipt.TxHash)
	}

	if harmonyReceipt.Status != ethReceipt.Status {
		return nil, fmt.Errorf("expected tx index %d, got %d", harmonyReceipt.Status, ethReceipt.Status)
	}

	if harmonyReceipt.CumulativeGasUsed != ethReceipt.CumulativeGasUsed {
		return nil, fmt.Errorf("expected index %d, got %d", harmonyReceipt.CumulativeGasUsed, ethReceipt.CumulativeGasUsed)
	}

	if harmonyReceipt.GasUsed != ethReceipt.GasUsed {
		return nil, fmt.Errorf("expected index %d, got %d", harmonyReceipt.GasUsed, ethReceipt.GasUsed)
	}

	if len(harmonyReceipt.Logs) != len(ethReceipt.Logs) {
		return nil, fmt.Errorf("expected %d logs, got %d", len(harmonyReceipt.Logs), len(ethReceipt.Logs))
	}

	for i := 0; i < len(harmonyReceipt.Logs); i++ {
		ethReceipt.Logs[i].TxHash = ethReceipt.TxHash
	}

	receiptLogsMarshall, err := json.Marshal(ethReceipt.Logs)
	if err != nil {
		return nil, fmt.Errorf("could not marshal eth receipt: %w", err)
	}

	var fields map[string]json.RawMessage
	err = json.Unmarshal(rpcMessage.Result, &fields)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal fields: %w", err)
	}

	fields["logs"] = json.RawMessage(receiptLogsMarshall)
	rpcMessage.Result, err = json.Marshal(fields)
	if err != nil {
		return nil, fmt.Errorf("could not marshal fields: %w", err)
	}

	rawResp, err = json.Marshal(rpcMessage)
	if err != nil {
		return nil, fmt.Errorf("could not marshal rpc message: %w", err)
	}

	return rawResp, nil
}

func (r *SimpleProxy) getLogsHarmonyVerify(parentCtx context.Context, query ethereum.FilterQuery, rawBody []byte) (rawResp []byte, err error) {
	ctx, span := r.tracer.Start(parentCtx, "getLogsHarmonyVerify")

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	hmyClient, err := client.DialHarmonyBackend(ctx, r.proxyURL, r.handler, client.Capture(true))
	if err != nil {
		return nil, fmt.Errorf("could not dial harmony backend: %w", err)
	}

	var ethLogs []types.Log
	var rpcMessage JSONRPCMessage

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		web3Version, err := hmyClient.Web3Version(gCtx)
		if err != nil {
			return fmt.Errorf("could not get web3 version: %w", err)
		}

		if !strings.Contains(web3Version, expectedVersion) {
			return fmt.Errorf("expected version %s, got %s", expectedVersion, web3Version)
		}
		return nil
	})

	g.Go(func() error {
		rawResp, err = r.makeReq(ctx, rawBody)
		if err != nil {
			return fmt.Errorf("could not make req: %w", err)
		}

		err = json.Unmarshal(rawResp, &rpcMessage)
		if err != nil {
			return fmt.Errorf("could not unmarshal: %w", err)
		}

		err = json.Unmarshal(rpcMessage.Result, &ethLogs)
		if err != nil {
			return fmt.Errorf("could not unmarshal eth receipt: %w", err)
		}

		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could not get logs: %w", err)
	}

	uniqueHashes := sets.NewString()
	for i := 0; i < len(ethLogs); i++ {
		uniqueHashes.Insert(ethLogs[i].TxHash.String())
	}

	g, gCtx = errgroup.WithContext(ctx)
	var logs []*types.Log
	var mux sync.Mutex
	for _, hash := range uniqueHashes.List() {
		g.Go(func() error {
			rawReqBody, err := json.Marshal(rpc.Request{
				ID:     1,
				Method: client.TransactionReceiptByHashMethod.String(),
				Params: []json.RawMessage{json.RawMessage(fmt.Sprintf("\"%s\"", hash))},
			})

			resp, err := r.getHarmonyReceiptVerify(gCtx, common.HexToHash(hash), rawReqBody, false)
			if err != nil {
				return fmt.Errorf("could not get harmony receipt: %w", err)
			}

			var rpcMessage JSONRPCMessage
			err = json.Unmarshal(resp, &rpcMessage)
			if err != nil {
				return fmt.Errorf("could not unmarshal: %w", err)
			}

			var receipt types.Receipt
			err = json.Unmarshal(rpcMessage.Result, &receipt)
			if err != nil {
				return fmt.Errorf("could not unmarshal: %w", err)
			}

			mux.Lock()
			logs = append(logs, receipt.Logs...)
			mux.Unlock()
			return nil
		})
	}

	err = g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could not get logs: %w", err)
	}

	filteredLogs := filterLogs(logs, query.FromBlock, query.ToBlock, query.Addresses, query.Topics)
	if err != nil {
		return nil, fmt.Errorf("could not filter logs: %w", err)
	}

	rpcMessage.Result, err = json.Marshal(filteredLogs)
	if err != nil {
		return nil, fmt.Errorf("could not marshal fields: %w", err)
	}

	return rawResp, nil
}

// filterLogs creates a slice of logs matching the given criteria.
func filterLogs(logs []*types.Log, fromBlock, toBlock *big.Int, addresses []common.Address, topics [][]common.Hash) []*types.Log {
	var ret []*types.Log
Logs:
	for _, log := range logs {
		if fromBlock != nil && fromBlock.Int64() >= 0 && fromBlock.Uint64() > log.BlockNumber {
			continue
		}
		if toBlock != nil && toBlock.Int64() >= 0 && toBlock.Uint64() < log.BlockNumber {
			continue
		}

		if len(addresses) > 0 && !includes(addresses, log.Address) {
			continue
		}
		// If the to filtered topics is greater than the amount of topics in logs, skip.
		if len(topics) > len(log.Topics) {
			continue
		}
		for i, sub := range topics {
			match := len(sub) == 0 // empty rule set == wildcard
			for _, topic := range sub {
				if log.Topics[i] == topic {
					match = true
					break
				}
			}
			if !match {
				continue Logs
			}
		}
		ret = append(ret, log)
	}
	return ret
}

func includes(addresses []common.Address, a common.Address) bool {
	for _, addr := range addresses {
		if addr == a {
			return true
		}
	}

	return false
}
