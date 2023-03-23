package proxy

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/services/omnirpc/config"
	omniHTTP "github.com/synapsecns/sanguine/services/omnirpc/http"
	"io"
	"net/http"
	"sync"
)

type SubmitProxy struct {
	// port is the port the server is run on
	port uint16
	// client contains the http client
	client omniHTTP.Client
	// submitChan is the channel for submitting requests
	submitChan chan RPCRequest
	// config is the config to use
	config config.SubmitOnlyConfig
	// lastNonce is the last nonce used
	lastNonce uint16
}

// NewSubmitProxy creates a new submit proxy.
func NewSubmitProxy(cfg config.SubmitOnlyConfig, port uint16, client omniHTTP.ClientType) *SubmitProxy {
	return &SubmitProxy{
		port:       port,
		client:     omniHTTP.NewClient(client),
		config:     cfg,
		submitChan: make(chan RPCRequest, 2000),
		lastNonce:  0,
	}
}

// Run runs the proxy.
func (s *SubmitProxy) Run(ctx context.Context) {
	router := ginhelper.New(logger)
	log.SetAllLoggers(log.LevelDebug)

	go func() {
		s.processQueue(ctx)
	}()

	router.POST("", func(c *gin.Context) {
		err := s.Forward(c)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		}
	})
}

type SubmitForwarder struct {
	// proxy is the proxy that submits txes
	proxy *SubmitProxy
	// c is the gin context for the request
	c *gin.Context
	// body is the body of the request
	body []byte
	// rpcRequests is the list of rpc requests
	rpcRequest RPCRequests
	// requestID is the request ID
	requestID []byte
}

// Forward forwards the request to the read node.
func (s *SubmitProxy) Forward(c *gin.Context) (err error) {
	forwarder := SubmitForwarder{
		c:         c,
		proxy:     s,
		requestID: []byte(c.GetHeader(omniHTTP.XRequestIDString)),
	}

	forwarder.body, err = io.ReadAll(c.Request.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}

	forwarder.rpcRequest, err = parseRPCPayload(forwarder.body)
	if err != nil {
		return fmt.Errorf("failed to parse rpc payload: %w", err)
	}

	for _, request := range forwarder.rpcRequest {
		if RPCMethod(request.Method) == SendRawTransactionMethod {
			go func() {
				s.submitChan <- request
			}()
		}
	}

	req := s.client.NewRequest()
	resp, err := req.
		SetContext(c).
		SetRequestURI(s.config.ReadURL).
		SetBody(forwarder.body).
		SetHeaderBytes(omniHTTP.XRequestID, forwarder.requestID).
		SetHeaderBytes(omniHTTP.XForwardedFor, omniHTTP.OmniRPCValue).
		SetHeaderBytes(omniHTTP.ContentType, omniHTTP.JSONType).
		SetHeaderBytes(omniHTTP.Accept, omniHTTP.JSONType).
		Do()
	if err != nil {
		return fmt.Errorf("could not get response from %s: %w", c.Request.Body, err)
	}

	c.Data(http.StatusOK, gin.MIMEJSON, resp.Body())
	return nil
}

func (s *SubmitProxy) processQueue(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case request := <-s.submitChan:
			// TODO: add some kind of nonce helper to avoid submitting nonces too far forward
			err := s.submitTx(ctx, request)
			if err != nil {
				logger.Warn(err)
			}
		}
	}
}

func (s *SubmitProxy) submitTx(ctx context.Context, request RPCRequest) error {
	var wg sync.WaitGroup

	marshalledReq, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("could not marshal request: %w", err)
	}
	for _, rpcURL := range s.config.WriteURLS {
		rpcURL := rpcURL // capture func literal
		wg.Add(1)
		go func() {
			defer wg.Done()
			req := s.client.NewRequest()
			resp, err := req.
				SetContext(ctx).
				SetRequestURI(rpcURL).
				SetBody(marshalledReq).
				SetHeaderBytes(omniHTTP.XRequestID, []byte(gofakeit.UUID())).
				SetHeaderBytes(omniHTTP.XForwardedFor, omniHTTP.OmniRPCValue).
				SetHeaderBytes(omniHTTP.ContentType, omniHTTP.JSONType).
				SetHeaderBytes(omniHTTP.Accept, omniHTTP.JSONType).
				Do()
			if err != nil {
				logger.Errorf("could not get response from %s: %w", rpcURL, err)
				return
			}
			if resp.StatusCode() != http.StatusOK {
				logger.Errorf("got non-200 status code from %s: %d", rpcURL, resp.StatusCode())
				return
			}
		}()
	}
	wg.Wait()
	return nil
}
