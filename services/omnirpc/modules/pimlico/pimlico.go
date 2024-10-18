package pimlico

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
	experimentalLogger "github.com/synapsecns/sanguine/core/metrics/logger"
	"github.com/synapsecns/sanguine/ethergo/parser/rpc"
	omniHTTP "github.com/synapsecns/sanguine/services/omnirpc/http"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// PimlicoProxy is the interface for the Pimlico proxy.
type PimlicoProxy interface {
	Run(ctx context.Context) error
}

type pimlicoProxyImpl struct {
	port       uint16
	client     omniHTTP.Client
	handler    metrics.Handler
	pimlicoAPI string
	ethAPI     string
	logger     experimentalLogger.ExperimentalLogger
}

func NewProxy(pimlicoAPI, ethAPI string, handler metrics.Handler, port int) PimlicoProxy {

	return &pimlicoProxyImpl{
		pimlicoAPI: pimlicoAPI,
		ethAPI:     ethAPI,
		handler:    handler,
		port:       uint16(port),
		client:     omniHTTP.NewRestyClient(),
		logger:     handler.ExperimentalLogger(),
	}
}

func (p *pimlicoProxyImpl) Run(ctx context.Context) error {
	router := ginhelper.NewWithExperimentalLogger(ctx, p.handler.ExperimentalLogger())
	router.Use(p.handler.Gin()...)

	router.POST("/", func(c *gin.Context) {
		err := p.ProxyRequest(c)
		if err != nil {
			_ = c.Error(err)
		}
	})

	p.logger.Infof(ctx, "running on port %d", p.port)
	err := router.Run(fmt.Sprintf("0.0.0.0:%d", p.port))
	if err != nil {
		return fmt.Errorf("could not run: %w", err)
	}
	return nil
}

func (p *pimlicoProxyImpl) ProxyRequest(c *gin.Context) (err error) {
	ctx, span := p.handler.Tracer().Start(c, "ProxyRequest")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	requestID := []byte(c.GetHeader(omniHTTP.XRequestIDString))

	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return fmt.Errorf("could not read request body: %w", err)
	}

	if rpc.IsBatch(rawBody) {
		return fmt.Errorf("batch requests are not supported")
	}

	rpcRequests, err := rpc.ParseRPCPayload(rawBody)
	if err != nil {
		return fmt.Errorf("could not parse payload: %w", err)
	}

	rpcRequest := rpcRequests[0]
	resp, err := p.processRequest(ctx, c, rpcRequest, requestID)
	if err != nil {
		return err
	}

	if resp.StatusCode() > 300 {
		fmt.Println("f")
	}

	c.Data(resp.StatusCode(), gin.MIMEJSON, resp.Body())
	return nil
}

func (p *pimlicoProxyImpl) processRequest(ctx context.Context, c *gin.Context, rpcRequest rpc.Request, requestID []byte) (resp omniHTTP.Response, err error) {
	ctx, span := p.handler.Tracer().Start(ctx, "processRequest")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	pimlicoMethods := map[string]bool{
		"eth_chainId":                          true,
		"eth_supportedEntryPoints":             true,
		"eth_coinbase":                         true,
		"eth_estimateUserOperationGas":         true,
		"eth_sendUserOperation":                true,
		"eth_getUserOperationByHash":           true,
		"eth_getUserOperationReceipt":          true,
		"pm_supportedEntryPoints":              true,
		"pm_sponsorUserOperation":              true,
		"pm_getPaymasterStubData":              true,
		"pm_getPaymasterData":                  true,
		"pm_validateSponsorshipPolicies":       true,
		"web3_clientVersion":                   true,
		"pimlico_getTokenQuotesAndEstimateGas": true,
		"pimlico_getTokenQuotes":               true,
		"pimlico_sendCompressedUserOperation":  true,
		"pimlico_getUserOperationStatus":       true,
		"pimlico_getUserOperationGasPrice":     true,
		"pimlico_getBalance":                   true,
	}

	targetURL := p.ethAPI
	if pimlicoMethods[rpcRequest.Method] {
		targetURL = p.pimlicoAPI
	}

	span.SetAttributes(attribute.String("target_url", targetURL))

	req := p.client.NewRequest()
	body, err := json.Marshal(rpcRequest)
	if err != nil {
		return nil, fmt.Errorf("could not marshal request: %w", err)
	}

	for key, value := range c.Request.Header {
		req.SetHeader(key, value[0])
	}

	resp, err = req.
		SetContext(ctx).
		SetRequestURI(targetURL).
		SetBody(body).
		SetHeader("User-Agent", "omnirpc").
		SetHeaderBytes(omniHTTP.XRequestID, requestID).
		SetHeaderBytes(omniHTTP.XForwardedFor, []byte(targetURL)).
		SetHeaderBytes(omniHTTP.ContentType, omniHTTP.JSONType).
		SetHeaderBytes(omniHTTP.Accept, omniHTTP.JSONType).
		Do()

	if err != nil {
		return nil, fmt.Errorf("could not get response from API %s: %w", targetURL, err)
	}

	span.AddEvent("response returned", trace.WithAttributes(attribute.String("body", string(resp.Body()))))

	return resp, nil
}
