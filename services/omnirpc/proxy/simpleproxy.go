package proxy

import (
	"bytes"
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
	"net/http"
	"strings"
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

func (r *SimpleProxy) Run(ctx context.Context) error {
	router := ginhelper.New(logger)
	router.Use(r.handler.Gin())

	router.POST("/", func(c *gin.Context) {
		err := r.ProxyRequest(c)
		if err != nil {
			_ = c.Error(err)
		}
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

		resp, err = r.getHarmonyReceiptVerify(ctx, txHash, rawBody)
		if err != nil {
			return true, resp, fmt.Errorf("could not get receipt: %w", err)
		}

		return true, resp, nil

	}
	return false, []byte{}, nil
}

func (r *SimpleProxy) makeReq(ctx context.Context, body []byte) ([]byte, error) {
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

	return resp.Body(), nil
}

const expectedVersion = "Harmony (C) 2023. harmony, version v8196-v2023.4.2-0-g8717ccf6 (@ 2023-12-19T10:09:52+0000)"

func (r *SimpleProxy) getHarmonyReceiptVerify(parentCtx context.Context, txHash common.Hash, rawBody []byte) (_ []byte, err error) {
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

	g.Go(func() error {
		web3Version, err := hmyClient.Web3Version(gCtx)
		if err != nil {
			return fmt.Errorf("could not get web3 version: %w", err)
		}

		if web3Version != expectedVersion {
			return fmt.Errorf("expected version %s, got %s", expectedVersion, web3Version)
		}
		return nil
	})

	g.Go(func() error {
		rawResp, err = r.makeReq(ctx, rawBody)
		if err != nil {
			return fmt.Errorf("could not make req: %w", err)
		}

		var rpcMessage JSONRPCMessage
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
		_, err := compareTX(*harmonyReceipt.Logs[i], *ethReceipt.Logs[i])
		if err != nil {
			return nil, fmt.Errorf("could not compare tx: %w", err)
		}
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

	var harmonyLogs, ethLogs []types.Log

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		harmonyLogs, err = hmyClient.FilterHarmonyLogs(gCtx, query)
		if err != nil {
			return fmt.Errorf("could not get harmony logs: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		web3Version, err := hmyClient.Web3Version(gCtx)
		if err != nil {
			return fmt.Errorf("could not get web3 version: %w", err)
		}

		if web3Version != expectedVersion {
			return fmt.Errorf("expected version %s, got %s", expectedVersion, web3Version)
		}
		return nil
	})

	g.Go(func() error {
		rawResp, err = r.makeReq(ctx, rawBody)
		if err != nil {
			return fmt.Errorf("could not make req: %w", err)
		}

		var rpcMessage JSONRPCMessage
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

	if len(harmonyLogs) != len(ethLogs) {
		return nil, fmt.Errorf("expected %d logs, got %d", len(harmonyLogs), len(ethLogs))
	}

	for i := 0; i < len(harmonyLogs); i++ {
		_, err := compareTX(harmonyLogs[i], ethLogs[i])
		if err != nil {
			return nil, fmt.Errorf("could not compare tx: %w", err)
		}

	}

	return rawResp, nil
}

// compareTX makes sure all logs are equal except for the tx hash
// which is epected to be different owing to Ethereum Hash and Hash which both exist on harmony
func compareTX(harmonyLog, ethLog types.Log) (bool, error) {
	if harmonyLog.BlockHash != ethLog.BlockHash {
		return false, fmt.Errorf("expected block hash %s, got %s", harmonyLog.BlockHash, ethLog.BlockHash)
	}

	if harmonyLog.TxHash == ethLog.TxHash {
		return false, fmt.Errorf("expected different tx hashes %s, got %s", harmonyLog.TxHash, ethLog.TxHash)
	}

	if harmonyLog.TxIndex != ethLog.TxIndex {
		return false, fmt.Errorf("expected tx index %d, got %d", harmonyLog.TxIndex, ethLog.TxIndex)
	}

	if harmonyLog.Index != ethLog.Index {
		return false, fmt.Errorf("expected index %d, got %d", harmonyLog.Index, ethLog.Index)
	}

	if len(harmonyLog.Topics) != len(ethLog.Topics) {
		return false, fmt.Errorf("expected %d topics, got %d", len(harmonyLog.Topics), len(ethLog.Topics))
	}

	for j := 0; j < len(harmonyLog.Topics); j++ {
		harmonyTopic := harmonyLog.Topics[j]
		ethTopic := ethLog.Topics[j]

		if harmonyTopic != ethTopic {
			return false, fmt.Errorf("expected topic %s, got %s", harmonyTopic, ethTopic)
		}
	}

	if !bytes.Equal(harmonyLog.Data, ethLog.Data) {
		return false, fmt.Errorf("expected data %s, got %s", harmonyLog.Data, ethLog.Data)
	}
	return true, nil
}
