package rewiteconfirmable

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/flowchartsman/swaggerui"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
	experimentalLogger "github.com/synapsecns/sanguine/core/metrics/logger"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/parser/rpc"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/omnirpc/collection"
	omniHTTP "github.com/synapsecns/sanguine/services/omnirpc/http"
	"github.com/synapsecns/sanguine/services/omnirpc/proxy"
	"github.com/synapsecns/sanguine/services/omnirpc/swagger"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"io"
	"math/big"
	"net/http"
)

type RewriteConfirmableProxy interface {
	Run(ctx context.Context) error
}

type rewriteConfirmableProxy struct {
	// port is the port the server is run on
	port uint16
	// client contains the http client
	client omniHTTP.Client
	// handler is the metrics handler
	handler metrics.Handler
	// omnirpcURL is the proxy url to proxy to
	omnirpcURL string
	// omnirpcConfirmationCount is the number of confirmations to wait for
	omnirpcConfirmationCount int
	// chainID is the chain id
	chainID int
	// omniClient
	omniClient omnirpcClient.RPCClient
	// logger is the logger
	logger experimentalLogger.ExperimentalLogger
}

// NewProxy creates a new simply proxy.
func NewProxy(omnirpcURL string, handler metrics.Handler, port, omnirpcConfirmationCount, chainID int) RewriteConfirmableProxy {
	return &rewriteConfirmableProxy{
		omnirpcURL:               omnirpcURL,
		handler:                  handler,
		port:                     uint16(port),
		omnirpcConfirmationCount: omnirpcConfirmationCount,
		chainID:                  chainID,
		client:                   omniHTTP.NewRestyClient(),
		logger:                   handler.ExperimentalLogger(),
		omniClient:               omnirpcClient.NewOmnirpcClient(omnirpcURL, handler),
	}
}

func (r *rewriteConfirmableProxy) Run(ctx context.Context) error {
	router := ginhelper.NewWithExperimentalLogger(ctx, r.handler.ExperimentalLogger())
	router.Use(r.handler.Gin()...)

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

	router.Any("/swagger/*any", gin.WrapH(http.StripPrefix("/swagger", swaggerui.Handler(swagger.OpenAPI))))

	r.logger.Infof(ctx, "running on port %d", r.port)
	err := router.Run(fmt.Sprintf("0.0.0.0:%d", r.port))
	if err != nil {
		return fmt.Errorf("could not run: %w", err)
	}
	return nil
}

var errBatchNotSupported = errors.New("rewritable proxy batch requests are not supported")

func (r *rewriteConfirmableProxy) ProxyRequest(c *gin.Context) (err error) {
	ctx, span := r.handler.Tracer().Start(c, "ProxyRequest",
		trace.WithAttributes(attribute.String("endpoint", r.omnirpcURL)),
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
		err = c.Error(errBatchNotSupported)
		// nolint: wrapcheck
		return err
	}

	rpcRequests, err := rpc.ParseRPCPayload(rawBody)
	if err != nil {
		return fmt.Errorf("could not parse payload: %w", err)
	}

	rpcRequest := rpcRequests[0]

	resp, err := r.processRequest(ctx, rpcRequest, requestID)
	if err != nil {
		return err
	}

	c.Data(resp.StatusCode(), gin.MIMEJSON, resp.Body())

	return nil
}

// rewriteBodyIfNeeded rewrites the body if needed.
func (r *rewriteConfirmableProxy) rewriteBodyIfNeeded(ctx context.Context, rpcRequest rpc.Request) (body []byte, err error) {
	body, err = json.Marshal(rpcRequest)
	if err != nil {
		return nil, errors.New("could not marshal request")
	}

	var isConfirmable bool
	isConfirmable, err = proxy.IsConfirmable(rpcRequest)
	if err != nil {
		return nil, errors.New("could not marshal request")
	}

	// it's already confirmable, nothing to do here.
	// also, we only rewrite getLogsRequests.
	if isConfirmable || client.RPCMethod(rpcRequest.Method) != client.GetLogsMethod {
		return body, nil
	}

	// it's not confirmable, we need to rewrite the request.
	// at this point we already know it's a getLogs request that's using rpc.LatestBlockNumber.
	// so we're gonna need to get the latest block number.

	filterCriteria := filters.FilterCriteria{}
	err = filterCriteria.UnmarshalJSON(rpcRequest.Params[0])
	if err != nil {
		return nil, errors.New("could not marshal request")
	}

	latestClient, err := r.omniClient.GetChainClient(ctx, r.chainID)
	if err != nil {
		return nil, errors.New("could not marshal request")
	}

	// get latest block number
	var toBlock uint64
	toBlock, err = latestClient.BlockNumber(ctx)
	if err != nil {
		return nil, errors.New("could not marshal request")
	}

	filterCriteria.ToBlock = big.NewInt(int64(toBlock))

	// rewrite the request to use latest block number
	var rewrittenCriteria []byte
	rewrittenCriteria, err = json.Marshal(filterCriteria)
	if err != nil {
		return nil, errors.New("could not marshal request")
	}

	rpcRequest.Params[0] = rewrittenCriteria

	// rewrite the body
	body, err = json.Marshal(rpcRequest)
	if err != nil {
		return nil, errors.New("could not marshal request")
	}

	return body, nil
}

func (r *rewriteConfirmableProxy) processRequest(ctx context.Context, rpcRequest rpc.Request, requestID []byte) (resp omniHTTP.Response, err error) {
	req := r.client.NewRequest()
	body, err := json.Marshal(rpcRequest)
	if err != nil {
		return nil, errors.New("could not marshal request")
	}

	//nolint: exhaustive
	switch client.RPCMethod(rpcRequest.Method) {
	case client.GetLogsMethod:
		var isConfirmable bool
		isConfirmable, err = proxy.IsConfirmable(rpcRequest)
		if err != nil {
			return nil, errors.New("could not marshal request")
		}

		// it's already confirmable, nothing to do here.
		if isConfirmable {
			break
		}

		body, err = r.rewriteBodyIfNeeded(ctx, rpcRequest)
		if err != nil {
			return nil, fmt.Errorf("could not rewrite body: %w", err)
		}

		// it's not confirmable, we need to rewrite the request.
	}

	resp, err = req.
		SetContext(ctx).
		SetRequestURI(r.omniClient.GetEndpoint(r.chainID, r.omnirpcConfirmationCount)).
		SetBody(body).
		SetHeaderBytes(omniHTTP.XRequestID, requestID).
		//SetHeaderBytes(omniHTTP.XForwardedFor, []byte(r.om)). // TODO: enable this
		SetHeaderBytes(omniHTTP.ContentType, omniHTTP.JSONType).
		SetHeaderBytes(omniHTTP.Accept, omniHTTP.JSONType).
		Do()

	if err != nil {
		return nil, fmt.Errorf("could not get response from RPC %s: %w", r.omnirpcURL, err)
	}
	return resp, nil
}
