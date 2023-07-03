package proxy

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/ethergo/parser/rpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/Soft/iter"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/puzpuzpuz/xsync"
	"github.com/synapsecns/sanguine/core/threaditer"
	"github.com/synapsecns/sanguine/services/omnirpc/chainmanager"
	omniHTTP "github.com/synapsecns/sanguine/services/omnirpc/http"
	"k8s.io/apimachinery/pkg/util/sets"
)

// Forwarder creates a request forwarder.
type Forwarder struct {
	// r is the parent rpc proxy object
	r *RPCProxy
	// c is the gin context for the request
	c *gin.Context
	// chain is the chain from the chain manager
	chain chainmanager.Chain
	// body is the body of the request
	body []byte
	// requiredConfirmations is the number of required confirmations for the request to go through
	requiredConfirmations uint16
	// requestID is the request id
	requestID []byte
	// client is the client used for fasthttp
	client omniHTTP.Client
	// resMap is the res map
	// Note: because we use an array here, this is not thread safe for writes
	resMap *xsync.MapOf[[]rawResponse]
	// failedForwards is a map of failed forwards
	failedForwards *xsync.MapOf[error]
	// rpcRequest is the parsed rpc request
	rpcRequest rpc.Requests
	// mux is used to track the release of the forwarder. This should only be used in async methods
	// as RLock
	mux sync.RWMutex
	// span is the span for the request
	span trace.Span
	// tracer is the tracer for the request
	tracer trace.Tracer
}

// Reset resets the forwarder so it can be reused.
func (f *Forwarder) Reset() {
	// try to acquire the lock. this is
	f.mux.Lock()
	defer f.mux.Unlock()
	// client and forwarder can stay the same
	f.c = nil
	f.chain = nil
	f.body = nil
	f.requiredConfirmations = 0
	f.requestID = nil
	f.resMap = nil
	f.failedForwards = nil
	f.rpcRequest = nil
	f.span = nil
}

// AcquireForwarder allocates a forwarder and allows it to be released when not in use
// this allows forwarder cycling reducing GC overhead.
func (r *RPCProxy) AcquireForwarder() *Forwarder {
	v := r.forwarderPool.Get()
	if v == nil {
		return &Forwarder{
			r:      r,
			client: r.client,
			tracer: r.tracer,
		}
	}
	//nolint: forcetypeassert
	return v.(*Forwarder)
}

// ReleaseForwarder releases a forwarder object for reuse.
func (r *RPCProxy) ReleaseForwarder(f *Forwarder) {
	f.Reset()
	r.forwarderPool.Put(f)
}

// Forward forwards the rpc request to the servers and makes assertions around confirmation thresholds.
// required confirmations can be used to override the required confirmations count.
func (r *RPCProxy) Forward(c *gin.Context, chainID uint32, requiredConfirmationsOverride *uint16) {
	ctx, span := r.tracer.Start(c, "rpcRequest",
		trace.WithAttributes(attribute.Int("chainID", int(chainID))),
	)

	forwarder := r.AcquireForwarder()
	defer func() {
		span.End()
		r.ReleaseForwarder(forwarder)
	}()

	forwarder.c = c
	forwarder.span = span
	forwarder.resMap = xsync.NewMapOf[[]rawResponse]()
	forwarder.failedForwards = xsync.NewMapOf[error]()
	if requiredConfirmationsOverride != nil {
		forwarder.requiredConfirmations = *requiredConfirmationsOverride
	}

	if ok := forwarder.fillAndValidate(chainID); !ok {
		return
	}

	forwarder.attemptForwardAndValidate(ctx)
}

// attemptForwardAndValidate attempts to forward the request and
// makes sure it is valid
// TODO: maybe the context shouldn't be used from a struct here?
//
//nolint:gocognit,cyclop
func (f *Forwarder) attemptForwardAndValidate(ctx context.Context) {
	urlIter := threaditer.ThreadSafe(iter.Slice(f.chain.URLs()))

	// setup the channels we use for confirmation
	errChan := make(chan FailedForward)
	resChan := make(chan rawResponse)

	forwardCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	// start requiredConfirmations workers
	for i := uint16(0); i < f.requiredConfirmations; i++ {
		go func() {
			f.mux.RLock()
			defer f.mux.RUnlock()

			for {
				select {
				case <-forwardCtx.Done():
					return
				default:
					done := f.attemptForward(forwardCtx, errChan, resChan, urlIter)
					// if there's nothing else we can end the goroutine
					if done {
						return
					}
				}
			}
		}()
	}

	totalResponses := 0

	for {
		select {
		// request timeout
		case <-f.c.Done():
			return
		case failedForward := <-errChan:
			totalResponses++

			f.failedForwards.Store(failedForward.URL, failedForward.Err)

			// if we've checked every url
			if totalResponses == len(f.chain.URLs()) {
				if done := f.checkResponses(totalResponses); done {
					return
				}
			}
		case res := <-resChan:
			totalResponses++

			// add the response to resmap
			responses, _ := f.resMap.Load(res.hash)
			responses = append(responses, res)
			f.resMap.Store(res.hash, responses)

			// if we've checked every url or the number of non-error responses is greater than or equal to the
			// number of confirmations
			if totalResponses == len(f.chain.URLs()) || uint16(f.resMap.Size()) >= f.requiredConfirmations {
				if done := f.checkResponses(totalResponses); done {
					return
				}
			}
		}
	}
}

// urlConfirmationsHeader is a header specifying which urls were checked.
const urlConfirmationsHeader = "x-checked-urls"

// jsonHashHeader is the hash of the returned json.
const jsonHashHeader = "x-json-hash"

// forwardedFrom the actual url the json was forwarded from.
const forwardedFrom = "x-forwarded-from"

