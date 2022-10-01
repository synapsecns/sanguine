package proxy

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Soft/iter"
	"github.com/gin-gonic/gin"
	"github.com/puzpuzpuz/xsync"
	"github.com/synapsecns/sanguine/core/threaditer"
	"github.com/synapsecns/sanguine/services/omnirpc/chainmanager"
	omniHTTP "github.com/synapsecns/sanguine/services/omnirpc/http"
	"io"
	"k8s.io/apimachinery/pkg/util/sets"
	"net/http"
	"strconv"
	"strings"
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
	requestID string
	// client is the client used for fasthttp
	client omniHTTP.Client
	// resMap is the res map
	// Note: because we use an array here, this is not thread safe for writes
	resMap *xsync.MapOf[[]rawResponse]
	// rpcRequest is the parsed rpc request
	rpcRequest *RPCRequest
}

// Reset resets the forwarder so it can be reused.
func (f *Forwarder) Reset() {
	// client and forwarder can stay the same
	f.c = nil
	f.chain = nil
	f.body = nil
	f.requiredConfirmations = 0
	f.requestID = ""
	f.resMap = nil
	f.rpcRequest = nil
}

// AcquireForwarder allocates a forwarder and allows it to be released when not in use
// this allows forwarder cycling reducing GC overhead.
func (r *RPCProxy) AcquireForwarder() *Forwarder {
	v := r.forwarderPool.Get()
	if v == nil {
		return &Forwarder{
			r:      r,
			client: r.client,
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
func (r *RPCProxy) Forward(c *gin.Context, chainID uint32) {
	forwarder := r.AcquireForwarder()
	defer r.ReleaseForwarder(forwarder)

	forwarder.c = c
	forwarder.resMap = xsync.NewMapOf[[]rawResponse]()

	if ok := forwarder.fillAndValidate(chainID); !ok {
		return
	}

	forwarder.attemptForwardAndValidate()
}

// attemptForwardAndValidate attempts to forward the request and
// makes sure it is valid
// TODO: maybe the context shouldn't be used from a struct here?
// nolint: gocognit, cyclop
func (f *Forwarder) attemptForwardAndValidate() {
	urlIter := threaditer.ThreadSafe(iter.Slice(f.chain.URLs()))

	// setup the channels we use for confirmation
	errChan := make(chan error)
	resChan := make(chan rawResponse)

	forwardCtx, cancel := context.WithCancel(f.c)
	defer cancel()

	// start requiredConfirmations workers
	for i := uint16(0); i < f.requiredConfirmations; i++ {
		go func() {
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
		case <-errChan:
			totalResponses++
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
// thisis mostly used for debugging.
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
func (f *Forwarder) attemptForward(ctx context.Context, errChan chan error, resChan chan rawResponse, urlIter iter.Iterator[string]) (done bool) {
	nextURL := urlIter.Next()
	if nextURL.IsNone() {
		return true
	}

	url := nextURL.Unwrap()

	res, err := f.forwardRequest(ctx, url, f.requestID)
	if err != nil {
		// check if we're done, otherwise add to errchan
		select {
		case <-ctx.Done():
			return true
		case errChan <- err:
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

	f.requestID = f.c.GetHeader(omniHTTP.XRequestIDString)

	if ok := f.checkAndSetConfirmability(); !ok {
		return false
	}

	return true
}

// checkAndSetConfirmability checks the confirmability of the request body and makes sure
// we have enough urls to validate the request.
func (f *Forwarder) checkAndSetConfirmability() (ok bool) {
	f.requiredConfirmations = f.chain.ConfirmationsThreshold()

	var err error
	f.rpcRequest, err = parseRPCPayload(f.body)
	if err != nil {
		f.c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return false
	}

	confirmable, err := f.rpcRequest.isConfirmable()
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

	// make sure we have enough urls to hit the required confirmation threshold
	if len(f.chain.URLs()) < int(f.requiredConfirmations) {
		f.c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("not enough endpoints for chain %d: found %d needed %d", f.chain.ID(), len(f.chain.URLs()), f.requiredConfirmations),
		})
		return false
	}

	return true
}
