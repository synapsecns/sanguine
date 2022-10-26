package proxy

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/puzpuzpuz/xsync"
	"github.com/synapsecns/sanguine/services/omnirpc/chainmanager"
	omniHTTP "github.com/synapsecns/sanguine/services/omnirpc/http"
)

// IsConfirmable exports isConfirmable for testing.
func IsConfirmable(body []byte) (bool, error) {
	parsedPayload, err := parseRPCPayload(body)
	if err != nil {
		return false, fmt.Errorf("could not parse payload: %w", err)
	}
	//nolint: wrapcheck
	return parsedPayload.isConfirmable()
}

// ParseRPCPayload exports parseRPCPayload for testing.
func ParseRPCPayload(body []byte) (_ *RPCRequest, err error) {
	//nolint: wrapcheck
	return parseRPCPayload(body)
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

func (f *Forwarder) RPCRequest() *RPCRequest {
	return f.rpcRequest
}

func (f *Forwarder) SetRPCRequest(rpcRequest *RPCRequest) {
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

func StandardizeResponse(method string, body []byte) ([]byte, error) {
	var rpcMessage JSONRPCMessage
	err := json.Unmarshal(body, &rpcMessage)
	if err != nil {
		//nolint: wrapcheck
		return nil, err
	}
	params := []json.RawMessage{rpcMessage.Params}

	// Handle BlockByHash, BlockByNumber, and HeaderByNumber events.
	if method == "eth_getBlockByHash" || method == "eth_getBlockByNumber" {
		blockNumber := "0x1"
		// TODO change for block by header
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
	rpcRequest := RPCRequest{
		ID:     rpcMessage.ID,
		Method: method,
		Params: params,
	}
	return standardizeResponse(context.Background(), rpcRequest, rpcMessage)
}

// CheckAndSetConfirmability exports checkAndSetConfirmability for testing.
func (f *Forwarder) CheckAndSetConfirmability() (ok bool) {
	return f.checkAndSetConfirmability()
}
