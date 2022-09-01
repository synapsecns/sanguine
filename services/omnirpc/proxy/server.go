package proxy

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/services/omnirpc/config"
	"github.com/synapsecns/sanguine/services/omnirpc/rpcmap"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// RPCProxy proxies rpc request to the fastest endpoint. Requests fallback in cases where data is not available.
type RPCProxy struct {
	// rpcMap contains a list of [chainid]->[]hosts in order of altency
	// this list may not be updated at
	rpcMap *rpcmap.RPCMap
	// config contains the config for each chain
	config config.Config
}

// NewProxy creates a new rpc proxy.
func NewProxy(rpcMap *rpcmap.RPCMap, config config.Config) *RPCProxy {
	return &RPCProxy{
		rpcMap: rpcMap,
		config: config,
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

	logger.Infof("running on port %d", r.config.Port)
	err := router.Run(fmt.Sprintf("0.0.0.0:%d", r.config.Port))
	if err != nil {
		logger.Warn(err)
	}
}

// rpcTimeout is how long to wait for a response.
const rpcTimeout = time.Second * 5

// scanInterval is how long to wait between latency scans.
const scanInterval = time.Second * 60

func (r *RPCProxy) startProxyLoop(ctx context.Context) {
	// TODO(trajan): jitter if not first run
	var waitTime time.Duration

	for {
		select {
		// parent loop terminated
		case <-ctx.Done():
			return
		case <-time.After(waitTime):
			var wg sync.WaitGroup

			for _, chainID := range r.rpcMap.GetChainIDs() {
				wg.Add(1)

				go func(chainID int) {
					r.reorderRPCs(ctx, chainID)

					wg.Done()
				}(chainID)
			}

			wg.Wait()

			waitTime = scanInterval
		}
	}
}
