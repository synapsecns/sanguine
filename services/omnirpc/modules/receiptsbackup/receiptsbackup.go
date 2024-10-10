package receiptsbackup

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/synapsecns/sanguine/services/omnirpc/modules/mixins"
	"io"
	"net/http"
	"time"

	"github.com/flowchartsman/swaggerui"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
	experimentalLogger "github.com/synapsecns/sanguine/core/metrics/logger"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/parser/rpc"
	"github.com/synapsecns/sanguine/services/omnirpc/collection"
	omniHTTP "github.com/synapsecns/sanguine/services/omnirpc/http"
	"github.com/synapsecns/sanguine/services/omnirpc/swagger"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// ReceiptsProxy is the interface for the receipts proxy.
type ReceiptsProxy interface {
	// Run runs the proxy.
	Run(ctx context.Context) error
}

// receiptsProxyImpl handles simple proxy requests to omnirpc.
type receiptsProxyImpl struct {
	// port is the port the server is run on
	port uint16
	// client contains the http client
	client omniHTTP.Client
	// handler is the metrics handler
	handler metrics.Handler
	// proxyURL is the proxy url to proxy to
	proxyURL string
	// backupURL is the fallback for hanging receipts requests
	backupURL string
	// logger is the logger
	logger experimentalLogger.ExperimentalLogger
	// chainID is the chain id
	chainID int
	// receiptTimeout is the timeout for receipt requests before switching to backup
	receiptTimeout time.Duration
}

// NewProxy creates a new simply proxy.
func NewProxy(proxyURL, backupURL string, receiptTimeout time.Duration, handler metrics.Handler, port, chainID int) ReceiptsProxy {
	return &receiptsProxyImpl{
		proxyURL:       proxyURL,
		backupURL:      backupURL,
		handler:        handler,
		port:           uint16(port),
		client:         omniHTTP.NewRestyClient(),
		logger:         handler.ExperimentalLogger(),
		chainID:        chainID,
		receiptTimeout: receiptTimeout,
	}
}

func (r *receiptsProxyImpl) Run(ctx context.Context) error {
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

var errBatchNotSupported = errors.New("simple proxy batch requests are not supported")

// ProxyRequest proxies a request to the proxyURL.
func (r *receiptsProxyImpl) ProxyRequest(c *gin.Context) (err error) {
	ctx, span := r.handler.Tracer().Start(c, "ProxyRequest",
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

func (r *receiptsProxyImpl) processRequest(ctx context.Context, rpcRequest rpc.Request, requestID []byte) (resp omniHTTP.Response, err error) {
	ctx, span := r.handler.Tracer().Start(ctx, "proxyrequest")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	mixins.TxSubmitMixin(ctx, r.handler, rpcRequest)

	req := r.client.NewRequest()
	body, err := json.Marshal(rpcRequest)

	span.AddEvent("request marshalled", trace.WithAttributes(attribute.String("body", string(body))))

	//nolint: exhaustive
	switch client.RPCMethod(rpcRequest.Method) {
	case client.TransactionReceiptByHashMethod:
		if err != nil {
			return nil, errors.New("could not marshal request")
		}

		ctxWithTimeout, cancel := context.WithTimeout(ctx, r.receiptTimeout)
		defer cancel()

		resp, err := req.
			SetContext(ctxWithTimeout).
			SetRequestURI(r.proxyURL).
			SetBody(body).
			SetHeaderBytes(omniHTTP.XRequestID, requestID).
			SetHeaderBytes(omniHTTP.XForwardedFor, []byte(r.proxyURL)).
			SetHeaderBytes(omniHTTP.ContentType, omniHTTP.JSONType).
			SetHeaderBytes(omniHTTP.Accept, omniHTTP.JSONType).
			Do()
		if err != nil || rpc.IsNullResponse(resp.Body()) {
			// do backup request
			resp, err = req.
				SetContext(ctxWithTimeout).
				SetRequestURI(r.backupURL).
				SetBody(body).
				SetHeaderBytes(omniHTTP.XRequestID, requestID).
				SetHeaderBytes(omniHTTP.XForwardedFor, []byte(r.proxyURL)).
				SetHeaderBytes(omniHTTP.ContentType, omniHTTP.JSONType).
				SetHeaderBytes(omniHTTP.Accept, omniHTTP.JSONType).
				Do()
			if err != nil {
				return nil, fmt.Errorf("could not get response from backup RPC %s: %w", r.proxyURL, err)
			}
		}
		return resp, nil
	default:
		resp, err = req.
			SetContext(ctx).
			SetRequestURI(r.proxyURL).
			SetBody(body).
			SetHeaderBytes(omniHTTP.XRequestID, requestID).
			SetHeaderBytes(omniHTTP.XForwardedFor, []byte(r.proxyURL)).
			SetHeaderBytes(omniHTTP.ContentType, omniHTTP.JSONType).
			SetHeaderBytes(omniHTTP.Accept, omniHTTP.JSONType).
			Do()

		if err != nil {
			return nil, fmt.Errorf("could not get response from RPC %s: %w", r.proxyURL, err)
		}

		span.AddEvent("response returned", trace.WithAttributes(attribute.String("body", string(resp.Body()))))

		return resp, nil
	}
}
