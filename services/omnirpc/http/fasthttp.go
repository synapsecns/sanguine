package http

import (
	"context"
	"fmt"
	"github.com/ImVexed/fasturl"
	"github.com/puzpuzpuz/xsync"
	http2 "github.com/synapsecns/fasthttp-http2"
	"github.com/valyala/fasthttp"
	"sync"
	"time"
)

// dialer is an allocated fasthttp dialer for increasing dns cache time.
var dialer = &fasthttp.TCPDialer{
	Concurrency:      4096,
	DNSCacheDuration: time.Hour,
}

type fastHTTPClient struct {
	// defaultClient is used when host client cannot be used
	defaultClient *fasthttp.Client

	clients *xsync.MapOf[FastClient]
	// reqPool stores Request instances that may be passed to ReleaseRequest when it is
	// no longer needed. This allows Request recycling, reduces GC pressure
	// and usually improves performance.
	reqPool sync.Pool
}

// FastClient is an interface for storing both fasthttp.Clients and fasthttp.HostClients.
type FastClient interface {
	Do(req *fasthttp.Request, resp *fasthttp.Response) error
	DoDeadline(req *fasthttp.Request, resp *fasthttp.Response, deadline time.Time) error
	DoTimeout(req *fasthttp.Request, resp *fasthttp.Response, deadline time.Duration) error
}

var _ FastClient = &fasthttp.Client{}
var _ FastClient = &fasthttp.HostClient{}

// NewFastHTTPClient creates a new fasthttp client.
// while substantially faster than resty, this can be a bad choice in certain cases:
//   - Context Cancellation not respected: fasthttp does not support context cancellation, so we hardcode a timeout here
//     this is less than ideal and puts additional load on both the application and rpc servers since we pessimistically fetch
func NewFastHTTPClient() Client {
	return &fastHTTPClient{clients: xsync.NewMapOf[FastClient](), defaultClient: &fasthttp.Client{
		NoDefaultUserAgentHeader:      true,
		Dial:                          dialer.Dial,
		DialDualStack:                 false,
		ReadTimeout:                   time.Second * 30,
		WriteTimeout:                  time.Second * 30,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
	}}
}

type rawResponse struct {
	body       []byte
	statusCode int
}

func (r *rawResponse) StatusCode() int {
	return r.statusCode
}

func newRawResponse(body []byte, statusCode int) *rawResponse {
	newBody := make([]byte, len(body))
	// copy to avoid a reallocation
	copy(newBody, body)

	return &rawResponse{body: newBody, statusCode: statusCode}
}

func (r *rawResponse) Body() []byte {
	return r.body
}

// GetClient gets the client based on string
// there are possible race conditions with clients being created twice here
// but these will be picked up by garbage collection and clients per url
// will reach 1 as server runs.
//
// this is important to note because fasthttp client level rate limiters are not enforcable
// one other note: in order to avoid url parsing every time (another struct alloc), we
// use the full endpoint url as the key here. This could result in multiple clients per host.
func (f *fastHTTPClient) GetClient(url string) FastClient {
	res, ok := f.clients.Load(url)
	if ok {
		return res
	}

	parsedURL, err := fasturl.ParseURL(url)
	// note: this should never happen because we parse urls
	// prior to this for validity, but if it does, this is used as a fallback
	// we trigger a log here
	if err != nil {
		logger.Errorf("got err: %v, not using host client", err)
		f.clients.Store(url, f.defaultClient)
		return f.defaultClient
	}

	if parsedURL.Port == "" {
		switch parsedURL.Protocol {
		case "https":
			parsedURL.Port = "443"
		case "http":
			parsedURL.Port = "80"
		}
	}

	newClient := &fasthttp.HostClient{
		Addr:                          fmt.Sprintf("%s:%s", parsedURL.Host, parsedURL.Port),
		NoDefaultUserAgentHeader:      true,
		IsTLS:                         parsedURL.Port == "443",
		Dial:                          dialer.Dial,
		DialDualStack:                 false,
		ReadTimeout:                   time.Second * 30,
		WriteTimeout:                  time.Second * 30,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
	}

	if err := http2.ConfigureClient(newClient, http2.ClientOpts{}); err != nil {
		logger.Debugf("%s does not support http2: %v", newClient.Addr, err)
	}

	f.clients.Store(url, newClient)

	return newClient
}

// AcquireRequest acquires a fasthttp request.
// this allows request cycling reducing GC overhead.
func (f *fastHTTPClient) AcquireRequest() *fastHTTPRequest {
	v := f.reqPool.Get()
	if v == nil {
		return &fastHTTPRequest{
			&fasthttp.Request{},
			f,
			nil,
		}
	}
	//nolint: forcetypeassert
	return v.(*fastHTTPRequest)
}

// ReleaseRequest releases a request object for re-use.
func (f *fastHTTPClient) ReleaseRequest(req *fastHTTPRequest) {
	req.Reset()
	req.context = nil
	f.reqPool.Put(req)
}

// fastHTTPRequest wraps fasthttp request for new methods.
type fastHTTPRequest struct {
	*fasthttp.Request
	client *fastHTTPClient
	// we need to respect context cancellation even after response
	//nolint: containedctx
	context context.Context
}

// Reset clears request contents.
func (f *fastHTTPRequest) Reset() {
	// client can stay the same
	f.Request.Reset()
}

func (f *fastHTTPRequest) SetBody(body []byte) Request {
	f.Request.SetBodyRaw(body)
	return f
}

// SetContext does nothing on fasthttp request.
func (f *fastHTTPRequest) SetContext(ctx context.Context) Request {
	f.context = ctx
	return f
}

func (f *fastHTTPRequest) SetHeader(key, value string) Request {
	f.Request.Header.Set(key, value)
	return f
}

func (f *fastHTTPRequest) SetHeaderBytes(key, value []byte) Request {
	f.Request.Header.SetBytesKV(key, value)
	return f
}

func (f *fastHTTPRequest) SetRequestURI(uri string) Request {
	f.Request.SetRequestURI(uri)
	return f
}

func (f *fastHTTPRequest) Do() (Response, error) {
	defer f.Reset()

	uri := f.Request.URI()
	if uri == nil {
		return nil, fasthttp.ErrorInvalidURI
	}

	hostClient := f.client.GetClient(f.Request.URI().String())
	f.Request.Header.SetBytesKV(Encoding, EncodingTypes)
	f.Request.Header.SetMethodBytes(PostType)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := hostClient.DoTimeout(f.Request, resp, time.Second*30)
	if err != nil {
		return nil, fmt.Errorf("could not get response: %w", err)
	}

	realResponse, err := resp.BodyUncompressed()
	if err != nil {
		return nil, fmt.Errorf("could not get response: %w", err)
	}

	select {
	case <-f.context.Done():
		return nil, fmt.Errorf("could not get context: %w", f.context.Err())
	default:
		return newRawResponse(realResponse, resp.StatusCode()), nil
	}
}

func (f *fastHTTPClient) NewRequest() Request {
	req := f.AcquireRequest()
	return req
}
