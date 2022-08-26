package proxy

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/serivces/omnirpc/latency"
	"github.com/synapsecns/sanguine/serivces/omnirpc/rpcmap"
	"golang.org/x/sync/errgroup"
	"net/http"
	"sort"
	"strconv"
	"time"
)

// RPCProxy proxies rpc request to the fastest endpoint. Requests fallback in cases where data is not available.
type RPCProxy struct {
	// port is the port the rpc proxy is running on
	port uint32
	// rpcMap contains a list of [chainid]->[]hosts in order of altency
	// this list may not be updated at
	rpcMap *rpcmap.RPCMap
}

// NewProxy creates a new rpc proxy.
func NewProxy(port uint32, rpcMap *rpcmap.RPCMap) *RPCProxy {
	return &RPCProxy{
		port:   port,
		rpcMap: rpcMap,
	}
}

// Run runs the rpc server until context cancellation.
func (r *RPCProxy) Run(ctx context.Context) {
	go r.startProxyLoop(ctx)

	router := gin.Default()

	router.POST("/rpc/:id", func(c *gin.Context) {
		chainID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("chainid must be a number: %d", chainID),
			})
		}
		r.serveRPCReq(c, chainID)
	})

	logger.Infof("running on port %d", r.port)
	err := router.Run(fmt.Sprintf("0.0.0.0:%d", r.port))
	if err != nil {
		logger.Warn(err)
	}
}

// rpcTimeout is how long to wait for a response.
const rpcTimeout = time.Second * 5

// scanInterval is how long to wait between latency scans.
const scanInterval = time.Second * 60

func (r *RPCProxy) startProxyLoop(parentCtx context.Context) {
	// TODO(trajan): jitter if not first run
	var waitTime time.Duration

	for {
		select {
		// parent loop terminated
		case <-parentCtx.Done():
			return
		case <-time.After(waitTime):
			g, ctx := errgroup.WithContext(parentCtx)
			for _, chainID := range r.rpcMap.GetChainIDs() {
				// capture func literal
				chainID := chainID
				rpcList := r.rpcMap.ChainID(chainID)

				// no need to setup a goroutine
				if len(rpcList) == 0 {
					continue
				}

				g.Go(func() error {
					latencyList := latency.GetRPCLatency(ctx, rpcTimeout, rpcList)

					// sort loweset->highest latency
					sort.Slice(latencyList, func(i, j int) bool {
						// ignore latencies with an error
						if latencyList[i].HasError {
							return false
						}
						return latencyList[i].Latency < latencyList[j].Latency
					})

					var newOrder []string
					for _, rpcItem := range latencyList {
						newOrder = append(newOrder, rpcItem.URL)
					}

					r.rpcMap.PutChainID(chainID, newOrder)

					return nil
				})
			}

			_ = g.Wait()

			waitTime = scanInterval
		}
	}
}
