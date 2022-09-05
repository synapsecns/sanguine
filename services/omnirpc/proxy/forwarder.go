package proxy

import (
	"context"
	"fmt"
	"github.com/Soft/iter"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/threaditer"
	"github.com/synapsecns/sanguine/services/omnirpc/chainmanager"
	"io"
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
}

// Forward forwards the rpc request to the servers and makes assertions around confirmation thresholds.
func (r *RPCProxy) Forward(c *gin.Context, chainID uint32) {
	forwarder := &Forwarder{
		r: r,
		c: c,
	}

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
	var prevResponses []rawResponse

	for {
		select {
		// request timeout
		case <-f.c.Done():
			return
		case <-errChan:
			totalResponses++
			// if we've checked every url
			if totalResponses == len(f.chain.URLs()) {
				if done := f.checkResponses(totalResponses, prevResponses); done {
					return
				}
			}
		case res := <-resChan:
			prevResponses = append(prevResponses, res)
			// if we've checked every url or the number of non-error responses is greater than or equal to the
			// number of confirmations
			if totalResponses == len(f.chain.URLs()) || uint16(len(prevResponses)) >= f.requiredConfirmations {
				if done := f.checkResponses(totalResponses, prevResponses); done {
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

// forwardedFrom the actual url the json was forwared from.
const forwardedFrom = "x-forwarded-from"

func (f *Forwarder) checkResponses(responseCount int, prevResponses []rawResponse) (done bool) {
	// hash-> rawResponse
	resMap := make(map[[32]byte][]rawResponse)

	for _, res := range prevResponses {
		resMap[res.hash] = append(resMap[res.hash], res)
	}

	// check for a valid response
	for key, responses := range resMap {
		if uint16(len(responses)) >= f.requiredConfirmations {
			var responseUrls []string
			for _, url := range responses {
				responseUrls = append(responseUrls, url.url)
			}

			// use the first response, they're both the same
			f.c.Header(urlConfirmationsHeader, strings.Join(responseUrls, ","))
			f.c.Header(jsonHashHeader, common.Bytes2Hex(core.BytesToSlice(key)))
			f.c.Header(forwardedFrom, responses[0].url)

			f.c.Data(http.StatusOK, gin.MIMEJSON, responses[0].body)
			return true
		}
	}

	// every urls been checked, we need to error
	if responseCount == len(f.chain.URLs()) {
		f.c.JSON(http.StatusBadGateway, gin.H{
			"error": "could not get consistent response",
		})

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

	res, err := forwardRequest(ctx, f.body, url, f.requestID)
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

	f.requestID = f.c.GetHeader(requestIDKey)

	if ok := f.checkAndSetConfirmability(); !ok {
		return false
	}

	return true
}

// checkAndSetConfirmability checks the confirmability of the request body and makes sure
// we have enough urls to validate the request.
func (f *Forwarder) checkAndSetConfirmability() (ok bool) {
	f.requiredConfirmations = f.chain.ConfirmationsThreshold()

	confirmable, err := isConfirmable(f.body)
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
