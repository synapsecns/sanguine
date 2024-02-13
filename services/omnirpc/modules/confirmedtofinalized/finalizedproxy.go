package confirmedtofinalized

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
	experimentalLogger "github.com/synapsecns/sanguine/core/metrics/logger"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/parser/rpc"
	ethergoRPC "github.com/synapsecns/sanguine/ethergo/parser/rpc"
	"github.com/synapsecns/sanguine/services/omnirpc/collection"
	omniHTTP "github.com/synapsecns/sanguine/services/omnirpc/http"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
)

// FinalizedProxy is the interface for the finalized proxy.
type FinalizedProxy interface {
	// Run runs the proxy.
	Run(ctx context.Context) error
}

// finalizedProxyImpl handles simple rxoy requests to omnirpc.
type finalizedProxyImpl struct {
	tracer trace.Tracer
	// port is the port the server is run on
	port uint16
	// client contains the http client
	client omniHTTP.Client
	// handler is the metrics handler
	handler metrics.Handler
	// proxyURL is the proxy url to proxy to
	proxyURL string
	// logger is the logger
	logger experimentalLogger.ExperimentalLogger
}

// NewProxy creates a new simply proxy.
func NewProxy(proxyURL string, handler metrics.Handler, port int) FinalizedProxy {
	return &finalizedProxyImpl{
		proxyURL: proxyURL,
		handler:  handler,
		port:     uint16(port),
		client:   omniHTTP.NewRestyClient(),
		tracer:   handler.Tracer(),
		logger:   handler.ExperimentalLogger(),
	}
}

func (r *finalizedProxyImpl) Run(ctx context.Context) error {
	router := ginhelper.NewWithExperimentalLogger(ctx, r.handler.ExperimentalLogger())
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

	r.logger.Infof(ctx, "running on port %d", r.port)
	err := router.Run(fmt.Sprintf("0.0.0.0:%d", r.port))
	if err != nil {
		return fmt.Errorf("could not run: %w", err)
	}
	return nil
}

var batchErr = errors.New("simple proxy batch requests are not supported")

// ProxyRequest proxies a request to the proxyURL.
func (r *finalizedProxyImpl) ProxyRequest(c *gin.Context) (err error) {
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

	rpcRequest = rewriteConfirmableRequest(rpcRequest)

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

	c.Data(resp.StatusCode(), gin.MIMEJSON, resp.Body())
	return nil
}

func rewriteConfirmableRequest(r ethergoRPC.Request) ethergoRPC.Request {
	switch client.RPCMethod(r.Method) {
	case client.BlockByNumberMethod:
		r.Params[0] = bytes.Replace(r.Params[0], latestBlock, finalizedBlock, 1)
	case client.BlockNumberMethod:
		r.Params[0] = bytes.Replace(r.Params[0], latestBlock, finalizedBlock, 1)
	}
	return r
}
