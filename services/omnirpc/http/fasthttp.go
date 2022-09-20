package http

import (
	"context"
	"fmt"
	"github.com/ImVexed/fasturl"
	"github.com/dgrr/http2"
	"github.com/puzpuzpuz/xsync"
	"github.com/valyala/fasthttp"
	"sync"
	"time"
)

// dialer is an allocated fasthttp dialer for increasing dns cache time.
var dialer = &fasthttp.TCPDialer{
	Concurrency:      4096,
	DNSCacheDuration: time.Hour,
}

type FastHTTPClient struct {
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

func NewFastHTTPClient() Client {
	return &FastHTTPClient{clients: xsync.NewMapOf[FastClient](), defaultClient: &fasthttp.Client{
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
	body []byte
}

func newRawResponse(body []byte) *rawResponse {
	return &rawResponse{body: body}
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
func (f *FastHTTPClient) GetClient(url string) FastClient {
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
func (f *FastHTTPClient) AcquireRequest() *fastHTTPRequest {
	v := f.reqPool.Get()
	if v == nil {
		return &fastHTTPRequest{
			&fasthttp.Request{},
			f,
		}
	}
	return v.(*fastHTTPRequest)
}

// ReleaseRequest releases a request object for re-use.
func (f *FastHTTPClient) ReleaseRequest(req *fastHTTPRequest) {
	req.Reset()
	f.reqPool.Put(req)
}

// fastHTTPRequest wraps fasthttp request for new methods.
type fastHTTPRequest struct {
	*fasthttp.Request
	client *FastHTTPClient
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
func (f *fastHTTPRequest) SetContext(_ context.Context) Request {
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

	err := hostClient.Do(f.Request, resp)
	if err != nil {
		return nil, fmt.Errorf("could not get response: %w", err)
	}

	realResponse, err := resp.BodyUncompressed()
	if err != nil {
		return nil, fmt.Errorf("could not get response: %w", err)
	}

	return newRawResponse(realResponse), nil
}

func (f *FastHTTPClient) NewRequest() Request {
	req := f.AcquireRequest()
	return req
}
