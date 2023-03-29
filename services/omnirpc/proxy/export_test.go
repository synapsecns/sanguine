package proxy

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/puzpuzpuz/xsync"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/parser/rpc"
	"github.com/synapsecns/sanguine/services/omnirpc/chainmanager"
	omniHTTP "github.com/synapsecns/sanguine/services/omnirpc/http"
	"go.opentelemetry.io/otel/trace"
)

// IsConfirmable exports isConfirmable for testing.
func IsConfirmable(body []byte) (bool, error) {
	parsedPayload, err := rpc.ParseRPCPayload(body)
	if err != nil {
		return false, fmt.Errorf("could not parse payload: %w", err)
	}

	//nolint: wrapcheck
	return areConfirmable(parsedPayload)
}

// SetClient allows overriding the client on the rpc proxy.
func (r *RPCProxy) SetClient(client omniHTTP.Client) {
	r.client = client
}

// RawResponse exports rawResponse for testing.
type RawResponse interface {
	Body() []byte
	URL() string
	Hash() string
}

// ForwardRequest exports forward request for testing.
func (f *Forwarder) ForwardRequest(ctx context.Context, endpoint string) (RawResponse, error) {
	//nolint: wrapcheck
	return f.forwardRequest(ctx, endpoint)
}

func (r rawResponse) Body() []byte {
	return r.body
}

func (r rawResponse) URL() string {
	return r.url
}

func (f *Forwarder) R() *RPCProxy {
	return f.r
}

func (f *Forwarder) SetR(r *RPCProxy) {
	f.r = r
}

func (f *Forwarder) C() *gin.Context {
	return f.c
}

func (f *Forwarder) SetC(c *gin.Context) {
	f.c = c
}

func (f *Forwarder) Chain() chainmanager.Chain {
	return f.chain
}

func (f *Forwarder) SetChain(chain chainmanager.Chain) {
	f.chain = chain
}

func (f *Forwarder) Body() []byte {
	return f.body
}

func (f *Forwarder) SetBody(body []byte) {
	f.body = body
}

func (f *Forwarder) SetSpan(span trace.Span) {
	f.span = span
}

func (f *Forwarder) RequiredConfirmations() uint16 {
	return f.requiredConfirmations
}

func (f *Forwarder) SetRequiredConfirmations(requiredConfirmations uint16) {
	f.requiredConfirmations = requiredConfirmations
}

func (f *Forwarder) RequestID() []byte {
	return f.requestID
}

func (f *Forwarder) SetRequestID(requestID []byte) {
	f.requestID = requestID
}

func (f *Forwarder) Client() omniHTTP.Client {
	return f.client
}

func (f *Forwarder) SetClient(client omniHTTP.Client) {
	f.client = client
}

func (f *Forwarder) ResMap() *xsync.MapOf[[]rawResponse] {
	return f.resMap
}

func (f *Forwarder) SetResMap(resMap *xsync.MapOf[[]rawResponse]) {
	f.resMap = resMap
}

func (f *Forwarder) RPCRequest() []rpc.Request {
	return f.rpcRequest
}

func (f *Forwarder) SetRPCRequest(rpcRequest []rpc.Request) {
	f.rpcRequest = rpcRequest
}

func (r rawResponse) Hash() string {
	return r.hash
}

var _ RawResponse = rawResponse{}

// SetBlankResMap sets a forwarder to a new res map for testing.
func (f *Forwarder) SetBlankResMap() {
	f.SetResMap(xsync.NewMapOf[[]rawResponse]())
}

func StandardizeResponse(ctx context.Context, req []rpc.Request, body []byte) ([]byte, error) {
	var rpcMessage JSONRPCMessage
	err := json.Unmarshal(body, &rpcMessage)
	if err != nil {
		//nolint: wrapcheck
		return nil, err
	}
	var standardizedResponses []byte
	for i := range req {
		standardizedResponse, err := standardizeResponse(ctx, &req[i], rpcMessage)
		if err != nil {
			return nil, fmt.Errorf("could not standardize response: %w", err)
		}
		standardizedResponses = append(standardizedResponses, standardizedResponse...)
	}
	return standardizedResponses, nil
}

// StandardizeResponseFalseParams exports standardizeResponseFalseParams for testing.
// this is only used when params[1] is false when calling eth_getBlockByNumber or eth_getBlockByHash.
func StandardizeResponseFalseParams(ctx context.Context, req []rpc.Request, body []byte) ([]byte, error) {
	var rpcMessage JSONRPCMessage
	err := json.Unmarshal(body, &rpcMessage)
	if err != nil {
		//nolint: wrapcheck
		return nil, err
	}
	params := []json.RawMessage{rpcMessage.Params}

	// Handle BlockByHash, BlockByNumber, and HeaderByNumber events.
	if req[0].Method == string(client.BlockByHashMethod) || req[0].Method == string(client.BlockByNumberMethod) {
		blockNumber := "0x1"
		flag := true
		jsonBlockNumber, err := json.Marshal(&blockNumber)
		if err != nil {
			//nolint: wrapcheck
			return nil, err
		}
		jsonFlag, err := json.Marshal(&flag)
		if err != nil {
			//nolint: wrapcheck
			return nil, err
		}
		jsonRawParams := []json.RawMessage{jsonBlockNumber, jsonFlag}
		params = jsonRawParams
	}
	var standardizedResponses []byte
	for i := range req {
		rpcRequest := rpc.Request{
			ID:     rpcMessage.ID,
			Method: req[i].Method,
			Params: params,
		}
		standardizedResponse, err := standardizeResponse(ctx, &rpcRequest, rpcMessage)
		if err != nil {
			return nil, fmt.Errorf("could not standardize response: %w", err)
		}
		standardizedResponses = append(standardizedResponses, standardizedResponse...)
	}

	return standardizedResponses, nil
}

// CheckAndSetConfirmability exports checkAndSetConfirmability for testing.
func (f *Forwarder) CheckAndSetConfirmability() (ok bool) {
	return f.checkAndSetConfirmability()
}
