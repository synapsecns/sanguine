package proxy

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/omnirpc/chainmanager"
	"github.com/synapsecns/sanguine/services/omnirpc/collection"
	"github.com/synapsecns/sanguine/services/omnirpc/config"
	omniHTTP "github.com/synapsecns/sanguine/services/omnirpc/http"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// RPCProxy proxies rpc request to the fastest endpoint. Requests fallback in cases where data is not available.
type RPCProxy struct {
	// tracer is the tracer for the proxy
	tracer trace.Tracer
	// chainManager contains a list of chains and latency ordered rpcs
	chainManager chainmanager.ChainManager
	// config contains the config for each chain
	refreshInterval time.Duration
	// port is the por the server is run on
	port uint16
	// forwarderPool is used for allocating forwarders
	forwarderPool sync.Pool
	// client contains the http client
	client omniHTTP.Client
	// handler is the metrics handler
	handler metrics.Handler
}

// defaultInterval is the default refresh interval.
const defaultInterval = 30

// NewProxy creates a new rpc proxy.
func NewProxy(config config.Config, handler metrics.Handler) *RPCProxy {
	if config.RefreshInterval == 0 {
		logger.Warn("no refresh interval set (or interval is 0), using default of %d seconds", defaultInterval)
	}

	return &RPCProxy{
		chainManager:    chainmanager.NewChainManagerFromConfig(config, handler),
		refreshInterval: time.Second * time.Duration(config.RefreshInterval),
		port:            config.Port,
		client:          omniHTTP.NewClient(omniHTTP.ClientTypeFromString(config.ClientType)),
		handler:         handler,
		tracer:          handler.Tracer(),
	}
}

// Run runs the rpc server until context cancellation.
func (r *RPCProxy) Run(ctx context.Context) {
	go r.startProxyLoop(ctx)

	router := ginhelper.New(logger)
	router.Use(r.handler.Gin())

	router.POST("/rpc/:id", func(c *gin.Context) {
		chainID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("chainid must be a number: %d", chainID),
			})
		}
		r.Forward(c, uint32(chainID), nil)
	})

	router.POST("/confirmations/:confirmations/rpc/:id", func(c *gin.Context) {
		chainID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("chainid must be a number: %d", chainID),
			})
		}
		realConfs, err := strconv.Atoi(c.Param("confirmations"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("confirmations must be a number: %d", chainID),
			})
		}

		confirmations := uint16(realConfs)

		r.Forward(c, uint32(chainID), &confirmations)
	})

	// gets a list of chain-ids
	// TODO: this needs to be added to the collection.json
	router.GET("/chain-ids", func(c *gin.Context) {
		c.JSON(http.StatusOK, r.chainManager.GetChainIDs())
	})

	router.GET("/collection.json", func(c *gin.Context) {
		res, err := collection.CreateCollection()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("could not parse collection: %v", err),
			})
		}
		c.Data(http.StatusOK, gin.MIMEJSON, res)
	})

	logger.Infof("running on port %d", r.port)
	err := router.Run(fmt.Sprintf("0.0.0.0:%d", r.port))
	if err != nil {
		logger.Warn(err)
	}
}

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
			r.benchmarkProxies(ctx)

			waitTime = scanInterval
		}
	}
}

// benchmarkProxies benchmarks all proxies.
func (r *RPCProxy) benchmarkProxies(parentCtx context.Context) {
	ctx, span := r.handler.Tracer().Start(parentCtx, "benchmarkProxies")
	defer func() {
		span.End()
	}()

	var wg sync.WaitGroup

	for _, chainID := range r.chainManager.GetChainIDs() {
		wg.Add(1)

		go func(ctx context.Context, chainID uint32) {
			r.chainManager.RefreshRPCInfo(ctx, chainID)

			wg.Done()
		}(ctx, chainID)
	}

	wg.Wait()
}

// Port gets the port the proxy is running on.
func (r *RPCProxy) Port() uint16 {
	return r.port
}