// ErroredRPCResponse contains an errored rpc response
// this is mostly used for debugging.
type ErroredRPCResponse struct {
	Raw json.RawMessage `json:"json_response"`
	URL string          `json:"url"`
}

// ErrorResponse contains error response used for debugging.
type ErrorResponse struct {
	Hashes map[string][]ErroredRPCResponse `json:"hashes"`
	Error  string                          `json:"error"`
	// ErroredURLS returned no response at all
	ErroredURLS []string `json:"errored_urls"`
	// FailedForwards stores lower level json errors where no response could be returned at all
	FailedForwards map[string]string `json:"failed_forwards"`
}

// FailedForward contains a failed forward.
type FailedForward struct {
	// Err is the error returned
	Err error
	// URL is the url of the error
	URL string
}

func (f *Forwarder) checkResponses(responseCount int) (done bool) {
	var valid bool

	f.resMap.Range(func(key string, responses []rawResponse) bool {
		if uint16(len(responses)) >= f.requiredConfirmations {
			responseURLS := make([]string, len(responses))

			for i, url := range responses {
				responseURLS[i] = url.url
			}

			f.c.Header(urlConfirmationsHeader, strings.Join(responseURLS, ","))
			f.c.Header(jsonHashHeader, responses[0].hash)
			f.c.Header(forwardedFrom, responses[0].url)

			f.c.Data(http.StatusOK, gin.MIMEJSON, responses[0].body)
			valid = true

			return false
		}

		return true
	})

	if valid {
		return true
	}

	// every urls been checked, we need to error
	if responseCount == len(f.chain.URLs()) {
		erroredUrls := sets.NewString(f.chain.URLs()...)

		errResponse := ErrorResponse{
			Error:  "could not get consistent response",
			Hashes: make(map[string][]ErroredRPCResponse),
		}

		f.resMap.Range(func(key string, responses []rawResponse) bool {
			for _, response := range responses {
				erroredUrls.Delete(response.url)
				rpcErr := ErroredRPCResponse{
					URL: response.url,
					Raw: response.body,
				}

				errResponse.Hashes[key] = append(errResponse.Hashes[key], rpcErr)
			}
			return true
		})

		errResponse.FailedForwards = make(map[string]string)
		f.failedForwards.Range(func(key string, value error) bool {
			errResponse.FailedForwards[key] = value.Error()
			return true
		})

		errResponse.ErroredURLS = erroredUrls.List()

		f.c.JSON(http.StatusBadGateway, errResponse)

		return true
	}
	return false
}

// attemptForward attempts to forward a request. If it runs out of urls to process
// or context is canceled, done is returned as true
//
// otherwise errors are added to an errChan and responses are added to the response chan.
func (f *Forwarder) attemptForward(ctx context.Context, errChan chan FailedForward, resChan chan rawResponse, urlIter iter.Iterator[string]) (done bool) {
	nextURL := urlIter.Next()
	if nextURL.IsNone() {
		return true
	}

	url := nextURL.Unwrap()

	res, err := f.forwardRequest(ctx, url)
	if err != nil {
		// check if we're done, otherwise add to errchan
		select {
		case <-ctx.Done():
			return true
		case errChan <- FailedForward{Err: err, URL: url}:
			return false
		}
	}

	// request was successful, add the body to the raw response channel for processing
	select {
	case <-ctx.Done():
		return true
	case resChan <- *res:
		return false
	}
}

// fillAndValidate fills request fields and validates fields.
func (f *Forwarder) fillAndValidate(chainID uint32) (ok bool) {
	var err error

	f.chain = f.r.chainManager.GetChain(chainID)
	if f.chain == nil {
		f.c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("chain %d not found", chainID),
		})
		return false
	}

	f.body, err = io.ReadAll(f.c.Request.Body)
	if err != nil {
		f.c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return false
	}

	f.requestID = []byte(f.c.GetHeader(omniHTTP.XRequestIDString))
	f.span.SetAttributes(attribute.String("request_id", string(f.requestID)))

	if ok := f.checkAndSetConfirmability(); !ok {
		return false
	}

	return true
}

// checkAndSetConfirmability checks the confirmability of the request body and makes sure
// we have enough urls to validate the request.
func (f *Forwarder) checkAndSetConfirmability() (ok bool) {
	// if we overrided required confirmations above, use that
	if f.requiredConfirmations == 0 {
		f.requiredConfirmations = f.chain.ConfirmationsThreshold()
	}
	var err error
	f.rpcRequest, err = rpc.ParseRPCPayload(f.body)
	if err != nil {
		f.c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return false
	}

	// If any request ina  batch is not confirmable, the entire batch is marks as non-confirmable
	confirmable, err := areConfirmable(f.rpcRequest)
	if err != nil {
		f.c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return false
	}

	// non-confirmable requests must use 1
	if !confirmable {
		f.requiredConfirmations = 1
	}

	// set the headers
	f.c.Header("x-confirmable", strconv.FormatBool(confirmable))
	// this will be 1 if not confirmable
	f.c.Header("x-required-confirmations", strconv.Itoa(int(f.requiredConfirmations)))

	f.span.SetAttributes(attribute.Int("required_confirmations", int(f.requiredConfirmations)))
	f.span.SetAttributes(attribute.Bool("confirmable", confirmable))
	f.span.SetAttributes(attribute.String("method", f.rpcRequest.Method()))

	// make sure we have enough urls to hit the required confirmation threshold
	if len(f.chain.URLs()) < int(f.requiredConfirmations) {
		f.c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("not enough endpoints for chain %d: found %d needed %d", f.chain.ID(), len(f.chain.URLs()), f.requiredConfirmations),
		})
		return false
	}

	return true
}
